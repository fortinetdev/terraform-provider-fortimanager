// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure AntiVirus quarantine settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectAntivirusProfileNacQuar() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectAntivirusProfileNacQuarUpdate,
		Read:   resourceObjectAntivirusProfileNacQuarRead,
		Update: resourceObjectAntivirusProfileNacQuarUpdate,
		Delete: resourceObjectAntivirusProfileNacQuarDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"expiry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"infected": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectAntivirusProfileNacQuarUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectAntivirusProfileNacQuar(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectAntivirusProfileNacQuar resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectAntivirusProfileNacQuar(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectAntivirusProfileNacQuar resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectAntivirusProfileNacQuar")

	return resourceObjectAntivirusProfileNacQuarRead(d, m)
}

func resourceObjectAntivirusProfileNacQuarDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	err = c.DeleteObjectAntivirusProfileNacQuar(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectAntivirusProfileNacQuar resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectAntivirusProfileNacQuarRead(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadObjectAntivirusProfileNacQuar(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectAntivirusProfileNacQuar resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectAntivirusProfileNacQuar(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectAntivirusProfileNacQuar resource from API: %v", err)
	}
	return nil
}

func flattenObjectAntivirusProfileNacQuarExpiry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileNacQuarInfected2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileNacQuarLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectAntivirusProfileNacQuar(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("expiry", flattenObjectAntivirusProfileNacQuarExpiry2edl(o["expiry"], d, "expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["expiry"], "ObjectAntivirusProfileNacQuar-Expiry"); ok {
			if err = d.Set("expiry", vv); err != nil {
				return fmt.Errorf("Error reading expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expiry: %v", err)
		}
	}

	if err = d.Set("infected", flattenObjectAntivirusProfileNacQuarInfected2edl(o["infected"], d, "infected")); err != nil {
		if vv, ok := fortiAPIPatch(o["infected"], "ObjectAntivirusProfileNacQuar-Infected"); ok {
			if err = d.Set("infected", vv); err != nil {
				return fmt.Errorf("Error reading infected: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading infected: %v", err)
		}
	}

	if err = d.Set("log", flattenObjectAntivirusProfileNacQuarLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ObjectAntivirusProfileNacQuar-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	return nil
}

func flattenObjectAntivirusProfileNacQuarFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectAntivirusProfileNacQuarExpiry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileNacQuarInfected2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileNacQuarLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectAntivirusProfileNacQuar(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("expiry"); ok || d.HasChange("expiry") {
		t, err := expandObjectAntivirusProfileNacQuarExpiry2edl(d, v, "expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expiry"] = t
		}
	}

	if v, ok := d.GetOk("infected"); ok || d.HasChange("infected") {
		t, err := expandObjectAntivirusProfileNacQuarInfected2edl(d, v, "infected")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["infected"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandObjectAntivirusProfileNacQuarLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	return &obj, nil
}
