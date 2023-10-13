// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: WTP LAN port mapping.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerWtpProfileLan() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerWtpProfileLanUpdate,
		Read:   resourceObjectWirelessControllerWtpProfileLanRead,
		Update: resourceObjectWirelessControllerWtpProfileLanUpdate,
		Delete: resourceObjectWirelessControllerWtpProfileLanDelete,

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
			"port_esl_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_esl_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port1_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port1_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port2_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port2_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port3_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port3_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port4_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port4_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port5_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port5_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port6_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port6_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port7_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port7_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port8_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port8_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectWirelessControllerWtpProfileLanUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerWtpProfileLan(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWtpProfileLan resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerWtpProfileLan(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWtpProfileLan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectWirelessControllerWtpProfileLan")

	return resourceObjectWirelessControllerWtpProfileLanRead(d, m)
}

func resourceObjectWirelessControllerWtpProfileLanDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWirelessControllerWtpProfileLan(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerWtpProfileLan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerWtpProfileLanRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWirelessControllerWtpProfileLan(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWtpProfileLan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerWtpProfileLan(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWtpProfileLan resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerWtpProfileLanPortEslMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPortEslSsid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWirelessControllerWtpProfileLanPortMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPortSsid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort1Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort1Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort2Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort2Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort3Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort3Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort4Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort4Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort5Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort5Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort6Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort6Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort7Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort7Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort8Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort8Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerWtpProfileLan(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("port_esl_mode", flattenObjectWirelessControllerWtpProfileLanPortEslMode2edl(o["port-esl-mode"], d, "port_esl_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-esl-mode"], "ObjectWirelessControllerWtpProfileLan-PortEslMode"); ok {
			if err = d.Set("port_esl_mode", vv); err != nil {
				return fmt.Errorf("Error reading port_esl_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_esl_mode: %v", err)
		}
	}

	if err = d.Set("port_esl_ssid", flattenObjectWirelessControllerWtpProfileLanPortEslSsid2edl(o["port-esl-ssid"], d, "port_esl_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-esl-ssid"], "ObjectWirelessControllerWtpProfileLan-PortEslSsid"); ok {
			if err = d.Set("port_esl_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port_esl_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_esl_ssid: %v", err)
		}
	}

	if err = d.Set("port_mode", flattenObjectWirelessControllerWtpProfileLanPortMode2edl(o["port-mode"], d, "port_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-mode"], "ObjectWirelessControllerWtpProfileLan-PortMode"); ok {
			if err = d.Set("port_mode", vv); err != nil {
				return fmt.Errorf("Error reading port_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_mode: %v", err)
		}
	}

	if err = d.Set("port_ssid", flattenObjectWirelessControllerWtpProfileLanPortSsid2edl(o["port-ssid"], d, "port_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-ssid"], "ObjectWirelessControllerWtpProfileLan-PortSsid"); ok {
			if err = d.Set("port_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_ssid: %v", err)
		}
	}

	if err = d.Set("port1_mode", flattenObjectWirelessControllerWtpProfileLanPort1Mode2edl(o["port1-mode"], d, "port1_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port1-mode"], "ObjectWirelessControllerWtpProfileLan-Port1Mode"); ok {
			if err = d.Set("port1_mode", vv); err != nil {
				return fmt.Errorf("Error reading port1_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port1_mode: %v", err)
		}
	}

	if err = d.Set("port1_ssid", flattenObjectWirelessControllerWtpProfileLanPort1Ssid2edl(o["port1-ssid"], d, "port1_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port1-ssid"], "ObjectWirelessControllerWtpProfileLan-Port1Ssid"); ok {
			if err = d.Set("port1_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port1_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port1_ssid: %v", err)
		}
	}

	if err = d.Set("port2_mode", flattenObjectWirelessControllerWtpProfileLanPort2Mode2edl(o["port2-mode"], d, "port2_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port2-mode"], "ObjectWirelessControllerWtpProfileLan-Port2Mode"); ok {
			if err = d.Set("port2_mode", vv); err != nil {
				return fmt.Errorf("Error reading port2_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port2_mode: %v", err)
		}
	}

	if err = d.Set("port2_ssid", flattenObjectWirelessControllerWtpProfileLanPort2Ssid2edl(o["port2-ssid"], d, "port2_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port2-ssid"], "ObjectWirelessControllerWtpProfileLan-Port2Ssid"); ok {
			if err = d.Set("port2_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port2_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port2_ssid: %v", err)
		}
	}

	if err = d.Set("port3_mode", flattenObjectWirelessControllerWtpProfileLanPort3Mode2edl(o["port3-mode"], d, "port3_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port3-mode"], "ObjectWirelessControllerWtpProfileLan-Port3Mode"); ok {
			if err = d.Set("port3_mode", vv); err != nil {
				return fmt.Errorf("Error reading port3_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port3_mode: %v", err)
		}
	}

	if err = d.Set("port3_ssid", flattenObjectWirelessControllerWtpProfileLanPort3Ssid2edl(o["port3-ssid"], d, "port3_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port3-ssid"], "ObjectWirelessControllerWtpProfileLan-Port3Ssid"); ok {
			if err = d.Set("port3_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port3_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port3_ssid: %v", err)
		}
	}

	if err = d.Set("port4_mode", flattenObjectWirelessControllerWtpProfileLanPort4Mode2edl(o["port4-mode"], d, "port4_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port4-mode"], "ObjectWirelessControllerWtpProfileLan-Port4Mode"); ok {
			if err = d.Set("port4_mode", vv); err != nil {
				return fmt.Errorf("Error reading port4_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port4_mode: %v", err)
		}
	}

	if err = d.Set("port4_ssid", flattenObjectWirelessControllerWtpProfileLanPort4Ssid2edl(o["port4-ssid"], d, "port4_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port4-ssid"], "ObjectWirelessControllerWtpProfileLan-Port4Ssid"); ok {
			if err = d.Set("port4_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port4_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port4_ssid: %v", err)
		}
	}

	if err = d.Set("port5_mode", flattenObjectWirelessControllerWtpProfileLanPort5Mode2edl(o["port5-mode"], d, "port5_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port5-mode"], "ObjectWirelessControllerWtpProfileLan-Port5Mode"); ok {
			if err = d.Set("port5_mode", vv); err != nil {
				return fmt.Errorf("Error reading port5_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port5_mode: %v", err)
		}
	}

	if err = d.Set("port5_ssid", flattenObjectWirelessControllerWtpProfileLanPort5Ssid2edl(o["port5-ssid"], d, "port5_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port5-ssid"], "ObjectWirelessControllerWtpProfileLan-Port5Ssid"); ok {
			if err = d.Set("port5_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port5_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port5_ssid: %v", err)
		}
	}

	if err = d.Set("port6_mode", flattenObjectWirelessControllerWtpProfileLanPort6Mode2edl(o["port6-mode"], d, "port6_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port6-mode"], "ObjectWirelessControllerWtpProfileLan-Port6Mode"); ok {
			if err = d.Set("port6_mode", vv); err != nil {
				return fmt.Errorf("Error reading port6_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port6_mode: %v", err)
		}
	}

	if err = d.Set("port6_ssid", flattenObjectWirelessControllerWtpProfileLanPort6Ssid2edl(o["port6-ssid"], d, "port6_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port6-ssid"], "ObjectWirelessControllerWtpProfileLan-Port6Ssid"); ok {
			if err = d.Set("port6_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port6_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port6_ssid: %v", err)
		}
	}

	if err = d.Set("port7_mode", flattenObjectWirelessControllerWtpProfileLanPort7Mode2edl(o["port7-mode"], d, "port7_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port7-mode"], "ObjectWirelessControllerWtpProfileLan-Port7Mode"); ok {
			if err = d.Set("port7_mode", vv); err != nil {
				return fmt.Errorf("Error reading port7_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port7_mode: %v", err)
		}
	}

	if err = d.Set("port7_ssid", flattenObjectWirelessControllerWtpProfileLanPort7Ssid2edl(o["port7-ssid"], d, "port7_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port7-ssid"], "ObjectWirelessControllerWtpProfileLan-Port7Ssid"); ok {
			if err = d.Set("port7_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port7_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port7_ssid: %v", err)
		}
	}

	if err = d.Set("port8_mode", flattenObjectWirelessControllerWtpProfileLanPort8Mode2edl(o["port8-mode"], d, "port8_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port8-mode"], "ObjectWirelessControllerWtpProfileLan-Port8Mode"); ok {
			if err = d.Set("port8_mode", vv); err != nil {
				return fmt.Errorf("Error reading port8_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port8_mode: %v", err)
		}
	}

	if err = d.Set("port8_ssid", flattenObjectWirelessControllerWtpProfileLanPort8Ssid2edl(o["port8-ssid"], d, "port8_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port8-ssid"], "ObjectWirelessControllerWtpProfileLan-Port8Ssid"); ok {
			if err = d.Set("port8_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port8_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port8_ssid: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerWtpProfileLanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerWtpProfileLanPortEslMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPortEslSsid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPortMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPortSsid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort1Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort1Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort2Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort2Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort3Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort3Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort4Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort4Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort5Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort5Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort6Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort6Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort7Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort7Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort8Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort8Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerWtpProfileLan(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("port_esl_mode"); ok || d.HasChange("port_esl_mode") {
		t, err := expandObjectWirelessControllerWtpProfileLanPortEslMode2edl(d, v, "port_esl_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-esl-mode"] = t
		}
	}

	if v, ok := d.GetOk("port_esl_ssid"); ok || d.HasChange("port_esl_ssid") {
		t, err := expandObjectWirelessControllerWtpProfileLanPortEslSsid2edl(d, v, "port_esl_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-esl-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port_mode"); ok || d.HasChange("port_mode") {
		t, err := expandObjectWirelessControllerWtpProfileLanPortMode2edl(d, v, "port_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-mode"] = t
		}
	}

	if v, ok := d.GetOk("port_ssid"); ok || d.HasChange("port_ssid") {
		t, err := expandObjectWirelessControllerWtpProfileLanPortSsid2edl(d, v, "port_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port1_mode"); ok || d.HasChange("port1_mode") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort1Mode2edl(d, v, "port1_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port1-mode"] = t
		}
	}

	if v, ok := d.GetOk("port1_ssid"); ok || d.HasChange("port1_ssid") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort1Ssid2edl(d, v, "port1_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port1-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port2_mode"); ok || d.HasChange("port2_mode") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort2Mode2edl(d, v, "port2_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port2-mode"] = t
		}
	}

	if v, ok := d.GetOk("port2_ssid"); ok || d.HasChange("port2_ssid") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort2Ssid2edl(d, v, "port2_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port2-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port3_mode"); ok || d.HasChange("port3_mode") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort3Mode2edl(d, v, "port3_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port3-mode"] = t
		}
	}

	if v, ok := d.GetOk("port3_ssid"); ok || d.HasChange("port3_ssid") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort3Ssid2edl(d, v, "port3_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port3-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port4_mode"); ok || d.HasChange("port4_mode") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort4Mode2edl(d, v, "port4_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port4-mode"] = t
		}
	}

	if v, ok := d.GetOk("port4_ssid"); ok || d.HasChange("port4_ssid") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort4Ssid2edl(d, v, "port4_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port4-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port5_mode"); ok || d.HasChange("port5_mode") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort5Mode2edl(d, v, "port5_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port5-mode"] = t
		}
	}

	if v, ok := d.GetOk("port5_ssid"); ok || d.HasChange("port5_ssid") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort5Ssid2edl(d, v, "port5_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port5-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port6_mode"); ok || d.HasChange("port6_mode") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort6Mode2edl(d, v, "port6_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port6-mode"] = t
		}
	}

	if v, ok := d.GetOk("port6_ssid"); ok || d.HasChange("port6_ssid") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort6Ssid2edl(d, v, "port6_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port6-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port7_mode"); ok || d.HasChange("port7_mode") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort7Mode2edl(d, v, "port7_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port7-mode"] = t
		}
	}

	if v, ok := d.GetOk("port7_ssid"); ok || d.HasChange("port7_ssid") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort7Ssid2edl(d, v, "port7_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port7-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port8_mode"); ok || d.HasChange("port8_mode") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort8Mode2edl(d, v, "port8_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port8-mode"] = t
		}
	}

	if v, ok := d.GetOk("port8_ssid"); ok || d.HasChange("port8_ssid") {
		t, err := expandObjectWirelessControllerWtpProfileLanPort8Ssid2edl(d, v, "port8_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port8-ssid"] = t
		}
	}

	return &obj, nil
}
