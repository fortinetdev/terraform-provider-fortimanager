// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPv4 ranges in the disable entry.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallInternetServiceExtensionDisableEntryIpRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallInternetServiceExtensionDisableEntryIpRangeCreate,
		Read:   resourceObjectFirewallInternetServiceExtensionDisableEntryIpRangeRead,
		Update: resourceObjectFirewallInternetServiceExtensionDisableEntryIpRangeUpdate,
		Delete: resourceObjectFirewallInternetServiceExtensionDisableEntryIpRangeDelete,

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
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallInternetServiceExtensionDisableEntryIpRangeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallInternetServiceExtensionDisableEntryIpRange(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallInternetServiceExtensionDisableEntryIpRange resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectFirewallInternetServiceExtensionDisableEntryIpRange(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectFirewallInternetServiceExtensionDisableEntryIpRange(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectFirewallInternetServiceExtensionDisableEntryIpRange resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectFirewallInternetServiceExtensionDisableEntryIpRange(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectFirewallInternetServiceExtensionDisableEntryIpRange resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallInternetServiceExtensionDisableEntryIpRangeRead(d, m)
}

func resourceObjectFirewallInternetServiceExtensionDisableEntryIpRangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallInternetServiceExtensionDisableEntryIpRange(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceExtensionDisableEntryIpRange resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallInternetServiceExtensionDisableEntryIpRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceExtensionDisableEntryIpRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallInternetServiceExtensionDisableEntryIpRangeRead(d, m)
}

func resourceObjectFirewallInternetServiceExtensionDisableEntryIpRangeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallInternetServiceExtensionDisableEntryIpRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallInternetServiceExtensionDisableEntryIpRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallInternetServiceExtensionDisableEntryIpRangeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallInternetServiceExtensionDisableEntryIpRange(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceExtensionDisableEntryIpRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallInternetServiceExtensionDisableEntryIpRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceExtensionDisableEntryIpRange resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallInternetServiceExtensionDisableEntryIpRangeEndIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceExtensionDisableEntryIpRangeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceExtensionDisableEntryIpRangeStartIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallInternetServiceExtensionDisableEntryIpRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("end_ip", flattenObjectFirewallInternetServiceExtensionDisableEntryIpRangeEndIp3rdl(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "ObjectFirewallInternetServiceExtensionDisableEntryIpRange-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallInternetServiceExtensionDisableEntryIpRangeId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallInternetServiceExtensionDisableEntryIpRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenObjectFirewallInternetServiceExtensionDisableEntryIpRangeStartIp3rdl(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "ObjectFirewallInternetServiceExtensionDisableEntryIpRange-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallInternetServiceExtensionDisableEntryIpRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallInternetServiceExtensionDisableEntryIpRangeEndIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceExtensionDisableEntryIpRangeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceExtensionDisableEntryIpRangeStartIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallInternetServiceExtensionDisableEntryIpRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandObjectFirewallInternetServiceExtensionDisableEntryIpRangeEndIp3rdl(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallInternetServiceExtensionDisableEntryIpRangeId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandObjectFirewallInternetServiceExtensionDisableEntryIpRangeStartIp3rdl(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	return &obj, nil
}
