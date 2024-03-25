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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			},
			"esp_minlen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gre_csum_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gtpu_plen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"sctp_clen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sctp_crc_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sctp_l4len_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"unknproto_minlen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vxlan_minlen_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemNpuFpAnomalyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuFpAnomaly(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuFpAnomaly resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuFpAnomaly(obj, mkey, paradict)
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	err = c.DeleteObjectSystemNpuFpAnomaly(mkey, paradict)
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	o, err := c.ReadObjectSystemNpuFpAnomaly(mkey, paradict)
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

func flattenObjectSystemNpuFpAnomalyCapwapMinlenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyEspMinlenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyGreCsumErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyGtpuPlenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpCsumErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpFrag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpLand2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpMinlenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4CsumErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4IhlErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Land2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4LenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4OptErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Optlsrr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Optrr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Optsecurity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Optssrr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Optstream2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Opttimestamp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4ProtoErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4TtlzeroErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Unknopt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4VerErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6DaddrErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6ExthdrLenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6ExthdrOrderErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6IhlErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Land2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Optendpid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Opthomeaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Optinvld2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Optjumbo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Optnsap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Optralert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Opttunnel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6PlenZero2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6ProtoErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6SaddrErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Unknopt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6VerErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyNvgreMinlenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalySctpClenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalySctpCrcErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalySctpL4LenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpCsumErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpFinNoack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpFinOnly2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpHlenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpHlenvsl4LenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpLand2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpNoFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpPlenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpSynData2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpSynFin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpWinnuke2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpCsumErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpHlenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpLand2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpLenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpPlenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpliteCoverErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpliteCsumErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUespMinlenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUnknprotoMinlenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyVxlanMinlenErr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuFpAnomaly(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("capwap_minlen_err", flattenObjectSystemNpuFpAnomalyCapwapMinlenErr2edl(o["capwap-minlen-err"], d, "capwap_minlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["capwap-minlen-err"], "ObjectSystemNpuFpAnomaly-CapwapMinlenErr"); ok {
			if err = d.Set("capwap_minlen_err", vv); err != nil {
				return fmt.Errorf("Error reading capwap_minlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capwap_minlen_err: %v", err)
		}
	}

	if err = d.Set("esp_minlen_err", flattenObjectSystemNpuFpAnomalyEspMinlenErr2edl(o["esp-minlen-err"], d, "esp_minlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["esp-minlen-err"], "ObjectSystemNpuFpAnomaly-EspMinlenErr"); ok {
			if err = d.Set("esp_minlen_err", vv); err != nil {
				return fmt.Errorf("Error reading esp_minlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading esp_minlen_err: %v", err)
		}
	}

	if err = d.Set("gre_csum_err", flattenObjectSystemNpuFpAnomalyGreCsumErr2edl(o["gre-csum-err"], d, "gre_csum_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["gre-csum-err"], "ObjectSystemNpuFpAnomaly-GreCsumErr"); ok {
			if err = d.Set("gre_csum_err", vv); err != nil {
				return fmt.Errorf("Error reading gre_csum_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gre_csum_err: %v", err)
		}
	}

	if err = d.Set("gtpu_plen_err", flattenObjectSystemNpuFpAnomalyGtpuPlenErr2edl(o["gtpu-plen-err"], d, "gtpu_plen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtpu-plen-err"], "ObjectSystemNpuFpAnomaly-GtpuPlenErr"); ok {
			if err = d.Set("gtpu_plen_err", vv); err != nil {
				return fmt.Errorf("Error reading gtpu_plen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtpu_plen_err: %v", err)
		}
	}

	if err = d.Set("icmp_csum_err", flattenObjectSystemNpuFpAnomalyIcmpCsumErr2edl(o["icmp-csum-err"], d, "icmp_csum_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-csum-err"], "ObjectSystemNpuFpAnomaly-IcmpCsumErr"); ok {
			if err = d.Set("icmp_csum_err", vv); err != nil {
				return fmt.Errorf("Error reading icmp_csum_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_csum_err: %v", err)
		}
	}

	if err = d.Set("icmp_frag", flattenObjectSystemNpuFpAnomalyIcmpFrag2edl(o["icmp-frag"], d, "icmp_frag")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-frag"], "ObjectSystemNpuFpAnomaly-IcmpFrag"); ok {
			if err = d.Set("icmp_frag", vv); err != nil {
				return fmt.Errorf("Error reading icmp_frag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_frag: %v", err)
		}
	}

	if err = d.Set("icmp_land", flattenObjectSystemNpuFpAnomalyIcmpLand2edl(o["icmp-land"], d, "icmp_land")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-land"], "ObjectSystemNpuFpAnomaly-IcmpLand"); ok {
			if err = d.Set("icmp_land", vv); err != nil {
				return fmt.Errorf("Error reading icmp_land: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_land: %v", err)
		}
	}

	if err = d.Set("icmp_minlen_err", flattenObjectSystemNpuFpAnomalyIcmpMinlenErr2edl(o["icmp-minlen-err"], d, "icmp_minlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-minlen-err"], "ObjectSystemNpuFpAnomaly-IcmpMinlenErr"); ok {
			if err = d.Set("icmp_minlen_err", vv); err != nil {
				return fmt.Errorf("Error reading icmp_minlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_minlen_err: %v", err)
		}
	}

	if err = d.Set("ipv4_csum_err", flattenObjectSystemNpuFpAnomalyIpv4CsumErr2edl(o["ipv4-csum-err"], d, "ipv4_csum_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-csum-err"], "ObjectSystemNpuFpAnomaly-Ipv4CsumErr"); ok {
			if err = d.Set("ipv4_csum_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_csum_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_csum_err: %v", err)
		}
	}

	if err = d.Set("ipv4_ihl_err", flattenObjectSystemNpuFpAnomalyIpv4IhlErr2edl(o["ipv4-ihl-err"], d, "ipv4_ihl_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-ihl-err"], "ObjectSystemNpuFpAnomaly-Ipv4IhlErr"); ok {
			if err = d.Set("ipv4_ihl_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_ihl_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_ihl_err: %v", err)
		}
	}

	if err = d.Set("ipv4_land", flattenObjectSystemNpuFpAnomalyIpv4Land2edl(o["ipv4-land"], d, "ipv4_land")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-land"], "ObjectSystemNpuFpAnomaly-Ipv4Land"); ok {
			if err = d.Set("ipv4_land", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_land: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_land: %v", err)
		}
	}

	if err = d.Set("ipv4_len_err", flattenObjectSystemNpuFpAnomalyIpv4LenErr2edl(o["ipv4-len-err"], d, "ipv4_len_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-len-err"], "ObjectSystemNpuFpAnomaly-Ipv4LenErr"); ok {
			if err = d.Set("ipv4_len_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_len_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_len_err: %v", err)
		}
	}

	if err = d.Set("ipv4_opt_err", flattenObjectSystemNpuFpAnomalyIpv4OptErr2edl(o["ipv4-opt-err"], d, "ipv4_opt_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-opt-err"], "ObjectSystemNpuFpAnomaly-Ipv4OptErr"); ok {
			if err = d.Set("ipv4_opt_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_opt_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_opt_err: %v", err)
		}
	}

	if err = d.Set("ipv4_optlsrr", flattenObjectSystemNpuFpAnomalyIpv4Optlsrr2edl(o["ipv4-optlsrr"], d, "ipv4_optlsrr")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-optlsrr"], "ObjectSystemNpuFpAnomaly-Ipv4Optlsrr"); ok {
			if err = d.Set("ipv4_optlsrr", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_optlsrr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_optlsrr: %v", err)
		}
	}

	if err = d.Set("ipv4_optrr", flattenObjectSystemNpuFpAnomalyIpv4Optrr2edl(o["ipv4-optrr"], d, "ipv4_optrr")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-optrr"], "ObjectSystemNpuFpAnomaly-Ipv4Optrr"); ok {
			if err = d.Set("ipv4_optrr", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_optrr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_optrr: %v", err)
		}
	}

	if err = d.Set("ipv4_optsecurity", flattenObjectSystemNpuFpAnomalyIpv4Optsecurity2edl(o["ipv4-optsecurity"], d, "ipv4_optsecurity")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-optsecurity"], "ObjectSystemNpuFpAnomaly-Ipv4Optsecurity"); ok {
			if err = d.Set("ipv4_optsecurity", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_optsecurity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_optsecurity: %v", err)
		}
	}

	if err = d.Set("ipv4_optssrr", flattenObjectSystemNpuFpAnomalyIpv4Optssrr2edl(o["ipv4-optssrr"], d, "ipv4_optssrr")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-optssrr"], "ObjectSystemNpuFpAnomaly-Ipv4Optssrr"); ok {
			if err = d.Set("ipv4_optssrr", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_optssrr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_optssrr: %v", err)
		}
	}

	if err = d.Set("ipv4_optstream", flattenObjectSystemNpuFpAnomalyIpv4Optstream2edl(o["ipv4-optstream"], d, "ipv4_optstream")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-optstream"], "ObjectSystemNpuFpAnomaly-Ipv4Optstream"); ok {
			if err = d.Set("ipv4_optstream", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_optstream: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_optstream: %v", err)
		}
	}

	if err = d.Set("ipv4_opttimestamp", flattenObjectSystemNpuFpAnomalyIpv4Opttimestamp2edl(o["ipv4-opttimestamp"], d, "ipv4_opttimestamp")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-opttimestamp"], "ObjectSystemNpuFpAnomaly-Ipv4Opttimestamp"); ok {
			if err = d.Set("ipv4_opttimestamp", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_opttimestamp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_opttimestamp: %v", err)
		}
	}

	if err = d.Set("ipv4_proto_err", flattenObjectSystemNpuFpAnomalyIpv4ProtoErr2edl(o["ipv4-proto-err"], d, "ipv4_proto_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-proto-err"], "ObjectSystemNpuFpAnomaly-Ipv4ProtoErr"); ok {
			if err = d.Set("ipv4_proto_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_proto_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_proto_err: %v", err)
		}
	}

	if err = d.Set("ipv4_ttlzero_err", flattenObjectSystemNpuFpAnomalyIpv4TtlzeroErr2edl(o["ipv4-ttlzero-err"], d, "ipv4_ttlzero_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-ttlzero-err"], "ObjectSystemNpuFpAnomaly-Ipv4TtlzeroErr"); ok {
			if err = d.Set("ipv4_ttlzero_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_ttlzero_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_ttlzero_err: %v", err)
		}
	}

	if err = d.Set("ipv4_unknopt", flattenObjectSystemNpuFpAnomalyIpv4Unknopt2edl(o["ipv4-unknopt"], d, "ipv4_unknopt")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-unknopt"], "ObjectSystemNpuFpAnomaly-Ipv4Unknopt"); ok {
			if err = d.Set("ipv4_unknopt", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_unknopt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_unknopt: %v", err)
		}
	}

	if err = d.Set("ipv4_ver_err", flattenObjectSystemNpuFpAnomalyIpv4VerErr2edl(o["ipv4-ver-err"], d, "ipv4_ver_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-ver-err"], "ObjectSystemNpuFpAnomaly-Ipv4VerErr"); ok {
			if err = d.Set("ipv4_ver_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_ver_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_ver_err: %v", err)
		}
	}

	if err = d.Set("ipv6_daddr_err", flattenObjectSystemNpuFpAnomalyIpv6DaddrErr2edl(o["ipv6-daddr-err"], d, "ipv6_daddr_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-daddr-err"], "ObjectSystemNpuFpAnomaly-Ipv6DaddrErr"); ok {
			if err = d.Set("ipv6_daddr_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_daddr_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_daddr_err: %v", err)
		}
	}

	if err = d.Set("ipv6_exthdr_len_err", flattenObjectSystemNpuFpAnomalyIpv6ExthdrLenErr2edl(o["ipv6-exthdr-len-err"], d, "ipv6_exthdr_len_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-exthdr-len-err"], "ObjectSystemNpuFpAnomaly-Ipv6ExthdrLenErr"); ok {
			if err = d.Set("ipv6_exthdr_len_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_exthdr_len_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_exthdr_len_err: %v", err)
		}
	}

	if err = d.Set("ipv6_exthdr_order_err", flattenObjectSystemNpuFpAnomalyIpv6ExthdrOrderErr2edl(o["ipv6-exthdr-order-err"], d, "ipv6_exthdr_order_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-exthdr-order-err"], "ObjectSystemNpuFpAnomaly-Ipv6ExthdrOrderErr"); ok {
			if err = d.Set("ipv6_exthdr_order_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_exthdr_order_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_exthdr_order_err: %v", err)
		}
	}

	if err = d.Set("ipv6_ihl_err", flattenObjectSystemNpuFpAnomalyIpv6IhlErr2edl(o["ipv6-ihl-err"], d, "ipv6_ihl_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-ihl-err"], "ObjectSystemNpuFpAnomaly-Ipv6IhlErr"); ok {
			if err = d.Set("ipv6_ihl_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_ihl_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_ihl_err: %v", err)
		}
	}

	if err = d.Set("ipv6_land", flattenObjectSystemNpuFpAnomalyIpv6Land2edl(o["ipv6-land"], d, "ipv6_land")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-land"], "ObjectSystemNpuFpAnomaly-Ipv6Land"); ok {
			if err = d.Set("ipv6_land", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_land: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_land: %v", err)
		}
	}

	if err = d.Set("ipv6_optendpid", flattenObjectSystemNpuFpAnomalyIpv6Optendpid2edl(o["ipv6-optendpid"], d, "ipv6_optendpid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-optendpid"], "ObjectSystemNpuFpAnomaly-Ipv6Optendpid"); ok {
			if err = d.Set("ipv6_optendpid", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_optendpid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_optendpid: %v", err)
		}
	}

	if err = d.Set("ipv6_opthomeaddr", flattenObjectSystemNpuFpAnomalyIpv6Opthomeaddr2edl(o["ipv6-opthomeaddr"], d, "ipv6_opthomeaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-opthomeaddr"], "ObjectSystemNpuFpAnomaly-Ipv6Opthomeaddr"); ok {
			if err = d.Set("ipv6_opthomeaddr", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_opthomeaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_opthomeaddr: %v", err)
		}
	}

	if err = d.Set("ipv6_optinvld", flattenObjectSystemNpuFpAnomalyIpv6Optinvld2edl(o["ipv6-optinvld"], d, "ipv6_optinvld")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-optinvld"], "ObjectSystemNpuFpAnomaly-Ipv6Optinvld"); ok {
			if err = d.Set("ipv6_optinvld", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_optinvld: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_optinvld: %v", err)
		}
	}

	if err = d.Set("ipv6_optjumbo", flattenObjectSystemNpuFpAnomalyIpv6Optjumbo2edl(o["ipv6-optjumbo"], d, "ipv6_optjumbo")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-optjumbo"], "ObjectSystemNpuFpAnomaly-Ipv6Optjumbo"); ok {
			if err = d.Set("ipv6_optjumbo", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_optjumbo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_optjumbo: %v", err)
		}
	}

	if err = d.Set("ipv6_optnsap", flattenObjectSystemNpuFpAnomalyIpv6Optnsap2edl(o["ipv6-optnsap"], d, "ipv6_optnsap")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-optnsap"], "ObjectSystemNpuFpAnomaly-Ipv6Optnsap"); ok {
			if err = d.Set("ipv6_optnsap", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_optnsap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_optnsap: %v", err)
		}
	}

	if err = d.Set("ipv6_optralert", flattenObjectSystemNpuFpAnomalyIpv6Optralert2edl(o["ipv6-optralert"], d, "ipv6_optralert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-optralert"], "ObjectSystemNpuFpAnomaly-Ipv6Optralert"); ok {
			if err = d.Set("ipv6_optralert", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_optralert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_optralert: %v", err)
		}
	}

	if err = d.Set("ipv6_opttunnel", flattenObjectSystemNpuFpAnomalyIpv6Opttunnel2edl(o["ipv6-opttunnel"], d, "ipv6_opttunnel")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-opttunnel"], "ObjectSystemNpuFpAnomaly-Ipv6Opttunnel"); ok {
			if err = d.Set("ipv6_opttunnel", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_opttunnel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_opttunnel: %v", err)
		}
	}

	if err = d.Set("ipv6_plen_zero", flattenObjectSystemNpuFpAnomalyIpv6PlenZero2edl(o["ipv6-plen-zero"], d, "ipv6_plen_zero")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-plen-zero"], "ObjectSystemNpuFpAnomaly-Ipv6PlenZero"); ok {
			if err = d.Set("ipv6_plen_zero", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_plen_zero: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_plen_zero: %v", err)
		}
	}

	if err = d.Set("ipv6_proto_err", flattenObjectSystemNpuFpAnomalyIpv6ProtoErr2edl(o["ipv6-proto-err"], d, "ipv6_proto_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-proto-err"], "ObjectSystemNpuFpAnomaly-Ipv6ProtoErr"); ok {
			if err = d.Set("ipv6_proto_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_proto_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_proto_err: %v", err)
		}
	}

	if err = d.Set("ipv6_saddr_err", flattenObjectSystemNpuFpAnomalyIpv6SaddrErr2edl(o["ipv6-saddr-err"], d, "ipv6_saddr_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-saddr-err"], "ObjectSystemNpuFpAnomaly-Ipv6SaddrErr"); ok {
			if err = d.Set("ipv6_saddr_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_saddr_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_saddr_err: %v", err)
		}
	}

	if err = d.Set("ipv6_unknopt", flattenObjectSystemNpuFpAnomalyIpv6Unknopt2edl(o["ipv6-unknopt"], d, "ipv6_unknopt")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-unknopt"], "ObjectSystemNpuFpAnomaly-Ipv6Unknopt"); ok {
			if err = d.Set("ipv6_unknopt", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_unknopt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_unknopt: %v", err)
		}
	}

	if err = d.Set("ipv6_ver_err", flattenObjectSystemNpuFpAnomalyIpv6VerErr2edl(o["ipv6-ver-err"], d, "ipv6_ver_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-ver-err"], "ObjectSystemNpuFpAnomaly-Ipv6VerErr"); ok {
			if err = d.Set("ipv6_ver_err", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_ver_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_ver_err: %v", err)
		}
	}

	if err = d.Set("nvgre_minlen_err", flattenObjectSystemNpuFpAnomalyNvgreMinlenErr2edl(o["nvgre-minlen-err"], d, "nvgre_minlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["nvgre-minlen-err"], "ObjectSystemNpuFpAnomaly-NvgreMinlenErr"); ok {
			if err = d.Set("nvgre_minlen_err", vv); err != nil {
				return fmt.Errorf("Error reading nvgre_minlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nvgre_minlen_err: %v", err)
		}
	}

	if err = d.Set("sctp_clen_err", flattenObjectSystemNpuFpAnomalySctpClenErr2edl(o["sctp-clen-err"], d, "sctp_clen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["sctp-clen-err"], "ObjectSystemNpuFpAnomaly-SctpClenErr"); ok {
			if err = d.Set("sctp_clen_err", vv); err != nil {
				return fmt.Errorf("Error reading sctp_clen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sctp_clen_err: %v", err)
		}
	}

	if err = d.Set("sctp_crc_err", flattenObjectSystemNpuFpAnomalySctpCrcErr2edl(o["sctp-crc-err"], d, "sctp_crc_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["sctp-crc-err"], "ObjectSystemNpuFpAnomaly-SctpCrcErr"); ok {
			if err = d.Set("sctp_crc_err", vv); err != nil {
				return fmt.Errorf("Error reading sctp_crc_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sctp_crc_err: %v", err)
		}
	}

	if err = d.Set("sctp_l4len_err", flattenObjectSystemNpuFpAnomalySctpL4LenErr2edl(o["sctp-l4len-err"], d, "sctp_l4len_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["sctp-l4len-err"], "ObjectSystemNpuFpAnomaly-SctpL4LenErr"); ok {
			if err = d.Set("sctp_l4len_err", vv); err != nil {
				return fmt.Errorf("Error reading sctp_l4len_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sctp_l4len_err: %v", err)
		}
	}

	if err = d.Set("tcp_csum_err", flattenObjectSystemNpuFpAnomalyTcpCsumErr2edl(o["tcp-csum-err"], d, "tcp_csum_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-csum-err"], "ObjectSystemNpuFpAnomaly-TcpCsumErr"); ok {
			if err = d.Set("tcp_csum_err", vv); err != nil {
				return fmt.Errorf("Error reading tcp_csum_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_csum_err: %v", err)
		}
	}

	if err = d.Set("tcp_fin_noack", flattenObjectSystemNpuFpAnomalyTcpFinNoack2edl(o["tcp-fin-noack"], d, "tcp_fin_noack")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-fin-noack"], "ObjectSystemNpuFpAnomaly-TcpFinNoack"); ok {
			if err = d.Set("tcp_fin_noack", vv); err != nil {
				return fmt.Errorf("Error reading tcp_fin_noack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_fin_noack: %v", err)
		}
	}

	if err = d.Set("tcp_fin_only", flattenObjectSystemNpuFpAnomalyTcpFinOnly2edl(o["tcp-fin-only"], d, "tcp_fin_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-fin-only"], "ObjectSystemNpuFpAnomaly-TcpFinOnly"); ok {
			if err = d.Set("tcp_fin_only", vv); err != nil {
				return fmt.Errorf("Error reading tcp_fin_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_fin_only: %v", err)
		}
	}

	if err = d.Set("tcp_hlen_err", flattenObjectSystemNpuFpAnomalyTcpHlenErr2edl(o["tcp-hlen-err"], d, "tcp_hlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-hlen-err"], "ObjectSystemNpuFpAnomaly-TcpHlenErr"); ok {
			if err = d.Set("tcp_hlen_err", vv); err != nil {
				return fmt.Errorf("Error reading tcp_hlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_hlen_err: %v", err)
		}
	}

	if err = d.Set("tcp_hlenvsl4len_err", flattenObjectSystemNpuFpAnomalyTcpHlenvsl4LenErr2edl(o["tcp-hlenvsl4len-err"], d, "tcp_hlenvsl4len_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-hlenvsl4len-err"], "ObjectSystemNpuFpAnomaly-TcpHlenvsl4LenErr"); ok {
			if err = d.Set("tcp_hlenvsl4len_err", vv); err != nil {
				return fmt.Errorf("Error reading tcp_hlenvsl4len_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_hlenvsl4len_err: %v", err)
		}
	}

	if err = d.Set("tcp_land", flattenObjectSystemNpuFpAnomalyTcpLand2edl(o["tcp-land"], d, "tcp_land")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-land"], "ObjectSystemNpuFpAnomaly-TcpLand"); ok {
			if err = d.Set("tcp_land", vv); err != nil {
				return fmt.Errorf("Error reading tcp_land: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_land: %v", err)
		}
	}

	if err = d.Set("tcp_no_flag", flattenObjectSystemNpuFpAnomalyTcpNoFlag2edl(o["tcp-no-flag"], d, "tcp_no_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-no-flag"], "ObjectSystemNpuFpAnomaly-TcpNoFlag"); ok {
			if err = d.Set("tcp_no_flag", vv); err != nil {
				return fmt.Errorf("Error reading tcp_no_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_no_flag: %v", err)
		}
	}

	if err = d.Set("tcp_plen_err", flattenObjectSystemNpuFpAnomalyTcpPlenErr2edl(o["tcp-plen-err"], d, "tcp_plen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-plen-err"], "ObjectSystemNpuFpAnomaly-TcpPlenErr"); ok {
			if err = d.Set("tcp_plen_err", vv); err != nil {
				return fmt.Errorf("Error reading tcp_plen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_plen_err: %v", err)
		}
	}

	if err = d.Set("tcp_syn_data", flattenObjectSystemNpuFpAnomalyTcpSynData2edl(o["tcp-syn-data"], d, "tcp_syn_data")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-syn-data"], "ObjectSystemNpuFpAnomaly-TcpSynData"); ok {
			if err = d.Set("tcp_syn_data", vv); err != nil {
				return fmt.Errorf("Error reading tcp_syn_data: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_syn_data: %v", err)
		}
	}

	if err = d.Set("tcp_syn_fin", flattenObjectSystemNpuFpAnomalyTcpSynFin2edl(o["tcp-syn-fin"], d, "tcp_syn_fin")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-syn-fin"], "ObjectSystemNpuFpAnomaly-TcpSynFin"); ok {
			if err = d.Set("tcp_syn_fin", vv); err != nil {
				return fmt.Errorf("Error reading tcp_syn_fin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_syn_fin: %v", err)
		}
	}

	if err = d.Set("tcp_winnuke", flattenObjectSystemNpuFpAnomalyTcpWinnuke2edl(o["tcp-winnuke"], d, "tcp_winnuke")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-winnuke"], "ObjectSystemNpuFpAnomaly-TcpWinnuke"); ok {
			if err = d.Set("tcp_winnuke", vv); err != nil {
				return fmt.Errorf("Error reading tcp_winnuke: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_winnuke: %v", err)
		}
	}

	if err = d.Set("udp_csum_err", flattenObjectSystemNpuFpAnomalyUdpCsumErr2edl(o["udp-csum-err"], d, "udp_csum_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-csum-err"], "ObjectSystemNpuFpAnomaly-UdpCsumErr"); ok {
			if err = d.Set("udp_csum_err", vv); err != nil {
				return fmt.Errorf("Error reading udp_csum_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_csum_err: %v", err)
		}
	}

	if err = d.Set("udp_hlen_err", flattenObjectSystemNpuFpAnomalyUdpHlenErr2edl(o["udp-hlen-err"], d, "udp_hlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-hlen-err"], "ObjectSystemNpuFpAnomaly-UdpHlenErr"); ok {
			if err = d.Set("udp_hlen_err", vv); err != nil {
				return fmt.Errorf("Error reading udp_hlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_hlen_err: %v", err)
		}
	}

	if err = d.Set("udp_land", flattenObjectSystemNpuFpAnomalyUdpLand2edl(o["udp-land"], d, "udp_land")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-land"], "ObjectSystemNpuFpAnomaly-UdpLand"); ok {
			if err = d.Set("udp_land", vv); err != nil {
				return fmt.Errorf("Error reading udp_land: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_land: %v", err)
		}
	}

	if err = d.Set("udp_len_err", flattenObjectSystemNpuFpAnomalyUdpLenErr2edl(o["udp-len-err"], d, "udp_len_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-len-err"], "ObjectSystemNpuFpAnomaly-UdpLenErr"); ok {
			if err = d.Set("udp_len_err", vv); err != nil {
				return fmt.Errorf("Error reading udp_len_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_len_err: %v", err)
		}
	}

	if err = d.Set("udp_plen_err", flattenObjectSystemNpuFpAnomalyUdpPlenErr2edl(o["udp-plen-err"], d, "udp_plen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-plen-err"], "ObjectSystemNpuFpAnomaly-UdpPlenErr"); ok {
			if err = d.Set("udp_plen_err", vv); err != nil {
				return fmt.Errorf("Error reading udp_plen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_plen_err: %v", err)
		}
	}

	if err = d.Set("udplite_cover_err", flattenObjectSystemNpuFpAnomalyUdpliteCoverErr2edl(o["udplite-cover-err"], d, "udplite_cover_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["udplite-cover-err"], "ObjectSystemNpuFpAnomaly-UdpliteCoverErr"); ok {
			if err = d.Set("udplite_cover_err", vv); err != nil {
				return fmt.Errorf("Error reading udplite_cover_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udplite_cover_err: %v", err)
		}
	}

	if err = d.Set("udplite_csum_err", flattenObjectSystemNpuFpAnomalyUdpliteCsumErr2edl(o["udplite-csum-err"], d, "udplite_csum_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["udplite-csum-err"], "ObjectSystemNpuFpAnomaly-UdpliteCsumErr"); ok {
			if err = d.Set("udplite_csum_err", vv); err != nil {
				return fmt.Errorf("Error reading udplite_csum_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udplite_csum_err: %v", err)
		}
	}

	if err = d.Set("uesp_minlen_err", flattenObjectSystemNpuFpAnomalyUespMinlenErr2edl(o["uesp-minlen-err"], d, "uesp_minlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["uesp-minlen-err"], "ObjectSystemNpuFpAnomaly-UespMinlenErr"); ok {
			if err = d.Set("uesp_minlen_err", vv); err != nil {
				return fmt.Errorf("Error reading uesp_minlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uesp_minlen_err: %v", err)
		}
	}

	if err = d.Set("unknproto_minlen_err", flattenObjectSystemNpuFpAnomalyUnknprotoMinlenErr2edl(o["unknproto-minlen-err"], d, "unknproto_minlen_err")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknproto-minlen-err"], "ObjectSystemNpuFpAnomaly-UnknprotoMinlenErr"); ok {
			if err = d.Set("unknproto_minlen_err", vv); err != nil {
				return fmt.Errorf("Error reading unknproto_minlen_err: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknproto_minlen_err: %v", err)
		}
	}

	if err = d.Set("vxlan_minlen_err", flattenObjectSystemNpuFpAnomalyVxlanMinlenErr2edl(o["vxlan-minlen-err"], d, "vxlan_minlen_err")); err != nil {
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

func expandObjectSystemNpuFpAnomalyCapwapMinlenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyEspMinlenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyGreCsumErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyGtpuPlenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpCsumErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpFrag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpLand2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpMinlenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4CsumErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4IhlErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Land2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4LenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4OptErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Optlsrr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Optrr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Optsecurity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Optssrr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Optstream2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Opttimestamp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4ProtoErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4TtlzeroErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Unknopt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4VerErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6DaddrErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6ExthdrLenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6ExthdrOrderErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6IhlErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Land2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Optendpid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Opthomeaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Optinvld2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Optjumbo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Optnsap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Optralert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Opttunnel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6PlenZero2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6ProtoErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6SaddrErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Unknopt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6VerErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyNvgreMinlenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalySctpClenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalySctpCrcErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalySctpL4LenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpCsumErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpFinNoack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpFinOnly2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpHlenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpHlenvsl4LenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpLand2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpNoFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpPlenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpSynData2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpSynFin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpWinnuke2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpCsumErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpHlenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpLand2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpLenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpPlenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpliteCoverErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpliteCsumErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUespMinlenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUnknprotoMinlenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyVxlanMinlenErr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuFpAnomaly(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("capwap_minlen_err"); ok || d.HasChange("capwap_minlen_err") {
		t, err := expandObjectSystemNpuFpAnomalyCapwapMinlenErr2edl(d, v, "capwap_minlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capwap-minlen-err"] = t
		}
	}

	if v, ok := d.GetOk("esp_minlen_err"); ok || d.HasChange("esp_minlen_err") {
		t, err := expandObjectSystemNpuFpAnomalyEspMinlenErr2edl(d, v, "esp_minlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["esp-minlen-err"] = t
		}
	}

	if v, ok := d.GetOk("gre_csum_err"); ok || d.HasChange("gre_csum_err") {
		t, err := expandObjectSystemNpuFpAnomalyGreCsumErr2edl(d, v, "gre_csum_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gre-csum-err"] = t
		}
	}

	if v, ok := d.GetOk("gtpu_plen_err"); ok || d.HasChange("gtpu_plen_err") {
		t, err := expandObjectSystemNpuFpAnomalyGtpuPlenErr2edl(d, v, "gtpu_plen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtpu-plen-err"] = t
		}
	}

	if v, ok := d.GetOk("icmp_csum_err"); ok || d.HasChange("icmp_csum_err") {
		t, err := expandObjectSystemNpuFpAnomalyIcmpCsumErr2edl(d, v, "icmp_csum_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-csum-err"] = t
		}
	}

	if v, ok := d.GetOk("icmp_frag"); ok || d.HasChange("icmp_frag") {
		t, err := expandObjectSystemNpuFpAnomalyIcmpFrag2edl(d, v, "icmp_frag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-frag"] = t
		}
	}

	if v, ok := d.GetOk("icmp_land"); ok || d.HasChange("icmp_land") {
		t, err := expandObjectSystemNpuFpAnomalyIcmpLand2edl(d, v, "icmp_land")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-land"] = t
		}
	}

	if v, ok := d.GetOk("icmp_minlen_err"); ok || d.HasChange("icmp_minlen_err") {
		t, err := expandObjectSystemNpuFpAnomalyIcmpMinlenErr2edl(d, v, "icmp_minlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-minlen-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_csum_err"); ok || d.HasChange("ipv4_csum_err") {
		t, err := expandObjectSystemNpuFpAnomalyIpv4CsumErr2edl(d, v, "ipv4_csum_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-csum-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_ihl_err"); ok || d.HasChange("ipv4_ihl_err") {
		t, err := expandObjectSystemNpuFpAnomalyIpv4IhlErr2edl(d, v, "ipv4_ihl_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-ihl-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_land"); ok || d.HasChange("ipv4_land") {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Land2edl(d, v, "ipv4_land")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-land"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_len_err"); ok || d.HasChange("ipv4_len_err") {
		t, err := expandObjectSystemNpuFpAnomalyIpv4LenErr2edl(d, v, "ipv4_len_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-len-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_opt_err"); ok || d.HasChange("ipv4_opt_err") {
		t, err := expandObjectSystemNpuFpAnomalyIpv4OptErr2edl(d, v, "ipv4_opt_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-opt-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_optlsrr"); ok || d.HasChange("ipv4_optlsrr") {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Optlsrr2edl(d, v, "ipv4_optlsrr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-optlsrr"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_optrr"); ok || d.HasChange("ipv4_optrr") {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Optrr2edl(d, v, "ipv4_optrr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-optrr"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_optsecurity"); ok || d.HasChange("ipv4_optsecurity") {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Optsecurity2edl(d, v, "ipv4_optsecurity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-optsecurity"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_optssrr"); ok || d.HasChange("ipv4_optssrr") {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Optssrr2edl(d, v, "ipv4_optssrr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-optssrr"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_optstream"); ok || d.HasChange("ipv4_optstream") {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Optstream2edl(d, v, "ipv4_optstream")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-optstream"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_opttimestamp"); ok || d.HasChange("ipv4_opttimestamp") {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Opttimestamp2edl(d, v, "ipv4_opttimestamp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-opttimestamp"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_proto_err"); ok || d.HasChange("ipv4_proto_err") {
		t, err := expandObjectSystemNpuFpAnomalyIpv4ProtoErr2edl(d, v, "ipv4_proto_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-proto-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_ttlzero_err"); ok || d.HasChange("ipv4_ttlzero_err") {
		t, err := expandObjectSystemNpuFpAnomalyIpv4TtlzeroErr2edl(d, v, "ipv4_ttlzero_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-ttlzero-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_unknopt"); ok || d.HasChange("ipv4_unknopt") {
		t, err := expandObjectSystemNpuFpAnomalyIpv4Unknopt2edl(d, v, "ipv4_unknopt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-unknopt"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_ver_err"); ok || d.HasChange("ipv4_ver_err") {
		t, err := expandObjectSystemNpuFpAnomalyIpv4VerErr2edl(d, v, "ipv4_ver_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-ver-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_daddr_err"); ok || d.HasChange("ipv6_daddr_err") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6DaddrErr2edl(d, v, "ipv6_daddr_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-daddr-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_exthdr_len_err"); ok || d.HasChange("ipv6_exthdr_len_err") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6ExthdrLenErr2edl(d, v, "ipv6_exthdr_len_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-exthdr-len-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_exthdr_order_err"); ok || d.HasChange("ipv6_exthdr_order_err") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6ExthdrOrderErr2edl(d, v, "ipv6_exthdr_order_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-exthdr-order-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_ihl_err"); ok || d.HasChange("ipv6_ihl_err") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6IhlErr2edl(d, v, "ipv6_ihl_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-ihl-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_land"); ok || d.HasChange("ipv6_land") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Land2edl(d, v, "ipv6_land")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-land"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_optendpid"); ok || d.HasChange("ipv6_optendpid") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Optendpid2edl(d, v, "ipv6_optendpid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-optendpid"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_opthomeaddr"); ok || d.HasChange("ipv6_opthomeaddr") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Opthomeaddr2edl(d, v, "ipv6_opthomeaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-opthomeaddr"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_optinvld"); ok || d.HasChange("ipv6_optinvld") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Optinvld2edl(d, v, "ipv6_optinvld")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-optinvld"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_optjumbo"); ok || d.HasChange("ipv6_optjumbo") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Optjumbo2edl(d, v, "ipv6_optjumbo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-optjumbo"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_optnsap"); ok || d.HasChange("ipv6_optnsap") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Optnsap2edl(d, v, "ipv6_optnsap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-optnsap"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_optralert"); ok || d.HasChange("ipv6_optralert") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Optralert2edl(d, v, "ipv6_optralert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-optralert"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_opttunnel"); ok || d.HasChange("ipv6_opttunnel") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Opttunnel2edl(d, v, "ipv6_opttunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-opttunnel"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_plen_zero"); ok || d.HasChange("ipv6_plen_zero") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6PlenZero2edl(d, v, "ipv6_plen_zero")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-plen-zero"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_proto_err"); ok || d.HasChange("ipv6_proto_err") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6ProtoErr2edl(d, v, "ipv6_proto_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-proto-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_saddr_err"); ok || d.HasChange("ipv6_saddr_err") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6SaddrErr2edl(d, v, "ipv6_saddr_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-saddr-err"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_unknopt"); ok || d.HasChange("ipv6_unknopt") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6Unknopt2edl(d, v, "ipv6_unknopt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-unknopt"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_ver_err"); ok || d.HasChange("ipv6_ver_err") {
		t, err := expandObjectSystemNpuFpAnomalyIpv6VerErr2edl(d, v, "ipv6_ver_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-ver-err"] = t
		}
	}

	if v, ok := d.GetOk("nvgre_minlen_err"); ok || d.HasChange("nvgre_minlen_err") {
		t, err := expandObjectSystemNpuFpAnomalyNvgreMinlenErr2edl(d, v, "nvgre_minlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nvgre-minlen-err"] = t
		}
	}

	if v, ok := d.GetOk("sctp_clen_err"); ok || d.HasChange("sctp_clen_err") {
		t, err := expandObjectSystemNpuFpAnomalySctpClenErr2edl(d, v, "sctp_clen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-clen-err"] = t
		}
	}

	if v, ok := d.GetOk("sctp_crc_err"); ok || d.HasChange("sctp_crc_err") {
		t, err := expandObjectSystemNpuFpAnomalySctpCrcErr2edl(d, v, "sctp_crc_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-crc-err"] = t
		}
	}

	if v, ok := d.GetOk("sctp_l4len_err"); ok || d.HasChange("sctp_l4len_err") {
		t, err := expandObjectSystemNpuFpAnomalySctpL4LenErr2edl(d, v, "sctp_l4len_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-l4len-err"] = t
		}
	}

	if v, ok := d.GetOk("tcp_csum_err"); ok || d.HasChange("tcp_csum_err") {
		t, err := expandObjectSystemNpuFpAnomalyTcpCsumErr2edl(d, v, "tcp_csum_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-csum-err"] = t
		}
	}

	if v, ok := d.GetOk("tcp_fin_noack"); ok || d.HasChange("tcp_fin_noack") {
		t, err := expandObjectSystemNpuFpAnomalyTcpFinNoack2edl(d, v, "tcp_fin_noack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-fin-noack"] = t
		}
	}

	if v, ok := d.GetOk("tcp_fin_only"); ok || d.HasChange("tcp_fin_only") {
		t, err := expandObjectSystemNpuFpAnomalyTcpFinOnly2edl(d, v, "tcp_fin_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-fin-only"] = t
		}
	}

	if v, ok := d.GetOk("tcp_hlen_err"); ok || d.HasChange("tcp_hlen_err") {
		t, err := expandObjectSystemNpuFpAnomalyTcpHlenErr2edl(d, v, "tcp_hlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-hlen-err"] = t
		}
	}

	if v, ok := d.GetOk("tcp_hlenvsl4len_err"); ok || d.HasChange("tcp_hlenvsl4len_err") {
		t, err := expandObjectSystemNpuFpAnomalyTcpHlenvsl4LenErr2edl(d, v, "tcp_hlenvsl4len_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-hlenvsl4len-err"] = t
		}
	}

	if v, ok := d.GetOk("tcp_land"); ok || d.HasChange("tcp_land") {
		t, err := expandObjectSystemNpuFpAnomalyTcpLand2edl(d, v, "tcp_land")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-land"] = t
		}
	}

	if v, ok := d.GetOk("tcp_no_flag"); ok || d.HasChange("tcp_no_flag") {
		t, err := expandObjectSystemNpuFpAnomalyTcpNoFlag2edl(d, v, "tcp_no_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-no-flag"] = t
		}
	}

	if v, ok := d.GetOk("tcp_plen_err"); ok || d.HasChange("tcp_plen_err") {
		t, err := expandObjectSystemNpuFpAnomalyTcpPlenErr2edl(d, v, "tcp_plen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-plen-err"] = t
		}
	}

	if v, ok := d.GetOk("tcp_syn_data"); ok || d.HasChange("tcp_syn_data") {
		t, err := expandObjectSystemNpuFpAnomalyTcpSynData2edl(d, v, "tcp_syn_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-syn-data"] = t
		}
	}

	if v, ok := d.GetOk("tcp_syn_fin"); ok || d.HasChange("tcp_syn_fin") {
		t, err := expandObjectSystemNpuFpAnomalyTcpSynFin2edl(d, v, "tcp_syn_fin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-syn-fin"] = t
		}
	}

	if v, ok := d.GetOk("tcp_winnuke"); ok || d.HasChange("tcp_winnuke") {
		t, err := expandObjectSystemNpuFpAnomalyTcpWinnuke2edl(d, v, "tcp_winnuke")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-winnuke"] = t
		}
	}

	if v, ok := d.GetOk("udp_csum_err"); ok || d.HasChange("udp_csum_err") {
		t, err := expandObjectSystemNpuFpAnomalyUdpCsumErr2edl(d, v, "udp_csum_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-csum-err"] = t
		}
	}

	if v, ok := d.GetOk("udp_hlen_err"); ok || d.HasChange("udp_hlen_err") {
		t, err := expandObjectSystemNpuFpAnomalyUdpHlenErr2edl(d, v, "udp_hlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-hlen-err"] = t
		}
	}

	if v, ok := d.GetOk("udp_land"); ok || d.HasChange("udp_land") {
		t, err := expandObjectSystemNpuFpAnomalyUdpLand2edl(d, v, "udp_land")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-land"] = t
		}
	}

	if v, ok := d.GetOk("udp_len_err"); ok || d.HasChange("udp_len_err") {
		t, err := expandObjectSystemNpuFpAnomalyUdpLenErr2edl(d, v, "udp_len_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-len-err"] = t
		}
	}

	if v, ok := d.GetOk("udp_plen_err"); ok || d.HasChange("udp_plen_err") {
		t, err := expandObjectSystemNpuFpAnomalyUdpPlenErr2edl(d, v, "udp_plen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-plen-err"] = t
		}
	}

	if v, ok := d.GetOk("udplite_cover_err"); ok || d.HasChange("udplite_cover_err") {
		t, err := expandObjectSystemNpuFpAnomalyUdpliteCoverErr2edl(d, v, "udplite_cover_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udplite-cover-err"] = t
		}
	}

	if v, ok := d.GetOk("udplite_csum_err"); ok || d.HasChange("udplite_csum_err") {
		t, err := expandObjectSystemNpuFpAnomalyUdpliteCsumErr2edl(d, v, "udplite_csum_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udplite-csum-err"] = t
		}
	}

	if v, ok := d.GetOk("uesp_minlen_err"); ok || d.HasChange("uesp_minlen_err") {
		t, err := expandObjectSystemNpuFpAnomalyUespMinlenErr2edl(d, v, "uesp_minlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uesp-minlen-err"] = t
		}
	}

	if v, ok := d.GetOk("unknproto_minlen_err"); ok || d.HasChange("unknproto_minlen_err") {
		t, err := expandObjectSystemNpuFpAnomalyUnknprotoMinlenErr2edl(d, v, "unknproto_minlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknproto-minlen-err"] = t
		}
	}

	if v, ok := d.GetOk("vxlan_minlen_err"); ok || d.HasChange("vxlan_minlen_err") {
		t, err := expandObjectSystemNpuFpAnomalyVxlanMinlenErr2edl(d, v, "vxlan_minlen_err")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vxlan-minlen-err"] = t
		}
	}

	return &obj, nil
}
