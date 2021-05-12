// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: HA configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"peer": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
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

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemHa(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHa(obj, adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemHa(adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemHa(adomv, mkey, nil)
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
	if v != nil {
		emap := map[int]string{
			0: "standalone",
			1: "primary",
			2: "secondary",
		}
		res := getEnumVal(v, emap)
		return res
	}
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
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectSystemHa(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("clusterid", flattenSystemHaClusteridSha(o["clusterid"], d, "clusterid")); err != nil {
		if vv, ok := fortiAPIPatch(o["clusterid"], "SystemHa-Clusterid"); ok {
			if err = d.Set("clusterid", vv); err != nil {
				return fmt.Errorf("Error reading clusterid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clusterid: %v", err)
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

	if err = d.Set("password", flattenSystemHaPasswordSha(o["password"], d, "password")); err != nil {
		if vv, ok := fortiAPIPatch(o["password"], "SystemHa-Password"); ok {
			if err = d.Set("password", vv); err != nil {
				return fmt.Errorf("Error reading password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password: %v", err)
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

func expandSystemHaPasswordSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaPeerSha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSystemHaPeerIdSha(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandSystemHaPeerIpSha(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip6"], _ = expandSystemHaPeerIp6Sha(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial_number"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["serial-number"], _ = expandSystemHaPeerSerialNumberSha(d, i["serial_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
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

func getObjectSystemHa(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("clusterid"); ok {
		t, err := expandSystemHaClusteridSha(d, v, "clusterid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clusterid"] = t
		}
	}

	if v, ok := d.GetOk("file_quota"); ok {
		t, err := expandSystemHaFileQuotaSha(d, v, "file_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-quota"] = t
		}
	}

	if v, ok := d.GetOk("hb_interval"); ok {
		t, err := expandSystemHaHbIntervalSha(d, v, "hb_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-interval"] = t
		}
	}

	if v, ok := d.GetOk("hb_lost_threshold"); ok {
		t, err := expandSystemHaHbLostThresholdSha(d, v, "hb_lost_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-lost-threshold"] = t
		}
	}

	if v, ok := d.GetOk("local_cert"); ok {
		t, err := expandSystemHaLocalCertSha(d, v, "local_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-cert"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSystemHaModeSha(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemHaPasswordSha(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("peer"); ok {
		t, err := expandSystemHaPeerSha(d, v, "peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	}

	return &obj, nil
}
