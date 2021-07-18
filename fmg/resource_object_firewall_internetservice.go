// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Show Internet Service application.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallInternetService() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallInternetServiceUpdate,
		Read:   resourceObjectFirewallInternetServiceRead,
		Update: resourceObjectFirewallInternetServiceUpdate,
		Delete: resourceObjectFirewallInternetServiceDelete,

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
			"city": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extra_ip_range_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"icon_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip_range_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"jitter_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"latency_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obsolete": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"packetloss_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reputation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"singularity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sld_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallInternetServiceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallInternetService(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetService resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallInternetService(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFirewallInternetService")

	return resourceObjectFirewallInternetServiceRead(d, m)
}

func resourceObjectFirewallInternetServiceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallInternetService(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallInternetService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallInternetServiceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallInternetService(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallInternetService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetService resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallInternetServiceCity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallInternetServiceCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallInternetServiceDatabase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceExtraIpRangeNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceIconId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceIpNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceIpRangeNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceJitterThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceLatencyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceObsolete(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallInternetServicePacketlossThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceReputation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceSingularity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceSldId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallInternetService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("city", flattenObjectFirewallInternetServiceCity(o["city"], d, "city")); err != nil {
		if vv, ok := fortiAPIPatch(o["city"], "ObjectFirewallInternetService-City"); ok {
			if err = d.Set("city", vv); err != nil {
				return fmt.Errorf("Error reading city: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading city: %v", err)
		}
	}

	if err = d.Set("country", flattenObjectFirewallInternetServiceCountry(o["country"], d, "country")); err != nil {
		if vv, ok := fortiAPIPatch(o["country"], "ObjectFirewallInternetService-Country"); ok {
			if err = d.Set("country", vv); err != nil {
				return fmt.Errorf("Error reading country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if err = d.Set("database", flattenObjectFirewallInternetServiceDatabase(o["database"], d, "database")); err != nil {
		if vv, ok := fortiAPIPatch(o["database"], "ObjectFirewallInternetService-Database"); ok {
			if err = d.Set("database", vv); err != nil {
				return fmt.Errorf("Error reading database: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading database: %v", err)
		}
	}

	if err = d.Set("direction", flattenObjectFirewallInternetServiceDirection(o["direction"], d, "direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["direction"], "ObjectFirewallInternetService-Direction"); ok {
			if err = d.Set("direction", vv); err != nil {
				return fmt.Errorf("Error reading direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading direction: %v", err)
		}
	}

	if err = d.Set("extra_ip_range_number", flattenObjectFirewallInternetServiceExtraIpRangeNumber(o["extra-ip-range-number"], d, "extra_ip_range_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["extra-ip-range-number"], "ObjectFirewallInternetService-ExtraIpRangeNumber"); ok {
			if err = d.Set("extra_ip_range_number", vv); err != nil {
				return fmt.Errorf("Error reading extra_ip_range_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extra_ip_range_number: %v", err)
		}
	}

	if err = d.Set("icon_id", flattenObjectFirewallInternetServiceIconId(o["icon-id"], d, "icon_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["icon-id"], "ObjectFirewallInternetService-IconId"); ok {
			if err = d.Set("icon_id", vv); err != nil {
				return fmt.Errorf("Error reading icon_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icon_id: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallInternetServiceId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallInternetService-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_number", flattenObjectFirewallInternetServiceIpNumber(o["ip-number"], d, "ip_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-number"], "ObjectFirewallInternetService-IpNumber"); ok {
			if err = d.Set("ip_number", vv); err != nil {
				return fmt.Errorf("Error reading ip_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_number: %v", err)
		}
	}

	if err = d.Set("ip_range_number", flattenObjectFirewallInternetServiceIpRangeNumber(o["ip-range-number"], d, "ip_range_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-range-number"], "ObjectFirewallInternetService-IpRangeNumber"); ok {
			if err = d.Set("ip_range_number", vv); err != nil {
				return fmt.Errorf("Error reading ip_range_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_range_number: %v", err)
		}
	}

	if err = d.Set("jitter_threshold", flattenObjectFirewallInternetServiceJitterThreshold(o["jitter-threshold"], d, "jitter_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["jitter-threshold"], "ObjectFirewallInternetService-JitterThreshold"); ok {
			if err = d.Set("jitter_threshold", vv); err != nil {
				return fmt.Errorf("Error reading jitter_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jitter_threshold: %v", err)
		}
	}

	if err = d.Set("latency_threshold", flattenObjectFirewallInternetServiceLatencyThreshold(o["latency-threshold"], d, "latency_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-threshold"], "ObjectFirewallInternetService-LatencyThreshold"); ok {
			if err = d.Set("latency_threshold", vv); err != nil {
				return fmt.Errorf("Error reading latency_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_threshold: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallInternetServiceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallInternetService-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("obsolete", flattenObjectFirewallInternetServiceObsolete(o["obsolete"], d, "obsolete")); err != nil {
		if vv, ok := fortiAPIPatch(o["obsolete"], "ObjectFirewallInternetService-Obsolete"); ok {
			if err = d.Set("obsolete", vv); err != nil {
				return fmt.Errorf("Error reading obsolete: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obsolete: %v", err)
		}
	}

	if err = d.Set("region", flattenObjectFirewallInternetServiceRegion(o["region"], d, "region")); err != nil {
		if vv, ok := fortiAPIPatch(o["region"], "ObjectFirewallInternetService-Region"); ok {
			if err = d.Set("region", vv); err != nil {
				return fmt.Errorf("Error reading region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading region: %v", err)
		}
	}

	if err = d.Set("packetloss_threshold", flattenObjectFirewallInternetServicePacketlossThreshold(o["packetloss-threshold"], d, "packetloss_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["packetloss-threshold"], "ObjectFirewallInternetService-PacketlossThreshold"); ok {
			if err = d.Set("packetloss_threshold", vv); err != nil {
				return fmt.Errorf("Error reading packetloss_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packetloss_threshold: %v", err)
		}
	}

	if err = d.Set("reputation", flattenObjectFirewallInternetServiceReputation(o["reputation"], d, "reputation")); err != nil {
		if vv, ok := fortiAPIPatch(o["reputation"], "ObjectFirewallInternetService-Reputation"); ok {
			if err = d.Set("reputation", vv); err != nil {
				return fmt.Errorf("Error reading reputation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reputation: %v", err)
		}
	}

	if err = d.Set("singularity", flattenObjectFirewallInternetServiceSingularity(o["singularity"], d, "singularity")); err != nil {
		if vv, ok := fortiAPIPatch(o["singularity"], "ObjectFirewallInternetService-Singularity"); ok {
			if err = d.Set("singularity", vv); err != nil {
				return fmt.Errorf("Error reading singularity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading singularity: %v", err)
		}
	}

	if err = d.Set("sld_id", flattenObjectFirewallInternetServiceSldId(o["sld-id"], d, "sld_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["sld-id"], "ObjectFirewallInternetService-SldId"); ok {
			if err = d.Set("sld_id", vv); err != nil {
				return fmt.Errorf("Error reading sld_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sld_id: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallInternetServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallInternetServiceCity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallInternetServiceCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallInternetServiceDatabase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceExtraIpRangeNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceIconId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceIpNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceIpRangeNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceJitterThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceLatencyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceObsolete(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallInternetServicePacketlossThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceReputation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceSingularity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceSldId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallInternetService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("city"); ok {
		t, err := expandObjectFirewallInternetServiceCity(d, v, "city")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["city"] = t
		}
	}

	if v, ok := d.GetOk("country"); ok {
		t, err := expandObjectFirewallInternetServiceCountry(d, v, "country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("database"); ok {
		t, err := expandObjectFirewallInternetServiceDatabase(d, v, "database")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database"] = t
		}
	}

	if v, ok := d.GetOk("direction"); ok {
		t, err := expandObjectFirewallInternetServiceDirection(d, v, "direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direction"] = t
		}
	}

	if v, ok := d.GetOk("extra_ip_range_number"); ok {
		t, err := expandObjectFirewallInternetServiceExtraIpRangeNumber(d, v, "extra_ip_range_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extra-ip-range-number"] = t
		}
	}

	if v, ok := d.GetOk("icon_id"); ok {
		t, err := expandObjectFirewallInternetServiceIconId(d, v, "icon_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icon-id"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandObjectFirewallInternetServiceId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_number"); ok {
		t, err := expandObjectFirewallInternetServiceIpNumber(d, v, "ip_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-number"] = t
		}
	}

	if v, ok := d.GetOk("ip_range_number"); ok {
		t, err := expandObjectFirewallInternetServiceIpRangeNumber(d, v, "ip_range_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range-number"] = t
		}
	}

	if v, ok := d.GetOk("jitter_threshold"); ok {
		t, err := expandObjectFirewallInternetServiceJitterThreshold(d, v, "jitter_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jitter-threshold"] = t
		}
	}

	if v, ok := d.GetOk("latency_threshold"); ok {
		t, err := expandObjectFirewallInternetServiceLatencyThreshold(d, v, "latency_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-threshold"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectFirewallInternetServiceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("obsolete"); ok {
		t, err := expandObjectFirewallInternetServiceObsolete(d, v, "obsolete")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obsolete"] = t
		}
	}

	if v, ok := d.GetOk("region"); ok {
		t, err := expandObjectFirewallInternetServiceRegion(d, v, "region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region"] = t
		}
	}

	if v, ok := d.GetOk("packetloss_threshold"); ok {
		t, err := expandObjectFirewallInternetServicePacketlossThreshold(d, v, "packetloss_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packetloss-threshold"] = t
		}
	}

	if v, ok := d.GetOk("reputation"); ok {
		t, err := expandObjectFirewallInternetServiceReputation(d, v, "reputation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation"] = t
		}
	}

	if v, ok := d.GetOk("singularity"); ok {
		t, err := expandObjectFirewallInternetServiceSingularity(d, v, "singularity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["singularity"] = t
		}
	}

	if v, ok := d.GetOk("sld_id"); ok {
		t, err := expandObjectFirewallInternetServiceSldId(d, v, "sld_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sld-id"] = t
		}
	}

	return &obj, nil
}
