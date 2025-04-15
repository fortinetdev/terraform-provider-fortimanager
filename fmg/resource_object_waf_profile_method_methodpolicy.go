// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: HTTP method policy.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWafProfileMethodMethodPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWafProfileMethodMethodPolicyCreate,
		Read:   resourceObjectWafProfileMethodMethodPolicyRead,
		Update: resourceObjectWafProfileMethodMethodPolicyUpdate,
		Delete: resourceObjectWafProfileMethodMethodPolicyDelete,

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
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allowed_methods": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWafProfileMethodMethodPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWafProfileMethodMethodPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWafProfileMethodMethodPolicy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWafProfileMethodMethodPolicy(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWafProfileMethodMethodPolicy resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWafProfileMethodMethodPolicyRead(d, m)
}

func resourceObjectWafProfileMethodMethodPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWafProfileMethodMethodPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfileMethodMethodPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWafProfileMethodMethodPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfileMethodMethodPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWafProfileMethodMethodPolicyRead(d, m)
}

func resourceObjectWafProfileMethodMethodPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWafProfileMethodMethodPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWafProfileMethodMethodPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWafProfileMethodMethodPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWafProfileMethodMethodPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfileMethodMethodPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWafProfileMethodMethodPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfileMethodMethodPolicy resource from API: %v", err)
	}
	return nil
}

func flattenObjectWafProfileMethodMethodPolicyAddress3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWafProfileMethodMethodPolicyAllowedMethods3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWafProfileMethodMethodPolicyId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileMethodMethodPolicyPattern3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileMethodMethodPolicyRegex3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWafProfileMethodMethodPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("address", flattenObjectWafProfileMethodMethodPolicyAddress3rdl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "ObjectWafProfileMethodMethodPolicy-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("allowed_methods", flattenObjectWafProfileMethodMethodPolicyAllowedMethods3rdl(o["allowed-methods"], d, "allowed_methods")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowed-methods"], "ObjectWafProfileMethodMethodPolicy-AllowedMethods"); ok {
			if err = d.Set("allowed_methods", vv); err != nil {
				return fmt.Errorf("Error reading allowed_methods: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowed_methods: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectWafProfileMethodMethodPolicyId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWafProfileMethodMethodPolicy-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("pattern", flattenObjectWafProfileMethodMethodPolicyPattern3rdl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "ObjectWafProfileMethodMethodPolicy-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("regex", flattenObjectWafProfileMethodMethodPolicyRegex3rdl(o["regex"], d, "regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["regex"], "ObjectWafProfileMethodMethodPolicy-Regex"); ok {
			if err = d.Set("regex", vv); err != nil {
				return fmt.Errorf("Error reading regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading regex: %v", err)
		}
	}

	return nil
}

func flattenObjectWafProfileMethodMethodPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWafProfileMethodMethodPolicyAddress3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWafProfileMethodMethodPolicyAllowedMethods3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWafProfileMethodMethodPolicyId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileMethodMethodPolicyPattern3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileMethodMethodPolicyRegex3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWafProfileMethodMethodPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandObjectWafProfileMethodMethodPolicyAddress3rdl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("allowed_methods"); ok || d.HasChange("allowed_methods") {
		t, err := expandObjectWafProfileMethodMethodPolicyAllowedMethods3rdl(d, v, "allowed_methods")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-methods"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWafProfileMethodMethodPolicyId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandObjectWafProfileMethodMethodPolicyPattern3rdl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("regex"); ok || d.HasChange("regex") {
		t, err := expandObjectWafProfileMethodMethodPolicyRegex3rdl(d, v, "regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["regex"] = t
		}
	}

	return &obj, nil
}
