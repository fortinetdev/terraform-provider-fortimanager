// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectCloud OrchestAws

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCloudOrchestAws() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCloudOrchestAwsCreate,
		Read:   resourceObjectCloudOrchestAwsRead,
		Update: resourceObjectCloudOrchestAwsUpdate,
		Delete: resourceObjectCloudOrchestAwsDelete,

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
			"connector": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"region_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"template_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectCloudOrchestAwsCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCloudOrchestAws(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCloudOrchestAws resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectCloudOrchestAws(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCloudOrchestAws resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCloudOrchestAwsRead(d, m)
}

func resourceObjectCloudOrchestAwsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCloudOrchestAws(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCloudOrchestAws resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCloudOrchestAws(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCloudOrchestAws resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCloudOrchestAwsRead(d, m)
}

func resourceObjectCloudOrchestAwsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectCloudOrchestAws(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCloudOrchestAws resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCloudOrchestAwsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectCloudOrchestAws(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCloudOrchestAws resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCloudOrchestAws(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCloudOrchestAws resource from API: %v", err)
	}
	return nil
}

func flattenObjectCloudOrchestAwsConnector(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectCloudOrchestAwsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwsRegionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwsTemplateConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectObjectCloudOrchestAws(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("connector", flattenObjectCloudOrchestAwsConnector(o["connector"], d, "connector")); err != nil {
		if vv, ok := fortiAPIPatch(o["connector"], "ObjectCloudOrchestAws-Connector"); ok {
			if err = d.Set("connector", vv); err != nil {
				return fmt.Errorf("Error reading connector: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading connector: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCloudOrchestAwsName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCloudOrchestAws-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("region_name", flattenObjectCloudOrchestAwsRegionName(o["region-name"], d, "region_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["region-name"], "ObjectCloudOrchestAws-RegionName"); ok {
			if err = d.Set("region_name", vv); err != nil {
				return fmt.Errorf("Error reading region_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading region_name: %v", err)
		}
	}

	if err = d.Set("template_configuration", flattenObjectCloudOrchestAwsTemplateConfiguration(o["template-configuration"], d, "template_configuration")); err != nil {
		if vv, ok := fortiAPIPatch(o["template-configuration"], "ObjectCloudOrchestAws-TemplateConfiguration"); ok {
			if err = d.Set("template_configuration", vv); err != nil {
				return fmt.Errorf("Error reading template_configuration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading template_configuration: %v", err)
		}
	}

	return nil
}

func flattenObjectCloudOrchestAwsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCloudOrchestAwsConnector(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectCloudOrchestAwsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwsRegionName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwsTemplateConfiguration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectObjectCloudOrchestAws(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("connector"); ok || d.HasChange("connector") {
		t, err := expandObjectCloudOrchestAwsConnector(d, v, "connector")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connector"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCloudOrchestAwsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("region_name"); ok || d.HasChange("region_name") {
		t, err := expandObjectCloudOrchestAwsRegionName(d, v, "region_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region-name"] = t
		}
	}

	if v, ok := d.GetOk("template_configuration"); ok || d.HasChange("template_configuration") {
		t, err := expandObjectCloudOrchestAwsTemplateConfiguration(d, v, "template_configuration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["template-configuration"] = t
		}
	}

	return &obj, nil
}
