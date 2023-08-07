// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure ACL groups to be applied on managed FortiSwitch ports.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerAclGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerAclGroupCreate,
		Read:   resourceObjectSwitchControllerAclGroupRead,
		Update: resourceObjectSwitchControllerAclGroupUpdate,
		Delete: resourceObjectSwitchControllerAclGroupDelete,

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
			"ingress": &schema.Schema{
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
		},
	}
}

func resourceObjectSwitchControllerAclGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSwitchControllerAclGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerAclGroup resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSwitchControllerAclGroup(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerAclGroup resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerAclGroupRead(d, m)
}

func resourceObjectSwitchControllerAclGroupUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSwitchControllerAclGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerAclGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerAclGroup(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerAclGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerAclGroupRead(d, m)
}

func resourceObjectSwitchControllerAclGroupDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSwitchControllerAclGroup(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerAclGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerAclGroupRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSwitchControllerAclGroup(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerAclGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerAclGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerAclGroup resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerAclGroupIngress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerAclGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerAclGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("ingress", flattenObjectSwitchControllerAclGroupIngress(o["ingress"], d, "ingress")); err != nil {
		if vv, ok := fortiAPIPatch(o["ingress"], "ObjectSwitchControllerAclGroup-Ingress"); ok {
			if err = d.Set("ingress", vv); err != nil {
				return fmt.Errorf("Error reading ingress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ingress: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerAclGroupName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerAclGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerAclGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerAclGroupIngress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerAclGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerAclGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ingress"); ok || d.HasChange("ingress") {
		t, err := expandObjectSwitchControllerAclGroupIngress(d, v, "ingress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerAclGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
