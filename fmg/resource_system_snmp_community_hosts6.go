// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Allow hosts configuration for IPv6.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSnmpCommunityHosts6() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnmpCommunityHosts6Create,
		Read:   resourceSystemSnmpCommunityHosts6Read,
		Update: resourceSystemSnmpCommunityHosts6Update,
		Delete: resourceSystemSnmpCommunityHosts6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"community": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSnmpCommunityHosts6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	community := d.Get("community").(string)
	paradict["community"] = community

	obj, err := getObjectSystemSnmpCommunityHosts6(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpCommunityHosts6 resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSnmpCommunityHosts6(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpCommunityHosts6 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSnmpCommunityHosts6Read(d, m)
}

func resourceSystemSnmpCommunityHosts6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	community := d.Get("community").(string)
	paradict["community"] = community

	obj, err := getObjectSystemSnmpCommunityHosts6(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpCommunityHosts6 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSnmpCommunityHosts6(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpCommunityHosts6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSnmpCommunityHosts6Read(d, m)
}

func resourceSystemSnmpCommunityHosts6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	community := d.Get("community").(string)
	paradict["community"] = community

	err = c.DeleteSystemSnmpCommunityHosts6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpCommunityHosts6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpCommunityHosts6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	community := d.Get("community").(string)
	if community == "" {
		community = importOptionChecking(m.(*FortiClient).Cfg, "community")
		if community == "" {
			return fmt.Errorf("Parameter community is missing")
		}
		if err = d.Set("community", community); err != nil {
			return fmt.Errorf("Error set params community: %v", err)
		}
	}
	paradict["community"] = community

	o, err := c.ReadSystemSnmpCommunityHosts6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpCommunityHosts6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpCommunityHosts6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpCommunityHosts6 resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnmpCommunityHosts6Id2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6Interface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6Ip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSnmpCommunityHosts6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemSnmpCommunityHosts6Id2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemSnmpCommunityHosts6-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemSnmpCommunityHosts6Interface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemSnmpCommunityHosts6-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemSnmpCommunityHosts6Ip2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SystemSnmpCommunityHosts6-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	return nil
}

func flattenSystemSnmpCommunityHosts6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSnmpCommunityHosts6Id2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6Interface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6Ip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpCommunityHosts6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemSnmpCommunityHosts6Id2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemSnmpCommunityHosts6Interface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSystemSnmpCommunityHosts6Ip2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	return &obj, nil
}
