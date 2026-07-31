// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Systemp SystemTemplateInterface

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystempSystemTemplateInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystempSystemTemplateInterfaceCreate,
		Read:   resourceSystempSystemTemplateInterfaceRead,
		Update: resourceSystempSystemTemplateInterfaceUpdate,
		Delete: resourceSystempSystemTemplateInterfaceDelete,

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
			"devprof": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dhcp_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface_members": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ipmask": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"monitor_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"seq": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wifi_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"wifi_security": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wifi_ssid": &schema.Schema{
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

func resourceSystempSystemTemplateInterfaceCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystempSystemTemplateInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating SystempSystemTemplateInterface resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("seq")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystempSystemTemplateInterface(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystempSystemTemplateInterface(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystempSystemTemplateInterface resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateSystempSystemTemplateInterface(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystempSystemTemplateInterface resource: %v", err)
		}

		if v != nil && v["seq"] != nil {
			if vidn, ok := v["seq"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceSystempSystemTemplateInterfaceRead(d, m)
			} else {
				return fmt.Errorf("Error creating SystempSystemTemplateInterface resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "seq")))

	return resourceSystempSystemTemplateInterfaceRead(d, m)
}

func resourceSystempSystemTemplateInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystempSystemTemplateInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemTemplateInterface resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	v, err := c.UpdateSystempSystemTemplateInterface(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemTemplateInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	if v != nil && v["seq"] != nil {
		if vidn, ok := v["seq"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceSystempSystemTemplateInterfaceRead(d, m)
		} else {
			return fmt.Errorf("Error updating SystempSystemTemplateInterface resource: %v", err)
		}
	}

	return resourceSystempSystemTemplateInterfaceRead(d, m)
}

func resourceSystempSystemTemplateInterfaceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystempSystemTemplateInterface(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystempSystemTemplateInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystempSystemTemplateInterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystempSystemTemplateInterface(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystempSystemTemplateInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystempSystemTemplateInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemTemplateInterface resource from API: %v", err)
	}
	return nil
}

func flattenSystempSystemTemplateInterfaceAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempSystemTemplateInterfaceDhcpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceInterfaceMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempSystemTemplateInterfaceIpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenSystempSystemTemplateInterfaceIpRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "SystempSystemTemplateInterface-IpRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystempSystemTemplateInterfaceIpRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystempSystemTemplateInterface-IpRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenSystempSystemTemplateInterfaceIpRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "SystempSystemTemplateInterface-IpRange-StartIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystempSystemTemplateInterfaceIpRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceIpRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceIpRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceIpmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempSystemTemplateInterfaceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceModel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceMonitorBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceSeq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceWifiSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceWifiSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystempSystemTemplateInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("action", flattenSystempSystemTemplateInterfaceAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "SystempSystemTemplateInterface-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("alias", flattenSystempSystemTemplateInterfaceAlias(o["alias"], d, "alias")); err != nil {
		if vv, ok := fortiAPIPatch(o["alias"], "SystempSystemTemplateInterface-Alias"); ok {
			if err = d.Set("alias", vv); err != nil {
				return fmt.Errorf("Error reading alias: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenSystempSystemTemplateInterfaceAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowaccess"], "SystempSystemTemplateInterface-Allowaccess"); ok {
			if err = d.Set("allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("dhcp_id", flattenSystempSystemTemplateInterfaceDhcpId(o["dhcp-id"], d, "dhcp_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-id"], "SystempSystemTemplateInterface-DhcpId"); ok {
			if err = d.Set("dhcp_id", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_id: %v", err)
		}
	}

	if err = d.Set("gateway", flattenSystempSystemTemplateInterfaceGateway(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "SystempSystemTemplateInterface-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystempSystemTemplateInterfaceInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystempSystemTemplateInterface-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_members", flattenSystempSystemTemplateInterfaceInterfaceMembers(o["interface-members"], d, "interface_members")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-members"], "SystempSystemTemplateInterface-InterfaceMembers"); ok {
			if err = d.Set("interface_members", vv); err != nil {
				return fmt.Errorf("Error reading interface_members: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_members: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip_range", flattenSystempSystemTemplateInterfaceIpRange(o["ip-range"], d, "ip_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-range"], "SystempSystemTemplateInterface-IpRange"); ok {
				if err = d.Set("ip_range", vv); err != nil {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_range"); ok {
			if err = d.Set("ip_range", flattenSystempSystemTemplateInterfaceIpRange(o["ip-range"], d, "ip_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-range"], "SystempSystemTemplateInterface-IpRange"); ok {
					if err = d.Set("ip_range", vv); err != nil {
						return fmt.Errorf("Error reading ip_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipmask", flattenSystempSystemTemplateInterfaceIpmask(o["ipmask"], d, "ipmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipmask"], "SystempSystemTemplateInterface-Ipmask"); ok {
			if err = d.Set("ipmask", vv); err != nil {
				return fmt.Errorf("Error reading ipmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipmask: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystempSystemTemplateInterfaceMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystempSystemTemplateInterface-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("model", flattenSystempSystemTemplateInterfaceModel(o["model"], d, "model")); err != nil {
		if vv, ok := fortiAPIPatch(o["model"], "SystempSystemTemplateInterface-Model"); ok {
			if err = d.Set("model", vv); err != nil {
				return fmt.Errorf("Error reading model: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading model: %v", err)
		}
	}

	if err = d.Set("monitor_bandwidth", flattenSystempSystemTemplateInterfaceMonitorBandwidth(o["monitor-bandwidth"], d, "monitor_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-bandwidth"], "SystempSystemTemplateInterface-MonitorBandwidth"); ok {
			if err = d.Set("monitor_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading monitor_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_bandwidth: %v", err)
		}
	}

	if err = d.Set("name", flattenSystempSystemTemplateInterfaceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystempSystemTemplateInterface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("netmask", flattenSystempSystemTemplateInterfaceNetmask(o["netmask"], d, "netmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["netmask"], "SystempSystemTemplateInterface-Netmask"); ok {
			if err = d.Set("netmask", vv); err != nil {
				return fmt.Errorf("Error reading netmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netmask: %v", err)
		}
	}

	if err = d.Set("role", flattenSystempSystemTemplateInterfaceRole(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "SystempSystemTemplateInterface-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("seq", flattenSystempSystemTemplateInterfaceSeq(o["seq"], d, "seq")); err != nil {
		if vv, ok := fortiAPIPatch(o["seq"], "SystempSystemTemplateInterface-Seq"); ok {
			if err = d.Set("seq", vv); err != nil {
				return fmt.Errorf("Error reading seq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seq: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystempSystemTemplateInterfaceVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystempSystemTemplateInterface-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("vlan_id", flattenSystempSystemTemplateInterfaceVlanId(o["vlan-id"], d, "vlan_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-id"], "SystempSystemTemplateInterface-VlanId"); ok {
			if err = d.Set("vlan_id", vv); err != nil {
				return fmt.Errorf("Error reading vlan_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_id: %v", err)
		}
	}

	if err = d.Set("wifi_security", flattenSystempSystemTemplateInterfaceWifiSecurity(o["wifi-security"], d, "wifi_security")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-security"], "SystempSystemTemplateInterface-WifiSecurity"); ok {
			if err = d.Set("wifi_security", vv); err != nil {
				return fmt.Errorf("Error reading wifi_security: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_security: %v", err)
		}
	}

	if err = d.Set("wifi_ssid", flattenSystempSystemTemplateInterfaceWifiSsid(o["wifi-ssid"], d, "wifi_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ssid"], "SystempSystemTemplateInterface-WifiSsid"); ok {
			if err = d.Set("wifi_ssid", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ssid: %v", err)
		}
	}

	return nil
}

func flattenSystempSystemTemplateInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystempSystemTemplateInterfaceAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceAlias(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemTemplateInterfaceDhcpId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceInterfaceMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemTemplateInterfaceIpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandSystempSystemTemplateInterfaceIpRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystempSystemTemplateInterfaceIpRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandSystempSystemTemplateInterfaceIpRangeStartIp(d, i["start_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystempSystemTemplateInterfaceIpRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceIpRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceIpRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceIpmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystempSystemTemplateInterfaceMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceModel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceMonitorBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceNetmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceSeq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceWifiKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemTemplateInterfaceWifiSecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceWifiSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystempSystemTemplateInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandSystempSystemTemplateInterfaceAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("alias"); ok || d.HasChange("alias") {
		t, err := expandSystempSystemTemplateInterfaceAlias(d, v, "alias")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok || d.HasChange("allowaccess") {
		t, err := expandSystempSystemTemplateInterfaceAllowaccess(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_id"); ok || d.HasChange("dhcp_id") {
		t, err := expandSystempSystemTemplateInterfaceDhcpId(d, v, "dhcp_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-id"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandSystempSystemTemplateInterfaceGateway(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystempSystemTemplateInterfaceInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_members"); ok || d.HasChange("interface_members") {
		t, err := expandSystempSystemTemplateInterfaceInterfaceMembers(d, v, "interface_members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-members"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok || d.HasChange("ip_range") {
		t, err := expandSystempSystemTemplateInterfaceIpRange(d, v, "ip_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	if v, ok := d.GetOk("ipmask"); ok || d.HasChange("ipmask") {
		t, err := expandSystempSystemTemplateInterfaceIpmask(d, v, "ipmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipmask"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystempSystemTemplateInterfaceMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("model"); ok || d.HasChange("model") {
		t, err := expandSystempSystemTemplateInterfaceModel(d, v, "model")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["model"] = t
		}
	}

	if v, ok := d.GetOk("monitor_bandwidth"); ok || d.HasChange("monitor_bandwidth") {
		t, err := expandSystempSystemTemplateInterfaceMonitorBandwidth(d, v, "monitor_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystempSystemTemplateInterfaceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("netmask"); ok || d.HasChange("netmask") {
		t, err := expandSystempSystemTemplateInterfaceNetmask(d, v, "netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netmask"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandSystempSystemTemplateInterfaceRole(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("seq"); ok || d.HasChange("seq") {
		t, err := expandSystempSystemTemplateInterfaceSeq(d, v, "seq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystempSystemTemplateInterfaceVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("vlan_id"); ok || d.HasChange("vlan_id") {
		t, err := expandSystempSystemTemplateInterfaceVlanId(d, v, "vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-id"] = t
		}
	}

	if v, ok := d.GetOk("wifi_key"); ok || d.HasChange("wifi_key") {
		t, err := expandSystempSystemTemplateInterfaceWifiKey(d, v, "wifi_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-key"] = t
		}
	}

	if v, ok := d.GetOk("wifi_security"); ok || d.HasChange("wifi_security") {
		t, err := expandSystempSystemTemplateInterfaceWifiSecurity(d, v, "wifi_security")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-security"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ssid"); ok || d.HasChange("wifi_ssid") {
		t, err := expandSystempSystemTemplateInterfaceWifiSsid(d, v, "wifi_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ssid"] = t
		}
	}

	return &obj, nil
}
