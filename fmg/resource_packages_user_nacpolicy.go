// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure NAC policy matching pattern to identify matching NAC devices.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePackagesUserNacPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesUserNacPolicyCreate,
		Read:   resourcePackagesUserNacPolicyRead,
		Update: resourcePackagesUserNacPolicyUpdate,
		Delete: resourcePackagesUserNacPolicyDelete,

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
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ems_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"family": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"firewall_address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fortivoice_tag": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hw_vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hw_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"match_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"os": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssid_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sw_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_fortilink": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"switch_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"switch_mac_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"switch_scope": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourcePackagesUserNacPolicyCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	obj, err := getObjectPackagesUserNacPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesUserNacPolicy resource while getting object: %v", err)
	}

	_, err = c.CreatePackagesUserNacPolicy(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating PackagesUserNacPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourcePackagesUserNacPolicyRead(d, m)
}

func resourcePackagesUserNacPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	obj, err := getObjectPackagesUserNacPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesUserNacPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesUserNacPolicy(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating PackagesUserNacPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourcePackagesUserNacPolicyRead(d, m)
}

func resourcePackagesUserNacPolicyDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	err = c.DeletePackagesUserNacPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesUserNacPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesUserNacPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadPackagesUserNacPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading PackagesUserNacPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesUserNacPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesUserNacPolicy resource from API: %v", err)
	}
	return nil
}

func flattenPackagesUserNacPolicyCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyEmsTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenPackagesUserNacPolicyFamily(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyFirewallAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesUserNacPolicyFortivoiceTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesUserNacPolicyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyHwVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyHwVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyMatchPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyMatchType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyOs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicySeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenPackagesUserNacPolicySrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicySsidPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenPackagesUserNacPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicySwVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicySwitchFortilink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesUserNacPolicySwitchGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesUserNacPolicySwitchMacPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesUserNacPolicySwitchScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesUserNacPolicyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesUserNacPolicyUserGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectPackagesUserNacPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("category", flattenPackagesUserNacPolicyCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "PackagesUserNacPolicy-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("description", flattenPackagesUserNacPolicyDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "PackagesUserNacPolicy-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("ems_tag", flattenPackagesUserNacPolicyEmsTag(o["ems-tag"], d, "ems_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ems-tag"], "PackagesUserNacPolicy-EmsTag"); ok {
			if err = d.Set("ems_tag", vv); err != nil {
				return fmt.Errorf("Error reading ems_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ems_tag: %v", err)
		}
	}

	if err = d.Set("family", flattenPackagesUserNacPolicyFamily(o["family"], d, "family")); err != nil {
		if vv, ok := fortiAPIPatch(o["family"], "PackagesUserNacPolicy-Family"); ok {
			if err = d.Set("family", vv); err != nil {
				return fmt.Errorf("Error reading family: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading family: %v", err)
		}
	}

	if err = d.Set("firewall_address", flattenPackagesUserNacPolicyFirewallAddress(o["firewall-address"], d, "firewall_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["firewall-address"], "PackagesUserNacPolicy-FirewallAddress"); ok {
			if err = d.Set("firewall_address", vv); err != nil {
				return fmt.Errorf("Error reading firewall_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firewall_address: %v", err)
		}
	}

	if err = d.Set("fortivoice_tag", flattenPackagesUserNacPolicyFortivoiceTag(o["fortivoice-tag"], d, "fortivoice_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortivoice-tag"], "PackagesUserNacPolicy-FortivoiceTag"); ok {
			if err = d.Set("fortivoice_tag", vv); err != nil {
				return fmt.Errorf("Error reading fortivoice_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortivoice_tag: %v", err)
		}
	}

	if err = d.Set("host", flattenPackagesUserNacPolicyHost(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "PackagesUserNacPolicy-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("hw_vendor", flattenPackagesUserNacPolicyHwVendor(o["hw-vendor"], d, "hw_vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-vendor"], "PackagesUserNacPolicy-HwVendor"); ok {
			if err = d.Set("hw_vendor", vv); err != nil {
				return fmt.Errorf("Error reading hw_vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_vendor: %v", err)
		}
	}

	if err = d.Set("hw_version", flattenPackagesUserNacPolicyHwVersion(o["hw-version"], d, "hw_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-version"], "PackagesUserNacPolicy-HwVersion"); ok {
			if err = d.Set("hw_version", vv); err != nil {
				return fmt.Errorf("Error reading hw_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_version: %v", err)
		}
	}

	if err = d.Set("mac", flattenPackagesUserNacPolicyMac(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "PackagesUserNacPolicy-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("match_period", flattenPackagesUserNacPolicyMatchPeriod(o["match-period"], d, "match_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-period"], "PackagesUserNacPolicy-MatchPeriod"); ok {
			if err = d.Set("match_period", vv); err != nil {
				return fmt.Errorf("Error reading match_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_period: %v", err)
		}
	}

	if err = d.Set("match_type", flattenPackagesUserNacPolicyMatchType(o["match-type"], d, "match_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-type"], "PackagesUserNacPolicy-MatchType"); ok {
			if err = d.Set("match_type", vv); err != nil {
				return fmt.Errorf("Error reading match_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_type: %v", err)
		}
	}

	if err = d.Set("name", flattenPackagesUserNacPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesUserNacPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("os", flattenPackagesUserNacPolicyOs(o["os"], d, "os")); err != nil {
		if vv, ok := fortiAPIPatch(o["os"], "PackagesUserNacPolicy-Os"); ok {
			if err = d.Set("os", vv); err != nil {
				return fmt.Errorf("Error reading os: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os: %v", err)
		}
	}

	if err = d.Set("severity", flattenPackagesUserNacPolicySeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "PackagesUserNacPolicy-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("src", flattenPackagesUserNacPolicySrc(o["src"], d, "src")); err != nil {
		if vv, ok := fortiAPIPatch(o["src"], "PackagesUserNacPolicy-Src"); ok {
			if err = d.Set("src", vv); err != nil {
				return fmt.Errorf("Error reading src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	if err = d.Set("ssid_policy", flattenPackagesUserNacPolicySsidPolicy(o["ssid-policy"], d, "ssid_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssid-policy"], "PackagesUserNacPolicy-SsidPolicy"); ok {
			if err = d.Set("ssid_policy", vv); err != nil {
				return fmt.Errorf("Error reading ssid_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssid_policy: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesUserNacPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesUserNacPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("sw_version", flattenPackagesUserNacPolicySwVersion(o["sw-version"], d, "sw_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["sw-version"], "PackagesUserNacPolicy-SwVersion"); ok {
			if err = d.Set("sw_version", vv); err != nil {
				return fmt.Errorf("Error reading sw_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sw_version: %v", err)
		}
	}

	if err = d.Set("switch_fortilink", flattenPackagesUserNacPolicySwitchFortilink(o["switch-fortilink"], d, "switch_fortilink")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-fortilink"], "PackagesUserNacPolicy-SwitchFortilink"); ok {
			if err = d.Set("switch_fortilink", vv); err != nil {
				return fmt.Errorf("Error reading switch_fortilink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_fortilink: %v", err)
		}
	}

	if err = d.Set("switch_group", flattenPackagesUserNacPolicySwitchGroup(o["switch-group"], d, "switch_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-group"], "PackagesUserNacPolicy-SwitchGroup"); ok {
			if err = d.Set("switch_group", vv); err != nil {
				return fmt.Errorf("Error reading switch_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_group: %v", err)
		}
	}

	if err = d.Set("switch_mac_policy", flattenPackagesUserNacPolicySwitchMacPolicy(o["switch-mac-policy"], d, "switch_mac_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-mac-policy"], "PackagesUserNacPolicy-SwitchMacPolicy"); ok {
			if err = d.Set("switch_mac_policy", vv); err != nil {
				return fmt.Errorf("Error reading switch_mac_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_mac_policy: %v", err)
		}
	}

	if err = d.Set("switch_scope", flattenPackagesUserNacPolicySwitchScope(o["switch-scope"], d, "switch_scope")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-scope"], "PackagesUserNacPolicy-SwitchScope"); ok {
			if err = d.Set("switch_scope", vv); err != nil {
				return fmt.Errorf("Error reading switch_scope: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_scope: %v", err)
		}
	}

	if err = d.Set("type", flattenPackagesUserNacPolicyType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "PackagesUserNacPolicy-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user", flattenPackagesUserNacPolicyUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "PackagesUserNacPolicy-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("user_group", flattenPackagesUserNacPolicyUserGroup(o["user-group"], d, "user_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-group"], "PackagesUserNacPolicy-UserGroup"); ok {
			if err = d.Set("user_group", vv); err != nil {
				return fmt.Errorf("Error reading user_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_group: %v", err)
		}
	}

	return nil
}

func flattenPackagesUserNacPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesUserNacPolicyCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyEmsTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandPackagesUserNacPolicyFamily(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyFirewallAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesUserNacPolicyFortivoiceTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesUserNacPolicyHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyHwVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyHwVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyMatchPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyMatchType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyOs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicySeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandPackagesUserNacPolicySrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicySsidPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandPackagesUserNacPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicySwVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicySwitchFortilink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesUserNacPolicySwitchGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesUserNacPolicySwitchMacPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesUserNacPolicySwitchScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesUserNacPolicyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesUserNacPolicyUserGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectPackagesUserNacPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandPackagesUserNacPolicyCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandPackagesUserNacPolicyDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("ems_tag"); ok || d.HasChange("ems_tag") {
		t, err := expandPackagesUserNacPolicyEmsTag(d, v, "ems_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-tag"] = t
		}
	}

	if v, ok := d.GetOk("family"); ok || d.HasChange("family") {
		t, err := expandPackagesUserNacPolicyFamily(d, v, "family")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["family"] = t
		}
	}

	if v, ok := d.GetOk("firewall_address"); ok || d.HasChange("firewall_address") {
		t, err := expandPackagesUserNacPolicyFirewallAddress(d, v, "firewall_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-address"] = t
		}
	}

	if v, ok := d.GetOk("fortivoice_tag"); ok || d.HasChange("fortivoice_tag") {
		t, err := expandPackagesUserNacPolicyFortivoiceTag(d, v, "fortivoice_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortivoice-tag"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandPackagesUserNacPolicyHost(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("hw_vendor"); ok || d.HasChange("hw_vendor") {
		t, err := expandPackagesUserNacPolicyHwVendor(d, v, "hw_vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-vendor"] = t
		}
	}

	if v, ok := d.GetOk("hw_version"); ok || d.HasChange("hw_version") {
		t, err := expandPackagesUserNacPolicyHwVersion(d, v, "hw_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-version"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandPackagesUserNacPolicyMac(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("match_period"); ok || d.HasChange("match_period") {
		t, err := expandPackagesUserNacPolicyMatchPeriod(d, v, "match_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-period"] = t
		}
	}

	if v, ok := d.GetOk("match_type"); ok || d.HasChange("match_type") {
		t, err := expandPackagesUserNacPolicyMatchType(d, v, "match_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandPackagesUserNacPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("os"); ok || d.HasChange("os") {
		t, err := expandPackagesUserNacPolicyOs(d, v, "os")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandPackagesUserNacPolicySeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok || d.HasChange("src") {
		t, err := expandPackagesUserNacPolicySrc(d, v, "src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("ssid_policy"); ok || d.HasChange("ssid_policy") {
		t, err := expandPackagesUserNacPolicySsidPolicy(d, v, "ssid_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssid-policy"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandPackagesUserNacPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("sw_version"); ok || d.HasChange("sw_version") {
		t, err := expandPackagesUserNacPolicySwVersion(d, v, "sw_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sw-version"] = t
		}
	}

	if v, ok := d.GetOk("switch_fortilink"); ok || d.HasChange("switch_fortilink") {
		t, err := expandPackagesUserNacPolicySwitchFortilink(d, v, "switch_fortilink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-fortilink"] = t
		}
	}

	if v, ok := d.GetOk("switch_group"); ok || d.HasChange("switch_group") {
		t, err := expandPackagesUserNacPolicySwitchGroup(d, v, "switch_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-group"] = t
		}
	}

	if v, ok := d.GetOk("switch_mac_policy"); ok || d.HasChange("switch_mac_policy") {
		t, err := expandPackagesUserNacPolicySwitchMacPolicy(d, v, "switch_mac_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-mac-policy"] = t
		}
	}

	if v, ok := d.GetOk("switch_scope"); ok || d.HasChange("switch_scope") {
		t, err := expandPackagesUserNacPolicySwitchScope(d, v, "switch_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-scope"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandPackagesUserNacPolicyType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandPackagesUserNacPolicyUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("user_group"); ok || d.HasChange("user_group") {
		t, err := expandPackagesUserNacPolicyUserGroup(d, v, "user_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group"] = t
		}
	}

	return &obj, nil
}
