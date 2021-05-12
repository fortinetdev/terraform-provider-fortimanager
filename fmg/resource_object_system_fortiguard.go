// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard services.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSystemFortiguard() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemFortiguardUpdate,
		Read:   resourceObjectSystemFortiguardRead,
		Update: resourceObjectSystemFortiguardUpdate,
		Delete: resourceObjectSystemFortiguardDelete,

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
			"antispam_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"antispam_cache_mpercent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"antispam_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"antispam_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"antispam_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"antispam_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"antispam_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"anycast_sdns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anycast_sdns_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_join_forticloud": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fortiguard_anycast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiguard_anycast_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"load_balance_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_cache_mpercent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"proxy_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proxy_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sandbox_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdns_options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sdns_server_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sdns_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"service_account_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_server_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"webfilter_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"webfilter_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"webfilter_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemFortiguardUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemFortiguard(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemFortiguard resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemFortiguard(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemFortiguard resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemFortiguard")

	return resourceObjectSystemFortiguardRead(d, m)
}

func resourceObjectSystemFortiguardDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemFortiguard(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemFortiguard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemFortiguardRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemFortiguard(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemFortiguard resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemFortiguard(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemFortiguard resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemFortiguardAntispamCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemFortiguardAntispamCacheMpercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardAntispamCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardAntispamExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardAntispamForceOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemFortiguardAntispamLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardAntispamTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardAnycastSdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardAnycastSdnsServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardAutoJoinForticloud(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemFortiguardDdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardDdnsServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardFortiguardAnycast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemFortiguardFortiguardAnycastSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "fortinet",
			1: "aws",
			2: "debug",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemFortiguardInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "auto",
			1: "sdwan",
			2: "specify",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemFortiguardLoadBalanceServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardOutbreakPreventionCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemFortiguardOutbreakPreventionCacheMpercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardOutbreakPreventionCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardOutbreakPreventionExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardOutbreakPreventionForceOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemFortiguardOutbreakPreventionLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardOutbreakPreventionTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			53:   "53",
			80:   "80",
			443:  "443",
			8888: "8888",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemFortiguardProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "udp",
			1: "http",
			2: "https",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemFortiguardProxyPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemFortiguardProxyServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardProxyServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardProxyUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardSandboxRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardSdnsOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "include-question-section",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemFortiguardSdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemFortiguardSdnsServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardServiceAccountId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardUpdateServerLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "any",
			1: "usa",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemFortiguardWebfilterCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemFortiguardWebfilterCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardWebfilterExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardWebfilterForceOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemFortiguardWebfilterLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemFortiguardWebfilterTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemFortiguard(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("antispam_cache", flattenObjectSystemFortiguardAntispamCache(o["antispam-cache"], d, "antispam_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["antispam-cache"], "ObjectSystemFortiguard-AntispamCache"); ok {
			if err = d.Set("antispam_cache", vv); err != nil {
				return fmt.Errorf("Error reading antispam_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antispam_cache: %v", err)
		}
	}

	if err = d.Set("antispam_cache_mpercent", flattenObjectSystemFortiguardAntispamCacheMpercent(o["antispam-cache-mpercent"], d, "antispam_cache_mpercent")); err != nil {
		if vv, ok := fortiAPIPatch(o["antispam-cache-mpercent"], "ObjectSystemFortiguard-AntispamCacheMpercent"); ok {
			if err = d.Set("antispam_cache_mpercent", vv); err != nil {
				return fmt.Errorf("Error reading antispam_cache_mpercent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antispam_cache_mpercent: %v", err)
		}
	}

	if err = d.Set("antispam_cache_ttl", flattenObjectSystemFortiguardAntispamCacheTtl(o["antispam-cache-ttl"], d, "antispam_cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["antispam-cache-ttl"], "ObjectSystemFortiguard-AntispamCacheTtl"); ok {
			if err = d.Set("antispam_cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading antispam_cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antispam_cache_ttl: %v", err)
		}
	}

	if err = d.Set("antispam_expiration", flattenObjectSystemFortiguardAntispamExpiration(o["antispam-expiration"], d, "antispam_expiration")); err != nil {
		if vv, ok := fortiAPIPatch(o["antispam-expiration"], "ObjectSystemFortiguard-AntispamExpiration"); ok {
			if err = d.Set("antispam_expiration", vv); err != nil {
				return fmt.Errorf("Error reading antispam_expiration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antispam_expiration: %v", err)
		}
	}

	if err = d.Set("antispam_force_off", flattenObjectSystemFortiguardAntispamForceOff(o["antispam-force-off"], d, "antispam_force_off")); err != nil {
		if vv, ok := fortiAPIPatch(o["antispam-force-off"], "ObjectSystemFortiguard-AntispamForceOff"); ok {
			if err = d.Set("antispam_force_off", vv); err != nil {
				return fmt.Errorf("Error reading antispam_force_off: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antispam_force_off: %v", err)
		}
	}

	if err = d.Set("antispam_license", flattenObjectSystemFortiguardAntispamLicense(o["antispam-license"], d, "antispam_license")); err != nil {
		if vv, ok := fortiAPIPatch(o["antispam-license"], "ObjectSystemFortiguard-AntispamLicense"); ok {
			if err = d.Set("antispam_license", vv); err != nil {
				return fmt.Errorf("Error reading antispam_license: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antispam_license: %v", err)
		}
	}

	if err = d.Set("antispam_timeout", flattenObjectSystemFortiguardAntispamTimeout(o["antispam-timeout"], d, "antispam_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["antispam-timeout"], "ObjectSystemFortiguard-AntispamTimeout"); ok {
			if err = d.Set("antispam_timeout", vv); err != nil {
				return fmt.Errorf("Error reading antispam_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antispam_timeout: %v", err)
		}
	}

	if err = d.Set("anycast_sdns_server_ip", flattenObjectSystemFortiguardAnycastSdnsServerIp(o["anycast-sdns-server-ip"], d, "anycast_sdns_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["anycast-sdns-server-ip"], "ObjectSystemFortiguard-AnycastSdnsServerIp"); ok {
			if err = d.Set("anycast_sdns_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading anycast_sdns_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anycast_sdns_server_ip: %v", err)
		}
	}

	if err = d.Set("anycast_sdns_server_port", flattenObjectSystemFortiguardAnycastSdnsServerPort(o["anycast-sdns-server-port"], d, "anycast_sdns_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["anycast-sdns-server-port"], "ObjectSystemFortiguard-AnycastSdnsServerPort"); ok {
			if err = d.Set("anycast_sdns_server_port", vv); err != nil {
				return fmt.Errorf("Error reading anycast_sdns_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anycast_sdns_server_port: %v", err)
		}
	}

	if err = d.Set("auto_join_forticloud", flattenObjectSystemFortiguardAutoJoinForticloud(o["auto-join-forticloud"], d, "auto_join_forticloud")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-join-forticloud"], "ObjectSystemFortiguard-AutoJoinForticloud"); ok {
			if err = d.Set("auto_join_forticloud", vv); err != nil {
				return fmt.Errorf("Error reading auto_join_forticloud: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_join_forticloud: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", flattenObjectSystemFortiguardDdnsServerIp(o["ddns-server-ip"], d, "ddns_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-server-ip"], "ObjectSystemFortiguard-DdnsServerIp"); ok {
			if err = d.Set("ddns_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading ddns_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_server_port", flattenObjectSystemFortiguardDdnsServerPort(o["ddns-server-port"], d, "ddns_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-server-port"], "ObjectSystemFortiguard-DdnsServerPort"); ok {
			if err = d.Set("ddns_server_port", vv); err != nil {
				return fmt.Errorf("Error reading ddns_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_server_port: %v", err)
		}
	}

	if err = d.Set("fortiguard_anycast", flattenObjectSystemFortiguardFortiguardAnycast(o["fortiguard-anycast"], d, "fortiguard_anycast")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiguard-anycast"], "ObjectSystemFortiguard-FortiguardAnycast"); ok {
			if err = d.Set("fortiguard_anycast", vv); err != nil {
				return fmt.Errorf("Error reading fortiguard_anycast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiguard_anycast: %v", err)
		}
	}

	if err = d.Set("fortiguard_anycast_source", flattenObjectSystemFortiguardFortiguardAnycastSource(o["fortiguard-anycast-source"], d, "fortiguard_anycast_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiguard-anycast-source"], "ObjectSystemFortiguard-FortiguardAnycastSource"); ok {
			if err = d.Set("fortiguard_anycast_source", vv); err != nil {
				return fmt.Errorf("Error reading fortiguard_anycast_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiguard_anycast_source: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectSystemFortiguardInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectSystemFortiguard-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenObjectSystemFortiguardInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "ObjectSystemFortiguard-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("load_balance_servers", flattenObjectSystemFortiguardLoadBalanceServers(o["load-balance-servers"], d, "load_balance_servers")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balance-servers"], "ObjectSystemFortiguard-LoadBalanceServers"); ok {
			if err = d.Set("load_balance_servers", vv); err != nil {
				return fmt.Errorf("Error reading load_balance_servers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balance_servers: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache", flattenObjectSystemFortiguardOutbreakPreventionCache(o["outbreak-prevention-cache"], d, "outbreak_prevention_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-cache"], "ObjectSystemFortiguard-OutbreakPreventionCache"); ok {
			if err = d.Set("outbreak_prevention_cache", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_cache: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache_mpercent", flattenObjectSystemFortiguardOutbreakPreventionCacheMpercent(o["outbreak-prevention-cache-mpercent"], d, "outbreak_prevention_cache_mpercent")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-cache-mpercent"], "ObjectSystemFortiguard-OutbreakPreventionCacheMpercent"); ok {
			if err = d.Set("outbreak_prevention_cache_mpercent", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_cache_mpercent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_cache_mpercent: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache_ttl", flattenObjectSystemFortiguardOutbreakPreventionCacheTtl(o["outbreak-prevention-cache-ttl"], d, "outbreak_prevention_cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-cache-ttl"], "ObjectSystemFortiguard-OutbreakPreventionCacheTtl"); ok {
			if err = d.Set("outbreak_prevention_cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_cache_ttl: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_expiration", flattenObjectSystemFortiguardOutbreakPreventionExpiration(o["outbreak-prevention-expiration"], d, "outbreak_prevention_expiration")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-expiration"], "ObjectSystemFortiguard-OutbreakPreventionExpiration"); ok {
			if err = d.Set("outbreak_prevention_expiration", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_expiration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_expiration: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_force_off", flattenObjectSystemFortiguardOutbreakPreventionForceOff(o["outbreak-prevention-force-off"], d, "outbreak_prevention_force_off")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-force-off"], "ObjectSystemFortiguard-OutbreakPreventionForceOff"); ok {
			if err = d.Set("outbreak_prevention_force_off", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_force_off: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_force_off: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_license", flattenObjectSystemFortiguardOutbreakPreventionLicense(o["outbreak-prevention-license"], d, "outbreak_prevention_license")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-license"], "ObjectSystemFortiguard-OutbreakPreventionLicense"); ok {
			if err = d.Set("outbreak_prevention_license", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_license: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_license: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_timeout", flattenObjectSystemFortiguardOutbreakPreventionTimeout(o["outbreak-prevention-timeout"], d, "outbreak_prevention_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-timeout"], "ObjectSystemFortiguard-OutbreakPreventionTimeout"); ok {
			if err = d.Set("outbreak_prevention_timeout", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_timeout: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectSystemFortiguardPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectSystemFortiguard-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("protocol", flattenObjectSystemFortiguardProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ObjectSystemFortiguard-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("proxy_password", flattenObjectSystemFortiguardProxyPassword(o["proxy-password"], d, "proxy_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-password"], "ObjectSystemFortiguard-ProxyPassword"); ok {
			if err = d.Set("proxy_password", vv); err != nil {
				return fmt.Errorf("Error reading proxy_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_password: %v", err)
		}
	}

	if err = d.Set("proxy_server_ip", flattenObjectSystemFortiguardProxyServerIp(o["proxy-server-ip"], d, "proxy_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-server-ip"], "ObjectSystemFortiguard-ProxyServerIp"); ok {
			if err = d.Set("proxy_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading proxy_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_server_ip: %v", err)
		}
	}

	if err = d.Set("proxy_server_port", flattenObjectSystemFortiguardProxyServerPort(o["proxy-server-port"], d, "proxy_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-server-port"], "ObjectSystemFortiguard-ProxyServerPort"); ok {
			if err = d.Set("proxy_server_port", vv); err != nil {
				return fmt.Errorf("Error reading proxy_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_server_port: %v", err)
		}
	}

	if err = d.Set("proxy_username", flattenObjectSystemFortiguardProxyUsername(o["proxy-username"], d, "proxy_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-username"], "ObjectSystemFortiguard-ProxyUsername"); ok {
			if err = d.Set("proxy_username", vv); err != nil {
				return fmt.Errorf("Error reading proxy_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_username: %v", err)
		}
	}

	if err = d.Set("sandbox_region", flattenObjectSystemFortiguardSandboxRegion(o["sandbox-region"], d, "sandbox_region")); err != nil {
		if vv, ok := fortiAPIPatch(o["sandbox-region"], "ObjectSystemFortiguard-SandboxRegion"); ok {
			if err = d.Set("sandbox_region", vv); err != nil {
				return fmt.Errorf("Error reading sandbox_region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sandbox_region: %v", err)
		}
	}

	if err = d.Set("sdns_options", flattenObjectSystemFortiguardSdnsOptions(o["sdns-options"], d, "sdns_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdns-options"], "ObjectSystemFortiguard-SdnsOptions"); ok {
			if err = d.Set("sdns_options", vv); err != nil {
				return fmt.Errorf("Error reading sdns_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdns_options: %v", err)
		}
	}

	if err = d.Set("sdns_server_ip", flattenObjectSystemFortiguardSdnsServerIp(o["sdns-server-ip"], d, "sdns_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdns-server-ip"], "ObjectSystemFortiguard-SdnsServerIp"); ok {
			if err = d.Set("sdns_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading sdns_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdns_server_ip: %v", err)
		}
	}

	if err = d.Set("sdns_server_port", flattenObjectSystemFortiguardSdnsServerPort(o["sdns-server-port"], d, "sdns_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdns-server-port"], "ObjectSystemFortiguard-SdnsServerPort"); ok {
			if err = d.Set("sdns_server_port", vv); err != nil {
				return fmt.Errorf("Error reading sdns_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdns_server_port: %v", err)
		}
	}

	if err = d.Set("service_account_id", flattenObjectSystemFortiguardServiceAccountId(o["service-account-id"], d, "service_account_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-account-id"], "ObjectSystemFortiguard-ServiceAccountId"); ok {
			if err = d.Set("service_account_id", vv); err != nil {
				return fmt.Errorf("Error reading service_account_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_account_id: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenObjectSystemFortiguardSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "ObjectSystemFortiguard-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenObjectSystemFortiguardSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip6"], "ObjectSystemFortiguard-SourceIp6"); ok {
			if err = d.Set("source_ip6", vv); err != nil {
				return fmt.Errorf("Error reading source_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("update_server_location", flattenObjectSystemFortiguardUpdateServerLocation(o["update-server-location"], d, "update_server_location")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-server-location"], "ObjectSystemFortiguard-UpdateServerLocation"); ok {
			if err = d.Set("update_server_location", vv); err != nil {
				return fmt.Errorf("Error reading update_server_location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_server_location: %v", err)
		}
	}

	if err = d.Set("webfilter_cache", flattenObjectSystemFortiguardWebfilterCache(o["webfilter-cache"], d, "webfilter_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-cache"], "ObjectSystemFortiguard-WebfilterCache"); ok {
			if err = d.Set("webfilter_cache", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_cache: %v", err)
		}
	}

	if err = d.Set("webfilter_cache_ttl", flattenObjectSystemFortiguardWebfilterCacheTtl(o["webfilter-cache-ttl"], d, "webfilter_cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-cache-ttl"], "ObjectSystemFortiguard-WebfilterCacheTtl"); ok {
			if err = d.Set("webfilter_cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_cache_ttl: %v", err)
		}
	}

	if err = d.Set("webfilter_expiration", flattenObjectSystemFortiguardWebfilterExpiration(o["webfilter-expiration"], d, "webfilter_expiration")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-expiration"], "ObjectSystemFortiguard-WebfilterExpiration"); ok {
			if err = d.Set("webfilter_expiration", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_expiration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_expiration: %v", err)
		}
	}

	if err = d.Set("webfilter_force_off", flattenObjectSystemFortiguardWebfilterForceOff(o["webfilter-force-off"], d, "webfilter_force_off")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-force-off"], "ObjectSystemFortiguard-WebfilterForceOff"); ok {
			if err = d.Set("webfilter_force_off", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_force_off: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_force_off: %v", err)
		}
	}

	if err = d.Set("webfilter_license", flattenObjectSystemFortiguardWebfilterLicense(o["webfilter-license"], d, "webfilter_license")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-license"], "ObjectSystemFortiguard-WebfilterLicense"); ok {
			if err = d.Set("webfilter_license", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_license: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_license: %v", err)
		}
	}

	if err = d.Set("webfilter_timeout", flattenObjectSystemFortiguardWebfilterTimeout(o["webfilter-timeout"], d, "webfilter_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-timeout"], "ObjectSystemFortiguard-WebfilterTimeout"); ok {
			if err = d.Set("webfilter_timeout", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_timeout: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemFortiguardFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemFortiguardAntispamCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardAntispamCacheMpercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardAntispamCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardAntispamExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardAntispamForceOff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardAntispamLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardAntispamTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardAnycastSdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardAnycastSdnsServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardAutoJoinForticloud(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardDdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardDdnsServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardFortiguardAnycast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardFortiguardAnycastSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardLoadBalanceServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardOutbreakPreventionCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardOutbreakPreventionCacheMpercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardOutbreakPreventionCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardOutbreakPreventionExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardOutbreakPreventionForceOff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardOutbreakPreventionLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardOutbreakPreventionTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardProxyPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemFortiguardProxyServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardProxyServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardProxyUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardSandboxRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardSdnsOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemFortiguardSdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemFortiguardSdnsServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardServiceAccountId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardUpdateServerLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardWebfilterCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardWebfilterCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardWebfilterExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardWebfilterForceOff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardWebfilterLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemFortiguardWebfilterTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemFortiguard(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("antispam_cache"); ok {
		t, err := expandObjectSystemFortiguardAntispamCache(d, v, "antispam_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-cache"] = t
		}
	}

	if v, ok := d.GetOk("antispam_cache_mpercent"); ok {
		t, err := expandObjectSystemFortiguardAntispamCacheMpercent(d, v, "antispam_cache_mpercent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-cache-mpercent"] = t
		}
	}

	if v, ok := d.GetOk("antispam_cache_ttl"); ok {
		t, err := expandObjectSystemFortiguardAntispamCacheTtl(d, v, "antispam_cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("antispam_expiration"); ok {
		t, err := expandObjectSystemFortiguardAntispamExpiration(d, v, "antispam_expiration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-expiration"] = t
		}
	}

	if v, ok := d.GetOk("antispam_force_off"); ok {
		t, err := expandObjectSystemFortiguardAntispamForceOff(d, v, "antispam_force_off")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-force-off"] = t
		}
	}

	if v, ok := d.GetOk("antispam_license"); ok {
		t, err := expandObjectSystemFortiguardAntispamLicense(d, v, "antispam_license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-license"] = t
		}
	}

	if v, ok := d.GetOk("antispam_timeout"); ok {
		t, err := expandObjectSystemFortiguardAntispamTimeout(d, v, "antispam_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-timeout"] = t
		}
	}

	if v, ok := d.GetOk("anycast_sdns_server_ip"); ok {
		t, err := expandObjectSystemFortiguardAnycastSdnsServerIp(d, v, "anycast_sdns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anycast-sdns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("anycast_sdns_server_port"); ok {
		t, err := expandObjectSystemFortiguardAnycastSdnsServerPort(d, v, "anycast_sdns_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anycast-sdns-server-port"] = t
		}
	}

	if v, ok := d.GetOk("auto_join_forticloud"); ok {
		t, err := expandObjectSystemFortiguardAutoJoinForticloud(d, v, "auto_join_forticloud")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-join-forticloud"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_ip"); ok {
		t, err := expandObjectSystemFortiguardDdnsServerIp(d, v, "ddns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_port"); ok {
		t, err := expandObjectSystemFortiguardDdnsServerPort(d, v, "ddns_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-port"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_anycast"); ok {
		t, err := expandObjectSystemFortiguardFortiguardAnycast(d, v, "fortiguard_anycast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-anycast"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_anycast_source"); ok {
		t, err := expandObjectSystemFortiguardFortiguardAnycastSource(d, v, "fortiguard_anycast_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-anycast-source"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandObjectSystemFortiguardInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		t, err := expandObjectSystemFortiguardInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("load_balance_servers"); ok {
		t, err := expandObjectSystemFortiguardLoadBalanceServers(d, v, "load_balance_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance-servers"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_cache"); ok {
		t, err := expandObjectSystemFortiguardOutbreakPreventionCache(d, v, "outbreak_prevention_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-cache"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_cache_mpercent"); ok {
		t, err := expandObjectSystemFortiguardOutbreakPreventionCacheMpercent(d, v, "outbreak_prevention_cache_mpercent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-cache-mpercent"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_cache_ttl"); ok {
		t, err := expandObjectSystemFortiguardOutbreakPreventionCacheTtl(d, v, "outbreak_prevention_cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_expiration"); ok {
		t, err := expandObjectSystemFortiguardOutbreakPreventionExpiration(d, v, "outbreak_prevention_expiration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-expiration"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_force_off"); ok {
		t, err := expandObjectSystemFortiguardOutbreakPreventionForceOff(d, v, "outbreak_prevention_force_off")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-force-off"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_license"); ok {
		t, err := expandObjectSystemFortiguardOutbreakPreventionLicense(d, v, "outbreak_prevention_license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-license"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_timeout"); ok {
		t, err := expandObjectSystemFortiguardOutbreakPreventionTimeout(d, v, "outbreak_prevention_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-timeout"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandObjectSystemFortiguardPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandObjectSystemFortiguardProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("proxy_password"); ok {
		t, err := expandObjectSystemFortiguardProxyPassword(d, v, "proxy_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-password"] = t
		}
	}

	if v, ok := d.GetOk("proxy_server_ip"); ok {
		t, err := expandObjectSystemFortiguardProxyServerIp(d, v, "proxy_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("proxy_server_port"); ok {
		t, err := expandObjectSystemFortiguardProxyServerPort(d, v, "proxy_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-server-port"] = t
		}
	}

	if v, ok := d.GetOk("proxy_username"); ok {
		t, err := expandObjectSystemFortiguardProxyUsername(d, v, "proxy_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-username"] = t
		}
	}

	if v, ok := d.GetOk("sandbox_region"); ok {
		t, err := expandObjectSystemFortiguardSandboxRegion(d, v, "sandbox_region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sandbox-region"] = t
		}
	}

	if v, ok := d.GetOk("sdns_options"); ok {
		t, err := expandObjectSystemFortiguardSdnsOptions(d, v, "sdns_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdns-options"] = t
		}
	}

	if v, ok := d.GetOk("sdns_server_ip"); ok {
		t, err := expandObjectSystemFortiguardSdnsServerIp(d, v, "sdns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("sdns_server_port"); ok {
		t, err := expandObjectSystemFortiguardSdnsServerPort(d, v, "sdns_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdns-server-port"] = t
		}
	}

	if v, ok := d.GetOk("service_account_id"); ok {
		t, err := expandObjectSystemFortiguardServiceAccountId(d, v, "service_account_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-account-id"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandObjectSystemFortiguardSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok {
		t, err := expandObjectSystemFortiguardSourceIp6(d, v, "source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("update_server_location"); ok {
		t, err := expandObjectSystemFortiguardUpdateServerLocation(d, v, "update_server_location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-server-location"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_cache"); ok {
		t, err := expandObjectSystemFortiguardWebfilterCache(d, v, "webfilter_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-cache"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_cache_ttl"); ok {
		t, err := expandObjectSystemFortiguardWebfilterCacheTtl(d, v, "webfilter_cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_expiration"); ok {
		t, err := expandObjectSystemFortiguardWebfilterExpiration(d, v, "webfilter_expiration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-expiration"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_force_off"); ok {
		t, err := expandObjectSystemFortiguardWebfilterForceOff(d, v, "webfilter_force_off")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-force-off"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_license"); ok {
		t, err := expandObjectSystemFortiguardWebfilterLicense(d, v, "webfilter_license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-license"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_timeout"); ok {
		t, err := expandObjectSystemFortiguardWebfilterTimeout(d, v, "webfilter_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-timeout"] = t
		}
	}

	return &obj, nil
}
