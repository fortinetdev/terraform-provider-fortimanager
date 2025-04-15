// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: YouTube channel filter.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebfilterProfileYoutubeChannelFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterProfileYoutubeChannelFilterCreate,
		Read:   resourceObjectWebfilterProfileYoutubeChannelFilterRead,
		Update: resourceObjectWebfilterProfileYoutubeChannelFilterUpdate,
		Delete: resourceObjectWebfilterProfileYoutubeChannelFilterDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"channel_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectWebfilterProfileYoutubeChannelFilterCreate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectWebfilterProfileYoutubeChannelFilter(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterProfileYoutubeChannelFilter resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWebfilterProfileYoutubeChannelFilter(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterProfileYoutubeChannelFilter resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWebfilterProfileYoutubeChannelFilterRead(d, m)
}

func resourceObjectWebfilterProfileYoutubeChannelFilterUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectWebfilterProfileYoutubeChannelFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileYoutubeChannelFilter resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWebfilterProfileYoutubeChannelFilter(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileYoutubeChannelFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWebfilterProfileYoutubeChannelFilterRead(d, m)
}

func resourceObjectWebfilterProfileYoutubeChannelFilterDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectWebfilterProfileYoutubeChannelFilter(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterProfileYoutubeChannelFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterProfileYoutubeChannelFilterRead(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadObjectWebfilterProfileYoutubeChannelFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileYoutubeChannelFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterProfileYoutubeChannelFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileYoutubeChannelFilter resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterProfileYoutubeChannelFilterChannelId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileYoutubeChannelFilterComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileYoutubeChannelFilterId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebfilterProfileYoutubeChannelFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("channel_id", flattenObjectWebfilterProfileYoutubeChannelFilterChannelId2edl(o["channel-id"], d, "channel_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["channel-id"], "ObjectWebfilterProfileYoutubeChannelFilter-ChannelId"); ok {
			if err = d.Set("channel_id", vv); err != nil {
				return fmt.Errorf("Error reading channel_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading channel_id: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectWebfilterProfileYoutubeChannelFilterComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWebfilterProfileYoutubeChannelFilter-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectWebfilterProfileYoutubeChannelFilterId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWebfilterProfileYoutubeChannelFilter-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterProfileYoutubeChannelFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterProfileYoutubeChannelFilterChannelId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileYoutubeChannelFilterComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileYoutubeChannelFilterId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebfilterProfileYoutubeChannelFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("channel_id"); ok || d.HasChange("channel_id") {
		t, err := expandObjectWebfilterProfileYoutubeChannelFilterChannelId2edl(d, v, "channel_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["channel-id"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectWebfilterProfileYoutubeChannelFilterComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWebfilterProfileYoutubeChannelFilterId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
