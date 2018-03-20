// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nhspd.proto

/*
Package ods is a generated protocol buffer package.

It is generated from these files:
	nhspd.proto

It has these top-level messages:
	NHSPD
*/
package ods

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Postcode user type
type NHSPD_UserType int32

const (
	NHSPD_SMALL NHSPD_UserType = 0
	NHSPD_LARGE NHSPD_UserType = 1
)

var NHSPD_UserType_name = map[int32]string{
	0: "SMALL",
	1: "LARGE",
}
var NHSPD_UserType_value = map[string]int32{
	"SMALL": 0,
	"LARGE": 1,
}

func (x NHSPD_UserType) String() string {
	return proto.EnumName(NHSPD_UserType_name, int32(x))
}
func (NHSPD_UserType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Grid reference positional quality indicator
type NHSPD_GridReferenceQuality int32

const (
	NHSPD_UNKNOWN_GRID_REF_QUALITY NHSPD_GridReferenceQuality = 0
	NHSPD_WITHIN_BUILDING          NHSPD_GridReferenceQuality = 1
	NHSPD_WITHIN_BUILDING_FROM_MAP NHSPD_GridReferenceQuality = 2
	NHSPD_WITHIN_50_METRES         NHSPD_GridReferenceQuality = 3
	NHSPD_POSTCODE_UNIT_MEAN       NHSPD_GridReferenceQuality = 4
	NHSPD_IMPUTED_BY_ONS           NHSPD_GridReferenceQuality = 5
	NHSPD_POSTCODE_SECTOR_MEAN     NHSPD_GridReferenceQuality = 6
	NHSPD_TERMINATED_LAST_KNOWN    NHSPD_GridReferenceQuality = 8
	NHSPD_NO_GRID_REFERENCE        NHSPD_GridReferenceQuality = 9
)

var NHSPD_GridReferenceQuality_name = map[int32]string{
	0: "UNKNOWN_GRID_REF_QUALITY",
	1: "WITHIN_BUILDING",
	2: "WITHIN_BUILDING_FROM_MAP",
	3: "WITHIN_50_METRES",
	4: "POSTCODE_UNIT_MEAN",
	5: "IMPUTED_BY_ONS",
	6: "POSTCODE_SECTOR_MEAN",
	8: "TERMINATED_LAST_KNOWN",
	9: "NO_GRID_REFERENCE",
}
var NHSPD_GridReferenceQuality_value = map[string]int32{
	"UNKNOWN_GRID_REF_QUALITY": 0,
	"WITHIN_BUILDING":          1,
	"WITHIN_BUILDING_FROM_MAP": 2,
	"WITHIN_50_METRES":         3,
	"POSTCODE_UNIT_MEAN":       4,
	"IMPUTED_BY_ONS":           5,
	"POSTCODE_SECTOR_MEAN":     6,
	"TERMINATED_LAST_KNOWN":    8,
	"NO_GRID_REFERENCE":        9,
}

func (x NHSPD_GridReferenceQuality) String() string {
	return proto.EnumName(NHSPD_GridReferenceQuality_name, int32(x))
}
func (NHSPD_GridReferenceQuality) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1}
}

// Census enumeration district positional quality indicator
type NHSPD_CensusCodeQuality int32

const (
	NHSPD_UNKNOWN_CENSUS_QUALITY            NHSPD_CensusCodeQuality = 0
	NHSPD_POINT_IN_POLYGON                  NHSPD_CensusCodeQuality = 1
	NHSPD_NORTHERN_IRELAND                  NHSPD_CensusCodeQuality = 6
	NHSPD_SCOTLAND_CHANNELISLANDS_ISLEOFMAN NHSPD_CensusCodeQuality = 9
)

var NHSPD_CensusCodeQuality_name = map[int32]string{
	0: "UNKNOWN_CENSUS_QUALITY",
	1: "POINT_IN_POLYGON",
	6: "NORTHERN_IRELAND",
	9: "SCOTLAND_CHANNELISLANDS_ISLEOFMAN",
}
var NHSPD_CensusCodeQuality_value = map[string]int32{
	"UNKNOWN_CENSUS_QUALITY":            0,
	"POINT_IN_POLYGON":                  1,
	"NORTHERN_IRELAND":                  6,
	"SCOTLAND_CHANNELISLANDS_ISLEOFMAN": 9,
}

func (x NHSPD_CensusCodeQuality) String() string {
	return proto.EnumName(NHSPD_CensusCodeQuality_name, int32(x))
}
func (NHSPD_CensusCodeQuality) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

// NHSPD is the NHS postcode directory
// See https://isd.digital.nhs.uk/trud3/user/authenticated/group/0/pack/5/subpack/242/releases
//
// TODO(wardle): check representation of OS grid references; switch from string to signed int64?
type NHSPD struct {
	Postcode                 string                     `protobuf:"bytes,1,opt,name=postcode" json:"postcode,omitempty"`
	PostcodeEgif             string                     `protobuf:"bytes,2,opt,name=postcode_egif,json=postcodeEgif" json:"postcode_egif,omitempty"`
	DateIntroduction         string                     `protobuf:"bytes,3,opt,name=date_introduction,json=dateIntroduction" json:"date_introduction,omitempty"`
	DateTermination          string                     `protobuf:"bytes,4,opt,name=date_termination,json=dateTermination" json:"date_termination,omitempty"`
	OsEasting_100M           string                     `protobuf:"bytes,5,opt,name=os_easting_100m,json=osEasting100m" json:"os_easting_100m,omitempty"`
	OsNorthing_100M          string                     `protobuf:"bytes,6,opt,name=os_northing_100m,json=osNorthing100m" json:"os_northing_100m,omitempty"`
	County                   string                     `protobuf:"bytes,7,opt,name=county" json:"county,omitempty"`
	LocalAuthorityOrg        string                     `protobuf:"bytes,8,opt,name=local_authority_org,json=localAuthorityOrg" json:"local_authority_org,omitempty"`
	LocalAuthorityDistrict   string                     `protobuf:"bytes,9,opt,name=local_authority_district,json=localAuthorityDistrict" json:"local_authority_district,omitempty"`
	ElectoralWard            string                     `protobuf:"bytes,10,opt,name=electoral_ward,json=electoralWard" json:"electoral_ward,omitempty"`
	UserType                 NHSPD_UserType             `protobuf:"varint,11,opt,name=user_type,json=userType,enum=ods.NHSPD_UserType" json:"user_type,omitempty"`
	GridRefereneQuality      NHSPD_GridReferenceQuality `protobuf:"varint,12,opt,name=grid_referene_quality,json=gridRefereneQuality,enum=ods.NHSPD_GridReferenceQuality" json:"grid_referene_quality,omitempty"`
	Country                  string                     `protobuf:"bytes,13,opt,name=country" json:"country,omitempty"`
	HealthAuthorityPre_2013  string                     `protobuf:"bytes,14,opt,name=health_authority_pre_2013,json=healthAuthorityPre2013" json:"health_authority_pre_2013,omitempty"`
	Region                   string                     `protobuf:"bytes,15,opt,name=region" json:"region,omitempty"`
	HealthAuthorityPre_2002  string                     `protobuf:"bytes,16,opt,name=health_authority_pre_2002,json=healthAuthorityPre2002" json:"health_authority_pre_2002,omitempty"`
	NhsEnglandRegion         string                     `protobuf:"bytes,17,opt,name=nhs_england_region,json=nhsEnglandRegion" json:"nhs_england_region,omitempty"`
	CcgLhbChpLcgPhd          string                     `protobuf:"bytes,18,opt,name=ccg_lhb_chp_lcg_phd,json=ccgLhbChpLcgPhd" json:"ccg_lhb_chp_lcg_phd,omitempty"`
	Census_1991Ogss          string                     `protobuf:"bytes,19,opt,name=census_1991_ogss,json=census1991Ogss" json:"census_1991_ogss,omitempty"`
	Census_1991Code          string                     `protobuf:"bytes,20,opt,name=census_1991_code,json=census1991Code" json:"census_1991_code,omitempty"`
	CensusCodeQuality        NHSPD_CensusCodeQuality    `protobuf:"varint,21,opt,name=census_code_quality,json=censusCodeQuality,enum=ods.NHSPD_CensusCodeQuality" json:"census_code_quality,omitempty"`
	Ward_1998                string                     `protobuf:"bytes,22,opt,name=ward_1998,json=ward1998" json:"ward_1998,omitempty"`
	Census_2001Area          string                     `protobuf:"bytes,23,opt,name=census_2001_area,json=census2001Area" json:"census_2001_area,omitempty"`
	NhsEnglandRlo            string                     `protobuf:"bytes,24,opt,name=nhs_england_rlo,json=nhsEnglandRlo" json:"nhs_england_rlo,omitempty"`
	FormerPanSha             string                     `protobuf:"bytes,25,opt,name=former_pan_sha,json=formerPanSha" json:"former_pan_sha,omitempty"`
	Census_2001Lsoa          string                     `protobuf:"bytes,26,opt,name=census_2001_lsoa,json=census2001Lsoa" json:"census_2001_lsoa,omitempty"`
	Census_2001Rural         string                     `protobuf:"bytes,27,opt,name=census_2001_rural,json=census2001Rural" json:"census_2001_rural,omitempty"`
	Census_2001Msoa          string                     `protobuf:"bytes,28,opt,name=census_2001_msoa,json=census2001Msoa" json:"census_2001_msoa,omitempty"`
	FormerCancerNetwork      string                     `protobuf:"bytes,29,opt,name=former_cancer_network,json=formerCancerNetwork" json:"former_cancer_network,omitempty"`
	StrategicClinicalNetwork string                     `protobuf:"bytes,30,opt,name=strategic_clinical_network,json=strategicClinicalNetwork" json:"strategic_clinical_network,omitempty"`
	HealthAuthorityPre_2006  string                     `protobuf:"bytes,31,opt,name=health_authority_pre_2006,json=healthAuthorityPre2006" json:"health_authority_pre_2006,omitempty"`
	PctPre_2006              string                     `protobuf:"bytes,32,opt,name=pct_pre_2006,json=pctPre2006" json:"pct_pre_2006,omitempty"`
	ItClusterPre_2007        string                     `protobuf:"bytes,33,opt,name=it_cluster_pre_2007,json=itClusterPre2007" json:"it_cluster_pre_2007,omitempty"`
	WestminsterConsituency   string                     `protobuf:"bytes,34,opt,name=westminster_consituency,json=westminsterConsituency" json:"westminster_consituency,omitempty"`
	CancerRegistry           string                     `protobuf:"bytes,35,opt,name=cancer_registry,json=cancerRegistry" json:"cancer_registry,omitempty"`
	Pct                      string                     `protobuf:"bytes,36,opt,name=pct" json:"pct,omitempty"`
	OsEasting_1M             string                     `protobuf:"bytes,37,opt,name=os_easting_1m,json=osEasting1m" json:"os_easting_1m,omitempty"`
	OsNorthing_1M            string                     `protobuf:"bytes,38,opt,name=os_northing_1m,json=osNorthing1m" json:"os_northing_1m,omitempty"`
	Census_2011Area          string                     `protobuf:"bytes,39,opt,name=census_2011_area,json=census2011Area" json:"census_2011_area,omitempty"`
	Census_2011Lsoa          string                     `protobuf:"bytes,40,opt,name=census_2011_lsoa,json=census2011Lsoa" json:"census_2011_lsoa,omitempty"`
	Census_2011Msoa          string                     `protobuf:"bytes,41,opt,name=census_2011_msoa,json=census2011Msoa" json:"census_2011_msoa,omitempty"`
}

func (m *NHSPD) Reset()                    { *m = NHSPD{} }
func (m *NHSPD) String() string            { return proto.CompactTextString(m) }
func (*NHSPD) ProtoMessage()               {}
func (*NHSPD) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NHSPD) GetPostcode() string {
	if m != nil {
		return m.Postcode
	}
	return ""
}

func (m *NHSPD) GetPostcodeEgif() string {
	if m != nil {
		return m.PostcodeEgif
	}
	return ""
}

func (m *NHSPD) GetDateIntroduction() string {
	if m != nil {
		return m.DateIntroduction
	}
	return ""
}

func (m *NHSPD) GetDateTermination() string {
	if m != nil {
		return m.DateTermination
	}
	return ""
}

func (m *NHSPD) GetOsEasting_100M() string {
	if m != nil {
		return m.OsEasting_100M
	}
	return ""
}

func (m *NHSPD) GetOsNorthing_100M() string {
	if m != nil {
		return m.OsNorthing_100M
	}
	return ""
}

func (m *NHSPD) GetCounty() string {
	if m != nil {
		return m.County
	}
	return ""
}

func (m *NHSPD) GetLocalAuthorityOrg() string {
	if m != nil {
		return m.LocalAuthorityOrg
	}
	return ""
}

func (m *NHSPD) GetLocalAuthorityDistrict() string {
	if m != nil {
		return m.LocalAuthorityDistrict
	}
	return ""
}

func (m *NHSPD) GetElectoralWard() string {
	if m != nil {
		return m.ElectoralWard
	}
	return ""
}

func (m *NHSPD) GetUserType() NHSPD_UserType {
	if m != nil {
		return m.UserType
	}
	return NHSPD_SMALL
}

func (m *NHSPD) GetGridRefereneQuality() NHSPD_GridReferenceQuality {
	if m != nil {
		return m.GridRefereneQuality
	}
	return NHSPD_UNKNOWN_GRID_REF_QUALITY
}

func (m *NHSPD) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *NHSPD) GetHealthAuthorityPre_2013() string {
	if m != nil {
		return m.HealthAuthorityPre_2013
	}
	return ""
}

func (m *NHSPD) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *NHSPD) GetHealthAuthorityPre_2002() string {
	if m != nil {
		return m.HealthAuthorityPre_2002
	}
	return ""
}

func (m *NHSPD) GetNhsEnglandRegion() string {
	if m != nil {
		return m.NhsEnglandRegion
	}
	return ""
}

func (m *NHSPD) GetCcgLhbChpLcgPhd() string {
	if m != nil {
		return m.CcgLhbChpLcgPhd
	}
	return ""
}

func (m *NHSPD) GetCensus_1991Ogss() string {
	if m != nil {
		return m.Census_1991Ogss
	}
	return ""
}

func (m *NHSPD) GetCensus_1991Code() string {
	if m != nil {
		return m.Census_1991Code
	}
	return ""
}

func (m *NHSPD) GetCensusCodeQuality() NHSPD_CensusCodeQuality {
	if m != nil {
		return m.CensusCodeQuality
	}
	return NHSPD_UNKNOWN_CENSUS_QUALITY
}

func (m *NHSPD) GetWard_1998() string {
	if m != nil {
		return m.Ward_1998
	}
	return ""
}

func (m *NHSPD) GetCensus_2001Area() string {
	if m != nil {
		return m.Census_2001Area
	}
	return ""
}

func (m *NHSPD) GetNhsEnglandRlo() string {
	if m != nil {
		return m.NhsEnglandRlo
	}
	return ""
}

func (m *NHSPD) GetFormerPanSha() string {
	if m != nil {
		return m.FormerPanSha
	}
	return ""
}

func (m *NHSPD) GetCensus_2001Lsoa() string {
	if m != nil {
		return m.Census_2001Lsoa
	}
	return ""
}

func (m *NHSPD) GetCensus_2001Rural() string {
	if m != nil {
		return m.Census_2001Rural
	}
	return ""
}

func (m *NHSPD) GetCensus_2001Msoa() string {
	if m != nil {
		return m.Census_2001Msoa
	}
	return ""
}

func (m *NHSPD) GetFormerCancerNetwork() string {
	if m != nil {
		return m.FormerCancerNetwork
	}
	return ""
}

func (m *NHSPD) GetStrategicClinicalNetwork() string {
	if m != nil {
		return m.StrategicClinicalNetwork
	}
	return ""
}

func (m *NHSPD) GetHealthAuthorityPre_2006() string {
	if m != nil {
		return m.HealthAuthorityPre_2006
	}
	return ""
}

func (m *NHSPD) GetPctPre_2006() string {
	if m != nil {
		return m.PctPre_2006
	}
	return ""
}

func (m *NHSPD) GetItClusterPre_2007() string {
	if m != nil {
		return m.ItClusterPre_2007
	}
	return ""
}

func (m *NHSPD) GetWestminsterConsituency() string {
	if m != nil {
		return m.WestminsterConsituency
	}
	return ""
}

func (m *NHSPD) GetCancerRegistry() string {
	if m != nil {
		return m.CancerRegistry
	}
	return ""
}

func (m *NHSPD) GetPct() string {
	if m != nil {
		return m.Pct
	}
	return ""
}

func (m *NHSPD) GetOsEasting_1M() string {
	if m != nil {
		return m.OsEasting_1M
	}
	return ""
}

func (m *NHSPD) GetOsNorthing_1M() string {
	if m != nil {
		return m.OsNorthing_1M
	}
	return ""
}

func (m *NHSPD) GetCensus_2011Area() string {
	if m != nil {
		return m.Census_2011Area
	}
	return ""
}

func (m *NHSPD) GetCensus_2011Lsoa() string {
	if m != nil {
		return m.Census_2011Lsoa
	}
	return ""
}

func (m *NHSPD) GetCensus_2011Msoa() string {
	if m != nil {
		return m.Census_2011Msoa
	}
	return ""
}

func init() {
	proto.RegisterType((*NHSPD)(nil), "ods.NHSPD")
	proto.RegisterEnum("ods.NHSPD_UserType", NHSPD_UserType_name, NHSPD_UserType_value)
	proto.RegisterEnum("ods.NHSPD_GridReferenceQuality", NHSPD_GridReferenceQuality_name, NHSPD_GridReferenceQuality_value)
	proto.RegisterEnum("ods.NHSPD_CensusCodeQuality", NHSPD_CensusCodeQuality_name, NHSPD_CensusCodeQuality_value)
}

func init() { proto.RegisterFile("nhspd.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x96, 0x5d, 0x6f, 0xdb, 0x36,
	0x17, 0xc7, 0xeb, 0xa6, 0x49, 0x13, 0xb6, 0x49, 0x64, 0x3a, 0x49, 0xd9, 0xb4, 0xcf, 0xd3, 0x34,
	0x7d, 0x4b, 0xb7, 0x2e, 0xb0, 0x53, 0xac, 0x2f, 0xc0, 0x6e, 0x5c, 0x5b, 0x4d, 0x84, 0xc9, 0x92,
	0x2b, 0xc9, 0x28, 0x7a, 0x45, 0xa8, 0x14, 0x23, 0x09, 0x93, 0x45, 0x8d, 0xa4, 0x51, 0xf8, 0x76,
	0x1f, 0x6a, 0xdf, 0x6c, 0xf7, 0x03, 0x49, 0xd9, 0x71, 0xb4, 0xad, 0x77, 0xe2, 0xff, 0xfc, 0x0e,
	0x45, 0x9d, 0xf3, 0xe7, 0x81, 0xc0, 0x9d, 0x32, 0x13, 0x55, 0x72, 0x5a, 0x71, 0x26, 0x19, 0x5c,
	0x63, 0x89, 0x38, 0xfe, 0xb3, 0x0d, 0xd6, 0xbd, 0x8b, 0x70, 0x3c, 0x84, 0x87, 0x60, 0xb3, 0x62,
	0x42, 0x12, 0x96, 0x50, 0xd4, 0x3a, 0x6a, 0x9d, 0x6c, 0x05, 0xcb, 0x35, 0x7c, 0x02, 0xb6, 0x17,
	0xcf, 0x98, 0xa6, 0xf9, 0x25, 0xba, 0xa9, 0x81, 0xbb, 0x0b, 0xd1, 0x4e, 0xf3, 0x4b, 0xf8, 0x23,
	0x68, 0x27, 0xb1, 0xa4, 0x38, 0x2f, 0x25, 0x67, 0xc9, 0x8c, 0xc8, 0x9c, 0x95, 0x68, 0x4d, 0x83,
	0x96, 0x0a, 0x38, 0x2b, 0x3a, 0x7c, 0x09, 0xb4, 0x86, 0x25, 0xe5, 0xd3, 0xbc, 0x8c, 0x35, 0x7b,
	0x4b, 0xb3, 0xbb, 0x4a, 0x8f, 0xae, 0x64, 0xf8, 0x1c, 0xec, 0x32, 0x81, 0x69, 0x2c, 0x64, 0x5e,
	0xa6, 0xb8, 0xd7, 0xed, 0x4e, 0xd1, 0xba, 0x26, 0xb7, 0x99, 0xb0, 0x8d, 0xaa, 0x44, 0x78, 0x02,
	0x2c, 0x26, 0x70, 0xc9, 0xb8, 0xcc, 0x96, 0xe0, 0x86, 0x06, 0x77, 0x98, 0xf0, 0x6a, 0x59, 0x93,
	0x07, 0x60, 0x83, 0xb0, 0x59, 0x29, 0xe7, 0xe8, 0xb6, 0x8e, 0xd7, 0x2b, 0x78, 0x0a, 0x3a, 0x05,
	0x23, 0x71, 0x81, 0xe3, 0x99, 0xcc, 0x18, 0xcf, 0xe5, 0x1c, 0x33, 0x9e, 0xa2, 0x4d, 0x0d, 0xb5,
	0x75, 0xa8, 0xbf, 0x88, 0xf8, 0x3c, 0x85, 0xef, 0x00, 0x6a, 0xf2, 0x49, 0x2e, 0x24, 0xcf, 0x89,
	0x44, 0x5b, 0x3a, 0xe9, 0xe0, 0x7a, 0xd2, 0xb0, 0x8e, 0xc2, 0x67, 0x60, 0x87, 0x16, 0x94, 0x48,
	0xc6, 0xe3, 0x02, 0x7f, 0x8b, 0x79, 0x82, 0x80, 0xf9, 0xa4, 0xa5, 0xfa, 0x39, 0xe6, 0x09, 0xec,
	0x82, 0xad, 0x99, 0xa0, 0x1c, 0xcb, 0x79, 0x45, 0xd1, 0x9d, 0xa3, 0xd6, 0xc9, 0xce, 0x59, 0xe7,
	0x94, 0x25, 0xe2, 0x54, 0xb7, 0xec, 0x74, 0x22, 0x28, 0x8f, 0xe6, 0x15, 0x0d, 0x36, 0x67, 0xf5,
	0x13, 0x0c, 0xc1, 0x7e, 0xca, 0xf3, 0x04, 0x73, 0x7a, 0x49, 0x39, 0x2d, 0x29, 0xfe, 0x7d, 0x16,
	0x17, 0xb9, 0x9c, 0xa3, 0xbb, 0x3a, 0xfb, 0xd1, 0x4a, 0xf6, 0x39, 0xcf, 0x93, 0xc0, 0x60, 0x84,
	0x7e, 0x32, 0x58, 0xd0, 0x49, 0xaf, 0xd4, 0x85, 0x08, 0x11, 0xb8, 0xad, 0x2b, 0xc4, 0xe7, 0x68,
	0x5b, 0x1f, 0x73, 0xb1, 0x84, 0xef, 0xc1, 0xfd, 0x8c, 0xc6, 0x85, 0xcc, 0x56, 0x4a, 0x50, 0x71,
	0x8a, 0xcf, 0xba, 0xbd, 0xd7, 0x68, 0xc7, 0x94, 0xc0, 0x00, 0xcb, 0x1a, 0x8c, 0x39, 0x55, 0x51,
	0xd5, 0x04, 0x4e, 0x53, 0xd5, 0xf7, 0x5d, 0xd3, 0x04, 0xb3, 0xfa, 0xce, 0x96, 0xdd, 0x33, 0x64,
	0xfd, 0xf7, 0x96, 0xdd, 0x33, 0xf8, 0x0a, 0xc0, 0x32, 0x13, 0x98, 0x96, 0x69, 0x11, 0x97, 0xaa,
	0x06, 0x7a, 0xfb, 0xb6, 0xb1, 0x60, 0x99, 0x09, 0xdb, 0x04, 0x02, 0xf3, 0xa2, 0x57, 0xa0, 0x43,
	0x48, 0x8a, 0x8b, 0xec, 0x2b, 0x26, 0x59, 0x85, 0x0b, 0x92, 0xe2, 0x2a, 0x4b, 0x10, 0x34, 0x2e,
	0x24, 0x24, 0x75, 0xb3, 0xaf, 0x83, 0xac, 0x72, 0x49, 0x3a, 0xce, 0x12, 0xe5, 0x2e, 0x42, 0x4b,
	0x31, 0x13, 0xb8, 0xf7, 0xfe, 0x7d, 0x0f, 0xb3, 0x54, 0x08, 0xd4, 0x31, 0xee, 0x32, 0xba, 0x92,
	0xfd, 0x54, 0x88, 0x26, 0xa9, 0x2f, 0xd4, 0x5e, 0x93, 0x1c, 0xa8, 0x6b, 0xe5, 0x82, 0x4e, 0x4d,
	0xea, 0x9b, 0xb5, 0x68, 0xd5, 0xbe, 0x6e, 0xd5, 0xc3, 0x95, 0x56, 0x0d, 0x34, 0xa5, 0x72, 0x16,
	0x7d, 0x6a, 0x93, 0xa6, 0x04, 0x1f, 0x80, 0x2d, 0xe5, 0x24, 0xf5, 0xd6, 0x77, 0xe8, 0xc0, 0xdc,
	0x60, 0x25, 0xa8, 0xf5, 0xca, 0xa1, 0xce, 0xba, 0xdd, 0x1e, 0x8e, 0x39, 0x8d, 0xd1, 0xbd, 0xd5,
	0x43, 0x29, 0xb9, 0xcf, 0x69, 0xac, 0xae, 0xdb, 0xb5, 0x22, 0x16, 0x0c, 0x21, 0xe3, 0xcd, 0x95,
	0x0a, 0x16, 0x0c, 0x3e, 0x05, 0x3b, 0x97, 0x8c, 0x4f, 0x29, 0xc7, 0x55, 0x5c, 0x62, 0x91, 0xc5,
	0xe8, 0xbe, 0x19, 0x0a, 0x46, 0x1d, 0xc7, 0x65, 0x98, 0xc5, 0xcd, 0xf7, 0x16, 0x82, 0xc5, 0xe8,
	0xb0, 0xf9, 0x5e, 0x57, 0xb0, 0x18, 0xfe, 0x00, 0xda, 0xab, 0x24, 0x9f, 0xf1, 0xb8, 0x40, 0x0f,
	0xea, 0x66, 0x2c, 0xd1, 0x40, 0xc9, 0xcd, 0x5d, 0xa7, 0x6a, 0xd7, 0x87, 0xcd, 0x5d, 0x47, 0x6a,
	0xd7, 0x33, 0xb0, 0x5f, 0x9f, 0x92, 0xc4, 0x25, 0xa1, 0x1c, 0x97, 0x54, 0x7e, 0x63, 0xfc, 0x37,
	0xf4, 0x3f, 0x8d, 0x77, 0x4c, 0x70, 0xa0, 0x63, 0x9e, 0x09, 0xc1, 0x5f, 0xc0, 0xa1, 0x90, 0x3c,
	0x96, 0x34, 0xcd, 0x09, 0x26, 0x45, 0x5e, 0xe6, 0xea, 0x8e, 0x2f, 0x12, 0xff, 0xaf, 0x13, 0xd1,
	0x92, 0x18, 0xd4, 0xc0, 0x22, 0xfb, 0x3b, 0xfe, 0x7d, 0x83, 0x1e, 0x7d, 0xc7, 0xbf, 0x6f, 0xe0,
	0x11, 0xb8, 0x5b, 0x11, 0x79, 0x45, 0x1f, 0x69, 0x1a, 0x54, 0x44, 0x2e, 0x88, 0x9f, 0x40, 0x27,
	0x97, 0x98, 0x14, 0x33, 0x21, 0x55, 0xe1, 0x0d, 0xf8, 0x16, 0x3d, 0x36, 0x16, 0xcf, 0xe5, 0xc0,
	0x44, 0x0c, 0xfe, 0x16, 0xbe, 0x05, 0xf7, 0xbe, 0x51, 0x21, 0xa7, 0x79, 0xa9, 0x79, 0xc2, 0x4a,
	0x91, 0xcb, 0x19, 0x2d, 0xc9, 0x1c, 0x1d, 0x9b, 0x93, 0xac, 0x84, 0x07, 0x57, 0x51, 0xf8, 0x02,
	0xec, 0xd6, 0xf5, 0x52, 0x97, 0x48, 0xa8, 0x9b, 0xff, 0xa4, 0xae, 0xaf, 0x96, 0x83, 0x5a, 0x85,
	0x16, 0x58, 0xab, 0x88, 0x44, 0x4f, 0x75, 0x50, 0x3d, 0xc2, 0x63, 0xb0, 0xbd, 0x3a, 0xae, 0xa7,
	0xe8, 0x99, 0x8e, 0xdd, 0xb9, 0x1a, 0xd6, 0x53, 0xe5, 0x9d, 0x6b, 0xa3, 0x7a, 0x8a, 0x9e, 0x1b,
	0xef, 0xac, 0x0c, 0xea, 0xe9, 0xb5, 0x2e, 0xf7, 0x6a, 0xcf, 0xbe, 0xb8, 0xde, 0xe5, 0x9e, 0xf1,
	0x6c, 0x83, 0xd4, 0x2e, 0x3b, 0x69, 0x92, 0xda, 0x65, 0x0d, 0x52, 0x3b, 0xe7, 0x65, 0x93, 0x54,
	0xce, 0x39, 0x3e, 0x02, 0x9b, 0x8b, 0xf9, 0x0a, 0xb7, 0xc0, 0x7a, 0x38, 0xea, 0xbb, 0xae, 0x75,
	0x43, 0x3d, 0xba, 0xfd, 0xe0, 0xdc, 0xb6, 0x5a, 0xc7, 0x7f, 0xb5, 0xc0, 0xde, 0xbf, 0x0d, 0x51,
	0xf8, 0x10, 0xa0, 0x89, 0xf7, 0xab, 0xe7, 0x7f, 0xf6, 0xf0, 0x79, 0xe0, 0x0c, 0x71, 0x60, 0x7f,
	0xc4, 0x9f, 0x26, 0x7d, 0xd7, 0x89, 0xbe, 0x58, 0x37, 0x60, 0x07, 0xec, 0x7e, 0x76, 0xa2, 0x0b,
	0xc7, 0xc3, 0x1f, 0x26, 0x8e, 0x3b, 0x74, 0xbc, 0x73, 0xab, 0xa5, 0x52, 0x1a, 0x22, 0xfe, 0x18,
	0xf8, 0x23, 0x3c, 0xea, 0x8f, 0xad, 0x9b, 0x70, 0x0f, 0x58, 0x75, 0xf4, 0xe7, 0x2e, 0x1e, 0xd9,
	0x51, 0x60, 0x87, 0xd6, 0x1a, 0x3c, 0x00, 0x70, 0xec, 0x87, 0xd1, 0xc0, 0x1f, 0xda, 0x78, 0xe2,
	0x39, 0x11, 0x1e, 0xd9, 0x7d, 0xcf, 0xba, 0x05, 0x21, 0xd8, 0x71, 0x46, 0xe3, 0x49, 0x64, 0x0f,
	0xf1, 0x87, 0x2f, 0xd8, 0xf7, 0x42, 0x6b, 0x1d, 0x22, 0xb0, 0xb7, 0x64, 0x43, 0x7b, 0x10, 0xf9,
	0x81, 0xa1, 0x37, 0xe0, 0x7d, 0xb0, 0x1f, 0xd9, 0xc1, 0xc8, 0xf1, 0xfa, 0x2a, 0xc1, 0xed, 0x87,
	0x11, 0xd6, 0x47, 0xb7, 0x36, 0xe1, 0x3e, 0x68, 0x7b, 0xfe, 0xf2, 0x13, 0xec, 0xc0, 0xf6, 0x06,
	0xb6, 0xb5, 0x75, 0xfc, 0x47, 0x0b, 0xb4, 0xff, 0x31, 0x91, 0xe0, 0x21, 0x38, 0x58, 0x7c, 0xf4,
	0xc0, 0xf6, 0xc2, 0x49, 0xb8, 0xf2, 0xc9, 0x7b, 0xc0, 0x1a, 0xfb, 0x8e, 0x17, 0x61, 0xc7, 0xc3,
	0x63, 0xdf, 0xfd, 0x72, 0xee, 0x7b, 0x56, 0x4b, 0xa9, 0x9e, 0x1f, 0x44, 0x17, 0x76, 0xe0, 0x61,
	0x27, 0xb0, 0xdd, 0xbe, 0x37, 0xb4, 0x36, 0xe0, 0x33, 0xf0, 0x38, 0x1c, 0xf8, 0x91, 0x5a, 0xe1,
	0xc1, 0x45, 0xdf, 0xf3, 0x6c, 0xd7, 0x09, 0xd5, 0x2a, 0xc4, 0x4e, 0xe8, 0xda, 0xfe, 0xc7, 0x51,
	0xdf, 0xb3, 0xb6, 0xbe, 0x6e, 0xe8, 0x9f, 0x98, 0xd7, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x7c,
	0x8c, 0xd8, 0x15, 0xd3, 0x08, 0x00, 0x00,
}
