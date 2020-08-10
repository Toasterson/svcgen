// Copyright (c) 2004, 2010, Oracle and/or its affiliates. All rights reserved.
//
// CDDL HEADER START
//
// The contents of this file are subject to the terms of the
// Common Development and Distribution License (the "License").
// You may not use this file except in compliance with the License.
//
// You can obtain a copy of the license at usr/src/OPENSOLARIS.LICENSE
// or http://www.opensolaris.org/os/licensing.
// See the License for the specific language governing permissions
// and limitations under the License.
//
// When distributing Covered Code, include this CDDL HEADER in each
// file and include the License file at usr/src/OPENSOLARIS.LICENSE.
// If applicable, add the following below this CDDL HEADER, with the
// fields enclosed by brackets "[]" replaced with your own identifying
// information: Portions Copyright [yyyy] [name of copyright owner]
//
// CDDL HEADER END
//
//
//
//  Service description DTD
//
//    Most attributes are string values (or an individual string from a
//    restricted set), but attributes with a specific type requirement are
//    noted in the comment describing the element.
//
//
//

// NOTE: the golang package has no support for that yet...

package svcgen

import (
	"encoding/xml"
)

// stability
//
//    This element associates an SMI stability level with the parent
//    element.  See attributes(5) for an explanation of interface
//    stability levels.
//
//    Its attribute is
//
//	value	The stability level of the parent element.
//          possible values are ( Standard | Stable | Evolving | Unstable |
//			External | Obsolete )
type Stability struct {
	XMLName xml.Name `xml:"stability"`
	Value   string   `xml:"value,attr"`
}

//
//     These entities are used for the property, propval and property_group
//     elements, that require type attributes for manifest, while for profiles
//     the type attributes are only implied.

// value_node
//
//    This element represents a single value within any of the typed
//    property value lists.
//
//    Its attribute is
//
//	value	The value for this node in the list.
type Valuenode struct {
	XMLName xml.Name `xml:"value_node"`
	Value   string   `xml:"value,attr"`
}

// count_list
//  integer_list
//  opaque_list
//  host_list
//  hostname_list
//  net_address_list
//  net_address_v4_list
//  net_address_v6_list
//  time_list
//  astring_list
//  ustring_list
//  boolean_list
//  fmri_list
//  uri_list
//
//    These elements represent the typed lists of values for a property.
//    Each contains one or more value_node elements representing each
//    value on the list.
//
//    None of these elements has attributes.
type Countlist struct {
	XMLName   xml.Name    `xml:"count_list"`
	Valuenode []Valuenode `xml:"value_node"`
}

type Integerlist struct {
	XMLName   xml.Name    `xml:"integer_list"`
	Valuenode []Valuenode `xml:"value_node"`
}

type Opaquelist struct {
	XMLName   xml.Name    `xml:"opaque_list"`
	Valuenode []Valuenode `xml:"value_node"`
}

type Hostlist struct {
	XMLName   xml.Name    `xml:"host_list"`
	Valuenode []Valuenode `xml:"value_node"`
}

type Hostnamelist struct {
	XMLName   xml.Name    `xml:"hostname_list"`
	Valuenode []Valuenode `xml:"value_node"`
}

type Netaddresslist struct {
	XMLName   xml.Name    `xml:"net_address_list"`
	Valuenode []Valuenode `xml:"value_node"`
}

type Netaddressv4list struct {
	XMLName   xml.Name    `xml:"net_address_v4_list"`
	Valuenode []Valuenode `xml:"value_node"`
}

type Netaddressv6list struct {
	XMLName   xml.Name    `xml:"net_address_v6_list"`
	Valuenode []Valuenode `xml:"value_node"`
}

type Timelist struct {
	XMLName   xml.Name    `xml:"time_list"`
	Valuenode []Valuenode `xml:"value_node"`
}

type Astringlist struct {
	XMLName   xml.Name    `xml:"astring_list"`
	Valuenode []Valuenode `xml:"value_node"`
}

type Ustringlist struct {
	XMLName   xml.Name    `xml:"ustring_list"`
	Valuenode []Valuenode `xml:"value_node"`
}

type Booleanlist struct {
	XMLName   xml.Name    `xml:"boolean_list"`
	Valuenode []Valuenode `xml:"value_node"`
}

type Fmrilist struct {
	XMLName   xml.Name    `xml:"fmri_list"`
	Valuenode []Valuenode `xml:"value_node"`
}

type Urilist struct {
	XMLName   xml.Name    `xml:"uri_list"`
	Valuenode []Valuenode `xml:"value_node"`
}

// property
//
//     This element is for a singly or multiply valued property within a
//     property group.  It contains an appropriate value list element,
//     which is expected to be consistent with the type attribute.
//
//     Its attributes are
//
//	name	The name of this property.
//
//	type	The data type for this property.
//          Possible types are count | integer | opaque | host | hostname |
//			net_address | net_address_v4 | net_address_v6 | time |
//			astring | ustring | boolean | fmri | uri
//
//	override These values should replace values already in the
//		repository.
type Property struct {
	XMLName  xml.Name `xml:"property"`
	Name     string   `xml:"name,attr"`
	Type     string   `xml:"type,attr"`
	Override bool     `xml:"override,attr,omitempty"`
}

// propval
//
//     This element is for a singly valued property within a property
//     group.  List-valued properties must use the property element above.
//
//     Its attributes are
//
//	name	The name of this property.
//
//	type	The data type for this property.
//
//	value	The value for this property.  Must match type
//		restriction of type attribute.
//
//	override This value should replace any values already in the
//		repository.
type PropVal struct {
	XMLName  xml.Name `xml:"propval"`
	Name     string   `xml:"name,attr"`
	Type     string   `xml:"type,attr"`
	Value    string   `xml:"value,attr"`
	Override bool     `xml:"override,attr,omitempty"`
}

// property_group
//
//    This element is for a set of related properties on a service or
//    instance.  It contains an optional stability element, as well as
//    zero or more property-containing elements.
//
//    Its attributes are
//
//	name	The name of this property group.
//
//	type	A category for this property group.  Groups of type
//		"framework", "implementation" or "template" are primarily
//		of interest to the service management facility, while
//		groups of type "application" are expected to be only of
//		interest to the service to which this group is attached.
//		Other types may be introduced using the service symbol
//		namespace conventions.
//
//	delete	If in the repository, this property group should be removed.
type PropertyGroup struct {
	XMLName   xml.Name   `xml:"property_group"`
	Name      string     `xml:"name,attr"`
	Type      string     `xml:"type,attr"`
	Delete    bool       `xml:"delete,attr,omitempty"`
	Stability *Stability `xml:"stability"`
	PropVal   []PropVal  `xml:"propval"`
	Property  []Property `xml:"property"`
}

// service_fmri
//
//    This element defines a reference to a service FMRI (for either a
//    service or an instance).
//
//    Its attribute is
//
//	value	The FMRI.
type ServiceFmri struct {
	XMLName xml.Name `xml:"service_fmri"`
	Value   string   `xml:"value,attr"`
}

// dependency
//
//    This element identifies a group of FMRIs upon which the service is
//    in some sense dependent.  Its interpretation is left to the
//    restarter to which a particular service instance is delegated.  It
//    contains a group of service FMRIs, as well as a block of properties.
//
//    Its attributes are
//
//	name	The name of this dependency.
//
//	grouping The relationship between the various FMRIs grouped
//		here; "require_all" of the FMRIs to be online, "require_any"
//		of the FMRIs to be online, or "exclude_all" of the FMRIs
//		from being online or in maintenance for the dependency to
//		be satisfied.  "optional_all" dependencies are satisfied
//		when all of the FMRIs are either online or unable to come
//		online (because they are disabled, misconfigured, or one
//		of their dependencies is unable to come online).
//
//	restart_on The type of events from the FMRIs that the service should
//		be restarted for.  "error" restarts the service if the
//		dependency is restarted due to hardware fault.  "restart"
//		restarts the service if the dependency is restarted for
//		any reason, including hardware fault.  "refresh" restarts
//		the service if the dependency is refreshed or restarted for
//		any reason.  "none" will never restart the service due to
//		dependency state changes.
//
//	type	The type of dependency: on another service ('service'), on
//		a filesystem path ('path'), or another dependency type.
//
//	delete	This dependency should be deleted.
type Dependency struct {
	XMLName     xml.Name      `xml:"dependency"`
	Name        string        `xml:"name,attr"`
	Grouping    string        `xml:"grouping,attr"`
	RestartOn   string        `xml:"restart_on,attr"`
	Type        string        `xml:"type,attr"`
	Delete      string        `xml:"delete,attr,omitempty"`
	ServiceFmri []ServiceFmri `xml:"service_fmri"`
	Stability   *Stability    `xml:"stability"`
	PropVal     []PropVal     `xml:"propval"`
	Property    []Property    `xml:"property"`
}

// dependent
//
//    This element identifies a service which should depend on this service.  It
//    corresponds to a dependency in the named service.  The grouping and type
//    attributes of that dependency are implied to be "require_all" and
//    "service", respectively.
//
//    Its attributes are
//
//	name	The name of the dependency property group to create in the
//		dependent entity.
//
//	grouping The grouping relationship of the dependency property
//		group to create in the dependent entity.  See "grouping"
//		attribute on the dependency element.
//
//	restart_on The type of events from this service that the named service
//		should be restarted for.
//
//	delete	True if this dependent should be deleted.
//
//	override Whether to replace an existing dependent of the same name.
type Dependent struct {
	XMLName     xml.Name    `xml:"dependent"`
	Name        string      `xml:"name,attr"`
	Grouping    string      `xml:"grouping,attr"`
	RestartOn   string      `xml:"restart_on,attr"`
	Delete      bool        `xml:"delete,attr,omitempty"`
	Override    bool        `xml:"override,attr,omitempty"`
	ServiceFmri ServiceFmri `xml:"service_fmri"`
	Stability   *Stability  `xml:"stability"`
	PropVal     []PropVal   `xml:"propval"`
	Property    []Property  `xml:"property"`
}

// envvar
//
//    An environment variable. It has two attributes:
//
//	name	The name of the environment variable.
//	value	The value of the environment variable.
type EnvVar struct {
	XMLName xml.Name `xml:"envvar"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

// method_environment
//
//    This element defines the environment for a method. It has no
//    attributes, and one or more envvar child elements.
type MethodEnvironment struct {
	XMLName xml.Name `xml:"method_environment"`
	Envvar  []EnvVar `xml:"envvar"`
}

//  method_profile
//
//    This element indicates which exec_attr(5) profile applies to the
//    method context being defined.
//
//    Its attribute is
//
//	name	The name of the profile.
type MethodProfile struct {
	XMLName xml.Name `xml:"method_profile"`
	Name    string   `xml:"name,attr"`
}

// method_credential
//
//    This element specifies credential attributes for the execution
//    method to use.
//
//    Its attributes are
//
//	user	The user ID, in numeric or text form.
//
//	group	The group ID, in numeric or text form.  If absent or
//		":default", the group associated with the user in the
//		passwd database.
//
//	supp_groups Supplementary group IDs to be associated with the
//		method, separated by commas or spaces.  If absent or
//		":default", initgroups(3C) will be used.
//
//	privileges An optional string specifying the privilege set.
//
//	limit_privileges An optional string specifying the limit
//		privilege set.
type MethodCredential struct {
	XMLName         xml.Name `xml:"method_credential"`
	User            string   `xml:"user,attr"`
	Group           string   `xml:"group,attr,omitempty"`
	SuppGroups      string   `xml:"supp_groups,attr,omitempty"`
	Privileges      string   `xml:"privileges,attr,omitempty"`
	LimitPrivileges string   `xml:"limit_privileges,attr,omitempty"`
}

// method_context
//
//    This element combines credential and resource management attributes
//    for execution methods.  It may contain a method_environment, or
//    a method_profile or method_credential element.
//
//    Its attributes are
//
//	working_directory The home directory to launch the method from.
//		":default" can be used as a token to indicate use of the
//		user specified by the credential or profile specified.
//
//	project	The project ID, in numeric or text form.  ":default" can
//		be used as a token to indicate use of the project
//		identified by getdefaultproj(3PROJECT) for the non-root
//		user specified by the credential or profile specified.
//		If the user is root, ":default" designates the project
//		the restarter is running in.
//
//	resource_pool The resource pool name to launch the method on.
//		":default" can be used as a token to indicate use of the
//		pool specified in the project(4) entry given in the
//		"project" attribute above.
type MethodContext struct {
	XMLName           xml.Name           `xml:"method_context"`
	SecurityFlags     string             `xml:"security_flags,attr,omitempty"`
	WorkingDirectory  string             `xml:"working_directory,attr,omitempty"`
	Project           string             `xml:"project,attr,omitempty"`
	ResourcePool      string             `xml:"resource_pool,attr,omitempty"`
	MethodProfile     *MethodProfile     `xml:"method_profile"`
	MethodCredential  *MethodCredential  `xml:"method_credential"`
	MethodEnvironment *MethodEnvironment `xml:"method_environment"`
}

// exec_method
//
//    This element describes one of the methods used by the designated
//    restarter to act on the service instance.  Its interpretation is
//    left to the restarter to which a particular service instance is
//    delegated.  It contains a set of attributes, an optional method
//    context, and an optional stability element for the optional
//    properties that can be included.
//
//    Its attributes are
//
//	type	The type of method, either "method" or "monitor".
//
//	name	Name of this execution method.  The method names are
//		usually a defined interface of the restarter to which an
//		instance of this service is delegated.
//
//	exec	The string identifying the action to take.  For
//		svc.startd(1M), this is a string suitable to pass to
//		exec(2).
//
//	timeout_seconds [integer] Duration, in seconds, to wait for this
//		method to complete.  A '0' or '-1' denotes an infinite
//		timeout.
//
//	delete	If in the repository, the property group for this method
//		should be removed.
type ExecMethod struct {
	XMLName        xml.Name       `xml:"exec_method"`
	Type           string         `xml:"type,attr"`
	Name           string         `xml:"name,attr"`
	Exec           string         `xml:"exec,attr"`
	TimeoutSeconds string         `xml:"timeout_seconds,attr"`
	Delete         bool           `xml:"delete,attr,omitempty"`
	MethodContext  *MethodContext `xml:"method_context"`
	Stability      *Stability     `xml:"stability"`
	PropVal        []PropVal      `xml:"propval"`
	Property       []Property     `xml:"property"`
}

// restarter
//
//    A flag element identifying the restarter to which this service or
//    service instance is delegated.  Contains the FMRI naming the
//    delegated restarter.
//
//    This element has no attributes.
type Restarter struct {
	XMLName     xml.Name    `xml:"restarter"`
	ServiceFmri ServiceFmri `xml:"service_fmri"`
}

// doc_link
//
//    The doc_link relates a resource described by the given URI to the
//    service described by the containing template.  The resource is
//    expected to be a documentation or elucidatory reference of some
//    kind.
//
//    Its attributes are
//
//      name      A label for this resource.
//
//      uri       A URI to the resource.
type DocLink struct {
	XMLName xml.Name `xml:"doc_link"`
	Name    string   `xml:"name,attr"`
	Uri     string   `xml:"uri,attr"`
}

// manpage
//
//    The manpage element connects the reference manual page to the
//    template's service.
//
//    Its attributes are
//
//      title     The manual page title.
//
//      section   The manual page's section.
//
//      manpath   The MANPATH environment variable, as described in man(1)
//                that is required to reach the named manual page
type ManPage struct {
	XMLName xml.Name `xml:"manpage"`
	Title   string   `xml:"title,attr"`
	Section string   `xml:"section,attr"`
	ManPath string   `xml:"manpath,attr,omitempty"`
}

// documentation
//
//    The documentation element groups an arbitrary number of doc_link
//    and manpage references.
//
//    It has no attributes.
type Documentation struct {
	XMLName xml.Name  `xml:"documentation"`
	Doclink []DocLink `xml:"doc_link"`
	Manpage []ManPage `xml:"manpage"`
}

// loctext
//
//    The loctext element is a container for localized text.
//
//    Its sole attribute is
//
//	xml:lang The name of the locale, in the form accepted by LC_ALL,
//		etc.  See locale(5).
type LocText struct {
	XMLName xml.Name `xml:"loctext"`
	XmlLang string   `xml:"xml:lang,attr"`
	Text    string   `xml:",chardata"`
}

// description
//
//    The description holds a set of potentially longer, localized strings that
//    consist of a short description of the service.
//
//    The description has no attributes.
type Description struct {
	XMLName  xml.Name  `xml:"description"`
	LocTexts []LocText `xml:"loctext"`
}

// common_name
//
//    The common_name holds a set of short, localized strings that
//    represent a well-known name for the service in the given locale.
//
//    The common_name has no attributes.
type CommonName struct {
	XMLName  xml.Name  `xml:"common_name"`
	LocTexts []LocText `xml:"loctext"`
}

// units
//
//    The units a numerical property is expressed in.
type Units struct {
	XMLName  xml.Name  `xml:"units"`
	LocTexts []LocText `xml:"loctext"`
}

// visibility
//
//    Expresses how a property is typically accessed.  This isn't
//    intended as access control, but as an indicator as to how a
//    property is used.
//
//    Its attributes are:
//
//      value     'hidden', 'readonly', or 'readwrite' indicating that
//		the property should be hidden from the user, shown but
//		read-only, or modifiable.
type Visibility struct {
	XMLName xml.Name `xml:"visibility"`
	Value   string   `xml:"value,attr"`
}

// value
//
//    Describes a legal value for a property value, and optionally contains a
//    human-readable name and description for the specified property
//    value.
//
//    Its attributes are:
//
//      name	A string representation of the value.
type Value struct {
	XMLName     xml.Name     `xml:"value"`
	Name        string       `xml:"name,attr"`
	CommonName  *CommonName  `xml:"common_name"`
	Description *Description `xml:"description"`
}

// values
//
//    Human-readable names and descriptions for valid values of a property.
type Values struct {
	XMLName xml.Name `xml:"values"`
	Value   []Value  `xml:"value"`
}

// cardinality
//
//    Places a constraint on the number of values the property can take
//    on.
//
//    Its attributes are:
//	min	minimum number of values 0
//	max	maximum number of values 18446744073709551615
//
//    Both attributes are optional.  If min is not specified, it defaults to
//    0.  If max is not specified it indicates an unlimited number of values.
//    If neither is specified this indicates 0 or more values.
type Cardinality struct {
	XMLName xml.Name `xml:"cardinality"`
	Min     string   `xml:"min,attr,omitempty"`
	Max     string   `xml:"max,attr,omitempty"`
}

// internal_separators
//
//    Indicates the separators used within a property's value used to
//    separate the actual values.  Used in situations where multiple
//    values are packed into a single property value instead of using a
//    multi-valued property.
type InternalSeparators struct {
	Body string `xml:",chardata"`
}

// range
//
//    Indicates a range of possible integer values.
//
//    Its attributes are:
//
//      min	The minimum value of the range (inclusive).
//      max	The maximum value of the range (inclusive).
type Range struct {
	XMLName xml.Name `xml:"range"`
	MinAttr string   `xml:"min,attr"`
	MaxAttr string   `xml:"max,attr"`
}

// constraints
//
//    Provides a set of constraints on the values a property can take on.
type Constraints struct {
	XMLName xml.Name `xml:"constraints"`
	Value   []Value  `xml:"value"`
	Range   []Range  `xml:"range"`
}

//include_values
//
//    Includes an entire set of values in the choices block.
//
//    Its attributes are:
//
//	type    Either "constraints" or "values", indicating an
//		inclusion of all values allowed by the property's
//		constraints or all values for which there are
//		human-readable names and descriptions, respectively.
type IncludeValues struct {
	XMLName xml.Name `xml:"include_values"`
	Type    string   `xml:"type,attr"`
}

// choices
//
//    Provides a set of common choices for the values a property can take
//    on.  Useful in those cases where the possibilities are unenumerable
//    or merely inconveniently legion, and a manageable subset is desired
//    for presentation in a user interface.
type Choices struct {
	XMLName       xml.Name        `xml:"choices"`
	Value         []Value         `xml:"value"`
	Range         []Range         `xml:"range"`
	IncludeValues []IncludeValues `xml:"include_values"`
}

// prop_pattern
//
//
//    The prop_pattern describes one property of the enclosing property group
//    pattern.
//
//    Its attributes are:
//
//	name    The property's name.
//	type    The property's type.
//	required
//		If the property group is present, this property is required.
//
//	type can be omitted if required is false.
//  possible values for type are ( count | integer | opaque | host | hostname |
//			net_address | net_address_v4 | net_address_v6 | time |
//			astring | ustring | boolean | fmri | uri )
type PropPattern struct {
	XMLName            xml.Name            `xml:"prop_pattern"`
	Name               string              `xml:"name,attr"`
	Type               string              `xml:"type,attr,omitempty"`
	Required           string              `xml:"required,attr,omitempty"`
	CommonName         *CommonName         `xml:"common_name"`
	Description        *Description        `xml:"description"`
	Units              *Units              `xml:"units"`
	Visibility         *Visibility         `xml:"visibility"`
	Cardinality        *Cardinality        `xml:"cardinality"`
	InternalSeparators *InternalSeparators `xml:"internal_separators"`
	Values             *Values             `xml:"values"`
	Constraints        *Constraints        `xml:"constraints"`
	Choices            *Choices            `xml:"choices"`
}

// pg_pattern
//
//    The pg_pattern describes one property group.
//    Depending on the element's attributes, these descriptions may apply
//    to just the enclosing service/instance, instances of the enclosing
//    service, delegates of the service (assuming it is a restarter), or
//    all services.
//
//    Its attributes are:
//
//	name    The property group's name.  If not specified, it
//		matches all property groups with the specified type.
//	type    The property group's type.  If not specified, it
//		matches all property groups with the specified name.
//	required
//		If the property group is required.
//	target	The scope of the pattern, which may be all, delegate,
//		instance, or this.  'all' is reserved for framework use
//		and applies the template to all services on the system.
//		'delegate' is reserved for restarters, and means the
//		template applies to all services which use the restarter.
//		'this' would refer to the defining service or instance.
//		'instance' can only be used in a service's template block,
//		and means the definition applies to all instances of this
//		service.
type PgPattern struct {
	XMLName     xml.Name      `xml:"pg_pattern"`
	Name        string        `xml:"name,attr,omitempty"`
	Type        string        `xml:"type,attr,omitempty"`
	Required    string        `xml:"required,attr,omitempty"`
	Target      string        `xml:"target,attr,omitempty"`
	CommonName  *CommonName   `xml:"common_name"`
	Description *Description  `xml:"description"`
	PropPattern []PropPattern `xml:"prop_pattern"`
}

// template
//
//    The template contains a collection of metadata about the service.
//    It contains a localizable string that serves as a common,
//    human-readable name for the service.  (This name should be less than
//    60 characters in a single byte locale.)  The template may optionally
//    contain a longer localizable description of the service, a
//    collection of links to documentation, either in the form of manual
//    pages or in the form of URI specifications to external documentation
//    sources (such as docs.sun.com).
//
//    The template has no attributes.
type Template struct {
	XMLName       xml.Name       `xml:"template"`
	CommonName    CommonName     `xml:"common_name"`
	Description   *Description   `xml:"description"`
	Documentation *Documentation `xml:"documentation"`
	PgPattern     []PgPattern    `xml:"pg_pattern"`
}

// Notification Parameters
type ParamVal struct {
	XMLName xml.Name `xml:"paramval"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

// Parameter ...
type Parameter struct {
	XMLName   xml.Name    `xml:"parameter"`
	Name      string      `xml:"name,attr"`
	ValueNode []Valuenode `xml:"value_node"`
}

// Event ...
type Event struct {
	XMLName xml.Name `xml:"event"`
	Value   string   `xml:"value,attr"`
}

// Type ...
type Type struct {
	XMLName   xml.Name    `xml:"type"`
	Name      string      `xml:"name,attr"`
	Active    bool        `xml:"active,attr,omitempty"`
	Parameter []Parameter `xml:"parameter"`
	ParamVal  []ParamVal  `xml:"paramval"`
}

// notification parameters
//
//    This element sets the notification parameters for Software Events and
//    Fault Management problem lifecycle events.
type NotificationParameters struct {
	XMLName xml.Name `xml:"notification_parameters"`
	Event   Event    `xml:"event"`
	Type    []Type   `xml:"type"`
}

// create_default_instance
//
//    A flag element indicating that an otherwise empty default instance
//    of this service (named "default") should be created at install, with
//    its enabled property set as given.
//
//    Its attribute is
//
//	enabled	[boolean] The initial value for the enabled state of
//		this instance.
type CreateDefaultInstance struct {
	XMLName xml.Name `xml:"create_default_instance"`
	Enabled bool     `xml:"enabled,attr"`
}

// single_instance
//
//    A flag element stating that this service can only have a single
//    instance on a particular system.
type SingleInstance struct {
	XMLName xml.Name `xml:"single_instance"`
}

// instance
//
//    The service instance is the object representing a software component
//    that will run on the system if enabled.  It contains an enabled
//    element, a set of dependencies on other services, potentially
//    customized methods or configuration data, an optional method
//    context, and a pointer to its restarter.  (If no restarter is
//    specified, the master restarter, svc.startd(1M), is assumed to be
//    responsible for the service.)
//
//    Its attributes are
//
//	name	The canonical name for this instance of the service.
//
//	enabled	[boolean] The initial value for the enabled state of
//		this instance.
type Instance struct {
	XMLName                xml.Name                 `xml:"instance"`
	Name                   string                   `xml:"name,attr"`
	Enabled                bool                     `xml:"enabled,attr"`
	Restarter              *Restarter               `xml:"restarter"`
	Dependencies           []Dependency             `xml:"dependency"`
	Dependents             []Dependent              `xml:"dependent"`
	MethodContexts         *MethodContext           `xml:"method_context"`
	ExecMethods            []ExecMethod             `xml:"exec_method"`
	NotificationParameters []NotificationParameters `xml:"notification_parameters"`
	PropertyGroups         []PropertyGroup          `xml:"property_group"`
	Template               *Template                `xml:"template"`
}

// service
//
//    The service contains the set of instances defined by default for
//    this service, an optional method execution context, any default
//    methods, the template, and various restrictions or advice applicable
//    at installation.  The method execution context and template elements
//    are required for service_bundle documents with type "manifest", but
//    are optional for "profile" or "archive" documents.
//
//    Its attributes are
//
//	name	The canonical name for the service.
//
//	version	[integer] The integer version for this service.
//
//	type	Whether this service is a simple service, a delegated
//		restarter, or a milestone (a synthetic service that
//		collects a group of dependencies).
//      possible type values are ( service | restarter | milestone )
type Service struct {
	XMLName                xml.Name                 `xml:"service"`
	Name                   string                   `xml:"name,attr"`
	Version                string                   `xml:"version,attr"`
	Type                   string                   `xml:"type,attr"`
	CreateDefaultInstance  *CreateDefaultInstance   `xml:"create_default_instance"`
	SingleInstance         *SingleInstance          `xml:"single_instance"`
	Restarter              *Restarter               `xml:"restarter"`
	Dependency             []Dependency             `xml:"dependency"`
	Dependent              []Dependent              `xml:"dependent"`
	MethodContext          *MethodContext           `xml:"method_context"`
	ExecMethod             []ExecMethod             `xml:"exec_method"`
	NotificationParameters []NotificationParameters `xml:"notification_parameters"`
	PropertyGroup          []PropertyGroup          `xml:"property_group"`
	Instance               []Instance               `xml:"instance"`
	Stability              *Stability               `xml:"stability"`
	Template               *Template                `xml:"template"`
}

// service_bundle
//
//    The bundle possesses two attributes:
//
//	type	How this file is to be understood by the framework (or
//		used in a non-framework compliant way). Standard types
//		are 'archive', 'manifest', and 'profile'.
//
//	name	A name for the bundle.  Manifests should be named after
//		the package which delivered them; profiles should be
//		named after the "feature set nickname" they intend to
//		enable.
type ServiceBundle struct {
	XMLName  xml.Name        `xml:"service_bundle"`
	Type     string          `xml:"type,attr"`
	Name     string          `xml:"name,attr"`
	Services []Service       `xml:"service"`
	Includes []XiInclude     `xml:"xi:include"`
	Bundles  []ServiceBundle `xml:"service_bundle"`
}

//
//  XInclude support
//
//    A series of service bundles may be composed via the xi:include tag.
//    smf(5) tools enforce that all bundles be of the same type.
//
//
//
//     These entities are used for the property, propval and property_group
//     elements, that require type attributes for manifest, while for profiles
//     the type attributes are only implied.
//
type XiInclude struct {
	XMLName   xml.Name   `xml:"xi:include"`
	Href      string     `xml:"href,attr"`
	Parse     string     `xml:"parse,attr"`
	Encoding  string     `xml:"encoding,attr,omitempty"`
	Namespace string     `xml:"xmlns:xi,attr"`
	Fallback  XiFallback `xml:"xi:fallback"`
}

type XiFallback struct {
	XMLName   xml.Name `xml:"xi:fallback"`
	Namespace string   `xml:"xmlns:xi,attr"`
	Children  []byte   `xml:",any"`
}
