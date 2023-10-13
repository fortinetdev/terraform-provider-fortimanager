// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Default network service entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectApplicationListDefaultNetworkServices() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectApplicationListDefaultNetworkServicesCreate,
		Read:   resourceObjectApplicationListDefaultNetworkServicesRead,
		Update: resourceObjectApplicationListDefaultNetworkServicesUpdate,
		Delete: resourceObjectApplicationListDefaultNetworkServicesDelete,

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
			"list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"services": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"violation_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectApplicationListDefaultNetworkServicesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	list := d.Get("list").(string)
	paradict["list"] = list

	obj, err := getObjectObjectApplicationListDefaultNetworkServices(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectApplicationListDefaultNetworkServices resource while getting object: %v", err)
	}

	_, err = c.CreateObjectApplicationListDefaultNetworkServices(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectApplicationListDefaultNetworkServices resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectApplicationListDefaultNetworkServicesRead(d, m)
}

func resourceObjectApplicationListDefaultNetworkServicesUpdate(d *schema.ResourceData, m interface{}) error {
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

	list := d.Get("list").(string)
	paradict["list"] = list

	obj, err := getObjectObjectApplicationListDefaultNetworkServices(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectApplicationListDefaultNetworkServices resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectApplicationListDefaultNetworkServices(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectApplicationListDefaultNetworkServices resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectApplicationListDefaultNetworkServicesRead(d, m)
}

func resourceObjectApplicationListDefaultNetworkServicesDelete(d *schema.ResourceData, m interface{}) error {
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

	list := d.Get("list").(string)
	paradict["list"] = list

	err = c.DeleteObjectApplicationListDefaultNetworkServices(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectApplicationListDefaultNetworkServices resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectApplicationListDefaultNetworkServicesRead(d *schema.ResourceData, m interface{}) error {
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

	list := d.Get("list").(string)
	if list == "" {
		list = importOptionChecking(m.(*FortiClient).Cfg, "list")
		if list == "" {
			return fmt.Errorf("Parameter list is missing")
		}
		if err = d.Set("list", list); err != nil {
			return fmt.Errorf("Error set params list: %v", err)
		}
	}
	paradict["list"] = list

	o, err := c.ReadObjectApplicationListDefaultNetworkServices(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectApplicationListDefaultNetworkServices resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectApplicationListDefaultNetworkServices(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectApplicationListDefaultNetworkServices resource from API: %v", err)
	}
	return nil
}

func flattenObjectApplicationListDefaultNetworkServicesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListDefaultNetworkServicesPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListDefaultNetworkServicesServices2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectApplicationListDefaultNetworkServicesViolationAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectApplicationListDefaultNetworkServices(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenObjectApplicationListDefaultNetworkServicesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectApplicationListDefaultNetworkServices-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectApplicationListDefaultNetworkServicesPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectApplicationListDefaultNetworkServices-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("services", flattenObjectApplicationListDefaultNetworkServicesServices2edl(o["services"], d, "services")); err != nil {
		if vv, ok := fortiAPIPatch(o["services"], "ObjectApplicationListDefaultNetworkServices-Services"); ok {
			if err = d.Set("services", vv); err != nil {
				return fmt.Errorf("Error reading services: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading services: %v", err)
		}
	}

	if err = d.Set("violation_action", flattenObjectApplicationListDefaultNetworkServicesViolationAction2edl(o["violation-action"], d, "violation_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["violation-action"], "ObjectApplicationListDefaultNetworkServices-ViolationAction"); ok {
			if err = d.Set("violation_action", vv); err != nil {
				return fmt.Errorf("Error reading violation_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading violation_action: %v", err)
		}
	}

	return nil
}

func flattenObjectApplicationListDefaultNetworkServicesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectApplicationListDefaultNetworkServicesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListDefaultNetworkServicesPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListDefaultNetworkServicesServices2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListDefaultNetworkServicesViolationAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectApplicationListDefaultNetworkServices(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectApplicationListDefaultNetworkServicesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectApplicationListDefaultNetworkServicesPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("services"); ok || d.HasChange("services") {
		t, err := expandObjectApplicationListDefaultNetworkServicesServices2edl(d, v, "services")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["services"] = t
		}
	}

	if v, ok := d.GetOk("violation_action"); ok || d.HasChange("violation_action") {
		t, err := expandObjectApplicationListDefaultNetworkServicesViolationAction2edl(d, v, "violation_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["violation-action"] = t
		}
	}

	return &obj, nil
}
