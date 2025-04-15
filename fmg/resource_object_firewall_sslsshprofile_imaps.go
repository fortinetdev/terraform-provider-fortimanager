// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IMAPS options.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallSslSshProfileImaps() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallSslSshProfileImapsUpdate,
		Read:   resourceObjectFirewallSslSshProfileImapsRead,
		Update: resourceObjectFirewallSslSshProfileImapsUpdate,
		Delete: resourceObjectFirewallSslSshProfileImapsDelete,

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
			"allow_invalid_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cert_validation_failure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_validation_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expired_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_allowed_ssl_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_cert_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"invalid_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"proxy_after_tcp_handshake": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"revoked_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sni_server_cert_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unsupported_ssl_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unsupported_ssl_negotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unsupported_ssl_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unsupported_ssl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"untrusted_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"untrusted_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallSslSshProfileImapsUpdate(d *schema.ResourceData, m interface{}) error {
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

	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	obj, err := getObjectObjectFirewallSslSshProfileImaps(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallSslSshProfileImaps resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallSslSshProfileImaps(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallSslSshProfileImaps resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFirewallSslSshProfileImaps")

	return resourceObjectFirewallSslSshProfileImapsRead(d, m)
}

func resourceObjectFirewallSslSshProfileImapsDelete(d *schema.ResourceData, m interface{}) error {
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

	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallSslSshProfileImaps(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallSslSshProfileImaps resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallSslSshProfileImapsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallSslSshProfileImaps(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallSslSshProfileImaps resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallSslSshProfileImaps(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallSslSshProfileImaps resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallSslSshProfileImapsAllowInvalidServerCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsCertValidationFailure2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsCertValidationTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsExpiredServerCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsMinAllowedSslVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsClientCertRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsInvalidServerCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsPorts2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallSslSshProfileImapsProxyAfterTcpHandshake2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsRevokedServerCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsSniServerCertCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsUnsupportedSslCipher2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsUnsupportedSslNegotiation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsUnsupportedSslVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsUnsupportedSsl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsUntrustedCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsUntrustedServerCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallSslSshProfileImaps(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("allow_invalid_server_cert", flattenObjectFirewallSslSshProfileImapsAllowInvalidServerCert2edl(o["allow-invalid-server-cert"], d, "allow_invalid_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-invalid-server-cert"], "ObjectFirewallSslSshProfileImaps-AllowInvalidServerCert"); ok {
			if err = d.Set("allow_invalid_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading allow_invalid_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_invalid_server_cert: %v", err)
		}
	}

	if err = d.Set("cert_validation_failure", flattenObjectFirewallSslSshProfileImapsCertValidationFailure2edl(o["cert-validation-failure"], d, "cert_validation_failure")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-validation-failure"], "ObjectFirewallSslSshProfileImaps-CertValidationFailure"); ok {
			if err = d.Set("cert_validation_failure", vv); err != nil {
				return fmt.Errorf("Error reading cert_validation_failure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_validation_failure: %v", err)
		}
	}

	if err = d.Set("cert_validation_timeout", flattenObjectFirewallSslSshProfileImapsCertValidationTimeout2edl(o["cert-validation-timeout"], d, "cert_validation_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-validation-timeout"], "ObjectFirewallSslSshProfileImaps-CertValidationTimeout"); ok {
			if err = d.Set("cert_validation_timeout", vv); err != nil {
				return fmt.Errorf("Error reading cert_validation_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_validation_timeout: %v", err)
		}
	}

	if err = d.Set("client_certificate", flattenObjectFirewallSslSshProfileImapsClientCertificate2edl(o["client-certificate"], d, "client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-certificate"], "ObjectFirewallSslSshProfileImaps-ClientCertificate"); ok {
			if err = d.Set("client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_certificate: %v", err)
		}
	}

	if err = d.Set("expired_server_cert", flattenObjectFirewallSslSshProfileImapsExpiredServerCert2edl(o["expired-server-cert"], d, "expired_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["expired-server-cert"], "ObjectFirewallSslSshProfileImaps-ExpiredServerCert"); ok {
			if err = d.Set("expired_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading expired_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expired_server_cert: %v", err)
		}
	}

	if err = d.Set("min_allowed_ssl_version", flattenObjectFirewallSslSshProfileImapsMinAllowedSslVersion2edl(o["min-allowed-ssl-version"], d, "min_allowed_ssl_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-allowed-ssl-version"], "ObjectFirewallSslSshProfileImaps-MinAllowedSslVersion"); ok {
			if err = d.Set("min_allowed_ssl_version", vv); err != nil {
				return fmt.Errorf("Error reading min_allowed_ssl_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_allowed_ssl_version: %v", err)
		}
	}

	if err = d.Set("client_cert_request", flattenObjectFirewallSslSshProfileImapsClientCertRequest2edl(o["client-cert-request"], d, "client_cert_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert-request"], "ObjectFirewallSslSshProfileImaps-ClientCertRequest"); ok {
			if err = d.Set("client_cert_request", vv); err != nil {
				return fmt.Errorf("Error reading client_cert_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert_request: %v", err)
		}
	}

	if err = d.Set("invalid_server_cert", flattenObjectFirewallSslSshProfileImapsInvalidServerCert2edl(o["invalid-server-cert"], d, "invalid_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["invalid-server-cert"], "ObjectFirewallSslSshProfileImaps-InvalidServerCert"); ok {
			if err = d.Set("invalid_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading invalid_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading invalid_server_cert: %v", err)
		}
	}

	if err = d.Set("ports", flattenObjectFirewallSslSshProfileImapsPorts2edl(o["ports"], d, "ports")); err != nil {
		if vv, ok := fortiAPIPatch(o["ports"], "ObjectFirewallSslSshProfileImaps-Ports"); ok {
			if err = d.Set("ports", vv); err != nil {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ports: %v", err)
		}
	}

	if err = d.Set("proxy_after_tcp_handshake", flattenObjectFirewallSslSshProfileImapsProxyAfterTcpHandshake2edl(o["proxy-after-tcp-handshake"], d, "proxy_after_tcp_handshake")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-after-tcp-handshake"], "ObjectFirewallSslSshProfileImaps-ProxyAfterTcpHandshake"); ok {
			if err = d.Set("proxy_after_tcp_handshake", vv); err != nil {
				return fmt.Errorf("Error reading proxy_after_tcp_handshake: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_after_tcp_handshake: %v", err)
		}
	}

	if err = d.Set("revoked_server_cert", flattenObjectFirewallSslSshProfileImapsRevokedServerCert2edl(o["revoked-server-cert"], d, "revoked_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["revoked-server-cert"], "ObjectFirewallSslSshProfileImaps-RevokedServerCert"); ok {
			if err = d.Set("revoked_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading revoked_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading revoked_server_cert: %v", err)
		}
	}

	if err = d.Set("sni_server_cert_check", flattenObjectFirewallSslSshProfileImapsSniServerCertCheck2edl(o["sni-server-cert-check"], d, "sni_server_cert_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["sni-server-cert-check"], "ObjectFirewallSslSshProfileImaps-SniServerCertCheck"); ok {
			if err = d.Set("sni_server_cert_check", vv); err != nil {
				return fmt.Errorf("Error reading sni_server_cert_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sni_server_cert_check: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectFirewallSslSshProfileImapsStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectFirewallSslSshProfileImaps-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("unsupported_ssl_cipher", flattenObjectFirewallSslSshProfileImapsUnsupportedSslCipher2edl(o["unsupported-ssl-cipher"], d, "unsupported_ssl_cipher")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsupported-ssl-cipher"], "ObjectFirewallSslSshProfileImaps-UnsupportedSslCipher"); ok {
			if err = d.Set("unsupported_ssl_cipher", vv); err != nil {
				return fmt.Errorf("Error reading unsupported_ssl_cipher: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsupported_ssl_cipher: %v", err)
		}
	}

	if err = d.Set("unsupported_ssl_negotiation", flattenObjectFirewallSslSshProfileImapsUnsupportedSslNegotiation2edl(o["unsupported-ssl-negotiation"], d, "unsupported_ssl_negotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsupported-ssl-negotiation"], "ObjectFirewallSslSshProfileImaps-UnsupportedSslNegotiation"); ok {
			if err = d.Set("unsupported_ssl_negotiation", vv); err != nil {
				return fmt.Errorf("Error reading unsupported_ssl_negotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsupported_ssl_negotiation: %v", err)
		}
	}

	if err = d.Set("unsupported_ssl_version", flattenObjectFirewallSslSshProfileImapsUnsupportedSslVersion2edl(o["unsupported-ssl-version"], d, "unsupported_ssl_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsupported-ssl-version"], "ObjectFirewallSslSshProfileImaps-UnsupportedSslVersion"); ok {
			if err = d.Set("unsupported_ssl_version", vv); err != nil {
				return fmt.Errorf("Error reading unsupported_ssl_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsupported_ssl_version: %v", err)
		}
	}

	if err = d.Set("unsupported_ssl", flattenObjectFirewallSslSshProfileImapsUnsupportedSsl2edl(o["unsupported-ssl"], d, "unsupported_ssl")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsupported-ssl"], "ObjectFirewallSslSshProfileImaps-UnsupportedSsl"); ok {
			if err = d.Set("unsupported_ssl", vv); err != nil {
				return fmt.Errorf("Error reading unsupported_ssl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsupported_ssl: %v", err)
		}
	}

	if err = d.Set("untrusted_cert", flattenObjectFirewallSslSshProfileImapsUntrustedCert2edl(o["untrusted-cert"], d, "untrusted_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["untrusted-cert"], "ObjectFirewallSslSshProfileImaps-UntrustedCert"); ok {
			if err = d.Set("untrusted_cert", vv); err != nil {
				return fmt.Errorf("Error reading untrusted_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading untrusted_cert: %v", err)
		}
	}

	if err = d.Set("untrusted_server_cert", flattenObjectFirewallSslSshProfileImapsUntrustedServerCert2edl(o["untrusted-server-cert"], d, "untrusted_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["untrusted-server-cert"], "ObjectFirewallSslSshProfileImaps-UntrustedServerCert"); ok {
			if err = d.Set("untrusted_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading untrusted_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading untrusted_server_cert: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallSslSshProfileImapsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallSslSshProfileImapsAllowInvalidServerCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsCertValidationFailure2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsCertValidationTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsExpiredServerCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsMinAllowedSslVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsClientCertRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsInvalidServerCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsPorts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallSslSshProfileImapsProxyAfterTcpHandshake2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsRevokedServerCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsSniServerCertCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsUnsupportedSslCipher2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsUnsupportedSslNegotiation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsUnsupportedSslVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsUnsupportedSsl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsUntrustedCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsUntrustedServerCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallSslSshProfileImaps(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allow_invalid_server_cert"); ok || d.HasChange("allow_invalid_server_cert") {
		t, err := expandObjectFirewallSslSshProfileImapsAllowInvalidServerCert2edl(d, v, "allow_invalid_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-invalid-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("cert_validation_failure"); ok || d.HasChange("cert_validation_failure") {
		t, err := expandObjectFirewallSslSshProfileImapsCertValidationFailure2edl(d, v, "cert_validation_failure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-validation-failure"] = t
		}
	}

	if v, ok := d.GetOk("cert_validation_timeout"); ok || d.HasChange("cert_validation_timeout") {
		t, err := expandObjectFirewallSslSshProfileImapsCertValidationTimeout2edl(d, v, "cert_validation_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-validation-timeout"] = t
		}
	}

	if v, ok := d.GetOk("client_certificate"); ok || d.HasChange("client_certificate") {
		t, err := expandObjectFirewallSslSshProfileImapsClientCertificate2edl(d, v, "client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("expired_server_cert"); ok || d.HasChange("expired_server_cert") {
		t, err := expandObjectFirewallSslSshProfileImapsExpiredServerCert2edl(d, v, "expired_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expired-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("min_allowed_ssl_version"); ok || d.HasChange("min_allowed_ssl_version") {
		t, err := expandObjectFirewallSslSshProfileImapsMinAllowedSslVersion2edl(d, v, "min_allowed_ssl_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-allowed-ssl-version"] = t
		}
	}

	if v, ok := d.GetOk("client_cert_request"); ok || d.HasChange("client_cert_request") {
		t, err := expandObjectFirewallSslSshProfileImapsClientCertRequest2edl(d, v, "client_cert_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert-request"] = t
		}
	}

	if v, ok := d.GetOk("invalid_server_cert"); ok || d.HasChange("invalid_server_cert") {
		t, err := expandObjectFirewallSslSshProfileImapsInvalidServerCert2edl(d, v, "invalid_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["invalid-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandObjectFirewallSslSshProfileImapsPorts2edl(d, v, "ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("proxy_after_tcp_handshake"); ok || d.HasChange("proxy_after_tcp_handshake") {
		t, err := expandObjectFirewallSslSshProfileImapsProxyAfterTcpHandshake2edl(d, v, "proxy_after_tcp_handshake")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-after-tcp-handshake"] = t
		}
	}

	if v, ok := d.GetOk("revoked_server_cert"); ok || d.HasChange("revoked_server_cert") {
		t, err := expandObjectFirewallSslSshProfileImapsRevokedServerCert2edl(d, v, "revoked_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["revoked-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("sni_server_cert_check"); ok || d.HasChange("sni_server_cert_check") {
		t, err := expandObjectFirewallSslSshProfileImapsSniServerCertCheck2edl(d, v, "sni_server_cert_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sni-server-cert-check"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectFirewallSslSshProfileImapsStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("unsupported_ssl_cipher"); ok || d.HasChange("unsupported_ssl_cipher") {
		t, err := expandObjectFirewallSslSshProfileImapsUnsupportedSslCipher2edl(d, v, "unsupported_ssl_cipher")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsupported-ssl-cipher"] = t
		}
	}

	if v, ok := d.GetOk("unsupported_ssl_negotiation"); ok || d.HasChange("unsupported_ssl_negotiation") {
		t, err := expandObjectFirewallSslSshProfileImapsUnsupportedSslNegotiation2edl(d, v, "unsupported_ssl_negotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsupported-ssl-negotiation"] = t
		}
	}

	if v, ok := d.GetOk("unsupported_ssl_version"); ok || d.HasChange("unsupported_ssl_version") {
		t, err := expandObjectFirewallSslSshProfileImapsUnsupportedSslVersion2edl(d, v, "unsupported_ssl_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsupported-ssl-version"] = t
		}
	}

	if v, ok := d.GetOk("unsupported_ssl"); ok || d.HasChange("unsupported_ssl") {
		t, err := expandObjectFirewallSslSshProfileImapsUnsupportedSsl2edl(d, v, "unsupported_ssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsupported-ssl"] = t
		}
	}

	if v, ok := d.GetOk("untrusted_cert"); ok || d.HasChange("untrusted_cert") {
		t, err := expandObjectFirewallSslSshProfileImapsUntrustedCert2edl(d, v, "untrusted_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["untrusted-cert"] = t
		}
	}

	if v, ok := d.GetOk("untrusted_server_cert"); ok || d.HasChange("untrusted_server_cert") {
		t, err := expandObjectFirewallSslSshProfileImapsUntrustedServerCert2edl(d, v, "untrusted_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["untrusted-server-cert"] = t
		}
	}

	return &obj, nil
}
