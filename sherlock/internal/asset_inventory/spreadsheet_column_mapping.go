package asset_inventory

// spreadsheetColumnMapping is how we interact with the assetRow ([]any) type that represents a row
// in the actual Google Sheet.
//
// Here's the deal--there's a lot of ways you could dice this onion. This approach here helps us do
// two cool things:
//  1. We actually get a lot of type safety here. We need a string to match to each header cell, but
//     having it as a struct tag means we can't mismatch (and accessing it later is super simple).
//  2. Just parsing indices and going from there makes this system incredibly tolerant of common
//     operations like reordering columns, adding new columns, editing cells, etc. By contrast, if
//     we fully parsed each row into a struct and back, we'd have to be a lot stricter about what
//     human behaviors we could tolerate.
type spreadsheetColumnMapping struct {
	TeamColumnIndex                          int `column:"Team"`
	TerraServiceComponentColumnIndex         int `column:"Terra Service Component"`
	GcpProjectOrAzureSubscriptionColumnIndex int `column:"GCP Project or Azure Subscription"`
	AssetShortNameColumnIndex                int `column:"Asset Short Name"`
	UniqueAssetIdentifierColumnIndex         int `column:"Unique Asset Identifier"`
	Ipv4OrIpv6AddressColumnIndex             int `column:"IPv4 or IPv6"`
	VirtualColumnIndex                       int `column:"Virtual"`
	PublicColumnIndex                        int `column:"Public"`
	DnsNameOrUrlColumnIndex                  int `column:"DNS Name or URL"`
	NetBiosNameColumnIndex                   int `column:"NetBIOS Name"`
	MacAddressColumnIndex                    int `column:"MAC Address"`
	AuthenticatedScanColumnIndex             int `column:"Authenticated Scan"`
	BaselineConfigurationNameColumnIndex     int `column:"Baseline Configuration Name"`
	OsNameAndVersionColumnIndex              int `column:"OS Name and Version"`
	LocationColumnIndex                      int `column:"Location"`
	AssetTypeColumnIndex                     int `column:"Asset Type"`
	HardwareMakeModelColumnIndex             int `column:"Hardware Make/Model"`
	InLatestScanColumnIndex                  int `column:"In Latest Scan"`
	SoftwareDatabaseVendorColumnIndex        int `column:"Software/Database Vendor"`
	SoftwareDatabaseNameVersionColumnIndex   int `column:"Software/Database Name & Version"`
	PatchLevelColumnIndex                    int `column:"Patch Level"`
	DiagramLabelColumnIndex                  int `column:"Diagram Label"`
	CommentsColumnIndex                      int `column:"Comments"`
	SerialAssetTagColumnIndex                int `column:"Serial #/Asset Tag#"`
	VlanNetworkIDColumnIndex                 int `column:"VLAN/Network ID"`
	SystemAdministratorOwnerColumnIndex      int `column:"System Administrator/Owner"`
	ApplicationAdministratorOwnerColumnIndex int `column:"Application Administrator/Owner"`
	FunctionColumnIndex                      int `column:"Function"`
	EndOfLifeDateColumnIndex                 int `column:"End-of-Life"`
}

// Don't bother editing these one by one. Multicursor is your friend here.

func (mapping *spreadsheetColumnMapping) getTeam(row assetRow) string {
	if value, ok := row[mapping.TeamColumnIndex].(string); !ok {
		return ""
	} else {
		return value
	}
}
func (mapping *spreadsheetColumnMapping) getTerraServiceComponent(row assetRow) string {
	if value, ok := row[mapping.TerraServiceComponentColumnIndex].(string); !ok {
		return ""
	} else {
		return value
	}
}
func (mapping *spreadsheetColumnMapping) getGcpProjectOrAzureSubscription(row assetRow) string {
	if value, ok := row[mapping.GcpProjectOrAzureSubscriptionColumnIndex].(string); !ok {
		return ""
	} else {
		return value
	}
}
func (mapping *spreadsheetColumnMapping) getUniqueAssetIdentifier(row assetRow) string {
	if value, ok := row[mapping.UniqueAssetIdentifierColumnIndex].(string); !ok {
		return ""
	} else {
		return value
	}
}
func (mapping *spreadsheetColumnMapping) getAssetType(row assetRow) string {
	if value, ok := row[mapping.AssetTypeColumnIndex].(string); !ok {
		return ""
	} else {
		return value
	}
}

func (mapping *spreadsheetColumnMapping) setTeam(row assetRow, value string) {
	row[mapping.TeamColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setTerraServiceComponent(row assetRow, value string) {
	row[mapping.TerraServiceComponentColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setAssetShortName(row assetRow, value string) {
	row[mapping.AssetShortNameColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setGcpProjectOrAzureSubscription(row assetRow, value string) {
	row[mapping.GcpProjectOrAzureSubscriptionColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setUniqueAssetIdentifier(row assetRow, value string) {
	row[mapping.UniqueAssetIdentifierColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setIpv4OrIpv6Address(row assetRow, value string) {
	row[mapping.Ipv4OrIpv6AddressColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setVirtual(row assetRow, value string) {
	row[mapping.VirtualColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setPublic(row assetRow, value string) {
	row[mapping.PublicColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setDnsNameOrUrl(row assetRow, value string) {
	row[mapping.DnsNameOrUrlColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setNetBiosName(row assetRow, value string) {
	row[mapping.NetBiosNameColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setMacAddress(row assetRow, value string) {
	row[mapping.MacAddressColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setAuthenticatedScan(row assetRow, value string) {
	row[mapping.AuthenticatedScanColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setBaselineConfigurationName(row assetRow, value string) {
	row[mapping.BaselineConfigurationNameColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setOsNameAndVersion(row assetRow, value string) {
	row[mapping.OsNameAndVersionColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setLocation(row assetRow, value string) {
	row[mapping.LocationColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setAssetType(row assetRow, value string) {
	row[mapping.AssetTypeColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setHardwareMakeModel(row assetRow, value string) {
	row[mapping.HardwareMakeModelColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setInLatestScan(row assetRow, value string) {
	row[mapping.InLatestScanColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setSoftwareDatabaseVendor(row assetRow, value string) {
	row[mapping.SoftwareDatabaseVendorColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setSoftwareDatabaseNameVersion(row assetRow, value string) {
	row[mapping.SoftwareDatabaseNameVersionColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setPatchLevel(row assetRow, value string) {
	row[mapping.PatchLevelColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setDiagramLabel(row assetRow, value string) {
	row[mapping.DiagramLabelColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setComments(row assetRow, value string) {
	row[mapping.CommentsColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setSerialAssetTag(row assetRow, value string) {
	row[mapping.SerialAssetTagColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setVlanNetworkID(row assetRow, value string) {
	row[mapping.VlanNetworkIDColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setSystemAdministratorOwner(row assetRow, value string) {
	row[mapping.SystemAdministratorOwnerColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setApplicationAdministratorOwner(row assetRow, value string) {
	row[mapping.ApplicationAdministratorOwnerColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setFunction(row assetRow, value string) {
	row[mapping.FunctionColumnIndex] = value
}
func (mapping *spreadsheetColumnMapping) setEndOfLifeDate(row assetRow, value string) {
	row[mapping.EndOfLifeDateColumnIndex] = value
}
