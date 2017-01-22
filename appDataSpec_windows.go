package appdataspec

import (
	"fmt"
)

const (
	globalPathTemplate  = "%v"
	perUserPathTemplate = "%v"
)

func (this appDataSpec) GlobalPath() string {
	programDataEnvVar := this.vos.Getenv("PROGRAMDATA")
	if "" == programDataEnvVar {
		panic("Unable to determine global app data path. Error was: PROGRAMDATA env var required")
	}
	return fmt.Sprintf(globalPathTemplate, programDataEnvVar)
}

func (this appDataSpec) PerUserPath() string {
	localAppDataEnvVar := this.vos.Getenv("LOCALAPPDATA")
	if "" == localAppDataEnvVar {
		panic("Unable to determine per user app data path. Error was: LOCALAPPDATA env var required")
	}
	return fmt.Sprintf(
		perUserPathTemplate,
		localAppDataEnvVar,
	)
}
