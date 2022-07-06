// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure multilayer mode.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFmupdateMultilayer() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateMultilayerUpdate,
		Read:   resourceFmupdateMultilayerRead,
		Update: resourceFmupdateMultilayerUpdate,
		Delete: resourceFmupdateMultilayerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"webspam_rating": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateMultilayerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateMultilayer(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateMultilayer resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateMultilayer(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateMultilayer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateMultilayer")

	return resourceFmupdateMultilayerRead(d, m)
}

func resourceFmupdateMultilayerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateMultilayer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateMultilayer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateMultilayerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateMultilayer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateMultilayer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateMultilayer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateMultilayer resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateMultilayerWebspamRating(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateMultilayer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("webspam_rating", flattenFmupdateMultilayerWebspamRating(o["webspam-rating"], d, "webspam_rating")); err != nil {
		if vv, ok := fortiAPIPatch(o["webspam-rating"], "FmupdateMultilayer-WebspamRating"); ok {
			if err = d.Set("webspam_rating", vv); err != nil {
				return fmt.Errorf("Error reading webspam_rating: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webspam_rating: %v", err)
		}
	}

	return nil
}

func flattenFmupdateMultilayerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateMultilayerWebspamRating(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateMultilayer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("webspam_rating"); ok || d.HasChange("webspam_rating") {
		t, err := expandFmupdateMultilayerWebspamRating(d, v, "webspam_rating")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webspam-rating"] = t
		}
	}

	return &obj, nil
}
