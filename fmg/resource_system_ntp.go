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

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemNtp(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemNtp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemNtp(obj, adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemNtp(adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemNtp(adomv, mkey, nil)
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

func flattenSystemNtpNtpserverSna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemNtpNtpserverAuthenticationSna(i["authentication"], d, pre_append)
			tmp["authentication"] = fortiAPISubPartPatch(v, "SystemNtp-Ntpserver-Authentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemNtpNtpserverIdSna(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemNtp-Ntpserver-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := i["key"]; ok {
			v := flattenSystemNtpNtpserverKeySna(i["key"], d, pre_append)
			tmp["key"] = fortiAPISubPartPatch(v, "SystemNtp-Ntpserver-Key")
			c := d.Get(pre_append).(*schema.Set)
			if c.Len() > 0 {
				tmp["key"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_id"
		if _, ok := i["key-id"]; ok {
			v := flattenSystemNtpNtpserverKeyIdSna(i["key-id"], d, pre_append)
			tmp["key_id"] = fortiAPISubPartPatch(v, "SystemNtp-Ntpserver-KeyId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maxpoll"
		if _, ok := i["maxpoll"]; ok {
			v := flattenSystemNtpNtpserverMaxpollSna(i["maxpoll"], d, pre_append)
			tmp["maxpoll"] = fortiAPISubPartPatch(v, "SystemNtp-Ntpserver-Maxpoll")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minpoll"
		if _, ok := i["minpoll"]; ok {
			v := flattenSystemNtpNtpserverMinpollSna(i["minpoll"], d, pre_append)
			tmp["minpoll"] = fortiAPISubPartPatch(v, "SystemNtp-Ntpserver-Minpoll")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ntpv3"
		if _, ok := i["ntpv3"]; ok {
			v := flattenSystemNtpNtpserverNtpv3Sna(i["ntpv3"], d, pre_append)
			tmp["ntpv3"] = fortiAPISubPartPatch(v, "SystemNtp-Ntpserver-Ntpv3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenSystemNtpNtpserverServerSna(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "SystemNtp-Ntpserver-Server")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemNtpNtpserverAuthenticationSna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverIdSna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverKeySna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemNtpNtpserverKeyIdSna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverMaxpollSna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverMinpollSna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverNtpv3Sna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverServerSna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpStatusSna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpSyncIntervalSna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemNtp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("ntpserver", flattenSystemNtpNtpserverSna(o["ntpserver"], d, "ntpserver")); err != nil {
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
			if err = d.Set("ntpserver", flattenSystemNtpNtpserverSna(o["ntpserver"], d, "ntpserver")); err != nil {
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

	if err = d.Set("status", flattenSystemNtpStatusSna(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemNtp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("sync_interval", flattenSystemNtpSyncIntervalSna(o["sync_interval"], d, "sync_interval")); err != nil {
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

func expandSystemNtpNtpserverSna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authentication"], _ = expandSystemNtpNtpserverAuthenticationSna(d, i["authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemNtpNtpserverIdSna(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key"], _ = expandSystemNtpNtpserverKeySna(d, i["key"], pre_append)
		} else {
			tmp["key"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-id"], _ = expandSystemNtpNtpserverKeyIdSna(d, i["key_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maxpoll"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maxpoll"], _ = expandSystemNtpNtpserverMaxpollSna(d, i["maxpoll"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minpoll"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["minpoll"], _ = expandSystemNtpNtpserverMinpollSna(d, i["minpoll"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ntpv3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ntpv3"], _ = expandSystemNtpNtpserverNtpv3Sna(d, i["ntpv3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandSystemNtpNtpserverServerSna(d, i["server"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemNtpNtpserverAuthenticationSna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverIdSna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverKeySna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemNtpNtpserverKeyIdSna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverMaxpollSna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverMinpollSna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverNtpv3Sna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverServerSna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpStatusSna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpSyncIntervalSna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNtp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ntpserver"); ok || d.HasChange("ntpserver") {
		t, err := expandSystemNtpNtpserverSna(d, v, "ntpserver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntpserver"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemNtpStatusSna(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("sync_interval"); ok || d.HasChange("sync_interval") {
		t, err := expandSystemNtpSyncIntervalSna(d, v, "sync_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync_interval"] = t
		}
	}

	return &obj, nil
}
