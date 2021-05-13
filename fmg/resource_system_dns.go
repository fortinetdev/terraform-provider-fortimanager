// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DNS configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDnsUpdate,
		Read:   resourceSystemDnsRead,
		Update: resourceSystemDnsUpdate,
		Delete: resourceSystemDnsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ip6_primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDnsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemDns(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDns(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemDns")

	return resourceSystemDnsRead(d, m)
}

func resourceSystemDnsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemDns(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDnsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemDns(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDns(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource from API: %v", err)
	}
	return nil
}

func flattenSystemDnsIp6Primary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsIp6Secondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsSecondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDns(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip6_primary", flattenSystemDnsIp6Primary(o["ip6-primary"], d, "ip6_primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-primary"], "SystemDns-Ip6Primary"); ok {
			if err = d.Set("ip6_primary", vv); err != nil {
				return fmt.Errorf("Error reading ip6_primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_primary: %v", err)
		}
	}

	if err = d.Set("ip6_secondary", flattenSystemDnsIp6Secondary(o["ip6-secondary"], d, "ip6_secondary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-secondary"], "SystemDns-Ip6Secondary"); ok {
			if err = d.Set("ip6_secondary", vv); err != nil {
				return fmt.Errorf("Error reading ip6_secondary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_secondary: %v", err)
		}
	}

	if err = d.Set("primary", flattenSystemDnsPrimary(o["primary"], d, "primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary"], "SystemDns-Primary"); ok {
			if err = d.Set("primary", vv); err != nil {
				return fmt.Errorf("Error reading primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary: %v", err)
		}
	}

	if err = d.Set("secondary", flattenSystemDnsSecondary(o["secondary"], d, "secondary")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary"], "SystemDns-Secondary"); ok {
			if err = d.Set("secondary", vv); err != nil {
				return fmt.Errorf("Error reading secondary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary: %v", err)
		}
	}

	return nil
}

func flattenSystemDnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDnsIp6Primary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsIp6Secondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsSecondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDns(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip6_primary"); ok {
		t, err := expandSystemDnsIp6Primary(d, v, "ip6_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-primary"] = t
		}
	}

	if v, ok := d.GetOk("ip6_secondary"); ok {
		t, err := expandSystemDnsIp6Secondary(d, v, "ip6_secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-secondary"] = t
		}
	}

	if v, ok := d.GetOk("primary"); ok {
		t, err := expandSystemDnsPrimary(d, v, "primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary"] = t
		}
	}

	if v, ok := d.GetOk("secondary"); ok {
		t, err := expandSystemDnsSecondary(d, v, "secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary"] = t
		}
	}

	return &obj, nil
}
