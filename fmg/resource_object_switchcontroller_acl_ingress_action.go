// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ACL actions.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerAclIngressAction() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerAclIngressActionUpdate,
		Read:   resourceObjectSwitchControllerAclIngressActionRead,
		Update: resourceObjectSwitchControllerAclIngressActionUpdate,
		Delete: resourceObjectSwitchControllerAclIngressActionDelete,

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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fmgcount": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"drop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSwitchControllerAclIngressActionUpdate(d *schema.ResourceData, m interface{}) error {
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

	ingress := d.Get("ingress").(string)
	paradict["ingress"] = ingress

	obj, err := getObjectObjectSwitchControllerAclIngressAction(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerAclIngressAction resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSwitchControllerAclIngressAction(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerAclIngressAction resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSwitchControllerAclIngressAction")

	return resourceObjectSwitchControllerAclIngressActionRead(d, m)
}

func resourceObjectSwitchControllerAclIngressActionDelete(d *schema.ResourceData, m interface{}) error {
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

	ingress := d.Get("ingress").(string)
	paradict["ingress"] = ingress

	wsParams["adom"] = adomv

	err = c.DeleteObjectSwitchControllerAclIngressAction(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerAclIngressAction resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerAclIngressActionRead(d *schema.ResourceData, m interface{}) error {
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

	ingress := d.Get("ingress").(string)
	if ingress == "" {
		ingress = importOptionChecking(m.(*FortiClient).Cfg, "ingress")
		if ingress == "" {
			return fmt.Errorf("Parameter ingress is missing")
		}
		if err = d.Set("ingress", ingress); err != nil {
			return fmt.Errorf("Error set params ingress: %v", err)
		}
	}
	paradict["ingress"] = ingress

	o, err := c.ReadObjectSwitchControllerAclIngressAction(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerAclIngressAction resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerAclIngressAction(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerAclIngressAction resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerAclIngressActionCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerAclIngressActionDrop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerAclIngressAction(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fmgcount", flattenObjectSwitchControllerAclIngressActionCount2edl(o["count"], d, "fmgcount")); err != nil {
		if vv, ok := fortiAPIPatch(o["count"], "ObjectSwitchControllerAclIngressAction-Count"); ok {
			if err = d.Set("fmgcount", vv); err != nil {
				return fmt.Errorf("Error reading fmgcount: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgcount: %v", err)
		}
	}

	if err = d.Set("drop", flattenObjectSwitchControllerAclIngressActionDrop2edl(o["drop"], d, "drop")); err != nil {
		if vv, ok := fortiAPIPatch(o["drop"], "ObjectSwitchControllerAclIngressAction-Drop"); ok {
			if err = d.Set("drop", vv); err != nil {
				return fmt.Errorf("Error reading drop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drop: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerAclIngressActionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerAclIngressActionCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerAclIngressActionDrop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerAclIngressAction(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgcount"); ok || d.HasChange("fmgcount") {
		t, err := expandObjectSwitchControllerAclIngressActionCount2edl(d, v, "fmgcount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["count"] = t
		}
	}

	if v, ok := d.GetOk("drop"); ok || d.HasChange("drop") {
		t, err := expandObjectSwitchControllerAclIngressActionDrop2edl(d, v, "drop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop"] = t
		}
	}

	return &obj, nil
}
