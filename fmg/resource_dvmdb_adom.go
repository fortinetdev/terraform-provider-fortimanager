// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ADOM table, most attributes are read-only and can only be changed internally.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDvmdbAdom() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmdbAdomCreate,
		Read:   resourceDvmdbAdomRead,
		Update: resourceDvmdbAdomUpdate,
		Delete: resourceDvmdbAdomDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"create_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"desc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"log_db_retention_hours": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_disk_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_disk_quota_alert_thres": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_disk_quota_split_ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_file_retention_hours": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"metafields": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mig_mr": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mig_os_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mr": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"os_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restricted_prds": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"workspace_mode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDvmdbAdomCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "", fmt.Errorf("")

	obj, err := getObjectDvmdbAdom(d)
	if err != nil {
		return fmt.Errorf("Error creating DvmdbAdom resource while getting object: %v", err)
	}

	_, err = c.CreateDvmdbAdom(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating DvmdbAdom resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceDvmdbAdomRead(d, m)
}

func resourceDvmdbAdomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "", fmt.Errorf("")

	obj, err := getObjectDvmdbAdom(d)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbAdom resource while getting object: %v", err)
	}

	_, err = c.UpdateDvmdbAdom(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbAdom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceDvmdbAdomRead(d, m)
}

func resourceDvmdbAdomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "", fmt.Errorf("")

	err = c.DeleteDvmdbAdom(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting DvmdbAdom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDvmdbAdomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "", fmt.Errorf("")

	o, err := c.ReadDvmdbAdom(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading DvmdbAdom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDvmdbAdom(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DvmdbAdom resource from API: %v", err)
	}
	return nil
}

func flattenDvmdbAdomCreateTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomDesc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDvmdbAdomLogDbRetentionHours(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomLogDiskQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomLogDiskQuotaAlertThres(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomLogDiskQuotaSplitRatio(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomLogFileRetentionHours(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomMetaFields(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomMigMr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomMigOsVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomMr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomOsVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomRestrictedPrds(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDvmdbAdomState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomWorkspaceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDvmdbAdom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("create_time", flattenDvmdbAdomCreateTime(o["create_time"], d, "create_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["create_time"], "DvmdbAdom-CreateTime"); ok {
			if err = d.Set("create_time", vv); err != nil {
				return fmt.Errorf("Error reading create_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading create_time: %v", err)
		}
	}

	if err = d.Set("desc", flattenDvmdbAdomDesc(o["desc"], d, "desc")); err != nil {
		if vv, ok := fortiAPIPatch(o["desc"], "DvmdbAdom-Desc"); ok {
			if err = d.Set("desc", vv); err != nil {
				return fmt.Errorf("Error reading desc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading desc: %v", err)
		}
	}

	if err = d.Set("flags", flattenDvmdbAdomFlags(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "DvmdbAdom-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if err = d.Set("log_db_retention_hours", flattenDvmdbAdomLogDbRetentionHours(o["log_db_retention_hours"], d, "log_db_retention_hours")); err != nil {
		if vv, ok := fortiAPIPatch(o["log_db_retention_hours"], "DvmdbAdom-LogDbRetentionHours"); ok {
			if err = d.Set("log_db_retention_hours", vv); err != nil {
				return fmt.Errorf("Error reading log_db_retention_hours: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_db_retention_hours: %v", err)
		}
	}

	if err = d.Set("log_disk_quota", flattenDvmdbAdomLogDiskQuota(o["log_disk_quota"], d, "log_disk_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["log_disk_quota"], "DvmdbAdom-LogDiskQuota"); ok {
			if err = d.Set("log_disk_quota", vv); err != nil {
				return fmt.Errorf("Error reading log_disk_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_disk_quota: %v", err)
		}
	}

	if err = d.Set("log_disk_quota_alert_thres", flattenDvmdbAdomLogDiskQuotaAlertThres(o["log_disk_quota_alert_thres"], d, "log_disk_quota_alert_thres")); err != nil {
		if vv, ok := fortiAPIPatch(o["log_disk_quota_alert_thres"], "DvmdbAdom-LogDiskQuotaAlertThres"); ok {
			if err = d.Set("log_disk_quota_alert_thres", vv); err != nil {
				return fmt.Errorf("Error reading log_disk_quota_alert_thres: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_disk_quota_alert_thres: %v", err)
		}
	}

	if err = d.Set("log_disk_quota_split_ratio", flattenDvmdbAdomLogDiskQuotaSplitRatio(o["log_disk_quota_split_ratio"], d, "log_disk_quota_split_ratio")); err != nil {
		if vv, ok := fortiAPIPatch(o["log_disk_quota_split_ratio"], "DvmdbAdom-LogDiskQuotaSplitRatio"); ok {
			if err = d.Set("log_disk_quota_split_ratio", vv); err != nil {
				return fmt.Errorf("Error reading log_disk_quota_split_ratio: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_disk_quota_split_ratio: %v", err)
		}
	}

	if err = d.Set("log_file_retention_hours", flattenDvmdbAdomLogFileRetentionHours(o["log_file_retention_hours"], d, "log_file_retention_hours")); err != nil {
		if vv, ok := fortiAPIPatch(o["log_file_retention_hours"], "DvmdbAdom-LogFileRetentionHours"); ok {
			if err = d.Set("log_file_retention_hours", vv); err != nil {
				return fmt.Errorf("Error reading log_file_retention_hours: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_file_retention_hours: %v", err)
		}
	}

	if err = d.Set("metafields", flattenDvmdbAdomMetaFields(o["meta fields"], d, "metafields")); err != nil {
		if vv, ok := fortiAPIPatch(o["meta fields"], "DvmdbAdom-MetaFields"); ok {
			if err = d.Set("metafields", vv); err != nil {
				return fmt.Errorf("Error reading metafields: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading metafields: %v", err)
		}
	}

	if err = d.Set("mig_mr", flattenDvmdbAdomMigMr(o["mig_mr"], d, "mig_mr")); err != nil {
		if vv, ok := fortiAPIPatch(o["mig_mr"], "DvmdbAdom-MigMr"); ok {
			if err = d.Set("mig_mr", vv); err != nil {
				return fmt.Errorf("Error reading mig_mr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mig_mr: %v", err)
		}
	}

	if err = d.Set("mig_os_ver", flattenDvmdbAdomMigOsVer(o["mig_os_ver"], d, "mig_os_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["mig_os_ver"], "DvmdbAdom-MigOsVer"); ok {
			if err = d.Set("mig_os_ver", vv); err != nil {
				return fmt.Errorf("Error reading mig_os_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mig_os_ver: %v", err)
		}
	}

	if err = d.Set("mode", flattenDvmdbAdomMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "DvmdbAdom-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("mr", flattenDvmdbAdomMr(o["mr"], d, "mr")); err != nil {
		if vv, ok := fortiAPIPatch(o["mr"], "DvmdbAdom-Mr"); ok {
			if err = d.Set("mr", vv); err != nil {
				return fmt.Errorf("Error reading mr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mr: %v", err)
		}
	}

	if err = d.Set("name", flattenDvmdbAdomName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DvmdbAdom-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("os_ver", flattenDvmdbAdomOsVer(o["os_ver"], d, "os_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["os_ver"], "DvmdbAdom-OsVer"); ok {
			if err = d.Set("os_ver", vv); err != nil {
				return fmt.Errorf("Error reading os_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os_ver: %v", err)
		}
	}

	if err = d.Set("restricted_prds", flattenDvmdbAdomRestrictedPrds(o["restricted_prds"], d, "restricted_prds")); err != nil {
		if vv, ok := fortiAPIPatch(o["restricted_prds"], "DvmdbAdom-RestrictedPrds"); ok {
			if err = d.Set("restricted_prds", vv); err != nil {
				return fmt.Errorf("Error reading restricted_prds: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restricted_prds: %v", err)
		}
	}

	if err = d.Set("state", flattenDvmdbAdomState(o["state"], d, "state")); err != nil {
		if vv, ok := fortiAPIPatch(o["state"], "DvmdbAdom-State"); ok {
			if err = d.Set("state", vv); err != nil {
				return fmt.Errorf("Error reading state: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading state: %v", err)
		}
	}

	if err = d.Set("uuid", flattenDvmdbAdomUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "DvmdbAdom-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("workspace_mode", flattenDvmdbAdomWorkspaceMode(o["workspace_mode"], d, "workspace_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["workspace_mode"], "DvmdbAdom-WorkspaceMode"); ok {
			if err = d.Set("workspace_mode", vv); err != nil {
				return fmt.Errorf("Error reading workspace_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading workspace_mode: %v", err)
		}
	}

	return nil
}

func flattenDvmdbAdomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmdbAdomCreateTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomDesc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDvmdbAdomLogDbRetentionHours(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomLogDiskQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomLogDiskQuotaAlertThres(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomLogDiskQuotaSplitRatio(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomLogFileRetentionHours(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomMetaFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomMigMr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomMigOsVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomMr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomOsVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomRestrictedPrds(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDvmdbAdomState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomWorkspaceMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDvmdbAdom(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("create_time"); ok {
		t, err := expandDvmdbAdomCreateTime(d, v, "create_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create_time"] = t
		}
	}

	if v, ok := d.GetOk("desc"); ok {
		t, err := expandDvmdbAdomDesc(d, v, "desc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["desc"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok {
		t, err := expandDvmdbAdomFlags(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("log_db_retention_hours"); ok {
		t, err := expandDvmdbAdomLogDbRetentionHours(d, v, "log_db_retention_hours")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log_db_retention_hours"] = t
		}
	}

	if v, ok := d.GetOk("log_disk_quota"); ok {
		t, err := expandDvmdbAdomLogDiskQuota(d, v, "log_disk_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log_disk_quota"] = t
		}
	}

	if v, ok := d.GetOk("log_disk_quota_alert_thres"); ok {
		t, err := expandDvmdbAdomLogDiskQuotaAlertThres(d, v, "log_disk_quota_alert_thres")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log_disk_quota_alert_thres"] = t
		}
	}

	if v, ok := d.GetOk("log_disk_quota_split_ratio"); ok {
		t, err := expandDvmdbAdomLogDiskQuotaSplitRatio(d, v, "log_disk_quota_split_ratio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log_disk_quota_split_ratio"] = t
		}
	}

	if v, ok := d.GetOk("log_file_retention_hours"); ok {
		t, err := expandDvmdbAdomLogFileRetentionHours(d, v, "log_file_retention_hours")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log_file_retention_hours"] = t
		}
	}

	if v, ok := d.GetOk("metafields"); ok {
		t, err := expandDvmdbAdomMetaFields(d, v, "metafields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["meta fields"] = t
		}
	}

	if v, ok := d.GetOk("mig_mr"); ok {
		t, err := expandDvmdbAdomMigMr(d, v, "mig_mr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mig_mr"] = t
		}
	}

	if v, ok := d.GetOk("mig_os_ver"); ok {
		t, err := expandDvmdbAdomMigOsVer(d, v, "mig_os_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mig_os_ver"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandDvmdbAdomMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("mr"); ok {
		t, err := expandDvmdbAdomMr(d, v, "mr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mr"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandDvmdbAdomName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("os_ver"); ok {
		t, err := expandDvmdbAdomOsVer(d, v, "os_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os_ver"] = t
		}
	}

	if v, ok := d.GetOk("restricted_prds"); ok {
		t, err := expandDvmdbAdomRestrictedPrds(d, v, "restricted_prds")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restricted_prds"] = t
		}
	}

	if v, ok := d.GetOk("state"); ok {
		t, err := expandDvmdbAdomState(d, v, "state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["state"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandDvmdbAdomUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("workspace_mode"); ok {
		t, err := expandDvmdbAdomWorkspaceMode(d, v, "workspace_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["workspace_mode"] = t
		}
	}

	return &obj, nil
}
