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

func resourceObjectExtensionControllerExtenderProfileLanExtension() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtensionControllerExtenderProfileLanExtensionUpdate,
		Read:   resourceObjectExtensionControllerExtenderProfileLanExtensionRead,
		Update: resourceObjectExtensionControllerExtenderProfileLanExtensionUpdate,
		Delete: resourceObjectExtensionControllerExtenderProfileLanExtensionDelete,

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
			"downlinks": &schema.Schema{
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
						},
						"pvid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
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

func resourceObjectExtensionControllerExtenderProfileLanExtensionUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectExtensionControllerExtenderProfileLanExtension(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileLanExtension resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectExtensionControllerExtenderProfileLanExtension(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileLanExtension resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectExtensionControllerExtenderProfileLanExtension")

	return resourceObjectExtensionControllerExtenderProfileLanExtensionRead(d, m)
}

func resourceObjectExtensionControllerExtenderProfileLanExtensionDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectExtensionControllerExtenderProfileLanExtension(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtensionControllerExtenderProfileLanExtension resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtensionControllerExtenderProfileLanExtensionRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectExtensionControllerExtenderProfileLanExtension(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileLanExtension resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtensionControllerExtenderProfileLanExtension(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileLanExtension resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaul2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileLanExtension-Backhaul-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulPort2edl(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileLanExtension-Backhaul-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulRole2edl(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileLanExtension-Backhaul-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulWeight2edl(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileLanExtension-Backhaul-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulRole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionDownlinks2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectExtensionControllerExtenderProfileLanExtensionDownlinksName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileLanExtension-Downlinks-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileLanExtensionDownlinksPort2edl(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileLanExtension-Downlinks-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pvid"
		if _, ok := i["pvid"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileLanExtensionDownlinksPvid2edl(i["pvid"], d, pre_append)
			tmp["pvid"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileLanExtension-Downlinks-Pvid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileLanExtensionDownlinksType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileLanExtension-Downlinks-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vap"
		if _, ok := i["vap"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileLanExtensionDownlinksVap2edl(i["vap"], d, pre_append)
			tmp["vap"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileLanExtension-Downlinks-Vap")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionDownlinksName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionDownlinksPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionDownlinksPvid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionDownlinksType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionDownlinksVap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionIpsecTunnel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionLinkLoadbalance2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtensionControllerExtenderProfileLanExtension(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("backhaul", flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaul2edl(o["backhaul"], d, "backhaul")); err != nil {
			if vv, ok := fortiAPIPatch(o["backhaul"], "ObjectExtensionControllerExtenderProfileLanExtension-Backhaul"); ok {
				if err = d.Set("backhaul", vv); err != nil {
					return fmt.Errorf("Error reading backhaul: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading backhaul: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("backhaul"); ok {
			if err = d.Set("backhaul", flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaul2edl(o["backhaul"], d, "backhaul")); err != nil {
				if vv, ok := fortiAPIPatch(o["backhaul"], "ObjectExtensionControllerExtenderProfileLanExtension-Backhaul"); ok {
					if err = d.Set("backhaul", vv); err != nil {
						return fmt.Errorf("Error reading backhaul: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading backhaul: %v", err)
				}
			}
		}
	}

	if err = d.Set("backhaul_interface", flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulInterface2edl(o["backhaul-interface"], d, "backhaul_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["backhaul-interface"], "ObjectExtensionControllerExtenderProfileLanExtension-BackhaulInterface"); ok {
			if err = d.Set("backhaul_interface", vv); err != nil {
				return fmt.Errorf("Error reading backhaul_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading backhaul_interface: %v", err)
		}
	}

	if err = d.Set("backhaul_ip", flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulIp2edl(o["backhaul-ip"], d, "backhaul_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["backhaul-ip"], "ObjectExtensionControllerExtenderProfileLanExtension-BackhaulIp"); ok {
			if err = d.Set("backhaul_ip", vv); err != nil {
				return fmt.Errorf("Error reading backhaul_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading backhaul_ip: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("downlinks", flattenObjectExtensionControllerExtenderProfileLanExtensionDownlinks2edl(o["downlinks"], d, "downlinks")); err != nil {
			if vv, ok := fortiAPIPatch(o["downlinks"], "ObjectExtensionControllerExtenderProfileLanExtension-Downlinks"); ok {
				if err = d.Set("downlinks", vv); err != nil {
					return fmt.Errorf("Error reading downlinks: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading downlinks: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("downlinks"); ok {
			if err = d.Set("downlinks", flattenObjectExtensionControllerExtenderProfileLanExtensionDownlinks2edl(o["downlinks"], d, "downlinks")); err != nil {
				if vv, ok := fortiAPIPatch(o["downlinks"], "ObjectExtensionControllerExtenderProfileLanExtension-Downlinks"); ok {
					if err = d.Set("downlinks", vv); err != nil {
						return fmt.Errorf("Error reading downlinks: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading downlinks: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipsec_tunnel", flattenObjectExtensionControllerExtenderProfileLanExtensionIpsecTunnel2edl(o["ipsec-tunnel"], d, "ipsec_tunnel")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-tunnel"], "ObjectExtensionControllerExtenderProfileLanExtension-IpsecTunnel"); ok {
			if err = d.Set("ipsec_tunnel", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_tunnel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_tunnel: %v", err)
		}
	}

	if err = d.Set("link_loadbalance", flattenObjectExtensionControllerExtenderProfileLanExtensionLinkLoadbalance2edl(o["link-loadbalance"], d, "link_loadbalance")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-loadbalance"], "ObjectExtensionControllerExtenderProfileLanExtension-LinkLoadbalance"); ok {
			if err = d.Set("link_loadbalance", vv); err != nil {
				return fmt.Errorf("Error reading link_loadbalance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_loadbalance: %v", err)
		}
	}

	return nil
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaul2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulPort2edl(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulRole2edl(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulWeight2edl(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulRole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionDownlinks2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionDownlinksName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionDownlinksPort2edl(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pvid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pvid"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionDownlinksPvid2edl(d, i["pvid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionDownlinksType2edl(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vap"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vap"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionDownlinksVap2edl(d, i["vap"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionDownlinksName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionDownlinksPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionDownlinksPvid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionDownlinksType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionDownlinksVap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionIpsecTunnel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionLinkLoadbalance2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtensionControllerExtenderProfileLanExtension(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("backhaul"); ok || d.HasChange("backhaul") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionBackhaul2edl(d, v, "backhaul")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backhaul"] = t
		}
	}

	if v, ok := d.GetOk("backhaul_interface"); ok || d.HasChange("backhaul_interface") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulInterface2edl(d, v, "backhaul_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backhaul-interface"] = t
		}
	}

	if v, ok := d.GetOk("backhaul_ip"); ok || d.HasChange("backhaul_ip") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulIp2edl(d, v, "backhaul_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backhaul-ip"] = t
		}
	}

	if v, ok := d.GetOk("downlinks"); ok || d.HasChange("downlinks") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionDownlinks2edl(d, v, "downlinks")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downlinks"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_tunnel"); ok || d.HasChange("ipsec_tunnel") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionIpsecTunnel2edl(d, v, "ipsec_tunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-tunnel"] = t
		}
	}

	if v, ok := d.GetOk("link_loadbalance"); ok || d.HasChange("link_loadbalance") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionLinkLoadbalance2edl(d, v, "link_loadbalance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-loadbalance"] = t
		}
	}

	return &obj, nil
}
