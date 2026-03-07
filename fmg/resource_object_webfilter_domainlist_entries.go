// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectWebfilter DomainListEntries

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebfilterDomainListEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterDomainListEntriesCreate,
		Read:   resourceObjectWebfilterDomainListEntriesRead,
		Update: resourceObjectWebfilterDomainListEntriesUpdate,
		Delete: resourceObjectWebfilterDomainListEntriesDelete,

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
			"domain_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectWebfilterDomainListEntriesCreate(d *schema.ResourceData, m interface{}) error {
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

	domain_list := d.Get("domain_list").(string)
	paradict["domain_list"] = domain_list

	obj, err := getObjectObjectWebfilterDomainListEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterDomainListEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWebfilterDomainListEntries(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterDomainListEntries resource: %v", err)
	}

	d.SetId(getStringKey(d, "domain"))

	return resourceObjectWebfilterDomainListEntriesRead(d, m)
}

func resourceObjectWebfilterDomainListEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	domain_list := d.Get("domain_list").(string)
	paradict["domain_list"] = domain_list

	obj, err := getObjectObjectWebfilterDomainListEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterDomainListEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWebfilterDomainListEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterDomainListEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "domain"))

	return resourceObjectWebfilterDomainListEntriesRead(d, m)
}

func resourceObjectWebfilterDomainListEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	domain_list := d.Get("domain_list").(string)
	paradict["domain_list"] = domain_list

	wsParams["adom"] = adomv

	err = c.DeleteObjectWebfilterDomainListEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterDomainListEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterDomainListEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	domain_list := d.Get("domain_list").(string)
	if domain_list == "" {
		domain_list = importOptionChecking(m.(*FortiClient).Cfg, "domain_list")
		if domain_list == "" {
			return fmt.Errorf("Parameter domain_list is missing")
		}
		if err = d.Set("domain_list", domain_list); err != nil {
			return fmt.Errorf("Error set params domain_list: %v", err)
		}
	}
	paradict["domain_list"] = domain_list

	o, err := c.ReadObjectWebfilterDomainListEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterDomainListEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterDomainListEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterDomainListEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterDomainListEntriesDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebfilterDomainListEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("domain", flattenObjectWebfilterDomainListEntriesDomain2edl(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "ObjectWebfilterDomainListEntries-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterDomainListEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterDomainListEntriesDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebfilterDomainListEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandObjectWebfilterDomainListEntriesDomain2edl(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	return &obj, nil
}
