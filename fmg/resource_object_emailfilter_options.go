// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure AntiSpam options.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectEmailfilterOptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectEmailfilterOptionsUpdate,
		Read:   resourceObjectEmailfilterOptionsRead,
		Update: resourceObjectEmailfilterOptionsUpdate,
		Delete: resourceObjectEmailfilterOptionsDelete,

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
			"dns_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectEmailfilterOptionsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectEmailfilterOptions(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEmailfilterOptions resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectEmailfilterOptions(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEmailfilterOptions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectEmailfilterOptions")

	return resourceObjectEmailfilterOptionsRead(d, m)
}

func resourceObjectEmailfilterOptionsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectEmailfilterOptions(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectEmailfilterOptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectEmailfilterOptionsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectEmailfilterOptions(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEmailfilterOptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectEmailfilterOptions(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEmailfilterOptions resource from API: %v", err)
	}
	return nil
}

func flattenObjectEmailfilterOptionsDnsTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectEmailfilterOptions(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dns_timeout", flattenObjectEmailfilterOptionsDnsTimeout(o["dns-timeout"], d, "dns_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-timeout"], "ObjectEmailfilterOptions-DnsTimeout"); ok {
			if err = d.Set("dns_timeout", vv); err != nil {
				return fmt.Errorf("Error reading dns_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_timeout: %v", err)
		}
	}

	return nil
}

func flattenObjectEmailfilterOptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectEmailfilterOptionsDnsTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectEmailfilterOptions(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dns_timeout"); ok {
		t, err := expandObjectEmailfilterOptionsDnsTimeout(d, v, "dns_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-timeout"] = t
		}
	}

	return &obj, nil
}
