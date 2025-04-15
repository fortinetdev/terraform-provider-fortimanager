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

func resourceObjectFirewallInternetServiceAdditionEntryPortRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallInternetServiceAdditionEntryPortRangeCreate,
		Read:   resourceObjectFirewallInternetServiceAdditionEntryPortRangeRead,
		Update: resourceObjectFirewallInternetServiceAdditionEntryPortRangeUpdate,
		Delete: resourceObjectFirewallInternetServiceAdditionEntryPortRangeDelete,

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
			"internet_service_addition": &schema.Schema{
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

func resourceObjectFirewallInternetServiceAdditionEntryPortRangeCreate(d *schema.ResourceData, m interface{}) error {
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

	internet_service_addition := d.Get("internet_service_addition").(string)
	entry := d.Get("entry").(string)
	paradict["internet_service_addition"] = internet_service_addition
	paradict["entry"] = entry

	obj, err := getObjectObjectFirewallInternetServiceAdditionEntryPortRange(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallInternetServiceAdditionEntryPortRange resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFirewallInternetServiceAdditionEntryPortRange(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallInternetServiceAdditionEntryPortRange resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallInternetServiceAdditionEntryPortRangeRead(d, m)
}

func resourceObjectFirewallInternetServiceAdditionEntryPortRangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	internet_service_addition := d.Get("internet_service_addition").(string)
	entry := d.Get("entry").(string)
	paradict["internet_service_addition"] = internet_service_addition
	paradict["entry"] = entry

	obj, err := getObjectObjectFirewallInternetServiceAdditionEntryPortRange(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceAdditionEntryPortRange resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallInternetServiceAdditionEntryPortRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceAdditionEntryPortRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallInternetServiceAdditionEntryPortRangeRead(d, m)
}

func resourceObjectFirewallInternetServiceAdditionEntryPortRangeDelete(d *schema.ResourceData, m interface{}) error {
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

	internet_service_addition := d.Get("internet_service_addition").(string)
	entry := d.Get("entry").(string)
	paradict["internet_service_addition"] = internet_service_addition
	paradict["entry"] = entry

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallInternetServiceAdditionEntryPortRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallInternetServiceAdditionEntryPortRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallInternetServiceAdditionEntryPortRangeRead(d *schema.ResourceData, m interface{}) error {
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

	internet_service_addition := d.Get("internet_service_addition").(string)
	entry := d.Get("entry").(string)
	if internet_service_addition == "" {
		internet_service_addition = importOptionChecking(m.(*FortiClient).Cfg, "internet_service_addition")
		if internet_service_addition == "" {
			return fmt.Errorf("Parameter internet_service_addition is missing")
		}
		if err = d.Set("internet_service_addition", internet_service_addition); err != nil {
			return fmt.Errorf("Error set params internet_service_addition: %v", err)
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
	paradict["internet_service_addition"] = internet_service_addition
	paradict["entry"] = entry

	o, err := c.ReadObjectFirewallInternetServiceAdditionEntryPortRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceAdditionEntryPortRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallInternetServiceAdditionEntryPortRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceAdditionEntryPortRange resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallInternetServiceAdditionEntryPortRangeEndPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceAdditionEntryPortRangeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceAdditionEntryPortRangeStartPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallInternetServiceAdditionEntryPortRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("end_port", flattenObjectFirewallInternetServiceAdditionEntryPortRangeEndPort3rdl(o["end-port"], d, "end_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-port"], "ObjectFirewallInternetServiceAdditionEntryPortRange-EndPort"); ok {
			if err = d.Set("end_port", vv); err != nil {
				return fmt.Errorf("Error reading end_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallInternetServiceAdditionEntryPortRangeId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallInternetServiceAdditionEntryPortRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_port", flattenObjectFirewallInternetServiceAdditionEntryPortRangeStartPort3rdl(o["start-port"], d, "start_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-port"], "ObjectFirewallInternetServiceAdditionEntryPortRange-StartPort"); ok {
			if err = d.Set("start_port", vv); err != nil {
				return fmt.Errorf("Error reading start_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallInternetServiceAdditionEntryPortRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallInternetServiceAdditionEntryPortRangeEndPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceAdditionEntryPortRangeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceAdditionEntryPortRangeStartPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallInternetServiceAdditionEntryPortRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_port"); ok || d.HasChange("end_port") {
		t, err := expandObjectFirewallInternetServiceAdditionEntryPortRangeEndPort3rdl(d, v, "end_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallInternetServiceAdditionEntryPortRangeId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_port"); ok || d.HasChange("start_port") {
		t, err := expandObjectFirewallInternetServiceAdditionEntryPortRangeStartPort3rdl(d, v, "start_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	}

	return &obj, nil
}
