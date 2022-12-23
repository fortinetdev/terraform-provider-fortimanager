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
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSaml(d, false)
	if err != nil {
		return fmt.Errorf("Error updating SystemSaml resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSaml(obj, mkey, paradict)
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
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSaml(d, true)

	if err != nil {
		return fmt.Errorf("Error updating SystemSaml resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSaml(obj, mkey, paradict)
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

func flattenSystemSamlAcsUrlSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlCertSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlDefaultProfileSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlEntityIdSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpSSa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemSamlFabricIdpDevIdSSa(i["dev-id"], d, pre_append)
			tmp["dev_id"] = fortiAPISubPartPatch(v, "SystemSaml-FabricIdp-DevId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_cert"
		if _, ok := i["idp-cert"]; ok {
			v := flattenSystemSamlFabricIdpIdpCertSSa(i["idp-cert"], d, pre_append)
			tmp["idp_cert"] = fortiAPISubPartPatch(v, "SystemSaml-FabricIdp-IdpCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_entity_id"
		if _, ok := i["idp-entity-id"]; ok {
			v := flattenSystemSamlFabricIdpIdpEntityIdSSa(i["idp-entity-id"], d, pre_append)
			tmp["idp_entity_id"] = fortiAPISubPartPatch(v, "SystemSaml-FabricIdp-IdpEntityId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := i["idp-single-logout-url"]; ok {
			v := flattenSystemSamlFabricIdpIdpSingleLogoutUrlSSa(i["idp-single-logout-url"], d, pre_append)
			tmp["idp_single_logout_url"] = fortiAPISubPartPatch(v, "SystemSaml-FabricIdp-IdpSingleLogoutUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := i["idp-single-sign-on-url"]; ok {
			v := flattenSystemSamlFabricIdpIdpSingleSignOnUrlSSa(i["idp-single-sign-on-url"], d, pre_append)
			tmp["idp_single_sign_on_url"] = fortiAPISubPartPatch(v, "SystemSaml-FabricIdp-IdpSingleSignOnUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_status"
		if _, ok := i["idp-status"]; ok {
			v := flattenSystemSamlFabricIdpIdpStatusSSa(i["idp-status"], d, pre_append)
			tmp["idp_status"] = fortiAPISubPartPatch(v, "SystemSaml-FabricIdp-IdpStatus")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSamlFabricIdpDevIdSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpIdpCertSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpIdpEntityIdSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpIdpSingleLogoutUrlSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpIdpSingleSignOnUrlSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpIdpStatusSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlForticloudSsoSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlIdpCertSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlIdpEntityIdSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlIdpSingleLogoutUrlSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlIdpSingleSignOnUrlSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlLoginAutoRedirectSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlRoleSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServerAddressSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSSa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemSamlServiceProvidersIdpEntityIdSSa(i["idp-entity-id"], d, pre_append)
			tmp["idp_entity_id"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-IdpEntityId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := i["idp-single-logout-url"]; ok {
			v := flattenSystemSamlServiceProvidersIdpSingleLogoutUrlSSa(i["idp-single-logout-url"], d, pre_append)
			tmp["idp_single_logout_url"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-IdpSingleLogoutUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := i["idp-single-sign-on-url"]; ok {
			v := flattenSystemSamlServiceProvidersIdpSingleSignOnUrlSSa(i["idp-single-sign-on-url"], d, pre_append)
			tmp["idp_single_sign_on_url"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-IdpSingleSignOnUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemSamlServiceProvidersNameSSa(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenSystemSamlServiceProvidersPrefixSSa(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_adom"
		if _, ok := i["sp-adom"]; ok {
			v := flattenSystemSamlServiceProvidersSpAdomSSa(i["sp-adom"], d, pre_append)
			tmp["sp_adom"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpAdom")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_cert"
		if _, ok := i["sp-cert"]; ok {
			v := flattenSystemSamlServiceProvidersSpCertSSa(i["sp-cert"], d, pre_append)
			tmp["sp_cert"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_entity_id"
		if _, ok := i["sp-entity-id"]; ok {
			v := flattenSystemSamlServiceProvidersSpEntityIdSSa(i["sp-entity-id"], d, pre_append)
			tmp["sp_entity_id"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpEntityId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_profile"
		if _, ok := i["sp-profile"]; ok {
			v := flattenSystemSamlServiceProvidersSpProfileSSa(i["sp-profile"], d, pre_append)
			tmp["sp_profile"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_logout_url"
		if _, ok := i["sp-single-logout-url"]; ok {
			v := flattenSystemSamlServiceProvidersSpSingleLogoutUrlSSa(i["sp-single-logout-url"], d, pre_append)
			tmp["sp_single_logout_url"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpSingleLogoutUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_sign_on_url"
		if _, ok := i["sp-single-sign-on-url"]; ok {
			v := flattenSystemSamlServiceProvidersSpSingleSignOnUrlSSa(i["sp-single-sign-on-url"], d, pre_append)
			tmp["sp_single_sign_on_url"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpSingleSignOnUrl")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSamlServiceProvidersIdpEntityIdSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersIdpSingleLogoutUrlSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersIdpSingleSignOnUrlSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersNameSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersPrefixSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpAdomSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpCertSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpEntityIdSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpProfileSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpSingleLogoutUrlSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpSingleSignOnUrlSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlSlsUrlSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlStatusSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlUserAutoCreateSSa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSaml(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("acs_url", flattenSystemSamlAcsUrlSSa(o["acs-url"], d, "acs_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["acs-url"], "SystemSaml-AcsUrl"); ok {
			if err = d.Set("acs_url", vv); err != nil {
				return fmt.Errorf("Error reading acs_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acs_url: %v", err)
		}
	}

	if err = d.Set("cert", flattenSystemSamlCertSSa(o["cert"], d, "cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert"], "SystemSaml-Cert"); ok {
			if err = d.Set("cert", vv); err != nil {
				return fmt.Errorf("Error reading cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("default_profile", flattenSystemSamlDefaultProfileSSa(o["default-profile"], d, "default_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-profile"], "SystemSaml-DefaultProfile"); ok {
			if err = d.Set("default_profile", vv); err != nil {
				return fmt.Errorf("Error reading default_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_profile: %v", err)
		}
	}

	if err = d.Set("entity_id", flattenSystemSamlEntityIdSSa(o["entity-id"], d, "entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["entity-id"], "SystemSaml-EntityId"); ok {
			if err = d.Set("entity_id", vv); err != nil {
				return fmt.Errorf("Error reading entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading entity_id: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fabric_idp", flattenSystemSamlFabricIdpSSa(o["fabric-idp"], d, "fabric_idp")); err != nil {
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
			if err = d.Set("fabric_idp", flattenSystemSamlFabricIdpSSa(o["fabric-idp"], d, "fabric_idp")); err != nil {
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

	if err = d.Set("forticloud_sso", flattenSystemSamlForticloudSsoSSa(o["forticloud-sso"], d, "forticloud_sso")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticloud-sso"], "SystemSaml-ForticloudSso"); ok {
			if err = d.Set("forticloud_sso", vv); err != nil {
				return fmt.Errorf("Error reading forticloud_sso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticloud_sso: %v", err)
		}
	}

	if err = d.Set("idp_cert", flattenSystemSamlIdpCertSSa(o["idp-cert"], d, "idp_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-cert"], "SystemSaml-IdpCert"); ok {
			if err = d.Set("idp_cert", vv); err != nil {
				return fmt.Errorf("Error reading idp_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_cert: %v", err)
		}
	}

	if err = d.Set("idp_entity_id", flattenSystemSamlIdpEntityIdSSa(o["idp-entity-id"], d, "idp_entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-entity-id"], "SystemSaml-IdpEntityId"); ok {
			if err = d.Set("idp_entity_id", vv); err != nil {
				return fmt.Errorf("Error reading idp_entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_entity_id: %v", err)
		}
	}

	if err = d.Set("idp_single_logout_url", flattenSystemSamlIdpSingleLogoutUrlSSa(o["idp-single-logout-url"], d, "idp_single_logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-logout-url"], "SystemSaml-IdpSingleLogoutUrl"); ok {
			if err = d.Set("idp_single_logout_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
		}
	}

	if err = d.Set("idp_single_sign_on_url", flattenSystemSamlIdpSingleSignOnUrlSSa(o["idp-single-sign-on-url"], d, "idp_single_sign_on_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-sign-on-url"], "SystemSaml-IdpSingleSignOnUrl"); ok {
			if err = d.Set("idp_single_sign_on_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("login_auto_redirect", flattenSystemSamlLoginAutoRedirectSSa(o["login-auto-redirect"], d, "login_auto_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-auto-redirect"], "SystemSaml-LoginAutoRedirect"); ok {
			if err = d.Set("login_auto_redirect", vv); err != nil {
				return fmt.Errorf("Error reading login_auto_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_auto_redirect: %v", err)
		}
	}

	if err = d.Set("role", flattenSystemSamlRoleSSa(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "SystemSaml-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("server_address", flattenSystemSamlServerAddressSSa(o["server-address"], d, "server_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-address"], "SystemSaml-ServerAddress"); ok {
			if err = d.Set("server_address", vv); err != nil {
				return fmt.Errorf("Error reading server_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_address: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service_providers", flattenSystemSamlServiceProvidersSSa(o["service-providers"], d, "service_providers")); err != nil {
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
			if err = d.Set("service_providers", flattenSystemSamlServiceProvidersSSa(o["service-providers"], d, "service_providers")); err != nil {
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

	if err = d.Set("sls_url", flattenSystemSamlSlsUrlSSa(o["sls-url"], d, "sls_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["sls-url"], "SystemSaml-SlsUrl"); ok {
			if err = d.Set("sls_url", vv); err != nil {
				return fmt.Errorf("Error reading sls_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sls_url: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSamlStatusSSa(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSaml-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("user_auto_create", flattenSystemSamlUserAutoCreateSSa(o["user-auto-create"], d, "user_auto_create")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-auto-create"], "SystemSaml-UserAutoCreate"); ok {
			if err = d.Set("user_auto_create", vv); err != nil {
				return fmt.Errorf("Error reading user_auto_create: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_auto_create: %v", err)
		}
	}

	return nil
}

func flattenSystemSamlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSamlAcsUrlSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlCertSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlDefaultProfileSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlEntityIdSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["dev-id"], _ = expandSystemSamlFabricIdpDevIdSSa(d, i["dev_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-cert"], _ = expandSystemSamlFabricIdpIdpCertSSa(d, i["idp_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_entity_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-entity-id"], _ = expandSystemSamlFabricIdpIdpEntityIdSSa(d, i["idp_entity_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-single-logout-url"], _ = expandSystemSamlFabricIdpIdpSingleLogoutUrlSSa(d, i["idp_single_logout_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-single-sign-on-url"], _ = expandSystemSamlFabricIdpIdpSingleSignOnUrlSSa(d, i["idp_single_sign_on_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-status"], _ = expandSystemSamlFabricIdpIdpStatusSSa(d, i["idp_status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSamlFabricIdpDevIdSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpIdpCertSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpIdpEntityIdSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpIdpSingleLogoutUrlSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpIdpSingleSignOnUrlSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpIdpStatusSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlForticloudSsoSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpCertSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpEntityIdSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpSingleLogoutUrlSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpSingleSignOnUrlSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlLoginAutoRedirectSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlRoleSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServerAddressSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["idp-entity-id"], _ = expandSystemSamlServiceProvidersIdpEntityIdSSa(d, i["idp_entity_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-single-logout-url"], _ = expandSystemSamlServiceProvidersIdpSingleLogoutUrlSSa(d, i["idp_single_logout_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-single-sign-on-url"], _ = expandSystemSamlServiceProvidersIdpSingleSignOnUrlSSa(d, i["idp_single_sign_on_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemSamlServiceProvidersNameSSa(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandSystemSamlServiceProvidersPrefixSSa(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_adom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-adom"], _ = expandSystemSamlServiceProvidersSpAdomSSa(d, i["sp_adom"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-cert"], _ = expandSystemSamlServiceProvidersSpCertSSa(d, i["sp_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_entity_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-entity-id"], _ = expandSystemSamlServiceProvidersSpEntityIdSSa(d, i["sp_entity_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-profile"], _ = expandSystemSamlServiceProvidersSpProfileSSa(d, i["sp_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_logout_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-single-logout-url"], _ = expandSystemSamlServiceProvidersSpSingleLogoutUrlSSa(d, i["sp_single_logout_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_sign_on_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-single-sign-on-url"], _ = expandSystemSamlServiceProvidersSpSingleSignOnUrlSSa(d, i["sp_single_sign_on_url"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSamlServiceProvidersIdpEntityIdSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersIdpSingleLogoutUrlSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersIdpSingleSignOnUrlSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersNameSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersPrefixSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpAdomSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpCertSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpEntityIdSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpProfileSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpSingleLogoutUrlSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpSingleSignOnUrlSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlSlsUrlSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlStatusSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlUserAutoCreateSSa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSaml(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("acs_url"); ok || d.HasChange("acs_url") {
		t, err := expandSystemSamlAcsUrlSSa(d, v, "acs_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acs-url"] = t
		}
	}

	if v, ok := d.GetOk("cert"); ok || d.HasChange("cert") {
		t, err := expandSystemSamlCertSSa(d, v, "cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("default_profile"); ok || d.HasChange("default_profile") {
		t, err := expandSystemSamlDefaultProfileSSa(d, v, "default_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-profile"] = t
		}
	}

	if v, ok := d.GetOk("entity_id"); ok || d.HasChange("entity_id") {
		t, err := expandSystemSamlEntityIdSSa(d, v, "entity_id")
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
			t, err := expandSystemSamlFabricIdpSSa(d, v, "fabric_idp")
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fabric-idp"] = t
			}
		}
	}

	if v, ok := d.GetOk("forticloud_sso"); ok || d.HasChange("forticloud_sso") {
		t, err := expandSystemSamlForticloudSsoSSa(d, v, "forticloud_sso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticloud-sso"] = t
		}
	}

	if v, ok := d.GetOk("idp_cert"); ok || d.HasChange("idp_cert") {
		t, err := expandSystemSamlIdpCertSSa(d, v, "idp_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-cert"] = t
		}
	}

	if v, ok := d.GetOk("idp_entity_id"); ok || d.HasChange("idp_entity_id") {
		t, err := expandSystemSamlIdpEntityIdSSa(d, v, "idp_entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-entity-id"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_logout_url"); ok || d.HasChange("idp_single_logout_url") {
		t, err := expandSystemSamlIdpSingleLogoutUrlSSa(d, v, "idp_single_logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_sign_on_url"); ok || d.HasChange("idp_single_sign_on_url") {
		t, err := expandSystemSamlIdpSingleSignOnUrlSSa(d, v, "idp_single_sign_on_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-sign-on-url"] = t
		}
	}

	if v, ok := d.GetOk("login_auto_redirect"); ok || d.HasChange("login_auto_redirect") {
		t, err := expandSystemSamlLoginAutoRedirectSSa(d, v, "login_auto_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-auto-redirect"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandSystemSamlRoleSSa(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("server_address"); ok || d.HasChange("server_address") {
		t, err := expandSystemSamlServerAddressSSa(d, v, "server_address")
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
			t, err := expandSystemSamlServiceProvidersSSa(d, v, "service_providers")
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["service-providers"] = t
			}
		}
	}

	if v, ok := d.GetOk("sls_url"); ok || d.HasChange("sls_url") {
		t, err := expandSystemSamlSlsUrlSSa(d, v, "sls_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sls-url"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSamlStatusSSa(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("user_auto_create"); ok || d.HasChange("user_auto_create") {
		t, err := expandSystemSamlUserAutoCreateSSa(d, v, "user_auto_create")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-auto-create"] = t
		}
	}

	return &obj, nil
}
