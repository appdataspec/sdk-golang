package appdataspec

import (
	vosPkg "github.com/appdataspec/sdk-golang/util/vos"
)

//go:generate counterfeiter -o ./fakeAppDataSpec.go --fake-name FakeAppDataSpec ./ AppDataSpec

type AppDataSpec interface {
	// returns the per user app data path; panics if required env vars missing
	GlobalPath() string
	// returns the per user app data path; panics if required env vars missing
	PerUserPath() string
}

func New() AppDataSpec {
	return NewWithVos(vosPkg.New())
}

// allows passing in virtual operating system to decouple from running OS (useful for testing)
func NewWithVos(
	vos vosPkg.Vos,
) AppDataSpec {
	return appDataSpec{
		vos: vos,
	}
}

type appDataSpec struct {
	vos vosPkg.Vos
}
