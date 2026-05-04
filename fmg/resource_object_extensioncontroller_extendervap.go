// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiExtender wifi vap configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtensionControllerExtenderVap() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtensionControllerExtenderVapCreate,
		Read:   resourceObjectExtensionControllerExtenderVapRead,
		Update: resourceObjectExtensionControllerExtenderVapUpdate,
		Delete: resourceObjectExtensionControllerExtenderVapDelete,

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
			"allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auth_server_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auth_server_secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"broadcast_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bss_color_partial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dtim": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"max_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mu_mimo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"passphrase": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"pmf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rts_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sae_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"security": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_wake_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectExtensionControllerExtenderVapCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectExtensionControllerExtenderVap(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtensionControllerExtenderVap resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectExtensionControllerExtenderVap(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectExtensionControllerExtenderVap(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectExtensionControllerExtenderVap resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectExtensionControllerExtenderVap(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectExtensionControllerExtenderVap resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtensionControllerExtenderVapRead(d, m)
}

func resourceObjectExtensionControllerExtenderVapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectExtensionControllerExtenderVap(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderVap resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectExtensionControllerExtenderVap(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderVap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtensionControllerExtenderVapRead(d, m)
}

func resourceObjectExtensionControllerExtenderVapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectExtensionControllerExtenderVap(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtensionControllerExtenderVap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtensionControllerExtenderVapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectExtensionControllerExtenderVap(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderVap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtensionControllerExtenderVap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderVap resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtensionControllerExtenderVapAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderVapAuthServerAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapAuthServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapAuthServerSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapBroadcastSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapBssColorPartial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapDtim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderVapMaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapMuMimo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapPmf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapRtsThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapTargetWakeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderVapType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtensionControllerExtenderVap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("allowaccess", flattenObjectExtensionControllerExtenderVapAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowaccess"], "ObjectExtensionControllerExtenderVap-Allowaccess"); ok {
			if err = d.Set("allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("auth_server_address", flattenObjectExtensionControllerExtenderVapAuthServerAddress(o["auth-server-address"], d, "auth_server_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-server-address"], "ObjectExtensionControllerExtenderVap-AuthServerAddress"); ok {
			if err = d.Set("auth_server_address", vv); err != nil {
				return fmt.Errorf("Error reading auth_server_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_server_address: %v", err)
		}
	}

	if err = d.Set("auth_server_port", flattenObjectExtensionControllerExtenderVapAuthServerPort(o["auth-server-port"], d, "auth_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-server-port"], "ObjectExtensionControllerExtenderVap-AuthServerPort"); ok {
			if err = d.Set("auth_server_port", vv); err != nil {
				return fmt.Errorf("Error reading auth_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_server_port: %v", err)
		}
	}

	if err = d.Set("auth_server_secret", flattenObjectExtensionControllerExtenderVapAuthServerSecret(o["auth-server-secret"], d, "auth_server_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-server-secret"], "ObjectExtensionControllerExtenderVap-AuthServerSecret"); ok {
			if err = d.Set("auth_server_secret", vv); err != nil {
				return fmt.Errorf("Error reading auth_server_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_server_secret: %v", err)
		}
	}

	if err = d.Set("broadcast_ssid", flattenObjectExtensionControllerExtenderVapBroadcastSsid(o["broadcast-ssid"], d, "broadcast_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["broadcast-ssid"], "ObjectExtensionControllerExtenderVap-BroadcastSsid"); ok {
			if err = d.Set("broadcast_ssid", vv); err != nil {
				return fmt.Errorf("Error reading broadcast_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading broadcast_ssid: %v", err)
		}
	}

	if err = d.Set("bss_color_partial", flattenObjectExtensionControllerExtenderVapBssColorPartial(o["bss-color-partial"], d, "bss_color_partial")); err != nil {
		if vv, ok := fortiAPIPatch(o["bss-color-partial"], "ObjectExtensionControllerExtenderVap-BssColorPartial"); ok {
			if err = d.Set("bss_color_partial", vv); err != nil {
				return fmt.Errorf("Error reading bss_color_partial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bss_color_partial: %v", err)
		}
	}

	if err = d.Set("dtim", flattenObjectExtensionControllerExtenderVapDtim(o["dtim"], d, "dtim")); err != nil {
		if vv, ok := fortiAPIPatch(o["dtim"], "ObjectExtensionControllerExtenderVap-Dtim"); ok {
			if err = d.Set("dtim", vv); err != nil {
				return fmt.Errorf("Error reading dtim: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dtim: %v", err)
		}
	}

	if err = d.Set("end_ip", flattenObjectExtensionControllerExtenderVapEndIp(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "ObjectExtensionControllerExtenderVap-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("ip_address", flattenObjectExtensionControllerExtenderVapIpAddress(o["ip-address"], d, "ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-address"], "ObjectExtensionControllerExtenderVap-IpAddress"); ok {
			if err = d.Set("ip_address", vv); err != nil {
				return fmt.Errorf("Error reading ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_address: %v", err)
		}
	}

	if err = d.Set("max_clients", flattenObjectExtensionControllerExtenderVapMaxClients(o["max-clients"], d, "max_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-clients"], "ObjectExtensionControllerExtenderVap-MaxClients"); ok {
			if err = d.Set("max_clients", vv); err != nil {
				return fmt.Errorf("Error reading max_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_clients: %v", err)
		}
	}

	if err = d.Set("mu_mimo", flattenObjectExtensionControllerExtenderVapMuMimo(o["mu-mimo"], d, "mu_mimo")); err != nil {
		if vv, ok := fortiAPIPatch(o["mu-mimo"], "ObjectExtensionControllerExtenderVap-MuMimo"); ok {
			if err = d.Set("mu_mimo", vv); err != nil {
				return fmt.Errorf("Error reading mu_mimo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mu_mimo: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectExtensionControllerExtenderVapName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectExtensionControllerExtenderVap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pmf", flattenObjectExtensionControllerExtenderVapPmf(o["pmf"], d, "pmf")); err != nil {
		if vv, ok := fortiAPIPatch(o["pmf"], "ObjectExtensionControllerExtenderVap-Pmf"); ok {
			if err = d.Set("pmf", vv); err != nil {
				return fmt.Errorf("Error reading pmf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pmf: %v", err)
		}
	}

	if err = d.Set("rts_threshold", flattenObjectExtensionControllerExtenderVapRtsThreshold(o["rts-threshold"], d, "rts_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["rts-threshold"], "ObjectExtensionControllerExtenderVap-RtsThreshold"); ok {
			if err = d.Set("rts_threshold", vv); err != nil {
				return fmt.Errorf("Error reading rts_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rts_threshold: %v", err)
		}
	}

	if err = d.Set("security", flattenObjectExtensionControllerExtenderVapSecurity(o["security"], d, "security")); err != nil {
		if vv, ok := fortiAPIPatch(o["security"], "ObjectExtensionControllerExtenderVap-Security"); ok {
			if err = d.Set("security", vv); err != nil {
				return fmt.Errorf("Error reading security: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("ssid", flattenObjectExtensionControllerExtenderVapSsid(o["ssid"], d, "ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssid"], "ObjectExtensionControllerExtenderVap-Ssid"); ok {
			if err = d.Set("ssid", vv); err != nil {
				return fmt.Errorf("Error reading ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssid: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenObjectExtensionControllerExtenderVapStartIp(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "ObjectExtensionControllerExtenderVap-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	if err = d.Set("target_wake_time", flattenObjectExtensionControllerExtenderVapTargetWakeTime(o["target-wake-time"], d, "target_wake_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["target-wake-time"], "ObjectExtensionControllerExtenderVap-TargetWakeTime"); ok {
			if err = d.Set("target_wake_time", vv); err != nil {
				return fmt.Errorf("Error reading target_wake_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target_wake_time: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectExtensionControllerExtenderVapType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectExtensionControllerExtenderVap-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectExtensionControllerExtenderVapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtensionControllerExtenderVapAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderVapAuthServerAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapAuthServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapAuthServerSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapBroadcastSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapBssColorPartial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapDtim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectExtensionControllerExtenderVapMaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapMuMimo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapPassphrase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderVapPmf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapRtsThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapSaePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderVapSecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapTargetWakeTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderVapType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtensionControllerExtenderVap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allowaccess"); ok || d.HasChange("allowaccess") {
		t, err := expandObjectExtensionControllerExtenderVapAllowaccess(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("auth_server_address"); ok || d.HasChange("auth_server_address") {
		t, err := expandObjectExtensionControllerExtenderVapAuthServerAddress(d, v, "auth_server_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-server-address"] = t
		}
	}

	if v, ok := d.GetOk("auth_server_port"); ok || d.HasChange("auth_server_port") {
		t, err := expandObjectExtensionControllerExtenderVapAuthServerPort(d, v, "auth_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-server-port"] = t
		}
	}

	if v, ok := d.GetOk("auth_server_secret"); ok || d.HasChange("auth_server_secret") {
		t, err := expandObjectExtensionControllerExtenderVapAuthServerSecret(d, v, "auth_server_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-server-secret"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_ssid"); ok || d.HasChange("broadcast_ssid") {
		t, err := expandObjectExtensionControllerExtenderVapBroadcastSsid(d, v, "broadcast_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-ssid"] = t
		}
	}

	if v, ok := d.GetOk("bss_color_partial"); ok || d.HasChange("bss_color_partial") {
		t, err := expandObjectExtensionControllerExtenderVapBssColorPartial(d, v, "bss_color_partial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bss-color-partial"] = t
		}
	}

	if v, ok := d.GetOk("dtim"); ok || d.HasChange("dtim") {
		t, err := expandObjectExtensionControllerExtenderVapDtim(d, v, "dtim")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtim"] = t
		}
	}

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandObjectExtensionControllerExtenderVapEndIp(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("ip_address"); ok || d.HasChange("ip_address") {
		t, err := expandObjectExtensionControllerExtenderVapIpAddress(d, v, "ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-address"] = t
		}
	}

	if v, ok := d.GetOk("max_clients"); ok || d.HasChange("max_clients") {
		t, err := expandObjectExtensionControllerExtenderVapMaxClients(d, v, "max_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients"] = t
		}
	}

	if v, ok := d.GetOk("mu_mimo"); ok || d.HasChange("mu_mimo") {
		t, err := expandObjectExtensionControllerExtenderVapMuMimo(d, v, "mu_mimo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mu-mimo"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectExtensionControllerExtenderVapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("passphrase"); ok || d.HasChange("passphrase") {
		t, err := expandObjectExtensionControllerExtenderVapPassphrase(d, v, "passphrase")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passphrase"] = t
		}
	}

	if v, ok := d.GetOk("pmf"); ok || d.HasChange("pmf") {
		t, err := expandObjectExtensionControllerExtenderVapPmf(d, v, "pmf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf"] = t
		}
	}

	if v, ok := d.GetOk("rts_threshold"); ok || d.HasChange("rts_threshold") {
		t, err := expandObjectExtensionControllerExtenderVapRtsThreshold(d, v, "rts_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rts-threshold"] = t
		}
	}

	if v, ok := d.GetOk("sae_password"); ok || d.HasChange("sae_password") {
		t, err := expandObjectExtensionControllerExtenderVapSaePassword(d, v, "sae_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-password"] = t
		}
	}

	if v, ok := d.GetOk("security"); ok || d.HasChange("security") {
		t, err := expandObjectExtensionControllerExtenderVapSecurity(d, v, "security")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security"] = t
		}
	}

	if v, ok := d.GetOk("ssid"); ok || d.HasChange("ssid") {
		t, err := expandObjectExtensionControllerExtenderVapSsid(d, v, "ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssid"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandObjectExtensionControllerExtenderVapStartIp(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	if v, ok := d.GetOk("target_wake_time"); ok || d.HasChange("target_wake_time") {
		t, err := expandObjectExtensionControllerExtenderVapTargetWakeTime(d, v, "target_wake_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-wake-time"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectExtensionControllerExtenderVapType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
