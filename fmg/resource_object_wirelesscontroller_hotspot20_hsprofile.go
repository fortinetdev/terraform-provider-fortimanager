// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure hotspot profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWirelessControllerHotspot20HsProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20HsProfileCreate,
		Read:   resourceObjectWirelessControllerHotspot20HsProfileRead,
		Update: resourceObjectWirelessControllerHotspot20HsProfileUpdate,
		Delete: resourceObjectWirelessControllerHotspot20HsProfileDelete,

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
			"n3gpp_plmn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_network_asra": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_network_esr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_network_internet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_network_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_network_uesa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anqp_domain_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bss_transition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conn_cap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deauth_request_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dgaf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gas_comeback_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gas_fragmentation_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hessid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l2tif": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nai_realm": &schema.Schema{
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
			"network_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oper_friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"osu_provider": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"osu_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pame_bi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_arp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qos_map": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"roaming_consortium": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"venue_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"venue_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"venue_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wan_metrics": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wnm_sleep_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20HsProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20HsProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20HsProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerHotspot20HsProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20HsProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20HsProfileRead(d, m)
}

func resourceObjectWirelessControllerHotspot20HsProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20HsProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20HsProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerHotspot20HsProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20HsProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20HsProfileRead(d, m)
}

func resourceObjectWirelessControllerHotspot20HsProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerHotspot20HsProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20HsProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20HsProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerHotspot20HsProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20HsProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20HsProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20HsProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20HsProfile3GppPlmn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileAccessNetworkAsra(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectWirelessControllerHotspot20HsProfileAccessNetworkEsr(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectWirelessControllerHotspot20HsProfileAccessNetworkInternet(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectWirelessControllerHotspot20HsProfileAccessNetworkType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "private-network",
			1: "private-network-with-guest-access",
			2: "chargeable-public-network",
			3: "free-public-network",
			4: "personal-device-network",
			5: "emergency-services-only-network",
			6: "test-or-experimental",
			7: "wildcard",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileAccessNetworkUesa(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectWirelessControllerHotspot20HsProfileAnqpDomainId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileBssTransition(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectWirelessControllerHotspot20HsProfileConnCap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileDeauthRequestTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileDgaf(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectWirelessControllerHotspot20HsProfileDomainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileGasComebackDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileGasFragmentationLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileHessid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileIpAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileL2Tif(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectWirelessControllerHotspot20HsProfileNaiRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileNetworkAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileOperFriendlyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileOsuProvider(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileOsuSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfilePameBi(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectWirelessControllerHotspot20HsProfileProxyArp(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectWirelessControllerHotspot20HsProfileQosMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileRoamingConsortium(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileVenueGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:  "unspecified",
			1:  "assembly",
			2:  "business",
			3:  "educational",
			4:  "factory",
			5:  "institutional",
			6:  "mercantile",
			7:  "residential",
			8:  "storage",
			9:  "utility",
			10: "vehicular",
			11: "outdoor",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileVenueName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileVenueType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:  "unspecified",
			1:  "arena",
			2:  "stadium",
			3:  "passenger-terminal",
			4:  "amphitheater",
			5:  "amusement-park",
			6:  "place-of-worship",
			7:  "convention-center",
			8:  "library",
			9:  "museum",
			10: "restaurant",
			11: "theater",
			12: "bar",
			13: "coffee-shop",
			14: "zoo-or-aquarium",
			15: "emergency-center",
			16: "doctor-office",
			17: "bank",
			18: "fire-station",
			19: "police-station",
			20: "post-office",
			21: "professional-office",
			22: "research-facility",
			23: "attorney-office",
			24: "primary-school",
			25: "secondary-school",
			26: "university-or-college",
			27: "factory",
			28: "hospital",
			29: "long-term-care-facility",
			30: "rehab-center",
			31: "group-home",
			32: "prison-or-jail",
			33: "retail-store",
			34: "grocery-market",
			35: "auto-service-station",
			36: "shopping-mall",
			37: "gas-station",
			38: "private",
			39: "hotel-or-motel",
			40: "dormitory",
			41: "boarding-house",
			42: "automobile",
			43: "airplane",
			44: "bus",
			45: "ferry",
			46: "ship-or-boat",
			47: "train",
			48: "motor-bike",
			49: "muni-mesh-network",
			50: "city-park",
			51: "rest-area",
			52: "traffic-control",
			53: "bus-stop",
			54: "kiosk",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileWanMetrics(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20HsProfileWnmSleepMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func refreshObjectObjectWirelessControllerHotspot20HsProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("n3gpp_plmn", flattenObjectWirelessControllerHotspot20HsProfile3GppPlmn(o["3gpp-plmn"], d, "n3gpp_plmn")); err != nil {
		if vv, ok := fortiAPIPatch(o["3gpp-plmn"], "ObjectWirelessControllerHotspot20HsProfile-3GppPlmn"); ok {
			if err = d.Set("n3gpp_plmn", vv); err != nil {
				return fmt.Errorf("Error reading n3gpp_plmn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n3gpp_plmn: %v", err)
		}
	}

	if err = d.Set("access_network_asra", flattenObjectWirelessControllerHotspot20HsProfileAccessNetworkAsra(o["access-network-asra"], d, "access_network_asra")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-network-asra"], "ObjectWirelessControllerHotspot20HsProfile-AccessNetworkAsra"); ok {
			if err = d.Set("access_network_asra", vv); err != nil {
				return fmt.Errorf("Error reading access_network_asra: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_network_asra: %v", err)
		}
	}

	if err = d.Set("access_network_esr", flattenObjectWirelessControllerHotspot20HsProfileAccessNetworkEsr(o["access-network-esr"], d, "access_network_esr")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-network-esr"], "ObjectWirelessControllerHotspot20HsProfile-AccessNetworkEsr"); ok {
			if err = d.Set("access_network_esr", vv); err != nil {
				return fmt.Errorf("Error reading access_network_esr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_network_esr: %v", err)
		}
	}

	if err = d.Set("access_network_internet", flattenObjectWirelessControllerHotspot20HsProfileAccessNetworkInternet(o["access-network-internet"], d, "access_network_internet")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-network-internet"], "ObjectWirelessControllerHotspot20HsProfile-AccessNetworkInternet"); ok {
			if err = d.Set("access_network_internet", vv); err != nil {
				return fmt.Errorf("Error reading access_network_internet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_network_internet: %v", err)
		}
	}

	if err = d.Set("access_network_type", flattenObjectWirelessControllerHotspot20HsProfileAccessNetworkType(o["access-network-type"], d, "access_network_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-network-type"], "ObjectWirelessControllerHotspot20HsProfile-AccessNetworkType"); ok {
			if err = d.Set("access_network_type", vv); err != nil {
				return fmt.Errorf("Error reading access_network_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_network_type: %v", err)
		}
	}

	if err = d.Set("access_network_uesa", flattenObjectWirelessControllerHotspot20HsProfileAccessNetworkUesa(o["access-network-uesa"], d, "access_network_uesa")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-network-uesa"], "ObjectWirelessControllerHotspot20HsProfile-AccessNetworkUesa"); ok {
			if err = d.Set("access_network_uesa", vv); err != nil {
				return fmt.Errorf("Error reading access_network_uesa: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_network_uesa: %v", err)
		}
	}

	if err = d.Set("anqp_domain_id", flattenObjectWirelessControllerHotspot20HsProfileAnqpDomainId(o["anqp-domain-id"], d, "anqp_domain_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["anqp-domain-id"], "ObjectWirelessControllerHotspot20HsProfile-AnqpDomainId"); ok {
			if err = d.Set("anqp_domain_id", vv); err != nil {
				return fmt.Errorf("Error reading anqp_domain_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anqp_domain_id: %v", err)
		}
	}

	if err = d.Set("bss_transition", flattenObjectWirelessControllerHotspot20HsProfileBssTransition(o["bss-transition"], d, "bss_transition")); err != nil {
		if vv, ok := fortiAPIPatch(o["bss-transition"], "ObjectWirelessControllerHotspot20HsProfile-BssTransition"); ok {
			if err = d.Set("bss_transition", vv); err != nil {
				return fmt.Errorf("Error reading bss_transition: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bss_transition: %v", err)
		}
	}

	if err = d.Set("conn_cap", flattenObjectWirelessControllerHotspot20HsProfileConnCap(o["conn-cap"], d, "conn_cap")); err != nil {
		if vv, ok := fortiAPIPatch(o["conn-cap"], "ObjectWirelessControllerHotspot20HsProfile-ConnCap"); ok {
			if err = d.Set("conn_cap", vv); err != nil {
				return fmt.Errorf("Error reading conn_cap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conn_cap: %v", err)
		}
	}

	if err = d.Set("deauth_request_timeout", flattenObjectWirelessControllerHotspot20HsProfileDeauthRequestTimeout(o["deauth-request-timeout"], d, "deauth_request_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["deauth-request-timeout"], "ObjectWirelessControllerHotspot20HsProfile-DeauthRequestTimeout"); ok {
			if err = d.Set("deauth_request_timeout", vv); err != nil {
				return fmt.Errorf("Error reading deauth_request_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deauth_request_timeout: %v", err)
		}
	}

	if err = d.Set("dgaf", flattenObjectWirelessControllerHotspot20HsProfileDgaf(o["dgaf"], d, "dgaf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dgaf"], "ObjectWirelessControllerHotspot20HsProfile-Dgaf"); ok {
			if err = d.Set("dgaf", vv); err != nil {
				return fmt.Errorf("Error reading dgaf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dgaf: %v", err)
		}
	}

	if err = d.Set("domain_name", flattenObjectWirelessControllerHotspot20HsProfileDomainName(o["domain-name"], d, "domain_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-name"], "ObjectWirelessControllerHotspot20HsProfile-DomainName"); ok {
			if err = d.Set("domain_name", vv); err != nil {
				return fmt.Errorf("Error reading domain_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if err = d.Set("gas_comeback_delay", flattenObjectWirelessControllerHotspot20HsProfileGasComebackDelay(o["gas-comeback-delay"], d, "gas_comeback_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["gas-comeback-delay"], "ObjectWirelessControllerHotspot20HsProfile-GasComebackDelay"); ok {
			if err = d.Set("gas_comeback_delay", vv); err != nil {
				return fmt.Errorf("Error reading gas_comeback_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gas_comeback_delay: %v", err)
		}
	}

	if err = d.Set("gas_fragmentation_limit", flattenObjectWirelessControllerHotspot20HsProfileGasFragmentationLimit(o["gas-fragmentation-limit"], d, "gas_fragmentation_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["gas-fragmentation-limit"], "ObjectWirelessControllerHotspot20HsProfile-GasFragmentationLimit"); ok {
			if err = d.Set("gas_fragmentation_limit", vv); err != nil {
				return fmt.Errorf("Error reading gas_fragmentation_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gas_fragmentation_limit: %v", err)
		}
	}

	if err = d.Set("hessid", flattenObjectWirelessControllerHotspot20HsProfileHessid(o["hessid"], d, "hessid")); err != nil {
		if vv, ok := fortiAPIPatch(o["hessid"], "ObjectWirelessControllerHotspot20HsProfile-Hessid"); ok {
			if err = d.Set("hessid", vv); err != nil {
				return fmt.Errorf("Error reading hessid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hessid: %v", err)
		}
	}

	if err = d.Set("ip_addr_type", flattenObjectWirelessControllerHotspot20HsProfileIpAddrType(o["ip-addr-type"], d, "ip_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-addr-type"], "ObjectWirelessControllerHotspot20HsProfile-IpAddrType"); ok {
			if err = d.Set("ip_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading ip_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_addr_type: %v", err)
		}
	}

	if err = d.Set("l2tif", flattenObjectWirelessControllerHotspot20HsProfileL2Tif(o["l2tif"], d, "l2tif")); err != nil {
		if vv, ok := fortiAPIPatch(o["l2tif"], "ObjectWirelessControllerHotspot20HsProfile-L2Tif"); ok {
			if err = d.Set("l2tif", vv); err != nil {
				return fmt.Errorf("Error reading l2tif: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l2tif: %v", err)
		}
	}

	if err = d.Set("nai_realm", flattenObjectWirelessControllerHotspot20HsProfileNaiRealm(o["nai-realm"], d, "nai_realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["nai-realm"], "ObjectWirelessControllerHotspot20HsProfile-NaiRealm"); ok {
			if err = d.Set("nai_realm", vv); err != nil {
				return fmt.Errorf("Error reading nai_realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nai_realm: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerHotspot20HsProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerHotspot20HsProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("network_auth", flattenObjectWirelessControllerHotspot20HsProfileNetworkAuth(o["network-auth"], d, "network_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-auth"], "ObjectWirelessControllerHotspot20HsProfile-NetworkAuth"); ok {
			if err = d.Set("network_auth", vv); err != nil {
				return fmt.Errorf("Error reading network_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_auth: %v", err)
		}
	}

	if err = d.Set("oper_friendly_name", flattenObjectWirelessControllerHotspot20HsProfileOperFriendlyName(o["oper-friendly-name"], d, "oper_friendly_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["oper-friendly-name"], "ObjectWirelessControllerHotspot20HsProfile-OperFriendlyName"); ok {
			if err = d.Set("oper_friendly_name", vv); err != nil {
				return fmt.Errorf("Error reading oper_friendly_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oper_friendly_name: %v", err)
		}
	}

	if err = d.Set("osu_provider", flattenObjectWirelessControllerHotspot20HsProfileOsuProvider(o["osu-provider"], d, "osu_provider")); err != nil {
		if vv, ok := fortiAPIPatch(o["osu-provider"], "ObjectWirelessControllerHotspot20HsProfile-OsuProvider"); ok {
			if err = d.Set("osu_provider", vv); err != nil {
				return fmt.Errorf("Error reading osu_provider: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading osu_provider: %v", err)
		}
	}

	if err = d.Set("osu_ssid", flattenObjectWirelessControllerHotspot20HsProfileOsuSsid(o["osu-ssid"], d, "osu_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["osu-ssid"], "ObjectWirelessControllerHotspot20HsProfile-OsuSsid"); ok {
			if err = d.Set("osu_ssid", vv); err != nil {
				return fmt.Errorf("Error reading osu_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading osu_ssid: %v", err)
		}
	}

	if err = d.Set("pame_bi", flattenObjectWirelessControllerHotspot20HsProfilePameBi(o["pame-bi"], d, "pame_bi")); err != nil {
		if vv, ok := fortiAPIPatch(o["pame-bi"], "ObjectWirelessControllerHotspot20HsProfile-PameBi"); ok {
			if err = d.Set("pame_bi", vv); err != nil {
				return fmt.Errorf("Error reading pame_bi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pame_bi: %v", err)
		}
	}

	if err = d.Set("proxy_arp", flattenObjectWirelessControllerHotspot20HsProfileProxyArp(o["proxy-arp"], d, "proxy_arp")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-arp"], "ObjectWirelessControllerHotspot20HsProfile-ProxyArp"); ok {
			if err = d.Set("proxy_arp", vv); err != nil {
				return fmt.Errorf("Error reading proxy_arp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_arp: %v", err)
		}
	}

	if err = d.Set("qos_map", flattenObjectWirelessControllerHotspot20HsProfileQosMap(o["qos-map"], d, "qos_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-map"], "ObjectWirelessControllerHotspot20HsProfile-QosMap"); ok {
			if err = d.Set("qos_map", vv); err != nil {
				return fmt.Errorf("Error reading qos_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_map: %v", err)
		}
	}

	if err = d.Set("roaming_consortium", flattenObjectWirelessControllerHotspot20HsProfileRoamingConsortium(o["roaming-consortium"], d, "roaming_consortium")); err != nil {
		if vv, ok := fortiAPIPatch(o["roaming-consortium"], "ObjectWirelessControllerHotspot20HsProfile-RoamingConsortium"); ok {
			if err = d.Set("roaming_consortium", vv); err != nil {
				return fmt.Errorf("Error reading roaming_consortium: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading roaming_consortium: %v", err)
		}
	}

	if err = d.Set("venue_group", flattenObjectWirelessControllerHotspot20HsProfileVenueGroup(o["venue-group"], d, "venue_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["venue-group"], "ObjectWirelessControllerHotspot20HsProfile-VenueGroup"); ok {
			if err = d.Set("venue_group", vv); err != nil {
				return fmt.Errorf("Error reading venue_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading venue_group: %v", err)
		}
	}

	if err = d.Set("venue_name", flattenObjectWirelessControllerHotspot20HsProfileVenueName(o["venue-name"], d, "venue_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["venue-name"], "ObjectWirelessControllerHotspot20HsProfile-VenueName"); ok {
			if err = d.Set("venue_name", vv); err != nil {
				return fmt.Errorf("Error reading venue_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading venue_name: %v", err)
		}
	}

	if err = d.Set("venue_type", flattenObjectWirelessControllerHotspot20HsProfileVenueType(o["venue-type"], d, "venue_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["venue-type"], "ObjectWirelessControllerHotspot20HsProfile-VenueType"); ok {
			if err = d.Set("venue_type", vv); err != nil {
				return fmt.Errorf("Error reading venue_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading venue_type: %v", err)
		}
	}

	if err = d.Set("wan_metrics", flattenObjectWirelessControllerHotspot20HsProfileWanMetrics(o["wan-metrics"], d, "wan_metrics")); err != nil {
		if vv, ok := fortiAPIPatch(o["wan-metrics"], "ObjectWirelessControllerHotspot20HsProfile-WanMetrics"); ok {
			if err = d.Set("wan_metrics", vv); err != nil {
				return fmt.Errorf("Error reading wan_metrics: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wan_metrics: %v", err)
		}
	}

	if err = d.Set("wnm_sleep_mode", flattenObjectWirelessControllerHotspot20HsProfileWnmSleepMode(o["wnm-sleep-mode"], d, "wnm_sleep_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["wnm-sleep-mode"], "ObjectWirelessControllerHotspot20HsProfile-WnmSleepMode"); ok {
			if err = d.Set("wnm_sleep_mode", vv); err != nil {
				return fmt.Errorf("Error reading wnm_sleep_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wnm_sleep_mode: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20HsProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20HsProfile3GppPlmn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileAccessNetworkAsra(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileAccessNetworkEsr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileAccessNetworkInternet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileAccessNetworkType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileAccessNetworkUesa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileAnqpDomainId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileBssTransition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileConnCap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileDeauthRequestTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileDgaf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileDomainName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileGasComebackDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileGasFragmentationLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileHessid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileIpAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileL2Tif(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileNaiRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileNetworkAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileOperFriendlyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileOsuProvider(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileOsuSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfilePameBi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileProxyArp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileQosMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileRoamingConsortium(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileVenueGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileVenueName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileVenueType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileWanMetrics(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20HsProfileWnmSleepMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20HsProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("n3gpp_plmn"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfile3GppPlmn(d, v, "n3gpp_plmn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["3gpp-plmn"] = t
		}
	}

	if v, ok := d.GetOk("access_network_asra"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileAccessNetworkAsra(d, v, "access_network_asra")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-network-asra"] = t
		}
	}

	if v, ok := d.GetOk("access_network_esr"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileAccessNetworkEsr(d, v, "access_network_esr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-network-esr"] = t
		}
	}

	if v, ok := d.GetOk("access_network_internet"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileAccessNetworkInternet(d, v, "access_network_internet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-network-internet"] = t
		}
	}

	if v, ok := d.GetOk("access_network_type"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileAccessNetworkType(d, v, "access_network_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-network-type"] = t
		}
	}

	if v, ok := d.GetOk("access_network_uesa"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileAccessNetworkUesa(d, v, "access_network_uesa")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-network-uesa"] = t
		}
	}

	if v, ok := d.GetOk("anqp_domain_id"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileAnqpDomainId(d, v, "anqp_domain_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anqp-domain-id"] = t
		}
	}

	if v, ok := d.GetOk("bss_transition"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileBssTransition(d, v, "bss_transition")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bss-transition"] = t
		}
	}

	if v, ok := d.GetOk("conn_cap"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileConnCap(d, v, "conn_cap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-cap"] = t
		}
	}

	if v, ok := d.GetOk("deauth_request_timeout"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileDeauthRequestTimeout(d, v, "deauth_request_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deauth-request-timeout"] = t
		}
	}

	if v, ok := d.GetOk("dgaf"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileDgaf(d, v, "dgaf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dgaf"] = t
		}
	}

	if v, ok := d.GetOk("domain_name"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileDomainName(d, v, "domain_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name"] = t
		}
	}

	if v, ok := d.GetOk("gas_comeback_delay"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileGasComebackDelay(d, v, "gas_comeback_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gas-comeback-delay"] = t
		}
	}

	if v, ok := d.GetOk("gas_fragmentation_limit"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileGasFragmentationLimit(d, v, "gas_fragmentation_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gas-fragmentation-limit"] = t
		}
	}

	if v, ok := d.GetOk("hessid"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileHessid(d, v, "hessid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hessid"] = t
		}
	}

	if v, ok := d.GetOk("ip_addr_type"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileIpAddrType(d, v, "ip_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("l2tif"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileL2Tif(d, v, "l2tif")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2tif"] = t
		}
	}

	if v, ok := d.GetOk("nai_realm"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileNaiRealm(d, v, "nai_realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nai-realm"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("network_auth"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileNetworkAuth(d, v, "network_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-auth"] = t
		}
	}

	if v, ok := d.GetOk("oper_friendly_name"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileOperFriendlyName(d, v, "oper_friendly_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oper-friendly-name"] = t
		}
	}

	if v, ok := d.GetOk("osu_provider"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileOsuProvider(d, v, "osu_provider")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osu-provider"] = t
		}
	}

	if v, ok := d.GetOk("osu_ssid"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileOsuSsid(d, v, "osu_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osu-ssid"] = t
		}
	}

	if v, ok := d.GetOk("pame_bi"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfilePameBi(d, v, "pame_bi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pame-bi"] = t
		}
	}

	if v, ok := d.GetOk("proxy_arp"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileProxyArp(d, v, "proxy_arp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-arp"] = t
		}
	}

	if v, ok := d.GetOk("qos_map"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileQosMap(d, v, "qos_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-map"] = t
		}
	}

	if v, ok := d.GetOk("roaming_consortium"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileRoamingConsortium(d, v, "roaming_consortium")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roaming-consortium"] = t
		}
	}

	if v, ok := d.GetOk("venue_group"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileVenueGroup(d, v, "venue_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["venue-group"] = t
		}
	}

	if v, ok := d.GetOk("venue_name"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileVenueName(d, v, "venue_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["venue-name"] = t
		}
	}

	if v, ok := d.GetOk("venue_type"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileVenueType(d, v, "venue_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["venue-type"] = t
		}
	}

	if v, ok := d.GetOk("wan_metrics"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileWanMetrics(d, v, "wan_metrics")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-metrics"] = t
		}
	}

	if v, ok := d.GetOk("wnm_sleep_mode"); ok {
		t, err := expandObjectWirelessControllerHotspot20HsProfileWnmSleepMode(d, v, "wnm_sleep_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wnm-sleep-mode"] = t
		}
	}

	return &obj, nil
}
