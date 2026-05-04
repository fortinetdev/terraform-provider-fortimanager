// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DNS.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystempSystemDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystempSystemDnsUpdate,
		Read:   resourceSystempSystemDnsRead,
		Update: resourceSystempSystemDnsUpdate,
		Delete: resourceSystempSystemDnsDelete,

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
			"devprof": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"alt_primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alt_secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cache_notfound_responses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_cache_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dns_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dns_over_tls": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fqdn_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fqdn_max_refresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fqdn_min_refresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hostname_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hostname_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"retry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"root_servers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_hostname": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystempSystemDnsUpdate(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	obj, err := getObjectSystempSystemDns(d, false)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemDns resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystempSystemDns(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemDns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystempSystemDns")

	return resourceSystempSystemDnsRead(d, m)
}

func resourceSystempSystemDnsDelete(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	obj, err := getObjectSystempSystemDns(d, true)

	if err != nil {
		return fmt.Errorf("Error updating SystempSystemDns resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystempSystemDns(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing SystempSystemDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystempSystemDnsRead(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	if devprof == "" {
		devprof = importOptionChecking(m.(*FortiClient).Cfg, "devprof")
		if devprof == "" {
			return fmt.Errorf("Parameter devprof is missing")
		}
		if err = d.Set("devprof", devprof); err != nil {
			return fmt.Errorf("Error set params devprof: %v", err)
		}
	}
	paradict["devprof"] = devprof

	o, err := c.ReadSystempSystemDns(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystempSystemDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystempSystemDns(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemDns resource from API: %v", err)
	}
	return nil
}

func flattenSystempSystemDnsAltPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsAltSecondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsCacheNotfoundResponses(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsDnsCacheLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsDnsCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsDnsOverTls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempSystemDnsFqdnCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsFqdnMaxRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsFqdnMinRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsHostnameLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsHostnameTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempSystemDnsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsIp6Primary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsIp6Secondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempSystemDnsRetry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsRootServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsSecondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsServerHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempSystemDnsServerSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsSourceIpInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempSystemDnsSslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempSystemDnsTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemDnsVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystempSystemDns(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("alt_primary", flattenSystempSystemDnsAltPrimary(o["alt-primary"], d, "alt_primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["alt-primary"], "SystempSystemDns-AltPrimary"); ok {
			if err = d.Set("alt_primary", vv); err != nil {
				return fmt.Errorf("Error reading alt_primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alt_primary: %v", err)
		}
	}

	if err = d.Set("alt_secondary", flattenSystempSystemDnsAltSecondary(o["alt-secondary"], d, "alt_secondary")); err != nil {
		if vv, ok := fortiAPIPatch(o["alt-secondary"], "SystempSystemDns-AltSecondary"); ok {
			if err = d.Set("alt_secondary", vv); err != nil {
				return fmt.Errorf("Error reading alt_secondary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alt_secondary: %v", err)
		}
	}

	if err = d.Set("cache_notfound_responses", flattenSystempSystemDnsCacheNotfoundResponses(o["cache-notfound-responses"], d, "cache_notfound_responses")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-notfound-responses"], "SystempSystemDns-CacheNotfoundResponses"); ok {
			if err = d.Set("cache_notfound_responses", vv); err != nil {
				return fmt.Errorf("Error reading cache_notfound_responses: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_notfound_responses: %v", err)
		}
	}

	if err = d.Set("dns_cache_limit", flattenSystempSystemDnsDnsCacheLimit(o["dns-cache-limit"], d, "dns_cache_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-cache-limit"], "SystempSystemDns-DnsCacheLimit"); ok {
			if err = d.Set("dns_cache_limit", vv); err != nil {
				return fmt.Errorf("Error reading dns_cache_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_cache_limit: %v", err)
		}
	}

	if err = d.Set("dns_cache_ttl", flattenSystempSystemDnsDnsCacheTtl(o["dns-cache-ttl"], d, "dns_cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-cache-ttl"], "SystempSystemDns-DnsCacheTtl"); ok {
			if err = d.Set("dns_cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading dns_cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_cache_ttl: %v", err)
		}
	}

	if err = d.Set("dns_over_tls", flattenSystempSystemDnsDnsOverTls(o["dns-over-tls"], d, "dns_over_tls")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-over-tls"], "SystempSystemDns-DnsOverTls"); ok {
			if err = d.Set("dns_over_tls", vv); err != nil {
				return fmt.Errorf("Error reading dns_over_tls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_over_tls: %v", err)
		}
	}

	if err = d.Set("domain", flattenSystempSystemDnsDomain(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "SystempSystemDns-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("fqdn_cache_ttl", flattenSystempSystemDnsFqdnCacheTtl(o["fqdn-cache-ttl"], d, "fqdn_cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn-cache-ttl"], "SystempSystemDns-FqdnCacheTtl"); ok {
			if err = d.Set("fqdn_cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading fqdn_cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn_cache_ttl: %v", err)
		}
	}

	if err = d.Set("fqdn_max_refresh", flattenSystempSystemDnsFqdnMaxRefresh(o["fqdn-max-refresh"], d, "fqdn_max_refresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn-max-refresh"], "SystempSystemDns-FqdnMaxRefresh"); ok {
			if err = d.Set("fqdn_max_refresh", vv); err != nil {
				return fmt.Errorf("Error reading fqdn_max_refresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn_max_refresh: %v", err)
		}
	}

	if err = d.Set("fqdn_min_refresh", flattenSystempSystemDnsFqdnMinRefresh(o["fqdn-min-refresh"], d, "fqdn_min_refresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn-min-refresh"], "SystempSystemDns-FqdnMinRefresh"); ok {
			if err = d.Set("fqdn_min_refresh", vv); err != nil {
				return fmt.Errorf("Error reading fqdn_min_refresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn_min_refresh: %v", err)
		}
	}

	if err = d.Set("hostname_limit", flattenSystempSystemDnsHostnameLimit(o["hostname-limit"], d, "hostname_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname-limit"], "SystempSystemDns-HostnameLimit"); ok {
			if err = d.Set("hostname_limit", vv); err != nil {
				return fmt.Errorf("Error reading hostname_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname_limit: %v", err)
		}
	}

	if err = d.Set("hostname_ttl", flattenSystempSystemDnsHostnameTtl(o["hostname-ttl"], d, "hostname_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname-ttl"], "SystempSystemDns-HostnameTtl"); ok {
			if err = d.Set("hostname_ttl", vv); err != nil {
				return fmt.Errorf("Error reading hostname_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname_ttl: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystempSystemDnsInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystempSystemDns-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystempSystemDnsInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "SystempSystemDns-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ip6_primary", flattenSystempSystemDnsIp6Primary(o["ip6-primary"], d, "ip6_primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-primary"], "SystempSystemDns-Ip6Primary"); ok {
			if err = d.Set("ip6_primary", vv); err != nil {
				return fmt.Errorf("Error reading ip6_primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_primary: %v", err)
		}
	}

	if err = d.Set("ip6_secondary", flattenSystempSystemDnsIp6Secondary(o["ip6-secondary"], d, "ip6_secondary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-secondary"], "SystempSystemDns-Ip6Secondary"); ok {
			if err = d.Set("ip6_secondary", vv); err != nil {
				return fmt.Errorf("Error reading ip6_secondary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_secondary: %v", err)
		}
	}

	if err = d.Set("log", flattenSystempSystemDnsLog(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "SystempSystemDns-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("primary", flattenSystempSystemDnsPrimary(o["primary"], d, "primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary"], "SystempSystemDns-Primary"); ok {
			if err = d.Set("primary", vv); err != nil {
				return fmt.Errorf("Error reading primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystempSystemDnsProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystempSystemDns-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("retry", flattenSystempSystemDnsRetry(o["retry"], d, "retry")); err != nil {
		if vv, ok := fortiAPIPatch(o["retry"], "SystempSystemDns-Retry"); ok {
			if err = d.Set("retry", vv); err != nil {
				return fmt.Errorf("Error reading retry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retry: %v", err)
		}
	}

	if err = d.Set("root_servers", flattenSystempSystemDnsRootServers(o["root-servers"], d, "root_servers")); err != nil {
		if vv, ok := fortiAPIPatch(o["root-servers"], "SystempSystemDns-RootServers"); ok {
			if err = d.Set("root_servers", vv); err != nil {
				return fmt.Errorf("Error reading root_servers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading root_servers: %v", err)
		}
	}

	if err = d.Set("secondary", flattenSystempSystemDnsSecondary(o["secondary"], d, "secondary")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary"], "SystempSystemDns-Secondary"); ok {
			if err = d.Set("secondary", vv); err != nil {
				return fmt.Errorf("Error reading secondary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary: %v", err)
		}
	}

	if err = d.Set("server_hostname", flattenSystempSystemDnsServerHostname(o["server-hostname"], d, "server_hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-hostname"], "SystempSystemDns-ServerHostname"); ok {
			if err = d.Set("server_hostname", vv); err != nil {
				return fmt.Errorf("Error reading server_hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_hostname: %v", err)
		}
	}

	if err = d.Set("server_select_method", flattenSystempSystemDnsServerSelectMethod(o["server-select-method"], d, "server_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-select-method"], "SystempSystemDns-ServerSelectMethod"); ok {
			if err = d.Set("server_select_method", vv); err != nil {
				return fmt.Errorf("Error reading server_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_select_method: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystempSystemDnsSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "SystempSystemDns-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip_interface", flattenSystempSystemDnsSourceIpInterface(o["source-ip-interface"], d, "source_ip_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-interface"], "SystempSystemDns-SourceIpInterface"); ok {
			if err = d.Set("source_ip_interface", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_interface: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenSystempSystemDnsSslCertificate(o["ssl-certificate"], d, "ssl_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-certificate"], "SystempSystemDns-SslCertificate"); ok {
			if err = d.Set("ssl_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if err = d.Set("timeout", flattenSystempSystemDnsTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "SystempSystemDns-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenSystempSystemDnsVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "SystempSystemDns-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenSystempSystemDnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystempSystemDnsAltPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsAltSecondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsCacheNotfoundResponses(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsDnsCacheLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsDnsCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsDnsOverTls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemDnsFqdnCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsFqdnMaxRefresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsFqdnMinRefresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsHostnameLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsHostnameTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemDnsInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsIp6Primary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsIp6Secondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemDnsRetry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsRootServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsSecondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsServerHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemDnsServerSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsSourceIpInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemDnsSslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemDnsTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemDnsVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystempSystemDns(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("alt_primary"); ok || d.HasChange("alt_primary") {
		t, err := expandSystempSystemDnsAltPrimary(d, v, "alt_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alt-primary"] = t
		}
	}

	if v, ok := d.GetOk("alt_secondary"); ok || d.HasChange("alt_secondary") {
		t, err := expandSystempSystemDnsAltSecondary(d, v, "alt_secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alt-secondary"] = t
		}
	}

	if v, ok := d.GetOk("cache_notfound_responses"); ok || d.HasChange("cache_notfound_responses") {
		t, err := expandSystempSystemDnsCacheNotfoundResponses(d, v, "cache_notfound_responses")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-notfound-responses"] = t
		}
	}

	if v, ok := d.GetOk("dns_cache_limit"); ok || d.HasChange("dns_cache_limit") {
		t, err := expandSystempSystemDnsDnsCacheLimit(d, v, "dns_cache_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-cache-limit"] = t
		}
	}

	if v, ok := d.GetOk("dns_cache_ttl"); ok || d.HasChange("dns_cache_ttl") {
		t, err := expandSystempSystemDnsDnsCacheTtl(d, v, "dns_cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("dns_over_tls"); ok || d.HasChange("dns_over_tls") {
		t, err := expandSystempSystemDnsDnsOverTls(d, v, "dns_over_tls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-over-tls"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandSystempSystemDnsDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("fqdn_cache_ttl"); ok || d.HasChange("fqdn_cache_ttl") {
		t, err := expandSystempSystemDnsFqdnCacheTtl(d, v, "fqdn_cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn-cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("fqdn_max_refresh"); ok || d.HasChange("fqdn_max_refresh") {
		t, err := expandSystempSystemDnsFqdnMaxRefresh(d, v, "fqdn_max_refresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn-max-refresh"] = t
		}
	}

	if v, ok := d.GetOk("fqdn_min_refresh"); ok || d.HasChange("fqdn_min_refresh") {
		t, err := expandSystempSystemDnsFqdnMinRefresh(d, v, "fqdn_min_refresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn-min-refresh"] = t
		}
	}

	if v, ok := d.GetOk("hostname_limit"); ok || d.HasChange("hostname_limit") {
		t, err := expandSystempSystemDnsHostnameLimit(d, v, "hostname_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname-limit"] = t
		}
	}

	if v, ok := d.GetOk("hostname_ttl"); ok || d.HasChange("hostname_ttl") {
		t, err := expandSystempSystemDnsHostnameTtl(d, v, "hostname_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname-ttl"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystempSystemDnsInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandSystempSystemDnsInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ip6_primary"); ok || d.HasChange("ip6_primary") {
		t, err := expandSystempSystemDnsIp6Primary(d, v, "ip6_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-primary"] = t
		}
	}

	if v, ok := d.GetOk("ip6_secondary"); ok || d.HasChange("ip6_secondary") {
		t, err := expandSystempSystemDnsIp6Secondary(d, v, "ip6_secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-secondary"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandSystempSystemDnsLog(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("primary"); ok || d.HasChange("primary") {
		t, err := expandSystempSystemDnsPrimary(d, v, "primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystempSystemDnsProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("retry"); ok || d.HasChange("retry") {
		t, err := expandSystempSystemDnsRetry(d, v, "retry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retry"] = t
		}
	}

	if v, ok := d.GetOk("root_servers"); ok || d.HasChange("root_servers") {
		t, err := expandSystempSystemDnsRootServers(d, v, "root_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["root-servers"] = t
		}
	}

	if v, ok := d.GetOk("secondary"); ok || d.HasChange("secondary") {
		t, err := expandSystempSystemDnsSecondary(d, v, "secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary"] = t
		}
	}

	if v, ok := d.GetOk("server_hostname"); ok || d.HasChange("server_hostname") {
		t, err := expandSystempSystemDnsServerHostname(d, v, "server_hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-hostname"] = t
		}
	}

	if v, ok := d.GetOk("server_select_method"); ok || d.HasChange("server_select_method") {
		t, err := expandSystempSystemDnsServerSelectMethod(d, v, "server_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-select-method"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandSystempSystemDnsSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_interface"); ok || d.HasChange("source_ip_interface") {
		t, err := expandSystempSystemDnsSourceIpInterface(d, v, "source_ip_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-interface"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok || d.HasChange("ssl_certificate") {
		t, err := expandSystempSystemDnsSslCertificate(d, v, "ssl_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandSystempSystemDnsTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandSystempSystemDnsVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
