// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Host protection engine configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuHpe() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuHpeUpdate,
		Read:   resourceObjectSystemNpuHpeRead,
		Update: resourceObjectSystemNpuHpeUpdate,
		Delete: resourceObjectSystemNpuHpeDelete,

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
			"all_protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"arp_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"enable_queue_shaper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_shaper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"esp_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"exception_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fragment_with_sess": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fragment_without_session": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"high_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"icmp_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip_frag_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip_others_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"l2_others_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pri_type_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"queue_shaper_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sctp_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcpfin_rst_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcpsyn_ack_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcpsyn_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"udp_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemNpuHpeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuHpe(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuHpe resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuHpe(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuHpe resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuHpe")

	return resourceObjectSystemNpuHpeRead(d, m)
}

func resourceObjectSystemNpuHpeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemNpuHpe(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuHpe resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuHpeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuHpe(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuHpe resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuHpe(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuHpe resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuHpeAllProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeArpMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeEnableQueueShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeEnableShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeEspMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeExceptionCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeFragmentWithSess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeFragmentWithoutSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeHighPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeIcmpMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeIpFragMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeIpOthersMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeL2OthersMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpePriTypeMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeQueueShaperMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeSctpMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeTcpMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeTcpfinRstMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeTcpsynAckMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeTcpsynMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeUdpMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuHpe(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("all_protocol", flattenObjectSystemNpuHpeAllProtocol(o["all-protocol"], d, "all_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["all-protocol"], "ObjectSystemNpuHpe-AllProtocol"); ok {
			if err = d.Set("all_protocol", vv); err != nil {
				return fmt.Errorf("Error reading all_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading all_protocol: %v", err)
		}
	}

	if err = d.Set("arp_max", flattenObjectSystemNpuHpeArpMax(o["arp-max"], d, "arp_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-max"], "ObjectSystemNpuHpe-ArpMax"); ok {
			if err = d.Set("arp_max", vv); err != nil {
				return fmt.Errorf("Error reading arp_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_max: %v", err)
		}
	}

	if err = d.Set("enable_queue_shaper", flattenObjectSystemNpuHpeEnableQueueShaper(o["enable-queue-shaper"], d, "enable_queue_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["enable-queue-shaper"], "ObjectSystemNpuHpe-EnableQueueShaper"); ok {
			if err = d.Set("enable_queue_shaper", vv); err != nil {
				return fmt.Errorf("Error reading enable_queue_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enable_queue_shaper: %v", err)
		}
	}

	if err = d.Set("enable_shaper", flattenObjectSystemNpuHpeEnableShaper(o["enable-shaper"], d, "enable_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["enable-shaper"], "ObjectSystemNpuHpe-EnableShaper"); ok {
			if err = d.Set("enable_shaper", vv); err != nil {
				return fmt.Errorf("Error reading enable_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enable_shaper: %v", err)
		}
	}

	if err = d.Set("esp_max", flattenObjectSystemNpuHpeEspMax(o["esp-max"], d, "esp_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["esp-max"], "ObjectSystemNpuHpe-EspMax"); ok {
			if err = d.Set("esp_max", vv); err != nil {
				return fmt.Errorf("Error reading esp_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading esp_max: %v", err)
		}
	}

	if err = d.Set("exception_code", flattenObjectSystemNpuHpeExceptionCode(o["exception-code"], d, "exception_code")); err != nil {
		if vv, ok := fortiAPIPatch(o["exception-code"], "ObjectSystemNpuHpe-ExceptionCode"); ok {
			if err = d.Set("exception_code", vv); err != nil {
				return fmt.Errorf("Error reading exception_code: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exception_code: %v", err)
		}
	}

	if err = d.Set("fragment_with_sess", flattenObjectSystemNpuHpeFragmentWithSess(o["fragment-with-sess"], d, "fragment_with_sess")); err != nil {
		if vv, ok := fortiAPIPatch(o["fragment-with-sess"], "ObjectSystemNpuHpe-FragmentWithSess"); ok {
			if err = d.Set("fragment_with_sess", vv); err != nil {
				return fmt.Errorf("Error reading fragment_with_sess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fragment_with_sess: %v", err)
		}
	}

	if err = d.Set("fragment_without_session", flattenObjectSystemNpuHpeFragmentWithoutSession(o["fragment-without-session"], d, "fragment_without_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["fragment-without-session"], "ObjectSystemNpuHpe-FragmentWithoutSession"); ok {
			if err = d.Set("fragment_without_session", vv); err != nil {
				return fmt.Errorf("Error reading fragment_without_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fragment_without_session: %v", err)
		}
	}

	if err = d.Set("high_priority", flattenObjectSystemNpuHpeHighPriority(o["high-priority"], d, "high_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["high-priority"], "ObjectSystemNpuHpe-HighPriority"); ok {
			if err = d.Set("high_priority", vv); err != nil {
				return fmt.Errorf("Error reading high_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading high_priority: %v", err)
		}
	}

	if err = d.Set("icmp_max", flattenObjectSystemNpuHpeIcmpMax(o["icmp-max"], d, "icmp_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-max"], "ObjectSystemNpuHpe-IcmpMax"); ok {
			if err = d.Set("icmp_max", vv); err != nil {
				return fmt.Errorf("Error reading icmp_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_max: %v", err)
		}
	}

	if err = d.Set("ip_frag_max", flattenObjectSystemNpuHpeIpFragMax(o["ip-frag-max"], d, "ip_frag_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-frag-max"], "ObjectSystemNpuHpe-IpFragMax"); ok {
			if err = d.Set("ip_frag_max", vv); err != nil {
				return fmt.Errorf("Error reading ip_frag_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_frag_max: %v", err)
		}
	}

	if err = d.Set("ip_others_max", flattenObjectSystemNpuHpeIpOthersMax(o["ip-others-max"], d, "ip_others_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-others-max"], "ObjectSystemNpuHpe-IpOthersMax"); ok {
			if err = d.Set("ip_others_max", vv); err != nil {
				return fmt.Errorf("Error reading ip_others_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_others_max: %v", err)
		}
	}

	if err = d.Set("l2_others_max", flattenObjectSystemNpuHpeL2OthersMax(o["l2-others-max"], d, "l2_others_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["l2-others-max"], "ObjectSystemNpuHpe-L2OthersMax"); ok {
			if err = d.Set("l2_others_max", vv); err != nil {
				return fmt.Errorf("Error reading l2_others_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l2_others_max: %v", err)
		}
	}

	if err = d.Set("pri_type_max", flattenObjectSystemNpuHpePriTypeMax(o["pri-type-max"], d, "pri_type_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["pri-type-max"], "ObjectSystemNpuHpe-PriTypeMax"); ok {
			if err = d.Set("pri_type_max", vv); err != nil {
				return fmt.Errorf("Error reading pri_type_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pri_type_max: %v", err)
		}
	}

	if err = d.Set("queue_shaper_max", flattenObjectSystemNpuHpeQueueShaperMax(o["queue-shaper-max"], d, "queue_shaper_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["queue-shaper-max"], "ObjectSystemNpuHpe-QueueShaperMax"); ok {
			if err = d.Set("queue_shaper_max", vv); err != nil {
				return fmt.Errorf("Error reading queue_shaper_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading queue_shaper_max: %v", err)
		}
	}

	if err = d.Set("sctp_max", flattenObjectSystemNpuHpeSctpMax(o["sctp-max"], d, "sctp_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["sctp-max"], "ObjectSystemNpuHpe-SctpMax"); ok {
			if err = d.Set("sctp_max", vv); err != nil {
				return fmt.Errorf("Error reading sctp_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sctp_max: %v", err)
		}
	}

	if err = d.Set("tcp_max", flattenObjectSystemNpuHpeTcpMax(o["tcp-max"], d, "tcp_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-max"], "ObjectSystemNpuHpe-TcpMax"); ok {
			if err = d.Set("tcp_max", vv); err != nil {
				return fmt.Errorf("Error reading tcp_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_max: %v", err)
		}
	}

	if err = d.Set("tcpfin_rst_max", flattenObjectSystemNpuHpeTcpfinRstMax(o["tcpfin-rst-max"], d, "tcpfin_rst_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcpfin-rst-max"], "ObjectSystemNpuHpe-TcpfinRstMax"); ok {
			if err = d.Set("tcpfin_rst_max", vv); err != nil {
				return fmt.Errorf("Error reading tcpfin_rst_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcpfin_rst_max: %v", err)
		}
	}

	if err = d.Set("tcpsyn_ack_max", flattenObjectSystemNpuHpeTcpsynAckMax(o["tcpsyn-ack-max"], d, "tcpsyn_ack_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcpsyn-ack-max"], "ObjectSystemNpuHpe-TcpsynAckMax"); ok {
			if err = d.Set("tcpsyn_ack_max", vv); err != nil {
				return fmt.Errorf("Error reading tcpsyn_ack_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcpsyn_ack_max: %v", err)
		}
	}

	if err = d.Set("tcpsyn_max", flattenObjectSystemNpuHpeTcpsynMax(o["tcpsyn-max"], d, "tcpsyn_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcpsyn-max"], "ObjectSystemNpuHpe-TcpsynMax"); ok {
			if err = d.Set("tcpsyn_max", vv); err != nil {
				return fmt.Errorf("Error reading tcpsyn_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcpsyn_max: %v", err)
		}
	}

	if err = d.Set("udp_max", flattenObjectSystemNpuHpeUdpMax(o["udp-max"], d, "udp_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-max"], "ObjectSystemNpuHpe-UdpMax"); ok {
			if err = d.Set("udp_max", vv); err != nil {
				return fmt.Errorf("Error reading udp_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_max: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuHpeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuHpeAllProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeArpMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeEnableQueueShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeEnableShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeEspMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeExceptionCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeFragmentWithSess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeFragmentWithoutSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeHighPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeIcmpMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeIpFragMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeIpOthersMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeL2OthersMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpePriTypeMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeQueueShaperMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeSctpMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeTcpMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeTcpfinRstMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeTcpsynAckMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeTcpsynMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeUdpMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuHpe(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("all_protocol"); ok || d.HasChange("all_protocol") {
		t, err := expandObjectSystemNpuHpeAllProtocol(d, v, "all_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["all-protocol"] = t
		}
	}

	if v, ok := d.GetOk("arp_max"); ok || d.HasChange("arp_max") {
		t, err := expandObjectSystemNpuHpeArpMax(d, v, "arp_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-max"] = t
		}
	}

	if v, ok := d.GetOk("enable_queue_shaper"); ok || d.HasChange("enable_queue_shaper") {
		t, err := expandObjectSystemNpuHpeEnableQueueShaper(d, v, "enable_queue_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enable-queue-shaper"] = t
		}
	}

	if v, ok := d.GetOk("enable_shaper"); ok || d.HasChange("enable_shaper") {
		t, err := expandObjectSystemNpuHpeEnableShaper(d, v, "enable_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enable-shaper"] = t
		}
	}

	if v, ok := d.GetOk("esp_max"); ok || d.HasChange("esp_max") {
		t, err := expandObjectSystemNpuHpeEspMax(d, v, "esp_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["esp-max"] = t
		}
	}

	if v, ok := d.GetOk("exception_code"); ok || d.HasChange("exception_code") {
		t, err := expandObjectSystemNpuHpeExceptionCode(d, v, "exception_code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exception-code"] = t
		}
	}

	if v, ok := d.GetOk("fragment_with_sess"); ok || d.HasChange("fragment_with_sess") {
		t, err := expandObjectSystemNpuHpeFragmentWithSess(d, v, "fragment_with_sess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fragment-with-sess"] = t
		}
	}

	if v, ok := d.GetOk("fragment_without_session"); ok || d.HasChange("fragment_without_session") {
		t, err := expandObjectSystemNpuHpeFragmentWithoutSession(d, v, "fragment_without_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fragment-without-session"] = t
		}
	}

	if v, ok := d.GetOk("high_priority"); ok || d.HasChange("high_priority") {
		t, err := expandObjectSystemNpuHpeHighPriority(d, v, "high_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["high-priority"] = t
		}
	}

	if v, ok := d.GetOk("icmp_max"); ok || d.HasChange("icmp_max") {
		t, err := expandObjectSystemNpuHpeIcmpMax(d, v, "icmp_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-max"] = t
		}
	}

	if v, ok := d.GetOk("ip_frag_max"); ok || d.HasChange("ip_frag_max") {
		t, err := expandObjectSystemNpuHpeIpFragMax(d, v, "ip_frag_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-frag-max"] = t
		}
	}

	if v, ok := d.GetOk("ip_others_max"); ok || d.HasChange("ip_others_max") {
		t, err := expandObjectSystemNpuHpeIpOthersMax(d, v, "ip_others_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-others-max"] = t
		}
	}

	if v, ok := d.GetOk("l2_others_max"); ok || d.HasChange("l2_others_max") {
		t, err := expandObjectSystemNpuHpeL2OthersMax(d, v, "l2_others_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2-others-max"] = t
		}
	}

	if v, ok := d.GetOk("pri_type_max"); ok || d.HasChange("pri_type_max") {
		t, err := expandObjectSystemNpuHpePriTypeMax(d, v, "pri_type_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pri-type-max"] = t
		}
	}

	if v, ok := d.GetOk("queue_shaper_max"); ok || d.HasChange("queue_shaper_max") {
		t, err := expandObjectSystemNpuHpeQueueShaperMax(d, v, "queue_shaper_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["queue-shaper-max"] = t
		}
	}

	if v, ok := d.GetOk("sctp_max"); ok || d.HasChange("sctp_max") {
		t, err := expandObjectSystemNpuHpeSctpMax(d, v, "sctp_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-max"] = t
		}
	}

	if v, ok := d.GetOk("tcp_max"); ok || d.HasChange("tcp_max") {
		t, err := expandObjectSystemNpuHpeTcpMax(d, v, "tcp_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-max"] = t
		}
	}

	if v, ok := d.GetOk("tcpfin_rst_max"); ok || d.HasChange("tcpfin_rst_max") {
		t, err := expandObjectSystemNpuHpeTcpfinRstMax(d, v, "tcpfin_rst_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcpfin-rst-max"] = t
		}
	}

	if v, ok := d.GetOk("tcpsyn_ack_max"); ok || d.HasChange("tcpsyn_ack_max") {
		t, err := expandObjectSystemNpuHpeTcpsynAckMax(d, v, "tcpsyn_ack_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcpsyn-ack-max"] = t
		}
	}

	if v, ok := d.GetOk("tcpsyn_max"); ok || d.HasChange("tcpsyn_max") {
		t, err := expandObjectSystemNpuHpeTcpsynMax(d, v, "tcpsyn_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcpsyn-max"] = t
		}
	}

	if v, ok := d.GetOk("udp_max"); ok || d.HasChange("udp_max") {
		t, err := expandObjectSystemNpuHpeUdpMax(d, v, "udp_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-max"] = t
		}
	}

	return &obj, nil
}
