// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Global FortiAnalyzer Cloud settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystempLogFortianalyzerCloudSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystempLogFortianalyzerCloudSettingUpdate,
		Read:   resourceSystempLogFortianalyzerCloudSettingRead,
		Update: resourceSystempLogFortianalyzerCloudSettingUpdate,
		Delete: resourceSystempLogFortianalyzerCloudSettingDelete,

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
			"__change_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"access_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
				Computed: true,
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
			"serial": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
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
				Computed: true,
			},
			"upload_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystempLogFortianalyzerCloudSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystempLogFortianalyzerCloudSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating SystempLogFortianalyzerCloudSetting resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystempLogFortianalyzerCloudSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystempLogFortianalyzerCloudSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystempLogFortianalyzerCloudSetting")

	return resourceSystempLogFortianalyzerCloudSettingRead(d, m)
}

func resourceSystempLogFortianalyzerCloudSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystempLogFortianalyzerCloudSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystempLogFortianalyzerCloudSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystempLogFortianalyzerCloudSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystempLogFortianalyzerCloudSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystempLogFortianalyzerCloudSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystempLogFortianalyzerCloudSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystempLogFortianalyzerCloudSetting resource from API: %v", err)
	}
	return nil
}

func flattenSystempLogFortianalyzerCloudSettingChangeIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingAccessConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenSystempLogFortianalyzerCloudSettingCertificateVerification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingConnTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingHmacAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenSystempLogFortianalyzerCloudSettingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingIpsArchive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingMaxLogRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingMonitorFailureRetryPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingMonitorKeepalivePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingPresharedKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempLogFortianalyzerCloudSettingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingUploadDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingUploadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingUploadOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingUploadTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempLogFortianalyzerCloudSettingVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystempLogFortianalyzerCloudSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("__change_ip", flattenSystempLogFortianalyzerCloudSettingChangeIp(o["__change_ip"], d, "__change_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["__change_ip"], "SystempLogFortianalyzerCloudSetting-ChangeIp"); ok {
			if err = d.Set("__change_ip", vv); err != nil {
				return fmt.Errorf("Error reading __change_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading __change_ip: %v", err)
		}
	}

	if err = d.Set("access_config", flattenSystempLogFortianalyzerCloudSettingAccessConfig(o["access-config"], d, "access_config")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-config"], "SystempLogFortianalyzerCloudSetting-AccessConfig"); ok {
			if err = d.Set("access_config", vv); err != nil {
				return fmt.Errorf("Error reading access_config: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_config: %v", err)
		}
	}

	if err = d.Set("certificate", flattenSystempLogFortianalyzerCloudSettingCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "SystempLogFortianalyzerCloudSetting-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("certificate_verification", flattenSystempLogFortianalyzerCloudSettingCertificateVerification(o["certificate-verification"], d, "certificate_verification")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate-verification"], "SystempLogFortianalyzerCloudSetting-CertificateVerification"); ok {
			if err = d.Set("certificate_verification", vv); err != nil {
				return fmt.Errorf("Error reading certificate_verification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate_verification: %v", err)
		}
	}

	if err = d.Set("conn_timeout", flattenSystempLogFortianalyzerCloudSettingConnTimeout(o["conn-timeout"], d, "conn_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["conn-timeout"], "SystempLogFortianalyzerCloudSetting-ConnTimeout"); ok {
			if err = d.Set("conn_timeout", vv); err != nil {
				return fmt.Errorf("Error reading conn_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conn_timeout: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenSystempLogFortianalyzerCloudSettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["enc-algorithm"], "SystempLogFortianalyzerCloudSetting-EncAlgorithm"); ok {
			if err = d.Set("enc_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading enc_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("hmac_algorithm", flattenSystempLogFortianalyzerCloudSettingHmacAlgorithm(o["hmac-algorithm"], d, "hmac_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["hmac-algorithm"], "SystempLogFortianalyzerCloudSetting-HmacAlgorithm"); ok {
			if err = d.Set("hmac_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading hmac_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hmac_algorithm: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystempLogFortianalyzerCloudSettingInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystempLogFortianalyzerCloudSetting-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystempLogFortianalyzerCloudSettingInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "SystempLogFortianalyzerCloudSetting-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ips_archive", flattenSystempLogFortianalyzerCloudSettingIpsArchive(o["ips-archive"], d, "ips_archive")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-archive"], "SystempLogFortianalyzerCloudSetting-IpsArchive"); ok {
			if err = d.Set("ips_archive", vv); err != nil {
				return fmt.Errorf("Error reading ips_archive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_archive: %v", err)
		}
	}

	if err = d.Set("max_log_rate", flattenSystempLogFortianalyzerCloudSettingMaxLogRate(o["max-log-rate"], d, "max_log_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-log-rate"], "SystempLogFortianalyzerCloudSetting-MaxLogRate"); ok {
			if err = d.Set("max_log_rate", vv); err != nil {
				return fmt.Errorf("Error reading max_log_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_log_rate: %v", err)
		}
	}

	if err = d.Set("monitor_failure_retry_period", flattenSystempLogFortianalyzerCloudSettingMonitorFailureRetryPeriod(o["monitor-failure-retry-period"], d, "monitor_failure_retry_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-failure-retry-period"], "SystempLogFortianalyzerCloudSetting-MonitorFailureRetryPeriod"); ok {
			if err = d.Set("monitor_failure_retry_period", vv); err != nil {
				return fmt.Errorf("Error reading monitor_failure_retry_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_failure_retry_period: %v", err)
		}
	}

	if err = d.Set("monitor_keepalive_period", flattenSystempLogFortianalyzerCloudSettingMonitorKeepalivePeriod(o["monitor-keepalive-period"], d, "monitor_keepalive_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-keepalive-period"], "SystempLogFortianalyzerCloudSetting-MonitorKeepalivePeriod"); ok {
			if err = d.Set("monitor_keepalive_period", vv); err != nil {
				return fmt.Errorf("Error reading monitor_keepalive_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_keepalive_period: %v", err)
		}
	}

	if err = d.Set("preshared_key", flattenSystempLogFortianalyzerCloudSettingPresharedKey(o["preshared-key"], d, "preshared_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["preshared-key"], "SystempLogFortianalyzerCloudSetting-PresharedKey"); ok {
			if err = d.Set("preshared_key", vv); err != nil {
				return fmt.Errorf("Error reading preshared_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preshared_key: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystempLogFortianalyzerCloudSettingPriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "SystempLogFortianalyzerCloudSetting-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("serial", flattenSystempLogFortianalyzerCloudSettingSerial(o["serial"], d, "serial")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial"], "SystempLogFortianalyzerCloudSetting-Serial"); ok {
			if err = d.Set("serial", vv); err != nil {
				return fmt.Errorf("Error reading serial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystempLogFortianalyzerCloudSettingSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "SystempLogFortianalyzerCloudSetting-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenSystempLogFortianalyzerCloudSettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-proto-version"], "SystempLogFortianalyzerCloudSetting-SslMinProtoVersion"); ok {
			if err = d.Set("ssl_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("status", flattenSystempLogFortianalyzerCloudSettingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystempLogFortianalyzerCloudSetting-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upload_day", flattenSystempLogFortianalyzerCloudSettingUploadDay(o["upload-day"], d, "upload_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-day"], "SystempLogFortianalyzerCloudSetting-UploadDay"); ok {
			if err = d.Set("upload_day", vv); err != nil {
				return fmt.Errorf("Error reading upload_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_day: %v", err)
		}
	}

	if err = d.Set("upload_interval", flattenSystempLogFortianalyzerCloudSettingUploadInterval(o["upload-interval"], d, "upload_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-interval"], "SystempLogFortianalyzerCloudSetting-UploadInterval"); ok {
			if err = d.Set("upload_interval", vv); err != nil {
				return fmt.Errorf("Error reading upload_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_interval: %v", err)
		}
	}

	if err = d.Set("upload_option", flattenSystempLogFortianalyzerCloudSettingUploadOption(o["upload-option"], d, "upload_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-option"], "SystempLogFortianalyzerCloudSetting-UploadOption"); ok {
			if err = d.Set("upload_option", vv); err != nil {
				return fmt.Errorf("Error reading upload_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_option: %v", err)
		}
	}

	if err = d.Set("upload_time", flattenSystempLogFortianalyzerCloudSettingUploadTime(o["upload-time"], d, "upload_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-time"], "SystempLogFortianalyzerCloudSetting-UploadTime"); ok {
			if err = d.Set("upload_time", vv); err != nil {
				return fmt.Errorf("Error reading upload_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_time: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenSystempLogFortianalyzerCloudSettingVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "SystempLogFortianalyzerCloudSetting-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenSystempLogFortianalyzerCloudSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystempLogFortianalyzerCloudSettingChangeIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingAccessConfig(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandSystempLogFortianalyzerCloudSettingCertificateVerification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingConnTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingHmacAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandSystempLogFortianalyzerCloudSettingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingIpsArchive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingMaxLogRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingMonitorFailureRetryPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingMonitorKeepalivePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingPresharedKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingSerial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempLogFortianalyzerCloudSettingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingUploadDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingUploadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingUploadOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingUploadTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempLogFortianalyzerCloudSettingVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystempLogFortianalyzerCloudSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("__change_ip"); ok || d.HasChange("__change_ip") {
		t, err := expandSystempLogFortianalyzerCloudSettingChangeIp(d, v, "__change_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["__change_ip"] = t
		}
	}

	if v, ok := d.GetOk("access_config"); ok || d.HasChange("access_config") {
		t, err := expandSystempLogFortianalyzerCloudSettingAccessConfig(d, v, "access_config")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-config"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandSystempLogFortianalyzerCloudSettingCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("certificate_verification"); ok || d.HasChange("certificate_verification") {
		t, err := expandSystempLogFortianalyzerCloudSettingCertificateVerification(d, v, "certificate_verification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate-verification"] = t
		}
	}

	if v, ok := d.GetOk("conn_timeout"); ok || d.HasChange("conn_timeout") {
		t, err := expandSystempLogFortianalyzerCloudSettingConnTimeout(d, v, "conn_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-timeout"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok || d.HasChange("enc_algorithm") {
		t, err := expandSystempLogFortianalyzerCloudSettingEncAlgorithm(d, v, "enc_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("hmac_algorithm"); ok || d.HasChange("hmac_algorithm") {
		t, err := expandSystempLogFortianalyzerCloudSettingHmacAlgorithm(d, v, "hmac_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hmac-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystempLogFortianalyzerCloudSettingInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandSystempLogFortianalyzerCloudSettingInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ips_archive"); ok || d.HasChange("ips_archive") {
		t, err := expandSystempLogFortianalyzerCloudSettingIpsArchive(d, v, "ips_archive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-archive"] = t
		}
	}

	if v, ok := d.GetOk("max_log_rate"); ok || d.HasChange("max_log_rate") {
		t, err := expandSystempLogFortianalyzerCloudSettingMaxLogRate(d, v, "max_log_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-rate"] = t
		}
	}

	if v, ok := d.GetOk("monitor_failure_retry_period"); ok || d.HasChange("monitor_failure_retry_period") {
		t, err := expandSystempLogFortianalyzerCloudSettingMonitorFailureRetryPeriod(d, v, "monitor_failure_retry_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-failure-retry-period"] = t
		}
	}

	if v, ok := d.GetOk("monitor_keepalive_period"); ok || d.HasChange("monitor_keepalive_period") {
		t, err := expandSystempLogFortianalyzerCloudSettingMonitorKeepalivePeriod(d, v, "monitor_keepalive_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-keepalive-period"] = t
		}
	}

	if v, ok := d.GetOk("preshared_key"); ok || d.HasChange("preshared_key") {
		t, err := expandSystempLogFortianalyzerCloudSettingPresharedKey(d, v, "preshared_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preshared-key"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandSystempLogFortianalyzerCloudSettingPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("serial"); ok || d.HasChange("serial") {
		t, err := expandSystempLogFortianalyzerCloudSettingSerial(d, v, "serial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandSystempLogFortianalyzerCloudSettingSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok || d.HasChange("ssl_min_proto_version") {
		t, err := expandSystempLogFortianalyzerCloudSettingSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystempLogFortianalyzerCloudSettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("upload_day"); ok || d.HasChange("upload_day") {
		t, err := expandSystempLogFortianalyzerCloudSettingUploadDay(d, v, "upload_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-day"] = t
		}
	}

	if v, ok := d.GetOk("upload_interval"); ok || d.HasChange("upload_interval") {
		t, err := expandSystempLogFortianalyzerCloudSettingUploadInterval(d, v, "upload_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-interval"] = t
		}
	}

	if v, ok := d.GetOk("upload_option"); ok || d.HasChange("upload_option") {
		t, err := expandSystempLogFortianalyzerCloudSettingUploadOption(d, v, "upload_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-option"] = t
		}
	}

	if v, ok := d.GetOk("upload_time"); ok || d.HasChange("upload_time") {
		t, err := expandSystempLogFortianalyzerCloudSettingUploadTime(d, v, "upload_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-time"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandSystempLogFortianalyzerCloudSettingVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
