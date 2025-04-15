// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Web content filtering settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebfilterProfileWeb() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterProfileWebUpdate,
		Read:   resourceObjectWebfilterProfileWebRead,
		Update: resourceObjectWebfilterProfileWebUpdate,
		Delete: resourceObjectWebfilterProfileWebDelete,

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
			},
			"bword_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"content_header_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"vimeo_restrict": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
	}
}

func resourceObjectWebfilterProfileWebUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWebfilterProfileWeb(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileWeb resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWebfilterProfileWeb(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileWeb resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectWebfilterProfileWeb")

	return resourceObjectWebfilterProfileWebRead(d, m)
}

func resourceObjectWebfilterProfileWebDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWebfilterProfileWeb(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterProfileWeb resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterProfileWebRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWebfilterProfileWeb(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileWeb resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterProfileWeb(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileWeb resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterProfileWebAllowlist2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileWebBlocklist2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebBlacklist2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebBwordTable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWebfilterProfileWebBwordThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebContentHeaderList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWebfilterProfileWebKeywordMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileWebLogSearch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebSafeSearch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileWebUrlfilterTable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWebfilterProfileWebVimeoRestrict2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileWebWhitelist2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileWebYoutubeRestrict2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebfilterProfileWeb(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("allowlist", flattenObjectWebfilterProfileWebAllowlist2edl(o["allowlist"], d, "allowlist")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowlist"], "ObjectWebfilterProfileWeb-Allowlist"); ok {
			if err = d.Set("allowlist", vv); err != nil {
				return fmt.Errorf("Error reading allowlist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowlist: %v", err)
		}
	}

	if err = d.Set("blocklist", flattenObjectWebfilterProfileWebBlocklist2edl(o["blocklist"], d, "blocklist")); err != nil {
		if vv, ok := fortiAPIPatch(o["blocklist"], "ObjectWebfilterProfileWeb-Blocklist"); ok {
			if err = d.Set("blocklist", vv); err != nil {
				return fmt.Errorf("Error reading blocklist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading blocklist: %v", err)
		}
	}

	if err = d.Set("blacklist", flattenObjectWebfilterProfileWebBlacklist2edl(o["blacklist"], d, "blacklist")); err != nil {
		if vv, ok := fortiAPIPatch(o["blacklist"], "ObjectWebfilterProfileWeb-Blacklist"); ok {
			if err = d.Set("blacklist", vv); err != nil {
				return fmt.Errorf("Error reading blacklist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading blacklist: %v", err)
		}
	}

	if err = d.Set("bword_table", flattenObjectWebfilterProfileWebBwordTable2edl(o["bword-table"], d, "bword_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["bword-table"], "ObjectWebfilterProfileWeb-BwordTable"); ok {
			if err = d.Set("bword_table", vv); err != nil {
				return fmt.Errorf("Error reading bword_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bword_table: %v", err)
		}
	}

	if err = d.Set("bword_threshold", flattenObjectWebfilterProfileWebBwordThreshold2edl(o["bword-threshold"], d, "bword_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["bword-threshold"], "ObjectWebfilterProfileWeb-BwordThreshold"); ok {
			if err = d.Set("bword_threshold", vv); err != nil {
				return fmt.Errorf("Error reading bword_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bword_threshold: %v", err)
		}
	}

	if err = d.Set("content_header_list", flattenObjectWebfilterProfileWebContentHeaderList2edl(o["content-header-list"], d, "content_header_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["content-header-list"], "ObjectWebfilterProfileWeb-ContentHeaderList"); ok {
			if err = d.Set("content_header_list", vv); err != nil {
				return fmt.Errorf("Error reading content_header_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading content_header_list: %v", err)
		}
	}

	if err = d.Set("keyword_match", flattenObjectWebfilterProfileWebKeywordMatch2edl(o["keyword-match"], d, "keyword_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["keyword-match"], "ObjectWebfilterProfileWeb-KeywordMatch"); ok {
			if err = d.Set("keyword_match", vv); err != nil {
				return fmt.Errorf("Error reading keyword_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keyword_match: %v", err)
		}
	}

	if err = d.Set("log_search", flattenObjectWebfilterProfileWebLogSearch2edl(o["log-search"], d, "log_search")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-search"], "ObjectWebfilterProfileWeb-LogSearch"); ok {
			if err = d.Set("log_search", vv); err != nil {
				return fmt.Errorf("Error reading log_search: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_search: %v", err)
		}
	}

	if err = d.Set("safe_search", flattenObjectWebfilterProfileWebSafeSearch2edl(o["safe-search"], d, "safe_search")); err != nil {
		if vv, ok := fortiAPIPatch(o["safe-search"], "ObjectWebfilterProfileWeb-SafeSearch"); ok {
			if err = d.Set("safe_search", vv); err != nil {
				return fmt.Errorf("Error reading safe_search: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading safe_search: %v", err)
		}
	}

	if err = d.Set("urlfilter_table", flattenObjectWebfilterProfileWebUrlfilterTable2edl(o["urlfilter-table"], d, "urlfilter_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["urlfilter-table"], "ObjectWebfilterProfileWeb-UrlfilterTable"); ok {
			if err = d.Set("urlfilter_table", vv); err != nil {
				return fmt.Errorf("Error reading urlfilter_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading urlfilter_table: %v", err)
		}
	}

	if err = d.Set("vimeo_restrict", flattenObjectWebfilterProfileWebVimeoRestrict2edl(o["vimeo-restrict"], d, "vimeo_restrict")); err != nil {
		if vv, ok := fortiAPIPatch(o["vimeo-restrict"], "ObjectWebfilterProfileWeb-VimeoRestrict"); ok {
			if err = d.Set("vimeo_restrict", vv); err != nil {
				return fmt.Errorf("Error reading vimeo_restrict: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vimeo_restrict: %v", err)
		}
	}

	if err = d.Set("whitelist", flattenObjectWebfilterProfileWebWhitelist2edl(o["whitelist"], d, "whitelist")); err != nil {
		if vv, ok := fortiAPIPatch(o["whitelist"], "ObjectWebfilterProfileWeb-Whitelist"); ok {
			if err = d.Set("whitelist", vv); err != nil {
				return fmt.Errorf("Error reading whitelist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading whitelist: %v", err)
		}
	}

	if err = d.Set("youtube_restrict", flattenObjectWebfilterProfileWebYoutubeRestrict2edl(o["youtube-restrict"], d, "youtube_restrict")); err != nil {
		if vv, ok := fortiAPIPatch(o["youtube-restrict"], "ObjectWebfilterProfileWeb-YoutubeRestrict"); ok {
			if err = d.Set("youtube_restrict", vv); err != nil {
				return fmt.Errorf("Error reading youtube_restrict: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading youtube_restrict: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterProfileWebFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterProfileWebAllowlist2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileWebBlocklist2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebBlacklist2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebBwordTable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWebfilterProfileWebBwordThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebContentHeaderList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWebfilterProfileWebKeywordMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileWebLogSearch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebSafeSearch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileWebUrlfilterTable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWebfilterProfileWebVimeoRestrict2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileWebWhitelist2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileWebYoutubeRestrict2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebfilterProfileWeb(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allowlist"); ok || d.HasChange("allowlist") {
		t, err := expandObjectWebfilterProfileWebAllowlist2edl(d, v, "allowlist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowlist"] = t
		}
	}

	if v, ok := d.GetOk("blocklist"); ok || d.HasChange("blocklist") {
		t, err := expandObjectWebfilterProfileWebBlocklist2edl(d, v, "blocklist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blocklist"] = t
		}
	}

	if v, ok := d.GetOk("blacklist"); ok || d.HasChange("blacklist") {
		t, err := expandObjectWebfilterProfileWebBlacklist2edl(d, v, "blacklist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blacklist"] = t
		}
	}

	if v, ok := d.GetOk("bword_table"); ok || d.HasChange("bword_table") {
		t, err := expandObjectWebfilterProfileWebBwordTable2edl(d, v, "bword_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bword-table"] = t
		}
	}

	if v, ok := d.GetOk("bword_threshold"); ok || d.HasChange("bword_threshold") {
		t, err := expandObjectWebfilterProfileWebBwordThreshold2edl(d, v, "bword_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bword-threshold"] = t
		}
	}

	if v, ok := d.GetOk("content_header_list"); ok || d.HasChange("content_header_list") {
		t, err := expandObjectWebfilterProfileWebContentHeaderList2edl(d, v, "content_header_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-header-list"] = t
		}
	}

	if v, ok := d.GetOk("keyword_match"); ok || d.HasChange("keyword_match") {
		t, err := expandObjectWebfilterProfileWebKeywordMatch2edl(d, v, "keyword_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyword-match"] = t
		}
	}

	if v, ok := d.GetOk("log_search"); ok || d.HasChange("log_search") {
		t, err := expandObjectWebfilterProfileWebLogSearch2edl(d, v, "log_search")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-search"] = t
		}
	}

	if v, ok := d.GetOk("safe_search"); ok || d.HasChange("safe_search") {
		t, err := expandObjectWebfilterProfileWebSafeSearch2edl(d, v, "safe_search")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["safe-search"] = t
		}
	}

	if v, ok := d.GetOk("urlfilter_table"); ok || d.HasChange("urlfilter_table") {
		t, err := expandObjectWebfilterProfileWebUrlfilterTable2edl(d, v, "urlfilter_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["urlfilter-table"] = t
		}
	}

	if v, ok := d.GetOk("vimeo_restrict"); ok || d.HasChange("vimeo_restrict") {
		t, err := expandObjectWebfilterProfileWebVimeoRestrict2edl(d, v, "vimeo_restrict")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vimeo-restrict"] = t
		}
	}

	if v, ok := d.GetOk("whitelist"); ok || d.HasChange("whitelist") {
		t, err := expandObjectWebfilterProfileWebWhitelist2edl(d, v, "whitelist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["whitelist"] = t
		}
	}

	if v, ok := d.GetOk("youtube_restrict"); ok || d.HasChange("youtube_restrict") {
		t, err := expandObjectWebfilterProfileWebYoutubeRestrict2edl(d, v, "youtube_restrict")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-restrict"] = t
		}
	}

	return &obj, nil
}
