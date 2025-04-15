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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			},
			"flags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"lock_override": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"log_db_retention_hours": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_disk_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"os_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary_dns_ip4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary_dns_ip6_1": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"primary_dns_ip6_2": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"primary_dns_ip6_3": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"primary_dns_ip6_4": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"restricted_prds": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"secondary_dns_ip4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary_dns_ip6_1": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"secondary_dns_ip6_2": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"secondary_dns_ip6_3": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"secondary_dns_ip6_4": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tz": &schema.Schema{
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	obj, err := getObjectDvmdbAdom(d)
	if err != nil {
		return fmt.Errorf("Error creating DvmdbAdom resource while getting object: %v", err)
	}

	wsParams["lockMode"] = "skip_lock"

	_, err = c.CreateDvmdbAdom(obj, paradict, wsParams)
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	obj, err := getObjectDvmdbAdom(d)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbAdom resource while getting object: %v", err)
	}

	adomv := "adom/" + d.Get("name").(string)
	wsParams["adom"] = adomv

	_, err = c.UpdateDvmdbAdom(obj, mkey, paradict, wsParams)
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "adom/"+d.Get("name").(string), fmt.Errorf("")
	wsParams["adom"] = adomv
	wsParams["lockMode"] = "skip_unlock"

	err = c.DeleteDvmdbAdom(mkey, paradict, wsParams)
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

	paradict := make(map[string]string)

	o, err := c.ReadDvmdbAdom(mkey, paradict)
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

func flattenDvmdbAdomLockOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
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

func flattenDvmdbAdomPrimaryDnsIp4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomPrimaryDnsIp61(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomPrimaryDnsIp62(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomPrimaryDnsIp63(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomPrimaryDnsIp64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomRestrictedPrds(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convstr2list(v, d.Get(pre))
}

func flattenDvmdbAdomSecondaryDnsIp4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomSecondaryDnsIp61(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomSecondaryDnsIp62(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomSecondaryDnsIp63(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomSecondaryDnsIp64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbAdomTz(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	if err = d.Set("lock_override", flattenDvmdbAdomLockOverride(o["lock_override"], d, "lock_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["lock_override"], "DvmdbAdom-LockOverride"); ok {
			if err = d.Set("lock_override", vv); err != nil {
				return fmt.Errorf("Error reading lock_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lock_override: %v", err)
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

	if err = d.Set("primary_dns_ip4", flattenDvmdbAdomPrimaryDnsIp4(o["primary_dns_ip4"], d, "primary_dns_ip4")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary_dns_ip4"], "DvmdbAdom-PrimaryDnsIp4"); ok {
			if err = d.Set("primary_dns_ip4", vv); err != nil {
				return fmt.Errorf("Error reading primary_dns_ip4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_dns_ip4: %v", err)
		}
	}

	if err = d.Set("primary_dns_ip6_1", flattenDvmdbAdomPrimaryDnsIp61(o["primary_dns_ip6_1"], d, "primary_dns_ip6_1")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary_dns_ip6_1"], "DvmdbAdom-PrimaryDnsIp61"); ok {
			if err = d.Set("primary_dns_ip6_1", vv); err != nil {
				return fmt.Errorf("Error reading primary_dns_ip6_1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_dns_ip6_1: %v", err)
		}
	}

	if err = d.Set("primary_dns_ip6_2", flattenDvmdbAdomPrimaryDnsIp62(o["primary_dns_ip6_2"], d, "primary_dns_ip6_2")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary_dns_ip6_2"], "DvmdbAdom-PrimaryDnsIp62"); ok {
			if err = d.Set("primary_dns_ip6_2", vv); err != nil {
				return fmt.Errorf("Error reading primary_dns_ip6_2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_dns_ip6_2: %v", err)
		}
	}

	if err = d.Set("primary_dns_ip6_3", flattenDvmdbAdomPrimaryDnsIp63(o["primary_dns_ip6_3"], d, "primary_dns_ip6_3")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary_dns_ip6_3"], "DvmdbAdom-PrimaryDnsIp63"); ok {
			if err = d.Set("primary_dns_ip6_3", vv); err != nil {
				return fmt.Errorf("Error reading primary_dns_ip6_3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_dns_ip6_3: %v", err)
		}
	}

	if err = d.Set("primary_dns_ip6_4", flattenDvmdbAdomPrimaryDnsIp64(o["primary_dns_ip6_4"], d, "primary_dns_ip6_4")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary_dns_ip6_4"], "DvmdbAdom-PrimaryDnsIp64"); ok {
			if err = d.Set("primary_dns_ip6_4", vv); err != nil {
				return fmt.Errorf("Error reading primary_dns_ip6_4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_dns_ip6_4: %v", err)
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

	if err = d.Set("secondary_dns_ip4", flattenDvmdbAdomSecondaryDnsIp4(o["secondary_dns_ip4"], d, "secondary_dns_ip4")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary_dns_ip4"], "DvmdbAdom-SecondaryDnsIp4"); ok {
			if err = d.Set("secondary_dns_ip4", vv); err != nil {
				return fmt.Errorf("Error reading secondary_dns_ip4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_dns_ip4: %v", err)
		}
	}

	if err = d.Set("secondary_dns_ip6_1", flattenDvmdbAdomSecondaryDnsIp61(o["secondary_dns_ip6_1"], d, "secondary_dns_ip6_1")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary_dns_ip6_1"], "DvmdbAdom-SecondaryDnsIp61"); ok {
			if err = d.Set("secondary_dns_ip6_1", vv); err != nil {
				return fmt.Errorf("Error reading secondary_dns_ip6_1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_dns_ip6_1: %v", err)
		}
	}

	if err = d.Set("secondary_dns_ip6_2", flattenDvmdbAdomSecondaryDnsIp62(o["secondary_dns_ip6_2"], d, "secondary_dns_ip6_2")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary_dns_ip6_2"], "DvmdbAdom-SecondaryDnsIp62"); ok {
			if err = d.Set("secondary_dns_ip6_2", vv); err != nil {
				return fmt.Errorf("Error reading secondary_dns_ip6_2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_dns_ip6_2: %v", err)
		}
	}

	if err = d.Set("secondary_dns_ip6_3", flattenDvmdbAdomSecondaryDnsIp63(o["secondary_dns_ip6_3"], d, "secondary_dns_ip6_3")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary_dns_ip6_3"], "DvmdbAdom-SecondaryDnsIp63"); ok {
			if err = d.Set("secondary_dns_ip6_3", vv); err != nil {
				return fmt.Errorf("Error reading secondary_dns_ip6_3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_dns_ip6_3: %v", err)
		}
	}

	if err = d.Set("secondary_dns_ip6_4", flattenDvmdbAdomSecondaryDnsIp64(o["secondary_dns_ip6_4"], d, "secondary_dns_ip6_4")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary_dns_ip6_4"], "DvmdbAdom-SecondaryDnsIp64"); ok {
			if err = d.Set("secondary_dns_ip6_4", vv); err != nil {
				return fmt.Errorf("Error reading secondary_dns_ip6_4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_dns_ip6_4: %v", err)
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

	if err = d.Set("tz", flattenDvmdbAdomTz(o["tz"], d, "tz")); err != nil {
		if vv, ok := fortiAPIPatch(o["tz"], "DvmdbAdom-Tz"); ok {
			if err = d.Set("tz", vv); err != nil {
				return fmt.Errorf("Error reading tz: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tz: %v", err)
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

func expandDvmdbAdomLockOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
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

func expandDvmdbAdomPrimaryDnsIp4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomPrimaryDnsIp61(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomPrimaryDnsIp62(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomPrimaryDnsIp63(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomPrimaryDnsIp64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomRestrictedPrds(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDvmdbAdomSecondaryDnsIp4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomSecondaryDnsIp61(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomSecondaryDnsIp62(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomSecondaryDnsIp63(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomSecondaryDnsIp64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbAdomTz(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

	if v, ok := d.GetOk("create_time"); ok || d.HasChange("create_time") {
		t, err := expandDvmdbAdomCreateTime(d, v, "create_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create_time"] = t
		}
	}

	if v, ok := d.GetOk("desc"); ok || d.HasChange("desc") {
		t, err := expandDvmdbAdomDesc(d, v, "desc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["desc"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandDvmdbAdomFlags(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("lock_override"); ok || d.HasChange("lock_override") {
		t, err := expandDvmdbAdomLockOverride(d, v, "lock_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lock_override"] = t
		}
	}

	if v, ok := d.GetOk("log_db_retention_hours"); ok || d.HasChange("log_db_retention_hours") {
		t, err := expandDvmdbAdomLogDbRetentionHours(d, v, "log_db_retention_hours")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log_db_retention_hours"] = t
		}
	}

	if v, ok := d.GetOk("log_disk_quota"); ok || d.HasChange("log_disk_quota") {
		t, err := expandDvmdbAdomLogDiskQuota(d, v, "log_disk_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log_disk_quota"] = t
		}
	}

	if v, ok := d.GetOk("log_disk_quota_alert_thres"); ok || d.HasChange("log_disk_quota_alert_thres") {
		t, err := expandDvmdbAdomLogDiskQuotaAlertThres(d, v, "log_disk_quota_alert_thres")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log_disk_quota_alert_thres"] = t
		}
	}

	if v, ok := d.GetOk("log_disk_quota_split_ratio"); ok || d.HasChange("log_disk_quota_split_ratio") {
		t, err := expandDvmdbAdomLogDiskQuotaSplitRatio(d, v, "log_disk_quota_split_ratio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log_disk_quota_split_ratio"] = t
		}
	}

	if v, ok := d.GetOk("log_file_retention_hours"); ok || d.HasChange("log_file_retention_hours") {
		t, err := expandDvmdbAdomLogFileRetentionHours(d, v, "log_file_retention_hours")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log_file_retention_hours"] = t
		}
	}

	if v, ok := d.GetOk("metafields"); ok || d.HasChange("metafields") {
		t, err := expandDvmdbAdomMetaFields(d, v, "metafields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["meta fields"] = t
		}
	}

	if v, ok := d.GetOk("mig_mr"); ok || d.HasChange("mig_mr") {
		t, err := expandDvmdbAdomMigMr(d, v, "mig_mr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mig_mr"] = t
		}
	}

	if v, ok := d.GetOk("mig_os_ver"); ok || d.HasChange("mig_os_ver") {
		t, err := expandDvmdbAdomMigOsVer(d, v, "mig_os_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mig_os_ver"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandDvmdbAdomMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOkExists("mr"); ok {
		t, err := expandDvmdbAdomMr(d, v, "mr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mr"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDvmdbAdomName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("os_ver"); ok || d.HasChange("os_ver") {
		t, err := expandDvmdbAdomOsVer(d, v, "os_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os_ver"] = t
		}
	}

	if v, ok := d.GetOk("primary_dns_ip4"); ok || d.HasChange("primary_dns_ip4") {
		t, err := expandDvmdbAdomPrimaryDnsIp4(d, v, "primary_dns_ip4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary_dns_ip4"] = t
		}
	}

	if v, ok := d.GetOk("primary_dns_ip6_1"); ok || d.HasChange("primary_dns_ip6_1") {
		t, err := expandDvmdbAdomPrimaryDnsIp61(d, v, "primary_dns_ip6_1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary_dns_ip6_1"] = t
		}
	}

	if v, ok := d.GetOk("primary_dns_ip6_2"); ok || d.HasChange("primary_dns_ip6_2") {
		t, err := expandDvmdbAdomPrimaryDnsIp62(d, v, "primary_dns_ip6_2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary_dns_ip6_2"] = t
		}
	}

	if v, ok := d.GetOk("primary_dns_ip6_3"); ok || d.HasChange("primary_dns_ip6_3") {
		t, err := expandDvmdbAdomPrimaryDnsIp63(d, v, "primary_dns_ip6_3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary_dns_ip6_3"] = t
		}
	}

	if v, ok := d.GetOk("primary_dns_ip6_4"); ok || d.HasChange("primary_dns_ip6_4") {
		t, err := expandDvmdbAdomPrimaryDnsIp64(d, v, "primary_dns_ip6_4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary_dns_ip6_4"] = t
		}
	}

	if v, ok := d.GetOk("restricted_prds"); ok || d.HasChange("restricted_prds") {
		t, err := expandDvmdbAdomRestrictedPrds(d, v, "restricted_prds")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restricted_prds"] = t
		}
	}

	if v, ok := d.GetOk("secondary_dns_ip4"); ok || d.HasChange("secondary_dns_ip4") {
		t, err := expandDvmdbAdomSecondaryDnsIp4(d, v, "secondary_dns_ip4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary_dns_ip4"] = t
		}
	}

	if v, ok := d.GetOk("secondary_dns_ip6_1"); ok || d.HasChange("secondary_dns_ip6_1") {
		t, err := expandDvmdbAdomSecondaryDnsIp61(d, v, "secondary_dns_ip6_1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary_dns_ip6_1"] = t
		}
	}

	if v, ok := d.GetOk("secondary_dns_ip6_2"); ok || d.HasChange("secondary_dns_ip6_2") {
		t, err := expandDvmdbAdomSecondaryDnsIp62(d, v, "secondary_dns_ip6_2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary_dns_ip6_2"] = t
		}
	}

	if v, ok := d.GetOk("secondary_dns_ip6_3"); ok || d.HasChange("secondary_dns_ip6_3") {
		t, err := expandDvmdbAdomSecondaryDnsIp63(d, v, "secondary_dns_ip6_3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary_dns_ip6_3"] = t
		}
	}

	if v, ok := d.GetOk("secondary_dns_ip6_4"); ok || d.HasChange("secondary_dns_ip6_4") {
		t, err := expandDvmdbAdomSecondaryDnsIp64(d, v, "secondary_dns_ip6_4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary_dns_ip6_4"] = t
		}
	}

	if v, ok := d.GetOk("state"); ok || d.HasChange("state") {
		t, err := expandDvmdbAdomState(d, v, "state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["state"] = t
		}
	}

	if v, ok := d.GetOk("tz"); ok || d.HasChange("tz") {
		t, err := expandDvmdbAdomTz(d, v, "tz")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tz"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandDvmdbAdomUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("workspace_mode"); ok || d.HasChange("workspace_mode") {
		t, err := expandDvmdbAdomWorkspaceMode(d, v, "workspace_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["workspace_mode"] = t
		}
	}

	return &obj, nil
}
