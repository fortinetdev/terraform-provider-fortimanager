// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectZtna ServiceConnector

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectZtnaServiceConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectZtnaServiceConnectorCreate,
		Read:   resourceObjectZtnaServiceConnectorRead,
		Update: resourceObjectZtnaServiceConnectorUpdate,
		Delete: resourceObjectZtnaServiceConnectorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"connection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forward_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forward_destination_cn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forward_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"health_check_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"relay_dev_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"relay_user_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusted_ca": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"url_map": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectZtnaServiceConnectorCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectZtnaServiceConnector(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectZtnaServiceConnector resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectZtnaServiceConnector(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectZtnaServiceConnector(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectZtnaServiceConnector resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectZtnaServiceConnector(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectZtnaServiceConnector resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectZtnaServiceConnectorRead(d, m)
}

func resourceObjectZtnaServiceConnectorUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectZtnaServiceConnector(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectZtnaServiceConnector resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectZtnaServiceConnector(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectZtnaServiceConnector resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectZtnaServiceConnectorRead(d, m)
}

func resourceObjectZtnaServiceConnectorDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectZtnaServiceConnector(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectZtnaServiceConnector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectZtnaServiceConnectorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectZtnaServiceConnector(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectZtnaServiceConnector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectZtnaServiceConnector(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectZtnaServiceConnector resource from API: %v", err)
	}
	return nil
}

func flattenObjectZtnaServiceConnectorCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectZtnaServiceConnectorConnectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaServiceConnectorEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaServiceConnectorForwardAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaServiceConnectorForwardDestinationCn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaServiceConnectorForwardPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaServiceConnectorHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaServiceConnectorLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaServiceConnectorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaServiceConnectorRelayDevInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaServiceConnectorRelayUserInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaServiceConnectorSslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaServiceConnectorSslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaServiceConnectorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaServiceConnectorTrustedCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectZtnaServiceConnectorUrlMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectZtnaServiceConnector(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("certificate", flattenObjectZtnaServiceConnectorCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "ObjectZtnaServiceConnector-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("connection_mode", flattenObjectZtnaServiceConnectorConnectionMode(o["connection-mode"], d, "connection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["connection-mode"], "ObjectZtnaServiceConnector-ConnectionMode"); ok {
			if err = d.Set("connection_mode", vv); err != nil {
				return fmt.Errorf("Error reading connection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading connection_mode: %v", err)
		}
	}

	if err = d.Set("encryption", flattenObjectZtnaServiceConnectorEncryption(o["encryption"], d, "encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["encryption"], "ObjectZtnaServiceConnector-Encryption"); ok {
			if err = d.Set("encryption", vv); err != nil {
				return fmt.Errorf("Error reading encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encryption: %v", err)
		}
	}

	if err = d.Set("forward_address", flattenObjectZtnaServiceConnectorForwardAddress(o["forward-address"], d, "forward_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-address"], "ObjectZtnaServiceConnector-ForwardAddress"); ok {
			if err = d.Set("forward_address", vv); err != nil {
				return fmt.Errorf("Error reading forward_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_address: %v", err)
		}
	}

	if err = d.Set("forward_destination_cn", flattenObjectZtnaServiceConnectorForwardDestinationCn(o["forward-destination-cn"], d, "forward_destination_cn")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-destination-cn"], "ObjectZtnaServiceConnector-ForwardDestinationCn"); ok {
			if err = d.Set("forward_destination_cn", vv); err != nil {
				return fmt.Errorf("Error reading forward_destination_cn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_destination_cn: %v", err)
		}
	}

	if err = d.Set("forward_port", flattenObjectZtnaServiceConnectorForwardPort(o["forward-port"], d, "forward_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-port"], "ObjectZtnaServiceConnector-ForwardPort"); ok {
			if err = d.Set("forward_port", vv); err != nil {
				return fmt.Errorf("Error reading forward_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_port: %v", err)
		}
	}

	if err = d.Set("health_check_interval", flattenObjectZtnaServiceConnectorHealthCheckInterval(o["health-check-interval"], d, "health_check_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check-interval"], "ObjectZtnaServiceConnector-HealthCheckInterval"); ok {
			if err = d.Set("health_check_interval", vv); err != nil {
				return fmt.Errorf("Error reading health_check_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check_interval: %v", err)
		}
	}

	if err = d.Set("log", flattenObjectZtnaServiceConnectorLog(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ObjectZtnaServiceConnector-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectZtnaServiceConnectorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectZtnaServiceConnector-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("relay_dev_info", flattenObjectZtnaServiceConnectorRelayDevInfo(o["relay-dev-info"], d, "relay_dev_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["relay-dev-info"], "ObjectZtnaServiceConnector-RelayDevInfo"); ok {
			if err = d.Set("relay_dev_info", vv); err != nil {
				return fmt.Errorf("Error reading relay_dev_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading relay_dev_info: %v", err)
		}
	}

	if err = d.Set("relay_user_info", flattenObjectZtnaServiceConnectorRelayUserInfo(o["relay-user-info"], d, "relay_user_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["relay-user-info"], "ObjectZtnaServiceConnector-RelayUserInfo"); ok {
			if err = d.Set("relay_user_info", vv); err != nil {
				return fmt.Errorf("Error reading relay_user_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading relay_user_info: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenObjectZtnaServiceConnectorSslMaxVersion(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "ObjectZtnaServiceConnector-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenObjectZtnaServiceConnectorSslMinVersion(o["ssl-min-version"], d, "ssl_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-version"], "ObjectZtnaServiceConnector-SslMinVersion"); ok {
			if err = d.Set("ssl_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectZtnaServiceConnectorStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectZtnaServiceConnector-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trusted_ca", flattenObjectZtnaServiceConnectorTrustedCa(o["trusted-ca"], d, "trusted_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusted-ca"], "ObjectZtnaServiceConnector-TrustedCa"); ok {
			if err = d.Set("trusted_ca", vv); err != nil {
				return fmt.Errorf("Error reading trusted_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusted_ca: %v", err)
		}
	}

	if err = d.Set("url_map", flattenObjectZtnaServiceConnectorUrlMap(o["url-map"], d, "url_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-map"], "ObjectZtnaServiceConnector-UrlMap"); ok {
			if err = d.Set("url_map", vv); err != nil {
				return fmt.Errorf("Error reading url_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_map: %v", err)
		}
	}

	return nil
}

func flattenObjectZtnaServiceConnectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectZtnaServiceConnectorCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaServiceConnectorConnectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaServiceConnectorEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaServiceConnectorForwardAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaServiceConnectorForwardDestinationCn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaServiceConnectorForwardPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaServiceConnectorHealthCheckInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaServiceConnectorLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaServiceConnectorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaServiceConnectorRelayDevInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaServiceConnectorRelayUserInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaServiceConnectorSslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaServiceConnectorSslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaServiceConnectorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaServiceConnectorTrustedCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaServiceConnectorUrlMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectZtnaServiceConnector(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandObjectZtnaServiceConnectorCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("connection_mode"); ok || d.HasChange("connection_mode") {
		t, err := expandObjectZtnaServiceConnectorConnectionMode(d, v, "connection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connection-mode"] = t
		}
	}

	if v, ok := d.GetOk("encryption"); ok || d.HasChange("encryption") {
		t, err := expandObjectZtnaServiceConnectorEncryption(d, v, "encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encryption"] = t
		}
	}

	if v, ok := d.GetOk("forward_address"); ok || d.HasChange("forward_address") {
		t, err := expandObjectZtnaServiceConnectorForwardAddress(d, v, "forward_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-address"] = t
		}
	}

	if v, ok := d.GetOk("forward_destination_cn"); ok || d.HasChange("forward_destination_cn") {
		t, err := expandObjectZtnaServiceConnectorForwardDestinationCn(d, v, "forward_destination_cn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-destination-cn"] = t
		}
	}

	if v, ok := d.GetOk("forward_port"); ok || d.HasChange("forward_port") {
		t, err := expandObjectZtnaServiceConnectorForwardPort(d, v, "forward_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-port"] = t
		}
	}

	if v, ok := d.GetOk("health_check_interval"); ok || d.HasChange("health_check_interval") {
		t, err := expandObjectZtnaServiceConnectorHealthCheckInterval(d, v, "health_check_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-interval"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandObjectZtnaServiceConnectorLog(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectZtnaServiceConnectorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("relay_dev_info"); ok || d.HasChange("relay_dev_info") {
		t, err := expandObjectZtnaServiceConnectorRelayDevInfo(d, v, "relay_dev_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["relay-dev-info"] = t
		}
	}

	if v, ok := d.GetOk("relay_user_info"); ok || d.HasChange("relay_user_info") {
		t, err := expandObjectZtnaServiceConnectorRelayUserInfo(d, v, "relay_user_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["relay-user-info"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandObjectZtnaServiceConnectorSslMaxVersion(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok || d.HasChange("ssl_min_version") {
		t, err := expandObjectZtnaServiceConnectorSslMinVersion(d, v, "ssl_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectZtnaServiceConnectorStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trusted_ca"); ok || d.HasChange("trusted_ca") {
		t, err := expandObjectZtnaServiceConnectorTrustedCa(d, v, "trusted_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted-ca"] = t
		}
	}

	if v, ok := d.GetOk("url_map"); ok || d.HasChange("url_map") {
		t, err := expandObjectZtnaServiceConnectorUrlMap(d, v, "url_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-map"] = t
		}
	}

	return &obj, nil
}
