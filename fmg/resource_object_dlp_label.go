// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure labels used by DLP blocking.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDlpLabel() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDlpLabelCreate,
		Read:   resourceObjectDlpLabelRead,
		Update: resourceObjectDlpLabelUpdate,
		Delete: resourceObjectDlpLabelDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			},
			"connector": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fortidata_label_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"guid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mpip_label_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"fabric_force_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_object_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mpip_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceObjectDlpLabelCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectDlpLabel(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpLabel resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectDlpLabel(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectDlpLabel(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectDlpLabel resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectDlpLabel(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectDlpLabel resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDlpLabelRead(d, m)
}

func resourceObjectDlpLabelUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectDlpLabel(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpLabel resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectDlpLabel(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpLabel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDlpLabelRead(d, m)
}

func resourceObjectDlpLabelDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectDlpLabel(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDlpLabel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDlpLabelRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectDlpLabel(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectDlpLabel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDlpLabel(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpLabel resource from API: %v", err)
	}
	return nil
}

func flattenObjectDlpLabelComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpLabelConnector(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectDlpLabelEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortidata_label_name"
		if _, ok := i["fortidata-label-name"]; ok {
			v := flattenObjectDlpLabelEntriesFortidataLabelName(i["fortidata-label-name"], d, pre_append)
			tmp["fortidata_label_name"] = fortiAPISubPartPatch(v, "ObjectDlpLabel-Entries-FortidataLabelName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guid"
		if _, ok := i["guid"]; ok {
			v := flattenObjectDlpLabelEntriesGuid(i["guid"], d, pre_append)
			tmp["guid"] = fortiAPISubPartPatch(v, "ObjectDlpLabel-Entries-Guid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectDlpLabelEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectDlpLabel-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpip_label_name"
		if _, ok := i["mpip-label-name"]; ok {
			v := flattenObjectDlpLabelEntriesMpipLabelName(i["mpip-label-name"], d, pre_append)
			tmp["mpip_label_name"] = fortiAPISubPartPatch(v, "ObjectDlpLabel-Entries-MpipLabelName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectDlpLabelEntriesFortidataLabelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpLabelEntriesGuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpLabelEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpLabelEntriesMpipLabelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpLabelFabricForceSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpLabelFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpLabelFabricObjectSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpLabelMpipType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpLabelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpLabelType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpLabelUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDlpLabel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenObjectDlpLabelComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectDlpLabel-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("connector", flattenObjectDlpLabelConnector(o["connector"], d, "connector")); err != nil {
		if vv, ok := fortiAPIPatch(o["connector"], "ObjectDlpLabel-Connector"); ok {
			if err = d.Set("connector", vv); err != nil {
				return fmt.Errorf("Error reading connector: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading connector: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenObjectDlpLabelEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "ObjectDlpLabel-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenObjectDlpLabelEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "ObjectDlpLabel-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("fabric_force_sync", flattenObjectDlpLabelFabricForceSync(o["fabric-force-sync"], d, "fabric_force_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-force-sync"], "ObjectDlpLabel-FabricForceSync"); ok {
			if err = d.Set("fabric_force_sync", vv); err != nil {
				return fmt.Errorf("Error reading fabric_force_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_force_sync: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenObjectDlpLabelFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "ObjectDlpLabel-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("fabric_object_source", flattenObjectDlpLabelFabricObjectSource(o["fabric-object-source"], d, "fabric_object_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object-source"], "ObjectDlpLabel-FabricObjectSource"); ok {
			if err = d.Set("fabric_object_source", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object_source: %v", err)
		}
	}

	if err = d.Set("mpip_type", flattenObjectDlpLabelMpipType(o["mpip-type"], d, "mpip_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpip-type"], "ObjectDlpLabel-MpipType"); ok {
			if err = d.Set("mpip_type", vv); err != nil {
				return fmt.Errorf("Error reading mpip_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpip_type: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectDlpLabelName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDlpLabel-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectDlpLabelType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectDlpLabel-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectDlpLabelUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectDlpLabel-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenObjectDlpLabelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDlpLabelComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpLabelConnector(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectDlpLabelEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortidata_label_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fortidata-label-name"], _ = expandObjectDlpLabelEntriesFortidataLabelName(d, i["fortidata_label_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["guid"], _ = expandObjectDlpLabelEntriesGuid(d, i["guid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectDlpLabelEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpip_label_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mpip-label-name"], _ = expandObjectDlpLabelEntriesMpipLabelName(d, i["mpip_label_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectDlpLabelEntriesFortidataLabelName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpLabelEntriesGuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpLabelEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpLabelEntriesMpipLabelName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpLabelFabricForceSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpLabelFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpLabelFabricObjectSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpLabelMpipType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpLabelName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpLabelType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpLabelUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDlpLabel(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectDlpLabelComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("connector"); ok || d.HasChange("connector") {
		t, err := expandObjectDlpLabelConnector(d, v, "connector")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connector"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandObjectDlpLabelEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("fabric_force_sync"); ok || d.HasChange("fabric_force_sync") {
		t, err := expandObjectDlpLabelFabricForceSync(d, v, "fabric_force_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-force-sync"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandObjectDlpLabelFabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object_source"); ok || d.HasChange("fabric_object_source") {
		t, err := expandObjectDlpLabelFabricObjectSource(d, v, "fabric_object_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object-source"] = t
		}
	}

	if v, ok := d.GetOk("mpip_type"); ok || d.HasChange("mpip_type") {
		t, err := expandObjectDlpLabelMpipType(d, v, "mpip_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpip-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectDlpLabelName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectDlpLabelType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandObjectDlpLabelUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
