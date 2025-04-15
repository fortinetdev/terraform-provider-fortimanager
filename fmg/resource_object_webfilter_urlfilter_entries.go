// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: URL filter entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebfilterUrlfilterEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterUrlfilterEntriesCreate,
		Read:   resourceObjectWebfilterUrlfilterEntriesRead,
		Update: resourceObjectWebfilterUrlfilterEntriesUpdate,
		Delete: resourceObjectWebfilterUrlfilterEntriesDelete,

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
			"urlfilter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"antiphish_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_address_family": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exempt": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"referrer_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_proxy_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectWebfilterUrlfilterEntriesCreate(d *schema.ResourceData, m interface{}) error {
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

	urlfilter := d.Get("urlfilter").(string)
	paradict["urlfilter"] = urlfilter

	obj, err := getObjectObjectWebfilterUrlfilterEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterUrlfilterEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWebfilterUrlfilterEntries(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterUrlfilterEntries resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWebfilterUrlfilterEntriesRead(d, m)
}

func resourceObjectWebfilterUrlfilterEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	urlfilter := d.Get("urlfilter").(string)
	paradict["urlfilter"] = urlfilter

	obj, err := getObjectObjectWebfilterUrlfilterEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterUrlfilterEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWebfilterUrlfilterEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterUrlfilterEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWebfilterUrlfilterEntriesRead(d, m)
}

func resourceObjectWebfilterUrlfilterEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	urlfilter := d.Get("urlfilter").(string)
	paradict["urlfilter"] = urlfilter

	wsParams["adom"] = adomv

	err = c.DeleteObjectWebfilterUrlfilterEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterUrlfilterEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterUrlfilterEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	urlfilter := d.Get("urlfilter").(string)
	if urlfilter == "" {
		urlfilter = importOptionChecking(m.(*FortiClient).Cfg, "urlfilter")
		if urlfilter == "" {
			return fmt.Errorf("Parameter urlfilter is missing")
		}
		if err = d.Set("urlfilter", urlfilter); err != nil {
			return fmt.Errorf("Error set params urlfilter: %v", err)
		}
	}
	paradict["urlfilter"] = urlfilter

	o, err := c.ReadObjectWebfilterUrlfilterEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterUrlfilterEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterUrlfilterEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterUrlfilterEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterUrlfilterEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesAntiphishAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesDnsAddressFamily2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesExempt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterUrlfilterEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesReferrerHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesWebProxyProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectObjectWebfilterUrlfilterEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectWebfilterUrlfilterEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectWebfilterUrlfilterEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("antiphish_action", flattenObjectWebfilterUrlfilterEntriesAntiphishAction2edl(o["antiphish-action"], d, "antiphish_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["antiphish-action"], "ObjectWebfilterUrlfilterEntries-AntiphishAction"); ok {
			if err = d.Set("antiphish_action", vv); err != nil {
				return fmt.Errorf("Error reading antiphish_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antiphish_action: %v", err)
		}
	}

	if err = d.Set("dns_address_family", flattenObjectWebfilterUrlfilterEntriesDnsAddressFamily2edl(o["dns-address-family"], d, "dns_address_family")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-address-family"], "ObjectWebfilterUrlfilterEntries-DnsAddressFamily"); ok {
			if err = d.Set("dns_address_family", vv); err != nil {
				return fmt.Errorf("Error reading dns_address_family: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_address_family: %v", err)
		}
	}

	if err = d.Set("exempt", flattenObjectWebfilterUrlfilterEntriesExempt2edl(o["exempt"], d, "exempt")); err != nil {
		if vv, ok := fortiAPIPatch(o["exempt"], "ObjectWebfilterUrlfilterEntries-Exempt"); ok {
			if err = d.Set("exempt", vv); err != nil {
				return fmt.Errorf("Error reading exempt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exempt: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectWebfilterUrlfilterEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWebfilterUrlfilterEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("referrer_host", flattenObjectWebfilterUrlfilterEntriesReferrerHost2edl(o["referrer-host"], d, "referrer_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["referrer-host"], "ObjectWebfilterUrlfilterEntries-ReferrerHost"); ok {
			if err = d.Set("referrer_host", vv); err != nil {
				return fmt.Errorf("Error reading referrer_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading referrer_host: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectWebfilterUrlfilterEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectWebfilterUrlfilterEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectWebfilterUrlfilterEntriesType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectWebfilterUrlfilterEntries-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("url", flattenObjectWebfilterUrlfilterEntriesUrl2edl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "ObjectWebfilterUrlfilterEntries-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("web_proxy_profile", flattenObjectWebfilterUrlfilterEntriesWebProxyProfile2edl(o["web-proxy-profile"], d, "web_proxy_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-proxy-profile"], "ObjectWebfilterUrlfilterEntries-WebProxyProfile"); ok {
			if err = d.Set("web_proxy_profile", vv); err != nil {
				return fmt.Errorf("Error reading web_proxy_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_proxy_profile: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterUrlfilterEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterUrlfilterEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesAntiphishAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesDnsAddressFamily2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesExempt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterUrlfilterEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesReferrerHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesWebProxyProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectObjectWebfilterUrlfilterEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectWebfilterUrlfilterEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("antiphish_action"); ok || d.HasChange("antiphish_action") {
		t, err := expandObjectWebfilterUrlfilterEntriesAntiphishAction2edl(d, v, "antiphish_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antiphish-action"] = t
		}
	}

	if v, ok := d.GetOk("dns_address_family"); ok || d.HasChange("dns_address_family") {
		t, err := expandObjectWebfilterUrlfilterEntriesDnsAddressFamily2edl(d, v, "dns_address_family")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-address-family"] = t
		}
	}

	if v, ok := d.GetOk("exempt"); ok || d.HasChange("exempt") {
		t, err := expandObjectWebfilterUrlfilterEntriesExempt2edl(d, v, "exempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exempt"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWebfilterUrlfilterEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("referrer_host"); ok || d.HasChange("referrer_host") {
		t, err := expandObjectWebfilterUrlfilterEntriesReferrerHost2edl(d, v, "referrer_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["referrer-host"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectWebfilterUrlfilterEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectWebfilterUrlfilterEntriesType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandObjectWebfilterUrlfilterEntriesUrl2edl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("web_proxy_profile"); ok || d.HasChange("web_proxy_profile") {
		t, err := expandObjectWebfilterUrlfilterEntriesWebProxyProfile2edl(d, v, "web_proxy_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-proxy-profile"] = t
		}
	}

	return &obj, nil
}
