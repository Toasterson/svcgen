package svcgen

import (
	"fmt"
	"os/exec"
)

const svcCfgBin string = "/usr/sbin/svccfg"

func Import(path string) error {
	c := exec.Command(svcCfgBin, "import", path)
	output, err := c.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command failed with: %s", output)
	}

	return nil
}
