// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard - AntiSpam.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectEmailfilterFortishield() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectEmailfilterFortishieldUpdate,
		Read:   resourceObjectEmailfilterFortishieldRead,
		Update: resourceObjectEmailfilterFortishieldUpdate,
		Delete: resourceObjectEmailfilterFortishieldDelete,

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
			"spam_submit_force": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_submit_srv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_submit_txt2htm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectEmailfilterFortishieldUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectEmailfilterFortishield(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEmailfilterFortishield resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectEmailfilterFortishield(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEmailfilterFortishield resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectEmailfilterFortishield")

	return resourceObjectEmailfilterFortishieldRead(d, m)
}

func resourceObjectEmailfilterFortishieldDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectEmailfilterFortishield(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectEmailfilterFortishield resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectEmailfilterFortishieldRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectEmailfilterFortishield(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEmailfilterFortishield resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectEmailfilterFortishield(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEmailfilterFortishield resource from API: %v", err)
	}
	return nil
}

func flattenObjectEmailfilterFortishieldSpamSubmitForce(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterFortishieldSpamSubmitSrv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterFortishieldSpamSubmitTxt2Htm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectEmailfilterFortishield(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("spam_submit_force", flattenObjectEmailfilterFortishieldSpamSubmitForce(o["spam-submit-force"], d, "spam_submit_force")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-submit-force"], "ObjectEmailfilterFortishield-SpamSubmitForce"); ok {
			if err = d.Set("spam_submit_force", vv); err != nil {
				return fmt.Errorf("Error reading spam_submit_force: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_submit_force: %v", err)
		}
	}

	if err = d.Set("spam_submit_srv", flattenObjectEmailfilterFortishieldSpamSubmitSrv(o["spam-submit-srv"], d, "spam_submit_srv")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-submit-srv"], "ObjectEmailfilterFortishield-SpamSubmitSrv"); ok {
			if err = d.Set("spam_submit_srv", vv); err != nil {
				return fmt.Errorf("Error reading spam_submit_srv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_submit_srv: %v", err)
		}
	}

	if err = d.Set("spam_submit_txt2htm", flattenObjectEmailfilterFortishieldSpamSubmitTxt2Htm(o["spam-submit-txt2htm"], d, "spam_submit_txt2htm")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-submit-txt2htm"], "ObjectEmailfilterFortishield-SpamSubmitTxt2Htm"); ok {
			if err = d.Set("spam_submit_txt2htm", vv); err != nil {
				return fmt.Errorf("Error reading spam_submit_txt2htm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_submit_txt2htm: %v", err)
		}
	}

	return nil
}

func flattenObjectEmailfilterFortishieldFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectEmailfilterFortishieldSpamSubmitForce(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterFortishieldSpamSubmitSrv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterFortishieldSpamSubmitTxt2Htm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectEmailfilterFortishield(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("spam_submit_force"); ok || d.HasChange("spam_submit_force") {
		t, err := expandObjectEmailfilterFortishieldSpamSubmitForce(d, v, "spam_submit_force")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-submit-force"] = t
		}
	}

	if v, ok := d.GetOk("spam_submit_srv"); ok || d.HasChange("spam_submit_srv") {
		t, err := expandObjectEmailfilterFortishieldSpamSubmitSrv(d, v, "spam_submit_srv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-submit-srv"] = t
		}
	}

	if v, ok := d.GetOk("spam_submit_txt2htm"); ok || d.HasChange("spam_submit_txt2htm") {
		t, err := expandObjectEmailfilterFortishieldSpamSubmitTxt2Htm(d, v, "spam_submit_txt2htm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-submit-txt2htm"] = t
		}
	}

	return &obj, nil
}
