// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SNMP community configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSnmpCommunity() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnmpCommunityCreate,
		Read:   resourceSystemSnmpCommunityRead,
		Update: resourceSystemSnmpCommunityUpdate,
		Delete: resourceSystemSnmpCommunityDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"events": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"hosts": &schema.Schema{
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
							Computed: true,
						},
					},
				},
			},
			"hosts6": &schema.Schema{
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
							Computed: true,
						},
					},
				},
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"query_v1_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"query_v1_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_v2c_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"query_v2c_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap_v1_rport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_v1_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap_v2c_rport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_v2c_status": &schema.Schema{
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

func resourceSystemSnmpCommunityCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSnmpCommunity(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpCommunity resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSnmpCommunity(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpCommunity resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSnmpCommunityRead(d, m)
}

func resourceSystemSnmpCommunityUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSnmpCommunity(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpCommunity resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSnmpCommunity(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpCommunity resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSnmpCommunityRead(d, m)
}

func resourceSystemSnmpCommunityDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemSnmpCommunity(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpCommunity resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpCommunityRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemSnmpCommunity(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpCommunity resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpCommunity(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpCommunity resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnmpCommunityEvents(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSnmpCommunityHosts(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemSnmpCommunityHostsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenSystemSnmpCommunityHostsInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemSnmpCommunityHostsIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts-Ip")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSnmpCommunityHostsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHostsInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHostsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenSystemSnmpCommunityHosts6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemSnmpCommunityHosts6Id(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts6-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenSystemSnmpCommunityHosts6Interface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts6-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemSnmpCommunityHosts6Ip(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts6-Ip")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSnmpCommunityHosts6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6Ip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityQueryV1Port(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityQueryV1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityQueryV2CPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityQueryV2CStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityTrapV1Rport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityTrapV1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityTrapV2CRport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityTrapV2CStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSnmpCommunity(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("events", flattenSystemSnmpCommunityEvents(o["events"], d, "events")); err != nil {
		if vv, ok := fortiAPIPatch(o["events"], "SystemSnmpCommunity-Events"); ok {
			if err = d.Set("events", vv); err != nil {
				return fmt.Errorf("Error reading events: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading events: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("hosts", flattenSystemSnmpCommunityHosts(o["hosts"], d, "hosts")); err != nil {
			if vv, ok := fortiAPIPatch(o["hosts"], "SystemSnmpCommunity-Hosts"); ok {
				if err = d.Set("hosts", vv); err != nil {
					return fmt.Errorf("Error reading hosts: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading hosts: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("hosts"); ok {
			if err = d.Set("hosts", flattenSystemSnmpCommunityHosts(o["hosts"], d, "hosts")); err != nil {
				if vv, ok := fortiAPIPatch(o["hosts"], "SystemSnmpCommunity-Hosts"); ok {
					if err = d.Set("hosts", vv); err != nil {
						return fmt.Errorf("Error reading hosts: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading hosts: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("hosts6", flattenSystemSnmpCommunityHosts6(o["hosts6"], d, "hosts6")); err != nil {
			if vv, ok := fortiAPIPatch(o["hosts6"], "SystemSnmpCommunity-Hosts6"); ok {
				if err = d.Set("hosts6", vv); err != nil {
					return fmt.Errorf("Error reading hosts6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading hosts6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("hosts6"); ok {
			if err = d.Set("hosts6", flattenSystemSnmpCommunityHosts6(o["hosts6"], d, "hosts6")); err != nil {
				if vv, ok := fortiAPIPatch(o["hosts6"], "SystemSnmpCommunity-Hosts6"); ok {
					if err = d.Set("hosts6", vv); err != nil {
						return fmt.Errorf("Error reading hosts6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading hosts6: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenSystemSnmpCommunityId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemSnmpCommunity-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSnmpCommunityName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSnmpCommunity-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("query_v1_port", flattenSystemSnmpCommunityQueryV1Port(o["query_v1_port"], d, "query_v1_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["query_v1_port"], "SystemSnmpCommunity-QueryV1Port"); ok {
			if err = d.Set("query_v1_port", vv); err != nil {
				return fmt.Errorf("Error reading query_v1_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v1_port: %v", err)
		}
	}

	if err = d.Set("query_v1_status", flattenSystemSnmpCommunityQueryV1Status(o["query_v1_status"], d, "query_v1_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["query_v1_status"], "SystemSnmpCommunity-QueryV1Status"); ok {
			if err = d.Set("query_v1_status", vv); err != nil {
				return fmt.Errorf("Error reading query_v1_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v1_status: %v", err)
		}
	}

	if err = d.Set("query_v2c_port", flattenSystemSnmpCommunityQueryV2CPort(o["query_v2c_port"], d, "query_v2c_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["query_v2c_port"], "SystemSnmpCommunity-QueryV2CPort"); ok {
			if err = d.Set("query_v2c_port", vv); err != nil {
				return fmt.Errorf("Error reading query_v2c_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v2c_port: %v", err)
		}
	}

	if err = d.Set("query_v2c_status", flattenSystemSnmpCommunityQueryV2CStatus(o["query_v2c_status"], d, "query_v2c_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["query_v2c_status"], "SystemSnmpCommunity-QueryV2CStatus"); ok {
			if err = d.Set("query_v2c_status", vv); err != nil {
				return fmt.Errorf("Error reading query_v2c_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v2c_status: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSnmpCommunityStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSnmpCommunity-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trap_v1_rport", flattenSystemSnmpCommunityTrapV1Rport(o["trap_v1_rport"], d, "trap_v1_rport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap_v1_rport"], "SystemSnmpCommunity-TrapV1Rport"); ok {
			if err = d.Set("trap_v1_rport", vv); err != nil {
				return fmt.Errorf("Error reading trap_v1_rport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v1_rport: %v", err)
		}
	}

	if err = d.Set("trap_v1_status", flattenSystemSnmpCommunityTrapV1Status(o["trap_v1_status"], d, "trap_v1_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap_v1_status"], "SystemSnmpCommunity-TrapV1Status"); ok {
			if err = d.Set("trap_v1_status", vv); err != nil {
				return fmt.Errorf("Error reading trap_v1_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v1_status: %v", err)
		}
	}

	if err = d.Set("trap_v2c_rport", flattenSystemSnmpCommunityTrapV2CRport(o["trap_v2c_rport"], d, "trap_v2c_rport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap_v2c_rport"], "SystemSnmpCommunity-TrapV2CRport"); ok {
			if err = d.Set("trap_v2c_rport", vv); err != nil {
				return fmt.Errorf("Error reading trap_v2c_rport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v2c_rport: %v", err)
		}
	}

	if err = d.Set("trap_v2c_status", flattenSystemSnmpCommunityTrapV2CStatus(o["trap_v2c_status"], d, "trap_v2c_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap_v2c_status"], "SystemSnmpCommunity-TrapV2CStatus"); ok {
			if err = d.Set("trap_v2c_status", vv); err != nil {
				return fmt.Errorf("Error reading trap_v2c_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v2c_status: %v", err)
		}
	}

	return nil
}

func flattenSystemSnmpCommunityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSnmpCommunityEvents(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSnmpCommunityHosts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemSnmpCommunityHostsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandSystemSnmpCommunityHostsInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSystemSnmpCommunityHostsIp(d, i["ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSnmpCommunityHostsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHostsInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHostsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandSystemSnmpCommunityHosts6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemSnmpCommunityHosts6Id(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandSystemSnmpCommunityHosts6Interface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSystemSnmpCommunityHosts6Ip(d, i["ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSnmpCommunityHosts6Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6Ip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityQueryV1Port(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityQueryV1Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityQueryV2CPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityQueryV2CStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV1Rport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV1Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV2CRport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV2CStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpCommunity(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("events"); ok || d.HasChange("events") {
		t, err := expandSystemSnmpCommunityEvents(d, v, "events")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["events"] = t
		}
	}

	if v, ok := d.GetOk("hosts"); ok || d.HasChange("hosts") {
		t, err := expandSystemSnmpCommunityHosts(d, v, "hosts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hosts"] = t
		}
	}

	if v, ok := d.GetOk("hosts6"); ok || d.HasChange("hosts6") {
		t, err := expandSystemSnmpCommunityHosts6(d, v, "hosts6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hosts6"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemSnmpCommunityId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSnmpCommunityName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("query_v1_port"); ok || d.HasChange("query_v1_port") {
		t, err := expandSystemSnmpCommunityQueryV1Port(d, v, "query_v1_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query_v1_port"] = t
		}
	}

	if v, ok := d.GetOk("query_v1_status"); ok || d.HasChange("query_v1_status") {
		t, err := expandSystemSnmpCommunityQueryV1Status(d, v, "query_v1_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query_v1_status"] = t
		}
	}

	if v, ok := d.GetOk("query_v2c_port"); ok || d.HasChange("query_v2c_port") {
		t, err := expandSystemSnmpCommunityQueryV2CPort(d, v, "query_v2c_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query_v2c_port"] = t
		}
	}

	if v, ok := d.GetOk("query_v2c_status"); ok || d.HasChange("query_v2c_status") {
		t, err := expandSystemSnmpCommunityQueryV2CStatus(d, v, "query_v2c_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query_v2c_status"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSnmpCommunityStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trap_v1_rport"); ok || d.HasChange("trap_v1_rport") {
		t, err := expandSystemSnmpCommunityTrapV1Rport(d, v, "trap_v1_rport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap_v1_rport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v1_status"); ok || d.HasChange("trap_v1_status") {
		t, err := expandSystemSnmpCommunityTrapV1Status(d, v, "trap_v1_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap_v1_status"] = t
		}
	}

	if v, ok := d.GetOk("trap_v2c_rport"); ok || d.HasChange("trap_v2c_rport") {
		t, err := expandSystemSnmpCommunityTrapV2CRport(d, v, "trap_v2c_rport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap_v2c_rport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v2c_status"); ok || d.HasChange("trap_v2c_status") {
		t, err := expandSystemSnmpCommunityTrapV2CStatus(d, v, "trap_v2c_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap_v2c_status"] = t
		}
	}

	return &obj, nil
}
