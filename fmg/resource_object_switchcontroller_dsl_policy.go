// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DSL policy.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerDslPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerDslPolicyCreate,
		Read:   resourceObjectSwitchControllerDslPolicyRead,
		Update: resourceObjectSwitchControllerDslPolicyUpdate,
		Delete: resourceObjectSwitchControllerDslPolicyDelete,

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
			"append_padding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpe_aele": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpe_aele_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cs": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ds_bitswap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"pause_frame": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"us_bitswap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSwitchControllerDslPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSwitchControllerDslPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerDslPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSwitchControllerDslPolicy(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerDslPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerDslPolicyRead(d, m)
}

func resourceObjectSwitchControllerDslPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSwitchControllerDslPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerDslPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerDslPolicy(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerDslPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerDslPolicyRead(d, m)
}

func resourceObjectSwitchControllerDslPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSwitchControllerDslPolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerDslPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerDslPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSwitchControllerDslPolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerDslPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerDslPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerDslPolicy resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerDslPolicyAppendPadding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDslPolicyCpeAele(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDslPolicyCpeAeleMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDslPolicyCs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerDslPolicyDsBitswap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDslPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDslPolicyPauseFrame(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDslPolicyProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDslPolicyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDslPolicyUsBitswap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerDslPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("append_padding", flattenObjectSwitchControllerDslPolicyAppendPadding(o["append_padding"], d, "append_padding")); err != nil {
		if vv, ok := fortiAPIPatch(o["append_padding"], "ObjectSwitchControllerDslPolicy-AppendPadding"); ok {
			if err = d.Set("append_padding", vv); err != nil {
				return fmt.Errorf("Error reading append_padding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading append_padding: %v", err)
		}
	}

	if err = d.Set("cpe_aele", flattenObjectSwitchControllerDslPolicyCpeAele(o["cpe-aele"], d, "cpe_aele")); err != nil {
		if vv, ok := fortiAPIPatch(o["cpe-aele"], "ObjectSwitchControllerDslPolicy-CpeAele"); ok {
			if err = d.Set("cpe_aele", vv); err != nil {
				return fmt.Errorf("Error reading cpe_aele: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cpe_aele: %v", err)
		}
	}

	if err = d.Set("cpe_aele_mode", flattenObjectSwitchControllerDslPolicyCpeAeleMode(o["cpe-aele-mode"], d, "cpe_aele_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["cpe-aele-mode"], "ObjectSwitchControllerDslPolicy-CpeAeleMode"); ok {
			if err = d.Set("cpe_aele_mode", vv); err != nil {
				return fmt.Errorf("Error reading cpe_aele_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cpe_aele_mode: %v", err)
		}
	}

	if err = d.Set("cs", flattenObjectSwitchControllerDslPolicyCs(o["cs"], d, "cs")); err != nil {
		if vv, ok := fortiAPIPatch(o["cs"], "ObjectSwitchControllerDslPolicy-Cs"); ok {
			if err = d.Set("cs", vv); err != nil {
				return fmt.Errorf("Error reading cs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cs: %v", err)
		}
	}

	if err = d.Set("ds_bitswap", flattenObjectSwitchControllerDslPolicyDsBitswap(o["ds-bitswap"], d, "ds_bitswap")); err != nil {
		if vv, ok := fortiAPIPatch(o["ds-bitswap"], "ObjectSwitchControllerDslPolicy-DsBitswap"); ok {
			if err = d.Set("ds_bitswap", vv); err != nil {
				return fmt.Errorf("Error reading ds_bitswap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ds_bitswap: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerDslPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerDslPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pause_frame", flattenObjectSwitchControllerDslPolicyPauseFrame(o["pause-frame"], d, "pause_frame")); err != nil {
		if vv, ok := fortiAPIPatch(o["pause-frame"], "ObjectSwitchControllerDslPolicy-PauseFrame"); ok {
			if err = d.Set("pause_frame", vv); err != nil {
				return fmt.Errorf("Error reading pause_frame: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pause_frame: %v", err)
		}
	}

	if err = d.Set("profile", flattenObjectSwitchControllerDslPolicyProfile(o["profile"], d, "profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile"], "ObjectSwitchControllerDslPolicy-Profile"); ok {
			if err = d.Set("profile", vv); err != nil {
				return fmt.Errorf("Error reading profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectSwitchControllerDslPolicyType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectSwitchControllerDslPolicy-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("us_bitswap", flattenObjectSwitchControllerDslPolicyUsBitswap(o["us-bitswap"], d, "us_bitswap")); err != nil {
		if vv, ok := fortiAPIPatch(o["us-bitswap"], "ObjectSwitchControllerDslPolicy-UsBitswap"); ok {
			if err = d.Set("us_bitswap", vv); err != nil {
				return fmt.Errorf("Error reading us_bitswap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading us_bitswap: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerDslPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerDslPolicyAppendPadding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDslPolicyCpeAele(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDslPolicyCpeAeleMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDslPolicyCs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerDslPolicyDsBitswap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDslPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDslPolicyPauseFrame(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDslPolicyProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDslPolicyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDslPolicyUsBitswap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerDslPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("append_padding"); ok || d.HasChange("append_padding") {
		t, err := expandObjectSwitchControllerDslPolicyAppendPadding(d, v, "append_padding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["append_padding"] = t
		}
	}

	if v, ok := d.GetOk("cpe_aele"); ok || d.HasChange("cpe_aele") {
		t, err := expandObjectSwitchControllerDslPolicyCpeAele(d, v, "cpe_aele")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpe-aele"] = t
		}
	}

	if v, ok := d.GetOk("cpe_aele_mode"); ok || d.HasChange("cpe_aele_mode") {
		t, err := expandObjectSwitchControllerDslPolicyCpeAeleMode(d, v, "cpe_aele_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpe-aele-mode"] = t
		}
	}

	if v, ok := d.GetOk("cs"); ok || d.HasChange("cs") {
		t, err := expandObjectSwitchControllerDslPolicyCs(d, v, "cs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cs"] = t
		}
	}

	if v, ok := d.GetOk("ds_bitswap"); ok || d.HasChange("ds_bitswap") {
		t, err := expandObjectSwitchControllerDslPolicyDsBitswap(d, v, "ds_bitswap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ds-bitswap"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerDslPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pause_frame"); ok || d.HasChange("pause_frame") {
		t, err := expandObjectSwitchControllerDslPolicyPauseFrame(d, v, "pause_frame")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pause-frame"] = t
		}
	}

	if v, ok := d.GetOk("profile"); ok || d.HasChange("profile") {
		t, err := expandObjectSwitchControllerDslPolicyProfile(d, v, "profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectSwitchControllerDslPolicyType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("us_bitswap"); ok || d.HasChange("us_bitswap") {
		t, err := expandObjectSwitchControllerDslPolicyUsBitswap(d, v, "us_bitswap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["us-bitswap"] = t
		}
	}

	return &obj, nil
}
