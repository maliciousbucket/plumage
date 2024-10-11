package helm

import "fmt"

func (c *helmClient) UninstallRelease(name string) error {
	return c.uninstallRelease(name)
}

func (c *helmClient) uninstallRelease(name string) error {
	err := c.Client.UninstallReleaseByName(name)
	if err != nil {
		return fmt.Errorf("error uninstalling %s: %s", name, err)
	}
	return nil
}
