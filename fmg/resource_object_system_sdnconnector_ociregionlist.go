// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure OCI region list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemSdnConnectorOciRegionList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemSdnConnectorOciRegionListCreate,
		Read:   resourceObjectSystemSdnConnectorOciRegionListRead,
		Update: resourceObjectSystemSdnConnectorOciRegionListUpdate,
		Delete: resourceObjectSystemSdnConnectorOciRegionListDelete,

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
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemSdnConnectorOciRegionListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	sdn_connector := d.Get("sdn_connector").(string)
	paradict["sdn_connector"] = sdn_connector

	obj, err := getObjectObjectSystemSdnConnectorOciRegionList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnectorOciRegionList resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemSdnConnectorOciRegionList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnectorOciRegionList resource: %v", err)
	}

	d.SetId(getStringKey(d, "region"))

	return resourceObjectSystemSdnConnectorOciRegionListRead(d, m)
}

func resourceObjectSystemSdnConnectorOciRegionListUpdate(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	paradict["sdn_connector"] = sdn_connector

	obj, err := getObjectObjectSystemSdnConnectorOciRegionList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnectorOciRegionList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemSdnConnectorOciRegionList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnectorOciRegionList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "region"))

	return resourceObjectSystemSdnConnectorOciRegionListRead(d, m)
}

func resourceObjectSystemSdnConnectorOciRegionListDelete(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	paradict["sdn_connector"] = sdn_connector

	err = c.DeleteObjectSystemSdnConnectorOciRegionList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemSdnConnectorOciRegionList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemSdnConnectorOciRegionListRead(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	if sdn_connector == "" {
		sdn_connector = importOptionChecking(m.(*FortiClient).Cfg, "sdn_connector")
		if sdn_connector == "" {
			return fmt.Errorf("Parameter sdn_connector is missing")
		}
		if err = d.Set("sdn_connector", sdn_connector); err != nil {
			return fmt.Errorf("Error set params sdn_connector: %v", err)
		}
	}
	paradict["sdn_connector"] = sdn_connector

	o, err := c.ReadObjectSystemSdnConnectorOciRegionList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnectorOciRegionList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemSdnConnectorOciRegionList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnectorOciRegionList resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemSdnConnectorOciRegionListRegion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemSdnConnectorOciRegionList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("region", flattenObjectSystemSdnConnectorOciRegionListRegion2edl(o["region"], d, "region")); err != nil {
		if vv, ok := fortiAPIPatch(o["region"], "ObjectSystemSdnConnectorOciRegionList-Region"); ok {
			if err = d.Set("region", vv); err != nil {
				return fmt.Errorf("Error reading region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading region: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemSdnConnectorOciRegionListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemSdnConnectorOciRegionListRegion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemSdnConnectorOciRegionList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("region"); ok || d.HasChange("region") {
		t, err := expandObjectSystemSdnConnectorOciRegionListRegion2edl(d, v, "region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region"] = t
		}
	}

	return &obj, nil
}
