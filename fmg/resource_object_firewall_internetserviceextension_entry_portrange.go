// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Port ranges in the custom entry.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallInternetServiceExtensionEntryPortRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallInternetServiceExtensionEntryPortRangeCreate,
		Read:   resourceObjectFirewallInternetServiceExtensionEntryPortRangeRead,
		Update: resourceObjectFirewallInternetServiceExtensionEntryPortRangeUpdate,
		Delete: resourceObjectFirewallInternetServiceExtensionEntryPortRangeDelete,

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
			"entry": &schema.Schema{
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

func resourceObjectFirewallInternetServiceExtensionEntryPortRangeCreate(d *schema.ResourceData, m interface{}) error {
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
	entry := d.Get("entry").(string)
	paradict["internet_service_extension"] = internet_service_extension
	paradict["entry"] = entry

	obj, err := getObjectObjectFirewallInternetServiceExtensionEntryPortRange(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallInternetServiceExtensionEntryPortRange resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectFirewallInternetServiceExtensionEntryPortRange(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectFirewallInternetServiceExtensionEntryPortRange(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectFirewallInternetServiceExtensionEntryPortRange resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectFirewallInternetServiceExtensionEntryPortRange(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectFirewallInternetServiceExtensionEntryPortRange resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallInternetServiceExtensionEntryPortRangeRead(d, m)
}

func resourceObjectFirewallInternetServiceExtensionEntryPortRangeUpdate(d *schema.ResourceData, m interface{}) error {
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
	entry := d.Get("entry").(string)
	paradict["internet_service_extension"] = internet_service_extension
	paradict["entry"] = entry

	obj, err := getObjectObjectFirewallInternetServiceExtensionEntryPortRange(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceExtensionEntryPortRange resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallInternetServiceExtensionEntryPortRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceExtensionEntryPortRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallInternetServiceExtensionEntryPortRangeRead(d, m)
}

func resourceObjectFirewallInternetServiceExtensionEntryPortRangeDelete(d *schema.ResourceData, m interface{}) error {
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
	entry := d.Get("entry").(string)
	paradict["internet_service_extension"] = internet_service_extension
	paradict["entry"] = entry

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallInternetServiceExtensionEntryPortRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallInternetServiceExtensionEntryPortRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallInternetServiceExtensionEntryPortRangeRead(d *schema.ResourceData, m interface{}) error {
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
	entry := d.Get("entry").(string)
	if internet_service_extension == "" {
		internet_service_extension = importOptionChecking(m.(*FortiClient).Cfg, "internet_service_extension")
		if internet_service_extension == "" {
			return fmt.Errorf("Parameter internet_service_extension is missing")
		}
		if err = d.Set("internet_service_extension", internet_service_extension); err != nil {
			return fmt.Errorf("Error set params internet_service_extension: %v", err)
		}
	}
	if entry == "" {
		entry = importOptionChecking(m.(*FortiClient).Cfg, "entry")
		if entry == "" {
			return fmt.Errorf("Parameter entry is missing")
		}
		if err = d.Set("entry", entry); err != nil {
			return fmt.Errorf("Error set params entry: %v", err)
		}
	}
	paradict["internet_service_extension"] = internet_service_extension
	paradict["entry"] = entry

	o, err := c.ReadObjectFirewallInternetServiceExtensionEntryPortRange(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceExtensionEntryPortRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallInternetServiceExtensionEntryPortRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceExtensionEntryPortRange resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallInternetServiceExtensionEntryPortRangeEndPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceExtensionEntryPortRangeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceExtensionEntryPortRangeStartPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallInternetServiceExtensionEntryPortRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("end_port", flattenObjectFirewallInternetServiceExtensionEntryPortRangeEndPort3rdl(o["end-port"], d, "end_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-port"], "ObjectFirewallInternetServiceExtensionEntryPortRange-EndPort"); ok {
			if err = d.Set("end_port", vv); err != nil {
				return fmt.Errorf("Error reading end_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallInternetServiceExtensionEntryPortRangeId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallInternetServiceExtensionEntryPortRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_port", flattenObjectFirewallInternetServiceExtensionEntryPortRangeStartPort3rdl(o["start-port"], d, "start_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-port"], "ObjectFirewallInternetServiceExtensionEntryPortRange-StartPort"); ok {
			if err = d.Set("start_port", vv); err != nil {
				return fmt.Errorf("Error reading start_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallInternetServiceExtensionEntryPortRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallInternetServiceExtensionEntryPortRangeEndPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceExtensionEntryPortRangeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceExtensionEntryPortRangeStartPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallInternetServiceExtensionEntryPortRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_port"); ok || d.HasChange("end_port") {
		t, err := expandObjectFirewallInternetServiceExtensionEntryPortRangeEndPort3rdl(d, v, "end_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallInternetServiceExtensionEntryPortRangeId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_port"); ok || d.HasChange("start_port") {
		t, err := expandObjectFirewallInternetServiceExtensionEntryPortRangeStartPort3rdl(d, v, "start_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	}

	return &obj, nil
}
