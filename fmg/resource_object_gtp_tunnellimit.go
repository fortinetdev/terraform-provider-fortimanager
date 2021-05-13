// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: GTP tunnel limiter.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectGtpTunnelLimit() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectGtpTunnelLimitCreate,
		Read:   resourceObjectGtpTunnelLimitRead,
		Update: resourceObjectGtpTunnelLimitUpdate,
		Delete: resourceObjectGtpTunnelLimitDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"tunnel_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectGtpTunnelLimitCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectGtpTunnelLimit(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectGtpTunnelLimit resource while getting object: %v", err)
	}

	_, err = c.CreateObjectGtpTunnelLimit(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectGtpTunnelLimit resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectGtpTunnelLimitRead(d, m)
}

func resourceObjectGtpTunnelLimitUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectGtpTunnelLimit(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGtpTunnelLimit resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectGtpTunnelLimit(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGtpTunnelLimit resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectGtpTunnelLimitRead(d, m)
}

func resourceObjectGtpTunnelLimitDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectGtpTunnelLimit(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectGtpTunnelLimit resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectGtpTunnelLimitRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectGtpTunnelLimit(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGtpTunnelLimit resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectGtpTunnelLimit(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGtpTunnelLimit resource from API: %v", err)
	}
	return nil
}

func flattenObjectGtpTunnelLimitName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpTunnelLimitTunnelLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectGtpTunnelLimit(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenObjectGtpTunnelLimitName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectGtpTunnelLimit-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("tunnel_limit", flattenObjectGtpTunnelLimitTunnelLimit(o["tunnel-limit"], d, "tunnel_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-limit"], "ObjectGtpTunnelLimit-TunnelLimit"); ok {
			if err = d.Set("tunnel_limit", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_limit: %v", err)
		}
	}

	return nil
}

func flattenObjectGtpTunnelLimitFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectGtpTunnelLimitName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpTunnelLimitTunnelLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectGtpTunnelLimit(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectGtpTunnelLimitName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_limit"); ok {
		t, err := expandObjectGtpTunnelLimitTunnelLimit(d, v, "tunnel_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-limit"] = t
		}
	}

	return &obj, nil
}
