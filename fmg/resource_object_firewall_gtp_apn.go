// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: APN.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallGtpApn() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallGtpApnCreate,
		Read:   resourceObjectFirewallGtpApnRead,
		Update: resourceObjectFirewallGtpApnUpdate,
		Delete: resourceObjectFirewallGtpApnDelete,

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
			"gtp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apnmember": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"selection_mode": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallGtpApnCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	gtp := d.Get("gtp").(string)
	paradict["gtp"] = gtp

	obj, err := getObjectObjectFirewallGtpApn(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallGtpApn resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallGtpApn(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallGtpApn resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallGtpApnRead(d, m)
}

func resourceObjectFirewallGtpApnUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	gtp := d.Get("gtp").(string)
	paradict["gtp"] = gtp

	obj, err := getObjectObjectFirewallGtpApn(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpApn resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallGtpApn(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpApn resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallGtpApnRead(d, m)
}

func resourceObjectFirewallGtpApnDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	gtp := d.Get("gtp").(string)
	paradict["gtp"] = gtp

	err = c.DeleteObjectFirewallGtpApn(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallGtpApn resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallGtpApnRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	gtp := d.Get("gtp").(string)
	if gtp == "" {
		gtp = importOptionChecking(m.(*FortiClient).Cfg, "gtp")
		if gtp == "" {
			return fmt.Errorf("Parameter gtp is missing")
		}
		if err = d.Set("gtp", gtp); err != nil {
			return fmt.Errorf("Error set params gtp: %v", err)
		}
	}
	paradict["gtp"] = gtp

	o, err := c.ReadObjectFirewallGtpApn(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpApn resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallGtpApn(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpApn resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallGtpApnAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpApnApnmember2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpApnId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpApnSelectionMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectFirewallGtpApn(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectFirewallGtpApnAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectFirewallGtpApn-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("apnmember", flattenObjectFirewallGtpApnApnmember2edl(o["apnmember"], d, "apnmember")); err != nil {
		if vv, ok := fortiAPIPatch(o["apnmember"], "ObjectFirewallGtpApn-Apnmember"); ok {
			if err = d.Set("apnmember", vv); err != nil {
				return fmt.Errorf("Error reading apnmember: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apnmember: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallGtpApnId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallGtpApn-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("selection_mode", flattenObjectFirewallGtpApnSelectionMode2edl(o["selection-mode"], d, "selection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["selection-mode"], "ObjectFirewallGtpApn-SelectionMode"); ok {
			if err = d.Set("selection_mode", vv); err != nil {
				return fmt.Errorf("Error reading selection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading selection_mode: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallGtpApnFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallGtpApnAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpApnApnmember2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpApnId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpApnSelectionMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectFirewallGtpApn(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectFirewallGtpApnAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("apnmember"); ok || d.HasChange("apnmember") {
		t, err := expandObjectFirewallGtpApnApnmember2edl(d, v, "apnmember")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apnmember"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallGtpApnId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("selection_mode"); ok || d.HasChange("selection_mode") {
		t, err := expandObjectFirewallGtpApnSelectionMode2edl(d, v, "selection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selection-mode"] = t
		}
	}

	return &obj, nil
}
