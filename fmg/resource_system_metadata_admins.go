// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure admins.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemMetadataAdmins() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemMetadataAdminsCreate,
		Read:   resourceSystemMetadataAdminsRead,
		Update: resourceSystemMetadataAdminsUpdate,
		Delete: resourceSystemMetadataAdminsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fieldlength": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fieldname": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"importance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemMetadataAdminsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemMetadataAdmins(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemMetadataAdmins resource while getting object: %v", err)
	}

	_, err = c.CreateSystemMetadataAdmins(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemMetadataAdmins resource: %v", err)
	}

	d.SetId(getStringKey(d, "fieldname"))

	return resourceSystemMetadataAdminsRead(d, m)
}

func resourceSystemMetadataAdminsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemMetadataAdmins(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemMetadataAdmins resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemMetadataAdmins(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemMetadataAdmins resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "fieldname"))

	return resourceSystemMetadataAdminsRead(d, m)
}

func resourceSystemMetadataAdminsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemMetadataAdmins(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemMetadataAdmins resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemMetadataAdminsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemMetadataAdmins(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemMetadataAdmins resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemMetadataAdmins(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemMetadataAdmins resource from API: %v", err)
	}
	return nil
}

func flattenSystemMetadataAdminsFieldlength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMetadataAdminsFieldname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMetadataAdminsImportance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMetadataAdminsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemMetadataAdmins(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fieldlength", flattenSystemMetadataAdminsFieldlength(o["fieldlength"], d, "fieldlength")); err != nil {
		if vv, ok := fortiAPIPatch(o["fieldlength"], "SystemMetadataAdmins-Fieldlength"); ok {
			if err = d.Set("fieldlength", vv); err != nil {
				return fmt.Errorf("Error reading fieldlength: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fieldlength: %v", err)
		}
	}

	if err = d.Set("fieldname", flattenSystemMetadataAdminsFieldname(o["fieldname"], d, "fieldname")); err != nil {
		if vv, ok := fortiAPIPatch(o["fieldname"], "SystemMetadataAdmins-Fieldname"); ok {
			if err = d.Set("fieldname", vv); err != nil {
				return fmt.Errorf("Error reading fieldname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fieldname: %v", err)
		}
	}

	if err = d.Set("importance", flattenSystemMetadataAdminsImportance(o["importance"], d, "importance")); err != nil {
		if vv, ok := fortiAPIPatch(o["importance"], "SystemMetadataAdmins-Importance"); ok {
			if err = d.Set("importance", vv); err != nil {
				return fmt.Errorf("Error reading importance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading importance: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemMetadataAdminsStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemMetadataAdmins-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemMetadataAdminsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemMetadataAdminsFieldlength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMetadataAdminsFieldname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMetadataAdminsImportance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMetadataAdminsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemMetadataAdmins(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fieldlength"); ok {
		t, err := expandSystemMetadataAdminsFieldlength(d, v, "fieldlength")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fieldlength"] = t
		}
	}

	if v, ok := d.GetOk("fieldname"); ok {
		t, err := expandSystemMetadataAdminsFieldname(d, v, "fieldname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fieldname"] = t
		}
	}

	if v, ok := d.GetOk("importance"); ok {
		t, err := expandSystemMetadataAdminsImportance(d, v, "importance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["importance"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemMetadataAdminsStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
