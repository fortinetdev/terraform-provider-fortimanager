// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ICAP response mode forward rules.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectIcapProfileRespmodForwardRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectIcapProfileRespmodForwardRulesCreate,
		Read:   resourceObjectIcapProfileRespmodForwardRulesRead,
		Update: resourceObjectIcapProfileRespmodForwardRulesUpdate,
		Delete: resourceObjectIcapProfileRespmodForwardRulesDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"case_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_resp_status_code": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectIcapProfileRespmodForwardRulesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectIcapProfileRespmodForwardRules(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIcapProfileRespmodForwardRules resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectIcapProfileRespmodForwardRules(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIcapProfileRespmodForwardRules resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectIcapProfileRespmodForwardRulesRead(d, m)
}

func resourceObjectIcapProfileRespmodForwardRulesUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectIcapProfileRespmodForwardRules(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIcapProfileRespmodForwardRules resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectIcapProfileRespmodForwardRules(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIcapProfileRespmodForwardRules resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectIcapProfileRespmodForwardRulesRead(d, m)
}

func resourceObjectIcapProfileRespmodForwardRulesDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectIcapProfileRespmodForwardRules(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectIcapProfileRespmodForwardRules resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectIcapProfileRespmodForwardRulesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectIcapProfileRespmodForwardRules(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIcapProfileRespmodForwardRules resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectIcapProfileRespmodForwardRules(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIcapProfileRespmodForwardRules resource from API: %v", err)
	}
	return nil
}

func flattenObjectIcapProfileRespmodForwardRulesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRespmodForwardRulesHeaderGroup2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := i["case-sensitivity"]; ok {
			v := flattenObjectIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity2edl(i["case-sensitivity"], d, pre_append)
			tmp["case_sensitivity"] = fortiAPISubPartPatch(v, "ObjectIcapProfileRespmodForwardRules-HeaderGroup-CaseSensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectIcapProfileRespmodForwardRulesHeaderGroupHeader2edl(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectIcapProfileRespmodForwardRules-HeaderGroup-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenObjectIcapProfileRespmodForwardRulesHeaderGroupHeaderName2edl(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "ObjectIcapProfileRespmodForwardRules-HeaderGroup-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectIcapProfileRespmodForwardRulesHeaderGroupId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectIcapProfileRespmodForwardRules-HeaderGroup-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRespmodForwardRulesHeaderGroupHeader2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRespmodForwardRulesHeaderGroupHeaderName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRespmodForwardRulesHeaderGroupId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRespmodForwardRulesHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectIcapProfileRespmodForwardRulesHttpRespStatusCode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectIcapProfileRespmodForwardRulesName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectIcapProfileRespmodForwardRules(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("action", flattenObjectIcapProfileRespmodForwardRulesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectIcapProfileRespmodForwardRules-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("header_group", flattenObjectIcapProfileRespmodForwardRulesHeaderGroup2edl(o["header-group"], d, "header_group")); err != nil {
			if vv, ok := fortiAPIPatch(o["header-group"], "ObjectIcapProfileRespmodForwardRules-HeaderGroup"); ok {
				if err = d.Set("header_group", vv); err != nil {
					return fmt.Errorf("Error reading header_group: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading header_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("header_group"); ok {
			if err = d.Set("header_group", flattenObjectIcapProfileRespmodForwardRulesHeaderGroup2edl(o["header-group"], d, "header_group")); err != nil {
				if vv, ok := fortiAPIPatch(o["header-group"], "ObjectIcapProfileRespmodForwardRules-HeaderGroup"); ok {
					if err = d.Set("header_group", vv); err != nil {
						return fmt.Errorf("Error reading header_group: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading header_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("host", flattenObjectIcapProfileRespmodForwardRulesHost2edl(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "ObjectIcapProfileRespmodForwardRules-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("http_resp_status_code", flattenObjectIcapProfileRespmodForwardRulesHttpRespStatusCode2edl(o["http-resp-status-code"], d, "http_resp_status_code")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-resp-status-code"], "ObjectIcapProfileRespmodForwardRules-HttpRespStatusCode"); ok {
			if err = d.Set("http_resp_status_code", vv); err != nil {
				return fmt.Errorf("Error reading http_resp_status_code: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_resp_status_code: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectIcapProfileRespmodForwardRulesName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectIcapProfileRespmodForwardRules-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectIcapProfileRespmodForwardRulesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectIcapProfileRespmodForwardRulesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRespmodForwardRulesHeaderGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitivity"], _ = expandObjectIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity2edl(d, i["case_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectIcapProfileRespmodForwardRulesHeaderGroupHeader2edl(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandObjectIcapProfileRespmodForwardRulesHeaderGroupHeaderName2edl(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectIcapProfileRespmodForwardRulesHeaderGroupId2edl(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRespmodForwardRulesHeaderGroupHeader2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRespmodForwardRulesHeaderGroupHeaderName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRespmodForwardRulesHeaderGroupId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRespmodForwardRulesHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectIcapProfileRespmodForwardRulesHttpRespStatusCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectIcapProfileRespmodForwardRulesName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectIcapProfileRespmodForwardRules(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectIcapProfileRespmodForwardRulesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("header_group"); ok || d.HasChange("header_group") {
		t, err := expandObjectIcapProfileRespmodForwardRulesHeaderGroup2edl(d, v, "header_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-group"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandObjectIcapProfileRespmodForwardRulesHost2edl(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("http_resp_status_code"); ok || d.HasChange("http_resp_status_code") {
		t, err := expandObjectIcapProfileRespmodForwardRulesHttpRespStatusCode2edl(d, v, "http_resp_status_code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-resp-status-code"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectIcapProfileRespmodForwardRulesName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
