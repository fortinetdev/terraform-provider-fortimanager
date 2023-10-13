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
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemHa(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHa(obj, mkey, paradict)
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
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemHa(mkey, paradict)
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

func flattenSystemHaClusteridSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaFailoverModeSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaFileQuotaSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHbIntervalSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHbLostThresholdSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaLocalCertSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaModeSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitoredInterfacesSha(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemHaMonitoredInterfacesInterfaceNameSha(i["interface-name"], d, pre_append)
			tmp["interface_name"] = fortiAPISubPartPatch(v, "SystemHa-MonitoredInterfaces-InterfaceName")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemHaMonitoredInterfacesInterfaceNameSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitoredIpsSha(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemHaMonitoredIpsIdSha(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemHa-MonitoredIps-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenSystemHaMonitoredIpsInterfaceSha(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "SystemHa-MonitoredIps-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemHaMonitoredIpsIpSha(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemHa-MonitoredIps-Ip")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemHaMonitoredIpsIdSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitoredIpsInterfaceSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitoredIpsIpSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPasswordSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaPeerSha(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemHaPeerIdSha(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemHa-Peer-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemHaPeerIpSha(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemHa-Peer-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenSystemHaPeerIp6Sha(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "SystemHa-Peer-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial_number"
		if _, ok := i["serial-number"]; ok {
			v := flattenSystemHaPeerSerialNumberSha(i["serial-number"], d, pre_append)
			tmp["serial_number"] = fortiAPISubPartPatch(v, "SystemHa-Peer-SerialNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSystemHaPeerStatusSha(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SystemHa-Peer-Status")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemHaPeerIdSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerIpSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerIp6Sha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerSerialNumberSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerStatusSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPrioritySha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUnicastSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVipSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVipInterfaceSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVrrpAdvIntervalSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVrrpInterfaceSha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemHa(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("clusterid", flattenSystemHaClusteridSha(o["clusterid"], d, "clusterid")); err != nil {
		if vv, ok := fortiAPIPatch(o["clusterid"], "SystemHa-Clusterid"); ok {
			if err = d.Set("clusterid", vv); err != nil {
				return fmt.Errorf("Error reading clusterid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clusterid: %v", err)
		}
	}

	if err = d.Set("failover_mode", flattenSystemHaFailoverModeSha(o["failover-mode"], d, "failover_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["failover-mode"], "SystemHa-FailoverMode"); ok {
			if err = d.Set("failover_mode", vv); err != nil {
				return fmt.Errorf("Error reading failover_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failover_mode: %v", err)
		}
	}

	if err = d.Set("file_quota", flattenSystemHaFileQuotaSha(o["file-quota"], d, "file_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-quota"], "SystemHa-FileQuota"); ok {
			if err = d.Set("file_quota", vv); err != nil {
				return fmt.Errorf("Error reading file_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_quota: %v", err)
		}
	}

	if err = d.Set("hb_interval", flattenSystemHaHbIntervalSha(o["hb-interval"], d, "hb_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["hb-interval"], "SystemHa-HbInterval"); ok {
			if err = d.Set("hb_interval", vv); err != nil {
				return fmt.Errorf("Error reading hb_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hb_interval: %v", err)
		}
	}

	if err = d.Set("hb_lost_threshold", flattenSystemHaHbLostThresholdSha(o["hb-lost-threshold"], d, "hb_lost_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["hb-lost-threshold"], "SystemHa-HbLostThreshold"); ok {
			if err = d.Set("hb_lost_threshold", vv); err != nil {
				return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
		}
	}

	if err = d.Set("local_cert", flattenSystemHaLocalCertSha(o["local-cert"], d, "local_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-cert"], "SystemHa-LocalCert"); ok {
			if err = d.Set("local_cert", vv); err != nil {
				return fmt.Errorf("Error reading local_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_cert: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemHaModeSha(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemHa-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("monitored_interfaces", flattenSystemHaMonitoredInterfacesSha(o["monitored-interfaces"], d, "monitored_interfaces")); err != nil {
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
			if err = d.Set("monitored_interfaces", flattenSystemHaMonitoredInterfacesSha(o["monitored-interfaces"], d, "monitored_interfaces")); err != nil {
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
		if err = d.Set("monitored_ips", flattenSystemHaMonitoredIpsSha(o["monitored-ips"], d, "monitored_ips")); err != nil {
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
			if err = d.Set("monitored_ips", flattenSystemHaMonitoredIpsSha(o["monitored-ips"], d, "monitored_ips")); err != nil {
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
		if err = d.Set("peer", flattenSystemHaPeerSha(o["peer"], d, "peer")); err != nil {
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
			if err = d.Set("peer", flattenSystemHaPeerSha(o["peer"], d, "peer")); err != nil {
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

	if err = d.Set("priority", flattenSystemHaPrioritySha(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "SystemHa-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("unicast", flattenSystemHaUnicastSha(o["unicast"], d, "unicast")); err != nil {
		if vv, ok := fortiAPIPatch(o["unicast"], "SystemHa-Unicast"); ok {
			if err = d.Set("unicast", vv); err != nil {
				return fmt.Errorf("Error reading unicast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unicast: %v", err)
		}
	}

	if err = d.Set("vip", flattenSystemHaVipSha(o["vip"], d, "vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip"], "SystemHa-Vip"); ok {
			if err = d.Set("vip", vv); err != nil {
				return fmt.Errorf("Error reading vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip: %v", err)
		}
	}

	if err = d.Set("vip_interface", flattenSystemHaVipInterfaceSha(o["vip-interface"], d, "vip_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip-interface"], "SystemHa-VipInterface"); ok {
			if err = d.Set("vip_interface", vv); err != nil {
				return fmt.Errorf("Error reading vip_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip_interface: %v", err)
		}
	}

	if err = d.Set("vrrp_adv_interval", flattenSystemHaVrrpAdvIntervalSha(o["vrrp-adv-interval"], d, "vrrp_adv_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrrp-adv-interval"], "SystemHa-VrrpAdvInterval"); ok {
			if err = d.Set("vrrp_adv_interval", vv); err != nil {
				return fmt.Errorf("Error reading vrrp_adv_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrrp_adv_interval: %v", err)
		}
	}

	if err = d.Set("vrrp_interface", flattenSystemHaVrrpInterfaceSha(o["vrrp-interface"], d, "vrrp_interface")); err != nil {
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

func expandSystemHaClusteridSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaFailoverModeSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaFileQuotaSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbIntervalSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbLostThresholdSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaLocalCertSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaModeSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitoredInterfacesSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["interface-name"], _ = expandSystemHaMonitoredInterfacesInterfaceNameSha(d, i["interface_name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemHaMonitoredInterfacesInterfaceNameSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitoredIpsSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemHaMonitoredIpsIdSha(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandSystemHaMonitoredIpsInterfaceSha(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSystemHaMonitoredIpsIpSha(d, i["ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemHaMonitoredIpsIdSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitoredIpsInterfaceSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitoredIpsIpSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPasswordSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaPeerSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemHaPeerIdSha(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSystemHaPeerIpSha(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandSystemHaPeerIp6Sha(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["serial-number"], _ = expandSystemHaPeerSerialNumberSha(d, i["serial_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSystemHaPeerStatusSha(d, i["status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemHaPeerIdSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerIpSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerIp6Sha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerSerialNumberSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerStatusSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPrioritySha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVipSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVipInterfaceSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVrrpAdvIntervalSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVrrpInterfaceSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHa(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("clusterid"); ok || d.HasChange("clusterid") {
		t, err := expandSystemHaClusteridSha(d, v, "clusterid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clusterid"] = t
		}
	}

	if v, ok := d.GetOk("failover_mode"); ok || d.HasChange("failover_mode") {
		t, err := expandSystemHaFailoverModeSha(d, v, "failover_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failover-mode"] = t
		}
	}

	if v, ok := d.GetOk("file_quota"); ok || d.HasChange("file_quota") {
		t, err := expandSystemHaFileQuotaSha(d, v, "file_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-quota"] = t
		}
	}

	if v, ok := d.GetOk("hb_interval"); ok || d.HasChange("hb_interval") {
		t, err := expandSystemHaHbIntervalSha(d, v, "hb_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-interval"] = t
		}
	}

	if v, ok := d.GetOk("hb_lost_threshold"); ok || d.HasChange("hb_lost_threshold") {
		t, err := expandSystemHaHbLostThresholdSha(d, v, "hb_lost_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-lost-threshold"] = t
		}
	}

	if v, ok := d.GetOk("local_cert"); ok || d.HasChange("local_cert") {
		t, err := expandSystemHaLocalCertSha(d, v, "local_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-cert"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemHaModeSha(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("monitored_interfaces"); ok || d.HasChange("monitored_interfaces") {
		t, err := expandSystemHaMonitoredInterfacesSha(d, v, "monitored_interfaces")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitored-interfaces"] = t
		}
	}

	if v, ok := d.GetOk("monitored_ips"); ok || d.HasChange("monitored_ips") {
		t, err := expandSystemHaMonitoredIpsSha(d, v, "monitored_ips")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitored-ips"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemHaPasswordSha(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("peer"); ok || d.HasChange("peer") {
		t, err := expandSystemHaPeerSha(d, v, "peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandSystemHaPrioritySha(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("unicast"); ok || d.HasChange("unicast") {
		t, err := expandSystemHaUnicastSha(d, v, "unicast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unicast"] = t
		}
	}

	if v, ok := d.GetOk("vip"); ok || d.HasChange("vip") {
		t, err := expandSystemHaVipSha(d, v, "vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip"] = t
		}
	}

	if v, ok := d.GetOk("vip_interface"); ok || d.HasChange("vip_interface") {
		t, err := expandSystemHaVipInterfaceSha(d, v, "vip_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip-interface"] = t
		}
	}

	if v, ok := d.GetOk("vrrp_adv_interval"); ok || d.HasChange("vrrp_adv_interval") {
		t, err := expandSystemHaVrrpAdvIntervalSha(d, v, "vrrp_adv_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrrp-adv-interval"] = t
		}
	}

	if v, ok := d.GetOk("vrrp_interface"); ok || d.HasChange("vrrp_interface") {
		t, err := expandSystemHaVrrpInterfaceSha(d, v, "vrrp_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrrp-interface"] = t
		}
	}

	return &obj, nil
}
