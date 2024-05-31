// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectCloud OrchestAwsconnector

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCloudOrchestAwsconnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCloudOrchestAwsconnectorCreate,
		Read:   resourceObjectCloudOrchestAwsconnectorRead,
		Update: resourceObjectCloudOrchestAwsconnectorUpdate,
		Delete: resourceObjectCloudOrchestAwsconnectorDelete,

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
			"access_key_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"access_key_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"use_metadata_iam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectCloudOrchestAwsconnectorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectCloudOrchestAwsconnector(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCloudOrchestAwsconnector resource while getting object: %v", err)
	}

	_, err = c.CreateObjectCloudOrchestAwsconnector(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectCloudOrchestAwsconnector resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCloudOrchestAwsconnectorRead(d, m)
}

func resourceObjectCloudOrchestAwsconnectorUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCloudOrchestAwsconnector(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCloudOrchestAwsconnector resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectCloudOrchestAwsconnector(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCloudOrchestAwsconnector resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCloudOrchestAwsconnectorRead(d, m)
}

func resourceObjectCloudOrchestAwsconnectorDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectCloudOrchestAwsconnector(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCloudOrchestAwsconnector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCloudOrchestAwsconnectorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectCloudOrchestAwsconnector(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCloudOrchestAwsconnector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCloudOrchestAwsconnector(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCloudOrchestAwsconnector resource from API: %v", err)
	}
	return nil
}

func flattenObjectCloudOrchestAwsconnectorAccessKeyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwsconnectorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestAwsconnectorUseMetadataIam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCloudOrchestAwsconnector(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("access_key_id", flattenObjectCloudOrchestAwsconnectorAccessKeyId(o["access-key-id"], d, "access_key_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-key-id"], "ObjectCloudOrchestAwsconnector-AccessKeyId"); ok {
			if err = d.Set("access_key_id", vv); err != nil {
				return fmt.Errorf("Error reading access_key_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_key_id: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCloudOrchestAwsconnectorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCloudOrchestAwsconnector-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("use_metadata_iam", flattenObjectCloudOrchestAwsconnectorUseMetadataIam(o["use-metadata-iam"], d, "use_metadata_iam")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-metadata-iam"], "ObjectCloudOrchestAwsconnector-UseMetadataIam"); ok {
			if err = d.Set("use_metadata_iam", vv); err != nil {
				return fmt.Errorf("Error reading use_metadata_iam: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_metadata_iam: %v", err)
		}
	}

	return nil
}

func flattenObjectCloudOrchestAwsconnectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCloudOrchestAwsconnectorAccessKeyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwsconnectorAccessKeySecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCloudOrchestAwsconnectorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestAwsconnectorUseMetadataIam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCloudOrchestAwsconnector(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_key_id"); ok || d.HasChange("access_key_id") {
		t, err := expandObjectCloudOrchestAwsconnectorAccessKeyId(d, v, "access_key_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-key-id"] = t
		}
	}

	if v, ok := d.GetOk("access_key_secret"); ok || d.HasChange("access_key_secret") {
		t, err := expandObjectCloudOrchestAwsconnectorAccessKeySecret(d, v, "access_key_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-key-secret"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCloudOrchestAwsconnectorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("use_metadata_iam"); ok || d.HasChange("use_metadata_iam") {
		t, err := expandObjectCloudOrchestAwsconnectorUseMetadataIam(d, v, "use_metadata_iam")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-metadata-iam"] = t
		}
	}

	return &obj, nil
}
