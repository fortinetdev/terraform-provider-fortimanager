// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: HA configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemHa() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHaUpdate,
		Read:   resourceSystemHaRead,
		Update: resourceSystemHaUpdate,
		Delete: resourceSystemHaDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"clusterid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"failover_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hb_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hb_lost_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitored_interfaces": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"monitored_ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"peer": &schema.Schema{
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
						"ip6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"unicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrrp_adv_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vrrp_interface": &schema.Schema{
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

func resourceSystemHaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemHa(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemHa(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemHa")

	return resourceSystemHaRead(d, m)
}

func resourceSystemHaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemHa(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemHa(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemHa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHa(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHa resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaClusterid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaFailoverMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaFileQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHbInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHbLostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaLocalCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitoredInterfaces(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := i["interface-name"]; ok {
			v := flattenSystemHaMonitoredInterfacesInterfaceName(i["interface-name"], d, pre_append)
			tmp["interface_name"] = fortiAPISubPartPatch(v, "SystemHa-MonitoredInterfaces-InterfaceName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemHaMonitoredInterfacesInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitoredIps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemHaMonitoredIpsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemHa-MonitoredIps-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenSystemHaMonitoredIpsInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "SystemHa-MonitoredIps-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemHaMonitoredIpsIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemHa-MonitoredIps-Ip")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemHaMonitoredIpsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitoredIpsInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitoredIpsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemHaPeerId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemHa-Peer-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemHaPeerIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemHa-Peer-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenSystemHaPeerIp6(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "SystemHa-Peer-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial_number"
		if _, ok := i["serial-number"]; ok {
			v := flattenSystemHaPeerSerialNumber(i["serial-number"], d, pre_append)
			tmp["serial_number"] = fortiAPISubPartPatch(v, "SystemHa-Peer-SerialNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSystemHaPeerStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SystemHa-Peer-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemHaPeerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUnicast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVipInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVrrpAdvInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVrrpInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemHa(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("clusterid", flattenSystemHaClusterid(o["clusterid"], d, "clusterid")); err != nil {
		if vv, ok := fortiAPIPatch(o["clusterid"], "SystemHa-Clusterid"); ok {
			if err = d.Set("clusterid", vv); err != nil {
				return fmt.Errorf("Error reading clusterid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clusterid: %v", err)
		}
	}

	if err = d.Set("failover_mode", flattenSystemHaFailoverMode(o["failover-mode"], d, "failover_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["failover-mode"], "SystemHa-FailoverMode"); ok {
			if err = d.Set("failover_mode", vv); err != nil {
				return fmt.Errorf("Error reading failover_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failover_mode: %v", err)
		}
	}

	if err = d.Set("file_quota", flattenSystemHaFileQuota(o["file-quota"], d, "file_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-quota"], "SystemHa-FileQuota"); ok {
			if err = d.Set("file_quota", vv); err != nil {
				return fmt.Errorf("Error reading file_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_quota: %v", err)
		}
	}

	if err = d.Set("hb_interval", flattenSystemHaHbInterval(o["hb-interval"], d, "hb_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["hb-interval"], "SystemHa-HbInterval"); ok {
			if err = d.Set("hb_interval", vv); err != nil {
				return fmt.Errorf("Error reading hb_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hb_interval: %v", err)
		}
	}

	if err = d.Set("hb_lost_threshold", flattenSystemHaHbLostThreshold(o["hb-lost-threshold"], d, "hb_lost_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["hb-lost-threshold"], "SystemHa-HbLostThreshold"); ok {
			if err = d.Set("hb_lost_threshold", vv); err != nil {
				return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
		}
	}

	if err = d.Set("local_cert", flattenSystemHaLocalCert(o["local-cert"], d, "local_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-cert"], "SystemHa-LocalCert"); ok {
			if err = d.Set("local_cert", vv); err != nil {
				return fmt.Errorf("Error reading local_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_cert: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemHaMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemHa-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("monitored_interfaces", flattenSystemHaMonitoredInterfaces(o["monitored-interfaces"], d, "monitored_interfaces")); err != nil {
			if vv, ok := fortiAPIPatch(o["monitored-interfaces"], "SystemHa-MonitoredInterfaces"); ok {
				if err = d.Set("monitored_interfaces", vv); err != nil {
					return fmt.Errorf("Error reading monitored_interfaces: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading monitored_interfaces: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("monitored_interfaces"); ok {
			if err = d.Set("monitored_interfaces", flattenSystemHaMonitoredInterfaces(o["monitored-interfaces"], d, "monitored_interfaces")); err != nil {
				if vv, ok := fortiAPIPatch(o["monitored-interfaces"], "SystemHa-MonitoredInterfaces"); ok {
					if err = d.Set("monitored_interfaces", vv); err != nil {
						return fmt.Errorf("Error reading monitored_interfaces: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading monitored_interfaces: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("monitored_ips", flattenSystemHaMonitoredIps(o["monitored-ips"], d, "monitored_ips")); err != nil {
			if vv, ok := fortiAPIPatch(o["monitored-ips"], "SystemHa-MonitoredIps"); ok {
				if err = d.Set("monitored_ips", vv); err != nil {
					return fmt.Errorf("Error reading monitored_ips: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading monitored_ips: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("monitored_ips"); ok {
			if err = d.Set("monitored_ips", flattenSystemHaMonitoredIps(o["monitored-ips"], d, "monitored_ips")); err != nil {
				if vv, ok := fortiAPIPatch(o["monitored-ips"], "SystemHa-MonitoredIps"); ok {
					if err = d.Set("monitored_ips", vv); err != nil {
						return fmt.Errorf("Error reading monitored_ips: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading monitored_ips: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("peer", flattenSystemHaPeer(o["peer"], d, "peer")); err != nil {
			if vv, ok := fortiAPIPatch(o["peer"], "SystemHa-Peer"); ok {
				if err = d.Set("peer", vv); err != nil {
					return fmt.Errorf("Error reading peer: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading peer: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("peer"); ok {
			if err = d.Set("peer", flattenSystemHaPeer(o["peer"], d, "peer")); err != nil {
				if vv, ok := fortiAPIPatch(o["peer"], "SystemHa-Peer"); ok {
					if err = d.Set("peer", vv); err != nil {
						return fmt.Errorf("Error reading peer: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading peer: %v", err)
				}
			}
		}
	}

	if err = d.Set("priority", flattenSystemHaPriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "SystemHa-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("unicast", flattenSystemHaUnicast(o["unicast"], d, "unicast")); err != nil {
		if vv, ok := fortiAPIPatch(o["unicast"], "SystemHa-Unicast"); ok {
			if err = d.Set("unicast", vv); err != nil {
				return fmt.Errorf("Error reading unicast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unicast: %v", err)
		}
	}

	if err = d.Set("vip", flattenSystemHaVip(o["vip"], d, "vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip"], "SystemHa-Vip"); ok {
			if err = d.Set("vip", vv); err != nil {
				return fmt.Errorf("Error reading vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip: %v", err)
		}
	}

	if err = d.Set("vip_interface", flattenSystemHaVipInterface(o["vip-interface"], d, "vip_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip-interface"], "SystemHa-VipInterface"); ok {
			if err = d.Set("vip_interface", vv); err != nil {
				return fmt.Errorf("Error reading vip_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip_interface: %v", err)
		}
	}

	if err = d.Set("vrrp_adv_interval", flattenSystemHaVrrpAdvInterval(o["vrrp-adv-interval"], d, "vrrp_adv_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrrp-adv-interval"], "SystemHa-VrrpAdvInterval"); ok {
			if err = d.Set("vrrp_adv_interval", vv); err != nil {
				return fmt.Errorf("Error reading vrrp_adv_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrrp_adv_interval: %v", err)
		}
	}

	if err = d.Set("vrrp_interface", flattenSystemHaVrrpInterface(o["vrrp-interface"], d, "vrrp_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrrp-interface"], "SystemHa-VrrpInterface"); ok {
			if err = d.Set("vrrp_interface", vv); err != nil {
				return fmt.Errorf("Error reading vrrp_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrrp_interface: %v", err)
		}
	}

	return nil
}

func flattenSystemHaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHaClusterid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaFailoverMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaFileQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbLostThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaLocalCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitoredInterfaces(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-name"], _ = expandSystemHaMonitoredInterfacesInterfaceName(d, i["interface_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemHaMonitoredInterfacesInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitoredIps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemHaMonitoredIpsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandSystemHaMonitoredIpsInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSystemHaMonitoredIpsIp(d, i["ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemHaMonitoredIpsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitoredIpsInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitoredIpsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemHaPeerId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSystemHaPeerIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandSystemHaPeerIp6(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["serial-number"], _ = expandSystemHaPeerSerialNumber(d, i["serial_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSystemHaPeerStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemHaPeerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerSerialNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVipInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVrrpAdvInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVrrpInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHa(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("clusterid"); ok || d.HasChange("clusterid") {
		t, err := expandSystemHaClusterid(d, v, "clusterid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clusterid"] = t
		}
	}

	if v, ok := d.GetOk("failover_mode"); ok || d.HasChange("failover_mode") {
		t, err := expandSystemHaFailoverMode(d, v, "failover_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failover-mode"] = t
		}
	}

	if v, ok := d.GetOk("file_quota"); ok || d.HasChange("file_quota") {
		t, err := expandSystemHaFileQuota(d, v, "file_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-quota"] = t
		}
	}

	if v, ok := d.GetOk("hb_interval"); ok || d.HasChange("hb_interval") {
		t, err := expandSystemHaHbInterval(d, v, "hb_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-interval"] = t
		}
	}

	if v, ok := d.GetOk("hb_lost_threshold"); ok || d.HasChange("hb_lost_threshold") {
		t, err := expandSystemHaHbLostThreshold(d, v, "hb_lost_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-lost-threshold"] = t
		}
	}

	if v, ok := d.GetOk("local_cert"); ok || d.HasChange("local_cert") {
		t, err := expandSystemHaLocalCert(d, v, "local_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-cert"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemHaMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("monitored_interfaces"); ok || d.HasChange("monitored_interfaces") {
		t, err := expandSystemHaMonitoredInterfaces(d, v, "monitored_interfaces")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitored-interfaces"] = t
		}
	}

	if v, ok := d.GetOk("monitored_ips"); ok || d.HasChange("monitored_ips") {
		t, err := expandSystemHaMonitoredIps(d, v, "monitored_ips")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitored-ips"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemHaPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("peer"); ok || d.HasChange("peer") {
		t, err := expandSystemHaPeer(d, v, "peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandSystemHaPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("unicast"); ok || d.HasChange("unicast") {
		t, err := expandSystemHaUnicast(d, v, "unicast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unicast"] = t
		}
	}

	if v, ok := d.GetOk("vip"); ok || d.HasChange("vip") {
		t, err := expandSystemHaVip(d, v, "vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip"] = t
		}
	}

	if v, ok := d.GetOk("vip_interface"); ok || d.HasChange("vip_interface") {
		t, err := expandSystemHaVipInterface(d, v, "vip_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip-interface"] = t
		}
	}

	if v, ok := d.GetOk("vrrp_adv_interval"); ok || d.HasChange("vrrp_adv_interval") {
		t, err := expandSystemHaVrrpAdvInterval(d, v, "vrrp_adv_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrrp-adv-interval"] = t
		}
	}

	if v, ok := d.GetOk("vrrp_interface"); ok || d.HasChange("vrrp_interface") {
		t, err := expandSystemHaVrrpInterface(d, v, "vrrp_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrrp-interface"] = t
		}
	}

	return &obj, nil
}
