// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Parameter tuple members.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectApplicationListEntriesParametersMembers() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectApplicationListEntriesParametersMembersCreate,
		Read:   resourceObjectApplicationListEntriesParametersMembersRead,
		Update: resourceObjectApplicationListEntriesParametersMembersUpdate,
		Delete: resourceObjectApplicationListEntriesParametersMembersDelete,

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
			"parameters": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
	}
}

func resourceObjectApplicationListEntriesParametersMembersCreate(d *schema.ResourceData, m interface{}) error {
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
	parameters := d.Get("parameters").(string)
	paradict["list"] = list
	paradict["entries"] = entries
	paradict["parameters"] = parameters

	obj, err := getObjectObjectApplicationListEntriesParametersMembers(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectApplicationListEntriesParametersMembers resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectApplicationListEntriesParametersMembers(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectApplicationListEntriesParametersMembers resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectApplicationListEntriesParametersMembersRead(d, m)
}

func resourceObjectApplicationListEntriesParametersMembersUpdate(d *schema.ResourceData, m interface{}) error {
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
	parameters := d.Get("parameters").(string)
	paradict["list"] = list
	paradict["entries"] = entries
	paradict["parameters"] = parameters

	obj, err := getObjectObjectApplicationListEntriesParametersMembers(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectApplicationListEntriesParametersMembers resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectApplicationListEntriesParametersMembers(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectApplicationListEntriesParametersMembers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectApplicationListEntriesParametersMembersRead(d, m)
}

func resourceObjectApplicationListEntriesParametersMembersDelete(d *schema.ResourceData, m interface{}) error {
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
	parameters := d.Get("parameters").(string)
	paradict["list"] = list
	paradict["entries"] = entries
	paradict["parameters"] = parameters

	wsParams["adom"] = adomv

	err = c.DeleteObjectApplicationListEntriesParametersMembers(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectApplicationListEntriesParametersMembers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectApplicationListEntriesParametersMembersRead(d *schema.ResourceData, m interface{}) error {
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
	parameters := d.Get("parameters").(string)
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
	if parameters == "" {
		parameters = importOptionChecking(m.(*FortiClient).Cfg, "parameters")
		if parameters == "" {
			return fmt.Errorf("Parameter parameters is missing")
		}
		if err = d.Set("parameters", parameters); err != nil {
			return fmt.Errorf("Error set params parameters: %v", err)
		}
	}
	paradict["list"] = list
	paradict["entries"] = entries
	paradict["parameters"] = parameters

	o, err := c.ReadObjectApplicationListEntriesParametersMembers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectApplicationListEntriesParametersMembers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectApplicationListEntriesParametersMembers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectApplicationListEntriesParametersMembers resource from API: %v", err)
	}
	return nil
}

func flattenObjectApplicationListEntriesParametersMembersId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParametersMembersName4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParametersMembersValue4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectApplicationListEntriesParametersMembers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenObjectApplicationListEntriesParametersMembersId4thl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectApplicationListEntriesParametersMembers-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectApplicationListEntriesParametersMembersName4thl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectApplicationListEntriesParametersMembers-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("value", flattenObjectApplicationListEntriesParametersMembersValue4thl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "ObjectApplicationListEntriesParametersMembers-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenObjectApplicationListEntriesParametersMembersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectApplicationListEntriesParametersMembersId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParametersMembersName4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParametersMembersValue4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectApplicationListEntriesParametersMembers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectApplicationListEntriesParametersMembersId4thl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectApplicationListEntriesParametersMembersName4thl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandObjectApplicationListEntriesParametersMembersValue4thl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
