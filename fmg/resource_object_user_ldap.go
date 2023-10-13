// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure LDAP server entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserLdap() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserLdapCreate,
		Read:   resourceObjectUserLdapRead,
		Update: resourceObjectUserLdapUpdate,
		Delete: resourceObjectUserLdapDelete,

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
			"account_key_cert_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_key_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_key_processing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_key_upn_san": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"antiphish": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ca_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_cert_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cnid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dn": &schema.Schema{
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
						"account_key_cert_field": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"account_key_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"account_key_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"account_key_processing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"account_key_upn_san": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"antiphish": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ca_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"client_cert": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"client_cert_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cnid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_member_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"group_object_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_object_search_base": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_search_base": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"interface_select_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_connections": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"member_attr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"obtain_user_info": &schema.Schema{
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
						"password_attr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password_expiry_warning": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"password_renewal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"retrieve_protection_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"search_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"secondary_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"secure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"server_identity_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_min_proto_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tertiary_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"two_factor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"two_factor_authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"two_factor_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"two_factor_notification": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"user_info_exchange_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"username": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"group_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_member_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_object_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_search_base": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"member_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"obtain_user_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"password_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"password_expiry_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password_renewal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"search_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"secondary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_identity_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tertiary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"two_factor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"two_factor_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"two_factor_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_info_exchange_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": &schema.Schema{
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

func resourceObjectUserLdapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectUserLdap(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserLdap resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserLdap(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserLdap resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserLdapRead(d, m)
}

func resourceObjectUserLdapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserLdap(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserLdap resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserLdap(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserLdap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserLdapRead(d, m)
}

func resourceObjectUserLdapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectUserLdap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserLdap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserLdapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectUserLdap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserLdap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserLdap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserLdap resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserLdapAccountKeyCertField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapAccountKeyFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapAccountKeyProcessing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapAccountKeyUpnSan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapAntiphish(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapClientCertAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapCnid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserLdapDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_cert_field"
		if _, ok := i["account-key-cert-field"]; ok {
			v := flattenObjectUserLdapDynamicMappingAccountKeyCertField(i["account-key-cert-field"], d, pre_append)
			tmp["account_key_cert_field"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-AccountKeyCertField")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_filter"
		if _, ok := i["account-key-filter"]; ok {
			v := flattenObjectUserLdapDynamicMappingAccountKeyFilter(i["account-key-filter"], d, pre_append)
			tmp["account_key_filter"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-AccountKeyFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_name"
		if _, ok := i["account-key-name"]; ok {
			v := flattenObjectUserLdapDynamicMappingAccountKeyName(i["account-key-name"], d, pre_append)
			tmp["account_key_name"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-AccountKeyName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_processing"
		if _, ok := i["account-key-processing"]; ok {
			v := flattenObjectUserLdapDynamicMappingAccountKeyProcessing(i["account-key-processing"], d, pre_append)
			tmp["account_key_processing"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-AccountKeyProcessing")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_upn_san"
		if _, ok := i["account-key-upn-san"]; ok {
			v := flattenObjectUserLdapDynamicMappingAccountKeyUpnSan(i["account-key-upn-san"], d, pre_append)
			tmp["account_key_upn_san"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-AccountKeyUpnSan")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "antiphish"
		if _, ok := i["antiphish"]; ok {
			v := flattenObjectUserLdapDynamicMappingAntiphish(i["antiphish"], d, pre_append)
			tmp["antiphish"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-Antiphish")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ca_cert"
		if _, ok := i["ca-cert"]; ok {
			v := flattenObjectUserLdapDynamicMappingCaCert(i["ca-cert"], d, pre_append)
			tmp["ca_cert"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-CaCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert"
		if _, ok := i["client-cert"]; ok {
			v := flattenObjectUserLdapDynamicMappingClientCert(i["client-cert"], d, pre_append)
			tmp["client_cert"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-ClientCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert_auth"
		if _, ok := i["client-cert-auth"]; ok {
			v := flattenObjectUserLdapDynamicMappingClientCertAuth(i["client-cert-auth"], d, pre_append)
			tmp["client_cert_auth"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-ClientCertAuth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cnid"
		if _, ok := i["cnid"]; ok {
			v := flattenObjectUserLdapDynamicMappingCnid(i["cnid"], d, pre_append)
			tmp["cnid"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-Cnid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dn"
		if _, ok := i["dn"]; ok {
			v := flattenObjectUserLdapDynamicMappingDn(i["dn"], d, pre_append)
			tmp["dn"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-Dn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {
			v := flattenObjectUserLdapDynamicMappingFilter(i["filter"], d, pre_append)
			tmp["filter"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-Filter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := i["group"]; ok {
			v := flattenObjectUserLdapDynamicMappingGroup(i["group"], d, pre_append)
			tmp["group"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-Group")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_filter"
		if _, ok := i["group-filter"]; ok {
			v := flattenObjectUserLdapDynamicMappingGroupFilter(i["group-filter"], d, pre_append)
			tmp["group_filter"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-GroupFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_member_check"
		if _, ok := i["group-member-check"]; ok {
			v := flattenObjectUserLdapDynamicMappingGroupMemberCheck(i["group-member-check"], d, pre_append)
			tmp["group_member_check"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-GroupMemberCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_object_filter"
		if _, ok := i["group-object-filter"]; ok {
			v := flattenObjectUserLdapDynamicMappingGroupObjectFilter(i["group-object-filter"], d, pre_append)
			tmp["group_object_filter"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-GroupObjectFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_object_search_base"
		if _, ok := i["group-object-search-base"]; ok {
			v := flattenObjectUserLdapDynamicMappingGroupObjectSearchBase(i["group-object-search-base"], d, pre_append)
			tmp["group_object_search_base"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-GroupObjectSearchBase")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_search_base"
		if _, ok := i["group-search-base"]; ok {
			v := flattenObjectUserLdapDynamicMappingGroupSearchBase(i["group-search-base"], d, pre_append)
			tmp["group_search_base"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-GroupSearchBase")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenObjectUserLdapDynamicMappingInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenObjectUserLdapDynamicMappingInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := i["max-connections"]; ok {
			v := flattenObjectUserLdapDynamicMappingMaxConnections(i["max-connections"], d, pre_append)
			tmp["max_connections"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-MaxConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_attr"
		if _, ok := i["member-attr"]; ok {
			v := flattenObjectUserLdapDynamicMappingMemberAttr(i["member-attr"], d, pre_append)
			tmp["member_attr"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-MemberAttr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obtain_user_info"
		if _, ok := i["obtain-user-info"]; ok {
			v := flattenObjectUserLdapDynamicMappingObtainUserInfo(i["obtain-user-info"], d, pre_append)
			tmp["obtain_user_info"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-ObtainUserInfo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			v := flattenObjectUserLdapDynamicMappingPassword(i["password"], d, pre_append)
			tmp["password"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-Password")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_attr"
		if _, ok := i["password-attr"]; ok {
			v := flattenObjectUserLdapDynamicMappingPasswordAttr(i["password-attr"], d, pre_append)
			tmp["password_attr"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-PasswordAttr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_expiry_warning"
		if _, ok := i["password-expiry-warning"]; ok {
			v := flattenObjectUserLdapDynamicMappingPasswordExpiryWarning(i["password-expiry-warning"], d, pre_append)
			tmp["password_expiry_warning"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-PasswordExpiryWarning")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_renewal"
		if _, ok := i["password-renewal"]; ok {
			v := flattenObjectUserLdapDynamicMappingPasswordRenewal(i["password-renewal"], d, pre_append)
			tmp["password_renewal"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-PasswordRenewal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectUserLdapDynamicMappingPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retrieve_protection_profile"
		if _, ok := i["retrieve-protection-profile"]; ok {
			v := flattenObjectUserLdapDynamicMappingRetrieveProtectionProfile(i["retrieve-protection-profile"], d, pre_append)
			tmp["retrieve_protection_profile"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-RetrieveProtectionProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "search_type"
		if _, ok := i["search-type"]; ok {
			v := flattenObjectUserLdapDynamicMappingSearchType(i["search-type"], d, pre_append)
			tmp["search_type"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-SearchType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_server"
		if _, ok := i["secondary-server"]; ok {
			v := flattenObjectUserLdapDynamicMappingSecondaryServer(i["secondary-server"], d, pre_append)
			tmp["secondary_server"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-SecondaryServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secure"
		if _, ok := i["secure"]; ok {
			v := flattenObjectUserLdapDynamicMappingSecure(i["secure"], d, pre_append)
			tmp["secure"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-Secure")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenObjectUserLdapDynamicMappingServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_identity_check"
		if _, ok := i["server-identity-check"]; ok {
			v := flattenObjectUserLdapDynamicMappingServerIdentityCheck(i["server-identity-check"], d, pre_append)
			tmp["server_identity_check"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-ServerIdentityCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenObjectUserLdapDynamicMappingSourceIp(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-SourceIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_port"
		if _, ok := i["source-port"]; ok {
			v := flattenObjectUserLdapDynamicMappingSourcePort(i["source-port"], d, pre_append)
			tmp["source_port"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-SourcePort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_proto_version"
		if _, ok := i["ssl-min-proto-version"]; ok {
			v := flattenObjectUserLdapDynamicMappingSslMinProtoVersion(i["ssl-min-proto-version"], d, pre_append)
			tmp["ssl_min_proto_version"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-SslMinProtoVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_server"
		if _, ok := i["tertiary-server"]; ok {
			v := flattenObjectUserLdapDynamicMappingTertiaryServer(i["tertiary-server"], d, pre_append)
			tmp["tertiary_server"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-TertiaryServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor"
		if _, ok := i["two-factor"]; ok {
			v := flattenObjectUserLdapDynamicMappingTwoFactor(i["two-factor"], d, pre_append)
			tmp["two_factor"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-TwoFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_authentication"
		if _, ok := i["two-factor-authentication"]; ok {
			v := flattenObjectUserLdapDynamicMappingTwoFactorAuthentication(i["two-factor-authentication"], d, pre_append)
			tmp["two_factor_authentication"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-TwoFactorAuthentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_filter"
		if _, ok := i["two-factor-filter"]; ok {
			v := flattenObjectUserLdapDynamicMappingTwoFactorFilter(i["two-factor-filter"], d, pre_append)
			tmp["two_factor_filter"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-TwoFactorFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_notification"
		if _, ok := i["two-factor-notification"]; ok {
			v := flattenObjectUserLdapDynamicMappingTwoFactorNotification(i["two-factor-notification"], d, pre_append)
			tmp["two_factor_notification"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-TwoFactorNotification")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectUserLdapDynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_info_exchange_server"
		if _, ok := i["user-info-exchange-server"]; ok {
			v := flattenObjectUserLdapDynamicMappingUserInfoExchangeServer(i["user-info-exchange-server"], d, pre_append)
			tmp["user_info_exchange_server"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-UserInfoExchangeServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "username"
		if _, ok := i["username"]; ok {
			v := flattenObjectUserLdapDynamicMappingUsername(i["username"], d, pre_append)
			tmp["username"] = fortiAPISubPartPatch(v, "ObjectUserLdap-DynamicMapping-Username")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectUserLdapDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserLdapDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserLdapDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectUserLdapDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectUserLdapDynamicMapping-Scope-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectUserLdapDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingAccountKeyCertField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingAccountKeyFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingAccountKeyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingAccountKeyProcessing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingAccountKeyUpnSan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingAntiphish(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLdapDynamicMappingClientCertAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingCnid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingDn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingGroupFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingGroupMemberCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingGroupObjectFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingGroupObjectSearchBase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingGroupSearchBase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingMemberAttr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingObtainUserInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLdapDynamicMappingPasswordAttr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingPasswordExpiryWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingPasswordRenewal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingRetrieveProtectionProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingSearchType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLdapDynamicMappingSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingServerIdentityCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingTertiaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingTwoFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingTwoFactorAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingTwoFactorFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingTwoFactorNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingUserInfoExchangeServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapGroupFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapGroupMemberCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapGroupObjectFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapGroupSearchBase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapMemberAttr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapObtainUserInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLdapPasswordAttr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapPasswordExpiryWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapPasswordRenewal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapSearchType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLdapSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapServerIdentityCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapTertiaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapTwoFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapTwoFactorAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapTwoFactorFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapTwoFactorNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapUserInfoExchangeServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserLdap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("account_key_cert_field", flattenObjectUserLdapAccountKeyCertField(o["account-key-cert-field"], d, "account_key_cert_field")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-cert-field"], "ObjectUserLdap-AccountKeyCertField"); ok {
			if err = d.Set("account_key_cert_field", vv); err != nil {
				return fmt.Errorf("Error reading account_key_cert_field: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_cert_field: %v", err)
		}
	}

	if err = d.Set("account_key_filter", flattenObjectUserLdapAccountKeyFilter(o["account-key-filter"], d, "account_key_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-filter"], "ObjectUserLdap-AccountKeyFilter"); ok {
			if err = d.Set("account_key_filter", vv); err != nil {
				return fmt.Errorf("Error reading account_key_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_filter: %v", err)
		}
	}

	if err = d.Set("account_key_processing", flattenObjectUserLdapAccountKeyProcessing(o["account-key-processing"], d, "account_key_processing")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-processing"], "ObjectUserLdap-AccountKeyProcessing"); ok {
			if err = d.Set("account_key_processing", vv); err != nil {
				return fmt.Errorf("Error reading account_key_processing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_processing: %v", err)
		}
	}

	if err = d.Set("account_key_upn_san", flattenObjectUserLdapAccountKeyUpnSan(o["account-key-upn-san"], d, "account_key_upn_san")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-upn-san"], "ObjectUserLdap-AccountKeyUpnSan"); ok {
			if err = d.Set("account_key_upn_san", vv); err != nil {
				return fmt.Errorf("Error reading account_key_upn_san: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_upn_san: %v", err)
		}
	}

	if err = d.Set("antiphish", flattenObjectUserLdapAntiphish(o["antiphish"], d, "antiphish")); err != nil {
		if vv, ok := fortiAPIPatch(o["antiphish"], "ObjectUserLdap-Antiphish"); ok {
			if err = d.Set("antiphish", vv); err != nil {
				return fmt.Errorf("Error reading antiphish: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antiphish: %v", err)
		}
	}

	if err = d.Set("ca_cert", flattenObjectUserLdapCaCert(o["ca-cert"], d, "ca_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca-cert"], "ObjectUserLdap-CaCert"); ok {
			if err = d.Set("ca_cert", vv); err != nil {
				return fmt.Errorf("Error reading ca_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca_cert: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenObjectUserLdapClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "ObjectUserLdap-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("client_cert_auth", flattenObjectUserLdapClientCertAuth(o["client-cert-auth"], d, "client_cert_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert-auth"], "ObjectUserLdap-ClientCertAuth"); ok {
			if err = d.Set("client_cert_auth", vv); err != nil {
				return fmt.Errorf("Error reading client_cert_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert_auth: %v", err)
		}
	}

	if err = d.Set("cnid", flattenObjectUserLdapCnid(o["cnid"], d, "cnid")); err != nil {
		if vv, ok := fortiAPIPatch(o["cnid"], "ObjectUserLdap-Cnid"); ok {
			if err = d.Set("cnid", vv); err != nil {
				return fmt.Errorf("Error reading cnid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cnid: %v", err)
		}
	}

	if err = d.Set("dn", flattenObjectUserLdapDn(o["dn"], d, "dn")); err != nil {
		if vv, ok := fortiAPIPatch(o["dn"], "ObjectUserLdap-Dn"); ok {
			if err = d.Set("dn", vv); err != nil {
				return fmt.Errorf("Error reading dn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dn: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectUserLdapDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectUserLdap-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectUserLdapDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectUserLdap-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("group_filter", flattenObjectUserLdapGroupFilter(o["group-filter"], d, "group_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-filter"], "ObjectUserLdap-GroupFilter"); ok {
			if err = d.Set("group_filter", vv); err != nil {
				return fmt.Errorf("Error reading group_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_filter: %v", err)
		}
	}

	if err = d.Set("group_member_check", flattenObjectUserLdapGroupMemberCheck(o["group-member-check"], d, "group_member_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-member-check"], "ObjectUserLdap-GroupMemberCheck"); ok {
			if err = d.Set("group_member_check", vv); err != nil {
				return fmt.Errorf("Error reading group_member_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_member_check: %v", err)
		}
	}

	if err = d.Set("group_object_filter", flattenObjectUserLdapGroupObjectFilter(o["group-object-filter"], d, "group_object_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-object-filter"], "ObjectUserLdap-GroupObjectFilter"); ok {
			if err = d.Set("group_object_filter", vv); err != nil {
				return fmt.Errorf("Error reading group_object_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_object_filter: %v", err)
		}
	}

	if err = d.Set("group_search_base", flattenObjectUserLdapGroupSearchBase(o["group-search-base"], d, "group_search_base")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-search-base"], "ObjectUserLdap-GroupSearchBase"); ok {
			if err = d.Set("group_search_base", vv); err != nil {
				return fmt.Errorf("Error reading group_search_base: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_search_base: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectUserLdapInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectUserLdap-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenObjectUserLdapInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "ObjectUserLdap-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("member_attr", flattenObjectUserLdapMemberAttr(o["member-attr"], d, "member_attr")); err != nil {
		if vv, ok := fortiAPIPatch(o["member-attr"], "ObjectUserLdap-MemberAttr"); ok {
			if err = d.Set("member_attr", vv); err != nil {
				return fmt.Errorf("Error reading member_attr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member_attr: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserLdapName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserLdap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("obtain_user_info", flattenObjectUserLdapObtainUserInfo(o["obtain-user-info"], d, "obtain_user_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["obtain-user-info"], "ObjectUserLdap-ObtainUserInfo"); ok {
			if err = d.Set("obtain_user_info", vv); err != nil {
				return fmt.Errorf("Error reading obtain_user_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obtain_user_info: %v", err)
		}
	}

	if err = d.Set("password_attr", flattenObjectUserLdapPasswordAttr(o["password-attr"], d, "password_attr")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-attr"], "ObjectUserLdap-PasswordAttr"); ok {
			if err = d.Set("password_attr", vv); err != nil {
				return fmt.Errorf("Error reading password_attr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_attr: %v", err)
		}
	}

	if err = d.Set("password_expiry_warning", flattenObjectUserLdapPasswordExpiryWarning(o["password-expiry-warning"], d, "password_expiry_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-expiry-warning"], "ObjectUserLdap-PasswordExpiryWarning"); ok {
			if err = d.Set("password_expiry_warning", vv); err != nil {
				return fmt.Errorf("Error reading password_expiry_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_expiry_warning: %v", err)
		}
	}

	if err = d.Set("password_renewal", flattenObjectUserLdapPasswordRenewal(o["password-renewal"], d, "password_renewal")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-renewal"], "ObjectUserLdap-PasswordRenewal"); ok {
			if err = d.Set("password_renewal", vv); err != nil {
				return fmt.Errorf("Error reading password_renewal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_renewal: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectUserLdapPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectUserLdap-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("search_type", flattenObjectUserLdapSearchType(o["search-type"], d, "search_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["search-type"], "ObjectUserLdap-SearchType"); ok {
			if err = d.Set("search_type", vv); err != nil {
				return fmt.Errorf("Error reading search_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading search_type: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenObjectUserLdapSecondaryServer(o["secondary-server"], d, "secondary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-server"], "ObjectUserLdap-SecondaryServer"); ok {
			if err = d.Set("secondary_server", vv); err != nil {
				return fmt.Errorf("Error reading secondary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("secure", flattenObjectUserLdapSecure(o["secure"], d, "secure")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure"], "ObjectUserLdap-Secure"); ok {
			if err = d.Set("secure", vv); err != nil {
				return fmt.Errorf("Error reading secure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure: %v", err)
		}
	}

	if err = d.Set("server", flattenObjectUserLdapServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "ObjectUserLdap-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server_identity_check", flattenObjectUserLdapServerIdentityCheck(o["server-identity-check"], d, "server_identity_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-identity-check"], "ObjectUserLdap-ServerIdentityCheck"); ok {
			if err = d.Set("server_identity_check", vv); err != nil {
				return fmt.Errorf("Error reading server_identity_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_identity_check: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenObjectUserLdapSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "ObjectUserLdap-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_port", flattenObjectUserLdapSourcePort(o["source-port"], d, "source_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-port"], "ObjectUserLdap-SourcePort"); ok {
			if err = d.Set("source_port", vv); err != nil {
				return fmt.Errorf("Error reading source_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_port: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenObjectUserLdapSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-proto-version"], "ObjectUserLdap-SslMinProtoVersion"); ok {
			if err = d.Set("ssl_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("tertiary_server", flattenObjectUserLdapTertiaryServer(o["tertiary-server"], d, "tertiary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tertiary-server"], "ObjectUserLdap-TertiaryServer"); ok {
			if err = d.Set("tertiary_server", vv); err != nil {
				return fmt.Errorf("Error reading tertiary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tertiary_server: %v", err)
		}
	}

	if err = d.Set("two_factor", flattenObjectUserLdapTwoFactor(o["two-factor"], d, "two_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor"], "ObjectUserLdap-TwoFactor"); ok {
			if err = d.Set("two_factor", vv); err != nil {
				return fmt.Errorf("Error reading two_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	if err = d.Set("two_factor_authentication", flattenObjectUserLdapTwoFactorAuthentication(o["two-factor-authentication"], d, "two_factor_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-authentication"], "ObjectUserLdap-TwoFactorAuthentication"); ok {
			if err = d.Set("two_factor_authentication", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_authentication: %v", err)
		}
	}

	if err = d.Set("two_factor_filter", flattenObjectUserLdapTwoFactorFilter(o["two-factor-filter"], d, "two_factor_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-filter"], "ObjectUserLdap-TwoFactorFilter"); ok {
			if err = d.Set("two_factor_filter", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_filter: %v", err)
		}
	}

	if err = d.Set("two_factor_notification", flattenObjectUserLdapTwoFactorNotification(o["two-factor-notification"], d, "two_factor_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-notification"], "ObjectUserLdap-TwoFactorNotification"); ok {
			if err = d.Set("two_factor_notification", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_notification: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectUserLdapType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectUserLdap-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user_info_exchange_server", flattenObjectUserLdapUserInfoExchangeServer(o["user-info-exchange-server"], d, "user_info_exchange_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-info-exchange-server"], "ObjectUserLdap-UserInfoExchangeServer"); ok {
			if err = d.Set("user_info_exchange_server", vv); err != nil {
				return fmt.Errorf("Error reading user_info_exchange_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_info_exchange_server: %v", err)
		}
	}

	if err = d.Set("username", flattenObjectUserLdapUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "ObjectUserLdap-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	return nil
}

func flattenObjectUserLdapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserLdapAccountKeyCertField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapAccountKeyFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapAccountKeyProcessing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapAccountKeyUpnSan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapAntiphish(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapClientCertAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapCnid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			t, err := expandObjectUserLdapDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_cert_field"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["account-key-cert-field"], _ = expandObjectUserLdapDynamicMappingAccountKeyCertField(d, i["account_key_cert_field"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["account-key-filter"], _ = expandObjectUserLdapDynamicMappingAccountKeyFilter(d, i["account_key_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["account-key-name"], _ = expandObjectUserLdapDynamicMappingAccountKeyName(d, i["account_key_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_processing"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["account-key-processing"], _ = expandObjectUserLdapDynamicMappingAccountKeyProcessing(d, i["account_key_processing"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_upn_san"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["account-key-upn-san"], _ = expandObjectUserLdapDynamicMappingAccountKeyUpnSan(d, i["account_key_upn_san"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "antiphish"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["antiphish"], _ = expandObjectUserLdapDynamicMappingAntiphish(d, i["antiphish"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ca_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ca-cert"], _ = expandObjectUserLdapDynamicMappingCaCert(d, i["ca_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-cert"], _ = expandObjectUserLdapDynamicMappingClientCert(d, i["client_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert_auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-cert-auth"], _ = expandObjectUserLdapDynamicMappingClientCertAuth(d, i["client_cert_auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cnid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cnid"], _ = expandObjectUserLdapDynamicMappingCnid(d, i["cnid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dn"], _ = expandObjectUserLdapDynamicMappingDn(d, i["dn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter"], _ = expandObjectUserLdapDynamicMappingFilter(d, i["filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group"], _ = expandObjectUserLdapDynamicMappingGroup(d, i["group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-filter"], _ = expandObjectUserLdapDynamicMappingGroupFilter(d, i["group_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_member_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-member-check"], _ = expandObjectUserLdapDynamicMappingGroupMemberCheck(d, i["group_member_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_object_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-object-filter"], _ = expandObjectUserLdapDynamicMappingGroupObjectFilter(d, i["group_object_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_object_search_base"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-object-search-base"], _ = expandObjectUserLdapDynamicMappingGroupObjectSearchBase(d, i["group_object_search_base"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_search_base"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-search-base"], _ = expandObjectUserLdapDynamicMappingGroupSearchBase(d, i["group_search_base"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandObjectUserLdapDynamicMappingInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-select-method"], _ = expandObjectUserLdapDynamicMappingInterfaceSelectMethod(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-connections"], _ = expandObjectUserLdapDynamicMappingMaxConnections(d, i["max_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_attr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member-attr"], _ = expandObjectUserLdapDynamicMappingMemberAttr(d, i["member_attr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obtain_user_info"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["obtain-user-info"], _ = expandObjectUserLdapDynamicMappingObtainUserInfo(d, i["obtain_user_info"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandObjectUserLdapDynamicMappingPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_attr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password-attr"], _ = expandObjectUserLdapDynamicMappingPasswordAttr(d, i["password_attr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_expiry_warning"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password-expiry-warning"], _ = expandObjectUserLdapDynamicMappingPasswordExpiryWarning(d, i["password_expiry_warning"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_renewal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password-renewal"], _ = expandObjectUserLdapDynamicMappingPasswordRenewal(d, i["password_renewal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectUserLdapDynamicMappingPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retrieve_protection_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["retrieve-protection-profile"], _ = expandObjectUserLdapDynamicMappingRetrieveProtectionProfile(d, i["retrieve_protection_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "search_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["search-type"], _ = expandObjectUserLdapDynamicMappingSearchType(d, i["search_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secondary-server"], _ = expandObjectUserLdapDynamicMappingSecondaryServer(d, i["secondary_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secure"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secure"], _ = expandObjectUserLdapDynamicMappingSecure(d, i["secure"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandObjectUserLdapDynamicMappingServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_identity_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-identity-check"], _ = expandObjectUserLdapDynamicMappingServerIdentityCheck(d, i["server_identity_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip"], _ = expandObjectUserLdapDynamicMappingSourceIp(d, i["source_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-port"], _ = expandObjectUserLdapDynamicMappingSourcePort(d, i["source_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_proto_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-min-proto-version"], _ = expandObjectUserLdapDynamicMappingSslMinProtoVersion(d, i["ssl_min_proto_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tertiary-server"], _ = expandObjectUserLdapDynamicMappingTertiaryServer(d, i["tertiary_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["two-factor"], _ = expandObjectUserLdapDynamicMappingTwoFactor(d, i["two_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_authentication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["two-factor-authentication"], _ = expandObjectUserLdapDynamicMappingTwoFactorAuthentication(d, i["two_factor_authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["two-factor-filter"], _ = expandObjectUserLdapDynamicMappingTwoFactorFilter(d, i["two_factor_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_notification"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["two-factor-notification"], _ = expandObjectUserLdapDynamicMappingTwoFactorNotification(d, i["two_factor_notification"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectUserLdapDynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_info_exchange_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-info-exchange-server"], _ = expandObjectUserLdapDynamicMappingUserInfoExchangeServer(d, i["user_info_exchange_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "username"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["username"], _ = expandObjectUserLdapDynamicMappingUsername(d, i["username"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectUserLdapDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectUserLdapDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectUserLdapDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectUserLdapDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingAccountKeyCertField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingAccountKeyFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingAccountKeyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingAccountKeyProcessing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingAccountKeyUpnSan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingAntiphish(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLdapDynamicMappingClientCertAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingCnid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingDn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingGroupFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingGroupMemberCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingGroupObjectFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingGroupObjectSearchBase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingGroupSearchBase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingMemberAttr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingObtainUserInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLdapDynamicMappingPasswordAttr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingPasswordExpiryWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingPasswordRenewal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingRetrieveProtectionProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingSearchType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLdapDynamicMappingSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingServerIdentityCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingSourcePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingTertiaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingTwoFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingTwoFactorAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingTwoFactorFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingTwoFactorNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingUserInfoExchangeServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapGroupFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapGroupMemberCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapGroupObjectFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapGroupSearchBase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapMemberAttr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapObtainUserInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLdapPasswordAttr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapPasswordExpiryWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapPasswordRenewal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapSearchType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLdapSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapServerIdentityCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapSourcePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapTertiaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapTwoFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapTwoFactorAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapTwoFactorFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapTwoFactorNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapUserInfoExchangeServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserLdap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("account_key_cert_field"); ok || d.HasChange("account_key_cert_field") {
		t, err := expandObjectUserLdapAccountKeyCertField(d, v, "account_key_cert_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-cert-field"] = t
		}
	}

	if v, ok := d.GetOk("account_key_filter"); ok || d.HasChange("account_key_filter") {
		t, err := expandObjectUserLdapAccountKeyFilter(d, v, "account_key_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-filter"] = t
		}
	}

	if v, ok := d.GetOk("account_key_processing"); ok || d.HasChange("account_key_processing") {
		t, err := expandObjectUserLdapAccountKeyProcessing(d, v, "account_key_processing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-processing"] = t
		}
	}

	if v, ok := d.GetOk("account_key_upn_san"); ok || d.HasChange("account_key_upn_san") {
		t, err := expandObjectUserLdapAccountKeyUpnSan(d, v, "account_key_upn_san")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-upn-san"] = t
		}
	}

	if v, ok := d.GetOk("antiphish"); ok || d.HasChange("antiphish") {
		t, err := expandObjectUserLdapAntiphish(d, v, "antiphish")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antiphish"] = t
		}
	}

	if v, ok := d.GetOk("ca_cert"); ok || d.HasChange("ca_cert") {
		t, err := expandObjectUserLdapCaCert(d, v, "ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandObjectUserLdapClientCert(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("client_cert_auth"); ok || d.HasChange("client_cert_auth") {
		t, err := expandObjectUserLdapClientCertAuth(d, v, "client_cert_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert-auth"] = t
		}
	}

	if v, ok := d.GetOk("cnid"); ok || d.HasChange("cnid") {
		t, err := expandObjectUserLdapCnid(d, v, "cnid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cnid"] = t
		}
	}

	if v, ok := d.GetOk("dn"); ok || d.HasChange("dn") {
		t, err := expandObjectUserLdapDn(d, v, "dn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dn"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectUserLdapDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("group_filter"); ok || d.HasChange("group_filter") {
		t, err := expandObjectUserLdapGroupFilter(d, v, "group_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-filter"] = t
		}
	}

	if v, ok := d.GetOk("group_member_check"); ok || d.HasChange("group_member_check") {
		t, err := expandObjectUserLdapGroupMemberCheck(d, v, "group_member_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-member-check"] = t
		}
	}

	if v, ok := d.GetOk("group_object_filter"); ok || d.HasChange("group_object_filter") {
		t, err := expandObjectUserLdapGroupObjectFilter(d, v, "group_object_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-object-filter"] = t
		}
	}

	if v, ok := d.GetOk("group_search_base"); ok || d.HasChange("group_search_base") {
		t, err := expandObjectUserLdapGroupSearchBase(d, v, "group_search_base")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-search-base"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectUserLdapInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandObjectUserLdapInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("member_attr"); ok || d.HasChange("member_attr") {
		t, err := expandObjectUserLdapMemberAttr(d, v, "member_attr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-attr"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserLdapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("obtain_user_info"); ok || d.HasChange("obtain_user_info") {
		t, err := expandObjectUserLdapObtainUserInfo(d, v, "obtain_user_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obtain-user-info"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectUserLdapPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("password_attr"); ok || d.HasChange("password_attr") {
		t, err := expandObjectUserLdapPasswordAttr(d, v, "password_attr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-attr"] = t
		}
	}

	if v, ok := d.GetOk("password_expiry_warning"); ok || d.HasChange("password_expiry_warning") {
		t, err := expandObjectUserLdapPasswordExpiryWarning(d, v, "password_expiry_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-expiry-warning"] = t
		}
	}

	if v, ok := d.GetOk("password_renewal"); ok || d.HasChange("password_renewal") {
		t, err := expandObjectUserLdapPasswordRenewal(d, v, "password_renewal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-renewal"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectUserLdapPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("search_type"); ok || d.HasChange("search_type") {
		t, err := expandObjectUserLdapSearchType(d, v, "search_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["search-type"] = t
		}
	}

	if v, ok := d.GetOk("secondary_server"); ok || d.HasChange("secondary_server") {
		t, err := expandObjectUserLdapSecondaryServer(d, v, "secondary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}

	if v, ok := d.GetOk("secure"); ok || d.HasChange("secure") {
		t, err := expandObjectUserLdapSecure(d, v, "secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandObjectUserLdapServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server_identity_check"); ok || d.HasChange("server_identity_check") {
		t, err := expandObjectUserLdapServerIdentityCheck(d, v, "server_identity_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-identity-check"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandObjectUserLdapSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_port"); ok || d.HasChange("source_port") {
		t, err := expandObjectUserLdapSourcePort(d, v, "source_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-port"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok || d.HasChange("ssl_min_proto_version") {
		t, err := expandObjectUserLdapSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_server"); ok || d.HasChange("tertiary_server") {
		t, err := expandObjectUserLdapTertiaryServer(d, v, "tertiary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-server"] = t
		}
	}

	if v, ok := d.GetOk("two_factor"); ok || d.HasChange("two_factor") {
		t, err := expandObjectUserLdapTwoFactor(d, v, "two_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_authentication"); ok || d.HasChange("two_factor_authentication") {
		t, err := expandObjectUserLdapTwoFactorAuthentication(d, v, "two_factor_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-authentication"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_filter"); ok || d.HasChange("two_factor_filter") {
		t, err := expandObjectUserLdapTwoFactorFilter(d, v, "two_factor_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-filter"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_notification"); ok || d.HasChange("two_factor_notification") {
		t, err := expandObjectUserLdapTwoFactorNotification(d, v, "two_factor_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-notification"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectUserLdapType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user_info_exchange_server"); ok || d.HasChange("user_info_exchange_server") {
		t, err := expandObjectUserLdapUserInfoExchangeServer(d, v, "user_info_exchange_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-info-exchange-server"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandObjectUserLdapUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	return &obj, nil
}
