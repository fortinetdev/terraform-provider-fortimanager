// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPv6 ranges in the disable entry.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallInternetServiceExtensionDisableEntryIp6Range() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallInternetServiceExtensionDisableEntryIp6RangeCreate,
		Read:   resourceObjectFirewallInternetServiceExtensionDisableEntryIp6RangeRead,
		Update: resourceObjectFirewallInternetServiceExtensionDisableEntryIp6RangeUpdate,
		Delete: resourceObjectFirewallInternetServiceExtensionDisableEntryIp6RangeDelete,

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
			"end_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"start_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallInternetServiceExtensionDisableEntryIp6RangeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallInternetServiceExtensionDisableEntryIp6Range(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallInternetServiceExtensionDisableEntryIp6Range resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectFirewallInternetServiceExtensionDisableEntryIp6Range(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectFirewallInternetServiceExtensionDisableEntryIp6Range(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectFirewallInternetServiceExtensionDisableEntryIp6Range resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectFirewallInternetServiceExtensionDisableEntryIp6Range(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectFirewallInternetServiceExtensionDisableEntryIp6Range resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallInternetServiceExtensionDisableEntryIp6RangeRead(d, m)
}

func resourceObjectFirewallInternetServiceExtensionDisableEntryIp6RangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallInternetServiceExtensionDisableEntryIp6Range(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceExtensionDisableEntryIp6Range resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallInternetServiceExtensionDisableEntryIp6Range(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceExtensionDisableEntryIp6Range resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallInternetServiceExtensionDisableEntryIp6RangeRead(d, m)
}

func resourceObjectFirewallInternetServiceExtensionDisableEntryIp6RangeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallInternetServiceExtensionDisableEntryIp6Range(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallInternetServiceExtensionDisableEntryIp6Range resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallInternetServiceExtensionDisableEntryIp6RangeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallInternetServiceExtensionDisableEntryIp6Range(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceExtensionDisableEntryIp6Range resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallInternetServiceExtensionDisableEntryIp6Range(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceExtensionDisableEntryIp6Range resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallInternetServiceExtensionDisableEntryIp6RangeEndIp63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceExtensionDisableEntryIp6RangeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceExtensionDisableEntryIp6RangeStartIp63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallInternetServiceExtensionDisableEntryIp6Range(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("end_ip6", flattenObjectFirewallInternetServiceExtensionDisableEntryIp6RangeEndIp63rdl(o["end-ip6"], d, "end_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip6"], "ObjectFirewallInternetServiceExtensionDisableEntryIp6Range-EndIp6"); ok {
			if err = d.Set("end_ip6", vv); err != nil {
				return fmt.Errorf("Error reading end_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip6: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallInternetServiceExtensionDisableEntryIp6RangeId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallInternetServiceExtensionDisableEntryIp6Range-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_ip6", flattenObjectFirewallInternetServiceExtensionDisableEntryIp6RangeStartIp63rdl(o["start-ip6"], d, "start_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip6"], "ObjectFirewallInternetServiceExtensionDisableEntryIp6Range-StartIp6"); ok {
			if err = d.Set("start_ip6", vv); err != nil {
				return fmt.Errorf("Error reading start_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip6: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallInternetServiceExtensionDisableEntryIp6RangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallInternetServiceExtensionDisableEntryIp6RangeEndIp63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceExtensionDisableEntryIp6RangeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceExtensionDisableEntryIp6RangeStartIp63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallInternetServiceExtensionDisableEntryIp6Range(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_ip6"); ok || d.HasChange("end_ip6") {
		t, err := expandObjectFirewallInternetServiceExtensionDisableEntryIp6RangeEndIp63rdl(d, v, "end_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip6"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallInternetServiceExtensionDisableEntryIp6RangeId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_ip6"); ok || d.HasChange("start_ip6") {
		t, err := expandObjectFirewallInternetServiceExtensionDisableEntryIp6RangeStartIp63rdl(d, v, "start_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip6"] = t
		}
	}

	return &obj, nil
}
