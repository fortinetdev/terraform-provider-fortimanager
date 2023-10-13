// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Group members.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAdminGroupMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminGroupMemberCreate,
		Read:   resourceSystemAdminGroupMemberRead,
		Update: resourceSystemAdminGroupMemberUpdate,
		Delete: resourceSystemAdminGroupMemberDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemAdminGroupMemberCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	group := d.Get("group").(string)
	paradict["group"] = group

	obj, err := getObjectSystemAdminGroupMember(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminGroupMember resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAdminGroupMember(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemAdminGroupMember resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAdminGroupMemberRead(d, m)
}

func resourceSystemAdminGroupMemberUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	group := d.Get("group").(string)
	paradict["group"] = group

	obj, err := getObjectSystemAdminGroupMember(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminGroupMember resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAdminGroupMember(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminGroupMember resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAdminGroupMemberRead(d, m)
}

func resourceSystemAdminGroupMemberDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	group := d.Get("group").(string)
	paradict["group"] = group

	err = c.DeleteSystemAdminGroupMember(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdminGroupMember resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminGroupMemberRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	group := d.Get("group").(string)
	if group == "" {
		group = importOptionChecking(m.(*FortiClient).Cfg, "group")
		if group == "" {
			return fmt.Errorf("Parameter group is missing")
		}
		if err = d.Set("group", group); err != nil {
			return fmt.Errorf("Error set params group: %v", err)
		}
	}
	paradict["group"] = group

	o, err := c.ReadSystemAdminGroupMember(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminGroupMember resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdminGroupMember(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminGroupMember resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminGroupMemberName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdminGroupMember(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemAdminGroupMemberName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemAdminGroupMember-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminGroupMemberFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminGroupMemberName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdminGroupMember(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemAdminGroupMemberName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
