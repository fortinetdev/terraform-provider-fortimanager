// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Syslog servers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSyslog() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSyslogCreate,
		Read:   resourceSystemSyslogRead,
		Update: resourceSystemSyslogUpdate,
		Delete: resourceSystemSyslogDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"peer_cert_cn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reliable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secure_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSyslogCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSyslog(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSyslog resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSyslog(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemSyslog resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSyslogRead(d, m)
}

func resourceSystemSyslogUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSyslog(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSyslog resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSyslog(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSyslog resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSyslogRead(d, m)
}

func resourceSystemSyslogDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemSyslog(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSyslog resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSyslogRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemSyslog(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSyslog resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSyslog(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSyslog resource from API: %v", err)
	}
	return nil
}

func flattenSystemSyslogIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSyslogLocalCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSyslogName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSyslogPeerCertCn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSyslogPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSyslogReliable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSyslogSecureConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSyslogSslProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSyslog(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip", flattenSystemSyslogIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SystemSyslog-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("local_cert", flattenSystemSyslogLocalCert(o["local-cert"], d, "local_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-cert"], "SystemSyslog-LocalCert"); ok {
			if err = d.Set("local_cert", vv); err != nil {
				return fmt.Errorf("Error reading local_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_cert: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSyslogName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSyslog-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("peer_cert_cn", flattenSystemSyslogPeerCertCn(o["peer-cert-cn"], d, "peer_cert_cn")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-cert-cn"], "SystemSyslog-PeerCertCn"); ok {
			if err = d.Set("peer_cert_cn", vv); err != nil {
				return fmt.Errorf("Error reading peer_cert_cn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_cert_cn: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemSyslogPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemSyslog-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("reliable", flattenSystemSyslogReliable(o["reliable"], d, "reliable")); err != nil {
		if vv, ok := fortiAPIPatch(o["reliable"], "SystemSyslog-Reliable"); ok {
			if err = d.Set("reliable", vv); err != nil {
				return fmt.Errorf("Error reading reliable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reliable: %v", err)
		}
	}

	if err = d.Set("secure_connection", flattenSystemSyslogSecureConnection(o["secure-connection"], d, "secure_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-connection"], "SystemSyslog-SecureConnection"); ok {
			if err = d.Set("secure_connection", vv); err != nil {
				return fmt.Errorf("Error reading secure_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_connection: %v", err)
		}
	}

	if err = d.Set("ssl_protocol", flattenSystemSyslogSslProtocol(o["ssl-protocol"], d, "ssl_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-protocol"], "SystemSyslog-SslProtocol"); ok {
			if err = d.Set("ssl_protocol", vv); err != nil {
				return fmt.Errorf("Error reading ssl_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_protocol: %v", err)
		}
	}

	return nil
}

func flattenSystemSyslogFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSyslogIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSyslogLocalCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSyslogName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSyslogPeerCertCn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSyslogPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSyslogReliable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSyslogSecureConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSyslogSslProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSyslog(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSystemSyslogIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("local_cert"); ok || d.HasChange("local_cert") {
		t, err := expandSystemSyslogLocalCert(d, v, "local_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-cert"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSyslogName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("peer_cert_cn"); ok || d.HasChange("peer_cert_cn") {
		t, err := expandSystemSyslogPeerCertCn(d, v, "peer_cert_cn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-cert-cn"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemSyslogPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("reliable"); ok || d.HasChange("reliable") {
		t, err := expandSystemSyslogReliable(d, v, "reliable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reliable"] = t
		}
	}

	if v, ok := d.GetOk("secure_connection"); ok || d.HasChange("secure_connection") {
		t, err := expandSystemSyslogSecureConnection(d, v, "secure_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-connection"] = t
		}
	}

	if v, ok := d.GetOk("ssl_protocol"); ok || d.HasChange("ssl_protocol") {
		t, err := expandSystemSyslogSslProtocol(d, v, "ssl_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-protocol"] = t
		}
	}

	return &obj, nil
}
