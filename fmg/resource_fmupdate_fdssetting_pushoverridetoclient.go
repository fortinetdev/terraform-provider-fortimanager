// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Enable/disable push updates, and override the default IP address and port used by FortiGuard to send antivirus and IPS push messages for clients.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFmupdateFdsSettingPushOverrideToClient() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateFdsSettingPushOverrideToClientUpdate,
		Read:   resourceFmupdateFdsSettingPushOverrideToClientRead,
		Update: resourceFmupdateFdsSettingPushOverrideToClientUpdate,
		Delete: resourceFmupdateFdsSettingPushOverrideToClientDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"announce_ip": &schema.Schema{
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
						"port": &schema.Schema{
							Type:     schema.TypeInt,
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

func resourceFmupdateFdsSettingPushOverrideToClientUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateFdsSettingPushOverrideToClient(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFdsSettingPushOverrideToClient resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateFdsSettingPushOverrideToClient(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFdsSettingPushOverrideToClient resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateFdsSettingPushOverrideToClient")

	return resourceFmupdateFdsSettingPushOverrideToClientRead(d, m)
}

func resourceFmupdateFdsSettingPushOverrideToClientDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateFdsSettingPushOverrideToClient(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateFdsSettingPushOverrideToClient resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateFdsSettingPushOverrideToClientRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateFdsSettingPushOverrideToClient(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFdsSettingPushOverrideToClient resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateFdsSettingPushOverrideToClient(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFdsSettingPushOverrideToClient resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpFfpb(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpIdFfpb(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingPushOverrideToClient-AnnounceIp-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpIpFfpb(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingPushOverrideToClient-AnnounceIp-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpPortFfpb(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingPushOverrideToClient-AnnounceIp-Port")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpIdFfpb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpIpFfpb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpPortFfpb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClientStatusFfpb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateFdsSettingPushOverrideToClient(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("announce_ip", flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpFfpb(o["announce-ip"], d, "announce_ip")); err != nil {
			if vv, ok := fortiAPIPatch(o["announce-ip"], "FmupdateFdsSettingPushOverrideToClient-AnnounceIp"); ok {
				if err = d.Set("announce_ip", vv); err != nil {
					return fmt.Errorf("Error reading announce_ip: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading announce_ip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("announce_ip"); ok {
			if err = d.Set("announce_ip", flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpFfpb(o["announce-ip"], d, "announce_ip")); err != nil {
				if vv, ok := fortiAPIPatch(o["announce-ip"], "FmupdateFdsSettingPushOverrideToClient-AnnounceIp"); ok {
					if err = d.Set("announce_ip", vv); err != nil {
						return fmt.Errorf("Error reading announce_ip: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading announce_ip: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenFmupdateFdsSettingPushOverrideToClientStatusFfpb(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FmupdateFdsSettingPushOverrideToClient-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenFmupdateFdsSettingPushOverrideToClientFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpFfpb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFmupdateFdsSettingPushOverrideToClientAnnounceIpIdFfpb(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandFmupdateFdsSettingPushOverrideToClientAnnounceIpIpFfpb(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandFmupdateFdsSettingPushOverrideToClientAnnounceIpPortFfpb(d, i["port"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpIdFfpb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpIpFfpb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpPortFfpb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClientStatusFfpb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateFdsSettingPushOverrideToClient(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("announce_ip"); ok {
		t, err := expandFmupdateFdsSettingPushOverrideToClientAnnounceIpFfpb(d, v, "announce_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["announce-ip"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandFmupdateFdsSettingPushOverrideToClientStatusFfpb(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
