package k8s

// Represents a Persistent Disk resource in Google Compute Engine.
//
// A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.
type GcePersistentDiskVolumeSource struct {
	// pdName is unique name of the PD resource in GCE.
	//
	// Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	PdName *string `field:"required" json:"pdName" yaml:"pdName"`
	// fsType is filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// partition is the partition in the volume that you want to mount.
	//
	// If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
	// readOnly here will force the ReadOnly setting in VolumeMounts.
	//
	// Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// Default: false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}
