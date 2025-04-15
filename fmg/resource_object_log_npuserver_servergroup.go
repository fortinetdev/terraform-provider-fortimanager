// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: create server group.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectLogNpuServerServerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectLogNpuServerServerGroupCreate,
		Read:   resourceObjectLogNpuServerServerGroupRead,
		Update: resourceObjectLogNpuServerServerGroupUpdate,
		Delete: resourceObjectLogNpuServerServerGroupDelete,

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
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"log_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_gen_event": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_tx_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_user_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_start_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sw_log_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectLogNpuServerServerGroupCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectLogNpuServerServerGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectLogNpuServerServerGroup resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectLogNpuServerServerGroup(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectLogNpuServerServerGroup resource: %v", err)
	}

	d.SetId(getStringKey(d, "group_name"))

	return resourceObjectLogNpuServerServerGroupRead(d, m)
}

func resourceObjectLogNpuServerServerGroupUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectLogNpuServerServerGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectLogNpuServerServerGroup resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectLogNpuServerServerGroup(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectLogNpuServerServerGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "group_name"))

	return resourceObjectLogNpuServerServerGroupRead(d, m)
}

func resourceObjectLogNpuServerServerGroupDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectLogNpuServerServerGroup(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectLogNpuServerServerGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectLogNpuServerServerGroupRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectLogNpuServerServerGroup(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectLogNpuServerServerGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectLogNpuServerServerGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectLogNpuServerServerGroup resource from API: %v", err)
	}
	return nil
}

func flattenObjectLogNpuServerServerGroupGroupName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupLogFormat2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupLogGenEvent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupLogMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupLogTxMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupLogUserInfo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupServerNumber2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupServerStartId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupSwLogFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectLogNpuServerServerGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("group_name", flattenObjectLogNpuServerServerGroupGroupName2edl(o["group-name"], d, "group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-name"], "ObjectLogNpuServerServerGroup-GroupName"); ok {
			if err = d.Set("group_name", vv); err != nil {
				return fmt.Errorf("Error reading group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("log_format", flattenObjectLogNpuServerServerGroupLogFormat2edl(o["log-format"], d, "log_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-format"], "ObjectLogNpuServerServerGroup-LogFormat"); ok {
			if err = d.Set("log_format", vv); err != nil {
				return fmt.Errorf("Error reading log_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_format: %v", err)
		}
	}

	if err = d.Set("log_gen_event", flattenObjectLogNpuServerServerGroupLogGenEvent2edl(o["log-gen-event"], d, "log_gen_event")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-gen-event"], "ObjectLogNpuServerServerGroup-LogGenEvent"); ok {
			if err = d.Set("log_gen_event", vv); err != nil {
				return fmt.Errorf("Error reading log_gen_event: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_gen_event: %v", err)
		}
	}

	if err = d.Set("log_mode", flattenObjectLogNpuServerServerGroupLogMode2edl(o["log-mode"], d, "log_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-mode"], "ObjectLogNpuServerServerGroup-LogMode"); ok {
			if err = d.Set("log_mode", vv); err != nil {
				return fmt.Errorf("Error reading log_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_mode: %v", err)
		}
	}

	if err = d.Set("log_tx_mode", flattenObjectLogNpuServerServerGroupLogTxMode2edl(o["log-tx-mode"], d, "log_tx_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-tx-mode"], "ObjectLogNpuServerServerGroup-LogTxMode"); ok {
			if err = d.Set("log_tx_mode", vv); err != nil {
				return fmt.Errorf("Error reading log_tx_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_tx_mode: %v", err)
		}
	}

	if err = d.Set("log_user_info", flattenObjectLogNpuServerServerGroupLogUserInfo2edl(o["log-user-info"], d, "log_user_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-user-info"], "ObjectLogNpuServerServerGroup-LogUserInfo"); ok {
			if err = d.Set("log_user_info", vv); err != nil {
				return fmt.Errorf("Error reading log_user_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_user_info: %v", err)
		}
	}

	if err = d.Set("server_number", flattenObjectLogNpuServerServerGroupServerNumber2edl(o["server-number"], d, "server_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-number"], "ObjectLogNpuServerServerGroup-ServerNumber"); ok {
			if err = d.Set("server_number", vv); err != nil {
				return fmt.Errorf("Error reading server_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_number: %v", err)
		}
	}

	if err = d.Set("server_start_id", flattenObjectLogNpuServerServerGroupServerStartId2edl(o["server-start-id"], d, "server_start_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-start-id"], "ObjectLogNpuServerServerGroup-ServerStartId"); ok {
			if err = d.Set("server_start_id", vv); err != nil {
				return fmt.Errorf("Error reading server_start_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_start_id: %v", err)
		}
	}

	if err = d.Set("sw_log_flags", flattenObjectLogNpuServerServerGroupSwLogFlags2edl(o["sw-log-flags"], d, "sw_log_flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["sw-log-flags"], "ObjectLogNpuServerServerGroup-SwLogFlags"); ok {
			if err = d.Set("sw_log_flags", vv); err != nil {
				return fmt.Errorf("Error reading sw_log_flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sw_log_flags: %v", err)
		}
	}

	return nil
}

func flattenObjectLogNpuServerServerGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectLogNpuServerServerGroupGroupName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupLogFormat2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupLogGenEvent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupLogMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupLogTxMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupLogUserInfo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupServerNumber2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupServerStartId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupSwLogFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectLogNpuServerServerGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("group_name"); ok || d.HasChange("group_name") {
		t, err := expandObjectLogNpuServerServerGroupGroupName2edl(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("log_format"); ok || d.HasChange("log_format") {
		t, err := expandObjectLogNpuServerServerGroupLogFormat2edl(d, v, "log_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-format"] = t
		}
	}

	if v, ok := d.GetOk("log_gen_event"); ok || d.HasChange("log_gen_event") {
		t, err := expandObjectLogNpuServerServerGroupLogGenEvent2edl(d, v, "log_gen_event")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-gen-event"] = t
		}
	}

	if v, ok := d.GetOk("log_mode"); ok || d.HasChange("log_mode") {
		t, err := expandObjectLogNpuServerServerGroupLogMode2edl(d, v, "log_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-mode"] = t
		}
	}

	if v, ok := d.GetOk("log_tx_mode"); ok || d.HasChange("log_tx_mode") {
		t, err := expandObjectLogNpuServerServerGroupLogTxMode2edl(d, v, "log_tx_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-tx-mode"] = t
		}
	}

	if v, ok := d.GetOk("log_user_info"); ok || d.HasChange("log_user_info") {
		t, err := expandObjectLogNpuServerServerGroupLogUserInfo2edl(d, v, "log_user_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-user-info"] = t
		}
	}

	if v, ok := d.GetOk("server_number"); ok || d.HasChange("server_number") {
		t, err := expandObjectLogNpuServerServerGroupServerNumber2edl(d, v, "server_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-number"] = t
		}
	}

	if v, ok := d.GetOk("server_start_id"); ok || d.HasChange("server_start_id") {
		t, err := expandObjectLogNpuServerServerGroupServerStartId2edl(d, v, "server_start_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-start-id"] = t
		}
	}

	if v, ok := d.GetOk("sw_log_flags"); ok || d.HasChange("sw_log_flags") {
		t, err := expandObjectLogNpuServerServerGroupSwLogFlags2edl(d, v, "sw_log_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sw-log-flags"] = t
		}
	}

	return &obj, nil
}
