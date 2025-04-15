// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Server SSH public key authentication settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallAccessProxyServerPubkeyAuthSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallAccessProxyServerPubkeyAuthSettingsUpdate,
		Read:   resourceObjectFirewallAccessProxyServerPubkeyAuthSettingsRead,
		Update: resourceObjectFirewallAccessProxyServerPubkeyAuthSettingsUpdate,
		Delete: resourceObjectFirewallAccessProxyServerPubkeyAuthSettingsDelete,

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
			"access_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"auth_ca": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cert_extension": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"critical": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"data": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"permit_agent_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_port_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_pty": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_user_rc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_x11_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectFirewallAccessProxyServerPubkeyAuthSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	access_proxy := d.Get("access_proxy").(string)
	paradict["access_proxy"] = access_proxy

	obj, err := getObjectObjectFirewallAccessProxyServerPubkeyAuthSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAccessProxyServerPubkeyAuthSettings resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallAccessProxyServerPubkeyAuthSettings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAccessProxyServerPubkeyAuthSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFirewallAccessProxyServerPubkeyAuthSettings")

	return resourceObjectFirewallAccessProxyServerPubkeyAuthSettingsRead(d, m)
}

func resourceObjectFirewallAccessProxyServerPubkeyAuthSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	access_proxy := d.Get("access_proxy").(string)
	paradict["access_proxy"] = access_proxy

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallAccessProxyServerPubkeyAuthSettings(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallAccessProxyServerPubkeyAuthSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallAccessProxyServerPubkeyAuthSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	access_proxy := d.Get("access_proxy").(string)
	if access_proxy == "" {
		access_proxy = importOptionChecking(m.(*FortiClient).Cfg, "access_proxy")
		if access_proxy == "" {
			return fmt.Errorf("Parameter access_proxy is missing")
		}
		if err = d.Set("access_proxy", access_proxy); err != nil {
			return fmt.Errorf("Error set params access_proxy: %v", err)
		}
	}
	paradict["access_proxy"] = access_proxy

	o, err := c.ReadObjectFirewallAccessProxyServerPubkeyAuthSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAccessProxyServerPubkeyAuthSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallAccessProxyServerPubkeyAuthSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAccessProxyServerPubkeyAuthSettings resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsAuthCa2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtension2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "critical"
		if _, ok := i["critical"]; ok {
			v := flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionCritical2edl(i["critical"], d, pre_append)
			tmp["critical"] = fortiAPISubPartPatch(v, "ObjectFirewallAccessProxyServerPubkeyAuthSettings-CertExtension-Critical")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := i["data"]; ok {
			v := flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionData2edl(i["data"], d, pre_append)
			tmp["data"] = fortiAPISubPartPatch(v, "ObjectFirewallAccessProxyServerPubkeyAuthSettings-CertExtension-Data")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallAccessProxyServerPubkeyAuthSettings-CertExtension-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFirewallAccessProxyServerPubkeyAuthSettings-CertExtension-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionCritical2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionData2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitAgentForwarding2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitPortForwarding2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitPty2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitUserRc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitX11Forwarding2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsSourceAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallAccessProxyServerPubkeyAuthSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("auth_ca", flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsAuthCa2edl(o["auth-ca"], d, "auth_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-ca"], "ObjectFirewallAccessProxyServerPubkeyAuthSettings-AuthCa"); ok {
			if err = d.Set("auth_ca", vv); err != nil {
				return fmt.Errorf("Error reading auth_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_ca: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("cert_extension", flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtension2edl(o["cert-extension"], d, "cert_extension")); err != nil {
			if vv, ok := fortiAPIPatch(o["cert-extension"], "ObjectFirewallAccessProxyServerPubkeyAuthSettings-CertExtension"); ok {
				if err = d.Set("cert_extension", vv); err != nil {
					return fmt.Errorf("Error reading cert_extension: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading cert_extension: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cert_extension"); ok {
			if err = d.Set("cert_extension", flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtension2edl(o["cert-extension"], d, "cert_extension")); err != nil {
				if vv, ok := fortiAPIPatch(o["cert-extension"], "ObjectFirewallAccessProxyServerPubkeyAuthSettings-CertExtension"); ok {
					if err = d.Set("cert_extension", vv); err != nil {
						return fmt.Errorf("Error reading cert_extension: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading cert_extension: %v", err)
				}
			}
		}
	}

	if err = d.Set("permit_agent_forwarding", flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitAgentForwarding2edl(o["permit-agent-forwarding"], d, "permit_agent_forwarding")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-agent-forwarding"], "ObjectFirewallAccessProxyServerPubkeyAuthSettings-PermitAgentForwarding"); ok {
			if err = d.Set("permit_agent_forwarding", vv); err != nil {
				return fmt.Errorf("Error reading permit_agent_forwarding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_agent_forwarding: %v", err)
		}
	}

	if err = d.Set("permit_port_forwarding", flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitPortForwarding2edl(o["permit-port-forwarding"], d, "permit_port_forwarding")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-port-forwarding"], "ObjectFirewallAccessProxyServerPubkeyAuthSettings-PermitPortForwarding"); ok {
			if err = d.Set("permit_port_forwarding", vv); err != nil {
				return fmt.Errorf("Error reading permit_port_forwarding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_port_forwarding: %v", err)
		}
	}

	if err = d.Set("permit_pty", flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitPty2edl(o["permit-pty"], d, "permit_pty")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-pty"], "ObjectFirewallAccessProxyServerPubkeyAuthSettings-PermitPty"); ok {
			if err = d.Set("permit_pty", vv); err != nil {
				return fmt.Errorf("Error reading permit_pty: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_pty: %v", err)
		}
	}

	if err = d.Set("permit_user_rc", flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitUserRc2edl(o["permit-user-rc"], d, "permit_user_rc")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-user-rc"], "ObjectFirewallAccessProxyServerPubkeyAuthSettings-PermitUserRc"); ok {
			if err = d.Set("permit_user_rc", vv); err != nil {
				return fmt.Errorf("Error reading permit_user_rc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_user_rc: %v", err)
		}
	}

	if err = d.Set("permit_x11_forwarding", flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitX11Forwarding2edl(o["permit-x11-forwarding"], d, "permit_x11_forwarding")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-x11-forwarding"], "ObjectFirewallAccessProxyServerPubkeyAuthSettings-PermitX11Forwarding"); ok {
			if err = d.Set("permit_x11_forwarding", vv); err != nil {
				return fmt.Errorf("Error reading permit_x11_forwarding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_x11_forwarding: %v", err)
		}
	}

	if err = d.Set("source_address", flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsSourceAddress2edl(o["source-address"], d, "source_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-address"], "ObjectFirewallAccessProxyServerPubkeyAuthSettings-SourceAddress"); ok {
			if err = d.Set("source_address", vv); err != nil {
				return fmt.Errorf("Error reading source_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_address: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallAccessProxyServerPubkeyAuthSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallAccessProxyServerPubkeyAuthSettingsAuthCa2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtension2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "critical"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["critical"], _ = expandObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionCritical2edl(d, i["critical"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["data"], _ = expandObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionData2edl(d, i["data"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionType2edl(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionCritical2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionData2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtensionType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitAgentForwarding2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitPortForwarding2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitPty2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitUserRc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitX11Forwarding2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyServerPubkeyAuthSettingsSourceAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallAccessProxyServerPubkeyAuthSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_ca"); ok || d.HasChange("auth_ca") {
		t, err := expandObjectFirewallAccessProxyServerPubkeyAuthSettingsAuthCa2edl(d, v, "auth_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-ca"] = t
		}
	}

	if v, ok := d.GetOk("cert_extension"); ok || d.HasChange("cert_extension") {
		t, err := expandObjectFirewallAccessProxyServerPubkeyAuthSettingsCertExtension2edl(d, v, "cert_extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-extension"] = t
		}
	}

	if v, ok := d.GetOk("permit_agent_forwarding"); ok || d.HasChange("permit_agent_forwarding") {
		t, err := expandObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitAgentForwarding2edl(d, v, "permit_agent_forwarding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-agent-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("permit_port_forwarding"); ok || d.HasChange("permit_port_forwarding") {
		t, err := expandObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitPortForwarding2edl(d, v, "permit_port_forwarding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-port-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("permit_pty"); ok || d.HasChange("permit_pty") {
		t, err := expandObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitPty2edl(d, v, "permit_pty")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-pty"] = t
		}
	}

	if v, ok := d.GetOk("permit_user_rc"); ok || d.HasChange("permit_user_rc") {
		t, err := expandObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitUserRc2edl(d, v, "permit_user_rc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-user-rc"] = t
		}
	}

	if v, ok := d.GetOk("permit_x11_forwarding"); ok || d.HasChange("permit_x11_forwarding") {
		t, err := expandObjectFirewallAccessProxyServerPubkeyAuthSettingsPermitX11Forwarding2edl(d, v, "permit_x11_forwarding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-x11-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("source_address"); ok || d.HasChange("source_address") {
		t, err := expandObjectFirewallAccessProxyServerPubkeyAuthSettingsSourceAddress2edl(d, v, "source_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address"] = t
		}
	}

	return &obj, nil
}
