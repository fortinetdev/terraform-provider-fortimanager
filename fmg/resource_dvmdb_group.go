// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device group table.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDvmdbGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmdbGroupCreate,
		Read:   resourceDvmdbGroupRead,
		Update: resourceDvmdbGroupUpdate,
		Delete: resourceDvmdbGroupDelete,

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
			"desc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metafields": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"os_type": &schema.Schema{
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

func resourceDvmdbGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectDvmdbGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating DvmdbGroup resource while getting object: %v", err)
	}

	_, err = c.CreateDvmdbGroup(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating DvmdbGroup resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceDvmdbGroupRead(d, m)
}

func resourceDvmdbGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectDvmdbGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateDvmdbGroup(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceDvmdbGroupRead(d, m)
}

func resourceDvmdbGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteDvmdbGroup(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting DvmdbGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDvmdbGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadDvmdbGroup(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading DvmdbGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDvmdbGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DvmdbGroup resource from API: %v", err)
	}
	return nil
}

func flattenDvmdbGroupDesc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbGroupMetaFields(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbGroupOsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			-1: "unknown",
			0:  "fos",
			1:  "fsw",
			2:  "foc",
			3:  "fml",
			4:  "faz",
			5:  "fwb",
			6:  "fch",
			7:  "fct",
			8:  "log",
			9:  "fmg",
			10: "fsa",
			11: "fdd",
			12: "fac",
			13: "fpx",
			14: "fna",
			15: "ffw",
			16: "fsr",
			17: "fad",
			18: "fdc",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenDvmdbGroupType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "normal",
			1: "default",
			2: "auto",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectDvmdbGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("desc", flattenDvmdbGroupDesc(o["desc"], d, "desc")); err != nil {
		if vv, ok := fortiAPIPatch(o["desc"], "DvmdbGroup-Desc"); ok {
			if err = d.Set("desc", vv); err != nil {
				return fmt.Errorf("Error reading desc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading desc: %v", err)
		}
	}

	if err = d.Set("metafields", flattenDvmdbGroupMetaFields(o["meta fields"], d, "metafields")); err != nil {
		if vv, ok := fortiAPIPatch(o["meta fields"], "DvmdbGroup-MetaFields"); ok {
			if err = d.Set("metafields", vv); err != nil {
				return fmt.Errorf("Error reading metafields: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading metafields: %v", err)
		}
	}

	if err = d.Set("name", flattenDvmdbGroupName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DvmdbGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("os_type", flattenDvmdbGroupOsType(o["os_type"], d, "os_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["os_type"], "DvmdbGroup-OsType"); ok {
			if err = d.Set("os_type", vv); err != nil {
				return fmt.Errorf("Error reading os_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os_type: %v", err)
		}
	}

	if err = d.Set("type", flattenDvmdbGroupType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "DvmdbGroup-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenDvmdbGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmdbGroupDesc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbGroupMetaFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbGroupOsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbGroupType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDvmdbGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("desc"); ok {
		t, err := expandDvmdbGroupDesc(d, v, "desc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["desc"] = t
		}
	}

	if v, ok := d.GetOk("metafields"); ok {
		t, err := expandDvmdbGroupMetaFields(d, v, "metafields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["meta fields"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandDvmdbGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("os_type"); ok {
		t, err := expandDvmdbGroupOsType(d, v, "os_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os_type"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandDvmdbGroupType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
