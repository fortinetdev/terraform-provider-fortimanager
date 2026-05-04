// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiTelemetry profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectTelemetryControllerProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectTelemetryControllerProfileCreate,
		Read:   resourceObjectTelemetryControllerProfileRead,
		Update: resourceObjectTelemetryControllerProfileUpdate,
		Delete: resourceObjectTelemetryControllerProfileDelete,

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
			"application": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
				},
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func resourceObjectTelemetryControllerProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectTelemetryControllerProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectTelemetryControllerProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectTelemetryControllerProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectTelemetryControllerProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectTelemetryControllerProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectTelemetryControllerProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectTelemetryControllerProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectTelemetryControllerProfileRead(d, m)
}

func resourceObjectTelemetryControllerProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectTelemetryControllerProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectTelemetryControllerProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectTelemetryControllerProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectTelemetryControllerProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectTelemetryControllerProfileRead(d, m)
}

func resourceObjectTelemetryControllerProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectTelemetryControllerProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectTelemetryControllerProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectTelemetryControllerProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectTelemetryControllerProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectTelemetryControllerProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectTelemetryControllerProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectTelemetryControllerProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectTelemetryControllerProfileApplication(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_name"
		if _, ok := i["app-name"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationAppName(i["app-name"], d, pre_append)
			tmp["app_name"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-AppName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_throughput"
		if _, ok := i["app-throughput"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationAppThroughput(i["app-throughput"], d, pre_append)
			tmp["app_throughput"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-AppThroughput")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "atdt_threshold"
		if _, ok := i["atdt-threshold"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationAtdtThreshold(i["atdt-threshold"], d, pre_append)
			tmp["atdt_threshold"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-AtdtThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_time_threshold"
		if _, ok := i["dns-time-threshold"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationDnsTimeThreshold(i["dns-time-threshold"], d, pre_append)
			tmp["dns_time_threshold"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-DnsTimeThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "experience_score_threshold"
		if _, ok := i["experience-score-threshold"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationExperienceScoreThreshold(i["experience-score-threshold"], d, pre_append)
			tmp["experience_score_threshold"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-ExperienceScoreThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failure_rate_threshold"
		if _, ok := i["failure-rate-threshold"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationFailureRateThreshold(i["failure-rate-threshold"], d, pre_append)
			tmp["failure_rate_threshold"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-FailureRateThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := i["interval"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationInterval(i["interval"], d, pre_append)
			tmp["interval"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-Interval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := i["jitter-threshold"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationJitterThreshold(i["jitter-threshold"], d, pre_append)
			tmp["jitter_threshold"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-JitterThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationLatencyThreshold(i["latency-threshold"], d, pre_append)
			tmp["latency_threshold"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-LatencyThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_threshold"
		if _, ok := i["packet-loss-threshold"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationPacketLossThreshold(i["packet-loss-threshold"], d, pre_append)
			tmp["packet_loss_threshold"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-PacketLossThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := i["sla"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationSla(i["sla"], d, pre_append)
			tmp["sla"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-Sla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tcp_rtt_threshold"
		if _, ok := i["tcp-rtt-threshold"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationTcpRttThreshold(i["tcp-rtt-threshold"], d, pre_append)
			tmp["tcp_rtt_threshold"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-TcpRttThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tls_time_threshold"
		if _, ok := i["tls-time-threshold"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationTlsTimeThreshold(i["tls-time-threshold"], d, pre_append)
			tmp["tls_time_threshold"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-TlsTimeThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttfb_threshold"
		if _, ok := i["ttfb-threshold"]; ok {
			v := flattenObjectTelemetryControllerProfileApplicationTtfbThreshold(i["ttfb-threshold"], d, pre_append)
			tmp["ttfb_threshold"] = fortiAPISubPartPatch(v, "ObjectTelemetryControllerProfile-Application-TtfbThreshold")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectTelemetryControllerProfileApplicationAppName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectTelemetryControllerProfileApplicationAppThroughput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationAtdtThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationDnsTimeThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationExperienceScoreThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationFailureRateThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationJitterThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationLatencyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationPacketLossThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSla(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "app_throughput_threshold"
	if _, ok := i["app-throughput-threshold"]; ok {
		result["app_throughput_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaAppThroughputThreshold(i["app-throughput-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "atdt_threshold"
	if _, ok := i["atdt-threshold"]; ok {
		result["atdt_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaAtdtThreshold(i["atdt-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "dns_time_threshold"
	if _, ok := i["dns-time-threshold"]; ok {
		result["dns_time_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaDnsTimeThreshold(i["dns-time-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "experience_score_threshold"
	if _, ok := i["experience-score-threshold"]; ok {
		result["experience_score_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaExperienceScoreThreshold(i["experience-score-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "failure_rate_threshold"
	if _, ok := i["failure-rate-threshold"]; ok {
		result["failure_rate_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaFailureRateThreshold(i["failure-rate-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "jitter_threshold"
	if _, ok := i["jitter-threshold"]; ok {
		result["jitter_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaJitterThreshold(i["jitter-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "latency_threshold"
	if _, ok := i["latency-threshold"]; ok {
		result["latency_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaLatencyThreshold(i["latency-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "packet_loss_threshold"
	if _, ok := i["packet-loss-threshold"]; ok {
		result["packet_loss_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaPacketLossThreshold(i["packet-loss-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "sla_factor"
	if _, ok := i["sla-factor"]; ok {
		result["sla_factor"] = flattenObjectTelemetryControllerProfileApplicationSlaSlaFactor(i["sla-factor"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_rtt_threshold"
	if _, ok := i["tcp-rtt-threshold"]; ok {
		result["tcp_rtt_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaTcpRttThreshold(i["tcp-rtt-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "tls_time_threshold"
	if _, ok := i["tls-time-threshold"]; ok {
		result["tls_time_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaTlsTimeThreshold(i["tls-time-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "ttfb_threshold"
	if _, ok := i["ttfb-threshold"]; ok {
		result["ttfb_threshold"] = flattenObjectTelemetryControllerProfileApplicationSlaTtfbThreshold(i["ttfb-threshold"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectTelemetryControllerProfileApplicationSlaAppThroughputThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaAtdtThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaDnsTimeThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaExperienceScoreThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaFailureRateThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaJitterThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaLatencyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaPacketLossThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaSlaFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectTelemetryControllerProfileApplicationSlaTcpRttThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaTlsTimeThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaTtfbThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationTcpRttThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationTlsTimeThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationTtfbThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectTelemetryControllerProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("application", flattenObjectTelemetryControllerProfileApplication(o["application"], d, "application")); err != nil {
			if vv, ok := fortiAPIPatch(o["application"], "ObjectTelemetryControllerProfile-Application"); ok {
				if err = d.Set("application", vv); err != nil {
					return fmt.Errorf("Error reading application: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("application"); ok {
			if err = d.Set("application", flattenObjectTelemetryControllerProfileApplication(o["application"], d, "application")); err != nil {
				if vv, ok := fortiAPIPatch(o["application"], "ObjectTelemetryControllerProfile-Application"); ok {
					if err = d.Set("application", vv); err != nil {
						return fmt.Errorf("Error reading application: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading application: %v", err)
				}
			}
		}
	}

	if err = d.Set("comment", flattenObjectTelemetryControllerProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectTelemetryControllerProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectTelemetryControllerProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectTelemetryControllerProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectTelemetryControllerProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectTelemetryControllerProfileApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["app-name"], _ = expandObjectTelemetryControllerProfileApplicationAppName(d, i["app_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_throughput"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["app-throughput"], _ = expandObjectTelemetryControllerProfileApplicationAppThroughput(d, i["app_throughput"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "atdt_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["atdt-threshold"], _ = expandObjectTelemetryControllerProfileApplicationAtdtThreshold(d, i["atdt_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_time_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-time-threshold"], _ = expandObjectTelemetryControllerProfileApplicationDnsTimeThreshold(d, i["dns_time_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "experience_score_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["experience-score-threshold"], _ = expandObjectTelemetryControllerProfileApplicationExperienceScoreThreshold(d, i["experience_score_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failure_rate_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["failure-rate-threshold"], _ = expandObjectTelemetryControllerProfileApplicationFailureRateThreshold(d, i["failure_rate_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectTelemetryControllerProfileApplicationId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interval"], _ = expandObjectTelemetryControllerProfileApplicationInterval(d, i["interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-threshold"], _ = expandObjectTelemetryControllerProfileApplicationJitterThreshold(d, i["jitter_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-threshold"], _ = expandObjectTelemetryControllerProfileApplicationLatencyThreshold(d, i["latency_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandObjectTelemetryControllerProfileApplicationMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-loss-threshold"], _ = expandObjectTelemetryControllerProfileApplicationPacketLossThreshold(d, i["packet_loss_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectTelemetryControllerProfileApplicationSla(d, i["sla"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sla"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tcp_rtt_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tcp-rtt-threshold"], _ = expandObjectTelemetryControllerProfileApplicationTcpRttThreshold(d, i["tcp_rtt_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tls_time_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tls-time-threshold"], _ = expandObjectTelemetryControllerProfileApplicationTlsTimeThreshold(d, i["tls_time_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttfb_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ttfb-threshold"], _ = expandObjectTelemetryControllerProfileApplicationTtfbThreshold(d, i["ttfb_threshold"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectTelemetryControllerProfileApplicationAppName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectTelemetryControllerProfileApplicationAppThroughput(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationAtdtThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationDnsTimeThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationExperienceScoreThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationFailureRateThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationJitterThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationLatencyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationPacketLossThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "app_throughput_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["app-throughput-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaAppThroughputThreshold(d, i["app_throughput_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "atdt_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["atdt-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaAtdtThreshold(d, i["atdt_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "dns_time_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dns-time-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaDnsTimeThreshold(d, i["dns_time_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "experience_score_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["experience-score-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaExperienceScoreThreshold(d, i["experience_score_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "failure_rate_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["failure-rate-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaFailureRateThreshold(d, i["failure_rate_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "jitter_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["jitter-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaJitterThreshold(d, i["jitter_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "latency_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["latency-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaLatencyThreshold(d, i["latency_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "packet_loss_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["packet-loss-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaPacketLossThreshold(d, i["packet_loss_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "sla_factor"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sla-factor"], _ = expandObjectTelemetryControllerProfileApplicationSlaSlaFactor(d, i["sla_factor"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_rtt_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-rtt-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaTcpRttThreshold(d, i["tcp_rtt_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "tls_time_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tls-time-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaTlsTimeThreshold(d, i["tls_time_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "ttfb_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ttfb-threshold"], _ = expandObjectTelemetryControllerProfileApplicationSlaTtfbThreshold(d, i["ttfb_threshold"], pre_append)
	}

	return result, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaAppThroughputThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaAtdtThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaDnsTimeThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaExperienceScoreThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaFailureRateThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaJitterThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaLatencyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaPacketLossThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaSlaFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectTelemetryControllerProfileApplicationSlaTcpRttThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaTlsTimeThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaTtfbThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationTcpRttThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationTlsTimeThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationTtfbThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectTelemetryControllerProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandObjectTelemetryControllerProfileApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectTelemetryControllerProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectTelemetryControllerProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
