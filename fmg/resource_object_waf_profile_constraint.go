// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: WAF HTTP protocol restrictions.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWafProfileConstraint() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWafProfileConstraintUpdate,
		Read:   resourceObjectWafProfileConstraintRead,
		Update: resourceObjectWafProfileConstraintUpdate,
		Delete: resourceObjectWafProfileConstraintDelete,

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
			"content_length": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"exception": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
				},
			},
			"header_length": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"line_length": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"malformed": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"max_cookie": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_cookie": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"max_header_line": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_header_line": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"max_range_segment": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_range_segment": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"max_url_param": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_url_param": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"method": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"param_length": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"url_param_length": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"version": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectWafProfileConstraintUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWafProfileConstraint(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfileConstraint resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWafProfileConstraint(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfileConstraint resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectWafProfileConstraint")

	return resourceObjectWafProfileConstraintRead(d, m)
}

func resourceObjectWafProfileConstraintDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWafProfileConstraint(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWafProfileConstraint resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWafProfileConstraintRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWafProfileConstraint(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfileConstraint resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWafProfileConstraint(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfileConstraint resource from API: %v", err)
	}
	return nil
}

func flattenObjectWafProfileConstraintContentLength2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintContentLengthAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenObjectWafProfileConstraintContentLengthLength2edl(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintContentLengthLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintContentLengthSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintContentLengthStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintContentLengthAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintContentLengthLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintContentLengthLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintContentLengthSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintContentLengthStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintException2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenObjectWafProfileConstraintExceptionAddress2edl(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content_length"
		if _, ok := i["content-length"]; ok {
			v := flattenObjectWafProfileConstraintExceptionContentLength2edl(i["content-length"], d, pre_append)
			tmp["content_length"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-ContentLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_length"
		if _, ok := i["header-length"]; ok {
			v := flattenObjectWafProfileConstraintExceptionHeaderLength2edl(i["header-length"], d, pre_append)
			tmp["header_length"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-HeaderLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := i["hostname"]; ok {
			v := flattenObjectWafProfileConstraintExceptionHostname2edl(i["hostname"], d, pre_append)
			tmp["hostname"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Hostname")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWafProfileConstraintExceptionId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "line_length"
		if _, ok := i["line-length"]; ok {
			v := flattenObjectWafProfileConstraintExceptionLineLength2edl(i["line-length"], d, pre_append)
			tmp["line_length"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-LineLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "malformed"
		if _, ok := i["malformed"]; ok {
			v := flattenObjectWafProfileConstraintExceptionMalformed2edl(i["malformed"], d, pre_append)
			tmp["malformed"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Malformed")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_cookie"
		if _, ok := i["max-cookie"]; ok {
			v := flattenObjectWafProfileConstraintExceptionMaxCookie2edl(i["max-cookie"], d, pre_append)
			tmp["max_cookie"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-MaxCookie")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_header_line"
		if _, ok := i["max-header-line"]; ok {
			v := flattenObjectWafProfileConstraintExceptionMaxHeaderLine2edl(i["max-header-line"], d, pre_append)
			tmp["max_header_line"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-MaxHeaderLine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_range_segment"
		if _, ok := i["max-range-segment"]; ok {
			v := flattenObjectWafProfileConstraintExceptionMaxRangeSegment2edl(i["max-range-segment"], d, pre_append)
			tmp["max_range_segment"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-MaxRangeSegment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_url_param"
		if _, ok := i["max-url-param"]; ok {
			v := flattenObjectWafProfileConstraintExceptionMaxUrlParam2edl(i["max-url-param"], d, pre_append)
			tmp["max_url_param"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-MaxUrlParam")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := i["method"]; ok {
			v := flattenObjectWafProfileConstraintExceptionMethod2edl(i["method"], d, pre_append)
			tmp["method"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Method")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "param_length"
		if _, ok := i["param-length"]; ok {
			v := flattenObjectWafProfileConstraintExceptionParamLength2edl(i["param-length"], d, pre_append)
			tmp["param_length"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-ParamLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenObjectWafProfileConstraintExceptionPattern2edl(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			v := flattenObjectWafProfileConstraintExceptionRegex2edl(i["regex"], d, pre_append)
			tmp["regex"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Regex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_param_length"
		if _, ok := i["url-param-length"]; ok {
			v := flattenObjectWafProfileConstraintExceptionUrlParamLength2edl(i["url-param-length"], d, pre_append)
			tmp["url_param_length"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-UrlParamLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			v := flattenObjectWafProfileConstraintExceptionVersion2edl(i["version"], d, pre_append)
			tmp["version"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Version")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectWafProfileConstraintExceptionAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWafProfileConstraintExceptionContentLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionHeaderLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionHostname2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionLineLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMalformed2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMaxCookie2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMaxHeaderLine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMaxRangeSegment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMaxUrlParam2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionParamLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionRegex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionUrlParamLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHeaderLength2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintHeaderLengthAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenObjectWafProfileConstraintHeaderLengthLength2edl(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintHeaderLengthLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintHeaderLengthSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintHeaderLengthStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintHeaderLengthAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHeaderLengthLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHeaderLengthLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHeaderLengthSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHeaderLengthStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHostname2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintHostnameAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintHostnameLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintHostnameSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintHostnameStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintHostnameAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHostnameLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHostnameSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHostnameStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintLineLength2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintLineLengthAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenObjectWafProfileConstraintLineLengthLength2edl(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintLineLengthLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintLineLengthSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintLineLengthStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintLineLengthAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintLineLengthLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintLineLengthLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintLineLengthSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintLineLengthStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMalformed2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintMalformedAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintMalformedLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintMalformedSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintMalformedStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintMalformedAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMalformedLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMalformedSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMalformedStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxCookie2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintMaxCookieAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintMaxCookieLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_cookie"
	if _, ok := i["max-cookie"]; ok {
		result["max_cookie"] = flattenObjectWafProfileConstraintMaxCookieMaxCookie2edl(i["max-cookie"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintMaxCookieSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintMaxCookieStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintMaxCookieAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxCookieLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxCookieMaxCookie2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxCookieSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxCookieStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxHeaderLine2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintMaxHeaderLineAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintMaxHeaderLineLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_header_line"
	if _, ok := i["max-header-line"]; ok {
		result["max_header_line"] = flattenObjectWafProfileConstraintMaxHeaderLineMaxHeaderLine2edl(i["max-header-line"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintMaxHeaderLineSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintMaxHeaderLineStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintMaxHeaderLineAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxHeaderLineLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxHeaderLineMaxHeaderLine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxHeaderLineSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxHeaderLineStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxRangeSegment2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintMaxRangeSegmentAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintMaxRangeSegmentLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := i["max-range-segment"]; ok {
		result["max_range_segment"] = flattenObjectWafProfileConstraintMaxRangeSegmentMaxRangeSegment2edl(i["max-range-segment"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintMaxRangeSegmentSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintMaxRangeSegmentStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintMaxRangeSegmentAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxRangeSegmentLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxRangeSegmentMaxRangeSegment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxRangeSegmentSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxRangeSegmentStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxUrlParam2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintMaxUrlParamAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintMaxUrlParamLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_url_param"
	if _, ok := i["max-url-param"]; ok {
		result["max_url_param"] = flattenObjectWafProfileConstraintMaxUrlParamMaxUrlParam2edl(i["max-url-param"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintMaxUrlParamSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintMaxUrlParamStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintMaxUrlParamAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxUrlParamLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxUrlParamMaxUrlParam2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxUrlParamSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxUrlParamStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMethod2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintMethodAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintMethodLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintMethodSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintMethodStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintMethodAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMethodLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMethodSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMethodStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintParamLength2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintParamLengthAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenObjectWafProfileConstraintParamLengthLength2edl(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintParamLengthLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintParamLengthSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintParamLengthStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintParamLengthAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintParamLengthLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintParamLengthLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintParamLengthSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintParamLengthStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintUrlParamLength2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintUrlParamLengthAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenObjectWafProfileConstraintUrlParamLengthLength2edl(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintUrlParamLengthLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintUrlParamLengthSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintUrlParamLengthStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintUrlParamLengthAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintUrlParamLengthLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintUrlParamLengthLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintUrlParamLengthSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintUrlParamLengthStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintVersion2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintVersionAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintVersionLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintVersionSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintVersionStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintVersionAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintVersionLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintVersionSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintVersionStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWafProfileConstraint(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("content_length", flattenObjectWafProfileConstraintContentLength2edl(o["content-length"], d, "content_length")); err != nil {
			if vv, ok := fortiAPIPatch(o["content-length"], "ObjectWafProfileConstraint-ContentLength"); ok {
				if err = d.Set("content_length", vv); err != nil {
					return fmt.Errorf("Error reading content_length: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading content_length: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("content_length"); ok {
			if err = d.Set("content_length", flattenObjectWafProfileConstraintContentLength2edl(o["content-length"], d, "content_length")); err != nil {
				if vv, ok := fortiAPIPatch(o["content-length"], "ObjectWafProfileConstraint-ContentLength"); ok {
					if err = d.Set("content_length", vv); err != nil {
						return fmt.Errorf("Error reading content_length: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading content_length: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("exception", flattenObjectWafProfileConstraintException2edl(o["exception"], d, "exception")); err != nil {
			if vv, ok := fortiAPIPatch(o["exception"], "ObjectWafProfileConstraint-Exception"); ok {
				if err = d.Set("exception", vv); err != nil {
					return fmt.Errorf("Error reading exception: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading exception: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("exception"); ok {
			if err = d.Set("exception", flattenObjectWafProfileConstraintException2edl(o["exception"], d, "exception")); err != nil {
				if vv, ok := fortiAPIPatch(o["exception"], "ObjectWafProfileConstraint-Exception"); ok {
					if err = d.Set("exception", vv); err != nil {
						return fmt.Errorf("Error reading exception: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading exception: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("header_length", flattenObjectWafProfileConstraintHeaderLength2edl(o["header-length"], d, "header_length")); err != nil {
			if vv, ok := fortiAPIPatch(o["header-length"], "ObjectWafProfileConstraint-HeaderLength"); ok {
				if err = d.Set("header_length", vv); err != nil {
					return fmt.Errorf("Error reading header_length: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading header_length: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("header_length"); ok {
			if err = d.Set("header_length", flattenObjectWafProfileConstraintHeaderLength2edl(o["header-length"], d, "header_length")); err != nil {
				if vv, ok := fortiAPIPatch(o["header-length"], "ObjectWafProfileConstraint-HeaderLength"); ok {
					if err = d.Set("header_length", vv); err != nil {
						return fmt.Errorf("Error reading header_length: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading header_length: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("hostname", flattenObjectWafProfileConstraintHostname2edl(o["hostname"], d, "hostname")); err != nil {
			if vv, ok := fortiAPIPatch(o["hostname"], "ObjectWafProfileConstraint-Hostname"); ok {
				if err = d.Set("hostname", vv); err != nil {
					return fmt.Errorf("Error reading hostname: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("hostname"); ok {
			if err = d.Set("hostname", flattenObjectWafProfileConstraintHostname2edl(o["hostname"], d, "hostname")); err != nil {
				if vv, ok := fortiAPIPatch(o["hostname"], "ObjectWafProfileConstraint-Hostname"); ok {
					if err = d.Set("hostname", vv); err != nil {
						return fmt.Errorf("Error reading hostname: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading hostname: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("line_length", flattenObjectWafProfileConstraintLineLength2edl(o["line-length"], d, "line_length")); err != nil {
			if vv, ok := fortiAPIPatch(o["line-length"], "ObjectWafProfileConstraint-LineLength"); ok {
				if err = d.Set("line_length", vv); err != nil {
					return fmt.Errorf("Error reading line_length: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading line_length: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("line_length"); ok {
			if err = d.Set("line_length", flattenObjectWafProfileConstraintLineLength2edl(o["line-length"], d, "line_length")); err != nil {
				if vv, ok := fortiAPIPatch(o["line-length"], "ObjectWafProfileConstraint-LineLength"); ok {
					if err = d.Set("line_length", vv); err != nil {
						return fmt.Errorf("Error reading line_length: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading line_length: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("malformed", flattenObjectWafProfileConstraintMalformed2edl(o["malformed"], d, "malformed")); err != nil {
			if vv, ok := fortiAPIPatch(o["malformed"], "ObjectWafProfileConstraint-Malformed"); ok {
				if err = d.Set("malformed", vv); err != nil {
					return fmt.Errorf("Error reading malformed: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading malformed: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("malformed"); ok {
			if err = d.Set("malformed", flattenObjectWafProfileConstraintMalformed2edl(o["malformed"], d, "malformed")); err != nil {
				if vv, ok := fortiAPIPatch(o["malformed"], "ObjectWafProfileConstraint-Malformed"); ok {
					if err = d.Set("malformed", vv); err != nil {
						return fmt.Errorf("Error reading malformed: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading malformed: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("max_cookie", flattenObjectWafProfileConstraintMaxCookie2edl(o["max-cookie"], d, "max_cookie")); err != nil {
			if vv, ok := fortiAPIPatch(o["max-cookie"], "ObjectWafProfileConstraint-MaxCookie"); ok {
				if err = d.Set("max_cookie", vv); err != nil {
					return fmt.Errorf("Error reading max_cookie: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading max_cookie: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("max_cookie"); ok {
			if err = d.Set("max_cookie", flattenObjectWafProfileConstraintMaxCookie2edl(o["max-cookie"], d, "max_cookie")); err != nil {
				if vv, ok := fortiAPIPatch(o["max-cookie"], "ObjectWafProfileConstraint-MaxCookie"); ok {
					if err = d.Set("max_cookie", vv); err != nil {
						return fmt.Errorf("Error reading max_cookie: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading max_cookie: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("max_header_line", flattenObjectWafProfileConstraintMaxHeaderLine2edl(o["max-header-line"], d, "max_header_line")); err != nil {
			if vv, ok := fortiAPIPatch(o["max-header-line"], "ObjectWafProfileConstraint-MaxHeaderLine"); ok {
				if err = d.Set("max_header_line", vv); err != nil {
					return fmt.Errorf("Error reading max_header_line: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading max_header_line: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("max_header_line"); ok {
			if err = d.Set("max_header_line", flattenObjectWafProfileConstraintMaxHeaderLine2edl(o["max-header-line"], d, "max_header_line")); err != nil {
				if vv, ok := fortiAPIPatch(o["max-header-line"], "ObjectWafProfileConstraint-MaxHeaderLine"); ok {
					if err = d.Set("max_header_line", vv); err != nil {
						return fmt.Errorf("Error reading max_header_line: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading max_header_line: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("max_range_segment", flattenObjectWafProfileConstraintMaxRangeSegment2edl(o["max-range-segment"], d, "max_range_segment")); err != nil {
			if vv, ok := fortiAPIPatch(o["max-range-segment"], "ObjectWafProfileConstraint-MaxRangeSegment"); ok {
				if err = d.Set("max_range_segment", vv); err != nil {
					return fmt.Errorf("Error reading max_range_segment: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading max_range_segment: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("max_range_segment"); ok {
			if err = d.Set("max_range_segment", flattenObjectWafProfileConstraintMaxRangeSegment2edl(o["max-range-segment"], d, "max_range_segment")); err != nil {
				if vv, ok := fortiAPIPatch(o["max-range-segment"], "ObjectWafProfileConstraint-MaxRangeSegment"); ok {
					if err = d.Set("max_range_segment", vv); err != nil {
						return fmt.Errorf("Error reading max_range_segment: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading max_range_segment: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("max_url_param", flattenObjectWafProfileConstraintMaxUrlParam2edl(o["max-url-param"], d, "max_url_param")); err != nil {
			if vv, ok := fortiAPIPatch(o["max-url-param"], "ObjectWafProfileConstraint-MaxUrlParam"); ok {
				if err = d.Set("max_url_param", vv); err != nil {
					return fmt.Errorf("Error reading max_url_param: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading max_url_param: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("max_url_param"); ok {
			if err = d.Set("max_url_param", flattenObjectWafProfileConstraintMaxUrlParam2edl(o["max-url-param"], d, "max_url_param")); err != nil {
				if vv, ok := fortiAPIPatch(o["max-url-param"], "ObjectWafProfileConstraint-MaxUrlParam"); ok {
					if err = d.Set("max_url_param", vv); err != nil {
						return fmt.Errorf("Error reading max_url_param: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading max_url_param: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("method", flattenObjectWafProfileConstraintMethod2edl(o["method"], d, "method")); err != nil {
			if vv, ok := fortiAPIPatch(o["method"], "ObjectWafProfileConstraint-Method"); ok {
				if err = d.Set("method", vv); err != nil {
					return fmt.Errorf("Error reading method: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading method: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("method"); ok {
			if err = d.Set("method", flattenObjectWafProfileConstraintMethod2edl(o["method"], d, "method")); err != nil {
				if vv, ok := fortiAPIPatch(o["method"], "ObjectWafProfileConstraint-Method"); ok {
					if err = d.Set("method", vv); err != nil {
						return fmt.Errorf("Error reading method: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading method: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("param_length", flattenObjectWafProfileConstraintParamLength2edl(o["param-length"], d, "param_length")); err != nil {
			if vv, ok := fortiAPIPatch(o["param-length"], "ObjectWafProfileConstraint-ParamLength"); ok {
				if err = d.Set("param_length", vv); err != nil {
					return fmt.Errorf("Error reading param_length: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading param_length: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("param_length"); ok {
			if err = d.Set("param_length", flattenObjectWafProfileConstraintParamLength2edl(o["param-length"], d, "param_length")); err != nil {
				if vv, ok := fortiAPIPatch(o["param-length"], "ObjectWafProfileConstraint-ParamLength"); ok {
					if err = d.Set("param_length", vv); err != nil {
						return fmt.Errorf("Error reading param_length: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading param_length: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("url_param_length", flattenObjectWafProfileConstraintUrlParamLength2edl(o["url-param-length"], d, "url_param_length")); err != nil {
			if vv, ok := fortiAPIPatch(o["url-param-length"], "ObjectWafProfileConstraint-UrlParamLength"); ok {
				if err = d.Set("url_param_length", vv); err != nil {
					return fmt.Errorf("Error reading url_param_length: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading url_param_length: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("url_param_length"); ok {
			if err = d.Set("url_param_length", flattenObjectWafProfileConstraintUrlParamLength2edl(o["url-param-length"], d, "url_param_length")); err != nil {
				if vv, ok := fortiAPIPatch(o["url-param-length"], "ObjectWafProfileConstraint-UrlParamLength"); ok {
					if err = d.Set("url_param_length", vv); err != nil {
						return fmt.Errorf("Error reading url_param_length: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading url_param_length: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("version", flattenObjectWafProfileConstraintVersion2edl(o["version"], d, "version")); err != nil {
			if vv, ok := fortiAPIPatch(o["version"], "ObjectWafProfileConstraint-Version"); ok {
				if err = d.Set("version", vv); err != nil {
					return fmt.Errorf("Error reading version: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading version: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("version"); ok {
			if err = d.Set("version", flattenObjectWafProfileConstraintVersion2edl(o["version"], d, "version")); err != nil {
				if vv, ok := fortiAPIPatch(o["version"], "ObjectWafProfileConstraint-Version"); ok {
					if err = d.Set("version", vv); err != nil {
						return fmt.Errorf("Error reading version: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading version: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectWafProfileConstraintFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWafProfileConstraintContentLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintContentLengthAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandObjectWafProfileConstraintContentLengthLength2edl(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintContentLengthLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintContentLengthSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintContentLengthStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintContentLengthAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintContentLengthLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintContentLengthLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintContentLengthSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintContentLengthStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintException2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandObjectWafProfileConstraintExceptionAddress2edl(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["content-length"], _ = expandObjectWafProfileConstraintExceptionContentLength2edl(d, i["content_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-length"], _ = expandObjectWafProfileConstraintExceptionHeaderLength2edl(d, i["header_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hostname"], _ = expandObjectWafProfileConstraintExceptionHostname2edl(d, i["hostname"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectWafProfileConstraintExceptionId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "line_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["line-length"], _ = expandObjectWafProfileConstraintExceptionLineLength2edl(d, i["line_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "malformed"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["malformed"], _ = expandObjectWafProfileConstraintExceptionMalformed2edl(d, i["malformed"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_cookie"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-cookie"], _ = expandObjectWafProfileConstraintExceptionMaxCookie2edl(d, i["max_cookie"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_header_line"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-header-line"], _ = expandObjectWafProfileConstraintExceptionMaxHeaderLine2edl(d, i["max_header_line"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_range_segment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-range-segment"], _ = expandObjectWafProfileConstraintExceptionMaxRangeSegment2edl(d, i["max_range_segment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_url_param"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-url-param"], _ = expandObjectWafProfileConstraintExceptionMaxUrlParam2edl(d, i["max_url_param"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["method"], _ = expandObjectWafProfileConstraintExceptionMethod2edl(d, i["method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "param_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["param-length"], _ = expandObjectWafProfileConstraintExceptionParamLength2edl(d, i["param_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandObjectWafProfileConstraintExceptionPattern2edl(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["regex"], _ = expandObjectWafProfileConstraintExceptionRegex2edl(d, i["regex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_param_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-param-length"], _ = expandObjectWafProfileConstraintExceptionUrlParamLength2edl(d, i["url_param_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["version"], _ = expandObjectWafProfileConstraintExceptionVersion2edl(d, i["version"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectWafProfileConstraintExceptionAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWafProfileConstraintExceptionContentLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionHeaderLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionHostname2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionLineLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMalformed2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMaxCookie2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMaxHeaderLine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMaxRangeSegment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMaxUrlParam2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionParamLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionRegex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionUrlParamLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHeaderLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintHeaderLengthAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandObjectWafProfileConstraintHeaderLengthLength2edl(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintHeaderLengthLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintHeaderLengthSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintHeaderLengthStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintHeaderLengthAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHeaderLengthLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHeaderLengthLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHeaderLengthSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHeaderLengthStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHostname2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintHostnameAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintHostnameLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintHostnameSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintHostnameStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintHostnameAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHostnameLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHostnameSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHostnameStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintLineLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintLineLengthAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandObjectWafProfileConstraintLineLengthLength2edl(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintLineLengthLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintLineLengthSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintLineLengthStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintLineLengthAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintLineLengthLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintLineLengthLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintLineLengthSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintLineLengthStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMalformed2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintMalformedAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintMalformedLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintMalformedSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintMalformedStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintMalformedAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMalformedLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMalformedSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMalformedStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxCookie2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintMaxCookieAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintMaxCookieLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_cookie"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-cookie"], _ = expandObjectWafProfileConstraintMaxCookieMaxCookie2edl(d, i["max_cookie"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintMaxCookieSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintMaxCookieStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintMaxCookieAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxCookieLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxCookieMaxCookie2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxCookieSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxCookieStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxHeaderLine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintMaxHeaderLineAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintMaxHeaderLineLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_header_line"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-header-line"], _ = expandObjectWafProfileConstraintMaxHeaderLineMaxHeaderLine2edl(d, i["max_header_line"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintMaxHeaderLineSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintMaxHeaderLineStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintMaxHeaderLineAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxHeaderLineLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxHeaderLineMaxHeaderLine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxHeaderLineSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxHeaderLineStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxRangeSegment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintMaxRangeSegmentAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintMaxRangeSegmentLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-range-segment"], _ = expandObjectWafProfileConstraintMaxRangeSegmentMaxRangeSegment2edl(d, i["max_range_segment"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintMaxRangeSegmentSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintMaxRangeSegmentStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintMaxRangeSegmentAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxRangeSegmentLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxRangeSegmentMaxRangeSegment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxRangeSegmentSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxRangeSegmentStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxUrlParam2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintMaxUrlParamAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintMaxUrlParamLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_url_param"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-url-param"], _ = expandObjectWafProfileConstraintMaxUrlParamMaxUrlParam2edl(d, i["max_url_param"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintMaxUrlParamSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintMaxUrlParamStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintMaxUrlParamAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxUrlParamLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxUrlParamMaxUrlParam2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxUrlParamSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxUrlParamStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintMethodAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintMethodLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintMethodSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintMethodStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintMethodAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMethodLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMethodSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMethodStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintParamLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintParamLengthAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandObjectWafProfileConstraintParamLengthLength2edl(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintParamLengthLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintParamLengthSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintParamLengthStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintParamLengthAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintParamLengthLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintParamLengthLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintParamLengthSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintParamLengthStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintUrlParamLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintUrlParamLengthAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandObjectWafProfileConstraintUrlParamLengthLength2edl(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintUrlParamLengthLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintUrlParamLengthSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintUrlParamLengthStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintUrlParamLengthAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintUrlParamLengthLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintUrlParamLengthLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintUrlParamLengthSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintUrlParamLengthStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintVersionAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintVersionLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintVersionSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintVersionStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintVersionAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintVersionLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintVersionSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintVersionStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWafProfileConstraint(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("content_length"); ok || d.HasChange("content_length") {
		t, err := expandObjectWafProfileConstraintContentLength2edl(d, v, "content_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-length"] = t
		}
	}

	if v, ok := d.GetOk("exception"); ok || d.HasChange("exception") {
		t, err := expandObjectWafProfileConstraintException2edl(d, v, "exception")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exception"] = t
		}
	}

	if v, ok := d.GetOk("header_length"); ok || d.HasChange("header_length") {
		t, err := expandObjectWafProfileConstraintHeaderLength2edl(d, v, "header_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-length"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok || d.HasChange("hostname") {
		t, err := expandObjectWafProfileConstraintHostname2edl(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("line_length"); ok || d.HasChange("line_length") {
		t, err := expandObjectWafProfileConstraintLineLength2edl(d, v, "line_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["line-length"] = t
		}
	}

	if v, ok := d.GetOk("malformed"); ok || d.HasChange("malformed") {
		t, err := expandObjectWafProfileConstraintMalformed2edl(d, v, "malformed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed"] = t
		}
	}

	if v, ok := d.GetOk("max_cookie"); ok || d.HasChange("max_cookie") {
		t, err := expandObjectWafProfileConstraintMaxCookie2edl(d, v, "max_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-cookie"] = t
		}
	}

	if v, ok := d.GetOk("max_header_line"); ok || d.HasChange("max_header_line") {
		t, err := expandObjectWafProfileConstraintMaxHeaderLine2edl(d, v, "max_header_line")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-header-line"] = t
		}
	}

	if v, ok := d.GetOk("max_range_segment"); ok || d.HasChange("max_range_segment") {
		t, err := expandObjectWafProfileConstraintMaxRangeSegment2edl(d, v, "max_range_segment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-range-segment"] = t
		}
	}

	if v, ok := d.GetOk("max_url_param"); ok || d.HasChange("max_url_param") {
		t, err := expandObjectWafProfileConstraintMaxUrlParam2edl(d, v, "max_url_param")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok || d.HasChange("method") {
		t, err := expandObjectWafProfileConstraintMethod2edl(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("param_length"); ok || d.HasChange("param_length") {
		t, err := expandObjectWafProfileConstraintParamLength2edl(d, v, "param_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["param-length"] = t
		}
	}

	if v, ok := d.GetOk("url_param_length"); ok || d.HasChange("url_param_length") {
		t, err := expandObjectWafProfileConstraintUrlParamLength2edl(d, v, "url_param_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-length"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandObjectWafProfileConstraintVersion2edl(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	return &obj, nil
}
