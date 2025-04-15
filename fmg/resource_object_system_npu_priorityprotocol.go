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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSystemNpuPriorityProtocol(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuPriorityProtocol resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemNpuPriorityProtocol(obj, mkey, paradict, wsParams)
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteObjectSystemNpuPriorityProtocol(mkey, paradict, wsParams)
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	o, err := c.ReadObjectSystemNpuPriorityProtocol(mkey, paradict)
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

func flattenObjectSystemNpuPriorityProtocolBfd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPriorityProtocolBgp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPriorityProtocolSlbc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuPriorityProtocol(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("bfd", flattenObjectSystemNpuPriorityProtocolBfd2edl(o["bfd"], d, "bfd")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd"], "ObjectSystemNpuPriorityProtocol-Bfd"); ok {
			if err = d.Set("bfd", vv); err != nil {
				return fmt.Errorf("Error reading bfd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("bgp", flattenObjectSystemNpuPriorityProtocolBgp2edl(o["bgp"], d, "bgp")); err != nil {
		if vv, ok := fortiAPIPatch(o["bgp"], "ObjectSystemNpuPriorityProtocol-Bgp"); ok {
			if err = d.Set("bgp", vv); err != nil {
				return fmt.Errorf("Error reading bgp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bgp: %v", err)
		}
	}

	if err = d.Set("slbc", flattenObjectSystemNpuPriorityProtocolSlbc2edl(o["slbc"], d, "slbc")); err != nil {
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

func expandObjectSystemNpuPriorityProtocolBfd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPriorityProtocolBgp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPriorityProtocolSlbc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuPriorityProtocol(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bfd"); ok || d.HasChange("bfd") {
		t, err := expandObjectSystemNpuPriorityProtocolBfd2edl(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("bgp"); ok || d.HasChange("bgp") {
		t, err := expandObjectSystemNpuPriorityProtocolBgp2edl(d, v, "bgp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bgp"] = t
		}
	}

	if v, ok := d.GetOk("slbc"); ok || d.HasChange("slbc") {
		t, err := expandObjectSystemNpuPriorityProtocolSlbc2edl(d, v, "slbc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slbc"] = t
		}
	}

	return &obj, nil
}
