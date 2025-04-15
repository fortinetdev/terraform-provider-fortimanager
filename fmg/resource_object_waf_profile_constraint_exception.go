// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: HTTP constraint exception.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWafProfileConstraintException() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWafProfileConstraintExceptionCreate,
		Read:   resourceObjectWafProfileConstraintExceptionRead,
		Update: resourceObjectWafProfileConstraintExceptionUpdate,
		Delete: resourceObjectWafProfileConstraintExceptionDelete,

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
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"content_length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"line_length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_header_line": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_range_segment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_url_param": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"param_length": &schema.Schema{
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
			"url_param_length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWafProfileConstraintExceptionCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWafProfileConstraintException(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWafProfileConstraintException resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWafProfileConstraintException(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWafProfileConstraintException resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWafProfileConstraintExceptionRead(d, m)
}

func resourceObjectWafProfileConstraintExceptionUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWafProfileConstraintException(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfileConstraintException resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWafProfileConstraintException(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfileConstraintException resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWafProfileConstraintExceptionRead(d, m)
}

func resourceObjectWafProfileConstraintExceptionDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWafProfileConstraintException(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWafProfileConstraintException resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWafProfileConstraintExceptionRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWafProfileConstraintException(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfileConstraintException resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWafProfileConstraintException(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfileConstraintException resource from API: %v", err)
	}
	return nil
}

func flattenObjectWafProfileConstraintExceptionAddress3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWafProfileConstraintExceptionContentLength3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionHeaderLength3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionHostname3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionLineLength3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMalformed3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMaxCookie3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMaxHeaderLine3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMaxRangeSegment3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMaxUrlParam3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMethod3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionParamLength3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionPattern3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionRegex3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionUrlParamLength3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionVersion3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWafProfileConstraintException(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("address", flattenObjectWafProfileConstraintExceptionAddress3rdl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "ObjectWafProfileConstraintException-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("content_length", flattenObjectWafProfileConstraintExceptionContentLength3rdl(o["content-length"], d, "content_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["content-length"], "ObjectWafProfileConstraintException-ContentLength"); ok {
			if err = d.Set("content_length", vv); err != nil {
				return fmt.Errorf("Error reading content_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading content_length: %v", err)
		}
	}

	if err = d.Set("header_length", flattenObjectWafProfileConstraintExceptionHeaderLength3rdl(o["header-length"], d, "header_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-length"], "ObjectWafProfileConstraintException-HeaderLength"); ok {
			if err = d.Set("header_length", vv); err != nil {
				return fmt.Errorf("Error reading header_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_length: %v", err)
		}
	}

	if err = d.Set("hostname", flattenObjectWafProfileConstraintExceptionHostname3rdl(o["hostname"], d, "hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname"], "ObjectWafProfileConstraintException-Hostname"); ok {
			if err = d.Set("hostname", vv); err != nil {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectWafProfileConstraintExceptionId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWafProfileConstraintException-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("line_length", flattenObjectWafProfileConstraintExceptionLineLength3rdl(o["line-length"], d, "line_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["line-length"], "ObjectWafProfileConstraintException-LineLength"); ok {
			if err = d.Set("line_length", vv); err != nil {
				return fmt.Errorf("Error reading line_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading line_length: %v", err)
		}
	}

	if err = d.Set("malformed", flattenObjectWafProfileConstraintExceptionMalformed3rdl(o["malformed"], d, "malformed")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed"], "ObjectWafProfileConstraintException-Malformed"); ok {
			if err = d.Set("malformed", vv); err != nil {
				return fmt.Errorf("Error reading malformed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed: %v", err)
		}
	}

	if err = d.Set("max_cookie", flattenObjectWafProfileConstraintExceptionMaxCookie3rdl(o["max-cookie"], d, "max_cookie")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-cookie"], "ObjectWafProfileConstraintException-MaxCookie"); ok {
			if err = d.Set("max_cookie", vv); err != nil {
				return fmt.Errorf("Error reading max_cookie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_cookie: %v", err)
		}
	}

	if err = d.Set("max_header_line", flattenObjectWafProfileConstraintExceptionMaxHeaderLine3rdl(o["max-header-line"], d, "max_header_line")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-header-line"], "ObjectWafProfileConstraintException-MaxHeaderLine"); ok {
			if err = d.Set("max_header_line", vv); err != nil {
				return fmt.Errorf("Error reading max_header_line: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_header_line: %v", err)
		}
	}

	if err = d.Set("max_range_segment", flattenObjectWafProfileConstraintExceptionMaxRangeSegment3rdl(o["max-range-segment"], d, "max_range_segment")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-range-segment"], "ObjectWafProfileConstraintException-MaxRangeSegment"); ok {
			if err = d.Set("max_range_segment", vv); err != nil {
				return fmt.Errorf("Error reading max_range_segment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_range_segment: %v", err)
		}
	}

	if err = d.Set("max_url_param", flattenObjectWafProfileConstraintExceptionMaxUrlParam3rdl(o["max-url-param"], d, "max_url_param")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-url-param"], "ObjectWafProfileConstraintException-MaxUrlParam"); ok {
			if err = d.Set("max_url_param", vv); err != nil {
				return fmt.Errorf("Error reading max_url_param: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_url_param: %v", err)
		}
	}

	if err = d.Set("method", flattenObjectWafProfileConstraintExceptionMethod3rdl(o["method"], d, "method")); err != nil {
		if vv, ok := fortiAPIPatch(o["method"], "ObjectWafProfileConstraintException-Method"); ok {
			if err = d.Set("method", vv); err != nil {
				return fmt.Errorf("Error reading method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("param_length", flattenObjectWafProfileConstraintExceptionParamLength3rdl(o["param-length"], d, "param_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["param-length"], "ObjectWafProfileConstraintException-ParamLength"); ok {
			if err = d.Set("param_length", vv); err != nil {
				return fmt.Errorf("Error reading param_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading param_length: %v", err)
		}
	}

	if err = d.Set("pattern", flattenObjectWafProfileConstraintExceptionPattern3rdl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "ObjectWafProfileConstraintException-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("regex", flattenObjectWafProfileConstraintExceptionRegex3rdl(o["regex"], d, "regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["regex"], "ObjectWafProfileConstraintException-Regex"); ok {
			if err = d.Set("regex", vv); err != nil {
				return fmt.Errorf("Error reading regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading regex: %v", err)
		}
	}

	if err = d.Set("url_param_length", flattenObjectWafProfileConstraintExceptionUrlParamLength3rdl(o["url-param-length"], d, "url_param_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-param-length"], "ObjectWafProfileConstraintException-UrlParamLength"); ok {
			if err = d.Set("url_param_length", vv); err != nil {
				return fmt.Errorf("Error reading url_param_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_param_length: %v", err)
		}
	}

	if err = d.Set("version", flattenObjectWafProfileConstraintExceptionVersion3rdl(o["version"], d, "version")); err != nil {
		if vv, ok := fortiAPIPatch(o["version"], "ObjectWafProfileConstraintException-Version"); ok {
			if err = d.Set("version", vv); err != nil {
				return fmt.Errorf("Error reading version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	return nil
}

func flattenObjectWafProfileConstraintExceptionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWafProfileConstraintExceptionAddress3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWafProfileConstraintExceptionContentLength3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionHeaderLength3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionHostname3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionLineLength3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMalformed3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMaxCookie3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMaxHeaderLine3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMaxRangeSegment3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMaxUrlParam3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMethod3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionParamLength3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionPattern3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionRegex3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionUrlParamLength3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionVersion3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWafProfileConstraintException(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandObjectWafProfileConstraintExceptionAddress3rdl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("content_length"); ok || d.HasChange("content_length") {
		t, err := expandObjectWafProfileConstraintExceptionContentLength3rdl(d, v, "content_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-length"] = t
		}
	}

	if v, ok := d.GetOk("header_length"); ok || d.HasChange("header_length") {
		t, err := expandObjectWafProfileConstraintExceptionHeaderLength3rdl(d, v, "header_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-length"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok || d.HasChange("hostname") {
		t, err := expandObjectWafProfileConstraintExceptionHostname3rdl(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWafProfileConstraintExceptionId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("line_length"); ok || d.HasChange("line_length") {
		t, err := expandObjectWafProfileConstraintExceptionLineLength3rdl(d, v, "line_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["line-length"] = t
		}
	}

	if v, ok := d.GetOk("malformed"); ok || d.HasChange("malformed") {
		t, err := expandObjectWafProfileConstraintExceptionMalformed3rdl(d, v, "malformed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed"] = t
		}
	}

	if v, ok := d.GetOk("max_cookie"); ok || d.HasChange("max_cookie") {
		t, err := expandObjectWafProfileConstraintExceptionMaxCookie3rdl(d, v, "max_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-cookie"] = t
		}
	}

	if v, ok := d.GetOk("max_header_line"); ok || d.HasChange("max_header_line") {
		t, err := expandObjectWafProfileConstraintExceptionMaxHeaderLine3rdl(d, v, "max_header_line")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-header-line"] = t
		}
	}

	if v, ok := d.GetOk("max_range_segment"); ok || d.HasChange("max_range_segment") {
		t, err := expandObjectWafProfileConstraintExceptionMaxRangeSegment3rdl(d, v, "max_range_segment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-range-segment"] = t
		}
	}

	if v, ok := d.GetOk("max_url_param"); ok || d.HasChange("max_url_param") {
		t, err := expandObjectWafProfileConstraintExceptionMaxUrlParam3rdl(d, v, "max_url_param")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok || d.HasChange("method") {
		t, err := expandObjectWafProfileConstraintExceptionMethod3rdl(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("param_length"); ok || d.HasChange("param_length") {
		t, err := expandObjectWafProfileConstraintExceptionParamLength3rdl(d, v, "param_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["param-length"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandObjectWafProfileConstraintExceptionPattern3rdl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("regex"); ok || d.HasChange("regex") {
		t, err := expandObjectWafProfileConstraintExceptionRegex3rdl(d, v, "regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["regex"] = t
		}
	}

	if v, ok := d.GetOk("url_param_length"); ok || d.HasChange("url_param_length") {
		t, err := expandObjectWafProfileConstraintExceptionUrlParamLength3rdl(d, v, "url_param_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-length"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandObjectWafProfileConstraintExceptionVersion3rdl(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	return &obj, nil
}
