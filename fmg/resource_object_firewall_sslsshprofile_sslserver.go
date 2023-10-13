// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SSL servers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallSslSshProfileSslServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallSslSshProfileSslServerCreate,
		Read:   resourceObjectFirewallSslSshProfileSslServerRead,
		Update: resourceObjectFirewallSslSshProfileSslServerUpdate,
		Delete: resourceObjectFirewallSslSshProfileSslServerDelete,

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
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ftps_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftps_client_cert_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"https_client_cert_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"imaps_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"imaps_client_cert_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pop3s_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smtps_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_other_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pop3s_client_cert_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"smtps_client_cert_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_other_client_cert_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallSslSshProfileSslServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	obj, err := getObjectObjectFirewallSslSshProfileSslServer(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallSslSshProfileSslServer resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallSslSshProfileSslServer(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallSslSshProfileSslServer resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallSslSshProfileSslServerRead(d, m)
}

func resourceObjectFirewallSslSshProfileSslServerUpdate(d *schema.ResourceData, m interface{}) error {
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

	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	obj, err := getObjectObjectFirewallSslSshProfileSslServer(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallSslSshProfileSslServer resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallSslSshProfileSslServer(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallSslSshProfileSslServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallSslSshProfileSslServerRead(d, m)
}

func resourceObjectFirewallSslSshProfileSslServerDelete(d *schema.ResourceData, m interface{}) error {
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

	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	err = c.DeleteObjectFirewallSslSshProfileSslServer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallSslSshProfileSslServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallSslSshProfileSslServerRead(d *schema.ResourceData, m interface{}) error {
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

	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	if ssl_ssh_profile == "" {
		ssl_ssh_profile = importOptionChecking(m.(*FortiClient).Cfg, "ssl_ssh_profile")
		if ssl_ssh_profile == "" {
			return fmt.Errorf("Parameter ssl_ssh_profile is missing")
		}
		if err = d.Set("ssl_ssh_profile", ssl_ssh_profile); err != nil {
			return fmt.Errorf("Error set params ssl_ssh_profile: %v", err)
		}
	}
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	o, err := c.ReadObjectFirewallSslSshProfileSslServer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallSslSshProfileSslServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallSslSshProfileSslServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallSslSshProfileSslServer resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallSslSshProfileSslServerFtpsClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerHttpsClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerFtpsClientCertRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerHttpsClientCertRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerImapsClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerImapsClientCertRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerPop3SClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerSmtpsClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerSslOtherClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerPop3SClientCertRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerSmtpsClientCertRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerSslOtherClientCertRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallSslSshProfileSslServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("ftps_client_certificate", flattenObjectFirewallSslSshProfileSslServerFtpsClientCertificate2edl(o["ftps-client-certificate"], d, "ftps_client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftps-client-certificate"], "ObjectFirewallSslSshProfileSslServer-FtpsClientCertificate"); ok {
			if err = d.Set("ftps_client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ftps_client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftps_client_certificate: %v", err)
		}
	}

	if err = d.Set("https_client_certificate", flattenObjectFirewallSslSshProfileSslServerHttpsClientCertificate2edl(o["https-client-certificate"], d, "https_client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-client-certificate"], "ObjectFirewallSslSshProfileSslServer-HttpsClientCertificate"); ok {
			if err = d.Set("https_client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading https_client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_client_certificate: %v", err)
		}
	}

	if err = d.Set("ftps_client_cert_request", flattenObjectFirewallSslSshProfileSslServerFtpsClientCertRequest2edl(o["ftps-client-cert-request"], d, "ftps_client_cert_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftps-client-cert-request"], "ObjectFirewallSslSshProfileSslServer-FtpsClientCertRequest"); ok {
			if err = d.Set("ftps_client_cert_request", vv); err != nil {
				return fmt.Errorf("Error reading ftps_client_cert_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftps_client_cert_request: %v", err)
		}
	}

	if err = d.Set("https_client_cert_request", flattenObjectFirewallSslSshProfileSslServerHttpsClientCertRequest2edl(o["https-client-cert-request"], d, "https_client_cert_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-client-cert-request"], "ObjectFirewallSslSshProfileSslServer-HttpsClientCertRequest"); ok {
			if err = d.Set("https_client_cert_request", vv); err != nil {
				return fmt.Errorf("Error reading https_client_cert_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_client_cert_request: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallSslSshProfileSslServerId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallSslSshProfileSslServer-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("imaps_client_certificate", flattenObjectFirewallSslSshProfileSslServerImapsClientCertificate2edl(o["imaps-client-certificate"], d, "imaps_client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["imaps-client-certificate"], "ObjectFirewallSslSshProfileSslServer-ImapsClientCertificate"); ok {
			if err = d.Set("imaps_client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading imaps_client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading imaps_client_certificate: %v", err)
		}
	}

	if err = d.Set("imaps_client_cert_request", flattenObjectFirewallSslSshProfileSslServerImapsClientCertRequest2edl(o["imaps-client-cert-request"], d, "imaps_client_cert_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["imaps-client-cert-request"], "ObjectFirewallSslSshProfileSslServer-ImapsClientCertRequest"); ok {
			if err = d.Set("imaps_client_cert_request", vv); err != nil {
				return fmt.Errorf("Error reading imaps_client_cert_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading imaps_client_cert_request: %v", err)
		}
	}

	if err = d.Set("ip", flattenObjectFirewallSslSshProfileSslServerIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectFirewallSslSshProfileSslServer-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("pop3s_client_certificate", flattenObjectFirewallSslSshProfileSslServerPop3SClientCertificate2edl(o["pop3s-client-certificate"], d, "pop3s_client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["pop3s-client-certificate"], "ObjectFirewallSslSshProfileSslServer-Pop3SClientCertificate"); ok {
			if err = d.Set("pop3s_client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading pop3s_client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pop3s_client_certificate: %v", err)
		}
	}

	if err = d.Set("smtps_client_certificate", flattenObjectFirewallSslSshProfileSslServerSmtpsClientCertificate2edl(o["smtps-client-certificate"], d, "smtps_client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["smtps-client-certificate"], "ObjectFirewallSslSshProfileSslServer-SmtpsClientCertificate"); ok {
			if err = d.Set("smtps_client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading smtps_client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smtps_client_certificate: %v", err)
		}
	}

	if err = d.Set("ssl_other_client_certificate", flattenObjectFirewallSslSshProfileSslServerSslOtherClientCertificate2edl(o["ssl-other-client-certificate"], d, "ssl_other_client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-other-client-certificate"], "ObjectFirewallSslSshProfileSslServer-SslOtherClientCertificate"); ok {
			if err = d.Set("ssl_other_client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_other_client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_other_client_certificate: %v", err)
		}
	}

	if err = d.Set("pop3s_client_cert_request", flattenObjectFirewallSslSshProfileSslServerPop3SClientCertRequest2edl(o["pop3s-client-cert-request"], d, "pop3s_client_cert_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["pop3s-client-cert-request"], "ObjectFirewallSslSshProfileSslServer-Pop3SClientCertRequest"); ok {
			if err = d.Set("pop3s_client_cert_request", vv); err != nil {
				return fmt.Errorf("Error reading pop3s_client_cert_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pop3s_client_cert_request: %v", err)
		}
	}

	if err = d.Set("smtps_client_cert_request", flattenObjectFirewallSslSshProfileSslServerSmtpsClientCertRequest2edl(o["smtps-client-cert-request"], d, "smtps_client_cert_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["smtps-client-cert-request"], "ObjectFirewallSslSshProfileSslServer-SmtpsClientCertRequest"); ok {
			if err = d.Set("smtps_client_cert_request", vv); err != nil {
				return fmt.Errorf("Error reading smtps_client_cert_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smtps_client_cert_request: %v", err)
		}
	}

	if err = d.Set("ssl_other_client_cert_request", flattenObjectFirewallSslSshProfileSslServerSslOtherClientCertRequest2edl(o["ssl-other-client-cert-request"], d, "ssl_other_client_cert_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-other-client-cert-request"], "ObjectFirewallSslSshProfileSslServer-SslOtherClientCertRequest"); ok {
			if err = d.Set("ssl_other_client_cert_request", vv); err != nil {
				return fmt.Errorf("Error reading ssl_other_client_cert_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_other_client_cert_request: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallSslSshProfileSslServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallSslSshProfileSslServerFtpsClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerHttpsClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerFtpsClientCertRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerHttpsClientCertRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerImapsClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerImapsClientCertRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerPop3SClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerSmtpsClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerSslOtherClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerPop3SClientCertRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerSmtpsClientCertRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerSslOtherClientCertRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallSslSshProfileSslServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ftps_client_certificate"); ok || d.HasChange("ftps_client_certificate") {
		t, err := expandObjectFirewallSslSshProfileSslServerFtpsClientCertificate2edl(d, v, "ftps_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftps-client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("https_client_certificate"); ok || d.HasChange("https_client_certificate") {
		t, err := expandObjectFirewallSslSshProfileSslServerHttpsClientCertificate2edl(d, v, "https_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ftps_client_cert_request"); ok || d.HasChange("ftps_client_cert_request") {
		t, err := expandObjectFirewallSslSshProfileSslServerFtpsClientCertRequest2edl(d, v, "ftps_client_cert_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftps-client-cert-request"] = t
		}
	}

	if v, ok := d.GetOk("https_client_cert_request"); ok || d.HasChange("https_client_cert_request") {
		t, err := expandObjectFirewallSslSshProfileSslServerHttpsClientCertRequest2edl(d, v, "https_client_cert_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-client-cert-request"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallSslSshProfileSslServerId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("imaps_client_certificate"); ok || d.HasChange("imaps_client_certificate") {
		t, err := expandObjectFirewallSslSshProfileSslServerImapsClientCertificate2edl(d, v, "imaps_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imaps-client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("imaps_client_cert_request"); ok || d.HasChange("imaps_client_cert_request") {
		t, err := expandObjectFirewallSslSshProfileSslServerImapsClientCertRequest2edl(d, v, "imaps_client_cert_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imaps-client-cert-request"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandObjectFirewallSslSshProfileSslServerIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("pop3s_client_certificate"); ok || d.HasChange("pop3s_client_certificate") {
		t, err := expandObjectFirewallSslSshProfileSslServerPop3SClientCertificate2edl(d, v, "pop3s_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3s-client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("smtps_client_certificate"); ok || d.HasChange("smtps_client_certificate") {
		t, err := expandObjectFirewallSslSshProfileSslServerSmtpsClientCertificate2edl(d, v, "smtps_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtps-client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ssl_other_client_certificate"); ok || d.HasChange("ssl_other_client_certificate") {
		t, err := expandObjectFirewallSslSshProfileSslServerSslOtherClientCertificate2edl(d, v, "ssl_other_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-other-client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("pop3s_client_cert_request"); ok || d.HasChange("pop3s_client_cert_request") {
		t, err := expandObjectFirewallSslSshProfileSslServerPop3SClientCertRequest2edl(d, v, "pop3s_client_cert_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3s-client-cert-request"] = t
		}
	}

	if v, ok := d.GetOk("smtps_client_cert_request"); ok || d.HasChange("smtps_client_cert_request") {
		t, err := expandObjectFirewallSslSshProfileSslServerSmtpsClientCertRequest2edl(d, v, "smtps_client_cert_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtps-client-cert-request"] = t
		}
	}

	if v, ok := d.GetOk("ssl_other_client_cert_request"); ok || d.HasChange("ssl_other_client_cert_request") {
		t, err := expandObjectFirewallSslSshProfileSslServerSslOtherClientCertRequest2edl(d, v, "ssl_other_client_cert_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-other-client-cert-request"] = t
		}
	}

	return &obj, nil
}
