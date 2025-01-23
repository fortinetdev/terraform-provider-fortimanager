// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure HTTP protocol options.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallProfileProtocolOptionsHttp() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallProfileProtocolOptionsHttpUpdate,
		Read:   resourceObjectFirewallProfileProtocolOptionsHttpRead,
		Update: resourceObjectFirewallProfileProtocolOptionsHttpUpdate,
		Delete: resourceObjectFirewallProfileProtocolOptionsHttpDelete,

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
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address_ip_rating": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_page_status_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comfort_amount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comfort_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"domain_fronting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"h2c": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortinet_bar": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortinet_bar_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"inspect_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"oversize_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"post_lang": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"proxy_after_tcp_handshake": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"range_block": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retry_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scan_bzip2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_offloaded": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stream_based_uncompressed_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"streaming_content_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strip_x_forwarded_for": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switching_protocols": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_window_maximum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_window_minimum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_window_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_window_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_non_http": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uncompressed_nest_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"uncompressed_oversize_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"unknown_content_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_http_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verify_dns_for_policy_matching": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallProfileProtocolOptionsHttpUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["profile_protocol_options"] = profile_protocol_options

	obj, err := getObjectObjectFirewallProfileProtocolOptionsHttp(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptionsHttp resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallProfileProtocolOptionsHttp(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptionsHttp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFirewallProfileProtocolOptionsHttp")

	return resourceObjectFirewallProfileProtocolOptionsHttpRead(d, m)
}

func resourceObjectFirewallProfileProtocolOptionsHttpDelete(d *schema.ResourceData, m interface{}) error {
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

	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["profile_protocol_options"] = profile_protocol_options

	err = c.DeleteObjectFirewallProfileProtocolOptionsHttp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallProfileProtocolOptionsHttp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallProfileProtocolOptionsHttpRead(d *schema.ResourceData, m interface{}) error {
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

	profile_protocol_options := d.Get("profile_protocol_options").(string)
	if profile_protocol_options == "" {
		profile_protocol_options = importOptionChecking(m.(*FortiClient).Cfg, "profile_protocol_options")
		if profile_protocol_options == "" {
			return fmt.Errorf("Parameter profile_protocol_options is missing")
		}
		if err = d.Set("profile_protocol_options", profile_protocol_options); err != nil {
			return fmt.Errorf("Error set params profile_protocol_options: %v", err)
		}
	}
	paradict["profile_protocol_options"] = profile_protocol_options

	o, err := c.ReadObjectFirewallProfileProtocolOptionsHttp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallProfileProtocolOptionsHttp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallProfileProtocolOptionsHttp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallProfileProtocolOptionsHttp resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallProfileProtocolOptionsHttpAddressIpRating2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpBlockPageStatusCode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpComfortAmount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpComfortInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpDomainFronting2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpH2C2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpFortinetBar2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpFortinetBarPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpHttpPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpInspectAll2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsHttpOversizeLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpPorts2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallProfileProtocolOptionsHttpPostLang2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsHttpProxyAfterTcpHandshake2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpRangeBlock2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpRetryCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpScanBzip22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpSslOffloaded2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpStreamBasedUncompressedLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpStreamingContentBypass2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpStripXForwardedFor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpSwitchingProtocols2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowMaximum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowMinimum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpTunnelNonHttp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpUncompressedNestLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpUnknownContentEncoding2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpUnknownHttpVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpVerifyDnsForPolicyMatching2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallProfileProtocolOptionsHttp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("address_ip_rating", flattenObjectFirewallProfileProtocolOptionsHttpAddressIpRating2edl(o["address-ip-rating"], d, "address_ip_rating")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-ip-rating"], "ObjectFirewallProfileProtocolOptionsHttp-AddressIpRating"); ok {
			if err = d.Set("address_ip_rating", vv); err != nil {
				return fmt.Errorf("Error reading address_ip_rating: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_ip_rating: %v", err)
		}
	}

	if err = d.Set("block_page_status_code", flattenObjectFirewallProfileProtocolOptionsHttpBlockPageStatusCode2edl(o["block-page-status-code"], d, "block_page_status_code")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-page-status-code"], "ObjectFirewallProfileProtocolOptionsHttp-BlockPageStatusCode"); ok {
			if err = d.Set("block_page_status_code", vv); err != nil {
				return fmt.Errorf("Error reading block_page_status_code: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_page_status_code: %v", err)
		}
	}

	if err = d.Set("comfort_amount", flattenObjectFirewallProfileProtocolOptionsHttpComfortAmount2edl(o["comfort-amount"], d, "comfort_amount")); err != nil {
		if vv, ok := fortiAPIPatch(o["comfort-amount"], "ObjectFirewallProfileProtocolOptionsHttp-ComfortAmount"); ok {
			if err = d.Set("comfort_amount", vv); err != nil {
				return fmt.Errorf("Error reading comfort_amount: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comfort_amount: %v", err)
		}
	}

	if err = d.Set("comfort_interval", flattenObjectFirewallProfileProtocolOptionsHttpComfortInterval2edl(o["comfort-interval"], d, "comfort_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["comfort-interval"], "ObjectFirewallProfileProtocolOptionsHttp-ComfortInterval"); ok {
			if err = d.Set("comfort_interval", vv); err != nil {
				return fmt.Errorf("Error reading comfort_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comfort_interval: %v", err)
		}
	}

	if err = d.Set("domain_fronting", flattenObjectFirewallProfileProtocolOptionsHttpDomainFronting2edl(o["domain-fronting"], d, "domain_fronting")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-fronting"], "ObjectFirewallProfileProtocolOptionsHttp-DomainFronting"); ok {
			if err = d.Set("domain_fronting", vv); err != nil {
				return fmt.Errorf("Error reading domain_fronting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_fronting: %v", err)
		}
	}

	if err = d.Set("h2c", flattenObjectFirewallProfileProtocolOptionsHttpH2C2edl(o["h2c"], d, "h2c")); err != nil {
		if vv, ok := fortiAPIPatch(o["h2c"], "ObjectFirewallProfileProtocolOptionsHttp-H2C"); ok {
			if err = d.Set("h2c", vv); err != nil {
				return fmt.Errorf("Error reading h2c: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h2c: %v", err)
		}
	}

	if err = d.Set("fortinet_bar", flattenObjectFirewallProfileProtocolOptionsHttpFortinetBar2edl(o["fortinet-bar"], d, "fortinet_bar")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortinet-bar"], "ObjectFirewallProfileProtocolOptionsHttp-FortinetBar"); ok {
			if err = d.Set("fortinet_bar", vv); err != nil {
				return fmt.Errorf("Error reading fortinet_bar: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortinet_bar: %v", err)
		}
	}

	if err = d.Set("fortinet_bar_port", flattenObjectFirewallProfileProtocolOptionsHttpFortinetBarPort2edl(o["fortinet-bar-port"], d, "fortinet_bar_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortinet-bar-port"], "ObjectFirewallProfileProtocolOptionsHttp-FortinetBarPort"); ok {
			if err = d.Set("fortinet_bar_port", vv); err != nil {
				return fmt.Errorf("Error reading fortinet_bar_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortinet_bar_port: %v", err)
		}
	}

	if err = d.Set("http_policy", flattenObjectFirewallProfileProtocolOptionsHttpHttpPolicy2edl(o["http-policy"], d, "http_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-policy"], "ObjectFirewallProfileProtocolOptionsHttp-HttpPolicy"); ok {
			if err = d.Set("http_policy", vv); err != nil {
				return fmt.Errorf("Error reading http_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_policy: %v", err)
		}
	}

	if err = d.Set("inspect_all", flattenObjectFirewallProfileProtocolOptionsHttpInspectAll2edl(o["inspect-all"], d, "inspect_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["inspect-all"], "ObjectFirewallProfileProtocolOptionsHttp-InspectAll"); ok {
			if err = d.Set("inspect_all", vv); err != nil {
				return fmt.Errorf("Error reading inspect_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inspect_all: %v", err)
		}
	}

	if err = d.Set("options", flattenObjectFirewallProfileProtocolOptionsHttpOptions2edl(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "ObjectFirewallProfileProtocolOptionsHttp-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("oversize_limit", flattenObjectFirewallProfileProtocolOptionsHttpOversizeLimit2edl(o["oversize-limit"], d, "oversize_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["oversize-limit"], "ObjectFirewallProfileProtocolOptionsHttp-OversizeLimit"); ok {
			if err = d.Set("oversize_limit", vv); err != nil {
				return fmt.Errorf("Error reading oversize_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oversize_limit: %v", err)
		}
	}

	if err = d.Set("ports", flattenObjectFirewallProfileProtocolOptionsHttpPorts2edl(o["ports"], d, "ports")); err != nil {
		if vv, ok := fortiAPIPatch(o["ports"], "ObjectFirewallProfileProtocolOptionsHttp-Ports"); ok {
			if err = d.Set("ports", vv); err != nil {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ports: %v", err)
		}
	}

	if err = d.Set("post_lang", flattenObjectFirewallProfileProtocolOptionsHttpPostLang2edl(o["post-lang"], d, "post_lang")); err != nil {
		if vv, ok := fortiAPIPatch(o["post-lang"], "ObjectFirewallProfileProtocolOptionsHttp-PostLang"); ok {
			if err = d.Set("post_lang", vv); err != nil {
				return fmt.Errorf("Error reading post_lang: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading post_lang: %v", err)
		}
	}

	if err = d.Set("proxy_after_tcp_handshake", flattenObjectFirewallProfileProtocolOptionsHttpProxyAfterTcpHandshake2edl(o["proxy-after-tcp-handshake"], d, "proxy_after_tcp_handshake")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-after-tcp-handshake"], "ObjectFirewallProfileProtocolOptionsHttp-ProxyAfterTcpHandshake"); ok {
			if err = d.Set("proxy_after_tcp_handshake", vv); err != nil {
				return fmt.Errorf("Error reading proxy_after_tcp_handshake: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_after_tcp_handshake: %v", err)
		}
	}

	if err = d.Set("range_block", flattenObjectFirewallProfileProtocolOptionsHttpRangeBlock2edl(o["range-block"], d, "range_block")); err != nil {
		if vv, ok := fortiAPIPatch(o["range-block"], "ObjectFirewallProfileProtocolOptionsHttp-RangeBlock"); ok {
			if err = d.Set("range_block", vv); err != nil {
				return fmt.Errorf("Error reading range_block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading range_block: %v", err)
		}
	}

	if err = d.Set("retry_count", flattenObjectFirewallProfileProtocolOptionsHttpRetryCount2edl(o["retry-count"], d, "retry_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["retry-count"], "ObjectFirewallProfileProtocolOptionsHttp-RetryCount"); ok {
			if err = d.Set("retry_count", vv); err != nil {
				return fmt.Errorf("Error reading retry_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retry_count: %v", err)
		}
	}

	if err = d.Set("scan_bzip2", flattenObjectFirewallProfileProtocolOptionsHttpScanBzip22edl(o["scan-bzip2"], d, "scan_bzip2")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-bzip2"], "ObjectFirewallProfileProtocolOptionsHttp-ScanBzip2"); ok {
			if err = d.Set("scan_bzip2", vv); err != nil {
				return fmt.Errorf("Error reading scan_bzip2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_bzip2: %v", err)
		}
	}

	if err = d.Set("ssl_offloaded", flattenObjectFirewallProfileProtocolOptionsHttpSslOffloaded2edl(o["ssl-offloaded"], d, "ssl_offloaded")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-offloaded"], "ObjectFirewallProfileProtocolOptionsHttp-SslOffloaded"); ok {
			if err = d.Set("ssl_offloaded", vv); err != nil {
				return fmt.Errorf("Error reading ssl_offloaded: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_offloaded: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectFirewallProfileProtocolOptionsHttpStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectFirewallProfileProtocolOptionsHttp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("stream_based_uncompressed_limit", flattenObjectFirewallProfileProtocolOptionsHttpStreamBasedUncompressedLimit2edl(o["stream-based-uncompressed-limit"], d, "stream_based_uncompressed_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["stream-based-uncompressed-limit"], "ObjectFirewallProfileProtocolOptionsHttp-StreamBasedUncompressedLimit"); ok {
			if err = d.Set("stream_based_uncompressed_limit", vv); err != nil {
				return fmt.Errorf("Error reading stream_based_uncompressed_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stream_based_uncompressed_limit: %v", err)
		}
	}

	if err = d.Set("streaming_content_bypass", flattenObjectFirewallProfileProtocolOptionsHttpStreamingContentBypass2edl(o["streaming-content-bypass"], d, "streaming_content_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["streaming-content-bypass"], "ObjectFirewallProfileProtocolOptionsHttp-StreamingContentBypass"); ok {
			if err = d.Set("streaming_content_bypass", vv); err != nil {
				return fmt.Errorf("Error reading streaming_content_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading streaming_content_bypass: %v", err)
		}
	}

	if err = d.Set("strip_x_forwarded_for", flattenObjectFirewallProfileProtocolOptionsHttpStripXForwardedFor2edl(o["strip-x-forwarded-for"], d, "strip_x_forwarded_for")); err != nil {
		if vv, ok := fortiAPIPatch(o["strip-x-forwarded-for"], "ObjectFirewallProfileProtocolOptionsHttp-StripXForwardedFor"); ok {
			if err = d.Set("strip_x_forwarded_for", vv); err != nil {
				return fmt.Errorf("Error reading strip_x_forwarded_for: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strip_x_forwarded_for: %v", err)
		}
	}

	if err = d.Set("switching_protocols", flattenObjectFirewallProfileProtocolOptionsHttpSwitchingProtocols2edl(o["switching-protocols"], d, "switching_protocols")); err != nil {
		if vv, ok := fortiAPIPatch(o["switching-protocols"], "ObjectFirewallProfileProtocolOptionsHttp-SwitchingProtocols"); ok {
			if err = d.Set("switching_protocols", vv); err != nil {
				return fmt.Errorf("Error reading switching_protocols: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switching_protocols: %v", err)
		}
	}

	if err = d.Set("tcp_window_maximum", flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowMaximum2edl(o["tcp-window-maximum"], d, "tcp_window_maximum")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-maximum"], "ObjectFirewallProfileProtocolOptionsHttp-TcpWindowMaximum"); ok {
			if err = d.Set("tcp_window_maximum", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_maximum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_maximum: %v", err)
		}
	}

	if err = d.Set("tcp_window_minimum", flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowMinimum2edl(o["tcp-window-minimum"], d, "tcp_window_minimum")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-minimum"], "ObjectFirewallProfileProtocolOptionsHttp-TcpWindowMinimum"); ok {
			if err = d.Set("tcp_window_minimum", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_minimum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_minimum: %v", err)
		}
	}

	if err = d.Set("tcp_window_size", flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowSize2edl(o["tcp-window-size"], d, "tcp_window_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-size"], "ObjectFirewallProfileProtocolOptionsHttp-TcpWindowSize"); ok {
			if err = d.Set("tcp_window_size", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_size: %v", err)
		}
	}

	if err = d.Set("tcp_window_type", flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowType2edl(o["tcp-window-type"], d, "tcp_window_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-type"], "ObjectFirewallProfileProtocolOptionsHttp-TcpWindowType"); ok {
			if err = d.Set("tcp_window_type", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_type: %v", err)
		}
	}

	if err = d.Set("tunnel_non_http", flattenObjectFirewallProfileProtocolOptionsHttpTunnelNonHttp2edl(o["tunnel-non-http"], d, "tunnel_non_http")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-non-http"], "ObjectFirewallProfileProtocolOptionsHttp-TunnelNonHttp"); ok {
			if err = d.Set("tunnel_non_http", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_non_http: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_non_http: %v", err)
		}
	}

	if err = d.Set("uncompressed_nest_limit", flattenObjectFirewallProfileProtocolOptionsHttpUncompressedNestLimit2edl(o["uncompressed-nest-limit"], d, "uncompressed_nest_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["uncompressed-nest-limit"], "ObjectFirewallProfileProtocolOptionsHttp-UncompressedNestLimit"); ok {
			if err = d.Set("uncompressed_nest_limit", vv); err != nil {
				return fmt.Errorf("Error reading uncompressed_nest_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uncompressed_nest_limit: %v", err)
		}
	}

	if err = d.Set("uncompressed_oversize_limit", flattenObjectFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit2edl(o["uncompressed-oversize-limit"], d, "uncompressed_oversize_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["uncompressed-oversize-limit"], "ObjectFirewallProfileProtocolOptionsHttp-UncompressedOversizeLimit"); ok {
			if err = d.Set("uncompressed_oversize_limit", vv); err != nil {
				return fmt.Errorf("Error reading uncompressed_oversize_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uncompressed_oversize_limit: %v", err)
		}
	}

	if err = d.Set("unknown_content_encoding", flattenObjectFirewallProfileProtocolOptionsHttpUnknownContentEncoding2edl(o["unknown-content-encoding"], d, "unknown_content_encoding")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-content-encoding"], "ObjectFirewallProfileProtocolOptionsHttp-UnknownContentEncoding"); ok {
			if err = d.Set("unknown_content_encoding", vv); err != nil {
				return fmt.Errorf("Error reading unknown_content_encoding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_content_encoding: %v", err)
		}
	}

	if err = d.Set("unknown_http_version", flattenObjectFirewallProfileProtocolOptionsHttpUnknownHttpVersion2edl(o["unknown-http-version"], d, "unknown_http_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-http-version"], "ObjectFirewallProfileProtocolOptionsHttp-UnknownHttpVersion"); ok {
			if err = d.Set("unknown_http_version", vv); err != nil {
				return fmt.Errorf("Error reading unknown_http_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_http_version: %v", err)
		}
	}

	if err = d.Set("verify_dns_for_policy_matching", flattenObjectFirewallProfileProtocolOptionsHttpVerifyDnsForPolicyMatching2edl(o["verify-dns-for-policy-matching"], d, "verify_dns_for_policy_matching")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-dns-for-policy-matching"], "ObjectFirewallProfileProtocolOptionsHttp-VerifyDnsForPolicyMatching"); ok {
			if err = d.Set("verify_dns_for_policy_matching", vv); err != nil {
				return fmt.Errorf("Error reading verify_dns_for_policy_matching: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_dns_for_policy_matching: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallProfileProtocolOptionsHttpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallProfileProtocolOptionsHttpAddressIpRating2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpBlockPageStatusCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpComfortAmount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpComfortInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpDomainFronting2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpH2C2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpFortinetBar2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpFortinetBarPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpHttpPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpInspectAll2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsHttpOversizeLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpPorts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsHttpPostLang2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsHttpProxyAfterTcpHandshake2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpRangeBlock2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpRetryCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpScanBzip22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpSslOffloaded2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpStreamBasedUncompressedLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpStreamingContentBypass2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpStripXForwardedFor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpSwitchingProtocols2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpTcpWindowMaximum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpTcpWindowMinimum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpTcpWindowSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpTcpWindowType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpTunnelNonHttp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpUncompressedNestLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpUnknownContentEncoding2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpUnknownHttpVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpVerifyDnsForPolicyMatching2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallProfileProtocolOptionsHttp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address_ip_rating"); ok || d.HasChange("address_ip_rating") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpAddressIpRating2edl(d, v, "address_ip_rating")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-ip-rating"] = t
		}
	}

	if v, ok := d.GetOk("block_page_status_code"); ok || d.HasChange("block_page_status_code") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpBlockPageStatusCode2edl(d, v, "block_page_status_code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-page-status-code"] = t
		}
	}

	if v, ok := d.GetOk("comfort_amount"); ok || d.HasChange("comfort_amount") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpComfortAmount2edl(d, v, "comfort_amount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comfort-amount"] = t
		}
	}

	if v, ok := d.GetOk("comfort_interval"); ok || d.HasChange("comfort_interval") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpComfortInterval2edl(d, v, "comfort_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comfort-interval"] = t
		}
	}

	if v, ok := d.GetOk("domain_fronting"); ok || d.HasChange("domain_fronting") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpDomainFronting2edl(d, v, "domain_fronting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-fronting"] = t
		}
	}

	if v, ok := d.GetOk("h2c"); ok || d.HasChange("h2c") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpH2C2edl(d, v, "h2c")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2c"] = t
		}
	}

	if v, ok := d.GetOk("fortinet_bar"); ok || d.HasChange("fortinet_bar") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpFortinetBar2edl(d, v, "fortinet_bar")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortinet-bar"] = t
		}
	}

	if v, ok := d.GetOk("fortinet_bar_port"); ok || d.HasChange("fortinet_bar_port") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpFortinetBarPort2edl(d, v, "fortinet_bar_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortinet-bar-port"] = t
		}
	}

	if v, ok := d.GetOk("http_policy"); ok || d.HasChange("http_policy") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpHttpPolicy2edl(d, v, "http_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-policy"] = t
		}
	}

	if v, ok := d.GetOk("inspect_all"); ok || d.HasChange("inspect_all") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpInspectAll2edl(d, v, "inspect_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspect-all"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("oversize_limit"); ok || d.HasChange("oversize_limit") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpOversizeLimit2edl(d, v, "oversize_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oversize-limit"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpPorts2edl(d, v, "ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("post_lang"); ok || d.HasChange("post_lang") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpPostLang2edl(d, v, "post_lang")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["post-lang"] = t
		}
	}

	if v, ok := d.GetOk("proxy_after_tcp_handshake"); ok || d.HasChange("proxy_after_tcp_handshake") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpProxyAfterTcpHandshake2edl(d, v, "proxy_after_tcp_handshake")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-after-tcp-handshake"] = t
		}
	}

	if v, ok := d.GetOk("range_block"); ok || d.HasChange("range_block") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpRangeBlock2edl(d, v, "range_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range-block"] = t
		}
	}

	if v, ok := d.GetOk("retry_count"); ok || d.HasChange("retry_count") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpRetryCount2edl(d, v, "retry_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retry-count"] = t
		}
	}

	if v, ok := d.GetOk("scan_bzip2"); ok || d.HasChange("scan_bzip2") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpScanBzip22edl(d, v, "scan_bzip2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-bzip2"] = t
		}
	}

	if v, ok := d.GetOk("ssl_offloaded"); ok || d.HasChange("ssl_offloaded") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpSslOffloaded2edl(d, v, "ssl_offloaded")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-offloaded"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("stream_based_uncompressed_limit"); ok || d.HasChange("stream_based_uncompressed_limit") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpStreamBasedUncompressedLimit2edl(d, v, "stream_based_uncompressed_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stream-based-uncompressed-limit"] = t
		}
	}

	if v, ok := d.GetOk("streaming_content_bypass"); ok || d.HasChange("streaming_content_bypass") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpStreamingContentBypass2edl(d, v, "streaming_content_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["streaming-content-bypass"] = t
		}
	}

	if v, ok := d.GetOk("strip_x_forwarded_for"); ok || d.HasChange("strip_x_forwarded_for") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpStripXForwardedFor2edl(d, v, "strip_x_forwarded_for")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strip-x-forwarded-for"] = t
		}
	}

	if v, ok := d.GetOk("switching_protocols"); ok || d.HasChange("switching_protocols") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpSwitchingProtocols2edl(d, v, "switching_protocols")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switching-protocols"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_maximum"); ok || d.HasChange("tcp_window_maximum") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpTcpWindowMaximum2edl(d, v, "tcp_window_maximum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-maximum"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_minimum"); ok || d.HasChange("tcp_window_minimum") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpTcpWindowMinimum2edl(d, v, "tcp_window_minimum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-minimum"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_size"); ok || d.HasChange("tcp_window_size") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpTcpWindowSize2edl(d, v, "tcp_window_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-size"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_type"); ok || d.HasChange("tcp_window_type") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpTcpWindowType2edl(d, v, "tcp_window_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-type"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_non_http"); ok || d.HasChange("tunnel_non_http") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpTunnelNonHttp2edl(d, v, "tunnel_non_http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-non-http"] = t
		}
	}

	if v, ok := d.GetOk("uncompressed_nest_limit"); ok || d.HasChange("uncompressed_nest_limit") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpUncompressedNestLimit2edl(d, v, "uncompressed_nest_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncompressed-nest-limit"] = t
		}
	}

	if v, ok := d.GetOk("uncompressed_oversize_limit"); ok || d.HasChange("uncompressed_oversize_limit") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit2edl(d, v, "uncompressed_oversize_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncompressed-oversize-limit"] = t
		}
	}

	if v, ok := d.GetOk("unknown_content_encoding"); ok || d.HasChange("unknown_content_encoding") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpUnknownContentEncoding2edl(d, v, "unknown_content_encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-content-encoding"] = t
		}
	}

	if v, ok := d.GetOk("unknown_http_version"); ok || d.HasChange("unknown_http_version") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpUnknownHttpVersion2edl(d, v, "unknown_http_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-http-version"] = t
		}
	}

	if v, ok := d.GetOk("verify_dns_for_policy_matching"); ok || d.HasChange("verify_dns_for_policy_matching") {
		t, err := expandObjectFirewallProfileProtocolOptionsHttpVerifyDnsForPolicyMatching2edl(d, v, "verify_dns_for_policy_matching")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-dns-for-policy-matching"] = t
		}
	}

	return &obj, nil
}
