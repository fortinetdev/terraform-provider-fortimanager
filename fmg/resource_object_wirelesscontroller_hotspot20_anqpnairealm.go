// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure network access identifier (NAI) realm.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerHotspot20AnqpNaiRealm() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20AnqpNaiRealmCreate,
		Read:   resourceObjectWirelessControllerHotspot20AnqpNaiRealmRead,
		Update: resourceObjectWirelessControllerHotspot20AnqpNaiRealmUpdate,
		Delete: resourceObjectWirelessControllerHotspot20AnqpNaiRealmDelete,

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
			"nai_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"eap_method": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_param": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"index": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"val": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"index": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"method": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"encoding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nai_realm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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

func resourceObjectWirelessControllerHotspot20AnqpNaiRealmCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpNaiRealm(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpNaiRealm resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerHotspot20AnqpNaiRealm(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpNaiRealm resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20AnqpNaiRealmRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpNaiRealmUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpNaiRealm(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpNaiRealm resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerHotspot20AnqpNaiRealm(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpNaiRealm resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20AnqpNaiRealmRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpNaiRealmDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerHotspot20AnqpNaiRealm(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20AnqpNaiRealm resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20AnqpNaiRealmRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerHotspot20AnqpNaiRealm(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpNaiRealm resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20AnqpNaiRealm(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpNaiRealm resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eap_method"
		if _, ok := i["eap-method"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(i["eap-method"], d, pre_append)
			tmp["eap_method"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealm-NaiList-EapMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encoding"
		if _, ok := i["encoding"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding(i["encoding"], d, pre_append)
			tmp["encoding"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealm-NaiList-Encoding")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nai_realm"
		if _, ok := i["nai-realm"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm(i["nai-realm"], d, pre_append)
			tmp["nai_realm"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealm-NaiList-NaiRealm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealm-NaiList-Name")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_param"
		if _, ok := i["auth-param"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(i["auth-param"], d, pre_append)
			tmp["auth_param"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod-AuthParam")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := i["method"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod(i["method"], d, pre_append)
			tmp["method"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod-Method")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "val"
		if _, ok := i["val"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal(i["val"], d, pre_append)
			tmp["val"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam-Val")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20AnqpNaiRealm(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("nai_list", flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(o["nai-list"], d, "nai_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["nai-list"], "ObjectWirelessControllerHotspot20AnqpNaiRealm-NaiList"); ok {
				if err = d.Set("nai_list", vv); err != nil {
					return fmt.Errorf("Error reading nai_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nai_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nai_list"); ok {
			if err = d.Set("nai_list", flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(o["nai-list"], d, "nai_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["nai-list"], "ObjectWirelessControllerHotspot20AnqpNaiRealm-NaiList"); ok {
					if err = d.Set("nai_list", vv); err != nil {
						return fmt.Errorf("Error reading nai_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nai_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerHotspot20AnqpNaiRealmName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerHotspot20AnqpNaiRealm-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eap_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["eap-method"], _ = expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(d, i["eap_method"], pre_append)
		} else {
			tmp["eap-method"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encoding"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["encoding"], _ = expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding(d, i["encoding"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nai_realm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nai-realm"], _ = expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm(d, i["nai_realm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_param"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-param"], _ = expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d, i["auth_param"], pre_append)
		} else {
			tmp["auth-param"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["method"], _ = expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod(d, i["method"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "val"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["val"], _ = expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal(d, i["val"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20AnqpNaiRealm(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("nai_list"); ok || d.HasChange("nai_list") {
		t, err := expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(d, v, "nai_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nai-list"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerHotspot20AnqpNaiRealmName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
