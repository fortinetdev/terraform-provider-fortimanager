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

func resourceFmupdateWebSpamFgdSettingServerOverrideServlist() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateWebSpamFgdSettingServerOverrideServlistCreate,
		Read:   resourceFmupdateWebSpamFgdSettingServerOverrideServlistRead,
		Update: resourceFmupdateWebSpamFgdSettingServerOverrideServlistUpdate,
		Delete: resourceFmupdateWebSpamFgdSettingServerOverrideServlistDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceFmupdateWebSpamFgdSettingServerOverrideServlistCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateWebSpamFgdSettingServerOverrideServlist(d)
	if err != nil {
		return fmt.Errorf("Error creating FmupdateWebSpamFgdSettingServerOverrideServlist resource while getting object: %v", err)
	}

	_, err = c.CreateFmupdateWebSpamFgdSettingServerOverrideServlist(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating FmupdateWebSpamFgdSettingServerOverrideServlist resource: %v", err)
	}

	d.SetId(getStringKey(d, ""))

	return resourceFmupdateWebSpamFgdSettingServerOverrideServlistRead(d, m)
}

func resourceFmupdateWebSpamFgdSettingServerOverrideServlistUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateWebSpamFgdSettingServerOverrideServlist(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateWebSpamFgdSettingServerOverrideServlist resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateWebSpamFgdSettingServerOverrideServlist(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateWebSpamFgdSettingServerOverrideServlist resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, ""))

	return resourceFmupdateWebSpamFgdSettingServerOverrideServlistRead(d, m)
}

func resourceFmupdateWebSpamFgdSettingServerOverrideServlistDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateWebSpamFgdSettingServerOverrideServlist(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateWebSpamFgdSettingServerOverrideServlist resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateWebSpamFgdSettingServerOverrideServlistRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateWebSpamFgdSettingServerOverrideServlist(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateWebSpamFgdSettingServerOverrideServlist resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateWebSpamFgdSettingServerOverrideServlist(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateWebSpamFgdSettingServerOverrideServlist resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistServiceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			256:  "fgd",
			512:  "fgc",
			1024: "fsa",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectFmupdateWebSpamFgdSettingServerOverrideServlist(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenFmupdateWebSpamFgdSettingServerOverrideServlistId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FmupdateWebSpamFgdSettingServerOverrideServlist-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenFmupdateWebSpamFgdSettingServerOverrideServlistIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "FmupdateWebSpamFgdSettingServerOverrideServlist-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ip6", flattenFmupdateWebSpamFgdSettingServerOverrideServlistIp6(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "FmupdateWebSpamFgdSettingServerOverrideServlist-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("port", flattenFmupdateWebSpamFgdSettingServerOverrideServlistPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "FmupdateWebSpamFgdSettingServerOverrideServlist-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("service_type", flattenFmupdateWebSpamFgdSettingServerOverrideServlistServiceType(o["service-type"], d, "service_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-type"], "FmupdateWebSpamFgdSettingServerOverrideServlist-ServiceType"); ok {
			if err = d.Set("service_type", vv); err != nil {
				return fmt.Errorf("Error reading service_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_type: %v", err)
		}
	}

	return nil
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistServiceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateWebSpamFgdSettingServerOverrideServlist(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandFmupdateWebSpamFgdSettingServerOverrideServlistId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandFmupdateWebSpamFgdSettingServerOverrideServlistIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {
		t, err := expandFmupdateWebSpamFgdSettingServerOverrideServlistIp6(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandFmupdateWebSpamFgdSettingServerOverrideServlistPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("service_type"); ok {
		t, err := expandFmupdateWebSpamFgdSettingServerOverrideServlistServiceType(d, v, "service_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-type"] = t
		}
	}

	return &obj, nil
}
