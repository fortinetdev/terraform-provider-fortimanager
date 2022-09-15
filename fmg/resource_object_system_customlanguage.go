// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure custom languages.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemCustomLanguage() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemCustomLanguageCreate,
		Read:   resourceObjectSystemCustomLanguageRead,
		Update: resourceObjectSystemCustomLanguageUpdate,
		Delete: resourceObjectSystemCustomLanguageDelete,

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
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemCustomLanguageCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemCustomLanguage(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemCustomLanguage resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemCustomLanguage(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemCustomLanguage resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemCustomLanguageRead(d, m)
}

func resourceObjectSystemCustomLanguageUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemCustomLanguage(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemCustomLanguage resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemCustomLanguage(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemCustomLanguage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemCustomLanguageRead(d, m)
}

func resourceObjectSystemCustomLanguageDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemCustomLanguage(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemCustomLanguage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemCustomLanguageRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemCustomLanguage(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemCustomLanguage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemCustomLanguage(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemCustomLanguage resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemCustomLanguageComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemCustomLanguageFilename(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemCustomLanguageName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemCustomLanguage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comments", flattenObjectSystemCustomLanguageComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "ObjectSystemCustomLanguage-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("filename", flattenObjectSystemCustomLanguageFilename(o["filename"], d, "filename")); err != nil {
		if vv, ok := fortiAPIPatch(o["filename"], "ObjectSystemCustomLanguage-Filename"); ok {
			if err = d.Set("filename", vv); err != nil {
				return fmt.Errorf("Error reading filename: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filename: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSystemCustomLanguageName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemCustomLanguage-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemCustomLanguageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemCustomLanguageComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemCustomLanguageFilename(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemCustomLanguageName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemCustomLanguage(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandObjectSystemCustomLanguageComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("filename"); ok || d.HasChange("filename") {
		t, err := expandObjectSystemCustomLanguageFilename(d, v, "filename")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filename"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSystemCustomLanguageName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
