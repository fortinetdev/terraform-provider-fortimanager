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

func resourceObjectUserSamlDynamicMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserSamlDynamicMappingCreate,
		Read:   resourceObjectUserSamlDynamicMappingRead,
		Update: resourceObjectUserSamlDynamicMappingUpdate,
		Delete: resourceObjectUserSamlDynamicMappingDelete,

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
			"saml": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectUserSamlDynamicMappingCreate(d *schema.ResourceData, m interface{}) error {
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

	saml := d.Get("saml").(string)
	paradict["saml"] = saml

	obj, err := getObjectObjectUserSamlDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserSamlDynamicMapping resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("_scope")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectUserSamlDynamicMapping(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectUserSamlDynamicMapping(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectUserSamlDynamicMapping resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectUserSamlDynamicMapping(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectUserSamlDynamicMapping resource: %v", err)
		}

	}

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectUserSamlDynamicMappingRead(d, m)
}

func resourceObjectUserSamlDynamicMappingUpdate(d *schema.ResourceData, m interface{}) error {
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

	saml := d.Get("saml").(string)
	paradict["saml"] = saml

	obj, err := getObjectObjectUserSamlDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserSamlDynamicMapping resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectUserSamlDynamicMapping(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserSamlDynamicMapping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectUserSamlDynamicMappingRead(d, m)
}

func resourceObjectUserSamlDynamicMappingDelete(d *schema.ResourceData, m interface{}) error {
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

	saml := d.Get("saml").(string)
	paradict["saml"] = saml

	wsParams["adom"] = adomv

	err = c.DeleteObjectUserSamlDynamicMapping(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserSamlDynamicMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserSamlDynamicMappingRead(d *schema.ResourceData, m interface{}) error {
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

	saml := d.Get("saml").(string)
	if saml == "" {
		saml = importOptionChecking(m.(*FortiClient).Cfg, "saml")
		if saml == "" {
			return fmt.Errorf("Parameter saml is missing")
		}
		if err = d.Set("saml", saml); err != nil {
			return fmt.Errorf("Error set params saml: %v", err)
		}
	}
	if mkey, err = checkScopeId(mkey); err != nil {
		return fmt.Errorf("Error set ID : %v", err)
	}
	paradict["saml"] = saml

	o, err := c.ReadObjectUserSamlDynamicMapping(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectUserSamlDynamicMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserSamlDynamicMapping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserSamlDynamicMapping resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserSamlDynamicMappingScope2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserSamlDynamicMappingScopeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserSamlDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectUserSamlDynamicMappingScopeVdom2edl(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectUserSamlDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserSamlDynamicMappingScopeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingScopeVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingAdfsClaim2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingAuthUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserSamlDynamicMappingClockTolerance2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingDigestMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingEntityId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingFabricForceSync2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingFabricObject2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingFabricObjectSource2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingGroupClaimType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingGroupName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingIdpCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserSamlDynamicMappingIdpEntityId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingIdpSingleLogoutUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingIdpSingleSignOnUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingLimitRelaystate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingRealm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingReauth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingRequireSignedRespAndAsrt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingScimClient2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserSamlDynamicMappingScimGroupAttrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingScimUserAttrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingServiceProviderAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingSingleLogoutUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingSingleSignOnUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingSsoAppId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingUserClaimType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingUserName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingUserSource2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingUuid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserSamlDynamicMapping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("_scope", flattenObjectUserSamlDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["_scope"], "ObjectUserSamlDynamicMapping-Scope"); ok {
				if err = d.Set("_scope", vv); err != nil {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading _scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("_scope"); ok {
			if err = d.Set("_scope", flattenObjectUserSamlDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["_scope"], "ObjectUserSamlDynamicMapping-Scope"); ok {
					if err = d.Set("_scope", vv); err != nil {
						return fmt.Errorf("Error reading _scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			}
		}
	}

	if err = d.Set("adfs_claim", flattenObjectUserSamlDynamicMappingAdfsClaim2edl(o["adfs-claim"], d, "adfs_claim")); err != nil {
		if vv, ok := fortiAPIPatch(o["adfs-claim"], "ObjectUserSamlDynamicMapping-AdfsClaim"); ok {
			if err = d.Set("adfs_claim", vv); err != nil {
				return fmt.Errorf("Error reading adfs_claim: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adfs_claim: %v", err)
		}
	}

	if err = d.Set("auth_url", flattenObjectUserSamlDynamicMappingAuthUrl2edl(o["auth-url"], d, "auth_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-url"], "ObjectUserSamlDynamicMapping-AuthUrl"); ok {
			if err = d.Set("auth_url", vv); err != nil {
				return fmt.Errorf("Error reading auth_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_url: %v", err)
		}
	}

	if err = d.Set("cert", flattenObjectUserSamlDynamicMappingCert2edl(o["cert"], d, "cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert"], "ObjectUserSamlDynamicMapping-Cert"); ok {
			if err = d.Set("cert", vv); err != nil {
				return fmt.Errorf("Error reading cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("clock_tolerance", flattenObjectUserSamlDynamicMappingClockTolerance2edl(o["clock-tolerance"], d, "clock_tolerance")); err != nil {
		if vv, ok := fortiAPIPatch(o["clock-tolerance"], "ObjectUserSamlDynamicMapping-ClockTolerance"); ok {
			if err = d.Set("clock_tolerance", vv); err != nil {
				return fmt.Errorf("Error reading clock_tolerance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clock_tolerance: %v", err)
		}
	}

	if err = d.Set("digest_method", flattenObjectUserSamlDynamicMappingDigestMethod2edl(o["digest-method"], d, "digest_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["digest-method"], "ObjectUserSamlDynamicMapping-DigestMethod"); ok {
			if err = d.Set("digest_method", vv); err != nil {
				return fmt.Errorf("Error reading digest_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading digest_method: %v", err)
		}
	}

	if err = d.Set("entity_id", flattenObjectUserSamlDynamicMappingEntityId2edl(o["entity-id"], d, "entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["entity-id"], "ObjectUserSamlDynamicMapping-EntityId"); ok {
			if err = d.Set("entity_id", vv); err != nil {
				return fmt.Errorf("Error reading entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading entity_id: %v", err)
		}
	}

	if err = d.Set("fabric_force_sync", flattenObjectUserSamlDynamicMappingFabricForceSync2edl(o["fabric-force-sync"], d, "fabric_force_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-force-sync"], "ObjectUserSamlDynamicMapping-FabricForceSync"); ok {
			if err = d.Set("fabric_force_sync", vv); err != nil {
				return fmt.Errorf("Error reading fabric_force_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_force_sync: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenObjectUserSamlDynamicMappingFabricObject2edl(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "ObjectUserSamlDynamicMapping-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("fabric_object_source", flattenObjectUserSamlDynamicMappingFabricObjectSource2edl(o["fabric-object-source"], d, "fabric_object_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object-source"], "ObjectUserSamlDynamicMapping-FabricObjectSource"); ok {
			if err = d.Set("fabric_object_source", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object_source: %v", err)
		}
	}

	if err = d.Set("group_claim_type", flattenObjectUserSamlDynamicMappingGroupClaimType2edl(o["group-claim-type"], d, "group_claim_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-claim-type"], "ObjectUserSamlDynamicMapping-GroupClaimType"); ok {
			if err = d.Set("group_claim_type", vv); err != nil {
				return fmt.Errorf("Error reading group_claim_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_claim_type: %v", err)
		}
	}

	if err = d.Set("group_name", flattenObjectUserSamlDynamicMappingGroupName2edl(o["group-name"], d, "group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-name"], "ObjectUserSamlDynamicMapping-GroupName"); ok {
			if err = d.Set("group_name", vv); err != nil {
				return fmt.Errorf("Error reading group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("idp_cert", flattenObjectUserSamlDynamicMappingIdpCert2edl(o["idp-cert"], d, "idp_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-cert"], "ObjectUserSamlDynamicMapping-IdpCert"); ok {
			if err = d.Set("idp_cert", vv); err != nil {
				return fmt.Errorf("Error reading idp_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_cert: %v", err)
		}
	}

	if err = d.Set("idp_entity_id", flattenObjectUserSamlDynamicMappingIdpEntityId2edl(o["idp-entity-id"], d, "idp_entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-entity-id"], "ObjectUserSamlDynamicMapping-IdpEntityId"); ok {
			if err = d.Set("idp_entity_id", vv); err != nil {
				return fmt.Errorf("Error reading idp_entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_entity_id: %v", err)
		}
	}

	if err = d.Set("idp_single_logout_url", flattenObjectUserSamlDynamicMappingIdpSingleLogoutUrl2edl(o["idp-single-logout-url"], d, "idp_single_logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-logout-url"], "ObjectUserSamlDynamicMapping-IdpSingleLogoutUrl"); ok {
			if err = d.Set("idp_single_logout_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
		}
	}

	if err = d.Set("idp_single_sign_on_url", flattenObjectUserSamlDynamicMappingIdpSingleSignOnUrl2edl(o["idp-single-sign-on-url"], d, "idp_single_sign_on_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-sign-on-url"], "ObjectUserSamlDynamicMapping-IdpSingleSignOnUrl"); ok {
			if err = d.Set("idp_single_sign_on_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("limit_relaystate", flattenObjectUserSamlDynamicMappingLimitRelaystate2edl(o["limit-relaystate"], d, "limit_relaystate")); err != nil {
		if vv, ok := fortiAPIPatch(o["limit-relaystate"], "ObjectUserSamlDynamicMapping-LimitRelaystate"); ok {
			if err = d.Set("limit_relaystate", vv); err != nil {
				return fmt.Errorf("Error reading limit_relaystate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading limit_relaystate: %v", err)
		}
	}

	if err = d.Set("realm", flattenObjectUserSamlDynamicMappingRealm2edl(o["realm"], d, "realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["realm"], "ObjectUserSamlDynamicMapping-Realm"); ok {
			if err = d.Set("realm", vv); err != nil {
				return fmt.Errorf("Error reading realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading realm: %v", err)
		}
	}

	if err = d.Set("reauth", flattenObjectUserSamlDynamicMappingReauth2edl(o["reauth"], d, "reauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["reauth"], "ObjectUserSamlDynamicMapping-Reauth"); ok {
			if err = d.Set("reauth", vv); err != nil {
				return fmt.Errorf("Error reading reauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reauth: %v", err)
		}
	}

	if err = d.Set("require_signed_resp_and_asrt", flattenObjectUserSamlDynamicMappingRequireSignedRespAndAsrt2edl(o["require-signed-resp-and-asrt"], d, "require_signed_resp_and_asrt")); err != nil {
		if vv, ok := fortiAPIPatch(o["require-signed-resp-and-asrt"], "ObjectUserSamlDynamicMapping-RequireSignedRespAndAsrt"); ok {
			if err = d.Set("require_signed_resp_and_asrt", vv); err != nil {
				return fmt.Errorf("Error reading require_signed_resp_and_asrt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading require_signed_resp_and_asrt: %v", err)
		}
	}

	if err = d.Set("scim_client", flattenObjectUserSamlDynamicMappingScimClient2edl(o["scim-client"], d, "scim_client")); err != nil {
		if vv, ok := fortiAPIPatch(o["scim-client"], "ObjectUserSamlDynamicMapping-ScimClient"); ok {
			if err = d.Set("scim_client", vv); err != nil {
				return fmt.Errorf("Error reading scim_client: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scim_client: %v", err)
		}
	}

	if err = d.Set("scim_group_attr_type", flattenObjectUserSamlDynamicMappingScimGroupAttrType2edl(o["scim-group-attr-type"], d, "scim_group_attr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["scim-group-attr-type"], "ObjectUserSamlDynamicMapping-ScimGroupAttrType"); ok {
			if err = d.Set("scim_group_attr_type", vv); err != nil {
				return fmt.Errorf("Error reading scim_group_attr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scim_group_attr_type: %v", err)
		}
	}

	if err = d.Set("scim_user_attr_type", flattenObjectUserSamlDynamicMappingScimUserAttrType2edl(o["scim-user-attr-type"], d, "scim_user_attr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["scim-user-attr-type"], "ObjectUserSamlDynamicMapping-ScimUserAttrType"); ok {
			if err = d.Set("scim_user_attr_type", vv); err != nil {
				return fmt.Errorf("Error reading scim_user_attr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scim_user_attr_type: %v", err)
		}
	}

	if err = d.Set("service_provider_address", flattenObjectUserSamlDynamicMappingServiceProviderAddress2edl(o["service-provider-address"], d, "service_provider_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-provider-address"], "ObjectUserSamlDynamicMapping-ServiceProviderAddress"); ok {
			if err = d.Set("service_provider_address", vv); err != nil {
				return fmt.Errorf("Error reading service_provider_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_provider_address: %v", err)
		}
	}

	if err = d.Set("single_logout_url", flattenObjectUserSamlDynamicMappingSingleLogoutUrl2edl(o["single-logout-url"], d, "single_logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["single-logout-url"], "ObjectUserSamlDynamicMapping-SingleLogoutUrl"); ok {
			if err = d.Set("single_logout_url", vv); err != nil {
				return fmt.Errorf("Error reading single_logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading single_logout_url: %v", err)
		}
	}

	if err = d.Set("single_sign_on_url", flattenObjectUserSamlDynamicMappingSingleSignOnUrl2edl(o["single-sign-on-url"], d, "single_sign_on_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["single-sign-on-url"], "ObjectUserSamlDynamicMapping-SingleSignOnUrl"); ok {
			if err = d.Set("single_sign_on_url", vv); err != nil {
				return fmt.Errorf("Error reading single_sign_on_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("sso_app_id", flattenObjectUserSamlDynamicMappingSsoAppId2edl(o["sso-app-id"], d, "sso_app_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-app-id"], "ObjectUserSamlDynamicMapping-SsoAppId"); ok {
			if err = d.Set("sso_app_id", vv); err != nil {
				return fmt.Errorf("Error reading sso_app_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_app_id: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectUserSamlDynamicMappingType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectUserSamlDynamicMapping-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user_claim_type", flattenObjectUserSamlDynamicMappingUserClaimType2edl(o["user-claim-type"], d, "user_claim_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-claim-type"], "ObjectUserSamlDynamicMapping-UserClaimType"); ok {
			if err = d.Set("user_claim_type", vv); err != nil {
				return fmt.Errorf("Error reading user_claim_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_claim_type: %v", err)
		}
	}

	if err = d.Set("user_name", flattenObjectUserSamlDynamicMappingUserName2edl(o["user-name"], d, "user_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-name"], "ObjectUserSamlDynamicMapping-UserName"); ok {
			if err = d.Set("user_name", vv); err != nil {
				return fmt.Errorf("Error reading user_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_name: %v", err)
		}
	}

	if err = d.Set("user_source", flattenObjectUserSamlDynamicMappingUserSource2edl(o["user-source"], d, "user_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-source"], "ObjectUserSamlDynamicMapping-UserSource"); ok {
			if err = d.Set("user_source", vv); err != nil {
				return fmt.Errorf("Error reading user_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_source: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectUserSamlDynamicMappingUuid2edl(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectUserSamlDynamicMapping-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenObjectUserSamlDynamicMappingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserSamlDynamicMappingScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectUserSamlDynamicMappingScopeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectUserSamlDynamicMappingScopeVdom2edl(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserSamlDynamicMappingScopeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingScopeVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingAdfsClaim2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingAuthUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserSamlDynamicMappingClockTolerance2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingDigestMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingEntityId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingFabricForceSync2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingFabricObject2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingFabricObjectSource2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingGroupClaimType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingGroupName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingIdpCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserSamlDynamicMappingIdpEntityId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingIdpSingleLogoutUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingIdpSingleSignOnUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingLimitRelaystate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingRealm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingReauth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingRequireSignedRespAndAsrt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingScimClient2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserSamlDynamicMappingScimGroupAttrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingScimUserAttrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingServiceProviderAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingSingleLogoutUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingSingleSignOnUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingSsoAppId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingUserClaimType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingUserName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingUserSource2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingUuid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserSamlDynamicMapping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_scope"); ok || d.HasChange("_scope") {
		t, err := expandObjectUserSamlDynamicMappingScope2edl(d, v, "_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_scope"] = t
		}
	}

	if v, ok := d.GetOk("adfs_claim"); ok || d.HasChange("adfs_claim") {
		t, err := expandObjectUserSamlDynamicMappingAdfsClaim2edl(d, v, "adfs_claim")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adfs-claim"] = t
		}
	}

	if v, ok := d.GetOk("auth_url"); ok || d.HasChange("auth_url") {
		t, err := expandObjectUserSamlDynamicMappingAuthUrl2edl(d, v, "auth_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-url"] = t
		}
	}

	if v, ok := d.GetOk("cert"); ok || d.HasChange("cert") {
		t, err := expandObjectUserSamlDynamicMappingCert2edl(d, v, "cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("clock_tolerance"); ok || d.HasChange("clock_tolerance") {
		t, err := expandObjectUserSamlDynamicMappingClockTolerance2edl(d, v, "clock_tolerance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clock-tolerance"] = t
		}
	}

	if v, ok := d.GetOk("digest_method"); ok || d.HasChange("digest_method") {
		t, err := expandObjectUserSamlDynamicMappingDigestMethod2edl(d, v, "digest_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digest-method"] = t
		}
	}

	if v, ok := d.GetOk("entity_id"); ok || d.HasChange("entity_id") {
		t, err := expandObjectUserSamlDynamicMappingEntityId2edl(d, v, "entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entity-id"] = t
		}
	}

	if v, ok := d.GetOk("fabric_force_sync"); ok || d.HasChange("fabric_force_sync") {
		t, err := expandObjectUserSamlDynamicMappingFabricForceSync2edl(d, v, "fabric_force_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-force-sync"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandObjectUserSamlDynamicMappingFabricObject2edl(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object_source"); ok || d.HasChange("fabric_object_source") {
		t, err := expandObjectUserSamlDynamicMappingFabricObjectSource2edl(d, v, "fabric_object_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object-source"] = t
		}
	}

	if v, ok := d.GetOk("group_claim_type"); ok || d.HasChange("group_claim_type") {
		t, err := expandObjectUserSamlDynamicMappingGroupClaimType2edl(d, v, "group_claim_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-claim-type"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok || d.HasChange("group_name") {
		t, err := expandObjectUserSamlDynamicMappingGroupName2edl(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("idp_cert"); ok || d.HasChange("idp_cert") {
		t, err := expandObjectUserSamlDynamicMappingIdpCert2edl(d, v, "idp_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-cert"] = t
		}
	}

	if v, ok := d.GetOk("idp_entity_id"); ok || d.HasChange("idp_entity_id") {
		t, err := expandObjectUserSamlDynamicMappingIdpEntityId2edl(d, v, "idp_entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-entity-id"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_logout_url"); ok || d.HasChange("idp_single_logout_url") {
		t, err := expandObjectUserSamlDynamicMappingIdpSingleLogoutUrl2edl(d, v, "idp_single_logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_sign_on_url"); ok || d.HasChange("idp_single_sign_on_url") {
		t, err := expandObjectUserSamlDynamicMappingIdpSingleSignOnUrl2edl(d, v, "idp_single_sign_on_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-sign-on-url"] = t
		}
	}

	if v, ok := d.GetOk("limit_relaystate"); ok || d.HasChange("limit_relaystate") {
		t, err := expandObjectUserSamlDynamicMappingLimitRelaystate2edl(d, v, "limit_relaystate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["limit-relaystate"] = t
		}
	}

	if v, ok := d.GetOk("realm"); ok || d.HasChange("realm") {
		t, err := expandObjectUserSamlDynamicMappingRealm2edl(d, v, "realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realm"] = t
		}
	}

	if v, ok := d.GetOk("reauth"); ok || d.HasChange("reauth") {
		t, err := expandObjectUserSamlDynamicMappingReauth2edl(d, v, "reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reauth"] = t
		}
	}

	if v, ok := d.GetOk("require_signed_resp_and_asrt"); ok || d.HasChange("require_signed_resp_and_asrt") {
		t, err := expandObjectUserSamlDynamicMappingRequireSignedRespAndAsrt2edl(d, v, "require_signed_resp_and_asrt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["require-signed-resp-and-asrt"] = t
		}
	}

	if v, ok := d.GetOk("scim_client"); ok || d.HasChange("scim_client") {
		t, err := expandObjectUserSamlDynamicMappingScimClient2edl(d, v, "scim_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scim-client"] = t
		}
	}

	if v, ok := d.GetOk("scim_group_attr_type"); ok || d.HasChange("scim_group_attr_type") {
		t, err := expandObjectUserSamlDynamicMappingScimGroupAttrType2edl(d, v, "scim_group_attr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scim-group-attr-type"] = t
		}
	}

	if v, ok := d.GetOk("scim_user_attr_type"); ok || d.HasChange("scim_user_attr_type") {
		t, err := expandObjectUserSamlDynamicMappingScimUserAttrType2edl(d, v, "scim_user_attr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scim-user-attr-type"] = t
		}
	}

	if v, ok := d.GetOk("service_provider_address"); ok || d.HasChange("service_provider_address") {
		t, err := expandObjectUserSamlDynamicMappingServiceProviderAddress2edl(d, v, "service_provider_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-provider-address"] = t
		}
	}

	if v, ok := d.GetOk("single_logout_url"); ok || d.HasChange("single_logout_url") {
		t, err := expandObjectUserSamlDynamicMappingSingleLogoutUrl2edl(d, v, "single_logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("single_sign_on_url"); ok || d.HasChange("single_sign_on_url") {
		t, err := expandObjectUserSamlDynamicMappingSingleSignOnUrl2edl(d, v, "single_sign_on_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-sign-on-url"] = t
		}
	}

	if v, ok := d.GetOk("sso_app_id"); ok || d.HasChange("sso_app_id") {
		t, err := expandObjectUserSamlDynamicMappingSsoAppId2edl(d, v, "sso_app_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-app-id"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectUserSamlDynamicMappingType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user_claim_type"); ok || d.HasChange("user_claim_type") {
		t, err := expandObjectUserSamlDynamicMappingUserClaimType2edl(d, v, "user_claim_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-claim-type"] = t
		}
	}

	if v, ok := d.GetOk("user_name"); ok || d.HasChange("user_name") {
		t, err := expandObjectUserSamlDynamicMappingUserName2edl(d, v, "user_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-name"] = t
		}
	}

	if v, ok := d.GetOk("user_source"); ok || d.HasChange("user_source") {
		t, err := expandObjectUserSamlDynamicMappingUserSource2edl(d, v, "user_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-source"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandObjectUserSamlDynamicMappingUuid2edl(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
