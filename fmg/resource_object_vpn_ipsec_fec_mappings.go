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
				ForceNew: true,
				Optional: true,
				Computed: true,
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

	_, err = c.CreateObjectVpnIpsecFecMappings(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnIpsecFecMappings resource: %v", err)
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

func flattenObjectVpnIpsecFecMappingsBandwidthDownThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsBandwidthUpThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsBase2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsLatencyThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsPacketLossThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsRedundant2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsSeqno2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnIpsecFecMappings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
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

	if err = d.Set("bandwidth_down_threshold", flattenObjectVpnIpsecFecMappingsBandwidthDownThreshold2edl(o["bandwidth-down-threshold"], d, "bandwidth_down_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-down-threshold"], "ObjectVpnIpsecFecMappings-BandwidthDownThreshold"); ok {
			if err = d.Set("bandwidth_down_threshold", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_down_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_down_threshold: %v", err)
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

	if err = d.Set("packet_loss_threshold", flattenObjectVpnIpsecFecMappingsPacketLossThreshold2edl(o["packet-loss-threshold"], d, "packet_loss_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-loss-threshold"], "ObjectVpnIpsecFecMappings-PacketLossThreshold"); ok {
			if err = d.Set("packet_loss_threshold", vv); err != nil {
				return fmt.Errorf("Error reading packet_loss_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_loss_threshold: %v", err)
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

func expandObjectVpnIpsecFecMappingsBandwidthDownThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsBandwidthUpThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsBase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsLatencyThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsPacketLossThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsRedundant2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsSeqno2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

	if v, ok := d.GetOk("bandwidth_down_threshold"); ok || d.HasChange("bandwidth_down_threshold") {
		t, err := expandObjectVpnIpsecFecMappingsBandwidthDownThreshold2edl(d, v, "bandwidth_down_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-down-threshold"] = t
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

	if v, ok := d.GetOk("packet_loss_threshold"); ok || d.HasChange("packet_loss_threshold") {
		t, err := expandObjectVpnIpsecFecMappingsPacketLossThreshold2edl(d, v, "packet_loss_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-loss-threshold"] = t
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

	return &obj, nil
}
