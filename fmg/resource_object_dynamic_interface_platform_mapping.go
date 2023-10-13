// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectDynamic InterfacePlatformMapping

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDynamicInterfacePlatformMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDynamicInterfacePlatformMappingCreate,
		Read:   resourceObjectDynamicInterfacePlatformMappingRead,
		Update: resourceObjectDynamicInterfacePlatformMappingUpdate,
		Delete: resourceObjectDynamicInterfacePlatformMappingDelete,

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
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"egress_shaping_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ingress_shaping_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"intf_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"intrazone_deny": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectDynamicInterfacePlatformMappingCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	var_interface := d.Get("interface").(string)
	paradict["interface"] = var_interface

	obj, err := getObjectObjectDynamicInterfacePlatformMapping(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicInterfacePlatformMapping resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDynamicInterfacePlatformMapping(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicInterfacePlatformMapping resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDynamicInterfacePlatformMappingRead(d, m)
}

func resourceObjectDynamicInterfacePlatformMappingUpdate(d *schema.ResourceData, m interface{}) error {
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

	var_interface := d.Get("interface").(string)
	paradict["interface"] = var_interface

	obj, err := getObjectObjectDynamicInterfacePlatformMapping(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicInterfacePlatformMapping resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDynamicInterfacePlatformMapping(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicInterfacePlatformMapping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDynamicInterfacePlatformMappingRead(d, m)
}

func resourceObjectDynamicInterfacePlatformMappingDelete(d *schema.ResourceData, m interface{}) error {
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

	var_interface := d.Get("interface").(string)
	paradict["interface"] = var_interface

	err = c.DeleteObjectDynamicInterfacePlatformMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDynamicInterfacePlatformMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDynamicInterfacePlatformMappingRead(d *schema.ResourceData, m interface{}) error {
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

	var_interface := d.Get("interface").(string)
	if var_interface == "" {
		var_interface = importOptionChecking(m.(*FortiClient).Cfg, "interface")
		if var_interface == "" {
			return fmt.Errorf("Parameter interface is missing")
		}
		if err = d.Set("interface", var_interface); err != nil {
			return fmt.Errorf("Error set params interface: %v", err)
		}
	}
	paradict["interface"] = var_interface

	o, err := c.ReadObjectDynamicInterfacePlatformMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicInterfacePlatformMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDynamicInterfacePlatformMapping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicInterfacePlatformMapping resource from API: %v", err)
	}
	return nil
}

func flattenObjectDynamicInterfacePlatformMappingEgressShapingProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfacePlatformMappingIngressShapingProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfacePlatformMappingIntfZone2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfacePlatformMappingIntrazoneDeny2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfacePlatformMappingName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDynamicInterfacePlatformMapping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("egress_shaping_profile", flattenObjectDynamicInterfacePlatformMappingEgressShapingProfile2edl(o["egress-shaping-profile"], d, "egress_shaping_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["egress-shaping-profile"], "ObjectDynamicInterfacePlatformMapping-EgressShapingProfile"); ok {
			if err = d.Set("egress_shaping_profile", vv); err != nil {
				return fmt.Errorf("Error reading egress_shaping_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading egress_shaping_profile: %v", err)
		}
	}

	if err = d.Set("ingress_shaping_profile", flattenObjectDynamicInterfacePlatformMappingIngressShapingProfile2edl(o["ingress-shaping-profile"], d, "ingress_shaping_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ingress-shaping-profile"], "ObjectDynamicInterfacePlatformMapping-IngressShapingProfile"); ok {
			if err = d.Set("ingress_shaping_profile", vv); err != nil {
				return fmt.Errorf("Error reading ingress_shaping_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ingress_shaping_profile: %v", err)
		}
	}

	if err = d.Set("intf_zone", flattenObjectDynamicInterfacePlatformMappingIntfZone2edl(o["intf-zone"], d, "intf_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["intf-zone"], "ObjectDynamicInterfacePlatformMapping-IntfZone"); ok {
			if err = d.Set("intf_zone", vv); err != nil {
				return fmt.Errorf("Error reading intf_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intf_zone: %v", err)
		}
	}

	if err = d.Set("intrazone_deny", flattenObjectDynamicInterfacePlatformMappingIntrazoneDeny2edl(o["intrazone-deny"], d, "intrazone_deny")); err != nil {
		if vv, ok := fortiAPIPatch(o["intrazone-deny"], "ObjectDynamicInterfacePlatformMapping-IntrazoneDeny"); ok {
			if err = d.Set("intrazone_deny", vv); err != nil {
				return fmt.Errorf("Error reading intrazone_deny: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intrazone_deny: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectDynamicInterfacePlatformMappingName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDynamicInterfacePlatformMapping-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectDynamicInterfacePlatformMappingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDynamicInterfacePlatformMappingEgressShapingProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfacePlatformMappingIngressShapingProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfacePlatformMappingIntfZone2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfacePlatformMappingIntrazoneDeny2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfacePlatformMappingName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDynamicInterfacePlatformMapping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("egress_shaping_profile"); ok || d.HasChange("egress_shaping_profile") {
		t, err := expandObjectDynamicInterfacePlatformMappingEgressShapingProfile2edl(d, v, "egress_shaping_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["egress-shaping-profile"] = t
		}
	}

	if v, ok := d.GetOk("ingress_shaping_profile"); ok || d.HasChange("ingress_shaping_profile") {
		t, err := expandObjectDynamicInterfacePlatformMappingIngressShapingProfile2edl(d, v, "ingress_shaping_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress-shaping-profile"] = t
		}
	}

	if v, ok := d.GetOk("intf_zone"); ok || d.HasChange("intf_zone") {
		t, err := expandObjectDynamicInterfacePlatformMappingIntfZone2edl(d, v, "intf_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intf-zone"] = t
		}
	}

	if v, ok := d.GetOk("intrazone_deny"); ok || d.HasChange("intrazone_deny") {
		t, err := expandObjectDynamicInterfacePlatformMappingIntrazoneDeny2edl(d, v, "intrazone_deny")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intrazone-deny"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectDynamicInterfacePlatformMappingName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
