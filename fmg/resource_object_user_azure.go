// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectUser Azure

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserAzure() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserAzureCreate,
		Read:   resourceObjectUserAzureRead,
		Update: resourceObjectUserAzureUpdate,
		Delete: resourceObjectUserAzureDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"spn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"kbconfig": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"page_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"proxy_enable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_passwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"proxy_scheme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"realm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rule": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"select_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenantid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upd_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"verifycert": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceObjectUserAzureCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserAzure(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserAzure resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectUserAzure(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectUserAzure(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectUserAzure resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectUserAzure(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectUserAzure resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserAzureRead(d, m)
}

func resourceObjectUserAzureUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserAzure(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserAzure resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectUserAzure(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserAzure resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserAzureRead(d, m)
}

func resourceObjectUserAzureDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectUserAzure(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserAzure resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserAzureRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectUserAzure(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectUserAzure resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserAzure(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserAzure resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserAzureSpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureKbconfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzurePageSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureProxyEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureProxyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureProxyScheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureProxyUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectUserAzureRuleName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserAzure-Rule-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule"
		if _, ok := i["rule"]; ok {
			v := flattenObjectUserAzureRuleRule(i["rule"], d, pre_append)
			tmp["rule"] = fortiAPISubPartPatch(v, "ObjectUserAzure-Rule-Rule")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserAzureRuleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureRuleRule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureSelectProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureTenantid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureUpdInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAzureVerifycert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserAzure(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("spn", flattenObjectUserAzureSpn(o["SPN"], d, "spn")); err != nil {
		if vv, ok := fortiAPIPatch(o["SPN"], "ObjectUserAzure-Spn"); ok {
			if err = d.Set("spn", vv); err != nil {
				return fmt.Errorf("Error reading spn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spn: %v", err)
		}
	}

	if err = d.Set("alias", flattenObjectUserAzureAlias(o["alias"], d, "alias")); err != nil {
		if vv, ok := fortiAPIPatch(o["alias"], "ObjectUserAzure-Alias"); ok {
			if err = d.Set("alias", vv); err != nil {
				return fmt.Errorf("Error reading alias: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("kbconfig", flattenObjectUserAzureKbconfig(o["kbconfig"], d, "kbconfig")); err != nil {
		if vv, ok := fortiAPIPatch(o["kbconfig"], "ObjectUserAzure-Kbconfig"); ok {
			if err = d.Set("kbconfig", vv); err != nil {
				return fmt.Errorf("Error reading kbconfig: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading kbconfig: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserAzureName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserAzure-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("page_size", flattenObjectUserAzurePageSize(o["page_size"], d, "page_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["page_size"], "ObjectUserAzure-PageSize"); ok {
			if err = d.Set("page_size", vv); err != nil {
				return fmt.Errorf("Error reading page_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading page_size: %v", err)
		}
	}

	if err = d.Set("proxy_enable", flattenObjectUserAzureProxyEnable(o["proxy_enable"], d, "proxy_enable")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy_enable"], "ObjectUserAzure-ProxyEnable"); ok {
			if err = d.Set("proxy_enable", vv); err != nil {
				return fmt.Errorf("Error reading proxy_enable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_enable: %v", err)
		}
	}

	if err = d.Set("proxy_host", flattenObjectUserAzureProxyHost(o["proxy_host"], d, "proxy_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy_host"], "ObjectUserAzure-ProxyHost"); ok {
			if err = d.Set("proxy_host", vv); err != nil {
				return fmt.Errorf("Error reading proxy_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_host: %v", err)
		}
	}

	if err = d.Set("proxy_scheme", flattenObjectUserAzureProxyScheme(o["proxy_scheme"], d, "proxy_scheme")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy_scheme"], "ObjectUserAzure-ProxyScheme"); ok {
			if err = d.Set("proxy_scheme", vv); err != nil {
				return fmt.Errorf("Error reading proxy_scheme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_scheme: %v", err)
		}
	}

	if err = d.Set("proxy_user", flattenObjectUserAzureProxyUser(o["proxy_user"], d, "proxy_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy_user"], "ObjectUserAzure-ProxyUser"); ok {
			if err = d.Set("proxy_user", vv); err != nil {
				return fmt.Errorf("Error reading proxy_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_user: %v", err)
		}
	}

	if err = d.Set("realm", flattenObjectUserAzureRealm(o["realm"], d, "realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["realm"], "ObjectUserAzure-Realm"); ok {
			if err = d.Set("realm", vv); err != nil {
				return fmt.Errorf("Error reading realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading realm: %v", err)
		}
	}

	if err = d.Set("region", flattenObjectUserAzureRegion(o["region"], d, "region")); err != nil {
		if vv, ok := fortiAPIPatch(o["region"], "ObjectUserAzure-Region"); ok {
			if err = d.Set("region", vv); err != nil {
				return fmt.Errorf("Error reading region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading region: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rule", flattenObjectUserAzureRule(o["rule"], d, "rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["rule"], "ObjectUserAzure-Rule"); ok {
				if err = d.Set("rule", vv); err != nil {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenObjectUserAzureRule(o["rule"], d, "rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["rule"], "ObjectUserAzure-Rule"); ok {
					if err = d.Set("rule", vv); err != nil {
						return fmt.Errorf("Error reading rule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			}
		}
	}

	if err = d.Set("select_proxy", flattenObjectUserAzureSelectProxy(o["select_proxy"], d, "select_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["select_proxy"], "ObjectUserAzure-SelectProxy"); ok {
			if err = d.Set("select_proxy", vv); err != nil {
				return fmt.Errorf("Error reading select_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading select_proxy: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectUserAzureStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectUserAzure-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tenantid", flattenObjectUserAzureTenantid(o["tenantid"], d, "tenantid")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenantid"], "ObjectUserAzure-Tenantid"); ok {
			if err = d.Set("tenantid", vv); err != nil {
				return fmt.Errorf("Error reading tenantid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenantid: %v", err)
		}
	}

	if err = d.Set("upd_interval", flattenObjectUserAzureUpdInterval(o["upd_interval"], d, "upd_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["upd_interval"], "ObjectUserAzure-UpdInterval"); ok {
			if err = d.Set("upd_interval", vv); err != nil {
				return fmt.Errorf("Error reading upd_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upd_interval: %v", err)
		}
	}

	if err = d.Set("user", flattenObjectUserAzureUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "ObjectUserAzure-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("verifycert", flattenObjectUserAzureVerifycert(o["verifycert"], d, "verifycert")); err != nil {
		if vv, ok := fortiAPIPatch(o["verifycert"], "ObjectUserAzure-Verifycert"); ok {
			if err = d.Set("verifycert", vv); err != nil {
				return fmt.Errorf("Error reading verifycert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verifycert: %v", err)
		}
	}

	return nil
}

func flattenObjectUserAzureFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserAzureSpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureAlias(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureKbconfig(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzurePageSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzurePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserAzureProxyEnable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureProxyHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureProxyPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserAzureProxyScheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureProxyUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectUserAzureRuleName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rule"], _ = expandObjectUserAzureRuleRule(d, i["rule"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserAzureRuleName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureRuleRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureSelectProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureTenantid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureUpdInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAzureVerifycert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserAzure(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("spn"); ok || d.HasChange("spn") {
		t, err := expandObjectUserAzureSpn(d, v, "spn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["SPN"] = t
		}
	}

	if v, ok := d.GetOk("alias"); ok || d.HasChange("alias") {
		t, err := expandObjectUserAzureAlias(d, v, "alias")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	}

	if v, ok := d.GetOk("kbconfig"); ok || d.HasChange("kbconfig") {
		t, err := expandObjectUserAzureKbconfig(d, v, "kbconfig")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["kbconfig"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserAzureName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("page_size"); ok || d.HasChange("page_size") {
		t, err := expandObjectUserAzurePageSize(d, v, "page_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["page_size"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectUserAzurePassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("proxy_enable"); ok || d.HasChange("proxy_enable") {
		t, err := expandObjectUserAzureProxyEnable(d, v, "proxy_enable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy_enable"] = t
		}
	}

	if v, ok := d.GetOk("proxy_host"); ok || d.HasChange("proxy_host") {
		t, err := expandObjectUserAzureProxyHost(d, v, "proxy_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy_host"] = t
		}
	}

	if v, ok := d.GetOk("proxy_passwd"); ok || d.HasChange("proxy_passwd") {
		t, err := expandObjectUserAzureProxyPasswd(d, v, "proxy_passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy_passwd"] = t
		}
	}

	if v, ok := d.GetOk("proxy_scheme"); ok || d.HasChange("proxy_scheme") {
		t, err := expandObjectUserAzureProxyScheme(d, v, "proxy_scheme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy_scheme"] = t
		}
	}

	if v, ok := d.GetOk("proxy_user"); ok || d.HasChange("proxy_user") {
		t, err := expandObjectUserAzureProxyUser(d, v, "proxy_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy_user"] = t
		}
	}

	if v, ok := d.GetOk("realm"); ok || d.HasChange("realm") {
		t, err := expandObjectUserAzureRealm(d, v, "realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realm"] = t
		}
	}

	if v, ok := d.GetOk("region"); ok || d.HasChange("region") {
		t, err := expandObjectUserAzureRegion(d, v, "region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandObjectUserAzureRule(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	if v, ok := d.GetOk("select_proxy"); ok || d.HasChange("select_proxy") {
		t, err := expandObjectUserAzureSelectProxy(d, v, "select_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["select_proxy"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectUserAzureStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tenantid"); ok || d.HasChange("tenantid") {
		t, err := expandObjectUserAzureTenantid(d, v, "tenantid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenantid"] = t
		}
	}

	if v, ok := d.GetOk("upd_interval"); ok || d.HasChange("upd_interval") {
		t, err := expandObjectUserAzureUpdInterval(d, v, "upd_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upd_interval"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandObjectUserAzureUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("verifycert"); ok || d.HasChange("verifycert") {
		t, err := expandObjectUserAzureVerifycert(d, v, "verifycert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verifycert"] = t
		}
	}

	return &obj, nil
}
