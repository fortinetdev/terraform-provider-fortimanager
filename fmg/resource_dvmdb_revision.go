// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ADOM revision table.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDvmdbRevision() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmdbRevisionCreate,
		Read:   resourceDvmdbRevisionRead,
		Update: resourceDvmdbRevisionUpdate,
		Delete: resourceDvmdbRevisionDelete,

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
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"created_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"desc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"locked": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceDvmdbRevisionCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectDvmdbRevision(d)
	if err != nil {
		return fmt.Errorf("Error creating DvmdbRevision resource while getting object: %v", err)
	}

	_, err = c.CreateDvmdbRevision(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating DvmdbRevision resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "version")))

	return resourceDvmdbRevisionRead(d, m)
}

func resourceDvmdbRevisionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectDvmdbRevision(d)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbRevision resource while getting object: %v", err)
	}

	_, err = c.UpdateDvmdbRevision(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbRevision resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "version")))

	return resourceDvmdbRevisionRead(d, m)
}

func resourceDvmdbRevisionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteDvmdbRevision(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting DvmdbRevision resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDvmdbRevisionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadDvmdbRevision(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading DvmdbRevision resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDvmdbRevision(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DvmdbRevision resource from API: %v", err)
	}
	return nil
}

func flattenDvmdbRevisionCreatedBy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbRevisionCreatedTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbRevisionDesc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbRevisionLocked(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbRevisionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbRevisionVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDvmdbRevision(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("created_by", flattenDvmdbRevisionCreatedBy(o["created_by"], d, "created_by")); err != nil {
		if vv, ok := fortiAPIPatch(o["created_by"], "DvmdbRevision-CreatedBy"); ok {
			if err = d.Set("created_by", vv); err != nil {
				return fmt.Errorf("Error reading created_by: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading created_by: %v", err)
		}
	}

	if err = d.Set("created_time", flattenDvmdbRevisionCreatedTime(o["created_time"], d, "created_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["created_time"], "DvmdbRevision-CreatedTime"); ok {
			if err = d.Set("created_time", vv); err != nil {
				return fmt.Errorf("Error reading created_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading created_time: %v", err)
		}
	}

	if err = d.Set("desc", flattenDvmdbRevisionDesc(o["desc"], d, "desc")); err != nil {
		if vv, ok := fortiAPIPatch(o["desc"], "DvmdbRevision-Desc"); ok {
			if err = d.Set("desc", vv); err != nil {
				return fmt.Errorf("Error reading desc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading desc: %v", err)
		}
	}

	if err = d.Set("locked", flattenDvmdbRevisionLocked(o["locked"], d, "locked")); err != nil {
		if vv, ok := fortiAPIPatch(o["locked"], "DvmdbRevision-Locked"); ok {
			if err = d.Set("locked", vv); err != nil {
				return fmt.Errorf("Error reading locked: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading locked: %v", err)
		}
	}

	if err = d.Set("name", flattenDvmdbRevisionName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DvmdbRevision-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("version", flattenDvmdbRevisionVersion(o["version"], d, "version")); err != nil {
		if vv, ok := fortiAPIPatch(o["version"], "DvmdbRevision-Version"); ok {
			if err = d.Set("version", vv); err != nil {
				return fmt.Errorf("Error reading version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	return nil
}

func flattenDvmdbRevisionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmdbRevisionCreatedBy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbRevisionCreatedTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbRevisionDesc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbRevisionLocked(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbRevisionName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbRevisionVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDvmdbRevision(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("created_by"); ok || d.HasChange("created_by") {
		t, err := expandDvmdbRevisionCreatedBy(d, v, "created_by")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["created_by"] = t
		}
	}

	if v, ok := d.GetOk("created_time"); ok || d.HasChange("created_time") {
		t, err := expandDvmdbRevisionCreatedTime(d, v, "created_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["created_time"] = t
		}
	}

	if v, ok := d.GetOk("desc"); ok || d.HasChange("desc") {
		t, err := expandDvmdbRevisionDesc(d, v, "desc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["desc"] = t
		}
	}

	if v, ok := d.GetOk("locked"); ok || d.HasChange("locked") {
		t, err := expandDvmdbRevisionLocked(d, v, "locked")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["locked"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDvmdbRevisionName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandDvmdbRevisionVersion(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	return &obj, nil
}
