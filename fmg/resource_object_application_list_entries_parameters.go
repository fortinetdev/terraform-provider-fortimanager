// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Application parameters.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectApplicationListEntriesParameters() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectApplicationListEntriesParametersCreate,
		Read:   resourceObjectApplicationListEntriesParametersRead,
		Update: resourceObjectApplicationListEntriesParametersUpdate,
		Delete: resourceObjectApplicationListEntriesParametersDelete,

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
			"list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"members": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
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
			"value": &schema.Schema{
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

func resourceObjectApplicationListEntriesParametersCreate(d *schema.ResourceData, m interface{}) error {
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

	list := d.Get("list").(string)
	entries := d.Get("entries").(string)
	paradict["list"] = list
	paradict["entries"] = entries

	obj, err := getObjectObjectApplicationListEntriesParameters(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectApplicationListEntriesParameters resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectApplicationListEntriesParameters(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectApplicationListEntriesParameters resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectApplicationListEntriesParametersRead(d, m)
}

func resourceObjectApplicationListEntriesParametersUpdate(d *schema.ResourceData, m interface{}) error {
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

	list := d.Get("list").(string)
	entries := d.Get("entries").(string)
	paradict["list"] = list
	paradict["entries"] = entries

	obj, err := getObjectObjectApplicationListEntriesParameters(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectApplicationListEntriesParameters resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectApplicationListEntriesParameters(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectApplicationListEntriesParameters resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectApplicationListEntriesParametersRead(d, m)
}

func resourceObjectApplicationListEntriesParametersDelete(d *schema.ResourceData, m interface{}) error {
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

	list := d.Get("list").(string)
	entries := d.Get("entries").(string)
	paradict["list"] = list
	paradict["entries"] = entries

	wsParams["adom"] = adomv

	err = c.DeleteObjectApplicationListEntriesParameters(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectApplicationListEntriesParameters resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectApplicationListEntriesParametersRead(d *schema.ResourceData, m interface{}) error {
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

	list := d.Get("list").(string)
	entries := d.Get("entries").(string)
	if list == "" {
		list = importOptionChecking(m.(*FortiClient).Cfg, "list")
		if list == "" {
			return fmt.Errorf("Parameter list is missing")
		}
		if err = d.Set("list", list); err != nil {
			return fmt.Errorf("Error set params list: %v", err)
		}
	}
	if entries == "" {
		entries = importOptionChecking(m.(*FortiClient).Cfg, "entries")
		if entries == "" {
			return fmt.Errorf("Parameter entries is missing")
		}
		if err = d.Set("entries", entries); err != nil {
			return fmt.Errorf("Error set params entries: %v", err)
		}
	}
	paradict["list"] = list
	paradict["entries"] = entries

	o, err := c.ReadObjectApplicationListEntriesParameters(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectApplicationListEntriesParameters resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectApplicationListEntriesParameters(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectApplicationListEntriesParameters resource from API: %v", err)
	}
	return nil
}

func flattenObjectApplicationListEntriesParametersId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParametersMembers3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectApplicationListEntriesParametersMembersId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectApplicationListEntriesParameters-Members-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectApplicationListEntriesParametersMembersName3rdl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectApplicationListEntriesParameters-Members-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectApplicationListEntriesParametersMembersValue3rdl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectApplicationListEntriesParameters-Members-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectApplicationListEntriesParametersMembersId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParametersMembersName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParametersMembersValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParametersValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectApplicationListEntriesParameters(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("fosid", flattenObjectApplicationListEntriesParametersId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectApplicationListEntriesParameters-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("members", flattenObjectApplicationListEntriesParametersMembers3rdl(o["members"], d, "members")); err != nil {
			if vv, ok := fortiAPIPatch(o["members"], "ObjectApplicationListEntriesParameters-Members"); ok {
				if err = d.Set("members", vv); err != nil {
					return fmt.Errorf("Error reading members: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading members: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("members"); ok {
			if err = d.Set("members", flattenObjectApplicationListEntriesParametersMembers3rdl(o["members"], d, "members")); err != nil {
				if vv, ok := fortiAPIPatch(o["members"], "ObjectApplicationListEntriesParameters-Members"); ok {
					if err = d.Set("members", vv); err != nil {
						return fmt.Errorf("Error reading members: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading members: %v", err)
				}
			}
		}
	}

	if err = d.Set("value", flattenObjectApplicationListEntriesParametersValue3rdl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "ObjectApplicationListEntriesParameters-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenObjectApplicationListEntriesParametersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectApplicationListEntriesParametersId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParametersMembers3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandObjectApplicationListEntriesParametersMembersId3rdl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectApplicationListEntriesParametersMembersName3rdl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectApplicationListEntriesParametersMembersValue3rdl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectApplicationListEntriesParametersMembersId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParametersMembersName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParametersMembersValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParametersValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectApplicationListEntriesParameters(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectApplicationListEntriesParametersId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("members"); ok || d.HasChange("members") {
		t, err := expandObjectApplicationListEntriesParametersMembers3rdl(d, v, "members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["members"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandObjectApplicationListEntriesParametersValue3rdl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
