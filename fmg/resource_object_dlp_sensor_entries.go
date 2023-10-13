// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DLP sensor entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDlpSensorEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDlpSensorEntriesCreate,
		Read:   resourceObjectDlpSensorEntriesRead,
		Update: resourceObjectDlpSensorEntriesUpdate,
		Delete: resourceObjectDlpSensorEntriesDelete,

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
			"sensor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fmgcount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dictionary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectDlpSensorEntriesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	sensor := d.Get("sensor").(string)
	paradict["sensor"] = sensor

	obj, err := getObjectObjectDlpSensorEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpSensorEntries resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDlpSensorEntries(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpSensorEntries resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDlpSensorEntriesRead(d, m)
}

func resourceObjectDlpSensorEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	paradict["sensor"] = sensor

	obj, err := getObjectObjectDlpSensorEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpSensorEntries resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDlpSensorEntries(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpSensorEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDlpSensorEntriesRead(d, m)
}

func resourceObjectDlpSensorEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	paradict["sensor"] = sensor

	err = c.DeleteObjectDlpSensorEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDlpSensorEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDlpSensorEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	if sensor == "" {
		sensor = importOptionChecking(m.(*FortiClient).Cfg, "sensor")
		if sensor == "" {
			return fmt.Errorf("Parameter sensor is missing")
		}
		if err = d.Set("sensor", sensor); err != nil {
			return fmt.Errorf("Error set params sensor: %v", err)
		}
	}
	paradict["sensor"] = sensor

	o, err := c.ReadObjectDlpSensorEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpSensorEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDlpSensorEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpSensorEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectDlpSensorEntriesCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorEntriesDictionary2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectDlpSensorEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDlpSensorEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fmgcount", flattenObjectDlpSensorEntriesCount2edl(o["count"], d, "fmgcount")); err != nil {
		if vv, ok := fortiAPIPatch(o["count"], "ObjectDlpSensorEntries-Count"); ok {
			if err = d.Set("fmgcount", vv); err != nil {
				return fmt.Errorf("Error reading fmgcount: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgcount: %v", err)
		}
	}

	if err = d.Set("dictionary", flattenObjectDlpSensorEntriesDictionary2edl(o["dictionary"], d, "dictionary")); err != nil {
		if vv, ok := fortiAPIPatch(o["dictionary"], "ObjectDlpSensorEntries-Dictionary"); ok {
			if err = d.Set("dictionary", vv); err != nil {
				return fmt.Errorf("Error reading dictionary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dictionary: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectDlpSensorEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectDlpSensorEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectDlpSensorEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectDlpSensorEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectDlpSensorEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDlpSensorEntriesCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorEntriesDictionary2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDlpSensorEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgcount"); ok || d.HasChange("fmgcount") {
		t, err := expandObjectDlpSensorEntriesCount2edl(d, v, "fmgcount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["count"] = t
		}
	}

	if v, ok := d.GetOk("dictionary"); ok || d.HasChange("dictionary") {
		t, err := expandObjectDlpSensorEntriesDictionary2edl(d, v, "dictionary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dictionary"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectDlpSensorEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectDlpSensorEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
