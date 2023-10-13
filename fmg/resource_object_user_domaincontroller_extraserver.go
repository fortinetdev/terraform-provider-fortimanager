// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: extra servers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserDomainControllerExtraServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserDomainControllerExtraServerCreate,
		Read:   resourceObjectUserDomainControllerExtraServerRead,
		Update: resourceObjectUserDomainControllerExtraServerUpdate,
		Delete: resourceObjectUserDomainControllerExtraServerDelete,

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
			"domain_controller": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"source_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectUserDomainControllerExtraServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	domain_controller := d.Get("domain_controller").(string)
	paradict["domain_controller"] = domain_controller

	obj, err := getObjectObjectUserDomainControllerExtraServer(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserDomainControllerExtraServer resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserDomainControllerExtraServer(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserDomainControllerExtraServer resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectUserDomainControllerExtraServerRead(d, m)
}

func resourceObjectUserDomainControllerExtraServerUpdate(d *schema.ResourceData, m interface{}) error {
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

	domain_controller := d.Get("domain_controller").(string)
	paradict["domain_controller"] = domain_controller

	obj, err := getObjectObjectUserDomainControllerExtraServer(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserDomainControllerExtraServer resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserDomainControllerExtraServer(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserDomainControllerExtraServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectUserDomainControllerExtraServerRead(d, m)
}

func resourceObjectUserDomainControllerExtraServerDelete(d *schema.ResourceData, m interface{}) error {
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

	domain_controller := d.Get("domain_controller").(string)
	paradict["domain_controller"] = domain_controller

	err = c.DeleteObjectUserDomainControllerExtraServer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserDomainControllerExtraServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserDomainControllerExtraServerRead(d *schema.ResourceData, m interface{}) error {
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

	domain_controller := d.Get("domain_controller").(string)
	if domain_controller == "" {
		domain_controller = importOptionChecking(m.(*FortiClient).Cfg, "domain_controller")
		if domain_controller == "" {
			return fmt.Errorf("Parameter domain_controller is missing")
		}
		if err = d.Set("domain_controller", domain_controller); err != nil {
			return fmt.Errorf("Error set params domain_controller: %v", err)
		}
	}
	paradict["domain_controller"] = domain_controller

	o, err := c.ReadObjectUserDomainControllerExtraServer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserDomainControllerExtraServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserDomainControllerExtraServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserDomainControllerExtraServer resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserDomainControllerExtraServerId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerExtraServerIpAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerExtraServerPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerExtraServerSourceIpAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerExtraServerSourcePort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserDomainControllerExtraServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenObjectUserDomainControllerExtraServerId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectUserDomainControllerExtraServer-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_address", flattenObjectUserDomainControllerExtraServerIpAddress2edl(o["ip-address"], d, "ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-address"], "ObjectUserDomainControllerExtraServer-IpAddress"); ok {
			if err = d.Set("ip_address", vv); err != nil {
				return fmt.Errorf("Error reading ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_address: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectUserDomainControllerExtraServerPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectUserDomainControllerExtraServer-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("source_ip_address", flattenObjectUserDomainControllerExtraServerSourceIpAddress2edl(o["source-ip-address"], d, "source_ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-address"], "ObjectUserDomainControllerExtraServer-SourceIpAddress"); ok {
			if err = d.Set("source_ip_address", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_address: %v", err)
		}
	}

	if err = d.Set("source_port", flattenObjectUserDomainControllerExtraServerSourcePort2edl(o["source-port"], d, "source_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-port"], "ObjectUserDomainControllerExtraServer-SourcePort"); ok {
			if err = d.Set("source_port", vv); err != nil {
				return fmt.Errorf("Error reading source_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_port: %v", err)
		}
	}

	return nil
}

func flattenObjectUserDomainControllerExtraServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserDomainControllerExtraServerId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerExtraServerIpAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerExtraServerPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerExtraServerSourceIpAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerExtraServerSourcePort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserDomainControllerExtraServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectUserDomainControllerExtraServerId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_address"); ok || d.HasChange("ip_address") {
		t, err := expandObjectUserDomainControllerExtraServerIpAddress2edl(d, v, "ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-address"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectUserDomainControllerExtraServerPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_address"); ok || d.HasChange("source_ip_address") {
		t, err := expandObjectUserDomainControllerExtraServerSourceIpAddress2edl(d, v, "source_ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-address"] = t
		}
	}

	if v, ok := d.GetOk("source_port"); ok || d.HasChange("source_port") {
		t, err := expandObjectUserDomainControllerExtraServerSourcePort2edl(d, v, "source_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-port"] = t
		}
	}

	return &obj, nil
}
