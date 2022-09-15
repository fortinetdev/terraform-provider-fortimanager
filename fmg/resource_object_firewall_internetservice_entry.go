// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Entries in the Internet Service database.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallInternetServiceEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallInternetServiceEntryCreate,
		Read:   resourceObjectFirewallInternetServiceEntryRead,
		Update: resourceObjectFirewallInternetServiceEntryUpdate,
		Delete: resourceObjectFirewallInternetServiceEntryDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip_range_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallInternetServiceEntryCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallInternetServiceEntry(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallInternetServiceEntry resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallInternetServiceEntry(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallInternetServiceEntry resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallInternetServiceEntryRead(d, m)
}

func resourceObjectFirewallInternetServiceEntryUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallInternetServiceEntry(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceEntry resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallInternetServiceEntry(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceEntry resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallInternetServiceEntryRead(d, m)
}

func resourceObjectFirewallInternetServiceEntryDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallInternetServiceEntry(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallInternetServiceEntry resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallInternetServiceEntryRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallInternetServiceEntry(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceEntry resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallInternetServiceEntry(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceEntry resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallInternetServiceEntryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceEntryIpNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceEntryIpRangeNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceEntryPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallInternetServiceEntryProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallInternetServiceEntry(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenObjectFirewallInternetServiceEntryId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallInternetServiceEntry-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_number", flattenObjectFirewallInternetServiceEntryIpNumber(o["ip-number"], d, "ip_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-number"], "ObjectFirewallInternetServiceEntry-IpNumber"); ok {
			if err = d.Set("ip_number", vv); err != nil {
				return fmt.Errorf("Error reading ip_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_number: %v", err)
		}
	}

	if err = d.Set("ip_range_number", flattenObjectFirewallInternetServiceEntryIpRangeNumber(o["ip-range-number"], d, "ip_range_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-range-number"], "ObjectFirewallInternetServiceEntry-IpRangeNumber"); ok {
			if err = d.Set("ip_range_number", vv); err != nil {
				return fmt.Errorf("Error reading ip_range_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_range_number: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectFirewallInternetServiceEntryPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectFirewallInternetServiceEntry-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("protocol", flattenObjectFirewallInternetServiceEntryProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ObjectFirewallInternetServiceEntry-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallInternetServiceEntryFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallInternetServiceEntryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceEntryIpNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceEntryIpRangeNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceEntryPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallInternetServiceEntryProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallInternetServiceEntry(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandObjectFirewallInternetServiceEntryId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_number"); ok || d.HasChange("ip_number") {
		t, err := expandObjectFirewallInternetServiceEntryIpNumber(d, v, "ip_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-number"] = t
		}
	}

	if v, ok := d.GetOk("ip_range_number"); ok || d.HasChange("ip_range_number") {
		t, err := expandObjectFirewallInternetServiceEntryIpRangeNumber(d, v, "ip_range_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range-number"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectFirewallInternetServiceEntryPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandObjectFirewallInternetServiceEntryProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	return &obj, nil
}
