// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure MS Exchange server entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserExchange() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserExchangeCreate,
		Read:   resourceObjectUserExchangeRead,
		Update: resourceObjectUserExchangeUpdate,
		Delete: resourceObjectUserExchangeDelete,

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
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discover_kdc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connect_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"kdc_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectUserExchangeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectUserExchange(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserExchange resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserExchange(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserExchange resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserExchangeRead(d, m)
}

func resourceObjectUserExchangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserExchange(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserExchange resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserExchange(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserExchange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserExchangeRead(d, m)
}

func resourceObjectUserExchangeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectUserExchange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserExchange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserExchangeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectUserExchange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserExchange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserExchange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserExchange resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserExchangeAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserExchangeAuthLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserExchangeAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserExchangeAutoDiscoverKdc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserExchangeConnectProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserExchangeDomainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserExchangeHttpAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserExchangeIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserExchangeIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserExchangeKdcIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserExchangeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserExchangeServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserExchangeSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserExchangeUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserExchange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("addr_type", flattenObjectUserExchangeAddrType(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "ObjectUserExchange-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("auth_level", flattenObjectUserExchangeAuthLevel(o["auth-level"], d, "auth_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-level"], "ObjectUserExchange-AuthLevel"); ok {
			if err = d.Set("auth_level", vv); err != nil {
				return fmt.Errorf("Error reading auth_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_level: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenObjectUserExchangeAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "ObjectUserExchange-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("auto_discover_kdc", flattenObjectUserExchangeAutoDiscoverKdc(o["auto-discover-kdc"], d, "auto_discover_kdc")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-discover-kdc"], "ObjectUserExchange-AutoDiscoverKdc"); ok {
			if err = d.Set("auto_discover_kdc", vv); err != nil {
				return fmt.Errorf("Error reading auto_discover_kdc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_discover_kdc: %v", err)
		}
	}

	if err = d.Set("connect_protocol", flattenObjectUserExchangeConnectProtocol(o["connect-protocol"], d, "connect_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["connect-protocol"], "ObjectUserExchange-ConnectProtocol"); ok {
			if err = d.Set("connect_protocol", vv); err != nil {
				return fmt.Errorf("Error reading connect_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading connect_protocol: %v", err)
		}
	}

	if err = d.Set("domain_name", flattenObjectUserExchangeDomainName(o["domain-name"], d, "domain_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-name"], "ObjectUserExchange-DomainName"); ok {
			if err = d.Set("domain_name", vv); err != nil {
				return fmt.Errorf("Error reading domain_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if err = d.Set("http_auth_type", flattenObjectUserExchangeHttpAuthType(o["http-auth-type"], d, "http_auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-auth-type"], "ObjectUserExchange-HttpAuthType"); ok {
			if err = d.Set("http_auth_type", vv); err != nil {
				return fmt.Errorf("Error reading http_auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_auth_type: %v", err)
		}
	}

	if err = d.Set("ip", flattenObjectUserExchangeIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectUserExchange-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ip6", flattenObjectUserExchangeIp6(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "ObjectUserExchange-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("kdc_ip", flattenObjectUserExchangeKdcIp(o["kdc-ip"], d, "kdc_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["kdc-ip"], "ObjectUserExchange-KdcIp"); ok {
			if err = d.Set("kdc_ip", vv); err != nil {
				return fmt.Errorf("Error reading kdc_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading kdc_ip: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserExchangeName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserExchange-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server_name", flattenObjectUserExchangeServerName(o["server-name"], d, "server_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-name"], "ObjectUserExchange-ServerName"); ok {
			if err = d.Set("server_name", vv); err != nil {
				return fmt.Errorf("Error reading server_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenObjectUserExchangeSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-proto-version"], "ObjectUserExchange-SslMinProtoVersion"); ok {
			if err = d.Set("ssl_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("username", flattenObjectUserExchangeUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "ObjectUserExchange-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	return nil
}

func flattenObjectUserExchangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserExchangeAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserExchangeAuthLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserExchangeAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserExchangeAutoDiscoverKdc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserExchangeConnectProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserExchangeDomainName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserExchangeHttpAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserExchangeIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserExchangeIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserExchangeKdcIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserExchangeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserExchangePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserExchangeServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserExchangeSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserExchangeUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserExchange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandObjectUserExchangeAddrType(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("auth_level"); ok || d.HasChange("auth_level") {
		t, err := expandObjectUserExchangeAuthLevel(d, v, "auth_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-level"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok || d.HasChange("auth_type") {
		t, err := expandObjectUserExchangeAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("auto_discover_kdc"); ok || d.HasChange("auto_discover_kdc") {
		t, err := expandObjectUserExchangeAutoDiscoverKdc(d, v, "auto_discover_kdc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discover-kdc"] = t
		}
	}

	if v, ok := d.GetOk("connect_protocol"); ok || d.HasChange("connect_protocol") {
		t, err := expandObjectUserExchangeConnectProtocol(d, v, "connect_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connect-protocol"] = t
		}
	}

	if v, ok := d.GetOk("domain_name"); ok || d.HasChange("domain_name") {
		t, err := expandObjectUserExchangeDomainName(d, v, "domain_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name"] = t
		}
	}

	if v, ok := d.GetOk("http_auth_type"); ok || d.HasChange("http_auth_type") {
		t, err := expandObjectUserExchangeHttpAuthType(d, v, "http_auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-auth-type"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandObjectUserExchangeIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok || d.HasChange("ip6") {
		t, err := expandObjectUserExchangeIp6(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("kdc_ip"); ok || d.HasChange("kdc_ip") {
		t, err := expandObjectUserExchangeKdcIp(d, v, "kdc_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["kdc-ip"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserExchangeName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectUserExchangePassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("server_name"); ok || d.HasChange("server_name") {
		t, err := expandObjectUserExchangeServerName(d, v, "server_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-name"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok || d.HasChange("ssl_min_proto_version") {
		t, err := expandObjectUserExchangeSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandObjectUserExchangeUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	return &obj, nil
}
