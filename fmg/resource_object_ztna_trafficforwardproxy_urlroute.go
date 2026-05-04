// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectZtna TrafficForwardProxyUrlRoute

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectZtnaTrafficForwardProxyUrlRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectZtnaTrafficForwardProxyUrlRouteCreate,
		Read:   resourceObjectZtnaTrafficForwardProxyUrlRouteRead,
		Update: resourceObjectZtnaTrafficForwardProxyUrlRouteUpdate,
		Delete: resourceObjectZtnaTrafficForwardProxyUrlRouteDelete,

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
			"traffic_forward_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"service_connector": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"url_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectZtnaTrafficForwardProxyUrlRouteCreate(d *schema.ResourceData, m interface{}) error {
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

	traffic_forward_proxy := d.Get("traffic_forward_proxy").(string)
	paradict["traffic_forward_proxy"] = traffic_forward_proxy

	obj, err := getObjectObjectZtnaTrafficForwardProxyUrlRoute(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectZtnaTrafficForwardProxyUrlRoute resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectZtnaTrafficForwardProxyUrlRoute(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectZtnaTrafficForwardProxyUrlRoute(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectZtnaTrafficForwardProxyUrlRoute resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectZtnaTrafficForwardProxyUrlRoute(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectZtnaTrafficForwardProxyUrlRoute resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectZtnaTrafficForwardProxyUrlRouteRead(d, m)
}

func resourceObjectZtnaTrafficForwardProxyUrlRouteUpdate(d *schema.ResourceData, m interface{}) error {
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

	traffic_forward_proxy := d.Get("traffic_forward_proxy").(string)
	paradict["traffic_forward_proxy"] = traffic_forward_proxy

	obj, err := getObjectObjectZtnaTrafficForwardProxyUrlRoute(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectZtnaTrafficForwardProxyUrlRoute resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectZtnaTrafficForwardProxyUrlRoute(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectZtnaTrafficForwardProxyUrlRoute resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectZtnaTrafficForwardProxyUrlRouteRead(d, m)
}

func resourceObjectZtnaTrafficForwardProxyUrlRouteDelete(d *schema.ResourceData, m interface{}) error {
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

	traffic_forward_proxy := d.Get("traffic_forward_proxy").(string)
	paradict["traffic_forward_proxy"] = traffic_forward_proxy

	wsParams["adom"] = adomv

	err = c.DeleteObjectZtnaTrafficForwardProxyUrlRoute(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectZtnaTrafficForwardProxyUrlRoute resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectZtnaTrafficForwardProxyUrlRouteRead(d *schema.ResourceData, m interface{}) error {
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

	traffic_forward_proxy := d.Get("traffic_forward_proxy").(string)
	if traffic_forward_proxy == "" {
		traffic_forward_proxy = importOptionChecking(m.(*FortiClient).Cfg, "traffic_forward_proxy")
		if traffic_forward_proxy == "" {
			return fmt.Errorf("Parameter traffic_forward_proxy is missing")
		}
		if err = d.Set("traffic_forward_proxy", traffic_forward_proxy); err != nil {
			return fmt.Errorf("Error set params traffic_forward_proxy: %v", err)
		}
	}
	paradict["traffic_forward_proxy"] = traffic_forward_proxy

	o, err := c.ReadObjectZtnaTrafficForwardProxyUrlRoute(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectZtnaTrafficForwardProxyUrlRoute resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectZtnaTrafficForwardProxyUrlRoute(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectZtnaTrafficForwardProxyUrlRoute resource from API: %v", err)
	}
	return nil
}

func flattenObjectZtnaTrafficForwardProxyUrlRouteName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaTrafficForwardProxyUrlRouteServiceConnector2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectZtnaTrafficForwardProxyUrlRouteUrlPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectZtnaTrafficForwardProxyUrlRoute(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectZtnaTrafficForwardProxyUrlRouteName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectZtnaTrafficForwardProxyUrlRoute-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("service_connector", flattenObjectZtnaTrafficForwardProxyUrlRouteServiceConnector2edl(o["service-connector"], d, "service_connector")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-connector"], "ObjectZtnaTrafficForwardProxyUrlRoute-ServiceConnector"); ok {
			if err = d.Set("service_connector", vv); err != nil {
				return fmt.Errorf("Error reading service_connector: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_connector: %v", err)
		}
	}

	if err = d.Set("url_pattern", flattenObjectZtnaTrafficForwardProxyUrlRouteUrlPattern2edl(o["url-pattern"], d, "url_pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-pattern"], "ObjectZtnaTrafficForwardProxyUrlRoute-UrlPattern"); ok {
			if err = d.Set("url_pattern", vv); err != nil {
				return fmt.Errorf("Error reading url_pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_pattern: %v", err)
		}
	}

	return nil
}

func flattenObjectZtnaTrafficForwardProxyUrlRouteFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectZtnaTrafficForwardProxyUrlRouteName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaTrafficForwardProxyUrlRouteServiceConnector2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaTrafficForwardProxyUrlRouteUrlPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectZtnaTrafficForwardProxyUrlRoute(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectZtnaTrafficForwardProxyUrlRouteName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("service_connector"); ok || d.HasChange("service_connector") {
		t, err := expandObjectZtnaTrafficForwardProxyUrlRouteServiceConnector2edl(d, v, "service_connector")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-connector"] = t
		}
	}

	if v, ok := d.GetOk("url_pattern"); ok || d.HasChange("url_pattern") {
		t, err := expandObjectZtnaTrafficForwardProxyUrlRouteUrlPattern2edl(d, v, "url_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-pattern"] = t
		}
	}

	return &obj, nil
}
