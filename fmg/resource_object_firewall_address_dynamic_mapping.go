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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallAddressDynamicMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallAddressDynamicMappingCreate,
		Read:   resourceObjectFirewallAddressDynamicMappingRead,
		Update: resourceObjectFirewallAddressDynamicMappingUpdate,
		Delete: resourceObjectFirewallAddressDynamicMappingDelete,

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
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"_image_base64": &schema.Schema{
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
			"allow_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"clearpass_spt": &schema.Schema{
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
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dirty": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"epg_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fsso_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"global_object": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hw_model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hw_vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"macaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			},
			"obj_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"organization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"os": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pattern_end": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pattern_start": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"policy_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"route_tag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sdn_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdn_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso_attribute_value": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sub_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sw_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag_detection_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wildcard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wildcard_fqdn": &schema.Schema{
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

func resourceObjectFirewallAddressDynamicMappingCreate(d *schema.ResourceData, m interface{}) error {
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

	address := d.Get("address").(string)
	paradict["address"] = address

	obj, err := getObjectObjectFirewallAddressDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAddressDynamicMapping resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFirewallAddressDynamicMapping(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAddressDynamicMapping resource: %v", err)
	}

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectFirewallAddressDynamicMappingRead(d, m)
}

func resourceObjectFirewallAddressDynamicMappingUpdate(d *schema.ResourceData, m interface{}) error {
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

	address := d.Get("address").(string)
	paradict["address"] = address

	obj, err := getObjectObjectFirewallAddressDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAddressDynamicMapping resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallAddressDynamicMapping(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAddressDynamicMapping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectFirewallAddressDynamicMappingRead(d, m)
}

func resourceObjectFirewallAddressDynamicMappingDelete(d *schema.ResourceData, m interface{}) error {
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

	address := d.Get("address").(string)
	paradict["address"] = address

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallAddressDynamicMapping(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallAddressDynamicMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallAddressDynamicMappingRead(d *schema.ResourceData, m interface{}) error {
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

	address := d.Get("address").(string)
	if address == "" {
		address = importOptionChecking(m.(*FortiClient).Cfg, "address")
		if address == "" {
			return fmt.Errorf("Parameter address is missing")
		}
		if err = d.Set("address", address); err != nil {
			return fmt.Errorf("Error set params address: %v", err)
		}
	}
	if mkey, err = checkScopeId(mkey); err != nil {
		return fmt.Errorf("Error set ID : %v", err)
	}
	paradict["address"] = address

	o, err := c.ReadObjectFirewallAddressDynamicMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAddressDynamicMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallAddressDynamicMapping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAddressDynamicMapping resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallAddressDynamicMappingImageBase642edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingScope2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallAddressDynamicMappingScopeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallAddressDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectFirewallAddressDynamicMappingScopeVdom2edl(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectFirewallAddressDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallAddressDynamicMappingScopeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingScopeVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingAllowRouting2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingAssociatedInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingCacheTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingClearpassSpt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingColor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingCountry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFirewallAddressDynamicMappingDirty2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingEndIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingEndMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingEpgName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingFabricObject2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingFssoGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFirewallAddressDynamicMappingGlobalObject2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingHwModel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingHwVendor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingMacaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallAddressDynamicMappingNodeIpOnly2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingObjId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingObjTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingObjType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingOrganization2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingOs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingPatternEnd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingPatternStart2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingPolicyGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingRouteTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingSdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFirewallAddressDynamicMappingSdnAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingSdnTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingSsoAttributeValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallAddressDynamicMappingStartIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingStartMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingSubType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingSubnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFirewallAddressDynamicMappingSubnetName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingSwVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingTagDetectionLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingTagType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingTags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFirewallAddressDynamicMappingTenant2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingUuid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingVisibility2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddressDynamicMappingWildcard2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFirewallAddressDynamicMappingWildcardFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallAddressDynamicMapping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_image_base64", flattenObjectFirewallAddressDynamicMappingImageBase642edl(o["_image-base64"], d, "_image_base64")); err != nil {
		if vv, ok := fortiAPIPatch(o["_image-base64"], "ObjectFirewallAddressDynamicMapping-ImageBase64"); ok {
			if err = d.Set("_image_base64", vv); err != nil {
				return fmt.Errorf("Error reading _image_base64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _image_base64: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("_scope", flattenObjectFirewallAddressDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["_scope"], "ObjectFirewallAddressDynamicMapping-Scope"); ok {
				if err = d.Set("_scope", vv); err != nil {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading _scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("_scope"); ok {
			if err = d.Set("_scope", flattenObjectFirewallAddressDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["_scope"], "ObjectFirewallAddressDynamicMapping-Scope"); ok {
					if err = d.Set("_scope", vv); err != nil {
						return fmt.Errorf("Error reading _scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			}
		}
	}

	if err = d.Set("allow_routing", flattenObjectFirewallAddressDynamicMappingAllowRouting2edl(o["allow-routing"], d, "allow_routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-routing"], "ObjectFirewallAddressDynamicMapping-AllowRouting"); ok {
			if err = d.Set("allow_routing", vv); err != nil {
				return fmt.Errorf("Error reading allow_routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_routing: %v", err)
		}
	}

	if err = d.Set("associated_interface", flattenObjectFirewallAddressDynamicMappingAssociatedInterface2edl(o["associated-interface"], d, "associated_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["associated-interface"], "ObjectFirewallAddressDynamicMapping-AssociatedInterface"); ok {
			if err = d.Set("associated_interface", vv); err != nil {
				return fmt.Errorf("Error reading associated_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading associated_interface: %v", err)
		}
	}

	if err = d.Set("cache_ttl", flattenObjectFirewallAddressDynamicMappingCacheTtl2edl(o["cache-ttl"], d, "cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-ttl"], "ObjectFirewallAddressDynamicMapping-CacheTtl"); ok {
			if err = d.Set("cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_ttl: %v", err)
		}
	}

	if err = d.Set("clearpass_spt", flattenObjectFirewallAddressDynamicMappingClearpassSpt2edl(o["clearpass-spt"], d, "clearpass_spt")); err != nil {
		if vv, ok := fortiAPIPatch(o["clearpass-spt"], "ObjectFirewallAddressDynamicMapping-ClearpassSpt"); ok {
			if err = d.Set("clearpass_spt", vv); err != nil {
				return fmt.Errorf("Error reading clearpass_spt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clearpass_spt: %v", err)
		}
	}

	if err = d.Set("color", flattenObjectFirewallAddressDynamicMappingColor2edl(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "ObjectFirewallAddressDynamicMapping-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectFirewallAddressDynamicMappingComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectFirewallAddressDynamicMapping-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("country", flattenObjectFirewallAddressDynamicMappingCountry2edl(o["country"], d, "country")); err != nil {
		if vv, ok := fortiAPIPatch(o["country"], "ObjectFirewallAddressDynamicMapping-Country"); ok {
			if err = d.Set("country", vv); err != nil {
				return fmt.Errorf("Error reading country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if err = d.Set("dirty", flattenObjectFirewallAddressDynamicMappingDirty2edl(o["dirty"], d, "dirty")); err != nil {
		if vv, ok := fortiAPIPatch(o["dirty"], "ObjectFirewallAddressDynamicMapping-Dirty"); ok {
			if err = d.Set("dirty", vv); err != nil {
				return fmt.Errorf("Error reading dirty: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dirty: %v", err)
		}
	}

	if err = d.Set("end_ip", flattenObjectFirewallAddressDynamicMappingEndIp2edl(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "ObjectFirewallAddressDynamicMapping-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("end_mac", flattenObjectFirewallAddressDynamicMappingEndMac2edl(o["end-mac"], d, "end_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-mac"], "ObjectFirewallAddressDynamicMapping-EndMac"); ok {
			if err = d.Set("end_mac", vv); err != nil {
				return fmt.Errorf("Error reading end_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_mac: %v", err)
		}
	}

	if err = d.Set("epg_name", flattenObjectFirewallAddressDynamicMappingEpgName2edl(o["epg-name"], d, "epg_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["epg-name"], "ObjectFirewallAddressDynamicMapping-EpgName"); ok {
			if err = d.Set("epg_name", vv); err != nil {
				return fmt.Errorf("Error reading epg_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading epg_name: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenObjectFirewallAddressDynamicMappingFabricObject2edl(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "ObjectFirewallAddressDynamicMapping-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("filter", flattenObjectFirewallAddressDynamicMappingFilter2edl(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "ObjectFirewallAddressDynamicMapping-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenObjectFirewallAddressDynamicMappingFqdn2edl(o["fqdn"], d, "fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn"], "ObjectFirewallAddressDynamicMapping-Fqdn"); ok {
			if err = d.Set("fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("fsso_group", flattenObjectFirewallAddressDynamicMappingFssoGroup2edl(o["fsso-group"], d, "fsso_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-group"], "ObjectFirewallAddressDynamicMapping-FssoGroup"); ok {
			if err = d.Set("fsso_group", vv); err != nil {
				return fmt.Errorf("Error reading fsso_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_group: %v", err)
		}
	}

	if err = d.Set("global_object", flattenObjectFirewallAddressDynamicMappingGlobalObject2edl(o["global-object"], d, "global_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-object"], "ObjectFirewallAddressDynamicMapping-GlobalObject"); ok {
			if err = d.Set("global_object", vv); err != nil {
				return fmt.Errorf("Error reading global_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_object: %v", err)
		}
	}

	if err = d.Set("hw_model", flattenObjectFirewallAddressDynamicMappingHwModel2edl(o["hw-model"], d, "hw_model")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-model"], "ObjectFirewallAddressDynamicMapping-HwModel"); ok {
			if err = d.Set("hw_model", vv); err != nil {
				return fmt.Errorf("Error reading hw_model: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_model: %v", err)
		}
	}

	if err = d.Set("hw_vendor", flattenObjectFirewallAddressDynamicMappingHwVendor2edl(o["hw-vendor"], d, "hw_vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-vendor"], "ObjectFirewallAddressDynamicMapping-HwVendor"); ok {
			if err = d.Set("hw_vendor", vv); err != nil {
				return fmt.Errorf("Error reading hw_vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_vendor: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectFirewallAddressDynamicMappingInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectFirewallAddressDynamicMapping-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("macaddr", flattenObjectFirewallAddressDynamicMappingMacaddr2edl(o["macaddr"], d, "macaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["macaddr"], "ObjectFirewallAddressDynamicMapping-Macaddr"); ok {
			if err = d.Set("macaddr", vv); err != nil {
				return fmt.Errorf("Error reading macaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading macaddr: %v", err)
		}
	}

	if err = d.Set("node_ip_only", flattenObjectFirewallAddressDynamicMappingNodeIpOnly2edl(o["node-ip-only"], d, "node_ip_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["node-ip-only"], "ObjectFirewallAddressDynamicMapping-NodeIpOnly"); ok {
			if err = d.Set("node_ip_only", vv); err != nil {
				return fmt.Errorf("Error reading node_ip_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading node_ip_only: %v", err)
		}
	}

	if err = d.Set("obj_id", flattenObjectFirewallAddressDynamicMappingObjId2edl(o["obj-id"], d, "obj_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["obj-id"], "ObjectFirewallAddressDynamicMapping-ObjId"); ok {
			if err = d.Set("obj_id", vv); err != nil {
				return fmt.Errorf("Error reading obj_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obj_id: %v", err)
		}
	}

	if err = d.Set("obj_tag", flattenObjectFirewallAddressDynamicMappingObjTag2edl(o["obj-tag"], d, "obj_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["obj-tag"], "ObjectFirewallAddressDynamicMapping-ObjTag"); ok {
			if err = d.Set("obj_tag", vv); err != nil {
				return fmt.Errorf("Error reading obj_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obj_tag: %v", err)
		}
	}

	if err = d.Set("obj_type", flattenObjectFirewallAddressDynamicMappingObjType2edl(o["obj-type"], d, "obj_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["obj-type"], "ObjectFirewallAddressDynamicMapping-ObjType"); ok {
			if err = d.Set("obj_type", vv); err != nil {
				return fmt.Errorf("Error reading obj_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obj_type: %v", err)
		}
	}

	if err = d.Set("organization", flattenObjectFirewallAddressDynamicMappingOrganization2edl(o["organization"], d, "organization")); err != nil {
		if vv, ok := fortiAPIPatch(o["organization"], "ObjectFirewallAddressDynamicMapping-Organization"); ok {
			if err = d.Set("organization", vv); err != nil {
				return fmt.Errorf("Error reading organization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading organization: %v", err)
		}
	}

	if err = d.Set("os", flattenObjectFirewallAddressDynamicMappingOs2edl(o["os"], d, "os")); err != nil {
		if vv, ok := fortiAPIPatch(o["os"], "ObjectFirewallAddressDynamicMapping-Os"); ok {
			if err = d.Set("os", vv); err != nil {
				return fmt.Errorf("Error reading os: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os: %v", err)
		}
	}

	if err = d.Set("pattern_end", flattenObjectFirewallAddressDynamicMappingPatternEnd2edl(o["pattern-end"], d, "pattern_end")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern-end"], "ObjectFirewallAddressDynamicMapping-PatternEnd"); ok {
			if err = d.Set("pattern_end", vv); err != nil {
				return fmt.Errorf("Error reading pattern_end: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern_end: %v", err)
		}
	}

	if err = d.Set("pattern_start", flattenObjectFirewallAddressDynamicMappingPatternStart2edl(o["pattern-start"], d, "pattern_start")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern-start"], "ObjectFirewallAddressDynamicMapping-PatternStart"); ok {
			if err = d.Set("pattern_start", vv); err != nil {
				return fmt.Errorf("Error reading pattern_start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern_start: %v", err)
		}
	}

	if err = d.Set("policy_group", flattenObjectFirewallAddressDynamicMappingPolicyGroup2edl(o["policy-group"], d, "policy_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-group"], "ObjectFirewallAddressDynamicMapping-PolicyGroup"); ok {
			if err = d.Set("policy_group", vv); err != nil {
				return fmt.Errorf("Error reading policy_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_group: %v", err)
		}
	}

	if err = d.Set("route_tag", flattenObjectFirewallAddressDynamicMappingRouteTag2edl(o["route-tag"], d, "route_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-tag"], "ObjectFirewallAddressDynamicMapping-RouteTag"); ok {
			if err = d.Set("route_tag", vv); err != nil {
				return fmt.Errorf("Error reading route_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_tag: %v", err)
		}
	}

	if err = d.Set("sdn", flattenObjectFirewallAddressDynamicMappingSdn2edl(o["sdn"], d, "sdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn"], "ObjectFirewallAddressDynamicMapping-Sdn"); ok {
			if err = d.Set("sdn", vv); err != nil {
				return fmt.Errorf("Error reading sdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn: %v", err)
		}
	}

	if err = d.Set("sdn_addr_type", flattenObjectFirewallAddressDynamicMappingSdnAddrType2edl(o["sdn-addr-type"], d, "sdn_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn-addr-type"], "ObjectFirewallAddressDynamicMapping-SdnAddrType"); ok {
			if err = d.Set("sdn_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading sdn_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn_addr_type: %v", err)
		}
	}

	if err = d.Set("sdn_tag", flattenObjectFirewallAddressDynamicMappingSdnTag2edl(o["sdn-tag"], d, "sdn_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn-tag"], "ObjectFirewallAddressDynamicMapping-SdnTag"); ok {
			if err = d.Set("sdn_tag", vv); err != nil {
				return fmt.Errorf("Error reading sdn_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn_tag: %v", err)
		}
	}

	if err = d.Set("sso_attribute_value", flattenObjectFirewallAddressDynamicMappingSsoAttributeValue2edl(o["sso-attribute-value"], d, "sso_attribute_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-attribute-value"], "ObjectFirewallAddressDynamicMapping-SsoAttributeValue"); ok {
			if err = d.Set("sso_attribute_value", vv); err != nil {
				return fmt.Errorf("Error reading sso_attribute_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_attribute_value: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenObjectFirewallAddressDynamicMappingStartIp2edl(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "ObjectFirewallAddressDynamicMapping-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	if err = d.Set("start_mac", flattenObjectFirewallAddressDynamicMappingStartMac2edl(o["start-mac"], d, "start_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-mac"], "ObjectFirewallAddressDynamicMapping-StartMac"); ok {
			if err = d.Set("start_mac", vv); err != nil {
				return fmt.Errorf("Error reading start_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_mac: %v", err)
		}
	}

	if err = d.Set("sub_type", flattenObjectFirewallAddressDynamicMappingSubType2edl(o["sub-type"], d, "sub_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["sub-type"], "ObjectFirewallAddressDynamicMapping-SubType"); ok {
			if err = d.Set("sub_type", vv); err != nil {
				return fmt.Errorf("Error reading sub_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sub_type: %v", err)
		}
	}

	if err = d.Set("subnet", flattenObjectFirewallAddressDynamicMappingSubnet2edl(o["subnet"], d, "subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet"], "ObjectFirewallAddressDynamicMapping-Subnet"); ok {
			if err = d.Set("subnet", vv); err != nil {
				return fmt.Errorf("Error reading subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet: %v", err)
		}
	}

	if err = d.Set("subnet_name", flattenObjectFirewallAddressDynamicMappingSubnetName2edl(o["subnet-name"], d, "subnet_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet-name"], "ObjectFirewallAddressDynamicMapping-SubnetName"); ok {
			if err = d.Set("subnet_name", vv); err != nil {
				return fmt.Errorf("Error reading subnet_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet_name: %v", err)
		}
	}

	if err = d.Set("sw_version", flattenObjectFirewallAddressDynamicMappingSwVersion2edl(o["sw-version"], d, "sw_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["sw-version"], "ObjectFirewallAddressDynamicMapping-SwVersion"); ok {
			if err = d.Set("sw_version", vv); err != nil {
				return fmt.Errorf("Error reading sw_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sw_version: %v", err)
		}
	}

	if err = d.Set("tag_detection_level", flattenObjectFirewallAddressDynamicMappingTagDetectionLevel2edl(o["tag-detection-level"], d, "tag_detection_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag-detection-level"], "ObjectFirewallAddressDynamicMapping-TagDetectionLevel"); ok {
			if err = d.Set("tag_detection_level", vv); err != nil {
				return fmt.Errorf("Error reading tag_detection_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag_detection_level: %v", err)
		}
	}

	if err = d.Set("tag_type", flattenObjectFirewallAddressDynamicMappingTagType2edl(o["tag-type"], d, "tag_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag-type"], "ObjectFirewallAddressDynamicMapping-TagType"); ok {
			if err = d.Set("tag_type", vv); err != nil {
				return fmt.Errorf("Error reading tag_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag_type: %v", err)
		}
	}

	if err = d.Set("tags", flattenObjectFirewallAddressDynamicMappingTags2edl(o["tags"], d, "tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["tags"], "ObjectFirewallAddressDynamicMapping-Tags"); ok {
			if err = d.Set("tags", vv); err != nil {
				return fmt.Errorf("Error reading tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tags: %v", err)
		}
	}

	if err = d.Set("tenant", flattenObjectFirewallAddressDynamicMappingTenant2edl(o["tenant"], d, "tenant")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenant"], "ObjectFirewallAddressDynamicMapping-Tenant"); ok {
			if err = d.Set("tenant", vv); err != nil {
				return fmt.Errorf("Error reading tenant: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenant: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFirewallAddressDynamicMappingType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallAddressDynamicMapping-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("url", flattenObjectFirewallAddressDynamicMappingUrl2edl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "ObjectFirewallAddressDynamicMapping-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectFirewallAddressDynamicMappingUuid2edl(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectFirewallAddressDynamicMapping-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("visibility", flattenObjectFirewallAddressDynamicMappingVisibility2edl(o["visibility"], d, "visibility")); err != nil {
		if vv, ok := fortiAPIPatch(o["visibility"], "ObjectFirewallAddressDynamicMapping-Visibility"); ok {
			if err = d.Set("visibility", vv); err != nil {
				return fmt.Errorf("Error reading visibility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading visibility: %v", err)
		}
	}

	if err = d.Set("wildcard", flattenObjectFirewallAddressDynamicMappingWildcard2edl(o["wildcard"], d, "wildcard")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard"], "ObjectFirewallAddressDynamicMapping-Wildcard"); ok {
			if err = d.Set("wildcard", vv); err != nil {
				return fmt.Errorf("Error reading wildcard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}

	if err = d.Set("wildcard_fqdn", flattenObjectFirewallAddressDynamicMappingWildcardFqdn2edl(o["wildcard-fqdn"], d, "wildcard_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard-fqdn"], "ObjectFirewallAddressDynamicMapping-WildcardFqdn"); ok {
			if err = d.Set("wildcard_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading wildcard_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard_fqdn: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallAddressDynamicMappingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallAddressDynamicMappingImageBase642edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallAddressDynamicMappingScopeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectFirewallAddressDynamicMappingScopeVdom2edl(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAddressDynamicMappingScopeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingScopeVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingAllowRouting2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingAssociatedInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingCacheTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingClearpassSpt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingColor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingCountry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallAddressDynamicMappingDirty2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingEndIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingEndMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingEpgName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingFabricObject2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingFssoGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallAddressDynamicMappingGlobalObject2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingHwModel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingHwVendor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingMacaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallAddressDynamicMappingNodeIpOnly2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingObjId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingObjTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingObjType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingOrganization2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingOs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingPatternEnd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingPatternStart2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingPolicyGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingRouteTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingSdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallAddressDynamicMappingSdnAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingSdnTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingSsoAttributeValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallAddressDynamicMappingStartIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingStartMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingSubType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingSubnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallAddressDynamicMappingSubnetName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingSwVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingTagDetectionLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingTagType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingTags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallAddressDynamicMappingTenant2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingUuid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingVisibility2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddressDynamicMappingWildcard2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallAddressDynamicMappingWildcardFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallAddressDynamicMapping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_image_base64"); ok || d.HasChange("_image_base64") {
		t, err := expandObjectFirewallAddressDynamicMappingImageBase642edl(d, v, "_image_base64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_image-base64"] = t
		}
	}

	if v, ok := d.GetOk("_scope"); ok || d.HasChange("_scope") {
		t, err := expandObjectFirewallAddressDynamicMappingScope2edl(d, v, "_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_scope"] = t
		}
	}

	if v, ok := d.GetOk("allow_routing"); ok || d.HasChange("allow_routing") {
		t, err := expandObjectFirewallAddressDynamicMappingAllowRouting2edl(d, v, "allow_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-routing"] = t
		}
	}

	if v, ok := d.GetOk("associated_interface"); ok || d.HasChange("associated_interface") {
		t, err := expandObjectFirewallAddressDynamicMappingAssociatedInterface2edl(d, v, "associated_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["associated-interface"] = t
		}
	}

	if v, ok := d.GetOk("cache_ttl"); ok || d.HasChange("cache_ttl") {
		t, err := expandObjectFirewallAddressDynamicMappingCacheTtl2edl(d, v, "cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("clearpass_spt"); ok || d.HasChange("clearpass_spt") {
		t, err := expandObjectFirewallAddressDynamicMappingClearpassSpt2edl(d, v, "clearpass_spt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clearpass-spt"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandObjectFirewallAddressDynamicMappingColor2edl(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectFirewallAddressDynamicMappingComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("country"); ok || d.HasChange("country") {
		t, err := expandObjectFirewallAddressDynamicMappingCountry2edl(d, v, "country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("dirty"); ok || d.HasChange("dirty") {
		t, err := expandObjectFirewallAddressDynamicMappingDirty2edl(d, v, "dirty")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dirty"] = t
		}
	}

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandObjectFirewallAddressDynamicMappingEndIp2edl(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("end_mac"); ok || d.HasChange("end_mac") {
		t, err := expandObjectFirewallAddressDynamicMappingEndMac2edl(d, v, "end_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-mac"] = t
		}
	}

	if v, ok := d.GetOk("epg_name"); ok || d.HasChange("epg_name") {
		t, err := expandObjectFirewallAddressDynamicMappingEpgName2edl(d, v, "epg_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["epg-name"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandObjectFirewallAddressDynamicMappingFabricObject2edl(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandObjectFirewallAddressDynamicMappingFilter2edl(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok || d.HasChange("fqdn") {
		t, err := expandObjectFirewallAddressDynamicMappingFqdn2edl(d, v, "fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("fsso_group"); ok || d.HasChange("fsso_group") {
		t, err := expandObjectFirewallAddressDynamicMappingFssoGroup2edl(d, v, "fsso_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-group"] = t
		}
	}

	if v, ok := d.GetOk("global_object"); ok || d.HasChange("global_object") {
		t, err := expandObjectFirewallAddressDynamicMappingGlobalObject2edl(d, v, "global_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-object"] = t
		}
	}

	if v, ok := d.GetOk("hw_model"); ok || d.HasChange("hw_model") {
		t, err := expandObjectFirewallAddressDynamicMappingHwModel2edl(d, v, "hw_model")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-model"] = t
		}
	}

	if v, ok := d.GetOk("hw_vendor"); ok || d.HasChange("hw_vendor") {
		t, err := expandObjectFirewallAddressDynamicMappingHwVendor2edl(d, v, "hw_vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-vendor"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectFirewallAddressDynamicMappingInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("macaddr"); ok || d.HasChange("macaddr") {
		t, err := expandObjectFirewallAddressDynamicMappingMacaddr2edl(d, v, "macaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macaddr"] = t
		}
	}

	if v, ok := d.GetOk("node_ip_only"); ok || d.HasChange("node_ip_only") {
		t, err := expandObjectFirewallAddressDynamicMappingNodeIpOnly2edl(d, v, "node_ip_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["node-ip-only"] = t
		}
	}

	if v, ok := d.GetOk("obj_id"); ok || d.HasChange("obj_id") {
		t, err := expandObjectFirewallAddressDynamicMappingObjId2edl(d, v, "obj_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-id"] = t
		}
	}

	if v, ok := d.GetOk("obj_tag"); ok || d.HasChange("obj_tag") {
		t, err := expandObjectFirewallAddressDynamicMappingObjTag2edl(d, v, "obj_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-tag"] = t
		}
	}

	if v, ok := d.GetOk("obj_type"); ok || d.HasChange("obj_type") {
		t, err := expandObjectFirewallAddressDynamicMappingObjType2edl(d, v, "obj_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-type"] = t
		}
	}

	if v, ok := d.GetOk("organization"); ok || d.HasChange("organization") {
		t, err := expandObjectFirewallAddressDynamicMappingOrganization2edl(d, v, "organization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["organization"] = t
		}
	}

	if v, ok := d.GetOk("os"); ok || d.HasChange("os") {
		t, err := expandObjectFirewallAddressDynamicMappingOs2edl(d, v, "os")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os"] = t
		}
	}

	if v, ok := d.GetOk("pattern_end"); ok || d.HasChange("pattern_end") {
		t, err := expandObjectFirewallAddressDynamicMappingPatternEnd2edl(d, v, "pattern_end")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern-end"] = t
		}
	}

	if v, ok := d.GetOk("pattern_start"); ok || d.HasChange("pattern_start") {
		t, err := expandObjectFirewallAddressDynamicMappingPatternStart2edl(d, v, "pattern_start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern-start"] = t
		}
	}

	if v, ok := d.GetOk("policy_group"); ok || d.HasChange("policy_group") {
		t, err := expandObjectFirewallAddressDynamicMappingPolicyGroup2edl(d, v, "policy_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-group"] = t
		}
	}

	if v, ok := d.GetOk("route_tag"); ok || d.HasChange("route_tag") {
		t, err := expandObjectFirewallAddressDynamicMappingRouteTag2edl(d, v, "route_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-tag"] = t
		}
	}

	if v, ok := d.GetOk("sdn"); ok || d.HasChange("sdn") {
		t, err := expandObjectFirewallAddressDynamicMappingSdn2edl(d, v, "sdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn"] = t
		}
	}

	if v, ok := d.GetOk("sdn_addr_type"); ok || d.HasChange("sdn_addr_type") {
		t, err := expandObjectFirewallAddressDynamicMappingSdnAddrType2edl(d, v, "sdn_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("sdn_tag"); ok || d.HasChange("sdn_tag") {
		t, err := expandObjectFirewallAddressDynamicMappingSdnTag2edl(d, v, "sdn_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-tag"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute_value"); ok || d.HasChange("sso_attribute_value") {
		t, err := expandObjectFirewallAddressDynamicMappingSsoAttributeValue2edl(d, v, "sso_attribute_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute-value"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandObjectFirewallAddressDynamicMappingStartIp2edl(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	if v, ok := d.GetOk("start_mac"); ok || d.HasChange("start_mac") {
		t, err := expandObjectFirewallAddressDynamicMappingStartMac2edl(d, v, "start_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-mac"] = t
		}
	}

	if v, ok := d.GetOk("sub_type"); ok || d.HasChange("sub_type") {
		t, err := expandObjectFirewallAddressDynamicMappingSubType2edl(d, v, "sub_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sub-type"] = t
		}
	}

	if v, ok := d.GetOk("subnet"); ok || d.HasChange("subnet") {
		t, err := expandObjectFirewallAddressDynamicMappingSubnet2edl(d, v, "subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet"] = t
		}
	}

	if v, ok := d.GetOk("subnet_name"); ok || d.HasChange("subnet_name") {
		t, err := expandObjectFirewallAddressDynamicMappingSubnetName2edl(d, v, "subnet_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-name"] = t
		}
	}

	if v, ok := d.GetOk("sw_version"); ok || d.HasChange("sw_version") {
		t, err := expandObjectFirewallAddressDynamicMappingSwVersion2edl(d, v, "sw_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sw-version"] = t
		}
	}

	if v, ok := d.GetOk("tag_detection_level"); ok || d.HasChange("tag_detection_level") {
		t, err := expandObjectFirewallAddressDynamicMappingTagDetectionLevel2edl(d, v, "tag_detection_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag-detection-level"] = t
		}
	}

	if v, ok := d.GetOk("tag_type"); ok || d.HasChange("tag_type") {
		t, err := expandObjectFirewallAddressDynamicMappingTagType2edl(d, v, "tag_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag-type"] = t
		}
	}

	if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
		t, err := expandObjectFirewallAddressDynamicMappingTags2edl(d, v, "tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tags"] = t
		}
	}

	if v, ok := d.GetOk("tenant"); ok || d.HasChange("tenant") {
		t, err := expandObjectFirewallAddressDynamicMappingTenant2edl(d, v, "tenant")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFirewallAddressDynamicMappingType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandObjectFirewallAddressDynamicMappingUrl2edl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandObjectFirewallAddressDynamicMappingUuid2edl(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("visibility"); ok || d.HasChange("visibility") {
		t, err := expandObjectFirewallAddressDynamicMappingVisibility2edl(d, v, "visibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["visibility"] = t
		}
	}

	if v, ok := d.GetOk("wildcard"); ok || d.HasChange("wildcard") {
		t, err := expandObjectFirewallAddressDynamicMappingWildcard2edl(d, v, "wildcard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard"] = t
		}
	}

	if v, ok := d.GetOk("wildcard_fqdn"); ok || d.HasChange("wildcard_fqdn") {
		t, err := expandObjectFirewallAddressDynamicMappingWildcardFqdn2edl(d, v, "wildcard_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard-fqdn"] = t
		}
	}

	return &obj, nil
}
