// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Domain filter settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDnsfilterProfileDomainFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDnsfilterProfileDomainFilterUpdate,
		Read:   resourceObjectDnsfilterProfileDomainFilterRead,
		Update: resourceObjectDnsfilterProfileDomainFilterUpdate,
		Delete: resourceObjectDnsfilterProfileDomainFilterDelete,

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
			"domain_filter_table": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectDnsfilterProfileDomainFilterUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectDnsfilterProfileDomainFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDnsfilterProfileDomainFilter resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectDnsfilterProfileDomainFilter(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDnsfilterProfileDomainFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectDnsfilterProfileDomainFilter")

	return resourceObjectDnsfilterProfileDomainFilterRead(d, m)
}

func resourceObjectDnsfilterProfileDomainFilterDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectDnsfilterProfileDomainFilter(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDnsfilterProfileDomainFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDnsfilterProfileDomainFilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectDnsfilterProfileDomainFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDnsfilterProfileDomainFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDnsfilterProfileDomainFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDnsfilterProfileDomainFilter resource from API: %v", err)
	}
	return nil
}

func flattenObjectDnsfilterProfileDomainFilterDomainFilterTable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2num(v)
}

func refreshObjectObjectDnsfilterProfileDomainFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("domain_filter_table", flattenObjectDnsfilterProfileDomainFilterDomainFilterTable2edl(o["domain-filter-table"], d, "domain_filter_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-filter-table"], "ObjectDnsfilterProfileDomainFilter-DomainFilterTable"); ok {
			if err = d.Set("domain_filter_table", vv); err != nil {
				return fmt.Errorf("Error reading domain_filter_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_filter_table: %v", err)
		}
	}

	return nil
}

func flattenObjectDnsfilterProfileDomainFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDnsfilterProfileDomainFilterDomainFilterTable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDnsfilterProfileDomainFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("domain_filter_table"); ok || d.HasChange("domain_filter_table") {
		t, err := expandObjectDnsfilterProfileDomainFilterDomainFilterTable2edl(d, v, "domain_filter_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-filter-table"] = t
		}
	}

	return &obj, nil
}
