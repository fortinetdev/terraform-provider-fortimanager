// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Create SD-WAN duplication rule.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWantempSystemSdwanDuplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceWantempSystemSdwanDuplicationCreate,
		Read:   resourceWantempSystemSdwanDuplicationRead,
		Update: resourceWantempSystemSdwanDuplicationUpdate,
		Delete: resourceWantempSystemSdwanDuplicationDelete,

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
			"wanprof": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"packet_de_duplication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"packet_duplication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sla_match_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWantempSystemSdwanDuplicationCreate(d *schema.ResourceData, m interface{}) error {
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

	wanprof := d.Get("wanprof").(string)
	paradict["wanprof"] = wanprof

	obj, err := getObjectWantempSystemSdwanDuplication(d)
	if err != nil {
		return fmt.Errorf("Error creating WantempSystemSdwanDuplication resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateWantempSystemSdwanDuplication(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WantempSystemSdwanDuplication resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWantempSystemSdwanDuplicationRead(d, m)
}

func resourceWantempSystemSdwanDuplicationUpdate(d *schema.ResourceData, m interface{}) error {
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

	wanprof := d.Get("wanprof").(string)
	paradict["wanprof"] = wanprof

	obj, err := getObjectWantempSystemSdwanDuplication(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwanDuplication resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWantempSystemSdwanDuplication(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwanDuplication resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWantempSystemSdwanDuplicationRead(d, m)
}

func resourceWantempSystemSdwanDuplicationDelete(d *schema.ResourceData, m interface{}) error {
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

	wanprof := d.Get("wanprof").(string)
	paradict["wanprof"] = wanprof

	wsParams["adom"] = adomv

	err = c.DeleteWantempSystemSdwanDuplication(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WantempSystemSdwanDuplication resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWantempSystemSdwanDuplicationRead(d *schema.ResourceData, m interface{}) error {
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

	wanprof := d.Get("wanprof").(string)
	if wanprof == "" {
		wanprof = importOptionChecking(m.(*FortiClient).Cfg, "wanprof")
		if wanprof == "" {
			return fmt.Errorf("Parameter wanprof is missing")
		}
		if err = d.Set("wanprof", wanprof); err != nil {
			return fmt.Errorf("Error set params wanprof: %v", err)
		}
	}
	paradict["wanprof"] = wanprof

	o, err := c.ReadWantempSystemSdwanDuplication(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwanDuplication resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWantempSystemSdwanDuplication(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwanDuplication resource from API: %v", err)
	}
	return nil
}

func flattenWantempSystemSdwanDuplicationDstaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanDuplicationDstaddr62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanDuplicationDstintf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanDuplicationId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationPacketDeDuplication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationPacketDuplication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanDuplicationServiceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanDuplicationSlaMatchService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationSrcaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanDuplicationSrcaddr62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanDuplicationSrcintf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectWantempSystemSdwanDuplication(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("dstaddr", flattenWantempSystemSdwanDuplicationDstaddr2edl(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "WantempSystemSdwanDuplication-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenWantempSystemSdwanDuplicationDstaddr62edl(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "WantempSystemSdwanDuplication-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenWantempSystemSdwanDuplicationDstintf2edl(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "WantempSystemSdwanDuplication-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWantempSystemSdwanDuplicationId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WantempSystemSdwanDuplication-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("packet_de_duplication", flattenWantempSystemSdwanDuplicationPacketDeDuplication2edl(o["packet-de-duplication"], d, "packet_de_duplication")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-de-duplication"], "WantempSystemSdwanDuplication-PacketDeDuplication"); ok {
			if err = d.Set("packet_de_duplication", vv); err != nil {
				return fmt.Errorf("Error reading packet_de_duplication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_de_duplication: %v", err)
		}
	}

	if err = d.Set("packet_duplication", flattenWantempSystemSdwanDuplicationPacketDuplication2edl(o["packet-duplication"], d, "packet_duplication")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-duplication"], "WantempSystemSdwanDuplication-PacketDuplication"); ok {
			if err = d.Set("packet_duplication", vv); err != nil {
				return fmt.Errorf("Error reading packet_duplication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_duplication: %v", err)
		}
	}

	if err = d.Set("service", flattenWantempSystemSdwanDuplicationService2edl(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "WantempSystemSdwanDuplication-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("service_id", flattenWantempSystemSdwanDuplicationServiceId2edl(o["service-id"], d, "service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-id"], "WantempSystemSdwanDuplication-ServiceId"); ok {
			if err = d.Set("service_id", vv); err != nil {
				return fmt.Errorf("Error reading service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_id: %v", err)
		}
	}

	if err = d.Set("sla_match_service", flattenWantempSystemSdwanDuplicationSlaMatchService2edl(o["sla-match-service"], d, "sla_match_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-match-service"], "WantempSystemSdwanDuplication-SlaMatchService"); ok {
			if err = d.Set("sla_match_service", vv); err != nil {
				return fmt.Errorf("Error reading sla_match_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_match_service: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenWantempSystemSdwanDuplicationSrcaddr2edl(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "WantempSystemSdwanDuplication-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenWantempSystemSdwanDuplicationSrcaddr62edl(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "WantempSystemSdwanDuplication-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenWantempSystemSdwanDuplicationSrcintf2edl(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "WantempSystemSdwanDuplication-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	return nil
}

func flattenWantempSystemSdwanDuplicationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWantempSystemSdwanDuplicationDstaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanDuplicationDstaddr62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanDuplicationDstintf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanDuplicationId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationPacketDeDuplication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationPacketDuplication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanDuplicationServiceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanDuplicationSlaMatchService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationSrcaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanDuplicationSrcaddr62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanDuplicationSrcintf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectWantempSystemSdwanDuplication(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandWantempSystemSdwanDuplicationDstaddr2edl(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandWantempSystemSdwanDuplicationDstaddr62edl(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandWantempSystemSdwanDuplicationDstintf2edl(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWantempSystemSdwanDuplicationId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("packet_de_duplication"); ok || d.HasChange("packet_de_duplication") {
		t, err := expandWantempSystemSdwanDuplicationPacketDeDuplication2edl(d, v, "packet_de_duplication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-de-duplication"] = t
		}
	}

	if v, ok := d.GetOk("packet_duplication"); ok || d.HasChange("packet_duplication") {
		t, err := expandWantempSystemSdwanDuplicationPacketDuplication2edl(d, v, "packet_duplication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-duplication"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandWantempSystemSdwanDuplicationService2edl(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("service_id"); ok || d.HasChange("service_id") {
		t, err := expandWantempSystemSdwanDuplicationServiceId2edl(d, v, "service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-id"] = t
		}
	}

	if v, ok := d.GetOk("sla_match_service"); ok || d.HasChange("sla_match_service") {
		t, err := expandWantempSystemSdwanDuplicationSlaMatchService2edl(d, v, "sla_match_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-match-service"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandWantempSystemSdwanDuplicationSrcaddr2edl(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandWantempSystemSdwanDuplicationSrcaddr62edl(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandWantempSystemSdwanDuplicationSrcintf2edl(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	return &obj, nil
}
