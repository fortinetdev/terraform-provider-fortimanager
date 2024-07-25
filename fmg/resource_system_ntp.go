// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: NTP settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNtpUpdate,
		Read:   resourceSystemNtpRead,
		Update: resourceSystemNtpUpdate,
		Delete: resourceSystemNtpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ntpserver": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"key": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"key_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"maxpoll": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"minpoll": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ntpv3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sync_interval": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSystemNtpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemNtp(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemNtp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemNtp(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemNtp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemNtp")

	return resourceSystemNtpRead(d, m)
}

func resourceSystemNtpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemNtp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNtpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemNtp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemNtp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNtp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemNtp resource from API: %v", err)
	}
	return nil
}

func flattenSystemNtpNtpserver(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			v := flattenSystemNtpNtpserverAuthentication(i["authentication"], d, pre_append)
			tmp["authentication"] = fortiAPISubPartPatch(v, "SystemNtp-Ntpserver-Authentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemNtpNtpserverId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemNtp-Ntpserver-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_id"
		if _, ok := i["key-id"]; ok {
			v := flattenSystemNtpNtpserverKeyId(i["key-id"], d, pre_append)
			tmp["key_id"] = fortiAPISubPartPatch(v, "SystemNtp-Ntpserver-KeyId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maxpoll"
		if _, ok := i["maxpoll"]; ok {
			v := flattenSystemNtpNtpserverMaxpoll(i["maxpoll"], d, pre_append)
			tmp["maxpoll"] = fortiAPISubPartPatch(v, "SystemNtp-Ntpserver-Maxpoll")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minpoll"
		if _, ok := i["minpoll"]; ok {
			v := flattenSystemNtpNtpserverMinpoll(i["minpoll"], d, pre_append)
			tmp["minpoll"] = fortiAPISubPartPatch(v, "SystemNtp-Ntpserver-Minpoll")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ntpv3"
		if _, ok := i["ntpv3"]; ok {
			v := flattenSystemNtpNtpserverNtpv3(i["ntpv3"], d, pre_append)
			tmp["ntpv3"] = fortiAPISubPartPatch(v, "SystemNtp-Ntpserver-Ntpv3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenSystemNtpNtpserverServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "SystemNtp-Ntpserver-Server")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemNtpNtpserverAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverKeyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverMaxpoll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverMinpoll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverNtpv3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpSyncInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemNtp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("ntpserver", flattenSystemNtpNtpserver(o["ntpserver"], d, "ntpserver")); err != nil {
			if vv, ok := fortiAPIPatch(o["ntpserver"], "SystemNtp-Ntpserver"); ok {
				if err = d.Set("ntpserver", vv); err != nil {
					return fmt.Errorf("Error reading ntpserver: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ntpserver: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ntpserver"); ok {
			if err = d.Set("ntpserver", flattenSystemNtpNtpserver(o["ntpserver"], d, "ntpserver")); err != nil {
				if vv, ok := fortiAPIPatch(o["ntpserver"], "SystemNtp-Ntpserver"); ok {
					if err = d.Set("ntpserver", vv); err != nil {
						return fmt.Errorf("Error reading ntpserver: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ntpserver: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenSystemNtpStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemNtp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("sync_interval", flattenSystemNtpSyncInterval(o["sync_interval"], d, "sync_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync_interval"], "SystemNtp-SyncInterval"); ok {
			if err = d.Set("sync_interval", vv); err != nil {
				return fmt.Errorf("Error reading sync_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_interval: %v", err)
		}
	}

	return nil
}

func flattenSystemNtpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemNtpNtpserver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authentication"], _ = expandSystemNtpNtpserverAuthentication(d, i["authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemNtpNtpserverId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key"], _ = expandSystemNtpNtpserverKey(d, i["key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-id"], _ = expandSystemNtpNtpserverKeyId(d, i["key_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maxpoll"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maxpoll"], _ = expandSystemNtpNtpserverMaxpoll(d, i["maxpoll"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minpoll"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["minpoll"], _ = expandSystemNtpNtpserverMinpoll(d, i["minpoll"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ntpv3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ntpv3"], _ = expandSystemNtpNtpserverNtpv3(d, i["ntpv3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandSystemNtpNtpserverServer(d, i["server"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemNtpNtpserverAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemNtpNtpserverKeyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverMaxpoll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverMinpoll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverNtpv3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpSyncInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNtp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ntpserver"); ok || d.HasChange("ntpserver") {
		t, err := expandSystemNtpNtpserver(d, v, "ntpserver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntpserver"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemNtpStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("sync_interval"); ok || d.HasChange("sync_interval") {
		t, err := expandSystemNtpSyncInterval(d, v, "sync_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync_interval"] = t
		}
	}

	return &obj, nil
}
