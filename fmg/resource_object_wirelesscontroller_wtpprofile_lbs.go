// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Set various location based service (LBS) options.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerWtpProfileLbs() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerWtpProfileLbsUpdate,
		Read:   resourceObjectWirelessControllerWtpProfileLbsRead,
		Update: resourceObjectWirelessControllerWtpProfileLbsUpdate,
		Delete: resourceObjectWirelessControllerWtpProfileLbsDelete,

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
			"wtp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"aeroscout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aeroscout_ap_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aeroscout_mmu_report": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aeroscout_mu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aeroscout_mu_factor": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"aeroscout_mu_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"aeroscout_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aeroscout_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ekahau_blink_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ekahau_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"erc_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"erc_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fortipresence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortipresence_ble": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortipresence_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fortipresence_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fortipresence_project": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortipresence_rogue": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortipresence_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"fortipresence_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortipresence_server_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortipresence_server_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortipresence_unassoc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"polestar": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"polestar_accumulation_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"polestar_asset_addrgrp_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"polestar_asset_uuid_list1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"polestar_asset_uuid_list2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"polestar_asset_uuid_list3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"polestar_asset_uuid_list4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"polestar_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"polestar_reporting_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"polestar_server_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"polestar_server_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"polestar_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"polestar_server_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"station_locate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerWtpProfileLbsUpdate(d *schema.ResourceData, m interface{}) error {
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

	wtp_profile := d.Get("wtp_profile").(string)
	paradict["wtp_profile"] = wtp_profile

	obj, err := getObjectObjectWirelessControllerWtpProfileLbs(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWtpProfileLbs resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerWtpProfileLbs(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWtpProfileLbs resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectWirelessControllerWtpProfileLbs")

	return resourceObjectWirelessControllerWtpProfileLbsRead(d, m)
}

func resourceObjectWirelessControllerWtpProfileLbsDelete(d *schema.ResourceData, m interface{}) error {
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

	wtp_profile := d.Get("wtp_profile").(string)
	paradict["wtp_profile"] = wtp_profile

	err = c.DeleteObjectWirelessControllerWtpProfileLbs(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerWtpProfileLbs resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerWtpProfileLbsRead(d *schema.ResourceData, m interface{}) error {
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

	wtp_profile := d.Get("wtp_profile").(string)
	if wtp_profile == "" {
		wtp_profile = importOptionChecking(m.(*FortiClient).Cfg, "wtp_profile")
		if wtp_profile == "" {
			return fmt.Errorf("Parameter wtp_profile is missing")
		}
		if err = d.Set("wtp_profile", wtp_profile); err != nil {
			return fmt.Errorf("Error set params wtp_profile: %v", err)
		}
	}
	paradict["wtp_profile"] = wtp_profile

	o, err := c.ReadObjectWirelessControllerWtpProfileLbs(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWtpProfileLbs resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerWtpProfileLbs(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWtpProfileLbs resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscoutApMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscoutMmuReport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscoutMu2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscoutMuFactor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscoutMuTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscoutServerIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscoutServerPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsEkahauBlinkMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsEkahauTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsErcServerIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsErcServerPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresence2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresenceBle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresenceFrequency2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresencePort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresenceProject2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresenceRogue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresenceServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresenceServerAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresenceServerFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresenceUnassoc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsPolestar2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsPolestarAccumulationInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsPolestarAssetAddrgrpList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsPolestarProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsPolestarReportingInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsPolestarServerFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsPolestarServerPath2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsPolestarServerPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsPolestarServerToken2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsStationLocate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerWtpProfileLbs(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("aeroscout", flattenObjectWirelessControllerWtpProfileLbsAeroscout2edl(o["aeroscout"], d, "aeroscout")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout"], "ObjectWirelessControllerWtpProfileLbs-Aeroscout"); ok {
			if err = d.Set("aeroscout", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout: %v", err)
		}
	}

	if err = d.Set("aeroscout_ap_mac", flattenObjectWirelessControllerWtpProfileLbsAeroscoutApMac2edl(o["aeroscout-ap-mac"], d, "aeroscout_ap_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout-ap-mac"], "ObjectWirelessControllerWtpProfileLbs-AeroscoutApMac"); ok {
			if err = d.Set("aeroscout_ap_mac", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout_ap_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout_ap_mac: %v", err)
		}
	}

	if err = d.Set("aeroscout_mmu_report", flattenObjectWirelessControllerWtpProfileLbsAeroscoutMmuReport2edl(o["aeroscout-mmu-report"], d, "aeroscout_mmu_report")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout-mmu-report"], "ObjectWirelessControllerWtpProfileLbs-AeroscoutMmuReport"); ok {
			if err = d.Set("aeroscout_mmu_report", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout_mmu_report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout_mmu_report: %v", err)
		}
	}

	if err = d.Set("aeroscout_mu", flattenObjectWirelessControllerWtpProfileLbsAeroscoutMu2edl(o["aeroscout-mu"], d, "aeroscout_mu")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout-mu"], "ObjectWirelessControllerWtpProfileLbs-AeroscoutMu"); ok {
			if err = d.Set("aeroscout_mu", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout_mu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout_mu: %v", err)
		}
	}

	if err = d.Set("aeroscout_mu_factor", flattenObjectWirelessControllerWtpProfileLbsAeroscoutMuFactor2edl(o["aeroscout-mu-factor"], d, "aeroscout_mu_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout-mu-factor"], "ObjectWirelessControllerWtpProfileLbs-AeroscoutMuFactor"); ok {
			if err = d.Set("aeroscout_mu_factor", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout_mu_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout_mu_factor: %v", err)
		}
	}

	if err = d.Set("aeroscout_mu_timeout", flattenObjectWirelessControllerWtpProfileLbsAeroscoutMuTimeout2edl(o["aeroscout-mu-timeout"], d, "aeroscout_mu_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout-mu-timeout"], "ObjectWirelessControllerWtpProfileLbs-AeroscoutMuTimeout"); ok {
			if err = d.Set("aeroscout_mu_timeout", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout_mu_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout_mu_timeout: %v", err)
		}
	}

	if err = d.Set("aeroscout_server_ip", flattenObjectWirelessControllerWtpProfileLbsAeroscoutServerIp2edl(o["aeroscout-server-ip"], d, "aeroscout_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout-server-ip"], "ObjectWirelessControllerWtpProfileLbs-AeroscoutServerIp"); ok {
			if err = d.Set("aeroscout_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout_server_ip: %v", err)
		}
	}

	if err = d.Set("aeroscout_server_port", flattenObjectWirelessControllerWtpProfileLbsAeroscoutServerPort2edl(o["aeroscout-server-port"], d, "aeroscout_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout-server-port"], "ObjectWirelessControllerWtpProfileLbs-AeroscoutServerPort"); ok {
			if err = d.Set("aeroscout_server_port", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout_server_port: %v", err)
		}
	}

	if err = d.Set("ekahau_blink_mode", flattenObjectWirelessControllerWtpProfileLbsEkahauBlinkMode2edl(o["ekahau-blink-mode"], d, "ekahau_blink_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ekahau-blink-mode"], "ObjectWirelessControllerWtpProfileLbs-EkahauBlinkMode"); ok {
			if err = d.Set("ekahau_blink_mode", vv); err != nil {
				return fmt.Errorf("Error reading ekahau_blink_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ekahau_blink_mode: %v", err)
		}
	}

	if err = d.Set("ekahau_tag", flattenObjectWirelessControllerWtpProfileLbsEkahauTag2edl(o["ekahau-tag"], d, "ekahau_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ekahau-tag"], "ObjectWirelessControllerWtpProfileLbs-EkahauTag"); ok {
			if err = d.Set("ekahau_tag", vv); err != nil {
				return fmt.Errorf("Error reading ekahau_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ekahau_tag: %v", err)
		}
	}

	if err = d.Set("erc_server_ip", flattenObjectWirelessControllerWtpProfileLbsErcServerIp2edl(o["erc-server-ip"], d, "erc_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["erc-server-ip"], "ObjectWirelessControllerWtpProfileLbs-ErcServerIp"); ok {
			if err = d.Set("erc_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading erc_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading erc_server_ip: %v", err)
		}
	}

	if err = d.Set("erc_server_port", flattenObjectWirelessControllerWtpProfileLbsErcServerPort2edl(o["erc-server-port"], d, "erc_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["erc-server-port"], "ObjectWirelessControllerWtpProfileLbs-ErcServerPort"); ok {
			if err = d.Set("erc_server_port", vv); err != nil {
				return fmt.Errorf("Error reading erc_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading erc_server_port: %v", err)
		}
	}

	if err = d.Set("fortipresence", flattenObjectWirelessControllerWtpProfileLbsFortipresence2edl(o["fortipresence"], d, "fortipresence")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence"], "ObjectWirelessControllerWtpProfileLbs-Fortipresence"); ok {
			if err = d.Set("fortipresence", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence: %v", err)
		}
	}

	if err = d.Set("fortipresence_ble", flattenObjectWirelessControllerWtpProfileLbsFortipresenceBle2edl(o["fortipresence-ble"], d, "fortipresence_ble")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-ble"], "ObjectWirelessControllerWtpProfileLbs-FortipresenceBle"); ok {
			if err = d.Set("fortipresence_ble", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_ble: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_ble: %v", err)
		}
	}

	if err = d.Set("fortipresence_frequency", flattenObjectWirelessControllerWtpProfileLbsFortipresenceFrequency2edl(o["fortipresence-frequency"], d, "fortipresence_frequency")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-frequency"], "ObjectWirelessControllerWtpProfileLbs-FortipresenceFrequency"); ok {
			if err = d.Set("fortipresence_frequency", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_frequency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_frequency: %v", err)
		}
	}

	if err = d.Set("fortipresence_port", flattenObjectWirelessControllerWtpProfileLbsFortipresencePort2edl(o["fortipresence-port"], d, "fortipresence_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-port"], "ObjectWirelessControllerWtpProfileLbs-FortipresencePort"); ok {
			if err = d.Set("fortipresence_port", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_port: %v", err)
		}
	}

	if err = d.Set("fortipresence_project", flattenObjectWirelessControllerWtpProfileLbsFortipresenceProject2edl(o["fortipresence-project"], d, "fortipresence_project")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-project"], "ObjectWirelessControllerWtpProfileLbs-FortipresenceProject"); ok {
			if err = d.Set("fortipresence_project", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_project: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_project: %v", err)
		}
	}

	if err = d.Set("fortipresence_rogue", flattenObjectWirelessControllerWtpProfileLbsFortipresenceRogue2edl(o["fortipresence-rogue"], d, "fortipresence_rogue")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-rogue"], "ObjectWirelessControllerWtpProfileLbs-FortipresenceRogue"); ok {
			if err = d.Set("fortipresence_rogue", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_rogue: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_rogue: %v", err)
		}
	}

	if err = d.Set("fortipresence_server", flattenObjectWirelessControllerWtpProfileLbsFortipresenceServer2edl(o["fortipresence-server"], d, "fortipresence_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-server"], "ObjectWirelessControllerWtpProfileLbs-FortipresenceServer"); ok {
			if err = d.Set("fortipresence_server", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_server: %v", err)
		}
	}

	if err = d.Set("fortipresence_server_addr_type", flattenObjectWirelessControllerWtpProfileLbsFortipresenceServerAddrType2edl(o["fortipresence-server-addr-type"], d, "fortipresence_server_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-server-addr-type"], "ObjectWirelessControllerWtpProfileLbs-FortipresenceServerAddrType"); ok {
			if err = d.Set("fortipresence_server_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_server_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_server_addr_type: %v", err)
		}
	}

	if err = d.Set("fortipresence_server_fqdn", flattenObjectWirelessControllerWtpProfileLbsFortipresenceServerFqdn2edl(o["fortipresence-server-fqdn"], d, "fortipresence_server_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-server-fqdn"], "ObjectWirelessControllerWtpProfileLbs-FortipresenceServerFqdn"); ok {
			if err = d.Set("fortipresence_server_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_server_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_server_fqdn: %v", err)
		}
	}

	if err = d.Set("fortipresence_unassoc", flattenObjectWirelessControllerWtpProfileLbsFortipresenceUnassoc2edl(o["fortipresence-unassoc"], d, "fortipresence_unassoc")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-unassoc"], "ObjectWirelessControllerWtpProfileLbs-FortipresenceUnassoc"); ok {
			if err = d.Set("fortipresence_unassoc", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_unassoc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_unassoc: %v", err)
		}
	}

	if err = d.Set("polestar", flattenObjectWirelessControllerWtpProfileLbsPolestar2edl(o["polestar"], d, "polestar")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar"], "ObjectWirelessControllerWtpProfileLbs-Polestar"); ok {
			if err = d.Set("polestar", vv); err != nil {
				return fmt.Errorf("Error reading polestar: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar: %v", err)
		}
	}

	if err = d.Set("polestar_accumulation_interval", flattenObjectWirelessControllerWtpProfileLbsPolestarAccumulationInterval2edl(o["polestar-accumulation-interval"], d, "polestar_accumulation_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-accumulation-interval"], "ObjectWirelessControllerWtpProfileLbs-PolestarAccumulationInterval"); ok {
			if err = d.Set("polestar_accumulation_interval", vv); err != nil {
				return fmt.Errorf("Error reading polestar_accumulation_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_accumulation_interval: %v", err)
		}
	}

	if err = d.Set("polestar_asset_addrgrp_list", flattenObjectWirelessControllerWtpProfileLbsPolestarAssetAddrgrpList2edl(o["polestar-asset-addrgrp-list"], d, "polestar_asset_addrgrp_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-asset-addrgrp-list"], "ObjectWirelessControllerWtpProfileLbs-PolestarAssetAddrgrpList"); ok {
			if err = d.Set("polestar_asset_addrgrp_list", vv); err != nil {
				return fmt.Errorf("Error reading polestar_asset_addrgrp_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_asset_addrgrp_list: %v", err)
		}
	}

	if err = d.Set("polestar_asset_uuid_list1", flattenObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList12edl(o["polestar-asset-uuid-list1"], d, "polestar_asset_uuid_list1")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-asset-uuid-list1"], "ObjectWirelessControllerWtpProfileLbs-PolestarAssetUuidList1"); ok {
			if err = d.Set("polestar_asset_uuid_list1", vv); err != nil {
				return fmt.Errorf("Error reading polestar_asset_uuid_list1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_asset_uuid_list1: %v", err)
		}
	}

	if err = d.Set("polestar_asset_uuid_list2", flattenObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList22edl(o["polestar-asset-uuid-list2"], d, "polestar_asset_uuid_list2")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-asset-uuid-list2"], "ObjectWirelessControllerWtpProfileLbs-PolestarAssetUuidList2"); ok {
			if err = d.Set("polestar_asset_uuid_list2", vv); err != nil {
				return fmt.Errorf("Error reading polestar_asset_uuid_list2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_asset_uuid_list2: %v", err)
		}
	}

	if err = d.Set("polestar_asset_uuid_list3", flattenObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList32edl(o["polestar-asset-uuid-list3"], d, "polestar_asset_uuid_list3")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-asset-uuid-list3"], "ObjectWirelessControllerWtpProfileLbs-PolestarAssetUuidList3"); ok {
			if err = d.Set("polestar_asset_uuid_list3", vv); err != nil {
				return fmt.Errorf("Error reading polestar_asset_uuid_list3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_asset_uuid_list3: %v", err)
		}
	}

	if err = d.Set("polestar_asset_uuid_list4", flattenObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList42edl(o["polestar-asset-uuid-list4"], d, "polestar_asset_uuid_list4")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-asset-uuid-list4"], "ObjectWirelessControllerWtpProfileLbs-PolestarAssetUuidList4"); ok {
			if err = d.Set("polestar_asset_uuid_list4", vv); err != nil {
				return fmt.Errorf("Error reading polestar_asset_uuid_list4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_asset_uuid_list4: %v", err)
		}
	}

	if err = d.Set("polestar_protocol", flattenObjectWirelessControllerWtpProfileLbsPolestarProtocol2edl(o["polestar-protocol"], d, "polestar_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-protocol"], "ObjectWirelessControllerWtpProfileLbs-PolestarProtocol"); ok {
			if err = d.Set("polestar_protocol", vv); err != nil {
				return fmt.Errorf("Error reading polestar_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_protocol: %v", err)
		}
	}

	if err = d.Set("polestar_reporting_interval", flattenObjectWirelessControllerWtpProfileLbsPolestarReportingInterval2edl(o["polestar-reporting-interval"], d, "polestar_reporting_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-reporting-interval"], "ObjectWirelessControllerWtpProfileLbs-PolestarReportingInterval"); ok {
			if err = d.Set("polestar_reporting_interval", vv); err != nil {
				return fmt.Errorf("Error reading polestar_reporting_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_reporting_interval: %v", err)
		}
	}

	if err = d.Set("polestar_server_fqdn", flattenObjectWirelessControllerWtpProfileLbsPolestarServerFqdn2edl(o["polestar-server-fqdn"], d, "polestar_server_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-server-fqdn"], "ObjectWirelessControllerWtpProfileLbs-PolestarServerFqdn"); ok {
			if err = d.Set("polestar_server_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading polestar_server_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_server_fqdn: %v", err)
		}
	}

	if err = d.Set("polestar_server_path", flattenObjectWirelessControllerWtpProfileLbsPolestarServerPath2edl(o["polestar-server-path"], d, "polestar_server_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-server-path"], "ObjectWirelessControllerWtpProfileLbs-PolestarServerPath"); ok {
			if err = d.Set("polestar_server_path", vv); err != nil {
				return fmt.Errorf("Error reading polestar_server_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_server_path: %v", err)
		}
	}

	if err = d.Set("polestar_server_port", flattenObjectWirelessControllerWtpProfileLbsPolestarServerPort2edl(o["polestar-server-port"], d, "polestar_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-server-port"], "ObjectWirelessControllerWtpProfileLbs-PolestarServerPort"); ok {
			if err = d.Set("polestar_server_port", vv); err != nil {
				return fmt.Errorf("Error reading polestar_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_server_port: %v", err)
		}
	}

	if err = d.Set("polestar_server_token", flattenObjectWirelessControllerWtpProfileLbsPolestarServerToken2edl(o["polestar-server-token"], d, "polestar_server_token")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-server-token"], "ObjectWirelessControllerWtpProfileLbs-PolestarServerToken"); ok {
			if err = d.Set("polestar_server_token", vv); err != nil {
				return fmt.Errorf("Error reading polestar_server_token: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_server_token: %v", err)
		}
	}

	if err = d.Set("station_locate", flattenObjectWirelessControllerWtpProfileLbsStationLocate2edl(o["station-locate"], d, "station_locate")); err != nil {
		if vv, ok := fortiAPIPatch(o["station-locate"], "ObjectWirelessControllerWtpProfileLbs-StationLocate"); ok {
			if err = d.Set("station_locate", vv); err != nil {
				return fmt.Errorf("Error reading station_locate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading station_locate: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerWtpProfileLbsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerWtpProfileLbsAeroscout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsAeroscoutApMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsAeroscoutMmuReport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsAeroscoutMu2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsAeroscoutMuFactor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsAeroscoutMuTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsAeroscoutServerIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsAeroscoutServerPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsEkahauBlinkMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsEkahauTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsErcServerIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsErcServerPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresence2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceBle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceFrequency2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresencePort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceProject2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceRogue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceSecret2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceServerAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceServerFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceUnassoc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsPolestar2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsPolestarAccumulationInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsPolestarAssetAddrgrpList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsPolestarProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsPolestarReportingInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsPolestarServerFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsPolestarServerPath2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsPolestarServerPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsPolestarServerToken2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsStationLocate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerWtpProfileLbs(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aeroscout"); ok || d.HasChange("aeroscout") {
		t, err := expandObjectWirelessControllerWtpProfileLbsAeroscout2edl(d, v, "aeroscout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout"] = t
		}
	}

	if v, ok := d.GetOk("aeroscout_ap_mac"); ok || d.HasChange("aeroscout_ap_mac") {
		t, err := expandObjectWirelessControllerWtpProfileLbsAeroscoutApMac2edl(d, v, "aeroscout_ap_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout-ap-mac"] = t
		}
	}

	if v, ok := d.GetOk("aeroscout_mmu_report"); ok || d.HasChange("aeroscout_mmu_report") {
		t, err := expandObjectWirelessControllerWtpProfileLbsAeroscoutMmuReport2edl(d, v, "aeroscout_mmu_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout-mmu-report"] = t
		}
	}

	if v, ok := d.GetOk("aeroscout_mu"); ok || d.HasChange("aeroscout_mu") {
		t, err := expandObjectWirelessControllerWtpProfileLbsAeroscoutMu2edl(d, v, "aeroscout_mu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout-mu"] = t
		}
	}

	if v, ok := d.GetOk("aeroscout_mu_factor"); ok || d.HasChange("aeroscout_mu_factor") {
		t, err := expandObjectWirelessControllerWtpProfileLbsAeroscoutMuFactor2edl(d, v, "aeroscout_mu_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout-mu-factor"] = t
		}
	}

	if v, ok := d.GetOk("aeroscout_mu_timeout"); ok || d.HasChange("aeroscout_mu_timeout") {
		t, err := expandObjectWirelessControllerWtpProfileLbsAeroscoutMuTimeout2edl(d, v, "aeroscout_mu_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout-mu-timeout"] = t
		}
	}

	if v, ok := d.GetOk("aeroscout_server_ip"); ok || d.HasChange("aeroscout_server_ip") {
		t, err := expandObjectWirelessControllerWtpProfileLbsAeroscoutServerIp2edl(d, v, "aeroscout_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("aeroscout_server_port"); ok || d.HasChange("aeroscout_server_port") {
		t, err := expandObjectWirelessControllerWtpProfileLbsAeroscoutServerPort2edl(d, v, "aeroscout_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout-server-port"] = t
		}
	}

	if v, ok := d.GetOk("ekahau_blink_mode"); ok || d.HasChange("ekahau_blink_mode") {
		t, err := expandObjectWirelessControllerWtpProfileLbsEkahauBlinkMode2edl(d, v, "ekahau_blink_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ekahau-blink-mode"] = t
		}
	}

	if v, ok := d.GetOk("ekahau_tag"); ok || d.HasChange("ekahau_tag") {
		t, err := expandObjectWirelessControllerWtpProfileLbsEkahauTag2edl(d, v, "ekahau_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ekahau-tag"] = t
		}
	}

	if v, ok := d.GetOk("erc_server_ip"); ok || d.HasChange("erc_server_ip") {
		t, err := expandObjectWirelessControllerWtpProfileLbsErcServerIp2edl(d, v, "erc_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["erc-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("erc_server_port"); ok || d.HasChange("erc_server_port") {
		t, err := expandObjectWirelessControllerWtpProfileLbsErcServerPort2edl(d, v, "erc_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["erc-server-port"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence"); ok || d.HasChange("fortipresence") {
		t, err := expandObjectWirelessControllerWtpProfileLbsFortipresence2edl(d, v, "fortipresence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_ble"); ok || d.HasChange("fortipresence_ble") {
		t, err := expandObjectWirelessControllerWtpProfileLbsFortipresenceBle2edl(d, v, "fortipresence_ble")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-ble"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_frequency"); ok || d.HasChange("fortipresence_frequency") {
		t, err := expandObjectWirelessControllerWtpProfileLbsFortipresenceFrequency2edl(d, v, "fortipresence_frequency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-frequency"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_port"); ok || d.HasChange("fortipresence_port") {
		t, err := expandObjectWirelessControllerWtpProfileLbsFortipresencePort2edl(d, v, "fortipresence_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-port"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_project"); ok || d.HasChange("fortipresence_project") {
		t, err := expandObjectWirelessControllerWtpProfileLbsFortipresenceProject2edl(d, v, "fortipresence_project")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-project"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_rogue"); ok || d.HasChange("fortipresence_rogue") {
		t, err := expandObjectWirelessControllerWtpProfileLbsFortipresenceRogue2edl(d, v, "fortipresence_rogue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-rogue"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_secret"); ok || d.HasChange("fortipresence_secret") {
		t, err := expandObjectWirelessControllerWtpProfileLbsFortipresenceSecret2edl(d, v, "fortipresence_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-secret"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_server"); ok || d.HasChange("fortipresence_server") {
		t, err := expandObjectWirelessControllerWtpProfileLbsFortipresenceServer2edl(d, v, "fortipresence_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-server"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_server_addr_type"); ok || d.HasChange("fortipresence_server_addr_type") {
		t, err := expandObjectWirelessControllerWtpProfileLbsFortipresenceServerAddrType2edl(d, v, "fortipresence_server_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-server-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_server_fqdn"); ok || d.HasChange("fortipresence_server_fqdn") {
		t, err := expandObjectWirelessControllerWtpProfileLbsFortipresenceServerFqdn2edl(d, v, "fortipresence_server_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-server-fqdn"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_unassoc"); ok || d.HasChange("fortipresence_unassoc") {
		t, err := expandObjectWirelessControllerWtpProfileLbsFortipresenceUnassoc2edl(d, v, "fortipresence_unassoc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-unassoc"] = t
		}
	}

	if v, ok := d.GetOk("polestar"); ok || d.HasChange("polestar") {
		t, err := expandObjectWirelessControllerWtpProfileLbsPolestar2edl(d, v, "polestar")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar"] = t
		}
	}

	if v, ok := d.GetOk("polestar_accumulation_interval"); ok || d.HasChange("polestar_accumulation_interval") {
		t, err := expandObjectWirelessControllerWtpProfileLbsPolestarAccumulationInterval2edl(d, v, "polestar_accumulation_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-accumulation-interval"] = t
		}
	}

	if v, ok := d.GetOk("polestar_asset_addrgrp_list"); ok || d.HasChange("polestar_asset_addrgrp_list") {
		t, err := expandObjectWirelessControllerWtpProfileLbsPolestarAssetAddrgrpList2edl(d, v, "polestar_asset_addrgrp_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-asset-addrgrp-list"] = t
		}
	}

	if v, ok := d.GetOk("polestar_asset_uuid_list1"); ok || d.HasChange("polestar_asset_uuid_list1") {
		t, err := expandObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList12edl(d, v, "polestar_asset_uuid_list1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-asset-uuid-list1"] = t
		}
	}

	if v, ok := d.GetOk("polestar_asset_uuid_list2"); ok || d.HasChange("polestar_asset_uuid_list2") {
		t, err := expandObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList22edl(d, v, "polestar_asset_uuid_list2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-asset-uuid-list2"] = t
		}
	}

	if v, ok := d.GetOk("polestar_asset_uuid_list3"); ok || d.HasChange("polestar_asset_uuid_list3") {
		t, err := expandObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList32edl(d, v, "polestar_asset_uuid_list3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-asset-uuid-list3"] = t
		}
	}

	if v, ok := d.GetOk("polestar_asset_uuid_list4"); ok || d.HasChange("polestar_asset_uuid_list4") {
		t, err := expandObjectWirelessControllerWtpProfileLbsPolestarAssetUuidList42edl(d, v, "polestar_asset_uuid_list4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-asset-uuid-list4"] = t
		}
	}

	if v, ok := d.GetOk("polestar_protocol"); ok || d.HasChange("polestar_protocol") {
		t, err := expandObjectWirelessControllerWtpProfileLbsPolestarProtocol2edl(d, v, "polestar_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-protocol"] = t
		}
	}

	if v, ok := d.GetOk("polestar_reporting_interval"); ok || d.HasChange("polestar_reporting_interval") {
		t, err := expandObjectWirelessControllerWtpProfileLbsPolestarReportingInterval2edl(d, v, "polestar_reporting_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-reporting-interval"] = t
		}
	}

	if v, ok := d.GetOk("polestar_server_fqdn"); ok || d.HasChange("polestar_server_fqdn") {
		t, err := expandObjectWirelessControllerWtpProfileLbsPolestarServerFqdn2edl(d, v, "polestar_server_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-server-fqdn"] = t
		}
	}

	if v, ok := d.GetOk("polestar_server_path"); ok || d.HasChange("polestar_server_path") {
		t, err := expandObjectWirelessControllerWtpProfileLbsPolestarServerPath2edl(d, v, "polestar_server_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-server-path"] = t
		}
	}

	if v, ok := d.GetOk("polestar_server_port"); ok || d.HasChange("polestar_server_port") {
		t, err := expandObjectWirelessControllerWtpProfileLbsPolestarServerPort2edl(d, v, "polestar_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-server-port"] = t
		}
	}

	if v, ok := d.GetOk("polestar_server_token"); ok || d.HasChange("polestar_server_token") {
		t, err := expandObjectWirelessControllerWtpProfileLbsPolestarServerToken2edl(d, v, "polestar_server_token")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-server-token"] = t
		}
	}

	if v, ok := d.GetOk("station_locate"); ok || d.HasChange("station_locate") {
		t, err := expandObjectWirelessControllerWtpProfileLbsStationLocate2edl(d, v, "station_locate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["station-locate"] = t
		}
	}

	return &obj, nil
}
