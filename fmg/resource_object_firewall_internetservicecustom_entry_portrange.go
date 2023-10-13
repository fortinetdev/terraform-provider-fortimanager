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

func resourceObjectFirewallInternetServiceCustomEntryPortRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallInternetServiceCustomEntryPortRangeCreate,
		Read:   resourceObjectFirewallInternetServiceCustomEntryPortRangeRead,
		Update: resourceObjectFirewallInternetServiceCustomEntryPortRangeUpdate,
		Delete: resourceObjectFirewallInternetServiceCustomEntryPortRangeDelete,

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
			"internet_service_custom": &schema.Schema{
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
				Computed: true,
			},
			"start_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallInternetServiceCustomEntryPortRangeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	internet_service_custom := d.Get("internet_service_custom").(string)
	entry := d.Get("entry").(string)
	paradict["internet_service_custom"] = internet_service_custom
	paradict["entry"] = entry

	obj, err := getObjectObjectFirewallInternetServiceCustomEntryPortRange(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallInternetServiceCustomEntryPortRange resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallInternetServiceCustomEntryPortRange(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallInternetServiceCustomEntryPortRange resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallInternetServiceCustomEntryPortRangeRead(d, m)
}

func resourceObjectFirewallInternetServiceCustomEntryPortRangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	internet_service_custom := d.Get("internet_service_custom").(string)
	entry := d.Get("entry").(string)
	paradict["internet_service_custom"] = internet_service_custom
	paradict["entry"] = entry

	obj, err := getObjectObjectFirewallInternetServiceCustomEntryPortRange(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceCustomEntryPortRange resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallInternetServiceCustomEntryPortRange(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceCustomEntryPortRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallInternetServiceCustomEntryPortRangeRead(d, m)
}

func resourceObjectFirewallInternetServiceCustomEntryPortRangeDelete(d *schema.ResourceData, m interface{}) error {
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

	internet_service_custom := d.Get("internet_service_custom").(string)
	entry := d.Get("entry").(string)
	paradict["internet_service_custom"] = internet_service_custom
	paradict["entry"] = entry

	err = c.DeleteObjectFirewallInternetServiceCustomEntryPortRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallInternetServiceCustomEntryPortRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallInternetServiceCustomEntryPortRangeRead(d *schema.ResourceData, m interface{}) error {
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

	internet_service_custom := d.Get("internet_service_custom").(string)
	entry := d.Get("entry").(string)
	if internet_service_custom == "" {
		internet_service_custom = importOptionChecking(m.(*FortiClient).Cfg, "internet_service_custom")
		if internet_service_custom == "" {
			return fmt.Errorf("Parameter internet_service_custom is missing")
		}
		if err = d.Set("internet_service_custom", internet_service_custom); err != nil {
			return fmt.Errorf("Error set params internet_service_custom: %v", err)
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
	paradict["internet_service_custom"] = internet_service_custom
	paradict["entry"] = entry

	o, err := c.ReadObjectFirewallInternetServiceCustomEntryPortRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceCustomEntryPortRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallInternetServiceCustomEntryPortRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceCustomEntryPortRange resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallInternetServiceCustomEntryPortRangeEndPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceCustomEntryPortRangeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceCustomEntryPortRangeStartPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallInternetServiceCustomEntryPortRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("end_port", flattenObjectFirewallInternetServiceCustomEntryPortRangeEndPort3rdl(o["end-port"], d, "end_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-port"], "ObjectFirewallInternetServiceCustomEntryPortRange-EndPort"); ok {
			if err = d.Set("end_port", vv); err != nil {
				return fmt.Errorf("Error reading end_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallInternetServiceCustomEntryPortRangeId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallInternetServiceCustomEntryPortRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_port", flattenObjectFirewallInternetServiceCustomEntryPortRangeStartPort3rdl(o["start-port"], d, "start_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-port"], "ObjectFirewallInternetServiceCustomEntryPortRange-StartPort"); ok {
			if err = d.Set("start_port", vv); err != nil {
				return fmt.Errorf("Error reading start_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallInternetServiceCustomEntryPortRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallInternetServiceCustomEntryPortRangeEndPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceCustomEntryPortRangeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceCustomEntryPortRangeStartPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallInternetServiceCustomEntryPortRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_port"); ok || d.HasChange("end_port") {
		t, err := expandObjectFirewallInternetServiceCustomEntryPortRangeEndPort3rdl(d, v, "end_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallInternetServiceCustomEntryPortRangeId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_port"); ok || d.HasChange("start_port") {
		t, err := expandObjectFirewallInternetServiceCustomEntryPortRangeStartPort3rdl(d, v, "start_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	}

	return &obj, nil
}
