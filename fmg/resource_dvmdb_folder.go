// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Dvmdb Folder

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDvmdbFolder() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmdbFolderCreate,
		Read:   resourceDvmdbFolderRead,
		Update: resourceDvmdbFolderUpdate,
		Delete: resourceDvmdbFolderDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"parent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDvmdbFolderCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectDvmdbFolder(d)
	if err != nil {
		return fmt.Errorf("Error creating DvmdbFolder resource while getting object: %v", err)
	}

	_, err = c.CreateDvmdbFolder(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating DvmdbFolder resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceDvmdbFolderRead(d, m)
}

func resourceDvmdbFolderUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectDvmdbFolder(d)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbFolder resource while getting object: %v", err)
	}

	_, err = c.UpdateDvmdbFolder(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbFolder resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceDvmdbFolderRead(d, m)
}

func resourceDvmdbFolderDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteDvmdbFolder(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting DvmdbFolder resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDvmdbFolderRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadDvmdbFolder(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading DvmdbFolder resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDvmdbFolder(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DvmdbFolder resource from API: %v", err)
	}
	return nil
}

func flattenDvmdbFolderDesc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbFolderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbFolderParent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDvmdbFolder(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("desc", flattenDvmdbFolderDesc(o["desc"], d, "desc")); err != nil {
		if vv, ok := fortiAPIPatch(o["desc"], "DvmdbFolder-Desc"); ok {
			if err = d.Set("desc", vv); err != nil {
				return fmt.Errorf("Error reading desc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading desc: %v", err)
		}
	}

	if err = d.Set("name", flattenDvmdbFolderName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DvmdbFolder-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("parent", flattenDvmdbFolderParent(o["parent"], d, "parent")); err != nil {
		if vv, ok := fortiAPIPatch(o["parent"], "DvmdbFolder-Parent"); ok {
			if err = d.Set("parent", vv); err != nil {
				return fmt.Errorf("Error reading parent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading parent: %v", err)
		}
	}

	return nil
}

func flattenDvmdbFolderFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmdbFolderDesc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbFolderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbFolderParent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDvmdbFolder(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("desc"); ok {
		t, err := expandDvmdbFolderDesc(d, v, "desc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["desc"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandDvmdbFolderName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("parent"); ok {
		t, err := expandDvmdbFolderParent(d, v, "parent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parent"] = t
		}
	}

	return &obj, nil
}
