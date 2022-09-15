// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Wireless Termination Points (WTP) system log server profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerSyslogProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerSyslogProfileCreate,
		Read:   resourceObjectWirelessControllerSyslogProfileRead,
		Update: resourceObjectWirelessControllerSyslogProfileUpdate,
		Delete: resourceObjectWirelessControllerSyslogProfileDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"server_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerSyslogProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerSyslogProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerSyslogProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerSyslogProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerSyslogProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerSyslogProfileRead(d, m)
}

func resourceObjectWirelessControllerSyslogProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerSyslogProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerSyslogProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerSyslogProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerSyslogProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerSyslogProfileRead(d, m)
}

func resourceObjectWirelessControllerSyslogProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerSyslogProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerSyslogProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerSyslogProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerSyslogProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerSyslogProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerSyslogProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerSyslogProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerSyslogProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerSyslogProfileLogLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerSyslogProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerSyslogProfileServerAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerSyslogProfileServerFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerSyslogProfileServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerSyslogProfileServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerSyslogProfileServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerSyslogProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comment", flattenObjectWirelessControllerSyslogProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWirelessControllerSyslogProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("log_level", flattenObjectWirelessControllerSyslogProfileLogLevel(o["log-level"], d, "log_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-level"], "ObjectWirelessControllerSyslogProfile-LogLevel"); ok {
			if err = d.Set("log_level", vv); err != nil {
				return fmt.Errorf("Error reading log_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_level: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerSyslogProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerSyslogProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server_addr_type", flattenObjectWirelessControllerSyslogProfileServerAddrType(o["server-addr-type"], d, "server_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-addr-type"], "ObjectWirelessControllerSyslogProfile-ServerAddrType"); ok {
			if err = d.Set("server_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading server_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_addr_type: %v", err)
		}
	}

	if err = d.Set("server_fqdn", flattenObjectWirelessControllerSyslogProfileServerFqdn(o["server-fqdn"], d, "server_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-fqdn"], "ObjectWirelessControllerSyslogProfile-ServerFqdn"); ok {
			if err = d.Set("server_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading server_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_fqdn: %v", err)
		}
	}

	if err = d.Set("server_ip", flattenObjectWirelessControllerSyslogProfileServerIp(o["server-ip"], d, "server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-ip"], "ObjectWirelessControllerSyslogProfile-ServerIp"); ok {
			if err = d.Set("server_ip", vv); err != nil {
				return fmt.Errorf("Error reading server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_ip: %v", err)
		}
	}

	if err = d.Set("server_port", flattenObjectWirelessControllerSyslogProfileServerPort(o["server-port"], d, "server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-port"], "ObjectWirelessControllerSyslogProfile-ServerPort"); ok {
			if err = d.Set("server_port", vv); err != nil {
				return fmt.Errorf("Error reading server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_port: %v", err)
		}
	}

	if err = d.Set("server_status", flattenObjectWirelessControllerSyslogProfileServerStatus(o["server-status"], d, "server_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-status"], "ObjectWirelessControllerSyslogProfile-ServerStatus"); ok {
			if err = d.Set("server_status", vv); err != nil {
				return fmt.Errorf("Error reading server_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_status: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerSyslogProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerSyslogProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerSyslogProfileLogLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerSyslogProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerSyslogProfileServerAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerSyslogProfileServerFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerSyslogProfileServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerSyslogProfileServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerSyslogProfileServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerSyslogProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectWirelessControllerSyslogProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("log_level"); ok || d.HasChange("log_level") {
		t, err := expandObjectWirelessControllerSyslogProfileLogLevel(d, v, "log_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-level"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerSyslogProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server_addr_type"); ok || d.HasChange("server_addr_type") {
		t, err := expandObjectWirelessControllerSyslogProfileServerAddrType(d, v, "server_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("server_fqdn"); ok || d.HasChange("server_fqdn") {
		t, err := expandObjectWirelessControllerSyslogProfileServerFqdn(d, v, "server_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-fqdn"] = t
		}
	}

	if v, ok := d.GetOk("server_ip"); ok || d.HasChange("server_ip") {
		t, err := expandObjectWirelessControllerSyslogProfileServerIp(d, v, "server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-ip"] = t
		}
	}

	if v, ok := d.GetOk("server_port"); ok || d.HasChange("server_port") {
		t, err := expandObjectWirelessControllerSyslogProfileServerPort(d, v, "server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-port"] = t
		}
	}

	if v, ok := d.GetOk("server_status"); ok || d.HasChange("server_status") {
		t, err := expandObjectWirelessControllerSyslogProfileServerStatus(d, v, "server_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-status"] = t
		}
	}

	return &obj, nil
}
