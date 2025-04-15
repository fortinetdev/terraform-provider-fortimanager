// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure GCP project list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemSdnConnectorGcpProjectList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemSdnConnectorGcpProjectListCreate,
		Read:   resourceObjectSystemSdnConnectorGcpProjectListRead,
		Update: resourceObjectSystemSdnConnectorGcpProjectListUpdate,
		Delete: resourceObjectSystemSdnConnectorGcpProjectListDelete,

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
			"gcp_zone_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemSdnConnectorGcpProjectListCreate(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	paradict["sdn_connector"] = sdn_connector

	obj, err := getObjectObjectSystemSdnConnectorGcpProjectList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnectorGcpProjectList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSystemSdnConnectorGcpProjectList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnectorGcpProjectList resource: %v", err)
	}

	d.SetId(getStringKey(d, "fosid"))

	return resourceObjectSystemSdnConnectorGcpProjectListRead(d, m)
}

func resourceObjectSystemSdnConnectorGcpProjectListUpdate(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	paradict["sdn_connector"] = sdn_connector

	obj, err := getObjectObjectSystemSdnConnectorGcpProjectList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnectorGcpProjectList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemSdnConnectorGcpProjectList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnectorGcpProjectList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "fosid"))

	return resourceObjectSystemSdnConnectorGcpProjectListRead(d, m)
}

func resourceObjectSystemSdnConnectorGcpProjectListDelete(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	paradict["sdn_connector"] = sdn_connector

	wsParams["adom"] = adomv

	err = c.DeleteObjectSystemSdnConnectorGcpProjectList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemSdnConnectorGcpProjectList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemSdnConnectorGcpProjectListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemSdnConnectorGcpProjectList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnectorGcpProjectList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemSdnConnectorGcpProjectList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnectorGcpProjectList resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemSdnConnectorGcpProjectListGcpZoneList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemSdnConnectorGcpProjectListId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemSdnConnectorGcpProjectList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("gcp_zone_list", flattenObjectSystemSdnConnectorGcpProjectListGcpZoneList2edl(o["gcp-zone-list"], d, "gcp_zone_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["gcp-zone-list"], "ObjectSystemSdnConnectorGcpProjectList-GcpZoneList"); ok {
			if err = d.Set("gcp_zone_list", vv); err != nil {
				return fmt.Errorf("Error reading gcp_zone_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gcp_zone_list: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectSystemSdnConnectorGcpProjectListId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSystemSdnConnectorGcpProjectList-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemSdnConnectorGcpProjectListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemSdnConnectorGcpProjectListGcpZoneList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemSdnConnectorGcpProjectListId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemSdnConnectorGcpProjectList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("gcp_zone_list"); ok || d.HasChange("gcp_zone_list") {
		t, err := expandObjectSystemSdnConnectorGcpProjectListGcpZoneList2edl(d, v, "gcp_zone_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gcp-zone-list"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectSystemSdnConnectorGcpProjectListId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
