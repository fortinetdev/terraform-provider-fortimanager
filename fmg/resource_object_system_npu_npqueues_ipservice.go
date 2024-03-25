// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure a NP7 QoS IP Service.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuNpQueuesIpService() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuNpQueuesIpServiceCreate,
		Read:   resourceObjectSystemNpuNpQueuesIpServiceRead,
		Update: resourceObjectSystemNpuNpQueuesIpServiceUpdate,
		Delete: resourceObjectSystemNpuNpQueuesIpServiceDelete,

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
			"dport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"queue": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemNpuNpQueuesIpServiceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSystemNpuNpQueuesIpService(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuNpQueuesIpService resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemNpuNpQueuesIpService(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuNpQueuesIpService resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemNpuNpQueuesIpServiceRead(d, m)
}

func resourceObjectSystemNpuNpQueuesIpServiceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuNpQueuesIpService(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpQueuesIpService resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuNpQueuesIpService(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpQueuesIpService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemNpuNpQueuesIpServiceRead(d, m)
}

func resourceObjectSystemNpuNpQueuesIpServiceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemNpuNpQueuesIpService(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuNpQueuesIpService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuNpQueuesIpServiceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuNpQueuesIpService(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpQueuesIpService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuNpQueuesIpService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpQueuesIpService resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuNpQueuesIpServiceDport3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceProtocol3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceQueue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceSport3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceWeight3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuNpQueuesIpService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("dport", flattenObjectSystemNpuNpQueuesIpServiceDport3rdl(o["dport"], d, "dport")); err != nil {
		if vv, ok := fortiAPIPatch(o["dport"], "ObjectSystemNpuNpQueuesIpService-Dport"); ok {
			if err = d.Set("dport", vv); err != nil {
				return fmt.Errorf("Error reading dport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dport: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSystemNpuNpQueuesIpServiceName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemNpuNpQueuesIpService-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("protocol", flattenObjectSystemNpuNpQueuesIpServiceProtocol3rdl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ObjectSystemNpuNpQueuesIpService-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("queue", flattenObjectSystemNpuNpQueuesIpServiceQueue3rdl(o["queue"], d, "queue")); err != nil {
		if vv, ok := fortiAPIPatch(o["queue"], "ObjectSystemNpuNpQueuesIpService-Queue"); ok {
			if err = d.Set("queue", vv); err != nil {
				return fmt.Errorf("Error reading queue: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading queue: %v", err)
		}
	}

	if err = d.Set("sport", flattenObjectSystemNpuNpQueuesIpServiceSport3rdl(o["sport"], d, "sport")); err != nil {
		if vv, ok := fortiAPIPatch(o["sport"], "ObjectSystemNpuNpQueuesIpService-Sport"); ok {
			if err = d.Set("sport", vv); err != nil {
				return fmt.Errorf("Error reading sport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sport: %v", err)
		}
	}

	if err = d.Set("weight", flattenObjectSystemNpuNpQueuesIpServiceWeight3rdl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ObjectSystemNpuNpQueuesIpService-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuNpQueuesIpServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuNpQueuesIpServiceDport3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceProtocol3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceQueue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceSport3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceWeight3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuNpQueuesIpService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dport"); ok || d.HasChange("dport") {
		t, err := expandObjectSystemNpuNpQueuesIpServiceDport3rdl(d, v, "dport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dport"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSystemNpuNpQueuesIpServiceName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandObjectSystemNpuNpQueuesIpServiceProtocol3rdl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("queue"); ok || d.HasChange("queue") {
		t, err := expandObjectSystemNpuNpQueuesIpServiceQueue3rdl(d, v, "queue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["queue"] = t
		}
	}

	if v, ok := d.GetOk("sport"); ok || d.HasChange("sport") {
		t, err := expandObjectSystemNpuNpQueuesIpServiceSport3rdl(d, v, "sport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sport"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandObjectSystemNpuNpQueuesIpServiceWeight3rdl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
