// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: VRRP configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFspVlanInterfaceVrrp() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFspVlanInterfaceVrrpCreate,
		Read:   resourceObjectFspVlanInterfaceVrrpRead,
		Update: resourceObjectFspVlanInterfaceVrrpUpdate,
		Delete: resourceObjectFspVlanInterfaceVrrpDelete,

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
			"vlan": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"accept_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adv_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ignore_default_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preempt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proxy_arp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrdst": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vrdst_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrgrp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"vrip": &schema.Schema{
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

func resourceObjectFspVlanInterfaceVrrpCreate(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	obj, err := getObjectObjectFspVlanInterfaceVrrp(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanInterfaceVrrp resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFspVlanInterfaceVrrp(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanInterfaceVrrp resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "vrid")))

	return resourceObjectFspVlanInterfaceVrrpRead(d, m)
}

func resourceObjectFspVlanInterfaceVrrpUpdate(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	obj, err := getObjectObjectFspVlanInterfaceVrrp(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterfaceVrrp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFspVlanInterfaceVrrp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterfaceVrrp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "vrid")))

	return resourceObjectFspVlanInterfaceVrrpRead(d, m)
}

func resourceObjectFspVlanInterfaceVrrpDelete(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	wsParams["adom"] = adomv

	err = c.DeleteObjectFspVlanInterfaceVrrp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFspVlanInterfaceVrrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFspVlanInterfaceVrrpRead(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	if vlan == "" {
		vlan = importOptionChecking(m.(*FortiClient).Cfg, "vlan")
		if vlan == "" {
			return fmt.Errorf("Parameter vlan is missing")
		}
		if err = d.Set("vlan", vlan); err != nil {
			return fmt.Errorf("Error set params vlan: %v", err)
		}
	}
	paradict["vlan"] = vlan

	o, err := c.ReadObjectFspVlanInterfaceVrrp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterfaceVrrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFspVlanInterfaceVrrp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterfaceVrrp resource from API: %v", err)
	}
	return nil
}

func flattenObjectFspVlanInterfaceVrrpAcceptMode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpAdvInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpIgnoreDefaultRoute3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpPreempt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpPriority3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpProxyArp3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpProxyArpId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceVrrp-ProxyArp-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanInterfaceVrrpProxyArpIp3rdl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanInterfaceVrrp-ProxyArp-Ip")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanInterfaceVrrpProxyArpId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpProxyArpIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpStartTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVersion3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVrdst3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceVrrpVrdstPriority3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVrgrp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVrid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpVrip3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFspVlanInterfaceVrrp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("accept_mode", flattenObjectFspVlanInterfaceVrrpAcceptMode3rdl(o["accept-mode"], d, "accept_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["accept-mode"], "ObjectFspVlanInterfaceVrrp-AcceptMode"); ok {
			if err = d.Set("accept_mode", vv); err != nil {
				return fmt.Errorf("Error reading accept_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accept_mode: %v", err)
		}
	}

	if err = d.Set("adv_interval", flattenObjectFspVlanInterfaceVrrpAdvInterval3rdl(o["adv-interval"], d, "adv_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["adv-interval"], "ObjectFspVlanInterfaceVrrp-AdvInterval"); ok {
			if err = d.Set("adv_interval", vv); err != nil {
				return fmt.Errorf("Error reading adv_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adv_interval: %v", err)
		}
	}

	if err = d.Set("ignore_default_route", flattenObjectFspVlanInterfaceVrrpIgnoreDefaultRoute3rdl(o["ignore-default-route"], d, "ignore_default_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-default-route"], "ObjectFspVlanInterfaceVrrp-IgnoreDefaultRoute"); ok {
			if err = d.Set("ignore_default_route", vv); err != nil {
				return fmt.Errorf("Error reading ignore_default_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_default_route: %v", err)
		}
	}

	if err = d.Set("preempt", flattenObjectFspVlanInterfaceVrrpPreempt3rdl(o["preempt"], d, "preempt")); err != nil {
		if vv, ok := fortiAPIPatch(o["preempt"], "ObjectFspVlanInterfaceVrrp-Preempt"); ok {
			if err = d.Set("preempt", vv); err != nil {
				return fmt.Errorf("Error reading preempt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preempt: %v", err)
		}
	}

	if err = d.Set("priority", flattenObjectFspVlanInterfaceVrrpPriority3rdl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "ObjectFspVlanInterfaceVrrp-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("proxy_arp", flattenObjectFspVlanInterfaceVrrpProxyArp3rdl(o["proxy-arp"], d, "proxy_arp")); err != nil {
			if vv, ok := fortiAPIPatch(o["proxy-arp"], "ObjectFspVlanInterfaceVrrp-ProxyArp"); ok {
				if err = d.Set("proxy_arp", vv); err != nil {
					return fmt.Errorf("Error reading proxy_arp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading proxy_arp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("proxy_arp"); ok {
			if err = d.Set("proxy_arp", flattenObjectFspVlanInterfaceVrrpProxyArp3rdl(o["proxy-arp"], d, "proxy_arp")); err != nil {
				if vv, ok := fortiAPIPatch(o["proxy-arp"], "ObjectFspVlanInterfaceVrrp-ProxyArp"); ok {
					if err = d.Set("proxy_arp", vv); err != nil {
						return fmt.Errorf("Error reading proxy_arp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading proxy_arp: %v", err)
				}
			}
		}
	}

	if err = d.Set("start_time", flattenObjectFspVlanInterfaceVrrpStartTime3rdl(o["start-time"], d, "start_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-time"], "ObjectFspVlanInterfaceVrrp-StartTime"); ok {
			if err = d.Set("start_time", vv); err != nil {
				return fmt.Errorf("Error reading start_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_time: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectFspVlanInterfaceVrrpStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectFspVlanInterfaceVrrp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("version", flattenObjectFspVlanInterfaceVrrpVersion3rdl(o["version"], d, "version")); err != nil {
		if vv, ok := fortiAPIPatch(o["version"], "ObjectFspVlanInterfaceVrrp-Version"); ok {
			if err = d.Set("version", vv); err != nil {
				return fmt.Errorf("Error reading version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	if err = d.Set("vrdst", flattenObjectFspVlanInterfaceVrrpVrdst3rdl(o["vrdst"], d, "vrdst")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrdst"], "ObjectFspVlanInterfaceVrrp-Vrdst"); ok {
			if err = d.Set("vrdst", vv); err != nil {
				return fmt.Errorf("Error reading vrdst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrdst: %v", err)
		}
	}

	if err = d.Set("vrdst_priority", flattenObjectFspVlanInterfaceVrrpVrdstPriority3rdl(o["vrdst-priority"], d, "vrdst_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrdst-priority"], "ObjectFspVlanInterfaceVrrp-VrdstPriority"); ok {
			if err = d.Set("vrdst_priority", vv); err != nil {
				return fmt.Errorf("Error reading vrdst_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrdst_priority: %v", err)
		}
	}

	if err = d.Set("vrgrp", flattenObjectFspVlanInterfaceVrrpVrgrp3rdl(o["vrgrp"], d, "vrgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrgrp"], "ObjectFspVlanInterfaceVrrp-Vrgrp"); ok {
			if err = d.Set("vrgrp", vv); err != nil {
				return fmt.Errorf("Error reading vrgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrgrp: %v", err)
		}
	}

	if err = d.Set("vrid", flattenObjectFspVlanInterfaceVrrpVrid3rdl(o["vrid"], d, "vrid")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrid"], "ObjectFspVlanInterfaceVrrp-Vrid"); ok {
			if err = d.Set("vrid", vv); err != nil {
				return fmt.Errorf("Error reading vrid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrid: %v", err)
		}
	}

	if err = d.Set("vrip", flattenObjectFspVlanInterfaceVrrpVrip3rdl(o["vrip"], d, "vrip")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrip"], "ObjectFspVlanInterfaceVrrp-Vrip"); ok {
			if err = d.Set("vrip", vv); err != nil {
				return fmt.Errorf("Error reading vrip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrip: %v", err)
		}
	}

	return nil
}

func flattenObjectFspVlanInterfaceVrrpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFspVlanInterfaceVrrpAcceptMode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpAdvInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpIgnoreDefaultRoute3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpPreempt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpPriority3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpProxyArp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanInterfaceVrrpProxyArpId3rdl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanInterfaceVrrpProxyArpIp3rdl(d, i["ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanInterfaceVrrpProxyArpId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpProxyArpIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpStartTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVersion3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVrdst3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceVrrpVrdstPriority3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVrgrp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVrid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpVrip3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFspVlanInterfaceVrrp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accept_mode"); ok || d.HasChange("accept_mode") {
		t, err := expandObjectFspVlanInterfaceVrrpAcceptMode3rdl(d, v, "accept_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accept-mode"] = t
		}
	}

	if v, ok := d.GetOk("adv_interval"); ok || d.HasChange("adv_interval") {
		t, err := expandObjectFspVlanInterfaceVrrpAdvInterval3rdl(d, v, "adv_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-interval"] = t
		}
	}

	if v, ok := d.GetOk("ignore_default_route"); ok || d.HasChange("ignore_default_route") {
		t, err := expandObjectFspVlanInterfaceVrrpIgnoreDefaultRoute3rdl(d, v, "ignore_default_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-default-route"] = t
		}
	}

	if v, ok := d.GetOk("preempt"); ok || d.HasChange("preempt") {
		t, err := expandObjectFspVlanInterfaceVrrpPreempt3rdl(d, v, "preempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preempt"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandObjectFspVlanInterfaceVrrpPriority3rdl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("proxy_arp"); ok || d.HasChange("proxy_arp") {
		t, err := expandObjectFspVlanInterfaceVrrpProxyArp3rdl(d, v, "proxy_arp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-arp"] = t
		}
	}

	if v, ok := d.GetOk("start_time"); ok || d.HasChange("start_time") {
		t, err := expandObjectFspVlanInterfaceVrrpStartTime3rdl(d, v, "start_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-time"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectFspVlanInterfaceVrrpStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandObjectFspVlanInterfaceVrrpVersion3rdl(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	if v, ok := d.GetOk("vrdst"); ok || d.HasChange("vrdst") {
		t, err := expandObjectFspVlanInterfaceVrrpVrdst3rdl(d, v, "vrdst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrdst"] = t
		}
	}

	if v, ok := d.GetOk("vrdst_priority"); ok || d.HasChange("vrdst_priority") {
		t, err := expandObjectFspVlanInterfaceVrrpVrdstPriority3rdl(d, v, "vrdst_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrdst-priority"] = t
		}
	}

	if v, ok := d.GetOk("vrgrp"); ok || d.HasChange("vrgrp") {
		t, err := expandObjectFspVlanInterfaceVrrpVrgrp3rdl(d, v, "vrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrgrp"] = t
		}
	}

	if v, ok := d.GetOk("vrid"); ok || d.HasChange("vrid") {
		t, err := expandObjectFspVlanInterfaceVrrpVrid3rdl(d, v, "vrid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrid"] = t
		}
	}

	if v, ok := d.GetOk("vrip"); ok || d.HasChange("vrip") {
		t, err := expandObjectFspVlanInterfaceVrrpVrip3rdl(d, v, "vrip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrip"] = t
		}
	}

	return &obj, nil
}
