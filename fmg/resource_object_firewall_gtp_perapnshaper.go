// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Per APN shaper.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallGtpPerApnShaper() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallGtpPerApnShaperCreate,
		Read:   resourceObjectFirewallGtpPerApnShaperRead,
		Update: resourceObjectFirewallGtpPerApnShaperUpdate,
		Delete: resourceObjectFirewallGtpPerApnShaperDelete,

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
			"gtp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"apn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"rate_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallGtpPerApnShaperCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	gtp := d.Get("gtp").(string)
	paradict["gtp"] = gtp

	obj, err := getObjectObjectFirewallGtpPerApnShaper(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallGtpPerApnShaper resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallGtpPerApnShaper(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallGtpPerApnShaper resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallGtpPerApnShaperRead(d, m)
}

func resourceObjectFirewallGtpPerApnShaperUpdate(d *schema.ResourceData, m interface{}) error {
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

	gtp := d.Get("gtp").(string)
	paradict["gtp"] = gtp

	obj, err := getObjectObjectFirewallGtpPerApnShaper(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpPerApnShaper resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallGtpPerApnShaper(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpPerApnShaper resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallGtpPerApnShaperRead(d, m)
}

func resourceObjectFirewallGtpPerApnShaperDelete(d *schema.ResourceData, m interface{}) error {
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

	gtp := d.Get("gtp").(string)
	paradict["gtp"] = gtp

	err = c.DeleteObjectFirewallGtpPerApnShaper(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallGtpPerApnShaper resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallGtpPerApnShaperRead(d *schema.ResourceData, m interface{}) error {
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

	gtp := d.Get("gtp").(string)
	if gtp == "" {
		gtp = importOptionChecking(m.(*FortiClient).Cfg, "gtp")
		if gtp == "" {
			return fmt.Errorf("Parameter gtp is missing")
		}
		if err = d.Set("gtp", gtp); err != nil {
			return fmt.Errorf("Error set params gtp: %v", err)
		}
	}
	paradict["gtp"] = gtp

	o, err := c.ReadObjectFirewallGtpPerApnShaper(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpPerApnShaper resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallGtpPerApnShaper(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpPerApnShaper resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallGtpPerApnShaperApn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPerApnShaperId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPerApnShaperRateLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpPerApnShaperVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallGtpPerApnShaper(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("apn", flattenObjectFirewallGtpPerApnShaperApn2edl(o["apn"], d, "apn")); err != nil {
		if vv, ok := fortiAPIPatch(o["apn"], "ObjectFirewallGtpPerApnShaper-Apn"); ok {
			if err = d.Set("apn", vv); err != nil {
				return fmt.Errorf("Error reading apn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apn: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallGtpPerApnShaperId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallGtpPerApnShaper-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("rate_limit", flattenObjectFirewallGtpPerApnShaperRateLimit2edl(o["rate-limit"], d, "rate_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-limit"], "ObjectFirewallGtpPerApnShaper-RateLimit"); ok {
			if err = d.Set("rate_limit", vv); err != nil {
				return fmt.Errorf("Error reading rate_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_limit: %v", err)
		}
	}

	if err = d.Set("version", flattenObjectFirewallGtpPerApnShaperVersion2edl(o["version"], d, "version")); err != nil {
		if vv, ok := fortiAPIPatch(o["version"], "ObjectFirewallGtpPerApnShaper-Version"); ok {
			if err = d.Set("version", vv); err != nil {
				return fmt.Errorf("Error reading version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallGtpPerApnShaperFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallGtpPerApnShaperApn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPerApnShaperId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPerApnShaperRateLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpPerApnShaperVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallGtpPerApnShaper(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("apn"); ok || d.HasChange("apn") {
		t, err := expandObjectFirewallGtpPerApnShaperApn2edl(d, v, "apn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apn"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallGtpPerApnShaperId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("rate_limit"); ok || d.HasChange("rate_limit") {
		t, err := expandObjectFirewallGtpPerApnShaperRateLimit2edl(d, v, "rate_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-limit"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandObjectFirewallGtpPerApnShaperVersion2edl(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	return &obj, nil
}
