// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure local FortiGuard Web Filter local ratings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWebfilterFtgdLocalRating() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterFtgdLocalRatingCreate,
		Read:   resourceObjectWebfilterFtgdLocalRatingRead,
		Update: resourceObjectWebfilterFtgdLocalRatingUpdate,
		Delete: resourceObjectWebfilterFtgdLocalRatingDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rating": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWebfilterFtgdLocalRatingCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWebfilterFtgdLocalRating(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterFtgdLocalRating resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWebfilterFtgdLocalRating(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterFtgdLocalRating resource: %v", err)
	}

	d.SetId(getStringKey(d, "url"))

	return resourceObjectWebfilterFtgdLocalRatingRead(d, m)
}

func resourceObjectWebfilterFtgdLocalRatingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWebfilterFtgdLocalRating(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterFtgdLocalRating resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWebfilterFtgdLocalRating(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterFtgdLocalRating resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "url"))

	return resourceObjectWebfilterFtgdLocalRatingRead(d, m)
}

func resourceObjectWebfilterFtgdLocalRatingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWebfilterFtgdLocalRating(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterFtgdLocalRating resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterFtgdLocalRatingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWebfilterFtgdLocalRating(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterFtgdLocalRating resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterFtgdLocalRating(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterFtgdLocalRating resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterFtgdLocalRatingComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterFtgdLocalRatingRating(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterFtgdLocalRatingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterFtgdLocalRatingUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebfilterFtgdLocalRating(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comment", flattenObjectWebfilterFtgdLocalRatingComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWebfilterFtgdLocalRating-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("rating", flattenObjectWebfilterFtgdLocalRatingRating(o["rating"], d, "rating")); err != nil {
		if vv, ok := fortiAPIPatch(o["rating"], "ObjectWebfilterFtgdLocalRating-Rating"); ok {
			if err = d.Set("rating", vv); err != nil {
				return fmt.Errorf("Error reading rating: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rating: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectWebfilterFtgdLocalRatingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectWebfilterFtgdLocalRating-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("url", flattenObjectWebfilterFtgdLocalRatingUrl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "ObjectWebfilterFtgdLocalRating-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterFtgdLocalRatingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterFtgdLocalRatingComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterFtgdLocalRatingRating(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectWebfilterFtgdLocalRatingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterFtgdLocalRatingUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebfilterFtgdLocalRating(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectWebfilterFtgdLocalRatingComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("rating"); ok {
		t, err := expandObjectWebfilterFtgdLocalRatingRating(d, v, "rating")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rating"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandObjectWebfilterFtgdLocalRatingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok {
		t, err := expandObjectWebfilterFtgdLocalRatingUrl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	return &obj, nil
}
