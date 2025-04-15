// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Spam filter trusted IP addresses.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectEmailfilterIptrustEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectEmailfilterIptrustEntriesCreate,
		Read:   resourceObjectEmailfilterIptrustEntriesRead,
		Update: resourceObjectEmailfilterIptrustEntriesUpdate,
		Delete: resourceObjectEmailfilterIptrustEntriesDelete,

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
			"iptrust": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"addr_type": &schema.Schema{
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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectEmailfilterIptrustEntriesCreate(d *schema.ResourceData, m interface{}) error {
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

	iptrust := d.Get("iptrust").(string)
	paradict["iptrust"] = iptrust

	obj, err := getObjectObjectEmailfilterIptrustEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectEmailfilterIptrustEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectEmailfilterIptrustEntries(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectEmailfilterIptrustEntries resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectEmailfilterIptrustEntriesRead(d, m)
}

func resourceObjectEmailfilterIptrustEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	iptrust := d.Get("iptrust").(string)
	paradict["iptrust"] = iptrust

	obj, err := getObjectObjectEmailfilterIptrustEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEmailfilterIptrustEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectEmailfilterIptrustEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEmailfilterIptrustEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectEmailfilterIptrustEntriesRead(d, m)
}

func resourceObjectEmailfilterIptrustEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	iptrust := d.Get("iptrust").(string)
	paradict["iptrust"] = iptrust

	wsParams["adom"] = adomv

	err = c.DeleteObjectEmailfilterIptrustEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectEmailfilterIptrustEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectEmailfilterIptrustEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	iptrust := d.Get("iptrust").(string)
	if iptrust == "" {
		iptrust = importOptionChecking(m.(*FortiClient).Cfg, "iptrust")
		if iptrust == "" {
			return fmt.Errorf("Parameter iptrust is missing")
		}
		if err = d.Set("iptrust", iptrust); err != nil {
			return fmt.Errorf("Error set params iptrust: %v", err)
		}
	}
	paradict["iptrust"] = iptrust

	o, err := c.ReadObjectEmailfilterIptrustEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEmailfilterIptrustEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectEmailfilterIptrustEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEmailfilterIptrustEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectEmailfilterIptrustEntriesAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterIptrustEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterIptrustEntriesIp4Subnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectEmailfilterIptrustEntriesIp6Subnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterIptrustEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectEmailfilterIptrustEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("addr_type", flattenObjectEmailfilterIptrustEntriesAddrType2edl(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "ObjectEmailfilterIptrustEntries-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectEmailfilterIptrustEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectEmailfilterIptrustEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip4_subnet", flattenObjectEmailfilterIptrustEntriesIp4Subnet2edl(o["ip4-subnet"], d, "ip4_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip4-subnet"], "ObjectEmailfilterIptrustEntries-Ip4Subnet"); ok {
			if err = d.Set("ip4_subnet", vv); err != nil {
				return fmt.Errorf("Error reading ip4_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip4_subnet: %v", err)
		}
	}

	if err = d.Set("ip6_subnet", flattenObjectEmailfilterIptrustEntriesIp6Subnet2edl(o["ip6-subnet"], d, "ip6_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-subnet"], "ObjectEmailfilterIptrustEntries-Ip6Subnet"); ok {
			if err = d.Set("ip6_subnet", vv); err != nil {
				return fmt.Errorf("Error reading ip6_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_subnet: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectEmailfilterIptrustEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectEmailfilterIptrustEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectEmailfilterIptrustEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectEmailfilterIptrustEntriesAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterIptrustEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterIptrustEntriesIp4Subnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectEmailfilterIptrustEntriesIp6Subnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterIptrustEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectEmailfilterIptrustEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandObjectEmailfilterIptrustEntriesAddrType2edl(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectEmailfilterIptrustEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip4_subnet"); ok || d.HasChange("ip4_subnet") {
		t, err := expandObjectEmailfilterIptrustEntriesIp4Subnet2edl(d, v, "ip4_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip4-subnet"] = t
		}
	}

	if v, ok := d.GetOk("ip6_subnet"); ok || d.HasChange("ip6_subnet") {
		t, err := expandObjectEmailfilterIptrustEntriesIp6Subnet2edl(d, v, "ip6_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-subnet"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectEmailfilterIptrustEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
