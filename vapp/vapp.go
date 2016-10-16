package vapp

type VApp struct {
	Deployed              bool   `xml:"deployed,attr,omitempty"`
	Href                  string `xml:"href,attr,omitempty"`
	Id                    string `xml:"id,attr,omitempty"`
	Name                  string `xml:"name,attr,omitempty"`
	OvfDescriptorUploaded bool   `xml:"ovfDescriptorUploaded,attr,omitempty"`
	Status                int    `xml:"status,attr,omitempty"`
	Type                  string `xml:"type,attr,omitempty"`

	Links                []Link               `xml:"Link"`
	Description          string               `xml:"Description"`
	LeaseSettingsSection LeaseSettingsSection `xml:"LeaseSettingsSection"`
	StartupSection       StartupSection       `xml:"http://schemas.dmtf.org/ovf/envelope/1 StartupSection"`
	NetworkConfigSection NetworkConfigSection `xml:"NetworkConfigSection"`
	Owner                Owner                `xml:"Owner"`
	InMaintenanceMode    bool                 `xml:"InMaintenanceMode"`
	Vms                  []Vm                 `xml:"Children>Vm"`
}

type User struct {
	Type string `xml:"type,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
	Href string `xml:"href,attr,omitempty"`
}

type Owner struct {
	Users []User `xml:"User"`
}

type LeaseSettingsSection struct {
	Type string `xml:"type,attr,omitempty"`
	Href string `xml:"href,attr,omitempty"`

	Links                     []Link `xml:"Link"`
	DeploymentLeaseInSeconds  int64  `xml:"DeploymentLeaseInSeconds"`
	StorageLeaseInSeconds     int64  `xml:"StorageLeaseInSeconds"`
	DeploymentLeaseExpiration string `xml:"DeploymentLeaseExpiration"`
}

type StartupSection struct {
	Href  string `xml:"http://www.vmware.com/vcloud/v1.5 href,attr,omitempty"`
	Type  string `xml:"http://www.vmware.com/vcloud/v1.5 type,attr,omitempty"`
	Info  string `xml:"http://schemas.dmtf.org/ovf/envelope/1 Info"`
	Items []Item `xml:"http://schemas.dmtf.org/ovf/envelope/1 Item"`
	Links []Link `xml:"http://www.vmware.com/vcloud/v1.5 Link"`
}

type NetworkConfigSection struct {
	Info string `xml:"OVF Info"`
	Type string `xml:"type,attr,omitempty"`
	Href string `xml:"href,attr,omitempty"`

	Links          []Link          `xml:"Link"`
	NetworkConfigs []NetworkConfig `xml:"NetworkConfig"`
	Description    string          `xml:"Description"`
}

type Link struct {
	Rel  string `xml:"rel,attr,omitempty"`
	Href string `xml:"href,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

type NetworkConfig struct {
	NetworkName                    string        `xml:"networkName,attr,omitempty"`
	Links                          []Link        `xml:"Link"`
	Description                    string        `xml:"Description"`
	IpScopes                       []IpScope     `xml:"Configuration>IpScope"`
	ParentNetwork                  ParentNetwork `xml:"Configuration>ParentNetwork,omitempty"`
	FenceMode                      string        `xml:"Configuration>FenceMode"`
	RetainNetInfoAcrossDeployments bool          `xml:"Configuration>RetainNetInfoAcrossDeployments"`
	Features                       []Feature     `xml:"Configuration>Features"`
	IsDeployed                     bool          `xml:"IsDeployed"`
}

type IpScope struct {
	IsInherited         bool      `xml:"IsInherited"`
	Gateway             string    `xml:"Gateway"`
	Netmask             string    `xml:"Netmask"`
	IpRanges            []IpRange `xml:"IpRanges>IpRange"`
	AllocatedIpAdresses []string  `xml:"AllocatedIpAddresses>IpAddress"`
}

type IpRange struct {
	StartAddress string `xml:"StartAddress"`
	EndAddress   string `xml:"EndAddress"`
}
type ParentNetwork struct {
	Name string `xml:"id,attr,omitempty"`
	Id   string `xml:"name,attr,omitempty"`
	Href string `xml:"href,attr,omitempty"`
}

type Item struct {
	StopDelay   int    `xml:"stopDelay,attr,omitempty"`
	StopAction  string `xml:"stopAction,attr,omitempty"`
	StartDelay  int    `xml:"startDelay,attr,omitempty"`
	StartAction string `xml:"startAction,attr,omitempty"`
	Order       string `xml:"order,attr,omitempty"`
	Id          string `xml:"id,attr,omitempty"`
}

type Feature struct {
	DhcpService     DhcpService     `xml:"DhcpService"`
	NatService      NatService      `xml:"NatService"`
	FirewallService FirewallService `xml:"FirewallService,omitempty"`
}

type DhcpService struct {
	IsEnabled        bool  `xml:"IsEnabled"`
	DefaultLeaseTime int64 `xml:"DefaultLeaseTime"`
	MaxLeaseTime     int64 `xml:"MaxLeaseTime"`
}

type NatService struct {
	IsEnabled bool      `xml:"IsEnabled"`
	NatType   string    `xml:"NatType"`
	Policy    string    `xml:"Policy"`
	NatRules  []NatRule `xml:"NatRule"`
}

type FirewallService struct {
	IsEnabled        bool           `xml:"IsEnabled,omitempty"`
	DefaultAction    string         `xml:"DefaultAction,omitempty"`
	LogDefaultAction string         `xml:"LogDefaultAction,omitempty"`
	FirewallRules    []FirewallRule `xml:"FirewallRule"`
}

type NatRule struct {
	OneToOneVmRules []OneToOneVmRule `xml:"OneToOneVmRule"`
}

type OneToOneVmRule struct {
	MappingMode       string `xml:"MappingMode"`
	ExternalIpAddress string `xml:"ExternalIpAddress"`
	VAppScopedVmId    string `xml:"VAppScopedVmId"`
	VmNicId           string `xml:"VmNicId"`
}

type FirewallRule struct {
	IsEnabled     bool       `xml:"IsEnabled"`
	Description   string     `xml:"Description"`
	Policy        string     `xml:"Policy"`
	Port          int        `xml:"Port"`
	DestinationIp string     `xml:"DestinationIp"`
	SourcePort    int        `xml:"SourcePort"`
	SourceIp      string     `xml:"SourceIp"`
	EnableLogging bool       `xml:"EnableLogging"`
	Protocols     []Protocol `xml:"Protocols"`
}

type Protocol struct {
	Tcp bool `xml:"Tcp"`
}

type Vm struct {
	Deployed           string `xml:"deployed,attr"`
	Href               string `xml:"href,attr"`
	Id                 string `xml:"id,attr"`
	Name               string `xml:"name,attr"`
	NeedsCustomization string `xml:"needsCustomization,attr"`
	Status             string `xml:"status,attr"`
	Type               string `xml:"type,attr"`

	Links                     []Link                    `xml:"Link"`
	Description               string                    `xml:"Description"`
	VirtualHardwareSection    VirtualHardwareSection    `xml:"http://schemas.dmtf.org/ovf/envelope/1 VirtualHardwareSection"`
	OperatingSystemSection    OperatingSystemSection    `xml:"http://schemas.dmtf.org/ovf/envelope/1 OperatingSystemSection"`
	NetworkConnectionSection  NetworkConnectionSection  `xml:"NetworkConnectionSection"`
	GuestCustomizationSection GuestCustomizationSection `xml:"GuestCustomizationSection"`
	RuntimeInfoSection        RuntimeInfoSection        `xml:"RuntimeInfoSection"`
	VAppScopedLocalId         string                    `xml:"VAppScopedLocalId"`
	Environment               Environment               `xml:"http://schemas.dmtf.org/ovf/environment/1 Environment"`
}

type VirtualHardwareSection struct {
	Href   string `xml:"http://www.vmware.com/vcloud/v1.5 href,attr,omitempty"`
	Type   string `xml:"http://www.vmware.com/vcloud/v1.5 type,attr,omitempty"`
	Info   string `xml:"Info"`
	System System `xml:"System"`
	//Items []Item `xml:"http://schemas.dmtf.org/ovf/envelope/1 Item"` //TODO
	Links []Link `xml:"http://www.vmware.com/vcloud/v1.5 Link"`
}

type System struct {
	ElementName             string `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_VirtualSystemSettingData ElementName"`
	InstanceID              int64  `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_VirtualSystemSettingData InstanceID"`
	VirtualSystemIdentifier string `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_VirtualSystemSettingData VirtualSystemIdentifier"`
	VirtualSystemType       string `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_VirtualSystemSettingData VirtualSystemType"`
}

type OperatingSystemSection struct {
	Href        string `xml:"http://www.vmware.com/vcloud/v1.5 href,attr,omitempty"`
	Type        string `xml:"http://www.vmware.com/vcloud/v1.5 type,attr,omitempty"`
	Id          string `xml:"id,attr,omitempty"`
	OsType      string `xml:"http://www.vmware.com/schema/ovf osType,attr,omitempty"`
	Info        string `xml:"Info"`
	Description string `xml:"Description"`
	Links       []Link `xml:"http://www.vmware.com/vcloud/v1.5 Link"`
}

type NetworkConnectionSection struct {
	Href                          string              `xml:"href,attr,omitempty"`
	Type                          string              `xml:"type,attr,omitempty"`
	Info                          string              `xml:"http://schemas.dmtf.org/ovf/environment/1 Info"`
	PrimaryNetworkConnectionIndex string              `xml:"PrimaryNetworkConnectionIndex,omitempty"`
	NetworkConnections            []NetworkConnection `xml:"NetworkConnection,omitempty"`
	Links                         []Link              `xml:"Link"`
}

type NetworkConnection struct {
	Network                 string `xml:"network,attr,omitempty"`
	NeedsCustomization      bool   `xml:"needsCustomization,attr,omitempty"`
	NetworkConnectionIndex  string `xml:"NetworkConnectionIndex"`
	IpAddress               string `xml:"IpAddress"`
	ExternalIpAddress       string `xml:"ExternalIpAddress"`
	IsConnected             bool   `xml:"IsConnected"`
	MACAddress              string `xml:"MACAddress"`
	IpAddressAllocationMode string `xml:"IpAddressAllocationMode"`
}

type GuestCustomizationSection struct {
	Href                  string `xml:"href,attr,omitempty"`
	Type                  string `xml:"type,attr,omitempty"`
	Info                  string `xml:"http://schemas.dmtf.org/ovf/environment/1 Info"`
	Enabled               bool   `xml:"Enabled"`
	ChangeSid             bool   `xml:"ChangeSid"`
	VirtualMachineId      string `xml:"VirtualMachineId"`
	JoinDomainEnabled     bool   `xml:"JoinDomainEnabled"`
	UseOrgSettings        bool   `xml:"UseOrgSettings"`
	AdminPasswordEnabled  bool   `xml:"AdminPasswordEnabled"`
	AdminPasswordAuto     bool   `xml:"AdminPasswordAuto"`
	ResetPasswordRequired bool   `xml:"ResetPasswordRequired"`
	CustomizationScript   string `xml:"CustomizationScript"`
	ComputerName          string   `xml:"ComputerName"`
	Links                 []Link `xml:"Link"`
}

type RuntimeInfoSection struct {
	Href        string      `xml:"href,attr,omitempty"`
	Type        string      `xml:"type,attr,omitempty"`
	Info        string      `xml:"http://schemas.dmtf.org/ovf/environment/1 Info"`
	VMWareTools VMWareTools `xml:"VMWareTools"`
}

type VMWareTools struct {
	version string `xml:"version,attr"`
}

type Environment struct {
}
