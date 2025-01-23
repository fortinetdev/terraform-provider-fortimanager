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

func resourceObjectUserLdapDynamicMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserLdapDynamicMappingCreate,
		Read:   resourceObjectUserLdapDynamicMappingRead,
		Update: resourceObjectUserLdapDynamicMappingUpdate,
		Delete: resourceObjectUserLdapDynamicMappingDelete,

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
			"ldap": &schema.Schema{
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
			"source_ip_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_max_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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

func resourceObjectUserLdapDynamicMappingCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	ldap := d.Get("ldap").(string)
	paradict["ldap"] = ldap

	obj, err := getObjectObjectUserLdapDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserLdapDynamicMapping resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserLdapDynamicMapping(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserLdapDynamicMapping resource: %v", err)
	}

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectUserLdapDynamicMappingRead(d, m)
}

func resourceObjectUserLdapDynamicMappingUpdate(d *schema.ResourceData, m interface{}) error {
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

	ldap := d.Get("ldap").(string)
	paradict["ldap"] = ldap

	obj, err := getObjectObjectUserLdapDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserLdapDynamicMapping resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserLdapDynamicMapping(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserLdapDynamicMapping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectUserLdapDynamicMappingRead(d, m)
}

func resourceObjectUserLdapDynamicMappingDelete(d *schema.ResourceData, m interface{}) error {
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

	ldap := d.Get("ldap").(string)
	paradict["ldap"] = ldap

	err = c.DeleteObjectUserLdapDynamicMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserLdapDynamicMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserLdapDynamicMappingRead(d *schema.ResourceData, m interface{}) error {
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

	ldap := d.Get("ldap").(string)
	if ldap == "" {
		ldap = importOptionChecking(m.(*FortiClient).Cfg, "ldap")
		if ldap == "" {
			return fmt.Errorf("Parameter ldap is missing")
		}
		if err = d.Set("ldap", ldap); err != nil {
			return fmt.Errorf("Error set params ldap: %v", err)
		}
	}
	if mkey, err = checkScopeId(mkey); err != nil {
		return fmt.Errorf("Error set ID : %v", err)
	}
	paradict["ldap"] = ldap

	o, err := c.ReadObjectUserLdapDynamicMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserLdapDynamicMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserLdapDynamicMapping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserLdapDynamicMapping resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserLdapDynamicMappingScope2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserLdapDynamicMappingScopeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserLdapDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectUserLdapDynamicMappingScopeVdom2edl(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectUserLdapDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserLdapDynamicMappingScopeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingScopeVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingAccountKeyCertField2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingAccountKeyFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingAccountKeyName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingAccountKeyProcessing2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingAccountKeyUpnSan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingAntiphish2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingCaCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingClientCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convstr2list(v, d.Get(pre))
}

func flattenObjectUserLdapDynamicMappingClientCertAuth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingCnid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingDn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingGroupFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingGroupMemberCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingGroupObjectFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingGroupObjectSearchBase2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingGroupSearchBase2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserLdapDynamicMappingInterfaceSelectMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingMaxConnections2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingMemberAttr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingObtainUserInfo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingPasswordAttr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingPasswordExpiryWarning2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingPasswordRenewal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingRetrieveProtectionProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingSearchType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLdapDynamicMappingSecondaryServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingSecure2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingServerIdentityCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingSourceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingSourceIpInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLdapDynamicMappingSourcePort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingSslMaxProtoVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingSslMinProtoVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingStatusTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingTertiaryServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingTwoFactor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingTwoFactorAuthentication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingTwoFactorFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingTwoFactorNotification2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLdapDynamicMappingUserInfoExchangeServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserLdapDynamicMappingUsername2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserLdapDynamicMapping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("_scope", flattenObjectUserLdapDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["_scope"], "ObjectUserLdapDynamicMapping-Scope"); ok {
				if err = d.Set("_scope", vv); err != nil {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading _scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("_scope"); ok {
			if err = d.Set("_scope", flattenObjectUserLdapDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["_scope"], "ObjectUserLdapDynamicMapping-Scope"); ok {
					if err = d.Set("_scope", vv); err != nil {
						return fmt.Errorf("Error reading _scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			}
		}
	}

	if err = d.Set("account_key_cert_field", flattenObjectUserLdapDynamicMappingAccountKeyCertField2edl(o["account-key-cert-field"], d, "account_key_cert_field")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-cert-field"], "ObjectUserLdapDynamicMapping-AccountKeyCertField"); ok {
			if err = d.Set("account_key_cert_field", vv); err != nil {
				return fmt.Errorf("Error reading account_key_cert_field: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_cert_field: %v", err)
		}
	}

	if err = d.Set("account_key_filter", flattenObjectUserLdapDynamicMappingAccountKeyFilter2edl(o["account-key-filter"], d, "account_key_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-filter"], "ObjectUserLdapDynamicMapping-AccountKeyFilter"); ok {
			if err = d.Set("account_key_filter", vv); err != nil {
				return fmt.Errorf("Error reading account_key_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_filter: %v", err)
		}
	}

	if err = d.Set("account_key_name", flattenObjectUserLdapDynamicMappingAccountKeyName2edl(o["account-key-name"], d, "account_key_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-name"], "ObjectUserLdapDynamicMapping-AccountKeyName"); ok {
			if err = d.Set("account_key_name", vv); err != nil {
				return fmt.Errorf("Error reading account_key_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_name: %v", err)
		}
	}

	if err = d.Set("account_key_processing", flattenObjectUserLdapDynamicMappingAccountKeyProcessing2edl(o["account-key-processing"], d, "account_key_processing")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-processing"], "ObjectUserLdapDynamicMapping-AccountKeyProcessing"); ok {
			if err = d.Set("account_key_processing", vv); err != nil {
				return fmt.Errorf("Error reading account_key_processing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_processing: %v", err)
		}
	}

	if err = d.Set("account_key_upn_san", flattenObjectUserLdapDynamicMappingAccountKeyUpnSan2edl(o["account-key-upn-san"], d, "account_key_upn_san")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-upn-san"], "ObjectUserLdapDynamicMapping-AccountKeyUpnSan"); ok {
			if err = d.Set("account_key_upn_san", vv); err != nil {
				return fmt.Errorf("Error reading account_key_upn_san: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_upn_san: %v", err)
		}
	}

	if err = d.Set("antiphish", flattenObjectUserLdapDynamicMappingAntiphish2edl(o["antiphish"], d, "antiphish")); err != nil {
		if vv, ok := fortiAPIPatch(o["antiphish"], "ObjectUserLdapDynamicMapping-Antiphish"); ok {
			if err = d.Set("antiphish", vv); err != nil {
				return fmt.Errorf("Error reading antiphish: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antiphish: %v", err)
		}
	}

	if err = d.Set("ca_cert", flattenObjectUserLdapDynamicMappingCaCert2edl(o["ca-cert"], d, "ca_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca-cert"], "ObjectUserLdapDynamicMapping-CaCert"); ok {
			if err = d.Set("ca_cert", vv); err != nil {
				return fmt.Errorf("Error reading ca_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca_cert: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenObjectUserLdapDynamicMappingClientCert2edl(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "ObjectUserLdapDynamicMapping-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("client_cert_auth", flattenObjectUserLdapDynamicMappingClientCertAuth2edl(o["client-cert-auth"], d, "client_cert_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert-auth"], "ObjectUserLdapDynamicMapping-ClientCertAuth"); ok {
			if err = d.Set("client_cert_auth", vv); err != nil {
				return fmt.Errorf("Error reading client_cert_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert_auth: %v", err)
		}
	}

	if err = d.Set("cnid", flattenObjectUserLdapDynamicMappingCnid2edl(o["cnid"], d, "cnid")); err != nil {
		if vv, ok := fortiAPIPatch(o["cnid"], "ObjectUserLdapDynamicMapping-Cnid"); ok {
			if err = d.Set("cnid", vv); err != nil {
				return fmt.Errorf("Error reading cnid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cnid: %v", err)
		}
	}

	if err = d.Set("dn", flattenObjectUserLdapDynamicMappingDn2edl(o["dn"], d, "dn")); err != nil {
		if vv, ok := fortiAPIPatch(o["dn"], "ObjectUserLdapDynamicMapping-Dn"); ok {
			if err = d.Set("dn", vv); err != nil {
				return fmt.Errorf("Error reading dn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dn: %v", err)
		}
	}

	if err = d.Set("filter", flattenObjectUserLdapDynamicMappingFilter2edl(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "ObjectUserLdapDynamicMapping-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("group", flattenObjectUserLdapDynamicMappingGroup2edl(o["group"], d, "group")); err != nil {
		if vv, ok := fortiAPIPatch(o["group"], "ObjectUserLdapDynamicMapping-Group"); ok {
			if err = d.Set("group", vv); err != nil {
				return fmt.Errorf("Error reading group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group: %v", err)
		}
	}

	if err = d.Set("group_filter", flattenObjectUserLdapDynamicMappingGroupFilter2edl(o["group-filter"], d, "group_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-filter"], "ObjectUserLdapDynamicMapping-GroupFilter"); ok {
			if err = d.Set("group_filter", vv); err != nil {
				return fmt.Errorf("Error reading group_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_filter: %v", err)
		}
	}

	if err = d.Set("group_member_check", flattenObjectUserLdapDynamicMappingGroupMemberCheck2edl(o["group-member-check"], d, "group_member_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-member-check"], "ObjectUserLdapDynamicMapping-GroupMemberCheck"); ok {
			if err = d.Set("group_member_check", vv); err != nil {
				return fmt.Errorf("Error reading group_member_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_member_check: %v", err)
		}
	}

	if err = d.Set("group_object_filter", flattenObjectUserLdapDynamicMappingGroupObjectFilter2edl(o["group-object-filter"], d, "group_object_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-object-filter"], "ObjectUserLdapDynamicMapping-GroupObjectFilter"); ok {
			if err = d.Set("group_object_filter", vv); err != nil {
				return fmt.Errorf("Error reading group_object_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_object_filter: %v", err)
		}
	}

	if err = d.Set("group_object_search_base", flattenObjectUserLdapDynamicMappingGroupObjectSearchBase2edl(o["group-object-search-base"], d, "group_object_search_base")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-object-search-base"], "ObjectUserLdapDynamicMapping-GroupObjectSearchBase"); ok {
			if err = d.Set("group_object_search_base", vv); err != nil {
				return fmt.Errorf("Error reading group_object_search_base: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_object_search_base: %v", err)
		}
	}

	if err = d.Set("group_search_base", flattenObjectUserLdapDynamicMappingGroupSearchBase2edl(o["group-search-base"], d, "group_search_base")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-search-base"], "ObjectUserLdapDynamicMapping-GroupSearchBase"); ok {
			if err = d.Set("group_search_base", vv); err != nil {
				return fmt.Errorf("Error reading group_search_base: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_search_base: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectUserLdapDynamicMappingInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectUserLdapDynamicMapping-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenObjectUserLdapDynamicMappingInterfaceSelectMethod2edl(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "ObjectUserLdapDynamicMapping-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("max_connections", flattenObjectUserLdapDynamicMappingMaxConnections2edl(o["max-connections"], d, "max_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-connections"], "ObjectUserLdapDynamicMapping-MaxConnections"); ok {
			if err = d.Set("max_connections", vv); err != nil {
				return fmt.Errorf("Error reading max_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_connections: %v", err)
		}
	}

	if err = d.Set("member_attr", flattenObjectUserLdapDynamicMappingMemberAttr2edl(o["member-attr"], d, "member_attr")); err != nil {
		if vv, ok := fortiAPIPatch(o["member-attr"], "ObjectUserLdapDynamicMapping-MemberAttr"); ok {
			if err = d.Set("member_attr", vv); err != nil {
				return fmt.Errorf("Error reading member_attr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member_attr: %v", err)
		}
	}

	if err = d.Set("obtain_user_info", flattenObjectUserLdapDynamicMappingObtainUserInfo2edl(o["obtain-user-info"], d, "obtain_user_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["obtain-user-info"], "ObjectUserLdapDynamicMapping-ObtainUserInfo"); ok {
			if err = d.Set("obtain_user_info", vv); err != nil {
				return fmt.Errorf("Error reading obtain_user_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obtain_user_info: %v", err)
		}
	}

	if err = d.Set("password_attr", flattenObjectUserLdapDynamicMappingPasswordAttr2edl(o["password-attr"], d, "password_attr")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-attr"], "ObjectUserLdapDynamicMapping-PasswordAttr"); ok {
			if err = d.Set("password_attr", vv); err != nil {
				return fmt.Errorf("Error reading password_attr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_attr: %v", err)
		}
	}

	if err = d.Set("password_expiry_warning", flattenObjectUserLdapDynamicMappingPasswordExpiryWarning2edl(o["password-expiry-warning"], d, "password_expiry_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-expiry-warning"], "ObjectUserLdapDynamicMapping-PasswordExpiryWarning"); ok {
			if err = d.Set("password_expiry_warning", vv); err != nil {
				return fmt.Errorf("Error reading password_expiry_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_expiry_warning: %v", err)
		}
	}

	if err = d.Set("password_renewal", flattenObjectUserLdapDynamicMappingPasswordRenewal2edl(o["password-renewal"], d, "password_renewal")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-renewal"], "ObjectUserLdapDynamicMapping-PasswordRenewal"); ok {
			if err = d.Set("password_renewal", vv); err != nil {
				return fmt.Errorf("Error reading password_renewal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_renewal: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectUserLdapDynamicMappingPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectUserLdapDynamicMapping-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("retrieve_protection_profile", flattenObjectUserLdapDynamicMappingRetrieveProtectionProfile2edl(o["retrieve-protection-profile"], d, "retrieve_protection_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["retrieve-protection-profile"], "ObjectUserLdapDynamicMapping-RetrieveProtectionProfile"); ok {
			if err = d.Set("retrieve_protection_profile", vv); err != nil {
				return fmt.Errorf("Error reading retrieve_protection_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retrieve_protection_profile: %v", err)
		}
	}

	if err = d.Set("search_type", flattenObjectUserLdapDynamicMappingSearchType2edl(o["search-type"], d, "search_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["search-type"], "ObjectUserLdapDynamicMapping-SearchType"); ok {
			if err = d.Set("search_type", vv); err != nil {
				return fmt.Errorf("Error reading search_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading search_type: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenObjectUserLdapDynamicMappingSecondaryServer2edl(o["secondary-server"], d, "secondary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-server"], "ObjectUserLdapDynamicMapping-SecondaryServer"); ok {
			if err = d.Set("secondary_server", vv); err != nil {
				return fmt.Errorf("Error reading secondary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("secure", flattenObjectUserLdapDynamicMappingSecure2edl(o["secure"], d, "secure")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure"], "ObjectUserLdapDynamicMapping-Secure"); ok {
			if err = d.Set("secure", vv); err != nil {
				return fmt.Errorf("Error reading secure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure: %v", err)
		}
	}

	if err = d.Set("server", flattenObjectUserLdapDynamicMappingServer2edl(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "ObjectUserLdapDynamicMapping-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server_identity_check", flattenObjectUserLdapDynamicMappingServerIdentityCheck2edl(o["server-identity-check"], d, "server_identity_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-identity-check"], "ObjectUserLdapDynamicMapping-ServerIdentityCheck"); ok {
			if err = d.Set("server_identity_check", vv); err != nil {
				return fmt.Errorf("Error reading server_identity_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_identity_check: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenObjectUserLdapDynamicMappingSourceIp2edl(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "ObjectUserLdapDynamicMapping-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip_interface", flattenObjectUserLdapDynamicMappingSourceIpInterface2edl(o["source-ip-interface"], d, "source_ip_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-interface"], "ObjectUserLdapDynamicMapping-SourceIpInterface"); ok {
			if err = d.Set("source_ip_interface", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_interface: %v", err)
		}
	}

	if err = d.Set("source_port", flattenObjectUserLdapDynamicMappingSourcePort2edl(o["source-port"], d, "source_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-port"], "ObjectUserLdapDynamicMapping-SourcePort"); ok {
			if err = d.Set("source_port", vv); err != nil {
				return fmt.Errorf("Error reading source_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_port: %v", err)
		}
	}

	if err = d.Set("ssl_max_proto_version", flattenObjectUserLdapDynamicMappingSslMaxProtoVersion2edl(o["ssl-max-proto-version"], d, "ssl_max_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-proto-version"], "ObjectUserLdapDynamicMapping-SslMaxProtoVersion"); ok {
			if err = d.Set("ssl_max_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_proto_version: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenObjectUserLdapDynamicMappingSslMinProtoVersion2edl(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-proto-version"], "ObjectUserLdapDynamicMapping-SslMinProtoVersion"); ok {
			if err = d.Set("ssl_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("status_ttl", flattenObjectUserLdapDynamicMappingStatusTtl2edl(o["status-ttl"], d, "status_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["status-ttl"], "ObjectUserLdapDynamicMapping-StatusTtl"); ok {
			if err = d.Set("status_ttl", vv); err != nil {
				return fmt.Errorf("Error reading status_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status_ttl: %v", err)
		}
	}

	if err = d.Set("tertiary_server", flattenObjectUserLdapDynamicMappingTertiaryServer2edl(o["tertiary-server"], d, "tertiary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tertiary-server"], "ObjectUserLdapDynamicMapping-TertiaryServer"); ok {
			if err = d.Set("tertiary_server", vv); err != nil {
				return fmt.Errorf("Error reading tertiary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tertiary_server: %v", err)
		}
	}

	if err = d.Set("two_factor", flattenObjectUserLdapDynamicMappingTwoFactor2edl(o["two-factor"], d, "two_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor"], "ObjectUserLdapDynamicMapping-TwoFactor"); ok {
			if err = d.Set("two_factor", vv); err != nil {
				return fmt.Errorf("Error reading two_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	if err = d.Set("two_factor_authentication", flattenObjectUserLdapDynamicMappingTwoFactorAuthentication2edl(o["two-factor-authentication"], d, "two_factor_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-authentication"], "ObjectUserLdapDynamicMapping-TwoFactorAuthentication"); ok {
			if err = d.Set("two_factor_authentication", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_authentication: %v", err)
		}
	}

	if err = d.Set("two_factor_filter", flattenObjectUserLdapDynamicMappingTwoFactorFilter2edl(o["two-factor-filter"], d, "two_factor_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-filter"], "ObjectUserLdapDynamicMapping-TwoFactorFilter"); ok {
			if err = d.Set("two_factor_filter", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_filter: %v", err)
		}
	}

	if err = d.Set("two_factor_notification", flattenObjectUserLdapDynamicMappingTwoFactorNotification2edl(o["two-factor-notification"], d, "two_factor_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-notification"], "ObjectUserLdapDynamicMapping-TwoFactorNotification"); ok {
			if err = d.Set("two_factor_notification", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_notification: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectUserLdapDynamicMappingType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectUserLdapDynamicMapping-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user_info_exchange_server", flattenObjectUserLdapDynamicMappingUserInfoExchangeServer2edl(o["user-info-exchange-server"], d, "user_info_exchange_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-info-exchange-server"], "ObjectUserLdapDynamicMapping-UserInfoExchangeServer"); ok {
			if err = d.Set("user_info_exchange_server", vv); err != nil {
				return fmt.Errorf("Error reading user_info_exchange_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_info_exchange_server: %v", err)
		}
	}

	if err = d.Set("username", flattenObjectUserLdapDynamicMappingUsername2edl(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "ObjectUserLdapDynamicMapping-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	return nil
}

func flattenObjectUserLdapDynamicMappingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserLdapDynamicMappingScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectUserLdapDynamicMappingScopeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectUserLdapDynamicMappingScopeVdom2edl(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserLdapDynamicMappingScopeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingScopeVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingAccountKeyCertField2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingAccountKeyFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingAccountKeyName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingAccountKeyProcessing2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingAccountKeyUpnSan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingAntiphish2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingCaCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingClientCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLdapDynamicMappingClientCertAuth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingCnid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingDn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingGroupFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingGroupMemberCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingGroupObjectFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingGroupObjectSearchBase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingGroupSearchBase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserLdapDynamicMappingInterfaceSelectMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingMaxConnections2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingMemberAttr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingObtainUserInfo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLdapDynamicMappingPasswordAttr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingPasswordExpiryWarning2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingPasswordRenewal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingRetrieveProtectionProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingSearchType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLdapDynamicMappingSecondaryServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingSecure2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingServerIdentityCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingSourceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingSourceIpInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLdapDynamicMappingSourcePort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingSslMaxProtoVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingSslMinProtoVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingStatusTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingTertiaryServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingTwoFactor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingTwoFactorAuthentication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingTwoFactorFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingTwoFactorNotification2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLdapDynamicMappingUserInfoExchangeServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserLdapDynamicMappingUsername2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserLdapDynamicMapping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_scope"); ok || d.HasChange("_scope") {
		t, err := expandObjectUserLdapDynamicMappingScope2edl(d, v, "_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_scope"] = t
		}
	}

	if v, ok := d.GetOk("account_key_cert_field"); ok || d.HasChange("account_key_cert_field") {
		t, err := expandObjectUserLdapDynamicMappingAccountKeyCertField2edl(d, v, "account_key_cert_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-cert-field"] = t
		}
	}

	if v, ok := d.GetOk("account_key_filter"); ok || d.HasChange("account_key_filter") {
		t, err := expandObjectUserLdapDynamicMappingAccountKeyFilter2edl(d, v, "account_key_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-filter"] = t
		}
	}

	if v, ok := d.GetOk("account_key_name"); ok || d.HasChange("account_key_name") {
		t, err := expandObjectUserLdapDynamicMappingAccountKeyName2edl(d, v, "account_key_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-name"] = t
		}
	}

	if v, ok := d.GetOk("account_key_processing"); ok || d.HasChange("account_key_processing") {
		t, err := expandObjectUserLdapDynamicMappingAccountKeyProcessing2edl(d, v, "account_key_processing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-processing"] = t
		}
	}

	if v, ok := d.GetOk("account_key_upn_san"); ok || d.HasChange("account_key_upn_san") {
		t, err := expandObjectUserLdapDynamicMappingAccountKeyUpnSan2edl(d, v, "account_key_upn_san")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-upn-san"] = t
		}
	}

	if v, ok := d.GetOk("antiphish"); ok || d.HasChange("antiphish") {
		t, err := expandObjectUserLdapDynamicMappingAntiphish2edl(d, v, "antiphish")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antiphish"] = t
		}
	}

	if v, ok := d.GetOk("ca_cert"); ok || d.HasChange("ca_cert") {
		t, err := expandObjectUserLdapDynamicMappingCaCert2edl(d, v, "ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandObjectUserLdapDynamicMappingClientCert2edl(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("client_cert_auth"); ok || d.HasChange("client_cert_auth") {
		t, err := expandObjectUserLdapDynamicMappingClientCertAuth2edl(d, v, "client_cert_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert-auth"] = t
		}
	}

	if v, ok := d.GetOk("cnid"); ok || d.HasChange("cnid") {
		t, err := expandObjectUserLdapDynamicMappingCnid2edl(d, v, "cnid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cnid"] = t
		}
	}

	if v, ok := d.GetOk("dn"); ok || d.HasChange("dn") {
		t, err := expandObjectUserLdapDynamicMappingDn2edl(d, v, "dn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dn"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandObjectUserLdapDynamicMappingFilter2edl(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("group"); ok || d.HasChange("group") {
		t, err := expandObjectUserLdapDynamicMappingGroup2edl(d, v, "group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group"] = t
		}
	}

	if v, ok := d.GetOk("group_filter"); ok || d.HasChange("group_filter") {
		t, err := expandObjectUserLdapDynamicMappingGroupFilter2edl(d, v, "group_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-filter"] = t
		}
	}

	if v, ok := d.GetOk("group_member_check"); ok || d.HasChange("group_member_check") {
		t, err := expandObjectUserLdapDynamicMappingGroupMemberCheck2edl(d, v, "group_member_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-member-check"] = t
		}
	}

	if v, ok := d.GetOk("group_object_filter"); ok || d.HasChange("group_object_filter") {
		t, err := expandObjectUserLdapDynamicMappingGroupObjectFilter2edl(d, v, "group_object_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-object-filter"] = t
		}
	}

	if v, ok := d.GetOk("group_object_search_base"); ok || d.HasChange("group_object_search_base") {
		t, err := expandObjectUserLdapDynamicMappingGroupObjectSearchBase2edl(d, v, "group_object_search_base")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-object-search-base"] = t
		}
	}

	if v, ok := d.GetOk("group_search_base"); ok || d.HasChange("group_search_base") {
		t, err := expandObjectUserLdapDynamicMappingGroupSearchBase2edl(d, v, "group_search_base")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-search-base"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectUserLdapDynamicMappingInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandObjectUserLdapDynamicMappingInterfaceSelectMethod2edl(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("max_connections"); ok || d.HasChange("max_connections") {
		t, err := expandObjectUserLdapDynamicMappingMaxConnections2edl(d, v, "max_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-connections"] = t
		}
	}

	if v, ok := d.GetOk("member_attr"); ok || d.HasChange("member_attr") {
		t, err := expandObjectUserLdapDynamicMappingMemberAttr2edl(d, v, "member_attr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-attr"] = t
		}
	}

	if v, ok := d.GetOk("obtain_user_info"); ok || d.HasChange("obtain_user_info") {
		t, err := expandObjectUserLdapDynamicMappingObtainUserInfo2edl(d, v, "obtain_user_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obtain-user-info"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectUserLdapDynamicMappingPassword2edl(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("password_attr"); ok || d.HasChange("password_attr") {
		t, err := expandObjectUserLdapDynamicMappingPasswordAttr2edl(d, v, "password_attr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-attr"] = t
		}
	}

	if v, ok := d.GetOk("password_expiry_warning"); ok || d.HasChange("password_expiry_warning") {
		t, err := expandObjectUserLdapDynamicMappingPasswordExpiryWarning2edl(d, v, "password_expiry_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-expiry-warning"] = t
		}
	}

	if v, ok := d.GetOk("password_renewal"); ok || d.HasChange("password_renewal") {
		t, err := expandObjectUserLdapDynamicMappingPasswordRenewal2edl(d, v, "password_renewal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-renewal"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectUserLdapDynamicMappingPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("retrieve_protection_profile"); ok || d.HasChange("retrieve_protection_profile") {
		t, err := expandObjectUserLdapDynamicMappingRetrieveProtectionProfile2edl(d, v, "retrieve_protection_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retrieve-protection-profile"] = t
		}
	}

	if v, ok := d.GetOk("search_type"); ok || d.HasChange("search_type") {
		t, err := expandObjectUserLdapDynamicMappingSearchType2edl(d, v, "search_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["search-type"] = t
		}
	}

	if v, ok := d.GetOk("secondary_server"); ok || d.HasChange("secondary_server") {
		t, err := expandObjectUserLdapDynamicMappingSecondaryServer2edl(d, v, "secondary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}

	if v, ok := d.GetOk("secure"); ok || d.HasChange("secure") {
		t, err := expandObjectUserLdapDynamicMappingSecure2edl(d, v, "secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandObjectUserLdapDynamicMappingServer2edl(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server_identity_check"); ok || d.HasChange("server_identity_check") {
		t, err := expandObjectUserLdapDynamicMappingServerIdentityCheck2edl(d, v, "server_identity_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-identity-check"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandObjectUserLdapDynamicMappingSourceIp2edl(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_interface"); ok || d.HasChange("source_ip_interface") {
		t, err := expandObjectUserLdapDynamicMappingSourceIpInterface2edl(d, v, "source_ip_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-interface"] = t
		}
	}

	if v, ok := d.GetOk("source_port"); ok || d.HasChange("source_port") {
		t, err := expandObjectUserLdapDynamicMappingSourcePort2edl(d, v, "source_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-port"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_proto_version"); ok || d.HasChange("ssl_max_proto_version") {
		t, err := expandObjectUserLdapDynamicMappingSslMaxProtoVersion2edl(d, v, "ssl_max_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok || d.HasChange("ssl_min_proto_version") {
		t, err := expandObjectUserLdapDynamicMappingSslMinProtoVersion2edl(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("status_ttl"); ok || d.HasChange("status_ttl") {
		t, err := expandObjectUserLdapDynamicMappingStatusTtl2edl(d, v, "status_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status-ttl"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_server"); ok || d.HasChange("tertiary_server") {
		t, err := expandObjectUserLdapDynamicMappingTertiaryServer2edl(d, v, "tertiary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-server"] = t
		}
	}

	if v, ok := d.GetOk("two_factor"); ok || d.HasChange("two_factor") {
		t, err := expandObjectUserLdapDynamicMappingTwoFactor2edl(d, v, "two_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_authentication"); ok || d.HasChange("two_factor_authentication") {
		t, err := expandObjectUserLdapDynamicMappingTwoFactorAuthentication2edl(d, v, "two_factor_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-authentication"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_filter"); ok || d.HasChange("two_factor_filter") {
		t, err := expandObjectUserLdapDynamicMappingTwoFactorFilter2edl(d, v, "two_factor_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-filter"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_notification"); ok || d.HasChange("two_factor_notification") {
		t, err := expandObjectUserLdapDynamicMappingTwoFactorNotification2edl(d, v, "two_factor_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-notification"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectUserLdapDynamicMappingType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user_info_exchange_server"); ok || d.HasChange("user_info_exchange_server") {
		t, err := expandObjectUserLdapDynamicMappingUserInfoExchangeServer2edl(d, v, "user_info_exchange_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-info-exchange-server"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandObjectUserLdapDynamicMappingUsername2edl(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	return &obj, nil
}
