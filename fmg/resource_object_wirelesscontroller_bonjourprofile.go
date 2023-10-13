// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Bonjour profiles. Bonjour is Apple's zero configuration networking protocol. Bonjour profiles allow APs and FortiAPs to connnect to networks using Bonjour.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerBonjourProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerBonjourProfileCreate,
		Read:   resourceObjectWirelessControllerBonjourProfileRead,
		Update: resourceObjectWirelessControllerBonjourProfileUpdate,
		Delete: resourceObjectWirelessControllerBonjourProfileDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"policy_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"from_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"policy_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"services": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"to_vlan": &schema.Schema{
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

func resourceObjectWirelessControllerBonjourProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectWirelessControllerBonjourProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerBonjourProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerBonjourProfile(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerBonjourProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerBonjourProfileRead(d, m)
}

func resourceObjectWirelessControllerBonjourProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerBonjourProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerBonjourProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerBonjourProfile(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerBonjourProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerBonjourProfileRead(d, m)
}

func resourceObjectWirelessControllerBonjourProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWirelessControllerBonjourProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerBonjourProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerBonjourProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWirelessControllerBonjourProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerBonjourProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerBonjourProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerBonjourProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerBonjourProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBonjourProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBonjourProfilePolicyList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenObjectWirelessControllerBonjourProfilePolicyListDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerBonjourProfile-PolicyList-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from_vlan"
		if _, ok := i["from-vlan"]; ok {
			v := flattenObjectWirelessControllerBonjourProfilePolicyListFromVlan(i["from-vlan"], d, pre_append)
			tmp["from_vlan"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerBonjourProfile-PolicyList-FromVlan")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy_id"
		if _, ok := i["policy-id"]; ok {
			v := flattenObjectWirelessControllerBonjourProfilePolicyListPolicyId(i["policy-id"], d, pre_append)
			tmp["policy_id"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerBonjourProfile-PolicyList-PolicyId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "services"
		if _, ok := i["services"]; ok {
			v := flattenObjectWirelessControllerBonjourProfilePolicyListServices(i["services"], d, pre_append)
			tmp["services"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerBonjourProfile-PolicyList-Services")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "to_vlan"
		if _, ok := i["to-vlan"]; ok {
			v := flattenObjectWirelessControllerBonjourProfilePolicyListToVlan(i["to-vlan"], d, pre_append)
			tmp["to_vlan"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerBonjourProfile-PolicyList-ToVlan")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerBonjourProfilePolicyListDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBonjourProfilePolicyListFromVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBonjourProfilePolicyListPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBonjourProfilePolicyListServices(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerBonjourProfilePolicyListToVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerBonjourProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenObjectWirelessControllerBonjourProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWirelessControllerBonjourProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerBonjourProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerBonjourProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("policy_list", flattenObjectWirelessControllerBonjourProfilePolicyList(o["policy-list"], d, "policy_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["policy-list"], "ObjectWirelessControllerBonjourProfile-PolicyList"); ok {
				if err = d.Set("policy_list", vv); err != nil {
					return fmt.Errorf("Error reading policy_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading policy_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("policy_list"); ok {
			if err = d.Set("policy_list", flattenObjectWirelessControllerBonjourProfilePolicyList(o["policy-list"], d, "policy_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["policy-list"], "ObjectWirelessControllerBonjourProfile-PolicyList"); ok {
					if err = d.Set("policy_list", vv); err != nil {
						return fmt.Errorf("Error reading policy_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading policy_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectWirelessControllerBonjourProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerBonjourProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBonjourProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBonjourProfilePolicyList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandObjectWirelessControllerBonjourProfilePolicyListDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from_vlan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["from-vlan"], _ = expandObjectWirelessControllerBonjourProfilePolicyListFromVlan(d, i["from_vlan"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["policy-id"], _ = expandObjectWirelessControllerBonjourProfilePolicyListPolicyId(d, i["policy_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "services"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["services"], _ = expandObjectWirelessControllerBonjourProfilePolicyListServices(d, i["services"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "to_vlan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["to-vlan"], _ = expandObjectWirelessControllerBonjourProfilePolicyListToVlan(d, i["to_vlan"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerBonjourProfilePolicyListDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBonjourProfilePolicyListFromVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBonjourProfilePolicyListPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBonjourProfilePolicyListServices(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerBonjourProfilePolicyListToVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerBonjourProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectWirelessControllerBonjourProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerBonjourProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("policy_list"); ok || d.HasChange("policy_list") {
		t, err := expandObjectWirelessControllerBonjourProfilePolicyList(d, v, "policy_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-list"] = t
		}
	}

	return &obj, nil
}
