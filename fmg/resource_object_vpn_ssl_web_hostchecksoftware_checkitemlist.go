// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Check item list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnSslWebHostCheckSoftwareCheckItemList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnSslWebHostCheckSoftwareCheckItemListCreate,
		Read:   resourceObjectVpnSslWebHostCheckSoftwareCheckItemListRead,
		Update: resourceObjectVpnSslWebHostCheckSoftwareCheckItemListUpdate,
		Delete: resourceObjectVpnSslWebHostCheckSoftwareCheckItemListDelete,

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
			"host_check_software": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
	}
}

func resourceObjectVpnSslWebHostCheckSoftwareCheckItemListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	host_check_software := d.Get("host_check_software").(string)
	paradict["host_check_software"] = host_check_software

	obj, err := getObjectObjectVpnSslWebHostCheckSoftwareCheckItemList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebHostCheckSoftwareCheckItemList resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVpnSslWebHostCheckSoftwareCheckItemList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebHostCheckSoftwareCheckItemList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectVpnSslWebHostCheckSoftwareCheckItemListRead(d, m)
}

func resourceObjectVpnSslWebHostCheckSoftwareCheckItemListUpdate(d *schema.ResourceData, m interface{}) error {
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

	host_check_software := d.Get("host_check_software").(string)
	paradict["host_check_software"] = host_check_software

	obj, err := getObjectObjectVpnSslWebHostCheckSoftwareCheckItemList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebHostCheckSoftwareCheckItemList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVpnSslWebHostCheckSoftwareCheckItemList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebHostCheckSoftwareCheckItemList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectVpnSslWebHostCheckSoftwareCheckItemListRead(d, m)
}

func resourceObjectVpnSslWebHostCheckSoftwareCheckItemListDelete(d *schema.ResourceData, m interface{}) error {
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

	host_check_software := d.Get("host_check_software").(string)
	paradict["host_check_software"] = host_check_software

	err = c.DeleteObjectVpnSslWebHostCheckSoftwareCheckItemList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnSslWebHostCheckSoftwareCheckItemList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnSslWebHostCheckSoftwareCheckItemListRead(d *schema.ResourceData, m interface{}) error {
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

	host_check_software := d.Get("host_check_software").(string)
	if host_check_software == "" {
		host_check_software = importOptionChecking(m.(*FortiClient).Cfg, "host_check_software")
		if host_check_software == "" {
			return fmt.Errorf("Parameter host_check_software is missing")
		}
		if err = d.Set("host_check_software", host_check_software); err != nil {
			return fmt.Errorf("Error set params host_check_software: %v", err)
		}
	}
	paradict["host_check_software"] = host_check_software

	o, err := c.ReadObjectVpnSslWebHostCheckSoftwareCheckItemList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebHostCheckSoftwareCheckItemList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnSslWebHostCheckSoftwareCheckItemList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebHostCheckSoftwareCheckItemList resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnSslWebHostCheckSoftwareCheckItemListAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebHostCheckSoftwareCheckItemListId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebHostCheckSoftwareCheckItemListMd5S2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnSslWebHostCheckSoftwareCheckItemListTarget2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebHostCheckSoftwareCheckItemListType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebHostCheckSoftwareCheckItemListVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnSslWebHostCheckSoftwareCheckItemList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectVpnSslWebHostCheckSoftwareCheckItemListAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectVpnSslWebHostCheckSoftwareCheckItemList-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectVpnSslWebHostCheckSoftwareCheckItemListId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectVpnSslWebHostCheckSoftwareCheckItemList-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("md5s", flattenObjectVpnSslWebHostCheckSoftwareCheckItemListMd5S2edl(o["md5s"], d, "md5s")); err != nil {
		if vv, ok := fortiAPIPatch(o["md5s"], "ObjectVpnSslWebHostCheckSoftwareCheckItemList-Md5S"); ok {
			if err = d.Set("md5s", vv); err != nil {
				return fmt.Errorf("Error reading md5s: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading md5s: %v", err)
		}
	}

	if err = d.Set("target", flattenObjectVpnSslWebHostCheckSoftwareCheckItemListTarget2edl(o["target"], d, "target")); err != nil {
		if vv, ok := fortiAPIPatch(o["target"], "ObjectVpnSslWebHostCheckSoftwareCheckItemList-Target"); ok {
			if err = d.Set("target", vv); err != nil {
				return fmt.Errorf("Error reading target: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectVpnSslWebHostCheckSoftwareCheckItemListType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectVpnSslWebHostCheckSoftwareCheckItemList-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("version", flattenObjectVpnSslWebHostCheckSoftwareCheckItemListVersion2edl(o["version"], d, "version")); err != nil {
		if vv, ok := fortiAPIPatch(o["version"], "ObjectVpnSslWebHostCheckSoftwareCheckItemList-Version"); ok {
			if err = d.Set("version", vv); err != nil {
				return fmt.Errorf("Error reading version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnSslWebHostCheckSoftwareCheckItemListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnSslWebHostCheckSoftwareCheckItemListAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebHostCheckSoftwareCheckItemListId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebHostCheckSoftwareCheckItemListMd5S2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnSslWebHostCheckSoftwareCheckItemListTarget2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebHostCheckSoftwareCheckItemListType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebHostCheckSoftwareCheckItemListVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnSslWebHostCheckSoftwareCheckItemList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectVpnSslWebHostCheckSoftwareCheckItemListAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectVpnSslWebHostCheckSoftwareCheckItemListId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("md5s"); ok || d.HasChange("md5s") {
		t, err := expandObjectVpnSslWebHostCheckSoftwareCheckItemListMd5S2edl(d, v, "md5s")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["md5s"] = t
		}
	}

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandObjectVpnSslWebHostCheckSoftwareCheckItemListTarget2edl(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectVpnSslWebHostCheckSoftwareCheckItemListType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandObjectVpnSslWebHostCheckSoftwareCheckItemListVersion2edl(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	return &obj, nil
}
