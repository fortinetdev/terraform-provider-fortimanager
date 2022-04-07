// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure NPU priority protocol.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSystemNpuPriorityProtocol() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuPriorityProtocolUpdate,
		Read:   resourceObjectSystemNpuPriorityProtocolRead,
		Update: resourceObjectSystemNpuPriorityProtocolUpdate,
		Delete: resourceObjectSystemNpuPriorityProtocolDelete,

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
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bgp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"slbc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemNpuPriorityProtocolUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemNpuPriorityProtocol(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuPriorityProtocol resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuPriorityProtocol(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuPriorityProtocol resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuPriorityProtocol")

	return resourceObjectSystemNpuPriorityProtocolRead(d, m)
}

func resourceObjectSystemNpuPriorityProtocolDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemNpuPriorityProtocol(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuPriorityProtocol resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuPriorityProtocolRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemNpuPriorityProtocol(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuPriorityProtocol resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuPriorityProtocol(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuPriorityProtocol resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuPriorityProtocolBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPriorityProtocolBgp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPriorityProtocolSlbc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuPriorityProtocol(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("bfd", flattenObjectSystemNpuPriorityProtocolBfd(o["bfd"], d, "bfd")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd"], "ObjectSystemNpuPriorityProtocol-Bfd"); ok {
			if err = d.Set("bfd", vv); err != nil {
				return fmt.Errorf("Error reading bfd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("bgp", flattenObjectSystemNpuPriorityProtocolBgp(o["bgp"], d, "bgp")); err != nil {
		if vv, ok := fortiAPIPatch(o["bgp"], "ObjectSystemNpuPriorityProtocol-Bgp"); ok {
			if err = d.Set("bgp", vv); err != nil {
				return fmt.Errorf("Error reading bgp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bgp: %v", err)
		}
	}

	if err = d.Set("slbc", flattenObjectSystemNpuPriorityProtocolSlbc(o["slbc"], d, "slbc")); err != nil {
		if vv, ok := fortiAPIPatch(o["slbc"], "ObjectSystemNpuPriorityProtocol-Slbc"); ok {
			if err = d.Set("slbc", vv); err != nil {
				return fmt.Errorf("Error reading slbc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading slbc: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuPriorityProtocolFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuPriorityProtocolBfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPriorityProtocolBgp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPriorityProtocolSlbc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuPriorityProtocol(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bfd"); ok {
		t, err := expandObjectSystemNpuPriorityProtocolBfd(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("bgp"); ok {
		t, err := expandObjectSystemNpuPriorityProtocolBgp(d, v, "bgp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bgp"] = t
		}
	}

	if v, ok := d.GetOk("slbc"); ok {
		t, err := expandObjectSystemNpuPriorityProtocolSlbc(d, v, "slbc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slbc"] = t
		}
	}

	return &obj, nil
}
