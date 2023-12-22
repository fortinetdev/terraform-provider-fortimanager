// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPv6 of interface.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFspVlanInterfaceIpv6() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFspVlanInterfaceIpv6Update,
		Read:   resourceObjectFspVlanInterfaceIpv6Read,
		Update: resourceObjectFspVlanInterfaceIpv6Update,
		Delete: resourceObjectFspVlanInterfaceIpv6Delete,

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
			"vlan": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"autoconf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cli_conn6_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dhcp6_client_options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dhcp6_information_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_prefix_delegation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_prefix_hint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_prefix_hint_plt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhcp6_prefix_hint_vlt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhcp6_relay_interface_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp6_relay_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp6_relay_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_relay_source_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_relay_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_relay_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp6_send_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip6_default_life": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip6_delegated_prefix_iaid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip6_delegated_prefix_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"autonomous_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"delegated_prefix_iaid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"onlink_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"prefix_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rdnss": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rdnss_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"subnet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"upstream_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ip6_dns_server_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_extra_addr": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ip6_hop_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip6_link_mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip6_manage_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_max_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip6_min_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip6_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_other_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_prefix_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"autonomous_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dnssl": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"onlink_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"preferred_life_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rdnss": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"valid_life_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"ip6_prefix_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_reachable_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip6_retrans_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip6_send_adv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_upstream_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nd_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nd_cga_modifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nd_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nd_security_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nd_timestamp_delta": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nd_timestamp_fuzz": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ra_send_mtu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unique_autoconf_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrip6_link_local": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrrp_virtual_mac6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrrp6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accept_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"adv_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"preempt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"start_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vrdst6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vrgrp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vrid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vrip6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectFspVlanInterfaceIpv6Update(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	obj, err := getObjectObjectFspVlanInterfaceIpv6(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterfaceIpv6 resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFspVlanInterfaceIpv6(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterfaceIpv6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFspVlanInterfaceIpv6")

	return resourceObjectFspVlanInterfaceIpv6Read(d, m)
}

func resourceObjectFspVlanInterfaceIpv6Delete(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	err = c.DeleteObjectFspVlanInterfaceIpv6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFspVlanInterfaceIpv6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFspVlanInterfaceIpv6Read(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	if vlan == "" {
		vlan = importOptionChecking(m.(*FortiClient).Cfg, "vlan")
		if vlan == "" {
			return fmt.Errorf("Parameter vlan is missing")
		}
		if err = d.Set("vlan", vlan); err != nil {
			return fmt.Errorf("Error set params vlan: %v", err)
		}
	}
	paradict["vlan"] = vlan

	o, err := c.ReadObjectFspVlanInterfaceIpv6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterfaceIpv6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFspVlanInterfaceIpv6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterfaceIpv6 resource from API: %v", err)
	}
	return nil
}

func flattenObjectFspVlanInterfaceIpv6Autoconf3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6CliConn6Status3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6ClientOptions3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6InformationRequest3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixDelegation3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHint3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHintPlt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHintVlt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6RelayInterfaceId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6RelayIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6RelayService3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6RelaySourceInterface3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6RelaySourceIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Dhcp6RelayType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Icmp6SendRedirect3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6InterfaceIdentifier3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6Address3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6Allowaccess3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Ip6DefaultLife3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixIaid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := i["autonomous-flag"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag3rdl(i["autonomous-flag"], d, pre_append)
			tmp["autonomous_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-AutonomousFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delegated_prefix_iaid"
		if _, ok := i["delegated-prefix-iaid"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid3rdl(i["delegated-prefix-iaid"], d, pre_append)
			tmp["delegated_prefix_iaid"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-DelegatedPrefixIaid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag3rdl(i["onlink-flag"], d, pre_append)
			tmp["onlink_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-OnlinkFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_id"
		if _, ok := i["prefix-id"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId3rdl(i["prefix-id"], d, pre_append)
			tmp["prefix_id"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-PrefixId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss3rdl(i["rdnss"], d, pre_append)
			tmp["rdnss"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-Rdnss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_service"
		if _, ok := i["rdnss-service"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService3rdl(i["rdnss-service"], d, pre_append)
			tmp["rdnss_service"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-RdnssService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet3rdl(i["subnet"], d, pre_append)
			tmp["subnet"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upstream_interface"
		if _, ok := i["upstream-interface"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface3rdl(i["upstream-interface"], d, pre_append)
			tmp["upstream_interface"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList-UpstreamInterface")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DnsServerOverride3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddr3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix3rdl(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6ExtraAddr-Prefix")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6HopLimit3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6LinkMtu3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6ManageFlag3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6MaxInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6MinInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6Mode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6OtherFlag3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixList3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := i["autonomous-flag"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListAutonomousFlag3rdl(i["autonomous-flag"], d, pre_append)
			tmp["autonomous_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-AutonomousFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl"
		if _, ok := i["dnssl"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListDnssl3rdl(i["dnssl"], d, pre_append)
			tmp["dnssl"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-Dnssl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListOnlinkFlag3rdl(i["onlink-flag"], d, pre_append)
			tmp["onlink_flag"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-OnlinkFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := i["preferred-life-time"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListPreferredLifeTime3rdl(i["preferred-life-time"], d, pre_append)
			tmp["preferred_life_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-PreferredLifeTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListPrefix3rdl(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListRdnss3rdl(i["rdnss"], d, pre_append)
			tmp["rdnss"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-Rdnss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := i["valid-life-time"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Ip6PrefixListValidLifeTime3rdl(i["valid-life-time"], d, pre_append)
			tmp["valid_life_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Ip6PrefixList-ValidLifeTime")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListAutonomousFlag3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListDnssl3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListOnlinkFlag3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListPreferredLifeTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListPrefix3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListRdnss3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixListValidLifeTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6PrefixMode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6ReachableTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6RetransTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6SendAdv3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6Subnet3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6UpstreamInterface3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFspVlanInterfaceIpv6NdCert3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdCgaModifier3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdMode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdSecurityLevel3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdTimestampDelta3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6NdTimestampFuzz3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6RaSendMtu3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6UniqueAutoconfAddr3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrip6LinkLocal3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6VrrpVirtualMac63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp63rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := i["accept-mode"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6AcceptMode3rdl(i["accept-mode"], d, pre_append)
			tmp["accept_mode"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-AcceptMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := i["adv-interval"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6AdvInterval3rdl(i["adv-interval"], d, pre_append)
			tmp["adv_interval"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-AdvInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := i["preempt"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Preempt3rdl(i["preempt"], d, pre_append)
			tmp["preempt"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Preempt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Priority3rdl(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := i["start-time"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6StartTime3rdl(i["start-time"], d, pre_append)
			tmp["start_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-StartTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Status3rdl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := i["vrdst6"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Vrdst63rdl(i["vrdst6"], d, pre_append)
			tmp["vrdst6"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Vrdst6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := i["vrgrp"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Vrgrp3rdl(i["vrgrp"], d, pre_append)
			tmp["vrgrp"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Vrgrp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := i["vrid"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Vrid3rdl(i["vrid"], d, pre_append)
			tmp["vrid"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Vrid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := i["vrip6"]; ok {
			v := flattenObjectFspVlanInterfaceIpv6Vrrp6Vrip63rdl(i["vrip6"], d, pre_append)
			tmp["vrip6"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceIpv6-Vrrp6-Vrip6")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6AcceptMode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6AdvInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Preempt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Priority3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6StartTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Status3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrdst63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrgrp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrip63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFspVlanInterfaceIpv6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("autoconf", flattenObjectFspVlanInterfaceIpv6Autoconf3rdl(o["autoconf"], d, "autoconf")); err != nil {
		if vv, ok := fortiAPIPatch(o["autoconf"], "ObjectFspVlanInterfaceIpv6-Autoconf"); ok {
			if err = d.Set("autoconf", vv); err != nil {
				return fmt.Errorf("Error reading autoconf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading autoconf: %v", err)
		}
	}

	if err = d.Set("cli_conn6_status", flattenObjectFspVlanInterfaceIpv6CliConn6Status3rdl(o["cli-conn6-status"], d, "cli_conn6_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["cli-conn6-status"], "ObjectFspVlanInterfaceIpv6-CliConn6Status"); ok {
			if err = d.Set("cli_conn6_status", vv); err != nil {
				return fmt.Errorf("Error reading cli_conn6_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cli_conn6_status: %v", err)
		}
	}

	if err = d.Set("dhcp6_client_options", flattenObjectFspVlanInterfaceIpv6Dhcp6ClientOptions3rdl(o["dhcp6-client-options"], d, "dhcp6_client_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-client-options"], "ObjectFspVlanInterfaceIpv6-Dhcp6ClientOptions"); ok {
			if err = d.Set("dhcp6_client_options", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_client_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_client_options: %v", err)
		}
	}

	if err = d.Set("dhcp6_information_request", flattenObjectFspVlanInterfaceIpv6Dhcp6InformationRequest3rdl(o["dhcp6-information-request"], d, "dhcp6_information_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-information-request"], "ObjectFspVlanInterfaceIpv6-Dhcp6InformationRequest"); ok {
			if err = d.Set("dhcp6_information_request", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_information_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_information_request: %v", err)
		}
	}

	if err = d.Set("dhcp6_prefix_delegation", flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixDelegation3rdl(o["dhcp6-prefix-delegation"], d, "dhcp6_prefix_delegation")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-prefix-delegation"], "ObjectFspVlanInterfaceIpv6-Dhcp6PrefixDelegation"); ok {
			if err = d.Set("dhcp6_prefix_delegation", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_prefix_delegation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_prefix_delegation: %v", err)
		}
	}

	if err = d.Set("dhcp6_prefix_hint", flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHint3rdl(o["dhcp6-prefix-hint"], d, "dhcp6_prefix_hint")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-prefix-hint"], "ObjectFspVlanInterfaceIpv6-Dhcp6PrefixHint"); ok {
			if err = d.Set("dhcp6_prefix_hint", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_prefix_hint: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_prefix_hint: %v", err)
		}
	}

	if err = d.Set("dhcp6_prefix_hint_plt", flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHintPlt3rdl(o["dhcp6-prefix-hint-plt"], d, "dhcp6_prefix_hint_plt")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-prefix-hint-plt"], "ObjectFspVlanInterfaceIpv6-Dhcp6PrefixHintPlt"); ok {
			if err = d.Set("dhcp6_prefix_hint_plt", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_prefix_hint_plt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_prefix_hint_plt: %v", err)
		}
	}

	if err = d.Set("dhcp6_prefix_hint_vlt", flattenObjectFspVlanInterfaceIpv6Dhcp6PrefixHintVlt3rdl(o["dhcp6-prefix-hint-vlt"], d, "dhcp6_prefix_hint_vlt")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-prefix-hint-vlt"], "ObjectFspVlanInterfaceIpv6-Dhcp6PrefixHintVlt"); ok {
			if err = d.Set("dhcp6_prefix_hint_vlt", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_prefix_hint_vlt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_prefix_hint_vlt: %v", err)
		}
	}

	if err = d.Set("dhcp6_relay_interface_id", flattenObjectFspVlanInterfaceIpv6Dhcp6RelayInterfaceId3rdl(o["dhcp6-relay-interface-id"], d, "dhcp6_relay_interface_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-relay-interface-id"], "ObjectFspVlanInterfaceIpv6-Dhcp6RelayInterfaceId"); ok {
			if err = d.Set("dhcp6_relay_interface_id", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_relay_interface_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_relay_interface_id: %v", err)
		}
	}

	if err = d.Set("dhcp6_relay_ip", flattenObjectFspVlanInterfaceIpv6Dhcp6RelayIp3rdl(o["dhcp6-relay-ip"], d, "dhcp6_relay_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-relay-ip"], "ObjectFspVlanInterfaceIpv6-Dhcp6RelayIp"); ok {
			if err = d.Set("dhcp6_relay_ip", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_relay_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_relay_ip: %v", err)
		}
	}

	if err = d.Set("dhcp6_relay_service", flattenObjectFspVlanInterfaceIpv6Dhcp6RelayService3rdl(o["dhcp6-relay-service"], d, "dhcp6_relay_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-relay-service"], "ObjectFspVlanInterfaceIpv6-Dhcp6RelayService"); ok {
			if err = d.Set("dhcp6_relay_service", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_relay_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_relay_service: %v", err)
		}
	}

	if err = d.Set("dhcp6_relay_source_interface", flattenObjectFspVlanInterfaceIpv6Dhcp6RelaySourceInterface3rdl(o["dhcp6-relay-source-interface"], d, "dhcp6_relay_source_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-relay-source-interface"], "ObjectFspVlanInterfaceIpv6-Dhcp6RelaySourceInterface"); ok {
			if err = d.Set("dhcp6_relay_source_interface", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_relay_source_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_relay_source_interface: %v", err)
		}
	}

	if err = d.Set("dhcp6_relay_source_ip", flattenObjectFspVlanInterfaceIpv6Dhcp6RelaySourceIp3rdl(o["dhcp6-relay-source-ip"], d, "dhcp6_relay_source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-relay-source-ip"], "ObjectFspVlanInterfaceIpv6-Dhcp6RelaySourceIp"); ok {
			if err = d.Set("dhcp6_relay_source_ip", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_relay_source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_relay_source_ip: %v", err)
		}
	}

	if err = d.Set("dhcp6_relay_type", flattenObjectFspVlanInterfaceIpv6Dhcp6RelayType3rdl(o["dhcp6-relay-type"], d, "dhcp6_relay_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-relay-type"], "ObjectFspVlanInterfaceIpv6-Dhcp6RelayType"); ok {
			if err = d.Set("dhcp6_relay_type", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_relay_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_relay_type: %v", err)
		}
	}

	if err = d.Set("icmp6_send_redirect", flattenObjectFspVlanInterfaceIpv6Icmp6SendRedirect3rdl(o["icmp6-send-redirect"], d, "icmp6_send_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp6-send-redirect"], "ObjectFspVlanInterfaceIpv6-Icmp6SendRedirect"); ok {
			if err = d.Set("icmp6_send_redirect", vv); err != nil {
				return fmt.Errorf("Error reading icmp6_send_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp6_send_redirect: %v", err)
		}
	}

	if err = d.Set("interface_identifier", flattenObjectFspVlanInterfaceIpv6InterfaceIdentifier3rdl(o["interface-identifier"], d, "interface_identifier")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-identifier"], "ObjectFspVlanInterfaceIpv6-InterfaceIdentifier"); ok {
			if err = d.Set("interface_identifier", vv); err != nil {
				return fmt.Errorf("Error reading interface_identifier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_identifier: %v", err)
		}
	}

	if err = d.Set("ip6_address", flattenObjectFspVlanInterfaceIpv6Ip6Address3rdl(o["ip6-address"], d, "ip6_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-address"], "ObjectFspVlanInterfaceIpv6-Ip6Address"); ok {
			if err = d.Set("ip6_address", vv); err != nil {
				return fmt.Errorf("Error reading ip6_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_address: %v", err)
		}
	}

	if err = d.Set("ip6_allowaccess", flattenObjectFspVlanInterfaceIpv6Ip6Allowaccess3rdl(o["ip6-allowaccess"], d, "ip6_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-allowaccess"], "ObjectFspVlanInterfaceIpv6-Ip6Allowaccess"); ok {
			if err = d.Set("ip6_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading ip6_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_allowaccess: %v", err)
		}
	}

	if err = d.Set("ip6_default_life", flattenObjectFspVlanInterfaceIpv6Ip6DefaultLife3rdl(o["ip6-default-life"], d, "ip6_default_life")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-default-life"], "ObjectFspVlanInterfaceIpv6-Ip6DefaultLife"); ok {
			if err = d.Set("ip6_default_life", vv); err != nil {
				return fmt.Errorf("Error reading ip6_default_life: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_default_life: %v", err)
		}
	}

	if err = d.Set("ip6_delegated_prefix_iaid", flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixIaid3rdl(o["ip6-delegated-prefix-iaid"], d, "ip6_delegated_prefix_iaid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-delegated-prefix-iaid"], "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixIaid"); ok {
			if err = d.Set("ip6_delegated_prefix_iaid", vv); err != nil {
				return fmt.Errorf("Error reading ip6_delegated_prefix_iaid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_delegated_prefix_iaid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip6_delegated_prefix_list", flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList3rdl(o["ip6-delegated-prefix-list"], d, "ip6_delegated_prefix_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip6-delegated-prefix-list"], "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList"); ok {
				if err = d.Set("ip6_delegated_prefix_list", vv); err != nil {
					return fmt.Errorf("Error reading ip6_delegated_prefix_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip6_delegated_prefix_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip6_delegated_prefix_list"); ok {
			if err = d.Set("ip6_delegated_prefix_list", flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList3rdl(o["ip6-delegated-prefix-list"], d, "ip6_delegated_prefix_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip6-delegated-prefix-list"], "ObjectFspVlanInterfaceIpv6-Ip6DelegatedPrefixList"); ok {
					if err = d.Set("ip6_delegated_prefix_list", vv); err != nil {
						return fmt.Errorf("Error reading ip6_delegated_prefix_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip6_delegated_prefix_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("ip6_dns_server_override", flattenObjectFspVlanInterfaceIpv6Ip6DnsServerOverride3rdl(o["ip6-dns-server-override"], d, "ip6_dns_server_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-dns-server-override"], "ObjectFspVlanInterfaceIpv6-Ip6DnsServerOverride"); ok {
			if err = d.Set("ip6_dns_server_override", vv); err != nil {
				return fmt.Errorf("Error reading ip6_dns_server_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_dns_server_override: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip6_extra_addr", flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddr3rdl(o["ip6-extra-addr"], d, "ip6_extra_addr")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip6-extra-addr"], "ObjectFspVlanInterfaceIpv6-Ip6ExtraAddr"); ok {
				if err = d.Set("ip6_extra_addr", vv); err != nil {
					return fmt.Errorf("Error reading ip6_extra_addr: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip6_extra_addr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip6_extra_addr"); ok {
			if err = d.Set("ip6_extra_addr", flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddr3rdl(o["ip6-extra-addr"], d, "ip6_extra_addr")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip6-extra-addr"], "ObjectFspVlanInterfaceIpv6-Ip6ExtraAddr"); ok {
					if err = d.Set("ip6_extra_addr", vv); err != nil {
						return fmt.Errorf("Error reading ip6_extra_addr: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip6_extra_addr: %v", err)
				}
			}
		}
	}

	if err = d.Set("ip6_hop_limit", flattenObjectFspVlanInterfaceIpv6Ip6HopLimit3rdl(o["ip6-hop-limit"], d, "ip6_hop_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-hop-limit"], "ObjectFspVlanInterfaceIpv6-Ip6HopLimit"); ok {
			if err = d.Set("ip6_hop_limit", vv); err != nil {
				return fmt.Errorf("Error reading ip6_hop_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_hop_limit: %v", err)
		}
	}

	if err = d.Set("ip6_link_mtu", flattenObjectFspVlanInterfaceIpv6Ip6LinkMtu3rdl(o["ip6-link-mtu"], d, "ip6_link_mtu")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-link-mtu"], "ObjectFspVlanInterfaceIpv6-Ip6LinkMtu"); ok {
			if err = d.Set("ip6_link_mtu", vv); err != nil {
				return fmt.Errorf("Error reading ip6_link_mtu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_link_mtu: %v", err)
		}
	}

	if err = d.Set("ip6_manage_flag", flattenObjectFspVlanInterfaceIpv6Ip6ManageFlag3rdl(o["ip6-manage-flag"], d, "ip6_manage_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-manage-flag"], "ObjectFspVlanInterfaceIpv6-Ip6ManageFlag"); ok {
			if err = d.Set("ip6_manage_flag", vv); err != nil {
				return fmt.Errorf("Error reading ip6_manage_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_manage_flag: %v", err)
		}
	}

	if err = d.Set("ip6_max_interval", flattenObjectFspVlanInterfaceIpv6Ip6MaxInterval3rdl(o["ip6-max-interval"], d, "ip6_max_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-max-interval"], "ObjectFspVlanInterfaceIpv6-Ip6MaxInterval"); ok {
			if err = d.Set("ip6_max_interval", vv); err != nil {
				return fmt.Errorf("Error reading ip6_max_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_max_interval: %v", err)
		}
	}

	if err = d.Set("ip6_min_interval", flattenObjectFspVlanInterfaceIpv6Ip6MinInterval3rdl(o["ip6-min-interval"], d, "ip6_min_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-min-interval"], "ObjectFspVlanInterfaceIpv6-Ip6MinInterval"); ok {
			if err = d.Set("ip6_min_interval", vv); err != nil {
				return fmt.Errorf("Error reading ip6_min_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_min_interval: %v", err)
		}
	}

	if err = d.Set("ip6_mode", flattenObjectFspVlanInterfaceIpv6Ip6Mode3rdl(o["ip6-mode"], d, "ip6_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-mode"], "ObjectFspVlanInterfaceIpv6-Ip6Mode"); ok {
			if err = d.Set("ip6_mode", vv); err != nil {
				return fmt.Errorf("Error reading ip6_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_mode: %v", err)
		}
	}

	if err = d.Set("ip6_other_flag", flattenObjectFspVlanInterfaceIpv6Ip6OtherFlag3rdl(o["ip6-other-flag"], d, "ip6_other_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-other-flag"], "ObjectFspVlanInterfaceIpv6-Ip6OtherFlag"); ok {
			if err = d.Set("ip6_other_flag", vv); err != nil {
				return fmt.Errorf("Error reading ip6_other_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_other_flag: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip6_prefix_list", flattenObjectFspVlanInterfaceIpv6Ip6PrefixList3rdl(o["ip6-prefix-list"], d, "ip6_prefix_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip6-prefix-list"], "ObjectFspVlanInterfaceIpv6-Ip6PrefixList"); ok {
				if err = d.Set("ip6_prefix_list", vv); err != nil {
					return fmt.Errorf("Error reading ip6_prefix_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip6_prefix_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip6_prefix_list"); ok {
			if err = d.Set("ip6_prefix_list", flattenObjectFspVlanInterfaceIpv6Ip6PrefixList3rdl(o["ip6-prefix-list"], d, "ip6_prefix_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip6-prefix-list"], "ObjectFspVlanInterfaceIpv6-Ip6PrefixList"); ok {
					if err = d.Set("ip6_prefix_list", vv); err != nil {
						return fmt.Errorf("Error reading ip6_prefix_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip6_prefix_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("ip6_prefix_mode", flattenObjectFspVlanInterfaceIpv6Ip6PrefixMode3rdl(o["ip6-prefix-mode"], d, "ip6_prefix_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-prefix-mode"], "ObjectFspVlanInterfaceIpv6-Ip6PrefixMode"); ok {
			if err = d.Set("ip6_prefix_mode", vv); err != nil {
				return fmt.Errorf("Error reading ip6_prefix_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_prefix_mode: %v", err)
		}
	}

	if err = d.Set("ip6_reachable_time", flattenObjectFspVlanInterfaceIpv6Ip6ReachableTime3rdl(o["ip6-reachable-time"], d, "ip6_reachable_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-reachable-time"], "ObjectFspVlanInterfaceIpv6-Ip6ReachableTime"); ok {
			if err = d.Set("ip6_reachable_time", vv); err != nil {
				return fmt.Errorf("Error reading ip6_reachable_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_reachable_time: %v", err)
		}
	}

	if err = d.Set("ip6_retrans_time", flattenObjectFspVlanInterfaceIpv6Ip6RetransTime3rdl(o["ip6-retrans-time"], d, "ip6_retrans_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-retrans-time"], "ObjectFspVlanInterfaceIpv6-Ip6RetransTime"); ok {
			if err = d.Set("ip6_retrans_time", vv); err != nil {
				return fmt.Errorf("Error reading ip6_retrans_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_retrans_time: %v", err)
		}
	}

	if err = d.Set("ip6_send_adv", flattenObjectFspVlanInterfaceIpv6Ip6SendAdv3rdl(o["ip6-send-adv"], d, "ip6_send_adv")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-send-adv"], "ObjectFspVlanInterfaceIpv6-Ip6SendAdv"); ok {
			if err = d.Set("ip6_send_adv", vv); err != nil {
				return fmt.Errorf("Error reading ip6_send_adv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_send_adv: %v", err)
		}
	}

	if err = d.Set("ip6_subnet", flattenObjectFspVlanInterfaceIpv6Ip6Subnet3rdl(o["ip6-subnet"], d, "ip6_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-subnet"], "ObjectFspVlanInterfaceIpv6-Ip6Subnet"); ok {
			if err = d.Set("ip6_subnet", vv); err != nil {
				return fmt.Errorf("Error reading ip6_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_subnet: %v", err)
		}
	}

	if err = d.Set("ip6_upstream_interface", flattenObjectFspVlanInterfaceIpv6Ip6UpstreamInterface3rdl(o["ip6-upstream-interface"], d, "ip6_upstream_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-upstream-interface"], "ObjectFspVlanInterfaceIpv6-Ip6UpstreamInterface"); ok {
			if err = d.Set("ip6_upstream_interface", vv); err != nil {
				return fmt.Errorf("Error reading ip6_upstream_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_upstream_interface: %v", err)
		}
	}

	if err = d.Set("nd_cert", flattenObjectFspVlanInterfaceIpv6NdCert3rdl(o["nd-cert"], d, "nd_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["nd-cert"], "ObjectFspVlanInterfaceIpv6-NdCert"); ok {
			if err = d.Set("nd_cert", vv); err != nil {
				return fmt.Errorf("Error reading nd_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nd_cert: %v", err)
		}
	}

	if err = d.Set("nd_cga_modifier", flattenObjectFspVlanInterfaceIpv6NdCgaModifier3rdl(o["nd-cga-modifier"], d, "nd_cga_modifier")); err != nil {
		if vv, ok := fortiAPIPatch(o["nd-cga-modifier"], "ObjectFspVlanInterfaceIpv6-NdCgaModifier"); ok {
			if err = d.Set("nd_cga_modifier", vv); err != nil {
				return fmt.Errorf("Error reading nd_cga_modifier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nd_cga_modifier: %v", err)
		}
	}

	if err = d.Set("nd_mode", flattenObjectFspVlanInterfaceIpv6NdMode3rdl(o["nd-mode"], d, "nd_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["nd-mode"], "ObjectFspVlanInterfaceIpv6-NdMode"); ok {
			if err = d.Set("nd_mode", vv); err != nil {
				return fmt.Errorf("Error reading nd_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nd_mode: %v", err)
		}
	}

	if err = d.Set("nd_security_level", flattenObjectFspVlanInterfaceIpv6NdSecurityLevel3rdl(o["nd-security-level"], d, "nd_security_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["nd-security-level"], "ObjectFspVlanInterfaceIpv6-NdSecurityLevel"); ok {
			if err = d.Set("nd_security_level", vv); err != nil {
				return fmt.Errorf("Error reading nd_security_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nd_security_level: %v", err)
		}
	}

	if err = d.Set("nd_timestamp_delta", flattenObjectFspVlanInterfaceIpv6NdTimestampDelta3rdl(o["nd-timestamp-delta"], d, "nd_timestamp_delta")); err != nil {
		if vv, ok := fortiAPIPatch(o["nd-timestamp-delta"], "ObjectFspVlanInterfaceIpv6-NdTimestampDelta"); ok {
			if err = d.Set("nd_timestamp_delta", vv); err != nil {
				return fmt.Errorf("Error reading nd_timestamp_delta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nd_timestamp_delta: %v", err)
		}
	}

	if err = d.Set("nd_timestamp_fuzz", flattenObjectFspVlanInterfaceIpv6NdTimestampFuzz3rdl(o["nd-timestamp-fuzz"], d, "nd_timestamp_fuzz")); err != nil {
		if vv, ok := fortiAPIPatch(o["nd-timestamp-fuzz"], "ObjectFspVlanInterfaceIpv6-NdTimestampFuzz"); ok {
			if err = d.Set("nd_timestamp_fuzz", vv); err != nil {
				return fmt.Errorf("Error reading nd_timestamp_fuzz: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nd_timestamp_fuzz: %v", err)
		}
	}

	if err = d.Set("ra_send_mtu", flattenObjectFspVlanInterfaceIpv6RaSendMtu3rdl(o["ra-send-mtu"], d, "ra_send_mtu")); err != nil {
		if vv, ok := fortiAPIPatch(o["ra-send-mtu"], "ObjectFspVlanInterfaceIpv6-RaSendMtu"); ok {
			if err = d.Set("ra_send_mtu", vv); err != nil {
				return fmt.Errorf("Error reading ra_send_mtu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ra_send_mtu: %v", err)
		}
	}

	if err = d.Set("unique_autoconf_addr", flattenObjectFspVlanInterfaceIpv6UniqueAutoconfAddr3rdl(o["unique-autoconf-addr"], d, "unique_autoconf_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["unique-autoconf-addr"], "ObjectFspVlanInterfaceIpv6-UniqueAutoconfAddr"); ok {
			if err = d.Set("unique_autoconf_addr", vv); err != nil {
				return fmt.Errorf("Error reading unique_autoconf_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unique_autoconf_addr: %v", err)
		}
	}

	if err = d.Set("vrip6_link_local", flattenObjectFspVlanInterfaceIpv6Vrip6LinkLocal3rdl(o["vrip6_link_local"], d, "vrip6_link_local")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrip6_link_local"], "ObjectFspVlanInterfaceIpv6-Vrip6LinkLocal"); ok {
			if err = d.Set("vrip6_link_local", vv); err != nil {
				return fmt.Errorf("Error reading vrip6_link_local: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrip6_link_local: %v", err)
		}
	}

	if err = d.Set("vrrp_virtual_mac6", flattenObjectFspVlanInterfaceIpv6VrrpVirtualMac63rdl(o["vrrp-virtual-mac6"], d, "vrrp_virtual_mac6")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrrp-virtual-mac6"], "ObjectFspVlanInterfaceIpv6-VrrpVirtualMac6"); ok {
			if err = d.Set("vrrp_virtual_mac6", vv); err != nil {
				return fmt.Errorf("Error reading vrrp_virtual_mac6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrrp_virtual_mac6: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vrrp6", flattenObjectFspVlanInterfaceIpv6Vrrp63rdl(o["vrrp6"], d, "vrrp6")); err != nil {
			if vv, ok := fortiAPIPatch(o["vrrp6"], "ObjectFspVlanInterfaceIpv6-Vrrp6"); ok {
				if err = d.Set("vrrp6", vv); err != nil {
					return fmt.Errorf("Error reading vrrp6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vrrp6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vrrp6"); ok {
			if err = d.Set("vrrp6", flattenObjectFspVlanInterfaceIpv6Vrrp63rdl(o["vrrp6"], d, "vrrp6")); err != nil {
				if vv, ok := fortiAPIPatch(o["vrrp6"], "ObjectFspVlanInterfaceIpv6-Vrrp6"); ok {
					if err = d.Set("vrrp6", vv); err != nil {
						return fmt.Errorf("Error reading vrrp6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vrrp6: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectFspVlanInterfaceIpv6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFspVlanInterfaceIpv6Autoconf3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6CliConn6Status3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6ClientOptions3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6InformationRequest3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6PrefixDelegation3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHint3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHintPlt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHintVlt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6RelayInterfaceId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6RelayIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6RelayService3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6RelaySourceInterface3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6RelaySourceIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Dhcp6RelayType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Icmp6SendRedirect3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6InterfaceIdentifier3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6Address3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6Allowaccess3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DefaultLife3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixIaid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["autonomous-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag3rdl(d, i["autonomous_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delegated_prefix_iaid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["delegated-prefix-iaid"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid3rdl(d, i["delegated_prefix_iaid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["onlink-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag3rdl(d, i["onlink_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-id"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId3rdl(d, i["prefix_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss3rdl(d, i["rdnss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss-service"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService3rdl(d, i["rdnss_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet3rdl(d, i["subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upstream_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["upstream-interface"], _ = expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface3rdl(d, i["upstream_interface"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DnsServerOverride3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6ExtraAddr3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix3rdl(d, i["prefix"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6HopLimit3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6LinkMtu3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6ManageFlag3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6MaxInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6MinInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6Mode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6OtherFlag3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixList3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["autonomous-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListAutonomousFlag3rdl(d, i["autonomous_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dnssl"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListDnssl3rdl(d, i["dnssl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["onlink-flag"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListOnlinkFlag3rdl(d, i["onlink_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preferred-life-time"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListPreferredLifeTime3rdl(d, i["preferred_life_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListPrefix3rdl(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListRdnss3rdl(d, i["rdnss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["valid-life-time"], _ = expandObjectFspVlanInterfaceIpv6Ip6PrefixListValidLifeTime3rdl(d, i["valid_life_time"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListAutonomousFlag3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListDnssl3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListOnlinkFlag3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListPreferredLifeTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListPrefix3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListRdnss3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixListValidLifeTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6PrefixMode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6ReachableTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6RetransTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6SendAdv3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6Subnet3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6UpstreamInterface3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdCert3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdCgaModifier3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdMode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdSecurityLevel3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdTimestampDelta3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6NdTimestampFuzz3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6RaSendMtu3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6UniqueAutoconfAddr3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrip6LinkLocal3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6VrrpVirtualMac63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["accept-mode"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6AcceptMode3rdl(d, i["accept_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-interval"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6AdvInterval3rdl(d, i["adv_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preempt"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Preempt3rdl(d, i["preempt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Priority3rdl(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-time"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6StartTime3rdl(d, i["start_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Status3rdl(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrdst6"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Vrdst63rdl(d, i["vrdst6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrgrp"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Vrgrp3rdl(d, i["vrgrp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrid"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Vrid3rdl(d, i["vrid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrip6"], _ = expandObjectFspVlanInterfaceIpv6Vrrp6Vrip63rdl(d, i["vrip6"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6AcceptMode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6AdvInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Preempt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Priority3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6StartTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Status3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrdst63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrgrp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrip63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFspVlanInterfaceIpv6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("autoconf"); ok || d.HasChange("autoconf") {
		t, err := expandObjectFspVlanInterfaceIpv6Autoconf3rdl(d, v, "autoconf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["autoconf"] = t
		}
	}

	if v, ok := d.GetOk("cli_conn6_status"); ok || d.HasChange("cli_conn6_status") {
		t, err := expandObjectFspVlanInterfaceIpv6CliConn6Status3rdl(d, v, "cli_conn6_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cli-conn6-status"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_client_options"); ok || d.HasChange("dhcp6_client_options") {
		t, err := expandObjectFspVlanInterfaceIpv6Dhcp6ClientOptions3rdl(d, v, "dhcp6_client_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-client-options"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_information_request"); ok || d.HasChange("dhcp6_information_request") {
		t, err := expandObjectFspVlanInterfaceIpv6Dhcp6InformationRequest3rdl(d, v, "dhcp6_information_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-information-request"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_prefix_delegation"); ok || d.HasChange("dhcp6_prefix_delegation") {
		t, err := expandObjectFspVlanInterfaceIpv6Dhcp6PrefixDelegation3rdl(d, v, "dhcp6_prefix_delegation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-prefix-delegation"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_prefix_hint"); ok || d.HasChange("dhcp6_prefix_hint") {
		t, err := expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHint3rdl(d, v, "dhcp6_prefix_hint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-prefix-hint"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_prefix_hint_plt"); ok || d.HasChange("dhcp6_prefix_hint_plt") {
		t, err := expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHintPlt3rdl(d, v, "dhcp6_prefix_hint_plt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-prefix-hint-plt"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_prefix_hint_vlt"); ok || d.HasChange("dhcp6_prefix_hint_vlt") {
		t, err := expandObjectFspVlanInterfaceIpv6Dhcp6PrefixHintVlt3rdl(d, v, "dhcp6_prefix_hint_vlt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-prefix-hint-vlt"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_relay_interface_id"); ok || d.HasChange("dhcp6_relay_interface_id") {
		t, err := expandObjectFspVlanInterfaceIpv6Dhcp6RelayInterfaceId3rdl(d, v, "dhcp6_relay_interface_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-relay-interface-id"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_relay_ip"); ok || d.HasChange("dhcp6_relay_ip") {
		t, err := expandObjectFspVlanInterfaceIpv6Dhcp6RelayIp3rdl(d, v, "dhcp6_relay_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-relay-ip"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_relay_service"); ok || d.HasChange("dhcp6_relay_service") {
		t, err := expandObjectFspVlanInterfaceIpv6Dhcp6RelayService3rdl(d, v, "dhcp6_relay_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-relay-service"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_relay_source_interface"); ok || d.HasChange("dhcp6_relay_source_interface") {
		t, err := expandObjectFspVlanInterfaceIpv6Dhcp6RelaySourceInterface3rdl(d, v, "dhcp6_relay_source_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-relay-source-interface"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_relay_source_ip"); ok || d.HasChange("dhcp6_relay_source_ip") {
		t, err := expandObjectFspVlanInterfaceIpv6Dhcp6RelaySourceIp3rdl(d, v, "dhcp6_relay_source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-relay-source-ip"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_relay_type"); ok || d.HasChange("dhcp6_relay_type") {
		t, err := expandObjectFspVlanInterfaceIpv6Dhcp6RelayType3rdl(d, v, "dhcp6_relay_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-relay-type"] = t
		}
	}

	if v, ok := d.GetOk("icmp6_send_redirect"); ok || d.HasChange("icmp6_send_redirect") {
		t, err := expandObjectFspVlanInterfaceIpv6Icmp6SendRedirect3rdl(d, v, "icmp6_send_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp6-send-redirect"] = t
		}
	}

	if v, ok := d.GetOk("interface_identifier"); ok || d.HasChange("interface_identifier") {
		t, err := expandObjectFspVlanInterfaceIpv6InterfaceIdentifier3rdl(d, v, "interface_identifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-identifier"] = t
		}
	}

	if v, ok := d.GetOk("ip6_address"); ok || d.HasChange("ip6_address") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6Address3rdl(d, v, "ip6_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-address"] = t
		}
	}

	if v, ok := d.GetOk("ip6_allowaccess"); ok || d.HasChange("ip6_allowaccess") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6Allowaccess3rdl(d, v, "ip6_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("ip6_default_life"); ok || d.HasChange("ip6_default_life") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6DefaultLife3rdl(d, v, "ip6_default_life")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-default-life"] = t
		}
	}

	if v, ok := d.GetOk("ip6_delegated_prefix_iaid"); ok || d.HasChange("ip6_delegated_prefix_iaid") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixIaid3rdl(d, v, "ip6_delegated_prefix_iaid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-delegated-prefix-iaid"] = t
		}
	}

	if v, ok := d.GetOk("ip6_delegated_prefix_list"); ok || d.HasChange("ip6_delegated_prefix_list") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList3rdl(d, v, "ip6_delegated_prefix_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-delegated-prefix-list"] = t
		}
	}

	if v, ok := d.GetOk("ip6_dns_server_override"); ok || d.HasChange("ip6_dns_server_override") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6DnsServerOverride3rdl(d, v, "ip6_dns_server_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-dns-server-override"] = t
		}
	}

	if v, ok := d.GetOk("ip6_extra_addr"); ok || d.HasChange("ip6_extra_addr") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6ExtraAddr3rdl(d, v, "ip6_extra_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-extra-addr"] = t
		}
	}

	if v, ok := d.GetOk("ip6_hop_limit"); ok || d.HasChange("ip6_hop_limit") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6HopLimit3rdl(d, v, "ip6_hop_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-hop-limit"] = t
		}
	}

	if v, ok := d.GetOk("ip6_link_mtu"); ok || d.HasChange("ip6_link_mtu") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6LinkMtu3rdl(d, v, "ip6_link_mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-link-mtu"] = t
		}
	}

	if v, ok := d.GetOk("ip6_manage_flag"); ok || d.HasChange("ip6_manage_flag") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6ManageFlag3rdl(d, v, "ip6_manage_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-manage-flag"] = t
		}
	}

	if v, ok := d.GetOk("ip6_max_interval"); ok || d.HasChange("ip6_max_interval") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6MaxInterval3rdl(d, v, "ip6_max_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-max-interval"] = t
		}
	}

	if v, ok := d.GetOk("ip6_min_interval"); ok || d.HasChange("ip6_min_interval") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6MinInterval3rdl(d, v, "ip6_min_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-min-interval"] = t
		}
	}

	if v, ok := d.GetOk("ip6_mode"); ok || d.HasChange("ip6_mode") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6Mode3rdl(d, v, "ip6_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-mode"] = t
		}
	}

	if v, ok := d.GetOk("ip6_other_flag"); ok || d.HasChange("ip6_other_flag") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6OtherFlag3rdl(d, v, "ip6_other_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-other-flag"] = t
		}
	}

	if v, ok := d.GetOk("ip6_prefix_list"); ok || d.HasChange("ip6_prefix_list") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6PrefixList3rdl(d, v, "ip6_prefix_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-prefix-list"] = t
		}
	}

	if v, ok := d.GetOk("ip6_prefix_mode"); ok || d.HasChange("ip6_prefix_mode") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6PrefixMode3rdl(d, v, "ip6_prefix_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-prefix-mode"] = t
		}
	}

	if v, ok := d.GetOk("ip6_reachable_time"); ok || d.HasChange("ip6_reachable_time") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6ReachableTime3rdl(d, v, "ip6_reachable_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-reachable-time"] = t
		}
	}

	if v, ok := d.GetOk("ip6_retrans_time"); ok || d.HasChange("ip6_retrans_time") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6RetransTime3rdl(d, v, "ip6_retrans_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-retrans-time"] = t
		}
	}

	if v, ok := d.GetOk("ip6_send_adv"); ok || d.HasChange("ip6_send_adv") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6SendAdv3rdl(d, v, "ip6_send_adv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-send-adv"] = t
		}
	}

	if v, ok := d.GetOk("ip6_subnet"); ok || d.HasChange("ip6_subnet") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6Subnet3rdl(d, v, "ip6_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-subnet"] = t
		}
	}

	if v, ok := d.GetOk("ip6_upstream_interface"); ok || d.HasChange("ip6_upstream_interface") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6UpstreamInterface3rdl(d, v, "ip6_upstream_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-upstream-interface"] = t
		}
	}

	if v, ok := d.GetOk("nd_cert"); ok || d.HasChange("nd_cert") {
		t, err := expandObjectFspVlanInterfaceIpv6NdCert3rdl(d, v, "nd_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nd-cert"] = t
		}
	}

	if v, ok := d.GetOk("nd_cga_modifier"); ok || d.HasChange("nd_cga_modifier") {
		t, err := expandObjectFspVlanInterfaceIpv6NdCgaModifier3rdl(d, v, "nd_cga_modifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nd-cga-modifier"] = t
		}
	}

	if v, ok := d.GetOk("nd_mode"); ok || d.HasChange("nd_mode") {
		t, err := expandObjectFspVlanInterfaceIpv6NdMode3rdl(d, v, "nd_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nd-mode"] = t
		}
	}

	if v, ok := d.GetOk("nd_security_level"); ok || d.HasChange("nd_security_level") {
		t, err := expandObjectFspVlanInterfaceIpv6NdSecurityLevel3rdl(d, v, "nd_security_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nd-security-level"] = t
		}
	}

	if v, ok := d.GetOk("nd_timestamp_delta"); ok || d.HasChange("nd_timestamp_delta") {
		t, err := expandObjectFspVlanInterfaceIpv6NdTimestampDelta3rdl(d, v, "nd_timestamp_delta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nd-timestamp-delta"] = t
		}
	}

	if v, ok := d.GetOk("nd_timestamp_fuzz"); ok || d.HasChange("nd_timestamp_fuzz") {
		t, err := expandObjectFspVlanInterfaceIpv6NdTimestampFuzz3rdl(d, v, "nd_timestamp_fuzz")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nd-timestamp-fuzz"] = t
		}
	}

	if v, ok := d.GetOk("ra_send_mtu"); ok || d.HasChange("ra_send_mtu") {
		t, err := expandObjectFspVlanInterfaceIpv6RaSendMtu3rdl(d, v, "ra_send_mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ra-send-mtu"] = t
		}
	}

	if v, ok := d.GetOk("unique_autoconf_addr"); ok || d.HasChange("unique_autoconf_addr") {
		t, err := expandObjectFspVlanInterfaceIpv6UniqueAutoconfAddr3rdl(d, v, "unique_autoconf_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unique-autoconf-addr"] = t
		}
	}

	if v, ok := d.GetOk("vrip6_link_local"); ok || d.HasChange("vrip6_link_local") {
		t, err := expandObjectFspVlanInterfaceIpv6Vrip6LinkLocal3rdl(d, v, "vrip6_link_local")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrip6_link_local"] = t
		}
	}

	if v, ok := d.GetOk("vrrp_virtual_mac6"); ok || d.HasChange("vrrp_virtual_mac6") {
		t, err := expandObjectFspVlanInterfaceIpv6VrrpVirtualMac63rdl(d, v, "vrrp_virtual_mac6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrrp-virtual-mac6"] = t
		}
	}

	if v, ok := d.GetOk("vrrp6"); ok || d.HasChange("vrrp6") {
		t, err := expandObjectFspVlanInterfaceIpv6Vrrp63rdl(d, v, "vrrp6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrrp6"] = t
		}
	}

	return &obj, nil
}
