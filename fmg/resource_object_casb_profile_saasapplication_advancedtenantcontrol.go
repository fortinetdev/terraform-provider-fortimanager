// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB profile advanced tenant control.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbProfileSaasApplicationAdvancedTenantControl() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbProfileSaasApplicationAdvancedTenantControlCreate,
		Read:   resourceObjectCasbProfileSaasApplicationAdvancedTenantControlRead,
		Update: resourceObjectCasbProfileSaasApplicationAdvancedTenantControlUpdate,
		Delete: resourceObjectCasbProfileSaasApplicationAdvancedTenantControlDelete,

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
			"saas_application": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"attribute": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"input": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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

func resourceObjectCasbProfileSaasApplicationAdvancedTenantControlCreate(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application

	obj, err := getObjectObjectCasbProfileSaasApplicationAdvancedTenantControl(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfileSaasApplicationAdvancedTenantControl resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectCasbProfileSaasApplicationAdvancedTenantControl(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfileSaasApplicationAdvancedTenantControl resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbProfileSaasApplicationAdvancedTenantControlRead(d, m)
}

func resourceObjectCasbProfileSaasApplicationAdvancedTenantControlUpdate(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application

	obj, err := getObjectObjectCasbProfileSaasApplicationAdvancedTenantControl(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfileSaasApplicationAdvancedTenantControl resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbProfileSaasApplicationAdvancedTenantControl(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfileSaasApplicationAdvancedTenantControl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbProfileSaasApplicationAdvancedTenantControlRead(d, m)
}

func resourceObjectCasbProfileSaasApplicationAdvancedTenantControlDelete(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application

	wsParams["adom"] = adomv

	err = c.DeleteObjectCasbProfileSaasApplicationAdvancedTenantControl(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbProfileSaasApplicationAdvancedTenantControl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbProfileSaasApplicationAdvancedTenantControlRead(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	if saas_application == "" {
		saas_application = importOptionChecking(m.(*FortiClient).Cfg, "saas_application")
		if saas_application == "" {
			return fmt.Errorf("Parameter saas_application is missing")
		}
		if err = d.Set("saas_application", saas_application); err != nil {
			return fmt.Errorf("Error set params saas_application: %v", err)
		}
	}
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application

	o, err := c.ReadObjectCasbProfileSaasApplicationAdvancedTenantControl(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfileSaasApplicationAdvancedTenantControl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbProfileSaasApplicationAdvancedTenantControl(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfileSaasApplicationAdvancedTenantControl resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input"
		if _, ok := i["input"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput3rdl(i["input"], d, pre_append)
			tmp["input"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationAdvancedTenantControl-Attribute-Input")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName3rdl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCasbProfileSaasApplicationAdvancedTenantControl-Attribute-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCasbProfileSaasApplicationAdvancedTenantControl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("attribute", flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute3rdl(o["attribute"], d, "attribute")); err != nil {
			if vv, ok := fortiAPIPatch(o["attribute"], "ObjectCasbProfileSaasApplicationAdvancedTenantControl-Attribute"); ok {
				if err = d.Set("attribute", vv); err != nil {
					return fmt.Errorf("Error reading attribute: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading attribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("attribute"); ok {
			if err = d.Set("attribute", flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute3rdl(o["attribute"], d, "attribute")); err != nil {
				if vv, ok := fortiAPIPatch(o["attribute"], "ObjectCasbProfileSaasApplicationAdvancedTenantControl-Attribute"); ok {
					if err = d.Set("attribute", vv); err != nil {
						return fmt.Errorf("Error reading attribute: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading attribute: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectCasbProfileSaasApplicationAdvancedTenantControlName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCasbProfileSaasApplicationAdvancedTenantControl-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input"], _ = expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput3rdl(d, i["input"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName3rdl(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControlName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCasbProfileSaasApplicationAdvancedTenantControl(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("attribute"); ok || d.HasChange("attribute") {
		t, err := expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute3rdl(d, v, "attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCasbProfileSaasApplicationAdvancedTenantControlName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
