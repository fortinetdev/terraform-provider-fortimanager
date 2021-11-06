// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure SSL/SSH protocol options.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallSslSshProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallSslSshProfileCreate,
		Read:   resourceObjectFirewallSslSshProfileRead,
		Update: resourceObjectFirewallSslSshProfileUpdate,
		Delete: resourceObjectFirewallSslSshProfileDelete,

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
			"allowlist": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_blocklisted_certificates": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_blacklisted_certificates": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"caname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dot": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ftps": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
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
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"https": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_probe_failure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"imaps": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mapi_over_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"pop3s": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"rpc_over_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_cert_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smtps": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssh": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"ssh_algorithm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssh_tun_policy_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"inspect_all": &schema.Schema{
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
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssl_anomalies_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_exempt": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"address6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiguard_category": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"regex": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wildcard_fqdn": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssl_exemptions_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_negotiation_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							Computed: true,
						},
						"https_client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"imaps_client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"imaps_client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"smtps_client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_other_client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"supported_alpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"untrusted_caname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_ssl_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"whitelist": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectFirewallSslSshProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallSslSshProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallSslSshProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallSslSshProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallSslSshProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallSslSshProfileRead(d, m)
}

func resourceObjectFirewallSslSshProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallSslSshProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallSslSshProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallSslSshProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallSslSshProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallSslSshProfileRead(d, m)
}

func resourceObjectFirewallSslSshProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallSslSshProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallSslSshProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallSslSshProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallSslSshProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallSslSshProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallSslSshProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallSslSshProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallSslSshProfileAllowlist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileBlockBlocklistedCertificates(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileBlockBlacklistedCertificates(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileCaname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileDot(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenObjectFirewallSslSshProfileDotCertValidationFailure(i["cert-validation-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenObjectFirewallSslSshProfileDotCertValidationTimeout(i["cert-validation-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenObjectFirewallSslSshProfileDotClientCertificate(i["client-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenObjectFirewallSslSshProfileDotExpiredServerCert(i["expired-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenObjectFirewallSslSshProfileDotProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenObjectFirewallSslSshProfileDotRevokedServerCert(i["revoked-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenObjectFirewallSslSshProfileDotSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallSslSshProfileDotStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenObjectFirewallSslSshProfileDotUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenObjectFirewallSslSshProfileDotUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenObjectFirewallSslSshProfileDotUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallSslSshProfileDotCertValidationFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileDotCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileDotClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileDotExpiredServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileDotProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileDotRevokedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileDotSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileDotStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileDotUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileDotUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileDotUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileFtps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenObjectFirewallSslSshProfileFtpsCertValidationFailure(i["cert-validation-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenObjectFirewallSslSshProfileFtpsCertValidationTimeout(i["cert-validation-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenObjectFirewallSslSshProfileFtpsClientCertificate(i["client-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenObjectFirewallSslSshProfileFtpsExpiredServerCert(i["expired-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenObjectFirewallSslSshProfileFtpsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenObjectFirewallSslSshProfileFtpsRevokedServerCert(i["revoked-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenObjectFirewallSslSshProfileFtpsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallSslSshProfileFtpsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenObjectFirewallSslSshProfileFtpsUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenObjectFirewallSslSshProfileFtpsUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenObjectFirewallSslSshProfileFtpsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallSslSshProfileFtpsCertValidationFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileFtpsCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileFtpsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileFtpsExpiredServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileFtpsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallSslSshProfileFtpsRevokedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileFtpsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileFtpsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileFtpsUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileFtpsUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileFtpsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileHttps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_probe_failure"
	if _, ok := i["cert-probe-failure"]; ok {
		result["cert_probe_failure"] = flattenObjectFirewallSslSshProfileHttpsCertProbeFailure(i["cert-probe-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenObjectFirewallSslSshProfileHttpsCertValidationFailure(i["cert-validation-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenObjectFirewallSslSshProfileHttpsCertValidationTimeout(i["cert-validation-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenObjectFirewallSslSshProfileHttpsClientCertificate(i["client-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenObjectFirewallSslSshProfileHttpsExpiredServerCert(i["expired-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenObjectFirewallSslSshProfileHttpsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenObjectFirewallSslSshProfileHttpsProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenObjectFirewallSslSshProfileHttpsRevokedServerCert(i["revoked-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenObjectFirewallSslSshProfileHttpsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallSslSshProfileHttpsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenObjectFirewallSslSshProfileHttpsUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenObjectFirewallSslSshProfileHttpsUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenObjectFirewallSslSshProfileHttpsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallSslSshProfileHttpsCertProbeFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileHttpsCertValidationFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileHttpsCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileHttpsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileHttpsExpiredServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileHttpsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallSslSshProfileHttpsProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileHttpsRevokedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileHttpsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileHttpsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileHttpsUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileHttpsUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileHttpsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImaps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenObjectFirewallSslSshProfileImapsCertValidationFailure(i["cert-validation-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenObjectFirewallSslSshProfileImapsCertValidationTimeout(i["cert-validation-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenObjectFirewallSslSshProfileImapsClientCertificate(i["client-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenObjectFirewallSslSshProfileImapsExpiredServerCert(i["expired-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenObjectFirewallSslSshProfileImapsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenObjectFirewallSslSshProfileImapsProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenObjectFirewallSslSshProfileImapsRevokedServerCert(i["revoked-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenObjectFirewallSslSshProfileImapsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallSslSshProfileImapsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenObjectFirewallSslSshProfileImapsUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenObjectFirewallSslSshProfileImapsUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenObjectFirewallSslSshProfileImapsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallSslSshProfileImapsCertValidationFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsExpiredServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallSslSshProfileImapsProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsRevokedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileImapsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileMapiOverHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfilePop3S(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenObjectFirewallSslSshProfilePop3SCertValidationFailure(i["cert-validation-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenObjectFirewallSslSshProfilePop3SCertValidationTimeout(i["cert-validation-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenObjectFirewallSslSshProfilePop3SClientCertificate(i["client-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenObjectFirewallSslSshProfilePop3SExpiredServerCert(i["expired-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenObjectFirewallSslSshProfilePop3SPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenObjectFirewallSslSshProfilePop3SProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenObjectFirewallSslSshProfilePop3SRevokedServerCert(i["revoked-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenObjectFirewallSslSshProfilePop3SSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallSslSshProfilePop3SStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenObjectFirewallSslSshProfilePop3SUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenObjectFirewallSslSshProfilePop3SUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenObjectFirewallSslSshProfilePop3SUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallSslSshProfilePop3SCertValidationFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfilePop3SCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfilePop3SClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfilePop3SExpiredServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfilePop3SPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallSslSshProfilePop3SProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfilePop3SRevokedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfilePop3SSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfilePop3SStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfilePop3SUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfilePop3SUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfilePop3SUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileRpcOverHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileServerCertMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSmtps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenObjectFirewallSslSshProfileSmtpsCertValidationFailure(i["cert-validation-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenObjectFirewallSslSshProfileSmtpsCertValidationTimeout(i["cert-validation-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenObjectFirewallSslSshProfileSmtpsClientCertificate(i["client-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenObjectFirewallSslSshProfileSmtpsExpiredServerCert(i["expired-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenObjectFirewallSslSshProfileSmtpsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenObjectFirewallSslSshProfileSmtpsProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenObjectFirewallSslSshProfileSmtpsRevokedServerCert(i["revoked-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenObjectFirewallSslSshProfileSmtpsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallSslSshProfileSmtpsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenObjectFirewallSslSshProfileSmtpsUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenObjectFirewallSslSshProfileSmtpsUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenObjectFirewallSslSshProfileSmtpsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallSslSshProfileSmtpsCertValidationFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSmtpsCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSmtpsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSmtpsExpiredServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSmtpsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallSslSshProfileSmtpsProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSmtpsRevokedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSmtpsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSmtpsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSmtpsUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSmtpsUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSmtpsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSsh(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenObjectFirewallSslSshProfileSshInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenObjectFirewallSslSshProfileSshPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenObjectFirewallSslSshProfileSshProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssh_algorithm"
	if _, ok := i["ssh-algorithm"]; ok {
		result["ssh_algorithm"] = flattenObjectFirewallSslSshProfileSshSshAlgorithm(i["ssh-algorithm"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssh_tun_policy_check"
	if _, ok := i["ssh-tun-policy-check"]; ok {
		result["ssh_tun_policy_check"] = flattenObjectFirewallSslSshProfileSshSshTunPolicyCheck(i["ssh-tun-policy-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallSslSshProfileSshStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_version"
	if _, ok := i["unsupported-version"]; ok {
		result["unsupported_version"] = flattenObjectFirewallSslSshProfileSshUnsupportedVersion(i["unsupported-version"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallSslSshProfileSshInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSshPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallSslSshProfileSshProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSshSshAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSshSshTunPolicyCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSshStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSshUnsupportedVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSsl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenObjectFirewallSslSshProfileSslCertValidationFailure(i["cert-validation-failure"], d, pre_append)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenObjectFirewallSslSshProfileSslCertValidationTimeout(i["cert-validation-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenObjectFirewallSslSshProfileSslClientCertificate(i["client-certificate"], d, pre_append)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenObjectFirewallSslSshProfileSslExpiredServerCert(i["expired-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenObjectFirewallSslSshProfileSslInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenObjectFirewallSslSshProfileSslRevokedServerCert(i["revoked-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenObjectFirewallSslSshProfileSslSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenObjectFirewallSslSshProfileSslUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenObjectFirewallSslSshProfileSslUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenObjectFirewallSslSshProfileSslUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallSslSshProfileSslCertValidationFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslExpiredServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslRevokedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslAnomaliesLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslExempt(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenObjectFirewallSslSshProfileSslExemptAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslExempt-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address6"
		if _, ok := i["address6"]; ok {
			v := flattenObjectFirewallSslSshProfileSslExemptAddress6(i["address6"], d, pre_append)
			tmp["address6"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslExempt-Address6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if _, ok := i["fortiguard-category"]; ok {
			v := flattenObjectFirewallSslSshProfileSslExemptFortiguardCategory(i["fortiguard-category"], d, pre_append)
			tmp["fortiguard_category"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslExempt-FortiguardCategory")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallSslSshProfileSslExemptId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslExempt-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			v := flattenObjectFirewallSslSshProfileSslExemptRegex(i["regex"], d, pre_append)
			tmp["regex"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslExempt-Regex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFirewallSslSshProfileSslExemptType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslExempt-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard_fqdn"
		if _, ok := i["wildcard-fqdn"]; ok {
			v := flattenObjectFirewallSslSshProfileSslExemptWildcardFqdn(i["wildcard-fqdn"], d, pre_append)
			tmp["wildcard_fqdn"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslExempt-WildcardFqdn")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallSslSshProfileSslExemptAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslExemptAddress6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslExemptFortiguardCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallSslSshProfileSslExemptId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslExemptRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslExemptType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslExemptWildcardFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallSslSshProfileSslExemptionsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslNegotiationLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftps_client_certificate"
		if _, ok := i["ftps-client-certificate"]; ok {
			v := flattenObjectFirewallSslSshProfileSslServerFtpsClientCertificate(i["ftps-client-certificate"], d, pre_append)
			tmp["ftps_client_certificate"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslServer-FtpsClientCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_client_certificate"
		if _, ok := i["https-client-certificate"]; ok {
			v := flattenObjectFirewallSslSshProfileSslServerHttpsClientCertificate(i["https-client-certificate"], d, pre_append)
			tmp["https_client_certificate"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslServer-HttpsClientCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftps_client_cert_request"
		if _, ok := i["ftps-client-cert-request"]; ok {
			v := flattenObjectFirewallSslSshProfileSslServerFtpsClientCertRequest(i["ftps-client-cert-request"], d, pre_append)
			tmp["ftps_client_cert_request"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslServer-FtpsClientCertRequest")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_client_cert_request"
		if _, ok := i["https-client-cert-request"]; ok {
			v := flattenObjectFirewallSslSshProfileSslServerHttpsClientCertRequest(i["https-client-cert-request"], d, pre_append)
			tmp["https_client_cert_request"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslServer-HttpsClientCertRequest")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallSslSshProfileSslServerId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslServer-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imaps_client_certificate"
		if _, ok := i["imaps-client-certificate"]; ok {
			v := flattenObjectFirewallSslSshProfileSslServerImapsClientCertificate(i["imaps-client-certificate"], d, pre_append)
			tmp["imaps_client_certificate"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslServer-ImapsClientCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imaps_client_cert_request"
		if _, ok := i["imaps-client-cert-request"]; ok {
			v := flattenObjectFirewallSslSshProfileSslServerImapsClientCertRequest(i["imaps-client-cert-request"], d, pre_append)
			tmp["imaps_client_cert_request"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslServer-ImapsClientCertRequest")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFirewallSslSshProfileSslServerIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslServer-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pop3s_client_certificate"
		if _, ok := i["pop3s-client-certificate"]; ok {
			v := flattenObjectFirewallSslSshProfileSslServerPop3SClientCertificate(i["pop3s-client-certificate"], d, pre_append)
			tmp["pop3s_client_certificate"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslServer-Pop3SClientCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smtps_client_certificate"
		if _, ok := i["smtps-client-certificate"]; ok {
			v := flattenObjectFirewallSslSshProfileSslServerSmtpsClientCertificate(i["smtps-client-certificate"], d, pre_append)
			tmp["smtps_client_certificate"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslServer-SmtpsClientCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_other_client_certificate"
		if _, ok := i["ssl-other-client-certificate"]; ok {
			v := flattenObjectFirewallSslSshProfileSslServerSslOtherClientCertificate(i["ssl-other-client-certificate"], d, pre_append)
			tmp["ssl_other_client_certificate"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslServer-SslOtherClientCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pop3s_client_cert_request"
		if _, ok := i["pop3s-client-cert-request"]; ok {
			v := flattenObjectFirewallSslSshProfileSslServerPop3SClientCertRequest(i["pop3s-client-cert-request"], d, pre_append)
			tmp["pop3s_client_cert_request"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslServer-Pop3SClientCertRequest")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smtps_client_cert_request"
		if _, ok := i["smtps-client-cert-request"]; ok {
			v := flattenObjectFirewallSslSshProfileSslServerSmtpsClientCertRequest(i["smtps-client-cert-request"], d, pre_append)
			tmp["smtps_client_cert_request"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslServer-SmtpsClientCertRequest")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_other_client_cert_request"
		if _, ok := i["ssl-other-client-cert-request"]; ok {
			v := flattenObjectFirewallSslSshProfileSslServerSslOtherClientCertRequest(i["ssl-other-client-cert-request"], d, pre_append)
			tmp["ssl_other_client_cert_request"] = fortiAPISubPartPatch(v, "ObjectFirewallSslSshProfile-SslServer-SslOtherClientCertRequest")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallSslSshProfileSslServerFtpsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerHttpsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerFtpsClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerHttpsClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerImapsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerImapsClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerPop3SClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerSmtpsClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerSslOtherClientCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerPop3SClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerSmtpsClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslServerSslOtherClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSupportedAlpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileUntrustedCaname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileUseSslServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileWhitelist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallSslSshProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("allowlist", flattenObjectFirewallSslSshProfileAllowlist(o["allowlist"], d, "allowlist")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowlist"], "ObjectFirewallSslSshProfile-Allowlist"); ok {
			if err = d.Set("allowlist", vv); err != nil {
				return fmt.Errorf("Error reading allowlist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowlist: %v", err)
		}
	}

	if err = d.Set("block_blocklisted_certificates", flattenObjectFirewallSslSshProfileBlockBlocklistedCertificates(o["block-blocklisted-certificates"], d, "block_blocklisted_certificates")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-blocklisted-certificates"], "ObjectFirewallSslSshProfile-BlockBlocklistedCertificates"); ok {
			if err = d.Set("block_blocklisted_certificates", vv); err != nil {
				return fmt.Errorf("Error reading block_blocklisted_certificates: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_blocklisted_certificates: %v", err)
		}
	}

	if err = d.Set("block_blacklisted_certificates", flattenObjectFirewallSslSshProfileBlockBlacklistedCertificates(o["block-blacklisted-certificates"], d, "block_blacklisted_certificates")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-blacklisted-certificates"], "ObjectFirewallSslSshProfile-BlockBlacklistedCertificates"); ok {
			if err = d.Set("block_blacklisted_certificates", vv); err != nil {
				return fmt.Errorf("Error reading block_blacklisted_certificates: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_blacklisted_certificates: %v", err)
		}
	}

	if err = d.Set("caname", flattenObjectFirewallSslSshProfileCaname(o["caname"], d, "caname")); err != nil {
		if vv, ok := fortiAPIPatch(o["caname"], "ObjectFirewallSslSshProfile-Caname"); ok {
			if err = d.Set("caname", vv); err != nil {
				return fmt.Errorf("Error reading caname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading caname: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectFirewallSslSshProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectFirewallSslSshProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dot", flattenObjectFirewallSslSshProfileDot(o["dot"], d, "dot")); err != nil {
			if vv, ok := fortiAPIPatch(o["dot"], "ObjectFirewallSslSshProfile-Dot"); ok {
				if err = d.Set("dot", vv); err != nil {
					return fmt.Errorf("Error reading dot: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dot: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dot"); ok {
			if err = d.Set("dot", flattenObjectFirewallSslSshProfileDot(o["dot"], d, "dot")); err != nil {
				if vv, ok := fortiAPIPatch(o["dot"], "ObjectFirewallSslSshProfile-Dot"); ok {
					if err = d.Set("dot", vv); err != nil {
						return fmt.Errorf("Error reading dot: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dot: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ftps", flattenObjectFirewallSslSshProfileFtps(o["ftps"], d, "ftps")); err != nil {
			if vv, ok := fortiAPIPatch(o["ftps"], "ObjectFirewallSslSshProfile-Ftps"); ok {
				if err = d.Set("ftps", vv); err != nil {
					return fmt.Errorf("Error reading ftps: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ftps: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftps"); ok {
			if err = d.Set("ftps", flattenObjectFirewallSslSshProfileFtps(o["ftps"], d, "ftps")); err != nil {
				if vv, ok := fortiAPIPatch(o["ftps"], "ObjectFirewallSslSshProfile-Ftps"); ok {
					if err = d.Set("ftps", vv); err != nil {
						return fmt.Errorf("Error reading ftps: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ftps: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("https", flattenObjectFirewallSslSshProfileHttps(o["https"], d, "https")); err != nil {
			if vv, ok := fortiAPIPatch(o["https"], "ObjectFirewallSslSshProfile-Https"); ok {
				if err = d.Set("https", vv); err != nil {
					return fmt.Errorf("Error reading https: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading https: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("https"); ok {
			if err = d.Set("https", flattenObjectFirewallSslSshProfileHttps(o["https"], d, "https")); err != nil {
				if vv, ok := fortiAPIPatch(o["https"], "ObjectFirewallSslSshProfile-Https"); ok {
					if err = d.Set("https", vv); err != nil {
						return fmt.Errorf("Error reading https: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading https: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("imaps", flattenObjectFirewallSslSshProfileImaps(o["imaps"], d, "imaps")); err != nil {
			if vv, ok := fortiAPIPatch(o["imaps"], "ObjectFirewallSslSshProfile-Imaps"); ok {
				if err = d.Set("imaps", vv); err != nil {
					return fmt.Errorf("Error reading imaps: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading imaps: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("imaps"); ok {
			if err = d.Set("imaps", flattenObjectFirewallSslSshProfileImaps(o["imaps"], d, "imaps")); err != nil {
				if vv, ok := fortiAPIPatch(o["imaps"], "ObjectFirewallSslSshProfile-Imaps"); ok {
					if err = d.Set("imaps", vv); err != nil {
						return fmt.Errorf("Error reading imaps: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading imaps: %v", err)
				}
			}
		}
	}

	if err = d.Set("mapi_over_https", flattenObjectFirewallSslSshProfileMapiOverHttps(o["mapi-over-https"], d, "mapi_over_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["mapi-over-https"], "ObjectFirewallSslSshProfile-MapiOverHttps"); ok {
			if err = d.Set("mapi_over_https", vv); err != nil {
				return fmt.Errorf("Error reading mapi_over_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mapi_over_https: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallSslSshProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallSslSshProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("pop3s", flattenObjectFirewallSslSshProfilePop3S(o["pop3s"], d, "pop3s")); err != nil {
			if vv, ok := fortiAPIPatch(o["pop3s"], "ObjectFirewallSslSshProfile-Pop3S"); ok {
				if err = d.Set("pop3s", vv); err != nil {
					return fmt.Errorf("Error reading pop3s: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pop3s: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pop3s"); ok {
			if err = d.Set("pop3s", flattenObjectFirewallSslSshProfilePop3S(o["pop3s"], d, "pop3s")); err != nil {
				if vv, ok := fortiAPIPatch(o["pop3s"], "ObjectFirewallSslSshProfile-Pop3S"); ok {
					if err = d.Set("pop3s", vv); err != nil {
						return fmt.Errorf("Error reading pop3s: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pop3s: %v", err)
				}
			}
		}
	}

	if err = d.Set("rpc_over_https", flattenObjectFirewallSslSshProfileRpcOverHttps(o["rpc-over-https"], d, "rpc_over_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["rpc-over-https"], "ObjectFirewallSslSshProfile-RpcOverHttps"); ok {
			if err = d.Set("rpc_over_https", vv); err != nil {
				return fmt.Errorf("Error reading rpc_over_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rpc_over_https: %v", err)
		}
	}

	if err = d.Set("server_cert", flattenObjectFirewallSslSshProfileServerCert(o["server-cert"], d, "server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-cert"], "ObjectFirewallSslSshProfile-ServerCert"); ok {
			if err = d.Set("server_cert", vv); err != nil {
				return fmt.Errorf("Error reading server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_cert: %v", err)
		}
	}

	if err = d.Set("server_cert_mode", flattenObjectFirewallSslSshProfileServerCertMode(o["server-cert-mode"], d, "server_cert_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-cert-mode"], "ObjectFirewallSslSshProfile-ServerCertMode"); ok {
			if err = d.Set("server_cert_mode", vv); err != nil {
				return fmt.Errorf("Error reading server_cert_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_cert_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("smtps", flattenObjectFirewallSslSshProfileSmtps(o["smtps"], d, "smtps")); err != nil {
			if vv, ok := fortiAPIPatch(o["smtps"], "ObjectFirewallSslSshProfile-Smtps"); ok {
				if err = d.Set("smtps", vv); err != nil {
					return fmt.Errorf("Error reading smtps: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading smtps: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("smtps"); ok {
			if err = d.Set("smtps", flattenObjectFirewallSslSshProfileSmtps(o["smtps"], d, "smtps")); err != nil {
				if vv, ok := fortiAPIPatch(o["smtps"], "ObjectFirewallSslSshProfile-Smtps"); ok {
					if err = d.Set("smtps", vv); err != nil {
						return fmt.Errorf("Error reading smtps: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading smtps: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ssh", flattenObjectFirewallSslSshProfileSsh(o["ssh"], d, "ssh")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssh"], "ObjectFirewallSslSshProfile-Ssh"); ok {
				if err = d.Set("ssh", vv); err != nil {
					return fmt.Errorf("Error reading ssh: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssh: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssh"); ok {
			if err = d.Set("ssh", flattenObjectFirewallSslSshProfileSsh(o["ssh"], d, "ssh")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssh"], "ObjectFirewallSslSshProfile-Ssh"); ok {
					if err = d.Set("ssh", vv); err != nil {
						return fmt.Errorf("Error reading ssh: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssh: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ssl", flattenObjectFirewallSslSshProfileSsl(o["ssl"], d, "ssl")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl"], "ObjectFirewallSslSshProfile-Ssl"); ok {
				if err = d.Set("ssl", vv); err != nil {
					return fmt.Errorf("Error reading ssl: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl"); ok {
			if err = d.Set("ssl", flattenObjectFirewallSslSshProfileSsl(o["ssl"], d, "ssl")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl"], "ObjectFirewallSslSshProfile-Ssl"); ok {
					if err = d.Set("ssl", vv); err != nil {
						return fmt.Errorf("Error reading ssl: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_anomalies_log", flattenObjectFirewallSslSshProfileSslAnomaliesLog(o["ssl-anomalies-log"], d, "ssl_anomalies_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-anomalies-log"], "ObjectFirewallSslSshProfile-SslAnomaliesLog"); ok {
			if err = d.Set("ssl_anomalies_log", vv); err != nil {
				return fmt.Errorf("Error reading ssl_anomalies_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_anomalies_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_exempt", flattenObjectFirewallSslSshProfileSslExempt(o["ssl-exempt"], d, "ssl_exempt")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-exempt"], "ObjectFirewallSslSshProfile-SslExempt"); ok {
				if err = d.Set("ssl_exempt", vv); err != nil {
					return fmt.Errorf("Error reading ssl_exempt: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_exempt: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_exempt"); ok {
			if err = d.Set("ssl_exempt", flattenObjectFirewallSslSshProfileSslExempt(o["ssl-exempt"], d, "ssl_exempt")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-exempt"], "ObjectFirewallSslSshProfile-SslExempt"); ok {
					if err = d.Set("ssl_exempt", vv); err != nil {
						return fmt.Errorf("Error reading ssl_exempt: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_exempt: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_exemptions_log", flattenObjectFirewallSslSshProfileSslExemptionsLog(o["ssl-exemptions-log"], d, "ssl_exemptions_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-exemptions-log"], "ObjectFirewallSslSshProfile-SslExemptionsLog"); ok {
			if err = d.Set("ssl_exemptions_log", vv); err != nil {
				return fmt.Errorf("Error reading ssl_exemptions_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_exemptions_log: %v", err)
		}
	}

	if err = d.Set("ssl_negotiation_log", flattenObjectFirewallSslSshProfileSslNegotiationLog(o["ssl-negotiation-log"], d, "ssl_negotiation_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-negotiation-log"], "ObjectFirewallSslSshProfile-SslNegotiationLog"); ok {
			if err = d.Set("ssl_negotiation_log", vv); err != nil {
				return fmt.Errorf("Error reading ssl_negotiation_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_negotiation_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_server", flattenObjectFirewallSslSshProfileSslServer(o["ssl-server"], d, "ssl_server")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssl-server"], "ObjectFirewallSslSshProfile-SslServer"); ok {
				if err = d.Set("ssl_server", vv); err != nil {
					return fmt.Errorf("Error reading ssl_server: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssl_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_server"); ok {
			if err = d.Set("ssl_server", flattenObjectFirewallSslSshProfileSslServer(o["ssl-server"], d, "ssl_server")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssl-server"], "ObjectFirewallSslSshProfile-SslServer"); ok {
					if err = d.Set("ssl_server", vv); err != nil {
						return fmt.Errorf("Error reading ssl_server: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssl_server: %v", err)
				}
			}
		}
	}

	if err = d.Set("supported_alpn", flattenObjectFirewallSslSshProfileSupportedAlpn(o["supported-alpn"], d, "supported_alpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["supported-alpn"], "ObjectFirewallSslSshProfile-SupportedAlpn"); ok {
			if err = d.Set("supported_alpn", vv); err != nil {
				return fmt.Errorf("Error reading supported_alpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading supported_alpn: %v", err)
		}
	}

	if err = d.Set("untrusted_caname", flattenObjectFirewallSslSshProfileUntrustedCaname(o["untrusted-caname"], d, "untrusted_caname")); err != nil {
		if vv, ok := fortiAPIPatch(o["untrusted-caname"], "ObjectFirewallSslSshProfile-UntrustedCaname"); ok {
			if err = d.Set("untrusted_caname", vv); err != nil {
				return fmt.Errorf("Error reading untrusted_caname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading untrusted_caname: %v", err)
		}
	}

	if err = d.Set("use_ssl_server", flattenObjectFirewallSslSshProfileUseSslServer(o["use-ssl-server"], d, "use_ssl_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-ssl-server"], "ObjectFirewallSslSshProfile-UseSslServer"); ok {
			if err = d.Set("use_ssl_server", vv); err != nil {
				return fmt.Errorf("Error reading use_ssl_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_ssl_server: %v", err)
		}
	}

	if err = d.Set("whitelist", flattenObjectFirewallSslSshProfileWhitelist(o["whitelist"], d, "whitelist")); err != nil {
		if vv, ok := fortiAPIPatch(o["whitelist"], "ObjectFirewallSslSshProfile-Whitelist"); ok {
			if err = d.Set("whitelist", vv); err != nil {
				return fmt.Errorf("Error reading whitelist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading whitelist: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallSslSshProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallSslSshProfileAllowlist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileBlockBlocklistedCertificates(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileBlockBlacklistedCertificates(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileCaname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileDot(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-failure"], _ = expandObjectFirewallSslSshProfileDotCertValidationFailure(d, i["cert_validation_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-timeout"], _ = expandObjectFirewallSslSshProfileDotCertValidationTimeout(d, i["cert_validation_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-certificate"], _ = expandObjectFirewallSslSshProfileDotClientCertificate(d, i["client_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["expired-server-cert"], _ = expandObjectFirewallSslSshProfileDotExpiredServerCert(d, i["expired_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandObjectFirewallSslSshProfileDotProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["revoked-server-cert"], _ = expandObjectFirewallSslSshProfileDotRevokedServerCert(d, i["revoked_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandObjectFirewallSslSshProfileDotSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallSslSshProfileDotStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-cipher"], _ = expandObjectFirewallSslSshProfileDotUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-negotiation"], _ = expandObjectFirewallSslSshProfileDotUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandObjectFirewallSslSshProfileDotUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallSslSshProfileDotCertValidationFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileDotCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileDotClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileDotExpiredServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileDotProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileDotRevokedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileDotSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileDotStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileDotUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileDotUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileDotUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileFtps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-failure"], _ = expandObjectFirewallSslSshProfileFtpsCertValidationFailure(d, i["cert_validation_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-timeout"], _ = expandObjectFirewallSslSshProfileFtpsCertValidationTimeout(d, i["cert_validation_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-certificate"], _ = expandObjectFirewallSslSshProfileFtpsClientCertificate(d, i["client_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["expired-server-cert"], _ = expandObjectFirewallSslSshProfileFtpsExpiredServerCert(d, i["expired_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandObjectFirewallSslSshProfileFtpsPorts(d, i["ports"], pre_append)
	} else {
		result["ports"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["revoked-server-cert"], _ = expandObjectFirewallSslSshProfileFtpsRevokedServerCert(d, i["revoked_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandObjectFirewallSslSshProfileFtpsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallSslSshProfileFtpsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-cipher"], _ = expandObjectFirewallSslSshProfileFtpsUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-negotiation"], _ = expandObjectFirewallSslSshProfileFtpsUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandObjectFirewallSslSshProfileFtpsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallSslSshProfileFtpsCertValidationFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileFtpsCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileFtpsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileFtpsExpiredServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileFtpsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallSslSshProfileFtpsRevokedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileFtpsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileFtpsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileFtpsUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileFtpsUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileFtpsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_probe_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-probe-failure"], _ = expandObjectFirewallSslSshProfileHttpsCertProbeFailure(d, i["cert_probe_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-failure"], _ = expandObjectFirewallSslSshProfileHttpsCertValidationFailure(d, i["cert_validation_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-timeout"], _ = expandObjectFirewallSslSshProfileHttpsCertValidationTimeout(d, i["cert_validation_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-certificate"], _ = expandObjectFirewallSslSshProfileHttpsClientCertificate(d, i["client_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["expired-server-cert"], _ = expandObjectFirewallSslSshProfileHttpsExpiredServerCert(d, i["expired_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandObjectFirewallSslSshProfileHttpsPorts(d, i["ports"], pre_append)
	} else {
		result["ports"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandObjectFirewallSslSshProfileHttpsProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["revoked-server-cert"], _ = expandObjectFirewallSslSshProfileHttpsRevokedServerCert(d, i["revoked_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandObjectFirewallSslSshProfileHttpsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallSslSshProfileHttpsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-cipher"], _ = expandObjectFirewallSslSshProfileHttpsUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-negotiation"], _ = expandObjectFirewallSslSshProfileHttpsUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandObjectFirewallSslSshProfileHttpsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallSslSshProfileHttpsCertProbeFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileHttpsCertValidationFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileHttpsCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileHttpsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileHttpsExpiredServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileHttpsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallSslSshProfileHttpsProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileHttpsRevokedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileHttpsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileHttpsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileHttpsUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileHttpsUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileHttpsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-failure"], _ = expandObjectFirewallSslSshProfileImapsCertValidationFailure(d, i["cert_validation_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-timeout"], _ = expandObjectFirewallSslSshProfileImapsCertValidationTimeout(d, i["cert_validation_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-certificate"], _ = expandObjectFirewallSslSshProfileImapsClientCertificate(d, i["client_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["expired-server-cert"], _ = expandObjectFirewallSslSshProfileImapsExpiredServerCert(d, i["expired_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandObjectFirewallSslSshProfileImapsPorts(d, i["ports"], pre_append)
	} else {
		result["ports"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandObjectFirewallSslSshProfileImapsProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["revoked-server-cert"], _ = expandObjectFirewallSslSshProfileImapsRevokedServerCert(d, i["revoked_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandObjectFirewallSslSshProfileImapsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallSslSshProfileImapsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-cipher"], _ = expandObjectFirewallSslSshProfileImapsUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-negotiation"], _ = expandObjectFirewallSslSshProfileImapsUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandObjectFirewallSslSshProfileImapsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallSslSshProfileImapsCertValidationFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsExpiredServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallSslSshProfileImapsProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsRevokedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileImapsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileMapiOverHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfilePop3S(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-failure"], _ = expandObjectFirewallSslSshProfilePop3SCertValidationFailure(d, i["cert_validation_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-timeout"], _ = expandObjectFirewallSslSshProfilePop3SCertValidationTimeout(d, i["cert_validation_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-certificate"], _ = expandObjectFirewallSslSshProfilePop3SClientCertificate(d, i["client_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["expired-server-cert"], _ = expandObjectFirewallSslSshProfilePop3SExpiredServerCert(d, i["expired_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandObjectFirewallSslSshProfilePop3SPorts(d, i["ports"], pre_append)
	} else {
		result["ports"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandObjectFirewallSslSshProfilePop3SProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["revoked-server-cert"], _ = expandObjectFirewallSslSshProfilePop3SRevokedServerCert(d, i["revoked_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandObjectFirewallSslSshProfilePop3SSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallSslSshProfilePop3SStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-cipher"], _ = expandObjectFirewallSslSshProfilePop3SUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-negotiation"], _ = expandObjectFirewallSslSshProfilePop3SUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandObjectFirewallSslSshProfilePop3SUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallSslSshProfilePop3SCertValidationFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfilePop3SCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfilePop3SClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfilePop3SExpiredServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfilePop3SPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallSslSshProfilePop3SProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfilePop3SRevokedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfilePop3SSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfilePop3SStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfilePop3SUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfilePop3SUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfilePop3SUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileRpcOverHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileServerCertMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSmtps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-failure"], _ = expandObjectFirewallSslSshProfileSmtpsCertValidationFailure(d, i["cert_validation_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-timeout"], _ = expandObjectFirewallSslSshProfileSmtpsCertValidationTimeout(d, i["cert_validation_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-certificate"], _ = expandObjectFirewallSslSshProfileSmtpsClientCertificate(d, i["client_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["expired-server-cert"], _ = expandObjectFirewallSslSshProfileSmtpsExpiredServerCert(d, i["expired_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandObjectFirewallSslSshProfileSmtpsPorts(d, i["ports"], pre_append)
	} else {
		result["ports"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandObjectFirewallSslSshProfileSmtpsProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["revoked-server-cert"], _ = expandObjectFirewallSslSshProfileSmtpsRevokedServerCert(d, i["revoked_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandObjectFirewallSslSshProfileSmtpsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallSslSshProfileSmtpsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-cipher"], _ = expandObjectFirewallSslSshProfileSmtpsUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-negotiation"], _ = expandObjectFirewallSslSshProfileSmtpsUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandObjectFirewallSslSshProfileSmtpsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallSslSshProfileSmtpsCertValidationFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSmtpsCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSmtpsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSmtpsExpiredServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSmtpsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallSslSshProfileSmtpsProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSmtpsRevokedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSmtpsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSmtpsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSmtpsUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSmtpsUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSmtpsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSsh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandObjectFirewallSslSshProfileSshInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandObjectFirewallSslSshProfileSshPorts(d, i["ports"], pre_append)
	} else {
		result["ports"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandObjectFirewallSslSshProfileSshProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "ssh_algorithm"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssh-algorithm"], _ = expandObjectFirewallSslSshProfileSshSshAlgorithm(d, i["ssh_algorithm"], pre_append)
	}
	pre_append = pre + ".0." + "ssh_tun_policy_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssh-tun-policy-check"], _ = expandObjectFirewallSslSshProfileSshSshTunPolicyCheck(d, i["ssh_tun_policy_check"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallSslSshProfileSshStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-version"], _ = expandObjectFirewallSslSshProfileSshUnsupportedVersion(d, i["unsupported_version"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallSslSshProfileSshInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSshPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallSslSshProfileSshProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSshSshAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSshSshTunPolicyCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSshStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSshUnsupportedVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-failure"], _ = expandObjectFirewallSslSshProfileSslCertValidationFailure(d, i["cert_validation_failure"], pre_append)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-timeout"], _ = expandObjectFirewallSslSshProfileSslCertValidationTimeout(d, i["cert_validation_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-certificate"], _ = expandObjectFirewallSslSshProfileSslClientCertificate(d, i["client_certificate"], pre_append)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["expired-server-cert"], _ = expandObjectFirewallSslSshProfileSslExpiredServerCert(d, i["expired_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandObjectFirewallSslSshProfileSslInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["revoked-server-cert"], _ = expandObjectFirewallSslSshProfileSslRevokedServerCert(d, i["revoked_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandObjectFirewallSslSshProfileSslSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-cipher"], _ = expandObjectFirewallSslSshProfileSslUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-negotiation"], _ = expandObjectFirewallSslSshProfileSslUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandObjectFirewallSslSshProfileSslUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallSslSshProfileSslCertValidationFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslExpiredServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslRevokedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslAnomaliesLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslExempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address"], _ = expandObjectFirewallSslSshProfileSslExemptAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address6"], _ = expandObjectFirewallSslSshProfileSslExemptAddress6(d, i["address6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fortiguard-category"], _ = expandObjectFirewallSslSshProfileSslExemptFortiguardCategory(d, i["fortiguard_category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectFirewallSslSshProfileSslExemptId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["regex"], _ = expandObjectFirewallSslSshProfileSslExemptRegex(d, i["regex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandObjectFirewallSslSshProfileSslExemptType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard_fqdn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["wildcard-fqdn"], _ = expandObjectFirewallSslSshProfileSslExemptWildcardFqdn(d, i["wildcard_fqdn"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallSslSshProfileSslExemptAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslExemptAddress6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslExemptFortiguardCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectFirewallSslSshProfileSslExemptId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslExemptRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslExemptType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslExemptWildcardFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectFirewallSslSshProfileSslExemptionsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslNegotiationLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftps_client_certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ftps-client-certificate"], _ = expandObjectFirewallSslSshProfileSslServerFtpsClientCertificate(d, i["ftps_client_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_client_certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["https-client-certificate"], _ = expandObjectFirewallSslSshProfileSslServerHttpsClientCertificate(d, i["https_client_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftps_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ftps-client-cert-request"], _ = expandObjectFirewallSslSshProfileSslServerFtpsClientCertRequest(d, i["ftps_client_cert_request"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["https-client-cert-request"], _ = expandObjectFirewallSslSshProfileSslServerHttpsClientCertRequest(d, i["https_client_cert_request"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectFirewallSslSshProfileSslServerId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imaps_client_certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["imaps-client-certificate"], _ = expandObjectFirewallSslSshProfileSslServerImapsClientCertificate(d, i["imaps_client_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imaps_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["imaps-client-cert-request"], _ = expandObjectFirewallSslSshProfileSslServerImapsClientCertRequest(d, i["imaps_client_cert_request"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandObjectFirewallSslSshProfileSslServerIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pop3s_client_certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pop3s-client-certificate"], _ = expandObjectFirewallSslSshProfileSslServerPop3SClientCertificate(d, i["pop3s_client_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smtps_client_certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["smtps-client-certificate"], _ = expandObjectFirewallSslSshProfileSslServerSmtpsClientCertificate(d, i["smtps_client_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_other_client_certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-other-client-certificate"], _ = expandObjectFirewallSslSshProfileSslServerSslOtherClientCertificate(d, i["ssl_other_client_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pop3s_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pop3s-client-cert-request"], _ = expandObjectFirewallSslSshProfileSslServerPop3SClientCertRequest(d, i["pop3s_client_cert_request"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smtps_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["smtps-client-cert-request"], _ = expandObjectFirewallSslSshProfileSslServerSmtpsClientCertRequest(d, i["smtps_client_cert_request"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_other_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-other-client-cert-request"], _ = expandObjectFirewallSslSshProfileSslServerSslOtherClientCertRequest(d, i["ssl_other_client_cert_request"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallSslSshProfileSslServerFtpsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerHttpsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerFtpsClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerHttpsClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerImapsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerImapsClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerPop3SClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerSmtpsClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerSslOtherClientCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerPop3SClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerSmtpsClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslServerSslOtherClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSupportedAlpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileUntrustedCaname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileUseSslServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileWhitelist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallSslSshProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allowlist"); ok {
		t, err := expandObjectFirewallSslSshProfileAllowlist(d, v, "allowlist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowlist"] = t
		}
	}

	if v, ok := d.GetOk("block_blocklisted_certificates"); ok {
		t, err := expandObjectFirewallSslSshProfileBlockBlocklistedCertificates(d, v, "block_blocklisted_certificates")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-blocklisted-certificates"] = t
		}
	}

	if v, ok := d.GetOk("block_blacklisted_certificates"); ok {
		t, err := expandObjectFirewallSslSshProfileBlockBlacklistedCertificates(d, v, "block_blacklisted_certificates")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-blacklisted-certificates"] = t
		}
	}

	if v, ok := d.GetOk("caname"); ok {
		t, err := expandObjectFirewallSslSshProfileCaname(d, v, "caname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["caname"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectFirewallSslSshProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dot"); ok {
		t, err := expandObjectFirewallSslSshProfileDot(d, v, "dot")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dot"] = t
		}
	}

	if v, ok := d.GetOk("ftps"); ok {
		t, err := expandObjectFirewallSslSshProfileFtps(d, v, "ftps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftps"] = t
		}
	}

	if v, ok := d.GetOk("https"); ok {
		t, err := expandObjectFirewallSslSshProfileHttps(d, v, "https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https"] = t
		}
	}

	if v, ok := d.GetOk("imaps"); ok {
		t, err := expandObjectFirewallSslSshProfileImaps(d, v, "imaps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imaps"] = t
		}
	}

	if v, ok := d.GetOk("mapi_over_https"); ok {
		t, err := expandObjectFirewallSslSshProfileMapiOverHttps(d, v, "mapi_over_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi-over-https"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectFirewallSslSshProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pop3s"); ok {
		t, err := expandObjectFirewallSslSshProfilePop3S(d, v, "pop3s")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3s"] = t
		}
	}

	if v, ok := d.GetOk("rpc_over_https"); ok {
		t, err := expandObjectFirewallSslSshProfileRpcOverHttps(d, v, "rpc_over_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-over-https"] = t
		}
	}

	if v, ok := d.GetOk("server_cert"); ok {
		t, err := expandObjectFirewallSslSshProfileServerCert(d, v, "server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-cert"] = t
		}
	}

	if v, ok := d.GetOk("server_cert_mode"); ok {
		t, err := expandObjectFirewallSslSshProfileServerCertMode(d, v, "server_cert_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-cert-mode"] = t
		}
	}

	if v, ok := d.GetOk("smtps"); ok {
		t, err := expandObjectFirewallSslSshProfileSmtps(d, v, "smtps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtps"] = t
		}
	}

	if v, ok := d.GetOk("ssh"); ok {
		t, err := expandObjectFirewallSslSshProfileSsh(d, v, "ssh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh"] = t
		}
	}

	if v, ok := d.GetOk("ssl"); ok {
		t, err := expandObjectFirewallSslSshProfileSsl(d, v, "ssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}

	if v, ok := d.GetOk("ssl_anomalies_log"); ok {
		t, err := expandObjectFirewallSslSshProfileSslAnomaliesLog(d, v, "ssl_anomalies_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-anomalies-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_exempt"); ok {
		t, err := expandObjectFirewallSslSshProfileSslExempt(d, v, "ssl_exempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-exempt"] = t
		}
	}

	if v, ok := d.GetOk("ssl_exemptions_log"); ok {
		t, err := expandObjectFirewallSslSshProfileSslExemptionsLog(d, v, "ssl_exemptions_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-exemptions-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_negotiation_log"); ok {
		t, err := expandObjectFirewallSslSshProfileSslNegotiationLog(d, v, "ssl_negotiation_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-negotiation-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server"); ok {
		t, err := expandObjectFirewallSslSshProfileSslServer(d, v, "ssl_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server"] = t
		}
	}

	if v, ok := d.GetOk("supported_alpn"); ok {
		t, err := expandObjectFirewallSslSshProfileSupportedAlpn(d, v, "supported_alpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["supported-alpn"] = t
		}
	}

	if v, ok := d.GetOk("untrusted_caname"); ok {
		t, err := expandObjectFirewallSslSshProfileUntrustedCaname(d, v, "untrusted_caname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["untrusted-caname"] = t
		}
	}

	if v, ok := d.GetOk("use_ssl_server"); ok {
		t, err := expandObjectFirewallSslSshProfileUseSslServer(d, v, "use_ssl_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-ssl-server"] = t
		}
	}

	if v, ok := d.GetOk("whitelist"); ok {
		t, err := expandObjectFirewallSslSshProfileWhitelist(d, v, "whitelist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["whitelist"] = t
		}
	}

	return &obj, nil
}
