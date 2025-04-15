// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Global FortiAnalyzer settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystempLogFortianalyzerSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystempLogFortianalyzerSettingUpdate,
		Read:   resourceSystempLogFortianalyzerSettingRead,
		Update: resourceSystempLogFortianalyzerSettingUpdate,
		Delete: resourceSystempLogFortianalyzerSettingDelete,

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
			"devprof": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"access_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alt_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_verification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conn_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fallback_to_primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hmac_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ips_archive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_log_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"monitor_failure_retry_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"monitor_keepalive_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"preshared_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reliable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_cert_ca": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystempLogFortianalyzerSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	obj, err := getObjectSystempLogFortianalyzerSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating SystempLogFortianalyzerSetting resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystempLogFortianalyzerSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystempLogFortianalyzerSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystempLogFortianalyzerSetting")

	return resourceSystempLogFortianalyzerSettingRead(d, m)
}

func resourceSystempLogFortianalyzerSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	wsParams["adom"] = adomv

	err = c.DeleteSystempLogFortianalyzerSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystempLogFortianalyzerSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystempLogFortianalyzerSettingRead(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	if devprof == "" {
		devprof = importOptionChecking(m.(*FortiClient).Cfg, "devprof")
		if devprof == "" {
			return fmt.Errorf("Parameter devprof is missing")
		}
		if err = d.Set("devprof", devprof); err != nil {
			return fmt.Errorf("Error set params devprof: %v", err)
		}
	}
	paradict["devprof"] = devprof

	o, err := c.ReadSystempLogFortianalyzerSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystempLogFortianalyzerSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystempLogFortianalyzerSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystempLogFortianalyzerSetting resource from API: %v", err)
	}
	return nil
}

func flattenSystempLogFortianalyzerSettingAccessConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingAltServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenSystempLogFortianalyzerSettingCertificateVerification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingConnTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingFallbackToPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingHmacAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenSystempLogFortianalyzerSettingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingIpsArchive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingMaxLogRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingMonitorFailureRetryPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingMonitorKeepalivePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingPresharedKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingReliable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingServerCertCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenSystempLogFortianalyzerSettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingUploadDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingUploadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingUploadOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingUploadTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerSettingVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystempLogFortianalyzerSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("access_config", flattenSystempLogFortianalyzerSettingAccessConfig(o["access-config"], d, "access_config")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-config"], "SystempLogFortianalyzerSetting-AccessConfig"); ok {
			if err = d.Set("access_config", vv); err != nil {
				return fmt.Errorf("Error reading access_config: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_config: %v", err)
		}
	}

	if err = d.Set("alt_server", flattenSystempLogFortianalyzerSettingAltServer(o["alt-server"], d, "alt_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["alt-server"], "SystempLogFortianalyzerSetting-AltServer"); ok {
			if err = d.Set("alt_server", vv); err != nil {
				return fmt.Errorf("Error reading alt_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alt_server: %v", err)
		}
	}

	if err = d.Set("certificate", flattenSystempLogFortianalyzerSettingCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "SystempLogFortianalyzerSetting-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("certificate_verification", flattenSystempLogFortianalyzerSettingCertificateVerification(o["certificate-verification"], d, "certificate_verification")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate-verification"], "SystempLogFortianalyzerSetting-CertificateVerification"); ok {
			if err = d.Set("certificate_verification", vv); err != nil {
				return fmt.Errorf("Error reading certificate_verification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate_verification: %v", err)
		}
	}

	if err = d.Set("conn_timeout", flattenSystempLogFortianalyzerSettingConnTimeout(o["conn-timeout"], d, "conn_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["conn-timeout"], "SystempLogFortianalyzerSetting-ConnTimeout"); ok {
			if err = d.Set("conn_timeout", vv); err != nil {
				return fmt.Errorf("Error reading conn_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conn_timeout: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenSystempLogFortianalyzerSettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["enc-algorithm"], "SystempLogFortianalyzerSetting-EncAlgorithm"); ok {
			if err = d.Set("enc_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading enc_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("fallback_to_primary", flattenSystempLogFortianalyzerSettingFallbackToPrimary(o["fallback-to-primary"], d, "fallback_to_primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["fallback-to-primary"], "SystempLogFortianalyzerSetting-FallbackToPrimary"); ok {
			if err = d.Set("fallback_to_primary", vv); err != nil {
				return fmt.Errorf("Error reading fallback_to_primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fallback_to_primary: %v", err)
		}
	}

	if err = d.Set("hmac_algorithm", flattenSystempLogFortianalyzerSettingHmacAlgorithm(o["hmac-algorithm"], d, "hmac_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["hmac-algorithm"], "SystempLogFortianalyzerSetting-HmacAlgorithm"); ok {
			if err = d.Set("hmac_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading hmac_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hmac_algorithm: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystempLogFortianalyzerSettingInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystempLogFortianalyzerSetting-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystempLogFortianalyzerSettingInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "SystempLogFortianalyzerSetting-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ips_archive", flattenSystempLogFortianalyzerSettingIpsArchive(o["ips-archive"], d, "ips_archive")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-archive"], "SystempLogFortianalyzerSetting-IpsArchive"); ok {
			if err = d.Set("ips_archive", vv); err != nil {
				return fmt.Errorf("Error reading ips_archive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_archive: %v", err)
		}
	}

	if err = d.Set("max_log_rate", flattenSystempLogFortianalyzerSettingMaxLogRate(o["max-log-rate"], d, "max_log_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-log-rate"], "SystempLogFortianalyzerSetting-MaxLogRate"); ok {
			if err = d.Set("max_log_rate", vv); err != nil {
				return fmt.Errorf("Error reading max_log_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_log_rate: %v", err)
		}
	}

	if err = d.Set("monitor_failure_retry_period", flattenSystempLogFortianalyzerSettingMonitorFailureRetryPeriod(o["monitor-failure-retry-period"], d, "monitor_failure_retry_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-failure-retry-period"], "SystempLogFortianalyzerSetting-MonitorFailureRetryPeriod"); ok {
			if err = d.Set("monitor_failure_retry_period", vv); err != nil {
				return fmt.Errorf("Error reading monitor_failure_retry_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_failure_retry_period: %v", err)
		}
	}

	if err = d.Set("monitor_keepalive_period", flattenSystempLogFortianalyzerSettingMonitorKeepalivePeriod(o["monitor-keepalive-period"], d, "monitor_keepalive_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-keepalive-period"], "SystempLogFortianalyzerSetting-MonitorKeepalivePeriod"); ok {
			if err = d.Set("monitor_keepalive_period", vv); err != nil {
				return fmt.Errorf("Error reading monitor_keepalive_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_keepalive_period: %v", err)
		}
	}

	if err = d.Set("preshared_key", flattenSystempLogFortianalyzerSettingPresharedKey(o["preshared-key"], d, "preshared_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["preshared-key"], "SystempLogFortianalyzerSetting-PresharedKey"); ok {
			if err = d.Set("preshared_key", vv); err != nil {
				return fmt.Errorf("Error reading preshared_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preshared_key: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystempLogFortianalyzerSettingPriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "SystempLogFortianalyzerSetting-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("reliable", flattenSystempLogFortianalyzerSettingReliable(o["reliable"], d, "reliable")); err != nil {
		if vv, ok := fortiAPIPatch(o["reliable"], "SystempLogFortianalyzerSetting-Reliable"); ok {
			if err = d.Set("reliable", vv); err != nil {
				return fmt.Errorf("Error reading reliable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reliable: %v", err)
		}
	}

	if err = d.Set("server_cert_ca", flattenSystempLogFortianalyzerSettingServerCertCa(o["server-cert-ca"], d, "server_cert_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-cert-ca"], "SystempLogFortianalyzerSetting-ServerCertCa"); ok {
			if err = d.Set("server_cert_ca", vv); err != nil {
				return fmt.Errorf("Error reading server_cert_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_cert_ca: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenSystempLogFortianalyzerSettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-proto-version"], "SystempLogFortianalyzerSetting-SslMinProtoVersion"); ok {
			if err = d.Set("ssl_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("upload_day", flattenSystempLogFortianalyzerSettingUploadDay(o["upload-day"], d, "upload_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-day"], "SystempLogFortianalyzerSetting-UploadDay"); ok {
			if err = d.Set("upload_day", vv); err != nil {
				return fmt.Errorf("Error reading upload_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_day: %v", err)
		}
	}

	if err = d.Set("upload_interval", flattenSystempLogFortianalyzerSettingUploadInterval(o["upload-interval"], d, "upload_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-interval"], "SystempLogFortianalyzerSetting-UploadInterval"); ok {
			if err = d.Set("upload_interval", vv); err != nil {
				return fmt.Errorf("Error reading upload_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_interval: %v", err)
		}
	}

	if err = d.Set("upload_option", flattenSystempLogFortianalyzerSettingUploadOption(o["upload-option"], d, "upload_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-option"], "SystempLogFortianalyzerSetting-UploadOption"); ok {
			if err = d.Set("upload_option", vv); err != nil {
				return fmt.Errorf("Error reading upload_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_option: %v", err)
		}
	}

	if err = d.Set("upload_time", flattenSystempLogFortianalyzerSettingUploadTime(o["upload-time"], d, "upload_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-time"], "SystempLogFortianalyzerSetting-UploadTime"); ok {
			if err = d.Set("upload_time", vv); err != nil {
				return fmt.Errorf("Error reading upload_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_time: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenSystempLogFortianalyzerSettingVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "SystempLogFortianalyzerSetting-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenSystempLogFortianalyzerSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystempLogFortianalyzerSettingAccessConfig(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingAltServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandSystempLogFortianalyzerSettingCertificateVerification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingConnTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingFallbackToPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingHmacAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandSystempLogFortianalyzerSettingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingIpsArchive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingMaxLogRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingMonitorFailureRetryPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingMonitorKeepalivePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingPresharedKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingReliable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingServerCertCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandSystempLogFortianalyzerSettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingUploadDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingUploadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingUploadOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingUploadTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerSettingVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystempLogFortianalyzerSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_config"); ok || d.HasChange("access_config") {
		t, err := expandSystempLogFortianalyzerSettingAccessConfig(d, v, "access_config")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-config"] = t
		}
	}

	if v, ok := d.GetOk("alt_server"); ok || d.HasChange("alt_server") {
		t, err := expandSystempLogFortianalyzerSettingAltServer(d, v, "alt_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alt-server"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandSystempLogFortianalyzerSettingCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("certificate_verification"); ok || d.HasChange("certificate_verification") {
		t, err := expandSystempLogFortianalyzerSettingCertificateVerification(d, v, "certificate_verification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate-verification"] = t
		}
	}

	if v, ok := d.GetOk("conn_timeout"); ok || d.HasChange("conn_timeout") {
		t, err := expandSystempLogFortianalyzerSettingConnTimeout(d, v, "conn_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-timeout"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok || d.HasChange("enc_algorithm") {
		t, err := expandSystempLogFortianalyzerSettingEncAlgorithm(d, v, "enc_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("fallback_to_primary"); ok || d.HasChange("fallback_to_primary") {
		t, err := expandSystempLogFortianalyzerSettingFallbackToPrimary(d, v, "fallback_to_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fallback-to-primary"] = t
		}
	}

	if v, ok := d.GetOk("hmac_algorithm"); ok || d.HasChange("hmac_algorithm") {
		t, err := expandSystempLogFortianalyzerSettingHmacAlgorithm(d, v, "hmac_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hmac-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystempLogFortianalyzerSettingInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandSystempLogFortianalyzerSettingInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ips_archive"); ok || d.HasChange("ips_archive") {
		t, err := expandSystempLogFortianalyzerSettingIpsArchive(d, v, "ips_archive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-archive"] = t
		}
	}

	if v, ok := d.GetOk("max_log_rate"); ok || d.HasChange("max_log_rate") {
		t, err := expandSystempLogFortianalyzerSettingMaxLogRate(d, v, "max_log_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-rate"] = t
		}
	}

	if v, ok := d.GetOk("monitor_failure_retry_period"); ok || d.HasChange("monitor_failure_retry_period") {
		t, err := expandSystempLogFortianalyzerSettingMonitorFailureRetryPeriod(d, v, "monitor_failure_retry_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-failure-retry-period"] = t
		}
	}

	if v, ok := d.GetOk("monitor_keepalive_period"); ok || d.HasChange("monitor_keepalive_period") {
		t, err := expandSystempLogFortianalyzerSettingMonitorKeepalivePeriod(d, v, "monitor_keepalive_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-keepalive-period"] = t
		}
	}

	if v, ok := d.GetOk("preshared_key"); ok || d.HasChange("preshared_key") {
		t, err := expandSystempLogFortianalyzerSettingPresharedKey(d, v, "preshared_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preshared-key"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandSystempLogFortianalyzerSettingPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("reliable"); ok || d.HasChange("reliable") {
		t, err := expandSystempLogFortianalyzerSettingReliable(d, v, "reliable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reliable"] = t
		}
	}

	if v, ok := d.GetOk("server_cert_ca"); ok || d.HasChange("server_cert_ca") {
		t, err := expandSystempLogFortianalyzerSettingServerCertCa(d, v, "server_cert_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-cert-ca"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok || d.HasChange("ssl_min_proto_version") {
		t, err := expandSystempLogFortianalyzerSettingSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("upload_day"); ok || d.HasChange("upload_day") {
		t, err := expandSystempLogFortianalyzerSettingUploadDay(d, v, "upload_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-day"] = t
		}
	}

	if v, ok := d.GetOk("upload_interval"); ok || d.HasChange("upload_interval") {
		t, err := expandSystempLogFortianalyzerSettingUploadInterval(d, v, "upload_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-interval"] = t
		}
	}

	if v, ok := d.GetOk("upload_option"); ok || d.HasChange("upload_option") {
		t, err := expandSystempLogFortianalyzerSettingUploadOption(d, v, "upload_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-option"] = t
		}
	}

	if v, ok := d.GetOk("upload_time"); ok || d.HasChange("upload_time") {
		t, err := expandSystempLogFortianalyzerSettingUploadTime(d, v, "upload_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-time"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandSystempLogFortianalyzerSettingVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
