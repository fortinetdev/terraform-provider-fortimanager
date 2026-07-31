// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectZtna WebPortalBookmarkLlmSecureProxy

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectZtnaWebPortalBookmarkLlmSecureProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectZtnaWebPortalBookmarkLlmSecureProxyUpdate,
		Read:   resourceObjectZtnaWebPortalBookmarkLlmSecureProxyRead,
		Update: resourceObjectZtnaWebPortalBookmarkLlmSecureProxyUpdate,
		Delete: resourceObjectZtnaWebPortalBookmarkLlmSecureProxyDelete,

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
			"web_portal_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"all_llm_servers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"llm_servers": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectZtnaWebPortalBookmarkLlmSecureProxyUpdate(d *schema.ResourceData, m interface{}) error {
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

	web_portal_bookmark := d.Get("web_portal_bookmark").(string)
	paradict["web_portal_bookmark"] = web_portal_bookmark

	obj, err := getObjectObjectZtnaWebPortalBookmarkLlmSecureProxy(d, false)
	if err != nil {
		return fmt.Errorf("Error updating ObjectZtnaWebPortalBookmarkLlmSecureProxy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectZtnaWebPortalBookmarkLlmSecureProxy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectZtnaWebPortalBookmarkLlmSecureProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectZtnaWebPortalBookmarkLlmSecureProxy")

	return resourceObjectZtnaWebPortalBookmarkLlmSecureProxyRead(d, m)
}

func resourceObjectZtnaWebPortalBookmarkLlmSecureProxyDelete(d *schema.ResourceData, m interface{}) error {
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

	web_portal_bookmark := d.Get("web_portal_bookmark").(string)
	paradict["web_portal_bookmark"] = web_portal_bookmark

	obj, err := getObjectObjectZtnaWebPortalBookmarkLlmSecureProxy(d, true)

	if err != nil {
		return fmt.Errorf("Error updating ObjectZtnaWebPortalBookmarkLlmSecureProxy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectZtnaWebPortalBookmarkLlmSecureProxy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing ObjectZtnaWebPortalBookmarkLlmSecureProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectZtnaWebPortalBookmarkLlmSecureProxyRead(d *schema.ResourceData, m interface{}) error {
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

	web_portal_bookmark := d.Get("web_portal_bookmark").(string)
	if web_portal_bookmark == "" {
		web_portal_bookmark = importOptionChecking(m.(*FortiClient).Cfg, "web_portal_bookmark")
		if web_portal_bookmark == "" {
			return fmt.Errorf("Parameter web_portal_bookmark is missing")
		}
		if err = d.Set("web_portal_bookmark", web_portal_bookmark); err != nil {
			return fmt.Errorf("Error set params web_portal_bookmark: %v", err)
		}
	}
	paradict["web_portal_bookmark"] = web_portal_bookmark

	o, err := c.ReadObjectZtnaWebPortalBookmarkLlmSecureProxy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectZtnaWebPortalBookmarkLlmSecureProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectZtnaWebPortalBookmarkLlmSecureProxy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectZtnaWebPortalBookmarkLlmSecureProxy resource from API: %v", err)
	}
	return nil
}

func flattenObjectZtnaWebPortalBookmarkLlmSecureProxyAllLlmServers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkLlmSecureProxyLlmServers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectZtnaWebPortalBookmarkLlmSecureProxy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("all_llm_servers", flattenObjectZtnaWebPortalBookmarkLlmSecureProxyAllLlmServers2edl(o["all-llm-servers"], d, "all_llm_servers")); err != nil {
		if vv, ok := fortiAPIPatch(o["all-llm-servers"], "ObjectZtnaWebPortalBookmarkLlmSecureProxy-AllLlmServers"); ok {
			if err = d.Set("all_llm_servers", vv); err != nil {
				return fmt.Errorf("Error reading all_llm_servers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading all_llm_servers: %v", err)
		}
	}

	if err = d.Set("llm_servers", flattenObjectZtnaWebPortalBookmarkLlmSecureProxyLlmServers2edl(o["llm-servers"], d, "llm_servers")); err != nil {
		if vv, ok := fortiAPIPatch(o["llm-servers"], "ObjectZtnaWebPortalBookmarkLlmSecureProxy-LlmServers"); ok {
			if err = d.Set("llm_servers", vv); err != nil {
				return fmt.Errorf("Error reading llm_servers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading llm_servers: %v", err)
		}
	}

	return nil
}

func flattenObjectZtnaWebPortalBookmarkLlmSecureProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectZtnaWebPortalBookmarkLlmSecureProxyAllLlmServers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkLlmSecureProxyLlmServers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectZtnaWebPortalBookmarkLlmSecureProxy(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("all_llm_servers"); ok || d.HasChange("all_llm_servers") {
		t, err := expandObjectZtnaWebPortalBookmarkLlmSecureProxyAllLlmServers2edl(d, v, "all_llm_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["all-llm-servers"] = t
		}
	}

	if v, ok := d.GetOk("llm_servers"); ok || d.HasChange("llm_servers") {
		t, err := expandObjectZtnaWebPortalBookmarkLlmSecureProxyLlmServers2edl(d, v, "llm_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["llm-servers"] = t
		}
	}

	return &obj, nil
}
