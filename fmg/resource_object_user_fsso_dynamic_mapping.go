// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Fortinet Single Sign On (FSSO) agents.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserFssoDynamicMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserFssoDynamicMappingCreate,
		Read:   resourceObjectUserFssoDynamicMappingRead,
		Update: resourceObjectUserFssoDynamicMappingUpdate,
		Delete: resourceObjectUserFssoDynamicMappingDelete,

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
			"fsso": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"_gui_meta": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"group_poll_interval": &schema.Schema{
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
			"ldap_poll": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ldap_poll_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ldap_poll_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logon_timeout": &schema.Schema{
				Type:     schema.TypeInt,
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
			"password2": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"password3": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"password4": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"password5": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port2": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port3": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port4": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port5": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sni": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_host_ip_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_trusted_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_info_server": &schema.Schema{
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

func resourceObjectUserFssoDynamicMappingCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	fsso := d.Get("fsso").(string)
	paradict["fsso"] = fsso

	obj, err := getObjectObjectUserFssoDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserFssoDynamicMapping resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserFssoDynamicMapping(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserFssoDynamicMapping resource: %v", err)
	}

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectUserFssoDynamicMappingRead(d, m)
}

func resourceObjectUserFssoDynamicMappingUpdate(d *schema.ResourceData, m interface{}) error {
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

	fsso := d.Get("fsso").(string)
	paradict["fsso"] = fsso

	obj, err := getObjectObjectUserFssoDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserFssoDynamicMapping resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserFssoDynamicMapping(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserFssoDynamicMapping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectUserFssoDynamicMappingRead(d, m)
}

func resourceObjectUserFssoDynamicMappingDelete(d *schema.ResourceData, m interface{}) error {
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

	fsso := d.Get("fsso").(string)
	paradict["fsso"] = fsso

	err = c.DeleteObjectUserFssoDynamicMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserFssoDynamicMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserFssoDynamicMappingRead(d *schema.ResourceData, m interface{}) error {
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

	fsso := d.Get("fsso").(string)
	if fsso == "" {
		fsso = importOptionChecking(m.(*FortiClient).Cfg, "fsso")
		if fsso == "" {
			return fmt.Errorf("Parameter fsso is missing")
		}
		if err = d.Set("fsso", fsso); err != nil {
			return fmt.Errorf("Error set params fsso: %v", err)
		}
	}
	if mkey, err = checkScopeId(mkey); err != nil {
		return fmt.Errorf("Error set ID : %v", err)
	}
	paradict["fsso"] = fsso

	o, err := c.ReadObjectUserFssoDynamicMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserFssoDynamicMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserFssoDynamicMapping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserFssoDynamicMapping resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserFssoDynamicMappingGuiMeta2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingScope2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserFssoDynamicMappingScopeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserFssoDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectUserFssoDynamicMappingScopeVdom2edl(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectUserFssoDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserFssoDynamicMappingScopeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingScopeVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingGroupPollInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserFssoDynamicMappingInterfaceSelectMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingLdapPoll2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingLdapPollFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingLdapPollInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingLdapServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserFssoDynamicMappingLogonTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingPort22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingPort32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingPort42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingPort52edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingServer22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingServer32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingServer42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingServer52edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingSni2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingSourceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingSourceIp62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingSsl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingSslServerHostIpCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingSslTrustedCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserFssoDynamicMappingType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingUserInfoServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectObjectUserFssoDynamicMapping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_gui_meta", flattenObjectUserFssoDynamicMappingGuiMeta2edl(o["_gui_meta"], d, "_gui_meta")); err != nil {
		if vv, ok := fortiAPIPatch(o["_gui_meta"], "ObjectUserFssoDynamicMapping-GuiMeta"); ok {
			if err = d.Set("_gui_meta", vv); err != nil {
				return fmt.Errorf("Error reading _gui_meta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _gui_meta: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("_scope", flattenObjectUserFssoDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["_scope"], "ObjectUserFssoDynamicMapping-Scope"); ok {
				if err = d.Set("_scope", vv); err != nil {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading _scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("_scope"); ok {
			if err = d.Set("_scope", flattenObjectUserFssoDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["_scope"], "ObjectUserFssoDynamicMapping-Scope"); ok {
					if err = d.Set("_scope", vv); err != nil {
						return fmt.Errorf("Error reading _scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			}
		}
	}

	if err = d.Set("group_poll_interval", flattenObjectUserFssoDynamicMappingGroupPollInterval2edl(o["group-poll-interval"], d, "group_poll_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-poll-interval"], "ObjectUserFssoDynamicMapping-GroupPollInterval"); ok {
			if err = d.Set("group_poll_interval", vv); err != nil {
				return fmt.Errorf("Error reading group_poll_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_poll_interval: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectUserFssoDynamicMappingInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectUserFssoDynamicMapping-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenObjectUserFssoDynamicMappingInterfaceSelectMethod2edl(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "ObjectUserFssoDynamicMapping-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ldap_poll", flattenObjectUserFssoDynamicMappingLdapPoll2edl(o["ldap-poll"], d, "ldap_poll")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-poll"], "ObjectUserFssoDynamicMapping-LdapPoll"); ok {
			if err = d.Set("ldap_poll", vv); err != nil {
				return fmt.Errorf("Error reading ldap_poll: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_poll: %v", err)
		}
	}

	if err = d.Set("ldap_poll_filter", flattenObjectUserFssoDynamicMappingLdapPollFilter2edl(o["ldap-poll-filter"], d, "ldap_poll_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-poll-filter"], "ObjectUserFssoDynamicMapping-LdapPollFilter"); ok {
			if err = d.Set("ldap_poll_filter", vv); err != nil {
				return fmt.Errorf("Error reading ldap_poll_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_poll_filter: %v", err)
		}
	}

	if err = d.Set("ldap_poll_interval", flattenObjectUserFssoDynamicMappingLdapPollInterval2edl(o["ldap-poll-interval"], d, "ldap_poll_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-poll-interval"], "ObjectUserFssoDynamicMapping-LdapPollInterval"); ok {
			if err = d.Set("ldap_poll_interval", vv); err != nil {
				return fmt.Errorf("Error reading ldap_poll_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_poll_interval: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenObjectUserFssoDynamicMappingLdapServer2edl(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "ObjectUserFssoDynamicMapping-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("logon_timeout", flattenObjectUserFssoDynamicMappingLogonTimeout2edl(o["logon-timeout"], d, "logon_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["logon-timeout"], "ObjectUserFssoDynamicMapping-LogonTimeout"); ok {
			if err = d.Set("logon_timeout", vv); err != nil {
				return fmt.Errorf("Error reading logon_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logon_timeout: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectUserFssoDynamicMappingPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectUserFssoDynamicMapping-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("port2", flattenObjectUserFssoDynamicMappingPort22edl(o["port2"], d, "port2")); err != nil {
		if vv, ok := fortiAPIPatch(o["port2"], "ObjectUserFssoDynamicMapping-Port2"); ok {
			if err = d.Set("port2", vv); err != nil {
				return fmt.Errorf("Error reading port2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port2: %v", err)
		}
	}

	if err = d.Set("port3", flattenObjectUserFssoDynamicMappingPort32edl(o["port3"], d, "port3")); err != nil {
		if vv, ok := fortiAPIPatch(o["port3"], "ObjectUserFssoDynamicMapping-Port3"); ok {
			if err = d.Set("port3", vv); err != nil {
				return fmt.Errorf("Error reading port3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port3: %v", err)
		}
	}

	if err = d.Set("port4", flattenObjectUserFssoDynamicMappingPort42edl(o["port4"], d, "port4")); err != nil {
		if vv, ok := fortiAPIPatch(o["port4"], "ObjectUserFssoDynamicMapping-Port4"); ok {
			if err = d.Set("port4", vv); err != nil {
				return fmt.Errorf("Error reading port4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port4: %v", err)
		}
	}

	if err = d.Set("port5", flattenObjectUserFssoDynamicMappingPort52edl(o["port5"], d, "port5")); err != nil {
		if vv, ok := fortiAPIPatch(o["port5"], "ObjectUserFssoDynamicMapping-Port5"); ok {
			if err = d.Set("port5", vv); err != nil {
				return fmt.Errorf("Error reading port5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port5: %v", err)
		}
	}

	if err = d.Set("server", flattenObjectUserFssoDynamicMappingServer2edl(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "ObjectUserFssoDynamicMapping-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server2", flattenObjectUserFssoDynamicMappingServer22edl(o["server2"], d, "server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["server2"], "ObjectUserFssoDynamicMapping-Server2"); ok {
			if err = d.Set("server2", vv); err != nil {
				return fmt.Errorf("Error reading server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server2: %v", err)
		}
	}

	if err = d.Set("server3", flattenObjectUserFssoDynamicMappingServer32edl(o["server3"], d, "server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["server3"], "ObjectUserFssoDynamicMapping-Server3"); ok {
			if err = d.Set("server3", vv); err != nil {
				return fmt.Errorf("Error reading server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server3: %v", err)
		}
	}

	if err = d.Set("server4", flattenObjectUserFssoDynamicMappingServer42edl(o["server4"], d, "server4")); err != nil {
		if vv, ok := fortiAPIPatch(o["server4"], "ObjectUserFssoDynamicMapping-Server4"); ok {
			if err = d.Set("server4", vv); err != nil {
				return fmt.Errorf("Error reading server4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server4: %v", err)
		}
	}

	if err = d.Set("server5", flattenObjectUserFssoDynamicMappingServer52edl(o["server5"], d, "server5")); err != nil {
		if vv, ok := fortiAPIPatch(o["server5"], "ObjectUserFssoDynamicMapping-Server5"); ok {
			if err = d.Set("server5", vv); err != nil {
				return fmt.Errorf("Error reading server5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server5: %v", err)
		}
	}

	if err = d.Set("sni", flattenObjectUserFssoDynamicMappingSni2edl(o["sni"], d, "sni")); err != nil {
		if vv, ok := fortiAPIPatch(o["sni"], "ObjectUserFssoDynamicMapping-Sni"); ok {
			if err = d.Set("sni", vv); err != nil {
				return fmt.Errorf("Error reading sni: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sni: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenObjectUserFssoDynamicMappingSourceIp2edl(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "ObjectUserFssoDynamicMapping-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenObjectUserFssoDynamicMappingSourceIp62edl(o["source-ip6"], d, "source_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip6"], "ObjectUserFssoDynamicMapping-SourceIp6"); ok {
			if err = d.Set("source_ip6", vv); err != nil {
				return fmt.Errorf("Error reading source_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("ssl", flattenObjectUserFssoDynamicMappingSsl2edl(o["ssl"], d, "ssl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl"], "ObjectUserFssoDynamicMapping-Ssl"); ok {
			if err = d.Set("ssl", vv); err != nil {
				return fmt.Errorf("Error reading ssl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl: %v", err)
		}
	}

	if err = d.Set("ssl_server_host_ip_check", flattenObjectUserFssoDynamicMappingSslServerHostIpCheck2edl(o["ssl-server-host-ip-check"], d, "ssl_server_host_ip_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-host-ip-check"], "ObjectUserFssoDynamicMapping-SslServerHostIpCheck"); ok {
			if err = d.Set("ssl_server_host_ip_check", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_host_ip_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_host_ip_check: %v", err)
		}
	}

	if err = d.Set("ssl_trusted_cert", flattenObjectUserFssoDynamicMappingSslTrustedCert2edl(o["ssl-trusted-cert"], d, "ssl_trusted_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-trusted-cert"], "ObjectUserFssoDynamicMapping-SslTrustedCert"); ok {
			if err = d.Set("ssl_trusted_cert", vv); err != nil {
				return fmt.Errorf("Error reading ssl_trusted_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_trusted_cert: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectUserFssoDynamicMappingType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectUserFssoDynamicMapping-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user_info_server", flattenObjectUserFssoDynamicMappingUserInfoServer2edl(o["user-info-server"], d, "user_info_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-info-server"], "ObjectUserFssoDynamicMapping-UserInfoServer"); ok {
			if err = d.Set("user_info_server", vv); err != nil {
				return fmt.Errorf("Error reading user_info_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_info_server: %v", err)
		}
	}

	return nil
}

func flattenObjectUserFssoDynamicMappingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserFssoDynamicMappingGuiMeta2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectUserFssoDynamicMappingScopeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectUserFssoDynamicMappingScopeVdom2edl(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserFssoDynamicMappingScopeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingScopeVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingGroupPollInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserFssoDynamicMappingInterfaceSelectMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingLdapPoll2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingLdapPollFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingLdapPollInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingLdapServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserFssoDynamicMappingLogonTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoDynamicMappingPassword22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoDynamicMappingPassword32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoDynamicMappingPassword42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoDynamicMappingPassword52edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoDynamicMappingPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingPort22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingPort32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingPort42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingPort52edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingServer22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingServer32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingServer42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingServer52edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingSni2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingSourceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingSourceIp62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingSsl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingSslServerHostIpCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingSslTrustedCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserFssoDynamicMappingType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingUserInfoServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectObjectUserFssoDynamicMapping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_gui_meta"); ok || d.HasChange("_gui_meta") {
		t, err := expandObjectUserFssoDynamicMappingGuiMeta2edl(d, v, "_gui_meta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_gui_meta"] = t
		}
	}

	if v, ok := d.GetOk("_scope"); ok || d.HasChange("_scope") {
		t, err := expandObjectUserFssoDynamicMappingScope2edl(d, v, "_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_scope"] = t
		}
	}

	if v, ok := d.GetOk("group_poll_interval"); ok || d.HasChange("group_poll_interval") {
		t, err := expandObjectUserFssoDynamicMappingGroupPollInterval2edl(d, v, "group_poll_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-poll-interval"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectUserFssoDynamicMappingInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandObjectUserFssoDynamicMappingInterfaceSelectMethod2edl(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ldap_poll"); ok || d.HasChange("ldap_poll") {
		t, err := expandObjectUserFssoDynamicMappingLdapPoll2edl(d, v, "ldap_poll")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-poll"] = t
		}
	}

	if v, ok := d.GetOk("ldap_poll_filter"); ok || d.HasChange("ldap_poll_filter") {
		t, err := expandObjectUserFssoDynamicMappingLdapPollFilter2edl(d, v, "ldap_poll_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-poll-filter"] = t
		}
	}

	if v, ok := d.GetOk("ldap_poll_interval"); ok || d.HasChange("ldap_poll_interval") {
		t, err := expandObjectUserFssoDynamicMappingLdapPollInterval2edl(d, v, "ldap_poll_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-poll-interval"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandObjectUserFssoDynamicMappingLdapServer2edl(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("logon_timeout"); ok || d.HasChange("logon_timeout") {
		t, err := expandObjectUserFssoDynamicMappingLogonTimeout2edl(d, v, "logon_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-timeout"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectUserFssoDynamicMappingPassword2edl(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("password2"); ok || d.HasChange("password2") {
		t, err := expandObjectUserFssoDynamicMappingPassword22edl(d, v, "password2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password2"] = t
		}
	}

	if v, ok := d.GetOk("password3"); ok || d.HasChange("password3") {
		t, err := expandObjectUserFssoDynamicMappingPassword32edl(d, v, "password3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password3"] = t
		}
	}

	if v, ok := d.GetOk("password4"); ok || d.HasChange("password4") {
		t, err := expandObjectUserFssoDynamicMappingPassword42edl(d, v, "password4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password4"] = t
		}
	}

	if v, ok := d.GetOk("password5"); ok || d.HasChange("password5") {
		t, err := expandObjectUserFssoDynamicMappingPassword52edl(d, v, "password5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password5"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectUserFssoDynamicMappingPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("port2"); ok || d.HasChange("port2") {
		t, err := expandObjectUserFssoDynamicMappingPort22edl(d, v, "port2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port2"] = t
		}
	}

	if v, ok := d.GetOk("port3"); ok || d.HasChange("port3") {
		t, err := expandObjectUserFssoDynamicMappingPort32edl(d, v, "port3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port3"] = t
		}
	}

	if v, ok := d.GetOk("port4"); ok || d.HasChange("port4") {
		t, err := expandObjectUserFssoDynamicMappingPort42edl(d, v, "port4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port4"] = t
		}
	}

	if v, ok := d.GetOk("port5"); ok || d.HasChange("port5") {
		t, err := expandObjectUserFssoDynamicMappingPort52edl(d, v, "port5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port5"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandObjectUserFssoDynamicMappingServer2edl(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server2"); ok || d.HasChange("server2") {
		t, err := expandObjectUserFssoDynamicMappingServer22edl(d, v, "server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server2"] = t
		}
	}

	if v, ok := d.GetOk("server3"); ok || d.HasChange("server3") {
		t, err := expandObjectUserFssoDynamicMappingServer32edl(d, v, "server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server3"] = t
		}
	}

	if v, ok := d.GetOk("server4"); ok || d.HasChange("server4") {
		t, err := expandObjectUserFssoDynamicMappingServer42edl(d, v, "server4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server4"] = t
		}
	}

	if v, ok := d.GetOk("server5"); ok || d.HasChange("server5") {
		t, err := expandObjectUserFssoDynamicMappingServer52edl(d, v, "server5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server5"] = t
		}
	}

	if v, ok := d.GetOk("sni"); ok || d.HasChange("sni") {
		t, err := expandObjectUserFssoDynamicMappingSni2edl(d, v, "sni")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sni"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandObjectUserFssoDynamicMappingSourceIp2edl(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok || d.HasChange("source_ip6") {
		t, err := expandObjectUserFssoDynamicMappingSourceIp62edl(d, v, "source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("ssl"); ok || d.HasChange("ssl") {
		t, err := expandObjectUserFssoDynamicMappingSsl2edl(d, v, "ssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_host_ip_check"); ok || d.HasChange("ssl_server_host_ip_check") {
		t, err := expandObjectUserFssoDynamicMappingSslServerHostIpCheck2edl(d, v, "ssl_server_host_ip_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-host-ip-check"] = t
		}
	}

	if v, ok := d.GetOk("ssl_trusted_cert"); ok || d.HasChange("ssl_trusted_cert") {
		t, err := expandObjectUserFssoDynamicMappingSslTrustedCert2edl(d, v, "ssl_trusted_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-trusted-cert"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectUserFssoDynamicMappingType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user_info_server"); ok || d.HasChange("user_info_server") {
		t, err := expandObjectUserFssoDynamicMappingUserInfoServer2edl(d, v, "user_info_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-info-server"] = t
		}
	}

	return &obj, nil
}
