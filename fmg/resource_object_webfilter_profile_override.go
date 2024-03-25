// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Web Filter override settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebfilterProfileOverride() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterProfileOverrideUpdate,
		Read:   resourceObjectWebfilterProfileOverrideRead,
		Update: resourceObjectWebfilterProfileOverrideUpdate,
		Delete: resourceObjectWebfilterProfileOverrideDelete,

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
			"url_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ovrd_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ovrd_dur": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ovrd_dur_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ovrd_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ovrd_user_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"profile_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWebfilterProfileOverrideUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("url_profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectWebfilterProfileOverride(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileOverride resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWebfilterProfileOverride(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileOverride resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectWebfilterProfileOverride")

	return resourceObjectWebfilterProfileOverrideRead(d, m)
}

func resourceObjectWebfilterProfileOverrideDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("url_profile").(string)
	paradict["profile"] = profile

	err = c.DeleteObjectWebfilterProfileOverride(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterProfileOverride resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterProfileOverrideRead(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("url_profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "url_profile")
		if profile == "" {
			return fmt.Errorf("Parameter url_profile is missing")
		}
		if err = d.Set("url_profile", profile); err != nil {
			return fmt.Errorf("Error set params url_profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadObjectWebfilterProfileOverride(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileOverride resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterProfileOverride(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileOverride resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterProfileOverrideOvrdCookie2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileOverrideOvrdDur2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileOverrideOvrdDurMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileOverrideOvrdScope2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileOverrideOvrdUserGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileOverrideProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileOverrideProfileAttribute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileOverrideProfileType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebfilterProfileOverride(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("ovrd_cookie", flattenObjectWebfilterProfileOverrideOvrdCookie2edl(o["ovrd-cookie"], d, "ovrd_cookie")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-cookie"], "ObjectWebfilterProfileOverride-OvrdCookie"); ok {
			if err = d.Set("ovrd_cookie", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_cookie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_cookie: %v", err)
		}
	}

	if err = d.Set("ovrd_dur", flattenObjectWebfilterProfileOverrideOvrdDur2edl(o["ovrd-dur"], d, "ovrd_dur")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-dur"], "ObjectWebfilterProfileOverride-OvrdDur"); ok {
			if err = d.Set("ovrd_dur", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_dur: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_dur: %v", err)
		}
	}

	if err = d.Set("ovrd_dur_mode", flattenObjectWebfilterProfileOverrideOvrdDurMode2edl(o["ovrd-dur-mode"], d, "ovrd_dur_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-dur-mode"], "ObjectWebfilterProfileOverride-OvrdDurMode"); ok {
			if err = d.Set("ovrd_dur_mode", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_dur_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_dur_mode: %v", err)
		}
	}

	if err = d.Set("ovrd_scope", flattenObjectWebfilterProfileOverrideOvrdScope2edl(o["ovrd-scope"], d, "ovrd_scope")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-scope"], "ObjectWebfilterProfileOverride-OvrdScope"); ok {
			if err = d.Set("ovrd_scope", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_scope: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_scope: %v", err)
		}
	}

	if err = d.Set("ovrd_user_group", flattenObjectWebfilterProfileOverrideOvrdUserGroup2edl(o["ovrd-user-group"], d, "ovrd_user_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-user-group"], "ObjectWebfilterProfileOverride-OvrdUserGroup"); ok {
			if err = d.Set("ovrd_user_group", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_user_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_user_group: %v", err)
		}
	}

	if err = d.Set("profile", flattenObjectWebfilterProfileOverrideProfile2edl(o["profile"], d, "profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile"], "ObjectWebfilterProfileOverride-Profile"); ok {
			if err = d.Set("profile", vv); err != nil {
				return fmt.Errorf("Error reading profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile: %v", err)
		}
	}

	if err = d.Set("profile_attribute", flattenObjectWebfilterProfileOverrideProfileAttribute2edl(o["profile-attribute"], d, "profile_attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-attribute"], "ObjectWebfilterProfileOverride-ProfileAttribute"); ok {
			if err = d.Set("profile_attribute", vv); err != nil {
				return fmt.Errorf("Error reading profile_attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_attribute: %v", err)
		}
	}

	if err = d.Set("profile_type", flattenObjectWebfilterProfileOverrideProfileType2edl(o["profile-type"], d, "profile_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-type"], "ObjectWebfilterProfileOverride-ProfileType"); ok {
			if err = d.Set("profile_type", vv); err != nil {
				return fmt.Errorf("Error reading profile_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_type: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterProfileOverrideFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterProfileOverrideOvrdCookie2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileOverrideOvrdDur2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileOverrideOvrdDurMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileOverrideOvrdScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileOverrideOvrdUserGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileOverrideProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileOverrideProfileAttribute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileOverrideProfileType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebfilterProfileOverride(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ovrd_cookie"); ok || d.HasChange("ovrd_cookie") {
		t, err := expandObjectWebfilterProfileOverrideOvrdCookie2edl(d, v, "ovrd_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-cookie"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_dur"); ok || d.HasChange("ovrd_dur") {
		t, err := expandObjectWebfilterProfileOverrideOvrdDur2edl(d, v, "ovrd_dur")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-dur"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_dur_mode"); ok || d.HasChange("ovrd_dur_mode") {
		t, err := expandObjectWebfilterProfileOverrideOvrdDurMode2edl(d, v, "ovrd_dur_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-dur-mode"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_scope"); ok || d.HasChange("ovrd_scope") {
		t, err := expandObjectWebfilterProfileOverrideOvrdScope2edl(d, v, "ovrd_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-scope"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_user_group"); ok || d.HasChange("ovrd_user_group") {
		t, err := expandObjectWebfilterProfileOverrideOvrdUserGroup2edl(d, v, "ovrd_user_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-user-group"] = t
		}
	}

	if v, ok := d.GetOk("profile"); ok || d.HasChange("profile") {
		t, err := expandObjectWebfilterProfileOverrideProfile2edl(d, v, "profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile"] = t
		}
	}

	if v, ok := d.GetOk("profile_attribute"); ok || d.HasChange("profile_attribute") {
		t, err := expandObjectWebfilterProfileOverrideProfileAttribute2edl(d, v, "profile_attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-attribute"] = t
		}
	}

	if v, ok := d.GetOk("profile_type"); ok || d.HasChange("profile_type") {
		t, err := expandObjectWebfilterProfileOverrideProfileType2edl(d, v, "profile_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-type"] = t
		}
	}

	return &obj, nil
}
