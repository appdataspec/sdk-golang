package appdataspec

import (
	"fmt"
)

const (
	globalPathTemplate  = "/Library/Application Support"
	perUserPathTemplate = "%v/Library/Application Support"
)

func (this appDataSpec) GlobalPath() string {
	return globalPathTemplate
}

func (this appDataSpec) PerUserPath() string {
	localAppDataEnvVar := this.vos.Getenv("HOME")
	if "" == localAppDataEnvVar {
		panic("Unable to determine per user app data path. Error was: HOME env var required")
	}
	return fmt.Sprintf(
		perUserPathTemplate,
		localAppDataEnvVar,
	)
}
