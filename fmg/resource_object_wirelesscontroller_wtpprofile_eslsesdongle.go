// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ESL SES-imagotag dongle configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerWtpProfileEslSesDongle() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerWtpProfileEslSesDongleUpdate,
		Read:   resourceObjectWirelessControllerWtpProfileEslSesDongleRead,
		Update: resourceObjectWirelessControllerWtpProfileEslSesDongleUpdate,
		Delete: resourceObjectWirelessControllerWtpProfileEslSesDongleDelete,

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
			"wtp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"apc_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apc_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apc_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apc_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"coex_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compliance_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"esl_channel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"output_power": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scd_enable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_cert_verification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_fqdn_verification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerWtpProfileEslSesDongleUpdate(d *schema.ResourceData, m interface{}) error {
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

	wtp_profile := d.Get("wtp_profile").(string)
	paradict["wtp_profile"] = wtp_profile

	obj, err := getObjectObjectWirelessControllerWtpProfileEslSesDongle(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWtpProfileEslSesDongle resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerWtpProfileEslSesDongle(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWtpProfileEslSesDongle resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectWirelessControllerWtpProfileEslSesDongle")

	return resourceObjectWirelessControllerWtpProfileEslSesDongleRead(d, m)
}

func resourceObjectWirelessControllerWtpProfileEslSesDongleDelete(d *schema.ResourceData, m interface{}) error {
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

	wtp_profile := d.Get("wtp_profile").(string)
	paradict["wtp_profile"] = wtp_profile

	err = c.DeleteObjectWirelessControllerWtpProfileEslSesDongle(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerWtpProfileEslSesDongle resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerWtpProfileEslSesDongleRead(d *schema.ResourceData, m interface{}) error {
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

	wtp_profile := d.Get("wtp_profile").(string)
	if wtp_profile == "" {
		wtp_profile = importOptionChecking(m.(*FortiClient).Cfg, "wtp_profile")
		if wtp_profile == "" {
			return fmt.Errorf("Parameter wtp_profile is missing")
		}
		if err = d.Set("wtp_profile", wtp_profile); err != nil {
			return fmt.Errorf("Error set params wtp_profile: %v", err)
		}
	}
	paradict["wtp_profile"] = wtp_profile

	o, err := c.ReadObjectWirelessControllerWtpProfileEslSesDongle(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWtpProfileEslSesDongle resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerWtpProfileEslSesDongle(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWtpProfileEslSesDongle resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerWtpProfileEslSesDongleApcAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileEslSesDongleApcFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileEslSesDongleApcIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileEslSesDongleApcPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileEslSesDongleCoexLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileEslSesDongleComplianceLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileEslSesDongleEslChannel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileEslSesDongleOutputPower2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileEslSesDongleScdEnable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileEslSesDongleTlsCertVerification2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileEslSesDongleTlsFqdnVerification2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerWtpProfileEslSesDongle(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("apc_addr_type", flattenObjectWirelessControllerWtpProfileEslSesDongleApcAddrType2edl(o["apc-addr-type"], d, "apc_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["apc-addr-type"], "ObjectWirelessControllerWtpProfileEslSesDongle-ApcAddrType"); ok {
			if err = d.Set("apc_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading apc_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apc_addr_type: %v", err)
		}
	}

	if err = d.Set("apc_fqdn", flattenObjectWirelessControllerWtpProfileEslSesDongleApcFqdn2edl(o["apc-fqdn"], d, "apc_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["apc-fqdn"], "ObjectWirelessControllerWtpProfileEslSesDongle-ApcFqdn"); ok {
			if err = d.Set("apc_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading apc_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apc_fqdn: %v", err)
		}
	}

	if err = d.Set("apc_ip", flattenObjectWirelessControllerWtpProfileEslSesDongleApcIp2edl(o["apc-ip"], d, "apc_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["apc-ip"], "ObjectWirelessControllerWtpProfileEslSesDongle-ApcIp"); ok {
			if err = d.Set("apc_ip", vv); err != nil {
				return fmt.Errorf("Error reading apc_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apc_ip: %v", err)
		}
	}

	if err = d.Set("apc_port", flattenObjectWirelessControllerWtpProfileEslSesDongleApcPort2edl(o["apc-port"], d, "apc_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["apc-port"], "ObjectWirelessControllerWtpProfileEslSesDongle-ApcPort"); ok {
			if err = d.Set("apc_port", vv); err != nil {
				return fmt.Errorf("Error reading apc_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apc_port: %v", err)
		}
	}

	if err = d.Set("coex_level", flattenObjectWirelessControllerWtpProfileEslSesDongleCoexLevel2edl(o["coex-level"], d, "coex_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["coex-level"], "ObjectWirelessControllerWtpProfileEslSesDongle-CoexLevel"); ok {
			if err = d.Set("coex_level", vv); err != nil {
				return fmt.Errorf("Error reading coex_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading coex_level: %v", err)
		}
	}

	if err = d.Set("compliance_level", flattenObjectWirelessControllerWtpProfileEslSesDongleComplianceLevel2edl(o["compliance-level"], d, "compliance_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["compliance-level"], "ObjectWirelessControllerWtpProfileEslSesDongle-ComplianceLevel"); ok {
			if err = d.Set("compliance_level", vv); err != nil {
				return fmt.Errorf("Error reading compliance_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading compliance_level: %v", err)
		}
	}

	if err = d.Set("esl_channel", flattenObjectWirelessControllerWtpProfileEslSesDongleEslChannel2edl(o["esl-channel"], d, "esl_channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["esl-channel"], "ObjectWirelessControllerWtpProfileEslSesDongle-EslChannel"); ok {
			if err = d.Set("esl_channel", vv); err != nil {
				return fmt.Errorf("Error reading esl_channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading esl_channel: %v", err)
		}
	}

	if err = d.Set("output_power", flattenObjectWirelessControllerWtpProfileEslSesDongleOutputPower2edl(o["output-power"], d, "output_power")); err != nil {
		if vv, ok := fortiAPIPatch(o["output-power"], "ObjectWirelessControllerWtpProfileEslSesDongle-OutputPower"); ok {
			if err = d.Set("output_power", vv); err != nil {
				return fmt.Errorf("Error reading output_power: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading output_power: %v", err)
		}
	}

	if err = d.Set("scd_enable", flattenObjectWirelessControllerWtpProfileEslSesDongleScdEnable2edl(o["scd-enable"], d, "scd_enable")); err != nil {
		if vv, ok := fortiAPIPatch(o["scd-enable"], "ObjectWirelessControllerWtpProfileEslSesDongle-ScdEnable"); ok {
			if err = d.Set("scd_enable", vv); err != nil {
				return fmt.Errorf("Error reading scd_enable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scd_enable: %v", err)
		}
	}

	if err = d.Set("tls_cert_verification", flattenObjectWirelessControllerWtpProfileEslSesDongleTlsCertVerification2edl(o["tls-cert-verification"], d, "tls_cert_verification")); err != nil {
		if vv, ok := fortiAPIPatch(o["tls-cert-verification"], "ObjectWirelessControllerWtpProfileEslSesDongle-TlsCertVerification"); ok {
			if err = d.Set("tls_cert_verification", vv); err != nil {
				return fmt.Errorf("Error reading tls_cert_verification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tls_cert_verification: %v", err)
		}
	}

	if err = d.Set("tls_fqdn_verification", flattenObjectWirelessControllerWtpProfileEslSesDongleTlsFqdnVerification2edl(o["tls-fqdn-verification"], d, "tls_fqdn_verification")); err != nil {
		if vv, ok := fortiAPIPatch(o["tls-fqdn-verification"], "ObjectWirelessControllerWtpProfileEslSesDongle-TlsFqdnVerification"); ok {
			if err = d.Set("tls_fqdn_verification", vv); err != nil {
				return fmt.Errorf("Error reading tls_fqdn_verification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tls_fqdn_verification: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerWtpProfileEslSesDongleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerWtpProfileEslSesDongleApcAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileEslSesDongleApcFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileEslSesDongleApcIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileEslSesDongleApcPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileEslSesDongleCoexLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileEslSesDongleComplianceLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileEslSesDongleEslChannel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileEslSesDongleOutputPower2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileEslSesDongleScdEnable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileEslSesDongleTlsCertVerification2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileEslSesDongleTlsFqdnVerification2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerWtpProfileEslSesDongle(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("apc_addr_type"); ok || d.HasChange("apc_addr_type") {
		t, err := expandObjectWirelessControllerWtpProfileEslSesDongleApcAddrType2edl(d, v, "apc_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apc-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("apc_fqdn"); ok || d.HasChange("apc_fqdn") {
		t, err := expandObjectWirelessControllerWtpProfileEslSesDongleApcFqdn2edl(d, v, "apc_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apc-fqdn"] = t
		}
	}

	if v, ok := d.GetOk("apc_ip"); ok || d.HasChange("apc_ip") {
		t, err := expandObjectWirelessControllerWtpProfileEslSesDongleApcIp2edl(d, v, "apc_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apc-ip"] = t
		}
	}

	if v, ok := d.GetOk("apc_port"); ok || d.HasChange("apc_port") {
		t, err := expandObjectWirelessControllerWtpProfileEslSesDongleApcPort2edl(d, v, "apc_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apc-port"] = t
		}
	}

	if v, ok := d.GetOk("coex_level"); ok || d.HasChange("coex_level") {
		t, err := expandObjectWirelessControllerWtpProfileEslSesDongleCoexLevel2edl(d, v, "coex_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["coex-level"] = t
		}
	}

	if v, ok := d.GetOk("compliance_level"); ok || d.HasChange("compliance_level") {
		t, err := expandObjectWirelessControllerWtpProfileEslSesDongleComplianceLevel2edl(d, v, "compliance_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compliance-level"] = t
		}
	}

	if v, ok := d.GetOk("esl_channel"); ok || d.HasChange("esl_channel") {
		t, err := expandObjectWirelessControllerWtpProfileEslSesDongleEslChannel2edl(d, v, "esl_channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["esl-channel"] = t
		}
	}

	if v, ok := d.GetOk("output_power"); ok || d.HasChange("output_power") {
		t, err := expandObjectWirelessControllerWtpProfileEslSesDongleOutputPower2edl(d, v, "output_power")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["output-power"] = t
		}
	}

	if v, ok := d.GetOk("scd_enable"); ok || d.HasChange("scd_enable") {
		t, err := expandObjectWirelessControllerWtpProfileEslSesDongleScdEnable2edl(d, v, "scd_enable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scd-enable"] = t
		}
	}

	if v, ok := d.GetOk("tls_cert_verification"); ok || d.HasChange("tls_cert_verification") {
		t, err := expandObjectWirelessControllerWtpProfileEslSesDongleTlsCertVerification2edl(d, v, "tls_cert_verification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-cert-verification"] = t
		}
	}

	if v, ok := d.GetOk("tls_fqdn_verification"); ok || d.HasChange("tls_fqdn_verification") {
		t, err := expandObjectWirelessControllerWtpProfileEslSesDongleTlsFqdnVerification2edl(d, v, "tls_fqdn_verification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-fqdn-verification"] = t
		}
	}

	return &obj, nil
}
