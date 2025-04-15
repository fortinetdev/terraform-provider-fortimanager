// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 SNMP managers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystempSystemSnmpCommunityHosts6() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystempSystemSnmpCommunityHosts6Create,
		Read:   resourceSystempSystemSnmpCommunityHosts6Read,
		Update: resourceSystempSystemSnmpCommunityHosts6Update,
		Delete: resourceSystempSystemSnmpCommunityHosts6Delete,

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
			"devprof": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"community": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ha_direct": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystempSystemSnmpCommunityHosts6Create(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	community := d.Get("community").(string)
	paradict["devprof"] = devprof
	paradict["community"] = community

	obj, err := getObjectSystempSystemSnmpCommunityHosts6(d)
	if err != nil {
		return fmt.Errorf("Error creating SystempSystemSnmpCommunityHosts6 resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystempSystemSnmpCommunityHosts6(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystempSystemSnmpCommunityHosts6 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystempSystemSnmpCommunityHosts6Read(d, m)
}

func resourceSystempSystemSnmpCommunityHosts6Update(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	community := d.Get("community").(string)
	paradict["devprof"] = devprof
	paradict["community"] = community

	obj, err := getObjectSystempSystemSnmpCommunityHosts6(d)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemSnmpCommunityHosts6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystempSystemSnmpCommunityHosts6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemSnmpCommunityHosts6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystempSystemSnmpCommunityHosts6Read(d, m)
}

func resourceSystempSystemSnmpCommunityHosts6Delete(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	community := d.Get("community").(string)
	paradict["devprof"] = devprof
	paradict["community"] = community

	wsParams["adom"] = adomv

	err = c.DeleteSystempSystemSnmpCommunityHosts6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystempSystemSnmpCommunityHosts6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystempSystemSnmpCommunityHosts6Read(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	community := d.Get("community").(string)
	if devprof == "" {
		devprof = importOptionChecking(m.(*FortiClient).Cfg, "devprof")
		if devprof == "" {
			return fmt.Errorf("Parameter devprof is missing")
		}
		if err = d.Set("devprof", devprof); err != nil {
			return fmt.Errorf("Error set params devprof: %v", err)
		}
	}
	if community == "" {
		community = importOptionChecking(m.(*FortiClient).Cfg, "community")
		if community == "" {
			return fmt.Errorf("Parameter community is missing")
		}
		if err = d.Set("community", community); err != nil {
			return fmt.Errorf("Error set params community: %v", err)
		}
	}
	paradict["devprof"] = devprof
	paradict["community"] = community

	o, err := c.ReadSystempSystemSnmpCommunityHosts6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemSnmpCommunityHosts6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystempSystemSnmpCommunityHosts6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemSnmpCommunityHosts6 resource from API: %v", err)
	}
	return nil
}

func flattenSystempSystemSnmpCommunityHosts6HaDirect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpCommunityHosts6HostType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpCommunityHosts6Id2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpCommunityHosts6Interface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempSystemSnmpCommunityHosts6InterfaceSelectMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpCommunityHosts6Ipv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpCommunityHosts6SourceIpv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpCommunityHosts6VrfSelect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystempSystemSnmpCommunityHosts6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("ha_direct", flattenSystempSystemSnmpCommunityHosts6HaDirect2edl(o["ha-direct"], d, "ha_direct")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-direct"], "SystempSystemSnmpCommunityHosts6-HaDirect"); ok {
			if err = d.Set("ha_direct", vv); err != nil {
				return fmt.Errorf("Error reading ha_direct: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_direct: %v", err)
		}
	}

	if err = d.Set("host_type", flattenSystempSystemSnmpCommunityHosts6HostType2edl(o["host-type"], d, "host_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-type"], "SystempSystemSnmpCommunityHosts6-HostType"); ok {
			if err = d.Set("host_type", vv); err != nil {
				return fmt.Errorf("Error reading host_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_type: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystempSystemSnmpCommunityHosts6Id2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystempSystemSnmpCommunityHosts6-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystempSystemSnmpCommunityHosts6Interface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystempSystemSnmpCommunityHosts6-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystempSystemSnmpCommunityHosts6InterfaceSelectMethod2edl(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "SystempSystemSnmpCommunityHosts6-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ipv6", flattenSystempSystemSnmpCommunityHosts6Ipv62edl(o["ipv6"], d, "ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6"], "SystempSystemSnmpCommunityHosts6-Ipv6"); ok {
			if err = d.Set("ipv6", vv); err != nil {
				return fmt.Errorf("Error reading ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	if err = d.Set("source_ipv6", flattenSystempSystemSnmpCommunityHosts6SourceIpv62edl(o["source-ipv6"], d, "source_ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ipv6"], "SystempSystemSnmpCommunityHosts6-SourceIpv6"); ok {
			if err = d.Set("source_ipv6", vv); err != nil {
				return fmt.Errorf("Error reading source_ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ipv6: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenSystempSystemSnmpCommunityHosts6VrfSelect2edl(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "SystempSystemSnmpCommunityHosts6-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenSystempSystemSnmpCommunityHosts6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystempSystemSnmpCommunityHosts6HaDirect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpCommunityHosts6HostType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpCommunityHosts6Id2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpCommunityHosts6Interface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemSnmpCommunityHosts6InterfaceSelectMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpCommunityHosts6Ipv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpCommunityHosts6SourceIpv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpCommunityHosts6VrfSelect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystempSystemSnmpCommunityHosts6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ha_direct"); ok || d.HasChange("ha_direct") {
		t, err := expandSystempSystemSnmpCommunityHosts6HaDirect2edl(d, v, "ha_direct")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-direct"] = t
		}
	}

	if v, ok := d.GetOk("host_type"); ok || d.HasChange("host_type") {
		t, err := expandSystempSystemSnmpCommunityHosts6HostType2edl(d, v, "host_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-type"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystempSystemSnmpCommunityHosts6Id2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystempSystemSnmpCommunityHosts6Interface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandSystempSystemSnmpCommunityHosts6InterfaceSelectMethod2edl(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok || d.HasChange("ipv6") {
		t, err := expandSystempSystemSnmpCommunityHosts6Ipv62edl(d, v, "ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	if v, ok := d.GetOk("source_ipv6"); ok || d.HasChange("source_ipv6") {
		t, err := expandSystempSystemSnmpCommunityHosts6SourceIpv62edl(d, v, "source_ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ipv6"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandSystempSystemSnmpCommunityHosts6VrfSelect2edl(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
