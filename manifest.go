package svcgen

import (
	"encoding/xml"
	"os"
	"strings"
)

const OCIProcesSVCName = "application/oci-process"
const OCIEntryPointSVCName = "application/oci-entrypoint"
const Header = `<?xml version="1.0"?>
<!DOCTYPE service_bundle SYSTEM "/usr/share/lib/xml/dtd/service_bundle.dtd.1">
`

func WriteManifest(filePath string, bundle ServiceBundle) error {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.WriteString(Header); err != nil {
		return err
	}

	enc := xml.NewEncoder(f)
	enc.Indent("", "\t")
	if err = enc.Encode(bundle); err != nil {
		return err
	}

	return nil
}

func NewManifest(name string) ServiceBundle {
	return ServiceBundle{
		Type:     "manifest",
		Name:     name,
		Services: []Service{},
	}
}

func NewProfile(name string) ServiceBundle {
	return ServiceBundle{
		Type:     "profile",
		Name:     name,
		Services: []Service{},
	}
}

type Params struct {
	WorkingDirectory string
	UserName         string
	GroupName        string
	Groups           GroupList
	Env              Environment
	StartCommand     Command
	StopCommand      Command
	Type             ServiceType
}

type ServiceType int

const (
	ServiceTypeChild ServiceType = iota
	ServiceTypeTransient
	ServiceTypeContract
)

type Command []string

func (cmd Command) String() string {
	if cmd == nil {
		return ":kill"
	}

	return strings.Join(cmd, " ")
}

type Environment []string

func (env Environment) Convert() []EnvVar {
	if env == nil {
		return nil
	}

	v := make([]EnvVar, 0)
	for _, sE := range env {
		idx := strings.Index(sE, "=")
		if idx == -1 {
			continue
		}

		ev := EnvVar{
			Name:  sE[:idx],
			Value: sE[idx+1:],
		}

		v = append(v, ev)
	}

	return v
}

type GroupList []string

func (list GroupList) String() string {
	if list == nil {
		return ""
	}

	returnValue := ""
	for index, entry := range list {
		if index == 0 {
			returnValue = entry
		} else {
			returnValue += "," + entry
		}
	}
	return returnValue
}

func NewManifestWithParams(bundleName, svcName string, milestone string, params Params) ServiceBundle {
	if milestone == "" {
		milestone = "svc:/milestone/multi-user"
	}

	bundle := ServiceBundle{
		Type: "manifest",
		Name: bundleName,
		Services: []Service{
			{
				Dependency: []Dependency{
					{
						Name:      "milestone",
						Grouping:  "require_all",
						RestartOn: "none",
						Type:      "service",
						ServiceFmri: []ServiceFmri{
							{Value: milestone},
						},
					},
				},
				Name:    svcName,
				Version: "1",
				Type:    "service",
				CreateDefaultInstance: &CreateDefaultInstance{
					Enabled: true,
				},
				MethodContext: &MethodContext{
					WorkingDirectory: params.WorkingDirectory,
					MethodCredential: &MethodCredential{
						User:       params.UserName,
						Group:      params.GroupName,
						SuppGroups: params.Groups.String(),
					},
					MethodEnvironment: &MethodEnvironment{
						Envvar: params.Env.Convert(),
					},
				},
				ExecMethod: []ExecMethod{
					{
						Type:           "method",
						Name:           "start",
						Exec:           params.StartCommand.String(),
						TimeoutSeconds: "0",
					},
					{
						Type:           "method",
						Name:           "stop",
						Exec:           params.StopCommand.String(),
						TimeoutSeconds: "0",
					},
				},
			},
		},
	}

	for _, svc := range bundle.Services {
		svc.PropertyGroup = append(svc.PropertyGroup, getServiceTypePropertyGroup(params.Type))
	}

	return bundle
}

func getServiceTypePropertyGroup(t ServiceType) PropertyGroup {
	switch t {
	case ServiceTypeChild:
		return PropertyGroup{
			Name: "startd",
			Type: "framework",
			PropVal: []PropVal{
				{
					Name:  "duration",
					Type:  "astring",
					Value: "child",
				},
			},
		}
	case ServiceTypeTransient:
		return PropertyGroup{Name: "startd",
			Type: "framework",
			PropVal: []PropVal{
				{
					Name:  "duration",
					Type:  "astring",
					Value: "transient",
				},
			},
		}
	case ServiceTypeContract:
		fallthrough
	default:
		return PropertyGroup{Name: "startd",
			Type: "framework",
			PropVal: []PropVal{
				{
					Name:  "duration",
					Type:  "astring",
					Value: "contract",
				},
			},
		}
	}
}
