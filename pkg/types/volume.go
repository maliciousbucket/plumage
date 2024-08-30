package types

import (
	"errors"
	"fmt"
	compose "github.com/compose-spec/compose-go/v2/types"
)

//TODO: Set names

type ContainerVolume interface {
	Name() string
	Target() string
	//Validate() error
}

type VolumeFromDirectory struct {
	Source     string
	TargetDir  string
	VolumeName string
	Readonly   bool
}

func ParseContainerVolume(c compose.ServiceVolumeConfig) (*ContainerVolume, error) {
	if c.Source == "" {
		emptyVolume, err := ParseEmptyVolume(c)
		if err != nil {
			return nil, err
		}
		var cv ContainerVolume = emptyVolume
		return &cv, nil
	}

	return nil, fmt.Errorf("unknown volume source %s", c.Source)
}

func ParseVolumeFromDirectory(config compose.ServiceVolumeConfig) (*VolumeFromDirectory, error) {

	if config.Source == "" {
		return nil, errors.New("volume source must be specified")
	}

	if config.Target == "" {
		return nil, errors.New("volume target must be specified")
	}

	volume := &VolumeFromDirectory{
		Source:    config.Source,
		TargetDir: config.Target,
		Readonly:  config.ReadOnly,
	}

	return volume, nil
}

func (v *VolumeFromDirectory) Name() string {
	return v.VolumeName
}

func (v *VolumeFromDirectory) Directory() string {
	return v.Source
}

func (v *VolumeFromDirectory) Target() string {
	return v.TargetDir
}

type VolumeFromFile struct {
	Source     string
	TargetFile string
	VolumeName string
	Readonly   bool
}

func ParseVolumeFromFile(config compose.ServiceVolumeConfig) (*VolumeFromFile, error) {

	if config.Source == "" {
		return nil, errors.New("volume source must be specified")
	}

	if config.Target == "" {
		return nil, errors.New("volume target must be specified")
	}

	volume := &VolumeFromFile{
		Source:     config.Source,
		TargetFile: config.Target,
		Readonly:   config.ReadOnly,
	}

	return volume, nil
}

func (v *VolumeFromFile) Name() string {
	return v.VolumeName
}

func (v *VolumeFromFile) File() string {
	return v.Source
}
func (v *VolumeFromFile) Target() string {
	return v.TargetFile
}

type EmptyVolume struct {
	TargetDir  string
	VolumeName string
}

func ParseEmptyVolume(config compose.ServiceVolumeConfig) (*EmptyVolume, error) {
	if config.Target == "" {
		return nil, errors.New("volume source must be specified")
	}
	volume := &EmptyVolume{
		TargetDir: config.Target,
	}
	return volume, nil
}
func (v *EmptyVolume) Name() string {
	return v.VolumeName
}
func (v *EmptyVolume) Target() string {
	return v.TargetDir
}
