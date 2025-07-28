package main

type InfoRetriever interface {
	SystemInfo() (*SystemInfo, error)
	LiveStats() (*LiveStats, error)
}

type SystemInfo struct {
	OS            string         `json:"os"`
	OSVersion     OptionalString `json:"os_version,omitzero"`
	OSVersionName OptionalString `json:"os_version_name,omitzero"`

	KernelVersion OptionalString `json:"kernel_version,omitzero"`
	Distribution  OptionalString `json:"distribution,omitzero"`

	CPUModel        OptionalString `json:"cpu_model,omitzero"`
	CPUCores        Optional[int]  `json:"cpu_cores,omitzero"`
	CPUThreads      Optional[int]  `json:"cpu_threads,omitzero"`
	CPUSpeedMHz     Optional[int]  `json:"cpu_speed_mhz,omitzero"`
	CPUArchitecture string         `json:"cpu_architecture"`

	MemoryTotalMB int `json:"memory_total_mb"`

	NetworkAdaptors []Adaptor `json:"network_adaptors"`

	Hostname string `json:"hostname"`
}

type Adaptor struct {
	Name           string             `json:"name,omitzero"`            // network adaptor name
	MacAddress     OptionalString     `json:"mac_address,omitzero"`     // MAC address
	ConfiguredIPs  Optional[[]string] `json:"configured_ips,omitzero"`  // configured IP addresses
	SubnetMask     OptionalString     `json:"subnet_mask,omitzero"`     // subnet mask
	DefaultGateway OptionalString     `json:"default_gateway,omitzero"` // default gateway
}

type LiveStats struct {
	CpuUtilization float64           `json:"cpu_utilization,omitzero"` // as a percentage
	CPUTemperature Optional[float64] `json:"cpu_temperature,omitzero"` // in Celsius

	MemoryUsageMB int `json:"memory_usage"` // current memory usage in MB

	Disks Optional[[]Disk] `json:"disks"` // disk usage statistics

	Processes Optional[[]Process] `json:"processes"` // running processes

	UptimeSeconds int `json:"uptime_seconds"` // system uptime in seconds
}

type Disk struct {
	Name         string         `json:"name"`                    // disk name
	UsedSpaceMB  Optional[int]  `json:"used_space_mb,omitzero"`  // used space in MB
	TotalSpaceMB Optional[int]  `json:"total_space_mb,omitzero"` // total space in MB
	MountPoint   OptionalString `json:"mount_point,omitzero"`    // mount point
}

type Process struct {
	PID  int            `json:"pid"`           // process ID
	Name OptionalString `json:"name,omitzero"` // process name

	CpuUsage Optional[float64] `json:"cpu_usage,omitzero"` // as a percentage
	MemoryMB Optional[int]     `json:"memory_mb,omitzero"` // memory usage in MB

	Threads Optional[int] `json:"threads,omitzero"` // number of threads
}
