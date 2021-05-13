// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CA certificate.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemCertificateCa() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCertificateCaCreate,
		Read:   resourceSystemCertificateCaRead,
		Update: resourceSystemCertificateCaUpdate,
		Delete: resourceSystemCertificateCaDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ca": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemCertificateCaCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemCertificateCa(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCa resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateCa(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCa resource: %v", err)
	}

	d.SetId(getStringKey(d, ""))

	return resourceSystemCertificateCaRead(d, m)
}

func resourceSystemCertificateCaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemCertificateCa(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCa resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateCa(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCa resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, ""))

	return resourceSystemCertificateCaRead(d, m)
}

func resourceSystemCertificateCaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemCertificateCa(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateCa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateCaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemCertificateCa(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateCa(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCa resource from API: %v", err)
	}
	return nil
}

func flattenSystemCertificateCaCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCertificateCaComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCertificateCaName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemCertificateCa(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ca", flattenSystemCertificateCaCa(o["ca"], d, "ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca"], "SystemCertificateCa-Ca"); ok {
			if err = d.Set("ca", vv); err != nil {
				return fmt.Errorf("Error reading ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}

	if err = d.Set("comment", flattenSystemCertificateCaComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "SystemCertificateCa-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemCertificateCaName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemCertificateCa-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemCertificateCaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCertificateCaCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCertificateCaComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCaName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateCa(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ca"); ok {
		t, err := expandSystemCertificateCaCa(d, v, "ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandSystemCertificateCaComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemCertificateCaName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
