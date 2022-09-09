// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Forward Error Correction (FEC) mapping profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectVpnIpsecFec() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnIpsecFecCreate,
		Read:   resourceObjectVpnIpsecFecRead,
		Update: resourceObjectVpnIpsecFecUpdate,
		Delete: resourceObjectVpnIpsecFecDelete,

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
			"mappings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bandwidth_bi_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"bandwidth_down_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"bandwidth_up_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"base": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"latency_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"packet_loss_threshold": &schema.Schema{
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
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectVpnIpsecFecCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVpnIpsecFec(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnIpsecFec resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVpnIpsecFec(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnIpsecFec resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnIpsecFecRead(d, m)
}

func resourceObjectVpnIpsecFecUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVpnIpsecFec(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnIpsecFec resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVpnIpsecFec(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnIpsecFec resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnIpsecFecRead(d, m)
}

func resourceObjectVpnIpsecFecDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectVpnIpsecFec(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnIpsecFec resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnIpsecFecRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectVpnIpsecFec(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnIpsecFec resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnIpsecFec(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnIpsecFec resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnIpsecFecMappings(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_bi_threshold"
		if _, ok := i["bandwidth-bi-threshold"]; ok {
			v := flattenObjectVpnIpsecFecMappingsBandwidthBiThreshold(i["bandwidth-bi-threshold"], d, pre_append)
			tmp["bandwidth_bi_threshold"] = fortiAPISubPartPatch(v, "ObjectVpnIpsecFec-Mappings-BandwidthBiThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_down_threshold"
		if _, ok := i["bandwidth-down-threshold"]; ok {
			v := flattenObjectVpnIpsecFecMappingsBandwidthDownThreshold(i["bandwidth-down-threshold"], d, pre_append)
			tmp["bandwidth_down_threshold"] = fortiAPISubPartPatch(v, "ObjectVpnIpsecFec-Mappings-BandwidthDownThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_up_threshold"
		if _, ok := i["bandwidth-up-threshold"]; ok {
			v := flattenObjectVpnIpsecFecMappingsBandwidthUpThreshold(i["bandwidth-up-threshold"], d, pre_append)
			tmp["bandwidth_up_threshold"] = fortiAPISubPartPatch(v, "ObjectVpnIpsecFec-Mappings-BandwidthUpThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base"
		if _, ok := i["base"]; ok {
			v := flattenObjectVpnIpsecFecMappingsBase(i["base"], d, pre_append)
			tmp["base"] = fortiAPISubPartPatch(v, "ObjectVpnIpsecFec-Mappings-Base")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {
			v := flattenObjectVpnIpsecFecMappingsLatencyThreshold(i["latency-threshold"], d, pre_append)
			tmp["latency_threshold"] = fortiAPISubPartPatch(v, "ObjectVpnIpsecFec-Mappings-LatencyThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_threshold"
		if _, ok := i["packet-loss-threshold"]; ok {
			v := flattenObjectVpnIpsecFecMappingsPacketLossThreshold(i["packet-loss-threshold"], d, pre_append)
			tmp["packet_loss_threshold"] = fortiAPISubPartPatch(v, "ObjectVpnIpsecFec-Mappings-PacketLossThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redundant"
		if _, ok := i["redundant"]; ok {
			v := flattenObjectVpnIpsecFecMappingsRedundant(i["redundant"], d, pre_append)
			tmp["redundant"] = fortiAPISubPartPatch(v, "ObjectVpnIpsecFec-Mappings-Redundant")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seqno"
		if _, ok := i["seqno"]; ok {
			v := flattenObjectVpnIpsecFecMappingsSeqno(i["seqno"], d, pre_append)
			tmp["seqno"] = fortiAPISubPartPatch(v, "ObjectVpnIpsecFec-Mappings-Seqno")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectVpnIpsecFecMappingsBandwidthBiThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsBandwidthDownThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsBandwidthUpThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsBase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsLatencyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsPacketLossThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsRedundant(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsSeqno(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnIpsecFec(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("mappings", flattenObjectVpnIpsecFecMappings(o["mappings"], d, "mappings")); err != nil {
			if vv, ok := fortiAPIPatch(o["mappings"], "ObjectVpnIpsecFec-Mappings"); ok {
				if err = d.Set("mappings", vv); err != nil {
					return fmt.Errorf("Error reading mappings: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mappings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mappings"); ok {
			if err = d.Set("mappings", flattenObjectVpnIpsecFecMappings(o["mappings"], d, "mappings")); err != nil {
				if vv, ok := fortiAPIPatch(o["mappings"], "ObjectVpnIpsecFec-Mappings"); ok {
					if err = d.Set("mappings", vv); err != nil {
						return fmt.Errorf("Error reading mappings: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mappings: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectVpnIpsecFecName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVpnIpsecFec-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnIpsecFecFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnIpsecFecMappings(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_bi_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bandwidth-bi-threshold"], _ = expandObjectVpnIpsecFecMappingsBandwidthBiThreshold(d, i["bandwidth_bi_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_down_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bandwidth-down-threshold"], _ = expandObjectVpnIpsecFecMappingsBandwidthDownThreshold(d, i["bandwidth_down_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_up_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bandwidth-up-threshold"], _ = expandObjectVpnIpsecFecMappingsBandwidthUpThreshold(d, i["bandwidth_up_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["base"], _ = expandObjectVpnIpsecFecMappingsBase(d, i["base"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-threshold"], _ = expandObjectVpnIpsecFecMappingsLatencyThreshold(d, i["latency_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-loss-threshold"], _ = expandObjectVpnIpsecFecMappingsPacketLossThreshold(d, i["packet_loss_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redundant"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["redundant"], _ = expandObjectVpnIpsecFecMappingsRedundant(d, i["redundant"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seqno"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seqno"], _ = expandObjectVpnIpsecFecMappingsSeqno(d, i["seqno"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectVpnIpsecFecMappingsBandwidthBiThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsBandwidthDownThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsBandwidthUpThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsBase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsLatencyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsPacketLossThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsRedundant(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsSeqno(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnIpsecFec(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mappings"); ok || d.HasChange("mappings") {
		t, err := expandObjectVpnIpsecFecMappings(d, v, "mappings")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappings"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectVpnIpsecFecName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
