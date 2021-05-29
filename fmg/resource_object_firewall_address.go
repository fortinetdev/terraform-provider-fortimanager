// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv4 addresses.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallAddressCreate,
		Read:   resourceObjectFirewallAddressRead,
		Update: resourceObjectFirewallAddressUpdate,
		Delete: resourceObjectFirewallAddressDelete,

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
			"_image_base64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"clearpass_spt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_image_base64": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
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
						"allow_routing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"associated_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cache_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"clearpass_spt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"color": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"country": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"end_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"epg_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fabric_object": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fsso_group": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"global_object": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"node_ip_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"obj_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"obj_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"obj_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"organization": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"policy_group": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sdn_addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sdn_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"start_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sub_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subnet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subnet_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tenant": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"visibility": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wildcard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wildcard_fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"epg_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsso_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"global_object": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"net_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"obj_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"organization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdn_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdn_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sub_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"subnet_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tagging": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"tenant": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wildcard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wildcard_fqdn": &schema.Schema{
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

func resourceObjectFirewallAddressCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallAddress(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAddress resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallAddress(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAddress resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallAddressRead(d, m)
}

func resourceObjectFirewallAddressUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallAddress(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAddress resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallAddress(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallAddressRead(d, m)
}

func resourceObjectFirewallAddressDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallAddress(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallAddressRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallAddress(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAddress resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallAddressImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressAllowRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectFirewallAddressAssociatedInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressClearpassSpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "unknown",
			1: "healthy",
			2: "quarantine",
			3: "checkup",
			5: "infected",
			6: "transient",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallAddressColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_image_base64"
		if _, ok := i["_image-base64"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingImageBase64(i["_image-base64"], d, pre_append)
			tmp["_image_base64"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-ImageBase64")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allow_routing"
		if _, ok := i["allow-routing"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingAllowRouting(i["allow-routing"], d, pre_append)
			tmp["allow_routing"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-AllowRouting")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "associated_interface"
		if _, ok := i["associated-interface"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingAssociatedInterface(i["associated-interface"], d, pre_append)
			tmp["associated_interface"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-AssociatedInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cache_ttl"
		if _, ok := i["cache-ttl"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingCacheTtl(i["cache-ttl"], d, pre_append)
			tmp["cache_ttl"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-CacheTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "clearpass_spt"
		if _, ok := i["clearpass-spt"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingClearpassSpt(i["clearpass-spt"], d, pre_append)
			tmp["clearpass_spt"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-ClearpassSpt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := i["color"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingColor(i["color"], d, pre_append)
			tmp["color"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Color")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "country"
		if _, ok := i["country"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingCountry(i["country"], d, pre_append)
			tmp["country"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Country")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_mac"
		if _, ok := i["end-mac"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingEndMac(i["end-mac"], d, pre_append)
			tmp["end_mac"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-EndMac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "epg_name"
		if _, ok := i["epg-name"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingEpgName(i["epg-name"], d, pre_append)
			tmp["epg_name"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-EpgName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object"
		if _, ok := i["fabric-object"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingFabricObject(i["fabric-object"], d, pre_append)
			tmp["fabric_object"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-FabricObject")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingFilter(i["filter"], d, pre_append)
			tmp["filter"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Filter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := i["fqdn"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingFqdn(i["fqdn"], d, pre_append)
			tmp["fqdn"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Fqdn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fsso_group"
		if _, ok := i["fsso-group"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingFssoGroup(i["fsso-group"], d, pre_append)
			tmp["fsso_group"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-FssoGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "global_object"
		if _, ok := i["global-object"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingGlobalObject(i["global-object"], d, pre_append)
			tmp["global_object"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-GlobalObject")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "node_ip_only"
		if _, ok := i["node-ip-only"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingNodeIpOnly(i["node-ip-only"], d, pre_append)
			tmp["node_ip_only"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-NodeIpOnly")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := i["obj-id"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingObjId(i["obj-id"], d, pre_append)
			tmp["obj_id"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-ObjId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_tag"
		if _, ok := i["obj-tag"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingObjTag(i["obj-tag"], d, pre_append)
			tmp["obj_tag"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-ObjTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_type"
		if _, ok := i["obj-type"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingObjType(i["obj-type"], d, pre_append)
			tmp["obj_type"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-ObjType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "organization"
		if _, ok := i["organization"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingOrganization(i["organization"], d, pre_append)
			tmp["organization"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Organization")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy_group"
		if _, ok := i["policy-group"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingPolicyGroup(i["policy-group"], d, pre_append)
			tmp["policy_group"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-PolicyGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn"
		if _, ok := i["sdn"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingSdn(i["sdn"], d, pre_append)
			tmp["sdn"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Sdn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn_addr_type"
		if _, ok := i["sdn-addr-type"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingSdnAddrType(i["sdn-addr-type"], d, pre_append)
			tmp["sdn_addr_type"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-SdnAddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn_tag"
		if _, ok := i["sdn-tag"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingSdnTag(i["sdn-tag"], d, pre_append)
			tmp["sdn_tag"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-SdnTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_mac"
		if _, ok := i["start-mac"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingStartMac(i["start-mac"], d, pre_append)
			tmp["start_mac"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-StartMac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sub_type"
		if _, ok := i["sub-type"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingSubType(i["sub-type"], d, pre_append)
			tmp["sub_type"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-SubType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingSubnet(i["subnet"], d, pre_append)
			tmp["subnet"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet_name"
		if _, ok := i["subnet-name"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingSubnetName(i["subnet-name"], d, pre_append)
			tmp["subnet_name"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-SubnetName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingTags(i["tags"], d, pre_append)
			tmp["tags"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Tags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant"
		if _, ok := i["tenant"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingTenant(i["tenant"], d, pre_append)
			tmp["tenant"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Tenant")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := i["url"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingUrl(i["url"], d, pre_append)
			tmp["url"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Url")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := i["uuid"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingUuid(i["uuid"], d, pre_append)
			tmp["uuid"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Uuid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "visibility"
		if _, ok := i["visibility"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingVisibility(i["visibility"], d, pre_append)
			tmp["visibility"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Visibility")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard"
		if _, ok := i["wildcard"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingWildcard(i["wildcard"], d, pre_append)
			tmp["wildcard"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-Wildcard")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard_fqdn"
		if _, ok := i["wildcard-fqdn"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingWildcardFqdn(i["wildcard-fqdn"], d, pre_append)
			tmp["wildcard_fqdn"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-DynamicMapping-WildcardFqdn")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallAddressDynamicMappingImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallAddressDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallAddressDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectFirewallAddressDynamicMapping-Scope-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallAddressDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingAllowRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectFirewallAddressDynamicMappingAssociatedInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingClearpassSpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "unknown",
			1: "healthy",
			2: "quarantine",
			3: "checkup",
			5: "infected",
			6: "transient",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallAddressDynamicMappingColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingEndMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingEpgName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectFirewallAddressDynamicMappingFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingFssoGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingGlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingNodeIpOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectFirewallAddressDynamicMappingObjId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingObjTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingObjType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			9:  "ip",
			18: "mac",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallAddressDynamicMappingOrganization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingPolicyGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingSdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingSdnAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "private",
			1: "public",
			2: "all",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallAddressDynamicMappingSdnTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingStartMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingSubType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "sdn",
			1: "clearpass-spt",
			2: "fsso",
			3: "ems-tag",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallAddressDynamicMappingSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingSubnetName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingTenant(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:  "ipmask",
			1:  "iprange",
			2:  "fqdn",
			3:  "wildcard",
			4:  "geography",
			15: "dynamic",
			16: "interface-subnet",
			17: "mac",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallAddressDynamicMappingUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectFirewallAddressDynamicMappingWildcard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingWildcardFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressEndMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressEpgName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressFssoGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressGlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFirewallAddressListIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-List-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "net_id"
		if _, ok := i["net-id"]; ok {
			v := flattenObjectFirewallAddressListNetId(i["net-id"], d, pre_append)
			tmp["net_id"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-List-NetId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := i["obj-id"]; ok {
			v := flattenObjectFirewallAddressListObjId(i["obj-id"], d, pre_append)
			tmp["obj_id"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-List-ObjId")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallAddressListIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressListNetId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressListObjId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressObjId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressObjTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressObjType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			9:  "ip",
			18: "mac",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallAddressOrganization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressPolicyGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressSdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressSdnAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "private",
			1: "public",
			2: "all",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallAddressSdnTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressStartMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressSubType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "sdn",
			1: "clearpass-spt",
			2: "fsso",
			3: "ems-tag",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallAddressSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallAddressSubnetName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressTagging(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenObjectFirewallAddressTaggingCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-Tagging-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectFirewallAddressTaggingName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-Tagging-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			v := flattenObjectFirewallAddressTaggingTags(i["tags"], d, pre_append)
			tmp["tags"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress-Tagging-Tags")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallAddressTaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressTaggingName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressTaggingTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallAddressTenant(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:  "ipmask",
			1:  "iprange",
			2:  "fqdn",
			3:  "wildcard",
			4:  "geography",
			15: "dynamic",
			16: "interface-subnet",
			17: "mac",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectFirewallAddressUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressWildcard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressWildcardFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("_image_base64", flattenObjectFirewallAddressImageBase64(o["_image-base64"], d, "_image_base64")); err != nil {
		if vv, ok := fortiAPIPatch(o["_image-base64"], "ObjectFirewallAddress-ImageBase64"); ok {
			if err = d.Set("_image_base64", vv); err != nil {
				return fmt.Errorf("Error reading _image_base64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _image_base64: %v", err)
		}
	}

	if err = d.Set("allow_routing", flattenObjectFirewallAddressAllowRouting(o["allow-routing"], d, "allow_routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-routing"], "ObjectFirewallAddress-AllowRouting"); ok {
			if err = d.Set("allow_routing", vv); err != nil {
				return fmt.Errorf("Error reading allow_routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_routing: %v", err)
		}
	}

	if err = d.Set("associated_interface", flattenObjectFirewallAddressAssociatedInterface(o["associated-interface"], d, "associated_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["associated-interface"], "ObjectFirewallAddress-AssociatedInterface"); ok {
			if err = d.Set("associated_interface", vv); err != nil {
				return fmt.Errorf("Error reading associated_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading associated_interface: %v", err)
		}
	}

	if err = d.Set("cache_ttl", flattenObjectFirewallAddressCacheTtl(o["cache-ttl"], d, "cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-ttl"], "ObjectFirewallAddress-CacheTtl"); ok {
			if err = d.Set("cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_ttl: %v", err)
		}
	}

	if err = d.Set("clearpass_spt", flattenObjectFirewallAddressClearpassSpt(o["clearpass-spt"], d, "clearpass_spt")); err != nil {
		if vv, ok := fortiAPIPatch(o["clearpass-spt"], "ObjectFirewallAddress-ClearpassSpt"); ok {
			if err = d.Set("clearpass_spt", vv); err != nil {
				return fmt.Errorf("Error reading clearpass_spt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clearpass_spt: %v", err)
		}
	}

	if err = d.Set("color", flattenObjectFirewallAddressColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "ObjectFirewallAddress-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectFirewallAddressComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectFirewallAddress-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("country", flattenObjectFirewallAddressCountry(o["country"], d, "country")); err != nil {
		if vv, ok := fortiAPIPatch(o["country"], "ObjectFirewallAddress-Country"); ok {
			if err = d.Set("country", vv); err != nil {
				return fmt.Errorf("Error reading country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectFirewallAddressDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFirewallAddress-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectFirewallAddressDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFirewallAddress-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("end_ip", flattenObjectFirewallAddressEndIp(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "ObjectFirewallAddress-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("end_mac", flattenObjectFirewallAddressEndMac(o["end-mac"], d, "end_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-mac"], "ObjectFirewallAddress-EndMac"); ok {
			if err = d.Set("end_mac", vv); err != nil {
				return fmt.Errorf("Error reading end_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_mac: %v", err)
		}
	}

	if err = d.Set("epg_name", flattenObjectFirewallAddressEpgName(o["epg-name"], d, "epg_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["epg-name"], "ObjectFirewallAddress-EpgName"); ok {
			if err = d.Set("epg_name", vv); err != nil {
				return fmt.Errorf("Error reading epg_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading epg_name: %v", err)
		}
	}

	if err = d.Set("filter", flattenObjectFirewallAddressFilter(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "ObjectFirewallAddress-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenObjectFirewallAddressFqdn(o["fqdn"], d, "fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn"], "ObjectFirewallAddress-Fqdn"); ok {
			if err = d.Set("fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("fsso_group", flattenObjectFirewallAddressFssoGroup(o["fsso-group"], d, "fsso_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-group"], "ObjectFirewallAddress-FssoGroup"); ok {
			if err = d.Set("fsso_group", vv); err != nil {
				return fmt.Errorf("Error reading fsso_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_group: %v", err)
		}
	}

	if err = d.Set("global_object", flattenObjectFirewallAddressGlobalObject(o["global-object"], d, "global_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-object"], "ObjectFirewallAddress-GlobalObject"); ok {
			if err = d.Set("global_object", vv); err != nil {
				return fmt.Errorf("Error reading global_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_object: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectFirewallAddressInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectFirewallAddress-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("list", flattenObjectFirewallAddressList(o["list"], d, "list")); err != nil {
			if vv, ok := fortiAPIPatch(o["list"], "ObjectFirewallAddress-List"); ok {
				if err = d.Set("list", vv); err != nil {
					return fmt.Errorf("Error reading list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("list"); ok {
			if err = d.Set("list", flattenObjectFirewallAddressList(o["list"], d, "list")); err != nil {
				if vv, ok := fortiAPIPatch(o["list"], "ObjectFirewallAddress-List"); ok {
					if err = d.Set("list", vv); err != nil {
						return fmt.Errorf("Error reading list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading list: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectFirewallAddressName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallAddress-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("obj_id", flattenObjectFirewallAddressObjId(o["obj-id"], d, "obj_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["obj-id"], "ObjectFirewallAddress-ObjId"); ok {
			if err = d.Set("obj_id", vv); err != nil {
				return fmt.Errorf("Error reading obj_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obj_id: %v", err)
		}
	}

	if err = d.Set("obj_tag", flattenObjectFirewallAddressObjTag(o["obj-tag"], d, "obj_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["obj-tag"], "ObjectFirewallAddress-ObjTag"); ok {
			if err = d.Set("obj_tag", vv); err != nil {
				return fmt.Errorf("Error reading obj_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obj_tag: %v", err)
		}
	}

	if err = d.Set("obj_type", flattenObjectFirewallAddressObjType(o["obj-type"], d, "obj_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["obj-type"], "ObjectFirewallAddress-ObjType"); ok {
			if err = d.Set("obj_type", vv); err != nil {
				return fmt.Errorf("Error reading obj_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obj_type: %v", err)
		}
	}

	if err = d.Set("organization", flattenObjectFirewallAddressOrganization(o["organization"], d, "organization")); err != nil {
		if vv, ok := fortiAPIPatch(o["organization"], "ObjectFirewallAddress-Organization"); ok {
			if err = d.Set("organization", vv); err != nil {
				return fmt.Errorf("Error reading organization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading organization: %v", err)
		}
	}

	if err = d.Set("policy_group", flattenObjectFirewallAddressPolicyGroup(o["policy-group"], d, "policy_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-group"], "ObjectFirewallAddress-PolicyGroup"); ok {
			if err = d.Set("policy_group", vv); err != nil {
				return fmt.Errorf("Error reading policy_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_group: %v", err)
		}
	}

	if err = d.Set("sdn", flattenObjectFirewallAddressSdn(o["sdn"], d, "sdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn"], "ObjectFirewallAddress-Sdn"); ok {
			if err = d.Set("sdn", vv); err != nil {
				return fmt.Errorf("Error reading sdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn: %v", err)
		}
	}

	if err = d.Set("sdn_addr_type", flattenObjectFirewallAddressSdnAddrType(o["sdn-addr-type"], d, "sdn_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn-addr-type"], "ObjectFirewallAddress-SdnAddrType"); ok {
			if err = d.Set("sdn_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading sdn_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn_addr_type: %v", err)
		}
	}

	if err = d.Set("sdn_tag", flattenObjectFirewallAddressSdnTag(o["sdn-tag"], d, "sdn_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn-tag"], "ObjectFirewallAddress-SdnTag"); ok {
			if err = d.Set("sdn_tag", vv); err != nil {
				return fmt.Errorf("Error reading sdn_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn_tag: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenObjectFirewallAddressStartIp(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "ObjectFirewallAddress-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	if err = d.Set("start_mac", flattenObjectFirewallAddressStartMac(o["start-mac"], d, "start_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-mac"], "ObjectFirewallAddress-StartMac"); ok {
			if err = d.Set("start_mac", vv); err != nil {
				return fmt.Errorf("Error reading start_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_mac: %v", err)
		}
	}

	if err = d.Set("sub_type", flattenObjectFirewallAddressSubType(o["sub-type"], d, "sub_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["sub-type"], "ObjectFirewallAddress-SubType"); ok {
			if err = d.Set("sub_type", vv); err != nil {
				return fmt.Errorf("Error reading sub_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sub_type: %v", err)
		}
	}

	if err = d.Set("subnet", flattenObjectFirewallAddressSubnet(o["subnet"], d, "subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet"], "ObjectFirewallAddress-Subnet"); ok {
			if err = d.Set("subnet", vv); err != nil {
				return fmt.Errorf("Error reading subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet: %v", err)
		}
	}

	if err = d.Set("subnet_name", flattenObjectFirewallAddressSubnetName(o["subnet-name"], d, "subnet_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet-name"], "ObjectFirewallAddress-SubnetName"); ok {
			if err = d.Set("subnet_name", vv); err != nil {
				return fmt.Errorf("Error reading subnet_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet_name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tagging", flattenObjectFirewallAddressTagging(o["tagging"], d, "tagging")); err != nil {
			if vv, ok := fortiAPIPatch(o["tagging"], "ObjectFirewallAddress-Tagging"); ok {
				if err = d.Set("tagging", vv); err != nil {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tagging: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tagging"); ok {
			if err = d.Set("tagging", flattenObjectFirewallAddressTagging(o["tagging"], d, "tagging")); err != nil {
				if vv, ok := fortiAPIPatch(o["tagging"], "ObjectFirewallAddress-Tagging"); ok {
					if err = d.Set("tagging", vv); err != nil {
						return fmt.Errorf("Error reading tagging: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			}
		}
	}

	if err = d.Set("tenant", flattenObjectFirewallAddressTenant(o["tenant"], d, "tenant")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenant"], "ObjectFirewallAddress-Tenant"); ok {
			if err = d.Set("tenant", vv); err != nil {
				return fmt.Errorf("Error reading tenant: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenant: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFirewallAddressType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallAddress-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectFirewallAddressUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectFirewallAddress-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("wildcard", flattenObjectFirewallAddressWildcard(o["wildcard"], d, "wildcard")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard"], "ObjectFirewallAddress-Wildcard"); ok {
			if err = d.Set("wildcard", vv); err != nil {
				return fmt.Errorf("Error reading wildcard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}

	if err = d.Set("wildcard_fqdn", flattenObjectFirewallAddressWildcardFqdn(o["wildcard-fqdn"], d, "wildcard_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard-fqdn"], "ObjectFirewallAddress-WildcardFqdn"); ok {
			if err = d.Set("wildcard_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading wildcard_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard_fqdn: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallAddressImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressAllowRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressAssociatedInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressClearpassSpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_image_base64"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["_image-base64"], _ = expandObjectFirewallAddressDynamicMappingImageBase64(d, i["_image_base64"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["_scope"], _ = expandObjectFirewallAddressDynamicMappingScope(d, i["_scope"], pre_append)
		} else {
			tmp["_scope"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allow_routing"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allow-routing"], _ = expandObjectFirewallAddressDynamicMappingAllowRouting(d, i["allow_routing"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "associated_interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["associated-interface"], _ = expandObjectFirewallAddressDynamicMappingAssociatedInterface(d, i["associated_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cache_ttl"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cache-ttl"], _ = expandObjectFirewallAddressDynamicMappingCacheTtl(d, i["cache_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "clearpass_spt"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["clearpass-spt"], _ = expandObjectFirewallAddressDynamicMappingClearpassSpt(d, i["clearpass_spt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["color"], _ = expandObjectFirewallAddressDynamicMappingColor(d, i["color"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandObjectFirewallAddressDynamicMappingComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "country"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["country"], _ = expandObjectFirewallAddressDynamicMappingCountry(d, i["country"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-ip"], _ = expandObjectFirewallAddressDynamicMappingEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_mac"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-mac"], _ = expandObjectFirewallAddressDynamicMappingEndMac(d, i["end_mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "epg_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["epg-name"], _ = expandObjectFirewallAddressDynamicMappingEpgName(d, i["epg_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fabric-object"], _ = expandObjectFirewallAddressDynamicMappingFabricObject(d, i["fabric_object"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter"], _ = expandObjectFirewallAddressDynamicMappingFilter(d, i["filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fqdn"], _ = expandObjectFirewallAddressDynamicMappingFqdn(d, i["fqdn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fsso_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fsso-group"], _ = expandObjectFirewallAddressDynamicMappingFssoGroup(d, i["fsso_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "global_object"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["global-object"], _ = expandObjectFirewallAddressDynamicMappingGlobalObject(d, i["global_object"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandObjectFirewallAddressDynamicMappingInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "node_ip_only"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["node-ip-only"], _ = expandObjectFirewallAddressDynamicMappingNodeIpOnly(d, i["node_ip_only"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["obj-id"], _ = expandObjectFirewallAddressDynamicMappingObjId(d, i["obj_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_tag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["obj-tag"], _ = expandObjectFirewallAddressDynamicMappingObjTag(d, i["obj_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["obj-type"], _ = expandObjectFirewallAddressDynamicMappingObjType(d, i["obj_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "organization"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["organization"], _ = expandObjectFirewallAddressDynamicMappingOrganization(d, i["organization"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["policy-group"], _ = expandObjectFirewallAddressDynamicMappingPolicyGroup(d, i["policy_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sdn"], _ = expandObjectFirewallAddressDynamicMappingSdn(d, i["sdn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn_addr_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sdn-addr-type"], _ = expandObjectFirewallAddressDynamicMappingSdnAddrType(d, i["sdn_addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn_tag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sdn-tag"], _ = expandObjectFirewallAddressDynamicMappingSdnTag(d, i["sdn_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip"], _ = expandObjectFirewallAddressDynamicMappingStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_mac"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-mac"], _ = expandObjectFirewallAddressDynamicMappingStartMac(d, i["start_mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sub_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sub-type"], _ = expandObjectFirewallAddressDynamicMappingSubType(d, i["sub_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["subnet"], _ = expandObjectFirewallAddressDynamicMappingSubnet(d, i["subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["subnet-name"], _ = expandObjectFirewallAddressDynamicMappingSubnetName(d, i["subnet_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tags"], _ = expandObjectFirewallAddressDynamicMappingTags(d, i["tags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tenant"], _ = expandObjectFirewallAddressDynamicMappingTenant(d, i["tenant"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandObjectFirewallAddressDynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url"], _ = expandObjectFirewallAddressDynamicMappingUrl(d, i["url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["uuid"], _ = expandObjectFirewallAddressDynamicMappingUuid(d, i["uuid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "visibility"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["visibility"], _ = expandObjectFirewallAddressDynamicMappingVisibility(d, i["visibility"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["wildcard"], _ = expandObjectFirewallAddressDynamicMappingWildcard(d, i["wildcard"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard_fqdn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["wildcard-fqdn"], _ = expandObjectFirewallAddressDynamicMappingWildcardFqdn(d, i["wildcard_fqdn"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAddressDynamicMappingImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallAddressDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vdom"], _ = expandObjectFirewallAddressDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAddressDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingAllowRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingAssociatedInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingClearpassSpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingEndMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingEpgName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingFssoGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingGlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingNodeIpOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingObjId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingObjTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingObjType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingOrganization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingPolicyGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingSdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingSdnAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingSdnTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingStartMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingSubType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingSubnetName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingTenant(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingVisibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingWildcard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingWildcardFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressEndMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressEpgName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressFssoGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressGlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandObjectFirewallAddressListIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "net_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["net-id"], _ = expandObjectFirewallAddressListNetId(d, i["net_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["obj-id"], _ = expandObjectFirewallAddressListObjId(d, i["obj_id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAddressListIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressListNetId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressListObjId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressObjId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressObjTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressObjType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressOrganization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressPolicyGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressSdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressSdnAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressSdnTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressStartMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressSubType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectFirewallAddressSubnetName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressTagging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandObjectFirewallAddressTaggingCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandObjectFirewallAddressTaggingName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tags"], _ = expandObjectFirewallAddressTaggingTags(d, i["tags"], pre_append)
		} else {
			tmp["tags"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAddressTaggingCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressTaggingName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressTaggingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallAddressTenant(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressWildcard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressWildcardFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallAddress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_image_base64"); ok {
		t, err := expandObjectFirewallAddressImageBase64(d, v, "_image_base64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_image-base64"] = t
		}
	}

	if v, ok := d.GetOk("allow_routing"); ok {
		t, err := expandObjectFirewallAddressAllowRouting(d, v, "allow_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-routing"] = t
		}
	}

	if v, ok := d.GetOk("associated_interface"); ok {
		t, err := expandObjectFirewallAddressAssociatedInterface(d, v, "associated_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["associated-interface"] = t
		}
	}

	if v, ok := d.GetOk("cache_ttl"); ok {
		t, err := expandObjectFirewallAddressCacheTtl(d, v, "cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("clearpass_spt"); ok {
		t, err := expandObjectFirewallAddressClearpassSpt(d, v, "clearpass_spt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clearpass-spt"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok {
		t, err := expandObjectFirewallAddressColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectFirewallAddressComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("country"); ok {
		t, err := expandObjectFirewallAddressCountry(d, v, "country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok {
		t, err := expandObjectFirewallAddressDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("end_ip"); ok {
		t, err := expandObjectFirewallAddressEndIp(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("end_mac"); ok {
		t, err := expandObjectFirewallAddressEndMac(d, v, "end_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-mac"] = t
		}
	}

	if v, ok := d.GetOk("epg_name"); ok {
		t, err := expandObjectFirewallAddressEpgName(d, v, "epg_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["epg-name"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok {
		t, err := expandObjectFirewallAddressFilter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok {
		t, err := expandObjectFirewallAddressFqdn(d, v, "fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("fsso_group"); ok {
		t, err := expandObjectFirewallAddressFssoGroup(d, v, "fsso_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-group"] = t
		}
	}

	if v, ok := d.GetOk("global_object"); ok {
		t, err := expandObjectFirewallAddressGlobalObject(d, v, "global_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-object"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandObjectFirewallAddressInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("list"); ok {
		t, err := expandObjectFirewallAddressList(d, v, "list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["list"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectFirewallAddressName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("obj_id"); ok {
		t, err := expandObjectFirewallAddressObjId(d, v, "obj_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-id"] = t
		}
	}

	if v, ok := d.GetOk("obj_tag"); ok {
		t, err := expandObjectFirewallAddressObjTag(d, v, "obj_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-tag"] = t
		}
	}

	if v, ok := d.GetOk("obj_type"); ok {
		t, err := expandObjectFirewallAddressObjType(d, v, "obj_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-type"] = t
		}
	}

	if v, ok := d.GetOk("organization"); ok {
		t, err := expandObjectFirewallAddressOrganization(d, v, "organization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["organization"] = t
		}
	}

	if v, ok := d.GetOk("policy_group"); ok {
		t, err := expandObjectFirewallAddressPolicyGroup(d, v, "policy_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-group"] = t
		}
	}

	if v, ok := d.GetOk("sdn"); ok {
		t, err := expandObjectFirewallAddressSdn(d, v, "sdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn"] = t
		}
	}

	if v, ok := d.GetOk("sdn_addr_type"); ok {
		t, err := expandObjectFirewallAddressSdnAddrType(d, v, "sdn_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("sdn_tag"); ok {
		t, err := expandObjectFirewallAddressSdnTag(d, v, "sdn_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-tag"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok {
		t, err := expandObjectFirewallAddressStartIp(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	if v, ok := d.GetOk("start_mac"); ok {
		t, err := expandObjectFirewallAddressStartMac(d, v, "start_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-mac"] = t
		}
	}

	if v, ok := d.GetOk("sub_type"); ok {
		t, err := expandObjectFirewallAddressSubType(d, v, "sub_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sub-type"] = t
		}
	}

	if v, ok := d.GetOk("subnet"); ok {
		t, err := expandObjectFirewallAddressSubnet(d, v, "subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet"] = t
		}
	}

	if v, ok := d.GetOk("subnet_name"); ok {
		t, err := expandObjectFirewallAddressSubnetName(d, v, "subnet_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-name"] = t
		}
	}

	if v, ok := d.GetOk("tagging"); ok {
		t, err := expandObjectFirewallAddressTagging(d, v, "tagging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tagging"] = t
		}
	}

	if v, ok := d.GetOk("tenant"); ok {
		t, err := expandObjectFirewallAddressTenant(d, v, "tenant")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandObjectFirewallAddressType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandObjectFirewallAddressUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("wildcard"); ok {
		t, err := expandObjectFirewallAddressWildcard(d, v, "wildcard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard"] = t
		}
	}

	if v, ok := d.GetOk("wildcard_fqdn"); ok {
		t, err := expandObjectFirewallAddressWildcardFqdn(d, v, "wildcard_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard-fqdn"] = t
		}
	}

	return &obj, nil
}
