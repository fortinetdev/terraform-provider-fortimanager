// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Add this device to a Security Fabric or set up a new Security Fabric on this device.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCsf() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCsfUpdate,
		Read:   resourceSystemCsfRead,
		Update: resourceSystemCsfUpdate,
		Delete: resourceSystemCsfDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"accept_auth_by_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorization_request_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"configuration_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"downstream_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"downstream_accprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_connector": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accprofile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"configuration_write_access": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"serial": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"fabric_object_unification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_workers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"file_mgmt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"file_quota_warning": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fixed_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"forticloud_account_enforcement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"log_unification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"saml_configuration_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusted_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authorization_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"downstream_authorization": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ha_members": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"serial": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"upstream": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upstream_confirm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upstream_port": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSystemCsfUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemCsf(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsf resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemCsf(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsf resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemCsf")

	return resourceSystemCsfRead(d, m)
}

func resourceSystemCsfDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemCsf(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCsf resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCsfRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemCsf(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsf resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCsf(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsf resource from API: %v", err)
	}
	return nil
}

func flattenSystemCsfAcceptAuthByCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfAuthorizationRequestType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfConfigurationSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfDownstreamAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfDownstreamAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricConnector(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accprofile"
		if _, ok := i["accprofile"]; ok {
			v := flattenSystemCsfFabricConnectorAccprofile(i["accprofile"], d, pre_append)
			tmp["accprofile"] = fortiAPISubPartPatch(v, "SystemCsf-FabricConnector-Accprofile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "configuration_write_access"
		if _, ok := i["configuration-write-access"]; ok {
			v := flattenSystemCsfFabricConnectorConfigurationWriteAccess(i["configuration-write-access"], d, pre_append)
			tmp["configuration_write_access"] = fortiAPISubPartPatch(v, "SystemCsf-FabricConnector-ConfigurationWriteAccess")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := i["serial"]; ok {
			v := flattenSystemCsfFabricConnectorSerial(i["serial"], d, pre_append)
			tmp["serial"] = fortiAPISubPartPatch(v, "SystemCsf-FabricConnector-Serial")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemCsfFabricConnectorAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricConnectorConfigurationWriteAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricConnectorSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricObjectUnification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricWorkers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFileMgmt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFileQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFileQuotaWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfForticloudAccountEnforcement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfLogUnification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfSslProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfSamlConfigurationSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenSystemCsfTrustedListAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authorization_type"
		if _, ok := i["authorization-type"]; ok {
			v := flattenSystemCsfTrustedListAuthorizationType(i["authorization-type"], d, pre_append)
			tmp["authorization_type"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-AuthorizationType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "certificate"
		if _, ok := i["certificate"]; ok {
			v := flattenSystemCsfTrustedListCertificate(i["certificate"], d, pre_append)
			tmp["certificate"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-Certificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "downstream_authorization"
		if _, ok := i["downstream-authorization"]; ok {
			v := flattenSystemCsfTrustedListDownstreamAuthorization(i["downstream-authorization"], d, pre_append)
			tmp["downstream_authorization"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-DownstreamAuthorization")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_members"
		if _, ok := i["ha-members"]; ok {
			v := flattenSystemCsfTrustedListHaMembers(i["ha-members"], d, pre_append)
			tmp["ha_members"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-HaMembers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenSystemCsfTrustedListIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemCsfTrustedListName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := i["serial"]; ok {
			v := flattenSystemCsfTrustedListSerial(i["serial"], d, pre_append)
			tmp["serial"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-Serial")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemCsfTrustedListAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListAuthorizationType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListDownstreamAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListHaMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfUpstream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfUpstreamConfirm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfUpstreamPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemCsf(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("accept_auth_by_cert", flattenSystemCsfAcceptAuthByCert(o["accept-auth-by-cert"], d, "accept_auth_by_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["accept-auth-by-cert"], "SystemCsf-AcceptAuthByCert"); ok {
			if err = d.Set("accept_auth_by_cert", vv); err != nil {
				return fmt.Errorf("Error reading accept_auth_by_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accept_auth_by_cert: %v", err)
		}
	}

	if err = d.Set("authorization_request_type", flattenSystemCsfAuthorizationRequestType(o["authorization-request-type"], d, "authorization_request_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["authorization-request-type"], "SystemCsf-AuthorizationRequestType"); ok {
			if err = d.Set("authorization_request_type", vv); err != nil {
				return fmt.Errorf("Error reading authorization_request_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authorization_request_type: %v", err)
		}
	}

	if err = d.Set("certificate", flattenSystemCsfCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "SystemCsf-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("configuration_sync", flattenSystemCsfConfigurationSync(o["configuration-sync"], d, "configuration_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["configuration-sync"], "SystemCsf-ConfigurationSync"); ok {
			if err = d.Set("configuration_sync", vv); err != nil {
				return fmt.Errorf("Error reading configuration_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading configuration_sync: %v", err)
		}
	}

	if err = d.Set("downstream_access", flattenSystemCsfDownstreamAccess(o["downstream-access"], d, "downstream_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["downstream-access"], "SystemCsf-DownstreamAccess"); ok {
			if err = d.Set("downstream_access", vv); err != nil {
				return fmt.Errorf("Error reading downstream_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading downstream_access: %v", err)
		}
	}

	if err = d.Set("downstream_accprofile", flattenSystemCsfDownstreamAccprofile(o["downstream-accprofile"], d, "downstream_accprofile")); err != nil {
		if vv, ok := fortiAPIPatch(o["downstream-accprofile"], "SystemCsf-DownstreamAccprofile"); ok {
			if err = d.Set("downstream_accprofile", vv); err != nil {
				return fmt.Errorf("Error reading downstream_accprofile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading downstream_accprofile: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fabric_connector", flattenSystemCsfFabricConnector(o["fabric-connector"], d, "fabric_connector")); err != nil {
			if vv, ok := fortiAPIPatch(o["fabric-connector"], "SystemCsf-FabricConnector"); ok {
				if err = d.Set("fabric_connector", vv); err != nil {
					return fmt.Errorf("Error reading fabric_connector: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading fabric_connector: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fabric_connector"); ok {
			if err = d.Set("fabric_connector", flattenSystemCsfFabricConnector(o["fabric-connector"], d, "fabric_connector")); err != nil {
				if vv, ok := fortiAPIPatch(o["fabric-connector"], "SystemCsf-FabricConnector"); ok {
					if err = d.Set("fabric_connector", vv); err != nil {
						return fmt.Errorf("Error reading fabric_connector: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading fabric_connector: %v", err)
				}
			}
		}
	}

	if err = d.Set("fabric_object_unification", flattenSystemCsfFabricObjectUnification(o["fabric-object-unification"], d, "fabric_object_unification")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object-unification"], "SystemCsf-FabricObjectUnification"); ok {
			if err = d.Set("fabric_object_unification", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object_unification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object_unification: %v", err)
		}
	}

	if err = d.Set("fabric_workers", flattenSystemCsfFabricWorkers(o["fabric-workers"], d, "fabric_workers")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-workers"], "SystemCsf-FabricWorkers"); ok {
			if err = d.Set("fabric_workers", vv); err != nil {
				return fmt.Errorf("Error reading fabric_workers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_workers: %v", err)
		}
	}

	if err = d.Set("file_mgmt", flattenSystemCsfFileMgmt(o["file-mgmt"], d, "file_mgmt")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-mgmt"], "SystemCsf-FileMgmt"); ok {
			if err = d.Set("file_mgmt", vv); err != nil {
				return fmt.Errorf("Error reading file_mgmt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_mgmt: %v", err)
		}
	}

	if err = d.Set("file_quota", flattenSystemCsfFileQuota(o["file-quota"], d, "file_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-quota"], "SystemCsf-FileQuota"); ok {
			if err = d.Set("file_quota", vv); err != nil {
				return fmt.Errorf("Error reading file_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_quota: %v", err)
		}
	}

	if err = d.Set("file_quota_warning", flattenSystemCsfFileQuotaWarning(o["file-quota-warning"], d, "file_quota_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-quota-warning"], "SystemCsf-FileQuotaWarning"); ok {
			if err = d.Set("file_quota_warning", vv); err != nil {
				return fmt.Errorf("Error reading file_quota_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_quota_warning: %v", err)
		}
	}

	if err = d.Set("forticloud_account_enforcement", flattenSystemCsfForticloudAccountEnforcement(o["forticloud-account-enforcement"], d, "forticloud_account_enforcement")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticloud-account-enforcement"], "SystemCsf-ForticloudAccountEnforcement"); ok {
			if err = d.Set("forticloud_account_enforcement", vv); err != nil {
				return fmt.Errorf("Error reading forticloud_account_enforcement: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticloud_account_enforcement: %v", err)
		}
	}

	if err = d.Set("group_name", flattenSystemCsfGroupName(o["group-name"], d, "group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-name"], "SystemCsf-GroupName"); ok {
			if err = d.Set("group_name", vv); err != nil {
				return fmt.Errorf("Error reading group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("log_unification", flattenSystemCsfLogUnification(o["log-unification"], d, "log_unification")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-unification"], "SystemCsf-LogUnification"); ok {
			if err = d.Set("log_unification", vv); err != nil {
				return fmt.Errorf("Error reading log_unification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_unification: %v", err)
		}
	}

	if err = d.Set("ssl_protocol", flattenSystemCsfSslProtocol(o["ssl-protocol"], d, "ssl_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-protocol"], "SystemCsf-SslProtocol"); ok {
			if err = d.Set("ssl_protocol", vv); err != nil {
				return fmt.Errorf("Error reading ssl_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_protocol: %v", err)
		}
	}

	if err = d.Set("saml_configuration_sync", flattenSystemCsfSamlConfigurationSync(o["saml-configuration-sync"], d, "saml_configuration_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["saml-configuration-sync"], "SystemCsf-SamlConfigurationSync"); ok {
			if err = d.Set("saml_configuration_sync", vv); err != nil {
				return fmt.Errorf("Error reading saml_configuration_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saml_configuration_sync: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemCsfStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemCsf-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("trusted_list", flattenSystemCsfTrustedList(o["trusted-list"], d, "trusted_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["trusted-list"], "SystemCsf-TrustedList"); ok {
				if err = d.Set("trusted_list", vv); err != nil {
					return fmt.Errorf("Error reading trusted_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading trusted_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("trusted_list"); ok {
			if err = d.Set("trusted_list", flattenSystemCsfTrustedList(o["trusted-list"], d, "trusted_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["trusted-list"], "SystemCsf-TrustedList"); ok {
					if err = d.Set("trusted_list", vv); err != nil {
						return fmt.Errorf("Error reading trusted_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading trusted_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("upstream", flattenSystemCsfUpstream(o["upstream"], d, "upstream")); err != nil {
		if vv, ok := fortiAPIPatch(o["upstream"], "SystemCsf-Upstream"); ok {
			if err = d.Set("upstream", vv); err != nil {
				return fmt.Errorf("Error reading upstream: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upstream: %v", err)
		}
	}

	if err = d.Set("upstream_confirm", flattenSystemCsfUpstreamConfirm(o["upstream-confirm"], d, "upstream_confirm")); err != nil {
		if vv, ok := fortiAPIPatch(o["upstream-confirm"], "SystemCsf-UpstreamConfirm"); ok {
			if err = d.Set("upstream_confirm", vv); err != nil {
				return fmt.Errorf("Error reading upstream_confirm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upstream_confirm: %v", err)
		}
	}

	if err = d.Set("upstream_port", flattenSystemCsfUpstreamPort(o["upstream-port"], d, "upstream_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["upstream-port"], "SystemCsf-UpstreamPort"); ok {
			if err = d.Set("upstream_port", vv); err != nil {
				return fmt.Errorf("Error reading upstream_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upstream_port: %v", err)
		}
	}

	return nil
}

func flattenSystemCsfFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCsfAcceptAuthByCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfAuthorizationRequestType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfConfigurationSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfDownstreamAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfDownstreamAccprofile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricConnector(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accprofile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["accprofile"], _ = expandSystemCsfFabricConnectorAccprofile(d, i["accprofile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "configuration_write_access"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["configuration-write-access"], _ = expandSystemCsfFabricConnectorConfigurationWriteAccess(d, i["configuration_write_access"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["serial"], _ = expandSystemCsfFabricConnectorSerial(d, i["serial"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemCsfFabricConnectorAccprofile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricConnectorConfigurationWriteAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricConnectorSerial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricObjectUnification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricWorkers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFileMgmt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFileQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFileQuotaWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFixedKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCsfForticloudAccountEnforcement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfGroupPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCsfLogUnification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfSslProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfSamlConfigurationSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandSystemCsfTrustedListAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authorization_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authorization-type"], _ = expandSystemCsfTrustedListAuthorizationType(d, i["authorization_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "certificate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["certificate"], _ = expandSystemCsfTrustedListCertificate(d, i["certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "downstream_authorization"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["downstream-authorization"], _ = expandSystemCsfTrustedListDownstreamAuthorization(d, i["downstream_authorization"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ha-members"], _ = expandSystemCsfTrustedListHaMembers(d, i["ha_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandSystemCsfTrustedListIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemCsfTrustedListName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["serial"], _ = expandSystemCsfTrustedListSerial(d, i["serial"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemCsfTrustedListAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListAuthorizationType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListDownstreamAuthorization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListHaMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListSerial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUpstream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUpstreamConfirm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUpstreamPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCsf(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accept_auth_by_cert"); ok || d.HasChange("accept_auth_by_cert") {
		t, err := expandSystemCsfAcceptAuthByCert(d, v, "accept_auth_by_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accept-auth-by-cert"] = t
		}
	}

	if v, ok := d.GetOk("authorization_request_type"); ok || d.HasChange("authorization_request_type") {
		t, err := expandSystemCsfAuthorizationRequestType(d, v, "authorization_request_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorization-request-type"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandSystemCsfCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("configuration_sync"); ok || d.HasChange("configuration_sync") {
		t, err := expandSystemCsfConfigurationSync(d, v, "configuration_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["configuration-sync"] = t
		}
	}

	if v, ok := d.GetOk("downstream_access"); ok || d.HasChange("downstream_access") {
		t, err := expandSystemCsfDownstreamAccess(d, v, "downstream_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downstream-access"] = t
		}
	}

	if v, ok := d.GetOk("downstream_accprofile"); ok || d.HasChange("downstream_accprofile") {
		t, err := expandSystemCsfDownstreamAccprofile(d, v, "downstream_accprofile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downstream-accprofile"] = t
		}
	}

	if v, ok := d.GetOk("fabric_connector"); ok || d.HasChange("fabric_connector") {
		t, err := expandSystemCsfFabricConnector(d, v, "fabric_connector")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-connector"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object_unification"); ok || d.HasChange("fabric_object_unification") {
		t, err := expandSystemCsfFabricObjectUnification(d, v, "fabric_object_unification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object-unification"] = t
		}
	}

	if v, ok := d.GetOk("fabric_workers"); ok || d.HasChange("fabric_workers") {
		t, err := expandSystemCsfFabricWorkers(d, v, "fabric_workers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-workers"] = t
		}
	}

	if v, ok := d.GetOk("file_mgmt"); ok || d.HasChange("file_mgmt") {
		t, err := expandSystemCsfFileMgmt(d, v, "file_mgmt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-mgmt"] = t
		}
	}

	if v, ok := d.GetOk("file_quota"); ok || d.HasChange("file_quota") {
		t, err := expandSystemCsfFileQuota(d, v, "file_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-quota"] = t
		}
	}

	if v, ok := d.GetOk("file_quota_warning"); ok || d.HasChange("file_quota_warning") {
		t, err := expandSystemCsfFileQuotaWarning(d, v, "file_quota_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-quota-warning"] = t
		}
	}

	if v, ok := d.GetOk("fixed_key"); ok || d.HasChange("fixed_key") {
		t, err := expandSystemCsfFixedKey(d, v, "fixed_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fixed-key"] = t
		}
	}

	if v, ok := d.GetOk("forticloud_account_enforcement"); ok || d.HasChange("forticloud_account_enforcement") {
		t, err := expandSystemCsfForticloudAccountEnforcement(d, v, "forticloud_account_enforcement")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticloud-account-enforcement"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok || d.HasChange("group_name") {
		t, err := expandSystemCsfGroupName(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("group_password"); ok || d.HasChange("group_password") {
		t, err := expandSystemCsfGroupPassword(d, v, "group_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-password"] = t
		}
	}

	if v, ok := d.GetOk("log_unification"); ok || d.HasChange("log_unification") {
		t, err := expandSystemCsfLogUnification(d, v, "log_unification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-unification"] = t
		}
	}

	if v, ok := d.GetOk("ssl_protocol"); ok || d.HasChange("ssl_protocol") {
		t, err := expandSystemCsfSslProtocol(d, v, "ssl_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-protocol"] = t
		}
	}

	if v, ok := d.GetOk("saml_configuration_sync"); ok || d.HasChange("saml_configuration_sync") {
		t, err := expandSystemCsfSamlConfigurationSync(d, v, "saml_configuration_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-configuration-sync"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemCsfStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trusted_list"); ok || d.HasChange("trusted_list") {
		t, err := expandSystemCsfTrustedList(d, v, "trusted_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted-list"] = t
		}
	}

	if v, ok := d.GetOk("upstream"); ok || d.HasChange("upstream") {
		t, err := expandSystemCsfUpstream(d, v, "upstream")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upstream"] = t
		}
	}

	if v, ok := d.GetOk("upstream_confirm"); ok || d.HasChange("upstream_confirm") {
		t, err := expandSystemCsfUpstreamConfirm(d, v, "upstream_confirm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upstream-confirm"] = t
		}
	}

	if v, ok := d.GetOk("upstream_port"); ok || d.HasChange("upstream_port") {
		t, err := expandSystemCsfUpstreamPort(d, v, "upstream_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upstream-port"] = t
		}
	}

	return &obj, nil
}
