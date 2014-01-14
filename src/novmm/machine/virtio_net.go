package machine

type VirtioNetDevice struct {
    *VirtioDevice

    // The tap device file descriptor.
    Fd  int `json:"fd"`

    // The mac address.
    Mac string `json:"mac"`
}

func NewVirtioMmioNet(info *DeviceInfo) (Device, error) {
    device, err := NewMmioVirtioDevice(info, []uint{256, 256, 16}, VirtioTypeNet)
    return &VirtioNetDevice{VirtioDevice: device}, err
}

func NewVirtioPciNet(info *DeviceInfo) (Device, error) {
    device, err := NewPciVirtioDevice(info, []uint{256, 256, 16}, PciClassNetwork, VirtioTypeNet)
    return &VirtioNetDevice{VirtioDevice: device}, err
}
