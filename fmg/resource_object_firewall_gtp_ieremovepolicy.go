// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IE remove policy.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallGtpIeRemovePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallGtpIeRemovePolicyCreate,
		Read:   resourceObjectFirewallGtpIeRemovePolicyRead,
		Update: resourceObjectFirewallGtpIeRemovePolicyUpdate,
		Delete: resourceObjectFirewallGtpIeRemovePolicyDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"remove_ies": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sgsn_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sgsn_addr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallGtpIeRemovePolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallGtpIeRemovePolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallGtpIeRemovePolicy resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallGtpIeRemovePolicy(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallGtpIeRemovePolicy resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallGtpIeRemovePolicyRead(d, m)
}

func resourceObjectFirewallGtpIeRemovePolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallGtpIeRemovePolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpIeRemovePolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallGtpIeRemovePolicy(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpIeRemovePolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallGtpIeRemovePolicyRead(d, m)
}

func resourceObjectFirewallGtpIeRemovePolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallGtpIeRemovePolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallGtpIeRemovePolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallGtpIeRemovePolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallGtpIeRemovePolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpIeRemovePolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallGtpIeRemovePolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpIeRemovePolicy resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallGtpIeRemovePolicyId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeRemovePolicyRemoveIes2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallGtpIeRemovePolicySgsnAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFirewallGtpIeRemovePolicySgsnAddr62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectObjectFirewallGtpIeRemovePolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenObjectFirewallGtpIeRemovePolicyId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallGtpIeRemovePolicy-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("remove_ies", flattenObjectFirewallGtpIeRemovePolicyRemoveIes2edl(o["remove-ies"], d, "remove_ies")); err != nil {
		if vv, ok := fortiAPIPatch(o["remove-ies"], "ObjectFirewallGtpIeRemovePolicy-RemoveIes"); ok {
			if err = d.Set("remove_ies", vv); err != nil {
				return fmt.Errorf("Error reading remove_ies: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remove_ies: %v", err)
		}
	}

	if err = d.Set("sgsn_addr", flattenObjectFirewallGtpIeRemovePolicySgsnAddr2edl(o["sgsn-addr"], d, "sgsn_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["sgsn-addr"], "ObjectFirewallGtpIeRemovePolicy-SgsnAddr"); ok {
			if err = d.Set("sgsn_addr", vv); err != nil {
				return fmt.Errorf("Error reading sgsn_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sgsn_addr: %v", err)
		}
	}

	if err = d.Set("sgsn_addr6", flattenObjectFirewallGtpIeRemovePolicySgsnAddr62edl(o["sgsn-addr6"], d, "sgsn_addr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["sgsn-addr6"], "ObjectFirewallGtpIeRemovePolicy-SgsnAddr6"); ok {
			if err = d.Set("sgsn_addr6", vv); err != nil {
				return fmt.Errorf("Error reading sgsn_addr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sgsn_addr6: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallGtpIeRemovePolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallGtpIeRemovePolicyId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeRemovePolicyRemoveIes2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallGtpIeRemovePolicySgsnAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallGtpIeRemovePolicySgsnAddr62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectObjectFirewallGtpIeRemovePolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallGtpIeRemovePolicyId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("remove_ies"); ok || d.HasChange("remove_ies") {
		t, err := expandObjectFirewallGtpIeRemovePolicyRemoveIes2edl(d, v, "remove_ies")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-ies"] = t
		}
	}

	if v, ok := d.GetOk("sgsn_addr"); ok || d.HasChange("sgsn_addr") {
		t, err := expandObjectFirewallGtpIeRemovePolicySgsnAddr2edl(d, v, "sgsn_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sgsn-addr"] = t
		}
	}

	if v, ok := d.GetOk("sgsn_addr6"); ok || d.HasChange("sgsn_addr6") {
		t, err := expandObjectFirewallGtpIeRemovePolicySgsnAddr62edl(d, v, "sgsn_addr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sgsn-addr6"] = t
		}
	}

	return &obj, nil
}
