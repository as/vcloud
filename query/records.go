package query

import "encoding/xml"

type Link struct {
	Rel  string `xml:"rel,attr,omitempty"`
	Href string `xml:"href,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

// Maps the Record struct names to the URL parameter used to retrieve them
// with the Query service
var UriParams = map[string]string{
	"ApiDefinitionRecord":        "apiDefinition",
	"CatalogItemRecord":          "catalogItem",
	"CatalogRecord":              "catalog",
	"DiskRecord":                 "disk",
	"EventRecord":                "event",
	"FileDescriptorRecord":       "fileDescriptor",
	"GroupRecord":                "group",
	"MediaRecord":                "media",
	"OrgNetworkRecord":           "orgNetwork",
	"OrgVdcRecord":               "orgVdc",
	"OrgVdcStorageProfileRecord": "orgVdcStorageProfile",
	"ServiceRecord":              "service",
	"TaskRecord":                 "task",
	"UserRecord":                 "user",
	"VAppNetworkRecord":          "vAppNetwork",
	"VAppRecord":                 "vApp",
	"VAppTemplateRecord":         "vAppTemplate",
	"VMRecord":                   "vm",
	"VmDiskRelationRecord":       "vmDiskRelation",
}

type ResultRecords struct {
	XMLName  xml.Name `xml:"QueryResultRecords"`
	Type     string   `xml:"type,attr"`
	Name     string   `xml:"name,attr"`
	Href     string   `xml:"href,attr"`
	Total    int      `xml:"total,attr"`
	PageSize int      `xml:"pageSize,attr"`
	Page     int      `xml:"page,attr"`
	Links    Links    `xml:"Link"`

	// Listing of all supported Record types because we can't unmarshal
	// Through an interface
	// See issue 6836. https://code.google.com/p/go/issues/detail?id=6836
	ApiDefinitionRecords        []ApiDefinitionRecord        `xml:"ApiDefinitionRecord"`
	CatalogItemRecords          []CatalogItemRecord          `xml:"CatalogItemRecord"`
	CatalogRecords              []CatalogRecord              `xml:"CatalogRecord"`
	DiskRecords                 []DiskRecord                 `xml:"DiskRecord"`
	EventRecords                []EventRecord                `xml:"EventRecord"`
	FileDescriptorRecords       []FileDescriptorRecord       `xml:"FileDescriptorRecord"`
	GroupRecords                []GroupRecord                `xml:"GroupRecord"`
	MediaRecords                []MediaRecord                `xml:"MediaRecord"`
	OrgNetworkRecords           []OrgNetworkRecord           `xml:"OrgNetworkRecord"`
	OrgVdcRecords               []OrgVdcRecord               `xml:"OrgVdcRecord"`
	OrgVdcStorageProfileRecords []OrgVdcStorageProfileRecord `xml:"OrgVdcStorageProfileRecord"`
	ResultRecords               []ResultRecords              `xml:"QueryResultRecords"`
	ServiceRecords              []ServiceRecord              `xml:"ServiceRecord"`
	TaskRecords                 []TaskRecord                 `xml:"TaskRecord"`
	UserRecords                 []UserRecord                 `xml:"UserRecord"`
	VAppNetworkRecords          []VAppNetworkRecord          `xml:"VAppNetworkRecord"`
	VAppRecords                 []VAppRecord                 `xml:"VAppRecord"`
	VAppTemplateRecords         []VAppTemplateRecord         `xml:"VAppTemplateRecord"`
	VMRecords                   []VMRecord                   `xml:"VMRecord"`
	VmDiskRelationRecords       []VmDiskRelationRecord       `xml:"VmDiskRelationRecord"`

	//Records []Recordable // Deep copy of the relevant record. See issue 6836.
	Records interface{}
}

type ApiDefinitionRecord struct {
	Id               string `xml:"id,attr"`
	Type             string `xml:"type,attr"`
	ApiVendor        string `xml:"apiVendor,attr"`
	EntryPoint       string `xml:"entryPoint,attr"`
	Name             string `xml:"name,attr"`
	Namespace        string `xml:"namespace,attr"`
	Service          string `xml:"service,attr"`
	ServiceName      string `xml:"serviceName,attr"`
	ServiceNamespace string `xml:"serviceNamespace,attr"`
	ServiceVendor    string `xml:"serviceVendor,attr"`
	//Link/g LinkType	 `xml:"Link,attr"`
	//Metadata /g//MetadataType	 `xml:"//Metadata,attr"`
}

type CatalogRecord struct {
	Id                    string `xml:"id,attr"`
	Type                  string `xml:"type,attr"`
	CreationDate          Date   `xml:"creationDate,attr"`
	IsPublished           bool   `xml:"isPublished,attr"`
	IsShared              bool   `xml:"isShared,attr"`
	Name                  string `xml:"name,attr"`
	NumberOfMedia         int    `xml:"numberOfMedia,attr"`
	NumberOfVAppTemplates int    `xml:"numberOfVAppTemplates,attr"`
	OrgName               string `xml:"orgName,attr"`
	Owner                 string `xml:"owner,attr"`
	OwnerName             string `xml:"ownerName,attr"`
	//Link/g LinkType	 `xml:"Link,attr"`
	//Metadata /g//MetadataType	 `xml:"//Metadata,attr"`
}

type CatalogItemRecord struct {
	Id           string `xml:"id,attr"`
	Type         string `xml:"type,attr"`
	Catalog      string `xml:"catalog,attr"`
	CatalogName  string `xml:"catalogName,attr"`
	CreationDate Date   `xml:"creationDate,attr"`
	Entity       string `xml:"entity,attr"`
	EntityName   string `xml:"entityName,attr"`
	EntityType   string `xml:"entityType,attr"`
	IsExpired    bool   `xml:"isExpired,attr"`
	IsPublished  bool   `xml:"isPublished,attr"`
	IsVdcEnabled bool   `xml:"isVdcEnabled,attr"`
	Name         string `xml:"name,attr"`
	Owner        string `xml:"owner,attr"`
	OwnerName    string `xml:"ownerName,attr"`
	Status       string `xml:"status,attr"`
	Vdc          string `xml:"vdc,attr"`
	VdcName      string `xml:"vdcName,attr"`
	//Link/g LinkType	 `xml:"Link,attr"`
	//Metadata /g//MetadataType	 `xml:"//Metadata,attr"`
}

type DiskRecord struct {
	Id                 string `xml:"id,attr"`
	Type               string `xml:"type,attr"`
	BusSubType         string `xml:"busSubType,attr"`
	BusType            string `xml:"busType,attr"`
	BusTypeDesc        string `xml:"busTypeDesc,attr"`
	Datastore          string `xml:"datastore,attr"`
	DatastoreName      string `xml:"datastoreName,attr"`
	IsAttached         bool   `xml:"isAttached,attr"`
	Name               string `xml:"name,attr"`
	OwnerName          string `xml:"ownerName,attr"`
	SizeB              int64  `xml:"sizeB,attr"`
	Status             string `xml:"status,attr"`
	StorageProfile     string `xml:"storageProfile,attr"`
	StorageProfileName string `xml:"storageProfileName,attr"`
	Task               string `xml:"task,attr"`
	Vdc                string `xml:"vdc,attr"`
	VdcName            string `xml:"vdcName,attr"`
	//Link/g LinkType	 `xml:"Link,attr"`
	//Metadata /g//MetadataType	 `xml:"//Metadata,attr"`
}

type EventRecord struct {
	Id               string `xml:"id,attr"`
	Type             string `xml:"type,attr"`
	Entity           string `xml:"entity,attr"`
	EntityName       string `xml:"entityName,attr"`
	EntityType       string `xml:"entityType,attr"`
	EventStatus      int    `xml:"eventStatus,attr"`
	EventType        string `xml:"eventType,attr"`
	OrgName          string `xml:"orgName,attr"`
	ServiceNamespace string `xml:"serviceNamespace,attr"`
	TimeStamp        Date   `xml:"timeStamp,attr"`
	UserName         string `xml:"userName,attr"`
	//Link/g LinkType	 `xml:"Link,attr"`
	//Metadata /g//MetadataType	 `xml:"//Metadata,attr"`
}

type FileDescriptorRecord struct {
	Id               string `xml:"id,attr"`
	Type             string `xml:"type,attr"`
	ApiDefinition    string `xml:"apiDefinition,attr"`
	ApiName          string `xml:"apiName,attr"`
	ApiNamespace     string `xml:"apiNamespace,attr"`
	ApiVendor        string `xml:"apiVendor,attr"`
	FileMimeType     string `xml:"fileMimeType,attr"`
	FileUrl          string `xml:"fileUrl,attr"`
	Name             string `xml:"name,attr"`
	Service          string `xml:"service,attr"`
	ServiceName      string `xml:"serviceName,attr"`
	ServiceNamespace string `xml:"serviceNamespace,attr"`
	ServiceVendor    string `xml:"serviceVendor,attr"`
	//Link/g LinkType	 `xml:"Link,attr"`
	//Metadata /g//MetadataType	 `xml:"//Metadata,attr"`
}

type GroupRecord struct {
	Id                   string `xml:"id,attr"`
	Type                 string `xml:"type,attr"`
	IdentityProviderType string `xml:"identityProviderType,attr"`
	IsReadOnly           bool   `xml:"isReadOnly,attr"`
	Name                 string `xml:"name,attr"`
	RoleName             string `xml:"roleName,attr"`
	//Link/g LinkType	 `xml:"Link,attr"`
	//Metadata /g//MetadataType	 `xml:"//Metadata,attr"`
}
type MediaRecord struct {
	Id                 string `xml:"id,attr"`
	Type               string `xml:"type,attr"`
	Catalog            string `xml:"catalog,attr"`
	CatalogItem        string `xml:"catalogItem,attr"`
	CatalogName        string `xml:"catalogName,attr"`
	CreationDate       Date   `xml:"creationDate,attr"`
	IsBusy             bool   `xml:"isBusy,attr"`
	IsPublished        bool   `xml:"isPublished,attr"`
	Name               string `xml:"name,attr"`
	Org                string `xml:"org,attr"`
	Owner              string `xml:"owner,attr"`
	OwnerName          string `xml:"ownerName,attr"`
	Status             string `xml:"status,attr"`
	StorageB           int64  `xml:"storageB,attr"`
	StorageProfileName string `xml:"storageProfileName,attr"`
	Vdc                string `xml:"vdc,attr"`
	VdcName            string `xml:"vdcName,attr"`
	//Link/g LinkType	 `xml:"Link,attr"`
	//Metadata /g//MetadataType	 `xml:"//Metadata,attr"`
}
type OrgNetworkRecord struct {
	Dns1               string `xml:"dns1,attr"`
	Dns2               string `xml:"dns2,attr"`
	DnsSuffix          string `xml:"dnsSuffix,attr"`
	Gateway            string `xml:"gateway,attr"`
	Href               string `xml:"href,attr"`
	Id                 string `xml:"id,attr"`
	IpScopeId          string `xml:"ipScopeId,attr"`
	IsBusy             bool   `xml:"isBusy,attr"`
	IsIpScopeInherited bool   `xml:"isIpScopeInherited,attr"`
	IsLinked           string `xml:"isLinked,attr"`
	LinkNetworkName    string `xml:"linkNetworkName,attr"`
	Name               string `xml:"name,attr"`
	Netmask            string `xml:"netmask,attr"`
	NetworkPool        string `xml:"networkPool,attr"`
	NetworkPoolName    string `xml:"networkPoolName,attr"`
	Org                string `xml:"org,attr"`
	Type               string `xml:"type,attr"`

	//Link/g LinkType	 `xml:"Link,attr"`
	//Metadata /g//MetadataType	 `xml:"//Metadata,attr"`
}
type OrgVdcRecord struct {
	Id                      string `xml:"id,attr"`
	Type                    string `xml:"type,attr"`
	CpuAllocationMhz        int64  `xml:"cpuAllocationMhz,attr"`
	CpuLimitMhz             int64  `xml:"cpuLimitMhz,attr"`
	CpuUsedMhz              int64  `xml:"cpuUsedMhz,attr"`
	IsBusy                  bool   `xml:"isBusy,attr"`
	IsEnabled               bool   `xml:"isEnabled,attr"`
	IsSystemVdc             bool   `xml:"isSystemVdc,attr"`
	MemoryAllocationMB      int64  `xml:"memoryAllocationMB,attr"`
	MemoryLimitMB           int64  `xml:"memoryLimitMB,attr"`
	MemoryUsedMB            int64  `xml:"memoryUsedMB,attr"`
	Name                    string `xml:"name,attr"`
	Href                    string `xml:"href,attr"`
	NumberOfDatastores      int    `xml:"numberOfDatastores,attr"`
	NumberOfDisks           int    `xml:"numberOfDisks,attr"`
	NumberOfMedia           int    `xml:"numberOfMedia,attr"`
	NumberOfStorageProfiles int    `xml:"numberOfStorageProfiles,attr"`
	NumberOfVAppTemplates   int    `xml:"numberOfVAppTemplates,attr"`
	NumberOfVApps           int    `xml:"numberOfVApps,attr"`
	OrgName                 string `xml:"orgName,attr"`
	ProviderVdc             string `xml:"providerVdc,attr"`
	ProviderVdcName         string `xml:"providerVdcName,attr"`
	Status                  string `xml:"status,attr"`
	StorageAllocationMB     int64  `xml:"storageAllocationMB,attr"`
	StorageLimitMB          int64  `xml:"storageLimitMB,attr"`
	StorageUsedMB           int64  `xml:"storageUsedMB,attr"`
	//Link/g LinkType	 `xml:"Link,attr"`
	//Metadata /g//MetadataType	 `xml:"//Metadata,attr"`
}
type OrgVdcStorageProfileRecord struct {
	Id                      string `xml:"id,attr"`
	Type                    string `xml:"type,attr"`
	IsDefaultStorageProfile bool   `xml:"isDefaultStorageProfile,attr"`
	IsEnabled               bool   `xml:"isEnabled,attr"`
	IsVdcBusy               bool   `xml:"isVdcBusy,attr"`
	Name                    string `xml:"name,attr"`
	NumberOfConditions      int    `xml:"numberOfConditions,attr"`
	StorageLimitMB          int    `xml:"storageLimitMB,attr"`
	StorageUsedMB           int    `xml:"storageUsedMB,attr"`
	Vdc                     string `xml:"vdc,attr"`
	VdcName                 string `xml:"vdcName,attr"`
	//Link/g LinkType	 `xml:"Link,attr"`
	//Metadata /g//MetadataType	 `xml:"//Metadata,attr"`
}
type ServiceRecord struct {
	Id        string `xml:"id,attr"`
	Type      string `xml:"type,attr"`
	Name      string `xml:"name,attr"`
	Namespace string `xml:"namespace,attr"`
	Vendor    string `xml:"vendor,attr"`
	//Link/g LinkType	 `xml:"Link,attr"`
	//Metadata /g//MetadataType	 `xml:"//Metadata,attr"`
}

type TaskRecord struct {
	XMLName          xml.Name `xml:"TaskRecord"`
	Id               string   `xml:"id,attr"`
	Type             string   `xml:"type,attr"`
	EndDate          Date     `xml:"endDate,attr"`
	Name             string   `xml:"name,attr"`
	Object           string   `xml:"object,attr"`
	ObjectName       string   `xml:"objectName,attr"`
	ObjectType       string   `xml:"objectType,attr"`
	Org              string   `xml:"org,attr"`
	OrgName          string   `xml:"orgName,attr"`
	OwnerName        string   `xml:"ownerName,attr"`
	ServiceNamespace string   `xml:"serviceNamespace,attr"`
	StartDate        Date     `xml:"startDate,attr"`
	Status           string   `xml:"status,attr"`
	//Link/g LinkType	 `xml:"Link,attr"`
	//Metadata /g//MetadataType	 `xml:"//Metadata,attr"`
}
type UserRecord struct {
	XMLName             xml.Name `xml:"UserRecord"`
	Type                string   `xml:"type,attr"`
	Name                string   `xml:"name,attr"`
	Href                string   `xml:"href,attr"`
	StoredVMQuotaRank   string   `xml:"storedVMQuotaRank,attr"`
	StoredVMQuota       string   `xml:"storedVMQuota,attr"`
	NumberOfStoredVMs   int      `xml:"numberOfStoredVMs,attr"`
	NumberOfDeployedVMs int      `xml:"numberOfDeployedVMs,attr"`
	IsLdapUser          string   `xml:"isLdapUser,attr"`
	FullName            string   `xml:"fullName,attr"`
	IsEnabled           string   `xml:"isEnabled,attr"`
	DeployedVMQuotaRank string   `xml:"deployedVMQuotaRank,attr"`
	DeployedVMQuota     string   `xml:"deployedVMQuota,attr"`
}
type VAppNetworkRecord struct {
	Id                 string `xml:"id,attr"`
	Type               string `xml:"type,attr"`
	Dns1               string `xml:"dns1,attr"`
	Dns2               string `xml:"dns2,attr"`
	DnsSuffix          string `xml:"dnsSuffix,attr"`
	Gateway            string `xml:"gateway,attr"`
	IpScopeId          string `xml:"ipScopeId,attr"`
	IsBusy             bool   `xml:"isBusy,attr"`
	IsIpScopeInherited bool   `xml:"isIpScopeInherited,attr"`
	Name               string `xml:"name,attr"`
	Netmask            string `xml:"netmask,attr"`
	VApp               string `xml:"vApp,attr"`
	VAppName           string `xml:"vAppName,attr"`
	//Link/g LinkType	 `xml:"Link,attr"`
	//Metadata /g//MetadataType	 `xml:"//Metadata,attr"`
}
type VAppRecord struct {
	AutoUndeployDate    Date     `xml:"autoUndeployDate,attr"`
	CpuAllocationInMhz  int64    `xml:"cpuAllocationInMhz,attr"`
	CreationDate        Date     `xml:"creationDate,attr"`
	Href                string   `xml:"href,attr"`
	IsBusy              bool     `xml:"isBusy,attr"`
	IsDeployed          bool     `xml:"isDeployed,attr"`
	IsEnabled           bool     `xml:"isEnabled,attr"`
	IsExpired           bool     `xml:"isExpired,attr"`
	IsInMaintenanceMode bool     `xml:"isInMaintenanceMode,attr"`
	IsPublic            bool     `xml:"isPublic,attr"`
	MemoryAllocationMB  int64    `xml:"memoryAllocationMB,attr"`
	Name                string   `xml:"name,attr"`
	NumberOfVMs         int      `xml:"numberOfVMs,attr"`
	OwnerName           string   `xml:"ownerName,attr"`
	Status              string   `xml:"status,attr"`
	StorageKB           int64    `xml:"storageKB,attr"`
	Type                string   `xml:"type,attr"`
	Vdc                 string   `xml:"vdc,attr"`
	VdcName             string   `xml:"vdcName,attr"`
	XMLName             xml.Name `xml:"VAppRecord"`
}

type VAppTemplateRecord struct {
	CatalogName        string `xml:"catalogName,attr"`
	CpuAllocationInMhz int64  `xml:"cpuAllocationInMhz,attr"`
	CreationDate       Date   `xml:"creationDate,attr"`
	Href               string `xml:"href,attr"`
	Id                 string `xml:"id,attr"`
	IsBusy             bool   `xml:"isBusy,attr"`
	IsDeployed         bool   `xml:"isDeployed,attr"`
	IsEnabled          bool   `xml:"isEnabled,attr"`
	IsExpired          bool   `xml:"isExpired,attr"`
	IsGoldMaster       bool   `xml:"isGoldMaster,attr"`
	IsPublished        bool   `xml:"isPublished,attr"`
	MemoryAllocationMB int64  `xml:"memoryAllocationMB,attr"`
	Name               string `xml:"name,attr"`
	NumberOfVMs        int    `xml:"numberOfVMs,attr"`
	Org                string `xml:"org,attr"`
	OwnerName          string `xml:"ownerName,attr"`
	Status             string `xml:"status,attr"`
	StorageKB          int64  `xml:"storageKB,attr"`
	StorageProfileName string `xml:"storageProfileName,attr"`
	Type               string `xml:"type,attr"`
	Vdc                string `xml:"vdc,attr"`
	VdcName            string `xml:"vdcName,attr"`
}

type VmDiskRelationRecord struct {
	Id   string `xml:"id,attr"`
	Type string `xml:"type,attr"`
	Disk string `xml:"disk,attr"`
	Vdc  string `xml:"vdc,attr"`
	Vm   string `xml:"vm,attr"`
	//Link LinkType	 `xml:"Link,attr"`
	//Metadata //MetadataType	 `xml:"//Metadata,attr"`
}

type VMRecord struct {
	XMLName             xml.Name `xml:"VMRecord"`
	Id                  string   `xml:"id,attr"`
	Type                string   `xml:"type,attr"`
	Href                string   `xml:"href,attr"`
	CatalogName         string   `xml:"catalogName,attr"`
	Container           string   `xml:"container,attr"`
	ContainerName       string   `xml:"containerName,attr"`
	GuestOs             string   `xml:"guestOs,attr"`
	HardwareVersion     int      `xml:"hardwareVersion,attr"`
	IsBusy              bool     `xml:"isBusy,attr"`
	IsDeleted           bool     `xml:"isDeleted,attr"`
	IsDeployed          bool     `xml:"isDeployed,attr"`
	IsInMaintenanceMode bool     `xml:"isInMaintenanceMode,attr"`
	IsPublished         bool     `xml:"isPublished,attr"`
	IsVAppTemplate      bool     `xml:"isVAppTemplate,attr"`
	MemoryMB            int      `xml:"memoryMB,attr"`
	Name                string   `xml:"name,attr"`
	NumberOfCpus        int      `xml:"numberOfCpus,attr"`
	Status              string   `xml:"status,attr"`
	StorageProfileName  string   `xml:"storageProfileName,attr"`
	VmToolsVersion      int      `xml:"vmToolsVersion,attr"`
	Vdc                 string   `xml:"vdc,attr"`
	//Link LinkType	 `xml:"Link,attr"`
	//Metadata MetadataType	 `xml:"Metadata,attr"`
}
