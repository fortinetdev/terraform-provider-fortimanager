// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Access Proxy SSH client certificate.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallAccessProxySshClientCert() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallAccessProxySshClientCertCreate,
		Read:   resourceObjectFirewallAccessProxySshClientCertRead,
		Update: resourceObjectFirewallAccessProxySshClientCertUpdate,
		Delete: resourceObjectFirewallAccessProxySshClientCertDelete,

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
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"permit_agent_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"permit_port_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"permit_pty": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"permit_user_rc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"permit_x11_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_address": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceObjectFirewallAccessProxySshClientCertCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectFirewallAccessProxySshClientCert(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAccessProxySshClientCert resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallAccessProxySshClientCert(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAccessProxySshClientCert resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallAccessProxySshClientCertRead(d, m)
}

func resourceObjectFirewallAccessProxySshClientCertUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallAccessProxySshClientCert(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAccessProxySshClientCert resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallAccessProxySshClientCert(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAccessProxySshClientCert resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallAccessProxySshClientCertRead(d, m)
}

func resourceObjectFirewallAccessProxySshClientCertDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallAccessProxySshClientCert(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallAccessProxySshClientCert resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallAccessProxySshClientCertRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallAccessProxySshClientCert(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAccessProxySshClientCert resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallAccessProxySshClientCert(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAccessProxySshClientCert resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallAccessProxySshClientCertAuthCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxySshClientCertCertExtension(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallAccessProxySshClientCertCertExtensionCritical(i["critical"], d, pre_append)
			tmp["critical"] = fortiAPISubPartPatch(v, "ObjectFirewallAccessProxySshClientCert-CertExtension-Critical")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := i["data"]; ok {
			v := flattenObjectFirewallAccessProxySshClientCertCertExtensionData(i["data"], d, pre_append)
			tmp["data"] = fortiAPISubPartPatch(v, "ObjectFirewallAccessProxySshClientCert-CertExtension-Data")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectFirewallAccessProxySshClientCertCertExtensionName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallAccessProxySshClientCert-CertExtension-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFirewallAccessProxySshClientCertCertExtensionType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFirewallAccessProxySshClientCert-CertExtension-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallAccessProxySshClientCertCertExtensionCritical(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxySshClientCertCertExtensionData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxySshClientCertCertExtensionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxySshClientCertCertExtensionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxySshClientCertName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxySshClientCertPermitAgentForwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxySshClientCertPermitPortForwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxySshClientCertPermitPty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxySshClientCertPermitUserRc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxySshClientCertPermitX11Forwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxySshClientCertSourceAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallAccessProxySshClientCert(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("auth_ca", flattenObjectFirewallAccessProxySshClientCertAuthCa(o["auth-ca"], d, "auth_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-ca"], "ObjectFirewallAccessProxySshClientCert-AuthCa"); ok {
			if err = d.Set("auth_ca", vv); err != nil {
				return fmt.Errorf("Error reading auth_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_ca: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("cert_extension", flattenObjectFirewallAccessProxySshClientCertCertExtension(o["cert-extension"], d, "cert_extension")); err != nil {
			if vv, ok := fortiAPIPatch(o["cert-extension"], "ObjectFirewallAccessProxySshClientCert-CertExtension"); ok {
				if err = d.Set("cert_extension", vv); err != nil {
					return fmt.Errorf("Error reading cert_extension: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading cert_extension: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cert_extension"); ok {
			if err = d.Set("cert_extension", flattenObjectFirewallAccessProxySshClientCertCertExtension(o["cert-extension"], d, "cert_extension")); err != nil {
				if vv, ok := fortiAPIPatch(o["cert-extension"], "ObjectFirewallAccessProxySshClientCert-CertExtension"); ok {
					if err = d.Set("cert_extension", vv); err != nil {
						return fmt.Errorf("Error reading cert_extension: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading cert_extension: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectFirewallAccessProxySshClientCertName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallAccessProxySshClientCert-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("permit_agent_forwarding", flattenObjectFirewallAccessProxySshClientCertPermitAgentForwarding(o["permit-agent-forwarding"], d, "permit_agent_forwarding")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-agent-forwarding"], "ObjectFirewallAccessProxySshClientCert-PermitAgentForwarding"); ok {
			if err = d.Set("permit_agent_forwarding", vv); err != nil {
				return fmt.Errorf("Error reading permit_agent_forwarding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_agent_forwarding: %v", err)
		}
	}

	if err = d.Set("permit_port_forwarding", flattenObjectFirewallAccessProxySshClientCertPermitPortForwarding(o["permit-port-forwarding"], d, "permit_port_forwarding")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-port-forwarding"], "ObjectFirewallAccessProxySshClientCert-PermitPortForwarding"); ok {
			if err = d.Set("permit_port_forwarding", vv); err != nil {
				return fmt.Errorf("Error reading permit_port_forwarding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_port_forwarding: %v", err)
		}
	}

	if err = d.Set("permit_pty", flattenObjectFirewallAccessProxySshClientCertPermitPty(o["permit-pty"], d, "permit_pty")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-pty"], "ObjectFirewallAccessProxySshClientCert-PermitPty"); ok {
			if err = d.Set("permit_pty", vv); err != nil {
				return fmt.Errorf("Error reading permit_pty: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_pty: %v", err)
		}
	}

	if err = d.Set("permit_user_rc", flattenObjectFirewallAccessProxySshClientCertPermitUserRc(o["permit-user-rc"], d, "permit_user_rc")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-user-rc"], "ObjectFirewallAccessProxySshClientCert-PermitUserRc"); ok {
			if err = d.Set("permit_user_rc", vv); err != nil {
				return fmt.Errorf("Error reading permit_user_rc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_user_rc: %v", err)
		}
	}

	if err = d.Set("permit_x11_forwarding", flattenObjectFirewallAccessProxySshClientCertPermitX11Forwarding(o["permit-x11-forwarding"], d, "permit_x11_forwarding")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-x11-forwarding"], "ObjectFirewallAccessProxySshClientCert-PermitX11Forwarding"); ok {
			if err = d.Set("permit_x11_forwarding", vv); err != nil {
				return fmt.Errorf("Error reading permit_x11_forwarding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_x11_forwarding: %v", err)
		}
	}

	if err = d.Set("source_address", flattenObjectFirewallAccessProxySshClientCertSourceAddress(o["source-address"], d, "source_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-address"], "ObjectFirewallAccessProxySshClientCert-SourceAddress"); ok {
			if err = d.Set("source_address", vv); err != nil {
				return fmt.Errorf("Error reading source_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_address: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallAccessProxySshClientCertFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallAccessProxySshClientCertAuthCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxySshClientCertCertExtension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["critical"], _ = expandObjectFirewallAccessProxySshClientCertCertExtensionCritical(d, i["critical"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["data"], _ = expandObjectFirewallAccessProxySshClientCertCertExtensionData(d, i["data"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectFirewallAccessProxySshClientCertCertExtensionName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFirewallAccessProxySshClientCertCertExtensionType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAccessProxySshClientCertCertExtensionCritical(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxySshClientCertCertExtensionData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxySshClientCertCertExtensionName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxySshClientCertCertExtensionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxySshClientCertName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxySshClientCertPermitAgentForwarding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxySshClientCertPermitPortForwarding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxySshClientCertPermitPty(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxySshClientCertPermitUserRc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxySshClientCertPermitX11Forwarding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxySshClientCertSourceAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallAccessProxySshClientCert(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_ca"); ok || d.HasChange("auth_ca") {
		t, err := expandObjectFirewallAccessProxySshClientCertAuthCa(d, v, "auth_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-ca"] = t
		}
	}

	if v, ok := d.GetOk("cert_extension"); ok || d.HasChange("cert_extension") {
		t, err := expandObjectFirewallAccessProxySshClientCertCertExtension(d, v, "cert_extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-extension"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallAccessProxySshClientCertName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("permit_agent_forwarding"); ok || d.HasChange("permit_agent_forwarding") {
		t, err := expandObjectFirewallAccessProxySshClientCertPermitAgentForwarding(d, v, "permit_agent_forwarding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-agent-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("permit_port_forwarding"); ok || d.HasChange("permit_port_forwarding") {
		t, err := expandObjectFirewallAccessProxySshClientCertPermitPortForwarding(d, v, "permit_port_forwarding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-port-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("permit_pty"); ok || d.HasChange("permit_pty") {
		t, err := expandObjectFirewallAccessProxySshClientCertPermitPty(d, v, "permit_pty")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-pty"] = t
		}
	}

	if v, ok := d.GetOk("permit_user_rc"); ok || d.HasChange("permit_user_rc") {
		t, err := expandObjectFirewallAccessProxySshClientCertPermitUserRc(d, v, "permit_user_rc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-user-rc"] = t
		}
	}

	if v, ok := d.GetOk("permit_x11_forwarding"); ok || d.HasChange("permit_x11_forwarding") {
		t, err := expandObjectFirewallAccessProxySshClientCertPermitX11Forwarding(d, v, "permit_x11_forwarding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-x11-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("source_address"); ok || d.HasChange("source_address") {
		t, err := expandObjectFirewallAccessProxySshClientCertSourceAddress(d, v, "source_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address"] = t
		}
	}

	return &obj, nil
}
