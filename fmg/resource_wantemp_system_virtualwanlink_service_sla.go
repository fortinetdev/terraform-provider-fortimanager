// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Service level agreement (SLA).

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWantempSystemVirtualWanLinkServiceSla() *schema.Resource {
	return &schema.Resource{
		Create: resourceWantempSystemVirtualWanLinkServiceSlaCreate,
		Read:   resourceWantempSystemVirtualWanLinkServiceSlaRead,
		Update: resourceWantempSystemVirtualWanLinkServiceSlaUpdate,
		Delete: resourceWantempSystemVirtualWanLinkServiceSlaDelete,

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
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"health_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceWantempSystemVirtualWanLinkServiceSlaCreate(d *schema.ResourceData, m interface{}) error {
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
	service := d.Get("service").(string)
	paradict["wanprof"] = wanprof
	paradict["service"] = service

	obj, err := getObjectWantempSystemVirtualWanLinkServiceSla(d)
	if err != nil {
		return fmt.Errorf("Error creating WantempSystemVirtualWanLinkServiceSla resource while getting object: %v", err)
	}

	_, err = c.CreateWantempSystemVirtualWanLinkServiceSla(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WantempSystemVirtualWanLinkServiceSla resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWantempSystemVirtualWanLinkServiceSlaRead(d, m)
}

func resourceWantempSystemVirtualWanLinkServiceSlaUpdate(d *schema.ResourceData, m interface{}) error {
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
	service := d.Get("service").(string)
	paradict["wanprof"] = wanprof
	paradict["service"] = service

	obj, err := getObjectWantempSystemVirtualWanLinkServiceSla(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemVirtualWanLinkServiceSla resource while getting object: %v", err)
	}

	_, err = c.UpdateWantempSystemVirtualWanLinkServiceSla(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemVirtualWanLinkServiceSla resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWantempSystemVirtualWanLinkServiceSlaRead(d, m)
}

func resourceWantempSystemVirtualWanLinkServiceSlaDelete(d *schema.ResourceData, m interface{}) error {
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
	service := d.Get("service").(string)
	paradict["wanprof"] = wanprof
	paradict["service"] = service

	err = c.DeleteWantempSystemVirtualWanLinkServiceSla(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WantempSystemVirtualWanLinkServiceSla resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWantempSystemVirtualWanLinkServiceSlaRead(d *schema.ResourceData, m interface{}) error {
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
	service := d.Get("service").(string)
	if wanprof == "" {
		wanprof = importOptionChecking(m.(*FortiClient).Cfg, "wanprof")
		if wanprof == "" {
			return fmt.Errorf("Parameter wanprof is missing")
		}
		if err = d.Set("wanprof", wanprof); err != nil {
			return fmt.Errorf("Error set params wanprof: %v", err)
		}
	}
	if service == "" {
		service = importOptionChecking(m.(*FortiClient).Cfg, "service")
		if service == "" {
			return fmt.Errorf("Parameter service is missing")
		}
		if err = d.Set("service", service); err != nil {
			return fmt.Errorf("Error set params service: %v", err)
		}
	}
	paradict["wanprof"] = wanprof
	paradict["service"] = service

	o, err := c.ReadWantempSystemVirtualWanLinkServiceSla(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemVirtualWanLinkServiceSla resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWantempSystemVirtualWanLinkServiceSla(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemVirtualWanLinkServiceSla resource from API: %v", err)
	}
	return nil
}

func flattenWantempSystemVirtualWanLinkServiceSlaHealthCheck3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceSlaId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWantempSystemVirtualWanLinkServiceSla(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("health_check", flattenWantempSystemVirtualWanLinkServiceSlaHealthCheck3rdl(o["health-check"], d, "health_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check"], "WantempSystemVirtualWanLinkServiceSla-HealthCheck"); ok {
			if err = d.Set("health_check", vv); err != nil {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWantempSystemVirtualWanLinkServiceSlaId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WantempSystemVirtualWanLinkServiceSla-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenWantempSystemVirtualWanLinkServiceSlaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWantempSystemVirtualWanLinkServiceSlaHealthCheck3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceSlaId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWantempSystemVirtualWanLinkServiceSla(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandWantempSystemVirtualWanLinkServiceSlaHealthCheck3rdl(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWantempSystemVirtualWanLinkServiceSlaId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
