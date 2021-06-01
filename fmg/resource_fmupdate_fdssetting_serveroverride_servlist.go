// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Override server.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFmupdateFdsSettingServerOverrideServlist() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateFdsSettingServerOverrideServlistCreate,
		Read:   resourceFmupdateFdsSettingServerOverrideServlistRead,
		Update: resourceFmupdateFdsSettingServerOverrideServlistUpdate,
		Delete: resourceFmupdateFdsSettingServerOverrideServlistDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"service_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateFdsSettingServerOverrideServlistCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateFdsSettingServerOverrideServlist(d)
	if err != nil {
		return fmt.Errorf("Error creating FmupdateFdsSettingServerOverrideServlist resource while getting object: %v", err)
	}

	_, err = c.CreateFmupdateFdsSettingServerOverrideServlist(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating FmupdateFdsSettingServerOverrideServlist resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFmupdateFdsSettingServerOverrideServlistRead(d, m)
}

func resourceFmupdateFdsSettingServerOverrideServlistUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateFdsSettingServerOverrideServlist(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFdsSettingServerOverrideServlist resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateFdsSettingServerOverrideServlist(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFdsSettingServerOverrideServlist resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFmupdateFdsSettingServerOverrideServlistRead(d, m)
}

func resourceFmupdateFdsSettingServerOverrideServlistDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateFdsSettingServerOverrideServlist(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateFdsSettingServerOverrideServlist resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateFdsSettingServerOverrideServlistRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateFdsSettingServerOverrideServlist(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFdsSettingServerOverrideServlist resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateFdsSettingServerOverrideServlist(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFdsSettingServerOverrideServlist resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateFdsSettingServerOverrideServlistIdFfssb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistIpFfssb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistIp6Ffssb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistPortFfssb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingServerOverrideServlistServiceTypeFfssb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "fct",
			2: "fds",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectFmupdateFdsSettingServerOverrideServlist(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenFmupdateFdsSettingServerOverrideServlistIdFfssb(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FmupdateFdsSettingServerOverrideServlist-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenFmupdateFdsSettingServerOverrideServlistIpFfssb(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "FmupdateFdsSettingServerOverrideServlist-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ip6", flattenFmupdateFdsSettingServerOverrideServlistIp6Ffssb(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "FmupdateFdsSettingServerOverrideServlist-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("port", flattenFmupdateFdsSettingServerOverrideServlistPortFfssb(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "FmupdateFdsSettingServerOverrideServlist-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("service_type", flattenFmupdateFdsSettingServerOverrideServlistServiceTypeFfssb(o["service-type"], d, "service_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-type"], "FmupdateFdsSettingServerOverrideServlist-ServiceType"); ok {
			if err = d.Set("service_type", vv); err != nil {
				return fmt.Errorf("Error reading service_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_type: %v", err)
		}
	}

	return nil
}

func flattenFmupdateFdsSettingServerOverrideServlistFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateFdsSettingServerOverrideServlistIdFfssb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistIpFfssb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistIp6Ffssb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistPortFfssb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingServerOverrideServlistServiceTypeFfssb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateFdsSettingServerOverrideServlist(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandFmupdateFdsSettingServerOverrideServlistIdFfssb(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandFmupdateFdsSettingServerOverrideServlistIpFfssb(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {
		t, err := expandFmupdateFdsSettingServerOverrideServlistIp6Ffssb(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandFmupdateFdsSettingServerOverrideServlistPortFfssb(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("service_type"); ok {
		t, err := expandFmupdateFdsSettingServerOverrideServlistServiceTypeFfssb(d, v, "service_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-type"] = t
		}
	}

	return &obj, nil
}
