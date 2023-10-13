// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure AWS external account list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemSdnConnectorExternalAccountList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemSdnConnectorExternalAccountListCreate,
		Read:   resourceObjectSystemSdnConnectorExternalAccountListRead,
		Update: resourceObjectSystemSdnConnectorExternalAccountListUpdate,
		Delete: resourceObjectSystemSdnConnectorExternalAccountListDelete,

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
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"region_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"role_arn": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemSdnConnectorExternalAccountListCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemSdnConnectorExternalAccountList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnectorExternalAccountList resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemSdnConnectorExternalAccountList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnectorExternalAccountList resource: %v", err)
	}

	d.SetId(getStringKey(d, "role_arn"))

	return resourceObjectSystemSdnConnectorExternalAccountListRead(d, m)
}

func resourceObjectSystemSdnConnectorExternalAccountListUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemSdnConnectorExternalAccountList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnectorExternalAccountList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemSdnConnectorExternalAccountList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnectorExternalAccountList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "role_arn"))

	return resourceObjectSystemSdnConnectorExternalAccountListRead(d, m)
}

func resourceObjectSystemSdnConnectorExternalAccountListDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemSdnConnectorExternalAccountList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemSdnConnectorExternalAccountList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemSdnConnectorExternalAccountListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemSdnConnectorExternalAccountList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnectorExternalAccountList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemSdnConnectorExternalAccountList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnectorExternalAccountList resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemSdnConnectorExternalAccountListExternalId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorExternalAccountListRegionList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemSdnConnectorExternalAccountListRoleArn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemSdnConnectorExternalAccountList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("external_id", flattenObjectSystemSdnConnectorExternalAccountListExternalId2edl(o["external-id"], d, "external_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-id"], "ObjectSystemSdnConnectorExternalAccountList-ExternalId"); ok {
			if err = d.Set("external_id", vv); err != nil {
				return fmt.Errorf("Error reading external_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_id: %v", err)
		}
	}

	if err = d.Set("region_list", flattenObjectSystemSdnConnectorExternalAccountListRegionList2edl(o["region-list"], d, "region_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["region-list"], "ObjectSystemSdnConnectorExternalAccountList-RegionList"); ok {
			if err = d.Set("region_list", vv); err != nil {
				return fmt.Errorf("Error reading region_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading region_list: %v", err)
		}
	}

	if err = d.Set("role_arn", flattenObjectSystemSdnConnectorExternalAccountListRoleArn2edl(o["role-arn"], d, "role_arn")); err != nil {
		if vv, ok := fortiAPIPatch(o["role-arn"], "ObjectSystemSdnConnectorExternalAccountList-RoleArn"); ok {
			if err = d.Set("role_arn", vv); err != nil {
				return fmt.Errorf("Error reading role_arn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role_arn: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemSdnConnectorExternalAccountListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemSdnConnectorExternalAccountListExternalId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorExternalAccountListRegionList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemSdnConnectorExternalAccountListRoleArn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemSdnConnectorExternalAccountList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("external_id"); ok || d.HasChange("external_id") {
		t, err := expandObjectSystemSdnConnectorExternalAccountListExternalId2edl(d, v, "external_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-id"] = t
		}
	}

	if v, ok := d.GetOk("region_list"); ok || d.HasChange("region_list") {
		t, err := expandObjectSystemSdnConnectorExternalAccountListRegionList2edl(d, v, "region_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region-list"] = t
		}
	}

	if v, ok := d.GetOk("role_arn"); ok || d.HasChange("role_arn") {
		t, err := expandObjectSystemSdnConnectorExternalAccountListRoleArn2edl(d, v, "role_arn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role-arn"] = t
		}
	}

	return &obj, nil
}
