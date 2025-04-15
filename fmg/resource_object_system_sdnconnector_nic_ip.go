// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IP configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemSdnConnectorNicIp() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemSdnConnectorNicIpCreate,
		Read:   resourceObjectSystemSdnConnectorNicIpRead,
		Update: resourceObjectSystemSdnConnectorNicIpUpdate,
		Delete: resourceObjectSystemSdnConnectorNicIpDelete,

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
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"nic": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"private_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemSdnConnectorNicIpCreate(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	nic := d.Get("nic").(string)
	paradict["sdn_connector"] = sdn_connector
	paradict["nic"] = nic

	obj, err := getObjectObjectSystemSdnConnectorNicIp(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnectorNicIp resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSystemSdnConnectorNicIp(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnectorNicIp resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemSdnConnectorNicIpRead(d, m)
}

func resourceObjectSystemSdnConnectorNicIpUpdate(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	nic := d.Get("nic").(string)
	paradict["sdn_connector"] = sdn_connector
	paradict["nic"] = nic

	obj, err := getObjectObjectSystemSdnConnectorNicIp(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnectorNicIp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemSdnConnectorNicIp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnectorNicIp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemSdnConnectorNicIpRead(d, m)
}

func resourceObjectSystemSdnConnectorNicIpDelete(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	nic := d.Get("nic").(string)
	paradict["sdn_connector"] = sdn_connector
	paradict["nic"] = nic

	wsParams["adom"] = adomv

	err = c.DeleteObjectSystemSdnConnectorNicIp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemSdnConnectorNicIp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemSdnConnectorNicIpRead(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	nic := d.Get("nic").(string)
	if sdn_connector == "" {
		sdn_connector = importOptionChecking(m.(*FortiClient).Cfg, "sdn_connector")
		if sdn_connector == "" {
			return fmt.Errorf("Parameter sdn_connector is missing")
		}
		if err = d.Set("sdn_connector", sdn_connector); err != nil {
			return fmt.Errorf("Error set params sdn_connector: %v", err)
		}
	}
	if nic == "" {
		nic = importOptionChecking(m.(*FortiClient).Cfg, "nic")
		if nic == "" {
			return fmt.Errorf("Parameter nic is missing")
		}
		if err = d.Set("nic", nic); err != nil {
			return fmt.Errorf("Error set params nic: %v", err)
		}
	}
	paradict["sdn_connector"] = sdn_connector
	paradict["nic"] = nic

	o, err := c.ReadObjectSystemSdnConnectorNicIp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnectorNicIp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemSdnConnectorNicIp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnectorNicIp resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemSdnConnectorNicIpName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorNicIpPrivateIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorNicIpPublicIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorNicIpResourceGroup3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemSdnConnectorNicIp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectSystemSdnConnectorNicIpName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemSdnConnectorNicIp-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("private_ip", flattenObjectSystemSdnConnectorNicIpPrivateIp3rdl(o["private-ip"], d, "private_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-ip"], "ObjectSystemSdnConnectorNicIp-PrivateIp"); ok {
			if err = d.Set("private_ip", vv); err != nil {
				return fmt.Errorf("Error reading private_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_ip: %v", err)
		}
	}

	if err = d.Set("public_ip", flattenObjectSystemSdnConnectorNicIpPublicIp3rdl(o["public-ip"], d, "public_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["public-ip"], "ObjectSystemSdnConnectorNicIp-PublicIp"); ok {
			if err = d.Set("public_ip", vv); err != nil {
				return fmt.Errorf("Error reading public_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading public_ip: %v", err)
		}
	}

	if err = d.Set("resource_group", flattenObjectSystemSdnConnectorNicIpResourceGroup3rdl(o["resource-group"], d, "resource_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["resource-group"], "ObjectSystemSdnConnectorNicIp-ResourceGroup"); ok {
			if err = d.Set("resource_group", vv); err != nil {
				return fmt.Errorf("Error reading resource_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resource_group: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemSdnConnectorNicIpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemSdnConnectorNicIpName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorNicIpPrivateIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorNicIpPublicIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorNicIpResourceGroup3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemSdnConnectorNicIp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSystemSdnConnectorNicIpName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("private_ip"); ok || d.HasChange("private_ip") {
		t, err := expandObjectSystemSdnConnectorNicIpPrivateIp3rdl(d, v, "private_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-ip"] = t
		}
	}

	if v, ok := d.GetOk("public_ip"); ok || d.HasChange("public_ip") {
		t, err := expandObjectSystemSdnConnectorNicIpPublicIp3rdl(d, v, "public_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["public-ip"] = t
		}
	}

	if v, ok := d.GetOk("resource_group"); ok || d.HasChange("resource_group") {
		t, err := expandObjectSystemSdnConnectorNicIpResourceGroup3rdl(d, v, "resource_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource-group"] = t
		}
	}

	return &obj, nil
}
