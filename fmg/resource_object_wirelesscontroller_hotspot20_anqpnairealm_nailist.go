// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: NAI list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListCreate,
		Read:   resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListRead,
		Update: resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListUpdate,
		Delete: resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListDelete,

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
			"anqp_nai_realm": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
										Computed: true,
									},
									"index": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"val": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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
							Computed: true,
						},
					},
				},
			},
			"encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nai_realm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	paradict["anqp_nai_realm"] = anqp_nai_realm

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListUpdate(d *schema.ResourceData, m interface{}) error {
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

	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	paradict["anqp_nai_realm"] = anqp_nai_realm

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListDelete(d *schema.ResourceData, m interface{}) error {
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

	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	paradict["anqp_nai_realm"] = anqp_nai_realm

	err = c.DeleteObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListRead(d *schema.ResourceData, m interface{}) error {
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

	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	if anqp_nai_realm == "" {
		anqp_nai_realm = importOptionChecking(m.(*FortiClient).Cfg, "anqp_nai_realm")
		if anqp_nai_realm == "" {
			return fmt.Errorf("Parameter anqp_nai_realm is missing")
		}
		if err = d.Set("anqp_nai_realm", anqp_nai_realm); err != nil {
			return fmt.Errorf("Error set params anqp_nai_realm: %v", err)
		}
	}
	paradict["anqp_nai_realm"] = anqp_nai_realm

	o, err := c.ReadObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam2edl(i["auth-param"], d, pre_append)
			tmp["auth_param"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod-AuthParam")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex2edl(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := i["method"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod2edl(i["method"], d, pre_append)
			tmp["method"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod-Method")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex2edl(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "val"
		if _, ok := i["val"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal2edl(i["val"], d, pre_append)
			tmp["val"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam-Val")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("eap_method", flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod2edl(o["eap-method"], d, "eap_method")); err != nil {
			if vv, ok := fortiAPIPatch(o["eap-method"], "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod"); ok {
				if err = d.Set("eap_method", vv); err != nil {
					return fmt.Errorf("Error reading eap_method: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading eap_method: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("eap_method"); ok {
			if err = d.Set("eap_method", flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod2edl(o["eap-method"], d, "eap_method")); err != nil {
				if vv, ok := fortiAPIPatch(o["eap-method"], "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod"); ok {
					if err = d.Set("eap_method", vv); err != nil {
						return fmt.Errorf("Error reading eap_method: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading eap_method: %v", err)
				}
			}
		}
	}

	if err = d.Set("encoding", flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding2edl(o["encoding"], d, "encoding")); err != nil {
		if vv, ok := fortiAPIPatch(o["encoding"], "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList-Encoding"); ok {
			if err = d.Set("encoding", vv); err != nil {
				return fmt.Errorf("Error reading encoding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encoding: %v", err)
		}
	}

	if err = d.Set("nai_realm", flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm2edl(o["nai-realm"], d, "nai_realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["nai-realm"], "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList-NaiRealm"); ok {
			if err = d.Set("nai_realm", vv); err != nil {
				return fmt.Errorf("Error reading nai_realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nai_realm: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_param"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam2edl(d, i["auth_param"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["auth-param"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex2edl(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["method"], _ = expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod2edl(d, i["method"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex2edl(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "val"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["val"], _ = expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal2edl(d, i["val"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("eap_method"); ok || d.HasChange("eap_method") {
		t, err := expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod2edl(d, v, "eap_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-method"] = t
		}
	}

	if v, ok := d.GetOk("encoding"); ok || d.HasChange("encoding") {
		t, err := expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding2edl(d, v, "encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encoding"] = t
		}
	}

	if v, ok := d.GetOk("nai_realm"); ok || d.HasChange("nai_realm") {
		t, err := expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm2edl(d, v, "nai_realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nai-realm"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
