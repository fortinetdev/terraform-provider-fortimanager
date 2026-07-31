// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SAML server entry configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserSaml() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserSamlCreate,
		Read:   resourceObjectUserSamlRead,
		Update: resourceObjectUserSamlUpdate,
		Delete: resourceObjectUserSamlDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"adfs_claim": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"clock_tolerance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"digest_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
						"adfs_claim": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"clock_tolerance": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"digest_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"entity_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fabric_force_sync": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fabric_object": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fabric_object_source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_claim_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"idp_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"idp_entity_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"idp_single_logout_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"idp_single_sign_on_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"limit_relaystate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"realm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"reauth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"require_signed_resp_and_asrt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scim_client": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"scim_group_attr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scim_user_attr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_provider_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"single_logout_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"single_sign_on_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sso_app_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_claim_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_force_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_object_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_claim_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"idp_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"idp_entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"idp_single_logout_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"idp_single_sign_on_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit_relaystate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"realm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"require_signed_resp_and_asrt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scim_client": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"scim_group_attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scim_user_attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_provider_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"single_logout_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"single_sign_on_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso_app_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_claim_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
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

func resourceObjectUserSamlCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserSaml(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserSaml resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectUserSaml(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectUserSaml(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectUserSaml resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectUserSaml(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectUserSaml resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserSamlRead(d, m)
}

func resourceObjectUserSamlUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserSaml(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserSaml resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectUserSaml(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserSaml resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserSamlRead(d, m)
}

func resourceObjectUserSamlDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectUserSaml(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserSaml resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserSamlRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectUserSaml(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectUserSaml resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserSaml(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserSaml resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserSamlAdfsClaim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlAuthUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserSamlClockTolerance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDigestMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserSamlDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adfs_claim"
		if _, ok := i["adfs-claim"]; ok {
			v := flattenObjectUserSamlDynamicMappingAdfsClaim(i["adfs-claim"], d, pre_append)
			tmp["adfs_claim"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-AdfsClaim")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_url"
		if _, ok := i["auth-url"]; ok {
			v := flattenObjectUserSamlDynamicMappingAuthUrl(i["auth-url"], d, pre_append)
			tmp["auth_url"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-AuthUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cert"
		if _, ok := i["cert"]; ok {
			v := flattenObjectUserSamlDynamicMappingCert(i["cert"], d, pre_append)
			tmp["cert"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-Cert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "clock_tolerance"
		if _, ok := i["clock-tolerance"]; ok {
			v := flattenObjectUserSamlDynamicMappingClockTolerance(i["clock-tolerance"], d, pre_append)
			tmp["clock_tolerance"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-ClockTolerance")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "digest_method"
		if _, ok := i["digest-method"]; ok {
			v := flattenObjectUserSamlDynamicMappingDigestMethod(i["digest-method"], d, pre_append)
			tmp["digest_method"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-DigestMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entity_id"
		if _, ok := i["entity-id"]; ok {
			v := flattenObjectUserSamlDynamicMappingEntityId(i["entity-id"], d, pre_append)
			tmp["entity_id"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-EntityId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_force_sync"
		if _, ok := i["fabric-force-sync"]; ok {
			v := flattenObjectUserSamlDynamicMappingFabricForceSync(i["fabric-force-sync"], d, pre_append)
			tmp["fabric_force_sync"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-FabricForceSync")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object"
		if _, ok := i["fabric-object"]; ok {
			v := flattenObjectUserSamlDynamicMappingFabricObject(i["fabric-object"], d, pre_append)
			tmp["fabric_object"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-FabricObject")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object_source"
		if _, ok := i["fabric-object-source"]; ok {
			v := flattenObjectUserSamlDynamicMappingFabricObjectSource(i["fabric-object-source"], d, pre_append)
			tmp["fabric_object_source"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-FabricObjectSource")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_claim_type"
		if _, ok := i["group-claim-type"]; ok {
			v := flattenObjectUserSamlDynamicMappingGroupClaimType(i["group-claim-type"], d, pre_append)
			tmp["group_claim_type"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-GroupClaimType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := i["group-name"]; ok {
			v := flattenObjectUserSamlDynamicMappingGroupName(i["group-name"], d, pre_append)
			tmp["group_name"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-GroupName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_cert"
		if _, ok := i["idp-cert"]; ok {
			v := flattenObjectUserSamlDynamicMappingIdpCert(i["idp-cert"], d, pre_append)
			tmp["idp_cert"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-IdpCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_entity_id"
		if _, ok := i["idp-entity-id"]; ok {
			v := flattenObjectUserSamlDynamicMappingIdpEntityId(i["idp-entity-id"], d, pre_append)
			tmp["idp_entity_id"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-IdpEntityId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := i["idp-single-logout-url"]; ok {
			v := flattenObjectUserSamlDynamicMappingIdpSingleLogoutUrl(i["idp-single-logout-url"], d, pre_append)
			tmp["idp_single_logout_url"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-IdpSingleLogoutUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := i["idp-single-sign-on-url"]; ok {
			v := flattenObjectUserSamlDynamicMappingIdpSingleSignOnUrl(i["idp-single-sign-on-url"], d, pre_append)
			tmp["idp_single_sign_on_url"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-IdpSingleSignOnUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "limit_relaystate"
		if _, ok := i["limit-relaystate"]; ok {
			v := flattenObjectUserSamlDynamicMappingLimitRelaystate(i["limit-relaystate"], d, pre_append)
			tmp["limit_relaystate"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-LimitRelaystate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realm"
		if _, ok := i["realm"]; ok {
			v := flattenObjectUserSamlDynamicMappingRealm(i["realm"], d, pre_append)
			tmp["realm"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-Realm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "reauth"
		if _, ok := i["reauth"]; ok {
			v := flattenObjectUserSamlDynamicMappingReauth(i["reauth"], d, pre_append)
			tmp["reauth"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-Reauth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "require_signed_resp_and_asrt"
		if _, ok := i["require-signed-resp-and-asrt"]; ok {
			v := flattenObjectUserSamlDynamicMappingRequireSignedRespAndAsrt(i["require-signed-resp-and-asrt"], d, pre_append)
			tmp["require_signed_resp_and_asrt"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-RequireSignedRespAndAsrt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scim_client"
		if _, ok := i["scim-client"]; ok {
			v := flattenObjectUserSamlDynamicMappingScimClient(i["scim-client"], d, pre_append)
			tmp["scim_client"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-ScimClient")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scim_group_attr_type"
		if _, ok := i["scim-group-attr-type"]; ok {
			v := flattenObjectUserSamlDynamicMappingScimGroupAttrType(i["scim-group-attr-type"], d, pre_append)
			tmp["scim_group_attr_type"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-ScimGroupAttrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scim_user_attr_type"
		if _, ok := i["scim-user-attr-type"]; ok {
			v := flattenObjectUserSamlDynamicMappingScimUserAttrType(i["scim-user-attr-type"], d, pre_append)
			tmp["scim_user_attr_type"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-ScimUserAttrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_provider_address"
		if _, ok := i["service-provider-address"]; ok {
			v := flattenObjectUserSamlDynamicMappingServiceProviderAddress(i["service-provider-address"], d, pre_append)
			tmp["service_provider_address"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-ServiceProviderAddress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "single_logout_url"
		if _, ok := i["single-logout-url"]; ok {
			v := flattenObjectUserSamlDynamicMappingSingleLogoutUrl(i["single-logout-url"], d, pre_append)
			tmp["single_logout_url"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-SingleLogoutUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "single_sign_on_url"
		if _, ok := i["single-sign-on-url"]; ok {
			v := flattenObjectUserSamlDynamicMappingSingleSignOnUrl(i["single-sign-on-url"], d, pre_append)
			tmp["single_sign_on_url"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-SingleSignOnUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_app_id"
		if _, ok := i["sso-app-id"]; ok {
			v := flattenObjectUserSamlDynamicMappingSsoAppId(i["sso-app-id"], d, pre_append)
			tmp["sso_app_id"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-SsoAppId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectUserSamlDynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_claim_type"
		if _, ok := i["user-claim-type"]; ok {
			v := flattenObjectUserSamlDynamicMappingUserClaimType(i["user-claim-type"], d, pre_append)
			tmp["user_claim_type"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-UserClaimType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_name"
		if _, ok := i["user-name"]; ok {
			v := flattenObjectUserSamlDynamicMappingUserName(i["user-name"], d, pre_append)
			tmp["user_name"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-UserName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_source"
		if _, ok := i["user-source"]; ok {
			v := flattenObjectUserSamlDynamicMappingUserSource(i["user-source"], d, pre_append)
			tmp["user_source"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-UserSource")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := i["uuid"]; ok {
			v := flattenObjectUserSamlDynamicMappingUuid(i["uuid"], d, pre_append)
			tmp["uuid"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-Uuid")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserSamlDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserSamlDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserSamlDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectUserSamlDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectUserSamlDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserSamlDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingAdfsClaim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingAuthUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserSamlDynamicMappingClockTolerance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingDigestMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingFabricForceSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingFabricObjectSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingGroupClaimType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingIdpCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserSamlDynamicMappingIdpEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingIdpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingIdpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingLimitRelaystate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingRequireSignedRespAndAsrt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingScimClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserSamlDynamicMappingScimGroupAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingScimUserAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingServiceProviderAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingSsoAppId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingUserClaimType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingUserSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlFabricForceSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlFabricObjectSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlGroupClaimType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlIdpCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserSamlIdpEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlIdpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlIdpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlLimitRelaystate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlRequireSignedRespAndAsrt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlScimClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserSamlScimGroupAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlScimUserAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlServiceProviderAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlSsoAppId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlUserClaimType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlUserSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserSaml(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("adfs_claim", flattenObjectUserSamlAdfsClaim(o["adfs-claim"], d, "adfs_claim")); err != nil {
		if vv, ok := fortiAPIPatch(o["adfs-claim"], "ObjectUserSaml-AdfsClaim"); ok {
			if err = d.Set("adfs_claim", vv); err != nil {
				return fmt.Errorf("Error reading adfs_claim: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adfs_claim: %v", err)
		}
	}

	if err = d.Set("auth_url", flattenObjectUserSamlAuthUrl(o["auth-url"], d, "auth_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-url"], "ObjectUserSaml-AuthUrl"); ok {
			if err = d.Set("auth_url", vv); err != nil {
				return fmt.Errorf("Error reading auth_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_url: %v", err)
		}
	}

	if err = d.Set("cert", flattenObjectUserSamlCert(o["cert"], d, "cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert"], "ObjectUserSaml-Cert"); ok {
			if err = d.Set("cert", vv); err != nil {
				return fmt.Errorf("Error reading cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("clock_tolerance", flattenObjectUserSamlClockTolerance(o["clock-tolerance"], d, "clock_tolerance")); err != nil {
		if vv, ok := fortiAPIPatch(o["clock-tolerance"], "ObjectUserSaml-ClockTolerance"); ok {
			if err = d.Set("clock_tolerance", vv); err != nil {
				return fmt.Errorf("Error reading clock_tolerance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clock_tolerance: %v", err)
		}
	}

	if err = d.Set("digest_method", flattenObjectUserSamlDigestMethod(o["digest-method"], d, "digest_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["digest-method"], "ObjectUserSaml-DigestMethod"); ok {
			if err = d.Set("digest_method", vv); err != nil {
				return fmt.Errorf("Error reading digest_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading digest_method: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectUserSamlDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectUserSaml-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectUserSamlDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectUserSaml-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("entity_id", flattenObjectUserSamlEntityId(o["entity-id"], d, "entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["entity-id"], "ObjectUserSaml-EntityId"); ok {
			if err = d.Set("entity_id", vv); err != nil {
				return fmt.Errorf("Error reading entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading entity_id: %v", err)
		}
	}

	if err = d.Set("fabric_force_sync", flattenObjectUserSamlFabricForceSync(o["fabric-force-sync"], d, "fabric_force_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-force-sync"], "ObjectUserSaml-FabricForceSync"); ok {
			if err = d.Set("fabric_force_sync", vv); err != nil {
				return fmt.Errorf("Error reading fabric_force_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_force_sync: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenObjectUserSamlFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "ObjectUserSaml-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("fabric_object_source", flattenObjectUserSamlFabricObjectSource(o["fabric-object-source"], d, "fabric_object_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object-source"], "ObjectUserSaml-FabricObjectSource"); ok {
			if err = d.Set("fabric_object_source", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object_source: %v", err)
		}
	}

	if err = d.Set("group_claim_type", flattenObjectUserSamlGroupClaimType(o["group-claim-type"], d, "group_claim_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-claim-type"], "ObjectUserSaml-GroupClaimType"); ok {
			if err = d.Set("group_claim_type", vv); err != nil {
				return fmt.Errorf("Error reading group_claim_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_claim_type: %v", err)
		}
	}

	if err = d.Set("group_name", flattenObjectUserSamlGroupName(o["group-name"], d, "group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-name"], "ObjectUserSaml-GroupName"); ok {
			if err = d.Set("group_name", vv); err != nil {
				return fmt.Errorf("Error reading group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("idp_cert", flattenObjectUserSamlIdpCert(o["idp-cert"], d, "idp_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-cert"], "ObjectUserSaml-IdpCert"); ok {
			if err = d.Set("idp_cert", vv); err != nil {
				return fmt.Errorf("Error reading idp_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_cert: %v", err)
		}
	}

	if err = d.Set("idp_entity_id", flattenObjectUserSamlIdpEntityId(o["idp-entity-id"], d, "idp_entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-entity-id"], "ObjectUserSaml-IdpEntityId"); ok {
			if err = d.Set("idp_entity_id", vv); err != nil {
				return fmt.Errorf("Error reading idp_entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_entity_id: %v", err)
		}
	}

	if err = d.Set("idp_single_logout_url", flattenObjectUserSamlIdpSingleLogoutUrl(o["idp-single-logout-url"], d, "idp_single_logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-logout-url"], "ObjectUserSaml-IdpSingleLogoutUrl"); ok {
			if err = d.Set("idp_single_logout_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
		}
	}

	if err = d.Set("idp_single_sign_on_url", flattenObjectUserSamlIdpSingleSignOnUrl(o["idp-single-sign-on-url"], d, "idp_single_sign_on_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-sign-on-url"], "ObjectUserSaml-IdpSingleSignOnUrl"); ok {
			if err = d.Set("idp_single_sign_on_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("limit_relaystate", flattenObjectUserSamlLimitRelaystate(o["limit-relaystate"], d, "limit_relaystate")); err != nil {
		if vv, ok := fortiAPIPatch(o["limit-relaystate"], "ObjectUserSaml-LimitRelaystate"); ok {
			if err = d.Set("limit_relaystate", vv); err != nil {
				return fmt.Errorf("Error reading limit_relaystate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading limit_relaystate: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserSamlName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserSaml-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("realm", flattenObjectUserSamlRealm(o["realm"], d, "realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["realm"], "ObjectUserSaml-Realm"); ok {
			if err = d.Set("realm", vv); err != nil {
				return fmt.Errorf("Error reading realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading realm: %v", err)
		}
	}

	if err = d.Set("reauth", flattenObjectUserSamlReauth(o["reauth"], d, "reauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["reauth"], "ObjectUserSaml-Reauth"); ok {
			if err = d.Set("reauth", vv); err != nil {
				return fmt.Errorf("Error reading reauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reauth: %v", err)
		}
	}

	if err = d.Set("require_signed_resp_and_asrt", flattenObjectUserSamlRequireSignedRespAndAsrt(o["require-signed-resp-and-asrt"], d, "require_signed_resp_and_asrt")); err != nil {
		if vv, ok := fortiAPIPatch(o["require-signed-resp-and-asrt"], "ObjectUserSaml-RequireSignedRespAndAsrt"); ok {
			if err = d.Set("require_signed_resp_and_asrt", vv); err != nil {
				return fmt.Errorf("Error reading require_signed_resp_and_asrt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading require_signed_resp_and_asrt: %v", err)
		}
	}

	if err = d.Set("scim_client", flattenObjectUserSamlScimClient(o["scim-client"], d, "scim_client")); err != nil {
		if vv, ok := fortiAPIPatch(o["scim-client"], "ObjectUserSaml-ScimClient"); ok {
			if err = d.Set("scim_client", vv); err != nil {
				return fmt.Errorf("Error reading scim_client: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scim_client: %v", err)
		}
	}

	if err = d.Set("scim_group_attr_type", flattenObjectUserSamlScimGroupAttrType(o["scim-group-attr-type"], d, "scim_group_attr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["scim-group-attr-type"], "ObjectUserSaml-ScimGroupAttrType"); ok {
			if err = d.Set("scim_group_attr_type", vv); err != nil {
				return fmt.Errorf("Error reading scim_group_attr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scim_group_attr_type: %v", err)
		}
	}

	if err = d.Set("scim_user_attr_type", flattenObjectUserSamlScimUserAttrType(o["scim-user-attr-type"], d, "scim_user_attr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["scim-user-attr-type"], "ObjectUserSaml-ScimUserAttrType"); ok {
			if err = d.Set("scim_user_attr_type", vv); err != nil {
				return fmt.Errorf("Error reading scim_user_attr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scim_user_attr_type: %v", err)
		}
	}

	if err = d.Set("service_provider_address", flattenObjectUserSamlServiceProviderAddress(o["service-provider-address"], d, "service_provider_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-provider-address"], "ObjectUserSaml-ServiceProviderAddress"); ok {
			if err = d.Set("service_provider_address", vv); err != nil {
				return fmt.Errorf("Error reading service_provider_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_provider_address: %v", err)
		}
	}

	if err = d.Set("single_logout_url", flattenObjectUserSamlSingleLogoutUrl(o["single-logout-url"], d, "single_logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["single-logout-url"], "ObjectUserSaml-SingleLogoutUrl"); ok {
			if err = d.Set("single_logout_url", vv); err != nil {
				return fmt.Errorf("Error reading single_logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading single_logout_url: %v", err)
		}
	}

	if err = d.Set("single_sign_on_url", flattenObjectUserSamlSingleSignOnUrl(o["single-sign-on-url"], d, "single_sign_on_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["single-sign-on-url"], "ObjectUserSaml-SingleSignOnUrl"); ok {
			if err = d.Set("single_sign_on_url", vv); err != nil {
				return fmt.Errorf("Error reading single_sign_on_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("sso_app_id", flattenObjectUserSamlSsoAppId(o["sso-app-id"], d, "sso_app_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-app-id"], "ObjectUserSaml-SsoAppId"); ok {
			if err = d.Set("sso_app_id", vv); err != nil {
				return fmt.Errorf("Error reading sso_app_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_app_id: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectUserSamlType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectUserSaml-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user_claim_type", flattenObjectUserSamlUserClaimType(o["user-claim-type"], d, "user_claim_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-claim-type"], "ObjectUserSaml-UserClaimType"); ok {
			if err = d.Set("user_claim_type", vv); err != nil {
				return fmt.Errorf("Error reading user_claim_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_claim_type: %v", err)
		}
	}

	if err = d.Set("user_name", flattenObjectUserSamlUserName(o["user-name"], d, "user_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-name"], "ObjectUserSaml-UserName"); ok {
			if err = d.Set("user_name", vv); err != nil {
				return fmt.Errorf("Error reading user_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_name: %v", err)
		}
	}

	if err = d.Set("user_source", flattenObjectUserSamlUserSource(o["user-source"], d, "user_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-source"], "ObjectUserSaml-UserSource"); ok {
			if err = d.Set("user_source", vv); err != nil {
				return fmt.Errorf("Error reading user_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_source: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectUserSamlUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectUserSaml-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenObjectUserSamlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserSamlAdfsClaim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlAuthUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserSamlClockTolerance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDigestMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			t, err := expandObjectUserSamlDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adfs_claim"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adfs-claim"], _ = expandObjectUserSamlDynamicMappingAdfsClaim(d, i["adfs_claim"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-url"], _ = expandObjectUserSamlDynamicMappingAuthUrl(d, i["auth_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cert"], _ = expandObjectUserSamlDynamicMappingCert(d, i["cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "clock_tolerance"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["clock-tolerance"], _ = expandObjectUserSamlDynamicMappingClockTolerance(d, i["clock_tolerance"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "digest_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["digest-method"], _ = expandObjectUserSamlDynamicMappingDigestMethod(d, i["digest_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entity_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["entity-id"], _ = expandObjectUserSamlDynamicMappingEntityId(d, i["entity_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_force_sync"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fabric-force-sync"], _ = expandObjectUserSamlDynamicMappingFabricForceSync(d, i["fabric_force_sync"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fabric-object"], _ = expandObjectUserSamlDynamicMappingFabricObject(d, i["fabric_object"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object_source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fabric-object-source"], _ = expandObjectUserSamlDynamicMappingFabricObjectSource(d, i["fabric_object_source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_claim_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-claim-type"], _ = expandObjectUserSamlDynamicMappingGroupClaimType(d, i["group_claim_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-name"], _ = expandObjectUserSamlDynamicMappingGroupName(d, i["group_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-cert"], _ = expandObjectUserSamlDynamicMappingIdpCert(d, i["idp_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_entity_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-entity-id"], _ = expandObjectUserSamlDynamicMappingIdpEntityId(d, i["idp_entity_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-single-logout-url"], _ = expandObjectUserSamlDynamicMappingIdpSingleLogoutUrl(d, i["idp_single_logout_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-single-sign-on-url"], _ = expandObjectUserSamlDynamicMappingIdpSingleSignOnUrl(d, i["idp_single_sign_on_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "limit_relaystate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["limit-relaystate"], _ = expandObjectUserSamlDynamicMappingLimitRelaystate(d, i["limit_relaystate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["realm"], _ = expandObjectUserSamlDynamicMappingRealm(d, i["realm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "reauth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["reauth"], _ = expandObjectUserSamlDynamicMappingReauth(d, i["reauth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "require_signed_resp_and_asrt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["require-signed-resp-and-asrt"], _ = expandObjectUserSamlDynamicMappingRequireSignedRespAndAsrt(d, i["require_signed_resp_and_asrt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scim_client"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["scim-client"], _ = expandObjectUserSamlDynamicMappingScimClient(d, i["scim_client"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scim_group_attr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["scim-group-attr-type"], _ = expandObjectUserSamlDynamicMappingScimGroupAttrType(d, i["scim_group_attr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scim_user_attr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["scim-user-attr-type"], _ = expandObjectUserSamlDynamicMappingScimUserAttrType(d, i["scim_user_attr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_provider_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-provider-address"], _ = expandObjectUserSamlDynamicMappingServiceProviderAddress(d, i["service_provider_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "single_logout_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["single-logout-url"], _ = expandObjectUserSamlDynamicMappingSingleLogoutUrl(d, i["single_logout_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "single_sign_on_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["single-sign-on-url"], _ = expandObjectUserSamlDynamicMappingSingleSignOnUrl(d, i["single_sign_on_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_app_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-app-id"], _ = expandObjectUserSamlDynamicMappingSsoAppId(d, i["sso_app_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectUserSamlDynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_claim_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-claim-type"], _ = expandObjectUserSamlDynamicMappingUserClaimType(d, i["user_claim_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-name"], _ = expandObjectUserSamlDynamicMappingUserName(d, i["user_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-source"], _ = expandObjectUserSamlDynamicMappingUserSource(d, i["user_source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uuid"], _ = expandObjectUserSamlDynamicMappingUuid(d, i["uuid"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserSamlDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectUserSamlDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectUserSamlDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserSamlDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingAdfsClaim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingAuthUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserSamlDynamicMappingClockTolerance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingDigestMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingFabricForceSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingFabricObjectSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingGroupClaimType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingIdpCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserSamlDynamicMappingIdpEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingIdpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingIdpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingLimitRelaystate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingRequireSignedRespAndAsrt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingScimClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserSamlDynamicMappingScimGroupAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingScimUserAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingServiceProviderAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingSsoAppId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingUserClaimType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingUserSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlFabricForceSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlFabricObjectSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlGroupClaimType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlIdpCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserSamlIdpEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlIdpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlIdpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlLimitRelaystate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlRequireSignedRespAndAsrt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlScimClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserSamlScimGroupAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlScimUserAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlServiceProviderAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlSsoAppId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlUserClaimType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlUserSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserSaml(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("adfs_claim"); ok || d.HasChange("adfs_claim") {
		t, err := expandObjectUserSamlAdfsClaim(d, v, "adfs_claim")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adfs-claim"] = t
		}
	}

	if v, ok := d.GetOk("auth_url"); ok || d.HasChange("auth_url") {
		t, err := expandObjectUserSamlAuthUrl(d, v, "auth_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-url"] = t
		}
	}

	if v, ok := d.GetOk("cert"); ok || d.HasChange("cert") {
		t, err := expandObjectUserSamlCert(d, v, "cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("clock_tolerance"); ok || d.HasChange("clock_tolerance") {
		t, err := expandObjectUserSamlClockTolerance(d, v, "clock_tolerance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clock-tolerance"] = t
		}
	}

	if v, ok := d.GetOk("digest_method"); ok || d.HasChange("digest_method") {
		t, err := expandObjectUserSamlDigestMethod(d, v, "digest_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digest-method"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectUserSamlDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("entity_id"); ok || d.HasChange("entity_id") {
		t, err := expandObjectUserSamlEntityId(d, v, "entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entity-id"] = t
		}
	}

	if v, ok := d.GetOk("fabric_force_sync"); ok || d.HasChange("fabric_force_sync") {
		t, err := expandObjectUserSamlFabricForceSync(d, v, "fabric_force_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-force-sync"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandObjectUserSamlFabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object_source"); ok || d.HasChange("fabric_object_source") {
		t, err := expandObjectUserSamlFabricObjectSource(d, v, "fabric_object_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object-source"] = t
		}
	}

	if v, ok := d.GetOk("group_claim_type"); ok || d.HasChange("group_claim_type") {
		t, err := expandObjectUserSamlGroupClaimType(d, v, "group_claim_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-claim-type"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok || d.HasChange("group_name") {
		t, err := expandObjectUserSamlGroupName(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("idp_cert"); ok || d.HasChange("idp_cert") {
		t, err := expandObjectUserSamlIdpCert(d, v, "idp_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-cert"] = t
		}
	}

	if v, ok := d.GetOk("idp_entity_id"); ok || d.HasChange("idp_entity_id") {
		t, err := expandObjectUserSamlIdpEntityId(d, v, "idp_entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-entity-id"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_logout_url"); ok || d.HasChange("idp_single_logout_url") {
		t, err := expandObjectUserSamlIdpSingleLogoutUrl(d, v, "idp_single_logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_sign_on_url"); ok || d.HasChange("idp_single_sign_on_url") {
		t, err := expandObjectUserSamlIdpSingleSignOnUrl(d, v, "idp_single_sign_on_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-sign-on-url"] = t
		}
	}

	if v, ok := d.GetOk("limit_relaystate"); ok || d.HasChange("limit_relaystate") {
		t, err := expandObjectUserSamlLimitRelaystate(d, v, "limit_relaystate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["limit-relaystate"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserSamlName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("realm"); ok || d.HasChange("realm") {
		t, err := expandObjectUserSamlRealm(d, v, "realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realm"] = t
		}
	}

	if v, ok := d.GetOk("reauth"); ok || d.HasChange("reauth") {
		t, err := expandObjectUserSamlReauth(d, v, "reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reauth"] = t
		}
	}

	if v, ok := d.GetOk("require_signed_resp_and_asrt"); ok || d.HasChange("require_signed_resp_and_asrt") {
		t, err := expandObjectUserSamlRequireSignedRespAndAsrt(d, v, "require_signed_resp_and_asrt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["require-signed-resp-and-asrt"] = t
		}
	}

	if v, ok := d.GetOk("scim_client"); ok || d.HasChange("scim_client") {
		t, err := expandObjectUserSamlScimClient(d, v, "scim_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scim-client"] = t
		}
	}

	if v, ok := d.GetOk("scim_group_attr_type"); ok || d.HasChange("scim_group_attr_type") {
		t, err := expandObjectUserSamlScimGroupAttrType(d, v, "scim_group_attr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scim-group-attr-type"] = t
		}
	}

	if v, ok := d.GetOk("scim_user_attr_type"); ok || d.HasChange("scim_user_attr_type") {
		t, err := expandObjectUserSamlScimUserAttrType(d, v, "scim_user_attr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scim-user-attr-type"] = t
		}
	}

	if v, ok := d.GetOk("service_provider_address"); ok || d.HasChange("service_provider_address") {
		t, err := expandObjectUserSamlServiceProviderAddress(d, v, "service_provider_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-provider-address"] = t
		}
	}

	if v, ok := d.GetOk("single_logout_url"); ok || d.HasChange("single_logout_url") {
		t, err := expandObjectUserSamlSingleLogoutUrl(d, v, "single_logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("single_sign_on_url"); ok || d.HasChange("single_sign_on_url") {
		t, err := expandObjectUserSamlSingleSignOnUrl(d, v, "single_sign_on_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-sign-on-url"] = t
		}
	}

	if v, ok := d.GetOk("sso_app_id"); ok || d.HasChange("sso_app_id") {
		t, err := expandObjectUserSamlSsoAppId(d, v, "sso_app_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-app-id"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectUserSamlType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user_claim_type"); ok || d.HasChange("user_claim_type") {
		t, err := expandObjectUserSamlUserClaimType(d, v, "user_claim_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-claim-type"] = t
		}
	}

	if v, ok := d.GetOk("user_name"); ok || d.HasChange("user_name") {
		t, err := expandObjectUserSamlUserName(d, v, "user_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-name"] = t
		}
	}

	if v, ok := d.GetOk("user_source"); ok || d.HasChange("user_source") {
		t, err := expandObjectUserSamlUserSource(d, v, "user_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-source"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandObjectUserSamlUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
