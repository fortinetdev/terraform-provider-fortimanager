// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure domain controller entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectUserDomainController() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserDomainControllerCreate,
		Read:   resourceObjectUserDomainControllerRead,
		Update: resourceObjectUserDomainControllerUpdate,
		Delete: resourceObjectUserDomainControllerDelete,

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
			"ad_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adlds_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adlds_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adlds_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adlds_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dns_srv_lookup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extra_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"source_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"hostname": &schema.Schema{
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
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"replication_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"source_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
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

func resourceObjectUserDomainControllerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserDomainController(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserDomainController resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserDomainController(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserDomainController resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserDomainControllerRead(d, m)
}

func resourceObjectUserDomainControllerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserDomainController(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserDomainController resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserDomainController(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserDomainController resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserDomainControllerRead(d, m)
}

func resourceObjectUserDomainControllerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectUserDomainController(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserDomainController resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserDomainControllerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectUserDomainController(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserDomainController resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserDomainController(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserDomainController resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserDomainControllerAdMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerAdldsDn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerAdldsIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerAdldsIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerAdldsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerDnsSrvLookup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerDomainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerExtraServer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserDomainControllerExtraServerId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectUserDomainController-ExtraServer-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := i["ip-address"]; ok {
			v := flattenObjectUserDomainControllerExtraServerIpAddress(i["ip-address"], d, pre_append)
			tmp["ip_address"] = fortiAPISubPartPatch(v, "ObjectUserDomainController-ExtraServer-IpAddress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectUserDomainControllerExtraServerPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectUserDomainController-ExtraServer-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip_address"
		if _, ok := i["source-ip-address"]; ok {
			v := flattenObjectUserDomainControllerExtraServerSourceIpAddress(i["source-ip-address"], d, pre_append)
			tmp["source_ip_address"] = fortiAPISubPartPatch(v, "ObjectUserDomainController-ExtraServer-SourceIpAddress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_port"
		if _, ok := i["source-port"]; ok {
			v := flattenObjectUserDomainControllerExtraServerSourcePort(i["source-port"], d, pre_append)
			tmp["source_port"] = fortiAPISubPartPatch(v, "ObjectUserDomainController-ExtraServer-SourcePort")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectUserDomainControllerExtraServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerExtraServerIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerExtraServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerExtraServerSourceIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerExtraServerSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserDomainControllerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserDomainControllerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerReplicationPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerSourceIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDomainControllerUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserDomainController(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ad_mode", flattenObjectUserDomainControllerAdMode(o["ad-mode"], d, "ad_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ad-mode"], "ObjectUserDomainController-AdMode"); ok {
			if err = d.Set("ad_mode", vv); err != nil {
				return fmt.Errorf("Error reading ad_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ad_mode: %v", err)
		}
	}

	if err = d.Set("adlds_dn", flattenObjectUserDomainControllerAdldsDn(o["adlds-dn"], d, "adlds_dn")); err != nil {
		if vv, ok := fortiAPIPatch(o["adlds-dn"], "ObjectUserDomainController-AdldsDn"); ok {
			if err = d.Set("adlds_dn", vv); err != nil {
				return fmt.Errorf("Error reading adlds_dn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adlds_dn: %v", err)
		}
	}

	if err = d.Set("adlds_ip_address", flattenObjectUserDomainControllerAdldsIpAddress(o["adlds-ip-address"], d, "adlds_ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["adlds-ip-address"], "ObjectUserDomainController-AdldsIpAddress"); ok {
			if err = d.Set("adlds_ip_address", vv); err != nil {
				return fmt.Errorf("Error reading adlds_ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adlds_ip_address: %v", err)
		}
	}

	if err = d.Set("adlds_ip6", flattenObjectUserDomainControllerAdldsIp6(o["adlds-ip6"], d, "adlds_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["adlds-ip6"], "ObjectUserDomainController-AdldsIp6"); ok {
			if err = d.Set("adlds_ip6", vv); err != nil {
				return fmt.Errorf("Error reading adlds_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adlds_ip6: %v", err)
		}
	}

	if err = d.Set("adlds_port", flattenObjectUserDomainControllerAdldsPort(o["adlds-port"], d, "adlds_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["adlds-port"], "ObjectUserDomainController-AdldsPort"); ok {
			if err = d.Set("adlds_port", vv); err != nil {
				return fmt.Errorf("Error reading adlds_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adlds_port: %v", err)
		}
	}

	if err = d.Set("dns_srv_lookup", flattenObjectUserDomainControllerDnsSrvLookup(o["dns-srv-lookup"], d, "dns_srv_lookup")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-srv-lookup"], "ObjectUserDomainController-DnsSrvLookup"); ok {
			if err = d.Set("dns_srv_lookup", vv); err != nil {
				return fmt.Errorf("Error reading dns_srv_lookup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_srv_lookup: %v", err)
		}
	}

	if err = d.Set("domain_name", flattenObjectUserDomainControllerDomainName(o["domain-name"], d, "domain_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-name"], "ObjectUserDomainController-DomainName"); ok {
			if err = d.Set("domain_name", vv); err != nil {
				return fmt.Errorf("Error reading domain_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("extra_server", flattenObjectUserDomainControllerExtraServer(o["extra-server"], d, "extra_server")); err != nil {
			if vv, ok := fortiAPIPatch(o["extra-server"], "ObjectUserDomainController-ExtraServer"); ok {
				if err = d.Set("extra_server", vv); err != nil {
					return fmt.Errorf("Error reading extra_server: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading extra_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("extra_server"); ok {
			if err = d.Set("extra_server", flattenObjectUserDomainControllerExtraServer(o["extra-server"], d, "extra_server")); err != nil {
				if vv, ok := fortiAPIPatch(o["extra-server"], "ObjectUserDomainController-ExtraServer"); ok {
					if err = d.Set("extra_server", vv); err != nil {
						return fmt.Errorf("Error reading extra_server: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading extra_server: %v", err)
				}
			}
		}
	}

	if err = d.Set("hostname", flattenObjectUserDomainControllerHostname(o["hostname"], d, "hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname"], "ObjectUserDomainController-Hostname"); ok {
			if err = d.Set("hostname", vv); err != nil {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectUserDomainControllerInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectUserDomainController-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenObjectUserDomainControllerInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "ObjectUserDomainController-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ip_address", flattenObjectUserDomainControllerIpAddress(o["ip-address"], d, "ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-address"], "ObjectUserDomainController-IpAddress"); ok {
			if err = d.Set("ip_address", vv); err != nil {
				return fmt.Errorf("Error reading ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_address: %v", err)
		}
	}

	if err = d.Set("ip6", flattenObjectUserDomainControllerIp6(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "ObjectUserDomainController-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenObjectUserDomainControllerLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "ObjectUserDomainController-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserDomainControllerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserDomainController-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("password", flattenObjectUserDomainControllerPassword(o["password"], d, "password")); err != nil {
		if vv, ok := fortiAPIPatch(o["password"], "ObjectUserDomainController-Password"); ok {
			if err = d.Set("password", vv); err != nil {
				return fmt.Errorf("Error reading password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectUserDomainControllerPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectUserDomainController-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("replication_port", flattenObjectUserDomainControllerReplicationPort(o["replication-port"], d, "replication_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["replication-port"], "ObjectUserDomainController-ReplicationPort"); ok {
			if err = d.Set("replication_port", vv); err != nil {
				return fmt.Errorf("Error reading replication_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replication_port: %v", err)
		}
	}

	if err = d.Set("source_ip_address", flattenObjectUserDomainControllerSourceIpAddress(o["source-ip-address"], d, "source_ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-address"], "ObjectUserDomainController-SourceIpAddress"); ok {
			if err = d.Set("source_ip_address", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_address: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenObjectUserDomainControllerSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip6"], "ObjectUserDomainController-SourceIp6"); ok {
			if err = d.Set("source_ip6", vv); err != nil {
				return fmt.Errorf("Error reading source_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("source_port", flattenObjectUserDomainControllerSourcePort(o["source-port"], d, "source_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-port"], "ObjectUserDomainController-SourcePort"); ok {
			if err = d.Set("source_port", vv); err != nil {
				return fmt.Errorf("Error reading source_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_port: %v", err)
		}
	}

	if err = d.Set("username", flattenObjectUserDomainControllerUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "ObjectUserDomainController-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	return nil
}

func flattenObjectUserDomainControllerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserDomainControllerAdMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerAdldsDn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerAdldsIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerAdldsIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerAdldsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerDnsSrvLookup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerDomainName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerExtraServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandObjectUserDomainControllerExtraServerId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip-address"], _ = expandObjectUserDomainControllerExtraServerIpAddress(d, i["ip_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandObjectUserDomainControllerExtraServerPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source-ip-address"], _ = expandObjectUserDomainControllerExtraServerSourceIpAddress(d, i["source_ip_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source-port"], _ = expandObjectUserDomainControllerExtraServerSourcePort(d, i["source_port"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectUserDomainControllerExtraServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerExtraServerIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerExtraServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerExtraServerSourceIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerExtraServerSourcePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectUserDomainControllerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserDomainControllerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerReplicationPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerSourceIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerSourcePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDomainControllerUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserDomainController(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ad_mode"); ok {
		t, err := expandObjectUserDomainControllerAdMode(d, v, "ad_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ad-mode"] = t
		}
	}

	if v, ok := d.GetOk("adlds_dn"); ok {
		t, err := expandObjectUserDomainControllerAdldsDn(d, v, "adlds_dn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adlds-dn"] = t
		}
	}

	if v, ok := d.GetOk("adlds_ip_address"); ok {
		t, err := expandObjectUserDomainControllerAdldsIpAddress(d, v, "adlds_ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adlds-ip-address"] = t
		}
	}

	if v, ok := d.GetOk("adlds_ip6"); ok {
		t, err := expandObjectUserDomainControllerAdldsIp6(d, v, "adlds_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adlds-ip6"] = t
		}
	}

	if v, ok := d.GetOk("adlds_port"); ok {
		t, err := expandObjectUserDomainControllerAdldsPort(d, v, "adlds_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adlds-port"] = t
		}
	}

	if v, ok := d.GetOk("dns_srv_lookup"); ok {
		t, err := expandObjectUserDomainControllerDnsSrvLookup(d, v, "dns_srv_lookup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-srv-lookup"] = t
		}
	}

	if v, ok := d.GetOk("domain_name"); ok {
		t, err := expandObjectUserDomainControllerDomainName(d, v, "domain_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name"] = t
		}
	}

	if v, ok := d.GetOk("extra_server"); ok {
		t, err := expandObjectUserDomainControllerExtraServer(d, v, "extra_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extra-server"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok {
		t, err := expandObjectUserDomainControllerHostname(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandObjectUserDomainControllerInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		t, err := expandObjectUserDomainControllerInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ip_address"); ok {
		t, err := expandObjectUserDomainControllerIpAddress(d, v, "ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-address"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {
		t, err := expandObjectUserDomainControllerIp6(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok {
		t, err := expandObjectUserDomainControllerLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectUserDomainControllerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandObjectUserDomainControllerPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandObjectUserDomainControllerPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("replication_port"); ok {
		t, err := expandObjectUserDomainControllerReplicationPort(d, v, "replication_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replication-port"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_address"); ok {
		t, err := expandObjectUserDomainControllerSourceIpAddress(d, v, "source_ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-address"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok {
		t, err := expandObjectUserDomainControllerSourceIp6(d, v, "source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("source_port"); ok {
		t, err := expandObjectUserDomainControllerSourcePort(d, v, "source_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-port"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandObjectUserDomainControllerUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	return &obj, nil
}
