// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectFirewall CasbProfileSaasApplicationCustomControl

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallCasbProfileSaasApplicationCustomControl() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallCasbProfileSaasApplicationCustomControlCreate,
		Read:   resourceObjectFirewallCasbProfileSaasApplicationCustomControlRead,
		Update: resourceObjectFirewallCasbProfileSaasApplicationCustomControlUpdate,
		Delete: resourceObjectFirewallCasbProfileSaasApplicationCustomControlDelete,

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
			"casb_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"saas_application": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"option": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_input": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceObjectFirewallCasbProfileSaasApplicationCustomControlCreate(d *schema.ResourceData, m interface{}) error {
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

	casb_profile := d.Get("casb_profile").(string)
	saas_application := d.Get("saas_application").(string)
	paradict["casb_profile"] = casb_profile
	paradict["saas_application"] = saas_application

	obj, err := getObjectObjectFirewallCasbProfileSaasApplicationCustomControl(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallCasbProfileSaasApplicationCustomControl resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFirewallCasbProfileSaasApplicationCustomControl(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallCasbProfileSaasApplicationCustomControl resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallCasbProfileSaasApplicationCustomControlRead(d, m)
}

func resourceObjectFirewallCasbProfileSaasApplicationCustomControlUpdate(d *schema.ResourceData, m interface{}) error {
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

	casb_profile := d.Get("casb_profile").(string)
	saas_application := d.Get("saas_application").(string)
	paradict["casb_profile"] = casb_profile
	paradict["saas_application"] = saas_application

	obj, err := getObjectObjectFirewallCasbProfileSaasApplicationCustomControl(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallCasbProfileSaasApplicationCustomControl resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallCasbProfileSaasApplicationCustomControl(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallCasbProfileSaasApplicationCustomControl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallCasbProfileSaasApplicationCustomControlRead(d, m)
}

func resourceObjectFirewallCasbProfileSaasApplicationCustomControlDelete(d *schema.ResourceData, m interface{}) error {
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

	casb_profile := d.Get("casb_profile").(string)
	saas_application := d.Get("saas_application").(string)
	paradict["casb_profile"] = casb_profile
	paradict["saas_application"] = saas_application

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallCasbProfileSaasApplicationCustomControl(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallCasbProfileSaasApplicationCustomControl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallCasbProfileSaasApplicationCustomControlRead(d *schema.ResourceData, m interface{}) error {
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

	casb_profile := d.Get("casb_profile").(string)
	saas_application := d.Get("saas_application").(string)
	if casb_profile == "" {
		casb_profile = importOptionChecking(m.(*FortiClient).Cfg, "casb_profile")
		if casb_profile == "" {
			return fmt.Errorf("Parameter casb_profile is missing")
		}
		if err = d.Set("casb_profile", casb_profile); err != nil {
			return fmt.Errorf("Error set params casb_profile: %v", err)
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
	paradict["casb_profile"] = casb_profile
	paradict["saas_application"] = saas_application

	o, err := c.ReadObjectFirewallCasbProfileSaasApplicationCustomControl(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallCasbProfileSaasApplicationCustomControl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallCasbProfileSaasApplicationCustomControl(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallCasbProfileSaasApplicationCustomControl resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlOption3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionName3rdl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplicationCustomControl-Option-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := i["user-input"]; ok {
			v := flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput3rdl(i["user-input"], d, pre_append)
			tmp["user_input"] = fortiAPISubPartPatch(v, "ObjectFirewallCasbProfileSaasApplicationCustomControl-Option-UserInput")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectFirewallCasbProfileSaasApplicationCustomControl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("name", flattenObjectFirewallCasbProfileSaasApplicationCustomControlName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallCasbProfileSaasApplicationCustomControl-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("option", flattenObjectFirewallCasbProfileSaasApplicationCustomControlOption3rdl(o["option"], d, "option")); err != nil {
			if vv, ok := fortiAPIPatch(o["option"], "ObjectFirewallCasbProfileSaasApplicationCustomControl-Option"); ok {
				if err = d.Set("option", vv); err != nil {
					return fmt.Errorf("Error reading option: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading option: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("option"); ok {
			if err = d.Set("option", flattenObjectFirewallCasbProfileSaasApplicationCustomControlOption3rdl(o["option"], d, "option")); err != nil {
				if vv, ok := fortiAPIPatch(o["option"], "ObjectFirewallCasbProfileSaasApplicationCustomControl-Option"); ok {
					if err = d.Set("option", vv); err != nil {
						return fmt.Errorf("Error reading option: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading option: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControlName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControlOption3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionName3rdl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-input"], _ = expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput3rdl(d, i["user_input"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectFirewallCasbProfileSaasApplicationCustomControl(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallCasbProfileSaasApplicationCustomControlName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandObjectFirewallCasbProfileSaasApplicationCustomControlOption3rdl(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	return &obj, nil
}
