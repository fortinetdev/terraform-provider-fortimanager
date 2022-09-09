// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectUser Nsx

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectUserNsx() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserNsxCreate,
		Read:   resourceObjectUserNsxRead,
		Update: resourceObjectUserNsxUpdate,
		Delete: resourceObjectUserNsxDelete,

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
			"fmgip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fmgpasswd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"fmguser": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"if_allgroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"integration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ref_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"service_manager_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_manager_rev": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
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

func resourceObjectUserNsxCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserNsx(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserNsx resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserNsx(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserNsx resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserNsxRead(d, m)
}

func resourceObjectUserNsxUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserNsx(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserNsx resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserNsx(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserNsx resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserNsxRead(d, m)
}

func resourceObjectUserNsxDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectUserNsx(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserNsx resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserNsxRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectUserNsx(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserNsx resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserNsx(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserNsx resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserNsxFmgip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxFmgpasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserNsxFmguser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxIfAllgroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserNsxServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxService(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserNsxServiceId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectUserNsx-Service-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "integration"
		if _, ok := i["integration"]; ok {
			v := flattenObjectUserNsxServiceIntegration(i["integration"], d, pre_append)
			tmp["integration"] = fortiAPISubPartPatch(v, "ObjectUserNsx-Service-Integration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectUserNsxServiceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserNsx-Service-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ref_id"
		if _, ok := i["ref-id"]; ok {
			v := flattenObjectUserNsxServiceRefId(i["ref-id"], d, pre_append)
			tmp["ref_id"] = fortiAPISubPartPatch(v, "ObjectUserNsx-Service-RefId")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectUserNsxServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxServiceIntegration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxServiceRefId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxServiceIdU(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserNsxServiceManagerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxServiceManagerRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserNsx(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("fmgip", flattenObjectUserNsxFmgip(o["fmgip"], d, "fmgip")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmgip"], "ObjectUserNsx-Fmgip"); ok {
			if err = d.Set("fmgip", vv); err != nil {
				return fmt.Errorf("Error reading fmgip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgip: %v", err)
		}
	}

	if err = d.Set("fmguser", flattenObjectUserNsxFmguser(o["fmguser"], d, "fmguser")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmguser"], "ObjectUserNsx-Fmguser"); ok {
			if err = d.Set("fmguser", vv); err != nil {
				return fmt.Errorf("Error reading fmguser: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmguser: %v", err)
		}
	}

	if err = d.Set("if_allgroup", flattenObjectUserNsxIfAllgroup(o["if-allgroup"], d, "if_allgroup")); err != nil {
		if vv, ok := fortiAPIPatch(o["if-allgroup"], "ObjectUserNsx-IfAllgroup"); ok {
			if err = d.Set("if_allgroup", vv); err != nil {
				return fmt.Errorf("Error reading if_allgroup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading if_allgroup: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserNsxName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserNsx-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server", flattenObjectUserNsxServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "ObjectUserNsx-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service", flattenObjectUserNsxService(o["service"], d, "service")); err != nil {
			if vv, ok := fortiAPIPatch(o["service"], "ObjectUserNsx-Service"); ok {
				if err = d.Set("service", vv); err != nil {
					return fmt.Errorf("Error reading service: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service"); ok {
			if err = d.Set("service", flattenObjectUserNsxService(o["service"], d, "service")); err != nil {
				if vv, ok := fortiAPIPatch(o["service"], "ObjectUserNsx-Service"); ok {
					if err = d.Set("service", vv); err != nil {
						return fmt.Errorf("Error reading service: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading service: %v", err)
				}
			}
		}
	}

	if err = d.Set("service_id", flattenObjectUserNsxServiceIdU(o["service-id"], d, "service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-id"], "ObjectUserNsx-ServiceId"); ok {
			if err = d.Set("service_id", vv); err != nil {
				return fmt.Errorf("Error reading service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_id: %v", err)
		}
	}

	if err = d.Set("service_manager_id", flattenObjectUserNsxServiceManagerId(o["service-manager-id"], d, "service_manager_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-manager-id"], "ObjectUserNsx-ServiceManagerId"); ok {
			if err = d.Set("service_manager_id", vv); err != nil {
				return fmt.Errorf("Error reading service_manager_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_manager_id: %v", err)
		}
	}

	if err = d.Set("service_manager_rev", flattenObjectUserNsxServiceManagerRev(o["service-manager-rev"], d, "service_manager_rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-manager-rev"], "ObjectUserNsx-ServiceManagerRev"); ok {
			if err = d.Set("service_manager_rev", vv); err != nil {
				return fmt.Errorf("Error reading service_manager_rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_manager_rev: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectUserNsxStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectUserNsx-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("user", flattenObjectUserNsxUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "ObjectUserNsx-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	return nil
}

func flattenObjectUserNsxFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserNsxFmgip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxFmgpasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserNsxFmguser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxIfAllgroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserNsxServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandObjectUserNsxServiceId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "integration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["integration"], _ = expandObjectUserNsxServiceIntegration(d, i["integration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectUserNsxServiceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ref_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ref-id"], _ = expandObjectUserNsxServiceRefId(d, i["ref_id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectUserNsxServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxServiceIntegration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxServiceRefId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxServiceIdU(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserNsxServiceManagerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxServiceManagerRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserNsx(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgip"); ok || d.HasChange("fmgip") {
		t, err := expandObjectUserNsxFmgip(d, v, "fmgip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmgip"] = t
		}
	}

	if v, ok := d.GetOk("fmgpasswd"); ok || d.HasChange("fmgpasswd") {
		t, err := expandObjectUserNsxFmgpasswd(d, v, "fmgpasswd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmgpasswd"] = t
		}
	}

	if v, ok := d.GetOk("fmguser"); ok || d.HasChange("fmguser") {
		t, err := expandObjectUserNsxFmguser(d, v, "fmguser")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmguser"] = t
		}
	}

	if v, ok := d.GetOk("if_allgroup"); ok || d.HasChange("if_allgroup") {
		t, err := expandObjectUserNsxIfAllgroup(d, v, "if_allgroup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["if-allgroup"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserNsxName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectUserNsxPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandObjectUserNsxServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandObjectUserNsxService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("service_id"); ok || d.HasChange("service_id") {
		t, err := expandObjectUserNsxServiceIdU(d, v, "service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-id"] = t
		}
	}

	if v, ok := d.GetOk("service_manager_id"); ok || d.HasChange("service_manager_id") {
		t, err := expandObjectUserNsxServiceManagerId(d, v, "service_manager_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-manager-id"] = t
		}
	}

	if v, ok := d.GetOk("service_manager_rev"); ok || d.HasChange("service_manager_rev") {
		t, err := expandObjectUserNsxServiceManagerRev(d, v, "service_manager_rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-manager-rev"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectUserNsxStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandObjectUserNsxUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	return &obj, nil
}
