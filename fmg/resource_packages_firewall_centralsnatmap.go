// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure central SNAT policies.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePackagesFirewallCentralSnatMap() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesFirewallCentralSnatMapCreate,
		Read:   resourcePackagesFirewallCentralSnatMapRead,
		Update: resourcePackagesFirewallCentralSnatMapUpdate,
		Delete: resourcePackagesFirewallCentralSnatMapDelete,

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
			"pkg_folder_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"pkg": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dst_addr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat_ippool": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"nat_ippool6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"nat_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"orig_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"orig_addr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"orig_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"port_preserve": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_random": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
		},
	}
}

func resourcePackagesFirewallCentralSnatMapCreate(d *schema.ResourceData, m interface{}) error {
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

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	obj, err := getObjectPackagesFirewallCentralSnatMap(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallCentralSnatMap resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	v, err := c.CreatePackagesFirewallCentralSnatMap(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallCentralSnatMap resource: %v", err)
	}

	if v != nil && v["policyid"] != nil {
		if vidn, ok := v["policyid"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourcePackagesFirewallCentralSnatMapRead(d, m)
		} else {
			return fmt.Errorf("Error creating PackagesFirewallCentralSnatMap resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallCentralSnatMapRead(d, m)
}

func resourcePackagesFirewallCentralSnatMapUpdate(d *schema.ResourceData, m interface{}) error {
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

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	obj, err := getObjectPackagesFirewallCentralSnatMap(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallCentralSnatMap resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdatePackagesFirewallCentralSnatMap(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallCentralSnatMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallCentralSnatMapRead(d, m)
}

func resourcePackagesFirewallCentralSnatMapDelete(d *schema.ResourceData, m interface{}) error {
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

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	wsParams["adom"] = adomv

	err = c.DeletePackagesFirewallCentralSnatMap(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesFirewallCentralSnatMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesFirewallCentralSnatMapRead(d *schema.ResourceData, m interface{}) error {
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

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	if pkg_folder_path == "" {
		pkg_folder_path = importOptionChecking(m.(*FortiClient).Cfg, "pkg_folder_path")
	}
	if pkg == "" {
		pkg = importOptionChecking(m.(*FortiClient).Cfg, "pkg")
		if pkg == "" {
			return fmt.Errorf("Parameter pkg is missing")
		}
		if err = d.Set("pkg", pkg); err != nil {
			return fmt.Errorf("Error set params pkg: %v", err)
		}
	}
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	o, err := c.ReadPackagesFirewallCentralSnatMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallCentralSnatMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesFirewallCentralSnatMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallCentralSnatMap resource from API: %v", err)
	}
	return nil
}

func flattenPackagesFirewallCentralSnatMapComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallCentralSnatMapDstAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallCentralSnatMapDstAddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenPackagesFirewallCentralSnatMapDstPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallCentralSnatMapDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallCentralSnatMapNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallCentralSnatMapNatIppool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallCentralSnatMapNatIppool6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallCentralSnatMapNatPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallCentralSnatMapNat46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallCentralSnatMapNat64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallCentralSnatMapOrigAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallCentralSnatMapOrigAddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenPackagesFirewallCentralSnatMapOrigPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenPackagesFirewallCentralSnatMapPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallCentralSnatMapPortPreserve(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallCentralSnatMapPortRandom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallCentralSnatMapProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallCentralSnatMapSrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallCentralSnatMapStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallCentralSnatMapType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallCentralSnatMapUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesFirewallCentralSnatMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comments", flattenPackagesFirewallCentralSnatMapComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "PackagesFirewallCentralSnatMap-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dst_addr", flattenPackagesFirewallCentralSnatMapDstAddr(o["dst-addr"], d, "dst_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-addr"], "PackagesFirewallCentralSnatMap-DstAddr"); ok {
			if err = d.Set("dst_addr", vv); err != nil {
				return fmt.Errorf("Error reading dst_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_addr: %v", err)
		}
	}

	if err = d.Set("dst_addr6", flattenPackagesFirewallCentralSnatMapDstAddr6(o["dst-addr6"], d, "dst_addr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-addr6"], "PackagesFirewallCentralSnatMap-DstAddr6"); ok {
			if err = d.Set("dst_addr6", vv); err != nil {
				return fmt.Errorf("Error reading dst_addr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_addr6: %v", err)
		}
	}

	if err = d.Set("dst_port", flattenPackagesFirewallCentralSnatMapDstPort(o["dst-port"], d, "dst_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-port"], "PackagesFirewallCentralSnatMap-DstPort"); ok {
			if err = d.Set("dst_port", vv); err != nil {
				return fmt.Errorf("Error reading dst_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_port: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenPackagesFirewallCentralSnatMapDstintf(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "PackagesFirewallCentralSnatMap-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("nat", flattenPackagesFirewallCentralSnatMapNat(o["nat"], d, "nat")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat"], "PackagesFirewallCentralSnatMap-Nat"); ok {
			if err = d.Set("nat", vv); err != nil {
				return fmt.Errorf("Error reading nat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat: %v", err)
		}
	}

	if err = d.Set("nat_ippool", flattenPackagesFirewallCentralSnatMapNatIppool(o["nat-ippool"], d, "nat_ippool")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-ippool"], "PackagesFirewallCentralSnatMap-NatIppool"); ok {
			if err = d.Set("nat_ippool", vv); err != nil {
				return fmt.Errorf("Error reading nat_ippool: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_ippool: %v", err)
		}
	}

	if err = d.Set("nat_ippool6", flattenPackagesFirewallCentralSnatMapNatIppool6(o["nat-ippool6"], d, "nat_ippool6")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-ippool6"], "PackagesFirewallCentralSnatMap-NatIppool6"); ok {
			if err = d.Set("nat_ippool6", vv); err != nil {
				return fmt.Errorf("Error reading nat_ippool6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_ippool6: %v", err)
		}
	}

	if err = d.Set("nat_port", flattenPackagesFirewallCentralSnatMapNatPort(o["nat-port"], d, "nat_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-port"], "PackagesFirewallCentralSnatMap-NatPort"); ok {
			if err = d.Set("nat_port", vv); err != nil {
				return fmt.Errorf("Error reading nat_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_port: %v", err)
		}
	}

	if err = d.Set("nat46", flattenPackagesFirewallCentralSnatMapNat46(o["nat46"], d, "nat46")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat46"], "PackagesFirewallCentralSnatMap-Nat46"); ok {
			if err = d.Set("nat46", vv); err != nil {
				return fmt.Errorf("Error reading nat46: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat46: %v", err)
		}
	}

	if err = d.Set("nat64", flattenPackagesFirewallCentralSnatMapNat64(o["nat64"], d, "nat64")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat64"], "PackagesFirewallCentralSnatMap-Nat64"); ok {
			if err = d.Set("nat64", vv); err != nil {
				return fmt.Errorf("Error reading nat64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat64: %v", err)
		}
	}

	if err = d.Set("orig_addr", flattenPackagesFirewallCentralSnatMapOrigAddr(o["orig-addr"], d, "orig_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["orig-addr"], "PackagesFirewallCentralSnatMap-OrigAddr"); ok {
			if err = d.Set("orig_addr", vv); err != nil {
				return fmt.Errorf("Error reading orig_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading orig_addr: %v", err)
		}
	}

	if err = d.Set("orig_addr6", flattenPackagesFirewallCentralSnatMapOrigAddr6(o["orig-addr6"], d, "orig_addr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["orig-addr6"], "PackagesFirewallCentralSnatMap-OrigAddr6"); ok {
			if err = d.Set("orig_addr6", vv); err != nil {
				return fmt.Errorf("Error reading orig_addr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading orig_addr6: %v", err)
		}
	}

	if err = d.Set("orig_port", flattenPackagesFirewallCentralSnatMapOrigPort(o["orig-port"], d, "orig_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["orig-port"], "PackagesFirewallCentralSnatMap-OrigPort"); ok {
			if err = d.Set("orig_port", vv); err != nil {
				return fmt.Errorf("Error reading orig_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading orig_port: %v", err)
		}
	}

	if err = d.Set("policyid", flattenPackagesFirewallCentralSnatMapPolicyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "PackagesFirewallCentralSnatMap-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("port_preserve", flattenPackagesFirewallCentralSnatMapPortPreserve(o["port-preserve"], d, "port_preserve")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-preserve"], "PackagesFirewallCentralSnatMap-PortPreserve"); ok {
			if err = d.Set("port_preserve", vv); err != nil {
				return fmt.Errorf("Error reading port_preserve: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_preserve: %v", err)
		}
	}

	if err = d.Set("port_random", flattenPackagesFirewallCentralSnatMapPortRandom(o["port-random"], d, "port_random")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-random"], "PackagesFirewallCentralSnatMap-PortRandom"); ok {
			if err = d.Set("port_random", vv); err != nil {
				return fmt.Errorf("Error reading port_random: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_random: %v", err)
		}
	}

	if err = d.Set("protocol", flattenPackagesFirewallCentralSnatMapProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "PackagesFirewallCentralSnatMap-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenPackagesFirewallCentralSnatMapSrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "PackagesFirewallCentralSnatMap-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesFirewallCentralSnatMapStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesFirewallCentralSnatMap-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenPackagesFirewallCentralSnatMapType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "PackagesFirewallCentralSnatMap-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenPackagesFirewallCentralSnatMapUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "PackagesFirewallCentralSnatMap-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenPackagesFirewallCentralSnatMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesFirewallCentralSnatMapComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallCentralSnatMapDstAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallCentralSnatMapDstAddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandPackagesFirewallCentralSnatMapDstPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallCentralSnatMapDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallCentralSnatMapNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallCentralSnatMapNatIppool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallCentralSnatMapNatIppool6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallCentralSnatMapNatPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallCentralSnatMapNat46(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallCentralSnatMapNat64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallCentralSnatMapOrigAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallCentralSnatMapOrigAddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandPackagesFirewallCentralSnatMapOrigPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallCentralSnatMapPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallCentralSnatMapPortPreserve(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallCentralSnatMapPortRandom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallCentralSnatMapProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallCentralSnatMapSrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallCentralSnatMapStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallCentralSnatMapType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallCentralSnatMapUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesFirewallCentralSnatMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandPackagesFirewallCentralSnatMapComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dst_addr"); ok || d.HasChange("dst_addr") {
		t, err := expandPackagesFirewallCentralSnatMapDstAddr(d, v, "dst_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-addr"] = t
		}
	}

	if v, ok := d.GetOk("dst_addr6"); ok || d.HasChange("dst_addr6") {
		t, err := expandPackagesFirewallCentralSnatMapDstAddr6(d, v, "dst_addr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-addr6"] = t
		}
	}

	if v, ok := d.GetOk("dst_port"); ok || d.HasChange("dst_port") {
		t, err := expandPackagesFirewallCentralSnatMapDstPort(d, v, "dst_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-port"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandPackagesFirewallCentralSnatMapDstintf(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("nat"); ok || d.HasChange("nat") {
		t, err := expandPackagesFirewallCentralSnatMapNat(d, v, "nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat"] = t
		}
	}

	if v, ok := d.GetOk("nat_ippool"); ok || d.HasChange("nat_ippool") {
		t, err := expandPackagesFirewallCentralSnatMapNatIppool(d, v, "nat_ippool")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-ippool"] = t
		}
	}

	if v, ok := d.GetOk("nat_ippool6"); ok || d.HasChange("nat_ippool6") {
		t, err := expandPackagesFirewallCentralSnatMapNatIppool6(d, v, "nat_ippool6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-ippool6"] = t
		}
	}

	if v, ok := d.GetOk("nat_port"); ok || d.HasChange("nat_port") {
		t, err := expandPackagesFirewallCentralSnatMapNatPort(d, v, "nat_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-port"] = t
		}
	}

	if v, ok := d.GetOk("nat46"); ok || d.HasChange("nat46") {
		t, err := expandPackagesFirewallCentralSnatMapNat46(d, v, "nat46")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46"] = t
		}
	}

	if v, ok := d.GetOk("nat64"); ok || d.HasChange("nat64") {
		t, err := expandPackagesFirewallCentralSnatMapNat64(d, v, "nat64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat64"] = t
		}
	}

	if v, ok := d.GetOk("orig_addr"); ok || d.HasChange("orig_addr") {
		t, err := expandPackagesFirewallCentralSnatMapOrigAddr(d, v, "orig_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["orig-addr"] = t
		}
	}

	if v, ok := d.GetOk("orig_addr6"); ok || d.HasChange("orig_addr6") {
		t, err := expandPackagesFirewallCentralSnatMapOrigAddr6(d, v, "orig_addr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["orig-addr6"] = t
		}
	}

	if v, ok := d.GetOk("orig_port"); ok || d.HasChange("orig_port") {
		t, err := expandPackagesFirewallCentralSnatMapOrigPort(d, v, "orig_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["orig-port"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandPackagesFirewallCentralSnatMapPolicyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("port_preserve"); ok || d.HasChange("port_preserve") {
		t, err := expandPackagesFirewallCentralSnatMapPortPreserve(d, v, "port_preserve")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-preserve"] = t
		}
	}

	if v, ok := d.GetOk("port_random"); ok || d.HasChange("port_random") {
		t, err := expandPackagesFirewallCentralSnatMapPortRandom(d, v, "port_random")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-random"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandPackagesFirewallCentralSnatMapProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandPackagesFirewallCentralSnatMapSrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandPackagesFirewallCentralSnatMapStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandPackagesFirewallCentralSnatMapType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandPackagesFirewallCentralSnatMapUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
