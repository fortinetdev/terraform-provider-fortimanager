// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Landing page options.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnSslWebPortalLandingPage() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnSslWebPortalLandingPageUpdate,
		Read:   resourceObjectVpnSslWebPortalLandingPageRead,
		Update: resourceObjectVpnSslWebPortalLandingPageUpdate,
		Delete: resourceObjectVpnSslWebPortalLandingPageDelete,

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
			"portal": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"form_data": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"logout_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso_credential": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"sso_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectVpnSslWebPortalLandingPageUpdate(d *schema.ResourceData, m interface{}) error {
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

	portal := d.Get("portal").(string)
	paradict["portal"] = portal

	obj, err := getObjectObjectVpnSslWebPortalLandingPage(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebPortalLandingPage resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVpnSslWebPortalLandingPage(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebPortalLandingPage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectVpnSslWebPortalLandingPage")

	return resourceObjectVpnSslWebPortalLandingPageRead(d, m)
}

func resourceObjectVpnSslWebPortalLandingPageDelete(d *schema.ResourceData, m interface{}) error {
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

	portal := d.Get("portal").(string)
	paradict["portal"] = portal

	err = c.DeleteObjectVpnSslWebPortalLandingPage(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnSslWebPortalLandingPage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnSslWebPortalLandingPageRead(d *schema.ResourceData, m interface{}) error {
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

	portal := d.Get("portal").(string)
	if portal == "" {
		portal = importOptionChecking(m.(*FortiClient).Cfg, "portal")
		if portal == "" {
			return fmt.Errorf("Parameter portal is missing")
		}
		if err = d.Set("portal", portal); err != nil {
			return fmt.Errorf("Error set params portal: %v", err)
		}
	}
	paradict["portal"] = portal

	o, err := c.ReadObjectVpnSslWebPortalLandingPage(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebPortalLandingPage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnSslWebPortalLandingPage(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebPortalLandingPage resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnSslWebPortalLandingPageFormData2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectVpnSslWebPortalLandingPageFormDataName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalLandingPage-FormData-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectVpnSslWebPortalLandingPageFormDataValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalLandingPage-FormData-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectVpnSslWebPortalLandingPageFormDataName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLandingPageFormDataValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLandingPageLogoutUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLandingPageSso2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLandingPageSsoCredential2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLandingPageSsoUsername2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLandingPageUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnSslWebPortalLandingPage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("form_data", flattenObjectVpnSslWebPortalLandingPageFormData2edl(o["form-data"], d, "form_data")); err != nil {
			if vv, ok := fortiAPIPatch(o["form-data"], "ObjectVpnSslWebPortalLandingPage-FormData"); ok {
				if err = d.Set("form_data", vv); err != nil {
					return fmt.Errorf("Error reading form_data: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading form_data: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("form_data"); ok {
			if err = d.Set("form_data", flattenObjectVpnSslWebPortalLandingPageFormData2edl(o["form-data"], d, "form_data")); err != nil {
				if vv, ok := fortiAPIPatch(o["form-data"], "ObjectVpnSslWebPortalLandingPage-FormData"); ok {
					if err = d.Set("form_data", vv); err != nil {
						return fmt.Errorf("Error reading form_data: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading form_data: %v", err)
				}
			}
		}
	}

	if err = d.Set("logout_url", flattenObjectVpnSslWebPortalLandingPageLogoutUrl2edl(o["logout-url"], d, "logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["logout-url"], "ObjectVpnSslWebPortalLandingPage-LogoutUrl"); ok {
			if err = d.Set("logout_url", vv); err != nil {
				return fmt.Errorf("Error reading logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logout_url: %v", err)
		}
	}

	if err = d.Set("sso", flattenObjectVpnSslWebPortalLandingPageSso2edl(o["sso"], d, "sso")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso"], "ObjectVpnSslWebPortalLandingPage-Sso"); ok {
			if err = d.Set("sso", vv); err != nil {
				return fmt.Errorf("Error reading sso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso: %v", err)
		}
	}

	if err = d.Set("sso_credential", flattenObjectVpnSslWebPortalLandingPageSsoCredential2edl(o["sso-credential"], d, "sso_credential")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-credential"], "ObjectVpnSslWebPortalLandingPage-SsoCredential"); ok {
			if err = d.Set("sso_credential", vv); err != nil {
				return fmt.Errorf("Error reading sso_credential: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_credential: %v", err)
		}
	}

	if err = d.Set("sso_username", flattenObjectVpnSslWebPortalLandingPageSsoUsername2edl(o["sso-username"], d, "sso_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-username"], "ObjectVpnSslWebPortalLandingPage-SsoUsername"); ok {
			if err = d.Set("sso_username", vv); err != nil {
				return fmt.Errorf("Error reading sso_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_username: %v", err)
		}
	}

	if err = d.Set("url", flattenObjectVpnSslWebPortalLandingPageUrl2edl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "ObjectVpnSslWebPortalLandingPage-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnSslWebPortalLandingPageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnSslWebPortalLandingPageFormData2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectVpnSslWebPortalLandingPageFormDataName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectVpnSslWebPortalLandingPageFormDataValue2edl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectVpnSslWebPortalLandingPageFormDataName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLandingPageFormDataValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLandingPageLogoutUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLandingPageSso2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLandingPageSsoCredential2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLandingPageSsoPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnSslWebPortalLandingPageSsoUsername2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLandingPageUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnSslWebPortalLandingPage(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("form_data"); ok || d.HasChange("form_data") {
		t, err := expandObjectVpnSslWebPortalLandingPageFormData2edl(d, v, "form_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["form-data"] = t
		}
	}

	if v, ok := d.GetOk("logout_url"); ok || d.HasChange("logout_url") {
		t, err := expandObjectVpnSslWebPortalLandingPageLogoutUrl2edl(d, v, "logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logout-url"] = t
		}
	}

	if v, ok := d.GetOk("sso"); ok || d.HasChange("sso") {
		t, err := expandObjectVpnSslWebPortalLandingPageSso2edl(d, v, "sso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso"] = t
		}
	}

	if v, ok := d.GetOk("sso_credential"); ok || d.HasChange("sso_credential") {
		t, err := expandObjectVpnSslWebPortalLandingPageSsoCredential2edl(d, v, "sso_credential")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-credential"] = t
		}
	}

	if v, ok := d.GetOk("sso_password"); ok || d.HasChange("sso_password") {
		t, err := expandObjectVpnSslWebPortalLandingPageSsoPassword2edl(d, v, "sso_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-password"] = t
		}
	}

	if v, ok := d.GetOk("sso_username"); ok || d.HasChange("sso_username") {
		t, err := expandObjectVpnSslWebPortalLandingPageSsoUsername2edl(d, v, "sso_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-username"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandObjectVpnSslWebPortalLandingPageUrl2edl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	return &obj, nil
}
