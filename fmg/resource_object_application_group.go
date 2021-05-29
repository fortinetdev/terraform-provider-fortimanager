// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure firewall application groups.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectApplicationGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectApplicationGroupCreate,
		Read:   resourceObjectApplicationGroupRead,
		Update: resourceObjectApplicationGroupUpdate,
		Delete: resourceObjectApplicationGroupDelete,

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
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectApplicationGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectApplicationGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectApplicationGroup resource while getting object: %v", err)
	}

	_, err = c.CreateObjectApplicationGroup(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectApplicationGroup resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectApplicationGroupRead(d, m)
}

func resourceObjectApplicationGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectApplicationGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectApplicationGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectApplicationGroup(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectApplicationGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectApplicationGroupRead(d, m)
}

func resourceObjectApplicationGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectApplicationGroup(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectApplicationGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectApplicationGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectApplicationGroup(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectApplicationGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectApplicationGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectApplicationGroup resource from API: %v", err)
	}
	return nil
}

func flattenObjectApplicationGroupApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectApplicationGroupCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectApplicationGroupComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationGroupType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "application",
			1: "category",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectObjectApplicationGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("application", flattenObjectApplicationGroupApplication(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "ObjectApplicationGroup-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("category", flattenObjectApplicationGroupCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "ObjectApplicationGroup-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectApplicationGroupComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectApplicationGroup-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectApplicationGroupName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectApplicationGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectApplicationGroupType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectApplicationGroup-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectApplicationGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectApplicationGroupApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationGroupCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationGroupComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationGroupType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectApplicationGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("application"); ok {
		t, err := expandObjectApplicationGroupApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok {
		t, err := expandObjectApplicationGroupCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectApplicationGroupComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectApplicationGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandObjectApplicationGroupType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
