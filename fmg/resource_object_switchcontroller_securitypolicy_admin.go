// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure fortiswitch's admin security-policy.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerSecurityPolicyAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerSecurityPolicyAdminCreate,
		Read:   resourceObjectSwitchControllerSecurityPolicyAdminRead,
		Update: resourceObjectSwitchControllerSecurityPolicyAdminUpdate,
		Delete: resourceObjectSwitchControllerSecurityPolicyAdminDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"auto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"trusthost1": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost10": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost2": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost3": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost4": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost5": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost6": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost7": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost8": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost9": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSwitchControllerSecurityPolicyAdminCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSwitchControllerSecurityPolicyAdmin(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerSecurityPolicyAdmin resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectSwitchControllerSecurityPolicyAdmin(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectSwitchControllerSecurityPolicyAdmin(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectSwitchControllerSecurityPolicyAdmin resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectSwitchControllerSecurityPolicyAdmin(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectSwitchControllerSecurityPolicyAdmin resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerSecurityPolicyAdminRead(d, m)
}

func resourceObjectSwitchControllerSecurityPolicyAdminUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSwitchControllerSecurityPolicyAdmin(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerSecurityPolicyAdmin resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSwitchControllerSecurityPolicyAdmin(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerSecurityPolicyAdmin resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerSecurityPolicyAdminRead(d, m)
}

func resourceObjectSwitchControllerSecurityPolicyAdminDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectSwitchControllerSecurityPolicyAdmin(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerSecurityPolicyAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerSecurityPolicyAdminRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSwitchControllerSecurityPolicyAdmin(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectSwitchControllerSecurityPolicyAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerSecurityPolicyAdmin(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerSecurityPolicyAdmin resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerSecurityPolicyAdminAuto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicyAdminName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicyAdminTrusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerSecurityPolicyAdminTrusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerSecurityPolicyAdminTrusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerSecurityPolicyAdminTrusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerSecurityPolicyAdminTrusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerSecurityPolicyAdminTrusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerSecurityPolicyAdminTrusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerSecurityPolicyAdminTrusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerSecurityPolicyAdminTrusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerSecurityPolicyAdminTrusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectSwitchControllerSecurityPolicyAdmin(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("auto", flattenObjectSwitchControllerSecurityPolicyAdminAuto(o["auto"], d, "auto")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto"], "ObjectSwitchControllerSecurityPolicyAdmin-Auto"); ok {
			if err = d.Set("auto", vv); err != nil {
				return fmt.Errorf("Error reading auto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost1", flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost1(o["ip6-trusthost1"], d, "ip6_trusthost1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost1"], "ObjectSwitchControllerSecurityPolicyAdmin-Ip6Trusthost1"); ok {
			if err = d.Set("ip6_trusthost1", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost1: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost10", flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost10(o["ip6-trusthost10"], d, "ip6_trusthost10")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost10"], "ObjectSwitchControllerSecurityPolicyAdmin-Ip6Trusthost10"); ok {
			if err = d.Set("ip6_trusthost10", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost10: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost10: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost2", flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost2(o["ip6-trusthost2"], d, "ip6_trusthost2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost2"], "ObjectSwitchControllerSecurityPolicyAdmin-Ip6Trusthost2"); ok {
			if err = d.Set("ip6_trusthost2", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost2: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost3", flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost3(o["ip6-trusthost3"], d, "ip6_trusthost3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost3"], "ObjectSwitchControllerSecurityPolicyAdmin-Ip6Trusthost3"); ok {
			if err = d.Set("ip6_trusthost3", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost3: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost4", flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost4(o["ip6-trusthost4"], d, "ip6_trusthost4")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost4"], "ObjectSwitchControllerSecurityPolicyAdmin-Ip6Trusthost4"); ok {
			if err = d.Set("ip6_trusthost4", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost4: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost5", flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost5(o["ip6-trusthost5"], d, "ip6_trusthost5")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost5"], "ObjectSwitchControllerSecurityPolicyAdmin-Ip6Trusthost5"); ok {
			if err = d.Set("ip6_trusthost5", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost5: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost6", flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost6(o["ip6-trusthost6"], d, "ip6_trusthost6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost6"], "ObjectSwitchControllerSecurityPolicyAdmin-Ip6Trusthost6"); ok {
			if err = d.Set("ip6_trusthost6", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost6: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost7", flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost7(o["ip6-trusthost7"], d, "ip6_trusthost7")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost7"], "ObjectSwitchControllerSecurityPolicyAdmin-Ip6Trusthost7"); ok {
			if err = d.Set("ip6_trusthost7", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost7: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost8", flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost8(o["ip6-trusthost8"], d, "ip6_trusthost8")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost8"], "ObjectSwitchControllerSecurityPolicyAdmin-Ip6Trusthost8"); ok {
			if err = d.Set("ip6_trusthost8", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost8: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost9", flattenObjectSwitchControllerSecurityPolicyAdminIp6Trusthost9(o["ip6-trusthost9"], d, "ip6_trusthost9")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost9"], "ObjectSwitchControllerSecurityPolicyAdmin-Ip6Trusthost9"); ok {
			if err = d.Set("ip6_trusthost9", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost9: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost9: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerSecurityPolicyAdminName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerSecurityPolicyAdmin-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("trusthost1", flattenObjectSwitchControllerSecurityPolicyAdminTrusthost1(o["trusthost1"], d, "trusthost1")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost1"], "ObjectSwitchControllerSecurityPolicyAdmin-Trusthost1"); ok {
			if err = d.Set("trusthost1", vv); err != nil {
				return fmt.Errorf("Error reading trusthost1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost1: %v", err)
		}
	}

	if err = d.Set("trusthost10", flattenObjectSwitchControllerSecurityPolicyAdminTrusthost10(o["trusthost10"], d, "trusthost10")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost10"], "ObjectSwitchControllerSecurityPolicyAdmin-Trusthost10"); ok {
			if err = d.Set("trusthost10", vv); err != nil {
				return fmt.Errorf("Error reading trusthost10: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost10: %v", err)
		}
	}

	if err = d.Set("trusthost2", flattenObjectSwitchControllerSecurityPolicyAdminTrusthost2(o["trusthost2"], d, "trusthost2")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost2"], "ObjectSwitchControllerSecurityPolicyAdmin-Trusthost2"); ok {
			if err = d.Set("trusthost2", vv); err != nil {
				return fmt.Errorf("Error reading trusthost2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost2: %v", err)
		}
	}

	if err = d.Set("trusthost3", flattenObjectSwitchControllerSecurityPolicyAdminTrusthost3(o["trusthost3"], d, "trusthost3")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost3"], "ObjectSwitchControllerSecurityPolicyAdmin-Trusthost3"); ok {
			if err = d.Set("trusthost3", vv); err != nil {
				return fmt.Errorf("Error reading trusthost3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost3: %v", err)
		}
	}

	if err = d.Set("trusthost4", flattenObjectSwitchControllerSecurityPolicyAdminTrusthost4(o["trusthost4"], d, "trusthost4")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost4"], "ObjectSwitchControllerSecurityPolicyAdmin-Trusthost4"); ok {
			if err = d.Set("trusthost4", vv); err != nil {
				return fmt.Errorf("Error reading trusthost4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost4: %v", err)
		}
	}

	if err = d.Set("trusthost5", flattenObjectSwitchControllerSecurityPolicyAdminTrusthost5(o["trusthost5"], d, "trusthost5")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost5"], "ObjectSwitchControllerSecurityPolicyAdmin-Trusthost5"); ok {
			if err = d.Set("trusthost5", vv); err != nil {
				return fmt.Errorf("Error reading trusthost5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost5: %v", err)
		}
	}

	if err = d.Set("trusthost6", flattenObjectSwitchControllerSecurityPolicyAdminTrusthost6(o["trusthost6"], d, "trusthost6")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost6"], "ObjectSwitchControllerSecurityPolicyAdmin-Trusthost6"); ok {
			if err = d.Set("trusthost6", vv); err != nil {
				return fmt.Errorf("Error reading trusthost6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost6: %v", err)
		}
	}

	if err = d.Set("trusthost7", flattenObjectSwitchControllerSecurityPolicyAdminTrusthost7(o["trusthost7"], d, "trusthost7")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost7"], "ObjectSwitchControllerSecurityPolicyAdmin-Trusthost7"); ok {
			if err = d.Set("trusthost7", vv); err != nil {
				return fmt.Errorf("Error reading trusthost7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost7: %v", err)
		}
	}

	if err = d.Set("trusthost8", flattenObjectSwitchControllerSecurityPolicyAdminTrusthost8(o["trusthost8"], d, "trusthost8")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost8"], "ObjectSwitchControllerSecurityPolicyAdmin-Trusthost8"); ok {
			if err = d.Set("trusthost8", vv); err != nil {
				return fmt.Errorf("Error reading trusthost8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost8: %v", err)
		}
	}

	if err = d.Set("trusthost9", flattenObjectSwitchControllerSecurityPolicyAdminTrusthost9(o["trusthost9"], d, "trusthost9")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost9"], "ObjectSwitchControllerSecurityPolicyAdmin-Trusthost9"); ok {
			if err = d.Set("trusthost9", vv); err != nil {
				return fmt.Errorf("Error reading trusthost9: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost9: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerSecurityPolicyAdminFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerSecurityPolicyAdminAuto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicyAdminName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicyAdminTrusthost1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectSwitchControllerSecurityPolicyAdminTrusthost10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectSwitchControllerSecurityPolicyAdminTrusthost2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectSwitchControllerSecurityPolicyAdminTrusthost3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectSwitchControllerSecurityPolicyAdminTrusthost4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectSwitchControllerSecurityPolicyAdminTrusthost5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectSwitchControllerSecurityPolicyAdminTrusthost6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectSwitchControllerSecurityPolicyAdminTrusthost7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectSwitchControllerSecurityPolicyAdminTrusthost8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectSwitchControllerSecurityPolicyAdminTrusthost9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectObjectSwitchControllerSecurityPolicyAdmin(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto"); ok || d.HasChange("auto") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminAuto(d, v, "auto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost1"); ok || d.HasChange("ip6_trusthost1") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost1(d, v, "ip6_trusthost1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost10"); ok || d.HasChange("ip6_trusthost10") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost10(d, v, "ip6_trusthost10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost10"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost2"); ok || d.HasChange("ip6_trusthost2") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost2(d, v, "ip6_trusthost2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost3"); ok || d.HasChange("ip6_trusthost3") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost3(d, v, "ip6_trusthost3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost4"); ok || d.HasChange("ip6_trusthost4") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost4(d, v, "ip6_trusthost4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost5"); ok || d.HasChange("ip6_trusthost5") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost5(d, v, "ip6_trusthost5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost6"); ok || d.HasChange("ip6_trusthost6") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost6(d, v, "ip6_trusthost6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost7"); ok || d.HasChange("ip6_trusthost7") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost7(d, v, "ip6_trusthost7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost8"); ok || d.HasChange("ip6_trusthost8") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost8(d, v, "ip6_trusthost8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost9"); ok || d.HasChange("ip6_trusthost9") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminIp6Trusthost9(d, v, "ip6_trusthost9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost9"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("trusthost1"); ok || d.HasChange("trusthost1") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminTrusthost1(d, v, "trusthost1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("trusthost10"); ok || d.HasChange("trusthost10") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminTrusthost10(d, v, "trusthost10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost10"] = t
		}
	}

	if v, ok := d.GetOk("trusthost2"); ok || d.HasChange("trusthost2") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminTrusthost2(d, v, "trusthost2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("trusthost3"); ok || d.HasChange("trusthost3") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminTrusthost3(d, v, "trusthost3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("trusthost4"); ok || d.HasChange("trusthost4") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminTrusthost4(d, v, "trusthost4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("trusthost5"); ok || d.HasChange("trusthost5") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminTrusthost5(d, v, "trusthost5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("trusthost6"); ok || d.HasChange("trusthost6") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminTrusthost6(d, v, "trusthost6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("trusthost7"); ok || d.HasChange("trusthost7") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminTrusthost7(d, v, "trusthost7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("trusthost8"); ok || d.HasChange("trusthost8") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminTrusthost8(d, v, "trusthost8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("trusthost9"); ok || d.HasChange("trusthost9") {
		t, err := expandObjectSwitchControllerSecurityPolicyAdminTrusthost9(d, v, "trusthost9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost9"] = t
		}
	}

	return &obj, nil
}
