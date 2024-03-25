// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure route maps.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectRouterRouteMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectRouterRouteMapCreate,
		Read:   resourceObjectRouterRouteMapRead,
		Update: resourceObjectRouterRouteMapUpdate,
		Delete: resourceObjectRouterRouteMapDelete,

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
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"match_as_path": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_community": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_community_exact": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_extcommunity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_extcommunity_exact": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_flags": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"match_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_ip_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_ip6_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_metric": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_origin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_route_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_vrf": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_aggregator_as": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_aggregator_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_aspath": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"set_aspath_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_atomic_aggregate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_community": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"set_community_additive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_community_delete": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_dampening_max_suppress": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_dampening_reachability_half_life": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_dampening_reuse": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_dampening_suppress": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_dampening_unreachability_half_life": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_extcommunity_rt": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"set_extcommunity_soo": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"set_flags": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_ip_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_ip_prefsrc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_ip6_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_ip6_nexthop_local": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_local_preference": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_metric": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_origin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_originator_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_route_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_vpnv4_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_vpnv6_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_vpnv6_nexthop_local": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_weight": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectRouterRouteMapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectRouterRouteMap(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterRouteMap resource while getting object: %v", err)
	}

	_, err = c.CreateObjectRouterRouteMap(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterRouteMap resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectRouterRouteMapRead(d, m)
}

func resourceObjectRouterRouteMapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectRouterRouteMap(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterRouteMap resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectRouterRouteMap(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterRouteMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectRouterRouteMapRead(d, m)
}

func resourceObjectRouterRouteMapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectRouterRouteMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectRouterRouteMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectRouterRouteMapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectRouterRouteMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterRouteMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectRouterRouteMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterRouteMap resource from API: %v", err)
	}
	return nil
}

func flattenObjectRouterRouteMapComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectRouterRouteMapRuleAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectRouterRouteMapRuleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_as_path"
		if _, ok := i["match-as-path"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchAsPath(i["match-as-path"], d, pre_append)
			tmp["match_as_path"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchAsPath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_community"
		if _, ok := i["match-community"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchCommunity(i["match-community"], d, pre_append)
			tmp["match_community"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchCommunity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_community_exact"
		if _, ok := i["match-community-exact"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchCommunityExact(i["match-community-exact"], d, pre_append)
			tmp["match_community_exact"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchCommunityExact")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_extcommunity"
		if _, ok := i["match-extcommunity"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchExtcommunity(i["match-extcommunity"], d, pre_append)
			tmp["match_extcommunity"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchExtcommunity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_extcommunity_exact"
		if _, ok := i["match-extcommunity-exact"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchExtcommunityExact(i["match-extcommunity-exact"], d, pre_append)
			tmp["match_extcommunity_exact"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchExtcommunityExact")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_flags"
		if _, ok := i["match-flags"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchFlags(i["match-flags"], d, pre_append)
			tmp["match_flags"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchFlags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_interface"
		if _, ok := i["match-interface"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchInterface(i["match-interface"], d, pre_append)
			tmp["match_interface"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip_address"
		if _, ok := i["match-ip-address"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchIpAddress(i["match-ip-address"], d, pre_append)
			tmp["match_ip_address"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchIpAddress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip_nexthop"
		if _, ok := i["match-ip-nexthop"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchIpNexthop(i["match-ip-nexthop"], d, pre_append)
			tmp["match_ip_nexthop"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchIpNexthop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip6_address"
		if _, ok := i["match-ip6-address"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchIp6Address(i["match-ip6-address"], d, pre_append)
			tmp["match_ip6_address"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchIp6Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip6_nexthop"
		if _, ok := i["match-ip6-nexthop"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchIp6Nexthop(i["match-ip6-nexthop"], d, pre_append)
			tmp["match_ip6_nexthop"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchIp6Nexthop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_metric"
		if _, ok := i["match-metric"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchMetric(i["match-metric"], d, pre_append)
			tmp["match_metric"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchMetric")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_origin"
		if _, ok := i["match-origin"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchOrigin(i["match-origin"], d, pre_append)
			tmp["match_origin"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchOrigin")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_route_type"
		if _, ok := i["match-route-type"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchRouteType(i["match-route-type"], d, pre_append)
			tmp["match_route_type"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchRouteType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_tag"
		if _, ok := i["match-tag"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchTag(i["match-tag"], d, pre_append)
			tmp["match_tag"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_vrf"
		if _, ok := i["match-vrf"]; ok {
			v := flattenObjectRouterRouteMapRuleMatchVrf(i["match-vrf"], d, pre_append)
			tmp["match_vrf"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-MatchVrf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aggregator_as"
		if _, ok := i["set-aggregator-as"]; ok {
			v := flattenObjectRouterRouteMapRuleSetAggregatorAs(i["set-aggregator-as"], d, pre_append)
			tmp["set_aggregator_as"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetAggregatorAs")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aggregator_ip"
		if _, ok := i["set-aggregator-ip"]; ok {
			v := flattenObjectRouterRouteMapRuleSetAggregatorIp(i["set-aggregator-ip"], d, pre_append)
			tmp["set_aggregator_ip"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetAggregatorIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aspath"
		if _, ok := i["set-aspath"]; ok {
			v := flattenObjectRouterRouteMapRuleSetAspath(i["set-aspath"], d, pre_append)
			tmp["set_aspath"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetAspath")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aspath_action"
		if _, ok := i["set-aspath-action"]; ok {
			v := flattenObjectRouterRouteMapRuleSetAspathAction(i["set-aspath-action"], d, pre_append)
			tmp["set_aspath_action"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetAspathAction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_atomic_aggregate"
		if _, ok := i["set-atomic-aggregate"]; ok {
			v := flattenObjectRouterRouteMapRuleSetAtomicAggregate(i["set-atomic-aggregate"], d, pre_append)
			tmp["set_atomic_aggregate"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetAtomicAggregate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community"
		if _, ok := i["set-community"]; ok {
			v := flattenObjectRouterRouteMapRuleSetCommunity(i["set-community"], d, pre_append)
			tmp["set_community"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetCommunity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community_additive"
		if _, ok := i["set-community-additive"]; ok {
			v := flattenObjectRouterRouteMapRuleSetCommunityAdditive(i["set-community-additive"], d, pre_append)
			tmp["set_community_additive"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetCommunityAdditive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community_delete"
		if _, ok := i["set-community-delete"]; ok {
			v := flattenObjectRouterRouteMapRuleSetCommunityDelete(i["set-community-delete"], d, pre_append)
			tmp["set_community_delete"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetCommunityDelete")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_max_suppress"
		if _, ok := i["set-dampening-max-suppress"]; ok {
			v := flattenObjectRouterRouteMapRuleSetDampeningMaxSuppress(i["set-dampening-max-suppress"], d, pre_append)
			tmp["set_dampening_max_suppress"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetDampeningMaxSuppress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_reachability_half_life"
		if _, ok := i["set-dampening-reachability-half-life"]; ok {
			v := flattenObjectRouterRouteMapRuleSetDampeningReachabilityHalfLife(i["set-dampening-reachability-half-life"], d, pre_append)
			tmp["set_dampening_reachability_half_life"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetDampeningReachabilityHalfLife")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_reuse"
		if _, ok := i["set-dampening-reuse"]; ok {
			v := flattenObjectRouterRouteMapRuleSetDampeningReuse(i["set-dampening-reuse"], d, pre_append)
			tmp["set_dampening_reuse"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetDampeningReuse")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_suppress"
		if _, ok := i["set-dampening-suppress"]; ok {
			v := flattenObjectRouterRouteMapRuleSetDampeningSuppress(i["set-dampening-suppress"], d, pre_append)
			tmp["set_dampening_suppress"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetDampeningSuppress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_unreachability_half_life"
		if _, ok := i["set-dampening-unreachability-half-life"]; ok {
			v := flattenObjectRouterRouteMapRuleSetDampeningUnreachabilityHalfLife(i["set-dampening-unreachability-half-life"], d, pre_append)
			tmp["set_dampening_unreachability_half_life"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetDampeningUnreachabilityHalfLife")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_extcommunity_rt"
		if _, ok := i["set-extcommunity-rt"]; ok {
			v := flattenObjectRouterRouteMapRuleSetExtcommunityRt(i["set-extcommunity-rt"], d, pre_append)
			tmp["set_extcommunity_rt"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetExtcommunityRt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_extcommunity_soo"
		if _, ok := i["set-extcommunity-soo"]; ok {
			v := flattenObjectRouterRouteMapRuleSetExtcommunitySoo(i["set-extcommunity-soo"], d, pre_append)
			tmp["set_extcommunity_soo"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetExtcommunitySoo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_flags"
		if _, ok := i["set-flags"]; ok {
			v := flattenObjectRouterRouteMapRuleSetFlags(i["set-flags"], d, pre_append)
			tmp["set_flags"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetFlags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip_nexthop"
		if _, ok := i["set-ip-nexthop"]; ok {
			v := flattenObjectRouterRouteMapRuleSetIpNexthop(i["set-ip-nexthop"], d, pre_append)
			tmp["set_ip_nexthop"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetIpNexthop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip_prefsrc"
		if _, ok := i["set-ip-prefsrc"]; ok {
			v := flattenObjectRouterRouteMapRuleSetIpPrefsrc(i["set-ip-prefsrc"], d, pre_append)
			tmp["set_ip_prefsrc"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetIpPrefsrc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip6_nexthop"
		if _, ok := i["set-ip6-nexthop"]; ok {
			v := flattenObjectRouterRouteMapRuleSetIp6Nexthop(i["set-ip6-nexthop"], d, pre_append)
			tmp["set_ip6_nexthop"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetIp6Nexthop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip6_nexthop_local"
		if _, ok := i["set-ip6-nexthop-local"]; ok {
			v := flattenObjectRouterRouteMapRuleSetIp6NexthopLocal(i["set-ip6-nexthop-local"], d, pre_append)
			tmp["set_ip6_nexthop_local"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetIp6NexthopLocal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_local_preference"
		if _, ok := i["set-local-preference"]; ok {
			v := flattenObjectRouterRouteMapRuleSetLocalPreference(i["set-local-preference"], d, pre_append)
			tmp["set_local_preference"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetLocalPreference")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_metric"
		if _, ok := i["set-metric"]; ok {
			v := flattenObjectRouterRouteMapRuleSetMetric(i["set-metric"], d, pre_append)
			tmp["set_metric"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetMetric")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_metric_type"
		if _, ok := i["set-metric-type"]; ok {
			v := flattenObjectRouterRouteMapRuleSetMetricType(i["set-metric-type"], d, pre_append)
			tmp["set_metric_type"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetMetricType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_origin"
		if _, ok := i["set-origin"]; ok {
			v := flattenObjectRouterRouteMapRuleSetOrigin(i["set-origin"], d, pre_append)
			tmp["set_origin"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetOrigin")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_originator_id"
		if _, ok := i["set-originator-id"]; ok {
			v := flattenObjectRouterRouteMapRuleSetOriginatorId(i["set-originator-id"], d, pre_append)
			tmp["set_originator_id"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetOriginatorId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_priority"
		if _, ok := i["set-priority"]; ok {
			v := flattenObjectRouterRouteMapRuleSetPriority(i["set-priority"], d, pre_append)
			tmp["set_priority"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_route_tag"
		if _, ok := i["set-route-tag"]; ok {
			v := flattenObjectRouterRouteMapRuleSetRouteTag(i["set-route-tag"], d, pre_append)
			tmp["set_route_tag"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetRouteTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_tag"
		if _, ok := i["set-tag"]; ok {
			v := flattenObjectRouterRouteMapRuleSetTag(i["set-tag"], d, pre_append)
			tmp["set_tag"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv4_nexthop"
		if _, ok := i["set-vpnv4-nexthop"]; ok {
			v := flattenObjectRouterRouteMapRuleSetVpnv4Nexthop(i["set-vpnv4-nexthop"], d, pre_append)
			tmp["set_vpnv4_nexthop"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetVpnv4Nexthop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv6_nexthop"
		if _, ok := i["set-vpnv6-nexthop"]; ok {
			v := flattenObjectRouterRouteMapRuleSetVpnv6Nexthop(i["set-vpnv6-nexthop"], d, pre_append)
			tmp["set_vpnv6_nexthop"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetVpnv6Nexthop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv6_nexthop_local"
		if _, ok := i["set-vpnv6-nexthop-local"]; ok {
			v := flattenObjectRouterRouteMapRuleSetVpnv6NexthopLocal(i["set-vpnv6-nexthop-local"], d, pre_append)
			tmp["set_vpnv6_nexthop_local"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetVpnv6NexthopLocal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_weight"
		if _, ok := i["set-weight"]; ok {
			v := flattenObjectRouterRouteMapRuleSetWeight(i["set-weight"], d, pre_append)
			tmp["set_weight"] = fortiAPISubPartPatch(v, "ObjectRouterRouteMap-Rule-SetWeight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectRouterRouteMapRuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchAsPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchCommunityExact(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchExtcommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchExtcommunityExact(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchIpNexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchIp6Nexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchOrigin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchRouteType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetAggregatorAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetAggregatorIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetAspath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectRouterRouteMapRuleSetAspathAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetAtomicAggregate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectRouterRouteMapRuleSetCommunityAdditive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetCommunityDelete(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetDampeningMaxSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetDampeningReachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetDampeningReuse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetDampeningSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetDampeningUnreachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetExtcommunityRt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectRouterRouteMapRuleSetExtcommunitySoo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectRouterRouteMapRuleSetFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetIpNexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetIpPrefsrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetIp6Nexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetIp6NexthopLocal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetLocalPreference(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetOrigin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetOriginatorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetRouteTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetVpnv4Nexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetVpnv6Nexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetVpnv6NexthopLocal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectRouterRouteMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comments", flattenObjectRouterRouteMapComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "ObjectRouterRouteMap-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectRouterRouteMapName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectRouterRouteMap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rule", flattenObjectRouterRouteMapRule(o["rule"], d, "rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["rule"], "ObjectRouterRouteMap-Rule"); ok {
				if err = d.Set("rule", vv); err != nil {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenObjectRouterRouteMapRule(o["rule"], d, "rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["rule"], "ObjectRouterRouteMap-Rule"); ok {
					if err = d.Set("rule", vv); err != nil {
						return fmt.Errorf("Error reading rule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectRouterRouteMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectRouterRouteMapComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectRouterRouteMapRuleAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectRouterRouteMapRuleId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_as_path"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-as-path"], _ = expandObjectRouterRouteMapRuleMatchAsPath(d, i["match_as_path"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_community"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-community"], _ = expandObjectRouterRouteMapRuleMatchCommunity(d, i["match_community"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_community_exact"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-community-exact"], _ = expandObjectRouterRouteMapRuleMatchCommunityExact(d, i["match_community_exact"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_extcommunity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-extcommunity"], _ = expandObjectRouterRouteMapRuleMatchExtcommunity(d, i["match_extcommunity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_extcommunity_exact"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-extcommunity-exact"], _ = expandObjectRouterRouteMapRuleMatchExtcommunityExact(d, i["match_extcommunity_exact"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_flags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-flags"], _ = expandObjectRouterRouteMapRuleMatchFlags(d, i["match_flags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-interface"], _ = expandObjectRouterRouteMapRuleMatchInterface(d, i["match_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-ip-address"], _ = expandObjectRouterRouteMapRuleMatchIpAddress(d, i["match_ip_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip_nexthop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-ip-nexthop"], _ = expandObjectRouterRouteMapRuleMatchIpNexthop(d, i["match_ip_nexthop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip6_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-ip6-address"], _ = expandObjectRouterRouteMapRuleMatchIp6Address(d, i["match_ip6_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip6_nexthop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-ip6-nexthop"], _ = expandObjectRouterRouteMapRuleMatchIp6Nexthop(d, i["match_ip6_nexthop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_metric"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-metric"], _ = expandObjectRouterRouteMapRuleMatchMetric(d, i["match_metric"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_origin"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-origin"], _ = expandObjectRouterRouteMapRuleMatchOrigin(d, i["match_origin"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_route_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-route-type"], _ = expandObjectRouterRouteMapRuleMatchRouteType(d, i["match_route_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-tag"], _ = expandObjectRouterRouteMapRuleMatchTag(d, i["match_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-vrf"], _ = expandObjectRouterRouteMapRuleMatchVrf(d, i["match_vrf"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aggregator_as"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-aggregator-as"], _ = expandObjectRouterRouteMapRuleSetAggregatorAs(d, i["set_aggregator_as"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aggregator_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-aggregator-ip"], _ = expandObjectRouterRouteMapRuleSetAggregatorIp(d, i["set_aggregator_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aspath"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-aspath"], _ = expandObjectRouterRouteMapRuleSetAspath(d, i["set_aspath"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aspath_action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-aspath-action"], _ = expandObjectRouterRouteMapRuleSetAspathAction(d, i["set_aspath_action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_atomic_aggregate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-atomic-aggregate"], _ = expandObjectRouterRouteMapRuleSetAtomicAggregate(d, i["set_atomic_aggregate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-community"], _ = expandObjectRouterRouteMapRuleSetCommunity(d, i["set_community"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community_additive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-community-additive"], _ = expandObjectRouterRouteMapRuleSetCommunityAdditive(d, i["set_community_additive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community_delete"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-community-delete"], _ = expandObjectRouterRouteMapRuleSetCommunityDelete(d, i["set_community_delete"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_max_suppress"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-dampening-max-suppress"], _ = expandObjectRouterRouteMapRuleSetDampeningMaxSuppress(d, i["set_dampening_max_suppress"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_reachability_half_life"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-dampening-reachability-half-life"], _ = expandObjectRouterRouteMapRuleSetDampeningReachabilityHalfLife(d, i["set_dampening_reachability_half_life"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_reuse"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-dampening-reuse"], _ = expandObjectRouterRouteMapRuleSetDampeningReuse(d, i["set_dampening_reuse"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_suppress"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-dampening-suppress"], _ = expandObjectRouterRouteMapRuleSetDampeningSuppress(d, i["set_dampening_suppress"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_unreachability_half_life"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-dampening-unreachability-half-life"], _ = expandObjectRouterRouteMapRuleSetDampeningUnreachabilityHalfLife(d, i["set_dampening_unreachability_half_life"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_extcommunity_rt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-extcommunity-rt"], _ = expandObjectRouterRouteMapRuleSetExtcommunityRt(d, i["set_extcommunity_rt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_extcommunity_soo"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-extcommunity-soo"], _ = expandObjectRouterRouteMapRuleSetExtcommunitySoo(d, i["set_extcommunity_soo"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_flags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-flags"], _ = expandObjectRouterRouteMapRuleSetFlags(d, i["set_flags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip_nexthop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-ip-nexthop"], _ = expandObjectRouterRouteMapRuleSetIpNexthop(d, i["set_ip_nexthop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip_prefsrc"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-ip-prefsrc"], _ = expandObjectRouterRouteMapRuleSetIpPrefsrc(d, i["set_ip_prefsrc"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip6_nexthop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-ip6-nexthop"], _ = expandObjectRouterRouteMapRuleSetIp6Nexthop(d, i["set_ip6_nexthop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip6_nexthop_local"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-ip6-nexthop-local"], _ = expandObjectRouterRouteMapRuleSetIp6NexthopLocal(d, i["set_ip6_nexthop_local"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_local_preference"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-local-preference"], _ = expandObjectRouterRouteMapRuleSetLocalPreference(d, i["set_local_preference"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_metric"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-metric"], _ = expandObjectRouterRouteMapRuleSetMetric(d, i["set_metric"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_metric_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-metric-type"], _ = expandObjectRouterRouteMapRuleSetMetricType(d, i["set_metric_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_origin"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-origin"], _ = expandObjectRouterRouteMapRuleSetOrigin(d, i["set_origin"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_originator_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-originator-id"], _ = expandObjectRouterRouteMapRuleSetOriginatorId(d, i["set_originator_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-priority"], _ = expandObjectRouterRouteMapRuleSetPriority(d, i["set_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_route_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-route-tag"], _ = expandObjectRouterRouteMapRuleSetRouteTag(d, i["set_route_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-tag"], _ = expandObjectRouterRouteMapRuleSetTag(d, i["set_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv4_nexthop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-vpnv4-nexthop"], _ = expandObjectRouterRouteMapRuleSetVpnv4Nexthop(d, i["set_vpnv4_nexthop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv6_nexthop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-vpnv6-nexthop"], _ = expandObjectRouterRouteMapRuleSetVpnv6Nexthop(d, i["set_vpnv6_nexthop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv6_nexthop_local"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-vpnv6-nexthop-local"], _ = expandObjectRouterRouteMapRuleSetVpnv6NexthopLocal(d, i["set_vpnv6_nexthop_local"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["set-weight"], _ = expandObjectRouterRouteMapRuleSetWeight(d, i["set_weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectRouterRouteMapRuleAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchAsPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchCommunity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchCommunityExact(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchExtcommunity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchExtcommunityExact(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchIpNexthop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchIp6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchIp6Nexthop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchOrigin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchRouteType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetAggregatorAs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetAggregatorIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetAspath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectRouterRouteMapRuleSetAspathAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetAtomicAggregate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetCommunity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectRouterRouteMapRuleSetCommunityAdditive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetCommunityDelete(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetDampeningMaxSuppress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetDampeningReachabilityHalfLife(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetDampeningReuse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetDampeningSuppress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetDampeningUnreachabilityHalfLife(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetExtcommunityRt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectRouterRouteMapRuleSetExtcommunitySoo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectRouterRouteMapRuleSetFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetIpNexthop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetIpPrefsrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetIp6Nexthop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetIp6NexthopLocal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetLocalPreference(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetMetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetOrigin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetOriginatorId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetRouteTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetVpnv4Nexthop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetVpnv6Nexthop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetVpnv6NexthopLocal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectRouterRouteMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandObjectRouterRouteMapComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectRouterRouteMapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandObjectRouterRouteMapRule(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	return &obj, nil
}
