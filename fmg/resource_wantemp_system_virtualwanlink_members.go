// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiGate interfaces added to the virtual-wan-link.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWantempSystemVirtualWanLinkMembers() *schema.Resource {
	return &schema.Resource{
		Create: resourceWantempSystemVirtualWanLinkMembersCreate,
		Read:   resourceWantempSystemVirtualWanLinkMembersRead,
		Update: resourceWantempSystemVirtualWanLinkMembersUpdate,
		Delete: resourceWantempSystemVirtualWanLinkMembersDelete,

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
			"_dynamic_member": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cost": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ingress_spillover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"seq_num": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"spillover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceWantempSystemVirtualWanLinkMembersCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["wanprof"] = wanprof

	obj, err := getObjectWantempSystemVirtualWanLinkMembers(d)
	if err != nil {
		return fmt.Errorf("Error creating WantempSystemVirtualWanLinkMembers resource while getting object: %v", err)
	}

	_, err = c.CreateWantempSystemVirtualWanLinkMembers(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WantempSystemVirtualWanLinkMembers resource: %v", err)
	}

	d.SetId(getStringKey(d, ""))

	return resourceWantempSystemVirtualWanLinkMembersRead(d, m)
}

func resourceWantempSystemVirtualWanLinkMembersUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["wanprof"] = wanprof

	obj, err := getObjectWantempSystemVirtualWanLinkMembers(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemVirtualWanLinkMembers resource while getting object: %v", err)
	}

	_, err = c.UpdateWantempSystemVirtualWanLinkMembers(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemVirtualWanLinkMembers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, ""))

	return resourceWantempSystemVirtualWanLinkMembersRead(d, m)
}

func resourceWantempSystemVirtualWanLinkMembersDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["wanprof"] = wanprof

	err = c.DeleteWantempSystemVirtualWanLinkMembers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WantempSystemVirtualWanLinkMembers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWantempSystemVirtualWanLinkMembersRead(d *schema.ResourceData, m interface{}) error {
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
		if err = d.Set("wanprof", wanprof); err != nil {
			return fmt.Errorf("Error set params wanprof: %v", err)
		}
	}
	paradict["wanprof"] = wanprof

	o, err := c.ReadWantempSystemVirtualWanLinkMembers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemVirtualWanLinkMembers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWantempSystemVirtualWanLinkMembers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemVirtualWanLinkMembers resource from API: %v", err)
	}
	return nil
}

func flattenWantempSystemVirtualWanLinkMembersDynamicMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersGateway6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersIngressSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersSeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersSource6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersVolumeRatio(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWantempSystemVirtualWanLinkMembers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("_dynamic_member", flattenWantempSystemVirtualWanLinkMembersDynamicMember(o["_dynamic-member"], d, "_dynamic_member")); err != nil {
		if vv, ok := fortiAPIPatch(o["_dynamic-member"], "WantempSystemVirtualWanLinkMembers-DynamicMember"); ok {
			if err = d.Set("_dynamic_member", vv); err != nil {
				return fmt.Errorf("Error reading _dynamic_member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _dynamic_member: %v", err)
		}
	}

	if err = d.Set("comment", flattenWantempSystemVirtualWanLinkMembersComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WantempSystemVirtualWanLinkMembers-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("cost", flattenWantempSystemVirtualWanLinkMembersCost(o["cost"], d, "cost")); err != nil {
		if vv, ok := fortiAPIPatch(o["cost"], "WantempSystemVirtualWanLinkMembers-Cost"); ok {
			if err = d.Set("cost", vv); err != nil {
				return fmt.Errorf("Error reading cost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cost: %v", err)
		}
	}

	if err = d.Set("gateway", flattenWantempSystemVirtualWanLinkMembersGateway(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "WantempSystemVirtualWanLinkMembers-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("gateway6", flattenWantempSystemVirtualWanLinkMembersGateway6(o["gateway6"], d, "gateway6")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway6"], "WantempSystemVirtualWanLinkMembers-Gateway6"); ok {
			if err = d.Set("gateway6", vv); err != nil {
				return fmt.Errorf("Error reading gateway6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway6: %v", err)
		}
	}

	if err = d.Set("ingress_spillover_threshold", flattenWantempSystemVirtualWanLinkMembersIngressSpilloverThreshold(o["ingress-spillover-threshold"], d, "ingress_spillover_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ingress-spillover-threshold"], "WantempSystemVirtualWanLinkMembers-IngressSpilloverThreshold"); ok {
			if err = d.Set("ingress_spillover_threshold", vv); err != nil {
				return fmt.Errorf("Error reading ingress_spillover_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ingress_spillover_threshold: %v", err)
		}
	}

	if err = d.Set("interface", flattenWantempSystemVirtualWanLinkMembersInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "WantempSystemVirtualWanLinkMembers-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("priority", flattenWantempSystemVirtualWanLinkMembersPriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "WantempSystemVirtualWanLinkMembers-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("seq_num", flattenWantempSystemVirtualWanLinkMembersSeqNum(o["seq-num"], d, "seq_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["seq-num"], "WantempSystemVirtualWanLinkMembers-SeqNum"); ok {
			if err = d.Set("seq_num", vv); err != nil {
				return fmt.Errorf("Error reading seq_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	if err = d.Set("source", flattenWantempSystemVirtualWanLinkMembersSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "WantempSystemVirtualWanLinkMembers-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("source6", flattenWantempSystemVirtualWanLinkMembersSource6(o["source6"], d, "source6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source6"], "WantempSystemVirtualWanLinkMembers-Source6"); ok {
			if err = d.Set("source6", vv); err != nil {
				return fmt.Errorf("Error reading source6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source6: %v", err)
		}
	}

	if err = d.Set("spillover_threshold", flattenWantempSystemVirtualWanLinkMembersSpilloverThreshold(o["spillover-threshold"], d, "spillover_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["spillover-threshold"], "WantempSystemVirtualWanLinkMembers-SpilloverThreshold"); ok {
			if err = d.Set("spillover_threshold", vv); err != nil {
				return fmt.Errorf("Error reading spillover_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spillover_threshold: %v", err)
		}
	}

	if err = d.Set("status", flattenWantempSystemVirtualWanLinkMembersStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WantempSystemVirtualWanLinkMembers-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("volume_ratio", flattenWantempSystemVirtualWanLinkMembersVolumeRatio(o["volume-ratio"], d, "volume_ratio")); err != nil {
		if vv, ok := fortiAPIPatch(o["volume-ratio"], "WantempSystemVirtualWanLinkMembers-VolumeRatio"); ok {
			if err = d.Set("volume_ratio", vv); err != nil {
				return fmt.Errorf("Error reading volume_ratio: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading volume_ratio: %v", err)
		}
	}

	if err = d.Set("weight", flattenWantempSystemVirtualWanLinkMembersWeight(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "WantempSystemVirtualWanLinkMembers-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenWantempSystemVirtualWanLinkMembersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWantempSystemVirtualWanLinkMembersDynamicMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersGateway6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersIngressSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersSeqNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersSource6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersVolumeRatio(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWantempSystemVirtualWanLinkMembers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_dynamic_member"); ok || d.HasChange("_dynamic_member") {
		t, err := expandWantempSystemVirtualWanLinkMembersDynamicMember(d, v, "_dynamic_member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_dynamic-member"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWantempSystemVirtualWanLinkMembersComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("cost"); ok || d.HasChange("cost") {
		t, err := expandWantempSystemVirtualWanLinkMembersCost(d, v, "cost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandWantempSystemVirtualWanLinkMembersGateway(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("gateway6"); ok || d.HasChange("gateway6") {
		t, err := expandWantempSystemVirtualWanLinkMembersGateway6(d, v, "gateway6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway6"] = t
		}
	}

	if v, ok := d.GetOk("ingress_spillover_threshold"); ok || d.HasChange("ingress_spillover_threshold") {
		t, err := expandWantempSystemVirtualWanLinkMembersIngressSpilloverThreshold(d, v, "ingress_spillover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress-spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandWantempSystemVirtualWanLinkMembersInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandWantempSystemVirtualWanLinkMembersPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("seq_num"); ok || d.HasChange("seq_num") {
		t, err := expandWantempSystemVirtualWanLinkMembersSeqNum(d, v, "seq_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq-num"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandWantempSystemVirtualWanLinkMembersSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("source6"); ok || d.HasChange("source6") {
		t, err := expandWantempSystemVirtualWanLinkMembersSource6(d, v, "source6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source6"] = t
		}
	}

	if v, ok := d.GetOk("spillover_threshold"); ok || d.HasChange("spillover_threshold") {
		t, err := expandWantempSystemVirtualWanLinkMembersSpilloverThreshold(d, v, "spillover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWantempSystemVirtualWanLinkMembersStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("volume_ratio"); ok || d.HasChange("volume_ratio") {
		t, err := expandWantempSystemVirtualWanLinkMembersVolumeRatio(d, v, "volume_ratio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["volume-ratio"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandWantempSystemVirtualWanLinkMembersWeight(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
