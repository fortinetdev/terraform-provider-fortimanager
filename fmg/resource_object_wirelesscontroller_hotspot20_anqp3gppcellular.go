// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure 3GPP public land mobile network (PLMN).

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWirelessControllerHotspot20Anqp3GppCellular() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20Anqp3GppCellularCreate,
		Read:   resourceObjectWirelessControllerHotspot20Anqp3GppCellularRead,
		Update: resourceObjectWirelessControllerHotspot20Anqp3GppCellularUpdate,
		Delete: resourceObjectWirelessControllerHotspot20Anqp3GppCellularDelete,

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
			"mcc_mnc_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mcc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mnc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20Anqp3GppCellularCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20Anqp3GppCellular(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20Anqp3GppCellular resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerHotspot20Anqp3GppCellular(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20Anqp3GppCellular resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20Anqp3GppCellularRead(d, m)
}

func resourceObjectWirelessControllerHotspot20Anqp3GppCellularUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20Anqp3GppCellular(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20Anqp3GppCellular resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerHotspot20Anqp3GppCellular(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20Anqp3GppCellular resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20Anqp3GppCellularRead(d, m)
}

func resourceObjectWirelessControllerHotspot20Anqp3GppCellularDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerHotspot20Anqp3GppCellular(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20Anqp3GppCellular resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20Anqp3GppCellularRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerHotspot20Anqp3GppCellular(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20Anqp3GppCellular resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20Anqp3GppCellular(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20Anqp3GppCellular resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20Anqp3GppCellular-MccMncList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcc"
		if _, ok := i["mcc"]; ok {
			v := flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc(i["mcc"], d, pre_append)
			tmp["mcc"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20Anqp3GppCellular-MccMncList-Mcc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mnc"
		if _, ok := i["mnc"]; ok {
			v := flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc(i["mnc"], d, pre_append)
			tmp["mnc"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20Anqp3GppCellular-MccMncList-Mnc")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20Anqp3GppCellularName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20Anqp3GppCellular(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("mcc_mnc_list", flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(o["mcc-mnc-list"], d, "mcc_mnc_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["mcc-mnc-list"], "ObjectWirelessControllerHotspot20Anqp3GppCellular-MccMncList"); ok {
				if err = d.Set("mcc_mnc_list", vv); err != nil {
					return fmt.Errorf("Error reading mcc_mnc_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mcc_mnc_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mcc_mnc_list"); ok {
			if err = d.Set("mcc_mnc_list", flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(o["mcc-mnc-list"], d, "mcc_mnc_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["mcc-mnc-list"], "ObjectWirelessControllerHotspot20Anqp3GppCellular-MccMncList"); ok {
					if err = d.Set("mcc_mnc_list", vv); err != nil {
						return fmt.Errorf("Error reading mcc_mnc_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mcc_mnc_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerHotspot20Anqp3GppCellularName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerHotspot20Anqp3GppCellular-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20Anqp3GppCellularFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcc"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mcc"], _ = expandObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc(d, i["mcc"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mnc"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mnc"], _ = expandObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc(d, i["mnc"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20Anqp3GppCellularName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20Anqp3GppCellular(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mcc_mnc_list"); ok || d.HasChange("mcc_mnc_list") {
		t, err := expandObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(d, v, "mcc_mnc_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcc-mnc-list"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerHotspot20Anqp3GppCellularName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
