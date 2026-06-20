package ai

import "strings"

var validMitreIDs = map[string]bool{
	"T1071": true, "T1071.001": true, "T1071.002": true, "T1071.003": true,
	"T1071.004": true, "T1005": true, "T1041": true, "T1048": true,
	"T1048.002": true, "T1048.003": true, "T1090": true, "T1090.001": true,
	"T1090.002": true, "T1090.003": true, "T1095": true, "T1102": true,
	"T1102.002": true, "T1105": true, "T1110": true, "T1114": true,
	"T1132": true, "T1132.001": true, "T1132.002": true, "T1190": true,
	"T1197": true, "T1203": true, "T1204": true, "T1204.002": true,
	"T1219": true, "T1485": true, "T1490": true, "T1497": true,
	"T1497.001": true, "T1497.002": true, "T1505": true, "T1505.003": true,
	"T1525": true, "T1526": true, "T1537": true, "T1543": true,
	"T1546": true, "T1547": true, "T1550": true, "T1550.002": true,
	"T1552": true, "T1553": true, "T1555": true, "T1560": true,
	"T1560.001": true, "T1560.002": true, "T1560.003": true, "T1562": true,
	"T1562.001": true, "T1566": true, "T1568": true, "T1568.002": true,
	"T1571": true, "T1572": true, "T1573": true, "T1573.001": true,
	"T1574": true, "T1580": true, "T1583": true, "T1584": true,
	"T1585": true, "T1586": true, "T1587": true, "T1588": true,
	"T1589": true, "T1590": true, "T1591": true, "T1592": true,
	"T1595": true, "T1598": true, "T1608": true, "T1609": true,
	"T1610": true, "T1611": true, "T1612": true, "T1613": true,
	"T1614": true, "T1614.001": true, "T1620": true, "T1622": true,
	"T1625": true, "T1647": true, "T1648": true, "T1649": true,
	"T1650": true, "T1651": true, "T1652": true, "T1653": true,
	"T1654": true, "T1655": true, "T1656": true, "T1657": true,
	"T1658": true, "T1659": true,
}

var mitrePrefix = map[string]bool{}

func init() {
	for id := range validMitreIDs {
		parts := strings.Split(id, ".")
		if len(parts) > 0 {
			mitrePrefix[parts[0]] = true
		}
	}
}

func SanitizeMitreIDs(ids []string) []string {
	if len(ids) == 0 {
		return nil
	}
	var valid []string
	for _, id := range ids {
		id = strings.TrimSpace(strings.ToUpper(id))
		if validMitreIDs[id] || mitrePrefix[id] {
			valid = append(valid, id)
		}
	}
	return valid
}
