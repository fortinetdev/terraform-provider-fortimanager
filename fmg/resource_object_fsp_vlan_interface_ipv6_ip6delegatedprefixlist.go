// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Advertised IPv6 delegated prefix list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListCreate,
		Read:   resourceObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRead,
		Update: resourceObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpdate,
		Delete: resourceObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelete,

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
			"vlan": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"autonomous_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"delegated_prefix_iaid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"onlink_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prefix_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"rdnss": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rdnss_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upstream_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListCreate(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	obj, err := getObjectObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "prefix_id")))

	return resourceObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRead(d, m)
}

func resourceObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpdate(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	obj, err := getObjectObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "prefix_id")))

	return resourceObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRead(d, m)
}

func resourceObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelete(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	wsParams["adom"] = adomv

	err = c.DeleteObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRead(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	if vlan == "" {
		vlan = importOptionChecking(m.(*FortiClient).Cfg, "vlan")
		if vlan == "" {
			return fmt.Errorf("Parameter vlan is missing")
		}
		if err = d.Set("vlan", vlan); err != nil {
			return fmt.Errorf("Error set params vlan: %v", err)
		}
	}
	paradict["vlan"] = vlan

	o, err := c.ReadObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList resource from API: %v", err)
	}
	return nil
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("autonomous_flag", flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag4thl(o["autonomous-flag"], d, "autonomous_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["autonomous-flag"], "ObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList-AutonomousFlag"); ok {
			if err = d.Set("autonomous_flag", vv); err != nil {
				return fmt.Errorf("Error reading autonomous_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading autonomous_flag: %v", err)
		}
	}

	if err = d.Set("delegated_prefix_iaid", flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid4thl(o["delegated-prefix-iaid"], d, "delegated_prefix_iaid")); err != nil {
		if vv, ok := fortiAPIPatch(o["delegated-prefix-iaid"], "ObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList-DelegatedPrefixIaid"); ok {
			if err = d.Set("delegated_prefix_iaid", vv); err != nil {
				return fmt.Errorf("Error reading delegated_prefix_iaid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delegated_prefix_iaid: %v", err)
		}
	}

	if err = d.Set("onlink_flag", flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag4thl(o["onlink-flag"], d, "onlink_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["onlink-flag"], "ObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList-OnlinkFlag"); ok {
			if err = d.Set("onlink_flag", vv); err != nil {
				return fmt.Errorf("Error reading onlink_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading onlink_flag: %v", err)
		}
	}

	if err = d.Set("prefix_id", flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId4thl(o["prefix-id"], d, "prefix_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-id"], "ObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList-PrefixId"); ok {
			if err = d.Set("prefix_id", vv); err != nil {
				return fmt.Errorf("Error reading prefix_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_id: %v", err)
		}
	}

	if err = d.Set("rdnss", flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss4thl(o["rdnss"], d, "rdnss")); err != nil {
		if vv, ok := fortiAPIPatch(o["rdnss"], "ObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList-Rdnss"); ok {
			if err = d.Set("rdnss", vv); err != nil {
				return fmt.Errorf("Error reading rdnss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rdnss: %v", err)
		}
	}

	if err = d.Set("rdnss_service", flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService4thl(o["rdnss-service"], d, "rdnss_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["rdnss-service"], "ObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList-RdnssService"); ok {
			if err = d.Set("rdnss_service", vv); err != nil {
				return fmt.Errorf("Error reading rdnss_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rdnss_service: %v", err)
		}
	}

	if err = d.Set("subnet", flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet4thl(o["subnet"], d, "subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet"], "ObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList-Subnet"); ok {
			if err = d.Set("subnet", vv); err != nil {
				return fmt.Errorf("Error reading subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet: %v", err)
		}
	}

	if err = d.Set("upstream_interface", flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface4thl(o["upstream-interface"], d, "upstream_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["upstream-interface"], "ObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList-UpstreamInterface"); ok {
			if err = d.Set("upstream_interface", vv); err != nil {
				return fmt.Errorf("Error reading upstream_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upstream_interface: %v", err)
		}
	}

	return nil
}

func flattenObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("autonomous_flag"); ok || d.HasChange("autonomous_flag") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag4thl(d, v, "autonomous_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["autonomous-flag"] = t
		}
	}

	if v, ok := d.GetOk("delegated_prefix_iaid"); ok || d.HasChange("delegated_prefix_iaid") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid4thl(d, v, "delegated_prefix_iaid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delegated-prefix-iaid"] = t
		}
	}

	if v, ok := d.GetOk("onlink_flag"); ok || d.HasChange("onlink_flag") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag4thl(d, v, "onlink_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["onlink-flag"] = t
		}
	}

	if v, ok := d.GetOk("prefix_id"); ok || d.HasChange("prefix_id") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListPrefixId4thl(d, v, "prefix_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-id"] = t
		}
	}

	if v, ok := d.GetOk("rdnss"); ok || d.HasChange("rdnss") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnss4thl(d, v, "rdnss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rdnss"] = t
		}
	}

	if v, ok := d.GetOk("rdnss_service"); ok || d.HasChange("rdnss_service") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListRdnssService4thl(d, v, "rdnss_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rdnss-service"] = t
		}
	}

	if v, ok := d.GetOk("subnet"); ok || d.HasChange("subnet") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListSubnet4thl(d, v, "subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet"] = t
		}
	}

	if v, ok := d.GetOk("upstream_interface"); ok || d.HasChange("upstream_interface") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface4thl(d, v, "upstream_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upstream-interface"] = t
		}
	}

	return &obj, nil
}
