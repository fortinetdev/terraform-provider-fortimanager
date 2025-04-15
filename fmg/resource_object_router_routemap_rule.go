// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Rule.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectRouterRouteMapRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectRouterRouteMapRuleCreate,
		Read:   resourceObjectRouterRouteMapRuleRead,
		Update: resourceObjectRouterRouteMapRuleUpdate,
		Delete: resourceObjectRouterRouteMapRuleDelete,

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
			"route_map": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
	}
}

func resourceObjectRouterRouteMapRuleCreate(d *schema.ResourceData, m interface{}) error {
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

	route_map := d.Get("route_map").(string)
	paradict["route_map"] = route_map

	obj, err := getObjectObjectRouterRouteMapRule(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterRouteMapRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectRouterRouteMapRule(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterRouteMapRule resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectRouterRouteMapRuleRead(d, m)
}

func resourceObjectRouterRouteMapRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	route_map := d.Get("route_map").(string)
	paradict["route_map"] = route_map

	obj, err := getObjectObjectRouterRouteMapRule(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterRouteMapRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectRouterRouteMapRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterRouteMapRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectRouterRouteMapRuleRead(d, m)
}

func resourceObjectRouterRouteMapRuleDelete(d *schema.ResourceData, m interface{}) error {
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

	route_map := d.Get("route_map").(string)
	paradict["route_map"] = route_map

	wsParams["adom"] = adomv

	err = c.DeleteObjectRouterRouteMapRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectRouterRouteMapRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectRouterRouteMapRuleRead(d *schema.ResourceData, m interface{}) error {
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

	route_map := d.Get("route_map").(string)
	if route_map == "" {
		route_map = importOptionChecking(m.(*FortiClient).Cfg, "route_map")
		if route_map == "" {
			return fmt.Errorf("Parameter route_map is missing")
		}
		if err = d.Set("route_map", route_map); err != nil {
			return fmt.Errorf("Error set params route_map: %v", err)
		}
	}
	paradict["route_map"] = route_map

	o, err := c.ReadObjectRouterRouteMapRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterRouteMapRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectRouterRouteMapRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterRouteMapRule resource from API: %v", err)
	}
	return nil
}

func flattenObjectRouterRouteMapRuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchAsPath2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectRouterRouteMapRuleMatchCommunity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectRouterRouteMapRuleMatchCommunityExact2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchExtcommunity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectRouterRouteMapRuleMatchExtcommunityExact2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectRouterRouteMapRuleMatchIpAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectRouterRouteMapRuleMatchIpNexthop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectRouterRouteMapRuleMatchIp6Address2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectRouterRouteMapRuleMatchIp6Nexthop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectRouterRouteMapRuleMatchMetric2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchOrigin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchRouteType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleMatchVrf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetAggregatorAs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetAggregatorIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetAspath2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectRouterRouteMapRuleSetAspathAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetAtomicAggregate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetCommunity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectRouterRouteMapRuleSetCommunityAdditive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetCommunityDelete2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectRouterRouteMapRuleSetDampeningMaxSuppress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetDampeningReachabilityHalfLife2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetDampeningReuse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetDampeningSuppress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetDampeningUnreachabilityHalfLife2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetExtcommunityRt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectRouterRouteMapRuleSetExtcommunitySoo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectRouterRouteMapRuleSetFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetIpNexthop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetIpPrefsrc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetIp6Nexthop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetIp6NexthopLocal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetLocalPreference2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetMetric2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetMetricType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetOrigin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetOriginatorId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetRouteTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetVpnv4Nexthop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetVpnv6Nexthop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetVpnv6NexthopLocal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterRouteMapRuleSetWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectRouterRouteMapRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectRouterRouteMapRuleAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectRouterRouteMapRule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectRouterRouteMapRuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectRouterRouteMapRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("match_as_path", flattenObjectRouterRouteMapRuleMatchAsPath2edl(o["match-as-path"], d, "match_as_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-as-path"], "ObjectRouterRouteMapRule-MatchAsPath"); ok {
			if err = d.Set("match_as_path", vv); err != nil {
				return fmt.Errorf("Error reading match_as_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_as_path: %v", err)
		}
	}

	if err = d.Set("match_community", flattenObjectRouterRouteMapRuleMatchCommunity2edl(o["match-community"], d, "match_community")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-community"], "ObjectRouterRouteMapRule-MatchCommunity"); ok {
			if err = d.Set("match_community", vv); err != nil {
				return fmt.Errorf("Error reading match_community: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_community: %v", err)
		}
	}

	if err = d.Set("match_community_exact", flattenObjectRouterRouteMapRuleMatchCommunityExact2edl(o["match-community-exact"], d, "match_community_exact")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-community-exact"], "ObjectRouterRouteMapRule-MatchCommunityExact"); ok {
			if err = d.Set("match_community_exact", vv); err != nil {
				return fmt.Errorf("Error reading match_community_exact: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_community_exact: %v", err)
		}
	}

	if err = d.Set("match_extcommunity", flattenObjectRouterRouteMapRuleMatchExtcommunity2edl(o["match-extcommunity"], d, "match_extcommunity")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-extcommunity"], "ObjectRouterRouteMapRule-MatchExtcommunity"); ok {
			if err = d.Set("match_extcommunity", vv); err != nil {
				return fmt.Errorf("Error reading match_extcommunity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_extcommunity: %v", err)
		}
	}

	if err = d.Set("match_extcommunity_exact", flattenObjectRouterRouteMapRuleMatchExtcommunityExact2edl(o["match-extcommunity-exact"], d, "match_extcommunity_exact")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-extcommunity-exact"], "ObjectRouterRouteMapRule-MatchExtcommunityExact"); ok {
			if err = d.Set("match_extcommunity_exact", vv); err != nil {
				return fmt.Errorf("Error reading match_extcommunity_exact: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_extcommunity_exact: %v", err)
		}
	}

	if err = d.Set("match_flags", flattenObjectRouterRouteMapRuleMatchFlags2edl(o["match-flags"], d, "match_flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-flags"], "ObjectRouterRouteMapRule-MatchFlags"); ok {
			if err = d.Set("match_flags", vv); err != nil {
				return fmt.Errorf("Error reading match_flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_flags: %v", err)
		}
	}

	if err = d.Set("match_interface", flattenObjectRouterRouteMapRuleMatchInterface2edl(o["match-interface"], d, "match_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-interface"], "ObjectRouterRouteMapRule-MatchInterface"); ok {
			if err = d.Set("match_interface", vv); err != nil {
				return fmt.Errorf("Error reading match_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_interface: %v", err)
		}
	}

	if err = d.Set("match_ip_address", flattenObjectRouterRouteMapRuleMatchIpAddress2edl(o["match-ip-address"], d, "match_ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-ip-address"], "ObjectRouterRouteMapRule-MatchIpAddress"); ok {
			if err = d.Set("match_ip_address", vv); err != nil {
				return fmt.Errorf("Error reading match_ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_ip_address: %v", err)
		}
	}

	if err = d.Set("match_ip_nexthop", flattenObjectRouterRouteMapRuleMatchIpNexthop2edl(o["match-ip-nexthop"], d, "match_ip_nexthop")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-ip-nexthop"], "ObjectRouterRouteMapRule-MatchIpNexthop"); ok {
			if err = d.Set("match_ip_nexthop", vv); err != nil {
				return fmt.Errorf("Error reading match_ip_nexthop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_ip_nexthop: %v", err)
		}
	}

	if err = d.Set("match_ip6_address", flattenObjectRouterRouteMapRuleMatchIp6Address2edl(o["match-ip6-address"], d, "match_ip6_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-ip6-address"], "ObjectRouterRouteMapRule-MatchIp6Address"); ok {
			if err = d.Set("match_ip6_address", vv); err != nil {
				return fmt.Errorf("Error reading match_ip6_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_ip6_address: %v", err)
		}
	}

	if err = d.Set("match_ip6_nexthop", flattenObjectRouterRouteMapRuleMatchIp6Nexthop2edl(o["match-ip6-nexthop"], d, "match_ip6_nexthop")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-ip6-nexthop"], "ObjectRouterRouteMapRule-MatchIp6Nexthop"); ok {
			if err = d.Set("match_ip6_nexthop", vv); err != nil {
				return fmt.Errorf("Error reading match_ip6_nexthop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_ip6_nexthop: %v", err)
		}
	}

	if err = d.Set("match_metric", flattenObjectRouterRouteMapRuleMatchMetric2edl(o["match-metric"], d, "match_metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-metric"], "ObjectRouterRouteMapRule-MatchMetric"); ok {
			if err = d.Set("match_metric", vv); err != nil {
				return fmt.Errorf("Error reading match_metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_metric: %v", err)
		}
	}

	if err = d.Set("match_origin", flattenObjectRouterRouteMapRuleMatchOrigin2edl(o["match-origin"], d, "match_origin")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-origin"], "ObjectRouterRouteMapRule-MatchOrigin"); ok {
			if err = d.Set("match_origin", vv); err != nil {
				return fmt.Errorf("Error reading match_origin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_origin: %v", err)
		}
	}

	if err = d.Set("match_route_type", flattenObjectRouterRouteMapRuleMatchRouteType2edl(o["match-route-type"], d, "match_route_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-route-type"], "ObjectRouterRouteMapRule-MatchRouteType"); ok {
			if err = d.Set("match_route_type", vv); err != nil {
				return fmt.Errorf("Error reading match_route_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_route_type: %v", err)
		}
	}

	if err = d.Set("match_tag", flattenObjectRouterRouteMapRuleMatchTag2edl(o["match-tag"], d, "match_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-tag"], "ObjectRouterRouteMapRule-MatchTag"); ok {
			if err = d.Set("match_tag", vv); err != nil {
				return fmt.Errorf("Error reading match_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_tag: %v", err)
		}
	}

	if err = d.Set("match_vrf", flattenObjectRouterRouteMapRuleMatchVrf2edl(o["match-vrf"], d, "match_vrf")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-vrf"], "ObjectRouterRouteMapRule-MatchVrf"); ok {
			if err = d.Set("match_vrf", vv); err != nil {
				return fmt.Errorf("Error reading match_vrf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_vrf: %v", err)
		}
	}

	if err = d.Set("set_aggregator_as", flattenObjectRouterRouteMapRuleSetAggregatorAs2edl(o["set-aggregator-as"], d, "set_aggregator_as")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-aggregator-as"], "ObjectRouterRouteMapRule-SetAggregatorAs"); ok {
			if err = d.Set("set_aggregator_as", vv); err != nil {
				return fmt.Errorf("Error reading set_aggregator_as: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_aggregator_as: %v", err)
		}
	}

	if err = d.Set("set_aggregator_ip", flattenObjectRouterRouteMapRuleSetAggregatorIp2edl(o["set-aggregator-ip"], d, "set_aggregator_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-aggregator-ip"], "ObjectRouterRouteMapRule-SetAggregatorIp"); ok {
			if err = d.Set("set_aggregator_ip", vv); err != nil {
				return fmt.Errorf("Error reading set_aggregator_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_aggregator_ip: %v", err)
		}
	}

	if err = d.Set("set_aspath", flattenObjectRouterRouteMapRuleSetAspath2edl(o["set-aspath"], d, "set_aspath")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-aspath"], "ObjectRouterRouteMapRule-SetAspath"); ok {
			if err = d.Set("set_aspath", vv); err != nil {
				return fmt.Errorf("Error reading set_aspath: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_aspath: %v", err)
		}
	}

	if err = d.Set("set_aspath_action", flattenObjectRouterRouteMapRuleSetAspathAction2edl(o["set-aspath-action"], d, "set_aspath_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-aspath-action"], "ObjectRouterRouteMapRule-SetAspathAction"); ok {
			if err = d.Set("set_aspath_action", vv); err != nil {
				return fmt.Errorf("Error reading set_aspath_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_aspath_action: %v", err)
		}
	}

	if err = d.Set("set_atomic_aggregate", flattenObjectRouterRouteMapRuleSetAtomicAggregate2edl(o["set-atomic-aggregate"], d, "set_atomic_aggregate")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-atomic-aggregate"], "ObjectRouterRouteMapRule-SetAtomicAggregate"); ok {
			if err = d.Set("set_atomic_aggregate", vv); err != nil {
				return fmt.Errorf("Error reading set_atomic_aggregate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_atomic_aggregate: %v", err)
		}
	}

	if err = d.Set("set_community", flattenObjectRouterRouteMapRuleSetCommunity2edl(o["set-community"], d, "set_community")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-community"], "ObjectRouterRouteMapRule-SetCommunity"); ok {
			if err = d.Set("set_community", vv); err != nil {
				return fmt.Errorf("Error reading set_community: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_community: %v", err)
		}
	}

	if err = d.Set("set_community_additive", flattenObjectRouterRouteMapRuleSetCommunityAdditive2edl(o["set-community-additive"], d, "set_community_additive")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-community-additive"], "ObjectRouterRouteMapRule-SetCommunityAdditive"); ok {
			if err = d.Set("set_community_additive", vv); err != nil {
				return fmt.Errorf("Error reading set_community_additive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_community_additive: %v", err)
		}
	}

	if err = d.Set("set_community_delete", flattenObjectRouterRouteMapRuleSetCommunityDelete2edl(o["set-community-delete"], d, "set_community_delete")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-community-delete"], "ObjectRouterRouteMapRule-SetCommunityDelete"); ok {
			if err = d.Set("set_community_delete", vv); err != nil {
				return fmt.Errorf("Error reading set_community_delete: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_community_delete: %v", err)
		}
	}

	if err = d.Set("set_dampening_max_suppress", flattenObjectRouterRouteMapRuleSetDampeningMaxSuppress2edl(o["set-dampening-max-suppress"], d, "set_dampening_max_suppress")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-dampening-max-suppress"], "ObjectRouterRouteMapRule-SetDampeningMaxSuppress"); ok {
			if err = d.Set("set_dampening_max_suppress", vv); err != nil {
				return fmt.Errorf("Error reading set_dampening_max_suppress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_dampening_max_suppress: %v", err)
		}
	}

	if err = d.Set("set_dampening_reachability_half_life", flattenObjectRouterRouteMapRuleSetDampeningReachabilityHalfLife2edl(o["set-dampening-reachability-half-life"], d, "set_dampening_reachability_half_life")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-dampening-reachability-half-life"], "ObjectRouterRouteMapRule-SetDampeningReachabilityHalfLife"); ok {
			if err = d.Set("set_dampening_reachability_half_life", vv); err != nil {
				return fmt.Errorf("Error reading set_dampening_reachability_half_life: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_dampening_reachability_half_life: %v", err)
		}
	}

	if err = d.Set("set_dampening_reuse", flattenObjectRouterRouteMapRuleSetDampeningReuse2edl(o["set-dampening-reuse"], d, "set_dampening_reuse")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-dampening-reuse"], "ObjectRouterRouteMapRule-SetDampeningReuse"); ok {
			if err = d.Set("set_dampening_reuse", vv); err != nil {
				return fmt.Errorf("Error reading set_dampening_reuse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_dampening_reuse: %v", err)
		}
	}

	if err = d.Set("set_dampening_suppress", flattenObjectRouterRouteMapRuleSetDampeningSuppress2edl(o["set-dampening-suppress"], d, "set_dampening_suppress")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-dampening-suppress"], "ObjectRouterRouteMapRule-SetDampeningSuppress"); ok {
			if err = d.Set("set_dampening_suppress", vv); err != nil {
				return fmt.Errorf("Error reading set_dampening_suppress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_dampening_suppress: %v", err)
		}
	}

	if err = d.Set("set_dampening_unreachability_half_life", flattenObjectRouterRouteMapRuleSetDampeningUnreachabilityHalfLife2edl(o["set-dampening-unreachability-half-life"], d, "set_dampening_unreachability_half_life")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-dampening-unreachability-half-life"], "ObjectRouterRouteMapRule-SetDampeningUnreachabilityHalfLife"); ok {
			if err = d.Set("set_dampening_unreachability_half_life", vv); err != nil {
				return fmt.Errorf("Error reading set_dampening_unreachability_half_life: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_dampening_unreachability_half_life: %v", err)
		}
	}

	if err = d.Set("set_extcommunity_rt", flattenObjectRouterRouteMapRuleSetExtcommunityRt2edl(o["set-extcommunity-rt"], d, "set_extcommunity_rt")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-extcommunity-rt"], "ObjectRouterRouteMapRule-SetExtcommunityRt"); ok {
			if err = d.Set("set_extcommunity_rt", vv); err != nil {
				return fmt.Errorf("Error reading set_extcommunity_rt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_extcommunity_rt: %v", err)
		}
	}

	if err = d.Set("set_extcommunity_soo", flattenObjectRouterRouteMapRuleSetExtcommunitySoo2edl(o["set-extcommunity-soo"], d, "set_extcommunity_soo")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-extcommunity-soo"], "ObjectRouterRouteMapRule-SetExtcommunitySoo"); ok {
			if err = d.Set("set_extcommunity_soo", vv); err != nil {
				return fmt.Errorf("Error reading set_extcommunity_soo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_extcommunity_soo: %v", err)
		}
	}

	if err = d.Set("set_flags", flattenObjectRouterRouteMapRuleSetFlags2edl(o["set-flags"], d, "set_flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-flags"], "ObjectRouterRouteMapRule-SetFlags"); ok {
			if err = d.Set("set_flags", vv); err != nil {
				return fmt.Errorf("Error reading set_flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_flags: %v", err)
		}
	}

	if err = d.Set("set_ip_nexthop", flattenObjectRouterRouteMapRuleSetIpNexthop2edl(o["set-ip-nexthop"], d, "set_ip_nexthop")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-ip-nexthop"], "ObjectRouterRouteMapRule-SetIpNexthop"); ok {
			if err = d.Set("set_ip_nexthop", vv); err != nil {
				return fmt.Errorf("Error reading set_ip_nexthop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_ip_nexthop: %v", err)
		}
	}

	if err = d.Set("set_ip_prefsrc", flattenObjectRouterRouteMapRuleSetIpPrefsrc2edl(o["set-ip-prefsrc"], d, "set_ip_prefsrc")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-ip-prefsrc"], "ObjectRouterRouteMapRule-SetIpPrefsrc"); ok {
			if err = d.Set("set_ip_prefsrc", vv); err != nil {
				return fmt.Errorf("Error reading set_ip_prefsrc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_ip_prefsrc: %v", err)
		}
	}

	if err = d.Set("set_ip6_nexthop", flattenObjectRouterRouteMapRuleSetIp6Nexthop2edl(o["set-ip6-nexthop"], d, "set_ip6_nexthop")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-ip6-nexthop"], "ObjectRouterRouteMapRule-SetIp6Nexthop"); ok {
			if err = d.Set("set_ip6_nexthop", vv); err != nil {
				return fmt.Errorf("Error reading set_ip6_nexthop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_ip6_nexthop: %v", err)
		}
	}

	if err = d.Set("set_ip6_nexthop_local", flattenObjectRouterRouteMapRuleSetIp6NexthopLocal2edl(o["set-ip6-nexthop-local"], d, "set_ip6_nexthop_local")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-ip6-nexthop-local"], "ObjectRouterRouteMapRule-SetIp6NexthopLocal"); ok {
			if err = d.Set("set_ip6_nexthop_local", vv); err != nil {
				return fmt.Errorf("Error reading set_ip6_nexthop_local: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_ip6_nexthop_local: %v", err)
		}
	}

	if err = d.Set("set_local_preference", flattenObjectRouterRouteMapRuleSetLocalPreference2edl(o["set-local-preference"], d, "set_local_preference")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-local-preference"], "ObjectRouterRouteMapRule-SetLocalPreference"); ok {
			if err = d.Set("set_local_preference", vv); err != nil {
				return fmt.Errorf("Error reading set_local_preference: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_local_preference: %v", err)
		}
	}

	if err = d.Set("set_metric", flattenObjectRouterRouteMapRuleSetMetric2edl(o["set-metric"], d, "set_metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-metric"], "ObjectRouterRouteMapRule-SetMetric"); ok {
			if err = d.Set("set_metric", vv); err != nil {
				return fmt.Errorf("Error reading set_metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_metric: %v", err)
		}
	}

	if err = d.Set("set_metric_type", flattenObjectRouterRouteMapRuleSetMetricType2edl(o["set-metric-type"], d, "set_metric_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-metric-type"], "ObjectRouterRouteMapRule-SetMetricType"); ok {
			if err = d.Set("set_metric_type", vv); err != nil {
				return fmt.Errorf("Error reading set_metric_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_metric_type: %v", err)
		}
	}

	if err = d.Set("set_origin", flattenObjectRouterRouteMapRuleSetOrigin2edl(o["set-origin"], d, "set_origin")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-origin"], "ObjectRouterRouteMapRule-SetOrigin"); ok {
			if err = d.Set("set_origin", vv); err != nil {
				return fmt.Errorf("Error reading set_origin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_origin: %v", err)
		}
	}

	if err = d.Set("set_originator_id", flattenObjectRouterRouteMapRuleSetOriginatorId2edl(o["set-originator-id"], d, "set_originator_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-originator-id"], "ObjectRouterRouteMapRule-SetOriginatorId"); ok {
			if err = d.Set("set_originator_id", vv); err != nil {
				return fmt.Errorf("Error reading set_originator_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_originator_id: %v", err)
		}
	}

	if err = d.Set("set_priority", flattenObjectRouterRouteMapRuleSetPriority2edl(o["set-priority"], d, "set_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-priority"], "ObjectRouterRouteMapRule-SetPriority"); ok {
			if err = d.Set("set_priority", vv); err != nil {
				return fmt.Errorf("Error reading set_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_priority: %v", err)
		}
	}

	if err = d.Set("set_route_tag", flattenObjectRouterRouteMapRuleSetRouteTag2edl(o["set-route-tag"], d, "set_route_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-route-tag"], "ObjectRouterRouteMapRule-SetRouteTag"); ok {
			if err = d.Set("set_route_tag", vv); err != nil {
				return fmt.Errorf("Error reading set_route_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_route_tag: %v", err)
		}
	}

	if err = d.Set("set_tag", flattenObjectRouterRouteMapRuleSetTag2edl(o["set-tag"], d, "set_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-tag"], "ObjectRouterRouteMapRule-SetTag"); ok {
			if err = d.Set("set_tag", vv); err != nil {
				return fmt.Errorf("Error reading set_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_tag: %v", err)
		}
	}

	if err = d.Set("set_vpnv4_nexthop", flattenObjectRouterRouteMapRuleSetVpnv4Nexthop2edl(o["set-vpnv4-nexthop"], d, "set_vpnv4_nexthop")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-vpnv4-nexthop"], "ObjectRouterRouteMapRule-SetVpnv4Nexthop"); ok {
			if err = d.Set("set_vpnv4_nexthop", vv); err != nil {
				return fmt.Errorf("Error reading set_vpnv4_nexthop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_vpnv4_nexthop: %v", err)
		}
	}

	if err = d.Set("set_vpnv6_nexthop", flattenObjectRouterRouteMapRuleSetVpnv6Nexthop2edl(o["set-vpnv6-nexthop"], d, "set_vpnv6_nexthop")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-vpnv6-nexthop"], "ObjectRouterRouteMapRule-SetVpnv6Nexthop"); ok {
			if err = d.Set("set_vpnv6_nexthop", vv); err != nil {
				return fmt.Errorf("Error reading set_vpnv6_nexthop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_vpnv6_nexthop: %v", err)
		}
	}

	if err = d.Set("set_vpnv6_nexthop_local", flattenObjectRouterRouteMapRuleSetVpnv6NexthopLocal2edl(o["set-vpnv6-nexthop-local"], d, "set_vpnv6_nexthop_local")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-vpnv6-nexthop-local"], "ObjectRouterRouteMapRule-SetVpnv6NexthopLocal"); ok {
			if err = d.Set("set_vpnv6_nexthop_local", vv); err != nil {
				return fmt.Errorf("Error reading set_vpnv6_nexthop_local: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_vpnv6_nexthop_local: %v", err)
		}
	}

	if err = d.Set("set_weight", flattenObjectRouterRouteMapRuleSetWeight2edl(o["set-weight"], d, "set_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["set-weight"], "ObjectRouterRouteMapRule-SetWeight"); ok {
			if err = d.Set("set_weight", vv); err != nil {
				return fmt.Errorf("Error reading set_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading set_weight: %v", err)
		}
	}

	return nil
}

func flattenObjectRouterRouteMapRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectRouterRouteMapRuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchAsPath2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectRouterRouteMapRuleMatchCommunity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectRouterRouteMapRuleMatchCommunityExact2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchExtcommunity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectRouterRouteMapRuleMatchExtcommunityExact2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectRouterRouteMapRuleMatchIpAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectRouterRouteMapRuleMatchIpNexthop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectRouterRouteMapRuleMatchIp6Address2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectRouterRouteMapRuleMatchIp6Nexthop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectRouterRouteMapRuleMatchMetric2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchOrigin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchRouteType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleMatchVrf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetAggregatorAs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetAggregatorIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetAspath2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectRouterRouteMapRuleSetAspathAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetAtomicAggregate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetCommunity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectRouterRouteMapRuleSetCommunityAdditive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetCommunityDelete2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectRouterRouteMapRuleSetDampeningMaxSuppress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetDampeningReachabilityHalfLife2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetDampeningReuse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetDampeningSuppress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetDampeningUnreachabilityHalfLife2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetExtcommunityRt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectRouterRouteMapRuleSetExtcommunitySoo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectRouterRouteMapRuleSetFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetIpNexthop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetIpPrefsrc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetIp6Nexthop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetIp6NexthopLocal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetLocalPreference2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetMetric2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetMetricType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetOrigin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetOriginatorId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetRouteTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetVpnv4Nexthop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetVpnv6Nexthop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetVpnv6NexthopLocal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterRouteMapRuleSetWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectRouterRouteMapRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectRouterRouteMapRuleAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectRouterRouteMapRuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("match_as_path"); ok || d.HasChange("match_as_path") {
		t, err := expandObjectRouterRouteMapRuleMatchAsPath2edl(d, v, "match_as_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-as-path"] = t
		}
	}

	if v, ok := d.GetOk("match_community"); ok || d.HasChange("match_community") {
		t, err := expandObjectRouterRouteMapRuleMatchCommunity2edl(d, v, "match_community")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-community"] = t
		}
	}

	if v, ok := d.GetOk("match_community_exact"); ok || d.HasChange("match_community_exact") {
		t, err := expandObjectRouterRouteMapRuleMatchCommunityExact2edl(d, v, "match_community_exact")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-community-exact"] = t
		}
	}

	if v, ok := d.GetOk("match_extcommunity"); ok || d.HasChange("match_extcommunity") {
		t, err := expandObjectRouterRouteMapRuleMatchExtcommunity2edl(d, v, "match_extcommunity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-extcommunity"] = t
		}
	}

	if v, ok := d.GetOk("match_extcommunity_exact"); ok || d.HasChange("match_extcommunity_exact") {
		t, err := expandObjectRouterRouteMapRuleMatchExtcommunityExact2edl(d, v, "match_extcommunity_exact")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-extcommunity-exact"] = t
		}
	}

	if v, ok := d.GetOk("match_flags"); ok || d.HasChange("match_flags") {
		t, err := expandObjectRouterRouteMapRuleMatchFlags2edl(d, v, "match_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-flags"] = t
		}
	}

	if v, ok := d.GetOk("match_interface"); ok || d.HasChange("match_interface") {
		t, err := expandObjectRouterRouteMapRuleMatchInterface2edl(d, v, "match_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-interface"] = t
		}
	}

	if v, ok := d.GetOk("match_ip_address"); ok || d.HasChange("match_ip_address") {
		t, err := expandObjectRouterRouteMapRuleMatchIpAddress2edl(d, v, "match_ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-ip-address"] = t
		}
	}

	if v, ok := d.GetOk("match_ip_nexthop"); ok || d.HasChange("match_ip_nexthop") {
		t, err := expandObjectRouterRouteMapRuleMatchIpNexthop2edl(d, v, "match_ip_nexthop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-ip-nexthop"] = t
		}
	}

	if v, ok := d.GetOk("match_ip6_address"); ok || d.HasChange("match_ip6_address") {
		t, err := expandObjectRouterRouteMapRuleMatchIp6Address2edl(d, v, "match_ip6_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-ip6-address"] = t
		}
	}

	if v, ok := d.GetOk("match_ip6_nexthop"); ok || d.HasChange("match_ip6_nexthop") {
		t, err := expandObjectRouterRouteMapRuleMatchIp6Nexthop2edl(d, v, "match_ip6_nexthop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-ip6-nexthop"] = t
		}
	}

	if v, ok := d.GetOk("match_metric"); ok || d.HasChange("match_metric") {
		t, err := expandObjectRouterRouteMapRuleMatchMetric2edl(d, v, "match_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-metric"] = t
		}
	}

	if v, ok := d.GetOk("match_origin"); ok || d.HasChange("match_origin") {
		t, err := expandObjectRouterRouteMapRuleMatchOrigin2edl(d, v, "match_origin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-origin"] = t
		}
	}

	if v, ok := d.GetOk("match_route_type"); ok || d.HasChange("match_route_type") {
		t, err := expandObjectRouterRouteMapRuleMatchRouteType2edl(d, v, "match_route_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-route-type"] = t
		}
	}

	if v, ok := d.GetOk("match_tag"); ok || d.HasChange("match_tag") {
		t, err := expandObjectRouterRouteMapRuleMatchTag2edl(d, v, "match_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-tag"] = t
		}
	}

	if v, ok := d.GetOk("match_vrf"); ok || d.HasChange("match_vrf") {
		t, err := expandObjectRouterRouteMapRuleMatchVrf2edl(d, v, "match_vrf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-vrf"] = t
		}
	}

	if v, ok := d.GetOk("set_aggregator_as"); ok || d.HasChange("set_aggregator_as") {
		t, err := expandObjectRouterRouteMapRuleSetAggregatorAs2edl(d, v, "set_aggregator_as")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-aggregator-as"] = t
		}
	}

	if v, ok := d.GetOk("set_aggregator_ip"); ok || d.HasChange("set_aggregator_ip") {
		t, err := expandObjectRouterRouteMapRuleSetAggregatorIp2edl(d, v, "set_aggregator_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-aggregator-ip"] = t
		}
	}

	if v, ok := d.GetOk("set_aspath"); ok || d.HasChange("set_aspath") {
		t, err := expandObjectRouterRouteMapRuleSetAspath2edl(d, v, "set_aspath")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-aspath"] = t
		}
	}

	if v, ok := d.GetOk("set_aspath_action"); ok || d.HasChange("set_aspath_action") {
		t, err := expandObjectRouterRouteMapRuleSetAspathAction2edl(d, v, "set_aspath_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-aspath-action"] = t
		}
	}

	if v, ok := d.GetOk("set_atomic_aggregate"); ok || d.HasChange("set_atomic_aggregate") {
		t, err := expandObjectRouterRouteMapRuleSetAtomicAggregate2edl(d, v, "set_atomic_aggregate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-atomic-aggregate"] = t
		}
	}

	if v, ok := d.GetOk("set_community"); ok || d.HasChange("set_community") {
		t, err := expandObjectRouterRouteMapRuleSetCommunity2edl(d, v, "set_community")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-community"] = t
		}
	}

	if v, ok := d.GetOk("set_community_additive"); ok || d.HasChange("set_community_additive") {
		t, err := expandObjectRouterRouteMapRuleSetCommunityAdditive2edl(d, v, "set_community_additive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-community-additive"] = t
		}
	}

	if v, ok := d.GetOk("set_community_delete"); ok || d.HasChange("set_community_delete") {
		t, err := expandObjectRouterRouteMapRuleSetCommunityDelete2edl(d, v, "set_community_delete")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-community-delete"] = t
		}
	}

	if v, ok := d.GetOk("set_dampening_max_suppress"); ok || d.HasChange("set_dampening_max_suppress") {
		t, err := expandObjectRouterRouteMapRuleSetDampeningMaxSuppress2edl(d, v, "set_dampening_max_suppress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-dampening-max-suppress"] = t
		}
	}

	if v, ok := d.GetOk("set_dampening_reachability_half_life"); ok || d.HasChange("set_dampening_reachability_half_life") {
		t, err := expandObjectRouterRouteMapRuleSetDampeningReachabilityHalfLife2edl(d, v, "set_dampening_reachability_half_life")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-dampening-reachability-half-life"] = t
		}
	}

	if v, ok := d.GetOk("set_dampening_reuse"); ok || d.HasChange("set_dampening_reuse") {
		t, err := expandObjectRouterRouteMapRuleSetDampeningReuse2edl(d, v, "set_dampening_reuse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-dampening-reuse"] = t
		}
	}

	if v, ok := d.GetOk("set_dampening_suppress"); ok || d.HasChange("set_dampening_suppress") {
		t, err := expandObjectRouterRouteMapRuleSetDampeningSuppress2edl(d, v, "set_dampening_suppress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-dampening-suppress"] = t
		}
	}

	if v, ok := d.GetOk("set_dampening_unreachability_half_life"); ok || d.HasChange("set_dampening_unreachability_half_life") {
		t, err := expandObjectRouterRouteMapRuleSetDampeningUnreachabilityHalfLife2edl(d, v, "set_dampening_unreachability_half_life")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-dampening-unreachability-half-life"] = t
		}
	}

	if v, ok := d.GetOk("set_extcommunity_rt"); ok || d.HasChange("set_extcommunity_rt") {
		t, err := expandObjectRouterRouteMapRuleSetExtcommunityRt2edl(d, v, "set_extcommunity_rt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-extcommunity-rt"] = t
		}
	}

	if v, ok := d.GetOk("set_extcommunity_soo"); ok || d.HasChange("set_extcommunity_soo") {
		t, err := expandObjectRouterRouteMapRuleSetExtcommunitySoo2edl(d, v, "set_extcommunity_soo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-extcommunity-soo"] = t
		}
	}

	if v, ok := d.GetOk("set_flags"); ok || d.HasChange("set_flags") {
		t, err := expandObjectRouterRouteMapRuleSetFlags2edl(d, v, "set_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-flags"] = t
		}
	}

	if v, ok := d.GetOk("set_ip_nexthop"); ok || d.HasChange("set_ip_nexthop") {
		t, err := expandObjectRouterRouteMapRuleSetIpNexthop2edl(d, v, "set_ip_nexthop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-ip-nexthop"] = t
		}
	}

	if v, ok := d.GetOk("set_ip_prefsrc"); ok || d.HasChange("set_ip_prefsrc") {
		t, err := expandObjectRouterRouteMapRuleSetIpPrefsrc2edl(d, v, "set_ip_prefsrc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-ip-prefsrc"] = t
		}
	}

	if v, ok := d.GetOk("set_ip6_nexthop"); ok || d.HasChange("set_ip6_nexthop") {
		t, err := expandObjectRouterRouteMapRuleSetIp6Nexthop2edl(d, v, "set_ip6_nexthop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-ip6-nexthop"] = t
		}
	}

	if v, ok := d.GetOk("set_ip6_nexthop_local"); ok || d.HasChange("set_ip6_nexthop_local") {
		t, err := expandObjectRouterRouteMapRuleSetIp6NexthopLocal2edl(d, v, "set_ip6_nexthop_local")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-ip6-nexthop-local"] = t
		}
	}

	if v, ok := d.GetOk("set_local_preference"); ok || d.HasChange("set_local_preference") {
		t, err := expandObjectRouterRouteMapRuleSetLocalPreference2edl(d, v, "set_local_preference")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-local-preference"] = t
		}
	}

	if v, ok := d.GetOk("set_metric"); ok || d.HasChange("set_metric") {
		t, err := expandObjectRouterRouteMapRuleSetMetric2edl(d, v, "set_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-metric"] = t
		}
	}

	if v, ok := d.GetOk("set_metric_type"); ok || d.HasChange("set_metric_type") {
		t, err := expandObjectRouterRouteMapRuleSetMetricType2edl(d, v, "set_metric_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-metric-type"] = t
		}
	}

	if v, ok := d.GetOk("set_origin"); ok || d.HasChange("set_origin") {
		t, err := expandObjectRouterRouteMapRuleSetOrigin2edl(d, v, "set_origin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-origin"] = t
		}
	}

	if v, ok := d.GetOk("set_originator_id"); ok || d.HasChange("set_originator_id") {
		t, err := expandObjectRouterRouteMapRuleSetOriginatorId2edl(d, v, "set_originator_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-originator-id"] = t
		}
	}

	if v, ok := d.GetOk("set_priority"); ok || d.HasChange("set_priority") {
		t, err := expandObjectRouterRouteMapRuleSetPriority2edl(d, v, "set_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-priority"] = t
		}
	}

	if v, ok := d.GetOk("set_route_tag"); ok || d.HasChange("set_route_tag") {
		t, err := expandObjectRouterRouteMapRuleSetRouteTag2edl(d, v, "set_route_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-route-tag"] = t
		}
	}

	if v, ok := d.GetOk("set_tag"); ok || d.HasChange("set_tag") {
		t, err := expandObjectRouterRouteMapRuleSetTag2edl(d, v, "set_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-tag"] = t
		}
	}

	if v, ok := d.GetOk("set_vpnv4_nexthop"); ok || d.HasChange("set_vpnv4_nexthop") {
		t, err := expandObjectRouterRouteMapRuleSetVpnv4Nexthop2edl(d, v, "set_vpnv4_nexthop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-vpnv4-nexthop"] = t
		}
	}

	if v, ok := d.GetOk("set_vpnv6_nexthop"); ok || d.HasChange("set_vpnv6_nexthop") {
		t, err := expandObjectRouterRouteMapRuleSetVpnv6Nexthop2edl(d, v, "set_vpnv6_nexthop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-vpnv6-nexthop"] = t
		}
	}

	if v, ok := d.GetOk("set_vpnv6_nexthop_local"); ok || d.HasChange("set_vpnv6_nexthop_local") {
		t, err := expandObjectRouterRouteMapRuleSetVpnv6NexthopLocal2edl(d, v, "set_vpnv6_nexthop_local")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-vpnv6-nexthop-local"] = t
		}
	}

	if v, ok := d.GetOk("set_weight"); ok || d.HasChange("set_weight") {
		t, err := expandObjectRouterRouteMapRuleSetWeight2edl(d, v, "set_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["set-weight"] = t
		}
	}

	return &obj, nil
}
