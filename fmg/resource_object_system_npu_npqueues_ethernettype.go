// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure a NP7 QoS Ethernet Type.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuNpQueuesEthernetType() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuNpQueuesEthernetTypeCreate,
		Read:   resourceObjectSystemNpuNpQueuesEthernetTypeRead,
		Update: resourceObjectSystemNpuNpQueuesEthernetTypeUpdate,
		Delete: resourceObjectSystemNpuNpQueuesEthernetTypeDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"queue": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemNpuNpQueuesEthernetTypeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuNpQueuesEthernetType(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuNpQueuesEthernetType resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSystemNpuNpQueuesEthernetType(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuNpQueuesEthernetType resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemNpuNpQueuesEthernetTypeRead(d, m)
}

func resourceObjectSystemNpuNpQueuesEthernetTypeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuNpQueuesEthernetType(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpQueuesEthernetType resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemNpuNpQueuesEthernetType(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpQueuesEthernetType resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemNpuNpQueuesEthernetTypeRead(d, m)
}

func resourceObjectSystemNpuNpQueuesEthernetTypeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemNpuNpQueuesEthernetType(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuNpQueuesEthernetType resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuNpQueuesEthernetTypeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuNpQueuesEthernetType(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpQueuesEthernetType resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuNpQueuesEthernetType(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpQueuesEthernetType resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuNpQueuesEthernetTypeName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesEthernetTypeQueue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesEthernetTypeType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesEthernetTypeWeight3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuNpQueuesEthernetType(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectSystemNpuNpQueuesEthernetTypeName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemNpuNpQueuesEthernetType-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("queue", flattenObjectSystemNpuNpQueuesEthernetTypeQueue3rdl(o["queue"], d, "queue")); err != nil {
		if vv, ok := fortiAPIPatch(o["queue"], "ObjectSystemNpuNpQueuesEthernetType-Queue"); ok {
			if err = d.Set("queue", vv); err != nil {
				return fmt.Errorf("Error reading queue: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading queue: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectSystemNpuNpQueuesEthernetTypeType3rdl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectSystemNpuNpQueuesEthernetType-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("weight", flattenObjectSystemNpuNpQueuesEthernetTypeWeight3rdl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ObjectSystemNpuNpQueuesEthernetType-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuNpQueuesEthernetTypeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuNpQueuesEthernetTypeName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeQueue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeWeight3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuNpQueuesEthernetType(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSystemNpuNpQueuesEthernetTypeName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("queue"); ok || d.HasChange("queue") {
		t, err := expandObjectSystemNpuNpQueuesEthernetTypeQueue3rdl(d, v, "queue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["queue"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectSystemNpuNpQueuesEthernetTypeType3rdl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandObjectSystemNpuNpQueuesEthernetTypeWeight3rdl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
