// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Web filter profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWebfilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterProfileCreate,
		Read:   resourceObjectWebfilterProfileRead,
		Update: resourceObjectWebfilterProfileUpdate,
		Delete: resourceObjectWebfilterProfileDelete,

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
			"antiphish": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"check_basic_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"check_uri": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"check_username_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"custom_patterns": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"category": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"pattern": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"default_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"domain_controller": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"inspection_entries": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"ldap": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_body_len": &schema.Schema{
							Type:     schema.TypeInt,
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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entries": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"comment": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"direction": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"encryption": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"file_type": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"filter": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"password_protected": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scan_archive_contents": &schema.Schema{
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
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftgd_wf": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exempt_quota": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"auth_usr_grp": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"category": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"override_replacemsg": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"warn_duration": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"warning_duration_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"warning_prompt": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"max_quota_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ovrd": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"quota": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"category": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"duration": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"override_replacemsg": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"unit": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"rate_crl_urls": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_css_urls": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_image_urls": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_javascript_urls": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"https_replacemsg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inspection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_all_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"profile": &schema.Schema{
							Type:     schema.TypeList,
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
				},
			},
			"ovrd_perm": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"post_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_extraction": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"redirect_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"redirect_no_content": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"redirect_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_fqdn": &schema.Schema{
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
			"web": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowlist": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"blacklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bword_table": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bword_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"content_header_list": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keyword_match": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"log_search": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"safe_search": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"urlfilter_table": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vimeo_restrict": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"whitelist": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"youtube_restrict": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"web_antiphishing_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_content_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_extended_all_action_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_activex_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_applet_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_command_block_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_cookie_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_cookie_removal_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_js_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_jscript_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_referer_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_unknown_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_vbs_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_ftgd_err_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_ftgd_quota_usage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_invalid_domain_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_url_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wisp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wisp_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wisp_servers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"youtube_channel_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"channel_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"youtube_channel_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectWebfilterProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWebfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWebfilterProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebfilterProfileRead(d, m)
}

func resourceObjectWebfilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWebfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWebfilterProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebfilterProfileRead(d, m)
}

func resourceObjectWebfilterProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWebfilterProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWebfilterProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterProfileAntiphish(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "authentication"
	if _, ok := i["authentication"]; ok {
		result["authentication"] = flattenObjectWebfilterProfileAntiphishAuthentication(i["authentication"], d, pre_append)
	}

	pre_append = pre + ".0." + "check_basic_auth"
	if _, ok := i["check-basic-auth"]; ok {
		result["check_basic_auth"] = flattenObjectWebfilterProfileAntiphishCheckBasicAuth(i["check-basic-auth"], d, pre_append)
	}

	pre_append = pre + ".0." + "check_uri"
	if _, ok := i["check-uri"]; ok {
		result["check_uri"] = flattenObjectWebfilterProfileAntiphishCheckUri(i["check-uri"], d, pre_append)
	}

	pre_append = pre + ".0." + "check_username_only"
	if _, ok := i["check-username-only"]; ok {
		result["check_username_only"] = flattenObjectWebfilterProfileAntiphishCheckUsernameOnly(i["check-username-only"], d, pre_append)
	}

	pre_append = pre + ".0." + "custom_patterns"
	if _, ok := i["custom-patterns"]; ok {
		result["custom_patterns"] = flattenObjectWebfilterProfileAntiphishCustomPatterns(i["custom-patterns"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_action"
	if _, ok := i["default-action"]; ok {
		result["default_action"] = flattenObjectWebfilterProfileAntiphishDefaultAction(i["default-action"], d, pre_append)
	}

	pre_append = pre + ".0." + "domain_controller"
	if _, ok := i["domain-controller"]; ok {
		result["domain_controller"] = flattenObjectWebfilterProfileAntiphishDomainController(i["domain-controller"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspection_entries"
	if _, ok := i["inspection-entries"]; ok {
		result["inspection_entries"] = flattenObjectWebfilterProfileAntiphishInspectionEntries(i["inspection-entries"], d, pre_append)
	}

	pre_append = pre + ".0." + "ldap"
	if _, ok := i["ldap"]; ok {
		result["ldap"] = flattenObjectWebfilterProfileAntiphishLdap(i["ldap"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_body_len"
	if _, ok := i["max-body-len"]; ok {
		result["max_body_len"] = flattenObjectWebfilterProfileAntiphishMaxBodyLen(i["max-body-len"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWebfilterProfileAntiphishStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWebfilterProfileAntiphishAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishCheckBasicAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishCheckUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishCheckUsernameOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishCustomPatterns(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenObjectWebfilterProfileAntiphishCustomPatternsCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileAntiphish-CustomPatterns-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenObjectWebfilterProfileAntiphishCustomPatternsPattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileAntiphish-CustomPatterns-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectWebfilterProfileAntiphishCustomPatternsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileAntiphish-CustomPatterns-Type")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWebfilterProfileAntiphishCustomPatternsCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishCustomPatternsPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishCustomPatternsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishDomainController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishInspectionEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWebfilterProfileAntiphishInspectionEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileAntiphish-InspectionEntries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if _, ok := i["fortiguard-category"]; ok {
			v := flattenObjectWebfilterProfileAntiphishInspectionEntriesFortiguardCategory(i["fortiguard-category"], d, pre_append)
			tmp["fortiguard_category"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileAntiphish-InspectionEntries-FortiguardCategory")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectWebfilterProfileAntiphishInspectionEntriesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileAntiphish-InspectionEntries-Name")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWebfilterProfileAntiphishInspectionEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishInspectionEntriesFortiguardCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileAntiphishInspectionEntriesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishLdap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishMaxBodyLen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileAntiphishStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFileFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "entries"
	if _, ok := i["entries"]; ok {
		result["entries"] = flattenObjectWebfilterProfileFileFilterEntries(i["entries"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectWebfilterProfileFileFilterLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_archive_contents"
	if _, ok := i["scan-archive-contents"]; ok {
		result["scan_archive_contents"] = flattenObjectWebfilterProfileFileFilterScanArchiveContents(i["scan-archive-contents"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWebfilterProfileFileFilterStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWebfilterProfileFileFilterEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWebfilterProfileFileFilterEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFileFilter-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenObjectWebfilterProfileFileFilterEntriesComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFileFilter-Entries-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenObjectWebfilterProfileFileFilterEntriesDirection(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFileFilter-Entries-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encryption"
		if _, ok := i["encryption"]; ok {
			v := flattenObjectWebfilterProfileFileFilterEntriesEncryption(i["encryption"], d, pre_append)
			tmp["encryption"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFileFilter-Entries-Encryption")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := i["file-type"]; ok {
			v := flattenObjectWebfilterProfileFileFilterEntriesFileType(i["file-type"], d, pre_append)
			tmp["file_type"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFileFilter-Entries-FileType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {
			v := flattenObjectWebfilterProfileFileFilterEntriesFilter(i["filter"], d, pre_append)
			tmp["filter"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFileFilter-Entries-Filter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_protected"
		if _, ok := i["password-protected"]; ok {
			v := flattenObjectWebfilterProfileFileFilterEntriesPasswordProtected(i["password-protected"], d, pre_append)
			tmp["password_protected"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFileFilter-Entries-PasswordProtected")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectWebfilterProfileFileFilterEntriesProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFileFilter-Entries-Protocol")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWebfilterProfileFileFilterEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFileFilterEntriesComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFileFilterEntriesDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFileFilterEntriesEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFileFilterEntriesFileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileFileFilterEntriesFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFileFilterEntriesPasswordProtected(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFileFilterEntriesProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileFileFilterLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFileFilterScanArchiveContents(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFileFilterStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFeatureSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "exempt_quota"
	if _, ok := i["exempt-quota"]; ok {
		result["exempt_quota"] = flattenObjectWebfilterProfileFtgdWfExemptQuota(i["exempt-quota"], d, pre_append)
	}

	pre_append = pre + ".0." + "filters"
	if _, ok := i["filters"]; ok {
		result["filters"] = flattenObjectWebfilterProfileFtgdWfFilters(i["filters"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_quota_timeout"
	if _, ok := i["max-quota-timeout"]; ok {
		result["max_quota_timeout"] = flattenObjectWebfilterProfileFtgdWfMaxQuotaTimeout(i["max-quota-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectWebfilterProfileFtgdWfOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "ovrd"
	if _, ok := i["ovrd"]; ok {
		result["ovrd"] = flattenObjectWebfilterProfileFtgdWfOvrd(i["ovrd"], d, pre_append)
	}

	pre_append = pre + ".0." + "quota"
	if _, ok := i["quota"]; ok {
		result["quota"] = flattenObjectWebfilterProfileFtgdWfQuota(i["quota"], d, pre_append)
	}

	pre_append = pre + ".0." + "rate_crl_urls"
	if _, ok := i["rate-crl-urls"]; ok {
		result["rate_crl_urls"] = flattenObjectWebfilterProfileFtgdWfRateCrlUrls(i["rate-crl-urls"], d, pre_append)
	}

	pre_append = pre + ".0." + "rate_css_urls"
	if _, ok := i["rate-css-urls"]; ok {
		result["rate_css_urls"] = flattenObjectWebfilterProfileFtgdWfRateCssUrls(i["rate-css-urls"], d, pre_append)
	}

	pre_append = pre + ".0." + "rate_image_urls"
	if _, ok := i["rate-image-urls"]; ok {
		result["rate_image_urls"] = flattenObjectWebfilterProfileFtgdWfRateImageUrls(i["rate-image-urls"], d, pre_append)
	}

	pre_append = pre + ".0." + "rate_javascript_urls"
	if _, ok := i["rate-javascript-urls"]; ok {
		result["rate_javascript_urls"] = flattenObjectWebfilterProfileFtgdWfRateJavascriptUrls(i["rate-javascript-urls"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWebfilterProfileFtgdWfExemptQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileFtgdWfFilters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWebfilterProfileFtgdWfFiltersAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_usr_grp"
		if _, ok := i["auth-usr-grp"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersAuthUsrGrp(i["auth-usr-grp"], d, pre_append)
			tmp["auth_usr_grp"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-AuthUsrGrp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := i["override-replacemsg"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersOverrideReplacemsg(i["override-replacemsg"], d, pre_append)
			tmp["override_replacemsg"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-OverrideReplacemsg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warn_duration"
		if _, ok := i["warn-duration"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersWarnDuration(i["warn-duration"], d, pre_append)
			tmp["warn_duration"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-WarnDuration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_duration_type"
		if _, ok := i["warning-duration-type"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersWarningDurationType(i["warning-duration-type"], d, pre_append)
			tmp["warning_duration_type"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-WarningDurationType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_prompt"
		if _, ok := i["warning-prompt"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfFiltersWarningPrompt(i["warning-prompt"], d, pre_append)
			tmp["warning_prompt"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Filters-WarningPrompt")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWebfilterProfileFtgdWfFiltersAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersAuthUsrGrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersOverrideReplacemsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersWarnDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersWarningDurationType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfFiltersWarningPrompt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfMaxQuotaTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileFtgdWfOvrd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileFtgdWfQuota(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfQuotaCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Quota-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "duration"
		if _, ok := i["duration"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfQuotaDuration(i["duration"], d, pre_append)
			tmp["duration"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Quota-Duration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfQuotaId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Quota-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := i["override-replacemsg"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfQuotaOverrideReplacemsg(i["override-replacemsg"], d, pre_append)
			tmp["override_replacemsg"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Quota-OverrideReplacemsg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfQuotaType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Quota-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unit"
		if _, ok := i["unit"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfQuotaUnit(i["unit"], d, pre_append)
			tmp["unit"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Quota-Unit")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectWebfilterProfileFtgdWfQuotaValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfileFtgdWf-Quota-Value")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWebfilterProfileFtgdWfQuotaCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfQuotaDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfQuotaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfQuotaOverrideReplacemsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfQuotaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfQuotaUnit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfQuotaValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfRateCrlUrls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfRateCssUrls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfRateImageUrls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFtgdWfRateJavascriptUrls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileHttpsReplacemsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileInspectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileLogAllUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileOverride(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ovrd_cookie"
	if _, ok := i["ovrd-cookie"]; ok {
		result["ovrd_cookie"] = flattenObjectWebfilterProfileOverrideOvrdCookie(i["ovrd-cookie"], d, pre_append)
	}

	pre_append = pre + ".0." + "ovrd_dur"
	if _, ok := i["ovrd-dur"]; ok {
		result["ovrd_dur"] = flattenObjectWebfilterProfileOverrideOvrdDur(i["ovrd-dur"], d, pre_append)
	}

	pre_append = pre + ".0." + "ovrd_dur_mode"
	if _, ok := i["ovrd-dur-mode"]; ok {
		result["ovrd_dur_mode"] = flattenObjectWebfilterProfileOverrideOvrdDurMode(i["ovrd-dur-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ovrd_scope"
	if _, ok := i["ovrd-scope"]; ok {
		result["ovrd_scope"] = flattenObjectWebfilterProfileOverrideOvrdScope(i["ovrd-scope"], d, pre_append)
	}

	pre_append = pre + ".0." + "ovrd_user_group"
	if _, ok := i["ovrd-user-group"]; ok {
		result["ovrd_user_group"] = flattenObjectWebfilterProfileOverrideOvrdUserGroup(i["ovrd-user-group"], d, pre_append)
	}

	pre_append = pre + ".0." + "profile"
	if _, ok := i["profile"]; ok {
		result["profile"] = flattenObjectWebfilterProfileOverrideProfile(i["profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "profile_attribute"
	if _, ok := i["profile-attribute"]; ok {
		result["profile_attribute"] = flattenObjectWebfilterProfileOverrideProfileAttribute(i["profile-attribute"], d, pre_append)
	}

	pre_append = pre + ".0." + "profile_type"
	if _, ok := i["profile-type"]; ok {
		result["profile_type"] = flattenObjectWebfilterProfileOverrideProfileType(i["profile-type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWebfilterProfileOverrideOvrdCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileOverrideOvrdDur(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileOverrideOvrdDurMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileOverrideOvrdScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileOverrideOvrdUserGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileOverrideProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileOverrideProfileAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileOverrideProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileOvrdPerm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfilePostAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileUrlExtraction(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "redirect_header"
	if _, ok := i["redirect-header"]; ok {
		result["redirect_header"] = flattenObjectWebfilterProfileUrlExtractionRedirectHeader(i["redirect-header"], d, pre_append)
	}

	pre_append = pre + ".0." + "redirect_no_content"
	if _, ok := i["redirect-no-content"]; ok {
		result["redirect_no_content"] = flattenObjectWebfilterProfileUrlExtractionRedirectNoContent(i["redirect-no-content"], d, pre_append)
	}

	pre_append = pre + ".0." + "redirect_url"
	if _, ok := i["redirect-url"]; ok {
		result["redirect_url"] = flattenObjectWebfilterProfileUrlExtractionRedirectUrl(i["redirect-url"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_fqdn"
	if _, ok := i["server-fqdn"]; ok {
		result["server_fqdn"] = flattenObjectWebfilterProfileUrlExtractionServerFqdn(i["server-fqdn"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWebfilterProfileUrlExtractionStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWebfilterProfileUrlExtractionRedirectHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileUrlExtractionRedirectNoContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileUrlExtractionRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileUrlExtractionServerFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileUrlExtractionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWeb(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "allowlist"
	if _, ok := i["allowlist"]; ok {
		result["allowlist"] = flattenObjectWebfilterProfileWebAllowlist(i["allowlist"], d, pre_append)
	}

	pre_append = pre + ".0." + "blocklist"
	if _, ok := i["blocklist"]; ok {
		result["blocklist"] = flattenObjectWebfilterProfileWebBlocklist(i["blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "blacklist"
	if _, ok := i["blacklist"]; ok {
		result["blacklist"] = flattenObjectWebfilterProfileWebBlacklist(i["blacklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "bword_table"
	if _, ok := i["bword-table"]; ok {
		result["bword_table"] = flattenObjectWebfilterProfileWebBwordTable(i["bword-table"], d, pre_append)
	}

	pre_append = pre + ".0." + "bword_threshold"
	if _, ok := i["bword-threshold"]; ok {
		result["bword_threshold"] = flattenObjectWebfilterProfileWebBwordThreshold(i["bword-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "content_header_list"
	if _, ok := i["content-header-list"]; ok {
		result["content_header_list"] = flattenObjectWebfilterProfileWebContentHeaderList(i["content-header-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "keyword_match"
	if _, ok := i["keyword-match"]; ok {
		result["keyword_match"] = flattenObjectWebfilterProfileWebKeywordMatch(i["keyword-match"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_search"
	if _, ok := i["log-search"]; ok {
		result["log_search"] = flattenObjectWebfilterProfileWebLogSearch(i["log-search"], d, pre_append)
	}

	pre_append = pre + ".0." + "safe_search"
	if _, ok := i["safe-search"]; ok {
		result["safe_search"] = flattenObjectWebfilterProfileWebSafeSearch(i["safe-search"], d, pre_append)
	}

	pre_append = pre + ".0." + "urlfilter_table"
	if _, ok := i["urlfilter-table"]; ok {
		result["urlfilter_table"] = flattenObjectWebfilterProfileWebUrlfilterTable(i["urlfilter-table"], d, pre_append)
	}

	pre_append = pre + ".0." + "vimeo_restrict"
	if _, ok := i["vimeo-restrict"]; ok {
		result["vimeo_restrict"] = flattenObjectWebfilterProfileWebVimeoRestrict(i["vimeo-restrict"], d, pre_append)
	}

	pre_append = pre + ".0." + "whitelist"
	if _, ok := i["whitelist"]; ok {
		result["whitelist"] = flattenObjectWebfilterProfileWebWhitelist(i["whitelist"], d, pre_append)
	}

	pre_append = pre + ".0." + "youtube_restrict"
	if _, ok := i["youtube-restrict"]; ok {
		result["youtube_restrict"] = flattenObjectWebfilterProfileWebYoutubeRestrict(i["youtube-restrict"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWebfilterProfileWebAllowlist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileWebBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebBlacklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebBwordTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWebfilterProfileWebBwordThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebContentHeaderList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWebfilterProfileWebKeywordMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileWebLogSearch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebSafeSearch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileWebUrlfilterTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWebfilterProfileWebVimeoRestrict(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebWhitelist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileWebYoutubeRestrict(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebAntiphishingLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebContentLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebExtendedAllActionLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebFilterActivexLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebFilterAppletLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebFilterCommandBlockLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebFilterCookieLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebFilterCookieRemovalLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebFilterJsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebFilterJscriptLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebFilterRefererLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebFilterUnknownLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebFilterVbsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebFtgdErrLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebFtgdQuotaUsage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebInvalidDomainLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebUrlLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWisp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWispAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWispServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileYoutubeChannelFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "channel_id"
		if _, ok := i["channel-id"]; ok {
			v := flattenObjectWebfilterProfileYoutubeChannelFilterChannelId(i["channel-id"], d, pre_append)
			tmp["channel_id"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfile-YoutubeChannelFilter-ChannelId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenObjectWebfilterProfileYoutubeChannelFilterComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfile-YoutubeChannelFilter-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWebfilterProfileYoutubeChannelFilterId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWebfilterProfile-YoutubeChannelFilter-Id")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWebfilterProfileYoutubeChannelFilterChannelId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileYoutubeChannelFilterComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileYoutubeChannelFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileYoutubeChannelStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebfilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("antiphish", flattenObjectWebfilterProfileAntiphish(o["antiphish"], d, "antiphish")); err != nil {
			if vv, ok := fortiAPIPatch(o["antiphish"], "ObjectWebfilterProfile-Antiphish"); ok {
				if err = d.Set("antiphish", vv); err != nil {
					return fmt.Errorf("Error reading antiphish: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading antiphish: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("antiphish"); ok {
			if err = d.Set("antiphish", flattenObjectWebfilterProfileAntiphish(o["antiphish"], d, "antiphish")); err != nil {
				if vv, ok := fortiAPIPatch(o["antiphish"], "ObjectWebfilterProfile-Antiphish"); ok {
					if err = d.Set("antiphish", vv); err != nil {
						return fmt.Errorf("Error reading antiphish: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading antiphish: %v", err)
				}
			}
		}
	}

	if err = d.Set("comment", flattenObjectWebfilterProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWebfilterProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenObjectWebfilterProfileExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "ObjectWebfilterProfile-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("file_filter", flattenObjectWebfilterProfileFileFilter(o["file-filter"], d, "file_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["file-filter"], "ObjectWebfilterProfile-FileFilter"); ok {
				if err = d.Set("file_filter", vv); err != nil {
					return fmt.Errorf("Error reading file_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading file_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("file_filter"); ok {
			if err = d.Set("file_filter", flattenObjectWebfilterProfileFileFilter(o["file-filter"], d, "file_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["file-filter"], "ObjectWebfilterProfile-FileFilter"); ok {
					if err = d.Set("file_filter", vv); err != nil {
						return fmt.Errorf("Error reading file_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading file_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("feature_set", flattenObjectWebfilterProfileFeatureSet(o["feature-set"], d, "feature_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["feature-set"], "ObjectWebfilterProfile-FeatureSet"); ok {
			if err = d.Set("feature_set", vv); err != nil {
				return fmt.Errorf("Error reading feature_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ftgd_wf", flattenObjectWebfilterProfileFtgdWf(o["ftgd-wf"], d, "ftgd_wf")); err != nil {
			if vv, ok := fortiAPIPatch(o["ftgd-wf"], "ObjectWebfilterProfile-FtgdWf"); ok {
				if err = d.Set("ftgd_wf", vv); err != nil {
					return fmt.Errorf("Error reading ftgd_wf: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ftgd_wf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftgd_wf"); ok {
			if err = d.Set("ftgd_wf", flattenObjectWebfilterProfileFtgdWf(o["ftgd-wf"], d, "ftgd_wf")); err != nil {
				if vv, ok := fortiAPIPatch(o["ftgd-wf"], "ObjectWebfilterProfile-FtgdWf"); ok {
					if err = d.Set("ftgd_wf", vv); err != nil {
						return fmt.Errorf("Error reading ftgd_wf: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ftgd_wf: %v", err)
				}
			}
		}
	}

	if err = d.Set("https_replacemsg", flattenObjectWebfilterProfileHttpsReplacemsg(o["https-replacemsg"], d, "https_replacemsg")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-replacemsg"], "ObjectWebfilterProfile-HttpsReplacemsg"); ok {
			if err = d.Set("https_replacemsg", vv); err != nil {
				return fmt.Errorf("Error reading https_replacemsg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_replacemsg: %v", err)
		}
	}

	if err = d.Set("inspection_mode", flattenObjectWebfilterProfileInspectionMode(o["inspection-mode"], d, "inspection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["inspection-mode"], "ObjectWebfilterProfile-InspectionMode"); ok {
			if err = d.Set("inspection_mode", vv); err != nil {
				return fmt.Errorf("Error reading inspection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inspection_mode: %v", err)
		}
	}

	if err = d.Set("log_all_url", flattenObjectWebfilterProfileLogAllUrl(o["log-all-url"], d, "log_all_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-all-url"], "ObjectWebfilterProfile-LogAllUrl"); ok {
			if err = d.Set("log_all_url", vv); err != nil {
				return fmt.Errorf("Error reading log_all_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_all_url: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWebfilterProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWebfilterProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("options", flattenObjectWebfilterProfileOptions(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "ObjectWebfilterProfile-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("override", flattenObjectWebfilterProfileOverride(o["override"], d, "override")); err != nil {
			if vv, ok := fortiAPIPatch(o["override"], "ObjectWebfilterProfile-Override"); ok {
				if err = d.Set("override", vv); err != nil {
					return fmt.Errorf("Error reading override: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading override: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("override"); ok {
			if err = d.Set("override", flattenObjectWebfilterProfileOverride(o["override"], d, "override")); err != nil {
				if vv, ok := fortiAPIPatch(o["override"], "ObjectWebfilterProfile-Override"); ok {
					if err = d.Set("override", vv); err != nil {
						return fmt.Errorf("Error reading override: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading override: %v", err)
				}
			}
		}
	}

	if err = d.Set("ovrd_perm", flattenObjectWebfilterProfileOvrdPerm(o["ovrd-perm"], d, "ovrd_perm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-perm"], "ObjectWebfilterProfile-OvrdPerm"); ok {
			if err = d.Set("ovrd_perm", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_perm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_perm: %v", err)
		}
	}

	if err = d.Set("post_action", flattenObjectWebfilterProfilePostAction(o["post-action"], d, "post_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["post-action"], "ObjectWebfilterProfile-PostAction"); ok {
			if err = d.Set("post_action", vv); err != nil {
				return fmt.Errorf("Error reading post_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading post_action: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenObjectWebfilterProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "ObjectWebfilterProfile-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("url_extraction", flattenObjectWebfilterProfileUrlExtraction(o["url-extraction"], d, "url_extraction")); err != nil {
			if vv, ok := fortiAPIPatch(o["url-extraction"], "ObjectWebfilterProfile-UrlExtraction"); ok {
				if err = d.Set("url_extraction", vv); err != nil {
					return fmt.Errorf("Error reading url_extraction: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading url_extraction: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("url_extraction"); ok {
			if err = d.Set("url_extraction", flattenObjectWebfilterProfileUrlExtraction(o["url-extraction"], d, "url_extraction")); err != nil {
				if vv, ok := fortiAPIPatch(o["url-extraction"], "ObjectWebfilterProfile-UrlExtraction"); ok {
					if err = d.Set("url_extraction", vv); err != nil {
						return fmt.Errorf("Error reading url_extraction: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading url_extraction: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("web", flattenObjectWebfilterProfileWeb(o["web"], d, "web")); err != nil {
			if vv, ok := fortiAPIPatch(o["web"], "ObjectWebfilterProfile-Web"); ok {
				if err = d.Set("web", vv); err != nil {
					return fmt.Errorf("Error reading web: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading web: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("web"); ok {
			if err = d.Set("web", flattenObjectWebfilterProfileWeb(o["web"], d, "web")); err != nil {
				if vv, ok := fortiAPIPatch(o["web"], "ObjectWebfilterProfile-Web"); ok {
					if err = d.Set("web", vv); err != nil {
						return fmt.Errorf("Error reading web: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading web: %v", err)
				}
			}
		}
	}

	if err = d.Set("web_antiphishing_log", flattenObjectWebfilterProfileWebAntiphishingLog(o["web-antiphishing-log"], d, "web_antiphishing_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-antiphishing-log"], "ObjectWebfilterProfile-WebAntiphishingLog"); ok {
			if err = d.Set("web_antiphishing_log", vv); err != nil {
				return fmt.Errorf("Error reading web_antiphishing_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_antiphishing_log: %v", err)
		}
	}

	if err = d.Set("web_content_log", flattenObjectWebfilterProfileWebContentLog(o["web-content-log"], d, "web_content_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-content-log"], "ObjectWebfilterProfile-WebContentLog"); ok {
			if err = d.Set("web_content_log", vv); err != nil {
				return fmt.Errorf("Error reading web_content_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_content_log: %v", err)
		}
	}

	if err = d.Set("web_extended_all_action_log", flattenObjectWebfilterProfileWebExtendedAllActionLog(o["web-extended-all-action-log"], d, "web_extended_all_action_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-extended-all-action-log"], "ObjectWebfilterProfile-WebExtendedAllActionLog"); ok {
			if err = d.Set("web_extended_all_action_log", vv); err != nil {
				return fmt.Errorf("Error reading web_extended_all_action_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_extended_all_action_log: %v", err)
		}
	}

	if err = d.Set("web_filter_activex_log", flattenObjectWebfilterProfileWebFilterActivexLog(o["web-filter-activex-log"], d, "web_filter_activex_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-activex-log"], "ObjectWebfilterProfile-WebFilterActivexLog"); ok {
			if err = d.Set("web_filter_activex_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_activex_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_activex_log: %v", err)
		}
	}

	if err = d.Set("web_filter_applet_log", flattenObjectWebfilterProfileWebFilterAppletLog(o["web-filter-applet-log"], d, "web_filter_applet_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-applet-log"], "ObjectWebfilterProfile-WebFilterAppletLog"); ok {
			if err = d.Set("web_filter_applet_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_applet_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_applet_log: %v", err)
		}
	}

	if err = d.Set("web_filter_command_block_log", flattenObjectWebfilterProfileWebFilterCommandBlockLog(o["web-filter-command-block-log"], d, "web_filter_command_block_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-command-block-log"], "ObjectWebfilterProfile-WebFilterCommandBlockLog"); ok {
			if err = d.Set("web_filter_command_block_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_command_block_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_command_block_log: %v", err)
		}
	}

	if err = d.Set("web_filter_cookie_log", flattenObjectWebfilterProfileWebFilterCookieLog(o["web-filter-cookie-log"], d, "web_filter_cookie_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-cookie-log"], "ObjectWebfilterProfile-WebFilterCookieLog"); ok {
			if err = d.Set("web_filter_cookie_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_cookie_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_cookie_log: %v", err)
		}
	}

	if err = d.Set("web_filter_cookie_removal_log", flattenObjectWebfilterProfileWebFilterCookieRemovalLog(o["web-filter-cookie-removal-log"], d, "web_filter_cookie_removal_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-cookie-removal-log"], "ObjectWebfilterProfile-WebFilterCookieRemovalLog"); ok {
			if err = d.Set("web_filter_cookie_removal_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_cookie_removal_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_cookie_removal_log: %v", err)
		}
	}

	if err = d.Set("web_filter_js_log", flattenObjectWebfilterProfileWebFilterJsLog(o["web-filter-js-log"], d, "web_filter_js_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-js-log"], "ObjectWebfilterProfile-WebFilterJsLog"); ok {
			if err = d.Set("web_filter_js_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_js_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_js_log: %v", err)
		}
	}

	if err = d.Set("web_filter_jscript_log", flattenObjectWebfilterProfileWebFilterJscriptLog(o["web-filter-jscript-log"], d, "web_filter_jscript_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-jscript-log"], "ObjectWebfilterProfile-WebFilterJscriptLog"); ok {
			if err = d.Set("web_filter_jscript_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_jscript_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_jscript_log: %v", err)
		}
	}

	if err = d.Set("web_filter_referer_log", flattenObjectWebfilterProfileWebFilterRefererLog(o["web-filter-referer-log"], d, "web_filter_referer_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-referer-log"], "ObjectWebfilterProfile-WebFilterRefererLog"); ok {
			if err = d.Set("web_filter_referer_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_referer_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_referer_log: %v", err)
		}
	}

	if err = d.Set("web_filter_unknown_log", flattenObjectWebfilterProfileWebFilterUnknownLog(o["web-filter-unknown-log"], d, "web_filter_unknown_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-unknown-log"], "ObjectWebfilterProfile-WebFilterUnknownLog"); ok {
			if err = d.Set("web_filter_unknown_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_unknown_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_unknown_log: %v", err)
		}
	}

	if err = d.Set("web_filter_vbs_log", flattenObjectWebfilterProfileWebFilterVbsLog(o["web-filter-vbs-log"], d, "web_filter_vbs_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-vbs-log"], "ObjectWebfilterProfile-WebFilterVbsLog"); ok {
			if err = d.Set("web_filter_vbs_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_vbs_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_vbs_log: %v", err)
		}
	}

	if err = d.Set("web_ftgd_err_log", flattenObjectWebfilterProfileWebFtgdErrLog(o["web-ftgd-err-log"], d, "web_ftgd_err_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-ftgd-err-log"], "ObjectWebfilterProfile-WebFtgdErrLog"); ok {
			if err = d.Set("web_ftgd_err_log", vv); err != nil {
				return fmt.Errorf("Error reading web_ftgd_err_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_ftgd_err_log: %v", err)
		}
	}

	if err = d.Set("web_ftgd_quota_usage", flattenObjectWebfilterProfileWebFtgdQuotaUsage(o["web-ftgd-quota-usage"], d, "web_ftgd_quota_usage")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-ftgd-quota-usage"], "ObjectWebfilterProfile-WebFtgdQuotaUsage"); ok {
			if err = d.Set("web_ftgd_quota_usage", vv); err != nil {
				return fmt.Errorf("Error reading web_ftgd_quota_usage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_ftgd_quota_usage: %v", err)
		}
	}

	if err = d.Set("web_invalid_domain_log", flattenObjectWebfilterProfileWebInvalidDomainLog(o["web-invalid-domain-log"], d, "web_invalid_domain_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-invalid-domain-log"], "ObjectWebfilterProfile-WebInvalidDomainLog"); ok {
			if err = d.Set("web_invalid_domain_log", vv); err != nil {
				return fmt.Errorf("Error reading web_invalid_domain_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_invalid_domain_log: %v", err)
		}
	}

	if err = d.Set("web_url_log", flattenObjectWebfilterProfileWebUrlLog(o["web-url-log"], d, "web_url_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-url-log"], "ObjectWebfilterProfile-WebUrlLog"); ok {
			if err = d.Set("web_url_log", vv); err != nil {
				return fmt.Errorf("Error reading web_url_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_url_log: %v", err)
		}
	}

	if err = d.Set("wisp", flattenObjectWebfilterProfileWisp(o["wisp"], d, "wisp")); err != nil {
		if vv, ok := fortiAPIPatch(o["wisp"], "ObjectWebfilterProfile-Wisp"); ok {
			if err = d.Set("wisp", vv); err != nil {
				return fmt.Errorf("Error reading wisp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wisp: %v", err)
		}
	}

	if err = d.Set("wisp_algorithm", flattenObjectWebfilterProfileWispAlgorithm(o["wisp-algorithm"], d, "wisp_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["wisp-algorithm"], "ObjectWebfilterProfile-WispAlgorithm"); ok {
			if err = d.Set("wisp_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading wisp_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wisp_algorithm: %v", err)
		}
	}

	if err = d.Set("wisp_servers", flattenObjectWebfilterProfileWispServers(o["wisp-servers"], d, "wisp_servers")); err != nil {
		if vv, ok := fortiAPIPatch(o["wisp-servers"], "ObjectWebfilterProfile-WispServers"); ok {
			if err = d.Set("wisp_servers", vv); err != nil {
				return fmt.Errorf("Error reading wisp_servers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wisp_servers: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("youtube_channel_filter", flattenObjectWebfilterProfileYoutubeChannelFilter(o["youtube-channel-filter"], d, "youtube_channel_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["youtube-channel-filter"], "ObjectWebfilterProfile-YoutubeChannelFilter"); ok {
				if err = d.Set("youtube_channel_filter", vv); err != nil {
					return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("youtube_channel_filter"); ok {
			if err = d.Set("youtube_channel_filter", flattenObjectWebfilterProfileYoutubeChannelFilter(o["youtube-channel-filter"], d, "youtube_channel_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["youtube-channel-filter"], "ObjectWebfilterProfile-YoutubeChannelFilter"); ok {
					if err = d.Set("youtube_channel_filter", vv); err != nil {
						return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("youtube_channel_status", flattenObjectWebfilterProfileYoutubeChannelStatus(o["youtube-channel-status"], d, "youtube_channel_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["youtube-channel-status"], "ObjectWebfilterProfile-YoutubeChannelStatus"); ok {
			if err = d.Set("youtube_channel_status", vv); err != nil {
				return fmt.Errorf("Error reading youtube_channel_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading youtube_channel_status: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterProfileAntiphish(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "authentication"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["authentication"], _ = expandObjectWebfilterProfileAntiphishAuthentication(d, i["authentication"], pre_append)
	}
	pre_append = pre + ".0." + "check_basic_auth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["check-basic-auth"], _ = expandObjectWebfilterProfileAntiphishCheckBasicAuth(d, i["check_basic_auth"], pre_append)
	}
	pre_append = pre + ".0." + "check_uri"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["check-uri"], _ = expandObjectWebfilterProfileAntiphishCheckUri(d, i["check_uri"], pre_append)
	}
	pre_append = pre + ".0." + "check_username_only"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["check-username-only"], _ = expandObjectWebfilterProfileAntiphishCheckUsernameOnly(d, i["check_username_only"], pre_append)
	}
	pre_append = pre + ".0." + "custom_patterns"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["custom-patterns"], _ = expandObjectWebfilterProfileAntiphishCustomPatterns(d, i["custom_patterns"], pre_append)
	} else {
		result["custom-patterns"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "default_action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-action"], _ = expandObjectWebfilterProfileAntiphishDefaultAction(d, i["default_action"], pre_append)
	}
	pre_append = pre + ".0." + "domain_controller"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["domain-controller"], _ = expandObjectWebfilterProfileAntiphishDomainController(d, i["domain_controller"], pre_append)
	}
	pre_append = pre + ".0." + "inspection_entries"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["inspection-entries"], _ = expandObjectWebfilterProfileAntiphishInspectionEntries(d, i["inspection_entries"], pre_append)
	} else {
		result["inspection-entries"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ldap"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ldap"], _ = expandObjectWebfilterProfileAntiphishLdap(d, i["ldap"], pre_append)
	}
	pre_append = pre + ".0." + "max_body_len"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-body-len"], _ = expandObjectWebfilterProfileAntiphishMaxBodyLen(d, i["max_body_len"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWebfilterProfileAntiphishStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWebfilterProfileAntiphishAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishCheckBasicAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishCheckUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishCheckUsernameOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishCustomPatterns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandObjectWebfilterProfileAntiphishCustomPatternsCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandObjectWebfilterProfileAntiphishCustomPatternsPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectWebfilterProfileAntiphishCustomPatternsType(d, i["type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWebfilterProfileAntiphishCustomPatternsCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishCustomPatternsPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishCustomPatternsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishDomainController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishInspectionEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectWebfilterProfileAntiphishInspectionEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fortiguard-category"], _ = expandObjectWebfilterProfileAntiphishInspectionEntriesFortiguardCategory(d, i["fortiguard_category"], pre_append)
		} else {
			tmp["fortiguard-category"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectWebfilterProfileAntiphishInspectionEntriesName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWebfilterProfileAntiphishInspectionEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishInspectionEntriesFortiguardCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileAntiphishInspectionEntriesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishLdap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishMaxBodyLen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileAntiphishStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFileFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "entries"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["entries"], _ = expandObjectWebfilterProfileFileFilterEntries(d, i["entries"], pre_append)
	} else {
		result["entries"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectWebfilterProfileFileFilterLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "scan_archive_contents"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["scan-archive-contents"], _ = expandObjectWebfilterProfileFileFilterScanArchiveContents(d, i["scan_archive_contents"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWebfilterProfileFileFilterStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWebfilterProfileFileFilterEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectWebfilterProfileFileFilterEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandObjectWebfilterProfileFileFilterEntriesComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandObjectWebfilterProfileFileFilterEntriesDirection(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encryption"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["encryption"], _ = expandObjectWebfilterProfileFileFilterEntriesEncryption(d, i["encryption"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["file-type"], _ = expandObjectWebfilterProfileFileFilterEntriesFileType(d, i["file_type"], pre_append)
		} else {
			tmp["file-type"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter"], _ = expandObjectWebfilterProfileFileFilterEntriesFilter(d, i["filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_protected"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password-protected"], _ = expandObjectWebfilterProfileFileFilterEntriesPasswordProtected(d, i["password_protected"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandObjectWebfilterProfileFileFilterEntriesProtocol(d, i["protocol"], pre_append)
		} else {
			tmp["protocol"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWebfilterProfileFileFilterEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFileFilterEntriesComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFileFilterEntriesDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFileFilterEntriesEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFileFilterEntriesFileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileFileFilterEntriesFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFileFilterEntriesPasswordProtected(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFileFilterEntriesProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileFileFilterLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFileFilterScanArchiveContents(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFileFilterStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFeatureSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "exempt_quota"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["exempt-quota"], _ = expandObjectWebfilterProfileFtgdWfExemptQuota(d, i["exempt_quota"], pre_append)
	}
	pre_append = pre + ".0." + "filters"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["filters"], _ = expandObjectWebfilterProfileFtgdWfFilters(d, i["filters"], pre_append)
	} else {
		result["filters"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "max_quota_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-quota-timeout"], _ = expandObjectWebfilterProfileFtgdWfMaxQuotaTimeout(d, i["max_quota_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandObjectWebfilterProfileFtgdWfOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ovrd"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ovrd"], _ = expandObjectWebfilterProfileFtgdWfOvrd(d, i["ovrd"], pre_append)
	}
	pre_append = pre + ".0." + "quota"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["quota"], _ = expandObjectWebfilterProfileFtgdWfQuota(d, i["quota"], pre_append)
	} else {
		result["quota"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "rate_crl_urls"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rate-crl-urls"], _ = expandObjectWebfilterProfileFtgdWfRateCrlUrls(d, i["rate_crl_urls"], pre_append)
	}
	pre_append = pre + ".0." + "rate_css_urls"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rate-css-urls"], _ = expandObjectWebfilterProfileFtgdWfRateCssUrls(d, i["rate_css_urls"], pre_append)
	}
	pre_append = pre + ".0." + "rate_image_urls"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rate-image-urls"], _ = expandObjectWebfilterProfileFtgdWfRateImageUrls(d, i["rate_image_urls"], pre_append)
	}
	pre_append = pre + ".0." + "rate_javascript_urls"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rate-javascript-urls"], _ = expandObjectWebfilterProfileFtgdWfRateJavascriptUrls(d, i["rate_javascript_urls"], pre_append)
	}

	return result, nil
}

func expandObjectWebfilterProfileFtgdWfExemptQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectWebfilterProfileFtgdWfFilters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectWebfilterProfileFtgdWfFiltersAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_usr_grp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-usr-grp"], _ = expandObjectWebfilterProfileFtgdWfFiltersAuthUsrGrp(d, i["auth_usr_grp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandObjectWebfilterProfileFtgdWfFiltersCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectWebfilterProfileFtgdWfFiltersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandObjectWebfilterProfileFtgdWfFiltersLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["override-replacemsg"], _ = expandObjectWebfilterProfileFtgdWfFiltersOverrideReplacemsg(d, i["override_replacemsg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warn_duration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["warn-duration"], _ = expandObjectWebfilterProfileFtgdWfFiltersWarnDuration(d, i["warn_duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_duration_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["warning-duration-type"], _ = expandObjectWebfilterProfileFtgdWfFiltersWarningDurationType(d, i["warning_duration_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_prompt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["warning-prompt"], _ = expandObjectWebfilterProfileFtgdWfFiltersWarningPrompt(d, i["warning_prompt"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersAuthUsrGrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersOverrideReplacemsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersWarnDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersWarningDurationType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfFiltersWarningPrompt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfMaxQuotaTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileFtgdWfOvrd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectWebfilterProfileFtgdWfQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandObjectWebfilterProfileFtgdWfQuotaCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "duration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["duration"], _ = expandObjectWebfilterProfileFtgdWfQuotaDuration(d, i["duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectWebfilterProfileFtgdWfQuotaId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["override-replacemsg"], _ = expandObjectWebfilterProfileFtgdWfQuotaOverrideReplacemsg(d, i["override_replacemsg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectWebfilterProfileFtgdWfQuotaType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unit"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["unit"], _ = expandObjectWebfilterProfileFtgdWfQuotaUnit(d, i["unit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectWebfilterProfileFtgdWfQuotaValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWebfilterProfileFtgdWfQuotaCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfQuotaDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfQuotaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfQuotaOverrideReplacemsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfQuotaType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfQuotaUnit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfQuotaValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfRateCrlUrls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfRateCssUrls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfRateImageUrls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFtgdWfRateJavascriptUrls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileHttpsReplacemsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileInspectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileLogAllUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ovrd_cookie"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ovrd-cookie"], _ = expandObjectWebfilterProfileOverrideOvrdCookie(d, i["ovrd_cookie"], pre_append)
	}
	pre_append = pre + ".0." + "ovrd_dur"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ovrd-dur"], _ = expandObjectWebfilterProfileOverrideOvrdDur(d, i["ovrd_dur"], pre_append)
	}
	pre_append = pre + ".0." + "ovrd_dur_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ovrd-dur-mode"], _ = expandObjectWebfilterProfileOverrideOvrdDurMode(d, i["ovrd_dur_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ovrd_scope"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ovrd-scope"], _ = expandObjectWebfilterProfileOverrideOvrdScope(d, i["ovrd_scope"], pre_append)
	}
	pre_append = pre + ".0." + "ovrd_user_group"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ovrd-user-group"], _ = expandObjectWebfilterProfileOverrideOvrdUserGroup(d, i["ovrd_user_group"], pre_append)
	}
	pre_append = pre + ".0." + "profile"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["profile"], _ = expandObjectWebfilterProfileOverrideProfile(d, i["profile"], pre_append)
	}
	pre_append = pre + ".0." + "profile_attribute"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["profile-attribute"], _ = expandObjectWebfilterProfileOverrideProfileAttribute(d, i["profile_attribute"], pre_append)
	}
	pre_append = pre + ".0." + "profile_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["profile-type"], _ = expandObjectWebfilterProfileOverrideProfileType(d, i["profile_type"], pre_append)
	}

	return result, nil
}

func expandObjectWebfilterProfileOverrideOvrdCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileOverrideOvrdDur(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileOverrideOvrdDurMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileOverrideOvrdScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileOverrideOvrdUserGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectWebfilterProfileOverrideProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectWebfilterProfileOverrideProfileAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileOverrideProfileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileOvrdPerm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfilePostAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileUrlExtraction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "redirect_header"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redirect-header"], _ = expandObjectWebfilterProfileUrlExtractionRedirectHeader(d, i["redirect_header"], pre_append)
	}
	pre_append = pre + ".0." + "redirect_no_content"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redirect-no-content"], _ = expandObjectWebfilterProfileUrlExtractionRedirectNoContent(d, i["redirect_no_content"], pre_append)
	}
	pre_append = pre + ".0." + "redirect_url"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redirect-url"], _ = expandObjectWebfilterProfileUrlExtractionRedirectUrl(d, i["redirect_url"], pre_append)
	}
	pre_append = pre + ".0." + "server_fqdn"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server-fqdn"], _ = expandObjectWebfilterProfileUrlExtractionServerFqdn(d, i["server_fqdn"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectWebfilterProfileUrlExtractionStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectWebfilterProfileUrlExtractionRedirectHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileUrlExtractionRedirectNoContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileUrlExtractionRedirectUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileUrlExtractionServerFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileUrlExtractionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWeb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "allowlist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["allowlist"], _ = expandObjectWebfilterProfileWebAllowlist(d, i["allowlist"], pre_append)
	} else {
		result["allowlist"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["blocklist"], _ = expandObjectWebfilterProfileWebBlocklist(d, i["blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "blacklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["blacklist"], _ = expandObjectWebfilterProfileWebBlacklist(d, i["blacklist"], pre_append)
	}
	pre_append = pre + ".0." + "bword_table"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bword-table"], _ = expandObjectWebfilterProfileWebBwordTable(d, i["bword_table"], pre_append)
	}
	pre_append = pre + ".0." + "bword_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bword-threshold"], _ = expandObjectWebfilterProfileWebBwordThreshold(d, i["bword_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "content_header_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["content-header-list"], _ = expandObjectWebfilterProfileWebContentHeaderList(d, i["content_header_list"], pre_append)
	}
	pre_append = pre + ".0." + "keyword_match"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keyword-match"], _ = expandObjectWebfilterProfileWebKeywordMatch(d, i["keyword_match"], pre_append)
	} else {
		result["keyword-match"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "log_search"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log-search"], _ = expandObjectWebfilterProfileWebLogSearch(d, i["log_search"], pre_append)
	}
	pre_append = pre + ".0." + "safe_search"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["safe-search"], _ = expandObjectWebfilterProfileWebSafeSearch(d, i["safe_search"], pre_append)
	} else {
		result["safe-search"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "urlfilter_table"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["urlfilter-table"], _ = expandObjectWebfilterProfileWebUrlfilterTable(d, i["urlfilter_table"], pre_append)
	}
	pre_append = pre + ".0." + "vimeo_restrict"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vimeo-restrict"], _ = expandObjectWebfilterProfileWebVimeoRestrict(d, i["vimeo_restrict"], pre_append)
	}
	pre_append = pre + ".0." + "whitelist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["whitelist"], _ = expandObjectWebfilterProfileWebWhitelist(d, i["whitelist"], pre_append)
	} else {
		result["whitelist"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "youtube_restrict"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["youtube-restrict"], _ = expandObjectWebfilterProfileWebYoutubeRestrict(d, i["youtube_restrict"], pre_append)
	}

	return result, nil
}

func expandObjectWebfilterProfileWebAllowlist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileWebBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebBlacklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebBwordTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebBwordThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebContentHeaderList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebKeywordMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileWebLogSearch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebSafeSearch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileWebUrlfilterTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebVimeoRestrict(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebWhitelist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileWebYoutubeRestrict(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebAntiphishingLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebContentLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebExtendedAllActionLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebFilterActivexLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebFilterAppletLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebFilterCommandBlockLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebFilterCookieLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebFilterCookieRemovalLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebFilterJsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebFilterJscriptLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebFilterRefererLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebFilterUnknownLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebFilterVbsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebFtgdErrLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebFtgdQuotaUsage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebInvalidDomainLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebUrlLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWisp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWispAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWispServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileYoutubeChannelFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "channel_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["channel-id"], _ = expandObjectWebfilterProfileYoutubeChannelFilterChannelId(d, i["channel_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandObjectWebfilterProfileYoutubeChannelFilterComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectWebfilterProfileYoutubeChannelFilterId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWebfilterProfileYoutubeChannelFilterChannelId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileYoutubeChannelFilterComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileYoutubeChannelFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileYoutubeChannelStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebfilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("antiphish"); ok || d.HasChange("antiphish") {
		t, err := expandObjectWebfilterProfileAntiphish(d, v, "antiphish")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antiphish"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectWebfilterProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandObjectWebfilterProfileExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("file_filter"); ok || d.HasChange("file_filter") {
		t, err := expandObjectWebfilterProfileFileFilter(d, v, "file_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok || d.HasChange("feature_set") {
		t, err := expandObjectWebfilterProfileFeatureSet(d, v, "feature_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("ftgd_wf"); ok || d.HasChange("ftgd_wf") {
		t, err := expandObjectWebfilterProfileFtgdWf(d, v, "ftgd_wf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftgd-wf"] = t
		}
	}

	if v, ok := d.GetOk("https_replacemsg"); ok || d.HasChange("https_replacemsg") {
		t, err := expandObjectWebfilterProfileHttpsReplacemsg(d, v, "https_replacemsg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-replacemsg"] = t
		}
	}

	if v, ok := d.GetOk("inspection_mode"); ok || d.HasChange("inspection_mode") {
		t, err := expandObjectWebfilterProfileInspectionMode(d, v, "inspection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspection-mode"] = t
		}
	}

	if v, ok := d.GetOk("log_all_url"); ok || d.HasChange("log_all_url") {
		t, err := expandObjectWebfilterProfileLogAllUrl(d, v, "log_all_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-all-url"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWebfilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandObjectWebfilterProfileOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("override"); ok || d.HasChange("override") {
		t, err := expandObjectWebfilterProfileOverride(d, v, "override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_perm"); ok || d.HasChange("ovrd_perm") {
		t, err := expandObjectWebfilterProfileOvrdPerm(d, v, "ovrd_perm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-perm"] = t
		}
	}

	if v, ok := d.GetOk("post_action"); ok || d.HasChange("post_action") {
		t, err := expandObjectWebfilterProfilePostAction(d, v, "post_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["post-action"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandObjectWebfilterProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("url_extraction"); ok || d.HasChange("url_extraction") {
		t, err := expandObjectWebfilterProfileUrlExtraction(d, v, "url_extraction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-extraction"] = t
		}
	}

	if v, ok := d.GetOk("web"); ok || d.HasChange("web") {
		t, err := expandObjectWebfilterProfileWeb(d, v, "web")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web"] = t
		}
	}

	if v, ok := d.GetOk("web_antiphishing_log"); ok || d.HasChange("web_antiphishing_log") {
		t, err := expandObjectWebfilterProfileWebAntiphishingLog(d, v, "web_antiphishing_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-antiphishing-log"] = t
		}
	}

	if v, ok := d.GetOk("web_content_log"); ok || d.HasChange("web_content_log") {
		t, err := expandObjectWebfilterProfileWebContentLog(d, v, "web_content_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-content-log"] = t
		}
	}

	if v, ok := d.GetOk("web_extended_all_action_log"); ok || d.HasChange("web_extended_all_action_log") {
		t, err := expandObjectWebfilterProfileWebExtendedAllActionLog(d, v, "web_extended_all_action_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-extended-all-action-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_activex_log"); ok || d.HasChange("web_filter_activex_log") {
		t, err := expandObjectWebfilterProfileWebFilterActivexLog(d, v, "web_filter_activex_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-activex-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_applet_log"); ok || d.HasChange("web_filter_applet_log") {
		t, err := expandObjectWebfilterProfileWebFilterAppletLog(d, v, "web_filter_applet_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-applet-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_command_block_log"); ok || d.HasChange("web_filter_command_block_log") {
		t, err := expandObjectWebfilterProfileWebFilterCommandBlockLog(d, v, "web_filter_command_block_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-command-block-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_cookie_log"); ok || d.HasChange("web_filter_cookie_log") {
		t, err := expandObjectWebfilterProfileWebFilterCookieLog(d, v, "web_filter_cookie_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-cookie-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_cookie_removal_log"); ok || d.HasChange("web_filter_cookie_removal_log") {
		t, err := expandObjectWebfilterProfileWebFilterCookieRemovalLog(d, v, "web_filter_cookie_removal_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-cookie-removal-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_js_log"); ok || d.HasChange("web_filter_js_log") {
		t, err := expandObjectWebfilterProfileWebFilterJsLog(d, v, "web_filter_js_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-js-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_jscript_log"); ok || d.HasChange("web_filter_jscript_log") {
		t, err := expandObjectWebfilterProfileWebFilterJscriptLog(d, v, "web_filter_jscript_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-jscript-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_referer_log"); ok || d.HasChange("web_filter_referer_log") {
		t, err := expandObjectWebfilterProfileWebFilterRefererLog(d, v, "web_filter_referer_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-referer-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_unknown_log"); ok || d.HasChange("web_filter_unknown_log") {
		t, err := expandObjectWebfilterProfileWebFilterUnknownLog(d, v, "web_filter_unknown_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-unknown-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_vbs_log"); ok || d.HasChange("web_filter_vbs_log") {
		t, err := expandObjectWebfilterProfileWebFilterVbsLog(d, v, "web_filter_vbs_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-vbs-log"] = t
		}
	}

	if v, ok := d.GetOk("web_ftgd_err_log"); ok || d.HasChange("web_ftgd_err_log") {
		t, err := expandObjectWebfilterProfileWebFtgdErrLog(d, v, "web_ftgd_err_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-ftgd-err-log"] = t
		}
	}

	if v, ok := d.GetOk("web_ftgd_quota_usage"); ok || d.HasChange("web_ftgd_quota_usage") {
		t, err := expandObjectWebfilterProfileWebFtgdQuotaUsage(d, v, "web_ftgd_quota_usage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-ftgd-quota-usage"] = t
		}
	}

	if v, ok := d.GetOk("web_invalid_domain_log"); ok || d.HasChange("web_invalid_domain_log") {
		t, err := expandObjectWebfilterProfileWebInvalidDomainLog(d, v, "web_invalid_domain_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-invalid-domain-log"] = t
		}
	}

	if v, ok := d.GetOk("web_url_log"); ok || d.HasChange("web_url_log") {
		t, err := expandObjectWebfilterProfileWebUrlLog(d, v, "web_url_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-url-log"] = t
		}
	}

	if v, ok := d.GetOk("wisp"); ok || d.HasChange("wisp") {
		t, err := expandObjectWebfilterProfileWisp(d, v, "wisp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wisp"] = t
		}
	}

	if v, ok := d.GetOk("wisp_algorithm"); ok || d.HasChange("wisp_algorithm") {
		t, err := expandObjectWebfilterProfileWispAlgorithm(d, v, "wisp_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wisp-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("wisp_servers"); ok || d.HasChange("wisp_servers") {
		t, err := expandObjectWebfilterProfileWispServers(d, v, "wisp_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wisp-servers"] = t
		}
	}

	if v, ok := d.GetOk("youtube_channel_filter"); ok || d.HasChange("youtube_channel_filter") {
		t, err := expandObjectWebfilterProfileYoutubeChannelFilter(d, v, "youtube_channel_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-channel-filter"] = t
		}
	}

	if v, ok := d.GetOk("youtube_channel_status"); ok || d.HasChange("youtube_channel_status") {
		t, err := expandObjectWebfilterProfileYoutubeChannelStatus(d, v, "youtube_channel_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-channel-status"] = t
		}
	}

	return &obj, nil
}
