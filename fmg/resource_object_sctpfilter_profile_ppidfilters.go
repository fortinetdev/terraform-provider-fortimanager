// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: PPID filters list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSctpFilterProfilePpidFilters() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSctpFilterProfilePpidFiltersCreate,
		Read:   resourceObjectSctpFilterProfilePpidFiltersRead,
		Update: resourceObjectSctpFilterProfilePpidFiltersUpdate,
		Delete: resourceObjectSctpFilterProfilePpidFiltersDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"ppid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectSctpFilterProfilePpidFiltersCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectSctpFilterProfilePpidFilters(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSctpFilterProfilePpidFilters resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	v, err := c.CreateObjectSctpFilterProfilePpidFilters(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSctpFilterProfilePpidFilters resource: %v", err)
	}

	if v != nil && v["id"] != nil {
		if vidn, ok := v["id"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceObjectSctpFilterProfilePpidFiltersRead(d, m)
		} else {
			return fmt.Errorf("Error creating ObjectSctpFilterProfilePpidFilters resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSctpFilterProfilePpidFiltersRead(d, m)
}

func resourceObjectSctpFilterProfilePpidFiltersUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectSctpFilterProfilePpidFilters(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSctpFilterProfilePpidFilters resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSctpFilterProfilePpidFilters(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSctpFilterProfilePpidFilters resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSctpFilterProfilePpidFiltersRead(d, m)
}

func resourceObjectSctpFilterProfilePpidFiltersDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectSctpFilterProfilePpidFilters(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSctpFilterProfilePpidFilters resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSctpFilterProfilePpidFiltersRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSctpFilterProfilePpidFilters(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSctpFilterProfilePpidFilters resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSctpFilterProfilePpidFilters(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSctpFilterProfilePpidFilters resource from API: %v", err)
	}
	return nil
}

func flattenObjectSctpFilterProfilePpidFiltersAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSctpFilterProfilePpidFiltersComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSctpFilterProfilePpidFiltersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSctpFilterProfilePpidFiltersPpid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSctpFilterProfilePpidFilters(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectSctpFilterProfilePpidFiltersAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectSctpFilterProfilePpidFilters-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectSctpFilterProfilePpidFiltersComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectSctpFilterProfilePpidFilters-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectSctpFilterProfilePpidFiltersId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSctpFilterProfilePpidFilters-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ppid", flattenObjectSctpFilterProfilePpidFiltersPpid2edl(o["ppid"], d, "ppid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ppid"], "ObjectSctpFilterProfilePpidFilters-Ppid"); ok {
			if err = d.Set("ppid", vv); err != nil {
				return fmt.Errorf("Error reading ppid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ppid: %v", err)
		}
	}

	return nil
}

func flattenObjectSctpFilterProfilePpidFiltersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSctpFilterProfilePpidFiltersAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSctpFilterProfilePpidFiltersComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSctpFilterProfilePpidFiltersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSctpFilterProfilePpidFiltersPpid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSctpFilterProfilePpidFilters(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectSctpFilterProfilePpidFiltersAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectSctpFilterProfilePpidFiltersComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectSctpFilterProfilePpidFiltersId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ppid"); ok || d.HasChange("ppid") {
		t, err := expandObjectSctpFilterProfilePpidFiltersPpid2edl(d, v, "ppid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppid"] = t
		}
	}

	return &obj, nil
}
