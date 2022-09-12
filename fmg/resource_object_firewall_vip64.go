// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 to IPv4 virtual IPs.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallVip64() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallVip64Create,
		Read:   resourceObjectFirewallVip64Read,
		Update: resourceObjectFirewallVip64Update,
		Delete: resourceObjectFirewallVip64Delete,

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
			"arp_reply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
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
						"arp_reply": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"color": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"extip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"extport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ldb_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mappedip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mappedport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"monitor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"portforward": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"server_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"src_filter": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"extip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"extport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ldb_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mappedip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mappedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"portforward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"realservers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"healthcheck": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"holddown_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_connections": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"monitor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
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

func resourceObjectFirewallVip64Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallVip64(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallVip64 resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallVip64(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallVip64 resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallVip64Read(d, m)
}

func resourceObjectFirewallVip64Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallVip64(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallVip64 resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallVip64(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallVip64 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallVip64Read(d, m)
}

func resourceObjectFirewallVip64Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallVip64(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallVip64 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallVip64Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallVip64(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallVip64 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallVip64(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallVip64 resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallVip64ArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64Color(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64Comment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallVip64DynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := i["arp-reply"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingArpReply(i["arp-reply"], d, pre_append)
			tmp["arp_reply"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-ArpReply")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := i["color"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingColor(i["color"], d, pre_append)
			tmp["color"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-Color")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extip"
		if _, ok := i["extip"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingExtip(i["extip"], d, pre_append)
			tmp["extip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-Extip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extport"
		if _, ok := i["extport"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingExtport(i["extport"], d, pre_append)
			tmp["extport"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-Extport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := i["ldb-method"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingLdbMethod(i["ldb-method"], d, pre_append)
			tmp["ldb_method"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-LdbMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedip"
		if _, ok := i["mappedip"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingMappedip(i["mappedip"], d, pre_append)
			tmp["mappedip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-Mappedip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := i["mappedport"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingMappedport(i["mappedport"], d, pre_append)
			tmp["mappedport"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-Mappedport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portforward"
		if _, ok := i["portforward"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingPortforward(i["portforward"], d, pre_append)
			tmp["portforward"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-Portforward")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := i["server-type"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingServerType(i["server-type"], d, pre_append)
			tmp["server_type"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-ServerType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_filter"
		if _, ok := i["src-filter"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingSrcFilter(i["src-filter"], d, pre_append)
			tmp["src_filter"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-SrcFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := i["uuid"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingUuid(i["uuid"], d, pre_append)
			tmp["uuid"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-DynamicMapping-Uuid")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallVip64DynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallVip64DynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64DynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectFirewallVip64DynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64DynamicMapping-Scope-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallVip64DynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingExtip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingExtport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingLdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingMappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingMappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingPortforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingSrcFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVip64DynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64DynamicMappingUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64Extip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64Extport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64LdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64Mappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64Mappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64Monitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVip64Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64Portforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64Protocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64Realservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := i["client-ip"]; ok {
			v := flattenObjectFirewallVip64RealserversClientIp(i["client-ip"], d, pre_append)
			tmp["client_ip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-Realservers-ClientIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := i["healthcheck"]; ok {
			v := flattenObjectFirewallVip64RealserversHealthcheck(i["healthcheck"], d, pre_append)
			tmp["healthcheck"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-Realservers-Healthcheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			v := flattenObjectFirewallVip64RealserversHolddownInterval(i["holddown-interval"], d, pre_append)
			tmp["holddown_interval"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-Realservers-HolddownInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallVip64RealserversId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-Realservers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFirewallVip64RealserversIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-Realservers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := i["max-connections"]; ok {
			v := flattenObjectFirewallVip64RealserversMaxConnections(i["max-connections"], d, pre_append)
			tmp["max_connections"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-Realservers-MaxConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenObjectFirewallVip64RealserversMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-Realservers-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectFirewallVip64RealserversPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-Realservers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectFirewallVip64RealserversStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-Realservers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectFirewallVip64RealserversWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectFirewallVip64-Realservers-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallVip64RealserversClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64RealserversHealthcheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64RealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64RealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64RealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64RealserversMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64RealserversMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64RealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64RealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64RealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64ServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64SrcFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallVip64Type(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVip64Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallVip64(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("arp_reply", flattenObjectFirewallVip64ArpReply(o["arp-reply"], d, "arp_reply")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-reply"], "ObjectFirewallVip64-ArpReply"); ok {
			if err = d.Set("arp_reply", vv); err != nil {
				return fmt.Errorf("Error reading arp_reply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_reply: %v", err)
		}
	}

	if err = d.Set("color", flattenObjectFirewallVip64Color(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "ObjectFirewallVip64-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectFirewallVip64Comment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectFirewallVip64-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectFirewallVip64DynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFirewallVip64-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectFirewallVip64DynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFirewallVip64-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("extip", flattenObjectFirewallVip64Extip(o["extip"], d, "extip")); err != nil {
		if vv, ok := fortiAPIPatch(o["extip"], "ObjectFirewallVip64-Extip"); ok {
			if err = d.Set("extip", vv); err != nil {
				return fmt.Errorf("Error reading extip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extip: %v", err)
		}
	}

	if err = d.Set("extport", flattenObjectFirewallVip64Extport(o["extport"], d, "extport")); err != nil {
		if vv, ok := fortiAPIPatch(o["extport"], "ObjectFirewallVip64-Extport"); ok {
			if err = d.Set("extport", vv); err != nil {
				return fmt.Errorf("Error reading extport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extport: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallVip64Id(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallVip64-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ldb_method", flattenObjectFirewallVip64LdbMethod(o["ldb-method"], d, "ldb_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldb-method"], "ObjectFirewallVip64-LdbMethod"); ok {
			if err = d.Set("ldb_method", vv); err != nil {
				return fmt.Errorf("Error reading ldb_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldb_method: %v", err)
		}
	}

	if err = d.Set("mappedip", flattenObjectFirewallVip64Mappedip(o["mappedip"], d, "mappedip")); err != nil {
		if vv, ok := fortiAPIPatch(o["mappedip"], "ObjectFirewallVip64-Mappedip"); ok {
			if err = d.Set("mappedip", vv); err != nil {
				return fmt.Errorf("Error reading mappedip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mappedip: %v", err)
		}
	}

	if err = d.Set("mappedport", flattenObjectFirewallVip64Mappedport(o["mappedport"], d, "mappedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["mappedport"], "ObjectFirewallVip64-Mappedport"); ok {
			if err = d.Set("mappedport", vv); err != nil {
				return fmt.Errorf("Error reading mappedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mappedport: %v", err)
		}
	}

	if err = d.Set("monitor", flattenObjectFirewallVip64Monitor(o["monitor"], d, "monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor"], "ObjectFirewallVip64-Monitor"); ok {
			if err = d.Set("monitor", vv); err != nil {
				return fmt.Errorf("Error reading monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallVip64Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallVip64-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("portforward", flattenObjectFirewallVip64Portforward(o["portforward"], d, "portforward")); err != nil {
		if vv, ok := fortiAPIPatch(o["portforward"], "ObjectFirewallVip64-Portforward"); ok {
			if err = d.Set("portforward", vv); err != nil {
				return fmt.Errorf("Error reading portforward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portforward: %v", err)
		}
	}

	if err = d.Set("protocol", flattenObjectFirewallVip64Protocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ObjectFirewallVip64-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("realservers", flattenObjectFirewallVip64Realservers(o["realservers"], d, "realservers")); err != nil {
			if vv, ok := fortiAPIPatch(o["realservers"], "ObjectFirewallVip64-Realservers"); ok {
				if err = d.Set("realservers", vv); err != nil {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading realservers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("realservers"); ok {
			if err = d.Set("realservers", flattenObjectFirewallVip64Realservers(o["realservers"], d, "realservers")); err != nil {
				if vv, ok := fortiAPIPatch(o["realservers"], "ObjectFirewallVip64-Realservers"); ok {
					if err = d.Set("realservers", vv); err != nil {
						return fmt.Errorf("Error reading realservers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_type", flattenObjectFirewallVip64ServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "ObjectFirewallVip64-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("src_filter", flattenObjectFirewallVip64SrcFilter(o["src-filter"], d, "src_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-filter"], "ObjectFirewallVip64-SrcFilter"); ok {
			if err = d.Set("src_filter", vv); err != nil {
				return fmt.Errorf("Error reading src_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_filter: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFirewallVip64Type(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallVip64-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectFirewallVip64Uuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectFirewallVip64-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallVip64FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallVip64ArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64Color(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64Comment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_scope"], _ = expandObjectFirewallVip64DynamicMappingScope(d, i["_scope"], pre_append)
		} else {
			tmp["_scope"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["arp-reply"], _ = expandObjectFirewallVip64DynamicMappingArpReply(d, i["arp_reply"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color"], _ = expandObjectFirewallVip64DynamicMappingColor(d, i["color"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandObjectFirewallVip64DynamicMappingComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extip"], _ = expandObjectFirewallVip64DynamicMappingExtip(d, i["extip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extport"], _ = expandObjectFirewallVip64DynamicMappingExtport(d, i["extport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFirewallVip64DynamicMappingId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldb-method"], _ = expandObjectFirewallVip64DynamicMappingLdbMethod(d, i["ldb_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedip"], _ = expandObjectFirewallVip64DynamicMappingMappedip(d, i["mappedip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mappedport"], _ = expandObjectFirewallVip64DynamicMappingMappedport(d, i["mappedport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandObjectFirewallVip64DynamicMappingMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portforward"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["portforward"], _ = expandObjectFirewallVip64DynamicMappingPortforward(d, i["portforward"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandObjectFirewallVip64DynamicMappingProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-type"], _ = expandObjectFirewallVip64DynamicMappingServerType(d, i["server_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-filter"], _ = expandObjectFirewallVip64DynamicMappingSrcFilter(d, i["src_filter"], pre_append)
		} else {
			tmp["src-filter"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFirewallVip64DynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uuid"], _ = expandObjectFirewallVip64DynamicMappingUuid(d, i["uuid"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVip64DynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectFirewallVip64DynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectFirewallVip64DynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVip64DynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingExtip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingExtport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingLdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingMappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingMappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingPortforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingSrcFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVip64DynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64DynamicMappingUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64Extip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64Extport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64LdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64Mappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64Mappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64Monitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectFirewallVip64Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64Portforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64Protocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64Realservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-ip"], _ = expandObjectFirewallVip64RealserversClientIp(d, i["client_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["healthcheck"], _ = expandObjectFirewallVip64RealserversHealthcheck(d, i["healthcheck"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["holddown-interval"], _ = expandObjectFirewallVip64RealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFirewallVip64RealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFirewallVip64RealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-connections"], _ = expandObjectFirewallVip64RealserversMaxConnections(d, i["max_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandObjectFirewallVip64RealserversMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectFirewallVip64RealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectFirewallVip64RealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectFirewallVip64RealserversWeight(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallVip64RealserversClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64RealserversHealthcheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64RealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64RealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64RealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64RealserversMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64RealserversMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64RealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64RealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64RealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64ServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64SrcFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallVip64Type(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVip64Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallVip64(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("arp_reply"); ok || d.HasChange("arp_reply") {
		t, err := expandObjectFirewallVip64ArpReply(d, v, "arp_reply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-reply"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandObjectFirewallVip64Color(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectFirewallVip64Comment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectFirewallVip64DynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("extip"); ok || d.HasChange("extip") {
		t, err := expandObjectFirewallVip64Extip(d, v, "extip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extip"] = t
		}
	}

	if v, ok := d.GetOk("extport"); ok || d.HasChange("extport") {
		t, err := expandObjectFirewallVip64Extport(d, v, "extport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extport"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandObjectFirewallVip64Id(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ldb_method"); ok || d.HasChange("ldb_method") {
		t, err := expandObjectFirewallVip64LdbMethod(d, v, "ldb_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-method"] = t
		}
	}

	if v, ok := d.GetOk("mappedip"); ok || d.HasChange("mappedip") {
		t, err := expandObjectFirewallVip64Mappedip(d, v, "mappedip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedip"] = t
		}
	}

	if v, ok := d.GetOk("mappedport"); ok || d.HasChange("mappedport") {
		t, err := expandObjectFirewallVip64Mappedport(d, v, "mappedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedport"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok || d.HasChange("monitor") {
		t, err := expandObjectFirewallVip64Monitor(d, v, "monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallVip64Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("portforward"); ok || d.HasChange("portforward") {
		t, err := expandObjectFirewallVip64Portforward(d, v, "portforward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portforward"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandObjectFirewallVip64Protocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("realservers"); ok || d.HasChange("realservers") {
		t, err := expandObjectFirewallVip64Realservers(d, v, "realservers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realservers"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandObjectFirewallVip64ServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("src_filter"); ok || d.HasChange("src_filter") {
		t, err := expandObjectFirewallVip64SrcFilter(d, v, "src_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-filter"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFirewallVip64Type(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandObjectFirewallVip64Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
