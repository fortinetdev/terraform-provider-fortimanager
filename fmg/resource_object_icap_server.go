// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure ICAP servers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectIcapServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectIcapServerCreate,
		Read:   resourceObjectIcapServerRead,
		Update: resourceObjectIcapServerUpdate,
		Delete: resourceObjectIcapServerDelete,

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
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"secure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectIcapServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectIcapServer(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIcapServer resource while getting object: %v", err)
	}

	_, err = c.CreateObjectIcapServer(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectIcapServer resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectIcapServerRead(d, m)
}

func resourceObjectIcapServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectIcapServer(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIcapServer resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectIcapServer(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIcapServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectIcapServerRead(d, m)
}

func resourceObjectIcapServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectIcapServer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectIcapServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectIcapServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectIcapServer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIcapServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectIcapServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIcapServer resource from API: %v", err)
	}
	return nil
}

func flattenObjectIcapServerIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapServerIpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapServerIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapServerMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapServerSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapServerSslCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectIcapServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("ip_address", flattenObjectIcapServerIpAddress(o["ip-address"], d, "ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-address"], "ObjectIcapServer-IpAddress"); ok {
			if err = d.Set("ip_address", vv); err != nil {
				return fmt.Errorf("Error reading ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_address: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenObjectIcapServerIpVersion(o["ip-version"], d, "ip_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-version"], "ObjectIcapServer-IpVersion"); ok {
			if err = d.Set("ip_version", vv); err != nil {
				return fmt.Errorf("Error reading ip_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("ip6_address", flattenObjectIcapServerIp6Address(o["ip6-address"], d, "ip6_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-address"], "ObjectIcapServer-Ip6Address"); ok {
			if err = d.Set("ip6_address", vv); err != nil {
				return fmt.Errorf("Error reading ip6_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_address: %v", err)
		}
	}

	if err = d.Set("max_connections", flattenObjectIcapServerMaxConnections(o["max-connections"], d, "max_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-connections"], "ObjectIcapServer-MaxConnections"); ok {
			if err = d.Set("max_connections", vv); err != nil {
				return fmt.Errorf("Error reading max_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_connections: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectIcapServerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectIcapServer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectIcapServerPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectIcapServer-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("secure", flattenObjectIcapServerSecure(o["secure"], d, "secure")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure"], "ObjectIcapServer-Secure"); ok {
			if err = d.Set("secure", vv); err != nil {
				return fmt.Errorf("Error reading secure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure: %v", err)
		}
	}

	if err = d.Set("ssl_cert", flattenObjectIcapServerSslCert(o["ssl-cert"], d, "ssl_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-cert"], "ObjectIcapServer-SslCert"); ok {
			if err = d.Set("ssl_cert", vv); err != nil {
				return fmt.Errorf("Error reading ssl_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_cert: %v", err)
		}
	}

	return nil
}

func flattenObjectIcapServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectIcapServerIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapServerIpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapServerIp6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapServerMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapServerSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapServerSslCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectIcapServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip_address"); ok || d.HasChange("ip_address") {
		t, err := expandObjectIcapServerIpAddress(d, v, "ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-address"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok || d.HasChange("ip_version") {
		t, err := expandObjectIcapServerIpVersion(d, v, "ip_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("ip6_address"); ok || d.HasChange("ip6_address") {
		t, err := expandObjectIcapServerIp6Address(d, v, "ip6_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-address"] = t
		}
	}

	if v, ok := d.GetOk("max_connections"); ok || d.HasChange("max_connections") {
		t, err := expandObjectIcapServerMaxConnections(d, v, "max_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-connections"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectIcapServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectIcapServerPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("secure"); ok || d.HasChange("secure") {
		t, err := expandObjectIcapServerSecure(d, v, "secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cert"); ok || d.HasChange("ssl_cert") {
		t, err := expandObjectIcapServerSslCert(d, v, "ssl_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cert"] = t
		}
	}

	return &obj, nil
}
