// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Server override configure.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFmupdateFdsSettingServerOverride() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateFdsSettingServerOverrideUpdate,
		Read:   resourceFmupdateFdsSettingServerOverrideRead,
		Update: resourceFmupdateFdsSettingServerOverrideUpdate,
		Delete: resourceFmupdateFdsSettingServerOverrideDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"servlist": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"service_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"status": &schema.Schema{
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

func resourceFmupdateFdsSettingServerOverrideUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateFdsSettingServerOverride(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFdsSettingServerOverride resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateFdsSettingServerOverride(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFdsSettingServerOverride resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateFdsSettingServerOverride")

	return resourceFmupdateFdsSettingServerOverrideRead(d, m)
}

func resourceFmupdateFdsSettingServerOverrideDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateFdsSettingServerOverride(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateFdsSettingServerOverride resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateFdsSettingServerOverrideRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateFdsSettingServerOverride(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFdsSettingServerOverride resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateFdsSettingServerOverride(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFdsSettingServerOverride resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateFdsSettingServerOverrideServlistFfsa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFmupdateFdsSettingServerOverrideServlistIdFfsa(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingServerOverride-Servlist-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFmupdateFdsSettingServerOverrideServlistIpFfsa(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingServerOverride-Servlist-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenFmupdateFdsSettingServerOverrideServlistIp6Ffsa(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingServerOverride-Servlist-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFmupdateFdsSettingServerOverrideServlistPortFfsa(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingServerOverride-Servlist-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_type"
		if _, ok := i["service-type"]; ok {
			v := flattenFmupdateFdsSettingServerOverrideServlistServiceTypeFfsa(i["service-type"], d, pre_append)
			tmp["service_type"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingServerOverride-Servlist-ServiceType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFmupdateFdsSettingServerOverrideServlistIdFfsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistIpFfsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistIp6Ffsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistPortFfsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistServiceTypeFfsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideStatusFfsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateFdsSettingServerOverride(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("servlist", flattenFmupdateFdsSettingServerOverrideServlistFfsa(o["servlist"], d, "servlist")); err != nil {
			if vv, ok := fortiAPIPatch(o["servlist"], "FmupdateFdsSettingServerOverride-Servlist"); ok {
				if err = d.Set("servlist", vv); err != nil {
					return fmt.Errorf("Error reading servlist: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading servlist: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("servlist"); ok {
			if err = d.Set("servlist", flattenFmupdateFdsSettingServerOverrideServlistFfsa(o["servlist"], d, "servlist")); err != nil {
				if vv, ok := fortiAPIPatch(o["servlist"], "FmupdateFdsSettingServerOverride-Servlist"); ok {
					if err = d.Set("servlist", vv); err != nil {
						return fmt.Errorf("Error reading servlist: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading servlist: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenFmupdateFdsSettingServerOverrideStatusFfsa(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FmupdateFdsSettingServerOverride-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenFmupdateFdsSettingServerOverrideFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateFdsSettingServerOverrideServlistFfsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFmupdateFdsSettingServerOverrideServlistIdFfsa(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFmupdateFdsSettingServerOverrideServlistIpFfsa(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandFmupdateFdsSettingServerOverrideServlistIp6Ffsa(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFmupdateFdsSettingServerOverrideServlistPortFfsa(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-type"], _ = expandFmupdateFdsSettingServerOverrideServlistServiceTypeFfsa(d, i["service_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFmupdateFdsSettingServerOverrideServlistIdFfsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistIpFfsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistIp6Ffsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistPortFfsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistServiceTypeFfsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideStatusFfsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateFdsSettingServerOverride(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("servlist"); ok || d.HasChange("servlist") {
		t, err := expandFmupdateFdsSettingServerOverrideServlistFfsa(d, v, "servlist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["servlist"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFmupdateFdsSettingServerOverrideStatusFfsa(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
