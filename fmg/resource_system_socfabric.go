// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SOC Fabric.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSocFabric() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSocFabricUpdate,
		Read:   resourceSystemSocFabricRead,
		Update: resourceSystemSocFabricUpdate,
		Delete: resourceSystemSocFabricDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"psk": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secure_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"supervisor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trusted_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"serial": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemSocFabricUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSocFabric(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSocFabric resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSocFabric(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSocFabric resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSocFabric")

	return resourceSystemSocFabricRead(d, m)
}

func resourceSystemSocFabricDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemSocFabric(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSocFabric resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSocFabricRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemSocFabric(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSocFabric resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSocFabric(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSocFabric resource from API: %v", err)
	}
	return nil
}

func flattenSystemSocFabricNameSsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocFabricPortSsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocFabricPskSsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSocFabricRoleSsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocFabricSecureConnectionSsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocFabricStatusSsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocFabricSupervisorSsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocFabricTrustedListSsa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemSocFabricTrustedListIdSsa(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSocFabric-TrustedList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := i["serial"]; ok {
			v := flattenSystemSocFabricTrustedListSerialSsa(i["serial"], d, pre_append)
			tmp["serial"] = fortiAPISubPartPatch(v, "SystemSocFabric-TrustedList-Serial")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSocFabricTrustedListIdSsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocFabricTrustedListSerialSsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSocFabric(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("name", flattenSystemSocFabricNameSsa(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSocFabric-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemSocFabricPortSsa(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemSocFabric-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("role", flattenSystemSocFabricRoleSsa(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "SystemSocFabric-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("secure_connection", flattenSystemSocFabricSecureConnectionSsa(o["secure-connection"], d, "secure_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-connection"], "SystemSocFabric-SecureConnection"); ok {
			if err = d.Set("secure_connection", vv); err != nil {
				return fmt.Errorf("Error reading secure_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_connection: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSocFabricStatusSsa(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSocFabric-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("supervisor", flattenSystemSocFabricSupervisorSsa(o["supervisor"], d, "supervisor")); err != nil {
		if vv, ok := fortiAPIPatch(o["supervisor"], "SystemSocFabric-Supervisor"); ok {
			if err = d.Set("supervisor", vv); err != nil {
				return fmt.Errorf("Error reading supervisor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading supervisor: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("trusted_list", flattenSystemSocFabricTrustedListSsa(o["trusted-list"], d, "trusted_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["trusted-list"], "SystemSocFabric-TrustedList"); ok {
				if err = d.Set("trusted_list", vv); err != nil {
					return fmt.Errorf("Error reading trusted_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading trusted_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("trusted_list"); ok {
			if err = d.Set("trusted_list", flattenSystemSocFabricTrustedListSsa(o["trusted-list"], d, "trusted_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["trusted-list"], "SystemSocFabric-TrustedList"); ok {
					if err = d.Set("trusted_list", vv); err != nil {
						return fmt.Errorf("Error reading trusted_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading trusted_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemSocFabricFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSocFabricNameSsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocFabricPortSsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocFabricPskSsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSocFabricRoleSsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocFabricSecureConnectionSsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocFabricStatusSsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocFabricSupervisorSsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocFabricTrustedListSsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemSocFabricTrustedListIdSsa(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["serial"], _ = expandSystemSocFabricTrustedListSerialSsa(d, i["serial"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSocFabricTrustedListIdSsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocFabricTrustedListSerialSsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSocFabric(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSocFabricNameSsa(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemSocFabricPortSsa(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("psk"); ok || d.HasChange("psk") {
		t, err := expandSystemSocFabricPskSsa(d, v, "psk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psk"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandSystemSocFabricRoleSsa(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("secure_connection"); ok || d.HasChange("secure_connection") {
		t, err := expandSystemSocFabricSecureConnectionSsa(d, v, "secure_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-connection"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSocFabricStatusSsa(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("supervisor"); ok || d.HasChange("supervisor") {
		t, err := expandSystemSocFabricSupervisorSsa(d, v, "supervisor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["supervisor"] = t
		}
	}

	if v, ok := d.GetOk("trusted_list"); ok || d.HasChange("trusted_list") {
		t, err := expandSystemSocFabricTrustedListSsa(d, v, "trusted_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted-list"] = t
		}
	}

	return &obj, nil
}
