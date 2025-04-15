// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DNS domain filter entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDnsfilterDomainFilterEntriesMove() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDnsfilterDomainFilterEntriesMoveUpdate,
		Read:   resourceObjectDnsfilterDomainFilterEntriesMoveRead,
		Update: resourceObjectDnsfilterDomainFilterEntriesMoveUpdate,
		Delete: resourceObjectDnsfilterDomainFilterEntriesMoveDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"state_pos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
			"domain_filter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"option": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceObjectDnsfilterDomainFilterEntriesMoveUpdate(d *schema.ResourceData, m interface{}) error {
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

	domain_filter := d.Get("domain_filter").(string)
	entries := d.Get("entries").(string)
	paradict["domain_filter"] = domain_filter
	paradict["entries"] = entries

	target := d.Get("target").(string)
	obj, err := getObjectObjectDnsfilterDomainFilterEntriesMove(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDnsfilterDomainFilterEntriesMove resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectDnsfilterDomainFilterEntriesMove(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDnsfilterDomainFilterEntriesMove resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectDnsfilterDomainFilterEntriesMove" + "_" + domain_filter + "_" + entries + "_" + target)

	return resourceObjectDnsfilterDomainFilterEntriesMoveRead(d, m)
}

func resourceObjectDnsfilterDomainFilterEntriesMoveDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceObjectDnsfilterDomainFilterEntriesMoveRead(d *schema.ResourceData, m interface{}) error {
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

	sid := d.Get("entries").(string)
	did := d.Get("target").(string)
	action := d.Get("option").(string)

	domain_filter := d.Get("domain_filter").(string)
	if domain_filter == "" {
		domain_filter = importOptionChecking(m.(*FortiClient).Cfg, "domain_filter")
		if domain_filter == "" {
			return fmt.Errorf("Parameter domain_filter is missing")
		}
		if err = d.Set("domain_filter", domain_filter); err != nil {
			return fmt.Errorf("Error set params domain_filter: %v", err)
		}
	}
	paradict["domain_filter"] = domain_filter

	o, err := c.ReadObjectDnsfilterDomainFilterEntriesMove(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectDnsfilterDomainFilterEntriesMove resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if o != nil {
		i := 1
		now_sid := -1
		now_did := -1

		for _, z := range o {
			if v, ok := z.(map[string]interface{}); ok {
				if _, ok := v["id"]; !ok {
					return fmt.Errorf("Error reading ObjectDnsfilterDomainFilterEntriesMove resource: id doesn't exist.")
				}

				vid := fmt.Sprintf("%v", v["id"])

				if vid == sid {
					now_sid = i
				}

				if vid == did {
					now_did = i
				}
			} else {
				return fmt.Errorf("Error reading ObjectDnsfilterDomainFilterEntriesMove resource: no valid map string.")
			}

			i += 1
		}

		state_pos := ""

		if now_sid == -1 || now_did == -1 {
			if now_sid == -1 && now_did == -1 {
				state_pos = "id(" + sid + ") and target(" + did + ") were deleted"
			} else if now_sid == -1 {
				state_pos = "id(" + sid + ") was deleted"
			} else if now_did == -1 {
				state_pos = "target(" + did + ") was deleted"
			}
		} else {
			bconsistent := true
			if action == "before" {
				if now_sid != now_did-1 {
					bconsistent = false
				}
			}

			if action == "after" {
				if now_sid != now_did+1 {
					bconsistent = false
				}
			}

			if bconsistent == false {
				relative_pos := now_sid - now_did

				if relative_pos > 0 {
					state_pos = "id(" + sid + ") is " + strconv.Itoa(relative_pos) + " behind target(" + did + ")"
				} else {
					state_pos = "id(" + sid + ") is " + strconv.Itoa(-relative_pos) + " ahead of target(" + did + ")"
				}
			}
		}

		d.Set("state_pos", state_pos)
	}

	return nil
}

func flattenObjectDnsfilterDomainFilterEntriesMoveFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDnsfilterDomainFilterEntriesMoveTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterDomainFilterEntriesMoveOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDnsfilterDomainFilterEntriesMove(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandObjectDnsfilterDomainFilterEntriesMoveTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandObjectDnsfilterDomainFilterEntriesMoveOption(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	return &obj, nil
}
