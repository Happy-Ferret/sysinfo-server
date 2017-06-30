package main

// SysInfo type anything
type SysInfo struct {
	Node    Node            `json:"node"`
	OS      OS              `json:"os"`
	Kernel  Kernel          `json:"kernel"`
	CPU     CPU             `json:"cpu"`
	Memory  Memory          `json:"memory"`
	Network []NetworkDevice `json:"network,omitempty"`
}

type Node struct {
	Hostname   string `json:"hostname,omitempty"`
	MachineID  string `json:"machineid,omitempty"`
	Hypervisor string `json:"hypervisor,omitempty"`
}

type OS struct {
	Vendor  string `json:"vendor,omitempty"`
	Version string `json:"version,omitempty"`
}

type Kernel struct {
	Release string `json:"release,omitempty"`
}

type CPU struct {
	Model string `json:"model,omitempty"`
	Cores uint   `json:"cores,omitempty"`
}

type Memory struct {
	Size uint `json:"size,omitempty"`
}

type NetworkDevice struct {
	Name string `json:"name,omitempty"`
	IP   string `json:"ip,omitempty"`
}
