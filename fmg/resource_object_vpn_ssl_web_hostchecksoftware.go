// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SSL-VPN host check software.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnSslWebHostCheckSoftware() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnSslWebHostCheckSoftwareCreate,
		Read:   resourceObjectVpnSslWebHostCheckSoftwareRead,
		Update: resourceObjectVpnSslWebHostCheckSoftwareUpdate,
		Delete: resourceObjectVpnSslWebHostCheckSoftwareDelete,

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
			"check_item_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"md5s": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"target": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"guid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"os_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectVpnSslWebHostCheckSoftwareCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectVpnSslWebHostCheckSoftware(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebHostCheckSoftware resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectVpnSslWebHostCheckSoftware(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebHostCheckSoftware resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnSslWebHostCheckSoftwareRead(d, m)
}

func resourceObjectVpnSslWebHostCheckSoftwareUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectVpnSslWebHostCheckSoftware(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebHostCheckSoftware resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectVpnSslWebHostCheckSoftware(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebHostCheckSoftware resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnSslWebHostCheckSoftwareRead(d, m)
}

func resourceObjectVpnSslWebHostCheckSoftwareDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectVpnSslWebHostCheckSoftware(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnSslWebHostCheckSoftware resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnSslWebHostCheckSoftwareRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectVpnSslWebHostCheckSoftware(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebHostCheckSoftware resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnSslWebHostCheckSoftware(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebHostCheckSoftware resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnSslWebHostCheckSoftwareCheckItemList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenObjectVpnSslWebHostCheckSoftwareCheckItemListAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebHostCheckSoftware-CheckItemList-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectVpnSslWebHostCheckSoftwareCheckItemListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebHostCheckSoftware-CheckItemList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5s"
		if _, ok := i["md5s"]; ok {
			v := flattenObjectVpnSslWebHostCheckSoftwareCheckItemListMd5S(i["md5s"], d, pre_append)
			tmp["md5s"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebHostCheckSoftware-CheckItemList-Md5S")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			v := flattenObjectVpnSslWebHostCheckSoftwareCheckItemListTarget(i["target"], d, pre_append)
			tmp["target"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebHostCheckSoftware-CheckItemList-Target")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectVpnSslWebHostCheckSoftwareCheckItemListType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebHostCheckSoftware-CheckItemList-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			v := flattenObjectVpnSslWebHostCheckSoftwareCheckItemListVersion(i["version"], d, pre_append)
			tmp["version"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebHostCheckSoftware-CheckItemList-Version")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectVpnSslWebHostCheckSoftwareCheckItemListAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebHostCheckSoftwareCheckItemListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebHostCheckSoftwareCheckItemListMd5S(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnSslWebHostCheckSoftwareCheckItemListTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebHostCheckSoftwareCheckItemListType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebHostCheckSoftwareCheckItemListVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebHostCheckSoftwareGuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebHostCheckSoftwareName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebHostCheckSoftwareOsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebHostCheckSoftwareType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebHostCheckSoftwareVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnSslWebHostCheckSoftware(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("check_item_list", flattenObjectVpnSslWebHostCheckSoftwareCheckItemList(o["check-item-list"], d, "check_item_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["check-item-list"], "ObjectVpnSslWebHostCheckSoftware-CheckItemList"); ok {
				if err = d.Set("check_item_list", vv); err != nil {
					return fmt.Errorf("Error reading check_item_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading check_item_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("check_item_list"); ok {
			if err = d.Set("check_item_list", flattenObjectVpnSslWebHostCheckSoftwareCheckItemList(o["check-item-list"], d, "check_item_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["check-item-list"], "ObjectVpnSslWebHostCheckSoftware-CheckItemList"); ok {
					if err = d.Set("check_item_list", vv); err != nil {
						return fmt.Errorf("Error reading check_item_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading check_item_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("guid", flattenObjectVpnSslWebHostCheckSoftwareGuid(o["guid"], d, "guid")); err != nil {
		if vv, ok := fortiAPIPatch(o["guid"], "ObjectVpnSslWebHostCheckSoftware-Guid"); ok {
			if err = d.Set("guid", vv); err != nil {
				return fmt.Errorf("Error reading guid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guid: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectVpnSslWebHostCheckSoftwareName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVpnSslWebHostCheckSoftware-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("os_type", flattenObjectVpnSslWebHostCheckSoftwareOsType(o["os-type"], d, "os_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["os-type"], "ObjectVpnSslWebHostCheckSoftware-OsType"); ok {
			if err = d.Set("os_type", vv); err != nil {
				return fmt.Errorf("Error reading os_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os_type: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectVpnSslWebHostCheckSoftwareType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectVpnSslWebHostCheckSoftware-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("version", flattenObjectVpnSslWebHostCheckSoftwareVersion(o["version"], d, "version")); err != nil {
		if vv, ok := fortiAPIPatch(o["version"], "ObjectVpnSslWebHostCheckSoftware-Version"); ok {
			if err = d.Set("version", vv); err != nil {
				return fmt.Errorf("Error reading version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnSslWebHostCheckSoftwareFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnSslWebHostCheckSoftwareCheckItemList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectVpnSslWebHostCheckSoftwareCheckItemListAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectVpnSslWebHostCheckSoftwareCheckItemListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5s"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["md5s"], _ = expandObjectVpnSslWebHostCheckSoftwareCheckItemListMd5S(d, i["md5s"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target"], _ = expandObjectVpnSslWebHostCheckSoftwareCheckItemListTarget(d, i["target"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectVpnSslWebHostCheckSoftwareCheckItemListType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["version"], _ = expandObjectVpnSslWebHostCheckSoftwareCheckItemListVersion(d, i["version"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectVpnSslWebHostCheckSoftwareCheckItemListAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebHostCheckSoftwareCheckItemListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebHostCheckSoftwareCheckItemListMd5S(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnSslWebHostCheckSoftwareCheckItemListTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebHostCheckSoftwareCheckItemListType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebHostCheckSoftwareCheckItemListVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebHostCheckSoftwareGuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebHostCheckSoftwareName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebHostCheckSoftwareOsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebHostCheckSoftwareType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebHostCheckSoftwareVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnSslWebHostCheckSoftware(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("check_item_list"); ok || d.HasChange("check_item_list") {
		t, err := expandObjectVpnSslWebHostCheckSoftwareCheckItemList(d, v, "check_item_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-item-list"] = t
		}
	}

	if v, ok := d.GetOk("guid"); ok || d.HasChange("guid") {
		t, err := expandObjectVpnSslWebHostCheckSoftwareGuid(d, v, "guid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guid"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectVpnSslWebHostCheckSoftwareName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("os_type"); ok || d.HasChange("os_type") {
		t, err := expandObjectVpnSslWebHostCheckSoftwareOsType(d, v, "os_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os-type"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectVpnSslWebHostCheckSoftwareType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandObjectVpnSslWebHostCheckSoftwareVersion(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	return &obj, nil
}
