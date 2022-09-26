// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure virtual wire pairs.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemVirtualWirePair() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemVirtualWirePairCreate,
		Read:   resourceObjectSystemVirtualWirePairRead,
		Update: resourceObjectSystemVirtualWirePairUpdate,
		Delete: resourceObjectSystemVirtualWirePairDelete,

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
			"member": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"poweroff_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"poweron_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wildcard_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemVirtualWirePairCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemVirtualWirePair(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemVirtualWirePair resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemVirtualWirePair(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemVirtualWirePair resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemVirtualWirePairRead(d, m)
}

func resourceObjectSystemVirtualWirePairUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemVirtualWirePair(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemVirtualWirePair resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemVirtualWirePair(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemVirtualWirePair resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemVirtualWirePairRead(d, m)
}

func resourceObjectSystemVirtualWirePairDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemVirtualWirePair(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemVirtualWirePair resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemVirtualWirePairRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemVirtualWirePair(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemVirtualWirePair resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemVirtualWirePair(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemVirtualWirePair resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemVirtualWirePairMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemVirtualWirePairName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemVirtualWirePairPoweroffBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemVirtualWirePairPoweronBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemVirtualWirePairVlanFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemVirtualWirePairWildcardVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemVirtualWirePair(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("member", flattenObjectSystemVirtualWirePairMember(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "ObjectSystemVirtualWirePair-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSystemVirtualWirePairName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemVirtualWirePair-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("poweroff_bypass", flattenObjectSystemVirtualWirePairPoweroffBypass(o["poweroff-bypass"], d, "poweroff_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["poweroff-bypass"], "ObjectSystemVirtualWirePair-PoweroffBypass"); ok {
			if err = d.Set("poweroff_bypass", vv); err != nil {
				return fmt.Errorf("Error reading poweroff_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poweroff_bypass: %v", err)
		}
	}

	if err = d.Set("poweron_bypass", flattenObjectSystemVirtualWirePairPoweronBypass(o["poweron-bypass"], d, "poweron_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["poweron-bypass"], "ObjectSystemVirtualWirePair-PoweronBypass"); ok {
			if err = d.Set("poweron_bypass", vv); err != nil {
				return fmt.Errorf("Error reading poweron_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poweron_bypass: %v", err)
		}
	}

	if err = d.Set("vlan_filter", flattenObjectSystemVirtualWirePairVlanFilter(o["vlan-filter"], d, "vlan_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-filter"], "ObjectSystemVirtualWirePair-VlanFilter"); ok {
			if err = d.Set("vlan_filter", vv); err != nil {
				return fmt.Errorf("Error reading vlan_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_filter: %v", err)
		}
	}

	if err = d.Set("wildcard_vlan", flattenObjectSystemVirtualWirePairWildcardVlan(o["wildcard-vlan"], d, "wildcard_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard-vlan"], "ObjectSystemVirtualWirePair-WildcardVlan"); ok {
			if err = d.Set("wildcard_vlan", vv); err != nil {
				return fmt.Errorf("Error reading wildcard_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard_vlan: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemVirtualWirePairFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemVirtualWirePairMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemVirtualWirePairName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemVirtualWirePairPoweroffBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemVirtualWirePairPoweronBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemVirtualWirePairVlanFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemVirtualWirePairWildcardVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemVirtualWirePair(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandObjectSystemVirtualWirePairMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSystemVirtualWirePairName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("poweroff_bypass"); ok || d.HasChange("poweroff_bypass") {
		t, err := expandObjectSystemVirtualWirePairPoweroffBypass(d, v, "poweroff_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poweroff-bypass"] = t
		}
	}

	if v, ok := d.GetOk("poweron_bypass"); ok || d.HasChange("poweron_bypass") {
		t, err := expandObjectSystemVirtualWirePairPoweronBypass(d, v, "poweron_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poweron-bypass"] = t
		}
	}

	if v, ok := d.GetOk("vlan_filter"); ok || d.HasChange("vlan_filter") {
		t, err := expandObjectSystemVirtualWirePairVlanFilter(d, v, "vlan_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-filter"] = t
		}
	}

	if v, ok := d.GetOk("wildcard_vlan"); ok || d.HasChange("wildcard_vlan") {
		t, err := expandObjectSystemVirtualWirePairWildcardVlan(d, v, "wildcard_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard-vlan"] = t
		}
	}

	return &obj, nil
}
