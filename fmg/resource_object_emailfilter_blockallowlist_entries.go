// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Anti-spam block/allow entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectEmailfilterBlockAllowListEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectEmailfilterBlockAllowListEntriesCreate,
		Read:   resourceObjectEmailfilterBlockAllowListEntriesRead,
		Update: resourceObjectEmailfilterBlockAllowListEntriesUpdate,
		Delete: resourceObjectEmailfilterBlockAllowListEntriesDelete,

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
			"block_allow_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip4_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pattern_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectEmailfilterBlockAllowListEntriesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	block_allow_list := d.Get("block_allow_list").(string)
	paradict["block_allow_list"] = block_allow_list

	obj, err := getObjectObjectEmailfilterBlockAllowListEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectEmailfilterBlockAllowListEntries resource while getting object: %v", err)
	}

	_, err = c.CreateObjectEmailfilterBlockAllowListEntries(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectEmailfilterBlockAllowListEntries resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectEmailfilterBlockAllowListEntriesRead(d, m)
}

func resourceObjectEmailfilterBlockAllowListEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	block_allow_list := d.Get("block_allow_list").(string)
	paradict["block_allow_list"] = block_allow_list

	obj, err := getObjectObjectEmailfilterBlockAllowListEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEmailfilterBlockAllowListEntries resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectEmailfilterBlockAllowListEntries(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEmailfilterBlockAllowListEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectEmailfilterBlockAllowListEntriesRead(d, m)
}

func resourceObjectEmailfilterBlockAllowListEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	block_allow_list := d.Get("block_allow_list").(string)
	paradict["block_allow_list"] = block_allow_list

	err = c.DeleteObjectEmailfilterBlockAllowListEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectEmailfilterBlockAllowListEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectEmailfilterBlockAllowListEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	block_allow_list := d.Get("block_allow_list").(string)
	if block_allow_list == "" {
		block_allow_list = importOptionChecking(m.(*FortiClient).Cfg, "block_allow_list")
		if block_allow_list == "" {
			return fmt.Errorf("Parameter block_allow_list is missing")
		}
		if err = d.Set("block_allow_list", block_allow_list); err != nil {
			return fmt.Errorf("Error set params block_allow_list: %v", err)
		}
	}
	paradict["block_allow_list"] = block_allow_list

	o, err := c.ReadObjectEmailfilterBlockAllowListEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEmailfilterBlockAllowListEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectEmailfilterBlockAllowListEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEmailfilterBlockAllowListEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectEmailfilterBlockAllowListEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterBlockAllowListEntriesAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterBlockAllowListEntriesEmailPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterBlockAllowListEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterBlockAllowListEntriesIp4Subnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenObjectEmailfilterBlockAllowListEntriesIp6Subnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterBlockAllowListEntriesPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterBlockAllowListEntriesPatternType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterBlockAllowListEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterBlockAllowListEntriesType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectEmailfilterBlockAllowListEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectEmailfilterBlockAllowListEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectEmailfilterBlockAllowListEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("addr_type", flattenObjectEmailfilterBlockAllowListEntriesAddrType2edl(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "ObjectEmailfilterBlockAllowListEntries-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("email_pattern", flattenObjectEmailfilterBlockAllowListEntriesEmailPattern2edl(o["email-pattern"], d, "email_pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-pattern"], "ObjectEmailfilterBlockAllowListEntries-EmailPattern"); ok {
			if err = d.Set("email_pattern", vv); err != nil {
				return fmt.Errorf("Error reading email_pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_pattern: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectEmailfilterBlockAllowListEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectEmailfilterBlockAllowListEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip4_subnet", flattenObjectEmailfilterBlockAllowListEntriesIp4Subnet2edl(o["ip4-subnet"], d, "ip4_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip4-subnet"], "ObjectEmailfilterBlockAllowListEntries-Ip4Subnet"); ok {
			if err = d.Set("ip4_subnet", vv); err != nil {
				return fmt.Errorf("Error reading ip4_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip4_subnet: %v", err)
		}
	}

	if err = d.Set("ip6_subnet", flattenObjectEmailfilterBlockAllowListEntriesIp6Subnet2edl(o["ip6-subnet"], d, "ip6_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-subnet"], "ObjectEmailfilterBlockAllowListEntries-Ip6Subnet"); ok {
			if err = d.Set("ip6_subnet", vv); err != nil {
				return fmt.Errorf("Error reading ip6_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_subnet: %v", err)
		}
	}

	if err = d.Set("pattern", flattenObjectEmailfilterBlockAllowListEntriesPattern2edl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "ObjectEmailfilterBlockAllowListEntries-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("pattern_type", flattenObjectEmailfilterBlockAllowListEntriesPatternType2edl(o["pattern-type"], d, "pattern_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern-type"], "ObjectEmailfilterBlockAllowListEntries-PatternType"); ok {
			if err = d.Set("pattern_type", vv); err != nil {
				return fmt.Errorf("Error reading pattern_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern_type: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectEmailfilterBlockAllowListEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectEmailfilterBlockAllowListEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectEmailfilterBlockAllowListEntriesType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectEmailfilterBlockAllowListEntries-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectEmailfilterBlockAllowListEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectEmailfilterBlockAllowListEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBlockAllowListEntriesAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBlockAllowListEntriesEmailPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBlockAllowListEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBlockAllowListEntriesIp4Subnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBlockAllowListEntriesIp6Subnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBlockAllowListEntriesPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBlockAllowListEntriesPatternType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBlockAllowListEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBlockAllowListEntriesType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectEmailfilterBlockAllowListEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectEmailfilterBlockAllowListEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandObjectEmailfilterBlockAllowListEntriesAddrType2edl(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("email_pattern"); ok || d.HasChange("email_pattern") {
		t, err := expandObjectEmailfilterBlockAllowListEntriesEmailPattern2edl(d, v, "email_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-pattern"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectEmailfilterBlockAllowListEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip4_subnet"); ok || d.HasChange("ip4_subnet") {
		t, err := expandObjectEmailfilterBlockAllowListEntriesIp4Subnet2edl(d, v, "ip4_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip4-subnet"] = t
		}
	}

	if v, ok := d.GetOk("ip6_subnet"); ok || d.HasChange("ip6_subnet") {
		t, err := expandObjectEmailfilterBlockAllowListEntriesIp6Subnet2edl(d, v, "ip6_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-subnet"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandObjectEmailfilterBlockAllowListEntriesPattern2edl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("pattern_type"); ok || d.HasChange("pattern_type") {
		t, err := expandObjectEmailfilterBlockAllowListEntriesPatternType2edl(d, v, "pattern_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern-type"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectEmailfilterBlockAllowListEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectEmailfilterBlockAllowListEntriesType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
