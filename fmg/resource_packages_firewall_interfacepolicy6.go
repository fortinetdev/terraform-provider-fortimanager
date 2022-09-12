// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 interface policies.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourcePackagesFirewallInterfacePolicy6() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesFirewallInterfacePolicy6Create,
		Read:   resourcePackagesFirewallInterfacePolicy6Read,
		Update: resourcePackagesFirewallInterfacePolicy6Update,
		Delete: resourcePackagesFirewallInterfacePolicy6Delete,

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
			"pkg": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_list_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"av_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dlp_sensor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dlp_sensor_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dsri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"emailfilter_profile": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"emailfilter_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"ips_sensor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ips_sensor_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service6": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"spamfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"spamfilter_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webfilter_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourcePackagesFirewallInterfacePolicy6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	obj, err := getObjectPackagesFirewallInterfacePolicy6(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallInterfacePolicy6 resource while getting object: %v", err)
	}

	_, err = c.CreatePackagesFirewallInterfacePolicy6(obj, adomv, paralist)

	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallInterfacePolicy6 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallInterfacePolicy6Read(d, m)
}

func resourcePackagesFirewallInterfacePolicy6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	obj, err := getObjectPackagesFirewallInterfacePolicy6(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallInterfacePolicy6 resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesFirewallInterfacePolicy6(obj, adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallInterfacePolicy6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallInterfacePolicy6Read(d, m)
}

func resourcePackagesFirewallInterfacePolicy6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	err = c.DeletePackagesFirewallInterfacePolicy6(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesFirewallInterfacePolicy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesFirewallInterfacePolicy6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	if pkg == "" {
		pkg = importOptionChecking(m.(*FortiClient).Cfg, "pkg")
		if err = d.Set("pkg", pkg); err != nil {
			return fmt.Errorf("Error set params pkg: %v", err)
		}
	}
	var paralist []string
	paralist = append(paralist, pkg)

	o, err := c.ReadPackagesFirewallInterfacePolicy6(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallInterfacePolicy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesFirewallInterfacePolicy6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallInterfacePolicy6 resource from API: %v", err)
	}
	return nil
}

func flattenPackagesFirewallInterfacePolicy6AddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6ApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6ApplicationListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6AvProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6AvProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6DlpSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6DlpSensorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6Dsri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6Dstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallInterfacePolicy6EmailfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallInterfacePolicy6EmailfilterProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallInterfacePolicy6IpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6IpsSensorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6Label(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6Logtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6Policyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6ScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6Service6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallInterfacePolicy6SpamfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6SpamfilterProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6Srcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallInterfacePolicy6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6WebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallInterfacePolicy6WebfilterProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesFirewallInterfacePolicy6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("address_type", flattenPackagesFirewallInterfacePolicy6AddressType(o["address-type"], d, "address_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-type"], "PackagesFirewallInterfacePolicy6-AddressType"); ok {
			if err = d.Set("address_type", vv); err != nil {
				return fmt.Errorf("Error reading address_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_type: %v", err)
		}
	}

	if err = d.Set("application_list", flattenPackagesFirewallInterfacePolicy6ApplicationList(o["application-list"], d, "application_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list"], "PackagesFirewallInterfacePolicy6-ApplicationList"); ok {
			if err = d.Set("application_list", vv); err != nil {
				return fmt.Errorf("Error reading application_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("application_list_status", flattenPackagesFirewallInterfacePolicy6ApplicationListStatus(o["application-list-status"], d, "application_list_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list-status"], "PackagesFirewallInterfacePolicy6-ApplicationListStatus"); ok {
			if err = d.Set("application_list_status", vv); err != nil {
				return fmt.Errorf("Error reading application_list_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list_status: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenPackagesFirewallInterfacePolicy6AvProfile(o["av-profile"], d, "av_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile"], "PackagesFirewallInterfacePolicy6-AvProfile"); ok {
			if err = d.Set("av_profile", vv); err != nil {
				return fmt.Errorf("Error reading av_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("av_profile_status", flattenPackagesFirewallInterfacePolicy6AvProfileStatus(o["av-profile-status"], d, "av_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile-status"], "PackagesFirewallInterfacePolicy6-AvProfileStatus"); ok {
			if err = d.Set("av_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading av_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile_status: %v", err)
		}
	}

	if err = d.Set("comments", flattenPackagesFirewallInterfacePolicy6Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "PackagesFirewallInterfacePolicy6-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenPackagesFirewallInterfacePolicy6DlpSensor(o["dlp-sensor"], d, "dlp_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor"], "PackagesFirewallInterfacePolicy6-DlpSensor"); ok {
			if err = d.Set("dlp_sensor", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("dlp_sensor_status", flattenPackagesFirewallInterfacePolicy6DlpSensorStatus(o["dlp-sensor-status"], d, "dlp_sensor_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor-status"], "PackagesFirewallInterfacePolicy6-DlpSensorStatus"); ok {
			if err = d.Set("dlp_sensor_status", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor_status: %v", err)
		}
	}

	if err = d.Set("dsri", flattenPackagesFirewallInterfacePolicy6Dsri(o["dsri"], d, "dsri")); err != nil {
		if vv, ok := fortiAPIPatch(o["dsri"], "PackagesFirewallInterfacePolicy6-Dsri"); ok {
			if err = d.Set("dsri", vv); err != nil {
				return fmt.Errorf("Error reading dsri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dsri: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenPackagesFirewallInterfacePolicy6Dstaddr6(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "PackagesFirewallInterfacePolicy6-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenPackagesFirewallInterfacePolicy6EmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["emailfilter-profile"], "PackagesFirewallInterfacePolicy6-EmailfilterProfile"); ok {
			if err = d.Set("emailfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading emailfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile_status", flattenPackagesFirewallInterfacePolicy6EmailfilterProfileStatus(o["emailfilter-profile-status"], d, "emailfilter_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["emailfilter-profile-status"], "PackagesFirewallInterfacePolicy6-EmailfilterProfileStatus"); ok {
			if err = d.Set("emailfilter_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading emailfilter_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emailfilter_profile_status: %v", err)
		}
	}

	if err = d.Set("interface", flattenPackagesFirewallInterfacePolicy6Interface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "PackagesFirewallInterfacePolicy6-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenPackagesFirewallInterfacePolicy6IpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor"], "PackagesFirewallInterfacePolicy6-IpsSensor"); ok {
			if err = d.Set("ips_sensor", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("ips_sensor_status", flattenPackagesFirewallInterfacePolicy6IpsSensorStatus(o["ips-sensor-status"], d, "ips_sensor_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor-status"], "PackagesFirewallInterfacePolicy6-IpsSensorStatus"); ok {
			if err = d.Set("ips_sensor_status", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor_status: %v", err)
		}
	}

	if err = d.Set("label", flattenPackagesFirewallInterfacePolicy6Label(o["label"], d, "label")); err != nil {
		if vv, ok := fortiAPIPatch(o["label"], "PackagesFirewallInterfacePolicy6-Label"); ok {
			if err = d.Set("label", vv); err != nil {
				return fmt.Errorf("Error reading label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenPackagesFirewallInterfacePolicy6Logtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic"], "PackagesFirewallInterfacePolicy6-Logtraffic"); ok {
			if err = d.Set("logtraffic", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("policyid", flattenPackagesFirewallInterfacePolicy6Policyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "PackagesFirewallInterfacePolicy6-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", flattenPackagesFirewallInterfacePolicy6ScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-botnet-connections"], "PackagesFirewallInterfacePolicy6-ScanBotnetConnections"); ok {
			if err = d.Set("scan_botnet_connections", vv); err != nil {
				return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	if err = d.Set("service6", flattenPackagesFirewallInterfacePolicy6Service6(o["service6"], d, "service6")); err != nil {
		if vv, ok := fortiAPIPatch(o["service6"], "PackagesFirewallInterfacePolicy6-Service6"); ok {
			if err = d.Set("service6", vv); err != nil {
				return fmt.Errorf("Error reading service6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service6: %v", err)
		}
	}

	if err = d.Set("spamfilter_profile", flattenPackagesFirewallInterfacePolicy6SpamfilterProfile(o["spamfilter-profile"], d, "spamfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["spamfilter-profile"], "PackagesFirewallInterfacePolicy6-SpamfilterProfile"); ok {
			if err = d.Set("spamfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading spamfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spamfilter_profile: %v", err)
		}
	}

	if err = d.Set("spamfilter_profile_status", flattenPackagesFirewallInterfacePolicy6SpamfilterProfileStatus(o["spamfilter-profile-status"], d, "spamfilter_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["spamfilter-profile-status"], "PackagesFirewallInterfacePolicy6-SpamfilterProfileStatus"); ok {
			if err = d.Set("spamfilter_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading spamfilter_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spamfilter_profile_status: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenPackagesFirewallInterfacePolicy6Srcaddr6(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "PackagesFirewallInterfacePolicy6-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesFirewallInterfacePolicy6Status(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesFirewallInterfacePolicy6-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("uuid", flattenPackagesFirewallInterfacePolicy6Uuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "PackagesFirewallInterfacePolicy6-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenPackagesFirewallInterfacePolicy6WebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile"], "PackagesFirewallInterfacePolicy6-WebfilterProfile"); ok {
			if err = d.Set("webfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile_status", flattenPackagesFirewallInterfacePolicy6WebfilterProfileStatus(o["webfilter-profile-status"], d, "webfilter_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile-status"], "PackagesFirewallInterfacePolicy6-WebfilterProfileStatus"); ok {
			if err = d.Set("webfilter_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile_status: %v", err)
		}
	}

	return nil
}

func flattenPackagesFirewallInterfacePolicy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesFirewallInterfacePolicy6AddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6ApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6ApplicationListStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6AvProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6AvProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6DlpSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6DlpSensorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6Dsri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6Dstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallInterfacePolicy6EmailfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallInterfacePolicy6EmailfilterProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallInterfacePolicy6IpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6IpsSensorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6Label(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6Logtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6Policyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6ScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6Service6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallInterfacePolicy6SpamfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6SpamfilterProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6Srcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallInterfacePolicy6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6WebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallInterfacePolicy6WebfilterProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesFirewallInterfacePolicy6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address_type"); ok || d.HasChange("address_type") {
		t, err := expandPackagesFirewallInterfacePolicy6AddressType(d, v, "address_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-type"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok || d.HasChange("application_list") {
		t, err := expandPackagesFirewallInterfacePolicy6ApplicationList(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("application_list_status"); ok || d.HasChange("application_list_status") {
		t, err := expandPackagesFirewallInterfacePolicy6ApplicationListStatus(d, v, "application_list_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list-status"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok || d.HasChange("av_profile") {
		t, err := expandPackagesFirewallInterfacePolicy6AvProfile(d, v, "av_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("av_profile_status"); ok || d.HasChange("av_profile_status") {
		t, err := expandPackagesFirewallInterfacePolicy6AvProfileStatus(d, v, "av_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandPackagesFirewallInterfacePolicy6Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok || d.HasChange("dlp_sensor") {
		t, err := expandPackagesFirewallInterfacePolicy6DlpSensor(d, v, "dlp_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor_status"); ok || d.HasChange("dlp_sensor_status") {
		t, err := expandPackagesFirewallInterfacePolicy6DlpSensorStatus(d, v, "dlp_sensor_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor-status"] = t
		}
	}

	if v, ok := d.GetOk("dsri"); ok || d.HasChange("dsri") {
		t, err := expandPackagesFirewallInterfacePolicy6Dsri(d, v, "dsri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsri"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandPackagesFirewallInterfacePolicy6Dstaddr6(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok || d.HasChange("emailfilter_profile") {
		t, err := expandPackagesFirewallInterfacePolicy6EmailfilterProfile(d, v, "emailfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile_status"); ok || d.HasChange("emailfilter_profile_status") {
		t, err := expandPackagesFirewallInterfacePolicy6EmailfilterProfileStatus(d, v, "emailfilter_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandPackagesFirewallInterfacePolicy6Interface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok || d.HasChange("ips_sensor") {
		t, err := expandPackagesFirewallInterfacePolicy6IpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor_status"); ok || d.HasChange("ips_sensor_status") {
		t, err := expandPackagesFirewallInterfacePolicy6IpsSensorStatus(d, v, "ips_sensor_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor-status"] = t
		}
	}

	if v, ok := d.GetOk("label"); ok || d.HasChange("label") {
		t, err := expandPackagesFirewallInterfacePolicy6Label(d, v, "label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["label"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok || d.HasChange("logtraffic") {
		t, err := expandPackagesFirewallInterfacePolicy6Logtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandPackagesFirewallInterfacePolicy6Policyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok || d.HasChange("scan_botnet_connections") {
		t, err := expandPackagesFirewallInterfacePolicy6ScanBotnetConnections(d, v, "scan_botnet_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	}

	if v, ok := d.GetOk("service6"); ok || d.HasChange("service6") {
		t, err := expandPackagesFirewallInterfacePolicy6Service6(d, v, "service6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service6"] = t
		}
	}

	if v, ok := d.GetOk("spamfilter_profile"); ok || d.HasChange("spamfilter_profile") {
		t, err := expandPackagesFirewallInterfacePolicy6SpamfilterProfile(d, v, "spamfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spamfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("spamfilter_profile_status"); ok || d.HasChange("spamfilter_profile_status") {
		t, err := expandPackagesFirewallInterfacePolicy6SpamfilterProfileStatus(d, v, "spamfilter_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spamfilter-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandPackagesFirewallInterfacePolicy6Srcaddr6(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandPackagesFirewallInterfacePolicy6Status(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandPackagesFirewallInterfacePolicy6Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok || d.HasChange("webfilter_profile") {
		t, err := expandPackagesFirewallInterfacePolicy6WebfilterProfile(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile_status"); ok || d.HasChange("webfilter_profile_status") {
		t, err := expandPackagesFirewallInterfacePolicy6WebfilterProfileStatus(d, v, "webfilter_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile-status"] = t
		}
	}

	return &obj, nil
}
