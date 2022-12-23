// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure SD-WAN zones.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWantempSystemSdwanZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceWantempSystemSdwanZoneCreate,
		Read:   resourceWantempSystemSdwanZoneRead,
		Update: resourceWantempSystemSdwanZoneUpdate,
		Delete: resourceWantempSystemSdwanZoneDelete,

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
			"wanprof": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"service_sla_tie_break": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWantempSystemSdwanZoneCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wanprof := d.Get("wanprof").(string)
	paradict["wanprof"] = wanprof

	obj, err := getObjectWantempSystemSdwanZone(d)
	if err != nil {
		return fmt.Errorf("Error creating WantempSystemSdwanZone resource while getting object: %v", err)
	}

	_, err = c.CreateWantempSystemSdwanZone(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WantempSystemSdwanZone resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWantempSystemSdwanZoneRead(d, m)
}

func resourceWantempSystemSdwanZoneUpdate(d *schema.ResourceData, m interface{}) error {
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

	wanprof := d.Get("wanprof").(string)
	paradict["wanprof"] = wanprof

	obj, err := getObjectWantempSystemSdwanZone(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwanZone resource while getting object: %v", err)
	}

	_, err = c.UpdateWantempSystemSdwanZone(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwanZone resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWantempSystemSdwanZoneRead(d, m)
}

func resourceWantempSystemSdwanZoneDelete(d *schema.ResourceData, m interface{}) error {
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

	wanprof := d.Get("wanprof").(string)
	paradict["wanprof"] = wanprof

	err = c.DeleteWantempSystemSdwanZone(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WantempSystemSdwanZone resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWantempSystemSdwanZoneRead(d *schema.ResourceData, m interface{}) error {
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

	wanprof := d.Get("wanprof").(string)
	if wanprof == "" {
		wanprof = importOptionChecking(m.(*FortiClient).Cfg, "wanprof")
		if err = d.Set("wanprof", wanprof); err != nil {
			return fmt.Errorf("Error set params wanprof: %v", err)
		}
	}
	paradict["wanprof"] = wanprof

	o, err := c.ReadWantempSystemSdwanZone(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwanZone resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWantempSystemSdwanZone(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwanZone resource from API: %v", err)
	}
	return nil
}

func flattenWantempSystemSdwanZoneName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanZoneServiceSlaTieBreak(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWantempSystemSdwanZone(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenWantempSystemSdwanZoneName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WantempSystemSdwanZone-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("service_sla_tie_break", flattenWantempSystemSdwanZoneServiceSlaTieBreak(o["service-sla-tie-break"], d, "service_sla_tie_break")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-sla-tie-break"], "WantempSystemSdwanZone-ServiceSlaTieBreak"); ok {
			if err = d.Set("service_sla_tie_break", vv); err != nil {
				return fmt.Errorf("Error reading service_sla_tie_break: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_sla_tie_break: %v", err)
		}
	}

	return nil
}

func flattenWantempSystemSdwanZoneFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWantempSystemSdwanZoneName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanZoneServiceSlaTieBreak(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWantempSystemSdwanZone(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWantempSystemSdwanZoneName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("service_sla_tie_break"); ok || d.HasChange("service_sla_tie_break") {
		t, err := expandWantempSystemSdwanZoneServiceSlaTieBreak(d, v, "service_sla_tie_break")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-sla-tie-break"] = t
		}
	}

	return &obj, nil
}
