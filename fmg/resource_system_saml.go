// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Global settings for SAML authentication.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSaml() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSamlUpdate,
		Read:   resourceSystemSamlRead,
		Update: resourceSystemSamlUpdate,
		Delete: resourceSystemSamlDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"acs_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_request_signed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_idp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dev_id": &schema.Schema{
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
						"idp_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"forticloud_sso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"login_auto_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_providers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sp_adom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sp_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sp_entity_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sp_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sp_single_logout_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sp_single_sign_on_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"sls_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_auto_create": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"want_assertions_signed": &schema.Schema{
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

func resourceSystemSamlUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSaml(d, false)
	if err != nil {
		return fmt.Errorf("Error updating SystemSaml resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemSaml(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSaml resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSaml")

	return resourceSystemSamlRead(d, m)
}

func resourceSystemSamlDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSaml(d, true)

	if err != nil {
		return fmt.Errorf("Error updating SystemSaml resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemSaml(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing SystemSaml resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSamlRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemSaml(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSaml resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSaml(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSaml resource from API: %v", err)
	}
	return nil
}

func flattenSystemSamlAcsUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlAuthRequestSigned(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlDefaultProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dev_id"
		if _, ok := i["dev-id"]; ok {
			v := flattenSystemSamlFabricIdpDevId(i["dev-id"], d, pre_append)
			tmp["dev_id"] = fortiAPISubPartPatch(v, "SystemSaml-FabricIdp-DevId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_cert"
		if _, ok := i["idp-cert"]; ok {
			v := flattenSystemSamlFabricIdpIdpCert(i["idp-cert"], d, pre_append)
			tmp["idp_cert"] = fortiAPISubPartPatch(v, "SystemSaml-FabricIdp-IdpCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_entity_id"
		if _, ok := i["idp-entity-id"]; ok {
			v := flattenSystemSamlFabricIdpIdpEntityId(i["idp-entity-id"], d, pre_append)
			tmp["idp_entity_id"] = fortiAPISubPartPatch(v, "SystemSaml-FabricIdp-IdpEntityId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := i["idp-single-logout-url"]; ok {
			v := flattenSystemSamlFabricIdpIdpSingleLogoutUrl(i["idp-single-logout-url"], d, pre_append)
			tmp["idp_single_logout_url"] = fortiAPISubPartPatch(v, "SystemSaml-FabricIdp-IdpSingleLogoutUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := i["idp-single-sign-on-url"]; ok {
			v := flattenSystemSamlFabricIdpIdpSingleSignOnUrl(i["idp-single-sign-on-url"], d, pre_append)
			tmp["idp_single_sign_on_url"] = fortiAPISubPartPatch(v, "SystemSaml-FabricIdp-IdpSingleSignOnUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_status"
		if _, ok := i["idp-status"]; ok {
			v := flattenSystemSamlFabricIdpIdpStatus(i["idp-status"], d, pre_append)
			tmp["idp_status"] = fortiAPISubPartPatch(v, "SystemSaml-FabricIdp-IdpStatus")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSamlFabricIdpDevId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpIdpCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpIdpEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpIdpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpIdpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpIdpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlForticloudSso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlIdpCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlIdpEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlIdpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlIdpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlLoginAutoRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServerAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProviders(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_entity_id"
		if _, ok := i["idp-entity-id"]; ok {
			v := flattenSystemSamlServiceProvidersIdpEntityId(i["idp-entity-id"], d, pre_append)
			tmp["idp_entity_id"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-IdpEntityId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := i["idp-single-logout-url"]; ok {
			v := flattenSystemSamlServiceProvidersIdpSingleLogoutUrl(i["idp-single-logout-url"], d, pre_append)
			tmp["idp_single_logout_url"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-IdpSingleLogoutUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := i["idp-single-sign-on-url"]; ok {
			v := flattenSystemSamlServiceProvidersIdpSingleSignOnUrl(i["idp-single-sign-on-url"], d, pre_append)
			tmp["idp_single_sign_on_url"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-IdpSingleSignOnUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemSamlServiceProvidersName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenSystemSamlServiceProvidersPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_adom"
		if _, ok := i["sp-adom"]; ok {
			v := flattenSystemSamlServiceProvidersSpAdom(i["sp-adom"], d, pre_append)
			tmp["sp_adom"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpAdom")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_cert"
		if _, ok := i["sp-cert"]; ok {
			v := flattenSystemSamlServiceProvidersSpCert(i["sp-cert"], d, pre_append)
			tmp["sp_cert"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_entity_id"
		if _, ok := i["sp-entity-id"]; ok {
			v := flattenSystemSamlServiceProvidersSpEntityId(i["sp-entity-id"], d, pre_append)
			tmp["sp_entity_id"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpEntityId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_profile"
		if _, ok := i["sp-profile"]; ok {
			v := flattenSystemSamlServiceProvidersSpProfile(i["sp-profile"], d, pre_append)
			tmp["sp_profile"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_logout_url"
		if _, ok := i["sp-single-logout-url"]; ok {
			v := flattenSystemSamlServiceProvidersSpSingleLogoutUrl(i["sp-single-logout-url"], d, pre_append)
			tmp["sp_single_logout_url"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpSingleLogoutUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_sign_on_url"
		if _, ok := i["sp-single-sign-on-url"]; ok {
			v := flattenSystemSamlServiceProvidersSpSingleSignOnUrl(i["sp-single-sign-on-url"], d, pre_append)
			tmp["sp_single_sign_on_url"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpSingleSignOnUrl")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSamlServiceProvidersIdpEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersIdpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersIdpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlSlsUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlUserAutoCreate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlWantAssertionsSigned(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSaml(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("acs_url", flattenSystemSamlAcsUrl(o["acs-url"], d, "acs_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["acs-url"], "SystemSaml-AcsUrl"); ok {
			if err = d.Set("acs_url", vv); err != nil {
				return fmt.Errorf("Error reading acs_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acs_url: %v", err)
		}
	}

	if err = d.Set("auth_request_signed", flattenSystemSamlAuthRequestSigned(o["auth-request-signed"], d, "auth_request_signed")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-request-signed"], "SystemSaml-AuthRequestSigned"); ok {
			if err = d.Set("auth_request_signed", vv); err != nil {
				return fmt.Errorf("Error reading auth_request_signed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_request_signed: %v", err)
		}
	}

	if err = d.Set("cert", flattenSystemSamlCert(o["cert"], d, "cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert"], "SystemSaml-Cert"); ok {
			if err = d.Set("cert", vv); err != nil {
				return fmt.Errorf("Error reading cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("default_profile", flattenSystemSamlDefaultProfile(o["default-profile"], d, "default_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-profile"], "SystemSaml-DefaultProfile"); ok {
			if err = d.Set("default_profile", vv); err != nil {
				return fmt.Errorf("Error reading default_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_profile: %v", err)
		}
	}

	if err = d.Set("entity_id", flattenSystemSamlEntityId(o["entity-id"], d, "entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["entity-id"], "SystemSaml-EntityId"); ok {
			if err = d.Set("entity_id", vv); err != nil {
				return fmt.Errorf("Error reading entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading entity_id: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fabric_idp", flattenSystemSamlFabricIdp(o["fabric-idp"], d, "fabric_idp")); err != nil {
			if vv, ok := fortiAPIPatch(o["fabric-idp"], "SystemSaml-FabricIdp"); ok {
				if err = d.Set("fabric_idp", vv); err != nil {
					return fmt.Errorf("Error reading fabric_idp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading fabric_idp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fabric_idp"); ok {
			if err = d.Set("fabric_idp", flattenSystemSamlFabricIdp(o["fabric-idp"], d, "fabric_idp")); err != nil {
				if vv, ok := fortiAPIPatch(o["fabric-idp"], "SystemSaml-FabricIdp"); ok {
					if err = d.Set("fabric_idp", vv); err != nil {
						return fmt.Errorf("Error reading fabric_idp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading fabric_idp: %v", err)
				}
			}
		}
	}

	if err = d.Set("forticloud_sso", flattenSystemSamlForticloudSso(o["forticloud-sso"], d, "forticloud_sso")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticloud-sso"], "SystemSaml-ForticloudSso"); ok {
			if err = d.Set("forticloud_sso", vv); err != nil {
				return fmt.Errorf("Error reading forticloud_sso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticloud_sso: %v", err)
		}
	}

	if err = d.Set("idp_cert", flattenSystemSamlIdpCert(o["idp-cert"], d, "idp_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-cert"], "SystemSaml-IdpCert"); ok {
			if err = d.Set("idp_cert", vv); err != nil {
				return fmt.Errorf("Error reading idp_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_cert: %v", err)
		}
	}

	if err = d.Set("idp_entity_id", flattenSystemSamlIdpEntityId(o["idp-entity-id"], d, "idp_entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-entity-id"], "SystemSaml-IdpEntityId"); ok {
			if err = d.Set("idp_entity_id", vv); err != nil {
				return fmt.Errorf("Error reading idp_entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_entity_id: %v", err)
		}
	}

	if err = d.Set("idp_single_logout_url", flattenSystemSamlIdpSingleLogoutUrl(o["idp-single-logout-url"], d, "idp_single_logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-logout-url"], "SystemSaml-IdpSingleLogoutUrl"); ok {
			if err = d.Set("idp_single_logout_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
		}
	}

	if err = d.Set("idp_single_sign_on_url", flattenSystemSamlIdpSingleSignOnUrl(o["idp-single-sign-on-url"], d, "idp_single_sign_on_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-sign-on-url"], "SystemSaml-IdpSingleSignOnUrl"); ok {
			if err = d.Set("idp_single_sign_on_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("login_auto_redirect", flattenSystemSamlLoginAutoRedirect(o["login-auto-redirect"], d, "login_auto_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-auto-redirect"], "SystemSaml-LoginAutoRedirect"); ok {
			if err = d.Set("login_auto_redirect", vv); err != nil {
				return fmt.Errorf("Error reading login_auto_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_auto_redirect: %v", err)
		}
	}

	if err = d.Set("role", flattenSystemSamlRole(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "SystemSaml-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("server_address", flattenSystemSamlServerAddress(o["server-address"], d, "server_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-address"], "SystemSaml-ServerAddress"); ok {
			if err = d.Set("server_address", vv); err != nil {
				return fmt.Errorf("Error reading server_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_address: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service_providers", flattenSystemSamlServiceProviders(o["service-providers"], d, "service_providers")); err != nil {
			if vv, ok := fortiAPIPatch(o["service-providers"], "SystemSaml-ServiceProviders"); ok {
				if err = d.Set("service_providers", vv); err != nil {
					return fmt.Errorf("Error reading service_providers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading service_providers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service_providers"); ok {
			if err = d.Set("service_providers", flattenSystemSamlServiceProviders(o["service-providers"], d, "service_providers")); err != nil {
				if vv, ok := fortiAPIPatch(o["service-providers"], "SystemSaml-ServiceProviders"); ok {
					if err = d.Set("service_providers", vv); err != nil {
						return fmt.Errorf("Error reading service_providers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading service_providers: %v", err)
				}
			}
		}
	}

	if err = d.Set("sls_url", flattenSystemSamlSlsUrl(o["sls-url"], d, "sls_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["sls-url"], "SystemSaml-SlsUrl"); ok {
			if err = d.Set("sls_url", vv); err != nil {
				return fmt.Errorf("Error reading sls_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sls_url: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSamlStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSaml-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("user_auto_create", flattenSystemSamlUserAutoCreate(o["user-auto-create"], d, "user_auto_create")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-auto-create"], "SystemSaml-UserAutoCreate"); ok {
			if err = d.Set("user_auto_create", vv); err != nil {
				return fmt.Errorf("Error reading user_auto_create: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_auto_create: %v", err)
		}
	}

	if err = d.Set("want_assertions_signed", flattenSystemSamlWantAssertionsSigned(o["want-assertions-signed"], d, "want_assertions_signed")); err != nil {
		if vv, ok := fortiAPIPatch(o["want-assertions-signed"], "SystemSaml-WantAssertionsSigned"); ok {
			if err = d.Set("want_assertions_signed", vv); err != nil {
				return fmt.Errorf("Error reading want_assertions_signed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading want_assertions_signed: %v", err)
		}
	}

	return nil
}

func flattenSystemSamlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSamlAcsUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlAuthRequestSigned(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlDefaultProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dev_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dev-id"], _ = expandSystemSamlFabricIdpDevId(d, i["dev_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-cert"], _ = expandSystemSamlFabricIdpIdpCert(d, i["idp_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_entity_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-entity-id"], _ = expandSystemSamlFabricIdpIdpEntityId(d, i["idp_entity_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-single-logout-url"], _ = expandSystemSamlFabricIdpIdpSingleLogoutUrl(d, i["idp_single_logout_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-single-sign-on-url"], _ = expandSystemSamlFabricIdpIdpSingleSignOnUrl(d, i["idp_single_sign_on_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-status"], _ = expandSystemSamlFabricIdpIdpStatus(d, i["idp_status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSamlFabricIdpDevId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpIdpCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpIdpEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpIdpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpIdpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpIdpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlForticloudSso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlLoginAutoRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServerAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProviders(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_entity_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-entity-id"], _ = expandSystemSamlServiceProvidersIdpEntityId(d, i["idp_entity_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-single-logout-url"], _ = expandSystemSamlServiceProvidersIdpSingleLogoutUrl(d, i["idp_single_logout_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-single-sign-on-url"], _ = expandSystemSamlServiceProvidersIdpSingleSignOnUrl(d, i["idp_single_sign_on_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemSamlServiceProvidersName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandSystemSamlServiceProvidersPrefix(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_adom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-adom"], _ = expandSystemSamlServiceProvidersSpAdom(d, i["sp_adom"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-cert"], _ = expandSystemSamlServiceProvidersSpCert(d, i["sp_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_entity_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-entity-id"], _ = expandSystemSamlServiceProvidersSpEntityId(d, i["sp_entity_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-profile"], _ = expandSystemSamlServiceProvidersSpProfile(d, i["sp_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_logout_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-single-logout-url"], _ = expandSystemSamlServiceProvidersSpSingleLogoutUrl(d, i["sp_single_logout_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_sign_on_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-single-sign-on-url"], _ = expandSystemSamlServiceProvidersSpSingleSignOnUrl(d, i["sp_single_sign_on_url"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSamlServiceProvidersIdpEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersIdpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersIdpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlSlsUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlUserAutoCreate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlWantAssertionsSigned(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSaml(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("acs_url"); ok || d.HasChange("acs_url") {
		t, err := expandSystemSamlAcsUrl(d, v, "acs_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acs-url"] = t
		}
	}

	if v, ok := d.GetOk("auth_request_signed"); ok || d.HasChange("auth_request_signed") {
		t, err := expandSystemSamlAuthRequestSigned(d, v, "auth_request_signed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-request-signed"] = t
		}
	}

	if v, ok := d.GetOk("cert"); ok || d.HasChange("cert") {
		t, err := expandSystemSamlCert(d, v, "cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("default_profile"); ok || d.HasChange("default_profile") {
		t, err := expandSystemSamlDefaultProfile(d, v, "default_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-profile"] = t
		}
	}

	if v, ok := d.GetOk("entity_id"); ok || d.HasChange("entity_id") {
		t, err := expandSystemSamlEntityId(d, v, "entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entity-id"] = t
		}
	}

	if bemptysontable {
		obj["fabric-idp"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("fabric_idp"); ok || d.HasChange("fabric_idp") {
			t, err := expandSystemSamlFabricIdp(d, v, "fabric_idp")
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fabric-idp"] = t
			}
		}
	}

	if v, ok := d.GetOk("forticloud_sso"); ok || d.HasChange("forticloud_sso") {
		t, err := expandSystemSamlForticloudSso(d, v, "forticloud_sso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticloud-sso"] = t
		}
	}

	if v, ok := d.GetOk("idp_cert"); ok || d.HasChange("idp_cert") {
		t, err := expandSystemSamlIdpCert(d, v, "idp_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-cert"] = t
		}
	}

	if v, ok := d.GetOk("idp_entity_id"); ok || d.HasChange("idp_entity_id") {
		t, err := expandSystemSamlIdpEntityId(d, v, "idp_entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-entity-id"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_logout_url"); ok || d.HasChange("idp_single_logout_url") {
		t, err := expandSystemSamlIdpSingleLogoutUrl(d, v, "idp_single_logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_sign_on_url"); ok || d.HasChange("idp_single_sign_on_url") {
		t, err := expandSystemSamlIdpSingleSignOnUrl(d, v, "idp_single_sign_on_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-sign-on-url"] = t
		}
	}

	if v, ok := d.GetOk("login_auto_redirect"); ok || d.HasChange("login_auto_redirect") {
		t, err := expandSystemSamlLoginAutoRedirect(d, v, "login_auto_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-auto-redirect"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandSystemSamlRole(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("server_address"); ok || d.HasChange("server_address") {
		t, err := expandSystemSamlServerAddress(d, v, "server_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-address"] = t
		}
	}

	if bemptysontable {
		obj["service-providers"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("service_providers"); ok || d.HasChange("service_providers") {
			t, err := expandSystemSamlServiceProviders(d, v, "service_providers")
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["service-providers"] = t
			}
		}
	}

	if v, ok := d.GetOk("sls_url"); ok || d.HasChange("sls_url") {
		t, err := expandSystemSamlSlsUrl(d, v, "sls_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sls-url"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSamlStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("user_auto_create"); ok || d.HasChange("user_auto_create") {
		t, err := expandSystemSamlUserAutoCreate(d, v, "user_auto_create")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-auto-create"] = t
		}
	}

	if v, ok := d.GetOk("want_assertions_signed"); ok || d.HasChange("want_assertions_signed") {
		t, err := expandSystemSamlWantAssertionsSigned(d, v, "want_assertions_signed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["want-assertions-signed"] = t
		}
	}

	return &obj, nil
}
