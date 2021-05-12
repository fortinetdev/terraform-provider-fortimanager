// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Probe a remote device and retrieve its device information and system status.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDvmCmdDiscoverDevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmCmdDiscoverDeviceUpdate,
		Read:   resourceDvmCmdDiscoverDeviceRead,
		Update: resourceDvmCmdDiscoverDeviceUpdate,
		Delete: resourceDvmCmdDiscoverDeviceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"force_recreate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"adm_pass": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"adm_usr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceDvmCmdDiscoverDeviceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "", fmt.Errorf("")

	obj, err := getObjectDvmCmdDiscoverDevice(d)
	if err != nil {
		return fmt.Errorf("Error updating DvmCmdDiscoverDevice resource while getting object: %v", err)
	}

	_, err = c.UpdateDvmCmdDiscoverDevice(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating DvmCmdDiscoverDevice resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("DvmCmdDiscoverDevice")

	return resourceDvmCmdDiscoverDeviceRead(d, m)
}

func resourceDvmCmdDiscoverDeviceDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceDvmCmdDiscoverDeviceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenDvmCmdDiscoverDeviceDevice(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "adm_pass"
	if _, ok := i["adm_pass"]; ok {
		result["adm_pass"] = flattenDvmCmdDiscoverDeviceDeviceAdmPass(i["adm_pass"], d, pre_append)
	}

	pre_append = pre + ".0." + "adm_usr"
	if _, ok := i["adm_usr"]; ok {
		result["adm_usr"] = flattenDvmCmdDiscoverDeviceDeviceAdmUsr(i["adm_usr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip"
	if _, ok := i["ip"]; ok {
		result["ip"] = flattenDvmCmdDiscoverDeviceDeviceIp(i["ip"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenDvmCmdDiscoverDeviceDeviceAdmPass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdDiscoverDeviceDeviceAdmUsr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdDiscoverDeviceDeviceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDvmCmdDiscoverDevice(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("device", flattenDvmCmdDiscoverDeviceDevice(o["device"], d, "device")); err != nil {
			if vv, ok := fortiAPIPatch(o["device"], "DvmCmdDiscoverDevice-Device"); ok {
				if err = d.Set("device", vv); err != nil {
					return fmt.Errorf("Error reading device: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading device: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("device"); ok {
			if err = d.Set("device", flattenDvmCmdDiscoverDeviceDevice(o["device"], d, "device")); err != nil {
				if vv, ok := fortiAPIPatch(o["device"], "DvmCmdDiscoverDevice-Device"); ok {
					if err = d.Set("device", vv); err != nil {
						return fmt.Errorf("Error reading device: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading device: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenDvmCmdDiscoverDeviceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmCmdDiscoverDeviceDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "adm_pass"
	if _, ok := d.GetOk(pre_append); ok {
		result["adm_pass"], _ = expandDvmCmdDiscoverDeviceDeviceAdmPass(d, i["adm_pass"], pre_append)
	}
	pre_append = pre + ".0." + "adm_usr"
	if _, ok := d.GetOk(pre_append); ok {
		result["adm_usr"], _ = expandDvmCmdDiscoverDeviceDeviceAdmUsr(d, i["adm_usr"], pre_append)
	}
	pre_append = pre + ".0." + "ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip"], _ = expandDvmCmdDiscoverDeviceDeviceIp(d, i["ip"], pre_append)
	}

	return result, nil
}

func expandDvmCmdDiscoverDeviceDeviceAdmPass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdDiscoverDeviceDeviceAdmUsr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdDiscoverDeviceDeviceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDvmCmdDiscoverDevice(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device"); ok {
		t, err := expandDvmCmdDiscoverDeviceDevice(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	return &obj, nil
}
