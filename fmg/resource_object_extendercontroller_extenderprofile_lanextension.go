// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiExtender lan extension configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtenderControllerExtenderProfileLanExtension() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtenderControllerExtenderProfileLanExtensionUpdate,
		Read:   resourceObjectExtenderControllerExtenderProfileLanExtensionRead,
		Update: resourceObjectExtenderControllerExtenderProfileLanExtensionUpdate,
		Delete: resourceObjectExtenderControllerExtenderProfileLanExtensionDelete,

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
			"extender_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"backhaul": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"backhaul_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"backhaul_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipsec_tunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"link_loadbalance": &schema.Schema{
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

func resourceObjectExtenderControllerExtenderProfileLanExtensionUpdate(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	paradict["extender_profile"] = extender_profile

	obj, err := getObjectObjectExtenderControllerExtenderProfileLanExtension(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerExtenderProfileLanExtension resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectExtenderControllerExtenderProfileLanExtension(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerExtenderProfileLanExtension resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectExtenderControllerExtenderProfileLanExtension")

	return resourceObjectExtenderControllerExtenderProfileLanExtensionRead(d, m)
}

func resourceObjectExtenderControllerExtenderProfileLanExtensionDelete(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	paradict["extender_profile"] = extender_profile

	err = c.DeleteObjectExtenderControllerExtenderProfileLanExtension(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtenderControllerExtenderProfileLanExtension resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtenderControllerExtenderProfileLanExtensionRead(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	if extender_profile == "" {
		extender_profile = importOptionChecking(m.(*FortiClient).Cfg, "extender_profile")
		if extender_profile == "" {
			return fmt.Errorf("Parameter extender_profile is missing")
		}
		if err = d.Set("extender_profile", extender_profile); err != nil {
			return fmt.Errorf("Error set params extender_profile: %v", err)
		}
	}
	paradict["extender_profile"] = extender_profile

	o, err := c.ReadObjectExtenderControllerExtenderProfileLanExtension(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerExtenderProfileLanExtension resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtenderControllerExtenderProfileLanExtension(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerExtenderProfileLanExtension resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaul2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectExtenderControllerExtenderProfileLanExtension-Backhaul-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulPort2edl(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectExtenderControllerExtenderProfileLanExtension-Backhaul-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulRole2edl(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "ObjectExtenderControllerExtenderProfileLanExtension-Backhaul-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulWeight2edl(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectExtenderControllerExtenderProfileLanExtension-Backhaul-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulRole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionIpsecTunnel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionLinkLoadbalance2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtenderControllerExtenderProfileLanExtension(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("backhaul", flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaul2edl(o["backhaul"], d, "backhaul")); err != nil {
			if vv, ok := fortiAPIPatch(o["backhaul"], "ObjectExtenderControllerExtenderProfileLanExtension-Backhaul"); ok {
				if err = d.Set("backhaul", vv); err != nil {
					return fmt.Errorf("Error reading backhaul: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading backhaul: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("backhaul"); ok {
			if err = d.Set("backhaul", flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaul2edl(o["backhaul"], d, "backhaul")); err != nil {
				if vv, ok := fortiAPIPatch(o["backhaul"], "ObjectExtenderControllerExtenderProfileLanExtension-Backhaul"); ok {
					if err = d.Set("backhaul", vv); err != nil {
						return fmt.Errorf("Error reading backhaul: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading backhaul: %v", err)
				}
			}
		}
	}

	if err = d.Set("backhaul_interface", flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulInterface2edl(o["backhaul-interface"], d, "backhaul_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["backhaul-interface"], "ObjectExtenderControllerExtenderProfileLanExtension-BackhaulInterface"); ok {
			if err = d.Set("backhaul_interface", vv); err != nil {
				return fmt.Errorf("Error reading backhaul_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading backhaul_interface: %v", err)
		}
	}

	if err = d.Set("backhaul_ip", flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulIp2edl(o["backhaul-ip"], d, "backhaul_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["backhaul-ip"], "ObjectExtenderControllerExtenderProfileLanExtension-BackhaulIp"); ok {
			if err = d.Set("backhaul_ip", vv); err != nil {
				return fmt.Errorf("Error reading backhaul_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading backhaul_ip: %v", err)
		}
	}

	if err = d.Set("ipsec_tunnel", flattenObjectExtenderControllerExtenderProfileLanExtensionIpsecTunnel2edl(o["ipsec-tunnel"], d, "ipsec_tunnel")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-tunnel"], "ObjectExtenderControllerExtenderProfileLanExtension-IpsecTunnel"); ok {
			if err = d.Set("ipsec_tunnel", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_tunnel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_tunnel: %v", err)
		}
	}

	if err = d.Set("link_loadbalance", flattenObjectExtenderControllerExtenderProfileLanExtensionLinkLoadbalance2edl(o["link-loadbalance"], d, "link_loadbalance")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-loadbalance"], "ObjectExtenderControllerExtenderProfileLanExtension-LinkLoadbalance"); ok {
			if err = d.Set("link_loadbalance", vv); err != nil {
				return fmt.Errorf("Error reading link_loadbalance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_loadbalance: %v", err)
		}
	}

	return nil
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaul2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulPort2edl(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulRole2edl(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulWeight2edl(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulRole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionIpsecTunnel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionLinkLoadbalance2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtenderControllerExtenderProfileLanExtension(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("backhaul"); ok || d.HasChange("backhaul") {
		t, err := expandObjectExtenderControllerExtenderProfileLanExtensionBackhaul2edl(d, v, "backhaul")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backhaul"] = t
		}
	}

	if v, ok := d.GetOk("backhaul_interface"); ok || d.HasChange("backhaul_interface") {
		t, err := expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulInterface2edl(d, v, "backhaul_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backhaul-interface"] = t
		}
	}

	if v, ok := d.GetOk("backhaul_ip"); ok || d.HasChange("backhaul_ip") {
		t, err := expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulIp2edl(d, v, "backhaul_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backhaul-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_tunnel"); ok || d.HasChange("ipsec_tunnel") {
		t, err := expandObjectExtenderControllerExtenderProfileLanExtensionIpsecTunnel2edl(d, v, "ipsec_tunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-tunnel"] = t
		}
	}

	if v, ok := d.GetOk("link_loadbalance"); ok || d.HasChange("link_loadbalance") {
		t, err := expandObjectExtenderControllerExtenderProfileLanExtensionLinkLoadbalance2edl(d, v, "link_loadbalance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-loadbalance"] = t
		}
	}

	return &obj, nil
}
