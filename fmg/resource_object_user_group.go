// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure user groups.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectUserGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserGroupCreate,
		Read:   resourceObjectUserGroupRead,
		Update: resourceObjectUserGroupUpdate,
		Delete: resourceObjectUserGroupDelete,

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
			"auth_concurrent_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_concurrent_value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"authtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"company": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expire": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"expire_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guest": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"company": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expiration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mobile_phone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
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
						"sponsor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"user_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"http_digest_realm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_gui_meta": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"group_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"server_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"max_accounts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mobile_phone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multiple_guest_add": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sms_custom_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sms_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sponsor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso_attribute_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_name": &schema.Schema{
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

func resourceObjectUserGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserGroup resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserGroup(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserGroup resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserGroupRead(d, m)
}

func resourceObjectUserGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserGroup(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserGroupRead(d, m)
}

func resourceObjectUserGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectUserGroup(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectUserGroup(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserGroup resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserGroupAuthConcurrentOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectUserGroupAuthConcurrentValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupAuthtimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupCompany(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "optional",
			1: "mandatory",
			2: "disabled",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserGroupEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectUserGroupExpire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupExpireType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "immediately",
			1: "first-successful-login",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserGroupGroupType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "firewall",
			6: "fsso-service",
			7: "guest",
			8: "rsso",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserGroupGuest(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserGroupGuestComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectUserGroup-Guest-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := i["company"]; ok {
			v := flattenObjectUserGroupGuestCompany(i["company"], d, pre_append)
			tmp["company"] = fortiAPISubPartPatch(v, "ObjectUserGroup-Guest-Company")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := i["email"]; ok {
			v := flattenObjectUserGroupGuestEmail(i["email"], d, pre_append)
			tmp["email"] = fortiAPISubPartPatch(v, "ObjectUserGroup-Guest-Email")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiration"
		if _, ok := i["expiration"]; ok {
			v := flattenObjectUserGroupGuestExpiration(i["expiration"], d, pre_append)
			tmp["expiration"] = fortiAPISubPartPatch(v, "ObjectUserGroup-Guest-Expiration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectUserGroupGuestId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectUserGroup-Guest-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := i["mobile-phone"]; ok {
			v := flattenObjectUserGroupGuestMobilePhone(i["mobile-phone"], d, pre_append)
			tmp["mobile_phone"] = fortiAPISubPartPatch(v, "ObjectUserGroup-Guest-MobilePhone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectUserGroupGuestName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserGroup-Guest-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			v := flattenObjectUserGroupGuestPassword(i["password"], d, pre_append)
			tmp["password"] = fortiAPISubPartPatch(v, "ObjectUserGroup-Guest-Password")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := i["sponsor"]; ok {
			v := flattenObjectUserGroupGuestSponsor(i["sponsor"], d, pre_append)
			tmp["sponsor"] = fortiAPISubPartPatch(v, "ObjectUserGroup-Guest-Sponsor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := i["user-id"]; ok {
			v := flattenObjectUserGroupGuestUserId(i["user-id"], d, pre_append)
			tmp["user_id"] = fortiAPISubPartPatch(v, "ObjectUserGroup-Guest-UserId")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectUserGroupGuestComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupGuestCompany(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupGuestEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupGuestExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupGuestId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupGuestMobilePhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupGuestName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupGuestPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserGroupGuestSponsor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupGuestUserId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupHttpDigestRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupMatch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_gui_meta"
		if _, ok := i["_gui_meta"]; ok {
			v := flattenObjectUserGroupMatchGuiMeta(i["_gui_meta"], d, pre_append)
			tmp["_gui_meta"] = fortiAPISubPartPatch(v, "ObjectUserGroup-Match-GuiMeta")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := i["group-name"]; ok {
			v := flattenObjectUserGroupMatchGroupName(i["group-name"], d, pre_append)
			tmp["group_name"] = fortiAPISubPartPatch(v, "ObjectUserGroup-Match-GroupName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectUserGroupMatchId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectUserGroup-Match-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_name"
		if _, ok := i["server-name"]; ok {
			v := flattenObjectUserGroupMatchServerName(i["server-name"], d, pre_append)
			tmp["server_name"] = fortiAPISubPartPatch(v, "ObjectUserGroup-Match-ServerName")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectUserGroupMatchGuiMeta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupMatchGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupMatchId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupMatchServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupMaxAccounts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupMobilePhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectUserGroupMultipleGuestAdd(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectUserGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "auto-generate",
			2: "specify",
			3: "disable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserGroupSmsCustomServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupSmsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "fortiguard",
			1: "custom",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserGroupSponsor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "optional",
			1: "mandatory",
			2: "disabled",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserGroupSsoAttributeValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupUserId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "email",
			1: "auto-generate",
			2: "specify",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserGroupUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func refreshObjectObjectUserGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_concurrent_override", flattenObjectUserGroupAuthConcurrentOverride(o["auth-concurrent-override"], d, "auth_concurrent_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-concurrent-override"], "ObjectUserGroup-AuthConcurrentOverride"); ok {
			if err = d.Set("auth_concurrent_override", vv); err != nil {
				return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
		}
	}

	if err = d.Set("auth_concurrent_value", flattenObjectUserGroupAuthConcurrentValue(o["auth-concurrent-value"], d, "auth_concurrent_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-concurrent-value"], "ObjectUserGroup-AuthConcurrentValue"); ok {
			if err = d.Set("auth_concurrent_value", vv); err != nil {
				return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
		}
	}

	if err = d.Set("authtimeout", flattenObjectUserGroupAuthtimeout(o["authtimeout"], d, "authtimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["authtimeout"], "ObjectUserGroup-Authtimeout"); ok {
			if err = d.Set("authtimeout", vv); err != nil {
				return fmt.Errorf("Error reading authtimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authtimeout: %v", err)
		}
	}

	if err = d.Set("company", flattenObjectUserGroupCompany(o["company"], d, "company")); err != nil {
		if vv, ok := fortiAPIPatch(o["company"], "ObjectUserGroup-Company"); ok {
			if err = d.Set("company", vv); err != nil {
				return fmt.Errorf("Error reading company: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading company: %v", err)
		}
	}

	if err = d.Set("email", flattenObjectUserGroupEmail(o["email"], d, "email")); err != nil {
		if vv, ok := fortiAPIPatch(o["email"], "ObjectUserGroup-Email"); ok {
			if err = d.Set("email", vv); err != nil {
				return fmt.Errorf("Error reading email: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email: %v", err)
		}
	}

	if err = d.Set("expire", flattenObjectUserGroupExpire(o["expire"], d, "expire")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire"], "ObjectUserGroup-Expire"); ok {
			if err = d.Set("expire", vv); err != nil {
				return fmt.Errorf("Error reading expire: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire: %v", err)
		}
	}

	if err = d.Set("expire_type", flattenObjectUserGroupExpireType(o["expire-type"], d, "expire_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire-type"], "ObjectUserGroup-ExpireType"); ok {
			if err = d.Set("expire_type", vv); err != nil {
				return fmt.Errorf("Error reading expire_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire_type: %v", err)
		}
	}

	if err = d.Set("group_type", flattenObjectUserGroupGroupType(o["group-type"], d, "group_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-type"], "ObjectUserGroup-GroupType"); ok {
			if err = d.Set("group_type", vv); err != nil {
				return fmt.Errorf("Error reading group_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("guest", flattenObjectUserGroupGuest(o["guest"], d, "guest")); err != nil {
			if vv, ok := fortiAPIPatch(o["guest"], "ObjectUserGroup-Guest"); ok {
				if err = d.Set("guest", vv); err != nil {
					return fmt.Errorf("Error reading guest: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading guest: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("guest"); ok {
			if err = d.Set("guest", flattenObjectUserGroupGuest(o["guest"], d, "guest")); err != nil {
				if vv, ok := fortiAPIPatch(o["guest"], "ObjectUserGroup-Guest"); ok {
					if err = d.Set("guest", vv); err != nil {
						return fmt.Errorf("Error reading guest: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading guest: %v", err)
				}
			}
		}
	}

	if err = d.Set("http_digest_realm", flattenObjectUserGroupHttpDigestRealm(o["http-digest-realm"], d, "http_digest_realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-digest-realm"], "ObjectUserGroup-HttpDigestRealm"); ok {
			if err = d.Set("http_digest_realm", vv); err != nil {
				return fmt.Errorf("Error reading http_digest_realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_digest_realm: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectUserGroupId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectUserGroup-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("match", flattenObjectUserGroupMatch(o["match"], d, "match")); err != nil {
			if vv, ok := fortiAPIPatch(o["match"], "ObjectUserGroup-Match"); ok {
				if err = d.Set("match", vv); err != nil {
					return fmt.Errorf("Error reading match: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading match: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("match"); ok {
			if err = d.Set("match", flattenObjectUserGroupMatch(o["match"], d, "match")); err != nil {
				if vv, ok := fortiAPIPatch(o["match"], "ObjectUserGroup-Match"); ok {
					if err = d.Set("match", vv); err != nil {
						return fmt.Errorf("Error reading match: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading match: %v", err)
				}
			}
		}
	}

	if err = d.Set("max_accounts", flattenObjectUserGroupMaxAccounts(o["max-accounts"], d, "max_accounts")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-accounts"], "ObjectUserGroup-MaxAccounts"); ok {
			if err = d.Set("max_accounts", vv); err != nil {
				return fmt.Errorf("Error reading max_accounts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_accounts: %v", err)
		}
	}

	if err = d.Set("member", flattenObjectUserGroupMember(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "ObjectUserGroup-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("mobile_phone", flattenObjectUserGroupMobilePhone(o["mobile-phone"], d, "mobile_phone")); err != nil {
		if vv, ok := fortiAPIPatch(o["mobile-phone"], "ObjectUserGroup-MobilePhone"); ok {
			if err = d.Set("mobile_phone", vv); err != nil {
				return fmt.Errorf("Error reading mobile_phone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mobile_phone: %v", err)
		}
	}

	if err = d.Set("multiple_guest_add", flattenObjectUserGroupMultipleGuestAdd(o["multiple-guest-add"], d, "multiple_guest_add")); err != nil {
		if vv, ok := fortiAPIPatch(o["multiple-guest-add"], "ObjectUserGroup-MultipleGuestAdd"); ok {
			if err = d.Set("multiple_guest_add", vv); err != nil {
				return fmt.Errorf("Error reading multiple_guest_add: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multiple_guest_add: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserGroupName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("password", flattenObjectUserGroupPassword(o["password"], d, "password")); err != nil {
		if vv, ok := fortiAPIPatch(o["password"], "ObjectUserGroup-Password"); ok {
			if err = d.Set("password", vv); err != nil {
				return fmt.Errorf("Error reading password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("sms_custom_server", flattenObjectUserGroupSmsCustomServer(o["sms-custom-server"], d, "sms_custom_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-custom-server"], "ObjectUserGroup-SmsCustomServer"); ok {
			if err = d.Set("sms_custom_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_custom_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_custom_server: %v", err)
		}
	}

	if err = d.Set("sms_server", flattenObjectUserGroupSmsServer(o["sms-server"], d, "sms_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-server"], "ObjectUserGroup-SmsServer"); ok {
			if err = d.Set("sms_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_server: %v", err)
		}
	}

	if err = d.Set("sponsor", flattenObjectUserGroupSponsor(o["sponsor"], d, "sponsor")); err != nil {
		if vv, ok := fortiAPIPatch(o["sponsor"], "ObjectUserGroup-Sponsor"); ok {
			if err = d.Set("sponsor", vv); err != nil {
				return fmt.Errorf("Error reading sponsor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sponsor: %v", err)
		}
	}

	if err = d.Set("sso_attribute_value", flattenObjectUserGroupSsoAttributeValue(o["sso-attribute-value"], d, "sso_attribute_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-attribute-value"], "ObjectUserGroup-SsoAttributeValue"); ok {
			if err = d.Set("sso_attribute_value", vv); err != nil {
				return fmt.Errorf("Error reading sso_attribute_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_attribute_value: %v", err)
		}
	}

	if err = d.Set("user_id", flattenObjectUserGroupUserId(o["user-id"], d, "user_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-id"], "ObjectUserGroup-UserId"); ok {
			if err = d.Set("user_id", vv); err != nil {
				return fmt.Errorf("Error reading user_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_id: %v", err)
		}
	}

	if err = d.Set("user_name", flattenObjectUserGroupUserName(o["user-name"], d, "user_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-name"], "ObjectUserGroup-UserName"); ok {
			if err = d.Set("user_name", vv); err != nil {
				return fmt.Errorf("Error reading user_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_name: %v", err)
		}
	}

	return nil
}

func flattenObjectUserGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserGroupAuthConcurrentOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupAuthConcurrentValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupAuthtimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupCompany(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupExpire(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupExpireType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupGroupType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupGuest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandObjectUserGroupGuestComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["company"], _ = expandObjectUserGroupGuestCompany(d, i["company"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["email"], _ = expandObjectUserGroupGuestEmail(d, i["email"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiration"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["expiration"], _ = expandObjectUserGroupGuestExpiration(d, i["expiration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectUserGroupGuestId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mobile-phone"], _ = expandObjectUserGroupGuestMobilePhone(d, i["mobile_phone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandObjectUserGroupGuestName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["password"], _ = expandObjectUserGroupGuestPassword(d, i["password"], pre_append)
		} else {
			tmp["password"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sponsor"], _ = expandObjectUserGroupGuestSponsor(d, i["sponsor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["user-id"], _ = expandObjectUserGroupGuestUserId(d, i["user_id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectUserGroupGuestComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupGuestCompany(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupGuestEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupGuestExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupGuestId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupGuestMobilePhone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupGuestName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupGuestPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserGroupGuestSponsor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupGuestUserId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupHttpDigestRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_gui_meta"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["_gui_meta"], _ = expandObjectUserGroupMatchGuiMeta(d, i["_gui_meta"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["group-name"], _ = expandObjectUserGroupMatchGroupName(d, i["group_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectUserGroupMatchId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["server-name"], _ = expandObjectUserGroupMatchServerName(d, i["server_name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectUserGroupMatchGuiMeta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupMatchGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupMatchId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupMatchServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupMaxAccounts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupMobilePhone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupMultipleGuestAdd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupSmsCustomServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupSmsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupSponsor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupSsoAttributeValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupUserId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_concurrent_override"); ok {
		t, err := expandObjectUserGroupAuthConcurrentOverride(d, v, "auth_concurrent_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-override"] = t
		}
	}

	if v, ok := d.GetOk("auth_concurrent_value"); ok {
		t, err := expandObjectUserGroupAuthConcurrentValue(d, v, "auth_concurrent_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-value"] = t
		}
	}

	if v, ok := d.GetOk("authtimeout"); ok {
		t, err := expandObjectUserGroupAuthtimeout(d, v, "authtimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authtimeout"] = t
		}
	}

	if v, ok := d.GetOk("company"); ok {
		t, err := expandObjectUserGroupCompany(d, v, "company")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["company"] = t
		}
	}

	if v, ok := d.GetOk("email"); ok {
		t, err := expandObjectUserGroupEmail(d, v, "email")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email"] = t
		}
	}

	if v, ok := d.GetOk("expire"); ok {
		t, err := expandObjectUserGroupExpire(d, v, "expire")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire"] = t
		}
	}

	if v, ok := d.GetOk("expire_type"); ok {
		t, err := expandObjectUserGroupExpireType(d, v, "expire_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-type"] = t
		}
	}

	if v, ok := d.GetOk("group_type"); ok {
		t, err := expandObjectUserGroupGroupType(d, v, "group_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-type"] = t
		}
	}

	if v, ok := d.GetOk("guest"); ok {
		t, err := expandObjectUserGroupGuest(d, v, "guest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest"] = t
		}
	}

	if v, ok := d.GetOk("http_digest_realm"); ok {
		t, err := expandObjectUserGroupHttpDigestRealm(d, v, "http_digest_realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-digest-realm"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandObjectUserGroupId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("match"); ok {
		t, err := expandObjectUserGroupMatch(d, v, "match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match"] = t
		}
	}

	if v, ok := d.GetOk("max_accounts"); ok {
		t, err := expandObjectUserGroupMaxAccounts(d, v, "max_accounts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-accounts"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok {
		t, err := expandObjectUserGroupMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("mobile_phone"); ok {
		t, err := expandObjectUserGroupMobilePhone(d, v, "mobile_phone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mobile-phone"] = t
		}
	}

	if v, ok := d.GetOk("multiple_guest_add"); ok {
		t, err := expandObjectUserGroupMultipleGuestAdd(d, v, "multiple_guest_add")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiple-guest-add"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectUserGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandObjectUserGroupPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("sms_custom_server"); ok {
		t, err := expandObjectUserGroupSmsCustomServer(d, v, "sms_custom_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-custom-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_server"); ok {
		t, err := expandObjectUserGroupSmsServer(d, v, "sms_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-server"] = t
		}
	}

	if v, ok := d.GetOk("sponsor"); ok {
		t, err := expandObjectUserGroupSponsor(d, v, "sponsor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sponsor"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute_value"); ok {
		t, err := expandObjectUserGroupSsoAttributeValue(d, v, "sso_attribute_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute-value"] = t
		}
	}

	if v, ok := d.GetOk("user_id"); ok {
		t, err := expandObjectUserGroupUserId(d, v, "user_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-id"] = t
		}
	}

	if v, ok := d.GetOk("user_name"); ok {
		t, err := expandObjectUserGroupUserName(d, v, "user_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-name"] = t
		}
	}

	return &obj, nil
}
