package main

// MachineID description
type MachineID struct {
	MachineID string `json:"machineid"`
	Hostname  string `json:"hostname"`
}

// Packages description
type Packages struct {
	Packages []string `json:"packages"`
}

// SysInfo description
type SysInfo struct {
	Node    Node            `json:"node,omitempty"`
	OS      OS              `json:"os,omitempty"`
	Kernel  Kernel          `json:"kernel,omitempty"`
	Product Product         `json:"product,omitempty"`
	CPU     CPU             `json:"cpu,omitempty"`
	Memory  Memory          `json:"memory,omitempty"`
	LVM     []LogicalVolume `json:"lvm,omitempty"`
	Network []NetworkDevice `json:"network,omitempty"`
}

// Node description
type Node struct {
	Hostname   string `json:"hostname,omitempty"`
	MachineID  string `json:"machineid,omitempty"`
	Hypervisor string `json:"hypervisor,omitempty"`
}

// OS description
type OS struct {
	Vendor  string `json:"vendor,omitempty"`
	Version string `json:"version,omitempty"`
}

// Kernel description
type Kernel struct {
	Release string `json:"release,omitempty"`
}

// CPU description
type CPU struct {
	Model string `json:"model,omitempty"`
	Cores uint   `json:"cores,omitempty"`
}

// Memory description
type Memory struct {
	Size uint `json:"size,omitempty"`
}

// LogicalVolume description
type LogicalVolume struct {
	LVName string  `json:"lvname,omitempty"`
	VGName string  `json:"vgname,omitempty"`
	LVSize float64 `json:"lvsize,omitempty"`
}

// NetworkDevice description
type NetworkDevice struct {
	Name       string `json:"name,omitempty"`
	MACAddress string `json:"macaddress,omitempty"`
	IP         string `json:"ip,omitempty"`
}

// Product description.
type Product struct {
	Name    string `json:"name,omitempty"`
	Vendor  string `json:"vendor,omitempty"`
	Version string `json:"version,omitempty"`
	Serial  string `json:"serial,omitempty"`
}
