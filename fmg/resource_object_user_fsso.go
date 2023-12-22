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

func resourceObjectUserFsso() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserFssoCreate,
		Read:   resourceObjectUserFssoRead,
		Update: resourceObjectUserFssoUpdate,
		Delete: resourceObjectUserFssoDelete,

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
			"_gui_meta": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
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

func resourceObjectUserFssoCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectUserFsso(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserFsso resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserFsso(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserFsso resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserFssoRead(d, m)
}

func resourceObjectUserFssoUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserFsso(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserFsso resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserFsso(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserFsso resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserFssoRead(d, m)
}

func resourceObjectUserFssoDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectUserFsso(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserFsso resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserFssoRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectUserFsso(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserFsso resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserFsso(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserFsso resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserFssoGuiMeta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_gui_meta"
		if _, ok := i["_gui_meta"]; ok {
			v := flattenObjectUserFssoDynamicMappingGuiMeta(i["_gui_meta"], d, pre_append)
			tmp["_gui_meta"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-GuiMeta")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenObjectUserFssoDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_poll_interval"
		if _, ok := i["group-poll-interval"]; ok {
			v := flattenObjectUserFssoDynamicMappingGroupPollInterval(i["group-poll-interval"], d, pre_append)
			tmp["group_poll_interval"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-GroupPollInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenObjectUserFssoDynamicMappingInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenObjectUserFssoDynamicMappingInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_poll"
		if _, ok := i["ldap-poll"]; ok {
			v := flattenObjectUserFssoDynamicMappingLdapPoll(i["ldap-poll"], d, pre_append)
			tmp["ldap_poll"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-LdapPoll")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_poll_filter"
		if _, ok := i["ldap-poll-filter"]; ok {
			v := flattenObjectUserFssoDynamicMappingLdapPollFilter(i["ldap-poll-filter"], d, pre_append)
			tmp["ldap_poll_filter"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-LdapPollFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_poll_interval"
		if _, ok := i["ldap-poll-interval"]; ok {
			v := flattenObjectUserFssoDynamicMappingLdapPollInterval(i["ldap-poll-interval"], d, pre_append)
			tmp["ldap_poll_interval"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-LdapPollInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_server"
		if _, ok := i["ldap-server"]; ok {
			v := flattenObjectUserFssoDynamicMappingLdapServer(i["ldap-server"], d, pre_append)
			tmp["ldap_server"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-LdapServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_timeout"
		if _, ok := i["logon-timeout"]; ok {
			v := flattenObjectUserFssoDynamicMappingLogonTimeout(i["logon-timeout"], d, pre_append)
			tmp["logon_timeout"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-LogonTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			v := flattenObjectUserFssoDynamicMappingPassword(i["password"], d, pre_append)
			tmp["password"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Password")
			c := d.Get(pre_append).(*schema.Set)
			if c.Len() > 0 {
				tmp["password"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password2"
		if _, ok := i["password2"]; ok {
			v := flattenObjectUserFssoDynamicMappingPassword2(i["password2"], d, pre_append)
			tmp["password2"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Password2")
			c := d.Get(pre_append).(*schema.Set)
			if c.Len() > 0 {
				tmp["password2"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password3"
		if _, ok := i["password3"]; ok {
			v := flattenObjectUserFssoDynamicMappingPassword3(i["password3"], d, pre_append)
			tmp["password3"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Password3")
			c := d.Get(pre_append).(*schema.Set)
			if c.Len() > 0 {
				tmp["password3"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password4"
		if _, ok := i["password4"]; ok {
			v := flattenObjectUserFssoDynamicMappingPassword4(i["password4"], d, pre_append)
			tmp["password4"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Password4")
			c := d.Get(pre_append).(*schema.Set)
			if c.Len() > 0 {
				tmp["password4"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password5"
		if _, ok := i["password5"]; ok {
			v := flattenObjectUserFssoDynamicMappingPassword5(i["password5"], d, pre_append)
			tmp["password5"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Password5")
			c := d.Get(pre_append).(*schema.Set)
			if c.Len() > 0 {
				tmp["password5"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectUserFssoDynamicMappingPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port2"
		if _, ok := i["port2"]; ok {
			v := flattenObjectUserFssoDynamicMappingPort2(i["port2"], d, pre_append)
			tmp["port2"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Port2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port3"
		if _, ok := i["port3"]; ok {
			v := flattenObjectUserFssoDynamicMappingPort3(i["port3"], d, pre_append)
			tmp["port3"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Port3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port4"
		if _, ok := i["port4"]; ok {
			v := flattenObjectUserFssoDynamicMappingPort4(i["port4"], d, pre_append)
			tmp["port4"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Port4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port5"
		if _, ok := i["port5"]; ok {
			v := flattenObjectUserFssoDynamicMappingPort5(i["port5"], d, pre_append)
			tmp["port5"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Port5")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenObjectUserFssoDynamicMappingServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server2"
		if _, ok := i["server2"]; ok {
			v := flattenObjectUserFssoDynamicMappingServer2(i["server2"], d, pre_append)
			tmp["server2"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Server2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server3"
		if _, ok := i["server3"]; ok {
			v := flattenObjectUserFssoDynamicMappingServer3(i["server3"], d, pre_append)
			tmp["server3"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Server3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server4"
		if _, ok := i["server4"]; ok {
			v := flattenObjectUserFssoDynamicMappingServer4(i["server4"], d, pre_append)
			tmp["server4"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Server4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server5"
		if _, ok := i["server5"]; ok {
			v := flattenObjectUserFssoDynamicMappingServer5(i["server5"], d, pre_append)
			tmp["server5"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Server5")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sni"
		if _, ok := i["sni"]; ok {
			v := flattenObjectUserFssoDynamicMappingSni(i["sni"], d, pre_append)
			tmp["sni"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Sni")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenObjectUserFssoDynamicMappingSourceIp(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-SourceIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip6"
		if _, ok := i["source-ip6"]; ok {
			v := flattenObjectUserFssoDynamicMappingSourceIp6(i["source-ip6"], d, pre_append)
			tmp["source_ip6"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-SourceIp6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl"
		if _, ok := i["ssl"]; ok {
			v := flattenObjectUserFssoDynamicMappingSsl(i["ssl"], d, pre_append)
			tmp["ssl"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Ssl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_host_ip_check"
		if _, ok := i["ssl-server-host-ip-check"]; ok {
			v := flattenObjectUserFssoDynamicMappingSslServerHostIpCheck(i["ssl-server-host-ip-check"], d, pre_append)
			tmp["ssl_server_host_ip_check"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-SslServerHostIpCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_trusted_cert"
		if _, ok := i["ssl-trusted-cert"]; ok {
			v := flattenObjectUserFssoDynamicMappingSslTrustedCert(i["ssl-trusted-cert"], d, pre_append)
			tmp["ssl_trusted_cert"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-SslTrustedCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectUserFssoDynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_info_server"
		if _, ok := i["user-info-server"]; ok {
			v := flattenObjectUserFssoDynamicMappingUserInfoServer(i["user-info-server"], d, pre_append)
			tmp["user_info_server"] = fortiAPISubPartPatch(v, "ObjectUserFsso-DynamicMapping-UserInfoServer")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserFssoDynamicMappingGuiMeta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserFssoDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserFssoDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectUserFssoDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectUserFssoDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserFssoDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingGroupPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingLdapPoll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingLdapPollFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingLdapPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingLogonTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserFssoDynamicMappingPassword2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserFssoDynamicMappingPassword3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserFssoDynamicMappingPassword4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserFssoDynamicMappingPassword5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserFssoDynamicMappingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingPort2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingPort3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingPort4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingPort5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingServer4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingServer5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingSni(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingSslServerHostIpCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingSslTrustedCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoDynamicMappingUserInfoServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoGroupPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoLdapPoll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoLdapPollFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoLdapPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoLogonTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserFssoPassword2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserFssoPassword3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserFssoPassword4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserFssoPassword5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserFssoPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPort2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPort3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPort4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPort5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoServer4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoServer5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoSni(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoSslServerHostIpCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoSslTrustedCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoUserInfoServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserFsso(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_gui_meta", flattenObjectUserFssoGuiMeta(o["_gui_meta"], d, "_gui_meta")); err != nil {
		if vv, ok := fortiAPIPatch(o["_gui_meta"], "ObjectUserFsso-GuiMeta"); ok {
			if err = d.Set("_gui_meta", vv); err != nil {
				return fmt.Errorf("Error reading _gui_meta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _gui_meta: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectUserFssoDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectUserFsso-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectUserFssoDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectUserFsso-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("group_poll_interval", flattenObjectUserFssoGroupPollInterval(o["group-poll-interval"], d, "group_poll_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-poll-interval"], "ObjectUserFsso-GroupPollInterval"); ok {
			if err = d.Set("group_poll_interval", vv); err != nil {
				return fmt.Errorf("Error reading group_poll_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_poll_interval: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectUserFssoInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectUserFsso-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenObjectUserFssoInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "ObjectUserFsso-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ldap_poll", flattenObjectUserFssoLdapPoll(o["ldap-poll"], d, "ldap_poll")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-poll"], "ObjectUserFsso-LdapPoll"); ok {
			if err = d.Set("ldap_poll", vv); err != nil {
				return fmt.Errorf("Error reading ldap_poll: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_poll: %v", err)
		}
	}

	if err = d.Set("ldap_poll_filter", flattenObjectUserFssoLdapPollFilter(o["ldap-poll-filter"], d, "ldap_poll_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-poll-filter"], "ObjectUserFsso-LdapPollFilter"); ok {
			if err = d.Set("ldap_poll_filter", vv); err != nil {
				return fmt.Errorf("Error reading ldap_poll_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_poll_filter: %v", err)
		}
	}

	if err = d.Set("ldap_poll_interval", flattenObjectUserFssoLdapPollInterval(o["ldap-poll-interval"], d, "ldap_poll_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-poll-interval"], "ObjectUserFsso-LdapPollInterval"); ok {
			if err = d.Set("ldap_poll_interval", vv); err != nil {
				return fmt.Errorf("Error reading ldap_poll_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_poll_interval: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenObjectUserFssoLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "ObjectUserFsso-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("logon_timeout", flattenObjectUserFssoLogonTimeout(o["logon-timeout"], d, "logon_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["logon-timeout"], "ObjectUserFsso-LogonTimeout"); ok {
			if err = d.Set("logon_timeout", vv); err != nil {
				return fmt.Errorf("Error reading logon_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logon_timeout: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserFssoName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserFsso-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectUserFssoPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectUserFsso-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("port2", flattenObjectUserFssoPort2(o["port2"], d, "port2")); err != nil {
		if vv, ok := fortiAPIPatch(o["port2"], "ObjectUserFsso-Port2"); ok {
			if err = d.Set("port2", vv); err != nil {
				return fmt.Errorf("Error reading port2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port2: %v", err)
		}
	}

	if err = d.Set("port3", flattenObjectUserFssoPort3(o["port3"], d, "port3")); err != nil {
		if vv, ok := fortiAPIPatch(o["port3"], "ObjectUserFsso-Port3"); ok {
			if err = d.Set("port3", vv); err != nil {
				return fmt.Errorf("Error reading port3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port3: %v", err)
		}
	}

	if err = d.Set("port4", flattenObjectUserFssoPort4(o["port4"], d, "port4")); err != nil {
		if vv, ok := fortiAPIPatch(o["port4"], "ObjectUserFsso-Port4"); ok {
			if err = d.Set("port4", vv); err != nil {
				return fmt.Errorf("Error reading port4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port4: %v", err)
		}
	}

	if err = d.Set("port5", flattenObjectUserFssoPort5(o["port5"], d, "port5")); err != nil {
		if vv, ok := fortiAPIPatch(o["port5"], "ObjectUserFsso-Port5"); ok {
			if err = d.Set("port5", vv); err != nil {
				return fmt.Errorf("Error reading port5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port5: %v", err)
		}
	}

	if err = d.Set("server", flattenObjectUserFssoServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "ObjectUserFsso-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server2", flattenObjectUserFssoServer2(o["server2"], d, "server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["server2"], "ObjectUserFsso-Server2"); ok {
			if err = d.Set("server2", vv); err != nil {
				return fmt.Errorf("Error reading server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server2: %v", err)
		}
	}

	if err = d.Set("server3", flattenObjectUserFssoServer3(o["server3"], d, "server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["server3"], "ObjectUserFsso-Server3"); ok {
			if err = d.Set("server3", vv); err != nil {
				return fmt.Errorf("Error reading server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server3: %v", err)
		}
	}

	if err = d.Set("server4", flattenObjectUserFssoServer4(o["server4"], d, "server4")); err != nil {
		if vv, ok := fortiAPIPatch(o["server4"], "ObjectUserFsso-Server4"); ok {
			if err = d.Set("server4", vv); err != nil {
				return fmt.Errorf("Error reading server4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server4: %v", err)
		}
	}

	if err = d.Set("server5", flattenObjectUserFssoServer5(o["server5"], d, "server5")); err != nil {
		if vv, ok := fortiAPIPatch(o["server5"], "ObjectUserFsso-Server5"); ok {
			if err = d.Set("server5", vv); err != nil {
				return fmt.Errorf("Error reading server5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server5: %v", err)
		}
	}

	if err = d.Set("sni", flattenObjectUserFssoSni(o["sni"], d, "sni")); err != nil {
		if vv, ok := fortiAPIPatch(o["sni"], "ObjectUserFsso-Sni"); ok {
			if err = d.Set("sni", vv); err != nil {
				return fmt.Errorf("Error reading sni: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sni: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenObjectUserFssoSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "ObjectUserFsso-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenObjectUserFssoSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip6"], "ObjectUserFsso-SourceIp6"); ok {
			if err = d.Set("source_ip6", vv); err != nil {
				return fmt.Errorf("Error reading source_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("ssl", flattenObjectUserFssoSsl(o["ssl"], d, "ssl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl"], "ObjectUserFsso-Ssl"); ok {
			if err = d.Set("ssl", vv); err != nil {
				return fmt.Errorf("Error reading ssl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl: %v", err)
		}
	}

	if err = d.Set("ssl_server_host_ip_check", flattenObjectUserFssoSslServerHostIpCheck(o["ssl-server-host-ip-check"], d, "ssl_server_host_ip_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-host-ip-check"], "ObjectUserFsso-SslServerHostIpCheck"); ok {
			if err = d.Set("ssl_server_host_ip_check", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_host_ip_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_host_ip_check: %v", err)
		}
	}

	if err = d.Set("ssl_trusted_cert", flattenObjectUserFssoSslTrustedCert(o["ssl-trusted-cert"], d, "ssl_trusted_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-trusted-cert"], "ObjectUserFsso-SslTrustedCert"); ok {
			if err = d.Set("ssl_trusted_cert", vv); err != nil {
				return fmt.Errorf("Error reading ssl_trusted_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_trusted_cert: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectUserFssoType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectUserFsso-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user_info_server", flattenObjectUserFssoUserInfoServer(o["user-info-server"], d, "user_info_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-info-server"], "ObjectUserFsso-UserInfoServer"); ok {
			if err = d.Set("user_info_server", vv); err != nil {
				return fmt.Errorf("Error reading user_info_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_info_server: %v", err)
		}
	}

	return nil
}

func flattenObjectUserFssoFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserFssoGuiMeta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_gui_meta"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_gui_meta"], _ = expandObjectUserFssoDynamicMappingGuiMeta(d, i["_gui_meta"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectUserFssoDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_poll_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-poll-interval"], _ = expandObjectUserFssoDynamicMappingGroupPollInterval(d, i["group_poll_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandObjectUserFssoDynamicMappingInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-select-method"], _ = expandObjectUserFssoDynamicMappingInterfaceSelectMethod(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_poll"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldap-poll"], _ = expandObjectUserFssoDynamicMappingLdapPoll(d, i["ldap_poll"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_poll_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldap-poll-filter"], _ = expandObjectUserFssoDynamicMappingLdapPollFilter(d, i["ldap_poll_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_poll_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldap-poll-interval"], _ = expandObjectUserFssoDynamicMappingLdapPollInterval(d, i["ldap_poll_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldap-server"], _ = expandObjectUserFssoDynamicMappingLdapServer(d, i["ldap_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logon-timeout"], _ = expandObjectUserFssoDynamicMappingLogonTimeout(d, i["logon_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandObjectUserFssoDynamicMappingPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password2"], _ = expandObjectUserFssoDynamicMappingPassword2(d, i["password2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password3"], _ = expandObjectUserFssoDynamicMappingPassword3(d, i["password3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password4"], _ = expandObjectUserFssoDynamicMappingPassword4(d, i["password4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password5"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password5"], _ = expandObjectUserFssoDynamicMappingPassword5(d, i["password5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectUserFssoDynamicMappingPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port2"], _ = expandObjectUserFssoDynamicMappingPort2(d, i["port2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port3"], _ = expandObjectUserFssoDynamicMappingPort3(d, i["port3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port4"], _ = expandObjectUserFssoDynamicMappingPort4(d, i["port4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port5"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port5"], _ = expandObjectUserFssoDynamicMappingPort5(d, i["port5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandObjectUserFssoDynamicMappingServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server2"], _ = expandObjectUserFssoDynamicMappingServer2(d, i["server2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server3"], _ = expandObjectUserFssoDynamicMappingServer3(d, i["server3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server4"], _ = expandObjectUserFssoDynamicMappingServer4(d, i["server4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server5"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server5"], _ = expandObjectUserFssoDynamicMappingServer5(d, i["server5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sni"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sni"], _ = expandObjectUserFssoDynamicMappingSni(d, i["sni"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip"], _ = expandObjectUserFssoDynamicMappingSourceIp(d, i["source_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip6"], _ = expandObjectUserFssoDynamicMappingSourceIp6(d, i["source_ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl"], _ = expandObjectUserFssoDynamicMappingSsl(d, i["ssl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_host_ip_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-host-ip-check"], _ = expandObjectUserFssoDynamicMappingSslServerHostIpCheck(d, i["ssl_server_host_ip_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_trusted_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-trusted-cert"], _ = expandObjectUserFssoDynamicMappingSslTrustedCert(d, i["ssl_trusted_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectUserFssoDynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_info_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-info-server"], _ = expandObjectUserFssoDynamicMappingUserInfoServer(d, i["user_info_server"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserFssoDynamicMappingGuiMeta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectUserFssoDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectUserFssoDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserFssoDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingGroupPollInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingLdapPoll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingLdapPollFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingLdapPollInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingLogonTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoDynamicMappingPassword2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoDynamicMappingPassword3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoDynamicMappingPassword4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoDynamicMappingPassword5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoDynamicMappingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingPort2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingPort3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingPort4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingPort5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingServer4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingServer5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingSni(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingSslServerHostIpCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingSslTrustedCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoDynamicMappingUserInfoServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoGroupPollInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoLdapPoll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoLdapPollFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoLdapPollInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoLogonTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoPassword2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoPassword3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoPassword4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoPassword5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPort2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPort3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPort4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPort5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoServer4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoServer5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoSni(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoSslServerHostIpCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoSslTrustedCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoUserInfoServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserFsso(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_gui_meta"); ok || d.HasChange("_gui_meta") {
		t, err := expandObjectUserFssoGuiMeta(d, v, "_gui_meta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_gui_meta"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectUserFssoDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("group_poll_interval"); ok || d.HasChange("group_poll_interval") {
		t, err := expandObjectUserFssoGroupPollInterval(d, v, "group_poll_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-poll-interval"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectUserFssoInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandObjectUserFssoInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ldap_poll"); ok || d.HasChange("ldap_poll") {
		t, err := expandObjectUserFssoLdapPoll(d, v, "ldap_poll")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-poll"] = t
		}
	}

	if v, ok := d.GetOk("ldap_poll_filter"); ok || d.HasChange("ldap_poll_filter") {
		t, err := expandObjectUserFssoLdapPollFilter(d, v, "ldap_poll_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-poll-filter"] = t
		}
	}

	if v, ok := d.GetOk("ldap_poll_interval"); ok || d.HasChange("ldap_poll_interval") {
		t, err := expandObjectUserFssoLdapPollInterval(d, v, "ldap_poll_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-poll-interval"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandObjectUserFssoLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("logon_timeout"); ok || d.HasChange("logon_timeout") {
		t, err := expandObjectUserFssoLogonTimeout(d, v, "logon_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-timeout"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserFssoName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectUserFssoPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("password2"); ok || d.HasChange("password2") {
		t, err := expandObjectUserFssoPassword2(d, v, "password2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password2"] = t
		}
	}

	if v, ok := d.GetOk("password3"); ok || d.HasChange("password3") {
		t, err := expandObjectUserFssoPassword3(d, v, "password3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password3"] = t
		}
	}

	if v, ok := d.GetOk("password4"); ok || d.HasChange("password4") {
		t, err := expandObjectUserFssoPassword4(d, v, "password4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password4"] = t
		}
	}

	if v, ok := d.GetOk("password5"); ok || d.HasChange("password5") {
		t, err := expandObjectUserFssoPassword5(d, v, "password5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password5"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectUserFssoPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("port2"); ok || d.HasChange("port2") {
		t, err := expandObjectUserFssoPort2(d, v, "port2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port2"] = t
		}
	}

	if v, ok := d.GetOk("port3"); ok || d.HasChange("port3") {
		t, err := expandObjectUserFssoPort3(d, v, "port3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port3"] = t
		}
	}

	if v, ok := d.GetOk("port4"); ok || d.HasChange("port4") {
		t, err := expandObjectUserFssoPort4(d, v, "port4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port4"] = t
		}
	}

	if v, ok := d.GetOk("port5"); ok || d.HasChange("port5") {
		t, err := expandObjectUserFssoPort5(d, v, "port5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port5"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandObjectUserFssoServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server2"); ok || d.HasChange("server2") {
		t, err := expandObjectUserFssoServer2(d, v, "server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server2"] = t
		}
	}

	if v, ok := d.GetOk("server3"); ok || d.HasChange("server3") {
		t, err := expandObjectUserFssoServer3(d, v, "server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server3"] = t
		}
	}

	if v, ok := d.GetOk("server4"); ok || d.HasChange("server4") {
		t, err := expandObjectUserFssoServer4(d, v, "server4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server4"] = t
		}
	}

	if v, ok := d.GetOk("server5"); ok || d.HasChange("server5") {
		t, err := expandObjectUserFssoServer5(d, v, "server5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server5"] = t
		}
	}

	if v, ok := d.GetOk("sni"); ok || d.HasChange("sni") {
		t, err := expandObjectUserFssoSni(d, v, "sni")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sni"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandObjectUserFssoSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok || d.HasChange("source_ip6") {
		t, err := expandObjectUserFssoSourceIp6(d, v, "source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("ssl"); ok || d.HasChange("ssl") {
		t, err := expandObjectUserFssoSsl(d, v, "ssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_host_ip_check"); ok || d.HasChange("ssl_server_host_ip_check") {
		t, err := expandObjectUserFssoSslServerHostIpCheck(d, v, "ssl_server_host_ip_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-host-ip-check"] = t
		}
	}

	if v, ok := d.GetOk("ssl_trusted_cert"); ok || d.HasChange("ssl_trusted_cert") {
		t, err := expandObjectUserFssoSslTrustedCert(d, v, "ssl_trusted_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-trusted-cert"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectUserFssoType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user_info_server"); ok || d.HasChange("user_info_server") {
		t, err := expandObjectUserFssoUserInfoServer(d, v, "user_info_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-info-server"] = t
		}
	}

	return &obj, nil
}
