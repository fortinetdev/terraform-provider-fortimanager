// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Physical interfaces that belong to the aggregate or redundant interface.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemInterfaceMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemInterfaceMemberCreate,
		Read:   resourceSystemInterfaceMemberRead,
		Update: resourceSystemInterfaceMemberUpdate,
		Delete: resourceSystemInterfaceMemberDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"interface_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemInterfaceMemberCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	var_interface := d.Get("interface").(string)
	paradict["interface"] = var_interface

	obj, err := getObjectSystemInterfaceMember(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceMember resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystemInterfaceMember(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceMember resource: %v", err)
	}

	d.SetId(getStringKey(d, "interface_name"))

	return resourceSystemInterfaceMemberRead(d, m)
}

func resourceSystemInterfaceMemberUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	var_interface := d.Get("interface").(string)
	paradict["interface"] = var_interface

	obj, err := getObjectSystemInterfaceMember(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceMember resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemInterfaceMember(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceMember resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "interface_name"))

	return resourceSystemInterfaceMemberRead(d, m)
}

func resourceSystemInterfaceMemberDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	var_interface := d.Get("interface").(string)
	paradict["interface"] = var_interface

	wsParams["adom"] = adomv

	err = c.DeleteSystemInterfaceMember(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterfaceMember resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfaceMemberRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
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

	o, err := c.ReadSystemInterfaceMember(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceMember resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterfaceMember(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceMember resource from API: %v", err)
	}
	return nil
}

func flattenSystemInterfaceMemberInterfaceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemInterfaceMember(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("interface_name", flattenSystemInterfaceMemberInterfaceName2edl(o["interface-name"], d, "interface_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-name"], "SystemInterfaceMember-InterfaceName"); ok {
			if err = d.Set("interface_name", vv); err != nil {
				return fmt.Errorf("Error reading interface_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_name: %v", err)
		}
	}

	return nil
}

func flattenSystemInterfaceMemberFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemInterfaceMemberInterfaceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterfaceMember(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("interface_name"); ok || d.HasChange("interface_name") {
		t, err := expandSystemInterfaceMemberInterfaceName2edl(d, v, "interface_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-name"] = t
		}
	}

	return &obj, nil
}
