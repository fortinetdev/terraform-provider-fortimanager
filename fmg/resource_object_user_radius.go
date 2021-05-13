// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure RADIUS server entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectUserRadius() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserRadiusCreate,
		Read:   resourceObjectUserRadiusRead,
		Update: resourceObjectUserRadiusUpdate,
		Delete: resourceObjectUserRadiusDelete,

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
			"accounting_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface_select_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"secret": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"server": &schema.Schema{
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
				Computed: true,
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
			"class": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
										Computed: true,
									},
									"vdom": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"accounting_server": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"interface": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"interface_select_method": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"secret": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"server": &schema.Schema{
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
							Computed: true,
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
						"class": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"dp_carrier_endpoint_attribute": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dp_carrier_endpoint_block_attribute": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dp_context_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"dp_flush_ip_session": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dp_hold_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"dp_http_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dp_http_header_fallback": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dp_http_header_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dp_http_header_suppress": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"dp_mem_percent": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"dp_profile_attribute": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dp_profile_attribute_key": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dp_radius_response": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dp_radius_server_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"dp_secret": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"dp_validate_request_secret": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dynamic_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"endpoint_translation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ep_carrier_endpoint_convert_hex": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ep_carrier_endpoint_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ep_carrier_endpoint_header_suppress": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ep_carrier_endpoint_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ep_carrier_endpoint_prefix_range_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ep_carrier_endpoint_prefix_range_min": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ep_carrier_endpoint_prefix_string": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ep_carrier_endpoint_source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ep_ip_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ep_ip_header_suppress": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ep_missing_header_fallback": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ep_profile_query_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"group_override_attr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"h3c_compatibility": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface_select_method": &schema.Schema{
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
							Computed: true,
						},
						"rsso_endpoint_attribute": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rsso_endpoint_block_attribute": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rsso_ep_one_ip_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rsso_flush_ip_session": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"rsso_radius_response": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rsso_radius_server_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rsso_secret": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rsso_validate_request_secret": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secondary_secret": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"secondary_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secret": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sso_attribute": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sso_attribute_key": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sso_attribute_value_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"switch_controller_acct_fast_framedip_detect": &schema.Schema{
							Type:     schema.TypeInt,
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
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"tertiary_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"use_group_for_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
					},
				},
			},
			"group_override_attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"h3c_compatibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"rsso_endpoint_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso_endpoint_block_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso_ep_one_ip_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso_flush_ip_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"rsso_radius_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso_radius_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rsso_secret": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rsso_validate_request_secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary_secret": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"secondary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secret": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso_attribute_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso_attribute_value_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_acct_fast_framedip_detect": &schema.Schema{
				Type:     schema.TypeInt,
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
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"tertiary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectUserRadiusCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserRadius(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserRadius resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserRadius(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserRadius resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserRadiusRead(d, m)
}

func resourceObjectUserRadiusUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserRadius(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserRadius resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserRadius(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserRadius resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserRadiusRead(d, m)
}

func resourceObjectUserRadiusDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectUserRadius(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserRadius resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserRadiusRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectUserRadius(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserRadius resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserRadius(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserRadius resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserRadiusAccountingServer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserRadiusAccountingServerId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectUserRadius-AccountingServer-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenObjectUserRadiusAccountingServerInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "ObjectUserRadius-AccountingServer-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenObjectUserRadiusAccountingServerInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "ObjectUserRadius-AccountingServer-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectUserRadiusAccountingServerPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectUserRadius-AccountingServer-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secret"
		if _, ok := i["secret"]; ok {
			v := flattenObjectUserRadiusAccountingServerSecret(i["secret"], d, pre_append)
			tmp["secret"] = fortiAPISubPartPatch(v, "ObjectUserRadius-AccountingServer-Secret")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenObjectUserRadiusAccountingServerServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "ObjectUserRadius-AccountingServer-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenObjectUserRadiusAccountingServerSourceIp(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "ObjectUserRadius-AccountingServer-SourceIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectUserRadiusAccountingServerStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectUserRadius-AccountingServer-Status")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectUserRadiusAccountingServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusAccountingServerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusAccountingServerInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "auto",
			1: "sdwan",
			2: "specify",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusAccountingServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusAccountingServerSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusAccountingServerServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusAccountingServerSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusAccountingServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusAcctAllServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusAcctInterimInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusAllUsergroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "pap",
			1: "chap",
			3: "ms_chap",
			4: "ms_chap_v2",
			6: "auto",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusClass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserRadiusDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accounting_server"
		if _, ok := i["accounting-server"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServer(i["accounting-server"], d, pre_append)
			tmp["accounting_server"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-AccountingServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acct_all_servers"
		if _, ok := i["acct-all-servers"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAcctAllServers(i["acct-all-servers"], d, pre_append)
			tmp["acct_all_servers"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-AcctAllServers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acct_interim_interval"
		if _, ok := i["acct-interim-interval"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAcctInterimInterval(i["acct-interim-interval"], d, pre_append)
			tmp["acct_interim_interval"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-AcctInterimInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "all_usergroup"
		if _, ok := i["all-usergroup"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAllUsergroup(i["all-usergroup"], d, pre_append)
			tmp["all_usergroup"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-AllUsergroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_type"
		if _, ok := i["auth-type"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAuthType(i["auth-type"], d, pre_append)
			tmp["auth_type"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-AuthType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class"
		if _, ok := i["class"]; ok {
			v := flattenObjectUserRadiusDynamicMappingClass(i["class"], d, pre_append)
			tmp["class"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-Class")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_carrier_endpoint_attribute"
		if _, ok := i["dp-carrier-endpoint-attribute"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpCarrierEndpointAttribute(i["dp-carrier-endpoint-attribute"], d, pre_append)
			tmp["dp_carrier_endpoint_attribute"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpCarrierEndpointAttribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_carrier_endpoint_block_attribute"
		if _, ok := i["dp-carrier-endpoint-block-attribute"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpCarrierEndpointBlockAttribute(i["dp-carrier-endpoint-block-attribute"], d, pre_append)
			tmp["dp_carrier_endpoint_block_attribute"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpCarrierEndpointBlockAttribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_context_timeout"
		if _, ok := i["dp-context-timeout"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpContextTimeout(i["dp-context-timeout"], d, pre_append)
			tmp["dp_context_timeout"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpContextTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_flush_ip_session"
		if _, ok := i["dp-flush-ip-session"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpFlushIpSession(i["dp-flush-ip-session"], d, pre_append)
			tmp["dp_flush_ip_session"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpFlushIpSession")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_hold_time"
		if _, ok := i["dp-hold-time"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpHoldTime(i["dp-hold-time"], d, pre_append)
			tmp["dp_hold_time"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpHoldTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header"
		if _, ok := i["dp-http-header"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpHttpHeader(i["dp-http-header"], d, pre_append)
			tmp["dp_http_header"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpHttpHeader")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header_fallback"
		if _, ok := i["dp-http-header-fallback"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpHttpHeaderFallback(i["dp-http-header-fallback"], d, pre_append)
			tmp["dp_http_header_fallback"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpHttpHeaderFallback")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header_status"
		if _, ok := i["dp-http-header-status"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpHttpHeaderStatus(i["dp-http-header-status"], d, pre_append)
			tmp["dp_http_header_status"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpHttpHeaderStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header_suppress"
		if _, ok := i["dp-http-header-suppress"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpHttpHeaderSuppress(i["dp-http-header-suppress"], d, pre_append)
			tmp["dp_http_header_suppress"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpHttpHeaderSuppress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_log_dyn_flags"
		if _, ok := i["dp-log-dyn_flags"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpLogDynFlags(i["dp-log-dyn_flags"], d, pre_append)
			tmp["dp_log_dyn_flags"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpLogDynFlags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_log_period"
		if _, ok := i["dp-log-period"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpLogPeriod(i["dp-log-period"], d, pre_append)
			tmp["dp_log_period"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_mem_percent"
		if _, ok := i["dp-mem-percent"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpMemPercent(i["dp-mem-percent"], d, pre_append)
			tmp["dp_mem_percent"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpMemPercent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_profile_attribute"
		if _, ok := i["dp-profile-attribute"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpProfileAttribute(i["dp-profile-attribute"], d, pre_append)
			tmp["dp_profile_attribute"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpProfileAttribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_profile_attribute_key"
		if _, ok := i["dp-profile-attribute-key"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpProfileAttributeKey(i["dp-profile-attribute-key"], d, pre_append)
			tmp["dp_profile_attribute_key"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpProfileAttributeKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_radius_response"
		if _, ok := i["dp-radius-response"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpRadiusResponse(i["dp-radius-response"], d, pre_append)
			tmp["dp_radius_response"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpRadiusResponse")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_radius_server_port"
		if _, ok := i["dp-radius-server-port"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpRadiusServerPort(i["dp-radius-server-port"], d, pre_append)
			tmp["dp_radius_server_port"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpRadiusServerPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_secret"
		if _, ok := i["dp-secret"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpSecret(i["dp-secret"], d, pre_append)
			tmp["dp_secret"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpSecret")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_validate_request_secret"
		if _, ok := i["dp-validate-request-secret"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDpValidateRequestSecret(i["dp-validate-request-secret"], d, pre_append)
			tmp["dp_validate_request_secret"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DpValidateRequestSecret")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dynamic_profile"
		if _, ok := i["dynamic-profile"]; ok {
			v := flattenObjectUserRadiusDynamicMappingDynamicProfile(i["dynamic-profile"], d, pre_append)
			tmp["dynamic_profile"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-DynamicProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endpoint_translation"
		if _, ok := i["endpoint-translation"]; ok {
			v := flattenObjectUserRadiusDynamicMappingEndpointTranslation(i["endpoint-translation"], d, pre_append)
			tmp["endpoint_translation"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-EndpointTranslation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_convert_hex"
		if _, ok := i["ep-carrier-endpoint-convert-hex"]; ok {
			v := flattenObjectUserRadiusDynamicMappingEpCarrierEndpointConvertHex(i["ep-carrier-endpoint-convert-hex"], d, pre_append)
			tmp["ep_carrier_endpoint_convert_hex"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-EpCarrierEndpointConvertHex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_header"
		if _, ok := i["ep-carrier-endpoint-header"]; ok {
			v := flattenObjectUserRadiusDynamicMappingEpCarrierEndpointHeader(i["ep-carrier-endpoint-header"], d, pre_append)
			tmp["ep_carrier_endpoint_header"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-EpCarrierEndpointHeader")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_header_suppress"
		if _, ok := i["ep-carrier-endpoint-header-suppress"]; ok {
			v := flattenObjectUserRadiusDynamicMappingEpCarrierEndpointHeaderSuppress(i["ep-carrier-endpoint-header-suppress"], d, pre_append)
			tmp["ep_carrier_endpoint_header_suppress"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-EpCarrierEndpointHeaderSuppress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix"
		if _, ok := i["ep-carrier-endpoint-prefix"]; ok {
			v := flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefix(i["ep-carrier-endpoint-prefix"], d, pre_append)
			tmp["ep_carrier_endpoint_prefix"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-EpCarrierEndpointPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix_range_max"
		if _, ok := i["ep-carrier-endpoint-prefix-range-max"]; ok {
			v := flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMax(i["ep-carrier-endpoint-prefix-range-max"], d, pre_append)
			tmp["ep_carrier_endpoint_prefix_range_max"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-EpCarrierEndpointPrefixRangeMax")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix_range_min"
		if _, ok := i["ep-carrier-endpoint-prefix-range-min"]; ok {
			v := flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMin(i["ep-carrier-endpoint-prefix-range-min"], d, pre_append)
			tmp["ep_carrier_endpoint_prefix_range_min"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-EpCarrierEndpointPrefixRangeMin")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix_string"
		if _, ok := i["ep-carrier-endpoint-prefix-string"]; ok {
			v := flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixString(i["ep-carrier-endpoint-prefix-string"], d, pre_append)
			tmp["ep_carrier_endpoint_prefix_string"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-EpCarrierEndpointPrefixString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_source"
		if _, ok := i["ep-carrier-endpoint-source"]; ok {
			v := flattenObjectUserRadiusDynamicMappingEpCarrierEndpointSource(i["ep-carrier-endpoint-source"], d, pre_append)
			tmp["ep_carrier_endpoint_source"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-EpCarrierEndpointSource")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_ip_header"
		if _, ok := i["ep-ip-header"]; ok {
			v := flattenObjectUserRadiusDynamicMappingEpIpHeader(i["ep-ip-header"], d, pre_append)
			tmp["ep_ip_header"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-EpIpHeader")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_ip_header_suppress"
		if _, ok := i["ep-ip-header-suppress"]; ok {
			v := flattenObjectUserRadiusDynamicMappingEpIpHeaderSuppress(i["ep-ip-header-suppress"], d, pre_append)
			tmp["ep_ip_header_suppress"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-EpIpHeaderSuppress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_missing_header_fallback"
		if _, ok := i["ep-missing-header-fallback"]; ok {
			v := flattenObjectUserRadiusDynamicMappingEpMissingHeaderFallback(i["ep-missing-header-fallback"], d, pre_append)
			tmp["ep_missing_header_fallback"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-EpMissingHeaderFallback")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_profile_query_type"
		if _, ok := i["ep-profile-query-type"]; ok {
			v := flattenObjectUserRadiusDynamicMappingEpProfileQueryType(i["ep-profile-query-type"], d, pre_append)
			tmp["ep_profile_query_type"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-EpProfileQueryType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_override_attr_type"
		if _, ok := i["group-override-attr-type"]; ok {
			v := flattenObjectUserRadiusDynamicMappingGroupOverrideAttrType(i["group-override-attr-type"], d, pre_append)
			tmp["group_override_attr_type"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-GroupOverrideAttrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3c_compatibility"
		if _, ok := i["h3c-compatibility"]; ok {
			v := flattenObjectUserRadiusDynamicMappingH3CCompatibility(i["h3c-compatibility"], d, pre_append)
			tmp["h3c_compatibility"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-H3CCompatibility")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenObjectUserRadiusDynamicMappingInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenObjectUserRadiusDynamicMappingInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nas_ip"
		if _, ok := i["nas-ip"]; ok {
			v := flattenObjectUserRadiusDynamicMappingNasIp(i["nas-ip"], d, pre_append)
			tmp["nas_ip"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-NasIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_encoding"
		if _, ok := i["password-encoding"]; ok {
			v := flattenObjectUserRadiusDynamicMappingPasswordEncoding(i["password-encoding"], d, pre_append)
			tmp["password_encoding"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-PasswordEncoding")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_renewal"
		if _, ok := i["password-renewal"]; ok {
			v := flattenObjectUserRadiusDynamicMappingPasswordRenewal(i["password-renewal"], d, pre_append)
			tmp["password_renewal"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-PasswordRenewal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_coa"
		if _, ok := i["radius-coa"]; ok {
			v := flattenObjectUserRadiusDynamicMappingRadiusCoa(i["radius-coa"], d, pre_append)
			tmp["radius_coa"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-RadiusCoa")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_port"
		if _, ok := i["radius-port"]; ok {
			v := flattenObjectUserRadiusDynamicMappingRadiusPort(i["radius-port"], d, pre_append)
			tmp["radius_port"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-RadiusPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso"
		if _, ok := i["rsso"]; ok {
			v := flattenObjectUserRadiusDynamicMappingRsso(i["rsso"], d, pre_append)
			tmp["rsso"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-Rsso")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_context_timeout"
		if _, ok := i["rsso-context-timeout"]; ok {
			v := flattenObjectUserRadiusDynamicMappingRssoContextTimeout(i["rsso-context-timeout"], d, pre_append)
			tmp["rsso_context_timeout"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-RssoContextTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_endpoint_attribute"
		if _, ok := i["rsso-endpoint-attribute"]; ok {
			v := flattenObjectUserRadiusDynamicMappingRssoEndpointAttribute(i["rsso-endpoint-attribute"], d, pre_append)
			tmp["rsso_endpoint_attribute"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-RssoEndpointAttribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_endpoint_block_attribute"
		if _, ok := i["rsso-endpoint-block-attribute"]; ok {
			v := flattenObjectUserRadiusDynamicMappingRssoEndpointBlockAttribute(i["rsso-endpoint-block-attribute"], d, pre_append)
			tmp["rsso_endpoint_block_attribute"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-RssoEndpointBlockAttribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_ep_one_ip_only"
		if _, ok := i["rsso-ep-one-ip-only"]; ok {
			v := flattenObjectUserRadiusDynamicMappingRssoEpOneIpOnly(i["rsso-ep-one-ip-only"], d, pre_append)
			tmp["rsso_ep_one_ip_only"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-RssoEpOneIpOnly")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_flush_ip_session"
		if _, ok := i["rsso-flush-ip-session"]; ok {
			v := flattenObjectUserRadiusDynamicMappingRssoFlushIpSession(i["rsso-flush-ip-session"], d, pre_append)
			tmp["rsso_flush_ip_session"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-RssoFlushIpSession")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_log_flags"
		if _, ok := i["rsso-log-flags"]; ok {
			v := flattenObjectUserRadiusDynamicMappingRssoLogFlags(i["rsso-log-flags"], d, pre_append)
			tmp["rsso_log_flags"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-RssoLogFlags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_log_period"
		if _, ok := i["rsso-log-period"]; ok {
			v := flattenObjectUserRadiusDynamicMappingRssoLogPeriod(i["rsso-log-period"], d, pre_append)
			tmp["rsso_log_period"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-RssoLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_radius_response"
		if _, ok := i["rsso-radius-response"]; ok {
			v := flattenObjectUserRadiusDynamicMappingRssoRadiusResponse(i["rsso-radius-response"], d, pre_append)
			tmp["rsso_radius_response"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-RssoRadiusResponse")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_radius_server_port"
		if _, ok := i["rsso-radius-server-port"]; ok {
			v := flattenObjectUserRadiusDynamicMappingRssoRadiusServerPort(i["rsso-radius-server-port"], d, pre_append)
			tmp["rsso_radius_server_port"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-RssoRadiusServerPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_secret"
		if _, ok := i["rsso-secret"]; ok {
			v := flattenObjectUserRadiusDynamicMappingRssoSecret(i["rsso-secret"], d, pre_append)
			tmp["rsso_secret"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-RssoSecret")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_validate_request_secret"
		if _, ok := i["rsso-validate-request-secret"]; ok {
			v := flattenObjectUserRadiusDynamicMappingRssoValidateRequestSecret(i["rsso-validate-request-secret"], d, pre_append)
			tmp["rsso_validate_request_secret"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-RssoValidateRequestSecret")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_secret"
		if _, ok := i["secondary-secret"]; ok {
			v := flattenObjectUserRadiusDynamicMappingSecondarySecret(i["secondary-secret"], d, pre_append)
			tmp["secondary_secret"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-SecondarySecret")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_server"
		if _, ok := i["secondary-server"]; ok {
			v := flattenObjectUserRadiusDynamicMappingSecondaryServer(i["secondary-server"], d, pre_append)
			tmp["secondary_server"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-SecondaryServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secret"
		if _, ok := i["secret"]; ok {
			v := flattenObjectUserRadiusDynamicMappingSecret(i["secret"], d, pre_append)
			tmp["secret"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-Secret")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenObjectUserRadiusDynamicMappingServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenObjectUserRadiusDynamicMappingSourceIp(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-SourceIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute"
		if _, ok := i["sso-attribute"]; ok {
			v := flattenObjectUserRadiusDynamicMappingSsoAttribute(i["sso-attribute"], d, pre_append)
			tmp["sso_attribute"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-SsoAttribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute_key"
		if _, ok := i["sso-attribute-key"]; ok {
			v := flattenObjectUserRadiusDynamicMappingSsoAttributeKey(i["sso-attribute-key"], d, pre_append)
			tmp["sso_attribute_key"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-SsoAttributeKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute_value_override"
		if _, ok := i["sso-attribute-value-override"]; ok {
			v := flattenObjectUserRadiusDynamicMappingSsoAttributeValueOverride(i["sso-attribute-value-override"], d, pre_append)
			tmp["sso_attribute_value_override"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-SsoAttributeValueOverride")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_controller_acct_fast_framedip_detect"
		if _, ok := i["switch-controller-acct-fast-framedip-detect"]; ok {
			v := flattenObjectUserRadiusDynamicMappingSwitchControllerAcctFastFramedipDetect(i["switch-controller-acct-fast-framedip-detect"], d, pre_append)
			tmp["switch_controller_acct_fast_framedip_detect"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-SwitchControllerAcctFastFramedipDetect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_controller_service_type"
		if _, ok := i["switch-controller-service-type"]; ok {
			v := flattenObjectUserRadiusDynamicMappingSwitchControllerServiceType(i["switch-controller-service-type"], d, pre_append)
			tmp["switch_controller_service_type"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-SwitchControllerServiceType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_secret"
		if _, ok := i["tertiary-secret"]; ok {
			v := flattenObjectUserRadiusDynamicMappingTertiarySecret(i["tertiary-secret"], d, pre_append)
			tmp["tertiary_secret"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-TertiarySecret")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_server"
		if _, ok := i["tertiary-server"]; ok {
			v := flattenObjectUserRadiusDynamicMappingTertiaryServer(i["tertiary-server"], d, pre_append)
			tmp["tertiary_server"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-TertiaryServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timeout"
		if _, ok := i["timeout"]; ok {
			v := flattenObjectUserRadiusDynamicMappingTimeout(i["timeout"], d, pre_append)
			tmp["timeout"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-Timeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_group_for_profile"
		if _, ok := i["use-group-for-profile"]; ok {
			v := flattenObjectUserRadiusDynamicMappingUseGroupForProfile(i["use-group-for-profile"], d, pre_append)
			tmp["use_group_for_profile"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-UseGroupForProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_management_vdom"
		if _, ok := i["use-management-vdom"]; ok {
			v := flattenObjectUserRadiusDynamicMappingUseManagementVdom(i["use-management-vdom"], d, pre_append)
			tmp["use_management_vdom"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-UseManagementVdom")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "username_case_sensitive"
		if _, ok := i["username-case-sensitive"]; ok {
			v := flattenObjectUserRadiusDynamicMappingUsernameCaseSensitive(i["username-case-sensitive"], d, pre_append)
			tmp["username_case_sensitive"] = fortiAPISubPartPatch(v, "ObjectUserRadius-DynamicMapping-UsernameCaseSensitive")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectUserRadiusDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserRadiusDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectUserRadiusDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-Scope-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectUserRadiusDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountingServer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserRadiusDynamicMappingAccountingServerId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServerInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServerInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServerPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secret"
		if _, ok := i["secret"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServerSecret(i["secret"], d, pre_append)
			tmp["secret"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-Secret")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServerServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServerSourceIp(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-SourceIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectUserRadiusDynamicMappingAccountingServerStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectUserRadiusDynamicMapping-AccountingServer-Status")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectUserRadiusDynamicMappingAccountingServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountingServerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountingServerInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "auto",
			1: "sdwan",
			2: "specify",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountingServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountingServerSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusDynamicMappingAccountingServerServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountingServerSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAccountingServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingAcctAllServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingAcctInterimInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingAllUsergroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "pap",
			1: "chap",
			3: "ms_chap",
			4: "ms_chap_v2",
			6: "auto",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingClass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusDynamicMappingDpCarrierEndpointAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:  "User-Name",
			2:  "User-Password",
			3:  "CHAP-Password",
			4:  "NAS-IP-Address",
			5:  "NAS-Port",
			6:  "Service-Type",
			7:  "Framed-Protocol",
			8:  "Framed-IP-Address",
			9:  "Framed-IP-Netmask",
			10: "Framed-Routing",
			11: "Filter-Id",
			12: "Framed-MTU",
			13: "Framed-Compression",
			14: "Login-IP-Host",
			15: "Login-Service",
			16: "Login-TCP-Port",
			18: "Reply-Message",
			19: "Callback-Number",
			20: "Callback-Id",
			22: "Framed-Route",
			23: "Framed-IPX-Network",
			24: "State",
			25: "Class",
			26: "Vendor-Specific",
			27: "Session-Timeout",
			28: "Idle-Timeout",
			29: "Termination-Action",
			30: "Called-Station-Id",
			31: "Calling-Station-Id",
			32: "NAS-Identifier",
			33: "Proxy-State",
			34: "Login-LAT-Service",
			35: "Login-LAT-Node",
			36: "Login-LAT-Group",
			37: "Framed-AppleTalk-Link",
			38: "Framed-AppleTalk-Network",
			39: "Framed-AppleTalk-Zone",
			40: "Acct-Status-Type",
			41: "Acct-Delay-Time",
			42: "Acct-Input-Octets",
			43: "Acct-Output-Octets",
			44: "Acct-Session-Id",
			45: "Acct-Authentic",
			46: "Acct-Session-Time",
			47: "Acct-Input-Packets",
			48: "Acct-Output-Packets",
			49: "Acct-Terminate-Cause",
			50: "Acct-Multi-Session-Id",
			51: "Acct-Link-Count",
			60: "CHAP-Challenge",
			61: "NAS-Port-Type",
			62: "Port-Limit",
			63: "Login-LAT-Port",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingDpCarrierEndpointBlockAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:  "User-Name",
			2:  "User-Password",
			3:  "CHAP-Password",
			4:  "NAS-IP-Address",
			5:  "NAS-Port",
			6:  "Service-Type",
			7:  "Framed-Protocol",
			8:  "Framed-IP-Address",
			9:  "Framed-IP-Netmask",
			10: "Framed-Routing",
			11: "Filter-Id",
			12: "Framed-MTU",
			13: "Framed-Compression",
			14: "Login-IP-Host",
			15: "Login-Service",
			16: "Login-TCP-Port",
			18: "Reply-Message",
			19: "Callback-Number",
			20: "Callback-Id",
			22: "Framed-Route",
			23: "Framed-IPX-Network",
			24: "State",
			25: "Class",
			26: "Vendor-Specific",
			27: "Session-Timeout",
			28: "Idle-Timeout",
			29: "Termination-Action",
			30: "Called-Station-Id",
			31: "Calling-Station-Id",
			32: "NAS-Identifier",
			33: "Proxy-State",
			34: "Login-LAT-Service",
			35: "Login-LAT-Node",
			36: "Login-LAT-Group",
			37: "Framed-AppleTalk-Link",
			38: "Framed-AppleTalk-Network",
			39: "Framed-AppleTalk-Zone",
			40: "Acct-Status-Type",
			41: "Acct-Delay-Time",
			42: "Acct-Input-Octets",
			43: "Acct-Output-Octets",
			44: "Acct-Session-Id",
			45: "Acct-Authentic",
			46: "Acct-Session-Time",
			47: "Acct-Input-Packets",
			48: "Acct-Output-Packets",
			49: "Acct-Terminate-Cause",
			50: "Acct-Multi-Session-Id",
			51: "Acct-Link-Count",
			60: "CHAP-Challenge",
			61: "NAS-Port-Type",
			62: "Port-Limit",
			63: "Login-LAT-Port",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingDpContextTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpFlushIpSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingDpHoldTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpHttpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpHttpHeaderFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "ip-header-address",
			1: "default-profile",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingDpHttpHeaderStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingDpHttpHeaderSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingDpLogDynFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:   "none",
			2:   "protocol-error",
			4:   "profile-missing",
			8:   "context-missing",
			16:  "accounting-stop-missed",
			32:  "accounting-event",
			64:  "radiusd-other",
			128: "endpoint-block",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingDpLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpMemPercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpProfileAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:  "User-Name",
			2:  "User-Password",
			3:  "CHAP-Password",
			4:  "NAS-IP-Address",
			5:  "NAS-Port",
			6:  "Service-Type",
			7:  "Framed-Protocol",
			8:  "Framed-IP-Address",
			9:  "Framed-IP-Netmask",
			10: "Framed-Routing",
			11: "Filter-Id",
			12: "Framed-MTU",
			13: "Framed-Compression",
			14: "Login-IP-Host",
			15: "Login-Service",
			16: "Login-TCP-Port",
			18: "Reply-Message",
			19: "Callback-Number",
			20: "Callback-Id",
			22: "Framed-Route",
			23: "Framed-IPX-Network",
			24: "State",
			25: "Class",
			26: "Vendor-Specific",
			27: "Session-Timeout",
			28: "Idle-Timeout",
			29: "Termination-Action",
			30: "Called-Station-Id",
			31: "Calling-Station-Id",
			32: "NAS-Identifier",
			33: "Proxy-State",
			34: "Login-LAT-Service",
			35: "Login-LAT-Node",
			36: "Login-LAT-Group",
			37: "Framed-AppleTalk-Link",
			38: "Framed-AppleTalk-Network",
			39: "Framed-AppleTalk-Zone",
			40: "Acct-Status-Type",
			41: "Acct-Delay-Time",
			42: "Acct-Input-Octets",
			43: "Acct-Output-Octets",
			44: "Acct-Session-Id",
			45: "Acct-Authentic",
			46: "Acct-Session-Time",
			47: "Acct-Input-Packets",
			48: "Acct-Output-Packets",
			49: "Acct-Terminate-Cause",
			50: "Acct-Multi-Session-Id",
			51: "Acct-Link-Count",
			60: "CHAP-Challenge",
			61: "NAS-Port-Type",
			62: "Port-Limit",
			63: "Login-LAT-Port",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingDpProfileAttributeKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpRadiusResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingDpRadiusServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingDpSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusDynamicMappingDpValidateRequestSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingDynamicProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingEndpointTranslation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointConvertHex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointHeaderSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpCarrierEndpointSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "http-header",
			2: "cookie",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingEpIpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingEpIpHeaderSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingEpMissingHeaderFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "session-ip",
			1: "policy-profile",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingEpProfileQueryType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "session-ip",
			2: "extract-ip",
			3: "extract-carrier-endpoint",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingGroupOverrideAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "filter-Id",
			2: "class",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingH3CCompatibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "auto",
			1: "sdwan",
			2: "specify",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingNasIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingPasswordEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			65536:   "ISO-8859-1",
			2097152: "auto",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingPasswordRenewal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingRadiusCoa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingRadiusPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRsso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoContextTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoEndpointAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:  "User-Name",
			4:  "NAS-IP-Address",
			8:  "Framed-IP-Address",
			9:  "Framed-IP-Netmask",
			11: "Filter-Id",
			14: "Login-IP-Host",
			18: "Reply-Message",
			19: "Callback-Number",
			20: "Callback-Id",
			22: "Framed-Route",
			23: "Framed-IPX-Network",
			25: "Class",
			30: "Called-Station-Id",
			31: "Calling-Station-Id",
			32: "NAS-Identifier",
			33: "Proxy-State",
			34: "Login-LAT-Service",
			35: "Login-LAT-Node",
			36: "Login-LAT-Group",
			39: "Framed-AppleTalk-Zone",
			44: "Acct-Session-Id",
			50: "Acct-Multi-Session-Id",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoEndpointBlockAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:  "User-Name",
			4:  "NAS-IP-Address",
			8:  "Framed-IP-Address",
			9:  "Framed-IP-Netmask",
			11: "Filter-Id",
			14: "Login-IP-Host",
			18: "Reply-Message",
			19: "Callback-Number",
			20: "Callback-Id",
			22: "Framed-Route",
			23: "Framed-IPX-Network",
			25: "Class",
			30: "Called-Station-Id",
			31: "Calling-Station-Id",
			32: "NAS-Identifier",
			33: "Proxy-State",
			34: "Login-LAT-Service",
			35: "Login-LAT-Node",
			36: "Login-LAT-Group",
			39: "Framed-AppleTalk-Zone",
			44: "Acct-Session-Id",
			50: "Acct-Multi-Session-Id",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoEpOneIpOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoFlushIpSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoLogFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:   "none",
			2:   "protocol-error",
			4:   "profile-missing",
			16:  "accounting-stop-missed",
			32:  "accounting-event",
			64:  "radiusd-other",
			128: "endpoint-block",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoRadiusResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoRadiusServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingRssoSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusDynamicMappingRssoValidateRequestSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingSecondarySecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusDynamicMappingSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusDynamicMappingServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingSsoAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:  "User-Name",
			4:  "NAS-IP-Address",
			8:  "Framed-IP-Address",
			9:  "Framed-IP-Netmask",
			11: "Filter-Id",
			14: "Login-IP-Host",
			18: "Reply-Message",
			19: "Callback-Number",
			20: "Callback-Id",
			22: "Framed-Route",
			23: "Framed-IPX-Network",
			25: "Class",
			30: "Called-Station-Id",
			31: "Calling-Station-Id",
			32: "NAS-Identifier",
			33: "Proxy-State",
			34: "Login-LAT-Service",
			35: "Login-LAT-Node",
			36: "Login-LAT-Group",
			39: "Framed-AppleTalk-Zone",
			44: "Acct-Session-Id",
			50: "Acct-Multi-Session-Id",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingSsoAttributeKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingSsoAttributeValueOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingSwitchControllerAcctFastFramedipDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingSwitchControllerServiceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:    "login",
			2:    "framed",
			4:    "callback-login",
			8:    "callback-framed",
			16:   "outbound",
			32:   "administrative",
			64:   "nas-prompt",
			128:  "authenticate-only",
			256:  "callback-nas-prompt",
			512:  "call-check",
			1024: "callback-administrative",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingTertiarySecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusDynamicMappingTertiaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusDynamicMappingUseGroupForProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingUseManagementVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusDynamicMappingUsernameCaseSensitive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusGroupOverrideAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "filter-Id",
			2: "class",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusH3CCompatibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "auto",
			1: "sdwan",
			2: "specify",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusNasIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusPasswordEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			65536:   "ISO-8859-1",
			2097152: "auto",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusPasswordRenewal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusRadiusCoa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusRadiusPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusRsso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusRssoContextTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusRssoEndpointAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:  "User-Name",
			4:  "NAS-IP-Address",
			8:  "Framed-IP-Address",
			9:  "Framed-IP-Netmask",
			11: "Filter-Id",
			14: "Login-IP-Host",
			18: "Reply-Message",
			19: "Callback-Number",
			20: "Callback-Id",
			22: "Framed-Route",
			23: "Framed-IPX-Network",
			25: "Class",
			30: "Called-Station-Id",
			31: "Calling-Station-Id",
			32: "NAS-Identifier",
			33: "Proxy-State",
			34: "Login-LAT-Service",
			35: "Login-LAT-Node",
			36: "Login-LAT-Group",
			39: "Framed-AppleTalk-Zone",
			44: "Acct-Session-Id",
			50: "Acct-Multi-Session-Id",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusRssoEndpointBlockAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:  "User-Name",
			4:  "NAS-IP-Address",
			8:  "Framed-IP-Address",
			9:  "Framed-IP-Netmask",
			11: "Filter-Id",
			14: "Login-IP-Host",
			18: "Reply-Message",
			19: "Callback-Number",
			20: "Callback-Id",
			22: "Framed-Route",
			23: "Framed-IPX-Network",
			25: "Class",
			30: "Called-Station-Id",
			31: "Calling-Station-Id",
			32: "NAS-Identifier",
			33: "Proxy-State",
			34: "Login-LAT-Service",
			35: "Login-LAT-Node",
			36: "Login-LAT-Group",
			39: "Framed-AppleTalk-Zone",
			44: "Acct-Session-Id",
			50: "Acct-Multi-Session-Id",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusRssoEpOneIpOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusRssoFlushIpSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusRssoLogFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:   "none",
			2:   "protocol-error",
			4:   "profile-missing",
			16:  "accounting-stop-missed",
			32:  "accounting-event",
			64:  "radiusd-other",
			128: "endpoint-block",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusRssoLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusRssoRadiusResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusRssoRadiusServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusRssoSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusRssoValidateRequestSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusSecondarySecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusSsoAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:  "User-Name",
			4:  "NAS-IP-Address",
			8:  "Framed-IP-Address",
			9:  "Framed-IP-Netmask",
			11: "Filter-Id",
			14: "Login-IP-Host",
			18: "Reply-Message",
			19: "Callback-Number",
			20: "Callback-Id",
			22: "Framed-Route",
			23: "Framed-IPX-Network",
			25: "Class",
			30: "Called-Station-Id",
			31: "Calling-Station-Id",
			32: "NAS-Identifier",
			33: "Proxy-State",
			34: "Login-LAT-Service",
			35: "Login-LAT-Node",
			36: "Login-LAT-Group",
			39: "Framed-AppleTalk-Zone",
			44: "Acct-Session-Id",
			50: "Acct-Multi-Session-Id",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusSsoAttributeKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusSsoAttributeValueOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusSwitchControllerAcctFastFramedipDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusSwitchControllerServiceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:    "login",
			2:    "framed",
			4:    "callback-login",
			8:    "callback-framed",
			16:   "outbound",
			32:   "administrative",
			64:   "nas-prompt",
			128:  "authenticate-only",
			256:  "callback-nas-prompt",
			512:  "call-check",
			1024: "callback-administrative",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusTertiarySecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserRadiusTertiaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserRadiusUseManagementVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserRadiusUsernameCaseSensitive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectObjectUserRadius(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("accounting_server", flattenObjectUserRadiusAccountingServer(o["accounting-server"], d, "accounting_server")); err != nil {
			if vv, ok := fortiAPIPatch(o["accounting-server"], "ObjectUserRadius-AccountingServer"); ok {
				if err = d.Set("accounting_server", vv); err != nil {
					return fmt.Errorf("Error reading accounting_server: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading accounting_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("accounting_server"); ok {
			if err = d.Set("accounting_server", flattenObjectUserRadiusAccountingServer(o["accounting-server"], d, "accounting_server")); err != nil {
				if vv, ok := fortiAPIPatch(o["accounting-server"], "ObjectUserRadius-AccountingServer"); ok {
					if err = d.Set("accounting_server", vv); err != nil {
						return fmt.Errorf("Error reading accounting_server: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading accounting_server: %v", err)
				}
			}
		}
	}

	if err = d.Set("acct_all_servers", flattenObjectUserRadiusAcctAllServers(o["acct-all-servers"], d, "acct_all_servers")); err != nil {
		if vv, ok := fortiAPIPatch(o["acct-all-servers"], "ObjectUserRadius-AcctAllServers"); ok {
			if err = d.Set("acct_all_servers", vv); err != nil {
				return fmt.Errorf("Error reading acct_all_servers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acct_all_servers: %v", err)
		}
	}

	if err = d.Set("acct_interim_interval", flattenObjectUserRadiusAcctInterimInterval(o["acct-interim-interval"], d, "acct_interim_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["acct-interim-interval"], "ObjectUserRadius-AcctInterimInterval"); ok {
			if err = d.Set("acct_interim_interval", vv); err != nil {
				return fmt.Errorf("Error reading acct_interim_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acct_interim_interval: %v", err)
		}
	}

	if err = d.Set("all_usergroup", flattenObjectUserRadiusAllUsergroup(o["all-usergroup"], d, "all_usergroup")); err != nil {
		if vv, ok := fortiAPIPatch(o["all-usergroup"], "ObjectUserRadius-AllUsergroup"); ok {
			if err = d.Set("all_usergroup", vv); err != nil {
				return fmt.Errorf("Error reading all_usergroup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading all_usergroup: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenObjectUserRadiusAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "ObjectUserRadius-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("class", flattenObjectUserRadiusClass(o["class"], d, "class")); err != nil {
		if vv, ok := fortiAPIPatch(o["class"], "ObjectUserRadius-Class"); ok {
			if err = d.Set("class", vv); err != nil {
				return fmt.Errorf("Error reading class: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectUserRadiusDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectUserRadius-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectUserRadiusDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectUserRadius-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("group_override_attr_type", flattenObjectUserRadiusGroupOverrideAttrType(o["group-override-attr-type"], d, "group_override_attr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-override-attr-type"], "ObjectUserRadius-GroupOverrideAttrType"); ok {
			if err = d.Set("group_override_attr_type", vv); err != nil {
				return fmt.Errorf("Error reading group_override_attr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_override_attr_type: %v", err)
		}
	}

	if err = d.Set("h3c_compatibility", flattenObjectUserRadiusH3CCompatibility(o["h3c-compatibility"], d, "h3c_compatibility")); err != nil {
		if vv, ok := fortiAPIPatch(o["h3c-compatibility"], "ObjectUserRadius-H3CCompatibility"); ok {
			if err = d.Set("h3c_compatibility", vv); err != nil {
				return fmt.Errorf("Error reading h3c_compatibility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h3c_compatibility: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectUserRadiusInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectUserRadius-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenObjectUserRadiusInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "ObjectUserRadius-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserRadiusName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserRadius-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nas_ip", flattenObjectUserRadiusNasIp(o["nas-ip"], d, "nas_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["nas-ip"], "ObjectUserRadius-NasIp"); ok {
			if err = d.Set("nas_ip", vv); err != nil {
				return fmt.Errorf("Error reading nas_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nas_ip: %v", err)
		}
	}

	if err = d.Set("password_encoding", flattenObjectUserRadiusPasswordEncoding(o["password-encoding"], d, "password_encoding")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-encoding"], "ObjectUserRadius-PasswordEncoding"); ok {
			if err = d.Set("password_encoding", vv); err != nil {
				return fmt.Errorf("Error reading password_encoding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_encoding: %v", err)
		}
	}

	if err = d.Set("password_renewal", flattenObjectUserRadiusPasswordRenewal(o["password-renewal"], d, "password_renewal")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-renewal"], "ObjectUserRadius-PasswordRenewal"); ok {
			if err = d.Set("password_renewal", vv); err != nil {
				return fmt.Errorf("Error reading password_renewal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_renewal: %v", err)
		}
	}

	if err = d.Set("radius_coa", flattenObjectUserRadiusRadiusCoa(o["radius-coa"], d, "radius_coa")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-coa"], "ObjectUserRadius-RadiusCoa"); ok {
			if err = d.Set("radius_coa", vv); err != nil {
				return fmt.Errorf("Error reading radius_coa: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_coa: %v", err)
		}
	}

	if err = d.Set("radius_port", flattenObjectUserRadiusRadiusPort(o["radius-port"], d, "radius_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-port"], "ObjectUserRadius-RadiusPort"); ok {
			if err = d.Set("radius_port", vv); err != nil {
				return fmt.Errorf("Error reading radius_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_port: %v", err)
		}
	}

	if err = d.Set("rsso", flattenObjectUserRadiusRsso(o["rsso"], d, "rsso")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso"], "ObjectUserRadius-Rsso"); ok {
			if err = d.Set("rsso", vv); err != nil {
				return fmt.Errorf("Error reading rsso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso: %v", err)
		}
	}

	if err = d.Set("rsso_context_timeout", flattenObjectUserRadiusRssoContextTimeout(o["rsso-context-timeout"], d, "rsso_context_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-context-timeout"], "ObjectUserRadius-RssoContextTimeout"); ok {
			if err = d.Set("rsso_context_timeout", vv); err != nil {
				return fmt.Errorf("Error reading rsso_context_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_context_timeout: %v", err)
		}
	}

	if err = d.Set("rsso_endpoint_attribute", flattenObjectUserRadiusRssoEndpointAttribute(o["rsso-endpoint-attribute"], d, "rsso_endpoint_attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-endpoint-attribute"], "ObjectUserRadius-RssoEndpointAttribute"); ok {
			if err = d.Set("rsso_endpoint_attribute", vv); err != nil {
				return fmt.Errorf("Error reading rsso_endpoint_attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_endpoint_attribute: %v", err)
		}
	}

	if err = d.Set("rsso_endpoint_block_attribute", flattenObjectUserRadiusRssoEndpointBlockAttribute(o["rsso-endpoint-block-attribute"], d, "rsso_endpoint_block_attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-endpoint-block-attribute"], "ObjectUserRadius-RssoEndpointBlockAttribute"); ok {
			if err = d.Set("rsso_endpoint_block_attribute", vv); err != nil {
				return fmt.Errorf("Error reading rsso_endpoint_block_attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_endpoint_block_attribute: %v", err)
		}
	}

	if err = d.Set("rsso_ep_one_ip_only", flattenObjectUserRadiusRssoEpOneIpOnly(o["rsso-ep-one-ip-only"], d, "rsso_ep_one_ip_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-ep-one-ip-only"], "ObjectUserRadius-RssoEpOneIpOnly"); ok {
			if err = d.Set("rsso_ep_one_ip_only", vv); err != nil {
				return fmt.Errorf("Error reading rsso_ep_one_ip_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_ep_one_ip_only: %v", err)
		}
	}

	if err = d.Set("rsso_flush_ip_session", flattenObjectUserRadiusRssoFlushIpSession(o["rsso-flush-ip-session"], d, "rsso_flush_ip_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-flush-ip-session"], "ObjectUserRadius-RssoFlushIpSession"); ok {
			if err = d.Set("rsso_flush_ip_session", vv); err != nil {
				return fmt.Errorf("Error reading rsso_flush_ip_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_flush_ip_session: %v", err)
		}
	}

	if err = d.Set("rsso_log_flags", flattenObjectUserRadiusRssoLogFlags(o["rsso-log-flags"], d, "rsso_log_flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-log-flags"], "ObjectUserRadius-RssoLogFlags"); ok {
			if err = d.Set("rsso_log_flags", vv); err != nil {
				return fmt.Errorf("Error reading rsso_log_flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_log_flags: %v", err)
		}
	}

	if err = d.Set("rsso_log_period", flattenObjectUserRadiusRssoLogPeriod(o["rsso-log-period"], d, "rsso_log_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-log-period"], "ObjectUserRadius-RssoLogPeriod"); ok {
			if err = d.Set("rsso_log_period", vv); err != nil {
				return fmt.Errorf("Error reading rsso_log_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_log_period: %v", err)
		}
	}

	if err = d.Set("rsso_radius_response", flattenObjectUserRadiusRssoRadiusResponse(o["rsso-radius-response"], d, "rsso_radius_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-radius-response"], "ObjectUserRadius-RssoRadiusResponse"); ok {
			if err = d.Set("rsso_radius_response", vv); err != nil {
				return fmt.Errorf("Error reading rsso_radius_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_radius_response: %v", err)
		}
	}

	if err = d.Set("rsso_radius_server_port", flattenObjectUserRadiusRssoRadiusServerPort(o["rsso-radius-server-port"], d, "rsso_radius_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-radius-server-port"], "ObjectUserRadius-RssoRadiusServerPort"); ok {
			if err = d.Set("rsso_radius_server_port", vv); err != nil {
				return fmt.Errorf("Error reading rsso_radius_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_radius_server_port: %v", err)
		}
	}

	if err = d.Set("rsso_secret", flattenObjectUserRadiusRssoSecret(o["rsso-secret"], d, "rsso_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-secret"], "ObjectUserRadius-RssoSecret"); ok {
			if err = d.Set("rsso_secret", vv); err != nil {
				return fmt.Errorf("Error reading rsso_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_secret: %v", err)
		}
	}

	if err = d.Set("rsso_validate_request_secret", flattenObjectUserRadiusRssoValidateRequestSecret(o["rsso-validate-request-secret"], d, "rsso_validate_request_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-validate-request-secret"], "ObjectUserRadius-RssoValidateRequestSecret"); ok {
			if err = d.Set("rsso_validate_request_secret", vv); err != nil {
				return fmt.Errorf("Error reading rsso_validate_request_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_validate_request_secret: %v", err)
		}
	}

	if err = d.Set("secondary_secret", flattenObjectUserRadiusSecondarySecret(o["secondary-secret"], d, "secondary_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-secret"], "ObjectUserRadius-SecondarySecret"); ok {
			if err = d.Set("secondary_secret", vv); err != nil {
				return fmt.Errorf("Error reading secondary_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_secret: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenObjectUserRadiusSecondaryServer(o["secondary-server"], d, "secondary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-server"], "ObjectUserRadius-SecondaryServer"); ok {
			if err = d.Set("secondary_server", vv); err != nil {
				return fmt.Errorf("Error reading secondary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("secret", flattenObjectUserRadiusSecret(o["secret"], d, "secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["secret"], "ObjectUserRadius-Secret"); ok {
			if err = d.Set("secret", vv); err != nil {
				return fmt.Errorf("Error reading secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secret: %v", err)
		}
	}

	if err = d.Set("server", flattenObjectUserRadiusServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "ObjectUserRadius-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenObjectUserRadiusSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "ObjectUserRadius-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("sso_attribute", flattenObjectUserRadiusSsoAttribute(o["sso-attribute"], d, "sso_attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-attribute"], "ObjectUserRadius-SsoAttribute"); ok {
			if err = d.Set("sso_attribute", vv); err != nil {
				return fmt.Errorf("Error reading sso_attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_attribute: %v", err)
		}
	}

	if err = d.Set("sso_attribute_key", flattenObjectUserRadiusSsoAttributeKey(o["sso-attribute-key"], d, "sso_attribute_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-attribute-key"], "ObjectUserRadius-SsoAttributeKey"); ok {
			if err = d.Set("sso_attribute_key", vv); err != nil {
				return fmt.Errorf("Error reading sso_attribute_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_attribute_key: %v", err)
		}
	}

	if err = d.Set("sso_attribute_value_override", flattenObjectUserRadiusSsoAttributeValueOverride(o["sso-attribute-value-override"], d, "sso_attribute_value_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-attribute-value-override"], "ObjectUserRadius-SsoAttributeValueOverride"); ok {
			if err = d.Set("sso_attribute_value_override", vv); err != nil {
				return fmt.Errorf("Error reading sso_attribute_value_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_attribute_value_override: %v", err)
		}
	}

	if err = d.Set("switch_controller_acct_fast_framedip_detect", flattenObjectUserRadiusSwitchControllerAcctFastFramedipDetect(o["switch-controller-acct-fast-framedip-detect"], d, "switch_controller_acct_fast_framedip_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-acct-fast-framedip-detect"], "ObjectUserRadius-SwitchControllerAcctFastFramedipDetect"); ok {
			if err = d.Set("switch_controller_acct_fast_framedip_detect", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_acct_fast_framedip_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_acct_fast_framedip_detect: %v", err)
		}
	}

	if err = d.Set("switch_controller_service_type", flattenObjectUserRadiusSwitchControllerServiceType(o["switch-controller-service-type"], d, "switch_controller_service_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-service-type"], "ObjectUserRadius-SwitchControllerServiceType"); ok {
			if err = d.Set("switch_controller_service_type", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_service_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_service_type: %v", err)
		}
	}

	if err = d.Set("tertiary_secret", flattenObjectUserRadiusTertiarySecret(o["tertiary-secret"], d, "tertiary_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["tertiary-secret"], "ObjectUserRadius-TertiarySecret"); ok {
			if err = d.Set("tertiary_secret", vv); err != nil {
				return fmt.Errorf("Error reading tertiary_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tertiary_secret: %v", err)
		}
	}

	if err = d.Set("tertiary_server", flattenObjectUserRadiusTertiaryServer(o["tertiary-server"], d, "tertiary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tertiary-server"], "ObjectUserRadius-TertiaryServer"); ok {
			if err = d.Set("tertiary_server", vv); err != nil {
				return fmt.Errorf("Error reading tertiary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tertiary_server: %v", err)
		}
	}

	if err = d.Set("timeout", flattenObjectUserRadiusTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "ObjectUserRadius-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("use_management_vdom", flattenObjectUserRadiusUseManagementVdom(o["use-management-vdom"], d, "use_management_vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-management-vdom"], "ObjectUserRadius-UseManagementVdom"); ok {
			if err = d.Set("use_management_vdom", vv); err != nil {
				return fmt.Errorf("Error reading use_management_vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_management_vdom: %v", err)
		}
	}

	if err = d.Set("username_case_sensitive", flattenObjectUserRadiusUsernameCaseSensitive(o["username-case-sensitive"], d, "username_case_sensitive")); err != nil {
		if vv, ok := fortiAPIPatch(o["username-case-sensitive"], "ObjectUserRadius-UsernameCaseSensitive"); ok {
			if err = d.Set("username_case_sensitive", vv); err != nil {
				return fmt.Errorf("Error reading username_case_sensitive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username_case_sensitive: %v", err)
		}
	}

	return nil
}

func flattenObjectUserRadiusFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserRadiusAccountingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectUserRadiusAccountingServerId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandObjectUserRadiusAccountingServerInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface-select-method"], _ = expandObjectUserRadiusAccountingServerInterfaceSelectMethod(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandObjectUserRadiusAccountingServerPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secret"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["secret"], _ = expandObjectUserRadiusAccountingServerSecret(d, i["secret"], pre_append)
		} else {
			tmp["secret"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["server"], _ = expandObjectUserRadiusAccountingServerServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source-ip"], _ = expandObjectUserRadiusAccountingServerSourceIp(d, i["source_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandObjectUserRadiusAccountingServerStatus(d, i["status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectUserRadiusAccountingServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusAccountingServerInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusAccountingServerInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusAccountingServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusAccountingServerSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusAccountingServerServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusAccountingServerSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusAccountingServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusAcctAllServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusAcctInterimInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusAllUsergroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusClass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["_scope"], _ = expandObjectUserRadiusDynamicMappingScope(d, i["_scope"], pre_append)
		} else {
			tmp["_scope"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accounting_server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["accounting-server"], _ = expandObjectUserRadiusDynamicMappingAccountingServer(d, i["accounting_server"], pre_append)
		} else {
			tmp["accounting-server"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acct_all_servers"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["acct-all-servers"], _ = expandObjectUserRadiusDynamicMappingAcctAllServers(d, i["acct_all_servers"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acct_interim_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["acct-interim-interval"], _ = expandObjectUserRadiusDynamicMappingAcctInterimInterval(d, i["acct_interim_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "all_usergroup"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["all-usergroup"], _ = expandObjectUserRadiusDynamicMappingAllUsergroup(d, i["all_usergroup"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["auth-type"], _ = expandObjectUserRadiusDynamicMappingAuthType(d, i["auth_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["class"], _ = expandObjectUserRadiusDynamicMappingClass(d, i["class"], pre_append)
		} else {
			tmp["class"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_carrier_endpoint_attribute"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-carrier-endpoint-attribute"], _ = expandObjectUserRadiusDynamicMappingDpCarrierEndpointAttribute(d, i["dp_carrier_endpoint_attribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_carrier_endpoint_block_attribute"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-carrier-endpoint-block-attribute"], _ = expandObjectUserRadiusDynamicMappingDpCarrierEndpointBlockAttribute(d, i["dp_carrier_endpoint_block_attribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_context_timeout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-context-timeout"], _ = expandObjectUserRadiusDynamicMappingDpContextTimeout(d, i["dp_context_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_flush_ip_session"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-flush-ip-session"], _ = expandObjectUserRadiusDynamicMappingDpFlushIpSession(d, i["dp_flush_ip_session"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_hold_time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-hold-time"], _ = expandObjectUserRadiusDynamicMappingDpHoldTime(d, i["dp_hold_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-http-header"], _ = expandObjectUserRadiusDynamicMappingDpHttpHeader(d, i["dp_http_header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header_fallback"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-http-header-fallback"], _ = expandObjectUserRadiusDynamicMappingDpHttpHeaderFallback(d, i["dp_http_header_fallback"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header_status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-http-header-status"], _ = expandObjectUserRadiusDynamicMappingDpHttpHeaderStatus(d, i["dp_http_header_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header_suppress"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-http-header-suppress"], _ = expandObjectUserRadiusDynamicMappingDpHttpHeaderSuppress(d, i["dp_http_header_suppress"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_log_dyn_flags"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-log-dyn_flags"], _ = expandObjectUserRadiusDynamicMappingDpLogDynFlags(d, i["dp_log_dyn_flags"], pre_append)
		} else {
			tmp["dp-log-dyn_flags"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_log_period"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-log-period"], _ = expandObjectUserRadiusDynamicMappingDpLogPeriod(d, i["dp_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_mem_percent"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-mem-percent"], _ = expandObjectUserRadiusDynamicMappingDpMemPercent(d, i["dp_mem_percent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_profile_attribute"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-profile-attribute"], _ = expandObjectUserRadiusDynamicMappingDpProfileAttribute(d, i["dp_profile_attribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_profile_attribute_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-profile-attribute-key"], _ = expandObjectUserRadiusDynamicMappingDpProfileAttributeKey(d, i["dp_profile_attribute_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_radius_response"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-radius-response"], _ = expandObjectUserRadiusDynamicMappingDpRadiusResponse(d, i["dp_radius_response"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_radius_server_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-radius-server-port"], _ = expandObjectUserRadiusDynamicMappingDpRadiusServerPort(d, i["dp_radius_server_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_secret"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-secret"], _ = expandObjectUserRadiusDynamicMappingDpSecret(d, i["dp_secret"], pre_append)
		} else {
			tmp["dp-secret"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_validate_request_secret"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dp-validate-request-secret"], _ = expandObjectUserRadiusDynamicMappingDpValidateRequestSecret(d, i["dp_validate_request_secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dynamic_profile"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dynamic-profile"], _ = expandObjectUserRadiusDynamicMappingDynamicProfile(d, i["dynamic_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endpoint_translation"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["endpoint-translation"], _ = expandObjectUserRadiusDynamicMappingEndpointTranslation(d, i["endpoint_translation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_convert_hex"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ep-carrier-endpoint-convert-hex"], _ = expandObjectUserRadiusDynamicMappingEpCarrierEndpointConvertHex(d, i["ep_carrier_endpoint_convert_hex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ep-carrier-endpoint-header"], _ = expandObjectUserRadiusDynamicMappingEpCarrierEndpointHeader(d, i["ep_carrier_endpoint_header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_header_suppress"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ep-carrier-endpoint-header-suppress"], _ = expandObjectUserRadiusDynamicMappingEpCarrierEndpointHeaderSuppress(d, i["ep_carrier_endpoint_header_suppress"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ep-carrier-endpoint-prefix"], _ = expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefix(d, i["ep_carrier_endpoint_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix_range_max"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ep-carrier-endpoint-prefix-range-max"], _ = expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMax(d, i["ep_carrier_endpoint_prefix_range_max"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix_range_min"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ep-carrier-endpoint-prefix-range-min"], _ = expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMin(d, i["ep_carrier_endpoint_prefix_range_min"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ep-carrier-endpoint-prefix-string"], _ = expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixString(d, i["ep_carrier_endpoint_prefix_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_source"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ep-carrier-endpoint-source"], _ = expandObjectUserRadiusDynamicMappingEpCarrierEndpointSource(d, i["ep_carrier_endpoint_source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_ip_header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ep-ip-header"], _ = expandObjectUserRadiusDynamicMappingEpIpHeader(d, i["ep_ip_header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_ip_header_suppress"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ep-ip-header-suppress"], _ = expandObjectUserRadiusDynamicMappingEpIpHeaderSuppress(d, i["ep_ip_header_suppress"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_missing_header_fallback"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ep-missing-header-fallback"], _ = expandObjectUserRadiusDynamicMappingEpMissingHeaderFallback(d, i["ep_missing_header_fallback"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_profile_query_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ep-profile-query-type"], _ = expandObjectUserRadiusDynamicMappingEpProfileQueryType(d, i["ep_profile_query_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_override_attr_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["group-override-attr-type"], _ = expandObjectUserRadiusDynamicMappingGroupOverrideAttrType(d, i["group_override_attr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3c_compatibility"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["h3c-compatibility"], _ = expandObjectUserRadiusDynamicMappingH3CCompatibility(d, i["h3c_compatibility"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandObjectUserRadiusDynamicMappingInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface-select-method"], _ = expandObjectUserRadiusDynamicMappingInterfaceSelectMethod(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nas_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nas-ip"], _ = expandObjectUserRadiusDynamicMappingNasIp(d, i["nas_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_encoding"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["password-encoding"], _ = expandObjectUserRadiusDynamicMappingPasswordEncoding(d, i["password_encoding"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_renewal"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["password-renewal"], _ = expandObjectUserRadiusDynamicMappingPasswordRenewal(d, i["password_renewal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_coa"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["radius-coa"], _ = expandObjectUserRadiusDynamicMappingRadiusCoa(d, i["radius_coa"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["radius-port"], _ = expandObjectUserRadiusDynamicMappingRadiusPort(d, i["radius_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rsso"], _ = expandObjectUserRadiusDynamicMappingRsso(d, i["rsso"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_context_timeout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rsso-context-timeout"], _ = expandObjectUserRadiusDynamicMappingRssoContextTimeout(d, i["rsso_context_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_endpoint_attribute"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rsso-endpoint-attribute"], _ = expandObjectUserRadiusDynamicMappingRssoEndpointAttribute(d, i["rsso_endpoint_attribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_endpoint_block_attribute"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rsso-endpoint-block-attribute"], _ = expandObjectUserRadiusDynamicMappingRssoEndpointBlockAttribute(d, i["rsso_endpoint_block_attribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_ep_one_ip_only"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rsso-ep-one-ip-only"], _ = expandObjectUserRadiusDynamicMappingRssoEpOneIpOnly(d, i["rsso_ep_one_ip_only"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_flush_ip_session"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rsso-flush-ip-session"], _ = expandObjectUserRadiusDynamicMappingRssoFlushIpSession(d, i["rsso_flush_ip_session"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_log_flags"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rsso-log-flags"], _ = expandObjectUserRadiusDynamicMappingRssoLogFlags(d, i["rsso_log_flags"], pre_append)
		} else {
			tmp["rsso-log-flags"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_log_period"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rsso-log-period"], _ = expandObjectUserRadiusDynamicMappingRssoLogPeriod(d, i["rsso_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_radius_response"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rsso-radius-response"], _ = expandObjectUserRadiusDynamicMappingRssoRadiusResponse(d, i["rsso_radius_response"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_radius_server_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rsso-radius-server-port"], _ = expandObjectUserRadiusDynamicMappingRssoRadiusServerPort(d, i["rsso_radius_server_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_secret"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rsso-secret"], _ = expandObjectUserRadiusDynamicMappingRssoSecret(d, i["rsso_secret"], pre_append)
		} else {
			tmp["rsso-secret"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_validate_request_secret"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rsso-validate-request-secret"], _ = expandObjectUserRadiusDynamicMappingRssoValidateRequestSecret(d, i["rsso_validate_request_secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_secret"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["secondary-secret"], _ = expandObjectUserRadiusDynamicMappingSecondarySecret(d, i["secondary_secret"], pre_append)
		} else {
			tmp["secondary-secret"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["secondary-server"], _ = expandObjectUserRadiusDynamicMappingSecondaryServer(d, i["secondary_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secret"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["secret"], _ = expandObjectUserRadiusDynamicMappingSecret(d, i["secret"], pre_append)
		} else {
			tmp["secret"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["server"], _ = expandObjectUserRadiusDynamicMappingServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source-ip"], _ = expandObjectUserRadiusDynamicMappingSourceIp(d, i["source_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso-attribute"], _ = expandObjectUserRadiusDynamicMappingSsoAttribute(d, i["sso_attribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso-attribute-key"], _ = expandObjectUserRadiusDynamicMappingSsoAttributeKey(d, i["sso_attribute_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute_value_override"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso-attribute-value-override"], _ = expandObjectUserRadiusDynamicMappingSsoAttributeValueOverride(d, i["sso_attribute_value_override"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_controller_acct_fast_framedip_detect"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["switch-controller-acct-fast-framedip-detect"], _ = expandObjectUserRadiusDynamicMappingSwitchControllerAcctFastFramedipDetect(d, i["switch_controller_acct_fast_framedip_detect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_controller_service_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["switch-controller-service-type"], _ = expandObjectUserRadiusDynamicMappingSwitchControllerServiceType(d, i["switch_controller_service_type"], pre_append)
		} else {
			tmp["switch-controller-service-type"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_secret"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tertiary-secret"], _ = expandObjectUserRadiusDynamicMappingTertiarySecret(d, i["tertiary_secret"], pre_append)
		} else {
			tmp["tertiary-secret"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tertiary-server"], _ = expandObjectUserRadiusDynamicMappingTertiaryServer(d, i["tertiary_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timeout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["timeout"], _ = expandObjectUserRadiusDynamicMappingTimeout(d, i["timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_group_for_profile"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["use-group-for-profile"], _ = expandObjectUserRadiusDynamicMappingUseGroupForProfile(d, i["use_group_for_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_management_vdom"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["use-management-vdom"], _ = expandObjectUserRadiusDynamicMappingUseManagementVdom(d, i["use_management_vdom"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "username_case_sensitive"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["username-case-sensitive"], _ = expandObjectUserRadiusDynamicMappingUsernameCaseSensitive(d, i["username_case_sensitive"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectUserRadiusDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandObjectUserRadiusDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vdom"], _ = expandObjectUserRadiusDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectUserRadiusDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectUserRadiusDynamicMappingAccountingServerId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandObjectUserRadiusDynamicMappingAccountingServerInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface-select-method"], _ = expandObjectUserRadiusDynamicMappingAccountingServerInterfaceSelectMethod(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandObjectUserRadiusDynamicMappingAccountingServerPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secret"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["secret"], _ = expandObjectUserRadiusDynamicMappingAccountingServerSecret(d, i["secret"], pre_append)
		} else {
			tmp["secret"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["server"], _ = expandObjectUserRadiusDynamicMappingAccountingServerServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source-ip"], _ = expandObjectUserRadiusDynamicMappingAccountingServerSourceIp(d, i["source_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandObjectUserRadiusDynamicMappingAccountingServerStatus(d, i["status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAccountingServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAcctAllServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAcctInterimInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAllUsergroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingClass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingDpCarrierEndpointAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpCarrierEndpointBlockAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpContextTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpFlushIpSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpHoldTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpHttpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpHttpHeaderFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpHttpHeaderStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpHttpHeaderSuppress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpLogDynFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingDpLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpMemPercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpProfileAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpProfileAttributeKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpRadiusResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpRadiusServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDpSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingDpValidateRequestSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingDynamicProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEndpointTranslation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointConvertHex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointHeaderSuppress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointPrefixString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpCarrierEndpointSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpIpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpIpHeaderSuppress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpMissingHeaderFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingEpProfileQueryType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingGroupOverrideAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingH3CCompatibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingNasIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingPasswordEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingPasswordRenewal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRadiusCoa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRadiusPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRsso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoContextTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoEndpointAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoEndpointBlockAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoEpOneIpOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoFlushIpSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoLogFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingRssoLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoRadiusResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoRadiusServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingRssoSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingRssoValidateRequestSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSecondarySecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSsoAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSsoAttributeKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSsoAttributeValueOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSwitchControllerAcctFastFramedipDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingSwitchControllerServiceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingTertiarySecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusDynamicMappingTertiaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingUseGroupForProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingUseManagementVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusDynamicMappingUsernameCaseSensitive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusGroupOverrideAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusH3CCompatibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusNasIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusPasswordEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusPasswordRenewal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusRadiusCoa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusRadiusPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusRsso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusRssoContextTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusRssoEndpointAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusRssoEndpointBlockAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusRssoEpOneIpOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusRssoFlushIpSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusRssoLogFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusRssoLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusRssoRadiusResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusRssoRadiusServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusRssoSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusRssoValidateRequestSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusSecondarySecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusSsoAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusSsoAttributeKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusSsoAttributeValueOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusSwitchControllerAcctFastFramedipDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusSwitchControllerServiceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusTertiarySecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserRadiusTertiaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusUseManagementVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserRadiusUsernameCaseSensitive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserRadius(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accounting_server"); ok {
		t, err := expandObjectUserRadiusAccountingServer(d, v, "accounting_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accounting-server"] = t
		}
	}

	if v, ok := d.GetOk("acct_all_servers"); ok {
		t, err := expandObjectUserRadiusAcctAllServers(d, v, "acct_all_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-all-servers"] = t
		}
	}

	if v, ok := d.GetOk("acct_interim_interval"); ok {
		t, err := expandObjectUserRadiusAcctInterimInterval(d, v, "acct_interim_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-interim-interval"] = t
		}
	}

	if v, ok := d.GetOk("all_usergroup"); ok {
		t, err := expandObjectUserRadiusAllUsergroup(d, v, "all_usergroup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["all-usergroup"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok {
		t, err := expandObjectUserRadiusAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("class"); ok {
		t, err := expandObjectUserRadiusClass(d, v, "class")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok {
		t, err := expandObjectUserRadiusDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("group_override_attr_type"); ok {
		t, err := expandObjectUserRadiusGroupOverrideAttrType(d, v, "group_override_attr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-override-attr-type"] = t
		}
	}

	if v, ok := d.GetOk("h3c_compatibility"); ok {
		t, err := expandObjectUserRadiusH3CCompatibility(d, v, "h3c_compatibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3c-compatibility"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandObjectUserRadiusInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		t, err := expandObjectUserRadiusInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectUserRadiusName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nas_ip"); ok {
		t, err := expandObjectUserRadiusNasIp(d, v, "nas_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-ip"] = t
		}
	}

	if v, ok := d.GetOk("password_encoding"); ok {
		t, err := expandObjectUserRadiusPasswordEncoding(d, v, "password_encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-encoding"] = t
		}
	}

	if v, ok := d.GetOk("password_renewal"); ok {
		t, err := expandObjectUserRadiusPasswordRenewal(d, v, "password_renewal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-renewal"] = t
		}
	}

	if v, ok := d.GetOk("radius_coa"); ok {
		t, err := expandObjectUserRadiusRadiusCoa(d, v, "radius_coa")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-coa"] = t
		}
	}

	if v, ok := d.GetOk("radius_port"); ok {
		t, err := expandObjectUserRadiusRadiusPort(d, v, "radius_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-port"] = t
		}
	}

	if v, ok := d.GetOk("rsso"); ok {
		t, err := expandObjectUserRadiusRsso(d, v, "rsso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso"] = t
		}
	}

	if v, ok := d.GetOk("rsso_context_timeout"); ok {
		t, err := expandObjectUserRadiusRssoContextTimeout(d, v, "rsso_context_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-context-timeout"] = t
		}
	}

	if v, ok := d.GetOk("rsso_endpoint_attribute"); ok {
		t, err := expandObjectUserRadiusRssoEndpointAttribute(d, v, "rsso_endpoint_attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-endpoint-attribute"] = t
		}
	}

	if v, ok := d.GetOk("rsso_endpoint_block_attribute"); ok {
		t, err := expandObjectUserRadiusRssoEndpointBlockAttribute(d, v, "rsso_endpoint_block_attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-endpoint-block-attribute"] = t
		}
	}

	if v, ok := d.GetOk("rsso_ep_one_ip_only"); ok {
		t, err := expandObjectUserRadiusRssoEpOneIpOnly(d, v, "rsso_ep_one_ip_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-ep-one-ip-only"] = t
		}
	}

	if v, ok := d.GetOk("rsso_flush_ip_session"); ok {
		t, err := expandObjectUserRadiusRssoFlushIpSession(d, v, "rsso_flush_ip_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-flush-ip-session"] = t
		}
	}

	if v, ok := d.GetOk("rsso_log_flags"); ok {
		t, err := expandObjectUserRadiusRssoLogFlags(d, v, "rsso_log_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-log-flags"] = t
		}
	}

	if v, ok := d.GetOk("rsso_log_period"); ok {
		t, err := expandObjectUserRadiusRssoLogPeriod(d, v, "rsso_log_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-log-period"] = t
		}
	}

	if v, ok := d.GetOk("rsso_radius_response"); ok {
		t, err := expandObjectUserRadiusRssoRadiusResponse(d, v, "rsso_radius_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-radius-response"] = t
		}
	}

	if v, ok := d.GetOk("rsso_radius_server_port"); ok {
		t, err := expandObjectUserRadiusRssoRadiusServerPort(d, v, "rsso_radius_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-radius-server-port"] = t
		}
	}

	if v, ok := d.GetOk("rsso_secret"); ok {
		t, err := expandObjectUserRadiusRssoSecret(d, v, "rsso_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-secret"] = t
		}
	}

	if v, ok := d.GetOk("rsso_validate_request_secret"); ok {
		t, err := expandObjectUserRadiusRssoValidateRequestSecret(d, v, "rsso_validate_request_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-validate-request-secret"] = t
		}
	}

	if v, ok := d.GetOk("secondary_secret"); ok {
		t, err := expandObjectUserRadiusSecondarySecret(d, v, "secondary_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-secret"] = t
		}
	}

	if v, ok := d.GetOk("secondary_server"); ok {
		t, err := expandObjectUserRadiusSecondaryServer(d, v, "secondary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}

	if v, ok := d.GetOk("secret"); ok {
		t, err := expandObjectUserRadiusSecret(d, v, "secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandObjectUserRadiusServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandObjectUserRadiusSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute"); ok {
		t, err := expandObjectUserRadiusSsoAttribute(d, v, "sso_attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute_key"); ok {
		t, err := expandObjectUserRadiusSsoAttributeKey(d, v, "sso_attribute_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute-key"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute_value_override"); ok {
		t, err := expandObjectUserRadiusSsoAttributeValueOverride(d, v, "sso_attribute_value_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute-value-override"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_acct_fast_framedip_detect"); ok {
		t, err := expandObjectUserRadiusSwitchControllerAcctFastFramedipDetect(d, v, "switch_controller_acct_fast_framedip_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-acct-fast-framedip-detect"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_service_type"); ok {
		t, err := expandObjectUserRadiusSwitchControllerServiceType(d, v, "switch_controller_service_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-service-type"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_secret"); ok {
		t, err := expandObjectUserRadiusTertiarySecret(d, v, "tertiary_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-secret"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_server"); ok {
		t, err := expandObjectUserRadiusTertiaryServer(d, v, "tertiary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-server"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok {
		t, err := expandObjectUserRadiusTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("use_management_vdom"); ok {
		t, err := expandObjectUserRadiusUseManagementVdom(d, v, "use_management_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-management-vdom"] = t
		}
	}

	if v, ok := d.GetOk("username_case_sensitive"); ok {
		t, err := expandObjectUserRadiusUsernameCaseSensitive(d, v, "username_case_sensitive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-case-sensitive"] = t
		}
	}

	return &obj, nil
}
