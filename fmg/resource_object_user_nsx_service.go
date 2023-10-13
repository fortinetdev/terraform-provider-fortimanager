// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectUser NsxService

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserNsxService() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserNsxServiceCreate,
		Read:   resourceObjectUserNsxServiceRead,
		Update: resourceObjectUserNsxServiceUpdate,
		Delete: resourceObjectUserNsxServiceDelete,

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
			"nsx": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"integration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ref_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectUserNsxServiceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	nsx := d.Get("nsx").(string)
	paradict["nsx"] = nsx

	obj, err := getObjectObjectUserNsxService(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserNsxService resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserNsxService(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserNsxService resource: %v", err)
	}

	d.SetId(getStringKey(d, "fosid"))

	return resourceObjectUserNsxServiceRead(d, m)
}

func resourceObjectUserNsxServiceUpdate(d *schema.ResourceData, m interface{}) error {
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

	nsx := d.Get("nsx").(string)
	paradict["nsx"] = nsx

	obj, err := getObjectObjectUserNsxService(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserNsxService resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserNsxService(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserNsxService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "fosid"))

	return resourceObjectUserNsxServiceRead(d, m)
}

func resourceObjectUserNsxServiceDelete(d *schema.ResourceData, m interface{}) error {
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

	nsx := d.Get("nsx").(string)
	paradict["nsx"] = nsx

	err = c.DeleteObjectUserNsxService(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserNsxService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserNsxServiceRead(d *schema.ResourceData, m interface{}) error {
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

	nsx := d.Get("nsx").(string)
	if nsx == "" {
		nsx = importOptionChecking(m.(*FortiClient).Cfg, "nsx")
		if nsx == "" {
			return fmt.Errorf("Parameter nsx is missing")
		}
		if err = d.Set("nsx", nsx); err != nil {
			return fmt.Errorf("Error set params nsx: %v", err)
		}
	}
	paradict["nsx"] = nsx

	o, err := c.ReadObjectUserNsxService(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserNsxService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserNsxService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserNsxService resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserNsxServiceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxServiceIntegration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxServiceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxServiceRefId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserNsxService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenObjectUserNsxServiceId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectUserNsxService-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("integration", flattenObjectUserNsxServiceIntegration2edl(o["integration"], d, "integration")); err != nil {
		if vv, ok := fortiAPIPatch(o["integration"], "ObjectUserNsxService-Integration"); ok {
			if err = d.Set("integration", vv); err != nil {
				return fmt.Errorf("Error reading integration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading integration: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserNsxServiceName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserNsxService-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ref_id", flattenObjectUserNsxServiceRefId2edl(o["ref-id"], d, "ref_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["ref-id"], "ObjectUserNsxService-RefId"); ok {
			if err = d.Set("ref_id", vv); err != nil {
				return fmt.Errorf("Error reading ref_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ref_id: %v", err)
		}
	}

	return nil
}

func flattenObjectUserNsxServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserNsxServiceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxServiceIntegration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxServiceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxServiceRefId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserNsxService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectUserNsxServiceId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("integration"); ok || d.HasChange("integration") {
		t, err := expandObjectUserNsxServiceIntegration2edl(d, v, "integration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["integration"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserNsxServiceName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ref_id"); ok || d.HasChange("ref_id") {
		t, err := expandObjectUserNsxServiceRefId2edl(d, v, "ref_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ref-id"] = t
		}
	}

	return &obj, nil
}
