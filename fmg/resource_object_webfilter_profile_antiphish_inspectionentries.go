// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: AntiPhishing entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebfilterProfileAntiphishInspectionEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterProfileAntiphishInspectionEntriesCreate,
		Read:   resourceObjectWebfilterProfileAntiphishInspectionEntriesRead,
		Update: resourceObjectWebfilterProfileAntiphishInspectionEntriesUpdate,
		Delete: resourceObjectWebfilterProfileAntiphishInspectionEntriesDelete,

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
			"fortiguard_category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectWebfilterProfileAntiphishInspectionEntriesCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWebfilterProfileAntiphishInspectionEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterProfileAntiphishInspectionEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWebfilterProfileAntiphishInspectionEntries(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterProfileAntiphishInspectionEntries resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebfilterProfileAntiphishInspectionEntriesRead(d, m)
}

func resourceObjectWebfilterProfileAntiphishInspectionEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWebfilterProfileAntiphishInspectionEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileAntiphishInspectionEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWebfilterProfileAntiphishInspectionEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileAntiphishInspectionEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebfilterProfileAntiphishInspectionEntriesRead(d, m)
}

func resourceObjectWebfilterProfileAntiphishInspectionEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWebfilterProfileAntiphishInspectionEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterProfileAntiphishInspectionEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterProfileAntiphishInspectionEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWebfilterProfileAntiphishInspectionEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileAntiphishInspectionEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterProfileAntiphishInspectionEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileAntiphishInspectionEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterProfileAntiphishInspectionEntriesAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishInspectionEntriesFortiguardCategory3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileAntiphishInspectionEntriesName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebfilterProfileAntiphishInspectionEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectWebfilterProfileAntiphishInspectionEntriesAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectWebfilterProfileAntiphishInspectionEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("fortiguard_category", flattenObjectWebfilterProfileAntiphishInspectionEntriesFortiguardCategory3rdl(o["fortiguard-category"], d, "fortiguard_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiguard-category"], "ObjectWebfilterProfileAntiphishInspectionEntries-FortiguardCategory"); ok {
			if err = d.Set("fortiguard_category", vv); err != nil {
				return fmt.Errorf("Error reading fortiguard_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiguard_category: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWebfilterProfileAntiphishInspectionEntriesName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWebfilterProfileAntiphishInspectionEntries-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterProfileAntiphishInspectionEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterProfileAntiphishInspectionEntriesAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishInspectionEntriesFortiguardCategory3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileAntiphishInspectionEntriesName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebfilterProfileAntiphishInspectionEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectWebfilterProfileAntiphishInspectionEntriesAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_category"); ok || d.HasChange("fortiguard_category") {
		t, err := expandObjectWebfilterProfileAntiphishInspectionEntriesFortiguardCategory3rdl(d, v, "fortiguard_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-category"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWebfilterProfileAntiphishInspectionEntriesName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
