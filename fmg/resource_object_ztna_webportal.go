// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure ztna web-portal.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectZtnaWebPortal() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectZtnaWebPortalCreate,
		Read:   resourceObjectZtnaWebPortalRead,
		Update: resourceObjectZtnaWebPortalUpdate,
		Delete: resourceObjectZtnaWebPortalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"auth_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_rule": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auth_virtual_host": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"clipboard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookie_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"customize_forticlient_download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"decrypted_traffic_mirror": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"default_window_height": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"default_window_width": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"display_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_history": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"focus_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_download": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_download_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"heading": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"log_blocked_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"macos_forticlient_download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"policy_auth_sso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vip6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"windows_forticlient_download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bookmarks": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"llm_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"llm_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectZtnaWebPortalCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectZtnaWebPortal(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectZtnaWebPortal resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectZtnaWebPortal(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectZtnaWebPortal(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectZtnaWebPortal resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectZtnaWebPortal(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectZtnaWebPortal resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectZtnaWebPortalRead(d, m)
}

func resourceObjectZtnaWebPortalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectZtnaWebPortal(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectZtnaWebPortal resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectZtnaWebPortal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectZtnaWebPortal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectZtnaWebPortalRead(d, m)
}

func resourceObjectZtnaWebPortalDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectZtnaWebPortal(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectZtnaWebPortal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectZtnaWebPortalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectZtnaWebPortal(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectZtnaWebPortal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectZtnaWebPortal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectZtnaWebPortal resource from API: %v", err)
	}
	return nil
}

func flattenObjectZtnaWebPortalAuthPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalAuthRule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectZtnaWebPortalAuthVirtualHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectZtnaWebPortalClipboard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalCustomizeForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalDecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectZtnaWebPortalDefaultWindowHeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalDefaultWindowWidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalDisplayBookmark(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalDisplayHistory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalDisplayStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalFocusBookmark(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalForticlientDownload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalForticlientDownloadMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalHeading(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectZtnaWebPortalHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectZtnaWebPortalLogBlockedTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalMacosForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalPolicyAuthSso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalTheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectZtnaWebPortalVip6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectZtnaWebPortalWindowsForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarks(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectZtnaWebPortalLlmProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectZtnaWebPortalLlmProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectZtnaWebPortal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("auth_portal", flattenObjectZtnaWebPortalAuthPortal(o["auth-portal"], d, "auth_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-portal"], "ObjectZtnaWebPortal-AuthPortal"); ok {
			if err = d.Set("auth_portal", vv); err != nil {
				return fmt.Errorf("Error reading auth_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_portal: %v", err)
		}
	}

	if err = d.Set("auth_rule", flattenObjectZtnaWebPortalAuthRule(o["auth-rule"], d, "auth_rule")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-rule"], "ObjectZtnaWebPortal-AuthRule"); ok {
			if err = d.Set("auth_rule", vv); err != nil {
				return fmt.Errorf("Error reading auth_rule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_rule: %v", err)
		}
	}

	if err = d.Set("auth_virtual_host", flattenObjectZtnaWebPortalAuthVirtualHost(o["auth-virtual-host"], d, "auth_virtual_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-virtual-host"], "ObjectZtnaWebPortal-AuthVirtualHost"); ok {
			if err = d.Set("auth_virtual_host", vv); err != nil {
				return fmt.Errorf("Error reading auth_virtual_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_virtual_host: %v", err)
		}
	}

	if err = d.Set("clipboard", flattenObjectZtnaWebPortalClipboard(o["clipboard"], d, "clipboard")); err != nil {
		if vv, ok := fortiAPIPatch(o["clipboard"], "ObjectZtnaWebPortal-Clipboard"); ok {
			if err = d.Set("clipboard", vv); err != nil {
				return fmt.Errorf("Error reading clipboard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clipboard: %v", err)
		}
	}

	if err = d.Set("cookie_age", flattenObjectZtnaWebPortalCookieAge(o["cookie-age"], d, "cookie_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["cookie-age"], "ObjectZtnaWebPortal-CookieAge"); ok {
			if err = d.Set("cookie_age", vv); err != nil {
				return fmt.Errorf("Error reading cookie_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cookie_age: %v", err)
		}
	}

	if err = d.Set("customize_forticlient_download_url", flattenObjectZtnaWebPortalCustomizeForticlientDownloadUrl(o["customize-forticlient-download-url"], d, "customize_forticlient_download_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["customize-forticlient-download-url"], "ObjectZtnaWebPortal-CustomizeForticlientDownloadUrl"); ok {
			if err = d.Set("customize_forticlient_download_url", vv); err != nil {
				return fmt.Errorf("Error reading customize_forticlient_download_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading customize_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenObjectZtnaWebPortalDecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror")); err != nil {
		if vv, ok := fortiAPIPatch(o["decrypted-traffic-mirror"], "ObjectZtnaWebPortal-DecryptedTrafficMirror"); ok {
			if err = d.Set("decrypted_traffic_mirror", vv); err != nil {
				return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if err = d.Set("default_window_height", flattenObjectZtnaWebPortalDefaultWindowHeight(o["default-window-height"], d, "default_window_height")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-window-height"], "ObjectZtnaWebPortal-DefaultWindowHeight"); ok {
			if err = d.Set("default_window_height", vv); err != nil {
				return fmt.Errorf("Error reading default_window_height: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_window_height: %v", err)
		}
	}

	if err = d.Set("default_window_width", flattenObjectZtnaWebPortalDefaultWindowWidth(o["default-window-width"], d, "default_window_width")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-window-width"], "ObjectZtnaWebPortal-DefaultWindowWidth"); ok {
			if err = d.Set("default_window_width", vv); err != nil {
				return fmt.Errorf("Error reading default_window_width: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_window_width: %v", err)
		}
	}

	if err = d.Set("display_bookmark", flattenObjectZtnaWebPortalDisplayBookmark(o["display-bookmark"], d, "display_bookmark")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-bookmark"], "ObjectZtnaWebPortal-DisplayBookmark"); ok {
			if err = d.Set("display_bookmark", vv); err != nil {
				return fmt.Errorf("Error reading display_bookmark: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_bookmark: %v", err)
		}
	}

	if err = d.Set("display_history", flattenObjectZtnaWebPortalDisplayHistory(o["display-history"], d, "display_history")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-history"], "ObjectZtnaWebPortal-DisplayHistory"); ok {
			if err = d.Set("display_history", vv); err != nil {
				return fmt.Errorf("Error reading display_history: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_history: %v", err)
		}
	}

	if err = d.Set("display_status", flattenObjectZtnaWebPortalDisplayStatus(o["display-status"], d, "display_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-status"], "ObjectZtnaWebPortal-DisplayStatus"); ok {
			if err = d.Set("display_status", vv); err != nil {
				return fmt.Errorf("Error reading display_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_status: %v", err)
		}
	}

	if err = d.Set("focus_bookmark", flattenObjectZtnaWebPortalFocusBookmark(o["focus-bookmark"], d, "focus_bookmark")); err != nil {
		if vv, ok := fortiAPIPatch(o["focus-bookmark"], "ObjectZtnaWebPortal-FocusBookmark"); ok {
			if err = d.Set("focus_bookmark", vv); err != nil {
				return fmt.Errorf("Error reading focus_bookmark: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading focus_bookmark: %v", err)
		}
	}

	if err = d.Set("forticlient_download", flattenObjectZtnaWebPortalForticlientDownload(o["forticlient-download"], d, "forticlient_download")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-download"], "ObjectZtnaWebPortal-ForticlientDownload"); ok {
			if err = d.Set("forticlient_download", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_download: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_download: %v", err)
		}
	}

	if err = d.Set("forticlient_download_method", flattenObjectZtnaWebPortalForticlientDownloadMethod(o["forticlient-download-method"], d, "forticlient_download_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-download-method"], "ObjectZtnaWebPortal-ForticlientDownloadMethod"); ok {
			if err = d.Set("forticlient_download_method", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_download_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_download_method: %v", err)
		}
	}

	if err = d.Set("heading", flattenObjectZtnaWebPortalHeading(o["heading"], d, "heading")); err != nil {
		if vv, ok := fortiAPIPatch(o["heading"], "ObjectZtnaWebPortal-Heading"); ok {
			if err = d.Set("heading", vv); err != nil {
				return fmt.Errorf("Error reading heading: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading heading: %v", err)
		}
	}

	if err = d.Set("host", flattenObjectZtnaWebPortalHost(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "ObjectZtnaWebPortal-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("log_blocked_traffic", flattenObjectZtnaWebPortalLogBlockedTraffic(o["log-blocked-traffic"], d, "log_blocked_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-blocked-traffic"], "ObjectZtnaWebPortal-LogBlockedTraffic"); ok {
			if err = d.Set("log_blocked_traffic", vv); err != nil {
				return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
		}
	}

	if err = d.Set("macos_forticlient_download_url", flattenObjectZtnaWebPortalMacosForticlientDownloadUrl(o["macos-forticlient-download-url"], d, "macos_forticlient_download_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["macos-forticlient-download-url"], "ObjectZtnaWebPortal-MacosForticlientDownloadUrl"); ok {
			if err = d.Set("macos_forticlient_download_url", vv); err != nil {
				return fmt.Errorf("Error reading macos_forticlient_download_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading macos_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectZtnaWebPortalName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectZtnaWebPortal-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("policy_auth_sso", flattenObjectZtnaWebPortalPolicyAuthSso(o["policy-auth-sso"], d, "policy_auth_sso")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-auth-sso"], "ObjectZtnaWebPortal-PolicyAuthSso"); ok {
			if err = d.Set("policy_auth_sso", vv); err != nil {
				return fmt.Errorf("Error reading policy_auth_sso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_auth_sso: %v", err)
		}
	}

	if err = d.Set("theme", flattenObjectZtnaWebPortalTheme(o["theme"], d, "theme")); err != nil {
		if vv, ok := fortiAPIPatch(o["theme"], "ObjectZtnaWebPortal-Theme"); ok {
			if err = d.Set("theme", vv); err != nil {
				return fmt.Errorf("Error reading theme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading theme: %v", err)
		}
	}

	if err = d.Set("vip", flattenObjectZtnaWebPortalVip(o["vip"], d, "vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip"], "ObjectZtnaWebPortal-Vip"); ok {
			if err = d.Set("vip", vv); err != nil {
				return fmt.Errorf("Error reading vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip: %v", err)
		}
	}

	if err = d.Set("vip6", flattenObjectZtnaWebPortalVip6(o["vip6"], d, "vip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip6"], "ObjectZtnaWebPortal-Vip6"); ok {
			if err = d.Set("vip6", vv); err != nil {
				return fmt.Errorf("Error reading vip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip6: %v", err)
		}
	}

	if err = d.Set("windows_forticlient_download_url", flattenObjectZtnaWebPortalWindowsForticlientDownloadUrl(o["windows-forticlient-download-url"], d, "windows_forticlient_download_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["windows-forticlient-download-url"], "ObjectZtnaWebPortal-WindowsForticlientDownloadUrl"); ok {
			if err = d.Set("windows_forticlient_download_url", vv); err != nil {
				return fmt.Errorf("Error reading windows_forticlient_download_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading windows_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("bookmarks", flattenObjectZtnaWebPortalBookmarks(o["bookmarks"], d, "bookmarks")); err != nil {
		if vv, ok := fortiAPIPatch(o["bookmarks"], "ObjectZtnaWebPortal-Bookmarks"); ok {
			if err = d.Set("bookmarks", vv); err != nil {
				return fmt.Errorf("Error reading bookmarks: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bookmarks: %v", err)
		}
	}

	if err = d.Set("llm_profile", flattenObjectZtnaWebPortalLlmProfile(o["llm-profile"], d, "llm_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["llm-profile"], "ObjectZtnaWebPortal-LlmProfile"); ok {
			if err = d.Set("llm_profile", vv); err != nil {
				return fmt.Errorf("Error reading llm_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading llm_profile: %v", err)
		}
	}

	if err = d.Set("llm_proxy", flattenObjectZtnaWebPortalLlmProxy(o["llm-proxy"], d, "llm_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["llm-proxy"], "ObjectZtnaWebPortal-LlmProxy"); ok {
			if err = d.Set("llm_proxy", vv); err != nil {
				return fmt.Errorf("Error reading llm_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading llm_proxy: %v", err)
		}
	}

	return nil
}

func flattenObjectZtnaWebPortalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectZtnaWebPortalAuthPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalAuthRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaWebPortalAuthVirtualHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaWebPortalClipboard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalCustomizeForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalDecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaWebPortalDefaultWindowHeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalDefaultWindowWidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalDisplayBookmark(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalDisplayHistory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalDisplayStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalFocusBookmark(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalForticlientDownload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalForticlientDownloadMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalHeading(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaWebPortalLogBlockedTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalMacosForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalPolicyAuthSso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalTheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaWebPortalVip6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaWebPortalWindowsForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaWebPortalLlmProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaWebPortalLlmProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectZtnaWebPortal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_portal"); ok || d.HasChange("auth_portal") {
		t, err := expandObjectZtnaWebPortalAuthPortal(d, v, "auth_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal"] = t
		}
	}

	if v, ok := d.GetOk("auth_rule"); ok || d.HasChange("auth_rule") {
		t, err := expandObjectZtnaWebPortalAuthRule(d, v, "auth_rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-rule"] = t
		}
	}

	if v, ok := d.GetOk("auth_virtual_host"); ok || d.HasChange("auth_virtual_host") {
		t, err := expandObjectZtnaWebPortalAuthVirtualHost(d, v, "auth_virtual_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-virtual-host"] = t
		}
	}

	if v, ok := d.GetOk("clipboard"); ok || d.HasChange("clipboard") {
		t, err := expandObjectZtnaWebPortalClipboard(d, v, "clipboard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clipboard"] = t
		}
	}

	if v, ok := d.GetOk("cookie_age"); ok || d.HasChange("cookie_age") {
		t, err := expandObjectZtnaWebPortalCookieAge(d, v, "cookie_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie-age"] = t
		}
	}

	if v, ok := d.GetOk("customize_forticlient_download_url"); ok || d.HasChange("customize_forticlient_download_url") {
		t, err := expandObjectZtnaWebPortalCustomizeForticlientDownloadUrl(d, v, "customize_forticlient_download_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["customize-forticlient-download-url"] = t
		}
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok || d.HasChange("decrypted_traffic_mirror") {
		t, err := expandObjectZtnaWebPortalDecryptedTrafficMirror(d, v, "decrypted_traffic_mirror")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	}

	if v, ok := d.GetOk("default_window_height"); ok || d.HasChange("default_window_height") {
		t, err := expandObjectZtnaWebPortalDefaultWindowHeight(d, v, "default_window_height")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-window-height"] = t
		}
	}

	if v, ok := d.GetOk("default_window_width"); ok || d.HasChange("default_window_width") {
		t, err := expandObjectZtnaWebPortalDefaultWindowWidth(d, v, "default_window_width")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-window-width"] = t
		}
	}

	if v, ok := d.GetOk("display_bookmark"); ok || d.HasChange("display_bookmark") {
		t, err := expandObjectZtnaWebPortalDisplayBookmark(d, v, "display_bookmark")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("display_history"); ok || d.HasChange("display_history") {
		t, err := expandObjectZtnaWebPortalDisplayHistory(d, v, "display_history")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-history"] = t
		}
	}

	if v, ok := d.GetOk("display_status"); ok || d.HasChange("display_status") {
		t, err := expandObjectZtnaWebPortalDisplayStatus(d, v, "display_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-status"] = t
		}
	}

	if v, ok := d.GetOk("focus_bookmark"); ok || d.HasChange("focus_bookmark") {
		t, err := expandObjectZtnaWebPortalFocusBookmark(d, v, "focus_bookmark")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["focus-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_download"); ok || d.HasChange("forticlient_download") {
		t, err := expandObjectZtnaWebPortalForticlientDownload(d, v, "forticlient_download")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-download"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_download_method"); ok || d.HasChange("forticlient_download_method") {
		t, err := expandObjectZtnaWebPortalForticlientDownloadMethod(d, v, "forticlient_download_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-download-method"] = t
		}
	}

	if v, ok := d.GetOk("heading"); ok || d.HasChange("heading") {
		t, err := expandObjectZtnaWebPortalHeading(d, v, "heading")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heading"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandObjectZtnaWebPortalHost(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("log_blocked_traffic"); ok || d.HasChange("log_blocked_traffic") {
		t, err := expandObjectZtnaWebPortalLogBlockedTraffic(d, v, "log_blocked_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-blocked-traffic"] = t
		}
	}

	if v, ok := d.GetOk("macos_forticlient_download_url"); ok || d.HasChange("macos_forticlient_download_url") {
		t, err := expandObjectZtnaWebPortalMacosForticlientDownloadUrl(d, v, "macos_forticlient_download_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macos-forticlient-download-url"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectZtnaWebPortalName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("policy_auth_sso"); ok || d.HasChange("policy_auth_sso") {
		t, err := expandObjectZtnaWebPortalPolicyAuthSso(d, v, "policy_auth_sso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-auth-sso"] = t
		}
	}

	if v, ok := d.GetOk("theme"); ok || d.HasChange("theme") {
		t, err := expandObjectZtnaWebPortalTheme(d, v, "theme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["theme"] = t
		}
	}

	if v, ok := d.GetOk("vip"); ok || d.HasChange("vip") {
		t, err := expandObjectZtnaWebPortalVip(d, v, "vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip"] = t
		}
	}

	if v, ok := d.GetOk("vip6"); ok || d.HasChange("vip6") {
		t, err := expandObjectZtnaWebPortalVip6(d, v, "vip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip6"] = t
		}
	}

	if v, ok := d.GetOk("windows_forticlient_download_url"); ok || d.HasChange("windows_forticlient_download_url") {
		t, err := expandObjectZtnaWebPortalWindowsForticlientDownloadUrl(d, v, "windows_forticlient_download_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["windows-forticlient-download-url"] = t
		}
	}

	if v, ok := d.GetOk("bookmarks"); ok || d.HasChange("bookmarks") {
		t, err := expandObjectZtnaWebPortalBookmarks(d, v, "bookmarks")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bookmarks"] = t
		}
	}

	if v, ok := d.GetOk("llm_profile"); ok || d.HasChange("llm_profile") {
		t, err := expandObjectZtnaWebPortalLlmProfile(d, v, "llm_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["llm-profile"] = t
		}
	}

	if v, ok := d.GetOk("llm_proxy"); ok || d.HasChange("llm_proxy") {
		t, err := expandObjectZtnaWebPortalLlmProxy(d, v, "llm_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["llm-proxy"] = t
		}
	}

	return &obj, nil
}
