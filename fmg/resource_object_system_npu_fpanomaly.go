// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: NP6Lite anomaly protection (packet drop or send trap to host).

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSystemNpuFpAnomaly() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuFpAnomalyUpdate,
		Read:   resourceObjectSystemNpuFpAnomalyRead,
		Update: resourceObjectSystemNpuFpAnomalyUpdate,
		Delete: resourceObjectSystemNpuFpAnomalyDelete,

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
			"capwap_minlen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"esp_minlen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gre_csum_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtpu_plen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp_csum_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp_frag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp_land": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp_minlen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_csum_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_ihl_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_land": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_len_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_opt_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_optlsrr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_optrr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_optsecurity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_optssrr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_optstream": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_opttimestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_proto_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_ttlzero_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_unknopt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_ver_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_daddr_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_exthdr_len_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_exthdr_order_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_ihl_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_land": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_optendpid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_opthomeaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_optinvld": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_optjumbo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_optnsap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_optralert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_opttunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_plen_zero": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_proto_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_saddr_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_unknopt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_ver_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nvgre_minlen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sctp_clen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sctp_crc_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sctp_l4len_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_csum_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_fin_noack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_fin_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_hlen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_hlenvsl4len_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_land": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_no_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_plen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_syn_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_syn_fin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_winnuke": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_csum_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_hlen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_land": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_len_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_plen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udplite_cover_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udplite_csum_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uesp_minlen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknproto_minlen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vxlan_minlen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemNpuFpAnomalyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemNpuFpAnomaly(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuFpAnomaly resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuFpAnomaly(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuFpAnomaly resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuFpAnomaly")

	return resourceObjectSystemNpuFpAnomalyRead(d, m)
}

func resourceObjectSystemNpuFpAnomalyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemNpuFpAnomaly(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuFpAnomaly resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuFpAnomalyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemNpuFpAnomaly(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuFpAnomaly resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuFpAnomaly(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuFpAnomaly resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuFpAnomalyCapwapMinlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyEspMinlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyGreCsumErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyGtpuPlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpCsumErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpFrag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpLand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpMinlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4CsumErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4IhlErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Land(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4LenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4OptErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Optlsrr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Optrr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Optsecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Optssrr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Optstream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Opttimestamp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4ProtoErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4TtlzeroErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Unknopt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4VerErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6DaddrErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6ExthdrLenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6ExthdrOrderErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6IhlErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Land(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Optendpid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Opthomeaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Optinvld(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Optjumbo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Optnsap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Optralert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Opttunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6PlenZero(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6ProtoErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6SaddrErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Unknopt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6VerErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyNvgreMinlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalySctpClenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalySctpCrcErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalySctpL4LenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpCsumErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpFinNoack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpFinOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpHlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpHlenvsl4LenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpLand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpNoFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpPlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpSynData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpSynFin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpWinnuke(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpCsumErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpHlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpLand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpLenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpPlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpliteCoverErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpliteCsumErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUespMinlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUnknprotoMinlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyVxlanMinlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuFpAnomaly(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("capwap_minlen_err", flattenObjectSystemNpuFpAnomalyCapwapMinlenErr(o["capwap-minlen-err"], d, "capwap_minlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["capwap-minlen-err"], "ObjectSystemNpuFpAnomaly-CapwapMinlenErr"); ok {
			if err = d.Set("capwap_minlen_err", vv); err != nil {
				return fmt.Errorf("Error reading capwap_minlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capwap_minlen_err: %v", err)
		}
	}

	if err = d.Set("esp_minlen_err", flattenObjectSystemNpuFpAnomalyEspMinlenErr(o["esp-minlen-err"], d, "esp_minlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["esp-minlen-err"], "ObjectSystemNpuFpAnomaly-EspMinlenErr"); ok {
			if err = d.Set("esp_minlen_err", vv); err != nil {
				return fmt.Errorf("Error reading esp_minlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading esp_minlen_err: %v", err)
		}
	}

	if err = d.Set("gre_csum_err", flattenObjectSystemNpuFpAnomalyGreCsumErr(o["gre-csum-err"], d, "gre_csum_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["gre-csum-err"], "ObjectSystemNpuFpAnomaly-GreCsumErr"); ok {
			if err = d.Set("gre_csum_err", vv); err != nil {
				return fmt.Errorf("Error reading gre_csum_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gre_csum_err: %v", err)
		}
	}

	if err = d.Set("gtpu_plen_err", flattenObjectSystemNpuFpAnomalyGtpuPlenErr(o["gtpu-plen-err"], d, "gtpu_plen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtpu-plen-err"], "ObjectSystemNpuFpAnomaly-GtpuPlenErr"); ok {
			if err = d.Set("gtpu_plen_err", vv); err != nil {
				return fmt.Errorf("Error reading gtpu_plen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtpu_plen_err: %v", err)
		}
	}

	if err = d.Set("icmp_csum_err", flattenObjectSystemNpuFpAnomalyIcmpCsumErr(o["icmp-csum-err"], d, "icmp_csum_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-csum-err"], "ObjectSystemNpuFpAnomaly-IcmpCsumErr"); ok {
			if err = d.Set("icmp_csum_err", vv); err != nil {
				return fmt.Errorf("Error reading icmp_csum_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_csum_err: %v", err)
		}
	}

	if err = d.Set("icmp_frag", flattenObjectSystemNpuFpAnomalyIcmpFrag(o["icmp-frag"], d, "icmp_frag")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-frag"], "ObjectSystemNpuFpAnomaly-IcmpFrag"); ok {
			if err = d.Set("icmp_frag", vv); err != nil {
				return fmt.Errorf("Error reading icmp_frag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_frag: %v", err)
		}
	}

	if err = d.Set("icmp_land", flattenObjectSystemNpuFpAnomalyIcmpLand(o["icmp-land"], d, "icmp_land")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-land"], "ObjectSystemNpuFpAnomaly-IcmpLand"); ok {
			if err = d.Set("icmp_land", vv); err != nil {
				return fmt.Errorf("Error reading icmp_land: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_land: %v", err)
		}
	}

	if err = d.Set("icmp_minlen_err", flattenObjectSystemNpuFpAnomalyIcmpMinlenErr(o["icmp-minlen-err"], d, "icmp_minlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-minlen-err"], "ObjectSystemNpuFpAnomaly-IcmpMinlenErr"); ok {
			if err = d.Set("icmp_minlen_err", vv); err != nil {
				return fmt.Errorf("Error reading icmp_minlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_minlen_err: %v", err)
		}
	}

	if err = d.Set("ipv4_csum_err", flattenObjectSystemNpuFpAnomalyIpv4CsumErr(o["ipv4-csum-err"], d, "ipv4_csum_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-csum-err"], "ObjectSystemNpuFpAnomaly-Ipv4CsumErr"); ok {
			if err = d.Set("ipv4_csum_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_csum_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_csum_err: %v", err)
		}
	}

	if err = d.Set("ipv4_ihl_err", flattenObjectSystemNpuFpAnomalyIpv4IhlErr(o["ipv4-ihl-err"], d, "ipv4_ihl_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-ihl-err"], "ObjectSystemNpuFpAnomaly-Ipv4IhlErr"); ok {
			if err = d.Set("ipv4_ihl_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_ihl_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_ihl_err: %v", err)
		}
	}

	if err = d.Set("ipv4_land", flattenObjectSystemNpuFpAnomalyIpv4Land(o["ipv4-land"], d, "ipv4_land")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-land"], "ObjectSystemNpuFpAnomaly-Ipv4Land"); ok {
			if err = d.Set("ipv4_land", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_land: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_land: %v", err)
		}
	}

	if err = d.Set("ipv4_len_err", flattenObjectSystemNpuFpAnomalyIpv4LenErr(o["ipv4-len-err"], d, "ipv4_len_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-len-err"], "ObjectSystemNpuFpAnomaly-Ipv4LenErr"); ok {
			if err = d.Set("ipv4_len_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_len_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_len_err: %v", err)
		}
	}

	if err = d.Set("ipv4_opt_err", flattenObjectSystemNpuFpAnomalyIpv4OptErr(o["ipv4-opt-err"], d, "ipv4_opt_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-opt-err"], "ObjectSystemNpuFpAnomaly-Ipv4OptErr"); ok {
			if err = d.Set("ipv4_opt_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_opt_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_opt_err: %v", err)
		}
	}

	if err = d.Set("ipv4_optlsrr", flattenObjectSystemNpuFpAnomalyIpv4Optlsrr(o["ipv4-optlsrr"], d, "ipv4_optlsrr")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-optlsrr"], "ObjectSystemNpuFpAnomaly-Ipv4Optlsrr"); ok {
			if err = d.Set("ipv4_optlsrr", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_optlsrr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_optlsrr: %v", err)
		}
	}

	if err = d.Set("ipv4_optrr", flattenObjectSystemNpuFpAnomalyIpv4Optrr(o["ipv4-optrr"], d, "ipv4_optrr")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-optrr"], "ObjectSystemNpuFpAnomaly-Ipv4Optrr"); ok {
			if err = d.Set("ipv4_optrr", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_optrr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_optrr: %v", err)
		}
	}

	if err = d.Set("ipv4_optsecurity", flattenObjectSystemNpuFpAnomalyIpv4Optsecurity(o["ipv4-optsecurity"], d, "ipv4_optsecurity")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-optsecurity"], "ObjectSystemNpuFpAnomaly-Ipv4Optsecurity"); ok {
			if err = d.Set("ipv4_optsecurity", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_optsecurity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_optsecurity: %v", err)
		}
	}

	if err = d.Set("ipv4_optssrr", flattenObjectSystemNpuFpAnomalyIpv4Optssrr(o["ipv4-optssrr"], d, "ipv4_optssrr")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-optssrr"], "ObjectSystemNpuFpAnomaly-Ipv4Optssrr"); ok {
			if err = d.Set("ipv4_optssrr", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_optssrr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_optssrr: %v", err)
		}
	}

	if err = d.Set("ipv4_optstream", flattenObjectSystemNpuFpAnomalyIpv4Optstream(o["ipv4-optstream"], d, "ipv4_optstream")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-optstream"], "ObjectSystemNpuFpAnomaly-Ipv4Optstream"); ok {
			if err = d.Set("ipv4_optstream", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_optstream: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_optstream: %v", err)
		}
	}

	if err = d.Set("ipv4_opttimestamp", flattenObjectSystemNpuFpAnomalyIpv4Opttimestamp(o["ipv4-opttimestamp"], d, "ipv4_opttimestamp")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-opttimestamp"], "ObjectSystemNpuFpAnomaly-Ipv4Opttimestamp"); ok {
			if err = d.Set("ipv4_opttimestamp", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_opttimestamp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_opttimestamp: %v", err)
		}
	}

	if err = d.Set("ipv4_proto_err", flattenObjectSystemNpuFpAnomalyIpv4ProtoErr(o["ipv4-proto-err"], d, "ipv4_proto_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-proto-err"], "ObjectSystemNpuFpAnomaly-Ipv4ProtoErr"); ok {
			if err = d.Set("ipv4_proto_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_proto_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_proto_err: %v", err)
		}
	}

	if err = d.Set("ipv4_ttlzero_err", flattenObjectSystemNpuFpAnomalyIpv4TtlzeroErr(o["ipv4-ttlzero-err"], d, "ipv4_ttlzero_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-ttlzero-err"], "ObjectSystemNpuFpAnomaly-Ipv4TtlzeroErr"); ok {
			if err = d.Set("ipv4_ttlzero_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_ttlzero_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_ttlzero_err: %v", err)
		}
	}

	if err = d.Set("ipv4_unknopt", flattenObjectSystemNpuFpAnomalyIpv4Unknopt(o["ipv4-unknopt"], d, "ipv4_unknopt")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-unknopt"], "ObjectSystemNpuFpAnomaly-Ipv4Unknopt"); ok {
			if err = d.Set("ipv4_unknopt", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_unknopt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_unknopt: %v", err)
		}
	}

	if err = d.Set("ipv4_ver_err", flattenObjectSystemNpuFpAnomalyIpv4VerErr(o["ipv4-ver-err"], d, "ipv4_ver_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-ver-err"], "ObjectSystemNpuFpAnomaly-Ipv4VerErr"); ok {
			if err = d.Set("ipv4_ver_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_ver_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_ver_err: %v", err)
		}
	}

	if err = d.Set("ipv6_daddr_err", flattenObjectSystemNpuFpAnomalyIpv6DaddrErr(o["ipv6-daddr-err"], d, "ipv6_daddr_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-daddr-err"], "ObjectSystemNpuFpAnomaly-Ipv6DaddrErr"); ok {
			if err = d.Set("ipv6_daddr_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_daddr_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_daddr_err: %v", err)
		}
	}

	if err = d.Set("ipv6_exthdr_len_err", flattenObjectSystemNpuFpAnomalyIpv6ExthdrLenErr(o["ipv6-exthdr-len-err"], d, "ipv6_exthdr_len_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-exthdr-len-err"], "ObjectSystemNpuFpAnomaly-Ipv6ExthdrLenErr"); ok {
			if err = d.Set("ipv6_exthdr_len_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_exthdr_len_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_exthdr_len_err: %v", err)
		}
	}

	if err = d.Set("ipv6_exthdr_order_err", flattenObjectSystemNpuFpAnomalyIpv6ExthdrOrderErr(o["ipv6-exthdr-order-err"], d, "ipv6_exthdr_order_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-exthdr-order-err"], "ObjectSystemNpuFpAnomaly-Ipv6ExthdrOrderErr"); ok {
			if err = d.Set("ipv6_exthdr_order_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_exthdr_order_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_exthdr_order_err: %v", err)
		}
	}

	if err = d.Set("ipv6_ihl_err", flattenObjectSystemNpuFpAnomalyIpv6IhlErr(o["ipv6-ihl-err"], d, "ipv6_ihl_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-ihl-err"], "ObjectSystemNpuFpAnomaly-Ipv6IhlErr"); ok {
			if err = d.Set("ipv6_ihl_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_ihl_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_ihl_err: %v", err)
		}
	}

	if err = d.Set("ipv6_land", flattenObjectSystemNpuFpAnomalyIpv6Land(o["ipv6-land"], d, "ipv6_land")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-land"], "ObjectSystemNpuFpAnomaly-Ipv6Land"); ok {
			if err = d.Set("ipv6_land", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_land: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_land: %v", err)
		}
	}

	if err = d.Set("ipv6_optendpid", flattenObjectSystemNpuFpAnomalyIpv6Optendpid(o["ipv6-optendpid"], d, "ipv6_optendpid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-optendpid"], "ObjectSystemNpuFpAnomaly-Ipv6Optendpid"); ok {
			if err = d.Set("ipv6_optendpid", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_optendpid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_optendpid: %v", err)
		}
	}

	if err = d.Set("ipv6_opthomeaddr", flattenObjectSystemNpuFpAnomalyIpv6Opthomeaddr(o["ipv6-opthomeaddr"], d, "ipv6_opthomeaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-opthomeaddr"], "ObjectSystemNpuFpAnomaly-Ipv6Opthomeaddr"); ok {
			if err = d.Set("ipv6_opthomeaddr", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_opthomeaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_opthomeaddr: %v", err)
		}
	}

	if err = d.Set("ipv6_optinvld", flattenObjectSystemNpuFpAnomalyIpv6Optinvld(o["ipv6-optinvld"], d, "ipv6_optinvld")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-optinvld"], "ObjectSystemNpuFpAnomaly-Ipv6Optinvld"); ok {
			if err = d.Set("ipv6_optinvld", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_optinvld: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_optinvld: %v", err)
		}
	}

	if err = d.Set("ipv6_optjumbo", flattenObjectSystemNpuFpAnomalyIpv6Optjumbo(o["ipv6-optjumbo"], d, "ipv6_optjumbo")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-optjumbo"], "ObjectSystemNpuFpAnomaly-Ipv6Optjumbo"); ok {
			if err = d.Set("ipv6_optjumbo", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_optjumbo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_optjumbo: %v", err)
		}
	}

	if err = d.Set("ipv6_optnsap", flattenObjectSystemNpuFpAnomalyIpv6Optnsap(o["ipv6-optnsap"], d, "ipv6_optnsap")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-optnsap"], "ObjectSystemNpuFpAnomaly-Ipv6Optnsap"); ok {
			if err = d.Set("ipv6_optnsap", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_optnsap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_optnsap: %v", err)
		}
	}

	if err = d.Set("ipv6_optralert", flattenObjectSystemNpuFpAnomalyIpv6Optralert(o["ipv6-optralert"], d, "ipv6_optralert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-optralert"], "ObjectSystemNpuFpAnomaly-Ipv6Optralert"); ok {
			if err = d.Set("ipv6_optralert", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_optralert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_optralert: %v", err)
		}
	}

	if err = d.Set("ipv6_opttunnel", flattenObjectSystemNpuFpAnomalyIpv6Opttunnel(o["ipv6-opttunnel"], d, "ipv6_opttunnel")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-opttunnel"], "ObjectSystemNpuFpAnomaly-Ipv6Opttunnel"); ok {
			if err = d.Set("ipv6_opttunnel", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_opttunnel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_opttunnel: %v", err)
		}
	}

	if err = d.Set("ipv6_plen_zero", flattenObjectSystemNpuFpAnomalyIpv6PlenZero(o["ipv6-plen-zero"], d, "ipv6_plen_zero")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-plen-zero"], "ObjectSystemNpuFpAnomaly-Ipv6PlenZero"); ok {
			if err = d.Set("ipv6_plen_zero", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_plen_zero: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_plen_zero: %v", err)
		}
	}

	if err = d.Set("ipv6_proto_err", flattenObjectSystemNpuFpAnomalyIpv6ProtoErr(o["ipv6-proto-err"], d, "ipv6_proto_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-proto-err"], "ObjectSystemNpuFpAnomaly-Ipv6ProtoErr"); ok {
			if err = d.Set("ipv6_proto_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_proto_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_proto_err: %v", err)
		}
	}

	if err = d.Set("ipv6_saddr_err", flattenObjectSystemNpuFpAnomalyIpv6SaddrErr(o["ipv6-saddr-err"], d, "ipv6_saddr_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-saddr-err"], "ObjectSystemNpuFpAnomaly-Ipv6SaddrErr"); ok {
			if err = d.Set("ipv6_saddr_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_saddr_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_saddr_err: %v", err)
		}
	}

	if err = d.Set("ipv6_unknopt", flattenObjectSystemNpuFpAnomalyIpv6Unknopt(o["ipv6-unknopt"], d, "ipv6_unknopt")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-unknopt"], "ObjectSystemNpuFpAnomaly-Ipv6Unknopt"); ok {
			if err = d.Set("ipv6_unknopt", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_unknopt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_unknopt: %v", err)
		}
	}

	if err = d.Set("ipv6_ver_err", flattenObjectSystemNpuFpAnomalyIpv6VerErr(o["ipv6-ver-err"], d, "ipv6_ver_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-ver-err"], "ObjectSystemNpuFpAnomaly-Ipv6VerErr"); ok {
			if err = d.Set("ipv6_ver_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_ver_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_ver_err: %v", err)
		}
	}

	if err = d.Set("nvgre_minlen_err", flattenObjectSystemNpuFpAnomalyNvgreMinlenErr(o["nvgre-minlen-err"], d, "nvgre_minlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["nvgre-minlen-err"], "ObjectSystemNpuFpAnomaly-NvgreMinlenErr"); ok {
			if err = d.Set("nvgre_minlen_err", vv); err != nil {
				return fmt.Errorf("Error reading nvgre_minlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nvgre_minlen_err: %v", err)
		}
	}

	if err = d.Set("sctp_clen_err", flattenObjectSystemNpuFpAnomalySctpClenErr(o["sctp-clen-err"], d, "sctp_clen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["sctp-clen-err"], "ObjectSystemNpuFpAnomaly-SctpClenErr"); ok {
			if err = d.Set("sctp_clen_err", vv); err != nil {
				return fmt.Errorf("Error reading sctp_clen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sctp_clen_err: %v", err)
		}
	}

	if err = d.Set("sctp_crc_err", flattenObjectSystemNpuFpAnomalySctpCrcErr(o["sctp-crc-err"], d, "sctp_crc_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["sctp-crc-err"], "ObjectSystemNpuFpAnomaly-SctpCrcErr"); ok {
			if err = d.Set("sctp_crc_err", vv); err != nil {
				return fmt.Errorf("Error reading sctp_crc_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sctp_crc_err: %v", err)
		}
	}

	if err = d.Set("sctp_l4len_err", flattenObjectSystemNpuFpAnomalySctpL4LenErr(o["sctp-l4len-err"], d, "sctp_l4len_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["sctp-l4len-err"], "ObjectSystemNpuFpAnomaly-SctpL4LenErr"); ok {
			if err = d.Set("sctp_l4len_err", vv); err != nil {
				return fmt.Errorf("Error reading sctp_l4len_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sctp_l4len_err: %v", err)
		}
	}

	if err = d.Set("tcp_csum_err", flattenObjectSystemNpuFpAnomalyTcpCsumErr(o["tcp-csum-err"], d, "tcp_csum_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-csum-err"], "ObjectSystemNpuFpAnomaly-TcpCsumErr"); ok {
			if err = d.Set("tcp_csum_err", vv); err != nil {
				return fmt.Errorf("Error reading tcp_csum_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_csum_err: %v", err)
		}
	}

	if err = d.Set("tcp_fin_noack", flattenObjectSystemNpuFpAnomalyTcpFinNoack(o["tcp-fin-noack"], d, "tcp_fin_noack")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-fin-noack"], "ObjectSystemNpuFpAnomaly-TcpFinNoack"); ok {
			if err = d.Set("tcp_fin_noack", vv); err != nil {
				return fmt.Errorf("Error reading tcp_fin_noack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_fin_noack: %v", err)
		}
	}

	if err = d.Set("tcp_fin_only", flattenObjectSystemNpuFpAnomalyTcpFinOnly(o["tcp-fin-only"], d, "tcp_fin_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-fin-only"], "ObjectSystemNpuFpAnomaly-TcpFinOnly"); ok {
			if err = d.Set("tcp_fin_only", vv); err != nil {
				return fmt.Errorf("Error reading tcp_fin_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_fin_only: %v", err)
		}
	}

	if err = d.Set("tcp_hlen_err", flattenObjectSystemNpuFpAnomalyTcpHlenErr(o["tcp-hlen-err"], d, "tcp_hlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-hlen-err"], "ObjectSystemNpuFpAnomaly-TcpHlenErr"); ok {
			if err = d.Set("tcp_hlen_err", vv); err != nil {
				return fmt.Errorf("Error reading tcp_hlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_hlen_err: %v", err)
		}
	}

	if err = d.Set("tcp_hlenvsl4len_err", flattenObjectSystemNpuFpAnomalyTcpHlenvsl4LenErr(o["tcp-hlenvsl4len-err"], d, "tcp_hlenvsl4len_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-hlenvsl4len-err"], "ObjectSystemNpuFpAnomaly-TcpHlenvsl4LenErr"); ok {
			if err = d.Set("tcp_hlenvsl4len_err", vv); err != nil {
				return fmt.Errorf("Error reading tcp_hlenvsl4len_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_hlenvsl4len_err: %v", err)
		}
	}

	if err = d.Set("tcp_land", flattenObjectSystemNpuFpAnomalyTcpLand(o["tcp-land"], d, "tcp_land")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-land"], "ObjectSystemNpuFpAnomaly-TcpLand"); ok {
			if err = d.Set("tcp_land", vv); err != nil {
				return fmt.Errorf("Error reading tcp_land: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_land: %v", err)
		}
	}

	if err = d.Set("tcp_no_flag", flattenObjectSystemNpuFpAnomalyTcpNoFlag(o["tcp-no-flag"], d, "tcp_no_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-no-flag"], "ObjectSystemNpuFpAnomaly-TcpNoFlag"); ok {
			if err = d.Set("tcp_no_flag", vv); err != nil {
				return fmt.Errorf("Error reading tcp_no_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_no_flag: %v", err)
		}
	}

	if err = d.Set("tcp_plen_err", flattenObjectSystemNpuFpAnomalyTcpPlenErr(o["tcp-plen-err"], d, "tcp_plen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-plen-err"], "ObjectSystemNpuFpAnomaly-TcpPlenErr"); ok {
			if err = d.Set("tcp_plen_err", vv); err != nil {
				return fmt.Errorf("Error reading tcp_plen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_plen_err: %v", err)
		}
	}

	if err = d.Set("tcp_syn_data", flattenObjectSystemNpuFpAnomalyTcpSynData(o["tcp-syn-data"], d, "tcp_syn_data")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-syn-data"], "ObjectSystemNpuFpAnomaly-TcpSynData"); ok {
			if err = d.Set("tcp_syn_data", vv); err != nil {
				return fmt.Errorf("Error reading tcp_syn_data: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_syn_data: %v", err)
		}
	}

	if err = d.Set("tcp_syn_fin", flattenObjectSystemNpuFpAnomalyTcpSynFin(o["tcp-syn-fin"], d, "tcp_syn_fin")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-syn-fin"], "ObjectSystemNpuFpAnomaly-TcpSynFin"); ok {
			if err = d.Set("tcp_syn_fin", vv); err != nil {
				return fmt.Errorf("Error reading tcp_syn_fin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_syn_fin: %v", err)
		}
	}

	if err = d.Set("tcp_winnuke", flattenObjectSystemNpuFpAnomalyTcpWinnuke(o["tcp-winnuke"], d, "tcp_winnuke")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-winnuke"], "ObjectSystemNpuFpAnomaly-TcpWinnuke"); ok {
			if err = d.Set("tcp_winnuke", vv); err != nil {
				return fmt.Errorf("Error reading tcp_winnuke: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_winnuke: %v", err)
		}
	}

	if err = d.Set("udp_csum_err", flattenObjectSystemNpuFpAnomalyUdpCsumErr(o["udp-csum-err"], d, "udp_csum_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-csum-err"], "ObjectSystemNpuFpAnomaly-UdpCsumErr"); ok {
			if err = d.Set("udp_csum_err", vv); err != nil {
				return fmt.Errorf("Error reading udp_csum_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_csum_err: %v", err)
		}
	}

	if err = d.Set("udp_hlen_err", flattenObjectSystemNpuFpAnomalyUdpHlenErr(o["udp-hlen-err"], d, "udp_hlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-hlen-err"], "ObjectSystemNpuFpAnomaly-UdpHlenErr"); ok {
			if err = d.Set("udp_hlen_err", vv); err != nil {
				return fmt.Errorf("Error reading udp_hlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_hlen_err: %v", err)
		}
	}

	if err = d.Set("udp_land", flattenObjectSystemNpuFpAnomalyUdpLand(o["udp-land"], d, "udp_land")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-land"], "ObjectSystemNpuFpAnomaly-UdpLand"); ok {
			if err = d.Set("udp_land", vv); err != nil {
				return fmt.Errorf("Error reading udp_land: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_land: %v", err)
		}
	}

	if err = d.Set("udp_len_err", flattenObjectSystemNpuFpAnomalyUdpLenErr(o["udp-len-err"], d, "udp_len_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-len-err"], "ObjectSystemNpuFpAnomaly-UdpLenErr"); ok {
			if err = d.Set("udp_len_err", vv); err != nil {
				return fmt.Errorf("Error reading udp_len_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_len_err: %v", err)
		}
	}

	if err = d.Set("udp_plen_err", flattenObjectSystemNpuFpAnomalyUdpPlenErr(o["udp-plen-err"], d, "udp_plen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-plen-err"], "ObjectSystemNpuFpAnomaly-UdpPlenErr"); ok {
			if err = d.Set("udp_plen_err", vv); err != nil {
				return fmt.Errorf("Error reading udp_plen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_plen_err: %v", err)
		}
	}

	if err = d.Set("udplite_cover_err", flattenObjectSystemNpuFpAnomalyUdpliteCoverErr(o["udplite-cover-err"], d, "udplite_cover_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["udplite-cover-err"], "ObjectSystemNpuFpAnomaly-UdpliteCoverErr"); ok {
			if err = d.Set("udplite_cover_err", vv); err != nil {
				return fmt.Errorf("Error reading udplite_cover_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udplite_cover_err: %v", err)
		}
	}

	if err = d.Set("udplite_csum_err", flattenObjectSystemNpuFpAnomalyUdpliteCsumErr(o["udplite-csum-err"], d, "udplite_csum_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["udplite-csum-err"], "ObjectSystemNpuFpAnomaly-UdpliteCsumErr"); ok {
			if err = d.Set("udplite_csum_err", vv); err != nil {
				return fmt.Errorf("Error reading udplite_csum_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udplite_csum_err: %v", err)
		}
	}

	if err = d.Set("uesp_minlen_err", flattenObjectSystemNpuFpAnomalyUespMinlenErr(o["uesp-minlen-err"], d, "uesp_minlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["uesp-minlen-err"], "ObjectSystemNpuFpAnomaly-UespMinlenErr"); ok {
			if err = d.Set("uesp_minlen_err", vv); err != nil {
				return fmt.Errorf("Error reading uesp_minlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uesp_minlen_err: %v", err)
		}
	}

	if err = d.Set("unknproto_minlen_err", flattenObjectSystemNpuFpAnomalyUnknprotoMinlenErr(o["unknproto-minlen-err"], d, "unknproto_minlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknproto-minlen-err"], "ObjectSystemNpuFpAnomaly-UnknprotoMinlenErr"); ok {
			if err = d.Set("unknproto_minlen_err", vv); err != nil {
				return fmt.Errorf("Error reading unknproto_minlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknproto_minlen_err: %v", err)
		}
	}

	if err = d.Set("vxlan_minlen_err", flattenObjectSystemNpuFpAnomalyVxlanMinlenErr(o["vxlan-minlen-err"], d, "vxlan_minlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["vxlan-minlen-err"], "ObjectSystemNpuFpAnomaly-VxlanMinlenErr"); ok {
			if err = d.Set("vxlan_minlen_err", vv); err != nil {
				return fmt.Errorf("Error reading vxlan_minlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vxlan_minlen_err: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuFpAnomalyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuFpAnomalyCapwapMinlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyEspMinlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyGreCsumErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyGtpuPlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpCsumErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpFrag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpLand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpMinlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4CsumErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4IhlErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Land(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4LenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4OptErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Optlsrr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Optrr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Optsecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Optssrr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Optstream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Opttimestamp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4ProtoErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4TtlzeroErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Unknopt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4VerErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6DaddrErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6ExthdrLenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6ExthdrOrderErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6IhlErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Land(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Optendpid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Opthomeaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Optinvld(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Optjumbo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Optnsap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Optralert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Opttunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6PlenZero(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6ProtoErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6SaddrErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Unknopt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6VerErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyNvgreMinlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalySctpClenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalySctpCrcErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalySctpL4LenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpCsumErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpFinNoack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpFinOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpHlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpHlenvsl4LenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpLand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpNoFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpPlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpSynData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpSynFin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpWinnuke(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpCsumErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpHlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpLand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpLenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpPlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpliteCoverErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpliteCsumErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUespMinlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUnknprotoMinlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyVxlanMinlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuFpAnomaly(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("capwap_minlen_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyCapwapMinlenErr(d, v, "capwap_minlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capwap-minlen-err"] = t
		}
	}

	if v, ok := d.GetOk("esp_minlen_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyEspMinlenErr(d, v, "esp_minlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["esp-minlen-err"] = t
		}
	}

	if v, ok := d.GetOk("gre_csum_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyGreCsumErr(d, v, "gre_csum_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gre-csum-err"] = t
		}
	}

	if v, ok := d.GetOk("gtpu_plen_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyGtpuPlenErr(d, v, "gtpu_plen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtpu-plen-err"] = t
		}
	}

	if v, ok := d.GetOk("icmp_csum_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIcmpCsumErr(d, v, "icmp_csum_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-csum-err"] = t
		}
	}

	if v, ok := d.GetOk("icmp_frag"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIcmpFrag(d, v, "icmp_frag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-frag"] = t
		}
	}

	if v, ok := d.GetOk("icmp_land"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIcmpLand(d, v, "icmp_land")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-land"] = t
		}
	}

	if v, ok := d.GetOk("icmp_minlen_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIcmpMinlenErr(d, v, "icmp_minlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-minlen-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_csum_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv4CsumErr(d, v, "ipv4_csum_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-csum-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_ihl_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv4IhlErr(d, v, "ipv4_ihl_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-ihl-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_land"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Land(d, v, "ipv4_land")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-land"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_len_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv4LenErr(d, v, "ipv4_len_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-len-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_opt_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv4OptErr(d, v, "ipv4_opt_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-opt-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_optlsrr"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Optlsrr(d, v, "ipv4_optlsrr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-optlsrr"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_optrr"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Optrr(d, v, "ipv4_optrr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-optrr"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_optsecurity"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Optsecurity(d, v, "ipv4_optsecurity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-optsecurity"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_optssrr"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Optssrr(d, v, "ipv4_optssrr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-optssrr"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_optstream"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Optstream(d, v, "ipv4_optstream")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-optstream"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_opttimestamp"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Opttimestamp(d, v, "ipv4_opttimestamp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-opttimestamp"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_proto_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv4ProtoErr(d, v, "ipv4_proto_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-proto-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_ttlzero_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv4TtlzeroErr(d, v, "ipv4_ttlzero_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-ttlzero-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_unknopt"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Unknopt(d, v, "ipv4_unknopt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-unknopt"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_ver_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv4VerErr(d, v, "ipv4_ver_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-ver-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_daddr_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6DaddrErr(d, v, "ipv6_daddr_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-daddr-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_exthdr_len_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6ExthdrLenErr(d, v, "ipv6_exthdr_len_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-exthdr-len-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_exthdr_order_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6ExthdrOrderErr(d, v, "ipv6_exthdr_order_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-exthdr-order-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_ihl_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6IhlErr(d, v, "ipv6_ihl_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-ihl-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_land"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Land(d, v, "ipv6_land")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-land"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_optendpid"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Optendpid(d, v, "ipv6_optendpid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-optendpid"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_opthomeaddr"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Opthomeaddr(d, v, "ipv6_opthomeaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-opthomeaddr"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_optinvld"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Optinvld(d, v, "ipv6_optinvld")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-optinvld"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_optjumbo"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Optjumbo(d, v, "ipv6_optjumbo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-optjumbo"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_optnsap"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Optnsap(d, v, "ipv6_optnsap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-optnsap"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_optralert"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Optralert(d, v, "ipv6_optralert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-optralert"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_opttunnel"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Opttunnel(d, v, "ipv6_opttunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-opttunnel"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_plen_zero"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6PlenZero(d, v, "ipv6_plen_zero")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-plen-zero"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_proto_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6ProtoErr(d, v, "ipv6_proto_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-proto-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_saddr_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6SaddrErr(d, v, "ipv6_saddr_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-saddr-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_unknopt"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Unknopt(d, v, "ipv6_unknopt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-unknopt"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_ver_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyIpv6VerErr(d, v, "ipv6_ver_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-ver-err"] = t
		}
	}

	if v, ok := d.GetOk("nvgre_minlen_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyNvgreMinlenErr(d, v, "nvgre_minlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nvgre-minlen-err"] = t
		}
	}

	if v, ok := d.GetOk("sctp_clen_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalySctpClenErr(d, v, "sctp_clen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-clen-err"] = t
		}
	}

	if v, ok := d.GetOk("sctp_crc_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalySctpCrcErr(d, v, "sctp_crc_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-crc-err"] = t
		}
	}

	if v, ok := d.GetOk("sctp_l4len_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalySctpL4LenErr(d, v, "sctp_l4len_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-l4len-err"] = t
		}
	}

	if v, ok := d.GetOk("tcp_csum_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyTcpCsumErr(d, v, "tcp_csum_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-csum-err"] = t
		}
	}

	if v, ok := d.GetOk("tcp_fin_noack"); ok {
		t, err := expandObjectSystemNpuFpAnomalyTcpFinNoack(d, v, "tcp_fin_noack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-fin-noack"] = t
		}
	}

	if v, ok := d.GetOk("tcp_fin_only"); ok {
		t, err := expandObjectSystemNpuFpAnomalyTcpFinOnly(d, v, "tcp_fin_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-fin-only"] = t
		}
	}

	if v, ok := d.GetOk("tcp_hlen_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyTcpHlenErr(d, v, "tcp_hlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-hlen-err"] = t
		}
	}

	if v, ok := d.GetOk("tcp_hlenvsl4len_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyTcpHlenvsl4LenErr(d, v, "tcp_hlenvsl4len_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-hlenvsl4len-err"] = t
		}
	}

	if v, ok := d.GetOk("tcp_land"); ok {
		t, err := expandObjectSystemNpuFpAnomalyTcpLand(d, v, "tcp_land")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-land"] = t
		}
	}

	if v, ok := d.GetOk("tcp_no_flag"); ok {
		t, err := expandObjectSystemNpuFpAnomalyTcpNoFlag(d, v, "tcp_no_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-no-flag"] = t
		}
	}

	if v, ok := d.GetOk("tcp_plen_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyTcpPlenErr(d, v, "tcp_plen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-plen-err"] = t
		}
	}

	if v, ok := d.GetOk("tcp_syn_data"); ok {
		t, err := expandObjectSystemNpuFpAnomalyTcpSynData(d, v, "tcp_syn_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-syn-data"] = t
		}
	}

	if v, ok := d.GetOk("tcp_syn_fin"); ok {
		t, err := expandObjectSystemNpuFpAnomalyTcpSynFin(d, v, "tcp_syn_fin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-syn-fin"] = t
		}
	}

	if v, ok := d.GetOk("tcp_winnuke"); ok {
		t, err := expandObjectSystemNpuFpAnomalyTcpWinnuke(d, v, "tcp_winnuke")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-winnuke"] = t
		}
	}

	if v, ok := d.GetOk("udp_csum_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyUdpCsumErr(d, v, "udp_csum_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-csum-err"] = t
		}
	}

	if v, ok := d.GetOk("udp_hlen_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyUdpHlenErr(d, v, "udp_hlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-hlen-err"] = t
		}
	}

	if v, ok := d.GetOk("udp_land"); ok {
		t, err := expandObjectSystemNpuFpAnomalyUdpLand(d, v, "udp_land")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-land"] = t
		}
	}

	if v, ok := d.GetOk("udp_len_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyUdpLenErr(d, v, "udp_len_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-len-err"] = t
		}
	}

	if v, ok := d.GetOk("udp_plen_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyUdpPlenErr(d, v, "udp_plen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-plen-err"] = t
		}
	}

	if v, ok := d.GetOk("udplite_cover_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyUdpliteCoverErr(d, v, "udplite_cover_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udplite-cover-err"] = t
		}
	}

	if v, ok := d.GetOk("udplite_csum_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyUdpliteCsumErr(d, v, "udplite_csum_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udplite-csum-err"] = t
		}
	}

	if v, ok := d.GetOk("uesp_minlen_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyUespMinlenErr(d, v, "uesp_minlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uesp-minlen-err"] = t
		}
	}

	if v, ok := d.GetOk("unknproto_minlen_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyUnknprotoMinlenErr(d, v, "unknproto_minlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknproto-minlen-err"] = t
		}
	}

	if v, ok := d.GetOk("vxlan_minlen_err"); ok {
		t, err := expandObjectSystemNpuFpAnomalyVxlanMinlenErr(d, v, "vxlan_minlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vxlan-minlen-err"] = t
		}
	}

	return &obj, nil
}
