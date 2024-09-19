package types

import (
	"errors"
	"fmt"
	compose "github.com/compose-spec/compose-go/v2/types"
	"os"
	"path"
	"path/filepath"
)

//TODO: Set names

type ContainerVolume interface {
	Name() string
	Target() string
	Type() string
	//Validate() error
}

type VolumeFromDirectory struct {
	Source     string
	TargetDir  string
	VolumeName string
	Readonly   bool
}

func ParseContainerVolume(projectDir string, svcName string, c compose.ServiceVolumeConfig) (*ContainerVolume, error) {
	var cv ContainerVolume
	if c.Source == "" {
		emptyVolume, err := ParseEmptyVolume(svcName, c)
		if err != nil {
			return nil, err
		}
		cv = emptyVolume
		return &cv, nil
	}
	//filePath := path.Join(projectDir, c.Source)
	filePath := c.Source
	file, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	switch mode := file.Mode(); {
	case mode.IsDir():
		volumeFromDir, err := ParseVolumeFromDirectory(projectDir, svcName, c)
		if err != nil {
			return nil, err
		}
		cv = volumeFromDir
		return &cv, nil
	case mode.IsRegular():
		volumeFromFile, err := ParseVolumeFromFile(projectDir, svcName, c)
		if err != nil {
			return nil, err
		}
		cv = volumeFromFile
		return &cv, nil
	}
	return nil, fmt.Errorf("unknown volume source %s", c.Source)
}

func ParseVolumeFromDirectory(projectDir string, svcName string, config compose.ServiceVolumeConfig) (*VolumeFromDirectory, error) {

	if config.Source == "" {
		return nil, errors.New("volume source must be specified")
	}

	if config.Target == "" {
		return nil, errors.New("volume target must be specified")
	}

	sourcePath := path.Join(projectDir, config.Source)
	_, err := filepath.Abs(sourcePath)
	if err != nil {
		return nil, err
	}

	volume := &VolumeFromDirectory{
		VolumeName: svcName,
		Source:     config.Source,
		TargetDir:  config.Target,
		Readonly:   config.ReadOnly,
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

func (v *VolumeFromDirectory) Type() string {
	return "directory"
}

type VolumeFromFile struct {
	Source     string
	TargetFile string
	VolumeName string
	Readonly   bool
}

func ParseVolumeFromFile(projectDir string, svcName string, config compose.ServiceVolumeConfig) (*VolumeFromFile, error) {

	if config.Source == "" {
		return nil, errors.New("volume source must be specified")
	}

	if config.Target == "" {
		return nil, errors.New("volume target must be specified")
	}

	sourcePath := path.Join(projectDir, config.Source)
	_, err := filepath.Abs(sourcePath)
	if err != nil {
		return nil, err
	}

	volume := &VolumeFromFile{
		VolumeName: svcName,
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
func (v *VolumeFromFile) Type() string {
	return "file"
}

type EmptyVolume struct {
	TargetDir  string
	VolumeName string
}

func ParseEmptyVolume(svcName string, config compose.ServiceVolumeConfig) (*EmptyVolume, error) {
	if config.Target == "" {
		return nil, errors.New("volume target must be specified")
	}
	volume := &EmptyVolume{
		VolumeName: svcName,
		TargetDir:  config.Target,
	}
	return volume, nil
}
func (v *EmptyVolume) Name() string {
	return v.VolumeName
}
func (v *EmptyVolume) Target() string {
	return v.TargetDir
}
func (v *EmptyVolume) Type() string {
	return "empty"
}
