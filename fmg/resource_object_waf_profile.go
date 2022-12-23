// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Web application firewall configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWafProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWafProfileCreate,
		Read:   resourceObjectWafProfileRead,
		Update: resourceObjectWafProfileUpdate,
		Delete: resourceObjectWafProfileDelete,

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
			"address_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"blocked_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"blocked_log": &schema.Schema{
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
						"trusted_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"constraint": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
									},
									"header_length": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"hostname": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"line_length": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"malformed": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"max_cookie": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"max_header_line": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"max_range_segment": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"max_url_param": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"method": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"param_length": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"pattern": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"regex": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"url_param_length": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
					},
				},
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_allowed_methods": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"method_policy": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"allowed_methods": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
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
								},
							},
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"signature": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credit_card_detection_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"custom_signature": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"case_sensitivity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"direction": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"pattern": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
									"target": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"disabled_signature": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"disabled_sub_class": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"main_class": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"url_access": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_pattern": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
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
							},
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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

func resourceObjectWafProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectWafProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWafProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWafProfile(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWafProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWafProfileRead(d, m)
}

func resourceObjectWafProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWafProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWafProfile(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWafProfileRead(d, m)
}

func resourceObjectWafProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWafProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWafProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWafProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWafProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWafProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectWafProfileAddressList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "blocked_address"
	if _, ok := i["blocked-address"]; ok {
		result["blocked_address"] = flattenObjectWafProfileAddressListBlockedAddress(i["blocked-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "blocked_log"
	if _, ok := i["blocked-log"]; ok {
		result["blocked_log"] = flattenObjectWafProfileAddressListBlockedLog(i["blocked-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileAddressListSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileAddressListStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "trusted_address"
	if _, ok := i["trusted-address"]; ok {
		result["trusted_address"] = flattenObjectWafProfileAddressListTrustedAddress(i["trusted-address"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileAddressListBlockedAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWafProfileAddressListBlockedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileAddressListSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileAddressListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileAddressListTrustedAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWafProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraint(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "content_length"
	if _, ok := i["content-length"]; ok {
		result["content_length"] = flattenObjectWafProfileConstraintContentLength(i["content-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "exception"
	if _, ok := i["exception"]; ok {
		result["exception"] = flattenObjectWafProfileConstraintException(i["exception"], d, pre_append)
	}

	pre_append = pre + ".0." + "header_length"
	if _, ok := i["header-length"]; ok {
		result["header_length"] = flattenObjectWafProfileConstraintHeaderLength(i["header-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "hostname"
	if _, ok := i["hostname"]; ok {
		result["hostname"] = flattenObjectWafProfileConstraintHostname(i["hostname"], d, pre_append)
	}

	pre_append = pre + ".0." + "line_length"
	if _, ok := i["line-length"]; ok {
		result["line_length"] = flattenObjectWafProfileConstraintLineLength(i["line-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed"
	if _, ok := i["malformed"]; ok {
		result["malformed"] = flattenObjectWafProfileConstraintMalformed(i["malformed"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_cookie"
	if _, ok := i["max-cookie"]; ok {
		result["max_cookie"] = flattenObjectWafProfileConstraintMaxCookie(i["max-cookie"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_header_line"
	if _, ok := i["max-header-line"]; ok {
		result["max_header_line"] = flattenObjectWafProfileConstraintMaxHeaderLine(i["max-header-line"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := i["max-range-segment"]; ok {
		result["max_range_segment"] = flattenObjectWafProfileConstraintMaxRangeSegment(i["max-range-segment"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_url_param"
	if _, ok := i["max-url-param"]; ok {
		result["max_url_param"] = flattenObjectWafProfileConstraintMaxUrlParam(i["max-url-param"], d, pre_append)
	}

	pre_append = pre + ".0." + "method"
	if _, ok := i["method"]; ok {
		result["method"] = flattenObjectWafProfileConstraintMethod(i["method"], d, pre_append)
	}

	pre_append = pre + ".0." + "param_length"
	if _, ok := i["param-length"]; ok {
		result["param_length"] = flattenObjectWafProfileConstraintParamLength(i["param-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "url_param_length"
	if _, ok := i["url-param-length"]; ok {
		result["url_param_length"] = flattenObjectWafProfileConstraintUrlParamLength(i["url-param-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "version"
	if _, ok := i["version"]; ok {
		result["version"] = flattenObjectWafProfileConstraintVersion(i["version"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintContentLength(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintContentLengthAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenObjectWafProfileConstraintContentLengthLength(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintContentLengthLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintContentLengthSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintContentLengthStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintContentLengthAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintContentLengthLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintContentLengthLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintContentLengthSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintContentLengthStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintException(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWafProfileConstraintExceptionAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content_length"
		if _, ok := i["content-length"]; ok {
			v := flattenObjectWafProfileConstraintExceptionContentLength(i["content-length"], d, pre_append)
			tmp["content_length"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-ContentLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_length"
		if _, ok := i["header-length"]; ok {
			v := flattenObjectWafProfileConstraintExceptionHeaderLength(i["header-length"], d, pre_append)
			tmp["header_length"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-HeaderLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := i["hostname"]; ok {
			v := flattenObjectWafProfileConstraintExceptionHostname(i["hostname"], d, pre_append)
			tmp["hostname"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Hostname")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWafProfileConstraintExceptionId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "line_length"
		if _, ok := i["line-length"]; ok {
			v := flattenObjectWafProfileConstraintExceptionLineLength(i["line-length"], d, pre_append)
			tmp["line_length"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-LineLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "malformed"
		if _, ok := i["malformed"]; ok {
			v := flattenObjectWafProfileConstraintExceptionMalformed(i["malformed"], d, pre_append)
			tmp["malformed"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Malformed")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_cookie"
		if _, ok := i["max-cookie"]; ok {
			v := flattenObjectWafProfileConstraintExceptionMaxCookie(i["max-cookie"], d, pre_append)
			tmp["max_cookie"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-MaxCookie")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_header_line"
		if _, ok := i["max-header-line"]; ok {
			v := flattenObjectWafProfileConstraintExceptionMaxHeaderLine(i["max-header-line"], d, pre_append)
			tmp["max_header_line"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-MaxHeaderLine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_range_segment"
		if _, ok := i["max-range-segment"]; ok {
			v := flattenObjectWafProfileConstraintExceptionMaxRangeSegment(i["max-range-segment"], d, pre_append)
			tmp["max_range_segment"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-MaxRangeSegment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_url_param"
		if _, ok := i["max-url-param"]; ok {
			v := flattenObjectWafProfileConstraintExceptionMaxUrlParam(i["max-url-param"], d, pre_append)
			tmp["max_url_param"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-MaxUrlParam")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := i["method"]; ok {
			v := flattenObjectWafProfileConstraintExceptionMethod(i["method"], d, pre_append)
			tmp["method"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Method")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "param_length"
		if _, ok := i["param-length"]; ok {
			v := flattenObjectWafProfileConstraintExceptionParamLength(i["param-length"], d, pre_append)
			tmp["param_length"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-ParamLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenObjectWafProfileConstraintExceptionPattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			v := flattenObjectWafProfileConstraintExceptionRegex(i["regex"], d, pre_append)
			tmp["regex"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Regex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_param_length"
		if _, ok := i["url-param-length"]; ok {
			v := flattenObjectWafProfileConstraintExceptionUrlParamLength(i["url-param-length"], d, pre_append)
			tmp["url_param_length"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-UrlParamLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			v := flattenObjectWafProfileConstraintExceptionVersion(i["version"], d, pre_append)
			tmp["version"] = fortiAPISubPartPatch(v, "ObjectWafProfileConstraint-Exception-Version")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWafProfileConstraintExceptionAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionContentLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionHeaderLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionLineLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMalformed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMaxCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMaxHeaderLine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMaxRangeSegment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMaxUrlParam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionParamLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionUrlParamLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintExceptionVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHeaderLength(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintHeaderLengthAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenObjectWafProfileConstraintHeaderLengthLength(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintHeaderLengthLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintHeaderLengthSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintHeaderLengthStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintHeaderLengthAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHeaderLengthLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHeaderLengthLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHeaderLengthSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHeaderLengthStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHostname(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintHostnameAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintHostnameLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintHostnameSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintHostnameStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintHostnameAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHostnameLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHostnameSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintHostnameStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintLineLength(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintLineLengthAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenObjectWafProfileConstraintLineLengthLength(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintLineLengthLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintLineLengthSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintLineLengthStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintLineLengthAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintLineLengthLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintLineLengthLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintLineLengthSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintLineLengthStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMalformed(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintMalformedAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintMalformedLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintMalformedSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintMalformedStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintMalformedAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMalformedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMalformedSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMalformedStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxCookie(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintMaxCookieAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintMaxCookieLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_cookie"
	if _, ok := i["max-cookie"]; ok {
		result["max_cookie"] = flattenObjectWafProfileConstraintMaxCookieMaxCookie(i["max-cookie"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintMaxCookieSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintMaxCookieStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintMaxCookieAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxCookieLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxCookieMaxCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxCookieSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxCookieStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxHeaderLine(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintMaxHeaderLineAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintMaxHeaderLineLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_header_line"
	if _, ok := i["max-header-line"]; ok {
		result["max_header_line"] = flattenObjectWafProfileConstraintMaxHeaderLineMaxHeaderLine(i["max-header-line"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintMaxHeaderLineSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintMaxHeaderLineStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintMaxHeaderLineAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxHeaderLineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxHeaderLineMaxHeaderLine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxHeaderLineSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxHeaderLineStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxRangeSegment(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintMaxRangeSegmentAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintMaxRangeSegmentLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := i["max-range-segment"]; ok {
		result["max_range_segment"] = flattenObjectWafProfileConstraintMaxRangeSegmentMaxRangeSegment(i["max-range-segment"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintMaxRangeSegmentSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintMaxRangeSegmentStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintMaxRangeSegmentAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxRangeSegmentLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxRangeSegmentMaxRangeSegment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxRangeSegmentSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxRangeSegmentStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxUrlParam(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintMaxUrlParamAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintMaxUrlParamLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_url_param"
	if _, ok := i["max-url-param"]; ok {
		result["max_url_param"] = flattenObjectWafProfileConstraintMaxUrlParamMaxUrlParam(i["max-url-param"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintMaxUrlParamSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintMaxUrlParamStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintMaxUrlParamAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxUrlParamLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxUrlParamMaxUrlParam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxUrlParamSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMaxUrlParamStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMethod(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintMethodAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintMethodLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintMethodSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintMethodStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintMethodAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMethodLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMethodSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintMethodStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintParamLength(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintParamLengthAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenObjectWafProfileConstraintParamLengthLength(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintParamLengthLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintParamLengthSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintParamLengthStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintParamLengthAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintParamLengthLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintParamLengthLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintParamLengthSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintParamLengthStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintUrlParamLength(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintUrlParamLengthAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenObjectWafProfileConstraintUrlParamLengthLength(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintUrlParamLengthLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintUrlParamLengthSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintUrlParamLengthStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintUrlParamLengthAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintUrlParamLengthLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintUrlParamLengthLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintUrlParamLengthSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintUrlParamLengthStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintVersion(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileConstraintVersionAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileConstraintVersionLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileConstraintVersionSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileConstraintVersionStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileConstraintVersionAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintVersionLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintVersionSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileConstraintVersionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileMethod(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "default_allowed_methods"
	if _, ok := i["default-allowed-methods"]; ok {
		result["default_allowed_methods"] = flattenObjectWafProfileMethodDefaultAllowedMethods(i["default-allowed-methods"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileMethodLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "method_policy"
	if _, ok := i["method-policy"]; ok {
		result["method_policy"] = flattenObjectWafProfileMethodMethodPolicy(i["method-policy"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileMethodSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileMethodStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileMethodDefaultAllowedMethods(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWafProfileMethodLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileMethodMethodPolicy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWafProfileMethodMethodPolicyAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "ObjectWafProfileMethod-MethodPolicy-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_methods"
		if _, ok := i["allowed-methods"]; ok {
			v := flattenObjectWafProfileMethodMethodPolicyAllowedMethods(i["allowed-methods"], d, pre_append)
			tmp["allowed_methods"] = fortiAPISubPartPatch(v, "ObjectWafProfileMethod-MethodPolicy-AllowedMethods")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWafProfileMethodMethodPolicyId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWafProfileMethod-MethodPolicy-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenObjectWafProfileMethodMethodPolicyPattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "ObjectWafProfileMethod-MethodPolicy-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			v := flattenObjectWafProfileMethodMethodPolicyRegex(i["regex"], d, pre_append)
			tmp["regex"] = fortiAPISubPartPatch(v, "ObjectWafProfileMethod-MethodPolicy-Regex")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWafProfileMethodMethodPolicyAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileMethodMethodPolicyAllowedMethods(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWafProfileMethodMethodPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileMethodMethodPolicyPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileMethodMethodPolicyRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileMethodSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileMethodStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignature(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "credit_card_detection_threshold"
	if _, ok := i["credit-card-detection-threshold"]; ok {
		result["credit_card_detection_threshold"] = flattenObjectWafProfileSignatureCreditCardDetectionThreshold(i["credit-card-detection-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "custom_signature"
	if _, ok := i["custom-signature"]; ok {
		result["custom_signature"] = flattenObjectWafProfileSignatureCustomSignature(i["custom-signature"], d, pre_append)
	}

	pre_append = pre + ".0." + "disabled_signature"
	if _, ok := i["disabled-signature"]; ok {
		result["disabled_signature"] = flattenObjectWafProfileSignatureDisabledSignature(i["disabled-signature"], d, pre_append)
	}

	pre_append = pre + ".0." + "disabled_sub_class"
	if _, ok := i["disabled-sub-class"]; ok {
		result["disabled_sub_class"] = flattenObjectWafProfileSignatureDisabledSubClass(i["disabled-sub-class"], d, pre_append)
	}

	pre_append = pre + ".0." + "main_class"
	if _, ok := i["main-class"]; ok {
		result["main_class"] = flattenObjectWafProfileSignatureMainClass(i["main-class"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileSignatureCreditCardDetectionThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignature(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := i["case-sensitivity"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureCaseSensitivity(i["case-sensitivity"], d, pre_append)
			tmp["case_sensitivity"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-CaseSensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureDirection(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignaturePattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureSeverity(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Severity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			v := flattenObjectWafProfileSignatureCustomSignatureTarget(i["target"], d, pre_append)
			tmp["target"] = fortiAPISubPartPatch(v, "ObjectWafProfileSignature-CustomSignature-Target")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWafProfileSignatureCustomSignatureAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignatureCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignatureDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignatureLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignatureName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignaturePattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignatureSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignatureStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureCustomSignatureTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWafProfileSignatureDisabledSignature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWafProfileSignatureDisabledSubClass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWafProfileSignatureMainClass(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectWafProfileSignatureMainClassAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenObjectWafProfileSignatureMainClassId(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWafProfileSignatureMainClassLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenObjectWafProfileSignatureMainClassSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWafProfileSignatureMainClassStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWafProfileSignatureMainClassAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureMainClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureMainClassLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureMainClassSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileSignatureMainClassStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccess(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_pattern"
		if _, ok := i["access-pattern"]; ok {
			v := flattenObjectWafProfileUrlAccessAccessPattern(i["access-pattern"], d, pre_append)
			tmp["access_pattern"] = fortiAPISubPartPatch(v, "ObjectWafProfile-UrlAccess-AccessPattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenObjectWafProfileUrlAccessAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectWafProfile-UrlAccess-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenObjectWafProfileUrlAccessAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "ObjectWafProfile-UrlAccess-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWafProfileUrlAccessId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWafProfile-UrlAccess-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectWafProfileUrlAccessLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectWafProfile-UrlAccess-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenObjectWafProfileUrlAccessSeverity(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "ObjectWafProfile-UrlAccess-Severity")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWafProfileUrlAccessAccessPattern(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWafProfileUrlAccessAccessPatternId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWafProfileUrlAccess-AccessPattern-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := i["negate"]; ok {
			v := flattenObjectWafProfileUrlAccessAccessPatternNegate(i["negate"], d, pre_append)
			tmp["negate"] = fortiAPISubPartPatch(v, "ObjectWafProfileUrlAccess-AccessPattern-Negate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenObjectWafProfileUrlAccessAccessPatternPattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "ObjectWafProfileUrlAccess-AccessPattern-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			v := flattenObjectWafProfileUrlAccessAccessPatternRegex(i["regex"], d, pre_append)
			tmp["regex"] = fortiAPISubPartPatch(v, "ObjectWafProfileUrlAccess-AccessPattern-Regex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {
			v := flattenObjectWafProfileUrlAccessAccessPatternSrcaddr(i["srcaddr"], d, pre_append)
			tmp["srcaddr"] = fortiAPISubPartPatch(v, "ObjectWafProfileUrlAccess-AccessPattern-Srcaddr")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWafProfileUrlAccessAccessPatternId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAccessPatternNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAccessPatternPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAccessPatternRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAccessPatternSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileUrlAccessSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWafProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("address_list", flattenObjectWafProfileAddressList(o["address-list"], d, "address_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["address-list"], "ObjectWafProfile-AddressList"); ok {
				if err = d.Set("address_list", vv); err != nil {
					return fmt.Errorf("Error reading address_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading address_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("address_list"); ok {
			if err = d.Set("address_list", flattenObjectWafProfileAddressList(o["address-list"], d, "address_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["address-list"], "ObjectWafProfile-AddressList"); ok {
					if err = d.Set("address_list", vv); err != nil {
						return fmt.Errorf("Error reading address_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading address_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("comment", flattenObjectWafProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWafProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("constraint", flattenObjectWafProfileConstraint(o["constraint"], d, "constraint")); err != nil {
			if vv, ok := fortiAPIPatch(o["constraint"], "ObjectWafProfile-Constraint"); ok {
				if err = d.Set("constraint", vv); err != nil {
					return fmt.Errorf("Error reading constraint: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading constraint: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("constraint"); ok {
			if err = d.Set("constraint", flattenObjectWafProfileConstraint(o["constraint"], d, "constraint")); err != nil {
				if vv, ok := fortiAPIPatch(o["constraint"], "ObjectWafProfile-Constraint"); ok {
					if err = d.Set("constraint", vv); err != nil {
						return fmt.Errorf("Error reading constraint: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading constraint: %v", err)
				}
			}
		}
	}

	if err = d.Set("extended_log", flattenObjectWafProfileExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "ObjectWafProfile-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("external", flattenObjectWafProfileExternal(o["external"], d, "external")); err != nil {
		if vv, ok := fortiAPIPatch(o["external"], "ObjectWafProfile-External"); ok {
			if err = d.Set("external", vv); err != nil {
				return fmt.Errorf("Error reading external: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("method", flattenObjectWafProfileMethod(o["method"], d, "method")); err != nil {
			if vv, ok := fortiAPIPatch(o["method"], "ObjectWafProfile-Method"); ok {
				if err = d.Set("method", vv); err != nil {
					return fmt.Errorf("Error reading method: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading method: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("method"); ok {
			if err = d.Set("method", flattenObjectWafProfileMethod(o["method"], d, "method")); err != nil {
				if vv, ok := fortiAPIPatch(o["method"], "ObjectWafProfile-Method"); ok {
					if err = d.Set("method", vv); err != nil {
						return fmt.Errorf("Error reading method: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading method: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectWafProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWafProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("signature", flattenObjectWafProfileSignature(o["signature"], d, "signature")); err != nil {
			if vv, ok := fortiAPIPatch(o["signature"], "ObjectWafProfile-Signature"); ok {
				if err = d.Set("signature", vv); err != nil {
					return fmt.Errorf("Error reading signature: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading signature: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("signature"); ok {
			if err = d.Set("signature", flattenObjectWafProfileSignature(o["signature"], d, "signature")); err != nil {
				if vv, ok := fortiAPIPatch(o["signature"], "ObjectWafProfile-Signature"); ok {
					if err = d.Set("signature", vv); err != nil {
						return fmt.Errorf("Error reading signature: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading signature: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("url_access", flattenObjectWafProfileUrlAccess(o["url-access"], d, "url_access")); err != nil {
			if vv, ok := fortiAPIPatch(o["url-access"], "ObjectWafProfile-UrlAccess"); ok {
				if err = d.Set("url_access", vv); err != nil {
					return fmt.Errorf("Error reading url_access: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading url_access: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("url_access"); ok {
			if err = d.Set("url_access", flattenObjectWafProfileUrlAccess(o["url-access"], d, "url_access")); err != nil {
				if vv, ok := fortiAPIPatch(o["url-access"], "ObjectWafProfile-UrlAccess"); ok {
					if err = d.Set("url_access", vv); err != nil {
						return fmt.Errorf("Error reading url_access: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading url_access: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectWafProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWafProfileAddressList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "blocked_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["blocked-address"], _ = expandObjectWafProfileAddressListBlockedAddress(d, i["blocked_address"], pre_append)
	}
	pre_append = pre + ".0." + "blocked_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["blocked-log"], _ = expandObjectWafProfileAddressListBlockedLog(d, i["blocked_log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileAddressListSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileAddressListStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "trusted_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["trusted-address"], _ = expandObjectWafProfileAddressListTrustedAddress(d, i["trusted_address"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileAddressListBlockedAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileAddressListBlockedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileAddressListSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileAddressListStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileAddressListTrustedAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "content_length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileConstraintContentLength(d, i["content_length"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["content-length"] = t
		}
	}
	pre_append = pre + ".0." + "exception"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileConstraintException(d, i["exception"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["exception"] = t
		}
	}
	pre_append = pre + ".0." + "header_length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileConstraintHeaderLength(d, i["header_length"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["header-length"] = t
		}
	}
	pre_append = pre + ".0." + "hostname"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileConstraintHostname(d, i["hostname"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["hostname"] = t
		}
	}
	pre_append = pre + ".0." + "line_length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileConstraintLineLength(d, i["line_length"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["line-length"] = t
		}
	}
	pre_append = pre + ".0." + "malformed"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileConstraintMalformed(d, i["malformed"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["malformed"] = t
		}
	}
	pre_append = pre + ".0." + "max_cookie"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileConstraintMaxCookie(d, i["max_cookie"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["max-cookie"] = t
		}
	}
	pre_append = pre + ".0." + "max_header_line"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileConstraintMaxHeaderLine(d, i["max_header_line"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["max-header-line"] = t
		}
	}
	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileConstraintMaxRangeSegment(d, i["max_range_segment"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["max-range-segment"] = t
		}
	}
	pre_append = pre + ".0." + "max_url_param"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileConstraintMaxUrlParam(d, i["max_url_param"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["max-url-param"] = t
		}
	}
	pre_append = pre + ".0." + "method"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileConstraintMethod(d, i["method"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["method"] = t
		}
	}
	pre_append = pre + ".0." + "param_length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileConstraintParamLength(d, i["param_length"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["param-length"] = t
		}
	}
	pre_append = pre + ".0." + "url_param_length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileConstraintUrlParamLength(d, i["url_param_length"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["url-param-length"] = t
		}
	}
	pre_append = pre + ".0." + "version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileConstraintVersion(d, i["version"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["version"] = t
		}
	}

	return result, nil
}

func expandObjectWafProfileConstraintContentLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintContentLengthAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandObjectWafProfileConstraintContentLengthLength(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintContentLengthLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintContentLengthSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintContentLengthStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintContentLengthAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintContentLengthLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintContentLengthLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintContentLengthSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintContentLengthStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintException(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["address"], _ = expandObjectWafProfileConstraintExceptionAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["content-length"], _ = expandObjectWafProfileConstraintExceptionContentLength(d, i["content_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-length"], _ = expandObjectWafProfileConstraintExceptionHeaderLength(d, i["header_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hostname"], _ = expandObjectWafProfileConstraintExceptionHostname(d, i["hostname"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectWafProfileConstraintExceptionId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "line_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["line-length"], _ = expandObjectWafProfileConstraintExceptionLineLength(d, i["line_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "malformed"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["malformed"], _ = expandObjectWafProfileConstraintExceptionMalformed(d, i["malformed"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_cookie"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-cookie"], _ = expandObjectWafProfileConstraintExceptionMaxCookie(d, i["max_cookie"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_header_line"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-header-line"], _ = expandObjectWafProfileConstraintExceptionMaxHeaderLine(d, i["max_header_line"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_range_segment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-range-segment"], _ = expandObjectWafProfileConstraintExceptionMaxRangeSegment(d, i["max_range_segment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_url_param"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-url-param"], _ = expandObjectWafProfileConstraintExceptionMaxUrlParam(d, i["max_url_param"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["method"], _ = expandObjectWafProfileConstraintExceptionMethod(d, i["method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "param_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["param-length"], _ = expandObjectWafProfileConstraintExceptionParamLength(d, i["param_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandObjectWafProfileConstraintExceptionPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["regex"], _ = expandObjectWafProfileConstraintExceptionRegex(d, i["regex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_param_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-param-length"], _ = expandObjectWafProfileConstraintExceptionUrlParamLength(d, i["url_param_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["version"], _ = expandObjectWafProfileConstraintExceptionVersion(d, i["version"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWafProfileConstraintExceptionAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionContentLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionHeaderLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionLineLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMalformed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMaxCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMaxHeaderLine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMaxRangeSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMaxUrlParam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionParamLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionUrlParamLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintExceptionVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHeaderLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintHeaderLengthAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandObjectWafProfileConstraintHeaderLengthLength(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintHeaderLengthLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintHeaderLengthSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintHeaderLengthStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintHeaderLengthAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHeaderLengthLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHeaderLengthLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHeaderLengthSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHeaderLengthStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintHostnameAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintHostnameLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintHostnameSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintHostnameStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintHostnameAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHostnameLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHostnameSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintHostnameStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintLineLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintLineLengthAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandObjectWafProfileConstraintLineLengthLength(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintLineLengthLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintLineLengthSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintLineLengthStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintLineLengthAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintLineLengthLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintLineLengthLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintLineLengthSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintLineLengthStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMalformed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintMalformedAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintMalformedLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintMalformedSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintMalformedStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintMalformedAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMalformedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMalformedSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMalformedStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintMaxCookieAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintMaxCookieLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_cookie"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-cookie"], _ = expandObjectWafProfileConstraintMaxCookieMaxCookie(d, i["max_cookie"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintMaxCookieSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintMaxCookieStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintMaxCookieAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxCookieLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxCookieMaxCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxCookieSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxCookieStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxHeaderLine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintMaxHeaderLineAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintMaxHeaderLineLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_header_line"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-header-line"], _ = expandObjectWafProfileConstraintMaxHeaderLineMaxHeaderLine(d, i["max_header_line"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintMaxHeaderLineSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintMaxHeaderLineStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintMaxHeaderLineAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxHeaderLineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxHeaderLineMaxHeaderLine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxHeaderLineSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxHeaderLineStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxRangeSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintMaxRangeSegmentAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintMaxRangeSegmentLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-range-segment"], _ = expandObjectWafProfileConstraintMaxRangeSegmentMaxRangeSegment(d, i["max_range_segment"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintMaxRangeSegmentSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintMaxRangeSegmentStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintMaxRangeSegmentAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxRangeSegmentLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxRangeSegmentMaxRangeSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxRangeSegmentSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxRangeSegmentStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxUrlParam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintMaxUrlParamAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintMaxUrlParamLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_url_param"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-url-param"], _ = expandObjectWafProfileConstraintMaxUrlParamMaxUrlParam(d, i["max_url_param"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintMaxUrlParamSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintMaxUrlParamStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintMaxUrlParamAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxUrlParamLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxUrlParamMaxUrlParam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxUrlParamSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMaxUrlParamStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintMethodAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintMethodLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintMethodSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintMethodStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintMethodAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMethodLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMethodSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintMethodStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintParamLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintParamLengthAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandObjectWafProfileConstraintParamLengthLength(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintParamLengthLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintParamLengthSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintParamLengthStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintParamLengthAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintParamLengthLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintParamLengthLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintParamLengthSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintParamLengthStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintUrlParamLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintUrlParamLengthAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandObjectWafProfileConstraintUrlParamLengthLength(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintUrlParamLengthLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintUrlParamLengthSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintUrlParamLengthStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintUrlParamLengthAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintUrlParamLengthLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintUrlParamLengthLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintUrlParamLengthSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintUrlParamLengthStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileConstraintVersionAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileConstraintVersionLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileConstraintVersionSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileConstraintVersionStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileConstraintVersionAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintVersionLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintVersionSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileConstraintVersionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileExternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "default_allowed_methods"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-allowed-methods"], _ = expandObjectWafProfileMethodDefaultAllowedMethods(d, i["default_allowed_methods"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileMethodLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "method_policy"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileMethodMethodPolicy(d, i["method_policy"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["method-policy"] = t
		}
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileMethodSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileMethodStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileMethodDefaultAllowedMethods(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWafProfileMethodLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileMethodMethodPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["address"], _ = expandObjectWafProfileMethodMethodPolicyAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_methods"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowed-methods"], _ = expandObjectWafProfileMethodMethodPolicyAllowedMethods(d, i["allowed_methods"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectWafProfileMethodMethodPolicyId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandObjectWafProfileMethodMethodPolicyPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["regex"], _ = expandObjectWafProfileMethodMethodPolicyRegex(d, i["regex"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWafProfileMethodMethodPolicyAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileMethodMethodPolicyAllowedMethods(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWafProfileMethodMethodPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileMethodMethodPolicyPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileMethodMethodPolicyRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileMethodSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileMethodStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "credit_card_detection_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["credit-card-detection-threshold"], _ = expandObjectWafProfileSignatureCreditCardDetectionThreshold(d, i["credit_card_detection_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "custom_signature"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileSignatureCustomSignature(d, i["custom_signature"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["custom-signature"] = t
		}
	}
	pre_append = pre + ".0." + "disabled_signature"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disabled-signature"], _ = expandObjectWafProfileSignatureDisabledSignature(d, i["disabled_signature"], pre_append)
	}
	pre_append = pre + ".0." + "disabled_sub_class"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disabled-sub-class"], _ = expandObjectWafProfileSignatureDisabledSubClass(d, i["disabled_sub_class"], pre_append)
	}
	pre_append = pre + ".0." + "main_class"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectWafProfileSignatureMainClass(d, i["main_class"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["main-class"] = t
		}
	}

	return result, nil
}

func expandObjectWafProfileSignatureCreditCardDetectionThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectWafProfileSignatureCustomSignatureAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitivity"], _ = expandObjectWafProfileSignatureCustomSignatureCaseSensitivity(d, i["case_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandObjectWafProfileSignatureCustomSignatureDirection(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandObjectWafProfileSignatureCustomSignatureLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectWafProfileSignatureCustomSignatureName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandObjectWafProfileSignatureCustomSignaturePattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["severity"], _ = expandObjectWafProfileSignatureCustomSignatureSeverity(d, i["severity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectWafProfileSignatureCustomSignatureStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target"], _ = expandObjectWafProfileSignatureCustomSignatureTarget(d, i["target"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWafProfileSignatureCustomSignatureAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignatureCaseSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignatureDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignatureLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignatureName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignaturePattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignatureSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignatureStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureCustomSignatureTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWafProfileSignatureDisabledSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureDisabledSubClass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureMainClass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectWafProfileSignatureMainClassAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandObjectWafProfileSignatureMainClassId(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWafProfileSignatureMainClassLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandObjectWafProfileSignatureMainClassSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWafProfileSignatureMainClassStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWafProfileSignatureMainClassAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureMainClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureMainClassLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureMainClassSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileSignatureMainClassStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectWafProfileUrlAccessAccessPattern(d, i["access_pattern"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["access-pattern"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectWafProfileUrlAccessAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandObjectWafProfileUrlAccessAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectWafProfileUrlAccessId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandObjectWafProfileUrlAccessLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["severity"], _ = expandObjectWafProfileUrlAccessSeverity(d, i["severity"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWafProfileUrlAccessAccessPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectWafProfileUrlAccessAccessPatternId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["negate"], _ = expandObjectWafProfileUrlAccessAccessPatternNegate(d, i["negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandObjectWafProfileUrlAccessAccessPatternPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["regex"], _ = expandObjectWafProfileUrlAccessAccessPatternRegex(d, i["regex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr"], _ = expandObjectWafProfileUrlAccessAccessPatternSrcaddr(d, i["srcaddr"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWafProfileUrlAccessAccessPatternId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAccessPatternNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAccessPatternPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAccessPatternRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAccessPatternSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileUrlAccessSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWafProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address_list"); ok || d.HasChange("address_list") {
		t, err := expandObjectWafProfileAddressList(d, v, "address_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-list"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectWafProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("constraint"); ok || d.HasChange("constraint") {
		t, err := expandObjectWafProfileConstraint(d, v, "constraint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["constraint"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandObjectWafProfileExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("external"); ok || d.HasChange("external") {
		t, err := expandObjectWafProfileExternal(d, v, "external")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok || d.HasChange("method") {
		t, err := expandObjectWafProfileMethod(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWafProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("signature"); ok || d.HasChange("signature") {
		t, err := expandObjectWafProfileSignature(d, v, "signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature"] = t
		}
	}

	if v, ok := d.GetOk("url_access"); ok || d.HasChange("url_access") {
		t, err := expandObjectWafProfileUrlAccess(d, v, "url_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-access"] = t
		}
	}

	return &obj, nil
}
