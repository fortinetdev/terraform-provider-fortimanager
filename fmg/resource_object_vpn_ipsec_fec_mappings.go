// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FEC redundancy mapping table.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnIpsecFecMappings() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnIpsecFecMappingsCreate,
		Read:   resourceObjectVpnIpsecFecMappingsRead,
		Update: resourceObjectVpnIpsecFecMappingsUpdate,
		Delete: resourceObjectVpnIpsecFecMappingsDelete,

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
			"fec": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bandwidth_bi_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bandwidth_bi_threshold_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth_down_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bandwidth_down_threshold_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth_up_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bandwidth_up_threshold_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"base": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_threshold_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"packet_loss_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"packet_loss_threshold_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redundant": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"seqno": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"base": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"redundant": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"seqno": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tos": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tos_mask": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

func resourceObjectVpnIpsecFecMappingsCreate(d *schema.ResourceData, m interface{}) error {
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

	fec := d.Get("fec").(string)
	paradict["fec"] = fec

	obj, err := getObjectObjectVpnIpsecFecMappings(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnIpsecFecMappings resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("seqno")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectVpnIpsecFecMappings(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectVpnIpsecFecMappings(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectVpnIpsecFecMappings resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectVpnIpsecFecMappings(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectVpnIpsecFecMappings resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "seqno")))

	return resourceObjectVpnIpsecFecMappingsRead(d, m)
}

func resourceObjectVpnIpsecFecMappingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	fec := d.Get("fec").(string)
	paradict["fec"] = fec

	obj, err := getObjectObjectVpnIpsecFecMappings(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnIpsecFecMappings resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectVpnIpsecFecMappings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnIpsecFecMappings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "seqno")))

	return resourceObjectVpnIpsecFecMappingsRead(d, m)
}

func resourceObjectVpnIpsecFecMappingsDelete(d *schema.ResourceData, m interface{}) error {
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

	fec := d.Get("fec").(string)
	paradict["fec"] = fec

	wsParams["adom"] = adomv

	err = c.DeleteObjectVpnIpsecFecMappings(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnIpsecFecMappings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnIpsecFecMappingsRead(d *schema.ResourceData, m interface{}) error {
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

	fec := d.Get("fec").(string)
	if fec == "" {
		fec = importOptionChecking(m.(*FortiClient).Cfg, "fec")
		if fec == "" {
			return fmt.Errorf("Parameter fec is missing")
		}
		if err = d.Set("fec", fec); err != nil {
			return fmt.Errorf("Error set params fec: %v", err)
		}
	}
	paradict["fec"] = fec

	o, err := c.ReadObjectVpnIpsecFecMappings(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectVpnIpsecFecMappings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnIpsecFecMappings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnIpsecFecMappings resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnIpsecFecMappingsBandwidthBiThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsBandwidthBiThresholdNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsBandwidthDownThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsBandwidthDownThresholdNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsBandwidthUpThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsBandwidthUpThresholdNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsBase2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsLatencyThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsLatencyThresholdNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsPacketLossThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsPacketLossThresholdNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsRedundant2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsSeqno2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsTos2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base"
		if _, ok := i["base"]; ok {
			v := flattenObjectVpnIpsecFecMappingsTosBase2edl(i["base"], d, pre_append)
			tmp["base"] = fortiAPISubPartPatch(v, "ObjectVpnIpsecFecMappings-Tos-Base")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redundant"
		if _, ok := i["redundant"]; ok {
			v := flattenObjectVpnIpsecFecMappingsTosRedundant2edl(i["redundant"], d, pre_append)
			tmp["redundant"] = fortiAPISubPartPatch(v, "ObjectVpnIpsecFecMappings-Tos-Redundant")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seqno"
		if _, ok := i["seqno"]; ok {
			v := flattenObjectVpnIpsecFecMappingsTosSeqno2edl(i["seqno"], d, pre_append)
			tmp["seqno"] = fortiAPISubPartPatch(v, "ObjectVpnIpsecFecMappings-Tos-Seqno")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := i["tos"]; ok {
			v := flattenObjectVpnIpsecFecMappingsTosTos2edl(i["tos"], d, pre_append)
			tmp["tos"] = fortiAPISubPartPatch(v, "ObjectVpnIpsecFecMappings-Tos-Tos")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := i["tos-mask"]; ok {
			v := flattenObjectVpnIpsecFecMappingsTosTosMask2edl(i["tos-mask"], d, pre_append)
			tmp["tos_mask"] = fortiAPISubPartPatch(v, "ObjectVpnIpsecFecMappings-Tos-TosMask")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectVpnIpsecFecMappingsTosBase2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsTosRedundant2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsTosSeqno2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsTosTos2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsTosTosMask2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnIpsecFecMappings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("bandwidth_bi_threshold", flattenObjectVpnIpsecFecMappingsBandwidthBiThreshold2edl(o["bandwidth-bi-threshold"], d, "bandwidth_bi_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-bi-threshold"], "ObjectVpnIpsecFecMappings-BandwidthBiThreshold"); ok {
			if err = d.Set("bandwidth_bi_threshold", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_bi_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_bi_threshold: %v", err)
		}
	}

	if err = d.Set("bandwidth_bi_threshold_negate", flattenObjectVpnIpsecFecMappingsBandwidthBiThresholdNegate2edl(o["bandwidth-bi-threshold-negate"], d, "bandwidth_bi_threshold_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-bi-threshold-negate"], "ObjectVpnIpsecFecMappings-BandwidthBiThresholdNegate"); ok {
			if err = d.Set("bandwidth_bi_threshold_negate", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_bi_threshold_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_bi_threshold_negate: %v", err)
		}
	}

	if err = d.Set("bandwidth_down_threshold", flattenObjectVpnIpsecFecMappingsBandwidthDownThreshold2edl(o["bandwidth-down-threshold"], d, "bandwidth_down_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-down-threshold"], "ObjectVpnIpsecFecMappings-BandwidthDownThreshold"); ok {
			if err = d.Set("bandwidth_down_threshold", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_down_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_down_threshold: %v", err)
		}
	}

	if err = d.Set("bandwidth_down_threshold_negate", flattenObjectVpnIpsecFecMappingsBandwidthDownThresholdNegate2edl(o["bandwidth-down-threshold-negate"], d, "bandwidth_down_threshold_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-down-threshold-negate"], "ObjectVpnIpsecFecMappings-BandwidthDownThresholdNegate"); ok {
			if err = d.Set("bandwidth_down_threshold_negate", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_down_threshold_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_down_threshold_negate: %v", err)
		}
	}

	if err = d.Set("bandwidth_up_threshold", flattenObjectVpnIpsecFecMappingsBandwidthUpThreshold2edl(o["bandwidth-up-threshold"], d, "bandwidth_up_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-up-threshold"], "ObjectVpnIpsecFecMappings-BandwidthUpThreshold"); ok {
			if err = d.Set("bandwidth_up_threshold", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_up_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_up_threshold: %v", err)
		}
	}

	if err = d.Set("bandwidth_up_threshold_negate", flattenObjectVpnIpsecFecMappingsBandwidthUpThresholdNegate2edl(o["bandwidth-up-threshold-negate"], d, "bandwidth_up_threshold_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-up-threshold-negate"], "ObjectVpnIpsecFecMappings-BandwidthUpThresholdNegate"); ok {
			if err = d.Set("bandwidth_up_threshold_negate", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_up_threshold_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_up_threshold_negate: %v", err)
		}
	}

	if err = d.Set("base", flattenObjectVpnIpsecFecMappingsBase2edl(o["base"], d, "base")); err != nil {
		if vv, ok := fortiAPIPatch(o["base"], "ObjectVpnIpsecFecMappings-Base"); ok {
			if err = d.Set("base", vv); err != nil {
				return fmt.Errorf("Error reading base: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading base: %v", err)
		}
	}

	if err = d.Set("latency_threshold", flattenObjectVpnIpsecFecMappingsLatencyThreshold2edl(o["latency-threshold"], d, "latency_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-threshold"], "ObjectVpnIpsecFecMappings-LatencyThreshold"); ok {
			if err = d.Set("latency_threshold", vv); err != nil {
				return fmt.Errorf("Error reading latency_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_threshold: %v", err)
		}
	}

	if err = d.Set("latency_threshold_negate", flattenObjectVpnIpsecFecMappingsLatencyThresholdNegate2edl(o["latency-threshold-negate"], d, "latency_threshold_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-threshold-negate"], "ObjectVpnIpsecFecMappings-LatencyThresholdNegate"); ok {
			if err = d.Set("latency_threshold_negate", vv); err != nil {
				return fmt.Errorf("Error reading latency_threshold_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_threshold_negate: %v", err)
		}
	}

	if err = d.Set("packet_loss_threshold", flattenObjectVpnIpsecFecMappingsPacketLossThreshold2edl(o["packet-loss-threshold"], d, "packet_loss_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-loss-threshold"], "ObjectVpnIpsecFecMappings-PacketLossThreshold"); ok {
			if err = d.Set("packet_loss_threshold", vv); err != nil {
				return fmt.Errorf("Error reading packet_loss_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_loss_threshold: %v", err)
		}
	}

	if err = d.Set("packet_loss_threshold_negate", flattenObjectVpnIpsecFecMappingsPacketLossThresholdNegate2edl(o["packet-loss-threshold-negate"], d, "packet_loss_threshold_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-loss-threshold-negate"], "ObjectVpnIpsecFecMappings-PacketLossThresholdNegate"); ok {
			if err = d.Set("packet_loss_threshold_negate", vv); err != nil {
				return fmt.Errorf("Error reading packet_loss_threshold_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_loss_threshold_negate: %v", err)
		}
	}

	if err = d.Set("redundant", flattenObjectVpnIpsecFecMappingsRedundant2edl(o["redundant"], d, "redundant")); err != nil {
		if vv, ok := fortiAPIPatch(o["redundant"], "ObjectVpnIpsecFecMappings-Redundant"); ok {
			if err = d.Set("redundant", vv); err != nil {
				return fmt.Errorf("Error reading redundant: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redundant: %v", err)
		}
	}

	if err = d.Set("seqno", flattenObjectVpnIpsecFecMappingsSeqno2edl(o["seqno"], d, "seqno")); err != nil {
		if vv, ok := fortiAPIPatch(o["seqno"], "ObjectVpnIpsecFecMappings-Seqno"); ok {
			if err = d.Set("seqno", vv); err != nil {
				return fmt.Errorf("Error reading seqno: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seqno: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tos", flattenObjectVpnIpsecFecMappingsTos2edl(o["tos"], d, "tos")); err != nil {
			if vv, ok := fortiAPIPatch(o["tos"], "ObjectVpnIpsecFecMappings-Tos"); ok {
				if err = d.Set("tos", vv); err != nil {
					return fmt.Errorf("Error reading tos: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tos: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tos"); ok {
			if err = d.Set("tos", flattenObjectVpnIpsecFecMappingsTos2edl(o["tos"], d, "tos")); err != nil {
				if vv, ok := fortiAPIPatch(o["tos"], "ObjectVpnIpsecFecMappings-Tos"); ok {
					if err = d.Set("tos", vv); err != nil {
						return fmt.Errorf("Error reading tos: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tos: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectVpnIpsecFecMappingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnIpsecFecMappingsBandwidthBiThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsBandwidthBiThresholdNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsBandwidthDownThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsBandwidthDownThresholdNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsBandwidthUpThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsBandwidthUpThresholdNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsBase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsLatencyThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsLatencyThresholdNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsPacketLossThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsPacketLossThresholdNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsRedundant2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsSeqno2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsTos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["base"], _ = expandObjectVpnIpsecFecMappingsTosBase2edl(d, i["base"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redundant"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["redundant"], _ = expandObjectVpnIpsecFecMappingsTosRedundant2edl(d, i["redundant"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seqno"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seqno"], _ = expandObjectVpnIpsecFecMappingsTosSeqno2edl(d, i["seqno"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tos"], _ = expandObjectVpnIpsecFecMappingsTosTos2edl(d, i["tos"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tos-mask"], _ = expandObjectVpnIpsecFecMappingsTosTosMask2edl(d, i["tos_mask"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectVpnIpsecFecMappingsTosBase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsTosRedundant2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsTosSeqno2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsTosTos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsTosTosMask2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnIpsecFecMappings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bandwidth_bi_threshold"); ok || d.HasChange("bandwidth_bi_threshold") {
		t, err := expandObjectVpnIpsecFecMappingsBandwidthBiThreshold2edl(d, v, "bandwidth_bi_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-bi-threshold"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_bi_threshold_negate"); ok || d.HasChange("bandwidth_bi_threshold_negate") {
		t, err := expandObjectVpnIpsecFecMappingsBandwidthBiThresholdNegate2edl(d, v, "bandwidth_bi_threshold_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-bi-threshold-negate"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_down_threshold"); ok || d.HasChange("bandwidth_down_threshold") {
		t, err := expandObjectVpnIpsecFecMappingsBandwidthDownThreshold2edl(d, v, "bandwidth_down_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-down-threshold"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_down_threshold_negate"); ok || d.HasChange("bandwidth_down_threshold_negate") {
		t, err := expandObjectVpnIpsecFecMappingsBandwidthDownThresholdNegate2edl(d, v, "bandwidth_down_threshold_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-down-threshold-negate"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_up_threshold"); ok || d.HasChange("bandwidth_up_threshold") {
		t, err := expandObjectVpnIpsecFecMappingsBandwidthUpThreshold2edl(d, v, "bandwidth_up_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-up-threshold"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_up_threshold_negate"); ok || d.HasChange("bandwidth_up_threshold_negate") {
		t, err := expandObjectVpnIpsecFecMappingsBandwidthUpThresholdNegate2edl(d, v, "bandwidth_up_threshold_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-up-threshold-negate"] = t
		}
	}

	if v, ok := d.GetOk("base"); ok || d.HasChange("base") {
		t, err := expandObjectVpnIpsecFecMappingsBase2edl(d, v, "base")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["base"] = t
		}
	}

	if v, ok := d.GetOk("latency_threshold"); ok || d.HasChange("latency_threshold") {
		t, err := expandObjectVpnIpsecFecMappingsLatencyThreshold2edl(d, v, "latency_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-threshold"] = t
		}
	}

	if v, ok := d.GetOk("latency_threshold_negate"); ok || d.HasChange("latency_threshold_negate") {
		t, err := expandObjectVpnIpsecFecMappingsLatencyThresholdNegate2edl(d, v, "latency_threshold_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-threshold-negate"] = t
		}
	}

	if v, ok := d.GetOk("packet_loss_threshold"); ok || d.HasChange("packet_loss_threshold") {
		t, err := expandObjectVpnIpsecFecMappingsPacketLossThreshold2edl(d, v, "packet_loss_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-loss-threshold"] = t
		}
	}

	if v, ok := d.GetOk("packet_loss_threshold_negate"); ok || d.HasChange("packet_loss_threshold_negate") {
		t, err := expandObjectVpnIpsecFecMappingsPacketLossThresholdNegate2edl(d, v, "packet_loss_threshold_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-loss-threshold-negate"] = t
		}
	}

	if v, ok := d.GetOk("redundant"); ok || d.HasChange("redundant") {
		t, err := expandObjectVpnIpsecFecMappingsRedundant2edl(d, v, "redundant")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant"] = t
		}
	}

	if v, ok := d.GetOk("seqno"); ok || d.HasChange("seqno") {
		t, err := expandObjectVpnIpsecFecMappingsSeqno2edl(d, v, "seqno")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seqno"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok || d.HasChange("tos") {
		t, err := expandObjectVpnIpsecFecMappingsTos2edl(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	return &obj, nil
}
