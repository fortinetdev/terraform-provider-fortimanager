// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Service level agreement (SLA).

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectTelemetryControllerProfileApplicationSla() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectTelemetryControllerProfileApplicationSlaUpdate,
		Read:   resourceObjectTelemetryControllerProfileApplicationSlaRead,
		Update: resourceObjectTelemetryControllerProfileApplicationSlaUpdate,
		Delete: resourceObjectTelemetryControllerProfileApplicationSlaDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
	}
}

func resourceObjectTelemetryControllerProfileApplicationSlaUpdate(d *schema.ResourceData, m interface{}) error {
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
	application := d.Get("application").(string)
	paradict["profile"] = profile
	paradict["application"] = application

	obj, err := getObjectObjectTelemetryControllerProfileApplicationSla(d, false)
	if err != nil {
		return fmt.Errorf("Error updating ObjectTelemetryControllerProfileApplicationSla resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectTelemetryControllerProfileApplicationSla(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectTelemetryControllerProfileApplicationSla resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectTelemetryControllerProfileApplicationSla")

	return resourceObjectTelemetryControllerProfileApplicationSlaRead(d, m)
}

func resourceObjectTelemetryControllerProfileApplicationSlaDelete(d *schema.ResourceData, m interface{}) error {
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
	application := d.Get("application").(string)
	paradict["profile"] = profile
	paradict["application"] = application

	obj, err := getObjectObjectTelemetryControllerProfileApplicationSla(d, true)

	if err != nil {
		return fmt.Errorf("Error updating ObjectTelemetryControllerProfileApplicationSla resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectTelemetryControllerProfileApplicationSla(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing ObjectTelemetryControllerProfileApplicationSla resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectTelemetryControllerProfileApplicationSlaRead(d *schema.ResourceData, m interface{}) error {
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
	application := d.Get("application").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	if application == "" {
		application = importOptionChecking(m.(*FortiClient).Cfg, "application")
		if application == "" {
			return fmt.Errorf("Parameter application is missing")
		}
		if err = d.Set("application", application); err != nil {
			return fmt.Errorf("Error set params application: %v", err)
		}
	}
	paradict["profile"] = profile
	paradict["application"] = application

	o, err := c.ReadObjectTelemetryControllerProfileApplicationSla(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectTelemetryControllerProfileApplicationSla resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectTelemetryControllerProfileApplicationSla(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectTelemetryControllerProfileApplicationSla resource from API: %v", err)
	}
	return nil
}

func flattenObjectTelemetryControllerProfileApplicationSlaAppThroughputThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaAtdtThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaDnsTimeThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaExperienceScoreThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaFailureRateThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaJitterThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaLatencyThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaPacketLossThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaSlaFactor3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectTelemetryControllerProfileApplicationSlaTcpRttThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaTlsTimeThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerProfileApplicationSlaTtfbThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectTelemetryControllerProfileApplicationSla(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("app_throughput_threshold", flattenObjectTelemetryControllerProfileApplicationSlaAppThroughputThreshold3rdl(o["app-throughput-threshold"], d, "app_throughput_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-throughput-threshold"], "ObjectTelemetryControllerProfileApplicationSla-AppThroughputThreshold"); ok {
			if err = d.Set("app_throughput_threshold", vv); err != nil {
				return fmt.Errorf("Error reading app_throughput_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_throughput_threshold: %v", err)
		}
	}

	if err = d.Set("atdt_threshold", flattenObjectTelemetryControllerProfileApplicationSlaAtdtThreshold3rdl(o["atdt-threshold"], d, "atdt_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["atdt-threshold"], "ObjectTelemetryControllerProfileApplicationSla-AtdtThreshold"); ok {
			if err = d.Set("atdt_threshold", vv); err != nil {
				return fmt.Errorf("Error reading atdt_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading atdt_threshold: %v", err)
		}
	}

	if err = d.Set("dns_time_threshold", flattenObjectTelemetryControllerProfileApplicationSlaDnsTimeThreshold3rdl(o["dns-time-threshold"], d, "dns_time_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-time-threshold"], "ObjectTelemetryControllerProfileApplicationSla-DnsTimeThreshold"); ok {
			if err = d.Set("dns_time_threshold", vv); err != nil {
				return fmt.Errorf("Error reading dns_time_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_time_threshold: %v", err)
		}
	}

	if err = d.Set("experience_score_threshold", flattenObjectTelemetryControllerProfileApplicationSlaExperienceScoreThreshold3rdl(o["experience-score-threshold"], d, "experience_score_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["experience-score-threshold"], "ObjectTelemetryControllerProfileApplicationSla-ExperienceScoreThreshold"); ok {
			if err = d.Set("experience_score_threshold", vv); err != nil {
				return fmt.Errorf("Error reading experience_score_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading experience_score_threshold: %v", err)
		}
	}

	if err = d.Set("failure_rate_threshold", flattenObjectTelemetryControllerProfileApplicationSlaFailureRateThreshold3rdl(o["failure-rate-threshold"], d, "failure_rate_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["failure-rate-threshold"], "ObjectTelemetryControllerProfileApplicationSla-FailureRateThreshold"); ok {
			if err = d.Set("failure_rate_threshold", vv); err != nil {
				return fmt.Errorf("Error reading failure_rate_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failure_rate_threshold: %v", err)
		}
	}

	if err = d.Set("jitter_threshold", flattenObjectTelemetryControllerProfileApplicationSlaJitterThreshold3rdl(o["jitter-threshold"], d, "jitter_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["jitter-threshold"], "ObjectTelemetryControllerProfileApplicationSla-JitterThreshold"); ok {
			if err = d.Set("jitter_threshold", vv); err != nil {
				return fmt.Errorf("Error reading jitter_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jitter_threshold: %v", err)
		}
	}

	if err = d.Set("latency_threshold", flattenObjectTelemetryControllerProfileApplicationSlaLatencyThreshold3rdl(o["latency-threshold"], d, "latency_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-threshold"], "ObjectTelemetryControllerProfileApplicationSla-LatencyThreshold"); ok {
			if err = d.Set("latency_threshold", vv); err != nil {
				return fmt.Errorf("Error reading latency_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_threshold: %v", err)
		}
	}

	if err = d.Set("packet_loss_threshold", flattenObjectTelemetryControllerProfileApplicationSlaPacketLossThreshold3rdl(o["packet-loss-threshold"], d, "packet_loss_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-loss-threshold"], "ObjectTelemetryControllerProfileApplicationSla-PacketLossThreshold"); ok {
			if err = d.Set("packet_loss_threshold", vv); err != nil {
				return fmt.Errorf("Error reading packet_loss_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_loss_threshold: %v", err)
		}
	}

	if err = d.Set("sla_factor", flattenObjectTelemetryControllerProfileApplicationSlaSlaFactor3rdl(o["sla-factor"], d, "sla_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-factor"], "ObjectTelemetryControllerProfileApplicationSla-SlaFactor"); ok {
			if err = d.Set("sla_factor", vv); err != nil {
				return fmt.Errorf("Error reading sla_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_factor: %v", err)
		}
	}

	if err = d.Set("tcp_rtt_threshold", flattenObjectTelemetryControllerProfileApplicationSlaTcpRttThreshold3rdl(o["tcp-rtt-threshold"], d, "tcp_rtt_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-rtt-threshold"], "ObjectTelemetryControllerProfileApplicationSla-TcpRttThreshold"); ok {
			if err = d.Set("tcp_rtt_threshold", vv); err != nil {
				return fmt.Errorf("Error reading tcp_rtt_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_rtt_threshold: %v", err)
		}
	}

	if err = d.Set("tls_time_threshold", flattenObjectTelemetryControllerProfileApplicationSlaTlsTimeThreshold3rdl(o["tls-time-threshold"], d, "tls_time_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["tls-time-threshold"], "ObjectTelemetryControllerProfileApplicationSla-TlsTimeThreshold"); ok {
			if err = d.Set("tls_time_threshold", vv); err != nil {
				return fmt.Errorf("Error reading tls_time_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tls_time_threshold: %v", err)
		}
	}

	if err = d.Set("ttfb_threshold", flattenObjectTelemetryControllerProfileApplicationSlaTtfbThreshold3rdl(o["ttfb-threshold"], d, "ttfb_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ttfb-threshold"], "ObjectTelemetryControllerProfileApplicationSla-TtfbThreshold"); ok {
			if err = d.Set("ttfb_threshold", vv); err != nil {
				return fmt.Errorf("Error reading ttfb_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ttfb_threshold: %v", err)
		}
	}

	return nil
}

func flattenObjectTelemetryControllerProfileApplicationSlaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectTelemetryControllerProfileApplicationSlaAppThroughputThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaAtdtThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaDnsTimeThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaExperienceScoreThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaFailureRateThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaJitterThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaLatencyThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaPacketLossThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaSlaFactor3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectTelemetryControllerProfileApplicationSlaTcpRttThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaTlsTimeThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerProfileApplicationSlaTtfbThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectTelemetryControllerProfileApplicationSla(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("app_throughput_threshold"); ok || d.HasChange("app_throughput_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationSlaAppThroughputThreshold3rdl(d, v, "app_throughput_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-throughput-threshold"] = t
		}
	}

	if v, ok := d.GetOk("atdt_threshold"); ok || d.HasChange("atdt_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationSlaAtdtThreshold3rdl(d, v, "atdt_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["atdt-threshold"] = t
		}
	}

	if v, ok := d.GetOk("dns_time_threshold"); ok || d.HasChange("dns_time_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationSlaDnsTimeThreshold3rdl(d, v, "dns_time_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-time-threshold"] = t
		}
	}

	if v, ok := d.GetOk("experience_score_threshold"); ok || d.HasChange("experience_score_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationSlaExperienceScoreThreshold3rdl(d, v, "experience_score_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["experience-score-threshold"] = t
		}
	}

	if v, ok := d.GetOk("failure_rate_threshold"); ok || d.HasChange("failure_rate_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationSlaFailureRateThreshold3rdl(d, v, "failure_rate_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failure-rate-threshold"] = t
		}
	}

	if v, ok := d.GetOk("jitter_threshold"); ok || d.HasChange("jitter_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationSlaJitterThreshold3rdl(d, v, "jitter_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jitter-threshold"] = t
		}
	}

	if v, ok := d.GetOk("latency_threshold"); ok || d.HasChange("latency_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationSlaLatencyThreshold3rdl(d, v, "latency_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-threshold"] = t
		}
	}

	if v, ok := d.GetOk("packet_loss_threshold"); ok || d.HasChange("packet_loss_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationSlaPacketLossThreshold3rdl(d, v, "packet_loss_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-loss-threshold"] = t
		}
	}

	if v, ok := d.GetOk("sla_factor"); ok || d.HasChange("sla_factor") {
		t, err := expandObjectTelemetryControllerProfileApplicationSlaSlaFactor3rdl(d, v, "sla_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-factor"] = t
		}
	}

	if v, ok := d.GetOk("tcp_rtt_threshold"); ok || d.HasChange("tcp_rtt_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationSlaTcpRttThreshold3rdl(d, v, "tcp_rtt_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-rtt-threshold"] = t
		}
	}

	if v, ok := d.GetOk("tls_time_threshold"); ok || d.HasChange("tls_time_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationSlaTlsTimeThreshold3rdl(d, v, "tls_time_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-time-threshold"] = t
		}
	}

	if v, ok := d.GetOk("ttfb_threshold"); ok || d.HasChange("ttfb_threshold") {
		t, err := expandObjectTelemetryControllerProfileApplicationSlaTtfbThreshold3rdl(d, v, "ttfb_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ttfb-threshold"] = t
		}
	}

	return &obj, nil
}
