// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure ZTNA destination.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectZtnaDestination() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectZtnaDestinationCreate,
		Read:   resourceObjectZtnaDestinationRead,
		Update: resourceObjectZtnaDestinationUpdate,
		Delete: resourceObjectZtnaDestinationDelete,

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
			"address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"conn_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mappedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"saas_application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_client_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_host_key": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_host_key_validation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectZtnaDestinationCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectZtnaDestination(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectZtnaDestination resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectZtnaDestination(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectZtnaDestination(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectZtnaDestination resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectZtnaDestination(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectZtnaDestination resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectZtnaDestinationRead(d, m)
}

func resourceObjectZtnaDestinationUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectZtnaDestination(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectZtnaDestination resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectZtnaDestination(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectZtnaDestination resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectZtnaDestinationRead(d, m)
}

func resourceObjectZtnaDestinationDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectZtnaDestination(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectZtnaDestination resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectZtnaDestinationRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectZtnaDestination(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectZtnaDestination resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectZtnaDestination(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectZtnaDestination resource from API: %v", err)
	}
	return nil
}

func flattenObjectZtnaDestinationAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectZtnaDestinationConnType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaDestinationDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaDestinationExternalAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaDestinationMappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaDestinationName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaDestinationProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaDestinationSaasApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectZtnaDestinationSshClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectZtnaDestinationSshHostKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectZtnaDestinationSshHostKeyValidation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaDestinationTunnelEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaDestinationType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaDestinationUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectZtnaDestination(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("address", flattenObjectZtnaDestinationAddress(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "ObjectZtnaDestination-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("conn_type", flattenObjectZtnaDestinationConnType(o["conn-type"], d, "conn_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["conn-type"], "ObjectZtnaDestination-ConnType"); ok {
			if err = d.Set("conn_type", vv); err != nil {
				return fmt.Errorf("Error reading conn_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conn_type: %v", err)
		}
	}

	if err = d.Set("domain", flattenObjectZtnaDestinationDomain(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "ObjectZtnaDestination-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("external_auth", flattenObjectZtnaDestinationExternalAuth(o["external-auth"], d, "external_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-auth"], "ObjectZtnaDestination-ExternalAuth"); ok {
			if err = d.Set("external_auth", vv); err != nil {
				return fmt.Errorf("Error reading external_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_auth: %v", err)
		}
	}

	if err = d.Set("mappedport", flattenObjectZtnaDestinationMappedport(o["mappedport"], d, "mappedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["mappedport"], "ObjectZtnaDestination-Mappedport"); ok {
			if err = d.Set("mappedport", vv); err != nil {
				return fmt.Errorf("Error reading mappedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mappedport: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectZtnaDestinationName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectZtnaDestination-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("protocol", flattenObjectZtnaDestinationProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ObjectZtnaDestination-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("saas_application", flattenObjectZtnaDestinationSaasApplication(o["saas-application"], d, "saas_application")); err != nil {
		if vv, ok := fortiAPIPatch(o["saas-application"], "ObjectZtnaDestination-SaasApplication"); ok {
			if err = d.Set("saas_application", vv); err != nil {
				return fmt.Errorf("Error reading saas_application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saas_application: %v", err)
		}
	}

	if err = d.Set("ssh_client_cert", flattenObjectZtnaDestinationSshClientCert(o["ssh-client-cert"], d, "ssh_client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-client-cert"], "ObjectZtnaDestination-SshClientCert"); ok {
			if err = d.Set("ssh_client_cert", vv); err != nil {
				return fmt.Errorf("Error reading ssh_client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_client_cert: %v", err)
		}
	}

	if err = d.Set("ssh_host_key", flattenObjectZtnaDestinationSshHostKey(o["ssh-host-key"], d, "ssh_host_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-host-key"], "ObjectZtnaDestination-SshHostKey"); ok {
			if err = d.Set("ssh_host_key", vv); err != nil {
				return fmt.Errorf("Error reading ssh_host_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_host_key: %v", err)
		}
	}

	if err = d.Set("ssh_host_key_validation", flattenObjectZtnaDestinationSshHostKeyValidation(o["ssh-host-key-validation"], d, "ssh_host_key_validation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-host-key-validation"], "ObjectZtnaDestination-SshHostKeyValidation"); ok {
			if err = d.Set("ssh_host_key_validation", vv); err != nil {
				return fmt.Errorf("Error reading ssh_host_key_validation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_host_key_validation: %v", err)
		}
	}

	if err = d.Set("tunnel_encryption", flattenObjectZtnaDestinationTunnelEncryption(o["tunnel-encryption"], d, "tunnel_encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-encryption"], "ObjectZtnaDestination-TunnelEncryption"); ok {
			if err = d.Set("tunnel_encryption", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_encryption: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectZtnaDestinationType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectZtnaDestination-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectZtnaDestinationUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectZtnaDestination-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenObjectZtnaDestinationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectZtnaDestinationAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaDestinationConnType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaDestinationDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaDestinationExternalAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaDestinationMappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaDestinationName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaDestinationProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaDestinationSaasApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaDestinationSshClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaDestinationSshHostKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaDestinationSshHostKeyValidation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaDestinationTunnelEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaDestinationType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaDestinationUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectZtnaDestination(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandObjectZtnaDestinationAddress(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("conn_type"); ok || d.HasChange("conn_type") {
		t, err := expandObjectZtnaDestinationConnType(d, v, "conn_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-type"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandObjectZtnaDestinationDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("external_auth"); ok || d.HasChange("external_auth") {
		t, err := expandObjectZtnaDestinationExternalAuth(d, v, "external_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-auth"] = t
		}
	}

	if v, ok := d.GetOk("mappedport"); ok || d.HasChange("mappedport") {
		t, err := expandObjectZtnaDestinationMappedport(d, v, "mappedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedport"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectZtnaDestinationName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandObjectZtnaDestinationProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("saas_application"); ok || d.HasChange("saas_application") {
		t, err := expandObjectZtnaDestinationSaasApplication(d, v, "saas_application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saas-application"] = t
		}
	}

	if v, ok := d.GetOk("ssh_client_cert"); ok || d.HasChange("ssh_client_cert") {
		t, err := expandObjectZtnaDestinationSshClientCert(d, v, "ssh_client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-client-cert"] = t
		}
	}

	if v, ok := d.GetOk("ssh_host_key"); ok || d.HasChange("ssh_host_key") {
		t, err := expandObjectZtnaDestinationSshHostKey(d, v, "ssh_host_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-host-key"] = t
		}
	}

	if v, ok := d.GetOk("ssh_host_key_validation"); ok || d.HasChange("ssh_host_key_validation") {
		t, err := expandObjectZtnaDestinationSshHostKeyValidation(d, v, "ssh_host_key_validation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-host-key-validation"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_encryption"); ok || d.HasChange("tunnel_encryption") {
		t, err := expandObjectZtnaDestinationTunnelEncryption(d, v, "tunnel_encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-encryption"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectZtnaDestinationType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandObjectZtnaDestinationUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
