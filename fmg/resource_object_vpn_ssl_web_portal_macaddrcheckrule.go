// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Client MAC address check rule.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnSslWebPortalMacAddrCheckRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnSslWebPortalMacAddrCheckRuleCreate,
		Read:   resourceObjectVpnSslWebPortalMacAddrCheckRuleRead,
		Update: resourceObjectVpnSslWebPortalMacAddrCheckRuleUpdate,
		Delete: resourceObjectVpnSslWebPortalMacAddrCheckRuleDelete,

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
			"portal": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mac_addr_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mac_addr_mask": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectVpnSslWebPortalMacAddrCheckRuleCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	portal := d.Get("portal").(string)
	paradict["portal"] = portal

	obj, err := getObjectObjectVpnSslWebPortalMacAddrCheckRule(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebPortalMacAddrCheckRule resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVpnSslWebPortalMacAddrCheckRule(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebPortalMacAddrCheckRule resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnSslWebPortalMacAddrCheckRuleRead(d, m)
}

func resourceObjectVpnSslWebPortalMacAddrCheckRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	portal := d.Get("portal").(string)
	paradict["portal"] = portal

	obj, err := getObjectObjectVpnSslWebPortalMacAddrCheckRule(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebPortalMacAddrCheckRule resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVpnSslWebPortalMacAddrCheckRule(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebPortalMacAddrCheckRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnSslWebPortalMacAddrCheckRuleRead(d, m)
}

func resourceObjectVpnSslWebPortalMacAddrCheckRuleDelete(d *schema.ResourceData, m interface{}) error {
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

	portal := d.Get("portal").(string)
	paradict["portal"] = portal

	err = c.DeleteObjectVpnSslWebPortalMacAddrCheckRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnSslWebPortalMacAddrCheckRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnSslWebPortalMacAddrCheckRuleRead(d *schema.ResourceData, m interface{}) error {
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

	portal := d.Get("portal").(string)
	if portal == "" {
		portal = importOptionChecking(m.(*FortiClient).Cfg, "portal")
		if portal == "" {
			return fmt.Errorf("Parameter portal is missing")
		}
		if err = d.Set("portal", portal); err != nil {
			return fmt.Errorf("Error set params portal: %v", err)
		}
	}
	paradict["portal"] = portal

	o, err := c.ReadObjectVpnSslWebPortalMacAddrCheckRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebPortalMacAddrCheckRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnSslWebPortalMacAddrCheckRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebPortalMacAddrCheckRule resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnSslWebPortalMacAddrCheckRuleMacAddrList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnSslWebPortalMacAddrCheckRuleMacAddrMask2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalMacAddrCheckRuleName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnSslWebPortalMacAddrCheckRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("mac_addr_list", flattenObjectVpnSslWebPortalMacAddrCheckRuleMacAddrList2edl(o["mac-addr-list"], d, "mac_addr_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-addr-list"], "ObjectVpnSslWebPortalMacAddrCheckRule-MacAddrList"); ok {
			if err = d.Set("mac_addr_list", vv); err != nil {
				return fmt.Errorf("Error reading mac_addr_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_addr_list: %v", err)
		}
	}

	if err = d.Set("mac_addr_mask", flattenObjectVpnSslWebPortalMacAddrCheckRuleMacAddrMask2edl(o["mac-addr-mask"], d, "mac_addr_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-addr-mask"], "ObjectVpnSslWebPortalMacAddrCheckRule-MacAddrMask"); ok {
			if err = d.Set("mac_addr_mask", vv); err != nil {
				return fmt.Errorf("Error reading mac_addr_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_addr_mask: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectVpnSslWebPortalMacAddrCheckRuleName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVpnSslWebPortalMacAddrCheckRule-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnSslWebPortalMacAddrCheckRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnSslWebPortalMacAddrCheckRuleMacAddrList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnSslWebPortalMacAddrCheckRuleMacAddrMask2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalMacAddrCheckRuleName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnSslWebPortalMacAddrCheckRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mac_addr_list"); ok || d.HasChange("mac_addr_list") {
		t, err := expandObjectVpnSslWebPortalMacAddrCheckRuleMacAddrList2edl(d, v, "mac_addr_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-addr-list"] = t
		}
	}

	if v, ok := d.GetOk("mac_addr_mask"); ok || d.HasChange("mac_addr_mask") {
		t, err := expandObjectVpnSslWebPortalMacAddrCheckRuleMacAddrMask2edl(d, v, "mac_addr_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-addr-mask"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectVpnSslWebPortalMacAddrCheckRuleName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
