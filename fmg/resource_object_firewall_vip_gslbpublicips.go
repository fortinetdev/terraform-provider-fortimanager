// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Publicly accessible IP addresses for the FortiGSLB service.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallVipGslbPublicIps() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallVipGslbPublicIpsCreate,
		Read:   resourceObjectFirewallVipGslbPublicIpsRead,
		Update: resourceObjectFirewallVipGslbPublicIpsUpdate,
		Delete: resourceObjectFirewallVipGslbPublicIpsDelete,

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
			"vip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallVipGslbPublicIpsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	vip := d.Get("vip").(string)
	paradict["vip"] = vip

	obj, err := getObjectObjectFirewallVipGslbPublicIps(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallVipGslbPublicIps resource while getting object: %v", err)
	}

	v, err := c.CreateObjectFirewallVipGslbPublicIps(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallVipGslbPublicIps resource: %v", err)
	}

	if v != nil && v["index"] != nil {
		if vidn, ok := v["index"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceObjectFirewallVipGslbPublicIpsRead(d, m)
		} else {
			return fmt.Errorf("Error creating ObjectFirewallVipGslbPublicIps resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectFirewallVipGslbPublicIpsRead(d, m)
}

func resourceObjectFirewallVipGslbPublicIpsUpdate(d *schema.ResourceData, m interface{}) error {
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

	vip := d.Get("vip").(string)
	paradict["vip"] = vip

	obj, err := getObjectObjectFirewallVipGslbPublicIps(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallVipGslbPublicIps resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallVipGslbPublicIps(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallVipGslbPublicIps resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectFirewallVipGslbPublicIpsRead(d, m)
}

func resourceObjectFirewallVipGslbPublicIpsDelete(d *schema.ResourceData, m interface{}) error {
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

	vip := d.Get("vip").(string)
	paradict["vip"] = vip

	err = c.DeleteObjectFirewallVipGslbPublicIps(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallVipGslbPublicIps resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallVipGslbPublicIpsRead(d *schema.ResourceData, m interface{}) error {
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

	vip := d.Get("vip").(string)
	if vip == "" {
		vip = importOptionChecking(m.(*FortiClient).Cfg, "vip")
		if vip == "" {
			return fmt.Errorf("Parameter vip is missing")
		}
		if err = d.Set("vip", vip); err != nil {
			return fmt.Errorf("Error set params vip: %v", err)
		}
	}
	paradict["vip"] = vip

	o, err := c.ReadObjectFirewallVipGslbPublicIps(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallVipGslbPublicIps resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallVipGslbPublicIps(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallVipGslbPublicIps resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallVipGslbPublicIpsIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipGslbPublicIpsIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallVipGslbPublicIps(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("index", flattenObjectFirewallVipGslbPublicIpsIndex2edl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "ObjectFirewallVipGslbPublicIps-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("ip", flattenObjectFirewallVipGslbPublicIpsIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectFirewallVipGslbPublicIps-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallVipGslbPublicIpsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallVipGslbPublicIpsIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipGslbPublicIpsIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallVipGslbPublicIps(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandObjectFirewallVipGslbPublicIpsIndex2edl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandObjectFirewallVipGslbPublicIpsIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	return &obj, nil
}
