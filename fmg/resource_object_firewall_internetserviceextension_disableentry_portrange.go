// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Port ranges in the disable entry.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallInternetServiceExtensionDisableEntryPortRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallInternetServiceExtensionDisableEntryPortRangeCreate,
		Read:   resourceObjectFirewallInternetServiceExtensionDisableEntryPortRangeRead,
		Update: resourceObjectFirewallInternetServiceExtensionDisableEntryPortRangeUpdate,
		Delete: resourceObjectFirewallInternetServiceExtensionDisableEntryPortRangeDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"internet_service_extension": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"disable_entry": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"end_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"start_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallInternetServiceExtensionDisableEntryPortRangeCreate(d *schema.ResourceData, m interface{}) error {
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

	internet_service_extension := d.Get("internet_service_extension").(string)
	disable_entry := d.Get("disable_entry").(string)
	paradict["internet_service_extension"] = internet_service_extension
	paradict["disable_entry"] = disable_entry

	obj, err := getObjectObjectFirewallInternetServiceExtensionDisableEntryPortRange(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallInternetServiceExtensionDisableEntryPortRange resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectFirewallInternetServiceExtensionDisableEntryPortRange(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectFirewallInternetServiceExtensionDisableEntryPortRange(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectFirewallInternetServiceExtensionDisableEntryPortRange resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectFirewallInternetServiceExtensionDisableEntryPortRange(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectFirewallInternetServiceExtensionDisableEntryPortRange resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallInternetServiceExtensionDisableEntryPortRangeRead(d, m)
}

func resourceObjectFirewallInternetServiceExtensionDisableEntryPortRangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	internet_service_extension := d.Get("internet_service_extension").(string)
	disable_entry := d.Get("disable_entry").(string)
	paradict["internet_service_extension"] = internet_service_extension
	paradict["disable_entry"] = disable_entry

	obj, err := getObjectObjectFirewallInternetServiceExtensionDisableEntryPortRange(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceExtensionDisableEntryPortRange resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallInternetServiceExtensionDisableEntryPortRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceExtensionDisableEntryPortRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallInternetServiceExtensionDisableEntryPortRangeRead(d, m)
}

func resourceObjectFirewallInternetServiceExtensionDisableEntryPortRangeDelete(d *schema.ResourceData, m interface{}) error {
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

	internet_service_extension := d.Get("internet_service_extension").(string)
	disable_entry := d.Get("disable_entry").(string)
	paradict["internet_service_extension"] = internet_service_extension
	paradict["disable_entry"] = disable_entry

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallInternetServiceExtensionDisableEntryPortRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallInternetServiceExtensionDisableEntryPortRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallInternetServiceExtensionDisableEntryPortRangeRead(d *schema.ResourceData, m interface{}) error {
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

	internet_service_extension := d.Get("internet_service_extension").(string)
	disable_entry := d.Get("disable_entry").(string)
	if internet_service_extension == "" {
		internet_service_extension = importOptionChecking(m.(*FortiClient).Cfg, "internet_service_extension")
		if internet_service_extension == "" {
			return fmt.Errorf("Parameter internet_service_extension is missing")
		}
		if err = d.Set("internet_service_extension", internet_service_extension); err != nil {
			return fmt.Errorf("Error set params internet_service_extension: %v", err)
		}
	}
	if disable_entry == "" {
		disable_entry = importOptionChecking(m.(*FortiClient).Cfg, "disable_entry")
		if disable_entry == "" {
			return fmt.Errorf("Parameter disable_entry is missing")
		}
		if err = d.Set("disable_entry", disable_entry); err != nil {
			return fmt.Errorf("Error set params disable_entry: %v", err)
		}
	}
	paradict["internet_service_extension"] = internet_service_extension
	paradict["disable_entry"] = disable_entry

	o, err := c.ReadObjectFirewallInternetServiceExtensionDisableEntryPortRange(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceExtensionDisableEntryPortRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallInternetServiceExtensionDisableEntryPortRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceExtensionDisableEntryPortRange resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallInternetServiceExtensionDisableEntryPortRangeEndPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceExtensionDisableEntryPortRangeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceExtensionDisableEntryPortRangeStartPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallInternetServiceExtensionDisableEntryPortRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("end_port", flattenObjectFirewallInternetServiceExtensionDisableEntryPortRangeEndPort3rdl(o["end-port"], d, "end_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-port"], "ObjectFirewallInternetServiceExtensionDisableEntryPortRange-EndPort"); ok {
			if err = d.Set("end_port", vv); err != nil {
				return fmt.Errorf("Error reading end_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallInternetServiceExtensionDisableEntryPortRangeId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallInternetServiceExtensionDisableEntryPortRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_port", flattenObjectFirewallInternetServiceExtensionDisableEntryPortRangeStartPort3rdl(o["start-port"], d, "start_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-port"], "ObjectFirewallInternetServiceExtensionDisableEntryPortRange-StartPort"); ok {
			if err = d.Set("start_port", vv); err != nil {
				return fmt.Errorf("Error reading start_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallInternetServiceExtensionDisableEntryPortRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallInternetServiceExtensionDisableEntryPortRangeEndPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceExtensionDisableEntryPortRangeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceExtensionDisableEntryPortRangeStartPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallInternetServiceExtensionDisableEntryPortRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_port"); ok || d.HasChange("end_port") {
		t, err := expandObjectFirewallInternetServiceExtensionDisableEntryPortRangeEndPort3rdl(d, v, "end_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallInternetServiceExtensionDisableEntryPortRangeId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_port"); ok || d.HasChange("start_port") {
		t, err := expandObjectFirewallInternetServiceExtensionDisableEntryPortRangeStartPort3rdl(d, v, "start_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	}

	return &obj, nil
}
