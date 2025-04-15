// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure RADIUS server entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserRadiusDynamicMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserRadiusDynamicMappingCreate,
		Read:   resourceObjectUserRadiusDynamicMappingRead,
		Update: resourceObjectUserRadiusDynamicMappingUpdate,
		Delete: resourceObjectUserRadiusDynamicMappingDelete,

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
			"radius": &schema.Schema{
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
			"account_key_processing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"accounting_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"secret": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vrf_select": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"acct_all_servers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acct_interim_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"all_usergroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ca_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"call_station_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"class": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dp_carrier_endpoint_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dp_carrier_endpoint_block_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dp_context_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dp_flush_ip_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dp_hold_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dp_http_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dp_http_header_fallback": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dp_http_header_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dp_http_header_suppress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dp_log_dyn_flags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dp_log_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dp_mem_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dp_profile_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dp_profile_attribute_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dp_radius_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dp_radius_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dp_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"dp_validate_request_secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"endpoint_translation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ep_carrier_endpoint_convert_hex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ep_carrier_endpoint_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ep_carrier_endpoint_header_suppress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ep_carrier_endpoint_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ep_carrier_endpoint_prefix_range_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ep_carrier_endpoint_prefix_range_min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ep_carrier_endpoint_prefix_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ep_carrier_endpoint_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ep_ip_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ep_ip_header_suppress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ep_missing_header_fallback": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ep_profile_query_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_override_attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"h3c_compatibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"mac_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_password_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_username_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nas_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nas_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nas_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password_renewal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_coa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"require_message_authenticator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso_context_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rsso_endpoint_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rsso_endpoint_block_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rsso_ep_one_ip_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso_flush_ip_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rsso_log_flags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rsso_log_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rsso_radius_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rsso_radius_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rsso_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"rsso_validate_request_secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"secondary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"server": &schema.Schema{
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
			},
			"source_ip_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sso_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso_attribute_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso_attribute_value_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"switch_controller_acct_fast_framedip_detect": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"switch_controller_nas_ip_dynamic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_service_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"tertiary_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"tertiary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tls_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transport_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_group_for_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_management_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username_case_sensitive": &schema.Schema{
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

func resourceObjectUserRadiusDynamicMappingCreate(d *schema.ResourceData, m interface{}) error {
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

	radius := d.Get("radius").(string)
	paradict["radius"] = radius

	obj, err := getObjectObjectUserRadiusDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserRadiusDynamicMapping resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectUserRadiusDynamicMapping(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserRadiusDynamicMapping resource: %v", err)
	}

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectUserRadiusDynamicMappingRead(d, m)
}

func resourceObjectUserRadiusDynamicMappingUpdate(d *schema.ResourceData, m interface{}) error {
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

	radius := d.Get("radius").(string)
	paradict["radius"] = radius

	obj, err := getObjectObjectUserRadiusDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserRadiusDynamicMapping resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectUserRadiusDynamicMapping(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserRadiusDynamicMapping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectUserRadiusDynamicMappingRead(d, m)
}

func resourceObjectUserRadiusDynamicMappingDelete(d *schema.ResourceData, m interface{}) error {
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

	radius := d.Get("radius").(string)
	paradict["radius"] = radius

	wsParams["adom"] = adomv

	err = c.DeleteObjectUserRadiusDynamicMapping(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserRadiusDynamicMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserRadiusDynamicMappingRead(d *schema.ResourceData, m interface{}) error {
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

	radius := d.Get("radius").(string)
	if radius == "" {
		radius = importOptionChecking(m.(*FortiClient).Cfg, "radius")
		if radius == "" {
			return fmt.Errorf("Parameter radius is missing")
		}
		if err = d.Set("radius", radius); err != nil {
			return fmt.Errorf("Error set params radius: %v", err)
		}
	}
	if mkey, err = checkScopeId(mkey); err != nil {
		return fmt.Errorf("Error set ID : %v", err)
	}
	paradict["radius"] = radius

	o, err := c.ReadObjectUserRadiusDynamicMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserRadiusDynamicMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserRadiusDynamicMapping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserRadiusDynamicMapping resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserRadiusDynamicMappingScope2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserRadiusDynamicMappingScopeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectUserRadiusDynamicMappingScopeVdom2edl(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserRadiusDynamicMappingScopeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingScopeVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountKeyCertField2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountKeyProcessing2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountingServer2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServerId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServerInterface2edl(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServerInterfaceSelectMethod2edl(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServerPort2edl(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServerServer2edl(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServerSourceIp2edl(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-SourceIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServerStatus2edl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := i["vrf-select"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServerVrfSelect2edl(i["vrf-select"], d, pre_append)
			tmp["vrf_select"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-VrfSelect")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserRadiusDynamicMappingAccountingServerId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountingServerInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserRadiusDynamicMappingAccountingServerInterfaceSelectMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountingServerPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountingServerServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountingServerSourceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountingServerStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountingServerVrfSelect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAcctAllServers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAcctInterimInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAllUsergroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAuthType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingCaCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserRadiusDynamicMappingCallStationIdType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingClass2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusDynamicMappingClientCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserRadiusDynamicMappingDelimiter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpCarrierEndpointAttribute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpCarrierEndpointBlockAttribute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpContextTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpFlushIpSession2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpHoldTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpHttpHeader2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpHttpHeaderFallback2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpHttpHeaderStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpHttpHeaderSuppress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpLogDynFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusDynamicMappingDpLogPeriod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpMemPercent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpProfileAttribute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpProfileAttributeKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpRadiusResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpRadiusServerPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpValidateRequestSecret2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDynamicProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEndpointTranslation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointConvertHex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointHeader2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointHeaderSuppress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMax2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointSource2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpIpHeader2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpIpHeaderSuppress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpMissingHeaderFallback2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpProfileQueryType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingGroupOverrideAttrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingH3CCompatibility2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserRadiusDynamicMappingInterfaceSelectMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingMacCase2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingMacPasswordDelimiter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingMacUsernameDelimiter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingNasId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingNasIdType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingNasIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingPasswordEncoding2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingPasswordRenewal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRadiusCoa2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRadiusPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRequireMessageAuthenticator2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRsso2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoContextTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoEndpointAttribute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoEndpointBlockAttribute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoEpOneIpOnly2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoFlushIpSession2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoLogFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusDynamicMappingRssoLogPeriod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoRadiusResponse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoRadiusServerPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoValidateRequestSecret2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingSecondaryServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingServerIdentityCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingSourceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingSourceIpInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusDynamicMappingSsoAttribute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingSsoAttributeKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingSsoAttributeValueOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingStatusTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingSwitchControllerAcctFastFramedipDetect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingSwitchControllerNasIpDynamic2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingSwitchControllerServiceType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusDynamicMappingTertiaryServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingTlsMinProtoVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingTransportProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingUseGroupForProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingUseManagementVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingUsernameCaseSensitive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingVrfSelect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserRadiusDynamicMapping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("_scope", flattenObjectUserRadiusDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["_scope"], "ObjectUserRadiusDynamicMapping-Scope"); ok {
				if err = d.Set("_scope", vv); err != nil {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading _scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("_scope"); ok {
			if err = d.Set("_scope", flattenObjectUserRadiusDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["_scope"], "ObjectUserRadiusDynamicMapping-Scope"); ok {
					if err = d.Set("_scope", vv); err != nil {
						return fmt.Errorf("Error reading _scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			}
		}
	}

	if err = d.Set("account_key_cert_field", flattenObjectUserRadiusDynamicMappingAccountKeyCertField2edl(o["account-key-cert-field"], d, "account_key_cert_field")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-cert-field"], "ObjectUserRadiusDynamicMapping-AccountKeyCertField"); ok {
			if err = d.Set("account_key_cert_field", vv); err != nil {
				return fmt.Errorf("Error reading account_key_cert_field: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_cert_field: %v", err)
		}
	}

	if err = d.Set("account_key_processing", flattenObjectUserRadiusDynamicMappingAccountKeyProcessing2edl(o["account-key-processing"], d, "account_key_processing")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-processing"], "ObjectUserRadiusDynamicMapping-AccountKeyProcessing"); ok {
			if err = d.Set("account_key_processing", vv); err != nil {
				return fmt.Errorf("Error reading account_key_processing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_processing: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("accounting_server", flattenObjectUserRadiusDynamicMappingAccountingServer2edl(o["accounting-server"], d, "accounting_server")); err != nil {
			if vv, ok := fortiAPIPatch(o["accounting-server"], "ObjectUserRadiusDynamicMapping-AccountingServer"); ok {
				if err = d.Set("accounting_server", vv); err != nil {
					return fmt.Errorf("Error reading accounting_server: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading accounting_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("accounting_server"); ok {
			if err = d.Set("accounting_server", flattenObjectUserRadiusDynamicMappingAccountingServer2edl(o["accounting-server"], d, "accounting_server")); err != nil {
				if vv, ok := fortiAPIPatch(o["accounting-server"], "ObjectUserRadiusDynamicMapping-AccountingServer"); ok {
					if err = d.Set("accounting_server", vv); err != nil {
						return fmt.Errorf("Error reading accounting_server: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading accounting_server: %v", err)
				}
			}
		}
	}

	if err = d.Set("acct_all_servers", flattenObjectUserRadiusDynamicMappingAcctAllServers2edl(o["acct-all-servers"], d, "acct_all_servers")); err != nil {
		if vv, ok := fortiAPIPatch(o["acct-all-servers"], "ObjectUserRadiusDynamicMapping-AcctAllServers"); ok {
			if err = d.Set("acct_all_servers", vv); err != nil {
				return fmt.Errorf("Error reading acct_all_servers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acct_all_servers: %v", err)
		}
	}

	if err = d.Set("acct_interim_interval", flattenObjectUserRadiusDynamicMappingAcctInterimInterval2edl(o["acct-interim-interval"], d, "acct_interim_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["acct-interim-interval"], "ObjectUserRadiusDynamicMapping-AcctInterimInterval"); ok {
			if err = d.Set("acct_interim_interval", vv); err != nil {
				return fmt.Errorf("Error reading acct_interim_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acct_interim_interval: %v", err)
		}
	}

	if err = d.Set("all_usergroup", flattenObjectUserRadiusDynamicMappingAllUsergroup2edl(o["all-usergroup"], d, "all_usergroup")); err != nil {
		if vv, ok := fortiAPIPatch(o["all-usergroup"], "ObjectUserRadiusDynamicMapping-AllUsergroup"); ok {
			if err = d.Set("all_usergroup", vv); err != nil {
				return fmt.Errorf("Error reading all_usergroup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading all_usergroup: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenObjectUserRadiusDynamicMappingAuthType2edl(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "ObjectUserRadiusDynamicMapping-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("ca_cert", flattenObjectUserRadiusDynamicMappingCaCert2edl(o["ca-cert"], d, "ca_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca-cert"], "ObjectUserRadiusDynamicMapping-CaCert"); ok {
			if err = d.Set("ca_cert", vv); err != nil {
				return fmt.Errorf("Error reading ca_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca_cert: %v", err)
		}
	}

	if err = d.Set("call_station_id_type", flattenObjectUserRadiusDynamicMappingCallStationIdType2edl(o["call-station-id-type"], d, "call_station_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["call-station-id-type"], "ObjectUserRadiusDynamicMapping-CallStationIdType"); ok {
			if err = d.Set("call_station_id_type", vv); err != nil {
				return fmt.Errorf("Error reading call_station_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading call_station_id_type: %v", err)
		}
	}

	if err = d.Set("class", flattenObjectUserRadiusDynamicMappingClass2edl(o["class"], d, "class")); err != nil {
		if vv, ok := fortiAPIPatch(o["class"], "ObjectUserRadiusDynamicMapping-Class"); ok {
			if err = d.Set("class", vv); err != nil {
				return fmt.Errorf("Error reading class: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenObjectUserRadiusDynamicMappingClientCert2edl(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "ObjectUserRadiusDynamicMapping-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("delimiter", flattenObjectUserRadiusDynamicMappingDelimiter2edl(o["delimiter"], d, "delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["delimiter"], "ObjectUserRadiusDynamicMapping-Delimiter"); ok {
			if err = d.Set("delimiter", vv); err != nil {
				return fmt.Errorf("Error reading delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delimiter: %v", err)
		}
	}

	if err = d.Set("dp_carrier_endpoint_attribute", flattenObjectUserRadiusDynamicMappingDpCarrierEndpointAttribute2edl(o["dp-carrier-endpoint-attribute"], d, "dp_carrier_endpoint_attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-carrier-endpoint-attribute"], "ObjectUserRadiusDynamicMapping-DpCarrierEndpointAttribute"); ok {
			if err = d.Set("dp_carrier_endpoint_attribute", vv); err != nil {
				return fmt.Errorf("Error reading dp_carrier_endpoint_attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_carrier_endpoint_attribute: %v", err)
		}
	}

	if err = d.Set("dp_carrier_endpoint_block_attribute", flattenObjectUserRadiusDynamicMappingDpCarrierEndpointBlockAttribute2edl(o["dp-carrier-endpoint-block-attribute"], d, "dp_carrier_endpoint_block_attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-carrier-endpoint-block-attribute"], "ObjectUserRadiusDynamicMapping-DpCarrierEndpointBlockAttribute"); ok {
			if err = d.Set("dp_carrier_endpoint_block_attribute", vv); err != nil {
				return fmt.Errorf("Error reading dp_carrier_endpoint_block_attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_carrier_endpoint_block_attribute: %v", err)
		}
	}

	if err = d.Set("dp_context_timeout", flattenObjectUserRadiusDynamicMappingDpContextTimeout2edl(o["dp-context-timeout"], d, "dp_context_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-context-timeout"], "ObjectUserRadiusDynamicMapping-DpContextTimeout"); ok {
			if err = d.Set("dp_context_timeout", vv); err != nil {
				return fmt.Errorf("Error reading dp_context_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_context_timeout: %v", err)
		}
	}

	if err = d.Set("dp_flush_ip_session", flattenObjectUserRadiusDynamicMappingDpFlushIpSession2edl(o["dp-flush-ip-session"], d, "dp_flush_ip_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-flush-ip-session"], "ObjectUserRadiusDynamicMapping-DpFlushIpSession"); ok {
			if err = d.Set("dp_flush_ip_session", vv); err != nil {
				return fmt.Errorf("Error reading dp_flush_ip_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_flush_ip_session: %v", err)
		}
	}

	if err = d.Set("dp_hold_time", flattenObjectUserRadiusDynamicMappingDpHoldTime2edl(o["dp-hold-time"], d, "dp_hold_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-hold-time"], "ObjectUserRadiusDynamicMapping-DpHoldTime"); ok {
			if err = d.Set("dp_hold_time", vv); err != nil {
				return fmt.Errorf("Error reading dp_hold_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_hold_time: %v", err)
		}
	}

	if err = d.Set("dp_http_header", flattenObjectUserRadiusDynamicMappingDpHttpHeader2edl(o["dp-http-header"], d, "dp_http_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-http-header"], "ObjectUserRadiusDynamicMapping-DpHttpHeader"); ok {
			if err = d.Set("dp_http_header", vv); err != nil {
				return fmt.Errorf("Error reading dp_http_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_http_header: %v", err)
		}
	}

	if err = d.Set("dp_http_header_fallback", flattenObjectUserRadiusDynamicMappingDpHttpHeaderFallback2edl(o["dp-http-header-fallback"], d, "dp_http_header_fallback")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-http-header-fallback"], "ObjectUserRadiusDynamicMapping-DpHttpHeaderFallback"); ok {
			if err = d.Set("dp_http_header_fallback", vv); err != nil {
				return fmt.Errorf("Error reading dp_http_header_fallback: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_http_header_fallback: %v", err)
		}
	}

	if err = d.Set("dp_http_header_status", flattenObjectUserRadiusDynamicMappingDpHttpHeaderStatus2edl(o["dp-http-header-status"], d, "dp_http_header_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-http-header-status"], "ObjectUserRadiusDynamicMapping-DpHttpHeaderStatus"); ok {
			if err = d.Set("dp_http_header_status", vv); err != nil {
				return fmt.Errorf("Error reading dp_http_header_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_http_header_status: %v", err)
		}
	}

	if err = d.Set("dp_http_header_suppress", flattenObjectUserRadiusDynamicMappingDpHttpHeaderSuppress2edl(o["dp-http-header-suppress"], d, "dp_http_header_suppress")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-http-header-suppress"], "ObjectUserRadiusDynamicMapping-DpHttpHeaderSuppress"); ok {
			if err = d.Set("dp_http_header_suppress", vv); err != nil {
				return fmt.Errorf("Error reading dp_http_header_suppress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_http_header_suppress: %v", err)
		}
	}

	if err = d.Set("dp_log_dyn_flags", flattenObjectUserRadiusDynamicMappingDpLogDynFlags2edl(o["dp-log-dyn_flags"], d, "dp_log_dyn_flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-log-dyn_flags"], "ObjectUserRadiusDynamicMapping-DpLogDynFlags"); ok {
			if err = d.Set("dp_log_dyn_flags", vv); err != nil {
				return fmt.Errorf("Error reading dp_log_dyn_flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_log_dyn_flags: %v", err)
		}
	}

	if err = d.Set("dp_log_period", flattenObjectUserRadiusDynamicMappingDpLogPeriod2edl(o["dp-log-period"], d, "dp_log_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-log-period"], "ObjectUserRadiusDynamicMapping-DpLogPeriod"); ok {
			if err = d.Set("dp_log_period", vv); err != nil {
				return fmt.Errorf("Error reading dp_log_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_log_period: %v", err)
		}
	}

	if err = d.Set("dp_mem_percent", flattenObjectUserRadiusDynamicMappingDpMemPercent2edl(o["dp-mem-percent"], d, "dp_mem_percent")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-mem-percent"], "ObjectUserRadiusDynamicMapping-DpMemPercent"); ok {
			if err = d.Set("dp_mem_percent", vv); err != nil {
				return fmt.Errorf("Error reading dp_mem_percent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_mem_percent: %v", err)
		}
	}

	if err = d.Set("dp_profile_attribute", flattenObjectUserRadiusDynamicMappingDpProfileAttribute2edl(o["dp-profile-attribute"], d, "dp_profile_attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-profile-attribute"], "ObjectUserRadiusDynamicMapping-DpProfileAttribute"); ok {
			if err = d.Set("dp_profile_attribute", vv); err != nil {
				return fmt.Errorf("Error reading dp_profile_attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_profile_attribute: %v", err)
		}
	}

	if err = d.Set("dp_profile_attribute_key", flattenObjectUserRadiusDynamicMappingDpProfileAttributeKey2edl(o["dp-profile-attribute-key"], d, "dp_profile_attribute_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-profile-attribute-key"], "ObjectUserRadiusDynamicMapping-DpProfileAttributeKey"); ok {
			if err = d.Set("dp_profile_attribute_key", vv); err != nil {
				return fmt.Errorf("Error reading dp_profile_attribute_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_profile_attribute_key: %v", err)
		}
	}

	if err = d.Set("dp_radius_response", flattenObjectUserRadiusDynamicMappingDpRadiusResponse2edl(o["dp-radius-response"], d, "dp_radius_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-radius-response"], "ObjectUserRadiusDynamicMapping-DpRadiusResponse"); ok {
			if err = d.Set("dp_radius_response", vv); err != nil {
				return fmt.Errorf("Error reading dp_radius_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_radius_response: %v", err)
		}
	}

	if err = d.Set("dp_radius_server_port", flattenObjectUserRadiusDynamicMappingDpRadiusServerPort2edl(o["dp-radius-server-port"], d, "dp_radius_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-radius-server-port"], "ObjectUserRadiusDynamicMapping-DpRadiusServerPort"); ok {
			if err = d.Set("dp_radius_server_port", vv); err != nil {
				return fmt.Errorf("Error reading dp_radius_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_radius_server_port: %v", err)
		}
	}

	if err = d.Set("dp_validate_request_secret", flattenObjectUserRadiusDynamicMappingDpValidateRequestSecret2edl(o["dp-validate-request-secret"], d, "dp_validate_request_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["dp-validate-request-secret"], "ObjectUserRadiusDynamicMapping-DpValidateRequestSecret"); ok {
			if err = d.Set("dp_validate_request_secret", vv); err != nil {
				return fmt.Errorf("Error reading dp_validate_request_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dp_validate_request_secret: %v", err)
		}
	}

	if err = d.Set("dynamic_profile", flattenObjectUserRadiusDynamicMappingDynamicProfile2edl(o["dynamic-profile"], d, "dynamic_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-profile"], "ObjectUserRadiusDynamicMapping-DynamicProfile"); ok {
			if err = d.Set("dynamic_profile", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_profile: %v", err)
		}
	}

	if err = d.Set("endpoint_translation", flattenObjectUserRadiusDynamicMappingEndpointTranslation2edl(o["endpoint-translation"], d, "endpoint_translation")); err != nil {
		if vv, ok := fortiAPIPatch(o["endpoint-translation"], "ObjectUserRadiusDynamicMapping-EndpointTranslation"); ok {
			if err = d.Set("endpoint_translation", vv); err != nil {
				return fmt.Errorf("Error reading endpoint_translation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading endpoint_translation: %v", err)
		}
	}

	if err = d.Set("ep_carrier_endpoint_convert_hex", flattenObjectUserRadiusDynamicMappingEpCarrierEndpointConvertHex2edl(o["ep-carrier-endpoint-convert-hex"], d, "ep_carrier_endpoint_convert_hex")); err != nil {
		if vv, ok := fortiAPIPatch(o["ep-carrier-endpoint-convert-hex"], "ObjectUserRadiusDynamicMapping-EpCarrierEndpointConvertHex"); ok {
			if err = d.Set("ep_carrier_endpoint_convert_hex", vv); err != nil {
				return fmt.Errorf("Error reading ep_carrier_endpoint_convert_hex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ep_carrier_endpoint_convert_hex: %v", err)
		}
	}

	if err = d.Set("ep_carrier_endpoint_header", flattenObjectUserRadiusDynamicMappingEpCarrierEndpointHeader2edl(o["ep-carrier-endpoint-header"], d, "ep_carrier_endpoint_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["ep-carrier-endpoint-header"], "ObjectUserRadiusDynamicMapping-EpCarrierEndpointHeader"); ok {
			if err = d.Set("ep_carrier_endpoint_header", vv); err != nil {
				return fmt.Errorf("Error reading ep_carrier_endpoint_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ep_carrier_endpoint_header: %v", err)
		}
	}

	if err = d.Set("ep_carrier_endpoint_header_suppress", flattenObjectUserRadiusDynamicMappingEpCarrierEndpointHeaderSuppress2edl(o["ep-carrier-endpoint-header-suppress"], d, "ep_carrier_endpoint_header_suppress")); err != nil {
		if vv, ok := fortiAPIPatch(o["ep-carrier-endpoint-header-suppress"], "ObjectUserRadiusDynamicMapping-EpCarrierEndpointHeaderSuppress"); ok {
			if err = d.Set("ep_carrier_endpoint_header_suppress", vv); err != nil {
				return fmt.Errorf("Error reading ep_carrier_endpoint_header_suppress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ep_carrier_endpoint_header_suppress: %v", err)
		}
	}

	if err = d.Set("ep_carrier_endpoint_prefix", flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefix2edl(o["ep-carrier-endpoint-prefix"], d, "ep_carrier_endpoint_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["ep-carrier-endpoint-prefix"], "ObjectUserRadiusDynamicMapping-EpCarrierEndpointPrefix"); ok {
			if err = d.Set("ep_carrier_endpoint_prefix", vv); err != nil {
				return fmt.Errorf("Error reading ep_carrier_endpoint_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ep_carrier_endpoint_prefix: %v", err)
		}
	}

	if err = d.Set("ep_carrier_endpoint_prefix_range_max", flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMax2edl(o["ep-carrier-endpoint-prefix-range-max"], d, "ep_carrier_endpoint_prefix_range_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["ep-carrier-endpoint-prefix-range-max"], "ObjectUserRadiusDynamicMapping-EpCarrierEndpointPrefixRangeMax"); ok {
			if err = d.Set("ep_carrier_endpoint_prefix_range_max", vv); err != nil {
				return fmt.Errorf("Error reading ep_carrier_endpoint_prefix_range_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ep_carrier_endpoint_prefix_range_max: %v", err)
		}
	}

	if err = d.Set("ep_carrier_endpoint_prefix_range_min", flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMin2edl(o["ep-carrier-endpoint-prefix-range-min"], d, "ep_carrier_endpoint_prefix_range_min")); err != nil {
		if vv, ok := fortiAPIPatch(o["ep-carrier-endpoint-prefix-range-min"], "ObjectUserRadiusDynamicMapping-EpCarrierEndpointPrefixRangeMin"); ok {
			if err = d.Set("ep_carrier_endpoint_prefix_range_min", vv); err != nil {
				return fmt.Errorf("Error reading ep_carrier_endpoint_prefix_range_min: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ep_carrier_endpoint_prefix_range_min: %v", err)
		}
	}

	if err = d.Set("ep_carrier_endpoint_prefix_string", flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixString2edl(o["ep-carrier-endpoint-prefix-string"], d, "ep_carrier_endpoint_prefix_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["ep-carrier-endpoint-prefix-string"], "ObjectUserRadiusDynamicMapping-EpCarrierEndpointPrefixString"); ok {
			if err = d.Set("ep_carrier_endpoint_prefix_string", vv); err != nil {
				return fmt.Errorf("Error reading ep_carrier_endpoint_prefix_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ep_carrier_endpoint_prefix_string: %v", err)
		}
	}

	if err = d.Set("ep_carrier_endpoint_source", flattenObjectUserRadiusDynamicMappingEpCarrierEndpointSource2edl(o["ep-carrier-endpoint-source"], d, "ep_carrier_endpoint_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["ep-carrier-endpoint-source"], "ObjectUserRadiusDynamicMapping-EpCarrierEndpointSource"); ok {
			if err = d.Set("ep_carrier_endpoint_source", vv); err != nil {
				return fmt.Errorf("Error reading ep_carrier_endpoint_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ep_carrier_endpoint_source: %v", err)
		}
	}

	if err = d.Set("ep_ip_header", flattenObjectUserRadiusDynamicMappingEpIpHeader2edl(o["ep-ip-header"], d, "ep_ip_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["ep-ip-header"], "ObjectUserRadiusDynamicMapping-EpIpHeader"); ok {
			if err = d.Set("ep_ip_header", vv); err != nil {
				return fmt.Errorf("Error reading ep_ip_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ep_ip_header: %v", err)
		}
	}

	if err = d.Set("ep_ip_header_suppress", flattenObjectUserRadiusDynamicMappingEpIpHeaderSuppress2edl(o["ep-ip-header-suppress"], d, "ep_ip_header_suppress")); err != nil {
		if vv, ok := fortiAPIPatch(o["ep-ip-header-suppress"], "ObjectUserRadiusDynamicMapping-EpIpHeaderSuppress"); ok {
			if err = d.Set("ep_ip_header_suppress", vv); err != nil {
				return fmt.Errorf("Error reading ep_ip_header_suppress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ep_ip_header_suppress: %v", err)
		}
	}

	if err = d.Set("ep_missing_header_fallback", flattenObjectUserRadiusDynamicMappingEpMissingHeaderFallback2edl(o["ep-missing-header-fallback"], d, "ep_missing_header_fallback")); err != nil {
		if vv, ok := fortiAPIPatch(o["ep-missing-header-fallback"], "ObjectUserRadiusDynamicMapping-EpMissingHeaderFallback"); ok {
			if err = d.Set("ep_missing_header_fallback", vv); err != nil {
				return fmt.Errorf("Error reading ep_missing_header_fallback: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ep_missing_header_fallback: %v", err)
		}
	}

	if err = d.Set("ep_profile_query_type", flattenObjectUserRadiusDynamicMappingEpProfileQueryType2edl(o["ep-profile-query-type"], d, "ep_profile_query_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ep-profile-query-type"], "ObjectUserRadiusDynamicMapping-EpProfileQueryType"); ok {
			if err = d.Set("ep_profile_query_type", vv); err != nil {
				return fmt.Errorf("Error reading ep_profile_query_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ep_profile_query_type: %v", err)
		}
	}

	if err = d.Set("group_override_attr_type", flattenObjectUserRadiusDynamicMappingGroupOverrideAttrType2edl(o["group-override-attr-type"], d, "group_override_attr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-override-attr-type"], "ObjectUserRadiusDynamicMapping-GroupOverrideAttrType"); ok {
			if err = d.Set("group_override_attr_type", vv); err != nil {
				return fmt.Errorf("Error reading group_override_attr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_override_attr_type: %v", err)
		}
	}

	if err = d.Set("h3c_compatibility", flattenObjectUserRadiusDynamicMappingH3CCompatibility2edl(o["h3c-compatibility"], d, "h3c_compatibility")); err != nil {
		if vv, ok := fortiAPIPatch(o["h3c-compatibility"], "ObjectUserRadiusDynamicMapping-H3CCompatibility"); ok {
			if err = d.Set("h3c_compatibility", vv); err != nil {
				return fmt.Errorf("Error reading h3c_compatibility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h3c_compatibility: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectUserRadiusDynamicMappingInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectUserRadiusDynamicMapping-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenObjectUserRadiusDynamicMappingInterfaceSelectMethod2edl(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "ObjectUserRadiusDynamicMapping-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("mac_case", flattenObjectUserRadiusDynamicMappingMacCase2edl(o["mac-case"], d, "mac_case")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-case"], "ObjectUserRadiusDynamicMapping-MacCase"); ok {
			if err = d.Set("mac_case", vv); err != nil {
				return fmt.Errorf("Error reading mac_case: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_case: %v", err)
		}
	}

	if err = d.Set("mac_password_delimiter", flattenObjectUserRadiusDynamicMappingMacPasswordDelimiter2edl(o["mac-password-delimiter"], d, "mac_password_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-password-delimiter"], "ObjectUserRadiusDynamicMapping-MacPasswordDelimiter"); ok {
			if err = d.Set("mac_password_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_username_delimiter", flattenObjectUserRadiusDynamicMappingMacUsernameDelimiter2edl(o["mac-username-delimiter"], d, "mac_username_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-username-delimiter"], "ObjectUserRadiusDynamicMapping-MacUsernameDelimiter"); ok {
			if err = d.Set("mac_username_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
		}
	}

	if err = d.Set("nas_id", flattenObjectUserRadiusDynamicMappingNasId2edl(o["nas-id"], d, "nas_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["nas-id"], "ObjectUserRadiusDynamicMapping-NasId"); ok {
			if err = d.Set("nas_id", vv); err != nil {
				return fmt.Errorf("Error reading nas_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nas_id: %v", err)
		}
	}

	if err = d.Set("nas_id_type", flattenObjectUserRadiusDynamicMappingNasIdType2edl(o["nas-id-type"], d, "nas_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["nas-id-type"], "ObjectUserRadiusDynamicMapping-NasIdType"); ok {
			if err = d.Set("nas_id_type", vv); err != nil {
				return fmt.Errorf("Error reading nas_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nas_id_type: %v", err)
		}
	}

	if err = d.Set("nas_ip", flattenObjectUserRadiusDynamicMappingNasIp2edl(o["nas-ip"], d, "nas_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["nas-ip"], "ObjectUserRadiusDynamicMapping-NasIp"); ok {
			if err = d.Set("nas_ip", vv); err != nil {
				return fmt.Errorf("Error reading nas_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nas_ip: %v", err)
		}
	}

	if err = d.Set("password_encoding", flattenObjectUserRadiusDynamicMappingPasswordEncoding2edl(o["password-encoding"], d, "password_encoding")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-encoding"], "ObjectUserRadiusDynamicMapping-PasswordEncoding"); ok {
			if err = d.Set("password_encoding", vv); err != nil {
				return fmt.Errorf("Error reading password_encoding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_encoding: %v", err)
		}
	}

	if err = d.Set("password_renewal", flattenObjectUserRadiusDynamicMappingPasswordRenewal2edl(o["password-renewal"], d, "password_renewal")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-renewal"], "ObjectUserRadiusDynamicMapping-PasswordRenewal"); ok {
			if err = d.Set("password_renewal", vv); err != nil {
				return fmt.Errorf("Error reading password_renewal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_renewal: %v", err)
		}
	}

	if err = d.Set("radius_coa", flattenObjectUserRadiusDynamicMappingRadiusCoa2edl(o["radius-coa"], d, "radius_coa")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-coa"], "ObjectUserRadiusDynamicMapping-RadiusCoa"); ok {
			if err = d.Set("radius_coa", vv); err != nil {
				return fmt.Errorf("Error reading radius_coa: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_coa: %v", err)
		}
	}

	if err = d.Set("radius_port", flattenObjectUserRadiusDynamicMappingRadiusPort2edl(o["radius-port"], d, "radius_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-port"], "ObjectUserRadiusDynamicMapping-RadiusPort"); ok {
			if err = d.Set("radius_port", vv); err != nil {
				return fmt.Errorf("Error reading radius_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_port: %v", err)
		}
	}

	if err = d.Set("require_message_authenticator", flattenObjectUserRadiusDynamicMappingRequireMessageAuthenticator2edl(o["require-message-authenticator"], d, "require_message_authenticator")); err != nil {
		if vv, ok := fortiAPIPatch(o["require-message-authenticator"], "ObjectUserRadiusDynamicMapping-RequireMessageAuthenticator"); ok {
			if err = d.Set("require_message_authenticator", vv); err != nil {
				return fmt.Errorf("Error reading require_message_authenticator: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading require_message_authenticator: %v", err)
		}
	}

	if err = d.Set("rsso", flattenObjectUserRadiusDynamicMappingRsso2edl(o["rsso"], d, "rsso")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso"], "ObjectUserRadiusDynamicMapping-Rsso"); ok {
			if err = d.Set("rsso", vv); err != nil {
				return fmt.Errorf("Error reading rsso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso: %v", err)
		}
	}

	if err = d.Set("rsso_context_timeout", flattenObjectUserRadiusDynamicMappingRssoContextTimeout2edl(o["rsso-context-timeout"], d, "rsso_context_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-context-timeout"], "ObjectUserRadiusDynamicMapping-RssoContextTimeout"); ok {
			if err = d.Set("rsso_context_timeout", vv); err != nil {
				return fmt.Errorf("Error reading rsso_context_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_context_timeout: %v", err)
		}
	}

	if err = d.Set("rsso_endpoint_attribute", flattenObjectUserRadiusDynamicMappingRssoEndpointAttribute2edl(o["rsso-endpoint-attribute"], d, "rsso_endpoint_attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-endpoint-attribute"], "ObjectUserRadiusDynamicMapping-RssoEndpointAttribute"); ok {
			if err = d.Set("rsso_endpoint_attribute", vv); err != nil {
				return fmt.Errorf("Error reading rsso_endpoint_attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_endpoint_attribute: %v", err)
		}
	}

	if err = d.Set("rsso_endpoint_block_attribute", flattenObjectUserRadiusDynamicMappingRssoEndpointBlockAttribute2edl(o["rsso-endpoint-block-attribute"], d, "rsso_endpoint_block_attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-endpoint-block-attribute"], "ObjectUserRadiusDynamicMapping-RssoEndpointBlockAttribute"); ok {
			if err = d.Set("rsso_endpoint_block_attribute", vv); err != nil {
				return fmt.Errorf("Error reading rsso_endpoint_block_attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_endpoint_block_attribute: %v", err)
		}
	}

	if err = d.Set("rsso_ep_one_ip_only", flattenObjectUserRadiusDynamicMappingRssoEpOneIpOnly2edl(o["rsso-ep-one-ip-only"], d, "rsso_ep_one_ip_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-ep-one-ip-only"], "ObjectUserRadiusDynamicMapping-RssoEpOneIpOnly"); ok {
			if err = d.Set("rsso_ep_one_ip_only", vv); err != nil {
				return fmt.Errorf("Error reading rsso_ep_one_ip_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_ep_one_ip_only: %v", err)
		}
	}

	if err = d.Set("rsso_flush_ip_session", flattenObjectUserRadiusDynamicMappingRssoFlushIpSession2edl(o["rsso-flush-ip-session"], d, "rsso_flush_ip_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-flush-ip-session"], "ObjectUserRadiusDynamicMapping-RssoFlushIpSession"); ok {
			if err = d.Set("rsso_flush_ip_session", vv); err != nil {
				return fmt.Errorf("Error reading rsso_flush_ip_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_flush_ip_session: %v", err)
		}
	}

	if err = d.Set("rsso_log_flags", flattenObjectUserRadiusDynamicMappingRssoLogFlags2edl(o["rsso-log-flags"], d, "rsso_log_flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-log-flags"], "ObjectUserRadiusDynamicMapping-RssoLogFlags"); ok {
			if err = d.Set("rsso_log_flags", vv); err != nil {
				return fmt.Errorf("Error reading rsso_log_flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_log_flags: %v", err)
		}
	}

	if err = d.Set("rsso_log_period", flattenObjectUserRadiusDynamicMappingRssoLogPeriod2edl(o["rsso-log-period"], d, "rsso_log_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-log-period"], "ObjectUserRadiusDynamicMapping-RssoLogPeriod"); ok {
			if err = d.Set("rsso_log_period", vv); err != nil {
				return fmt.Errorf("Error reading rsso_log_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_log_period: %v", err)
		}
	}

	if err = d.Set("rsso_radius_response", flattenObjectUserRadiusDynamicMappingRssoRadiusResponse2edl(o["rsso-radius-response"], d, "rsso_radius_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-radius-response"], "ObjectUserRadiusDynamicMapping-RssoRadiusResponse"); ok {
			if err = d.Set("rsso_radius_response", vv); err != nil {
				return fmt.Errorf("Error reading rsso_radius_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_radius_response: %v", err)
		}
	}

	if err = d.Set("rsso_radius_server_port", flattenObjectUserRadiusDynamicMappingRssoRadiusServerPort2edl(o["rsso-radius-server-port"], d, "rsso_radius_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-radius-server-port"], "ObjectUserRadiusDynamicMapping-RssoRadiusServerPort"); ok {
			if err = d.Set("rsso_radius_server_port", vv); err != nil {
				return fmt.Errorf("Error reading rsso_radius_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_radius_server_port: %v", err)
		}
	}

	if err = d.Set("rsso_validate_request_secret", flattenObjectUserRadiusDynamicMappingRssoValidateRequestSecret2edl(o["rsso-validate-request-secret"], d, "rsso_validate_request_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-validate-request-secret"], "ObjectUserRadiusDynamicMapping-RssoValidateRequestSecret"); ok {
			if err = d.Set("rsso_validate_request_secret", vv); err != nil {
				return fmt.Errorf("Error reading rsso_validate_request_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_validate_request_secret: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenObjectUserRadiusDynamicMappingSecondaryServer2edl(o["secondary-server"], d, "secondary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-server"], "ObjectUserRadiusDynamicMapping-SecondaryServer"); ok {
			if err = d.Set("secondary_server", vv); err != nil {
				return fmt.Errorf("Error reading secondary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("server", flattenObjectUserRadiusDynamicMappingServer2edl(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "ObjectUserRadiusDynamicMapping-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server_identity_check", flattenObjectUserRadiusDynamicMappingServerIdentityCheck2edl(o["server-identity-check"], d, "server_identity_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-identity-check"], "ObjectUserRadiusDynamicMapping-ServerIdentityCheck"); ok {
			if err = d.Set("server_identity_check", vv); err != nil {
				return fmt.Errorf("Error reading server_identity_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_identity_check: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenObjectUserRadiusDynamicMappingSourceIp2edl(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "ObjectUserRadiusDynamicMapping-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip_interface", flattenObjectUserRadiusDynamicMappingSourceIpInterface2edl(o["source-ip-interface"], d, "source_ip_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-interface"], "ObjectUserRadiusDynamicMapping-SourceIpInterface"); ok {
			if err = d.Set("source_ip_interface", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_interface: %v", err)
		}
	}

	if err = d.Set("sso_attribute", flattenObjectUserRadiusDynamicMappingSsoAttribute2edl(o["sso-attribute"], d, "sso_attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-attribute"], "ObjectUserRadiusDynamicMapping-SsoAttribute"); ok {
			if err = d.Set("sso_attribute", vv); err != nil {
				return fmt.Errorf("Error reading sso_attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_attribute: %v", err)
		}
	}

	if err = d.Set("sso_attribute_key", flattenObjectUserRadiusDynamicMappingSsoAttributeKey2edl(o["sso-attribute-key"], d, "sso_attribute_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-attribute-key"], "ObjectUserRadiusDynamicMapping-SsoAttributeKey"); ok {
			if err = d.Set("sso_attribute_key", vv); err != nil {
				return fmt.Errorf("Error reading sso_attribute_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_attribute_key: %v", err)
		}
	}

	if err = d.Set("sso_attribute_value_override", flattenObjectUserRadiusDynamicMappingSsoAttributeValueOverride2edl(o["sso-attribute-value-override"], d, "sso_attribute_value_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-attribute-value-override"], "ObjectUserRadiusDynamicMapping-SsoAttributeValueOverride"); ok {
			if err = d.Set("sso_attribute_value_override", vv); err != nil {
				return fmt.Errorf("Error reading sso_attribute_value_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_attribute_value_override: %v", err)
		}
	}

	if err = d.Set("status_ttl", flattenObjectUserRadiusDynamicMappingStatusTtl2edl(o["status-ttl"], d, "status_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["status-ttl"], "ObjectUserRadiusDynamicMapping-StatusTtl"); ok {
			if err = d.Set("status_ttl", vv); err != nil {
				return fmt.Errorf("Error reading status_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status_ttl: %v", err)
		}
	}

	if err = d.Set("switch_controller_acct_fast_framedip_detect", flattenObjectUserRadiusDynamicMappingSwitchControllerAcctFastFramedipDetect2edl(o["switch-controller-acct-fast-framedip-detect"], d, "switch_controller_acct_fast_framedip_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-acct-fast-framedip-detect"], "ObjectUserRadiusDynamicMapping-SwitchControllerAcctFastFramedipDetect"); ok {
			if err = d.Set("switch_controller_acct_fast_framedip_detect", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_acct_fast_framedip_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_acct_fast_framedip_detect: %v", err)
		}
	}

	if err = d.Set("switch_controller_nas_ip_dynamic", flattenObjectUserRadiusDynamicMappingSwitchControllerNasIpDynamic2edl(o["switch-controller-nas-ip-dynamic"], d, "switch_controller_nas_ip_dynamic")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-nas-ip-dynamic"], "ObjectUserRadiusDynamicMapping-SwitchControllerNasIpDynamic"); ok {
			if err = d.Set("switch_controller_nas_ip_dynamic", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_nas_ip_dynamic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_nas_ip_dynamic: %v", err)
		}
	}

	if err = d.Set("switch_controller_service_type", flattenObjectUserRadiusDynamicMappingSwitchControllerServiceType2edl(o["switch-controller-service-type"], d, "switch_controller_service_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-service-type"], "ObjectUserRadiusDynamicMapping-SwitchControllerServiceType"); ok {
			if err = d.Set("switch_controller_service_type", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_service_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_service_type: %v", err)
		}
	}

	if err = d.Set("tertiary_server", flattenObjectUserRadiusDynamicMappingTertiaryServer2edl(o["tertiary-server"], d, "tertiary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tertiary-server"], "ObjectUserRadiusDynamicMapping-TertiaryServer"); ok {
			if err = d.Set("tertiary_server", vv); err != nil {
				return fmt.Errorf("Error reading tertiary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tertiary_server: %v", err)
		}
	}

	if err = d.Set("timeout", flattenObjectUserRadiusDynamicMappingTimeout2edl(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "ObjectUserRadiusDynamicMapping-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("tls_min_proto_version", flattenObjectUserRadiusDynamicMappingTlsMinProtoVersion2edl(o["tls-min-proto-version"], d, "tls_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["tls-min-proto-version"], "ObjectUserRadiusDynamicMapping-TlsMinProtoVersion"); ok {
			if err = d.Set("tls_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading tls_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tls_min_proto_version: %v", err)
		}
	}

	if err = d.Set("transport_protocol", flattenObjectUserRadiusDynamicMappingTransportProtocol2edl(o["transport-protocol"], d, "transport_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["transport-protocol"], "ObjectUserRadiusDynamicMapping-TransportProtocol"); ok {
			if err = d.Set("transport_protocol", vv); err != nil {
				return fmt.Errorf("Error reading transport_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transport_protocol: %v", err)
		}
	}

	if err = d.Set("use_group_for_profile", flattenObjectUserRadiusDynamicMappingUseGroupForProfile2edl(o["use-group-for-profile"], d, "use_group_for_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-group-for-profile"], "ObjectUserRadiusDynamicMapping-UseGroupForProfile"); ok {
			if err = d.Set("use_group_for_profile", vv); err != nil {
				return fmt.Errorf("Error reading use_group_for_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_group_for_profile: %v", err)
		}
	}

	if err = d.Set("use_management_vdom", flattenObjectUserRadiusDynamicMappingUseManagementVdom2edl(o["use-management-vdom"], d, "use_management_vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-management-vdom"], "ObjectUserRadiusDynamicMapping-UseManagementVdom"); ok {
			if err = d.Set("use_management_vdom", vv); err != nil {
				return fmt.Errorf("Error reading use_management_vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_management_vdom: %v", err)
		}
	}

	if err = d.Set("username_case_sensitive", flattenObjectUserRadiusDynamicMappingUsernameCaseSensitive2edl(o["username-case-sensitive"], d, "username_case_sensitive")); err != nil {
		if vv, ok := fortiAPIPatch(o["username-case-sensitive"], "ObjectUserRadiusDynamicMapping-UsernameCaseSensitive"); ok {
			if err = d.Set("username_case_sensitive", vv); err != nil {
				return fmt.Errorf("Error reading username_case_sensitive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username_case_sensitive: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenObjectUserRadiusDynamicMappingVrfSelect2edl(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "ObjectUserRadiusDynamicMapping-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenObjectUserRadiusDynamicMappingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserRadiusDynamicMappingScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectUserRadiusDynamicMappingScopeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectUserRadiusDynamicMappingScopeVdom2edl(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserRadiusDynamicMappingScopeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingScopeVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountKeyCertField2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountKeyProcessing2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectUserRadiusDynamicMappingAccountingServerId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandObjectUserRadiusDynamicMappingAccountingServerInterface2edl(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-select-method"], _ = expandObjectUserRadiusDynamicMappingAccountingServerInterfaceSelectMethod2edl(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectUserRadiusDynamicMappingAccountingServerPort2edl(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secret"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secret"], _ = expandObjectUserRadiusDynamicMappingAccountingServerSecret2edl(d, i["secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandObjectUserRadiusDynamicMappingAccountingServerServer2edl(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip"], _ = expandObjectUserRadiusDynamicMappingAccountingServerSourceIp2edl(d, i["source_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectUserRadiusDynamicMappingAccountingServerStatus2edl(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf-select"], _ = expandObjectUserRadiusDynamicMappingAccountingServerVrfSelect2edl(d, i["vrf_select"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerInterfaceSelectMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerSecret2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerSourceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerVrfSelect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAcctAllServers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAcctInterimInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAllUsergroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAuthType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingCaCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserRadiusDynamicMappingCallStationIdType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingClass2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingClientCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserRadiusDynamicMappingDelimiter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpCarrierEndpointAttribute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpCarrierEndpointBlockAttribute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpContextTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpFlushIpSession2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpHoldTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpHttpHeader2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpHttpHeaderFallback2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpHttpHeaderStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpHttpHeaderSuppress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpLogDynFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingDpLogPeriod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpMemPercent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpProfileAttribute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpProfileAttributeKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpRadiusResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpRadiusServerPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpSecret2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingDpValidateRequestSecret2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDynamicProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEndpointTranslation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointConvertHex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointHeader2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointHeaderSuppress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMax2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointSource2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpIpHeader2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpIpHeaderSuppress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpMissingHeaderFallback2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpProfileQueryType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingGroupOverrideAttrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingH3CCompatibility2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserRadiusDynamicMappingInterfaceSelectMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingMacCase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingMacPasswordDelimiter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingMacUsernameDelimiter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingNasId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingNasIdType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingNasIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingPasswordEncoding2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingPasswordRenewal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRadiusCoa2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRadiusPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRequireMessageAuthenticator2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRsso2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoContextTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoEndpointAttribute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoEndpointBlockAttribute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoEpOneIpOnly2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoFlushIpSession2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoLogFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingRssoLogPeriod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoRadiusResponse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoRadiusServerPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoSecret2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingRssoValidateRequestSecret2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSecondarySecret2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingSecondaryServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSecret2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingServerIdentityCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSourceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSourceIpInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingSsoAttribute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSsoAttributeKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSsoAttributeValueOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingStatusTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSwitchControllerAcctFastFramedipDetect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSwitchControllerNasIpDynamic2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSwitchControllerServiceType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingTertiarySecret2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingTertiaryServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingTlsMinProtoVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingTransportProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingUseGroupForProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingUseManagementVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingUsernameCaseSensitive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingVrfSelect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserRadiusDynamicMapping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_scope"); ok || d.HasChange("_scope") {
		t, err := expandObjectUserRadiusDynamicMappingScope2edl(d, v, "_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_scope"] = t
		}
	}

	if v, ok := d.GetOk("account_key_cert_field"); ok || d.HasChange("account_key_cert_field") {
		t, err := expandObjectUserRadiusDynamicMappingAccountKeyCertField2edl(d, v, "account_key_cert_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-cert-field"] = t
		}
	}

	if v, ok := d.GetOk("account_key_processing"); ok || d.HasChange("account_key_processing") {
		t, err := expandObjectUserRadiusDynamicMappingAccountKeyProcessing2edl(d, v, "account_key_processing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-processing"] = t
		}
	}

	if v, ok := d.GetOk("accounting_server"); ok || d.HasChange("accounting_server") {
		t, err := expandObjectUserRadiusDynamicMappingAccountingServer2edl(d, v, "accounting_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accounting-server"] = t
		}
	}

	if v, ok := d.GetOk("acct_all_servers"); ok || d.HasChange("acct_all_servers") {
		t, err := expandObjectUserRadiusDynamicMappingAcctAllServers2edl(d, v, "acct_all_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-all-servers"] = t
		}
	}

	if v, ok := d.GetOk("acct_interim_interval"); ok || d.HasChange("acct_interim_interval") {
		t, err := expandObjectUserRadiusDynamicMappingAcctInterimInterval2edl(d, v, "acct_interim_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-interim-interval"] = t
		}
	}

	if v, ok := d.GetOk("all_usergroup"); ok || d.HasChange("all_usergroup") {
		t, err := expandObjectUserRadiusDynamicMappingAllUsergroup2edl(d, v, "all_usergroup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["all-usergroup"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok || d.HasChange("auth_type") {
		t, err := expandObjectUserRadiusDynamicMappingAuthType2edl(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("ca_cert"); ok || d.HasChange("ca_cert") {
		t, err := expandObjectUserRadiusDynamicMappingCaCert2edl(d, v, "ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("call_station_id_type"); ok || d.HasChange("call_station_id_type") {
		t, err := expandObjectUserRadiusDynamicMappingCallStationIdType2edl(d, v, "call_station_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-station-id-type"] = t
		}
	}

	if v, ok := d.GetOk("class"); ok || d.HasChange("class") {
		t, err := expandObjectUserRadiusDynamicMappingClass2edl(d, v, "class")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandObjectUserRadiusDynamicMappingClientCert2edl(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("delimiter"); ok || d.HasChange("delimiter") {
		t, err := expandObjectUserRadiusDynamicMappingDelimiter2edl(d, v, "delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delimiter"] = t
		}
	}

	if v, ok := d.GetOk("dp_carrier_endpoint_attribute"); ok || d.HasChange("dp_carrier_endpoint_attribute") {
		t, err := expandObjectUserRadiusDynamicMappingDpCarrierEndpointAttribute2edl(d, v, "dp_carrier_endpoint_attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-carrier-endpoint-attribute"] = t
		}
	}

	if v, ok := d.GetOk("dp_carrier_endpoint_block_attribute"); ok || d.HasChange("dp_carrier_endpoint_block_attribute") {
		t, err := expandObjectUserRadiusDynamicMappingDpCarrierEndpointBlockAttribute2edl(d, v, "dp_carrier_endpoint_block_attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-carrier-endpoint-block-attribute"] = t
		}
	}

	if v, ok := d.GetOk("dp_context_timeout"); ok || d.HasChange("dp_context_timeout") {
		t, err := expandObjectUserRadiusDynamicMappingDpContextTimeout2edl(d, v, "dp_context_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-context-timeout"] = t
		}
	}

	if v, ok := d.GetOk("dp_flush_ip_session"); ok || d.HasChange("dp_flush_ip_session") {
		t, err := expandObjectUserRadiusDynamicMappingDpFlushIpSession2edl(d, v, "dp_flush_ip_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-flush-ip-session"] = t
		}
	}

	if v, ok := d.GetOk("dp_hold_time"); ok || d.HasChange("dp_hold_time") {
		t, err := expandObjectUserRadiusDynamicMappingDpHoldTime2edl(d, v, "dp_hold_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-hold-time"] = t
		}
	}

	if v, ok := d.GetOk("dp_http_header"); ok || d.HasChange("dp_http_header") {
		t, err := expandObjectUserRadiusDynamicMappingDpHttpHeader2edl(d, v, "dp_http_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-http-header"] = t
		}
	}

	if v, ok := d.GetOk("dp_http_header_fallback"); ok || d.HasChange("dp_http_header_fallback") {
		t, err := expandObjectUserRadiusDynamicMappingDpHttpHeaderFallback2edl(d, v, "dp_http_header_fallback")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-http-header-fallback"] = t
		}
	}

	if v, ok := d.GetOk("dp_http_header_status"); ok || d.HasChange("dp_http_header_status") {
		t, err := expandObjectUserRadiusDynamicMappingDpHttpHeaderStatus2edl(d, v, "dp_http_header_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-http-header-status"] = t
		}
	}

	if v, ok := d.GetOk("dp_http_header_suppress"); ok || d.HasChange("dp_http_header_suppress") {
		t, err := expandObjectUserRadiusDynamicMappingDpHttpHeaderSuppress2edl(d, v, "dp_http_header_suppress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-http-header-suppress"] = t
		}
	}

	if v, ok := d.GetOk("dp_log_dyn_flags"); ok || d.HasChange("dp_log_dyn_flags") {
		t, err := expandObjectUserRadiusDynamicMappingDpLogDynFlags2edl(d, v, "dp_log_dyn_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-log-dyn_flags"] = t
		}
	}

	if v, ok := d.GetOk("dp_log_period"); ok || d.HasChange("dp_log_period") {
		t, err := expandObjectUserRadiusDynamicMappingDpLogPeriod2edl(d, v, "dp_log_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-log-period"] = t
		}
	}

	if v, ok := d.GetOk("dp_mem_percent"); ok || d.HasChange("dp_mem_percent") {
		t, err := expandObjectUserRadiusDynamicMappingDpMemPercent2edl(d, v, "dp_mem_percent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-mem-percent"] = t
		}
	}

	if v, ok := d.GetOk("dp_profile_attribute"); ok || d.HasChange("dp_profile_attribute") {
		t, err := expandObjectUserRadiusDynamicMappingDpProfileAttribute2edl(d, v, "dp_profile_attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-profile-attribute"] = t
		}
	}

	if v, ok := d.GetOk("dp_profile_attribute_key"); ok || d.HasChange("dp_profile_attribute_key") {
		t, err := expandObjectUserRadiusDynamicMappingDpProfileAttributeKey2edl(d, v, "dp_profile_attribute_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-profile-attribute-key"] = t
		}
	}

	if v, ok := d.GetOk("dp_radius_response"); ok || d.HasChange("dp_radius_response") {
		t, err := expandObjectUserRadiusDynamicMappingDpRadiusResponse2edl(d, v, "dp_radius_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-radius-response"] = t
		}
	}

	if v, ok := d.GetOk("dp_radius_server_port"); ok || d.HasChange("dp_radius_server_port") {
		t, err := expandObjectUserRadiusDynamicMappingDpRadiusServerPort2edl(d, v, "dp_radius_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-radius-server-port"] = t
		}
	}

	if v, ok := d.GetOk("dp_secret"); ok || d.HasChange("dp_secret") {
		t, err := expandObjectUserRadiusDynamicMappingDpSecret2edl(d, v, "dp_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-secret"] = t
		}
	}

	if v, ok := d.GetOk("dp_validate_request_secret"); ok || d.HasChange("dp_validate_request_secret") {
		t, err := expandObjectUserRadiusDynamicMappingDpValidateRequestSecret2edl(d, v, "dp_validate_request_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dp-validate-request-secret"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_profile"); ok || d.HasChange("dynamic_profile") {
		t, err := expandObjectUserRadiusDynamicMappingDynamicProfile2edl(d, v, "dynamic_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-profile"] = t
		}
	}

	if v, ok := d.GetOk("endpoint_translation"); ok || d.HasChange("endpoint_translation") {
		t, err := expandObjectUserRadiusDynamicMappingEndpointTranslation2edl(d, v, "endpoint_translation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endpoint-translation"] = t
		}
	}

	if v, ok := d.GetOk("ep_carrier_endpoint_convert_hex"); ok || d.HasChange("ep_carrier_endpoint_convert_hex") {
		t, err := expandObjectUserRadiusDynamicMappingEpCarrierEndpointConvertHex2edl(d, v, "ep_carrier_endpoint_convert_hex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ep-carrier-endpoint-convert-hex"] = t
		}
	}

	if v, ok := d.GetOk("ep_carrier_endpoint_header"); ok || d.HasChange("ep_carrier_endpoint_header") {
		t, err := expandObjectUserRadiusDynamicMappingEpCarrierEndpointHeader2edl(d, v, "ep_carrier_endpoint_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ep-carrier-endpoint-header"] = t
		}
	}

	if v, ok := d.GetOk("ep_carrier_endpoint_header_suppress"); ok || d.HasChange("ep_carrier_endpoint_header_suppress") {
		t, err := expandObjectUserRadiusDynamicMappingEpCarrierEndpointHeaderSuppress2edl(d, v, "ep_carrier_endpoint_header_suppress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ep-carrier-endpoint-header-suppress"] = t
		}
	}

	if v, ok := d.GetOk("ep_carrier_endpoint_prefix"); ok || d.HasChange("ep_carrier_endpoint_prefix") {
		t, err := expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefix2edl(d, v, "ep_carrier_endpoint_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ep-carrier-endpoint-prefix"] = t
		}
	}

	if v, ok := d.GetOk("ep_carrier_endpoint_prefix_range_max"); ok || d.HasChange("ep_carrier_endpoint_prefix_range_max") {
		t, err := expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMax2edl(d, v, "ep_carrier_endpoint_prefix_range_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ep-carrier-endpoint-prefix-range-max"] = t
		}
	}

	if v, ok := d.GetOk("ep_carrier_endpoint_prefix_range_min"); ok || d.HasChange("ep_carrier_endpoint_prefix_range_min") {
		t, err := expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMin2edl(d, v, "ep_carrier_endpoint_prefix_range_min")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ep-carrier-endpoint-prefix-range-min"] = t
		}
	}

	if v, ok := d.GetOk("ep_carrier_endpoint_prefix_string"); ok || d.HasChange("ep_carrier_endpoint_prefix_string") {
		t, err := expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixString2edl(d, v, "ep_carrier_endpoint_prefix_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ep-carrier-endpoint-prefix-string"] = t
		}
	}

	if v, ok := d.GetOk("ep_carrier_endpoint_source"); ok || d.HasChange("ep_carrier_endpoint_source") {
		t, err := expandObjectUserRadiusDynamicMappingEpCarrierEndpointSource2edl(d, v, "ep_carrier_endpoint_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ep-carrier-endpoint-source"] = t
		}
	}

	if v, ok := d.GetOk("ep_ip_header"); ok || d.HasChange("ep_ip_header") {
		t, err := expandObjectUserRadiusDynamicMappingEpIpHeader2edl(d, v, "ep_ip_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ep-ip-header"] = t
		}
	}

	if v, ok := d.GetOk("ep_ip_header_suppress"); ok || d.HasChange("ep_ip_header_suppress") {
		t, err := expandObjectUserRadiusDynamicMappingEpIpHeaderSuppress2edl(d, v, "ep_ip_header_suppress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ep-ip-header-suppress"] = t
		}
	}

	if v, ok := d.GetOk("ep_missing_header_fallback"); ok || d.HasChange("ep_missing_header_fallback") {
		t, err := expandObjectUserRadiusDynamicMappingEpMissingHeaderFallback2edl(d, v, "ep_missing_header_fallback")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ep-missing-header-fallback"] = t
		}
	}

	if v, ok := d.GetOk("ep_profile_query_type"); ok || d.HasChange("ep_profile_query_type") {
		t, err := expandObjectUserRadiusDynamicMappingEpProfileQueryType2edl(d, v, "ep_profile_query_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ep-profile-query-type"] = t
		}
	}

	if v, ok := d.GetOk("group_override_attr_type"); ok || d.HasChange("group_override_attr_type") {
		t, err := expandObjectUserRadiusDynamicMappingGroupOverrideAttrType2edl(d, v, "group_override_attr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-override-attr-type"] = t
		}
	}

	if v, ok := d.GetOk("h3c_compatibility"); ok || d.HasChange("h3c_compatibility") {
		t, err := expandObjectUserRadiusDynamicMappingH3CCompatibility2edl(d, v, "h3c_compatibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3c-compatibility"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectUserRadiusDynamicMappingInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandObjectUserRadiusDynamicMappingInterfaceSelectMethod2edl(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("mac_case"); ok || d.HasChange("mac_case") {
		t, err := expandObjectUserRadiusDynamicMappingMacCase2edl(d, v, "mac_case")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-case"] = t
		}
	}

	if v, ok := d.GetOk("mac_password_delimiter"); ok || d.HasChange("mac_password_delimiter") {
		t, err := expandObjectUserRadiusDynamicMappingMacPasswordDelimiter2edl(d, v, "mac_password_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-password-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_username_delimiter"); ok || d.HasChange("mac_username_delimiter") {
		t, err := expandObjectUserRadiusDynamicMappingMacUsernameDelimiter2edl(d, v, "mac_username_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-username-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("nas_id"); ok || d.HasChange("nas_id") {
		t, err := expandObjectUserRadiusDynamicMappingNasId2edl(d, v, "nas_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-id"] = t
		}
	}

	if v, ok := d.GetOk("nas_id_type"); ok || d.HasChange("nas_id_type") {
		t, err := expandObjectUserRadiusDynamicMappingNasIdType2edl(d, v, "nas_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-id-type"] = t
		}
	}

	if v, ok := d.GetOk("nas_ip"); ok || d.HasChange("nas_ip") {
		t, err := expandObjectUserRadiusDynamicMappingNasIp2edl(d, v, "nas_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-ip"] = t
		}
	}

	if v, ok := d.GetOk("password_encoding"); ok || d.HasChange("password_encoding") {
		t, err := expandObjectUserRadiusDynamicMappingPasswordEncoding2edl(d, v, "password_encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-encoding"] = t
		}
	}

	if v, ok := d.GetOk("password_renewal"); ok || d.HasChange("password_renewal") {
		t, err := expandObjectUserRadiusDynamicMappingPasswordRenewal2edl(d, v, "password_renewal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-renewal"] = t
		}
	}

	if v, ok := d.GetOk("radius_coa"); ok || d.HasChange("radius_coa") {
		t, err := expandObjectUserRadiusDynamicMappingRadiusCoa2edl(d, v, "radius_coa")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-coa"] = t
		}
	}

	if v, ok := d.GetOk("radius_port"); ok || d.HasChange("radius_port") {
		t, err := expandObjectUserRadiusDynamicMappingRadiusPort2edl(d, v, "radius_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-port"] = t
		}
	}

	if v, ok := d.GetOk("require_message_authenticator"); ok || d.HasChange("require_message_authenticator") {
		t, err := expandObjectUserRadiusDynamicMappingRequireMessageAuthenticator2edl(d, v, "require_message_authenticator")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["require-message-authenticator"] = t
		}
	}

	if v, ok := d.GetOk("rsso"); ok || d.HasChange("rsso") {
		t, err := expandObjectUserRadiusDynamicMappingRsso2edl(d, v, "rsso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso"] = t
		}
	}

	if v, ok := d.GetOk("rsso_context_timeout"); ok || d.HasChange("rsso_context_timeout") {
		t, err := expandObjectUserRadiusDynamicMappingRssoContextTimeout2edl(d, v, "rsso_context_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-context-timeout"] = t
		}
	}

	if v, ok := d.GetOk("rsso_endpoint_attribute"); ok || d.HasChange("rsso_endpoint_attribute") {
		t, err := expandObjectUserRadiusDynamicMappingRssoEndpointAttribute2edl(d, v, "rsso_endpoint_attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-endpoint-attribute"] = t
		}
	}

	if v, ok := d.GetOk("rsso_endpoint_block_attribute"); ok || d.HasChange("rsso_endpoint_block_attribute") {
		t, err := expandObjectUserRadiusDynamicMappingRssoEndpointBlockAttribute2edl(d, v, "rsso_endpoint_block_attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-endpoint-block-attribute"] = t
		}
	}

	if v, ok := d.GetOk("rsso_ep_one_ip_only"); ok || d.HasChange("rsso_ep_one_ip_only") {
		t, err := expandObjectUserRadiusDynamicMappingRssoEpOneIpOnly2edl(d, v, "rsso_ep_one_ip_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-ep-one-ip-only"] = t
		}
	}

	if v, ok := d.GetOk("rsso_flush_ip_session"); ok || d.HasChange("rsso_flush_ip_session") {
		t, err := expandObjectUserRadiusDynamicMappingRssoFlushIpSession2edl(d, v, "rsso_flush_ip_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-flush-ip-session"] = t
		}
	}

	if v, ok := d.GetOk("rsso_log_flags"); ok || d.HasChange("rsso_log_flags") {
		t, err := expandObjectUserRadiusDynamicMappingRssoLogFlags2edl(d, v, "rsso_log_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-log-flags"] = t
		}
	}

	if v, ok := d.GetOk("rsso_log_period"); ok || d.HasChange("rsso_log_period") {
		t, err := expandObjectUserRadiusDynamicMappingRssoLogPeriod2edl(d, v, "rsso_log_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-log-period"] = t
		}
	}

	if v, ok := d.GetOk("rsso_radius_response"); ok || d.HasChange("rsso_radius_response") {
		t, err := expandObjectUserRadiusDynamicMappingRssoRadiusResponse2edl(d, v, "rsso_radius_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-radius-response"] = t
		}
	}

	if v, ok := d.GetOk("rsso_radius_server_port"); ok || d.HasChange("rsso_radius_server_port") {
		t, err := expandObjectUserRadiusDynamicMappingRssoRadiusServerPort2edl(d, v, "rsso_radius_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-radius-server-port"] = t
		}
	}

	if v, ok := d.GetOk("rsso_secret"); ok || d.HasChange("rsso_secret") {
		t, err := expandObjectUserRadiusDynamicMappingRssoSecret2edl(d, v, "rsso_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-secret"] = t
		}
	}

	if v, ok := d.GetOk("rsso_validate_request_secret"); ok || d.HasChange("rsso_validate_request_secret") {
		t, err := expandObjectUserRadiusDynamicMappingRssoValidateRequestSecret2edl(d, v, "rsso_validate_request_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-validate-request-secret"] = t
		}
	}

	if v, ok := d.GetOk("secondary_secret"); ok || d.HasChange("secondary_secret") {
		t, err := expandObjectUserRadiusDynamicMappingSecondarySecret2edl(d, v, "secondary_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-secret"] = t
		}
	}

	if v, ok := d.GetOk("secondary_server"); ok || d.HasChange("secondary_server") {
		t, err := expandObjectUserRadiusDynamicMappingSecondaryServer2edl(d, v, "secondary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}

	if v, ok := d.GetOk("secret"); ok || d.HasChange("secret") {
		t, err := expandObjectUserRadiusDynamicMappingSecret2edl(d, v, "secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandObjectUserRadiusDynamicMappingServer2edl(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server_identity_check"); ok || d.HasChange("server_identity_check") {
		t, err := expandObjectUserRadiusDynamicMappingServerIdentityCheck2edl(d, v, "server_identity_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-identity-check"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandObjectUserRadiusDynamicMappingSourceIp2edl(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_interface"); ok || d.HasChange("source_ip_interface") {
		t, err := expandObjectUserRadiusDynamicMappingSourceIpInterface2edl(d, v, "source_ip_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-interface"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute"); ok || d.HasChange("sso_attribute") {
		t, err := expandObjectUserRadiusDynamicMappingSsoAttribute2edl(d, v, "sso_attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute_key"); ok || d.HasChange("sso_attribute_key") {
		t, err := expandObjectUserRadiusDynamicMappingSsoAttributeKey2edl(d, v, "sso_attribute_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute-key"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute_value_override"); ok || d.HasChange("sso_attribute_value_override") {
		t, err := expandObjectUserRadiusDynamicMappingSsoAttributeValueOverride2edl(d, v, "sso_attribute_value_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute-value-override"] = t
		}
	}

	if v, ok := d.GetOk("status_ttl"); ok || d.HasChange("status_ttl") {
		t, err := expandObjectUserRadiusDynamicMappingStatusTtl2edl(d, v, "status_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status-ttl"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_acct_fast_framedip_detect"); ok || d.HasChange("switch_controller_acct_fast_framedip_detect") {
		t, err := expandObjectUserRadiusDynamicMappingSwitchControllerAcctFastFramedipDetect2edl(d, v, "switch_controller_acct_fast_framedip_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-acct-fast-framedip-detect"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_nas_ip_dynamic"); ok || d.HasChange("switch_controller_nas_ip_dynamic") {
		t, err := expandObjectUserRadiusDynamicMappingSwitchControllerNasIpDynamic2edl(d, v, "switch_controller_nas_ip_dynamic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-nas-ip-dynamic"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_service_type"); ok || d.HasChange("switch_controller_service_type") {
		t, err := expandObjectUserRadiusDynamicMappingSwitchControllerServiceType2edl(d, v, "switch_controller_service_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-service-type"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_secret"); ok || d.HasChange("tertiary_secret") {
		t, err := expandObjectUserRadiusDynamicMappingTertiarySecret2edl(d, v, "tertiary_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-secret"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_server"); ok || d.HasChange("tertiary_server") {
		t, err := expandObjectUserRadiusDynamicMappingTertiaryServer2edl(d, v, "tertiary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-server"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandObjectUserRadiusDynamicMappingTimeout2edl(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("tls_min_proto_version"); ok || d.HasChange("tls_min_proto_version") {
		t, err := expandObjectUserRadiusDynamicMappingTlsMinProtoVersion2edl(d, v, "tls_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("transport_protocol"); ok || d.HasChange("transport_protocol") {
		t, err := expandObjectUserRadiusDynamicMappingTransportProtocol2edl(d, v, "transport_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transport-protocol"] = t
		}
	}

	if v, ok := d.GetOk("use_group_for_profile"); ok || d.HasChange("use_group_for_profile") {
		t, err := expandObjectUserRadiusDynamicMappingUseGroupForProfile2edl(d, v, "use_group_for_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-group-for-profile"] = t
		}
	}

	if v, ok := d.GetOk("use_management_vdom"); ok || d.HasChange("use_management_vdom") {
		t, err := expandObjectUserRadiusDynamicMappingUseManagementVdom2edl(d, v, "use_management_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-management-vdom"] = t
		}
	}

	if v, ok := d.GetOk("username_case_sensitive"); ok || d.HasChange("username_case_sensitive") {
		t, err := expandObjectUserRadiusDynamicMappingUsernameCaseSensitive2edl(d, v, "username_case_sensitive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-case-sensitive"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandObjectUserRadiusDynamicMappingVrfSelect2edl(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
