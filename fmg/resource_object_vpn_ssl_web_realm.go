// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Realm.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectVpnSslWebRealm() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnSslWebRealmCreate,
		Read:   resourceObjectVpnSslWebRealmRead,
		Update: resourceObjectVpnSslWebRealmUpdate,
		Delete: resourceObjectVpnSslWebRealmDelete,

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
			"login_page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_concurrent_user": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nas_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"radius_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_path": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"virtual_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_host_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectVpnSslWebRealmCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVpnSslWebRealm(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebRealm resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVpnSslWebRealm(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebRealm resource: %v", err)
	}

	d.SetId(getStringKey(d, "url_path"))

	return resourceObjectVpnSslWebRealmRead(d, m)
}

func resourceObjectVpnSslWebRealmUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVpnSslWebRealm(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebRealm resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVpnSslWebRealm(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebRealm resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "url_path"))

	return resourceObjectVpnSslWebRealmRead(d, m)
}

func resourceObjectVpnSslWebRealmDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectVpnSslWebRealm(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnSslWebRealm resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnSslWebRealmRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectVpnSslWebRealm(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebRealm resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnSslWebRealm(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebRealm resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnSslWebRealmLoginPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebRealmMaxConcurrentUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebRealmNasIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebRealmRadiusPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebRealmRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebRealmUrlPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebRealmVirtualHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebRealmVirtualHostOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectObjectVpnSslWebRealm(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("login_page", flattenObjectVpnSslWebRealmLoginPage(o["login-page"], d, "login_page")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-page"], "ObjectVpnSslWebRealm-LoginPage"); ok {
			if err = d.Set("login_page", vv); err != nil {
				return fmt.Errorf("Error reading login_page: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_page: %v", err)
		}
	}

	if err = d.Set("max_concurrent_user", flattenObjectVpnSslWebRealmMaxConcurrentUser(o["max-concurrent-user"], d, "max_concurrent_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-concurrent-user"], "ObjectVpnSslWebRealm-MaxConcurrentUser"); ok {
			if err = d.Set("max_concurrent_user", vv); err != nil {
				return fmt.Errorf("Error reading max_concurrent_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_concurrent_user: %v", err)
		}
	}

	if err = d.Set("nas_ip", flattenObjectVpnSslWebRealmNasIp(o["nas-ip"], d, "nas_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["nas-ip"], "ObjectVpnSslWebRealm-NasIp"); ok {
			if err = d.Set("nas_ip", vv); err != nil {
				return fmt.Errorf("Error reading nas_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nas_ip: %v", err)
		}
	}

	if err = d.Set("radius_port", flattenObjectVpnSslWebRealmRadiusPort(o["radius-port"], d, "radius_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-port"], "ObjectVpnSslWebRealm-RadiusPort"); ok {
			if err = d.Set("radius_port", vv); err != nil {
				return fmt.Errorf("Error reading radius_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_port: %v", err)
		}
	}

	if err = d.Set("radius_server", flattenObjectVpnSslWebRealmRadiusServer(o["radius-server"], d, "radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-server"], "ObjectVpnSslWebRealm-RadiusServer"); ok {
			if err = d.Set("radius_server", vv); err != nil {
				return fmt.Errorf("Error reading radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_server: %v", err)
		}
	}

	if err = d.Set("url_path", flattenObjectVpnSslWebRealmUrlPath(o["url-path"], d, "url_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-path"], "ObjectVpnSslWebRealm-UrlPath"); ok {
			if err = d.Set("url_path", vv); err != nil {
				return fmt.Errorf("Error reading url_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_path: %v", err)
		}
	}

	if err = d.Set("virtual_host", flattenObjectVpnSslWebRealmVirtualHost(o["virtual-host"], d, "virtual_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-host"], "ObjectVpnSslWebRealm-VirtualHost"); ok {
			if err = d.Set("virtual_host", vv); err != nil {
				return fmt.Errorf("Error reading virtual_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_host: %v", err)
		}
	}

	if err = d.Set("virtual_host_only", flattenObjectVpnSslWebRealmVirtualHostOnly(o["virtual-host-only"], d, "virtual_host_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-host-only"], "ObjectVpnSslWebRealm-VirtualHostOnly"); ok {
			if err = d.Set("virtual_host_only", vv); err != nil {
				return fmt.Errorf("Error reading virtual_host_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_host_only: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnSslWebRealmFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnSslWebRealmLoginPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebRealmMaxConcurrentUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebRealmNasIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebRealmRadiusPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebRealmRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebRealmUrlPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebRealmVirtualHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebRealmVirtualHostOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnSslWebRealm(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("login_page"); ok {
		t, err := expandObjectVpnSslWebRealmLoginPage(d, v, "login_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-page"] = t
		}
	}

	if v, ok := d.GetOk("max_concurrent_user"); ok {
		t, err := expandObjectVpnSslWebRealmMaxConcurrentUser(d, v, "max_concurrent_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-concurrent-user"] = t
		}
	}

	if v, ok := d.GetOk("nas_ip"); ok {
		t, err := expandObjectVpnSslWebRealmNasIp(d, v, "nas_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-ip"] = t
		}
	}

	if v, ok := d.GetOk("radius_port"); ok {
		t, err := expandObjectVpnSslWebRealmRadiusPort(d, v, "radius_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-port"] = t
		}
	}

	if v, ok := d.GetOk("radius_server"); ok {
		t, err := expandObjectVpnSslWebRealmRadiusServer(d, v, "radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server"] = t
		}
	}

	if v, ok := d.GetOk("url_path"); ok {
		t, err := expandObjectVpnSslWebRealmUrlPath(d, v, "url_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-path"] = t
		}
	}

	if v, ok := d.GetOk("virtual_host"); ok {
		t, err := expandObjectVpnSslWebRealmVirtualHost(d, v, "virtual_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-host"] = t
		}
	}

	if v, ok := d.GetOk("virtual_host_only"); ok {
		t, err := expandObjectVpnSslWebRealmVirtualHostOnly(d, v, "virtual_host_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-host-only"] = t
		}
	}

	return &obj, nil
}
