// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure OCI compartment list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemSdnConnectorCompartmentList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemSdnConnectorCompartmentListCreate,
		Read:   resourceObjectSystemSdnConnectorCompartmentListRead,
		Update: resourceObjectSystemSdnConnectorCompartmentListUpdate,
		Delete: resourceObjectSystemSdnConnectorCompartmentListDelete,

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
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemSdnConnectorCompartmentListCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemSdnConnectorCompartmentList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnectorCompartmentList resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemSdnConnectorCompartmentList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnectorCompartmentList resource: %v", err)
	}

	d.SetId(getStringKey(d, "compartment_id"))

	return resourceObjectSystemSdnConnectorCompartmentListRead(d, m)
}

func resourceObjectSystemSdnConnectorCompartmentListUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemSdnConnectorCompartmentList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnectorCompartmentList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemSdnConnectorCompartmentList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnectorCompartmentList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "compartment_id"))

	return resourceObjectSystemSdnConnectorCompartmentListRead(d, m)
}

func resourceObjectSystemSdnConnectorCompartmentListDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemSdnConnectorCompartmentList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemSdnConnectorCompartmentList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemSdnConnectorCompartmentListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemSdnConnectorCompartmentList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnectorCompartmentList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemSdnConnectorCompartmentList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnectorCompartmentList resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemSdnConnectorCompartmentListCompartmentId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemSdnConnectorCompartmentList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("compartment_id", flattenObjectSystemSdnConnectorCompartmentListCompartmentId2edl(o["compartment-id"], d, "compartment_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["compartment-id"], "ObjectSystemSdnConnectorCompartmentList-CompartmentId"); ok {
			if err = d.Set("compartment_id", vv); err != nil {
				return fmt.Errorf("Error reading compartment_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading compartment_id: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemSdnConnectorCompartmentListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemSdnConnectorCompartmentListCompartmentId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemSdnConnectorCompartmentList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("compartment_id"); ok || d.HasChange("compartment_id") {
		t, err := expandObjectSystemSdnConnectorCompartmentListCompartmentId2edl(d, v, "compartment_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compartment-id"] = t
		}
	}

	return &obj, nil
}
