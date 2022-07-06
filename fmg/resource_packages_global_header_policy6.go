// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 policies.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourcePackagesGlobalHeaderPolicy6() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesGlobalHeaderPolicy6Create,
		Read:   resourcePackagesGlobalHeaderPolicy6Read,
		Update: resourcePackagesGlobalHeaderPolicy6Update,
		Delete: resourcePackagesGlobalHeaderPolicy6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"pkg": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anti_replay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"application_charts": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"application_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"casi_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cgn_log_server_grp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cifs_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_log_fields": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"decrypted_traffic_mirror": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deep_inspection_options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_detection_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"devices": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode_rev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dlp_sensor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dsri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_profile_access": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dynamic_profile_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email_collection_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"emailfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_filter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firewall_session_dirty": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fixedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsae": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsso_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"global_label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_policy_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icap_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"identity_based": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"identity_from": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inspection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ippool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_sensor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mms_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"natinbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"natoutbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"np_accelation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"np_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_ip_shaper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"poolname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replacemsg_override_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_deny_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spamfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_filter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_policy_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_mirror": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_mirror_intf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_ccert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_mss_receiver": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_mss_sender": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_session_without_syn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout_send_rst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_shaper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_shaper_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"users": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utm_inspection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_cos_fwd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan_cos_rev": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voip_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpntunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"waf_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webcache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webcache_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webproxy_forward_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webproxy_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourcePackagesGlobalHeaderPolicy6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	obj, err := getObjectPackagesGlobalHeaderPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesGlobalHeaderPolicy6 resource while getting object: %v", err)
	}

	v, err := c.CreatePackagesGlobalHeaderPolicy6(obj, adomv, paralist)

	if err != nil {
		return fmt.Errorf("Error creating PackagesGlobalHeaderPolicy6 resource: %v", err)
	}

	if v != nil && v["policyid"] != nil {
		if vidn, ok := v["policyid"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourcePackagesGlobalHeaderPolicy6Read(d, m)
		} else {
			return fmt.Errorf("Error creating PackagesGlobalHeaderPolicy6 resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesGlobalHeaderPolicy6Read(d, m)
}

func resourcePackagesGlobalHeaderPolicy6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	obj, err := getObjectPackagesGlobalHeaderPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesGlobalHeaderPolicy6 resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesGlobalHeaderPolicy6(obj, adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error updating PackagesGlobalHeaderPolicy6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesGlobalHeaderPolicy6Read(d, m)
}

func resourcePackagesGlobalHeaderPolicy6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	err = c.DeletePackagesGlobalHeaderPolicy6(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesGlobalHeaderPolicy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesGlobalHeaderPolicy6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	pkg := d.Get("pkg").(string)
	if pkg == "" {
		pkg = importOptionChecking(m.(*FortiClient).Cfg, "pkg")
		if err = d.Set("pkg", pkg); err != nil {
			return fmt.Errorf("Error set params pkg: %v", err)
		}
	}
	var paralist []string
	paralist = append(paralist, pkg)

	o, err := c.ReadPackagesGlobalHeaderPolicy6(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error reading PackagesGlobalHeaderPolicy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesGlobalHeaderPolicy6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesGlobalHeaderPolicy6 resource from API: %v", err)
	}
	return nil
}

func flattenPackagesGlobalHeaderPolicy6Action(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6AntiReplay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6AppCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6AppGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Application(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenPackagesGlobalHeaderPolicy6ApplicationCharts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesGlobalHeaderPolicy6ApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6AutoAsicOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6AvProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6CasiProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6CgnLogServerGrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6CifsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6CustomLogFields(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6DecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6DeepInspectionOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6DeviceDetectionPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Devices(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6DiffservForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6DiffservReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6DiffservcodeForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6DiffservcodeRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6DlpSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6DnsfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6DscpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6DscpNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6DscpValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Dsri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Dstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6DstaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Dstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6DynamicProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6DynamicProfileAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesGlobalHeaderPolicy6DynamicProfileGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6EmailCollectionPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6EmailfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6FileFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6FirewallSessionDirty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Fixedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Fsae(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6FssoGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6GlobalLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Groups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6HttpPolicyRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6IcapProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6IdentityBased(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6IdentityFrom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Inbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6InspectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Ippool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6IpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Label(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Logtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6LogtrafficStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6MmsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Nat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Natinbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Natoutbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6NpAccelation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6NpAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Outbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6PerIpShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6PolicyOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Policyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Poolname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6ProfileGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6ProfileProtocolOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6ProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6ReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6ReplacemsgOverrideGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Rsso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Schedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6SendDenyPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Service(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6ServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6SessionTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6SpamfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Srcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6SrcaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Srcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6SshFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6SshPolicyRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6SslMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6SslMirrorIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6SslSshProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6SslvpnAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6SslvpnCcert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6SslvpnCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Tags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6TcpMssReceiver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6TcpMssSender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6TcpSessionWithoutSyn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6TimeoutSendRst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Tos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6TosMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6TosNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6TrafficShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6TrafficShaperReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6UrlCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Users(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6UtmInspectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6UtmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6VlanCosFwd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6VlanCosRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6VlanFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6VoipProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Vpntunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6WafProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6Webcache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6WebcacheHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6WebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6WebproxyForwardServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderPolicy6WebproxyProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesGlobalHeaderPolicy6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenPackagesGlobalHeaderPolicy6Action(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "PackagesGlobalHeaderPolicy6-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("anti_replay", flattenPackagesGlobalHeaderPolicy6AntiReplay(o["anti-replay"], d, "anti_replay")); err != nil {
		if vv, ok := fortiAPIPatch(o["anti-replay"], "PackagesGlobalHeaderPolicy6-AntiReplay"); ok {
			if err = d.Set("anti_replay", vv); err != nil {
				return fmt.Errorf("Error reading anti_replay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anti_replay: %v", err)
		}
	}

	if err = d.Set("app_category", flattenPackagesGlobalHeaderPolicy6AppCategory(o["app-category"], d, "app_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-category"], "PackagesGlobalHeaderPolicy6-AppCategory"); ok {
			if err = d.Set("app_category", vv); err != nil {
				return fmt.Errorf("Error reading app_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_category: %v", err)
		}
	}

	if err = d.Set("app_group", flattenPackagesGlobalHeaderPolicy6AppGroup(o["app-group"], d, "app_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-group"], "PackagesGlobalHeaderPolicy6-AppGroup"); ok {
			if err = d.Set("app_group", vv); err != nil {
				return fmt.Errorf("Error reading app_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_group: %v", err)
		}
	}

	if err = d.Set("application", flattenPackagesGlobalHeaderPolicy6Application(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "PackagesGlobalHeaderPolicy6-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("application_charts", flattenPackagesGlobalHeaderPolicy6ApplicationCharts(o["application-charts"], d, "application_charts")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-charts"], "PackagesGlobalHeaderPolicy6-ApplicationCharts"); ok {
			if err = d.Set("application_charts", vv); err != nil {
				return fmt.Errorf("Error reading application_charts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_charts: %v", err)
		}
	}

	if err = d.Set("application_list", flattenPackagesGlobalHeaderPolicy6ApplicationList(o["application-list"], d, "application_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list"], "PackagesGlobalHeaderPolicy6-ApplicationList"); ok {
			if err = d.Set("application_list", vv); err != nil {
				return fmt.Errorf("Error reading application_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("auto_asic_offload", flattenPackagesGlobalHeaderPolicy6AutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-asic-offload"], "PackagesGlobalHeaderPolicy6-AutoAsicOffload"); ok {
			if err = d.Set("auto_asic_offload", vv); err != nil {
				return fmt.Errorf("Error reading auto_asic_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenPackagesGlobalHeaderPolicy6AvProfile(o["av-profile"], d, "av_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile"], "PackagesGlobalHeaderPolicy6-AvProfile"); ok {
			if err = d.Set("av_profile", vv); err != nil {
				return fmt.Errorf("Error reading av_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("casi_profile", flattenPackagesGlobalHeaderPolicy6CasiProfile(o["casi-profile"], d, "casi_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["casi-profile"], "PackagesGlobalHeaderPolicy6-CasiProfile"); ok {
			if err = d.Set("casi_profile", vv); err != nil {
				return fmt.Errorf("Error reading casi_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading casi_profile: %v", err)
		}
	}

	if err = d.Set("cgn_log_server_grp", flattenPackagesGlobalHeaderPolicy6CgnLogServerGrp(o["cgn-log-server-grp"], d, "cgn_log_server_grp")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-log-server-grp"], "PackagesGlobalHeaderPolicy6-CgnLogServerGrp"); ok {
			if err = d.Set("cgn_log_server_grp", vv); err != nil {
				return fmt.Errorf("Error reading cgn_log_server_grp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_log_server_grp: %v", err)
		}
	}

	if err = d.Set("cifs_profile", flattenPackagesGlobalHeaderPolicy6CifsProfile(o["cifs-profile"], d, "cifs_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["cifs-profile"], "PackagesGlobalHeaderPolicy6-CifsProfile"); ok {
			if err = d.Set("cifs_profile", vv); err != nil {
				return fmt.Errorf("Error reading cifs_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cifs_profile: %v", err)
		}
	}

	if err = d.Set("comments", flattenPackagesGlobalHeaderPolicy6Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "PackagesGlobalHeaderPolicy6-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("custom_log_fields", flattenPackagesGlobalHeaderPolicy6CustomLogFields(o["custom-log-fields"], d, "custom_log_fields")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-log-fields"], "PackagesGlobalHeaderPolicy6-CustomLogFields"); ok {
			if err = d.Set("custom_log_fields", vv); err != nil {
				return fmt.Errorf("Error reading custom_log_fields: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_log_fields: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenPackagesGlobalHeaderPolicy6DecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror")); err != nil {
		if vv, ok := fortiAPIPatch(o["decrypted-traffic-mirror"], "PackagesGlobalHeaderPolicy6-DecryptedTrafficMirror"); ok {
			if err = d.Set("decrypted_traffic_mirror", vv); err != nil {
				return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if err = d.Set("deep_inspection_options", flattenPackagesGlobalHeaderPolicy6DeepInspectionOptions(o["deep-inspection-options"], d, "deep_inspection_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["deep-inspection-options"], "PackagesGlobalHeaderPolicy6-DeepInspectionOptions"); ok {
			if err = d.Set("deep_inspection_options", vv); err != nil {
				return fmt.Errorf("Error reading deep_inspection_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deep_inspection_options: %v", err)
		}
	}

	if err = d.Set("device_detection_portal", flattenPackagesGlobalHeaderPolicy6DeviceDetectionPortal(o["device-detection-portal"], d, "device_detection_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-detection-portal"], "PackagesGlobalHeaderPolicy6-DeviceDetectionPortal"); ok {
			if err = d.Set("device_detection_portal", vv); err != nil {
				return fmt.Errorf("Error reading device_detection_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_detection_portal: %v", err)
		}
	}

	if err = d.Set("devices", flattenPackagesGlobalHeaderPolicy6Devices(o["devices"], d, "devices")); err != nil {
		if vv, ok := fortiAPIPatch(o["devices"], "PackagesGlobalHeaderPolicy6-Devices"); ok {
			if err = d.Set("devices", vv); err != nil {
				return fmt.Errorf("Error reading devices: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devices: %v", err)
		}
	}

	if err = d.Set("diffserv_forward", flattenPackagesGlobalHeaderPolicy6DiffservForward(o["diffserv-forward"], d, "diffserv_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-forward"], "PackagesGlobalHeaderPolicy6-DiffservForward"); ok {
			if err = d.Set("diffserv_forward", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_forward: %v", err)
		}
	}

	if err = d.Set("diffserv_reverse", flattenPackagesGlobalHeaderPolicy6DiffservReverse(o["diffserv-reverse"], d, "diffserv_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-reverse"], "PackagesGlobalHeaderPolicy6-DiffservReverse"); ok {
			if err = d.Set("diffserv_reverse", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_reverse: %v", err)
		}
	}

	if err = d.Set("diffservcode_forward", flattenPackagesGlobalHeaderPolicy6DiffservcodeForward(o["diffservcode-forward"], d, "diffservcode_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode-forward"], "PackagesGlobalHeaderPolicy6-DiffservcodeForward"); ok {
			if err = d.Set("diffservcode_forward", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode_forward: %v", err)
		}
	}

	if err = d.Set("diffservcode_rev", flattenPackagesGlobalHeaderPolicy6DiffservcodeRev(o["diffservcode-rev"], d, "diffservcode_rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode-rev"], "PackagesGlobalHeaderPolicy6-DiffservcodeRev"); ok {
			if err = d.Set("diffservcode_rev", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode_rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode_rev: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenPackagesGlobalHeaderPolicy6DlpSensor(o["dlp-sensor"], d, "dlp_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor"], "PackagesGlobalHeaderPolicy6-DlpSensor"); ok {
			if err = d.Set("dlp_sensor", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("dnsfilter_profile", flattenPackagesGlobalHeaderPolicy6DnsfilterProfile(o["dnsfilter-profile"], d, "dnsfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dnsfilter-profile"], "PackagesGlobalHeaderPolicy6-DnsfilterProfile"); ok {
			if err = d.Set("dnsfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
		}
	}

	if err = d.Set("dscp_match", flattenPackagesGlobalHeaderPolicy6DscpMatch(o["dscp-match"], d, "dscp_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-match"], "PackagesGlobalHeaderPolicy6-DscpMatch"); ok {
			if err = d.Set("dscp_match", vv); err != nil {
				return fmt.Errorf("Error reading dscp_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_match: %v", err)
		}
	}

	if err = d.Set("dscp_negate", flattenPackagesGlobalHeaderPolicy6DscpNegate(o["dscp-negate"], d, "dscp_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-negate"], "PackagesGlobalHeaderPolicy6-DscpNegate"); ok {
			if err = d.Set("dscp_negate", vv); err != nil {
				return fmt.Errorf("Error reading dscp_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_negate: %v", err)
		}
	}

	if err = d.Set("dscp_value", flattenPackagesGlobalHeaderPolicy6DscpValue(o["dscp-value"], d, "dscp_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-value"], "PackagesGlobalHeaderPolicy6-DscpValue"); ok {
			if err = d.Set("dscp_value", vv); err != nil {
				return fmt.Errorf("Error reading dscp_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_value: %v", err)
		}
	}

	if err = d.Set("dsri", flattenPackagesGlobalHeaderPolicy6Dsri(o["dsri"], d, "dsri")); err != nil {
		if vv, ok := fortiAPIPatch(o["dsri"], "PackagesGlobalHeaderPolicy6-Dsri"); ok {
			if err = d.Set("dsri", vv); err != nil {
				return fmt.Errorf("Error reading dsri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dsri: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenPackagesGlobalHeaderPolicy6Dstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "PackagesGlobalHeaderPolicy6-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr_negate", flattenPackagesGlobalHeaderPolicy6DstaddrNegate(o["dstaddr-negate"], d, "dstaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr-negate"], "PackagesGlobalHeaderPolicy6-DstaddrNegate"); ok {
			if err = d.Set("dstaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr_negate: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenPackagesGlobalHeaderPolicy6Dstintf(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "PackagesGlobalHeaderPolicy6-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("dynamic_profile", flattenPackagesGlobalHeaderPolicy6DynamicProfile(o["dynamic-profile"], d, "dynamic_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-profile"], "PackagesGlobalHeaderPolicy6-DynamicProfile"); ok {
			if err = d.Set("dynamic_profile", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_profile: %v", err)
		}
	}

	if err = d.Set("dynamic_profile_access", flattenPackagesGlobalHeaderPolicy6DynamicProfileAccess(o["dynamic-profile-access"], d, "dynamic_profile_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-profile-access"], "PackagesGlobalHeaderPolicy6-DynamicProfileAccess"); ok {
			if err = d.Set("dynamic_profile_access", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_profile_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_profile_access: %v", err)
		}
	}

	if err = d.Set("dynamic_profile_group", flattenPackagesGlobalHeaderPolicy6DynamicProfileGroup(o["dynamic-profile-group"], d, "dynamic_profile_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-profile-group"], "PackagesGlobalHeaderPolicy6-DynamicProfileGroup"); ok {
			if err = d.Set("dynamic_profile_group", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_profile_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_profile_group: %v", err)
		}
	}

	if err = d.Set("email_collection_portal", flattenPackagesGlobalHeaderPolicy6EmailCollectionPortal(o["email-collection-portal"], d, "email_collection_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-collection-portal"], "PackagesGlobalHeaderPolicy6-EmailCollectionPortal"); ok {
			if err = d.Set("email_collection_portal", vv); err != nil {
				return fmt.Errorf("Error reading email_collection_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_collection_portal: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenPackagesGlobalHeaderPolicy6EmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["emailfilter-profile"], "PackagesGlobalHeaderPolicy6-EmailfilterProfile"); ok {
			if err = d.Set("emailfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading emailfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("file_filter_profile", flattenPackagesGlobalHeaderPolicy6FileFilterProfile(o["file-filter-profile"], d, "file_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-filter-profile"], "PackagesGlobalHeaderPolicy6-FileFilterProfile"); ok {
			if err = d.Set("file_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading file_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_filter_profile: %v", err)
		}
	}

	if err = d.Set("firewall_session_dirty", flattenPackagesGlobalHeaderPolicy6FirewallSessionDirty(o["firewall-session-dirty"], d, "firewall_session_dirty")); err != nil {
		if vv, ok := fortiAPIPatch(o["firewall-session-dirty"], "PackagesGlobalHeaderPolicy6-FirewallSessionDirty"); ok {
			if err = d.Set("firewall_session_dirty", vv); err != nil {
				return fmt.Errorf("Error reading firewall_session_dirty: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firewall_session_dirty: %v", err)
		}
	}

	if err = d.Set("fixedport", flattenPackagesGlobalHeaderPolicy6Fixedport(o["fixedport"], d, "fixedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["fixedport"], "PackagesGlobalHeaderPolicy6-Fixedport"); ok {
			if err = d.Set("fixedport", vv); err != nil {
				return fmt.Errorf("Error reading fixedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fixedport: %v", err)
		}
	}

	if err = d.Set("fsae", flattenPackagesGlobalHeaderPolicy6Fsae(o["fsae"], d, "fsae")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsae"], "PackagesGlobalHeaderPolicy6-Fsae"); ok {
			if err = d.Set("fsae", vv); err != nil {
				return fmt.Errorf("Error reading fsae: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsae: %v", err)
		}
	}

	if err = d.Set("fsso_groups", flattenPackagesGlobalHeaderPolicy6FssoGroups(o["fsso-groups"], d, "fsso_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-groups"], "PackagesGlobalHeaderPolicy6-FssoGroups"); ok {
			if err = d.Set("fsso_groups", vv); err != nil {
				return fmt.Errorf("Error reading fsso_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_groups: %v", err)
		}
	}

	if err = d.Set("global_label", flattenPackagesGlobalHeaderPolicy6GlobalLabel(o["global-label"], d, "global_label")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-label"], "PackagesGlobalHeaderPolicy6-GlobalLabel"); ok {
			if err = d.Set("global_label", vv); err != nil {
				return fmt.Errorf("Error reading global_label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_label: %v", err)
		}
	}

	if err = d.Set("groups", flattenPackagesGlobalHeaderPolicy6Groups(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "PackagesGlobalHeaderPolicy6-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("http_policy_redirect", flattenPackagesGlobalHeaderPolicy6HttpPolicyRedirect(o["http-policy-redirect"], d, "http_policy_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-policy-redirect"], "PackagesGlobalHeaderPolicy6-HttpPolicyRedirect"); ok {
			if err = d.Set("http_policy_redirect", vv); err != nil {
				return fmt.Errorf("Error reading http_policy_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_policy_redirect: %v", err)
		}
	}

	if err = d.Set("icap_profile", flattenPackagesGlobalHeaderPolicy6IcapProfile(o["icap-profile"], d, "icap_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["icap-profile"], "PackagesGlobalHeaderPolicy6-IcapProfile"); ok {
			if err = d.Set("icap_profile", vv); err != nil {
				return fmt.Errorf("Error reading icap_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icap_profile: %v", err)
		}
	}

	if err = d.Set("identity_based", flattenPackagesGlobalHeaderPolicy6IdentityBased(o["identity-based"], d, "identity_based")); err != nil {
		if vv, ok := fortiAPIPatch(o["identity-based"], "PackagesGlobalHeaderPolicy6-IdentityBased"); ok {
			if err = d.Set("identity_based", vv); err != nil {
				return fmt.Errorf("Error reading identity_based: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading identity_based: %v", err)
		}
	}

	if err = d.Set("identity_from", flattenPackagesGlobalHeaderPolicy6IdentityFrom(o["identity-from"], d, "identity_from")); err != nil {
		if vv, ok := fortiAPIPatch(o["identity-from"], "PackagesGlobalHeaderPolicy6-IdentityFrom"); ok {
			if err = d.Set("identity_from", vv); err != nil {
				return fmt.Errorf("Error reading identity_from: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading identity_from: %v", err)
		}
	}

	if err = d.Set("inbound", flattenPackagesGlobalHeaderPolicy6Inbound(o["inbound"], d, "inbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["inbound"], "PackagesGlobalHeaderPolicy6-Inbound"); ok {
			if err = d.Set("inbound", vv); err != nil {
				return fmt.Errorf("Error reading inbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inbound: %v", err)
		}
	}

	if err = d.Set("inspection_mode", flattenPackagesGlobalHeaderPolicy6InspectionMode(o["inspection-mode"], d, "inspection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["inspection-mode"], "PackagesGlobalHeaderPolicy6-InspectionMode"); ok {
			if err = d.Set("inspection_mode", vv); err != nil {
				return fmt.Errorf("Error reading inspection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inspection_mode: %v", err)
		}
	}

	if err = d.Set("ippool", flattenPackagesGlobalHeaderPolicy6Ippool(o["ippool"], d, "ippool")); err != nil {
		if vv, ok := fortiAPIPatch(o["ippool"], "PackagesGlobalHeaderPolicy6-Ippool"); ok {
			if err = d.Set("ippool", vv); err != nil {
				return fmt.Errorf("Error reading ippool: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ippool: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenPackagesGlobalHeaderPolicy6IpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor"], "PackagesGlobalHeaderPolicy6-IpsSensor"); ok {
			if err = d.Set("ips_sensor", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("label", flattenPackagesGlobalHeaderPolicy6Label(o["label"], d, "label")); err != nil {
		if vv, ok := fortiAPIPatch(o["label"], "PackagesGlobalHeaderPolicy6-Label"); ok {
			if err = d.Set("label", vv); err != nil {
				return fmt.Errorf("Error reading label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenPackagesGlobalHeaderPolicy6Logtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic"], "PackagesGlobalHeaderPolicy6-Logtraffic"); ok {
			if err = d.Set("logtraffic", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("logtraffic_start", flattenPackagesGlobalHeaderPolicy6LogtrafficStart(o["logtraffic-start"], d, "logtraffic_start")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic-start"], "PackagesGlobalHeaderPolicy6-LogtrafficStart"); ok {
			if err = d.Set("logtraffic_start", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic_start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic_start: %v", err)
		}
	}

	if err = d.Set("mms_profile", flattenPackagesGlobalHeaderPolicy6MmsProfile(o["mms-profile"], d, "mms_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["mms-profile"], "PackagesGlobalHeaderPolicy6-MmsProfile"); ok {
			if err = d.Set("mms_profile", vv); err != nil {
				return fmt.Errorf("Error reading mms_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mms_profile: %v", err)
		}
	}

	if err = d.Set("name", flattenPackagesGlobalHeaderPolicy6Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesGlobalHeaderPolicy6-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nat", flattenPackagesGlobalHeaderPolicy6Nat(o["nat"], d, "nat")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat"], "PackagesGlobalHeaderPolicy6-Nat"); ok {
			if err = d.Set("nat", vv); err != nil {
				return fmt.Errorf("Error reading nat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat: %v", err)
		}
	}

	if err = d.Set("natinbound", flattenPackagesGlobalHeaderPolicy6Natinbound(o["natinbound"], d, "natinbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["natinbound"], "PackagesGlobalHeaderPolicy6-Natinbound"); ok {
			if err = d.Set("natinbound", vv); err != nil {
				return fmt.Errorf("Error reading natinbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading natinbound: %v", err)
		}
	}

	if err = d.Set("natoutbound", flattenPackagesGlobalHeaderPolicy6Natoutbound(o["natoutbound"], d, "natoutbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["natoutbound"], "PackagesGlobalHeaderPolicy6-Natoutbound"); ok {
			if err = d.Set("natoutbound", vv); err != nil {
				return fmt.Errorf("Error reading natoutbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading natoutbound: %v", err)
		}
	}

	if err = d.Set("np_accelation", flattenPackagesGlobalHeaderPolicy6NpAccelation(o["np-accelation"], d, "np_accelation")); err != nil {
		if vv, ok := fortiAPIPatch(o["np-accelation"], "PackagesGlobalHeaderPolicy6-NpAccelation"); ok {
			if err = d.Set("np_accelation", vv); err != nil {
				return fmt.Errorf("Error reading np_accelation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading np_accelation: %v", err)
		}
	}

	if err = d.Set("np_acceleration", flattenPackagesGlobalHeaderPolicy6NpAcceleration(o["np-acceleration"], d, "np_acceleration")); err != nil {
		if vv, ok := fortiAPIPatch(o["np-acceleration"], "PackagesGlobalHeaderPolicy6-NpAcceleration"); ok {
			if err = d.Set("np_acceleration", vv); err != nil {
				return fmt.Errorf("Error reading np_acceleration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading np_acceleration: %v", err)
		}
	}

	if err = d.Set("outbound", flattenPackagesGlobalHeaderPolicy6Outbound(o["outbound"], d, "outbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbound"], "PackagesGlobalHeaderPolicy6-Outbound"); ok {
			if err = d.Set("outbound", vv); err != nil {
				return fmt.Errorf("Error reading outbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbound: %v", err)
		}
	}

	if err = d.Set("per_ip_shaper", flattenPackagesGlobalHeaderPolicy6PerIpShaper(o["per-ip-shaper"], d, "per_ip_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-ip-shaper"], "PackagesGlobalHeaderPolicy6-PerIpShaper"); ok {
			if err = d.Set("per_ip_shaper", vv); err != nil {
				return fmt.Errorf("Error reading per_ip_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_ip_shaper: %v", err)
		}
	}

	if err = d.Set("policy_offload", flattenPackagesGlobalHeaderPolicy6PolicyOffload(o["policy-offload"], d, "policy_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-offload"], "PackagesGlobalHeaderPolicy6-PolicyOffload"); ok {
			if err = d.Set("policy_offload", vv); err != nil {
				return fmt.Errorf("Error reading policy_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_offload: %v", err)
		}
	}

	if err = d.Set("policyid", flattenPackagesGlobalHeaderPolicy6Policyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "PackagesGlobalHeaderPolicy6-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("poolname", flattenPackagesGlobalHeaderPolicy6Poolname(o["poolname"], d, "poolname")); err != nil {
		if vv, ok := fortiAPIPatch(o["poolname"], "PackagesGlobalHeaderPolicy6-Poolname"); ok {
			if err = d.Set("poolname", vv); err != nil {
				return fmt.Errorf("Error reading poolname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poolname: %v", err)
		}
	}

	if err = d.Set("profile_group", flattenPackagesGlobalHeaderPolicy6ProfileGroup(o["profile-group"], d, "profile_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-group"], "PackagesGlobalHeaderPolicy6-ProfileGroup"); ok {
			if err = d.Set("profile_group", vv); err != nil {
				return fmt.Errorf("Error reading profile_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_group: %v", err)
		}
	}

	if err = d.Set("profile_protocol_options", flattenPackagesGlobalHeaderPolicy6ProfileProtocolOptions(o["profile-protocol-options"], d, "profile_protocol_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-protocol-options"], "PackagesGlobalHeaderPolicy6-ProfileProtocolOptions"); ok {
			if err = d.Set("profile_protocol_options", vv); err != nil {
				return fmt.Errorf("Error reading profile_protocol_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_protocol_options: %v", err)
		}
	}

	if err = d.Set("profile_type", flattenPackagesGlobalHeaderPolicy6ProfileType(o["profile-type"], d, "profile_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-type"], "PackagesGlobalHeaderPolicy6-ProfileType"); ok {
			if err = d.Set("profile_type", vv); err != nil {
				return fmt.Errorf("Error reading profile_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_type: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenPackagesGlobalHeaderPolicy6ReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "PackagesGlobalHeaderPolicy6-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("replacemsg_override_group", flattenPackagesGlobalHeaderPolicy6ReplacemsgOverrideGroup(o["replacemsg-override-group"], d, "replacemsg_override_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-override-group"], "PackagesGlobalHeaderPolicy6-ReplacemsgOverrideGroup"); ok {
			if err = d.Set("replacemsg_override_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
		}
	}

	if err = d.Set("rsso", flattenPackagesGlobalHeaderPolicy6Rsso(o["rsso"], d, "rsso")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso"], "PackagesGlobalHeaderPolicy6-Rsso"); ok {
			if err = d.Set("rsso", vv); err != nil {
				return fmt.Errorf("Error reading rsso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso: %v", err)
		}
	}

	if err = d.Set("schedule", flattenPackagesGlobalHeaderPolicy6Schedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "PackagesGlobalHeaderPolicy6-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("send_deny_packet", flattenPackagesGlobalHeaderPolicy6SendDenyPacket(o["send-deny-packet"], d, "send_deny_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-deny-packet"], "PackagesGlobalHeaderPolicy6-SendDenyPacket"); ok {
			if err = d.Set("send_deny_packet", vv); err != nil {
				return fmt.Errorf("Error reading send_deny_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_deny_packet: %v", err)
		}
	}

	if err = d.Set("service", flattenPackagesGlobalHeaderPolicy6Service(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "PackagesGlobalHeaderPolicy6-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("service_negate", flattenPackagesGlobalHeaderPolicy6ServiceNegate(o["service-negate"], d, "service_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-negate"], "PackagesGlobalHeaderPolicy6-ServiceNegate"); ok {
			if err = d.Set("service_negate", vv); err != nil {
				return fmt.Errorf("Error reading service_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_negate: %v", err)
		}
	}

	if err = d.Set("session_ttl", flattenPackagesGlobalHeaderPolicy6SessionTtl(o["session-ttl"], d, "session_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-ttl"], "PackagesGlobalHeaderPolicy6-SessionTtl"); ok {
			if err = d.Set("session_ttl", vv); err != nil {
				return fmt.Errorf("Error reading session_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_ttl: %v", err)
		}
	}

	if err = d.Set("spamfilter_profile", flattenPackagesGlobalHeaderPolicy6SpamfilterProfile(o["spamfilter-profile"], d, "spamfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["spamfilter-profile"], "PackagesGlobalHeaderPolicy6-SpamfilterProfile"); ok {
			if err = d.Set("spamfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading spamfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spamfilter_profile: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenPackagesGlobalHeaderPolicy6Srcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "PackagesGlobalHeaderPolicy6-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr_negate", flattenPackagesGlobalHeaderPolicy6SrcaddrNegate(o["srcaddr-negate"], d, "srcaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr-negate"], "PackagesGlobalHeaderPolicy6-SrcaddrNegate"); ok {
			if err = d.Set("srcaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr_negate: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenPackagesGlobalHeaderPolicy6Srcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "PackagesGlobalHeaderPolicy6-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("ssh_filter_profile", flattenPackagesGlobalHeaderPolicy6SshFilterProfile(o["ssh-filter-profile"], d, "ssh_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-filter-profile"], "PackagesGlobalHeaderPolicy6-SshFilterProfile"); ok {
			if err = d.Set("ssh_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
		}
	}

	if err = d.Set("ssh_policy_redirect", flattenPackagesGlobalHeaderPolicy6SshPolicyRedirect(o["ssh-policy-redirect"], d, "ssh_policy_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-policy-redirect"], "PackagesGlobalHeaderPolicy6-SshPolicyRedirect"); ok {
			if err = d.Set("ssh_policy_redirect", vv); err != nil {
				return fmt.Errorf("Error reading ssh_policy_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_policy_redirect: %v", err)
		}
	}

	if err = d.Set("ssl_mirror", flattenPackagesGlobalHeaderPolicy6SslMirror(o["ssl-mirror"], d, "ssl_mirror")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-mirror"], "PackagesGlobalHeaderPolicy6-SslMirror"); ok {
			if err = d.Set("ssl_mirror", vv); err != nil {
				return fmt.Errorf("Error reading ssl_mirror: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_mirror: %v", err)
		}
	}

	if err = d.Set("ssl_mirror_intf", flattenPackagesGlobalHeaderPolicy6SslMirrorIntf(o["ssl-mirror-intf"], d, "ssl_mirror_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-mirror-intf"], "PackagesGlobalHeaderPolicy6-SslMirrorIntf"); ok {
			if err = d.Set("ssl_mirror_intf", vv); err != nil {
				return fmt.Errorf("Error reading ssl_mirror_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_mirror_intf: %v", err)
		}
	}

	if err = d.Set("ssl_ssh_profile", flattenPackagesGlobalHeaderPolicy6SslSshProfile(o["ssl-ssh-profile"], d, "ssl_ssh_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-ssh-profile"], "PackagesGlobalHeaderPolicy6-SslSshProfile"); ok {
			if err = d.Set("ssl_ssh_profile", vv); err != nil {
				return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
		}
	}

	if err = d.Set("sslvpn_auth", flattenPackagesGlobalHeaderPolicy6SslvpnAuth(o["sslvpn-auth"], d, "sslvpn_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-auth"], "PackagesGlobalHeaderPolicy6-SslvpnAuth"); ok {
			if err = d.Set("sslvpn_auth", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_auth: %v", err)
		}
	}

	if err = d.Set("sslvpn_ccert", flattenPackagesGlobalHeaderPolicy6SslvpnCcert(o["sslvpn-ccert"], d, "sslvpn_ccert")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-ccert"], "PackagesGlobalHeaderPolicy6-SslvpnCcert"); ok {
			if err = d.Set("sslvpn_ccert", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_ccert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_ccert: %v", err)
		}
	}

	if err = d.Set("sslvpn_cipher", flattenPackagesGlobalHeaderPolicy6SslvpnCipher(o["sslvpn-cipher"], d, "sslvpn_cipher")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-cipher"], "PackagesGlobalHeaderPolicy6-SslvpnCipher"); ok {
			if err = d.Set("sslvpn_cipher", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_cipher: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_cipher: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesGlobalHeaderPolicy6Status(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesGlobalHeaderPolicy6-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tags", flattenPackagesGlobalHeaderPolicy6Tags(o["tags"], d, "tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["tags"], "PackagesGlobalHeaderPolicy6-Tags"); ok {
			if err = d.Set("tags", vv); err != nil {
				return fmt.Errorf("Error reading tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tags: %v", err)
		}
	}

	if err = d.Set("tcp_mss_receiver", flattenPackagesGlobalHeaderPolicy6TcpMssReceiver(o["tcp-mss-receiver"], d, "tcp_mss_receiver")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-mss-receiver"], "PackagesGlobalHeaderPolicy6-TcpMssReceiver"); ok {
			if err = d.Set("tcp_mss_receiver", vv); err != nil {
				return fmt.Errorf("Error reading tcp_mss_receiver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_mss_receiver: %v", err)
		}
	}

	if err = d.Set("tcp_mss_sender", flattenPackagesGlobalHeaderPolicy6TcpMssSender(o["tcp-mss-sender"], d, "tcp_mss_sender")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-mss-sender"], "PackagesGlobalHeaderPolicy6-TcpMssSender"); ok {
			if err = d.Set("tcp_mss_sender", vv); err != nil {
				return fmt.Errorf("Error reading tcp_mss_sender: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_mss_sender: %v", err)
		}
	}

	if err = d.Set("tcp_session_without_syn", flattenPackagesGlobalHeaderPolicy6TcpSessionWithoutSyn(o["tcp-session-without-syn"], d, "tcp_session_without_syn")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-session-without-syn"], "PackagesGlobalHeaderPolicy6-TcpSessionWithoutSyn"); ok {
			if err = d.Set("tcp_session_without_syn", vv); err != nil {
				return fmt.Errorf("Error reading tcp_session_without_syn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_session_without_syn: %v", err)
		}
	}

	if err = d.Set("timeout_send_rst", flattenPackagesGlobalHeaderPolicy6TimeoutSendRst(o["timeout-send-rst"], d, "timeout_send_rst")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout-send-rst"], "PackagesGlobalHeaderPolicy6-TimeoutSendRst"); ok {
			if err = d.Set("timeout_send_rst", vv); err != nil {
				return fmt.Errorf("Error reading timeout_send_rst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout_send_rst: %v", err)
		}
	}

	if err = d.Set("tos", flattenPackagesGlobalHeaderPolicy6Tos(o["tos"], d, "tos")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos"], "PackagesGlobalHeaderPolicy6-Tos"); ok {
			if err = d.Set("tos", vv); err != nil {
				return fmt.Errorf("Error reading tos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenPackagesGlobalHeaderPolicy6TosMask(o["tos-mask"], d, "tos_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos-mask"], "PackagesGlobalHeaderPolicy6-TosMask"); ok {
			if err = d.Set("tos_mask", vv); err != nil {
				return fmt.Errorf("Error reading tos_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("tos_negate", flattenPackagesGlobalHeaderPolicy6TosNegate(o["tos-negate"], d, "tos_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos-negate"], "PackagesGlobalHeaderPolicy6-TosNegate"); ok {
			if err = d.Set("tos_negate", vv); err != nil {
				return fmt.Errorf("Error reading tos_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos_negate: %v", err)
		}
	}

	if err = d.Set("traffic_shaper", flattenPackagesGlobalHeaderPolicy6TrafficShaper(o["traffic-shaper"], d, "traffic_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper"], "PackagesGlobalHeaderPolicy6-TrafficShaper"); ok {
			if err = d.Set("traffic_shaper", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper: %v", err)
		}
	}

	if err = d.Set("traffic_shaper_reverse", flattenPackagesGlobalHeaderPolicy6TrafficShaperReverse(o["traffic-shaper-reverse"], d, "traffic_shaper_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper-reverse"], "PackagesGlobalHeaderPolicy6-TrafficShaperReverse"); ok {
			if err = d.Set("traffic_shaper_reverse", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
		}
	}

	if err = d.Set("url_category", flattenPackagesGlobalHeaderPolicy6UrlCategory(o["url-category"], d, "url_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-category"], "PackagesGlobalHeaderPolicy6-UrlCategory"); ok {
			if err = d.Set("url_category", vv); err != nil {
				return fmt.Errorf("Error reading url_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_category: %v", err)
		}
	}

	if err = d.Set("users", flattenPackagesGlobalHeaderPolicy6Users(o["users"], d, "users")); err != nil {
		if vv, ok := fortiAPIPatch(o["users"], "PackagesGlobalHeaderPolicy6-Users"); ok {
			if err = d.Set("users", vv); err != nil {
				return fmt.Errorf("Error reading users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading users: %v", err)
		}
	}

	if err = d.Set("utm_inspection_mode", flattenPackagesGlobalHeaderPolicy6UtmInspectionMode(o["utm-inspection-mode"], d, "utm_inspection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-inspection-mode"], "PackagesGlobalHeaderPolicy6-UtmInspectionMode"); ok {
			if err = d.Set("utm_inspection_mode", vv); err != nil {
				return fmt.Errorf("Error reading utm_inspection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_inspection_mode: %v", err)
		}
	}

	if err = d.Set("utm_status", flattenPackagesGlobalHeaderPolicy6UtmStatus(o["utm-status"], d, "utm_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-status"], "PackagesGlobalHeaderPolicy6-UtmStatus"); ok {
			if err = d.Set("utm_status", vv); err != nil {
				return fmt.Errorf("Error reading utm_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("uuid", flattenPackagesGlobalHeaderPolicy6Uuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "PackagesGlobalHeaderPolicy6-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("vlan_cos_fwd", flattenPackagesGlobalHeaderPolicy6VlanCosFwd(o["vlan-cos-fwd"], d, "vlan_cos_fwd")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-cos-fwd"], "PackagesGlobalHeaderPolicy6-VlanCosFwd"); ok {
			if err = d.Set("vlan_cos_fwd", vv); err != nil {
				return fmt.Errorf("Error reading vlan_cos_fwd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_cos_fwd: %v", err)
		}
	}

	if err = d.Set("vlan_cos_rev", flattenPackagesGlobalHeaderPolicy6VlanCosRev(o["vlan-cos-rev"], d, "vlan_cos_rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-cos-rev"], "PackagesGlobalHeaderPolicy6-VlanCosRev"); ok {
			if err = d.Set("vlan_cos_rev", vv); err != nil {
				return fmt.Errorf("Error reading vlan_cos_rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_cos_rev: %v", err)
		}
	}

	if err = d.Set("vlan_filter", flattenPackagesGlobalHeaderPolicy6VlanFilter(o["vlan-filter"], d, "vlan_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-filter"], "PackagesGlobalHeaderPolicy6-VlanFilter"); ok {
			if err = d.Set("vlan_filter", vv); err != nil {
				return fmt.Errorf("Error reading vlan_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_filter: %v", err)
		}
	}

	if err = d.Set("voip_profile", flattenPackagesGlobalHeaderPolicy6VoipProfile(o["voip-profile"], d, "voip_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip-profile"], "PackagesGlobalHeaderPolicy6-VoipProfile"); ok {
			if err = d.Set("voip_profile", vv); err != nil {
				return fmt.Errorf("Error reading voip_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip_profile: %v", err)
		}
	}

	if err = d.Set("vpntunnel", flattenPackagesGlobalHeaderPolicy6Vpntunnel(o["vpntunnel"], d, "vpntunnel")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpntunnel"], "PackagesGlobalHeaderPolicy6-Vpntunnel"); ok {
			if err = d.Set("vpntunnel", vv); err != nil {
				return fmt.Errorf("Error reading vpntunnel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpntunnel: %v", err)
		}
	}

	if err = d.Set("waf_profile", flattenPackagesGlobalHeaderPolicy6WafProfile(o["waf-profile"], d, "waf_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["waf-profile"], "PackagesGlobalHeaderPolicy6-WafProfile"); ok {
			if err = d.Set("waf_profile", vv); err != nil {
				return fmt.Errorf("Error reading waf_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading waf_profile: %v", err)
		}
	}

	if err = d.Set("webcache", flattenPackagesGlobalHeaderPolicy6Webcache(o["webcache"], d, "webcache")); err != nil {
		if vv, ok := fortiAPIPatch(o["webcache"], "PackagesGlobalHeaderPolicy6-Webcache"); ok {
			if err = d.Set("webcache", vv); err != nil {
				return fmt.Errorf("Error reading webcache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webcache: %v", err)
		}
	}

	if err = d.Set("webcache_https", flattenPackagesGlobalHeaderPolicy6WebcacheHttps(o["webcache-https"], d, "webcache_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["webcache-https"], "PackagesGlobalHeaderPolicy6-WebcacheHttps"); ok {
			if err = d.Set("webcache_https", vv); err != nil {
				return fmt.Errorf("Error reading webcache_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webcache_https: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenPackagesGlobalHeaderPolicy6WebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile"], "PackagesGlobalHeaderPolicy6-WebfilterProfile"); ok {
			if err = d.Set("webfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("webproxy_forward_server", flattenPackagesGlobalHeaderPolicy6WebproxyForwardServer(o["webproxy-forward-server"], d, "webproxy_forward_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["webproxy-forward-server"], "PackagesGlobalHeaderPolicy6-WebproxyForwardServer"); ok {
			if err = d.Set("webproxy_forward_server", vv); err != nil {
				return fmt.Errorf("Error reading webproxy_forward_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webproxy_forward_server: %v", err)
		}
	}

	if err = d.Set("webproxy_profile", flattenPackagesGlobalHeaderPolicy6WebproxyProfile(o["webproxy-profile"], d, "webproxy_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webproxy-profile"], "PackagesGlobalHeaderPolicy6-WebproxyProfile"); ok {
			if err = d.Set("webproxy_profile", vv); err != nil {
				return fmt.Errorf("Error reading webproxy_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webproxy_profile: %v", err)
		}
	}

	return nil
}

func flattenPackagesGlobalHeaderPolicy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesGlobalHeaderPolicy6Action(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6AntiReplay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6AppCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6AppGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Application(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandPackagesGlobalHeaderPolicy6ApplicationCharts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesGlobalHeaderPolicy6ApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6AutoAsicOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6AvProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6CasiProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6CgnLogServerGrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6CifsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6CustomLogFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6DecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6DeepInspectionOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6DeviceDetectionPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Devices(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6DiffservForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6DiffservReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6DiffservcodeForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6DiffservcodeRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6DlpSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6DnsfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6DscpMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6DscpNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6DscpValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Dsri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Dstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6DstaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Dstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6DynamicProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6DynamicProfileAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesGlobalHeaderPolicy6DynamicProfileGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6EmailCollectionPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6EmailfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6FileFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6FirewallSessionDirty(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Fixedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Fsae(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6FssoGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6GlobalLabel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Groups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6HttpPolicyRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6IcapProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6IdentityBased(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6IdentityFrom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Inbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6InspectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Ippool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6IpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Label(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Logtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6LogtrafficStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6MmsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Nat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Natinbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Natoutbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6NpAccelation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6NpAcceleration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Outbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6PerIpShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6PolicyOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Policyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Poolname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6ProfileGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6ProfileProtocolOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6ProfileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6ReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6ReplacemsgOverrideGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Rsso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Schedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6SendDenyPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Service(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6ServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6SessionTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6SpamfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Srcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6SrcaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Srcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6SshFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6SshPolicyRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6SslMirror(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6SslMirrorIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6SslSshProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6SslvpnAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6SslvpnCcert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6SslvpnCipher(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Tags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6TcpMssReceiver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6TcpMssSender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6TcpSessionWithoutSyn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6TimeoutSendRst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Tos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6TosMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6TosNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6TrafficShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6TrafficShaperReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6UrlCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Users(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6UtmInspectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6UtmStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6VlanCosFwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6VlanCosRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6VlanFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6VoipProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Vpntunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6WafProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6Webcache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6WebcacheHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6WebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6WebproxyForwardServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderPolicy6WebproxyProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesGlobalHeaderPolicy6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandPackagesGlobalHeaderPolicy6Action(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("anti_replay"); ok || d.HasChange("anti_replay") {
		t, err := expandPackagesGlobalHeaderPolicy6AntiReplay(d, v, "anti_replay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anti-replay"] = t
		}
	}

	if v, ok := d.GetOk("app_category"); ok || d.HasChange("app_category") {
		t, err := expandPackagesGlobalHeaderPolicy6AppCategory(d, v, "app_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-category"] = t
		}
	}

	if v, ok := d.GetOk("app_group"); ok || d.HasChange("app_group") {
		t, err := expandPackagesGlobalHeaderPolicy6AppGroup(d, v, "app_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-group"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandPackagesGlobalHeaderPolicy6Application(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("application_charts"); ok || d.HasChange("application_charts") {
		t, err := expandPackagesGlobalHeaderPolicy6ApplicationCharts(d, v, "application_charts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-charts"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok || d.HasChange("application_list") {
		t, err := expandPackagesGlobalHeaderPolicy6ApplicationList(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("auto_asic_offload"); ok || d.HasChange("auto_asic_offload") {
		t, err := expandPackagesGlobalHeaderPolicy6AutoAsicOffload(d, v, "auto_asic_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok || d.HasChange("av_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6AvProfile(d, v, "av_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("casi_profile"); ok || d.HasChange("casi_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6CasiProfile(d, v, "casi_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casi-profile"] = t
		}
	}

	if v, ok := d.GetOk("cgn_log_server_grp"); ok || d.HasChange("cgn_log_server_grp") {
		t, err := expandPackagesGlobalHeaderPolicy6CgnLogServerGrp(d, v, "cgn_log_server_grp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-log-server-grp"] = t
		}
	}

	if v, ok := d.GetOk("cifs_profile"); ok || d.HasChange("cifs_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6CifsProfile(d, v, "cifs_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs-profile"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandPackagesGlobalHeaderPolicy6Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("custom_log_fields"); ok || d.HasChange("custom_log_fields") {
		t, err := expandPackagesGlobalHeaderPolicy6CustomLogFields(d, v, "custom_log_fields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-log-fields"] = t
		}
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok || d.HasChange("decrypted_traffic_mirror") {
		t, err := expandPackagesGlobalHeaderPolicy6DecryptedTrafficMirror(d, v, "decrypted_traffic_mirror")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	}

	if v, ok := d.GetOk("deep_inspection_options"); ok || d.HasChange("deep_inspection_options") {
		t, err := expandPackagesGlobalHeaderPolicy6DeepInspectionOptions(d, v, "deep_inspection_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deep-inspection-options"] = t
		}
	}

	if v, ok := d.GetOk("device_detection_portal"); ok || d.HasChange("device_detection_portal") {
		t, err := expandPackagesGlobalHeaderPolicy6DeviceDetectionPortal(d, v, "device_detection_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-detection-portal"] = t
		}
	}

	if v, ok := d.GetOk("devices"); ok || d.HasChange("devices") {
		t, err := expandPackagesGlobalHeaderPolicy6Devices(d, v, "devices")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devices"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_forward"); ok || d.HasChange("diffserv_forward") {
		t, err := expandPackagesGlobalHeaderPolicy6DiffservForward(d, v, "diffserv_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_reverse"); ok || d.HasChange("diffserv_reverse") {
		t, err := expandPackagesGlobalHeaderPolicy6DiffservReverse(d, v, "diffserv_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-reverse"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_forward"); ok || d.HasChange("diffservcode_forward") {
		t, err := expandPackagesGlobalHeaderPolicy6DiffservcodeForward(d, v, "diffservcode_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_rev"); ok || d.HasChange("diffservcode_rev") {
		t, err := expandPackagesGlobalHeaderPolicy6DiffservcodeRev(d, v, "diffservcode_rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-rev"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok || d.HasChange("dlp_sensor") {
		t, err := expandPackagesGlobalHeaderPolicy6DlpSensor(d, v, "dlp_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("dnsfilter_profile"); ok || d.HasChange("dnsfilter_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6DnsfilterProfile(d, v, "dnsfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dscp_match"); ok || d.HasChange("dscp_match") {
		t, err := expandPackagesGlobalHeaderPolicy6DscpMatch(d, v, "dscp_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-match"] = t
		}
	}

	if v, ok := d.GetOk("dscp_negate"); ok || d.HasChange("dscp_negate") {
		t, err := expandPackagesGlobalHeaderPolicy6DscpNegate(d, v, "dscp_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-negate"] = t
		}
	}

	if v, ok := d.GetOk("dscp_value"); ok || d.HasChange("dscp_value") {
		t, err := expandPackagesGlobalHeaderPolicy6DscpValue(d, v, "dscp_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-value"] = t
		}
	}

	if v, ok := d.GetOk("dsri"); ok || d.HasChange("dsri") {
		t, err := expandPackagesGlobalHeaderPolicy6Dsri(d, v, "dsri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsri"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandPackagesGlobalHeaderPolicy6Dstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr_negate"); ok || d.HasChange("dstaddr_negate") {
		t, err := expandPackagesGlobalHeaderPolicy6DstaddrNegate(d, v, "dstaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandPackagesGlobalHeaderPolicy6Dstintf(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_profile"); ok || d.HasChange("dynamic_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6DynamicProfile(d, v, "dynamic_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-profile"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_profile_access"); ok || d.HasChange("dynamic_profile_access") {
		t, err := expandPackagesGlobalHeaderPolicy6DynamicProfileAccess(d, v, "dynamic_profile_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-profile-access"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_profile_group"); ok || d.HasChange("dynamic_profile_group") {
		t, err := expandPackagesGlobalHeaderPolicy6DynamicProfileGroup(d, v, "dynamic_profile_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-profile-group"] = t
		}
	}

	if v, ok := d.GetOk("email_collection_portal"); ok || d.HasChange("email_collection_portal") {
		t, err := expandPackagesGlobalHeaderPolicy6EmailCollectionPortal(d, v, "email_collection_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-collection-portal"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok || d.HasChange("emailfilter_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6EmailfilterProfile(d, v, "emailfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("file_filter_profile"); ok || d.HasChange("file_filter_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6FileFilterProfile(d, v, "file_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("firewall_session_dirty"); ok || d.HasChange("firewall_session_dirty") {
		t, err := expandPackagesGlobalHeaderPolicy6FirewallSessionDirty(d, v, "firewall_session_dirty")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-session-dirty"] = t
		}
	}

	if v, ok := d.GetOk("fixedport"); ok || d.HasChange("fixedport") {
		t, err := expandPackagesGlobalHeaderPolicy6Fixedport(d, v, "fixedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fixedport"] = t
		}
	}

	if v, ok := d.GetOk("fsae"); ok || d.HasChange("fsae") {
		t, err := expandPackagesGlobalHeaderPolicy6Fsae(d, v, "fsae")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsae"] = t
		}
	}

	if v, ok := d.GetOk("fsso_groups"); ok || d.HasChange("fsso_groups") {
		t, err := expandPackagesGlobalHeaderPolicy6FssoGroups(d, v, "fsso_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-groups"] = t
		}
	}

	if v, ok := d.GetOk("global_label"); ok || d.HasChange("global_label") {
		t, err := expandPackagesGlobalHeaderPolicy6GlobalLabel(d, v, "global_label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-label"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandPackagesGlobalHeaderPolicy6Groups(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("http_policy_redirect"); ok || d.HasChange("http_policy_redirect") {
		t, err := expandPackagesGlobalHeaderPolicy6HttpPolicyRedirect(d, v, "http_policy_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("icap_profile"); ok || d.HasChange("icap_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6IcapProfile(d, v, "icap_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-profile"] = t
		}
	}

	if v, ok := d.GetOk("identity_based"); ok || d.HasChange("identity_based") {
		t, err := expandPackagesGlobalHeaderPolicy6IdentityBased(d, v, "identity_based")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["identity-based"] = t
		}
	}

	if v, ok := d.GetOk("identity_from"); ok || d.HasChange("identity_from") {
		t, err := expandPackagesGlobalHeaderPolicy6IdentityFrom(d, v, "identity_from")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["identity-from"] = t
		}
	}

	if v, ok := d.GetOk("inbound"); ok || d.HasChange("inbound") {
		t, err := expandPackagesGlobalHeaderPolicy6Inbound(d, v, "inbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound"] = t
		}
	}

	if v, ok := d.GetOk("inspection_mode"); ok || d.HasChange("inspection_mode") {
		t, err := expandPackagesGlobalHeaderPolicy6InspectionMode(d, v, "inspection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspection-mode"] = t
		}
	}

	if v, ok := d.GetOk("ippool"); ok || d.HasChange("ippool") {
		t, err := expandPackagesGlobalHeaderPolicy6Ippool(d, v, "ippool")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ippool"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok || d.HasChange("ips_sensor") {
		t, err := expandPackagesGlobalHeaderPolicy6IpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("label"); ok || d.HasChange("label") {
		t, err := expandPackagesGlobalHeaderPolicy6Label(d, v, "label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["label"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok || d.HasChange("logtraffic") {
		t, err := expandPackagesGlobalHeaderPolicy6Logtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic_start"); ok || d.HasChange("logtraffic_start") {
		t, err := expandPackagesGlobalHeaderPolicy6LogtrafficStart(d, v, "logtraffic_start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic-start"] = t
		}
	}

	if v, ok := d.GetOk("mms_profile"); ok || d.HasChange("mms_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6MmsProfile(d, v, "mms_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-profile"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandPackagesGlobalHeaderPolicy6Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nat"); ok || d.HasChange("nat") {
		t, err := expandPackagesGlobalHeaderPolicy6Nat(d, v, "nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat"] = t
		}
	}

	if v, ok := d.GetOk("natinbound"); ok || d.HasChange("natinbound") {
		t, err := expandPackagesGlobalHeaderPolicy6Natinbound(d, v, "natinbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natinbound"] = t
		}
	}

	if v, ok := d.GetOk("natoutbound"); ok || d.HasChange("natoutbound") {
		t, err := expandPackagesGlobalHeaderPolicy6Natoutbound(d, v, "natoutbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natoutbound"] = t
		}
	}

	if v, ok := d.GetOk("np_accelation"); ok || d.HasChange("np_accelation") {
		t, err := expandPackagesGlobalHeaderPolicy6NpAccelation(d, v, "np_accelation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["np-accelation"] = t
		}
	}

	if v, ok := d.GetOk("np_acceleration"); ok || d.HasChange("np_acceleration") {
		t, err := expandPackagesGlobalHeaderPolicy6NpAcceleration(d, v, "np_acceleration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["np-acceleration"] = t
		}
	}

	if v, ok := d.GetOk("outbound"); ok || d.HasChange("outbound") {
		t, err := expandPackagesGlobalHeaderPolicy6Outbound(d, v, "outbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbound"] = t
		}
	}

	if v, ok := d.GetOk("per_ip_shaper"); ok || d.HasChange("per_ip_shaper") {
		t, err := expandPackagesGlobalHeaderPolicy6PerIpShaper(d, v, "per_ip_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-ip-shaper"] = t
		}
	}

	if v, ok := d.GetOk("policy_offload"); ok || d.HasChange("policy_offload") {
		t, err := expandPackagesGlobalHeaderPolicy6PolicyOffload(d, v, "policy_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-offload"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandPackagesGlobalHeaderPolicy6Policyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("poolname"); ok || d.HasChange("poolname") {
		t, err := expandPackagesGlobalHeaderPolicy6Poolname(d, v, "poolname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poolname"] = t
		}
	}

	if v, ok := d.GetOk("profile_group"); ok || d.HasChange("profile_group") {
		t, err := expandPackagesGlobalHeaderPolicy6ProfileGroup(d, v, "profile_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-group"] = t
		}
	}

	if v, ok := d.GetOk("profile_protocol_options"); ok || d.HasChange("profile_protocol_options") {
		t, err := expandPackagesGlobalHeaderPolicy6ProfileProtocolOptions(d, v, "profile_protocol_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-protocol-options"] = t
		}
	}

	if v, ok := d.GetOk("profile_type"); ok || d.HasChange("profile_type") {
		t, err := expandPackagesGlobalHeaderPolicy6ProfileType(d, v, "profile_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-type"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandPackagesGlobalHeaderPolicy6ReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_override_group"); ok || d.HasChange("replacemsg_override_group") {
		t, err := expandPackagesGlobalHeaderPolicy6ReplacemsgOverrideGroup(d, v, "replacemsg_override_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-override-group"] = t
		}
	}

	if v, ok := d.GetOk("rsso"); ok || d.HasChange("rsso") {
		t, err := expandPackagesGlobalHeaderPolicy6Rsso(d, v, "rsso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandPackagesGlobalHeaderPolicy6Schedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("send_deny_packet"); ok || d.HasChange("send_deny_packet") {
		t, err := expandPackagesGlobalHeaderPolicy6SendDenyPacket(d, v, "send_deny_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-deny-packet"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandPackagesGlobalHeaderPolicy6Service(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("service_negate"); ok || d.HasChange("service_negate") {
		t, err := expandPackagesGlobalHeaderPolicy6ServiceNegate(d, v, "service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-negate"] = t
		}
	}

	if v, ok := d.GetOk("session_ttl"); ok || d.HasChange("session_ttl") {
		t, err := expandPackagesGlobalHeaderPolicy6SessionTtl(d, v, "session_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-ttl"] = t
		}
	}

	if v, ok := d.GetOk("spamfilter_profile"); ok || d.HasChange("spamfilter_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6SpamfilterProfile(d, v, "spamfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spamfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandPackagesGlobalHeaderPolicy6Srcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr_negate"); ok || d.HasChange("srcaddr_negate") {
		t, err := expandPackagesGlobalHeaderPolicy6SrcaddrNegate(d, v, "srcaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandPackagesGlobalHeaderPolicy6Srcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("ssh_filter_profile"); ok || d.HasChange("ssh_filter_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6SshFilterProfile(d, v, "ssh_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("ssh_policy_redirect"); ok || d.HasChange("ssh_policy_redirect") {
		t, err := expandPackagesGlobalHeaderPolicy6SshPolicyRedirect(d, v, "ssh_policy_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mirror"); ok || d.HasChange("ssl_mirror") {
		t, err := expandPackagesGlobalHeaderPolicy6SslMirror(d, v, "ssl_mirror")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mirror"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mirror_intf"); ok || d.HasChange("ssl_mirror_intf") {
		t, err := expandPackagesGlobalHeaderPolicy6SslMirrorIntf(d, v, "ssl_mirror_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mirror-intf"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ssh_profile"); ok || d.HasChange("ssl_ssh_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6SslSshProfile(d, v, "ssl_ssh_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ssh-profile"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_auth"); ok || d.HasChange("sslvpn_auth") {
		t, err := expandPackagesGlobalHeaderPolicy6SslvpnAuth(d, v, "sslvpn_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-auth"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_ccert"); ok || d.HasChange("sslvpn_ccert") {
		t, err := expandPackagesGlobalHeaderPolicy6SslvpnCcert(d, v, "sslvpn_ccert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-ccert"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_cipher"); ok || d.HasChange("sslvpn_cipher") {
		t, err := expandPackagesGlobalHeaderPolicy6SslvpnCipher(d, v, "sslvpn_cipher")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-cipher"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandPackagesGlobalHeaderPolicy6Status(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
		t, err := expandPackagesGlobalHeaderPolicy6Tags(d, v, "tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tags"] = t
		}
	}

	if v, ok := d.GetOk("tcp_mss_receiver"); ok || d.HasChange("tcp_mss_receiver") {
		t, err := expandPackagesGlobalHeaderPolicy6TcpMssReceiver(d, v, "tcp_mss_receiver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-receiver"] = t
		}
	}

	if v, ok := d.GetOk("tcp_mss_sender"); ok || d.HasChange("tcp_mss_sender") {
		t, err := expandPackagesGlobalHeaderPolicy6TcpMssSender(d, v, "tcp_mss_sender")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-sender"] = t
		}
	}

	if v, ok := d.GetOk("tcp_session_without_syn"); ok || d.HasChange("tcp_session_without_syn") {
		t, err := expandPackagesGlobalHeaderPolicy6TcpSessionWithoutSyn(d, v, "tcp_session_without_syn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-session-without-syn"] = t
		}
	}

	if v, ok := d.GetOk("timeout_send_rst"); ok || d.HasChange("timeout_send_rst") {
		t, err := expandPackagesGlobalHeaderPolicy6TimeoutSendRst(d, v, "timeout_send_rst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-send-rst"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok || d.HasChange("tos") {
		t, err := expandPackagesGlobalHeaderPolicy6Tos(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok || d.HasChange("tos_mask") {
		t, err := expandPackagesGlobalHeaderPolicy6TosMask(d, v, "tos_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	if v, ok := d.GetOk("tos_negate"); ok || d.HasChange("tos_negate") {
		t, err := expandPackagesGlobalHeaderPolicy6TosNegate(d, v, "tos_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-negate"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper"); ok || d.HasChange("traffic_shaper") {
		t, err := expandPackagesGlobalHeaderPolicy6TrafficShaper(d, v, "traffic_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper_reverse"); ok || d.HasChange("traffic_shaper_reverse") {
		t, err := expandPackagesGlobalHeaderPolicy6TrafficShaperReverse(d, v, "traffic_shaper_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper-reverse"] = t
		}
	}

	if v, ok := d.GetOk("url_category"); ok || d.HasChange("url_category") {
		t, err := expandPackagesGlobalHeaderPolicy6UrlCategory(d, v, "url_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-category"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok || d.HasChange("users") {
		t, err := expandPackagesGlobalHeaderPolicy6Users(d, v, "users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("utm_inspection_mode"); ok || d.HasChange("utm_inspection_mode") {
		t, err := expandPackagesGlobalHeaderPolicy6UtmInspectionMode(d, v, "utm_inspection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-inspection-mode"] = t
		}
	}

	if v, ok := d.GetOk("utm_status"); ok || d.HasChange("utm_status") {
		t, err := expandPackagesGlobalHeaderPolicy6UtmStatus(d, v, "utm_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandPackagesGlobalHeaderPolicy6Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("vlan_cos_fwd"); ok || d.HasChange("vlan_cos_fwd") {
		t, err := expandPackagesGlobalHeaderPolicy6VlanCosFwd(d, v, "vlan_cos_fwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-cos-fwd"] = t
		}
	}

	if v, ok := d.GetOk("vlan_cos_rev"); ok || d.HasChange("vlan_cos_rev") {
		t, err := expandPackagesGlobalHeaderPolicy6VlanCosRev(d, v, "vlan_cos_rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-cos-rev"] = t
		}
	}

	if v, ok := d.GetOk("vlan_filter"); ok || d.HasChange("vlan_filter") {
		t, err := expandPackagesGlobalHeaderPolicy6VlanFilter(d, v, "vlan_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-filter"] = t
		}
	}

	if v, ok := d.GetOk("voip_profile"); ok || d.HasChange("voip_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6VoipProfile(d, v, "voip_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-profile"] = t
		}
	}

	if v, ok := d.GetOk("vpntunnel"); ok || d.HasChange("vpntunnel") {
		t, err := expandPackagesGlobalHeaderPolicy6Vpntunnel(d, v, "vpntunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpntunnel"] = t
		}
	}

	if v, ok := d.GetOk("waf_profile"); ok || d.HasChange("waf_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6WafProfile(d, v, "waf_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["waf-profile"] = t
		}
	}

	if v, ok := d.GetOk("webcache"); ok || d.HasChange("webcache") {
		t, err := expandPackagesGlobalHeaderPolicy6Webcache(d, v, "webcache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache"] = t
		}
	}

	if v, ok := d.GetOk("webcache_https"); ok || d.HasChange("webcache_https") {
		t, err := expandPackagesGlobalHeaderPolicy6WebcacheHttps(d, v, "webcache_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache-https"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok || d.HasChange("webfilter_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6WebfilterProfile(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_forward_server"); ok || d.HasChange("webproxy_forward_server") {
		t, err := expandPackagesGlobalHeaderPolicy6WebproxyForwardServer(d, v, "webproxy_forward_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-forward-server"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_profile"); ok || d.HasChange("webproxy_profile") {
		t, err := expandPackagesGlobalHeaderPolicy6WebproxyProfile(d, v, "webproxy_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-profile"] = t
		}
	}

	return &obj, nil
}
