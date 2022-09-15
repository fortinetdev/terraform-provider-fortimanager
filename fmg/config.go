package fortimanager

import (
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func validateConvIPMask2CIDR(oNewIP, oOldIP string) string {
	if oNewIP != oOldIP && strings.Contains(oNewIP, "/") && strings.Contains(oOldIP, " ") {
		line := strings.Split(oOldIP, " ")
		if len(line) >= 2 {
			ip := line[0]
			mask := line[1]
			prefixSize, _ := net.IPMask(net.ParseIP(mask).To4()).Size()
			return ip + "/" + strconv.Itoa(prefixSize)
		}
	}
	return oOldIP
}

func fortiStringValue(t interface{}) string {
	if v, ok := t.(string); ok {
		return v
	} else {
		return ""
	}
}

func fortiIntValue(t interface{}) int {
	if v, ok := t.(float64); ok {
		return int(v)
	} else {
		return 0
	}
}

func expandStringList(c []interface{}) []string {
	vs := make([]string, 0, len(c))
	for _, v := range c {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, v.(string))
		}
	}
	return vs
}

func flattenStringList(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	if v1, ok := v.(string); ok {
		result := make([]string, 0, 1)
		result = append(result, v1)
		return result
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]string, 0, len(l))

	for _, r := range l {
		result = append(result, fortiStringValue(r))
	}
	return result
}

func expandIntegerList(c []interface{}) []int {
	vs := make([]int, 0, len(c))
	for _, v := range c {
		_, ok := v.(int)
		if ok {
			vs = append(vs, v.(int))
		}
	}
	return vs
}

func flattenIntegerList(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]int, 0, len(l))

	for _, r := range l {
		result = append(result, fortiIntValue(r))
	}
	return result
}

func getEnumVal(v interface{}, emap map[int]string) string {
	if v != nil {
		if v1, ok := v.(float64); ok {
			v2 := int(v1)
			if v3, ok := emap[v2]; ok {
				return v3
			}
		} else if v1, ok := v.(string); ok {
			for _, v3 := range emap {
				if v3 == v1 {
					return v3
				}
			}
		}
	}

	return "ErrorValue"
}

func getEnumValbyBit(v interface{}, emap map[int]string) interface{} {
	if v != nil {
		if v1, ok := v.(float64); ok {
			vs := make([]string, 0, 0)
			v2 := uint32(v1)
			for i := 0; i < 32; i++ {
				bit := uint32(math.Pow(2, float64(i)))
				if v2&bit != 0 {
					ind := int(bit)
					if _, ok := emap[ind]; ok {
						vs = append(vs, emap[ind])
					}
				}
			}
			return vs
		}
	}

	return v
}

func getStringKey(d *schema.ResourceData, field string) string {
	if v, ok := d.GetOkExists(field); ok {
		if v1, ok := v.(string); ok {
			return v1
		}
	}

	return ""
}

func getIntKey(d *schema.ResourceData, field string) int {
	if v, ok := d.GetOkExists(field); ok {
		if v1, ok := v.(int); ok {
			return v1
		}
	}

	return 0
}

func escapeFilter(filter string) string {
	filter = strings.ReplaceAll(filter, "_", "-")
	filter = strings.ReplaceAll(filter, "fosid", "id")
	filter = strings.ReplaceAll(filter, "&", "&filter=")
	filter = strings.ReplaceAll(filter, ".", "\\.")
	filter = strings.ReplaceAll(filter, "\\", "\\\\")
	filter = "filter=" + filter
	return filter
}

func adomChecking(c *Config, d *schema.ResourceData) (string, error) {
	cst := c.ScopeType
	cadom := c.Adom
	st := d.Get("scopetype").(string)
	adom := d.Get("adom").(string)

	if st == "inherit" || st == "" {
		if cst == "global" {
			return "global", nil
		}
		return "adom/" + cadom, nil
	} else if st == "global" {
		return "global", nil
	} else if st == "adom" {
		if adom == "" {
			err := fmt.Errorf("Empty adom")
			return "", err
		}
		return "adom/" + adom, nil
	}

	err := fmt.Errorf("unknown adom configuration error")
	return "", err
}

func importOptionChecking(c *Config, para string) string {
	v := c.ImportOptions.List()
	if len(v) == 0 {
		return ""
	}

	for _, v1 := range v {
		if v2, ok := v1.(string); ok {
			v3 := strings.Split(v2, "=")

			if len(v3) == 2 { // Example "pkg=default"
				if v3[0] == para {
					return v3[1]
				}
			}
		}
	}

	return ""
}

func dynamic_sort_subtable(result []map[string]interface{}, fieldname string, d *schema.ResourceData) {
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		if v.(string) == "true" {
			sort.Slice(result, func(i, j int) bool {
				v1 := fmt.Sprintf("%v", result[i][fieldname])
				v2 := fmt.Sprintf("%v", result[j][fieldname])

				return v1 < v2
			})
		}
	}
}

func fortiAPIInt2Bool(t interface{}) string {
	if v, ok := t.(string); ok {
		return v
	} else if v, ok := t.(float64); ok {
		v1 := int(v)
		if v1 == 0 {
			return "false"
		} else {
			return "true"
		}
	}

	return "false"
}

func convertV2Ipnetmaskstring(v interface{}) string {
	res := ""

	if v == nil {
		return res
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return res
	}

	for _, r := range l {
		if res != "" {
			res += " "
		}
		res += fortiStringValue(r)
	}

	return res
}

func getAdom(d *schema.ResourceData) string {
	if v, ok := d.GetOk("adom"); ok {
		if v1, ok := v.(string); ok {
			return v1
		}
	} else {
		// if global ..... adom exist?
		// ...
		return "global"
	}

	return "global"
}

func fortiAPISubPartPatch(t interface{}, lgname string) interface{} {
	if t == nil {
		return t
	} else if v, ok := t.([]interface{}); ok {
		if len(v) == 0 {
			log.Printf("shengh6671 - %v", lgname)
			return ""
		}
	}

	return t
}

func fortiAPIPatch(t interface{}, lgname string) (interface{}, bool) {
	if t == nil {
		return nil, false
		// } else if _, ok := t.(string); ok {
		// 	return true
		// } else if _, ok := t.(float64); ok {
		// 	return true
	} else if v, ok := t.([]interface{}); ok {
		if len(v) == 0 {
			log.Printf("shengh6670 - %v", lgname)
			return nil, true
		} else if len(v) == 1 {
			if vv, ok := v[0].(string); ok {
				return vv, true
			}
			// else if vv, ok := v[0].(float64); ok {
			// 	return int(vv), true
			// }
		}
	}

	return nil, false
}

func isImportTable() bool {
	itable := os.Getenv("FORTIMANAGER_IMPORT_TABLE")
	if itable == "true" {
		return true
	}
	return false
}

func convintflist2i(v interface{}) interface{} {
	if t, ok := v.([]interface{}); ok {
		if len(t) == 0 {
			return v
		}
		return t[0]
	}
	return v
}

func convipstringlist2ipmask(v interface{}) interface{} {
	if t, ok := v.([]interface{}); ok {
		if len(t) == 0 {
			return v
		} else if len(t) == 2 {
			v2 := fmt.Sprintf("%v %v", t[0], t[1])
			return v2
		}
	}
	return v
}

func convttlfloat642String(v interface{}) interface{} {
	if v1, ok := v.(float64); ok {
		return strconv.Itoa(int(v1))
	}

	return v
}

func conv2str(v interface{}) interface{} {
	if _, ok := v.(string); ok {
		return v
	} else if _, ok := v.([]interface{}); ok {
		return convintflist2str(v)
	} else if v1, ok := v.(float64); ok {
		return strconv.FormatFloat(v1, 'f', -1, 64)
	} else if v1, ok := v.(int); ok {
		return strconv.FormatInt(int64(v1), 10)
	}
	return ""
}

func convintflist2str(v interface{}) interface{} {
	res := ""
	if t, ok := v.([]interface{}); ok {
		if len(t) == 0 {
			return res
		}

		bFirst := true
		for _, v1 := range t {
			curVal := ""
			if t1, ok := v1.(float64); ok {
				curVal = strconv.FormatFloat(t1, 'f', -1, 64)
			} else if t1, ok := v1.(int); ok {
				curVal = strconv.FormatInt(int64(t1), 10)
			} else if t1, ok := v1.(string); ok {
				curVal = t1
			} else {
				continue
			}

			if bFirst == true {
				res += curVal
				bFirst = false
			} else {
				res += " "
				res += curVal
			}
		}
	}
	return res
}

func convstrlist2str(v interface{}) interface{} {
	res := ""
	if _, ok := v.(string); ok {
		return v
	}
	if t, ok := v.([]interface{}); ok {
		if len(t) == 0 {
			return res
		}

		bFirst := true
		for _, v1 := range t {
			if t1, ok := v1.(string); ok {
				if bFirst == true {
					res += t1
					bFirst = false
				} else {
					res += " "
					res += t1
				}
			}
		}
	}
	return res
}

func intBetweenWithZero(min, max int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if (v >= min && v <= max) || (v == 0) {
			return warnings, errors
		}

		errors = append(errors, fmt.Errorf("expected %s to be in the range (%d - %d) or equal to 0, got %d", k, min, max, v))

		return warnings, errors
	}
}

func checkVersionMatch(curVersion string, versionRequirement map[string][]string) (match bool, err error) {
	operations := versionRequirement["operations"]
	versions := versionRequirement["versions"]
	for _, curRequiredVersion := range versions {
		curCompare := compareVersion(curVersion, curRequiredVersion)
		for _, curOperation := range operations {
			if (curOperation == ">") && (curCompare > 0) {
				return true, nil
			}
			if (curOperation == "=") && (curCompare == 0) {
				return true, nil
			}
			if (curOperation == "<") && (curCompare < 0) {
				return true, nil
			}
		}
	}
	err = fmt.Errorf("requires FortiManager version: %s, your device version is: %s.", strings.Join(operations, "")+" "+strings.Join(versions, ","), curVersion)
	return false, err
}

func compareVersion(curVersion string, requiredVersion string) int {
	curVersionList := strings.Split(curVersion, ".")
	requiredVersionList := strings.Split(requiredVersion, ".")
	minLen := len(curVersionList)
	if len(curVersionList) > len(requiredVersionList) {
		minLen = len(requiredVersionList)
	}
	for i := 0; i < minLen; i++ {
		curVersionValue, _ := strconv.Atoi(curVersionList[i])
		requiredVersionValue, _ := strconv.Atoi(requiredVersionList[i])
		if curVersionValue > requiredVersionValue {
			return 1
		}
		if curVersionValue < requiredVersionValue {
			return -1
		}
	}
	return 0
}
