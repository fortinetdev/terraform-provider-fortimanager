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
						"reauth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"single_logout_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"single_sign_on_url": &schema.Schema{
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
					},
				},
			},
			"entity_id": &schema.Schema{
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
			"reauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"single_logout_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"single_sign_on_url": &schema.Schema{
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

	_, err = c.CreateObjectUserSaml(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserSaml resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserSamlRead(d, m)
}

func resourceObjectUserSamlUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserSaml(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserSaml resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserSaml(obj, mkey, paradict)
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
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	err = c.DeleteObjectUserSaml(mkey, paradict)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "reauth"
		if _, ok := i["reauth"]; ok {
			v := flattenObjectUserSamlDynamicMappingReauth(i["reauth"], d, pre_append)
			tmp["reauth"] = fortiAPISubPartPatch(v, "ObjectUserSaml-DynamicMapping-Reauth")
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

func flattenObjectUserSamlDynamicMappingReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingUserClaimType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlDynamicMappingUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectUserSamlReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlUserClaimType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSamlUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	if err = d.Set("reauth", flattenObjectUserSamlReauth(o["reauth"], d, "reauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["reauth"], "ObjectUserSaml-Reauth"); ok {
			if err = d.Set("reauth", vv); err != nil {
				return fmt.Errorf("Error reading reauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reauth: %v", err)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "reauth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["reauth"], _ = expandObjectUserSamlDynamicMappingReauth(d, i["reauth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "single_logout_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["single-logout-url"], _ = expandObjectUserSamlDynamicMappingSingleLogoutUrl(d, i["single_logout_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "single_sign_on_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["single-sign-on-url"], _ = expandObjectUserSamlDynamicMappingSingleSignOnUrl(d, i["single_sign_on_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_claim_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-claim-type"], _ = expandObjectUserSamlDynamicMappingUserClaimType(d, i["user_claim_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-name"], _ = expandObjectUserSamlDynamicMappingUserName(d, i["user_name"], pre_append)
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

func expandObjectUserSamlDynamicMappingReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingUserClaimType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlDynamicMappingUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

func expandObjectUserSamlReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlUserClaimType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSamlUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

	if v, ok := d.GetOk("reauth"); ok || d.HasChange("reauth") {
		t, err := expandObjectUserSamlReauth(d, v, "reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reauth"] = t
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

	return &obj, nil
}
