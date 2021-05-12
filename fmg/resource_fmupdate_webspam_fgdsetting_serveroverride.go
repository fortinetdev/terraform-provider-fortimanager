// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Server override configure.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFmupdateWebSpamFgdSettingServerOverride() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateWebSpamFgdSettingServerOverrideUpdate,
		Read:   resourceFmupdateWebSpamFgdSettingServerOverrideRead,
		Update: resourceFmupdateWebSpamFgdSettingServerOverrideUpdate,
		Delete: resourceFmupdateWebSpamFgdSettingServerOverrideDelete,

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
							Computed: true,
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

func resourceFmupdateWebSpamFgdSettingServerOverrideUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateWebSpamFgdSettingServerOverride(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateWebSpamFgdSettingServerOverride resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateWebSpamFgdSettingServerOverride(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateWebSpamFgdSettingServerOverride resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateWebSpamFgdSettingServerOverride")

	return resourceFmupdateWebSpamFgdSettingServerOverrideRead(d, m)
}

func resourceFmupdateWebSpamFgdSettingServerOverrideDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateWebSpamFgdSettingServerOverride(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateWebSpamFgdSettingServerOverride resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateWebSpamFgdSettingServerOverrideRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateWebSpamFgdSettingServerOverride(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateWebSpamFgdSettingServerOverride resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateWebSpamFgdSettingServerOverride(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateWebSpamFgdSettingServerOverride resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistFwfsb(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFmupdateWebSpamFgdSettingServerOverrideServlistIdFwfsb(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FmupdateWebSpamFgdSettingServerOverride-Servlist-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFmupdateWebSpamFgdSettingServerOverrideServlistIpFwfsb(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FmupdateWebSpamFgdSettingServerOverride-Servlist-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenFmupdateWebSpamFgdSettingServerOverrideServlistIp6Fwfsb(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "FmupdateWebSpamFgdSettingServerOverride-Servlist-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFmupdateWebSpamFgdSettingServerOverrideServlistPortFwfsb(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FmupdateWebSpamFgdSettingServerOverride-Servlist-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_type"
		if _, ok := i["service-type"]; ok {
			v := flattenFmupdateWebSpamFgdSettingServerOverrideServlistServiceTypeFwfsb(i["service-type"], d, pre_append)
			tmp["service_type"] = fortiAPISubPartPatch(v, "FmupdateWebSpamFgdSettingServerOverride-Servlist-ServiceType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistIdFwfsb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistIpFwfsb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistIp6Fwfsb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistPortFwfsb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistServiceTypeFwfsb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			256:  "fgd",
			512:  "fgc",
			1024: "fsa",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideStatusFwfsb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectFmupdateWebSpamFgdSettingServerOverride(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("servlist", flattenFmupdateWebSpamFgdSettingServerOverrideServlistFwfsb(o["servlist"], d, "servlist")); err != nil {
			if vv, ok := fortiAPIPatch(o["servlist"], "FmupdateWebSpamFgdSettingServerOverride-Servlist"); ok {
				if err = d.Set("servlist", vv); err != nil {
					return fmt.Errorf("Error reading servlist: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading servlist: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("servlist"); ok {
			if err = d.Set("servlist", flattenFmupdateWebSpamFgdSettingServerOverrideServlistFwfsb(o["servlist"], d, "servlist")); err != nil {
				if vv, ok := fortiAPIPatch(o["servlist"], "FmupdateWebSpamFgdSettingServerOverride-Servlist"); ok {
					if err = d.Set("servlist", vv); err != nil {
						return fmt.Errorf("Error reading servlist: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading servlist: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenFmupdateWebSpamFgdSettingServerOverrideStatusFwfsb(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FmupdateWebSpamFgdSettingServerOverride-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenFmupdateWebSpamFgdSettingServerOverrideFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistFwfsb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistIdFwfsb(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistIpFwfsb(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip6"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistIp6Fwfsb(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistPortFwfsb(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["service-type"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistServiceTypeFwfsb(d, i["service_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistIdFwfsb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistIpFwfsb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistIp6Fwfsb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistPortFwfsb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistServiceTypeFwfsb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideStatusFwfsb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateWebSpamFgdSettingServerOverride(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("servlist"); ok {
		t, err := expandFmupdateWebSpamFgdSettingServerOverrideServlistFwfsb(d, v, "servlist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["servlist"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandFmupdateWebSpamFgdSettingServerOverrideStatusFwfsb(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
