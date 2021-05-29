// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SSH certificates and keys.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemCertificateSsh() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCertificateSshCreate,
		Read:   resourceSystemCertificateSshRead,
		Update: resourceSystemCertificateSshUpdate,
		Delete: resourceSystemCertificateSshDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"certificate": &schema.Schema{
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
			"private_key": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemCertificateSshCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemCertificateSsh(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateSsh resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateSsh(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateSsh resource: %v", err)
	}

	d.SetId(getStringKey(d, ""))

	return resourceSystemCertificateSshRead(d, m)
}

func resourceSystemCertificateSshUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemCertificateSsh(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateSsh resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateSsh(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateSsh resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, ""))

	return resourceSystemCertificateSshRead(d, m)
}

func resourceSystemCertificateSshDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemCertificateSsh(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateSsh resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateSshRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemCertificateSsh(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateSsh resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateSsh(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateSsh resource from API: %v", err)
	}
	return nil
}

func flattenSystemCertificateSshCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCertificateSshComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCertificateSshName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCertificateSshPrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemCertificateSsh(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("certificate", flattenSystemCertificateSshCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "SystemCertificateSsh-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("comment", flattenSystemCertificateSshComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "SystemCertificateSsh-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemCertificateSshName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemCertificateSsh-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("private_key", flattenSystemCertificateSshPrivateKey(o["private-key"], d, "private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-key"], "SystemCertificateSsh-PrivateKey"); ok {
			if err = d.Set("private_key", vv); err != nil {
				return fmt.Errorf("Error reading private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_key: %v", err)
		}
	}

	return nil
}

func flattenSystemCertificateSshFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCertificateSshCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCertificateSshComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateSshName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateSshPrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemCertificateSsh(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("certificate"); ok {
		t, err := expandSystemCertificateSshCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandSystemCertificateSshComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemCertificateSshName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("private_key"); ok {
		t, err := expandSystemCertificateSshPrivateKey(d, v, "private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key"] = t
		}
	}

	return &obj, nil
}
