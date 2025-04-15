// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure external resource.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemExternalResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemExternalResourceCreate,
		Read:   resourceObjectSystemExternalResourceRead,
		Update: resourceObjectSystemExternalResourceUpdate,
		Delete: resourceObjectSystemExternalResourceDelete,

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
			"address_comment_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address_data_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address_name_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeInt,
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
			"comments": &schema.Schema{
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
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"namespace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"object_array_path": &schema.Schema{
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
			"refresh_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"resource": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_identity_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceObjectSystemExternalResourceCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemExternalResource(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemExternalResource resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSystemExternalResource(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemExternalResource resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemExternalResourceRead(d, m)
}

func resourceObjectSystemExternalResourceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemExternalResource(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemExternalResource resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemExternalResource(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemExternalResource resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemExternalResourceRead(d, m)
}

func resourceObjectSystemExternalResourceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemExternalResource(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemExternalResource resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemExternalResourceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemExternalResource(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemExternalResource resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemExternalResource(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemExternalResource resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemExternalResourceAddressCommentField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceAddressDataField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceAddressNameField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemExternalResourceClientCertAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemExternalResourceDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectSystemExternalResource-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenObjectSystemExternalResourceDynamicMappingSourceIp(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "ObjectSystemExternalResource-DynamicMapping-SourceIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemExternalResourceDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemExternalResourceDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemExternalResourceDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectSystemExternalResourceDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectSystemExternalResourceDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemExternalResourceDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceDynamicMappingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectSystemExternalResourceInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceNamespace(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceObjectArrayPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceRefreshRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceResource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceServerIdentityCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceUpdateMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceUserAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemExternalResourceVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemExternalResource(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("address_comment_field", flattenObjectSystemExternalResourceAddressCommentField(o["address-comment-field"], d, "address_comment_field")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-comment-field"], "ObjectSystemExternalResource-AddressCommentField"); ok {
			if err = d.Set("address_comment_field", vv); err != nil {
				return fmt.Errorf("Error reading address_comment_field: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_comment_field: %v", err)
		}
	}

	if err = d.Set("address_data_field", flattenObjectSystemExternalResourceAddressDataField(o["address-data-field"], d, "address_data_field")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-data-field"], "ObjectSystemExternalResource-AddressDataField"); ok {
			if err = d.Set("address_data_field", vv); err != nil {
				return fmt.Errorf("Error reading address_data_field: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_data_field: %v", err)
		}
	}

	if err = d.Set("address_name_field", flattenObjectSystemExternalResourceAddressNameField(o["address-name-field"], d, "address_name_field")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-name-field"], "ObjectSystemExternalResource-AddressNameField"); ok {
			if err = d.Set("address_name_field", vv); err != nil {
				return fmt.Errorf("Error reading address_name_field: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_name_field: %v", err)
		}
	}

	if err = d.Set("category", flattenObjectSystemExternalResourceCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "ObjectSystemExternalResource-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenObjectSystemExternalResourceClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "ObjectSystemExternalResource-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("client_cert_auth", flattenObjectSystemExternalResourceClientCertAuth(o["client-cert-auth"], d, "client_cert_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert-auth"], "ObjectSystemExternalResource-ClientCertAuth"); ok {
			if err = d.Set("client_cert_auth", vv); err != nil {
				return fmt.Errorf("Error reading client_cert_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert_auth: %v", err)
		}
	}

	if err = d.Set("comments", flattenObjectSystemExternalResourceComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "ObjectSystemExternalResource-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectSystemExternalResourceDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectSystemExternalResource-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectSystemExternalResourceDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectSystemExternalResource-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("interface", flattenObjectSystemExternalResourceInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectSystemExternalResource-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenObjectSystemExternalResourceInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "ObjectSystemExternalResource-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSystemExternalResourceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemExternalResource-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("namespace", flattenObjectSystemExternalResourceNamespace(o["namespace"], d, "namespace")); err != nil {
		if vv, ok := fortiAPIPatch(o["namespace"], "ObjectSystemExternalResource-Namespace"); ok {
			if err = d.Set("namespace", vv); err != nil {
				return fmt.Errorf("Error reading namespace: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading namespace: %v", err)
		}
	}

	if err = d.Set("object_array_path", flattenObjectSystemExternalResourceObjectArrayPath(o["object-array-path"], d, "object_array_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["object-array-path"], "ObjectSystemExternalResource-ObjectArrayPath"); ok {
			if err = d.Set("object_array_path", vv); err != nil {
				return fmt.Errorf("Error reading object_array_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading object_array_path: %v", err)
		}
	}

	if err = d.Set("refresh_rate", flattenObjectSystemExternalResourceRefreshRate(o["refresh-rate"], d, "refresh_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["refresh-rate"], "ObjectSystemExternalResource-RefreshRate"); ok {
			if err = d.Set("refresh_rate", vv); err != nil {
				return fmt.Errorf("Error reading refresh_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading refresh_rate: %v", err)
		}
	}

	if err = d.Set("resource", flattenObjectSystemExternalResourceResource(o["resource"], d, "resource")); err != nil {
		if vv, ok := fortiAPIPatch(o["resource"], "ObjectSystemExternalResource-Resource"); ok {
			if err = d.Set("resource", vv); err != nil {
				return fmt.Errorf("Error reading resource: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resource: %v", err)
		}
	}

	if err = d.Set("server_identity_check", flattenObjectSystemExternalResourceServerIdentityCheck(o["server-identity-check"], d, "server_identity_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-identity-check"], "ObjectSystemExternalResource-ServerIdentityCheck"); ok {
			if err = d.Set("server_identity_check", vv); err != nil {
				return fmt.Errorf("Error reading server_identity_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_identity_check: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenObjectSystemExternalResourceSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "ObjectSystemExternalResource-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectSystemExternalResourceStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectSystemExternalResource-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectSystemExternalResourceType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectSystemExternalResource-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("update_method", flattenObjectSystemExternalResourceUpdateMethod(o["update-method"], d, "update_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-method"], "ObjectSystemExternalResource-UpdateMethod"); ok {
			if err = d.Set("update_method", vv); err != nil {
				return fmt.Errorf("Error reading update_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_method: %v", err)
		}
	}

	if err = d.Set("user_agent", flattenObjectSystemExternalResourceUserAgent(o["user-agent"], d, "user_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-agent"], "ObjectSystemExternalResource-UserAgent"); ok {
			if err = d.Set("user_agent", vv); err != nil {
				return fmt.Errorf("Error reading user_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_agent: %v", err)
		}
	}

	if err = d.Set("username", flattenObjectSystemExternalResourceUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "ObjectSystemExternalResource-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectSystemExternalResourceUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectSystemExternalResource-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenObjectSystemExternalResourceVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "ObjectSystemExternalResource-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemExternalResourceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemExternalResourceAddressCommentField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceAddressDataField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceAddressNameField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemExternalResourceClientCertAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			t, err := expandObjectSystemExternalResourceDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip"], _ = expandObjectSystemExternalResourceDynamicMappingSourceIp(d, i["source_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemExternalResourceDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectSystemExternalResourceDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectSystemExternalResourceDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemExternalResourceDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceDynamicMappingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectSystemExternalResourceInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceNamespace(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceObjectArrayPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourcePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemExternalResourceRefreshRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceResource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceServerIdentityCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceUpdateMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceUserAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemExternalResourceVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemExternalResource(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address_comment_field"); ok || d.HasChange("address_comment_field") {
		t, err := expandObjectSystemExternalResourceAddressCommentField(d, v, "address_comment_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-comment-field"] = t
		}
	}

	if v, ok := d.GetOk("address_data_field"); ok || d.HasChange("address_data_field") {
		t, err := expandObjectSystemExternalResourceAddressDataField(d, v, "address_data_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-data-field"] = t
		}
	}

	if v, ok := d.GetOk("address_name_field"); ok || d.HasChange("address_name_field") {
		t, err := expandObjectSystemExternalResourceAddressNameField(d, v, "address_name_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-name-field"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandObjectSystemExternalResourceCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandObjectSystemExternalResourceClientCert(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("client_cert_auth"); ok || d.HasChange("client_cert_auth") {
		t, err := expandObjectSystemExternalResourceClientCertAuth(d, v, "client_cert_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert-auth"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandObjectSystemExternalResourceComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectSystemExternalResourceDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectSystemExternalResourceInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandObjectSystemExternalResourceInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSystemExternalResourceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("namespace"); ok || d.HasChange("namespace") {
		t, err := expandObjectSystemExternalResourceNamespace(d, v, "namespace")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["namespace"] = t
		}
	}

	if v, ok := d.GetOk("object_array_path"); ok || d.HasChange("object_array_path") {
		t, err := expandObjectSystemExternalResourceObjectArrayPath(d, v, "object_array_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object-array-path"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectSystemExternalResourcePassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("refresh_rate"); ok || d.HasChange("refresh_rate") {
		t, err := expandObjectSystemExternalResourceRefreshRate(d, v, "refresh_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["refresh-rate"] = t
		}
	}

	if v, ok := d.GetOk("resource"); ok || d.HasChange("resource") {
		t, err := expandObjectSystemExternalResourceResource(d, v, "resource")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource"] = t
		}
	}

	if v, ok := d.GetOk("server_identity_check"); ok || d.HasChange("server_identity_check") {
		t, err := expandObjectSystemExternalResourceServerIdentityCheck(d, v, "server_identity_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-identity-check"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandObjectSystemExternalResourceSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectSystemExternalResourceStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectSystemExternalResourceType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("update_method"); ok || d.HasChange("update_method") {
		t, err := expandObjectSystemExternalResourceUpdateMethod(d, v, "update_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-method"] = t
		}
	}

	if v, ok := d.GetOk("user_agent"); ok || d.HasChange("user_agent") {
		t, err := expandObjectSystemExternalResourceUserAgent(d, v, "user_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-agent"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandObjectSystemExternalResourceUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandObjectSystemExternalResourceUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandObjectSystemExternalResourceVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
