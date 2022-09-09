// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Config global/VDOM Wildcard FQDN address.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallWildcardFqdnCustom() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallWildcardFqdnCustomCreate,
		Read:   resourceObjectFirewallWildcardFqdnCustomRead,
		Update: resourceObjectFirewallWildcardFqdnCustomUpdate,
		Delete: resourceObjectFirewallWildcardFqdnCustomDelete,

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
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wildcard_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallWildcardFqdnCustomCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallWildcardFqdnCustom(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallWildcardFqdnCustom resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallWildcardFqdnCustom(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallWildcardFqdnCustom resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallWildcardFqdnCustomRead(d, m)
}

func resourceObjectFirewallWildcardFqdnCustomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallWildcardFqdnCustom(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallWildcardFqdnCustom resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallWildcardFqdnCustom(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallWildcardFqdnCustom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallWildcardFqdnCustomRead(d, m)
}

func resourceObjectFirewallWildcardFqdnCustomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallWildcardFqdnCustom(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallWildcardFqdnCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallWildcardFqdnCustomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallWildcardFqdnCustom(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallWildcardFqdnCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallWildcardFqdnCustom(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallWildcardFqdnCustom resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallWildcardFqdnCustomColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallWildcardFqdnCustomComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallWildcardFqdnCustomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallWildcardFqdnCustomUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallWildcardFqdnCustomVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallWildcardFqdnCustomWildcardFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallWildcardFqdnCustom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("color", flattenObjectFirewallWildcardFqdnCustomColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "ObjectFirewallWildcardFqdnCustom-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectFirewallWildcardFqdnCustomComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectFirewallWildcardFqdnCustom-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallWildcardFqdnCustomName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallWildcardFqdnCustom-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectFirewallWildcardFqdnCustomUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectFirewallWildcardFqdnCustom-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("visibility", flattenObjectFirewallWildcardFqdnCustomVisibility(o["visibility"], d, "visibility")); err != nil {
		if vv, ok := fortiAPIPatch(o["visibility"], "ObjectFirewallWildcardFqdnCustom-Visibility"); ok {
			if err = d.Set("visibility", vv); err != nil {
				return fmt.Errorf("Error reading visibility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading visibility: %v", err)
		}
	}

	if err = d.Set("wildcard_fqdn", flattenObjectFirewallWildcardFqdnCustomWildcardFqdn(o["wildcard-fqdn"], d, "wildcard_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard-fqdn"], "ObjectFirewallWildcardFqdnCustom-WildcardFqdn"); ok {
			if err = d.Set("wildcard_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading wildcard_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard_fqdn: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallWildcardFqdnCustomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallWildcardFqdnCustomColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallWildcardFqdnCustomComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallWildcardFqdnCustomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallWildcardFqdnCustomUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallWildcardFqdnCustomVisibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallWildcardFqdnCustomWildcardFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallWildcardFqdnCustom(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandObjectFirewallWildcardFqdnCustomColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectFirewallWildcardFqdnCustomComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallWildcardFqdnCustomName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandObjectFirewallWildcardFqdnCustomUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("visibility"); ok || d.HasChange("visibility") {
		t, err := expandObjectFirewallWildcardFqdnCustomVisibility(d, v, "visibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["visibility"] = t
		}
	}

	if v, ok := d.GetOk("wildcard_fqdn"); ok || d.HasChange("wildcard_fqdn") {
		t, err := expandObjectFirewallWildcardFqdnCustomWildcardFqdn(d, v, "wildcard_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard-fqdn"] = t
		}
	}

	return &obj, nil
}
