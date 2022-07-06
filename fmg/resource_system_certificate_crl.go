// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Certificate Revocation List.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemCertificateCrl() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCertificateCrlCreate,
		Read:   resourceSystemCertificateCrlRead,
		Update: resourceSystemCertificateCrlUpdate,
		Delete: resourceSystemCertificateCrlDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"crl": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"http_url": &schema.Schema{
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
			"update_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemCertificateCrlCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemCertificateCrl(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCrl resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateCrl(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCrl resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemCertificateCrlRead(d, m)
}

func resourceSystemCertificateCrlUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemCertificateCrl(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCrl resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateCrl(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCrl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemCertificateCrlRead(d, m)
}

func resourceSystemCertificateCrlDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemCertificateCrl(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateCrl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateCrlRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemCertificateCrl(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCrl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateCrl(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCrl resource from API: %v", err)
	}
	return nil
}

func flattenSystemCertificateCrlComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCertificateCrlCrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCertificateCrlHttpUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCertificateCrlName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCertificateCrlUpdateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemCertificateCrl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenSystemCertificateCrlComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "SystemCertificateCrl-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("crl", flattenSystemCertificateCrlCrl(o["crl"], d, "crl")); err != nil {
		if vv, ok := fortiAPIPatch(o["crl"], "SystemCertificateCrl-Crl"); ok {
			if err = d.Set("crl", vv); err != nil {
				return fmt.Errorf("Error reading crl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading crl: %v", err)
		}
	}

	if err = d.Set("http_url", flattenSystemCertificateCrlHttpUrl(o["http-url"], d, "http_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-url"], "SystemCertificateCrl-HttpUrl"); ok {
			if err = d.Set("http_url", vv); err != nil {
				return fmt.Errorf("Error reading http_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_url: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemCertificateCrlName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemCertificateCrl-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenSystemCertificateCrlUpdateInterval(o["update-interval"], d, "update_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-interval"], "SystemCertificateCrl-UpdateInterval"); ok {
			if err = d.Set("update_interval", vv); err != nil {
				return fmt.Errorf("Error reading update_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	return nil
}

func flattenSystemCertificateCrlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCertificateCrlComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlCrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCertificateCrlHttpUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlUpdateInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateCrl(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandSystemCertificateCrlComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("crl"); ok || d.HasChange("crl") {
		t, err := expandSystemCertificateCrlCrl(d, v, "crl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crl"] = t
		}
	}

	if v, ok := d.GetOk("http_url"); ok || d.HasChange("http_url") {
		t, err := expandSystemCertificateCrlHttpUrl(d, v, "http_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-url"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemCertificateCrlName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("update_interval"); ok || d.HasChange("update_interval") {
		t, err := expandSystemCertificateCrlUpdateInterval(d, v, "update_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-interval"] = t
		}
	}

	return &obj, nil
}
