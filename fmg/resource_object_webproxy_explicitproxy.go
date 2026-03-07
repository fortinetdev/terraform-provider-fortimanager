// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectWebProxy ExplicitProxy

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebProxyExplicitProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebProxyExplicitProxyCreate,
		Read:   resourceObjectWebProxyExplicitProxyRead,
		Update: resourceObjectWebProxyExplicitProxyUpdate,
		Delete: resourceObjectWebProxyExplicitProxyDelete,

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
			"detect_https_in_http_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstport_from_incoming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"empty_cert_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftp_incoming_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ftp_over_http": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"header_proxy_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_connection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_incoming_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"https_incoming_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"incoming_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv6_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"learn_dst_from_sni": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"pac_file_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pac_file_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pac_file_server_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pac_file_server_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pac_file_through_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pac_file_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pref_dns_result": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"realm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"return_to_sender": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sec_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secure_web_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secure_web_proxy_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"socks": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"socks_incoming_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_dh_bits": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_http_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_agent_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWebProxyExplicitProxyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWebProxyExplicitProxy(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebProxyExplicitProxy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWebProxyExplicitProxy(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebProxyExplicitProxy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebProxyExplicitProxyRead(d, m)
}

func resourceObjectWebProxyExplicitProxyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWebProxyExplicitProxy(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebProxyExplicitProxy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWebProxyExplicitProxy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebProxyExplicitProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebProxyExplicitProxyRead(d, m)
}

func resourceObjectWebProxyExplicitProxyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWebProxyExplicitProxy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebProxyExplicitProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebProxyExplicitProxyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWebProxyExplicitProxy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebProxyExplicitProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebProxyExplicitProxy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebProxyExplicitProxy resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebProxyExplicitProxyDetectHttpsInHttpRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyDnsMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyDstportFromIncoming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyEmptyCertAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyFtpIncomingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebProxyExplicitProxyFtpOverHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyHeaderProxyAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyHttpConnectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyHttpIncomingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebProxyExplicitProxyHttpsIncomingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebProxyExplicitProxyIncomingIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebProxyExplicitProxyIpv6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyLearnDstFromSni(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyPacFileData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyPacFileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyPacFileServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyPacFileServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyPacFileThroughHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyPacFileUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyPrefDnsResult(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyReturnToSender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxySecDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxySecureWebProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxySecureWebProxyCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebProxyExplicitProxySocks(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxySocksIncomingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebProxyExplicitProxySslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxySslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyUnknownHttpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyExplicitProxyUserAgentDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebProxyExplicitProxy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("detect_https_in_http_request", flattenObjectWebProxyExplicitProxyDetectHttpsInHttpRequest(o["detect-https-in-http-request"], d, "detect_https_in_http_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-https-in-http-request"], "ObjectWebProxyExplicitProxy-DetectHttpsInHttpRequest"); ok {
			if err = d.Set("detect_https_in_http_request", vv); err != nil {
				return fmt.Errorf("Error reading detect_https_in_http_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_https_in_http_request: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenObjectWebProxyExplicitProxyClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "ObjectWebProxyExplicitProxy-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("dns_mode", flattenObjectWebProxyExplicitProxyDnsMode(o["dns-mode"], d, "dns_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-mode"], "ObjectWebProxyExplicitProxy-DnsMode"); ok {
			if err = d.Set("dns_mode", vv); err != nil {
				return fmt.Errorf("Error reading dns_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_mode: %v", err)
		}
	}

	if err = d.Set("dstport_from_incoming", flattenObjectWebProxyExplicitProxyDstportFromIncoming(o["dstport-from-incoming"], d, "dstport_from_incoming")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstport-from-incoming"], "ObjectWebProxyExplicitProxy-DstportFromIncoming"); ok {
			if err = d.Set("dstport_from_incoming", vv); err != nil {
				return fmt.Errorf("Error reading dstport_from_incoming: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstport_from_incoming: %v", err)
		}
	}

	if err = d.Set("empty_cert_action", flattenObjectWebProxyExplicitProxyEmptyCertAction(o["empty-cert-action"], d, "empty_cert_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["empty-cert-action"], "ObjectWebProxyExplicitProxy-EmptyCertAction"); ok {
			if err = d.Set("empty_cert_action", vv); err != nil {
				return fmt.Errorf("Error reading empty_cert_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading empty_cert_action: %v", err)
		}
	}

	if err = d.Set("ftp_incoming_port", flattenObjectWebProxyExplicitProxyFtpIncomingPort(o["ftp-incoming-port"], d, "ftp_incoming_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftp-incoming-port"], "ObjectWebProxyExplicitProxy-FtpIncomingPort"); ok {
			if err = d.Set("ftp_incoming_port", vv); err != nil {
				return fmt.Errorf("Error reading ftp_incoming_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftp_incoming_port: %v", err)
		}
	}

	if err = d.Set("ftp_over_http", flattenObjectWebProxyExplicitProxyFtpOverHttp(o["ftp-over-http"], d, "ftp_over_http")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftp-over-http"], "ObjectWebProxyExplicitProxy-FtpOverHttp"); ok {
			if err = d.Set("ftp_over_http", vv); err != nil {
				return fmt.Errorf("Error reading ftp_over_http: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftp_over_http: %v", err)
		}
	}

	if err = d.Set("header_proxy_agent", flattenObjectWebProxyExplicitProxyHeaderProxyAgent(o["header-proxy-agent"], d, "header_proxy_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-proxy-agent"], "ObjectWebProxyExplicitProxy-HeaderProxyAgent"); ok {
			if err = d.Set("header_proxy_agent", vv); err != nil {
				return fmt.Errorf("Error reading header_proxy_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_proxy_agent: %v", err)
		}
	}

	if err = d.Set("http", flattenObjectWebProxyExplicitProxyHttp(o["http"], d, "http")); err != nil {
		if vv, ok := fortiAPIPatch(o["http"], "ObjectWebProxyExplicitProxy-Http"); ok {
			if err = d.Set("http", vv); err != nil {
				return fmt.Errorf("Error reading http: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http: %v", err)
		}
	}

	if err = d.Set("http_connection_mode", flattenObjectWebProxyExplicitProxyHttpConnectionMode(o["http-connection-mode"], d, "http_connection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-connection-mode"], "ObjectWebProxyExplicitProxy-HttpConnectionMode"); ok {
			if err = d.Set("http_connection_mode", vv); err != nil {
				return fmt.Errorf("Error reading http_connection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_connection_mode: %v", err)
		}
	}

	if err = d.Set("http_incoming_port", flattenObjectWebProxyExplicitProxyHttpIncomingPort(o["http-incoming-port"], d, "http_incoming_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-incoming-port"], "ObjectWebProxyExplicitProxy-HttpIncomingPort"); ok {
			if err = d.Set("http_incoming_port", vv); err != nil {
				return fmt.Errorf("Error reading http_incoming_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_incoming_port: %v", err)
		}
	}

	if err = d.Set("https_incoming_port", flattenObjectWebProxyExplicitProxyHttpsIncomingPort(o["https-incoming-port"], d, "https_incoming_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-incoming-port"], "ObjectWebProxyExplicitProxy-HttpsIncomingPort"); ok {
			if err = d.Set("https_incoming_port", vv); err != nil {
				return fmt.Errorf("Error reading https_incoming_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_incoming_port: %v", err)
		}
	}

	if err = d.Set("incoming_ip6", flattenObjectWebProxyExplicitProxyIncomingIp6(o["incoming-ip6"], d, "incoming_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["incoming-ip6"], "ObjectWebProxyExplicitProxy-IncomingIp6"); ok {
			if err = d.Set("incoming_ip6", vv); err != nil {
				return fmt.Errorf("Error reading incoming_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incoming_ip6: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectWebProxyExplicitProxyInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectWebProxyExplicitProxy-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ipv6_status", flattenObjectWebProxyExplicitProxyIpv6Status(o["ipv6-status"], d, "ipv6_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-status"], "ObjectWebProxyExplicitProxy-Ipv6Status"); ok {
			if err = d.Set("ipv6_status", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_status: %v", err)
		}
	}

	if err = d.Set("learn_dst_from_sni", flattenObjectWebProxyExplicitProxyLearnDstFromSni(o["learn-dst-from-sni"], d, "learn_dst_from_sni")); err != nil {
		if vv, ok := fortiAPIPatch(o["learn-dst-from-sni"], "ObjectWebProxyExplicitProxy-LearnDstFromSni"); ok {
			if err = d.Set("learn_dst_from_sni", vv); err != nil {
				return fmt.Errorf("Error reading learn_dst_from_sni: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading learn_dst_from_sni: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWebProxyExplicitProxyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWebProxyExplicitProxy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pac_file_data", flattenObjectWebProxyExplicitProxyPacFileData(o["pac-file-data"], d, "pac_file_data")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-data"], "ObjectWebProxyExplicitProxy-PacFileData"); ok {
			if err = d.Set("pac_file_data", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_data: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_data: %v", err)
		}
	}

	if err = d.Set("pac_file_name", flattenObjectWebProxyExplicitProxyPacFileName(o["pac-file-name"], d, "pac_file_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-name"], "ObjectWebProxyExplicitProxy-PacFileName"); ok {
			if err = d.Set("pac_file_name", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_name: %v", err)
		}
	}

	if err = d.Set("pac_file_server_port", flattenObjectWebProxyExplicitProxyPacFileServerPort(o["pac-file-server-port"], d, "pac_file_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-server-port"], "ObjectWebProxyExplicitProxy-PacFileServerPort"); ok {
			if err = d.Set("pac_file_server_port", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_server_port: %v", err)
		}
	}

	if err = d.Set("pac_file_server_status", flattenObjectWebProxyExplicitProxyPacFileServerStatus(o["pac-file-server-status"], d, "pac_file_server_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-server-status"], "ObjectWebProxyExplicitProxy-PacFileServerStatus"); ok {
			if err = d.Set("pac_file_server_status", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_server_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_server_status: %v", err)
		}
	}

	if err = d.Set("pac_file_through_https", flattenObjectWebProxyExplicitProxyPacFileThroughHttps(o["pac-file-through-https"], d, "pac_file_through_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-through-https"], "ObjectWebProxyExplicitProxy-PacFileThroughHttps"); ok {
			if err = d.Set("pac_file_through_https", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_through_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_through_https: %v", err)
		}
	}

	if err = d.Set("pac_file_url", flattenObjectWebProxyExplicitProxyPacFileUrl(o["pac-file-url"], d, "pac_file_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-url"], "ObjectWebProxyExplicitProxy-PacFileUrl"); ok {
			if err = d.Set("pac_file_url", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_url: %v", err)
		}
	}

	if err = d.Set("pref_dns_result", flattenObjectWebProxyExplicitProxyPrefDnsResult(o["pref-dns-result"], d, "pref_dns_result")); err != nil {
		if vv, ok := fortiAPIPatch(o["pref-dns-result"], "ObjectWebProxyExplicitProxy-PrefDnsResult"); ok {
			if err = d.Set("pref_dns_result", vv); err != nil {
				return fmt.Errorf("Error reading pref_dns_result: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pref_dns_result: %v", err)
		}
	}

	if err = d.Set("realm", flattenObjectWebProxyExplicitProxyRealm(o["realm"], d, "realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["realm"], "ObjectWebProxyExplicitProxy-Realm"); ok {
			if err = d.Set("realm", vv); err != nil {
				return fmt.Errorf("Error reading realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading realm: %v", err)
		}
	}

	if err = d.Set("return_to_sender", flattenObjectWebProxyExplicitProxyReturnToSender(o["return-to-sender"], d, "return_to_sender")); err != nil {
		if vv, ok := fortiAPIPatch(o["return-to-sender"], "ObjectWebProxyExplicitProxy-ReturnToSender"); ok {
			if err = d.Set("return_to_sender", vv); err != nil {
				return fmt.Errorf("Error reading return_to_sender: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading return_to_sender: %v", err)
		}
	}

	if err = d.Set("sec_default_action", flattenObjectWebProxyExplicitProxySecDefaultAction(o["sec-default-action"], d, "sec_default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["sec-default-action"], "ObjectWebProxyExplicitProxy-SecDefaultAction"); ok {
			if err = d.Set("sec_default_action", vv); err != nil {
				return fmt.Errorf("Error reading sec_default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sec_default_action: %v", err)
		}
	}

	if err = d.Set("secure_web_proxy", flattenObjectWebProxyExplicitProxySecureWebProxy(o["secure-web-proxy"], d, "secure_web_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-web-proxy"], "ObjectWebProxyExplicitProxy-SecureWebProxy"); ok {
			if err = d.Set("secure_web_proxy", vv); err != nil {
				return fmt.Errorf("Error reading secure_web_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_web_proxy: %v", err)
		}
	}

	if err = d.Set("secure_web_proxy_cert", flattenObjectWebProxyExplicitProxySecureWebProxyCert(o["secure-web-proxy-cert"], d, "secure_web_proxy_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-web-proxy-cert"], "ObjectWebProxyExplicitProxy-SecureWebProxyCert"); ok {
			if err = d.Set("secure_web_proxy_cert", vv); err != nil {
				return fmt.Errorf("Error reading secure_web_proxy_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_web_proxy_cert: %v", err)
		}
	}

	if err = d.Set("socks", flattenObjectWebProxyExplicitProxySocks(o["socks"], d, "socks")); err != nil {
		if vv, ok := fortiAPIPatch(o["socks"], "ObjectWebProxyExplicitProxy-Socks"); ok {
			if err = d.Set("socks", vv); err != nil {
				return fmt.Errorf("Error reading socks: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading socks: %v", err)
		}
	}

	if err = d.Set("socks_incoming_port", flattenObjectWebProxyExplicitProxySocksIncomingPort(o["socks-incoming-port"], d, "socks_incoming_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["socks-incoming-port"], "ObjectWebProxyExplicitProxy-SocksIncomingPort"); ok {
			if err = d.Set("socks_incoming_port", vv); err != nil {
				return fmt.Errorf("Error reading socks_incoming_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading socks_incoming_port: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenObjectWebProxyExplicitProxySslAlgorithm(o["ssl-algorithm"], d, "ssl_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-algorithm"], "ObjectWebProxyExplicitProxy-SslAlgorithm"); ok {
			if err = d.Set("ssl_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_dh_bits", flattenObjectWebProxyExplicitProxySslDhBits(o["ssl-dh-bits"], d, "ssl_dh_bits")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-dh-bits"], "ObjectWebProxyExplicitProxy-SslDhBits"); ok {
			if err = d.Set("ssl_dh_bits", vv); err != nil {
				return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectWebProxyExplicitProxyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectWebProxyExplicitProxy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("unknown_http_version", flattenObjectWebProxyExplicitProxyUnknownHttpVersion(o["unknown-http-version"], d, "unknown_http_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-http-version"], "ObjectWebProxyExplicitProxy-UnknownHttpVersion"); ok {
			if err = d.Set("unknown_http_version", vv); err != nil {
				return fmt.Errorf("Error reading unknown_http_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_http_version: %v", err)
		}
	}

	if err = d.Set("user_agent_detect", flattenObjectWebProxyExplicitProxyUserAgentDetect(o["user-agent-detect"], d, "user_agent_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-agent-detect"], "ObjectWebProxyExplicitProxy-UserAgentDetect"); ok {
			if err = d.Set("user_agent_detect", vv); err != nil {
				return fmt.Errorf("Error reading user_agent_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_agent_detect: %v", err)
		}
	}

	return nil
}

func flattenObjectWebProxyExplicitProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebProxyExplicitProxyDetectHttpsInHttpRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyDnsMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyDstportFromIncoming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyEmptyCertAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyFtpIncomingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebProxyExplicitProxyFtpOverHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyHeaderProxyAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyHttpConnectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyHttpIncomingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebProxyExplicitProxyHttpsIncomingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebProxyExplicitProxyIncomingIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebProxyExplicitProxyIpv6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyLearnDstFromSni(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyPacFileData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyPacFileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyPacFileServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyPacFileServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyPacFileThroughHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyPacFileUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyPrefDnsResult(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyReturnToSender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxySecDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxySecureWebProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxySecureWebProxyCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebProxyExplicitProxySocks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxySocksIncomingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebProxyExplicitProxySslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxySslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyUnknownHttpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyExplicitProxyUserAgentDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebProxyExplicitProxy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("detect_https_in_http_request"); ok || d.HasChange("detect_https_in_http_request") {
		t, err := expandObjectWebProxyExplicitProxyDetectHttpsInHttpRequest(d, v, "detect_https_in_http_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-https-in-http-request"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandObjectWebProxyExplicitProxyClientCert(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("dns_mode"); ok || d.HasChange("dns_mode") {
		t, err := expandObjectWebProxyExplicitProxyDnsMode(d, v, "dns_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-mode"] = t
		}
	}

	if v, ok := d.GetOk("dstport_from_incoming"); ok || d.HasChange("dstport_from_incoming") {
		t, err := expandObjectWebProxyExplicitProxyDstportFromIncoming(d, v, "dstport_from_incoming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstport-from-incoming"] = t
		}
	}

	if v, ok := d.GetOk("empty_cert_action"); ok || d.HasChange("empty_cert_action") {
		t, err := expandObjectWebProxyExplicitProxyEmptyCertAction(d, v, "empty_cert_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-cert-action"] = t
		}
	}

	if v, ok := d.GetOk("ftp_incoming_port"); ok || d.HasChange("ftp_incoming_port") {
		t, err := expandObjectWebProxyExplicitProxyFtpIncomingPort(d, v, "ftp_incoming_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("ftp_over_http"); ok || d.HasChange("ftp_over_http") {
		t, err := expandObjectWebProxyExplicitProxyFtpOverHttp(d, v, "ftp_over_http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-over-http"] = t
		}
	}

	if v, ok := d.GetOk("header_proxy_agent"); ok || d.HasChange("header_proxy_agent") {
		t, err := expandObjectWebProxyExplicitProxyHeaderProxyAgent(d, v, "header_proxy_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-proxy-agent"] = t
		}
	}

	if v, ok := d.GetOk("http"); ok || d.HasChange("http") {
		t, err := expandObjectWebProxyExplicitProxyHttp(d, v, "http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http"] = t
		}
	}

	if v, ok := d.GetOk("http_connection_mode"); ok || d.HasChange("http_connection_mode") {
		t, err := expandObjectWebProxyExplicitProxyHttpConnectionMode(d, v, "http_connection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-connection-mode"] = t
		}
	}

	if v, ok := d.GetOk("http_incoming_port"); ok || d.HasChange("http_incoming_port") {
		t, err := expandObjectWebProxyExplicitProxyHttpIncomingPort(d, v, "http_incoming_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("https_incoming_port"); ok || d.HasChange("https_incoming_port") {
		t, err := expandObjectWebProxyExplicitProxyHttpsIncomingPort(d, v, "https_incoming_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("incoming_ip6"); ok || d.HasChange("incoming_ip6") {
		t, err := expandObjectWebProxyExplicitProxyIncomingIp6(d, v, "incoming_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-ip6"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectWebProxyExplicitProxyInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_status"); ok || d.HasChange("ipv6_status") {
		t, err := expandObjectWebProxyExplicitProxyIpv6Status(d, v, "ipv6_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-status"] = t
		}
	}

	if v, ok := d.GetOk("learn_dst_from_sni"); ok || d.HasChange("learn_dst_from_sni") {
		t, err := expandObjectWebProxyExplicitProxyLearnDstFromSni(d, v, "learn_dst_from_sni")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learn-dst-from-sni"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWebProxyExplicitProxyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_data"); ok || d.HasChange("pac_file_data") {
		t, err := expandObjectWebProxyExplicitProxyPacFileData(d, v, "pac_file_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-data"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_name"); ok || d.HasChange("pac_file_name") {
		t, err := expandObjectWebProxyExplicitProxyPacFileName(d, v, "pac_file_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-name"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_server_port"); ok || d.HasChange("pac_file_server_port") {
		t, err := expandObjectWebProxyExplicitProxyPacFileServerPort(d, v, "pac_file_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-server-port"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_server_status"); ok || d.HasChange("pac_file_server_status") {
		t, err := expandObjectWebProxyExplicitProxyPacFileServerStatus(d, v, "pac_file_server_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-server-status"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_through_https"); ok || d.HasChange("pac_file_through_https") {
		t, err := expandObjectWebProxyExplicitProxyPacFileThroughHttps(d, v, "pac_file_through_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-through-https"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_url"); ok || d.HasChange("pac_file_url") {
		t, err := expandObjectWebProxyExplicitProxyPacFileUrl(d, v, "pac_file_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-url"] = t
		}
	}

	if v, ok := d.GetOk("pref_dns_result"); ok || d.HasChange("pref_dns_result") {
		t, err := expandObjectWebProxyExplicitProxyPrefDnsResult(d, v, "pref_dns_result")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pref-dns-result"] = t
		}
	}

	if v, ok := d.GetOk("realm"); ok || d.HasChange("realm") {
		t, err := expandObjectWebProxyExplicitProxyRealm(d, v, "realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realm"] = t
		}
	}

	if v, ok := d.GetOk("return_to_sender"); ok || d.HasChange("return_to_sender") {
		t, err := expandObjectWebProxyExplicitProxyReturnToSender(d, v, "return_to_sender")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["return-to-sender"] = t
		}
	}

	if v, ok := d.GetOk("sec_default_action"); ok || d.HasChange("sec_default_action") {
		t, err := expandObjectWebProxyExplicitProxySecDefaultAction(d, v, "sec_default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sec-default-action"] = t
		}
	}

	if v, ok := d.GetOk("secure_web_proxy"); ok || d.HasChange("secure_web_proxy") {
		t, err := expandObjectWebProxyExplicitProxySecureWebProxy(d, v, "secure_web_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-web-proxy"] = t
		}
	}

	if v, ok := d.GetOk("secure_web_proxy_cert"); ok || d.HasChange("secure_web_proxy_cert") {
		t, err := expandObjectWebProxyExplicitProxySecureWebProxyCert(d, v, "secure_web_proxy_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-web-proxy-cert"] = t
		}
	}

	if v, ok := d.GetOk("socks"); ok || d.HasChange("socks") {
		t, err := expandObjectWebProxyExplicitProxySocks(d, v, "socks")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["socks"] = t
		}
	}

	if v, ok := d.GetOk("socks_incoming_port"); ok || d.HasChange("socks_incoming_port") {
		t, err := expandObjectWebProxyExplicitProxySocksIncomingPort(d, v, "socks_incoming_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["socks-incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok || d.HasChange("ssl_algorithm") {
		t, err := expandObjectWebProxyExplicitProxySslAlgorithm(d, v, "ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok || d.HasChange("ssl_dh_bits") {
		t, err := expandObjectWebProxyExplicitProxySslDhBits(d, v, "ssl_dh_bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectWebProxyExplicitProxyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("unknown_http_version"); ok || d.HasChange("unknown_http_version") {
		t, err := expandObjectWebProxyExplicitProxyUnknownHttpVersion(d, v, "unknown_http_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-http-version"] = t
		}
	}

	if v, ok := d.GetOk("user_agent_detect"); ok || d.HasChange("user_agent_detect") {
		t, err := expandObjectWebProxyExplicitProxyUserAgentDetect(d, v, "user_agent_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-agent-detect"] = t
		}
	}

	return &obj, nil
}
