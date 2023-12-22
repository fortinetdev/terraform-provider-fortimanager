// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: URL access list

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWafProfileUrlAccess() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWafProfileUrlAccessCreate,
		Read:   resourceObjectWafProfileUrlAccessRead,
		Update: resourceObjectWafProfileUrlAccessUpdate,
		Delete: resourceObjectWafProfileUrlAccessDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"scopetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "inherit",
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"adom",
					"global",
					"inherit",
				}, false),
			},
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"access_pattern": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"regex": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"srcaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectWafProfileUrlAccessCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectWafProfileUrlAccess(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWafProfileUrlAccess resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWafProfileUrlAccess(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWafProfileUrlAccess resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWafProfileUrlAccessRead(d, m)
}

func resourceObjectWafProfileUrlAccessUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectWafProfileUrlAccess(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfileUrlAccess resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWafProfileUrlAccess(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfileUrlAccess resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWafProfileUrlAccessRead(d, m)
}

func resourceObjectWafProfileUrlAccessDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	err = c.DeleteObjectWafProfileUrlAccess(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWafProfileUrlAccess resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWafProfileUrlAccessRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadObjectWafProfileUrlAccess(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfileUrlAccess resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWafProfileUrlAccess(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfileUrlAccess resource from API: %v", err)
	}
	return nil
}

func flattenObjectWafProfileUrlAccessAccessPattern2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWafProfileUrlAccessAccessPatternId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWafProfileUrlAccess-AccessPattern-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := i["negate"]; ok {
			v := flattenObjectWafProfileUrlAccessAccessPatternNegate2edl(i["negate"], d, pre_append)
			tmp["negate"] = fortiAPISubPartPatch(v, "ObjectWafProfileUrlAccess-AccessPattern-Negate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenObjectWafProfileUrlAccessAccessPatternPattern2edl(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "ObjectWafProfileUrlAccess-AccessPattern-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			v := flattenObjectWafProfileUrlAccessAccessPatternRegex2edl(i["regex"], d, pre_append)
			tmp["regex"] = fortiAPISubPartPatch(v, "ObjectWafProfileUrlAccess-AccessPattern-Regex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {
			v := flattenObjectWafProfileUrlAccessAccessPatternSrcaddr2edl(i["srcaddr"], d, pre_append)
			tmp["srcaddr"] = fortiAPISubPartPatch(v, "ObjectWafProfileUrlAccess-AccessPattern-Srcaddr")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectWafProfileUrlAccessAccessPatternId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAccessPatternNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAccessPatternPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAccessPatternRegex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAccessPatternSrcaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWafProfileUrlAccess(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("access_pattern", flattenObjectWafProfileUrlAccessAccessPattern2edl(o["access-pattern"], d, "access_pattern")); err != nil {
			if vv, ok := fortiAPIPatch(o["access-pattern"], "ObjectWafProfileUrlAccess-AccessPattern"); ok {
				if err = d.Set("access_pattern", vv); err != nil {
					return fmt.Errorf("Error reading access_pattern: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading access_pattern: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("access_pattern"); ok {
			if err = d.Set("access_pattern", flattenObjectWafProfileUrlAccessAccessPattern2edl(o["access-pattern"], d, "access_pattern")); err != nil {
				if vv, ok := fortiAPIPatch(o["access-pattern"], "ObjectWafProfileUrlAccess-AccessPattern"); ok {
					if err = d.Set("access_pattern", vv); err != nil {
						return fmt.Errorf("Error reading access_pattern: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading access_pattern: %v", err)
				}
			}
		}
	}

	if err = d.Set("action", flattenObjectWafProfileUrlAccessAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectWafProfileUrlAccess-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("address", flattenObjectWafProfileUrlAccessAddress2edl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "ObjectWafProfileUrlAccess-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectWafProfileUrlAccessId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWafProfileUrlAccess-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("log", flattenObjectWafProfileUrlAccessLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ObjectWafProfileUrlAccess-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("severity", flattenObjectWafProfileUrlAccessSeverity2edl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "ObjectWafProfileUrlAccess-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	return nil
}

func flattenObjectWafProfileUrlAccessFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWafProfileUrlAccessAccessPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectWafProfileUrlAccessAccessPatternId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["negate"], _ = expandObjectWafProfileUrlAccessAccessPatternNegate2edl(d, i["negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandObjectWafProfileUrlAccessAccessPatternPattern2edl(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["regex"], _ = expandObjectWafProfileUrlAccessAccessPatternRegex2edl(d, i["regex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr"], _ = expandObjectWafProfileUrlAccessAccessPatternSrcaddr2edl(d, i["srcaddr"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectWafProfileUrlAccessAccessPatternId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAccessPatternNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAccessPatternPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAccessPatternRegex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAccessPatternSrcaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWafProfileUrlAccess(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_pattern"); ok || d.HasChange("access_pattern") {
		t, err := expandObjectWafProfileUrlAccessAccessPattern2edl(d, v, "access_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-pattern"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectWafProfileUrlAccessAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandObjectWafProfileUrlAccessAddress2edl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWafProfileUrlAccessId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandObjectWafProfileUrlAccessLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandObjectWafProfileUrlAccessSeverity2edl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	return &obj, nil
}
