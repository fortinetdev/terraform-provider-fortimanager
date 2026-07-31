// Copyright 2026 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu), Yue Wang (@yuew-ftnt)

// Description: Device group table.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDvmdbGroupObjectMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmdbGroupObjectMemberCreate,
		Read:   resourceDvmdbGroupObjectMemberRead,
		Update: resourceDvmdbGroupObjectMemberUpdate,
		Delete: resourceDvmdbGroupObjectMemberDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"scopetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "inherit",
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
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDvmdbGroupObjectMemberCreate(d *schema.ResourceData, m interface{}) error {
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

	group := d.Get("group").(string)
	paradict["group"] = group

	obj, err := getObjectDvmdbGroupObjectMember(d)
	if err != nil {
		return fmt.Errorf("Error creating DvmdbGroupObjectMember resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateDvmdbGroupObjectMember(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating DvmdbGroupObjectMember resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceDvmdbGroupObjectMemberRead(d, m)
}

func resourceDvmdbGroupObjectMemberUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDvmdbGroupObjectMemberDelete(d *schema.ResourceData, m interface{}) error {
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

	group := d.Get("group").(string)
	paradict["group"] = group
	obj, err := getObjectDvmdbGroupObjectMember(d)
	if err != nil {
		return fmt.Errorf("Error creating DvmdbGroupObjectMember resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.DeleteDvmdbGroupObjectMember(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DvmdbGroupObjectMember resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDvmdbGroupObjectMemberRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenDvmdbGroupObjectMemberName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbGroupObjectMemberVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDvmdbGroupObjectMember(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenDvmdbGroupObjectMemberName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DvmdbGroupObjectMember-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vdom", flattenDvmdbGroupObjectMemberVdom2edl(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "DvmdbGroupObjectMember-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	return nil
}

func flattenDvmdbGroupObjectMemberFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmdbGroupObjectMemberName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbGroupObjectMemberVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDvmdbGroupObjectMember(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDvmdbGroupObjectMemberName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandDvmdbGroupObjectMemberVdom2edl(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
