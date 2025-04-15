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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectFmupdateFdsSettingPushOverrideToClient(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFdsSettingPushOverrideToClient resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFmupdateFdsSettingPushOverrideToClient(obj, mkey, paradict, wsParams)
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteFmupdateFdsSettingPushOverrideToClient(mkey, paradict, wsParams)
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

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadFmupdateFdsSettingPushOverrideToClient(mkey, paradict)
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

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIp2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingPushOverrideToClient-AnnounceIp-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingPushOverrideToClient-AnnounceIp-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpPort2edl(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FmupdateFdsSettingPushOverrideToClient-AnnounceIp-Port")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClientStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateFdsSettingPushOverrideToClient(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("announce_ip", flattenFmupdateFdsSettingPushOverrideToClientAnnounceIp2edl(o["announce-ip"], d, "announce_ip")); err != nil {
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
			if err = d.Set("announce_ip", flattenFmupdateFdsSettingPushOverrideToClientAnnounceIp2edl(o["announce-ip"], d, "announce_ip")); err != nil {
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

	if err = d.Set("status", flattenFmupdateFdsSettingPushOverrideToClientStatus2edl(o["status"], d, "status")); err != nil {
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

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFmupdateFdsSettingPushOverrideToClientAnnounceIpId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFmupdateFdsSettingPushOverrideToClientAnnounceIpIp2edl(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFmupdateFdsSettingPushOverrideToClientAnnounceIpPort2edl(d, i["port"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClientStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateFdsSettingPushOverrideToClient(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("announce_ip"); ok || d.HasChange("announce_ip") {
		t, err := expandFmupdateFdsSettingPushOverrideToClientAnnounceIp2edl(d, v, "announce_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["announce-ip"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFmupdateFdsSettingPushOverrideToClientStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
