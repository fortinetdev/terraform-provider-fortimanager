// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: HTTP header group.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectIcapProfileRespmodForwardRulesHeaderGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectIcapProfileRespmodForwardRulesHeaderGroupCreate,
		Read:   resourceObjectIcapProfileRespmodForwardRulesHeaderGroupRead,
		Update: resourceObjectIcapProfileRespmodForwardRulesHeaderGroupUpdate,
		Delete: resourceObjectIcapProfileRespmodForwardRulesHeaderGroupDelete,

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
			"respmod_forward_rules": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"case_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"header_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectIcapProfileRespmodForwardRulesHeaderGroupCreate(d *schema.ResourceData, m interface{}) error {
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
	respmod_forward_rules := d.Get("respmod_forward_rules").(string)
	paradict["profile"] = profile
	paradict["respmod_forward_rules"] = respmod_forward_rules

	obj, err := getObjectObjectIcapProfileRespmodForwardRulesHeaderGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIcapProfileRespmodForwardRulesHeaderGroup resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectIcapProfileRespmodForwardRulesHeaderGroup(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIcapProfileRespmodForwardRulesHeaderGroup resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectIcapProfileRespmodForwardRulesHeaderGroupRead(d, m)
}

func resourceObjectIcapProfileRespmodForwardRulesHeaderGroupUpdate(d *schema.ResourceData, m interface{}) error {
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
	respmod_forward_rules := d.Get("respmod_forward_rules").(string)
	paradict["profile"] = profile
	paradict["respmod_forward_rules"] = respmod_forward_rules

	obj, err := getObjectObjectIcapProfileRespmodForwardRulesHeaderGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIcapProfileRespmodForwardRulesHeaderGroup resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectIcapProfileRespmodForwardRulesHeaderGroup(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIcapProfileRespmodForwardRulesHeaderGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectIcapProfileRespmodForwardRulesHeaderGroupRead(d, m)
}

func resourceObjectIcapProfileRespmodForwardRulesHeaderGroupDelete(d *schema.ResourceData, m interface{}) error {
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
	respmod_forward_rules := d.Get("respmod_forward_rules").(string)
	paradict["profile"] = profile
	paradict["respmod_forward_rules"] = respmod_forward_rules

	wsParams["adom"] = adomv

	err = c.DeleteObjectIcapProfileRespmodForwardRulesHeaderGroup(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectIcapProfileRespmodForwardRulesHeaderGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectIcapProfileRespmodForwardRulesHeaderGroupRead(d *schema.ResourceData, m interface{}) error {
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
	respmod_forward_rules := d.Get("respmod_forward_rules").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	if respmod_forward_rules == "" {
		respmod_forward_rules = importOptionChecking(m.(*FortiClient).Cfg, "respmod_forward_rules")
		if respmod_forward_rules == "" {
			return fmt.Errorf("Parameter respmod_forward_rules is missing")
		}
		if err = d.Set("respmod_forward_rules", respmod_forward_rules); err != nil {
			return fmt.Errorf("Error set params respmod_forward_rules: %v", err)
		}
	}
	paradict["profile"] = profile
	paradict["respmod_forward_rules"] = respmod_forward_rules

	o, err := c.ReadObjectIcapProfileRespmodForwardRulesHeaderGroup(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIcapProfileRespmodForwardRulesHeaderGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectIcapProfileRespmodForwardRulesHeaderGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIcapProfileRespmodForwardRulesHeaderGroup resource from API: %v", err)
	}
	return nil
}

func flattenObjectIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRespmodForwardRulesHeaderGroupHeader3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRespmodForwardRulesHeaderGroupHeaderName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRespmodForwardRulesHeaderGroupId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectIcapProfileRespmodForwardRulesHeaderGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("case_sensitivity", flattenObjectIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity3rdl(o["case-sensitivity"], d, "case_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["case-sensitivity"], "ObjectIcapProfileRespmodForwardRulesHeaderGroup-CaseSensitivity"); ok {
			if err = d.Set("case_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading case_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading case_sensitivity: %v", err)
		}
	}

	if err = d.Set("header", flattenObjectIcapProfileRespmodForwardRulesHeaderGroupHeader3rdl(o["header"], d, "header")); err != nil {
		if vv, ok := fortiAPIPatch(o["header"], "ObjectIcapProfileRespmodForwardRulesHeaderGroup-Header"); ok {
			if err = d.Set("header", vv); err != nil {
				return fmt.Errorf("Error reading header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if err = d.Set("header_name", flattenObjectIcapProfileRespmodForwardRulesHeaderGroupHeaderName3rdl(o["header-name"], d, "header_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-name"], "ObjectIcapProfileRespmodForwardRulesHeaderGroup-HeaderName"); ok {
			if err = d.Set("header_name", vv); err != nil {
				return fmt.Errorf("Error reading header_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectIcapProfileRespmodForwardRulesHeaderGroupId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectIcapProfileRespmodForwardRulesHeaderGroup-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenObjectIcapProfileRespmodForwardRulesHeaderGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRespmodForwardRulesHeaderGroupHeader3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRespmodForwardRulesHeaderGroupHeaderName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRespmodForwardRulesHeaderGroupId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectIcapProfileRespmodForwardRulesHeaderGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("case_sensitivity"); ok || d.HasChange("case_sensitivity") {
		t, err := expandObjectIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity3rdl(d, v, "case_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("header"); ok || d.HasChange("header") {
		t, err := expandObjectIcapProfileRespmodForwardRulesHeaderGroupHeader3rdl(d, v, "header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}

	if v, ok := d.GetOk("header_name"); ok || d.HasChange("header_name") {
		t, err := expandObjectIcapProfileRespmodForwardRulesHeaderGroupHeaderName3rdl(d, v, "header_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectIcapProfileRespmodForwardRulesHeaderGroupId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
