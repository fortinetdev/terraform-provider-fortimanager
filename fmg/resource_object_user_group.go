// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure user groups.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			},
			"authtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"company": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_scope": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vdom": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"auth_concurrent_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_concurrent_value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"authtimeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"company": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"expire": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"expire_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
									},
									"company": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"email": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"expiration": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"group": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"mobile_phone": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
									},
									"user_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"http_digest_realm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ldap_memberof": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"logic_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"_gui_meta": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"group_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"server_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"max_accounts": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"member": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mobile_phone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"multiple_guest_add": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"redir_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sms_custom_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sms_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sponsor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_bookmarks_group": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_cache_cleaner": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_client_check": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sslvpn_ftp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_http": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_os_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_os_check_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"latest_patch_level": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"tolerance": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"sslvpn_portal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_portal_heading": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_rdp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_samba": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_split_tunneling": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_ssh": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_telnet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_tunnel_endip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_tunnel_ip_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_tunnel_startip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_virtual_desktop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_vnc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_webapp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sso_attribute_value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"expire": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"expire_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
						},
						"company": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"expiration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mobile_phone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"user_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"http_digest_realm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_gui_meta": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"server_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"max_accounts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mobile_phone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"multiple_guest_add": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sms_custom_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sms_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sponsor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso_attribute_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_name": &schema.Schema{
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

func resourceObjectUserGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectUserGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserGroup resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserGroup(obj, paradict)

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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectUserGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserGroup(obj, mkey, paradict)
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	err = c.DeleteObjectUserGroup(mkey, paradict)
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	o, err := c.ReadObjectUserGroup(mkey, paradict)
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
	return v
}

func flattenObjectUserGroupAuthConcurrentValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupAuthtimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupCompany(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenObjectUserGroupDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_concurrent_override"
		if _, ok := i["auth-concurrent-override"]; ok {
			v := flattenObjectUserGroupDynamicMappingAuthConcurrentOverride(i["auth-concurrent-override"], d, pre_append)
			tmp["auth_concurrent_override"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-AuthConcurrentOverride")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_concurrent_value"
		if _, ok := i["auth-concurrent-value"]; ok {
			v := flattenObjectUserGroupDynamicMappingAuthConcurrentValue(i["auth-concurrent-value"], d, pre_append)
			tmp["auth_concurrent_value"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-AuthConcurrentValue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authtimeout"
		if _, ok := i["authtimeout"]; ok {
			v := flattenObjectUserGroupDynamicMappingAuthtimeout(i["authtimeout"], d, pre_append)
			tmp["authtimeout"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-Authtimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := i["company"]; ok {
			v := flattenObjectUserGroupDynamicMappingCompany(i["company"], d, pre_append)
			tmp["company"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-Company")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := i["email"]; ok {
			v := flattenObjectUserGroupDynamicMappingEmail(i["email"], d, pre_append)
			tmp["email"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-Email")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expire"
		if _, ok := i["expire"]; ok {
			v := flattenObjectUserGroupDynamicMappingExpire(i["expire"], d, pre_append)
			tmp["expire"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-Expire")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expire_type"
		if _, ok := i["expire-type"]; ok {
			v := flattenObjectUserGroupDynamicMappingExpireType(i["expire-type"], d, pre_append)
			tmp["expire_type"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-ExpireType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_type"
		if _, ok := i["group-type"]; ok {
			v := flattenObjectUserGroupDynamicMappingGroupType(i["group-type"], d, pre_append)
			tmp["group_type"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-GroupType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guest"
		if _, ok := i["guest"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuest(i["guest"], d, pre_append)
			tmp["guest"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-Guest")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_digest_realm"
		if _, ok := i["http-digest-realm"]; ok {
			v := flattenObjectUserGroupDynamicMappingHttpDigestRealm(i["http-digest-realm"], d, pre_append)
			tmp["http_digest_realm"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-HttpDigestRealm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectUserGroupDynamicMappingId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_memberof"
		if _, ok := i["ldap-memberof"]; ok {
			v := flattenObjectUserGroupDynamicMappingLdapMemberof(i["ldap-memberof"], d, pre_append)
			tmp["ldap_memberof"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-LdapMemberof")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logic_type"
		if _, ok := i["logic-type"]; ok {
			v := flattenObjectUserGroupDynamicMappingLogicType(i["logic-type"], d, pre_append)
			tmp["logic_type"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-LogicType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match"
		if _, ok := i["match"]; ok {
			v := flattenObjectUserGroupDynamicMappingMatch(i["match"], d, pre_append)
			tmp["match"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-Match")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_accounts"
		if _, ok := i["max-accounts"]; ok {
			v := flattenObjectUserGroupDynamicMappingMaxAccounts(i["max-accounts"], d, pre_append)
			tmp["max_accounts"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-MaxAccounts")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := i["member"]; ok {
			v := flattenObjectUserGroupDynamicMappingMember(i["member"], d, pre_append)
			tmp["member"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-Member")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := i["mobile-phone"]; ok {
			v := flattenObjectUserGroupDynamicMappingMobilePhone(i["mobile-phone"], d, pre_append)
			tmp["mobile_phone"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-MobilePhone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multiple_guest_add"
		if _, ok := i["multiple-guest-add"]; ok {
			v := flattenObjectUserGroupDynamicMappingMultipleGuestAdd(i["multiple-guest-add"], d, pre_append)
			tmp["multiple_guest_add"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-MultipleGuestAdd")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			v := flattenObjectUserGroupDynamicMappingPassword(i["password"], d, pre_append)
			tmp["password"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-Password")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redir_url"
		if _, ok := i["redir-url"]; ok {
			v := flattenObjectUserGroupDynamicMappingRedirUrl(i["redir-url"], d, pre_append)
			tmp["redir_url"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-RedirUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_custom_server"
		if _, ok := i["sms-custom-server"]; ok {
			v := flattenObjectUserGroupDynamicMappingSmsCustomServer(i["sms-custom-server"], d, pre_append)
			tmp["sms_custom_server"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SmsCustomServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_server"
		if _, ok := i["sms-server"]; ok {
			v := flattenObjectUserGroupDynamicMappingSmsServer(i["sms-server"], d, pre_append)
			tmp["sms_server"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SmsServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := i["sponsor"]; ok {
			v := flattenObjectUserGroupDynamicMappingSponsor(i["sponsor"], d, pre_append)
			tmp["sponsor"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-Sponsor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_bookmarks_group"
		if _, ok := i["sslvpn-bookmarks-group"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnBookmarksGroup(i["sslvpn-bookmarks-group"], d, pre_append)
			tmp["sslvpn_bookmarks_group"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnBookmarksGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_cache_cleaner"
		if _, ok := i["sslvpn-cache-cleaner"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnCacheCleaner(i["sslvpn-cache-cleaner"], d, pre_append)
			tmp["sslvpn_cache_cleaner"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnCacheCleaner")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_client_check"
		if _, ok := i["sslvpn-client-check"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnClientCheck(i["sslvpn-client-check"], d, pre_append)
			tmp["sslvpn_client_check"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnClientCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_ftp"
		if _, ok := i["sslvpn-ftp"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnFtp(i["sslvpn-ftp"], d, pre_append)
			tmp["sslvpn_ftp"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnFtp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_http"
		if _, ok := i["sslvpn-http"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnHttp(i["sslvpn-http"], d, pre_append)
			tmp["sslvpn_http"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnHttp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_os_check"
		if _, ok := i["sslvpn-os-check"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnOsCheck(i["sslvpn-os-check"], d, pre_append)
			tmp["sslvpn_os_check"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnOsCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_os_check_list"
		if _, ok := i["sslvpn-os-check-list"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnOsCheckList(i["sslvpn-os-check-list"], d, pre_append)
			tmp["sslvpn_os_check_list"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnOsCheckList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_portal"
		if _, ok := i["sslvpn-portal"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnPortal(i["sslvpn-portal"], d, pre_append)
			tmp["sslvpn_portal"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnPortal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_portal_heading"
		if _, ok := i["sslvpn-portal-heading"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnPortalHeading(i["sslvpn-portal-heading"], d, pre_append)
			tmp["sslvpn_portal_heading"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnPortalHeading")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_rdp"
		if _, ok := i["sslvpn-rdp"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnRdp(i["sslvpn-rdp"], d, pre_append)
			tmp["sslvpn_rdp"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnRdp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_samba"
		if _, ok := i["sslvpn-samba"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnSamba(i["sslvpn-samba"], d, pre_append)
			tmp["sslvpn_samba"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnSamba")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_split_tunneling"
		if _, ok := i["sslvpn-split-tunneling"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnSplitTunneling(i["sslvpn-split-tunneling"], d, pre_append)
			tmp["sslvpn_split_tunneling"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnSplitTunneling")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_ssh"
		if _, ok := i["sslvpn-ssh"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnSsh(i["sslvpn-ssh"], d, pre_append)
			tmp["sslvpn_ssh"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnSsh")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_telnet"
		if _, ok := i["sslvpn-telnet"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnTelnet(i["sslvpn-telnet"], d, pre_append)
			tmp["sslvpn_telnet"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnTelnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel"
		if _, ok := i["sslvpn-tunnel"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnTunnel(i["sslvpn-tunnel"], d, pre_append)
			tmp["sslvpn_tunnel"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnTunnel")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel_endip"
		if _, ok := i["sslvpn-tunnel-endip"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnTunnelEndip(i["sslvpn-tunnel-endip"], d, pre_append)
			tmp["sslvpn_tunnel_endip"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnTunnelEndip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel_ip_mode"
		if _, ok := i["sslvpn-tunnel-ip-mode"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnTunnelIpMode(i["sslvpn-tunnel-ip-mode"], d, pre_append)
			tmp["sslvpn_tunnel_ip_mode"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnTunnelIpMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel_startip"
		if _, ok := i["sslvpn-tunnel-startip"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnTunnelStartip(i["sslvpn-tunnel-startip"], d, pre_append)
			tmp["sslvpn_tunnel_startip"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnTunnelStartip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_virtual_desktop"
		if _, ok := i["sslvpn-virtual-desktop"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnVirtualDesktop(i["sslvpn-virtual-desktop"], d, pre_append)
			tmp["sslvpn_virtual_desktop"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnVirtualDesktop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_vnc"
		if _, ok := i["sslvpn-vnc"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnVnc(i["sslvpn-vnc"], d, pre_append)
			tmp["sslvpn_vnc"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnVnc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_webapp"
		if _, ok := i["sslvpn-webapp"]; ok {
			v := flattenObjectUserGroupDynamicMappingSslvpnWebapp(i["sslvpn-webapp"], d, pre_append)
			tmp["sslvpn_webapp"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SslvpnWebapp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute_value"
		if _, ok := i["sso-attribute-value"]; ok {
			v := flattenObjectUserGroupDynamicMappingSsoAttributeValue(i["sso-attribute-value"], d, pre_append)
			tmp["sso_attribute_value"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-SsoAttributeValue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := i["user-id"]; ok {
			v := flattenObjectUserGroupDynamicMappingUserId(i["user-id"], d, pre_append)
			tmp["user_id"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-UserId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_name"
		if _, ok := i["user-name"]; ok {
			v := flattenObjectUserGroupDynamicMappingUserName(i["user-name"], d, pre_append)
			tmp["user_name"] = fortiAPISubPartPatch(v, "ObjectUserGroup-DynamicMapping-UserName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserGroupDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectUserGroupDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectUserGroupDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserGroupDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingAuthConcurrentOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingAuthConcurrentValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingAuthtimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingCompany(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingExpire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingExpireType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGroupType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuest(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserGroupDynamicMappingGuestComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := i["company"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestCompany(i["company"], d, pre_append)
			tmp["company"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Company")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := i["email"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestEmail(i["email"], d, pre_append)
			tmp["email"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Email")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiration"
		if _, ok := i["expiration"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestExpiration(i["expiration"], d, pre_append)
			tmp["expiration"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Expiration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := i["group"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestGroup(i["group"], d, pre_append)
			tmp["group"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Group")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := i["mobile-phone"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestMobilePhone(i["mobile-phone"], d, pre_append)
			tmp["mobile_phone"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-MobilePhone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestPassword(i["password"], d, pre_append)
			tmp["password"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Password")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := i["sponsor"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestSponsor(i["sponsor"], d, pre_append)
			tmp["sponsor"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Sponsor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := i["user-id"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestUserId(i["user-id"], d, pre_append)
			tmp["user_id"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-UserId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserGroupDynamicMappingGuestComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestCompany(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestMobilePhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserGroupDynamicMappingGuestSponsor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestUserId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingHttpDigestRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingLdapMemberof(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingLogicType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMatch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserGroupDynamicMappingMatchGuiMeta(i["_gui_meta"], d, pre_append)
			tmp["_gui_meta"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Match-GuiMeta")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := i["group-name"]; ok {
			v := flattenObjectUserGroupDynamicMappingMatchGroupName(i["group-name"], d, pre_append)
			tmp["group_name"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Match-GroupName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectUserGroupDynamicMappingMatchId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Match-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_name"
		if _, ok := i["server-name"]; ok {
			v := flattenObjectUserGroupDynamicMappingMatchServerName(i["server-name"], d, pre_append)
			tmp["server_name"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Match-ServerName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserGroupDynamicMappingMatchGuiMeta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMatchGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMatchId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMatchServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMaxAccounts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserGroupDynamicMappingMobilePhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMultipleGuestAdd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingRedirUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSmsCustomServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSmsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSponsor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnBookmarksGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserGroupDynamicMappingSslvpnCacheCleaner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnClientCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserGroupDynamicMappingSslvpnFtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnOsCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnOsCheckList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectUserGroupDynamicMappingSslvpnOsCheckListAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "latest_patch_level"
	if _, ok := i["latest-patch-level"]; ok {
		result["latest_patch_level"] = flattenObjectUserGroupDynamicMappingSslvpnOsCheckListLatestPatchLevel(i["latest-patch-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "name"
	if _, ok := i["name"]; ok {
		result["name"] = flattenObjectUserGroupDynamicMappingSslvpnOsCheckListName(i["name"], d, pre_append)
	}

	pre_append = pre + ".0." + "tolerance"
	if _, ok := i["tolerance"]; ok {
		result["tolerance"] = flattenObjectUserGroupDynamicMappingSslvpnOsCheckListTolerance(i["tolerance"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectUserGroupDynamicMappingSslvpnOsCheckListAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnOsCheckListLatestPatchLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnOsCheckListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnOsCheckListTolerance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserGroupDynamicMappingSslvpnPortalHeading(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnRdp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnSamba(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnSplitTunneling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnSsh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnTelnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnTunnelEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnTunnelIpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnTunnelStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnVirtualDesktop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnVnc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnWebapp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSsoAttributeValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingUserId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupExpire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupExpireType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupGroupType(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

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

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

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
	return flattenStringList(v)
}

func flattenObjectUserGroupMobilePhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupMultipleGuestAdd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupSmsCustomServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupSmsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupSponsor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupSsoAttributeValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupUserId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

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

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectUserGroupDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectUserGroup-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectUserGroupDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectUserGroup-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
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

func expandObjectUserGroupDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectUserGroupDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_concurrent_override"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-concurrent-override"], _ = expandObjectUserGroupDynamicMappingAuthConcurrentOverride(d, i["auth_concurrent_override"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_concurrent_value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-concurrent-value"], _ = expandObjectUserGroupDynamicMappingAuthConcurrentValue(d, i["auth_concurrent_value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authtimeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authtimeout"], _ = expandObjectUserGroupDynamicMappingAuthtimeout(d, i["authtimeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["company"], _ = expandObjectUserGroupDynamicMappingCompany(d, i["company"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["email"], _ = expandObjectUserGroupDynamicMappingEmail(d, i["email"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expire"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["expire"], _ = expandObjectUserGroupDynamicMappingExpire(d, i["expire"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expire_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["expire-type"], _ = expandObjectUserGroupDynamicMappingExpireType(d, i["expire_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-type"], _ = expandObjectUserGroupDynamicMappingGroupType(d, i["group_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guest"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectUserGroupDynamicMappingGuest(d, i["guest"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["guest"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_digest_realm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-digest-realm"], _ = expandObjectUserGroupDynamicMappingHttpDigestRealm(d, i["http_digest_realm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectUserGroupDynamicMappingId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_memberof"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldap-memberof"], _ = expandObjectUserGroupDynamicMappingLdapMemberof(d, i["ldap_memberof"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logic_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logic-type"], _ = expandObjectUserGroupDynamicMappingLogicType(d, i["logic_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectUserGroupDynamicMappingMatch(d, i["match"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["match"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_accounts"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-accounts"], _ = expandObjectUserGroupDynamicMappingMaxAccounts(d, i["max_accounts"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member"], _ = expandObjectUserGroupDynamicMappingMember(d, i["member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mobile-phone"], _ = expandObjectUserGroupDynamicMappingMobilePhone(d, i["mobile_phone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multiple_guest_add"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["multiple-guest-add"], _ = expandObjectUserGroupDynamicMappingMultipleGuestAdd(d, i["multiple_guest_add"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandObjectUserGroupDynamicMappingPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redir_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["redir-url"], _ = expandObjectUserGroupDynamicMappingRedirUrl(d, i["redir_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_custom_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sms-custom-server"], _ = expandObjectUserGroupDynamicMappingSmsCustomServer(d, i["sms_custom_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sms-server"], _ = expandObjectUserGroupDynamicMappingSmsServer(d, i["sms_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sponsor"], _ = expandObjectUserGroupDynamicMappingSponsor(d, i["sponsor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_bookmarks_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-bookmarks-group"], _ = expandObjectUserGroupDynamicMappingSslvpnBookmarksGroup(d, i["sslvpn_bookmarks_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_cache_cleaner"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-cache-cleaner"], _ = expandObjectUserGroupDynamicMappingSslvpnCacheCleaner(d, i["sslvpn_cache_cleaner"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_client_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-client-check"], _ = expandObjectUserGroupDynamicMappingSslvpnClientCheck(d, i["sslvpn_client_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_ftp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-ftp"], _ = expandObjectUserGroupDynamicMappingSslvpnFtp(d, i["sslvpn_ftp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_http"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-http"], _ = expandObjectUserGroupDynamicMappingSslvpnHttp(d, i["sslvpn_http"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_os_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-os-check"], _ = expandObjectUserGroupDynamicMappingSslvpnOsCheck(d, i["sslvpn_os_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_os_check_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectUserGroupDynamicMappingSslvpnOsCheckList(d, i["sslvpn_os_check_list"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sslvpn-os-check-list"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_portal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-portal"], _ = expandObjectUserGroupDynamicMappingSslvpnPortal(d, i["sslvpn_portal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_portal_heading"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-portal-heading"], _ = expandObjectUserGroupDynamicMappingSslvpnPortalHeading(d, i["sslvpn_portal_heading"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_rdp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-rdp"], _ = expandObjectUserGroupDynamicMappingSslvpnRdp(d, i["sslvpn_rdp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_samba"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-samba"], _ = expandObjectUserGroupDynamicMappingSslvpnSamba(d, i["sslvpn_samba"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_split_tunneling"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-split-tunneling"], _ = expandObjectUserGroupDynamicMappingSslvpnSplitTunneling(d, i["sslvpn_split_tunneling"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_ssh"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-ssh"], _ = expandObjectUserGroupDynamicMappingSslvpnSsh(d, i["sslvpn_ssh"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_telnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-telnet"], _ = expandObjectUserGroupDynamicMappingSslvpnTelnet(d, i["sslvpn_telnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-tunnel"], _ = expandObjectUserGroupDynamicMappingSslvpnTunnel(d, i["sslvpn_tunnel"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel_endip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-tunnel-endip"], _ = expandObjectUserGroupDynamicMappingSslvpnTunnelEndip(d, i["sslvpn_tunnel_endip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel_ip_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-tunnel-ip-mode"], _ = expandObjectUserGroupDynamicMappingSslvpnTunnelIpMode(d, i["sslvpn_tunnel_ip_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel_startip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-tunnel-startip"], _ = expandObjectUserGroupDynamicMappingSslvpnTunnelStartip(d, i["sslvpn_tunnel_startip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_virtual_desktop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-virtual-desktop"], _ = expandObjectUserGroupDynamicMappingSslvpnVirtualDesktop(d, i["sslvpn_virtual_desktop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_vnc"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-vnc"], _ = expandObjectUserGroupDynamicMappingSslvpnVnc(d, i["sslvpn_vnc"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_webapp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-webapp"], _ = expandObjectUserGroupDynamicMappingSslvpnWebapp(d, i["sslvpn_webapp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute_value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-attribute-value"], _ = expandObjectUserGroupDynamicMappingSsoAttributeValue(d, i["sso_attribute_value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-id"], _ = expandObjectUserGroupDynamicMappingUserId(d, i["user_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-name"], _ = expandObjectUserGroupDynamicMappingUserName(d, i["user_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserGroupDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectUserGroupDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectUserGroupDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserGroupDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingAuthConcurrentOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingAuthConcurrentValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingAuthtimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingCompany(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingExpire(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingExpireType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGroupType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["comment"], _ = expandObjectUserGroupDynamicMappingGuestComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["company"], _ = expandObjectUserGroupDynamicMappingGuestCompany(d, i["company"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["email"], _ = expandObjectUserGroupDynamicMappingGuestEmail(d, i["email"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["expiration"], _ = expandObjectUserGroupDynamicMappingGuestExpiration(d, i["expiration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group"], _ = expandObjectUserGroupDynamicMappingGuestGroup(d, i["group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectUserGroupDynamicMappingGuestId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mobile-phone"], _ = expandObjectUserGroupDynamicMappingGuestMobilePhone(d, i["mobile_phone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectUserGroupDynamicMappingGuestName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandObjectUserGroupDynamicMappingGuestPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sponsor"], _ = expandObjectUserGroupDynamicMappingGuestSponsor(d, i["sponsor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-id"], _ = expandObjectUserGroupDynamicMappingGuestUserId(d, i["user_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserGroupDynamicMappingGuestComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestCompany(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestMobilePhone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserGroupDynamicMappingGuestSponsor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestUserId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingHttpDigestRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingLdapMemberof(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingLogicType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_gui_meta"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_gui_meta"], _ = expandObjectUserGroupDynamicMappingMatchGuiMeta(d, i["_gui_meta"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-name"], _ = expandObjectUserGroupDynamicMappingMatchGroupName(d, i["group_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectUserGroupDynamicMappingMatchId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-name"], _ = expandObjectUserGroupDynamicMappingMatchServerName(d, i["server_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserGroupDynamicMappingMatchGuiMeta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMatchGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMatchId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMatchServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMaxAccounts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserGroupDynamicMappingMobilePhone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMultipleGuestAdd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingRedirUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSmsCustomServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSmsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSponsor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnBookmarksGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserGroupDynamicMappingSslvpnCacheCleaner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnClientCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserGroupDynamicMappingSslvpnFtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnOsCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnOsCheckList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectUserGroupDynamicMappingSslvpnOsCheckListAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "latest_patch_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["latest-patch-level"], _ = expandObjectUserGroupDynamicMappingSslvpnOsCheckListLatestPatchLevel(d, i["latest_patch_level"], pre_append)
	}
	pre_append = pre + ".0." + "name"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["name"], _ = expandObjectUserGroupDynamicMappingSslvpnOsCheckListName(d, i["name"], pre_append)
	}
	pre_append = pre + ".0." + "tolerance"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tolerance"], _ = expandObjectUserGroupDynamicMappingSslvpnOsCheckListTolerance(d, i["tolerance"], pre_append)
	}

	return result, nil
}

func expandObjectUserGroupDynamicMappingSslvpnOsCheckListAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnOsCheckListLatestPatchLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnOsCheckListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnOsCheckListTolerance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserGroupDynamicMappingSslvpnPortalHeading(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnRdp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnSamba(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnSplitTunneling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnSsh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnTelnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnTunnelEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnTunnelIpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnTunnelStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnVirtualDesktop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnVnc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnWebapp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSsoAttributeValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingUserId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["comment"], _ = expandObjectUserGroupGuestComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["company"], _ = expandObjectUserGroupGuestCompany(d, i["company"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["email"], _ = expandObjectUserGroupGuestEmail(d, i["email"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["expiration"], _ = expandObjectUserGroupGuestExpiration(d, i["expiration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectUserGroupGuestId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mobile-phone"], _ = expandObjectUserGroupGuestMobilePhone(d, i["mobile_phone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectUserGroupGuestName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandObjectUserGroupGuestPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sponsor"], _ = expandObjectUserGroupGuestSponsor(d, i["sponsor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-id"], _ = expandObjectUserGroupGuestUserId(d, i["user_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

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
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_gui_meta"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_gui_meta"], _ = expandObjectUserGroupMatchGuiMeta(d, i["_gui_meta"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-name"], _ = expandObjectUserGroupMatchGroupName(d, i["group_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectUserGroupMatchId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-name"], _ = expandObjectUserGroupMatchServerName(d, i["server_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

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
	return expandStringList(v.(*schema.Set).List()), nil
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

	if v, ok := d.GetOk("auth_concurrent_override"); ok || d.HasChange("auth_concurrent_override") {
		t, err := expandObjectUserGroupAuthConcurrentOverride(d, v, "auth_concurrent_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-override"] = t
		}
	}

	if v, ok := d.GetOk("auth_concurrent_value"); ok || d.HasChange("auth_concurrent_value") {
		t, err := expandObjectUserGroupAuthConcurrentValue(d, v, "auth_concurrent_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-value"] = t
		}
	}

	if v, ok := d.GetOk("authtimeout"); ok || d.HasChange("authtimeout") {
		t, err := expandObjectUserGroupAuthtimeout(d, v, "authtimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authtimeout"] = t
		}
	}

	if v, ok := d.GetOk("company"); ok || d.HasChange("company") {
		t, err := expandObjectUserGroupCompany(d, v, "company")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["company"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectUserGroupDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("email"); ok || d.HasChange("email") {
		t, err := expandObjectUserGroupEmail(d, v, "email")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email"] = t
		}
	}

	if v, ok := d.GetOk("expire"); ok || d.HasChange("expire") {
		t, err := expandObjectUserGroupExpire(d, v, "expire")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire"] = t
		}
	}

	if v, ok := d.GetOk("expire_type"); ok || d.HasChange("expire_type") {
		t, err := expandObjectUserGroupExpireType(d, v, "expire_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-type"] = t
		}
	}

	if v, ok := d.GetOk("group_type"); ok || d.HasChange("group_type") {
		t, err := expandObjectUserGroupGroupType(d, v, "group_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-type"] = t
		}
	}

	if v, ok := d.GetOk("guest"); ok || d.HasChange("guest") {
		t, err := expandObjectUserGroupGuest(d, v, "guest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest"] = t
		}
	}

	if v, ok := d.GetOk("http_digest_realm"); ok || d.HasChange("http_digest_realm") {
		t, err := expandObjectUserGroupHttpDigestRealm(d, v, "http_digest_realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-digest-realm"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectUserGroupId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("match"); ok || d.HasChange("match") {
		t, err := expandObjectUserGroupMatch(d, v, "match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match"] = t
		}
	}

	if v, ok := d.GetOk("max_accounts"); ok || d.HasChange("max_accounts") {
		t, err := expandObjectUserGroupMaxAccounts(d, v, "max_accounts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-accounts"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandObjectUserGroupMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("mobile_phone"); ok || d.HasChange("mobile_phone") {
		t, err := expandObjectUserGroupMobilePhone(d, v, "mobile_phone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mobile-phone"] = t
		}
	}

	if v, ok := d.GetOk("multiple_guest_add"); ok || d.HasChange("multiple_guest_add") {
		t, err := expandObjectUserGroupMultipleGuestAdd(d, v, "multiple_guest_add")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiple-guest-add"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectUserGroupPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("sms_custom_server"); ok || d.HasChange("sms_custom_server") {
		t, err := expandObjectUserGroupSmsCustomServer(d, v, "sms_custom_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-custom-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_server"); ok || d.HasChange("sms_server") {
		t, err := expandObjectUserGroupSmsServer(d, v, "sms_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-server"] = t
		}
	}

	if v, ok := d.GetOk("sponsor"); ok || d.HasChange("sponsor") {
		t, err := expandObjectUserGroupSponsor(d, v, "sponsor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sponsor"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute_value"); ok || d.HasChange("sso_attribute_value") {
		t, err := expandObjectUserGroupSsoAttributeValue(d, v, "sso_attribute_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute-value"] = t
		}
	}

	if v, ok := d.GetOk("user_id"); ok || d.HasChange("user_id") {
		t, err := expandObjectUserGroupUserId(d, v, "user_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-id"] = t
		}
	}

	if v, ok := d.GetOk("user_name"); ok || d.HasChange("user_name") {
		t, err := expandObjectUserGroupUserName(d, v, "user_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-name"] = t
		}
	}

	return &obj, nil
}
