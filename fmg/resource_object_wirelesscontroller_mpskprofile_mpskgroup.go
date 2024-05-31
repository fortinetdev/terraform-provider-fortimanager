// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: List of multiple PSK groups.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerMpskProfileMpskGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerMpskProfileMpskGroupCreate,
		Read:   resourceObjectWirelessControllerMpskProfileMpskGroupRead,
		Update: resourceObjectWirelessControllerMpskProfileMpskGroupUpdate,
		Delete: resourceObjectWirelessControllerMpskProfileMpskGroupDelete,

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
			"mpsk_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
							Computed: true,
						},
						"concurrent_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"pmk": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vlan_type": &schema.Schema{
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

func resourceObjectWirelessControllerMpskProfileMpskGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	mpsk_profile := d.Get("mpsk_profile").(string)
	paradict["mpsk_profile"] = mpsk_profile

	obj, err := getObjectObjectWirelessControllerMpskProfileMpskGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerMpskProfileMpskGroup resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerMpskProfileMpskGroup(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerMpskProfileMpskGroup resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerMpskProfileMpskGroupRead(d, m)
}

func resourceObjectWirelessControllerMpskProfileMpskGroupUpdate(d *schema.ResourceData, m interface{}) error {
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

	mpsk_profile := d.Get("mpsk_profile").(string)
	paradict["mpsk_profile"] = mpsk_profile

	obj, err := getObjectObjectWirelessControllerMpskProfileMpskGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerMpskProfileMpskGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerMpskProfileMpskGroup(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerMpskProfileMpskGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerMpskProfileMpskGroupRead(d, m)
}

func resourceObjectWirelessControllerMpskProfileMpskGroupDelete(d *schema.ResourceData, m interface{}) error {
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

	mpsk_profile := d.Get("mpsk_profile").(string)
	paradict["mpsk_profile"] = mpsk_profile

	err = c.DeleteObjectWirelessControllerMpskProfileMpskGroup(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerMpskProfileMpskGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerMpskProfileMpskGroupRead(d *schema.ResourceData, m interface{}) error {
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

	mpsk_profile := d.Get("mpsk_profile").(string)
	if mpsk_profile == "" {
		mpsk_profile = importOptionChecking(m.(*FortiClient).Cfg, "mpsk_profile")
		if mpsk_profile == "" {
			return fmt.Errorf("Parameter mpsk_profile is missing")
		}
		if err = d.Set("mpsk_profile", mpsk_profile); err != nil {
			return fmt.Errorf("Error set params mpsk_profile: %v", err)
		}
	}
	paradict["mpsk_profile"] = mpsk_profile

	o, err := c.ReadObjectWirelessControllerMpskProfileMpskGroup(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerMpskProfileMpskGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerMpskProfileMpskGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerMpskProfileMpskGroup resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKey2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyComment2edl(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_client_limit_type"
		if _, ok := i["concurrent-client-limit-type"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType2edl(i["concurrent-client-limit-type"], d, pre_append)
			tmp["concurrent_client_limit_type"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey-ConcurrentClientLimitType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := i["concurrent-clients"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients2edl(i["concurrent-clients"], d, pre_append)
			tmp["concurrent_clients"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey-ConcurrentClients")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyMac2edl(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if _, ok := i["mpsk-schedules"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules2edl(i["mpsk-schedules"], d, pre_append)
			tmp["mpsk_schedules"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey-MpskSchedules")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupVlanId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupVlanType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerMpskProfileMpskGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("mpsk_key", flattenObjectWirelessControllerMpskProfileMpskGroupMpskKey2edl(o["mpsk-key"], d, "mpsk_key")); err != nil {
			if vv, ok := fortiAPIPatch(o["mpsk-key"], "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey"); ok {
				if err = d.Set("mpsk_key", vv); err != nil {
					return fmt.Errorf("Error reading mpsk_key: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mpsk_key: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mpsk_key"); ok {
			if err = d.Set("mpsk_key", flattenObjectWirelessControllerMpskProfileMpskGroupMpskKey2edl(o["mpsk-key"], d, "mpsk_key")); err != nil {
				if vv, ok := fortiAPIPatch(o["mpsk-key"], "ObjectWirelessControllerMpskProfileMpskGroup-MpskKey"); ok {
					if err = d.Set("mpsk_key", vv); err != nil {
						return fmt.Errorf("Error reading mpsk_key: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mpsk_key: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerMpskProfileMpskGroupName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerMpskProfileMpskGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vlan_id", flattenObjectWirelessControllerMpskProfileMpskGroupVlanId2edl(o["vlan-id"], d, "vlan_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-id"], "ObjectWirelessControllerMpskProfileMpskGroup-VlanId"); ok {
			if err = d.Set("vlan_id", vv); err != nil {
				return fmt.Errorf("Error reading vlan_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_id: %v", err)
		}
	}

	if err = d.Set("vlan_type", flattenObjectWirelessControllerMpskProfileMpskGroupVlanType2edl(o["vlan-type"], d, "vlan_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-type"], "ObjectWirelessControllerMpskProfileMpskGroup-VlanType"); ok {
			if err = d.Set("vlan_type", vv); err != nil {
				return fmt.Errorf("Error reading vlan_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_type: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerMpskProfileMpskGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyComment2edl(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_client_limit_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["concurrent-client-limit-type"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType2edl(d, i["concurrent_client_limit_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["concurrent-clients"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients2edl(d, i["concurrent_clients"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyMac2edl(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mpsk-schedules"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules2edl(d, i["mpsk_schedules"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passphrase"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase2edl(d, i["passphrase"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmk"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pmk"], _ = expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyPmk2edl(d, i["pmk"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyPmk2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupVlanId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupVlanType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerMpskProfileMpskGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mpsk_key"); ok || d.HasChange("mpsk_key") {
		t, err := expandObjectWirelessControllerMpskProfileMpskGroupMpskKey2edl(d, v, "mpsk_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-key"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerMpskProfileMpskGroupName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vlan_id"); ok || d.HasChange("vlan_id") {
		t, err := expandObjectWirelessControllerMpskProfileMpskGroupVlanId2edl(d, v, "vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-id"] = t
		}
	}

	if v, ok := d.GetOk("vlan_type"); ok || d.HasChange("vlan_type") {
		t, err := expandObjectWirelessControllerMpskProfileMpskGroupVlanType2edl(d, v, "vlan_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-type"] = t
		}
	}

	return &obj, nil
}
