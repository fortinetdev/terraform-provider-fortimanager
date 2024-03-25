// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiGuard filters.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebfilterProfileFtgdWfFilters() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterProfileFtgdWfFiltersCreate,
		Read:   resourceObjectWebfilterProfileFtgdWfFiltersRead,
		Update: resourceObjectWebfilterProfileFtgdWfFiltersUpdate,
		Delete: resourceObjectWebfilterProfileFtgdWfFiltersDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_usr_grp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_replacemsg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"warn_duration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"warning_duration_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"warning_prompt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectWebfilterProfileFtgdWfFiltersCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectWebfilterProfileFtgdWfFilters(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterProfileFtgdWfFilters resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWebfilterProfileFtgdWfFilters(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterProfileFtgdWfFilters resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWebfilterProfileFtgdWfFiltersRead(d, m)
}

func resourceObjectWebfilterProfileFtgdWfFiltersUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectWebfilterProfileFtgdWfFilters(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileFtgdWfFilters resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWebfilterProfileFtgdWfFilters(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileFtgdWfFilters resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWebfilterProfileFtgdWfFiltersRead(d, m)
}

func resourceObjectWebfilterProfileFtgdWfFiltersDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	err = c.DeleteObjectWebfilterProfileFtgdWfFilters(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterProfileFtgdWfFilters resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterProfileFtgdWfFiltersRead(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadObjectWebfilterProfileFtgdWfFilters(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileFtgdWfFilters resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterProfileFtgdWfFilters(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileFtgdWfFilters resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterProfileFtgdWfFiltersAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersAuthUsrGrp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWebfilterProfileFtgdWfFiltersCategory3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWebfilterProfileFtgdWfFiltersId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersLog3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersOverrideReplacemsg3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersWarnDuration3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersWarningDurationType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersWarningPrompt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebfilterProfileFtgdWfFilters(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectWebfilterProfileFtgdWfFiltersAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectWebfilterProfileFtgdWfFilters-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("auth_usr_grp", flattenObjectWebfilterProfileFtgdWfFiltersAuthUsrGrp3rdl(o["auth-usr-grp"], d, "auth_usr_grp")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-usr-grp"], "ObjectWebfilterProfileFtgdWfFilters-AuthUsrGrp"); ok {
			if err = d.Set("auth_usr_grp", vv); err != nil {
				return fmt.Errorf("Error reading auth_usr_grp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_usr_grp: %v", err)
		}
	}

	if err = d.Set("category", flattenObjectWebfilterProfileFtgdWfFiltersCategory3rdl(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "ObjectWebfilterProfileFtgdWfFilters-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectWebfilterProfileFtgdWfFiltersId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWebfilterProfileFtgdWfFilters-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("log", flattenObjectWebfilterProfileFtgdWfFiltersLog3rdl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ObjectWebfilterProfileFtgdWfFilters-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("override_replacemsg", flattenObjectWebfilterProfileFtgdWfFiltersOverrideReplacemsg3rdl(o["override-replacemsg"], d, "override_replacemsg")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-replacemsg"], "ObjectWebfilterProfileFtgdWfFilters-OverrideReplacemsg"); ok {
			if err = d.Set("override_replacemsg", vv); err != nil {
				return fmt.Errorf("Error reading override_replacemsg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_replacemsg: %v", err)
		}
	}

	if err = d.Set("warn_duration", flattenObjectWebfilterProfileFtgdWfFiltersWarnDuration3rdl(o["warn-duration"], d, "warn_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["warn-duration"], "ObjectWebfilterProfileFtgdWfFilters-WarnDuration"); ok {
			if err = d.Set("warn_duration", vv); err != nil {
				return fmt.Errorf("Error reading warn_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading warn_duration: %v", err)
		}
	}

	if err = d.Set("warning_duration_type", flattenObjectWebfilterProfileFtgdWfFiltersWarningDurationType3rdl(o["warning-duration-type"], d, "warning_duration_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["warning-duration-type"], "ObjectWebfilterProfileFtgdWfFilters-WarningDurationType"); ok {
			if err = d.Set("warning_duration_type", vv); err != nil {
				return fmt.Errorf("Error reading warning_duration_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading warning_duration_type: %v", err)
		}
	}

	if err = d.Set("warning_prompt", flattenObjectWebfilterProfileFtgdWfFiltersWarningPrompt3rdl(o["warning-prompt"], d, "warning_prompt")); err != nil {
		if vv, ok := fortiAPIPatch(o["warning-prompt"], "ObjectWebfilterProfileFtgdWfFilters-WarningPrompt"); ok {
			if err = d.Set("warning_prompt", vv); err != nil {
				return fmt.Errorf("Error reading warning_prompt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading warning_prompt: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterProfileFtgdWfFiltersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterProfileFtgdWfFiltersAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersAuthUsrGrp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWebfilterProfileFtgdWfFiltersCategory3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersLog3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersOverrideReplacemsg3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersWarnDuration3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersWarningDurationType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersWarningPrompt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebfilterProfileFtgdWfFilters(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectWebfilterProfileFtgdWfFiltersAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("auth_usr_grp"); ok || d.HasChange("auth_usr_grp") {
		t, err := expandObjectWebfilterProfileFtgdWfFiltersAuthUsrGrp3rdl(d, v, "auth_usr_grp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-usr-grp"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandObjectWebfilterProfileFtgdWfFiltersCategory3rdl(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWebfilterProfileFtgdWfFiltersId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandObjectWebfilterProfileFtgdWfFiltersLog3rdl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("override_replacemsg"); ok || d.HasChange("override_replacemsg") {
		t, err := expandObjectWebfilterProfileFtgdWfFiltersOverrideReplacemsg3rdl(d, v, "override_replacemsg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-replacemsg"] = t
		}
	}

	if v, ok := d.GetOk("warn_duration"); ok || d.HasChange("warn_duration") {
		t, err := expandObjectWebfilterProfileFtgdWfFiltersWarnDuration3rdl(d, v, "warn_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warn-duration"] = t
		}
	}

	if v, ok := d.GetOk("warning_duration_type"); ok || d.HasChange("warning_duration_type") {
		t, err := expandObjectWebfilterProfileFtgdWfFiltersWarningDurationType3rdl(d, v, "warning_duration_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warning-duration-type"] = t
		}
	}

	if v, ok := d.GetOk("warning_prompt"); ok || d.HasChange("warning_prompt") {
		t, err := expandObjectWebfilterProfileFtgdWfFiltersWarningPrompt3rdl(d, v, "warning_prompt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warning-prompt"] = t
		}
	}

	return &obj, nil
}
