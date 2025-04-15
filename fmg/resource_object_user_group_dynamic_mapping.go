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

func resourceObjectUserGroupDynamicMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserGroupDynamicMappingCreate,
		Read:   resourceObjectUserGroupDynamicMappingRead,
		Update: resourceObjectUserGroupDynamicMappingUpdate,
		Delete: resourceObjectUserGroupDynamicMappingDelete,

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
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectUserGroupDynamicMappingCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	group := d.Get("group").(string)
	paradict["group"] = group

	obj, err := getObjectObjectUserGroupDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserGroupDynamicMapping resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectUserGroupDynamicMapping(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserGroupDynamicMapping resource: %v", err)
	}

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectUserGroupDynamicMappingRead(d, m)
}

func resourceObjectUserGroupDynamicMappingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	group := d.Get("group").(string)
	paradict["group"] = group

	obj, err := getObjectObjectUserGroupDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserGroupDynamicMapping resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectUserGroupDynamicMapping(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserGroupDynamicMapping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectUserGroupDynamicMappingRead(d, m)
}

func resourceObjectUserGroupDynamicMappingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	group := d.Get("group").(string)
	paradict["group"] = group

	wsParams["adom"] = adomv

	err = c.DeleteObjectUserGroupDynamicMapping(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserGroupDynamicMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserGroupDynamicMappingRead(d *schema.ResourceData, m interface{}) error {
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

	group := d.Get("group").(string)
	if group == "" {
		group = importOptionChecking(m.(*FortiClient).Cfg, "group")
		if group == "" {
			return fmt.Errorf("Parameter group is missing")
		}
		if err = d.Set("group", group); err != nil {
			return fmt.Errorf("Error set params group: %v", err)
		}
	}
	if mkey, err = checkScopeId(mkey); err != nil {
		return fmt.Errorf("Error set ID : %v", err)
	}
	paradict["group"] = group

	o, err := c.ReadObjectUserGroupDynamicMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserGroupDynamicMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserGroupDynamicMapping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserGroupDynamicMapping resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserGroupDynamicMappingScope2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserGroupDynamicMappingScopeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectUserGroupDynamicMappingScopeVdom2edl(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserGroupDynamicMappingScopeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingScopeVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingAuthConcurrentOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingAuthConcurrentValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingAuthtimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingCompany2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingEmail2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingExpire2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingExpireType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGroupType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuest2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserGroupDynamicMappingGuestComment2edl(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := i["company"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestCompany2edl(i["company"], d, pre_append)
			tmp["company"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Company")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := i["email"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestEmail2edl(i["email"], d, pre_append)
			tmp["email"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Email")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiration"
		if _, ok := i["expiration"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestExpiration2edl(i["expiration"], d, pre_append)
			tmp["expiration"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Expiration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := i["group"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestGroup2edl(i["group"], d, pre_append)
			tmp["group"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Group")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := i["mobile-phone"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestMobilePhone2edl(i["mobile-phone"], d, pre_append)
			tmp["mobile_phone"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-MobilePhone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := i["sponsor"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestSponsor2edl(i["sponsor"], d, pre_append)
			tmp["sponsor"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-Sponsor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := i["user-id"]; ok {
			v := flattenObjectUserGroupDynamicMappingGuestUserId2edl(i["user-id"], d, pre_append)
			tmp["user_id"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Guest-UserId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserGroupDynamicMappingGuestComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestCompany2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestEmail2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestExpiration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestMobilePhone2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestSponsor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingGuestUserId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingHttpDigestRealm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingLdapMemberof2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingLogicType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMatch2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserGroupDynamicMappingMatchGuiMeta2edl(i["_gui_meta"], d, pre_append)
			tmp["_gui_meta"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Match-GuiMeta")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := i["group-name"]; ok {
			v := flattenObjectUserGroupDynamicMappingMatchGroupName2edl(i["group-name"], d, pre_append)
			tmp["group_name"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Match-GroupName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectUserGroupDynamicMappingMatchId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Match-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_name"
		if _, ok := i["server-name"]; ok {
			v := flattenObjectUserGroupDynamicMappingMatchServerName2edl(i["server-name"], d, pre_append)
			tmp["server_name"] = fortiAPISubPartPatch(v, "ObjectUserGroupDynamicMapping-Match-ServerName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserGroupDynamicMappingMatchGuiMeta2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMatchGroupName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMatchId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMatchServerName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserGroupDynamicMappingMaxAccounts2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMember2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserGroupDynamicMappingMobilePhone2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMultipleGuestAdd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingPassword2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingRedirUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSmsCustomServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserGroupDynamicMappingSmsServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSponsor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnBookmarksGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserGroupDynamicMappingSslvpnCacheCleaner2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnClientCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserGroupDynamicMappingSslvpnFtp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnHttp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnOsCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnOsCheckList2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectUserGroupDynamicMappingSslvpnOsCheckListAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "latest_patch_level"
	if _, ok := i["latest-patch-level"]; ok {
		result["latest_patch_level"] = flattenObjectUserGroupDynamicMappingSslvpnOsCheckListLatestPatchLevel2edl(i["latest-patch-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "name"
	if _, ok := i["name"]; ok {
		result["name"] = flattenObjectUserGroupDynamicMappingSslvpnOsCheckListName2edl(i["name"], d, pre_append)
	}

	pre_append = pre + ".0." + "tolerance"
	if _, ok := i["tolerance"]; ok {
		result["tolerance"] = flattenObjectUserGroupDynamicMappingSslvpnOsCheckListTolerance2edl(i["tolerance"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectUserGroupDynamicMappingSslvpnOsCheckListAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnOsCheckListLatestPatchLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnOsCheckListName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnOsCheckListTolerance2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnPortal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserGroupDynamicMappingSslvpnPortalHeading2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnRdp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnSamba2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnSplitTunneling2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnSsh2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnTelnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnTunnel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnTunnelEndip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnTunnelIpMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnTunnelStartip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnVirtualDesktop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnVnc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSslvpnWebapp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingSsoAttributeValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingUserId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingUserName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserGroupDynamicMapping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("_scope", flattenObjectUserGroupDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["_scope"], "ObjectUserGroupDynamicMapping-Scope"); ok {
				if err = d.Set("_scope", vv); err != nil {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading _scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("_scope"); ok {
			if err = d.Set("_scope", flattenObjectUserGroupDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["_scope"], "ObjectUserGroupDynamicMapping-Scope"); ok {
					if err = d.Set("_scope", vv); err != nil {
						return fmt.Errorf("Error reading _scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			}
		}
	}

	if err = d.Set("auth_concurrent_override", flattenObjectUserGroupDynamicMappingAuthConcurrentOverride2edl(o["auth-concurrent-override"], d, "auth_concurrent_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-concurrent-override"], "ObjectUserGroupDynamicMapping-AuthConcurrentOverride"); ok {
			if err = d.Set("auth_concurrent_override", vv); err != nil {
				return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
		}
	}

	if err = d.Set("auth_concurrent_value", flattenObjectUserGroupDynamicMappingAuthConcurrentValue2edl(o["auth-concurrent-value"], d, "auth_concurrent_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-concurrent-value"], "ObjectUserGroupDynamicMapping-AuthConcurrentValue"); ok {
			if err = d.Set("auth_concurrent_value", vv); err != nil {
				return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
		}
	}

	if err = d.Set("authtimeout", flattenObjectUserGroupDynamicMappingAuthtimeout2edl(o["authtimeout"], d, "authtimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["authtimeout"], "ObjectUserGroupDynamicMapping-Authtimeout"); ok {
			if err = d.Set("authtimeout", vv); err != nil {
				return fmt.Errorf("Error reading authtimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authtimeout: %v", err)
		}
	}

	if err = d.Set("company", flattenObjectUserGroupDynamicMappingCompany2edl(o["company"], d, "company")); err != nil {
		if vv, ok := fortiAPIPatch(o["company"], "ObjectUserGroupDynamicMapping-Company"); ok {
			if err = d.Set("company", vv); err != nil {
				return fmt.Errorf("Error reading company: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading company: %v", err)
		}
	}

	if err = d.Set("email", flattenObjectUserGroupDynamicMappingEmail2edl(o["email"], d, "email")); err != nil {
		if vv, ok := fortiAPIPatch(o["email"], "ObjectUserGroupDynamicMapping-Email"); ok {
			if err = d.Set("email", vv); err != nil {
				return fmt.Errorf("Error reading email: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email: %v", err)
		}
	}

	if err = d.Set("expire", flattenObjectUserGroupDynamicMappingExpire2edl(o["expire"], d, "expire")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire"], "ObjectUserGroupDynamicMapping-Expire"); ok {
			if err = d.Set("expire", vv); err != nil {
				return fmt.Errorf("Error reading expire: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire: %v", err)
		}
	}

	if err = d.Set("expire_type", flattenObjectUserGroupDynamicMappingExpireType2edl(o["expire-type"], d, "expire_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire-type"], "ObjectUserGroupDynamicMapping-ExpireType"); ok {
			if err = d.Set("expire_type", vv); err != nil {
				return fmt.Errorf("Error reading expire_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire_type: %v", err)
		}
	}

	if err = d.Set("group_type", flattenObjectUserGroupDynamicMappingGroupType2edl(o["group-type"], d, "group_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-type"], "ObjectUserGroupDynamicMapping-GroupType"); ok {
			if err = d.Set("group_type", vv); err != nil {
				return fmt.Errorf("Error reading group_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("guest", flattenObjectUserGroupDynamicMappingGuest2edl(o["guest"], d, "guest")); err != nil {
			if vv, ok := fortiAPIPatch(o["guest"], "ObjectUserGroupDynamicMapping-Guest"); ok {
				if err = d.Set("guest", vv); err != nil {
					return fmt.Errorf("Error reading guest: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading guest: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("guest"); ok {
			if err = d.Set("guest", flattenObjectUserGroupDynamicMappingGuest2edl(o["guest"], d, "guest")); err != nil {
				if vv, ok := fortiAPIPatch(o["guest"], "ObjectUserGroupDynamicMapping-Guest"); ok {
					if err = d.Set("guest", vv); err != nil {
						return fmt.Errorf("Error reading guest: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading guest: %v", err)
				}
			}
		}
	}

	if err = d.Set("http_digest_realm", flattenObjectUserGroupDynamicMappingHttpDigestRealm2edl(o["http-digest-realm"], d, "http_digest_realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-digest-realm"], "ObjectUserGroupDynamicMapping-HttpDigestRealm"); ok {
			if err = d.Set("http_digest_realm", vv); err != nil {
				return fmt.Errorf("Error reading http_digest_realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_digest_realm: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectUserGroupDynamicMappingId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectUserGroupDynamicMapping-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ldap_memberof", flattenObjectUserGroupDynamicMappingLdapMemberof2edl(o["ldap-memberof"], d, "ldap_memberof")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-memberof"], "ObjectUserGroupDynamicMapping-LdapMemberof"); ok {
			if err = d.Set("ldap_memberof", vv); err != nil {
				return fmt.Errorf("Error reading ldap_memberof: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_memberof: %v", err)
		}
	}

	if err = d.Set("logic_type", flattenObjectUserGroupDynamicMappingLogicType2edl(o["logic-type"], d, "logic_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["logic-type"], "ObjectUserGroupDynamicMapping-LogicType"); ok {
			if err = d.Set("logic_type", vv); err != nil {
				return fmt.Errorf("Error reading logic_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logic_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("match", flattenObjectUserGroupDynamicMappingMatch2edl(o["match"], d, "match")); err != nil {
			if vv, ok := fortiAPIPatch(o["match"], "ObjectUserGroupDynamicMapping-Match"); ok {
				if err = d.Set("match", vv); err != nil {
					return fmt.Errorf("Error reading match: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading match: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("match"); ok {
			if err = d.Set("match", flattenObjectUserGroupDynamicMappingMatch2edl(o["match"], d, "match")); err != nil {
				if vv, ok := fortiAPIPatch(o["match"], "ObjectUserGroupDynamicMapping-Match"); ok {
					if err = d.Set("match", vv); err != nil {
						return fmt.Errorf("Error reading match: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading match: %v", err)
				}
			}
		}
	}

	if err = d.Set("max_accounts", flattenObjectUserGroupDynamicMappingMaxAccounts2edl(o["max-accounts"], d, "max_accounts")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-accounts"], "ObjectUserGroupDynamicMapping-MaxAccounts"); ok {
			if err = d.Set("max_accounts", vv); err != nil {
				return fmt.Errorf("Error reading max_accounts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_accounts: %v", err)
		}
	}

	if err = d.Set("member", flattenObjectUserGroupDynamicMappingMember2edl(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "ObjectUserGroupDynamicMapping-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("mobile_phone", flattenObjectUserGroupDynamicMappingMobilePhone2edl(o["mobile-phone"], d, "mobile_phone")); err != nil {
		if vv, ok := fortiAPIPatch(o["mobile-phone"], "ObjectUserGroupDynamicMapping-MobilePhone"); ok {
			if err = d.Set("mobile_phone", vv); err != nil {
				return fmt.Errorf("Error reading mobile_phone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mobile_phone: %v", err)
		}
	}

	if err = d.Set("multiple_guest_add", flattenObjectUserGroupDynamicMappingMultipleGuestAdd2edl(o["multiple-guest-add"], d, "multiple_guest_add")); err != nil {
		if vv, ok := fortiAPIPatch(o["multiple-guest-add"], "ObjectUserGroupDynamicMapping-MultipleGuestAdd"); ok {
			if err = d.Set("multiple_guest_add", vv); err != nil {
				return fmt.Errorf("Error reading multiple_guest_add: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multiple_guest_add: %v", err)
		}
	}

	if err = d.Set("password", flattenObjectUserGroupDynamicMappingPassword2edl(o["password"], d, "password")); err != nil {
		if vv, ok := fortiAPIPatch(o["password"], "ObjectUserGroupDynamicMapping-Password"); ok {
			if err = d.Set("password", vv); err != nil {
				return fmt.Errorf("Error reading password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("redir_url", flattenObjectUserGroupDynamicMappingRedirUrl2edl(o["redir-url"], d, "redir_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["redir-url"], "ObjectUserGroupDynamicMapping-RedirUrl"); ok {
			if err = d.Set("redir_url", vv); err != nil {
				return fmt.Errorf("Error reading redir_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redir_url: %v", err)
		}
	}

	if err = d.Set("sms_custom_server", flattenObjectUserGroupDynamicMappingSmsCustomServer2edl(o["sms-custom-server"], d, "sms_custom_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-custom-server"], "ObjectUserGroupDynamicMapping-SmsCustomServer"); ok {
			if err = d.Set("sms_custom_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_custom_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_custom_server: %v", err)
		}
	}

	if err = d.Set("sms_server", flattenObjectUserGroupDynamicMappingSmsServer2edl(o["sms-server"], d, "sms_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-server"], "ObjectUserGroupDynamicMapping-SmsServer"); ok {
			if err = d.Set("sms_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_server: %v", err)
		}
	}

	if err = d.Set("sponsor", flattenObjectUserGroupDynamicMappingSponsor2edl(o["sponsor"], d, "sponsor")); err != nil {
		if vv, ok := fortiAPIPatch(o["sponsor"], "ObjectUserGroupDynamicMapping-Sponsor"); ok {
			if err = d.Set("sponsor", vv); err != nil {
				return fmt.Errorf("Error reading sponsor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sponsor: %v", err)
		}
	}

	if err = d.Set("sslvpn_bookmarks_group", flattenObjectUserGroupDynamicMappingSslvpnBookmarksGroup2edl(o["sslvpn-bookmarks-group"], d, "sslvpn_bookmarks_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-bookmarks-group"], "ObjectUserGroupDynamicMapping-SslvpnBookmarksGroup"); ok {
			if err = d.Set("sslvpn_bookmarks_group", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_bookmarks_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_bookmarks_group: %v", err)
		}
	}

	if err = d.Set("sslvpn_cache_cleaner", flattenObjectUserGroupDynamicMappingSslvpnCacheCleaner2edl(o["sslvpn-cache-cleaner"], d, "sslvpn_cache_cleaner")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-cache-cleaner"], "ObjectUserGroupDynamicMapping-SslvpnCacheCleaner"); ok {
			if err = d.Set("sslvpn_cache_cleaner", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_cache_cleaner: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_cache_cleaner: %v", err)
		}
	}

	if err = d.Set("sslvpn_client_check", flattenObjectUserGroupDynamicMappingSslvpnClientCheck2edl(o["sslvpn-client-check"], d, "sslvpn_client_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-client-check"], "ObjectUserGroupDynamicMapping-SslvpnClientCheck"); ok {
			if err = d.Set("sslvpn_client_check", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_client_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_client_check: %v", err)
		}
	}

	if err = d.Set("sslvpn_ftp", flattenObjectUserGroupDynamicMappingSslvpnFtp2edl(o["sslvpn-ftp"], d, "sslvpn_ftp")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-ftp"], "ObjectUserGroupDynamicMapping-SslvpnFtp"); ok {
			if err = d.Set("sslvpn_ftp", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_ftp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_ftp: %v", err)
		}
	}

	if err = d.Set("sslvpn_http", flattenObjectUserGroupDynamicMappingSslvpnHttp2edl(o["sslvpn-http"], d, "sslvpn_http")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-http"], "ObjectUserGroupDynamicMapping-SslvpnHttp"); ok {
			if err = d.Set("sslvpn_http", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_http: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_http: %v", err)
		}
	}

	if err = d.Set("sslvpn_os_check", flattenObjectUserGroupDynamicMappingSslvpnOsCheck2edl(o["sslvpn-os-check"], d, "sslvpn_os_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-os-check"], "ObjectUserGroupDynamicMapping-SslvpnOsCheck"); ok {
			if err = d.Set("sslvpn_os_check", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_os_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_os_check: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sslvpn_os_check_list", flattenObjectUserGroupDynamicMappingSslvpnOsCheckList2edl(o["sslvpn-os-check-list"], d, "sslvpn_os_check_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["sslvpn-os-check-list"], "ObjectUserGroupDynamicMapping-SslvpnOsCheckList"); ok {
				if err = d.Set("sslvpn_os_check_list", vv); err != nil {
					return fmt.Errorf("Error reading sslvpn_os_check_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sslvpn_os_check_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sslvpn_os_check_list"); ok {
			if err = d.Set("sslvpn_os_check_list", flattenObjectUserGroupDynamicMappingSslvpnOsCheckList2edl(o["sslvpn-os-check-list"], d, "sslvpn_os_check_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["sslvpn-os-check-list"], "ObjectUserGroupDynamicMapping-SslvpnOsCheckList"); ok {
					if err = d.Set("sslvpn_os_check_list", vv); err != nil {
						return fmt.Errorf("Error reading sslvpn_os_check_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sslvpn_os_check_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("sslvpn_portal", flattenObjectUserGroupDynamicMappingSslvpnPortal2edl(o["sslvpn-portal"], d, "sslvpn_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-portal"], "ObjectUserGroupDynamicMapping-SslvpnPortal"); ok {
			if err = d.Set("sslvpn_portal", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_portal: %v", err)
		}
	}

	if err = d.Set("sslvpn_portal_heading", flattenObjectUserGroupDynamicMappingSslvpnPortalHeading2edl(o["sslvpn-portal-heading"], d, "sslvpn_portal_heading")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-portal-heading"], "ObjectUserGroupDynamicMapping-SslvpnPortalHeading"); ok {
			if err = d.Set("sslvpn_portal_heading", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_portal_heading: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_portal_heading: %v", err)
		}
	}

	if err = d.Set("sslvpn_rdp", flattenObjectUserGroupDynamicMappingSslvpnRdp2edl(o["sslvpn-rdp"], d, "sslvpn_rdp")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-rdp"], "ObjectUserGroupDynamicMapping-SslvpnRdp"); ok {
			if err = d.Set("sslvpn_rdp", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_rdp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_rdp: %v", err)
		}
	}

	if err = d.Set("sslvpn_samba", flattenObjectUserGroupDynamicMappingSslvpnSamba2edl(o["sslvpn-samba"], d, "sslvpn_samba")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-samba"], "ObjectUserGroupDynamicMapping-SslvpnSamba"); ok {
			if err = d.Set("sslvpn_samba", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_samba: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_samba: %v", err)
		}
	}

	if err = d.Set("sslvpn_split_tunneling", flattenObjectUserGroupDynamicMappingSslvpnSplitTunneling2edl(o["sslvpn-split-tunneling"], d, "sslvpn_split_tunneling")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-split-tunneling"], "ObjectUserGroupDynamicMapping-SslvpnSplitTunneling"); ok {
			if err = d.Set("sslvpn_split_tunneling", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_split_tunneling: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_split_tunneling: %v", err)
		}
	}

	if err = d.Set("sslvpn_ssh", flattenObjectUserGroupDynamicMappingSslvpnSsh2edl(o["sslvpn-ssh"], d, "sslvpn_ssh")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-ssh"], "ObjectUserGroupDynamicMapping-SslvpnSsh"); ok {
			if err = d.Set("sslvpn_ssh", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_ssh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_ssh: %v", err)
		}
	}

	if err = d.Set("sslvpn_telnet", flattenObjectUserGroupDynamicMappingSslvpnTelnet2edl(o["sslvpn-telnet"], d, "sslvpn_telnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-telnet"], "ObjectUserGroupDynamicMapping-SslvpnTelnet"); ok {
			if err = d.Set("sslvpn_telnet", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_telnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_telnet: %v", err)
		}
	}

	if err = d.Set("sslvpn_tunnel", flattenObjectUserGroupDynamicMappingSslvpnTunnel2edl(o["sslvpn-tunnel"], d, "sslvpn_tunnel")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-tunnel"], "ObjectUserGroupDynamicMapping-SslvpnTunnel"); ok {
			if err = d.Set("sslvpn_tunnel", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_tunnel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_tunnel: %v", err)
		}
	}

	if err = d.Set("sslvpn_tunnel_endip", flattenObjectUserGroupDynamicMappingSslvpnTunnelEndip2edl(o["sslvpn-tunnel-endip"], d, "sslvpn_tunnel_endip")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-tunnel-endip"], "ObjectUserGroupDynamicMapping-SslvpnTunnelEndip"); ok {
			if err = d.Set("sslvpn_tunnel_endip", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_tunnel_endip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_tunnel_endip: %v", err)
		}
	}

	if err = d.Set("sslvpn_tunnel_ip_mode", flattenObjectUserGroupDynamicMappingSslvpnTunnelIpMode2edl(o["sslvpn-tunnel-ip-mode"], d, "sslvpn_tunnel_ip_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-tunnel-ip-mode"], "ObjectUserGroupDynamicMapping-SslvpnTunnelIpMode"); ok {
			if err = d.Set("sslvpn_tunnel_ip_mode", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_tunnel_ip_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_tunnel_ip_mode: %v", err)
		}
	}

	if err = d.Set("sslvpn_tunnel_startip", flattenObjectUserGroupDynamicMappingSslvpnTunnelStartip2edl(o["sslvpn-tunnel-startip"], d, "sslvpn_tunnel_startip")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-tunnel-startip"], "ObjectUserGroupDynamicMapping-SslvpnTunnelStartip"); ok {
			if err = d.Set("sslvpn_tunnel_startip", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_tunnel_startip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_tunnel_startip: %v", err)
		}
	}

	if err = d.Set("sslvpn_virtual_desktop", flattenObjectUserGroupDynamicMappingSslvpnVirtualDesktop2edl(o["sslvpn-virtual-desktop"], d, "sslvpn_virtual_desktop")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-virtual-desktop"], "ObjectUserGroupDynamicMapping-SslvpnVirtualDesktop"); ok {
			if err = d.Set("sslvpn_virtual_desktop", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_virtual_desktop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_virtual_desktop: %v", err)
		}
	}

	if err = d.Set("sslvpn_vnc", flattenObjectUserGroupDynamicMappingSslvpnVnc2edl(o["sslvpn-vnc"], d, "sslvpn_vnc")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-vnc"], "ObjectUserGroupDynamicMapping-SslvpnVnc"); ok {
			if err = d.Set("sslvpn_vnc", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_vnc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_vnc: %v", err)
		}
	}

	if err = d.Set("sslvpn_webapp", flattenObjectUserGroupDynamicMappingSslvpnWebapp2edl(o["sslvpn-webapp"], d, "sslvpn_webapp")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-webapp"], "ObjectUserGroupDynamicMapping-SslvpnWebapp"); ok {
			if err = d.Set("sslvpn_webapp", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_webapp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_webapp: %v", err)
		}
	}

	if err = d.Set("sso_attribute_value", flattenObjectUserGroupDynamicMappingSsoAttributeValue2edl(o["sso-attribute-value"], d, "sso_attribute_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-attribute-value"], "ObjectUserGroupDynamicMapping-SsoAttributeValue"); ok {
			if err = d.Set("sso_attribute_value", vv); err != nil {
				return fmt.Errorf("Error reading sso_attribute_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_attribute_value: %v", err)
		}
	}

	if err = d.Set("user_id", flattenObjectUserGroupDynamicMappingUserId2edl(o["user-id"], d, "user_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-id"], "ObjectUserGroupDynamicMapping-UserId"); ok {
			if err = d.Set("user_id", vv); err != nil {
				return fmt.Errorf("Error reading user_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_id: %v", err)
		}
	}

	if err = d.Set("user_name", flattenObjectUserGroupDynamicMappingUserName2edl(o["user-name"], d, "user_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-name"], "ObjectUserGroupDynamicMapping-UserName"); ok {
			if err = d.Set("user_name", vv); err != nil {
				return fmt.Errorf("Error reading user_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_name: %v", err)
		}
	}

	return nil
}

func flattenObjectUserGroupDynamicMappingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserGroupDynamicMappingScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectUserGroupDynamicMappingScopeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectUserGroupDynamicMappingScopeVdom2edl(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserGroupDynamicMappingScopeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingScopeVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingAuthConcurrentOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingAuthConcurrentValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingAuthtimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingCompany2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingEmail2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingExpire2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingExpireType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGroupType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["comment"], _ = expandObjectUserGroupDynamicMappingGuestComment2edl(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["company"], _ = expandObjectUserGroupDynamicMappingGuestCompany2edl(d, i["company"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["email"], _ = expandObjectUserGroupDynamicMappingGuestEmail2edl(d, i["email"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["expiration"], _ = expandObjectUserGroupDynamicMappingGuestExpiration2edl(d, i["expiration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group"], _ = expandObjectUserGroupDynamicMappingGuestGroup2edl(d, i["group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectUserGroupDynamicMappingGuestId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mobile-phone"], _ = expandObjectUserGroupDynamicMappingGuestMobilePhone2edl(d, i["mobile_phone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectUserGroupDynamicMappingGuestName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandObjectUserGroupDynamicMappingGuestPassword2edl(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sponsor"], _ = expandObjectUserGroupDynamicMappingGuestSponsor2edl(d, i["sponsor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-id"], _ = expandObjectUserGroupDynamicMappingGuestUserId2edl(d, i["user_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserGroupDynamicMappingGuestComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestCompany2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestEmail2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestExpiration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestMobilePhone2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserGroupDynamicMappingGuestSponsor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingGuestUserId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingHttpDigestRealm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingLdapMemberof2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingLogicType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["_gui_meta"], _ = expandObjectUserGroupDynamicMappingMatchGuiMeta2edl(d, i["_gui_meta"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-name"], _ = expandObjectUserGroupDynamicMappingMatchGroupName2edl(d, i["group_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectUserGroupDynamicMappingMatchId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-name"], _ = expandObjectUserGroupDynamicMappingMatchServerName2edl(d, i["server_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserGroupDynamicMappingMatchGuiMeta2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMatchGroupName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMatchId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMatchServerName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserGroupDynamicMappingMaxAccounts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMember2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserGroupDynamicMappingMobilePhone2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMultipleGuestAdd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingRedirUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSmsCustomServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserGroupDynamicMappingSmsServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSponsor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnBookmarksGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserGroupDynamicMappingSslvpnCacheCleaner2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnClientCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserGroupDynamicMappingSslvpnFtp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnHttp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnOsCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnOsCheckList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectUserGroupDynamicMappingSslvpnOsCheckListAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "latest_patch_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["latest-patch-level"], _ = expandObjectUserGroupDynamicMappingSslvpnOsCheckListLatestPatchLevel2edl(d, i["latest_patch_level"], pre_append)
	}
	pre_append = pre + ".0." + "name"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["name"], _ = expandObjectUserGroupDynamicMappingSslvpnOsCheckListName2edl(d, i["name"], pre_append)
	}
	pre_append = pre + ".0." + "tolerance"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tolerance"], _ = expandObjectUserGroupDynamicMappingSslvpnOsCheckListTolerance2edl(d, i["tolerance"], pre_append)
	}

	return result, nil
}

func expandObjectUserGroupDynamicMappingSslvpnOsCheckListAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnOsCheckListLatestPatchLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnOsCheckListName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnOsCheckListTolerance2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnPortal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserGroupDynamicMappingSslvpnPortalHeading2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnRdp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnSamba2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnSplitTunneling2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnSsh2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnTelnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnTunnel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnTunnelEndip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnTunnelIpMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnTunnelStartip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnVirtualDesktop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnVnc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSslvpnWebapp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingSsoAttributeValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingUserId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingUserName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserGroupDynamicMapping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_scope"); ok || d.HasChange("_scope") {
		t, err := expandObjectUserGroupDynamicMappingScope2edl(d, v, "_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_scope"] = t
		}
	}

	if v, ok := d.GetOk("auth_concurrent_override"); ok || d.HasChange("auth_concurrent_override") {
		t, err := expandObjectUserGroupDynamicMappingAuthConcurrentOverride2edl(d, v, "auth_concurrent_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-override"] = t
		}
	}

	if v, ok := d.GetOk("auth_concurrent_value"); ok || d.HasChange("auth_concurrent_value") {
		t, err := expandObjectUserGroupDynamicMappingAuthConcurrentValue2edl(d, v, "auth_concurrent_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-value"] = t
		}
	}

	if v, ok := d.GetOk("authtimeout"); ok || d.HasChange("authtimeout") {
		t, err := expandObjectUserGroupDynamicMappingAuthtimeout2edl(d, v, "authtimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authtimeout"] = t
		}
	}

	if v, ok := d.GetOk("company"); ok || d.HasChange("company") {
		t, err := expandObjectUserGroupDynamicMappingCompany2edl(d, v, "company")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["company"] = t
		}
	}

	if v, ok := d.GetOk("email"); ok || d.HasChange("email") {
		t, err := expandObjectUserGroupDynamicMappingEmail2edl(d, v, "email")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email"] = t
		}
	}

	if v, ok := d.GetOk("expire"); ok || d.HasChange("expire") {
		t, err := expandObjectUserGroupDynamicMappingExpire2edl(d, v, "expire")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire"] = t
		}
	}

	if v, ok := d.GetOk("expire_type"); ok || d.HasChange("expire_type") {
		t, err := expandObjectUserGroupDynamicMappingExpireType2edl(d, v, "expire_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-type"] = t
		}
	}

	if v, ok := d.GetOk("group_type"); ok || d.HasChange("group_type") {
		t, err := expandObjectUserGroupDynamicMappingGroupType2edl(d, v, "group_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-type"] = t
		}
	}

	if v, ok := d.GetOk("guest"); ok || d.HasChange("guest") {
		t, err := expandObjectUserGroupDynamicMappingGuest2edl(d, v, "guest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest"] = t
		}
	}

	if v, ok := d.GetOk("http_digest_realm"); ok || d.HasChange("http_digest_realm") {
		t, err := expandObjectUserGroupDynamicMappingHttpDigestRealm2edl(d, v, "http_digest_realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-digest-realm"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectUserGroupDynamicMappingId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ldap_memberof"); ok || d.HasChange("ldap_memberof") {
		t, err := expandObjectUserGroupDynamicMappingLdapMemberof2edl(d, v, "ldap_memberof")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-memberof"] = t
		}
	}

	if v, ok := d.GetOk("logic_type"); ok || d.HasChange("logic_type") {
		t, err := expandObjectUserGroupDynamicMappingLogicType2edl(d, v, "logic_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logic-type"] = t
		}
	}

	if v, ok := d.GetOk("match"); ok || d.HasChange("match") {
		t, err := expandObjectUserGroupDynamicMappingMatch2edl(d, v, "match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match"] = t
		}
	}

	if v, ok := d.GetOk("max_accounts"); ok || d.HasChange("max_accounts") {
		t, err := expandObjectUserGroupDynamicMappingMaxAccounts2edl(d, v, "max_accounts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-accounts"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandObjectUserGroupDynamicMappingMember2edl(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("mobile_phone"); ok || d.HasChange("mobile_phone") {
		t, err := expandObjectUserGroupDynamicMappingMobilePhone2edl(d, v, "mobile_phone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mobile-phone"] = t
		}
	}

	if v, ok := d.GetOk("multiple_guest_add"); ok || d.HasChange("multiple_guest_add") {
		t, err := expandObjectUserGroupDynamicMappingMultipleGuestAdd2edl(d, v, "multiple_guest_add")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiple-guest-add"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectUserGroupDynamicMappingPassword2edl(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("redir_url"); ok || d.HasChange("redir_url") {
		t, err := expandObjectUserGroupDynamicMappingRedirUrl2edl(d, v, "redir_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redir-url"] = t
		}
	}

	if v, ok := d.GetOk("sms_custom_server"); ok || d.HasChange("sms_custom_server") {
		t, err := expandObjectUserGroupDynamicMappingSmsCustomServer2edl(d, v, "sms_custom_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-custom-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_server"); ok || d.HasChange("sms_server") {
		t, err := expandObjectUserGroupDynamicMappingSmsServer2edl(d, v, "sms_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-server"] = t
		}
	}

	if v, ok := d.GetOk("sponsor"); ok || d.HasChange("sponsor") {
		t, err := expandObjectUserGroupDynamicMappingSponsor2edl(d, v, "sponsor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sponsor"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_bookmarks_group"); ok || d.HasChange("sslvpn_bookmarks_group") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnBookmarksGroup2edl(d, v, "sslvpn_bookmarks_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-bookmarks-group"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_cache_cleaner"); ok || d.HasChange("sslvpn_cache_cleaner") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnCacheCleaner2edl(d, v, "sslvpn_cache_cleaner")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-cache-cleaner"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_client_check"); ok || d.HasChange("sslvpn_client_check") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnClientCheck2edl(d, v, "sslvpn_client_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-client-check"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_ftp"); ok || d.HasChange("sslvpn_ftp") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnFtp2edl(d, v, "sslvpn_ftp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-ftp"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_http"); ok || d.HasChange("sslvpn_http") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnHttp2edl(d, v, "sslvpn_http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-http"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_os_check"); ok || d.HasChange("sslvpn_os_check") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnOsCheck2edl(d, v, "sslvpn_os_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-os-check"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_os_check_list"); ok || d.HasChange("sslvpn_os_check_list") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnOsCheckList2edl(d, v, "sslvpn_os_check_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-os-check-list"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_portal"); ok || d.HasChange("sslvpn_portal") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnPortal2edl(d, v, "sslvpn_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-portal"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_portal_heading"); ok || d.HasChange("sslvpn_portal_heading") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnPortalHeading2edl(d, v, "sslvpn_portal_heading")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-portal-heading"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_rdp"); ok || d.HasChange("sslvpn_rdp") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnRdp2edl(d, v, "sslvpn_rdp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-rdp"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_samba"); ok || d.HasChange("sslvpn_samba") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnSamba2edl(d, v, "sslvpn_samba")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-samba"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_split_tunneling"); ok || d.HasChange("sslvpn_split_tunneling") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnSplitTunneling2edl(d, v, "sslvpn_split_tunneling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-split-tunneling"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_ssh"); ok || d.HasChange("sslvpn_ssh") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnSsh2edl(d, v, "sslvpn_ssh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-ssh"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_telnet"); ok || d.HasChange("sslvpn_telnet") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnTelnet2edl(d, v, "sslvpn_telnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-telnet"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_tunnel"); ok || d.HasChange("sslvpn_tunnel") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnTunnel2edl(d, v, "sslvpn_tunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-tunnel"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_tunnel_endip"); ok || d.HasChange("sslvpn_tunnel_endip") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnTunnelEndip2edl(d, v, "sslvpn_tunnel_endip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-tunnel-endip"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_tunnel_ip_mode"); ok || d.HasChange("sslvpn_tunnel_ip_mode") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnTunnelIpMode2edl(d, v, "sslvpn_tunnel_ip_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-tunnel-ip-mode"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_tunnel_startip"); ok || d.HasChange("sslvpn_tunnel_startip") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnTunnelStartip2edl(d, v, "sslvpn_tunnel_startip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-tunnel-startip"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_virtual_desktop"); ok || d.HasChange("sslvpn_virtual_desktop") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnVirtualDesktop2edl(d, v, "sslvpn_virtual_desktop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-virtual-desktop"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_vnc"); ok || d.HasChange("sslvpn_vnc") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnVnc2edl(d, v, "sslvpn_vnc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-vnc"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_webapp"); ok || d.HasChange("sslvpn_webapp") {
		t, err := expandObjectUserGroupDynamicMappingSslvpnWebapp2edl(d, v, "sslvpn_webapp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-webapp"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute_value"); ok || d.HasChange("sso_attribute_value") {
		t, err := expandObjectUserGroupDynamicMappingSsoAttributeValue2edl(d, v, "sso_attribute_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute-value"] = t
		}
	}

	if v, ok := d.GetOk("user_id"); ok || d.HasChange("user_id") {
		t, err := expandObjectUserGroupDynamicMappingUserId2edl(d, v, "user_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-id"] = t
		}
	}

	if v, ok := d.GetOk("user_name"); ok || d.HasChange("user_name") {
		t, err := expandObjectUserGroupDynamicMappingUserName2edl(d, v, "user_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-name"] = t
		}
	}

	return &obj, nil
}
