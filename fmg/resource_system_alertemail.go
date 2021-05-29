// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure alertemail.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAlertemail() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAlertemailUpdate,
		Read:   resourceSystemAlertemailRead,
		Update: resourceSystemAlertemailUpdate,
		Delete: resourceSystemAlertemailDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fromaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fromname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smtppassword": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"smtpport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"smtpserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smtpuser": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAlertemailUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemAlertemail(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAlertemail resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAlertemail(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemAlertemail resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemAlertemail")

	return resourceSystemAlertemailRead(d, m)
}

func resourceSystemAlertemailDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemAlertemail(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAlertemail resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAlertemailRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemAlertemail(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemAlertemail resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAlertemail(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAlertemail resource from API: %v", err)
	}
	return nil
}

func flattenSystemAlertemailAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemAlertemailFromaddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAlertemailFromname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAlertemailSmtppassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAlertemailSmtpport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAlertemailSmtpserver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAlertemailSmtpuser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAlertemail(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("authentication", flattenSystemAlertemailAuthentication(o["authentication"], d, "authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["authentication"], "SystemAlertemail-Authentication"); ok {
			if err = d.Set("authentication", vv); err != nil {
				return fmt.Errorf("Error reading authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("fromaddress", flattenSystemAlertemailFromaddress(o["fromaddress"], d, "fromaddress")); err != nil {
		if vv, ok := fortiAPIPatch(o["fromaddress"], "SystemAlertemail-Fromaddress"); ok {
			if err = d.Set("fromaddress", vv); err != nil {
				return fmt.Errorf("Error reading fromaddress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fromaddress: %v", err)
		}
	}

	if err = d.Set("fromname", flattenSystemAlertemailFromname(o["fromname"], d, "fromname")); err != nil {
		if vv, ok := fortiAPIPatch(o["fromname"], "SystemAlertemail-Fromname"); ok {
			if err = d.Set("fromname", vv); err != nil {
				return fmt.Errorf("Error reading fromname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fromname: %v", err)
		}
	}

	if err = d.Set("smtppassword", flattenSystemAlertemailSmtppassword(o["smtppassword"], d, "smtppassword")); err != nil {
		if vv, ok := fortiAPIPatch(o["smtppassword"], "SystemAlertemail-Smtppassword"); ok {
			if err = d.Set("smtppassword", vv); err != nil {
				return fmt.Errorf("Error reading smtppassword: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smtppassword: %v", err)
		}
	}

	if err = d.Set("smtpport", flattenSystemAlertemailSmtpport(o["smtpport"], d, "smtpport")); err != nil {
		if vv, ok := fortiAPIPatch(o["smtpport"], "SystemAlertemail-Smtpport"); ok {
			if err = d.Set("smtpport", vv); err != nil {
				return fmt.Errorf("Error reading smtpport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smtpport: %v", err)
		}
	}

	if err = d.Set("smtpserver", flattenSystemAlertemailSmtpserver(o["smtpserver"], d, "smtpserver")); err != nil {
		if vv, ok := fortiAPIPatch(o["smtpserver"], "SystemAlertemail-Smtpserver"); ok {
			if err = d.Set("smtpserver", vv); err != nil {
				return fmt.Errorf("Error reading smtpserver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smtpserver: %v", err)
		}
	}

	if err = d.Set("smtpuser", flattenSystemAlertemailSmtpuser(o["smtpuser"], d, "smtpuser")); err != nil {
		if vv, ok := fortiAPIPatch(o["smtpuser"], "SystemAlertemail-Smtpuser"); ok {
			if err = d.Set("smtpuser", vv); err != nil {
				return fmt.Errorf("Error reading smtpuser: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smtpuser: %v", err)
		}
	}

	return nil
}

func flattenSystemAlertemailFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAlertemailAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertemailFromaddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertemailFromname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertemailSmtppassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAlertemailSmtpport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertemailSmtpserver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertemailSmtpuser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAlertemail(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authentication"); ok {
		t, err := expandSystemAlertemailAuthentication(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("fromaddress"); ok {
		t, err := expandSystemAlertemailFromaddress(d, v, "fromaddress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fromaddress"] = t
		}
	}

	if v, ok := d.GetOk("fromname"); ok {
		t, err := expandSystemAlertemailFromname(d, v, "fromname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fromname"] = t
		}
	}

	if v, ok := d.GetOk("smtppassword"); ok {
		t, err := expandSystemAlertemailSmtppassword(d, v, "smtppassword")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtppassword"] = t
		}
	}

	if v, ok := d.GetOk("smtpport"); ok {
		t, err := expandSystemAlertemailSmtpport(d, v, "smtpport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtpport"] = t
		}
	}

	if v, ok := d.GetOk("smtpserver"); ok {
		t, err := expandSystemAlertemailSmtpserver(d, v, "smtpserver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtpserver"] = t
		}
	}

	if v, ok := d.GetOk("smtpuser"); ok {
		t, err := expandSystemAlertemailSmtpuser(d, v, "smtpuser")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtpuser"] = t
		}
	}

	return &obj, nil
}
