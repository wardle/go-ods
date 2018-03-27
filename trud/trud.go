package trud

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"strconv"
	"strings"
)

// ODSFileType represents a unique file type within the ODS TRUD archive
type ODSFileType int

// Supported filetypes
const (
	ODSUnknownFileType ODSFileType = iota // Unknown filetype
	ODSGridAll                            // GridAll maps UK postcodes to different NHS organisational units
	ODSLastFileType                       // marker
)

var names = [...]string{
	"UNKNOWN",
	"GRIDALL",
	"LAST",
}
var archivePaths = [...]string{
	"",
	"/ONSdata/Data/gridall.zip/gridall.csv",
	"",
}

// property keys specify the primary key for an ODS file type
var propertyKeys = [...]string{
	"",
	"PostcodeEgif",
	"",
}

// GetODSFileType returns the filetype for the specified path
func GetODSFileType(path string) ODSFileType {
	for i, p := range archivePaths {
		if p != "" && strings.HasSuffix(path, p) {
			return ODSFileType(i)
		}
	}
	return ODSUnknownFileType
}

// GetName returns a name for this filetype
func (ft ODSFileType) GetName() string {
	return names[ft]
}

// GetArchivePath returns the file's path from the NHS TRUD ODS release
func (ft ODSFileType) GetArchivePath() string {
	return archivePaths[ft]
}

// GetPropertyKey returns the name of the property to be used as a primary key for this filetype
func (ft ODSFileType) GetPropertyKey() string {
	return propertyKeys[ft]
}

// Import provides a simple way to import a complete TRUD ODS distribution
func Import(path string, batchSize int, f func(filetype ODSFileType, result []proto.Message)) error {
	return Walk(path, func(wf *WalkedFile, err error) error {
		fmt.Printf("Processing path %s...", wf.Path)
		// fail if we have errors... but we could choose to log instead
		if err != nil {
			return err
		}
		ft := GetODSFileType(wf.Path)
		fmt.Printf("filetype: %d\n", ft)
		if ft != ODSUnknownFileType {
			r, err := wf.Open()
			if err != nil {
				return err
			}
			ProcessCSV(r, nil, batchSize, nil, func(rows [][]string) error {
				result := make([]proto.Message, 0, batchSize)
				for _, row := range rows {
					err = processRow(ft, row, &result)
				}
				f(ft, result)
				return nil
			})
		}
		return nil
	})
}

func processRow(ft ODSFileType, row []string, result *[]proto.Message) error {
	var r proto.Message
	var err error
	switch ft {
	case ODSGridAll:
		r, err = processNHSPD(row)

	}
	if err != nil {
		return err
	}
	*result = append(*result, r)
	return nil
}

//"PCD2","PCDS","DOINTR","DOTERM","OSEAST100M","OSNRTH100M","OSCTY","ODSLAUA","OSLAUA",
//"OSWARD","USERTYPE","OSGRDIND","CTRY","OSHLTHAU","GOR","OLDHA","NHSCR","CCG","PSED",
//"CENED","EDIND","WARD98","OA01","NHSRG","HRO","LSOA01","UR01IND","MSOA01","CANNET","SCN","OSHAPREV","OLDPCT","OLDHRO","PCON","CANREG","PCT","OSEAST1M","OSNRTH1M","OA11","LSOA11","MSOA11"
func processNHSPD(row []string) (*NHSPD, error) {
	r := new(NHSPD)
	errs := make([]error, 0)
	r.Postcode = row[0]
	r.PostcodeEgif = row[1]
	r.DateIntroduction = row[2]
	r.DateTermination = row[3]
	r.OsNorthing_100M = row[4]
	r.OsEasting_100M = row[5]
	r.County = row[6]
	r.LocalAuthorityOrg = row[7]
	r.LocalAuthorityDistrict = row[8]
	r.ElectoralWard = row[9]
	r.UserType = NHSPD_UserType(parseInt32(row[10], &errs))
	r.GridReferenceQuality = NHSPD_GridReferenceQuality(parseInt32(row[11], &errs))
	r.Country = row[12]
	r.HealthAuthorityPre_2013 = row[13]
	r.Region = row[14]
	r.HealthAuthorityPre_2002 = row[15]
	r.NhsEnglandRegion = row[16]
	r.CcgLhbChpLcgPhd = row[17]
	r.Census_1991Ogss = row[18]
	r.Census_1991Code = row[19]
	r.CensusCodeQuality = NHSPD_CensusCodeQuality(parseInt32(row[20], &errs))
	r.Ward_1998 = row[21]
	r.NhsEnglandRlo = row[22]
	r.FormerPanSha = row[23]
	r.Census_2001Lsoa = row[24]
	r.Census_2001Rural = row[25]
	r.Census_2001Msoa = row[26]
	r.FormerCancerNetwork = row[27]
	r.StrategicClinicalNetwork = row[28]
	r.HealthAuthorityPre_2006 = row[29]
	r.PctPre_2006 = row[30]
	r.ItClusterPre_2007 = row[31]
	r.WestminsterConsituency = row[32]
	r.CancerRegistry = row[33]
	r.Pct = row[34]
	r.OsNorthing_1M = row[35]
	r.Census_2011Area = row[36]
	r.Census_2011Lsoa = row[37]
	r.Census_2011Msoa = row[38]
	if len(errs) > 0 {
		return nil, errs[0] //TODO(wardle): concatenate errors
	}
	return r, nil
}

func parseInt32(s string, errs *[]error) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		*errs = append(*errs, err)
	}
	return int32(i)
}
