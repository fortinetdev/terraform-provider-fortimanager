// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Define known domain controller servers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectCredentialStoreDomainController() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCredentialStoreDomainControllerCreate,
		Read:   resourceObjectCredentialStoreDomainControllerRead,
		Update: resourceObjectCredentialStoreDomainControllerUpdate,
		Delete: resourceObjectCredentialStoreDomainControllerDelete,

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
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectCredentialStoreDomainControllerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectCredentialStoreDomainController(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCredentialStoreDomainController resource while getting object: %v", err)
	}

	_, err = c.CreateObjectCredentialStoreDomainController(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectCredentialStoreDomainController resource: %v", err)
	}

	d.SetId(getStringKey(d, "server_name"))

	return resourceObjectCredentialStoreDomainControllerRead(d, m)
}

func resourceObjectCredentialStoreDomainControllerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectCredentialStoreDomainController(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCredentialStoreDomainController resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectCredentialStoreDomainController(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCredentialStoreDomainController resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "server_name"))

	return resourceObjectCredentialStoreDomainControllerRead(d, m)
}

func resourceObjectCredentialStoreDomainControllerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectCredentialStoreDomainController(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCredentialStoreDomainController resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCredentialStoreDomainControllerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectCredentialStoreDomainController(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCredentialStoreDomainController resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCredentialStoreDomainController(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCredentialStoreDomainController resource from API: %v", err)
	}
	return nil
}

func flattenObjectCredentialStoreDomainControllerDomainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCredentialStoreDomainControllerHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCredentialStoreDomainControllerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCredentialStoreDomainControllerIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCredentialStoreDomainControllerPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCredentialStoreDomainControllerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCredentialStoreDomainControllerServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCredentialStoreDomainControllerUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCredentialStoreDomainController(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("domain_name", flattenObjectCredentialStoreDomainControllerDomainName(o["domain-name"], d, "domain_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-name"], "ObjectCredentialStoreDomainController-DomainName"); ok {
			if err = d.Set("domain_name", vv); err != nil {
				return fmt.Errorf("Error reading domain_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if err = d.Set("hostname", flattenObjectCredentialStoreDomainControllerHostname(o["hostname"], d, "hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname"], "ObjectCredentialStoreDomainController-Hostname"); ok {
			if err = d.Set("hostname", vv); err != nil {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("ip", flattenObjectCredentialStoreDomainControllerIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectCredentialStoreDomainController-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ip6", flattenObjectCredentialStoreDomainControllerIp6(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "ObjectCredentialStoreDomainController-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("password", flattenObjectCredentialStoreDomainControllerPassword(o["password"], d, "password")); err != nil {
		if vv, ok := fortiAPIPatch(o["password"], "ObjectCredentialStoreDomainController-Password"); ok {
			if err = d.Set("password", vv); err != nil {
				return fmt.Errorf("Error reading password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectCredentialStoreDomainControllerPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectCredentialStoreDomainController-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("server_name", flattenObjectCredentialStoreDomainControllerServerName(o["server-name"], d, "server_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-name"], "ObjectCredentialStoreDomainController-ServerName"); ok {
			if err = d.Set("server_name", vv); err != nil {
				return fmt.Errorf("Error reading server_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	if err = d.Set("username", flattenObjectCredentialStoreDomainControllerUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "ObjectCredentialStoreDomainController-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	return nil
}

func flattenObjectCredentialStoreDomainControllerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCredentialStoreDomainControllerDomainName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCredentialStoreDomainControllerHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCredentialStoreDomainControllerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCredentialStoreDomainControllerIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCredentialStoreDomainControllerPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCredentialStoreDomainControllerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCredentialStoreDomainControllerServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCredentialStoreDomainControllerUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCredentialStoreDomainController(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("domain_name"); ok || d.HasChange("domain_name") {
		t, err := expandObjectCredentialStoreDomainControllerDomainName(d, v, "domain_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok || d.HasChange("hostname") {
		t, err := expandObjectCredentialStoreDomainControllerHostname(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandObjectCredentialStoreDomainControllerIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok || d.HasChange("ip6") {
		t, err := expandObjectCredentialStoreDomainControllerIp6(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectCredentialStoreDomainControllerPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectCredentialStoreDomainControllerPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("server_name"); ok || d.HasChange("server_name") {
		t, err := expandObjectCredentialStoreDomainControllerServerName(d, v, "server_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-name"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandObjectCredentialStoreDomainControllerUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	return &obj, nil
}
