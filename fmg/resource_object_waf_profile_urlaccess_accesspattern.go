// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: URL access pattern.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWafProfileUrlAccessAccessPattern() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWafProfileUrlAccessAccessPatternCreate,
		Read:   resourceObjectWafProfileUrlAccessAccessPatternRead,
		Update: resourceObjectWafProfileUrlAccessAccessPatternUpdate,
		Delete: resourceObjectWafProfileUrlAccessAccessPatternDelete,

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
			"url_access": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectWafProfileUrlAccessAccessPatternCreate(d *schema.ResourceData, m interface{}) error {
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
	url_access := d.Get("url_access").(string)
	paradict["profile"] = profile
	paradict["url_access"] = url_access

	obj, err := getObjectObjectWafProfileUrlAccessAccessPattern(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWafProfileUrlAccessAccessPattern resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWafProfileUrlAccessAccessPattern(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWafProfileUrlAccessAccessPattern resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWafProfileUrlAccessAccessPatternRead(d, m)
}

func resourceObjectWafProfileUrlAccessAccessPatternUpdate(d *schema.ResourceData, m interface{}) error {
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
	url_access := d.Get("url_access").(string)
	paradict["profile"] = profile
	paradict["url_access"] = url_access

	obj, err := getObjectObjectWafProfileUrlAccessAccessPattern(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfileUrlAccessAccessPattern resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWafProfileUrlAccessAccessPattern(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfileUrlAccessAccessPattern resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWafProfileUrlAccessAccessPatternRead(d, m)
}

func resourceObjectWafProfileUrlAccessAccessPatternDelete(d *schema.ResourceData, m interface{}) error {
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
	url_access := d.Get("url_access").(string)
	paradict["profile"] = profile
	paradict["url_access"] = url_access

	err = c.DeleteObjectWafProfileUrlAccessAccessPattern(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWafProfileUrlAccessAccessPattern resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWafProfileUrlAccessAccessPatternRead(d *schema.ResourceData, m interface{}) error {
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
	url_access := d.Get("url_access").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	if url_access == "" {
		url_access = importOptionChecking(m.(*FortiClient).Cfg, "url_access")
		if url_access == "" {
			return fmt.Errorf("Parameter url_access is missing")
		}
		if err = d.Set("url_access", url_access); err != nil {
			return fmt.Errorf("Error set params url_access: %v", err)
		}
	}
	paradict["profile"] = profile
	paradict["url_access"] = url_access

	o, err := c.ReadObjectWafProfileUrlAccessAccessPattern(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfileUrlAccessAccessPattern resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWafProfileUrlAccessAccessPattern(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfileUrlAccessAccessPattern resource from API: %v", err)
	}
	return nil
}

func flattenObjectWafProfileUrlAccessAccessPatternId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAccessPatternNegate3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAccessPatternPattern3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAccessPatternRegex3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAccessPatternSrcaddr3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectObjectWafProfileUrlAccessAccessPattern(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenObjectWafProfileUrlAccessAccessPatternId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWafProfileUrlAccessAccessPattern-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("negate", flattenObjectWafProfileUrlAccessAccessPatternNegate3rdl(o["negate"], d, "negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["negate"], "ObjectWafProfileUrlAccessAccessPattern-Negate"); ok {
			if err = d.Set("negate", vv); err != nil {
				return fmt.Errorf("Error reading negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading negate: %v", err)
		}
	}

	if err = d.Set("pattern", flattenObjectWafProfileUrlAccessAccessPatternPattern3rdl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "ObjectWafProfileUrlAccessAccessPattern-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("regex", flattenObjectWafProfileUrlAccessAccessPatternRegex3rdl(o["regex"], d, "regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["regex"], "ObjectWafProfileUrlAccessAccessPattern-Regex"); ok {
			if err = d.Set("regex", vv); err != nil {
				return fmt.Errorf("Error reading regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading regex: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenObjectWafProfileUrlAccessAccessPatternSrcaddr3rdl(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "ObjectWafProfileUrlAccessAccessPattern-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	return nil
}

func flattenObjectWafProfileUrlAccessAccessPatternFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWafProfileUrlAccessAccessPatternId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAccessPatternNegate3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAccessPatternPattern3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAccessPatternRegex3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAccessPatternSrcaddr3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectObjectWafProfileUrlAccessAccessPattern(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWafProfileUrlAccessAccessPatternId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("negate"); ok || d.HasChange("negate") {
		t, err := expandObjectWafProfileUrlAccessAccessPatternNegate3rdl(d, v, "negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negate"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandObjectWafProfileUrlAccessAccessPatternPattern3rdl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("regex"); ok || d.HasChange("regex") {
		t, err := expandObjectWafProfileUrlAccessAccessPatternRegex3rdl(d, v, "regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["regex"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandObjectWafProfileUrlAccessAccessPatternSrcaddr3rdl(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	return &obj, nil
}
