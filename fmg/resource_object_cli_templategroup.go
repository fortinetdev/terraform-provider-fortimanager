// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectCli TemplateGroup

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectCliTemplateGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCliTemplateGroupCreate,
		Read:   resourceObjectCliTemplateGroupRead,
		Update: resourceObjectCliTemplateGroupUpdate,
		Delete: resourceObjectCliTemplateGroupDelete,

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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"modification_time": &schema.Schema{
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
		},
	}
}

func resourceObjectCliTemplateGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectCliTemplateGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCliTemplateGroup resource while getting object: %v", err)
	}

	_, err = c.CreateObjectCliTemplateGroup(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectCliTemplateGroup resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCliTemplateGroupRead(d, m)
}

func resourceObjectCliTemplateGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectCliTemplateGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCliTemplateGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectCliTemplateGroup(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCliTemplateGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCliTemplateGroupRead(d, m)
}

func resourceObjectCliTemplateGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectCliTemplateGroup(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCliTemplateGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCliTemplateGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectCliTemplateGroup(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCliTemplateGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCliTemplateGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCliTemplateGroup resource from API: %v", err)
	}
	return nil
}

func flattenObjectCliTemplateGroupDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCliTemplateGroupMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCliTemplateGroupModificationTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCliTemplateGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCliTemplateGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("description", flattenObjectCliTemplateGroupDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectCliTemplateGroup-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("member", flattenObjectCliTemplateGroupMember(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "ObjectCliTemplateGroup-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("modification_time", flattenObjectCliTemplateGroupModificationTime(o["modification-time"], d, "modification_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["modification-time"], "ObjectCliTemplateGroup-ModificationTime"); ok {
			if err = d.Set("modification_time", vv); err != nil {
				return fmt.Errorf("Error reading modification_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modification_time: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCliTemplateGroupName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCliTemplateGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectCliTemplateGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCliTemplateGroupDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCliTemplateGroupMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectCliTemplateGroupModificationTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCliTemplateGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCliTemplateGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok {
		t, err := expandObjectCliTemplateGroupDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok {
		t, err := expandObjectCliTemplateGroupMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("modification_time"); ok {
		t, err := expandObjectCliTemplateGroupModificationTime(d, v, "modification_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modification-time"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectCliTemplateGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
