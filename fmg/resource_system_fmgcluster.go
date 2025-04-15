// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: fmg clsuter.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFmgCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFmgClusterUpdate,
		Read:   resourceSystemFmgClusterRead,
		Update: resourceSystemFmgClusterUpdate,
		Delete: resourceSystemFmgClusterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

func resourceSystemFmgClusterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemFmgCluster(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFmgCluster resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemFmgCluster(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemFmgCluster resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemFmgCluster")

	return resourceSystemFmgClusterRead(d, m)
}

func resourceSystemFmgClusterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemFmgCluster(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFmgCluster resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFmgClusterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemFmgCluster(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemFmgCluster resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFmgCluster(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFmgCluster resource from API: %v", err)
	}
	return nil
}

func flattenSystemFmgClusterFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFmgClusterIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFmgClusterMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFmgClusterPeer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr"
		if _, ok := i["addr"]; ok {
			v := flattenSystemFmgClusterPeerAddr(i["addr"], d, pre_append)
			tmp["addr"] = fortiAPISubPartPatch(v, "SystemFmgCluster-Peer-Addr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := i["fqdn"]; ok {
			v := flattenSystemFmgClusterPeerFqdn(i["fqdn"], d, pre_append)
			tmp["fqdn"] = fortiAPISubPartPatch(v, "SystemFmgCluster-Peer-Fqdn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemFmgClusterPeerName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemFmgCluster-Peer-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sn"
		if _, ok := i["sn"]; ok {
			v := flattenSystemFmgClusterPeerSn(i["sn"], d, pre_append)
			tmp["sn"] = fortiAPISubPartPatch(v, "SystemFmgCluster-Peer-Sn")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemFmgClusterPeerAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFmgClusterPeerFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFmgClusterPeerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFmgClusterPeerSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemFmgCluster(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("fqdn", flattenSystemFmgClusterFqdn(o["fqdn"], d, "fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn"], "SystemFmgCluster-Fqdn"); ok {
			if err = d.Set("fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemFmgClusterIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SystemFmgCluster-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemFmgClusterMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemFmgCluster-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("peer", flattenSystemFmgClusterPeer(o["peer"], d, "peer")); err != nil {
			if vv, ok := fortiAPIPatch(o["peer"], "SystemFmgCluster-Peer"); ok {
				if err = d.Set("peer", vv); err != nil {
					return fmt.Errorf("Error reading peer: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading peer: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("peer"); ok {
			if err = d.Set("peer", flattenSystemFmgClusterPeer(o["peer"], d, "peer")); err != nil {
				if vv, ok := fortiAPIPatch(o["peer"], "SystemFmgCluster-Peer"); ok {
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

func flattenSystemFmgClusterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFmgClusterFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFmgClusterIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFmgClusterMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFmgClusterPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr"], _ = expandSystemFmgClusterPeerAddr(d, i["addr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fqdn"], _ = expandSystemFmgClusterPeerFqdn(d, i["fqdn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemFmgClusterPeerName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sn"], _ = expandSystemFmgClusterPeerSn(d, i["sn"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemFmgClusterPeerAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFmgClusterPeerFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFmgClusterPeerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFmgClusterPeerSn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFmgCluster(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fqdn"); ok || d.HasChange("fqdn") {
		t, err := expandSystemFmgClusterFqdn(d, v, "fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSystemFmgClusterIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemFmgClusterMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("peer"); ok || d.HasChange("peer") {
		t, err := expandSystemFmgClusterPeer(d, v, "peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	}

	return &obj, nil
}
