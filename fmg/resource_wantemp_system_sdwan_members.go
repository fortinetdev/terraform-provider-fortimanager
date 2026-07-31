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
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"billing_start_day": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"duplication_threshold_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"duplication_threshold_bibandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"duplication_threshold_dwbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"duplication_threshold_upbandwidth": &schema.Schema{
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
			"overage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"overage_cost": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"overage_volume_ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"overage_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
			"priority_in_sla": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority_out_sla": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority6": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"quota_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			"transport_group": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
	wsParams := make(map[string]string)
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
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("seq_num")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWantempSystemSdwanMembers(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWantempSystemSdwanMembers(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WantempSystemSdwanMembers resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWantempSystemSdwanMembers(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WantempSystemSdwanMembers resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "seq_num")))

	return resourceWantempSystemSdwanMembersRead(d, m)
}

func resourceWantempSystemSdwanMembersUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWantempSystemSdwanMembers(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwanMembers resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWantempSystemSdwanMembers(obj, mkey, paradict, wsParams)
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

	err = c.DeleteWantempSystemSdwanMembers(mkey, paradict, wsParams)
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
		d.SetId("")
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

func flattenWantempSystemSdwanMembersBillingStartDay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersDynamicMember2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersCost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenWantempSystemSdwanMembersDuplicationThresholdBandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersDuplicationThresholdBibandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersDuplicationThresholdDwbandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersDuplicationThresholdUpbandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersGateway2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersGateway62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersIngressSpilloverThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanMembersOverage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersOverageCost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersOverageVolumeRatio2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersOverageWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersPreferredSource2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersPriorityInSla2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersPriorityOutSla2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersPriority62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersQuotaLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSeqNum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSource2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSource62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSpilloverThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersTransportGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersVolumeRatio2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersZone2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectWantempSystemSdwanMembers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("billing_start_day", flattenWantempSystemSdwanMembersBillingStartDay2edl(o["billing-start-day"], d, "billing_start_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["billing-start-day"], "WantempSystemSdwanMembers-BillingStartDay"); ok {
			if err = d.Set("billing_start_day", vv); err != nil {
				return fmt.Errorf("Error reading billing_start_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading billing_start_day: %v", err)
		}
	}

	if err = d.Set("_dynamic_member", flattenWantempSystemSdwanMembersDynamicMember2edl(o["_dynamic-member"], d, "_dynamic_member")); err != nil {
		if vv, ok := fortiAPIPatch(o["_dynamic-member"], "WantempSystemSdwanMembers-DynamicMember"); ok {
			if err = d.Set("_dynamic_member", vv); err != nil {
				return fmt.Errorf("Error reading _dynamic_member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _dynamic_member: %v", err)
		}
	}

	if err = d.Set("comment", flattenWantempSystemSdwanMembersComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WantempSystemSdwanMembers-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("cost", flattenWantempSystemSdwanMembersCost2edl(o["cost"], d, "cost")); err != nil {
		if vv, ok := fortiAPIPatch(o["cost"], "WantempSystemSdwanMembers-Cost"); ok {
			if err = d.Set("cost", vv); err != nil {
				return fmt.Errorf("Error reading cost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cost: %v", err)
		}
	}

	if err = d.Set("duplication_threshold_bandwidth", flattenWantempSystemSdwanMembersDuplicationThresholdBandwidth2edl(o["duplication-threshold-bandwidth"], d, "duplication_threshold_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["duplication-threshold-bandwidth"], "WantempSystemSdwanMembers-DuplicationThresholdBandwidth"); ok {
			if err = d.Set("duplication_threshold_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading duplication_threshold_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading duplication_threshold_bandwidth: %v", err)
		}
	}

	if err = d.Set("duplication_threshold_bibandwidth", flattenWantempSystemSdwanMembersDuplicationThresholdBibandwidth2edl(o["duplication-threshold-bibandwidth"], d, "duplication_threshold_bibandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["duplication-threshold-bibandwidth"], "WantempSystemSdwanMembers-DuplicationThresholdBibandwidth"); ok {
			if err = d.Set("duplication_threshold_bibandwidth", vv); err != nil {
				return fmt.Errorf("Error reading duplication_threshold_bibandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading duplication_threshold_bibandwidth: %v", err)
		}
	}

	if err = d.Set("duplication_threshold_dwbandwidth", flattenWantempSystemSdwanMembersDuplicationThresholdDwbandwidth2edl(o["duplication-threshold-dwbandwidth"], d, "duplication_threshold_dwbandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["duplication-threshold-dwbandwidth"], "WantempSystemSdwanMembers-DuplicationThresholdDwbandwidth"); ok {
			if err = d.Set("duplication_threshold_dwbandwidth", vv); err != nil {
				return fmt.Errorf("Error reading duplication_threshold_dwbandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading duplication_threshold_dwbandwidth: %v", err)
		}
	}

	if err = d.Set("duplication_threshold_upbandwidth", flattenWantempSystemSdwanMembersDuplicationThresholdUpbandwidth2edl(o["duplication-threshold-upbandwidth"], d, "duplication_threshold_upbandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["duplication-threshold-upbandwidth"], "WantempSystemSdwanMembers-DuplicationThresholdUpbandwidth"); ok {
			if err = d.Set("duplication_threshold_upbandwidth", vv); err != nil {
				return fmt.Errorf("Error reading duplication_threshold_upbandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading duplication_threshold_upbandwidth: %v", err)
		}
	}

	if err = d.Set("gateway", flattenWantempSystemSdwanMembersGateway2edl(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "WantempSystemSdwanMembers-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("gateway6", flattenWantempSystemSdwanMembersGateway62edl(o["gateway6"], d, "gateway6")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway6"], "WantempSystemSdwanMembers-Gateway6"); ok {
			if err = d.Set("gateway6", vv); err != nil {
				return fmt.Errorf("Error reading gateway6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway6: %v", err)
		}
	}

	if err = d.Set("ingress_spillover_threshold", flattenWantempSystemSdwanMembersIngressSpilloverThreshold2edl(o["ingress-spillover-threshold"], d, "ingress_spillover_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ingress-spillover-threshold"], "WantempSystemSdwanMembers-IngressSpilloverThreshold"); ok {
			if err = d.Set("ingress_spillover_threshold", vv); err != nil {
				return fmt.Errorf("Error reading ingress_spillover_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ingress_spillover_threshold: %v", err)
		}
	}

	if err = d.Set("interface", flattenWantempSystemSdwanMembersInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "WantempSystemSdwanMembers-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("overage", flattenWantempSystemSdwanMembersOverage2edl(o["overage"], d, "overage")); err != nil {
		if vv, ok := fortiAPIPatch(o["overage"], "WantempSystemSdwanMembers-Overage"); ok {
			if err = d.Set("overage", vv); err != nil {
				return fmt.Errorf("Error reading overage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overage: %v", err)
		}
	}

	if err = d.Set("overage_cost", flattenWantempSystemSdwanMembersOverageCost2edl(o["overage-cost"], d, "overage_cost")); err != nil {
		if vv, ok := fortiAPIPatch(o["overage-cost"], "WantempSystemSdwanMembers-OverageCost"); ok {
			if err = d.Set("overage_cost", vv); err != nil {
				return fmt.Errorf("Error reading overage_cost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overage_cost: %v", err)
		}
	}

	if err = d.Set("overage_volume_ratio", flattenWantempSystemSdwanMembersOverageVolumeRatio2edl(o["overage-volume-ratio"], d, "overage_volume_ratio")); err != nil {
		if vv, ok := fortiAPIPatch(o["overage-volume-ratio"], "WantempSystemSdwanMembers-OverageVolumeRatio"); ok {
			if err = d.Set("overage_volume_ratio", vv); err != nil {
				return fmt.Errorf("Error reading overage_volume_ratio: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overage_volume_ratio: %v", err)
		}
	}

	if err = d.Set("overage_weight", flattenWantempSystemSdwanMembersOverageWeight2edl(o["overage-weight"], d, "overage_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["overage-weight"], "WantempSystemSdwanMembers-OverageWeight"); ok {
			if err = d.Set("overage_weight", vv); err != nil {
				return fmt.Errorf("Error reading overage_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overage_weight: %v", err)
		}
	}

	if err = d.Set("preferred_source", flattenWantempSystemSdwanMembersPreferredSource2edl(o["preferred-source"], d, "preferred_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["preferred-source"], "WantempSystemSdwanMembers-PreferredSource"); ok {
			if err = d.Set("preferred_source", vv); err != nil {
				return fmt.Errorf("Error reading preferred_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preferred_source: %v", err)
		}
	}

	if err = d.Set("priority", flattenWantempSystemSdwanMembersPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "WantempSystemSdwanMembers-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("priority_in_sla", flattenWantempSystemSdwanMembersPriorityInSla2edl(o["priority-in-sla"], d, "priority_in_sla")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-in-sla"], "WantempSystemSdwanMembers-PriorityInSla"); ok {
			if err = d.Set("priority_in_sla", vv); err != nil {
				return fmt.Errorf("Error reading priority_in_sla: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_in_sla: %v", err)
		}
	}

	if err = d.Set("priority_out_sla", flattenWantempSystemSdwanMembersPriorityOutSla2edl(o["priority-out-sla"], d, "priority_out_sla")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-out-sla"], "WantempSystemSdwanMembers-PriorityOutSla"); ok {
			if err = d.Set("priority_out_sla", vv); err != nil {
				return fmt.Errorf("Error reading priority_out_sla: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_out_sla: %v", err)
		}
	}

	if err = d.Set("priority6", flattenWantempSystemSdwanMembersPriority62edl(o["priority6"], d, "priority6")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority6"], "WantempSystemSdwanMembers-Priority6"); ok {
			if err = d.Set("priority6", vv); err != nil {
				return fmt.Errorf("Error reading priority6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority6: %v", err)
		}
	}

	if err = d.Set("quota_limit", flattenWantempSystemSdwanMembersQuotaLimit2edl(o["quota-limit"], d, "quota_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["quota-limit"], "WantempSystemSdwanMembers-QuotaLimit"); ok {
			if err = d.Set("quota_limit", vv); err != nil {
				return fmt.Errorf("Error reading quota_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quota_limit: %v", err)
		}
	}

	if err = d.Set("seq_num", flattenWantempSystemSdwanMembersSeqNum2edl(o["seq-num"], d, "seq_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["seq-num"], "WantempSystemSdwanMembers-SeqNum"); ok {
			if err = d.Set("seq_num", vv); err != nil {
				return fmt.Errorf("Error reading seq_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	if err = d.Set("source", flattenWantempSystemSdwanMembersSource2edl(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "WantempSystemSdwanMembers-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("source6", flattenWantempSystemSdwanMembersSource62edl(o["source6"], d, "source6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source6"], "WantempSystemSdwanMembers-Source6"); ok {
			if err = d.Set("source6", vv); err != nil {
				return fmt.Errorf("Error reading source6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source6: %v", err)
		}
	}

	if err = d.Set("spillover_threshold", flattenWantempSystemSdwanMembersSpilloverThreshold2edl(o["spillover-threshold"], d, "spillover_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["spillover-threshold"], "WantempSystemSdwanMembers-SpilloverThreshold"); ok {
			if err = d.Set("spillover_threshold", vv); err != nil {
				return fmt.Errorf("Error reading spillover_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spillover_threshold: %v", err)
		}
	}

	if err = d.Set("status", flattenWantempSystemSdwanMembersStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WantempSystemSdwanMembers-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("transport_group", flattenWantempSystemSdwanMembersTransportGroup2edl(o["transport-group"], d, "transport_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["transport-group"], "WantempSystemSdwanMembers-TransportGroup"); ok {
			if err = d.Set("transport_group", vv); err != nil {
				return fmt.Errorf("Error reading transport_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transport_group: %v", err)
		}
	}

	if err = d.Set("volume_ratio", flattenWantempSystemSdwanMembersVolumeRatio2edl(o["volume-ratio"], d, "volume_ratio")); err != nil {
		if vv, ok := fortiAPIPatch(o["volume-ratio"], "WantempSystemSdwanMembers-VolumeRatio"); ok {
			if err = d.Set("volume_ratio", vv); err != nil {
				return fmt.Errorf("Error reading volume_ratio: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading volume_ratio: %v", err)
		}
	}

	if err = d.Set("weight", flattenWantempSystemSdwanMembersWeight2edl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "WantempSystemSdwanMembers-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("zone", flattenWantempSystemSdwanMembersZone2edl(o["zone"], d, "zone")); err != nil {
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

func expandWantempSystemSdwanMembersBillingStartDay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersDynamicMember2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersCost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersDuplicationThresholdBandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersDuplicationThresholdBibandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersDuplicationThresholdDwbandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersDuplicationThresholdUpbandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersGateway2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersGateway62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersIngressSpilloverThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanMembersOverage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersOverageCost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersOverageVolumeRatio2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersOverageWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersPreferredSource2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersPriorityInSla2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersPriorityOutSla2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersPriority62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersQuotaLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSeqNum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSource2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSource62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSpilloverThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersTransportGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersVolumeRatio2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersZone2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectWantempSystemSdwanMembers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("billing_start_day"); ok || d.HasChange("billing_start_day") {
		t, err := expandWantempSystemSdwanMembersBillingStartDay2edl(d, v, "billing_start_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["billing-start-day"] = t
		}
	}

	if v, ok := d.GetOk("_dynamic_member"); ok || d.HasChange("_dynamic_member") {
		t, err := expandWantempSystemSdwanMembersDynamicMember2edl(d, v, "_dynamic_member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_dynamic-member"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWantempSystemSdwanMembersComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("cost"); ok || d.HasChange("cost") {
		t, err := expandWantempSystemSdwanMembersCost2edl(d, v, "cost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	}

	if v, ok := d.GetOk("duplication_threshold_bandwidth"); ok || d.HasChange("duplication_threshold_bandwidth") {
		t, err := expandWantempSystemSdwanMembersDuplicationThresholdBandwidth2edl(d, v, "duplication_threshold_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication-threshold-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("duplication_threshold_bibandwidth"); ok || d.HasChange("duplication_threshold_bibandwidth") {
		t, err := expandWantempSystemSdwanMembersDuplicationThresholdBibandwidth2edl(d, v, "duplication_threshold_bibandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication-threshold-bibandwidth"] = t
		}
	}

	if v, ok := d.GetOk("duplication_threshold_dwbandwidth"); ok || d.HasChange("duplication_threshold_dwbandwidth") {
		t, err := expandWantempSystemSdwanMembersDuplicationThresholdDwbandwidth2edl(d, v, "duplication_threshold_dwbandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication-threshold-dwbandwidth"] = t
		}
	}

	if v, ok := d.GetOk("duplication_threshold_upbandwidth"); ok || d.HasChange("duplication_threshold_upbandwidth") {
		t, err := expandWantempSystemSdwanMembersDuplicationThresholdUpbandwidth2edl(d, v, "duplication_threshold_upbandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication-threshold-upbandwidth"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandWantempSystemSdwanMembersGateway2edl(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("gateway6"); ok || d.HasChange("gateway6") {
		t, err := expandWantempSystemSdwanMembersGateway62edl(d, v, "gateway6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway6"] = t
		}
	}

	if v, ok := d.GetOk("ingress_spillover_threshold"); ok || d.HasChange("ingress_spillover_threshold") {
		t, err := expandWantempSystemSdwanMembersIngressSpilloverThreshold2edl(d, v, "ingress_spillover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress-spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandWantempSystemSdwanMembersInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("overage"); ok || d.HasChange("overage") {
		t, err := expandWantempSystemSdwanMembersOverage2edl(d, v, "overage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overage"] = t
		}
	}

	if v, ok := d.GetOk("overage_cost"); ok || d.HasChange("overage_cost") {
		t, err := expandWantempSystemSdwanMembersOverageCost2edl(d, v, "overage_cost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overage-cost"] = t
		}
	}

	if v, ok := d.GetOk("overage_volume_ratio"); ok || d.HasChange("overage_volume_ratio") {
		t, err := expandWantempSystemSdwanMembersOverageVolumeRatio2edl(d, v, "overage_volume_ratio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overage-volume-ratio"] = t
		}
	}

	if v, ok := d.GetOk("overage_weight"); ok || d.HasChange("overage_weight") {
		t, err := expandWantempSystemSdwanMembersOverageWeight2edl(d, v, "overage_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overage-weight"] = t
		}
	}

	if v, ok := d.GetOk("preferred_source"); ok || d.HasChange("preferred_source") {
		t, err := expandWantempSystemSdwanMembersPreferredSource2edl(d, v, "preferred_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-source"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandWantempSystemSdwanMembersPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("priority_in_sla"); ok || d.HasChange("priority_in_sla") {
		t, err := expandWantempSystemSdwanMembersPriorityInSla2edl(d, v, "priority_in_sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-in-sla"] = t
		}
	}

	if v, ok := d.GetOk("priority_out_sla"); ok || d.HasChange("priority_out_sla") {
		t, err := expandWantempSystemSdwanMembersPriorityOutSla2edl(d, v, "priority_out_sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-out-sla"] = t
		}
	}

	if v, ok := d.GetOk("priority6"); ok || d.HasChange("priority6") {
		t, err := expandWantempSystemSdwanMembersPriority62edl(d, v, "priority6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority6"] = t
		}
	}

	if v, ok := d.GetOk("quota_limit"); ok || d.HasChange("quota_limit") {
		t, err := expandWantempSystemSdwanMembersQuotaLimit2edl(d, v, "quota_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quota-limit"] = t
		}
	}

	if v, ok := d.GetOk("seq_num"); ok || d.HasChange("seq_num") {
		t, err := expandWantempSystemSdwanMembersSeqNum2edl(d, v, "seq_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq-num"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandWantempSystemSdwanMembersSource2edl(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("source6"); ok || d.HasChange("source6") {
		t, err := expandWantempSystemSdwanMembersSource62edl(d, v, "source6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source6"] = t
		}
	}

	if v, ok := d.GetOk("spillover_threshold"); ok || d.HasChange("spillover_threshold") {
		t, err := expandWantempSystemSdwanMembersSpilloverThreshold2edl(d, v, "spillover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWantempSystemSdwanMembersStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("transport_group"); ok || d.HasChange("transport_group") {
		t, err := expandWantempSystemSdwanMembersTransportGroup2edl(d, v, "transport_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transport-group"] = t
		}
	}

	if v, ok := d.GetOk("volume_ratio"); ok || d.HasChange("volume_ratio") {
		t, err := expandWantempSystemSdwanMembersVolumeRatio2edl(d, v, "volume_ratio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["volume-ratio"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandWantempSystemSdwanMembersWeight2edl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	if v, ok := d.GetOk("zone"); ok || d.HasChange("zone") {
		t, err := expandWantempSystemSdwanMembersZone2edl(d, v, "zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["zone"] = t
		}
	}

	return &obj, nil
}
