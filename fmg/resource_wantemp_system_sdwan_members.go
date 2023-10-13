// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiGate interfaces added to the SD-WAN.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWantempSystemSdwanMembers() *schema.Resource {
	return &schema.Resource{
		Create: resourceWantempSystemSdwanMembersCreate,
		Read:   resourceWantempSystemSdwanMembersRead,
		Update: resourceWantempSystemSdwanMembersUpdate,
		Delete: resourceWantempSystemSdwanMembersDelete,

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
				Computed: true,
			},
			"gateway6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ingress_spillover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"preferred_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"priority6": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"seq_num": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Required: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spillover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"volume_ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWantempSystemSdwanMembersCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWantempSystemSdwanMembers(d)
	if err != nil {
		return fmt.Errorf("Error creating WantempSystemSdwanMembers resource while getting object: %v", err)
	}

	_, err = c.CreateWantempSystemSdwanMembers(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WantempSystemSdwanMembers resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "seq_num")))

	return resourceWantempSystemSdwanMembersRead(d, m)
}

func resourceWantempSystemSdwanMembersUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWantempSystemSdwanMembers(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwanMembers resource while getting object: %v", err)
	}

	_, err = c.UpdateWantempSystemSdwanMembers(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwanMembers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "seq_num")))

	return resourceWantempSystemSdwanMembersRead(d, m)
}

func resourceWantempSystemSdwanMembersDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWantempSystemSdwanMembers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WantempSystemSdwanMembers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWantempSystemSdwanMembersRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWantempSystemSdwanMembers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwanMembers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWantempSystemSdwanMembers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwanMembers resource from API: %v", err)
	}
	return nil
}

func flattenWantempSystemSdwanMembersDynamicMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersGateway6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersIngressSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersPreferredSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersPriority6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSource6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersVolumeRatio(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWantempSystemSdwanMembers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("_dynamic_member", flattenWantempSystemSdwanMembersDynamicMember(o["_dynamic-member"], d, "_dynamic_member")); err != nil {
		if vv, ok := fortiAPIPatch(o["_dynamic-member"], "WantempSystemSdwanMembers-DynamicMember"); ok {
			if err = d.Set("_dynamic_member", vv); err != nil {
				return fmt.Errorf("Error reading _dynamic_member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _dynamic_member: %v", err)
		}
	}

	if err = d.Set("comment", flattenWantempSystemSdwanMembersComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WantempSystemSdwanMembers-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("cost", flattenWantempSystemSdwanMembersCost(o["cost"], d, "cost")); err != nil {
		if vv, ok := fortiAPIPatch(o["cost"], "WantempSystemSdwanMembers-Cost"); ok {
			if err = d.Set("cost", vv); err != nil {
				return fmt.Errorf("Error reading cost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cost: %v", err)
		}
	}

	if err = d.Set("gateway", flattenWantempSystemSdwanMembersGateway(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "WantempSystemSdwanMembers-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("gateway6", flattenWantempSystemSdwanMembersGateway6(o["gateway6"], d, "gateway6")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway6"], "WantempSystemSdwanMembers-Gateway6"); ok {
			if err = d.Set("gateway6", vv); err != nil {
				return fmt.Errorf("Error reading gateway6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway6: %v", err)
		}
	}

	if err = d.Set("ingress_spillover_threshold", flattenWantempSystemSdwanMembersIngressSpilloverThreshold(o["ingress-spillover-threshold"], d, "ingress_spillover_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ingress-spillover-threshold"], "WantempSystemSdwanMembers-IngressSpilloverThreshold"); ok {
			if err = d.Set("ingress_spillover_threshold", vv); err != nil {
				return fmt.Errorf("Error reading ingress_spillover_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ingress_spillover_threshold: %v", err)
		}
	}

	if err = d.Set("interface", flattenWantempSystemSdwanMembersInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "WantempSystemSdwanMembers-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("preferred_source", flattenWantempSystemSdwanMembersPreferredSource(o["preferred-source"], d, "preferred_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["preferred-source"], "WantempSystemSdwanMembers-PreferredSource"); ok {
			if err = d.Set("preferred_source", vv); err != nil {
				return fmt.Errorf("Error reading preferred_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preferred_source: %v", err)
		}
	}

	if err = d.Set("priority", flattenWantempSystemSdwanMembersPriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "WantempSystemSdwanMembers-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("priority6", flattenWantempSystemSdwanMembersPriority6(o["priority6"], d, "priority6")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority6"], "WantempSystemSdwanMembers-Priority6"); ok {
			if err = d.Set("priority6", vv); err != nil {
				return fmt.Errorf("Error reading priority6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority6: %v", err)
		}
	}

	if err = d.Set("seq_num", flattenWantempSystemSdwanMembersSeqNum(o["seq-num"], d, "seq_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["seq-num"], "WantempSystemSdwanMembers-SeqNum"); ok {
			if err = d.Set("seq_num", vv); err != nil {
				return fmt.Errorf("Error reading seq_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	if err = d.Set("source", flattenWantempSystemSdwanMembersSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "WantempSystemSdwanMembers-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("source6", flattenWantempSystemSdwanMembersSource6(o["source6"], d, "source6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source6"], "WantempSystemSdwanMembers-Source6"); ok {
			if err = d.Set("source6", vv); err != nil {
				return fmt.Errorf("Error reading source6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source6: %v", err)
		}
	}

	if err = d.Set("spillover_threshold", flattenWantempSystemSdwanMembersSpilloverThreshold(o["spillover-threshold"], d, "spillover_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["spillover-threshold"], "WantempSystemSdwanMembers-SpilloverThreshold"); ok {
			if err = d.Set("spillover_threshold", vv); err != nil {
				return fmt.Errorf("Error reading spillover_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spillover_threshold: %v", err)
		}
	}

	if err = d.Set("status", flattenWantempSystemSdwanMembersStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WantempSystemSdwanMembers-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("volume_ratio", flattenWantempSystemSdwanMembersVolumeRatio(o["volume-ratio"], d, "volume_ratio")); err != nil {
		if vv, ok := fortiAPIPatch(o["volume-ratio"], "WantempSystemSdwanMembers-VolumeRatio"); ok {
			if err = d.Set("volume_ratio", vv); err != nil {
				return fmt.Errorf("Error reading volume_ratio: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading volume_ratio: %v", err)
		}
	}

	if err = d.Set("weight", flattenWantempSystemSdwanMembersWeight(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "WantempSystemSdwanMembers-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("zone", flattenWantempSystemSdwanMembersZone(o["zone"], d, "zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["zone"], "WantempSystemSdwanMembers-Zone"); ok {
			if err = d.Set("zone", vv); err != nil {
				return fmt.Errorf("Error reading zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading zone: %v", err)
		}
	}

	return nil
}

func flattenWantempSystemSdwanMembersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWantempSystemSdwanMembersDynamicMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersGateway6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersIngressSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersPreferredSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersPriority6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSeqNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSource6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersVolumeRatio(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWantempSystemSdwanMembers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_dynamic_member"); ok || d.HasChange("_dynamic_member") {
		t, err := expandWantempSystemSdwanMembersDynamicMember(d, v, "_dynamic_member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_dynamic-member"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWantempSystemSdwanMembersComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("cost"); ok || d.HasChange("cost") {
		t, err := expandWantempSystemSdwanMembersCost(d, v, "cost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandWantempSystemSdwanMembersGateway(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("gateway6"); ok || d.HasChange("gateway6") {
		t, err := expandWantempSystemSdwanMembersGateway6(d, v, "gateway6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway6"] = t
		}
	}

	if v, ok := d.GetOk("ingress_spillover_threshold"); ok || d.HasChange("ingress_spillover_threshold") {
		t, err := expandWantempSystemSdwanMembersIngressSpilloverThreshold(d, v, "ingress_spillover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress-spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandWantempSystemSdwanMembersInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("preferred_source"); ok || d.HasChange("preferred_source") {
		t, err := expandWantempSystemSdwanMembersPreferredSource(d, v, "preferred_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-source"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandWantempSystemSdwanMembersPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("priority6"); ok || d.HasChange("priority6") {
		t, err := expandWantempSystemSdwanMembersPriority6(d, v, "priority6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority6"] = t
		}
	}

	if v, ok := d.GetOk("seq_num"); ok || d.HasChange("seq_num") {
		t, err := expandWantempSystemSdwanMembersSeqNum(d, v, "seq_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq-num"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandWantempSystemSdwanMembersSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("source6"); ok || d.HasChange("source6") {
		t, err := expandWantempSystemSdwanMembersSource6(d, v, "source6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source6"] = t
		}
	}

	if v, ok := d.GetOk("spillover_threshold"); ok || d.HasChange("spillover_threshold") {
		t, err := expandWantempSystemSdwanMembersSpilloverThreshold(d, v, "spillover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWantempSystemSdwanMembersStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("volume_ratio"); ok || d.HasChange("volume_ratio") {
		t, err := expandWantempSystemSdwanMembersVolumeRatio(d, v, "volume_ratio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["volume-ratio"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandWantempSystemSdwanMembersWeight(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	if v, ok := d.GetOk("zone"); ok || d.HasChange("zone") {
		t, err := expandWantempSystemSdwanMembersZone(d, v, "zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["zone"] = t
		}
	}

	return &obj, nil
}
