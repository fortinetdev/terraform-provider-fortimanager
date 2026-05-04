// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure applications.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectTelemetryControllerProfileApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectTelemetryControllerProfileApplicationCreate,
		Read:   resourceObjectTelemetryControllerProfileApplicationRead,
		Update: resourceObjectTelemetryControllerProfileApplicationUpdate,
		Delete: resourceObjectTelemetryControllerProfileApplicationDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"app_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"app_throughput": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"atdt_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dns_time_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"experience_score_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"failure_rate_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"jitter_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"packet_loss_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sla": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"app_throughput_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"atdt_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dns_time_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"experience_score_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"failure_rate_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"jitter_threshold": &schema.Schema{
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
						"sla_factor": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"tcp_rtt_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tls_time_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ttfb_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"tcp_rtt_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tls_time_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ttfb_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectTelemetryControllerProfileApplicationCreate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectTelemetryControllerProfileApplication(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectTelemetryControllerProfileApplication resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectTelemetryControllerProfileApplication(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectTelemetryControllerProfileApplication(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectTelemetryControllerProfileApplication resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateObjectTelemetryControllerProfileApplication(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectTelemetryControllerProfileApplication resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceObjectTelemetryControllerProfileApplicationRead(d, m)
			} else {
				return fmt.Errorf("Error creating ObjectTelemetryControllerProfileApplication resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectTelemetryControllerProfileApplicationRead(d, m)
}

func resourceObjectTelemetryControllerProfileApplicationUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectTelemetryControllerProfileApplication(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectTelemetryControllerProfileApplication resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectTelemetryControllerProfileApplication(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectTelemetryControllerProfileApplication resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectTelemetryControllerProfileApplicationRead(d, m)
}

func resourceObjectTelemetryControllerProfileApplicationDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectTelemetryControllerProfileApplication(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectTelemetryControllerProfileApplication resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectTelemetryControllerProfileApplicationRead(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadObjectTelemetryControllerProfileApplication(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectTelemetryControllerProfileApplication resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectTelemetryControllerProfileApplication(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectTelemetryControllerProfileApplication resource from API: %v", err)
	}
	return nil
}

func flattenObjectTelemetryControllerProfileApplicationAppName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectTelemetryControllerProfileApplicationAppThroughput2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationAtdtThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationDnsTimeThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationExperienceScoreThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationFailureRateThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationJitterThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationLatencyThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationMonitor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationPacketLossThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSla2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "app_throughput_threshold"
	if _, ok := i["app-throughput-threshold"]; ok {
		result["app_throughput_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaAppThroughputThreshold2edl(i["app-throughput-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "atdt_threshold"
	if _, ok := i["atdt-threshold"]; ok {
		result["atdt_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaAtdtThreshold2edl(i["atdt-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_time_threshold"
	if _, ok := i["dns-time-threshold"]; ok {
		result["dns_time_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaDnsTimeThreshold2edl(i["dns-time-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "experience_score_threshold"
	if _, ok := i["experience-score-threshold"]; ok {
		result["experience_score_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaExperienceScoreThreshold2edl(i["experience-score-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "failure_rate_threshold"
	if _, ok := i["failure-rate-threshold"]; ok {
		result["failure_rate_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaFailureRateThreshold2edl(i["failure-rate-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "jitter_threshold"
	if _, ok := i["jitter-threshold"]; ok {
		result["jitter_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaJitterThreshold2edl(i["jitter-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "latency_threshold"
	if _, ok := i["latency-threshold"]; ok {
		result["latency_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaLatencyThreshold2edl(i["latency-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "packet_loss_threshold"
	if _, ok := i["packet-loss-threshold"]; ok {
		result["packet_loss_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaPacketLossThreshold2edl(i["packet-loss-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "sla_factor"
	if _, ok := i["sla-factor"]; ok {
		result["sla_factor"] = flattenObjectTelemetryControllerProfileApplicationSlaSlaFactor2edl(i["sla-factor"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_rtt_threshold"
	if _, ok := i["tcp-rtt-threshold"]; ok {
		result["tcp_rtt_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaTcpRttThreshold2edl(i["tcp-rtt-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "tls_time_threshold"
	if _, ok := i["tls-time-threshold"]; ok {
		result["tls_time_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaTlsTimeThreshold2edl(i["tls-time-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "ttfb_threshold"
	if _, ok := i["ttfb-threshold"]; ok {
		result["ttfb_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaTtfbThreshold2edl(i["ttfb-threshold"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectTelemetryControllerProfileApplicationSlaAppThroughputThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaAtdtThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaDnsTimeThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaExperienceScoreThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaFailureRateThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaJitterThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaLatencyThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaPacketLossThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaSlaFactor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectTelemetryControllerProfileApplicationSlaTcpRttThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaTlsTimeThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaTtfbThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationTcpRttThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationTlsTimeThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationTtfbThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectTelemetryControllerProfileApplication(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("app_name", flattenObjectTelemetryControllerProfileApplicationAppName2edl(o["app-name"], d, "app_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-name"], "ObjectTelemetryControllerProfileApplication-AppName"); ok {
			if err = d.Set("app_name", vv); err != nil {
				return fmt.Errorf("Error reading app_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_name: %v", err)
		}
	}

	if err = d.Set("app_throughput", flattenObjectTelemetryControllerProfileApplicationAppThroughput2edl(o["app-throughput"], d, "app_throughput")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-throughput"], "ObjectTelemetryControllerProfileApplication-AppThroughput"); ok {
			if err = d.Set("app_throughput", vv); err != nil {
				return fmt.Errorf("Error reading app_throughput: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_throughput: %v", err)
		}
	}

	if err = d.Set("atdt_threshold", flattenObjectTelemetryControllerProfileApplicationAtdtThreshold2edl(o["atdt-threshold"], d, "atdt_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["atdt-threshold"], "ObjectTelemetryControllerProfileApplication-AtdtThreshold"); ok {
			if err = d.Set("atdt_threshold", vv); err != nil {
				return fmt.Errorf("Error reading atdt_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading atdt_threshold: %v", err)
		}
	}

	if err = d.Set("dns_time_threshold", flattenObjectTelemetryControllerProfileApplicationDnsTimeThreshold2edl(o["dns-time-threshold"], d, "dns_time_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-time-threshold"], "ObjectTelemetryControllerProfileApplication-DnsTimeThreshold"); ok {
			if err = d.Set("dns_time_threshold", vv); err != nil {
				return fmt.Errorf("Error reading dns_time_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_time_threshold: %v", err)
		}
	}

	if err = d.Set("experience_score_threshold", flattenObjectTelemetryControllerProfileApplicationExperienceScoreThreshold2edl(o["experience-score-threshold"], d, "experience_score_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["experience-score-threshold"], "ObjectTelemetryControllerProfileApplication-ExperienceScoreThreshold"); ok {
			if err = d.Set("experience_score_threshold", vv); err != nil {
				return fmt.Errorf("Error reading experience_score_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading experience_score_threshold: %v", err)
		}
	}

	if err = d.Set("failure_rate_threshold", flattenObjectTelemetryControllerProfileApplicationFailureRateThreshold2edl(o["failure-rate-threshold"], d, "failure_rate_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["failure-rate-threshold"], "ObjectTelemetryControllerProfileApplication-FailureRateThreshold"); ok {
			if err = d.Set("failure_rate_threshold", vv); err != nil {
				return fmt.Errorf("Error reading failure_rate_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failure_rate_threshold: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectTelemetryControllerProfileApplicationId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectTelemetryControllerProfileApplication-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interval", flattenObjectTelemetryControllerProfileApplicationInterval2edl(o["interval"], d, "interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["interval"], "ObjectTelemetryControllerProfileApplication-Interval"); ok {
			if err = d.Set("interval", vv); err != nil {
				return fmt.Errorf("Error reading interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("jitter_threshold", flattenObjectTelemetryControllerProfileApplicationJitterThreshold2edl(o["jitter-threshold"], d, "jitter_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["jitter-threshold"], "ObjectTelemetryControllerProfileApplication-JitterThreshold"); ok {
			if err = d.Set("jitter_threshold", vv); err != nil {
				return fmt.Errorf("Error reading jitter_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jitter_threshold: %v", err)
		}
	}

	if err = d.Set("latency_threshold", flattenObjectTelemetryControllerProfileApplicationLatencyThreshold2edl(o["latency-threshold"], d, "latency_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-threshold"], "ObjectTelemetryControllerProfileApplication-LatencyThreshold"); ok {
			if err = d.Set("latency_threshold", vv); err != nil {
				return fmt.Errorf("Error reading latency_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_threshold: %v", err)
		}
	}

	if err = d.Set("monitor", flattenObjectTelemetryControllerProfileApplicationMonitor2edl(o["monitor"], d, "monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor"], "ObjectTelemetryControllerProfileApplication-Monitor"); ok {
			if err = d.Set("monitor", vv); err != nil {
				return fmt.Errorf("Error reading monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("packet_loss_threshold", flattenObjectTelemetryControllerProfileApplicationPacketLossThreshold2edl(o["packet-loss-threshold"], d, "packet_loss_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-loss-threshold"], "ObjectTelemetryControllerProfileApplication-PacketLossThreshold"); ok {
			if err = d.Set("packet_loss_threshold", vv); err != nil {
				return fmt.Errorf("Error reading packet_loss_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_loss_threshold: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sla", flattenObjectTelemetryControllerProfileApplicationSla2edl(o["sla"], d, "sla")); err != nil {
			if vv, ok := fortiAPIPatch(o["sla"], "ObjectTelemetryControllerProfileApplication-Sla"); ok {
				if err = d.Set("sla", vv); err != nil {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sla: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sla"); ok {
			if err = d.Set("sla", flattenObjectTelemetryControllerProfileApplicationSla2edl(o["sla"], d, "sla")); err != nil {
				if vv, ok := fortiAPIPatch(o["sla"], "ObjectTelemetryControllerProfileApplication-Sla"); ok {
					if err = d.Set("sla", vv); err != nil {
						return fmt.Errorf("Error reading sla: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			}
		}
	}

	if err = d.Set("tcp_rtt_threshold", flattenObjectTelemetryControllerProfileApplicationTcpRttThreshold2edl(o["tcp-rtt-threshold"], d, "tcp_rtt_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-rtt-threshold"], "ObjectTelemetryControllerProfileApplication-TcpRttThreshold"); ok {
			if err = d.Set("tcp_rtt_threshold", vv); err != nil {
				return fmt.Errorf("Error reading tcp_rtt_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_rtt_threshold: %v", err)
		}
	}

	if err = d.Set("tls_time_threshold", flattenObjectTelemetryControllerProfileApplicationTlsTimeThreshold2edl(o["tls-time-threshold"], d, "tls_time_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["tls-time-threshold"], "ObjectTelemetryControllerProfileApplication-TlsTimeThreshold"); ok {
			if err = d.Set("tls_time_threshold", vv); err != nil {
				return fmt.Errorf("Error reading tls_time_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tls_time_threshold: %v", err)
		}
	}

	if err = d.Set("ttfb_threshold", flattenObjectTelemetryControllerProfileApplicationTtfbThreshold2edl(o["ttfb-threshold"], d, "ttfb_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ttfb-threshold"], "ObjectTelemetryControllerProfileApplication-TtfbThreshold"); ok {
			if err = d.Set("ttfb_threshold", vv); err != nil {
				return fmt.Errorf("Error reading ttfb_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ttfb_threshold: %v", err)
		}
	}

	return nil
}

func flattenObjectTelemetryControllerProfileApplicationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectTelemetryControllerProfileApplicationAppName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectTelemetryControllerProfileApplicationAppThroughput2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationAtdtThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationDnsTimeThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationExperienceScoreThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationFailureRateThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationJitterThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationLatencyThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationMonitor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationPacketLossThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSla2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "app_throughput_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["app-throughput-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaAppThroughputThreshold2edl(d, i["app_throughput_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "atdt_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["atdt-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaAtdtThreshold2edl(d, i["atdt_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "dns_time_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-time-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaDnsTimeThreshold2edl(d, i["dns_time_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "experience_score_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["experience-score-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaExperienceScoreThreshold2edl(d, i["experience_score_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "failure_rate_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["failure-rate-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaFailureRateThreshold2edl(d, i["failure_rate_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "jitter_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["jitter-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaJitterThreshold2edl(d, i["jitter_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "latency_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["latency-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaLatencyThreshold2edl(d, i["latency_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "packet_loss_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["packet-loss-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaPacketLossThreshold2edl(d, i["packet_loss_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "sla_factor"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sla-factor"], _ = expandObjectTelemetryControllerProfileApplicationSlaSlaFactor2edl(d, i["sla_factor"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_rtt_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-rtt-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaTcpRttThreshold2edl(d, i["tcp_rtt_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "tls_time_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tls-time-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaTlsTimeThreshold2edl(d, i["tls_time_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "ttfb_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ttfb-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaTtfbThreshold2edl(d, i["ttfb_threshold"], pre_append)
	}

	return result, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaAppThroughputThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaAtdtThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaDnsTimeThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaExperienceScoreThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaFailureRateThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaJitterThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaLatencyThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaPacketLossThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaSlaFactor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectTelemetryControllerProfileApplicationSlaTcpRttThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaTlsTimeThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaTtfbThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationTcpRttThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationTlsTimeThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationTtfbThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectTelemetryControllerProfileApplication(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("app_name"); ok || d.HasChange("app_name") {
		t, err := expandObjectTelemetryControllerProfileApplicationAppName2edl(d, v, "app_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-name"] = t
		}
	}

	if v, ok := d.GetOk("app_throughput"); ok || d.HasChange("app_throughput") {
		t, err := expandObjectTelemetryControllerProfileApplicationAppThroughput2edl(d, v, "app_throughput")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-throughput"] = t
		}
	}

	if v, ok := d.GetOk("atdt_threshold"); ok || d.HasChange("atdt_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationAtdtThreshold2edl(d, v, "atdt_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["atdt-threshold"] = t
		}
	}

	if v, ok := d.GetOk("dns_time_threshold"); ok || d.HasChange("dns_time_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationDnsTimeThreshold2edl(d, v, "dns_time_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-time-threshold"] = t
		}
	}

	if v, ok := d.GetOk("experience_score_threshold"); ok || d.HasChange("experience_score_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationExperienceScoreThreshold2edl(d, v, "experience_score_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["experience-score-threshold"] = t
		}
	}

	if v, ok := d.GetOk("failure_rate_threshold"); ok || d.HasChange("failure_rate_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationFailureRateThreshold2edl(d, v, "failure_rate_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failure-rate-threshold"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectTelemetryControllerProfileApplicationId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok || d.HasChange("interval") {
		t, err := expandObjectTelemetryControllerProfileApplicationInterval2edl(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("jitter_threshold"); ok || d.HasChange("jitter_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationJitterThreshold2edl(d, v, "jitter_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jitter-threshold"] = t
		}
	}

	if v, ok := d.GetOk("latency_threshold"); ok || d.HasChange("latency_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationLatencyThreshold2edl(d, v, "latency_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-threshold"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok || d.HasChange("monitor") {
		t, err := expandObjectTelemetryControllerProfileApplicationMonitor2edl(d, v, "monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOk("packet_loss_threshold"); ok || d.HasChange("packet_loss_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationPacketLossThreshold2edl(d, v, "packet_loss_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-loss-threshold"] = t
		}
	}

	if v, ok := d.GetOk("sla"); ok || d.HasChange("sla") {
		t, err := expandObjectTelemetryControllerProfileApplicationSla2edl(d, v, "sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla"] = t
		}
	}

	if v, ok := d.GetOk("tcp_rtt_threshold"); ok || d.HasChange("tcp_rtt_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationTcpRttThreshold2edl(d, v, "tcp_rtt_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-rtt-threshold"] = t
		}
	}

	if v, ok := d.GetOk("tls_time_threshold"); ok || d.HasChange("tls_time_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationTlsTimeThreshold2edl(d, v, "tls_time_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-time-threshold"] = t
		}
	}

	if v, ok := d.GetOk("ttfb_threshold"); ok || d.HasChange("ttfb_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationTtfbThreshold2edl(d, v, "ttfb_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ttfb-threshold"] = t
		}
	}

	return &obj, nil
}
