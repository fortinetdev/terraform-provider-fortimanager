// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure MPSK profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWirelessControllerMpskProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerMpskProfileCreate,
		Read:   resourceObjectWirelessControllerMpskProfileRead,
		Update: resourceObjectWirelessControllerMpskProfileUpdate,
		Delete: resourceObjectWirelessControllerMpskProfileDelete,

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
			"mpsk_concurrent_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mpsk_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mpsk_key": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"comment": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"concurrent_client_limit_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"concurrent_clients": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"mac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"mpsk_schedules": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"passphrase": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
									},
									"pmk": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vlan_type": &schema.Schema{
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
			"ssid": &schema.Schema{
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

func resourceObjectWirelessControllerMpskProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerMpskProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerMpskProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerMpskProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerMpskProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerMpskProfileRead(d, m)
}

func resourceObjectWirelessControllerMpskProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerMpskProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerMpskProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerMpskProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerMpskProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerMpskProfileRead(d, m)
}

func resourceObjectWirelessControllerMpskProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerMpskProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerMpskProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerMpskProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerMpskProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerMpskProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerMpskProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerMpskProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerMpskProfileMpskConcurrentClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_key"
		if _, ok := i["mpsk-key"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupMpskKey(i["mpsk-key"], d, pre_append)
			tmp["mpsk_key"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfile-MpskGroup-MpskKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfile-MpskGroup-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_id"
		if _, ok := i["vlan-id"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupVlanId(i["vlan-id"], d, pre_append)
			tmp["vlan_id"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfile-MpskGroup-VlanId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_type"
		if _, ok := i["vlan-type"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupVlanType(i["vlan-type"], d, pre_append)
			tmp["vlan_type"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfile-MpskGroup-VlanType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKey(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_client_limit_type"
		if _, ok := i["concurrent-client-limit-type"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType(i["concurrent-client-limit-type"], d, pre_append)
			tmp["concurrent_client_limit_type"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey-ConcurrentClientLimitType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := i["concurrent-clients"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients(i["concurrent-clients"], d, pre_append)
			tmp["concurrent_clients"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey-ConcurrentClients")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if _, ok := i["mpsk-schedules"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(i["mpsk-schedules"], d, pre_append)
			tmp["mpsk_schedules"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey-MpskSchedules")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := i["passphrase"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase(i["passphrase"], d, pre_append)
			tmp["passphrase"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey-Passphrase")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmk"
		if _, ok := i["pmk"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyPmk(i["pmk"], d, pre_append)
			tmp["pmk"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey-Pmk")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyPmk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerMpskProfileMpskGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupVlanType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerMpskProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("mpsk_concurrent_clients", flattenObjectWirelessControllerMpskProfileMpskConcurrentClients(o["mpsk-concurrent-clients"], d, "mpsk_concurrent_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk-concurrent-clients"], "ObjectWirelessControllerMpskProfile-MpskConcurrentClients"); ok {
			if err = d.Set("mpsk_concurrent_clients", vv); err != nil {
				return fmt.Errorf("Error reading mpsk_concurrent_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk_concurrent_clients: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mpsk_group", flattenObjectWirelessControllerMpskProfileMpskGroup(o["mpsk-group"], d, "mpsk_group")); err != nil {
			if vv, ok := fortiAPIPatch(o["mpsk-group"], "ObjectWirelessControllerMpskProfile-MpskGroup"); ok {
				if err = d.Set("mpsk_group", vv); err != nil {
					return fmt.Errorf("Error reading mpsk_group: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mpsk_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mpsk_group"); ok {
			if err = d.Set("mpsk_group", flattenObjectWirelessControllerMpskProfileMpskGroup(o["mpsk-group"], d, "mpsk_group")); err != nil {
				if vv, ok := fortiAPIPatch(o["mpsk-group"], "ObjectWirelessControllerMpskProfile-MpskGroup"); ok {
					if err = d.Set("mpsk_group", vv); err != nil {
						return fmt.Errorf("Error reading mpsk_group: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mpsk_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerMpskProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerMpskProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ssid", flattenObjectWirelessControllerMpskProfileSsid(o["ssid"], d, "ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssid"], "ObjectWirelessControllerMpskProfile-Ssid"); ok {
			if err = d.Set("ssid", vv); err != nil {
				return fmt.Errorf("Error reading ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssid: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerMpskProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerMpskProfileMpskConcurrentClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mpsk-key"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKey(d, i["mpsk_key"], pre_append)
		} else {
			tmp["mpsk-key"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectWirelessControllerMpskProfileMpskGroupName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-id"], _ = expandObjectWirelessControllerMpskProfileMpskGroupVlanId(d, i["vlan_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-type"], _ = expandObjectWirelessControllerMpskProfileMpskGroupVlanType(d, i["vlan_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_client_limit_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["concurrent-client-limit-type"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType(d, i["concurrent_client_limit_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["concurrent-clients"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients(d, i["concurrent_clients"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mpsk-schedules"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(d, i["mpsk_schedules"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passphrase"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase(d, i["passphrase"], pre_append)
		} else {
			tmp["passphrase"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmk"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pmk"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyPmk(d, i["pmk"], pre_append)
		} else {
			tmp["pmk"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyPmk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupVlanType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerMpskProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mpsk_concurrent_clients"); ok || d.HasChange("mpsk_concurrent_clients") {
		t, err := expandObjectWirelessControllerMpskProfileMpskConcurrentClients(d, v, "mpsk_concurrent_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-concurrent-clients"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_group"); ok || d.HasChange("mpsk_group") {
		t, err := expandObjectWirelessControllerMpskProfileMpskGroup(d, v, "mpsk_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-group"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerMpskProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ssid"); ok || d.HasChange("ssid") {
		t, err := expandObjectWirelessControllerMpskProfileSsid(d, v, "ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssid"] = t
		}
	}

	return &obj, nil
}
