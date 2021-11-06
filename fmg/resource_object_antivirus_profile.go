// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure AntiVirus profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectAntivirusProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectAntivirusProfileCreate,
		Read:   resourceObjectAntivirusProfileRead,
		Update: resourceObjectAntivirusProfileUpdate,
		Delete: resourceObjectAntivirusProfileDelete,

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
			"analytics_accept_filetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"analytics_bl_filetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"analytics_db": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"analytics_ignore_filetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"analytics_max_upload": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"analytics_wl_filetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_block_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_virus_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cifs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"content_disarm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cover_page": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"detect_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"error_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"office_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"office_dde": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"office_embed": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"office_hylink": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"office_linked": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"office_macro": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"original_file_destination": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pdf_act_form": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pdf_act_gotor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pdf_act_java": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pdf_act_launch": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pdf_act_movie": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pdf_act_sound": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pdf_embedfile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pdf_hyperlink": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pdf_javacode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ems_threat_feed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_blocklist": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"external_blocklist_archive_scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_blocklist_enable_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftgd_analytics": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"http": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content_disarm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"imap": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content_disarm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"executables": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mapi": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"executables": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mobile_malware_db": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nac_quar": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expiry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"infected": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
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
			"nntp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"outbreak_prevention_archive_scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ftgd_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"pop3": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content_disarm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"executables": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smtp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content_disarm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"executables": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssh": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceObjectAntivirusProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectAntivirusProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectAntivirusProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectAntivirusProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectAntivirusProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectAntivirusProfileRead(d, m)
}

func resourceObjectAntivirusProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectAntivirusProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectAntivirusProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectAntivirusProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectAntivirusProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectAntivirusProfileRead(d, m)
}

func resourceObjectAntivirusProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectAntivirusProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectAntivirusProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectAntivirusProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectAntivirusProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectAntivirusProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectAntivirusProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectAntivirusProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectAntivirusProfileAnalyticsAcceptFiletype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileAnalyticsBlFiletype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileAnalyticsDb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileAnalyticsIgnoreFiletype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileAnalyticsMaxUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileAnalyticsWlFiletype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileAvBlockLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileAvVirusLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileCifs(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenObjectAntivirusProfileCifsArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenObjectAntivirusProfileCifsArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenObjectAntivirusProfileCifsAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenObjectAntivirusProfileCifsEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenObjectAntivirusProfileCifsExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectAntivirusProfileCifsOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenObjectAntivirusProfileCifsOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenObjectAntivirusProfileCifsQuarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectAntivirusProfileCifsArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileCifsArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileCifsAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileCifsEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileCifsExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileCifsOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileCifsOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileCifsQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarm(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cover_page"
	if _, ok := i["cover-page"]; ok {
		result["cover_page"] = flattenObjectAntivirusProfileContentDisarmCoverPage(i["cover-page"], d, pre_append)
	}

	pre_append = pre + ".0." + "detect_only"
	if _, ok := i["detect-only"]; ok {
		result["detect_only"] = flattenObjectAntivirusProfileContentDisarmDetectOnly(i["detect-only"], d, pre_append)
	}

	pre_append = pre + ".0." + "error_action"
	if _, ok := i["error-action"]; ok {
		result["error_action"] = flattenObjectAntivirusProfileContentDisarmErrorAction(i["error-action"], d, pre_append)
	}

	pre_append = pre + ".0." + "office_action"
	if _, ok := i["office-action"]; ok {
		result["office_action"] = flattenObjectAntivirusProfileContentDisarmOfficeAction(i["office-action"], d, pre_append)
	}

	pre_append = pre + ".0." + "office_dde"
	if _, ok := i["office-dde"]; ok {
		result["office_dde"] = flattenObjectAntivirusProfileContentDisarmOfficeDde(i["office-dde"], d, pre_append)
	}

	pre_append = pre + ".0." + "office_embed"
	if _, ok := i["office-embed"]; ok {
		result["office_embed"] = flattenObjectAntivirusProfileContentDisarmOfficeEmbed(i["office-embed"], d, pre_append)
	}

	pre_append = pre + ".0." + "office_hylink"
	if _, ok := i["office-hylink"]; ok {
		result["office_hylink"] = flattenObjectAntivirusProfileContentDisarmOfficeHylink(i["office-hylink"], d, pre_append)
	}

	pre_append = pre + ".0." + "office_linked"
	if _, ok := i["office-linked"]; ok {
		result["office_linked"] = flattenObjectAntivirusProfileContentDisarmOfficeLinked(i["office-linked"], d, pre_append)
	}

	pre_append = pre + ".0." + "office_macro"
	if _, ok := i["office-macro"]; ok {
		result["office_macro"] = flattenObjectAntivirusProfileContentDisarmOfficeMacro(i["office-macro"], d, pre_append)
	}

	pre_append = pre + ".0." + "original_file_destination"
	if _, ok := i["original-file-destination"]; ok {
		result["original_file_destination"] = flattenObjectAntivirusProfileContentDisarmOriginalFileDestination(i["original-file-destination"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_act_form"
	if _, ok := i["pdf-act-form"]; ok {
		result["pdf_act_form"] = flattenObjectAntivirusProfileContentDisarmPdfActForm(i["pdf-act-form"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_act_gotor"
	if _, ok := i["pdf-act-gotor"]; ok {
		result["pdf_act_gotor"] = flattenObjectAntivirusProfileContentDisarmPdfActGotor(i["pdf-act-gotor"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_act_java"
	if _, ok := i["pdf-act-java"]; ok {
		result["pdf_act_java"] = flattenObjectAntivirusProfileContentDisarmPdfActJava(i["pdf-act-java"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_act_launch"
	if _, ok := i["pdf-act-launch"]; ok {
		result["pdf_act_launch"] = flattenObjectAntivirusProfileContentDisarmPdfActLaunch(i["pdf-act-launch"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_act_movie"
	if _, ok := i["pdf-act-movie"]; ok {
		result["pdf_act_movie"] = flattenObjectAntivirusProfileContentDisarmPdfActMovie(i["pdf-act-movie"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_act_sound"
	if _, ok := i["pdf-act-sound"]; ok {
		result["pdf_act_sound"] = flattenObjectAntivirusProfileContentDisarmPdfActSound(i["pdf-act-sound"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_embedfile"
	if _, ok := i["pdf-embedfile"]; ok {
		result["pdf_embedfile"] = flattenObjectAntivirusProfileContentDisarmPdfEmbedfile(i["pdf-embedfile"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_hyperlink"
	if _, ok := i["pdf-hyperlink"]; ok {
		result["pdf_hyperlink"] = flattenObjectAntivirusProfileContentDisarmPdfHyperlink(i["pdf-hyperlink"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_javacode"
	if _, ok := i["pdf-javacode"]; ok {
		result["pdf_javacode"] = flattenObjectAntivirusProfileContentDisarmPdfJavacode(i["pdf-javacode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectAntivirusProfileContentDisarmCoverPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmDetectOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmErrorAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmOfficeAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmOfficeDde(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmOfficeEmbed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmOfficeHylink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmOfficeLinked(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmOfficeMacro(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmOriginalFileDestination(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfActForm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfActGotor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfActJava(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfActLaunch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfActMovie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfActSound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfEmbedfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfHyperlink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileContentDisarmPdfJavacode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileEmsThreatFeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileExternalBlocklistArchiveScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileExternalBlocklistEnableAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFeatureSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFtgdAnalytics(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenObjectAntivirusProfileFtpArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenObjectAntivirusProfileFtpArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenObjectAntivirusProfileFtpAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenObjectAntivirusProfileFtpEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenObjectAntivirusProfileFtpExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectAntivirusProfileFtpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenObjectAntivirusProfileFtpOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenObjectAntivirusProfileFtpQuarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectAntivirusProfileFtpArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileFtpArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileFtpAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFtpEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFtpExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFtpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileFtpOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFtpQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileHttp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenObjectAntivirusProfileHttpArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenObjectAntivirusProfileHttpArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenObjectAntivirusProfileHttpAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "content_disarm"
	if _, ok := i["content-disarm"]; ok {
		result["content_disarm"] = flattenObjectAntivirusProfileHttpContentDisarm(i["content-disarm"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenObjectAntivirusProfileHttpEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenObjectAntivirusProfileHttpExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectAntivirusProfileHttpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenObjectAntivirusProfileHttpOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenObjectAntivirusProfileHttpQuarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectAntivirusProfileHttpArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileHttpArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileHttpAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileHttpContentDisarm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileHttpEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileHttpExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileHttpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileHttpOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileHttpQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileImap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenObjectAntivirusProfileImapArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenObjectAntivirusProfileImapArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenObjectAntivirusProfileImapAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "content_disarm"
	if _, ok := i["content-disarm"]; ok {
		result["content_disarm"] = flattenObjectAntivirusProfileImapContentDisarm(i["content-disarm"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenObjectAntivirusProfileImapEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "executables"
	if _, ok := i["executables"]; ok {
		result["executables"] = flattenObjectAntivirusProfileImapExecutables(i["executables"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenObjectAntivirusProfileImapExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectAntivirusProfileImapOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenObjectAntivirusProfileImapOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenObjectAntivirusProfileImapQuarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectAntivirusProfileImapArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileImapArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileImapAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileImapContentDisarm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileImapEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileImapExecutables(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileImapExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileImapOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileImapOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileImapQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileMapi(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenObjectAntivirusProfileMapiArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenObjectAntivirusProfileMapiArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenObjectAntivirusProfileMapiAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenObjectAntivirusProfileMapiEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "executables"
	if _, ok := i["executables"]; ok {
		result["executables"] = flattenObjectAntivirusProfileMapiExecutables(i["executables"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenObjectAntivirusProfileMapiExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectAntivirusProfileMapiOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenObjectAntivirusProfileMapiOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenObjectAntivirusProfileMapiQuarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectAntivirusProfileMapiArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileMapiArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileMapiAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileMapiEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileMapiExecutables(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileMapiExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileMapiOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileMapiOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileMapiQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileMobileMalwareDb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileNacQuar(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "expiry"
	if _, ok := i["expiry"]; ok {
		result["expiry"] = flattenObjectAntivirusProfileNacQuarExpiry(i["expiry"], d, pre_append)
	}

	pre_append = pre + ".0." + "infected"
	if _, ok := i["infected"]; ok {
		result["infected"] = flattenObjectAntivirusProfileNacQuarInfected(i["infected"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectAntivirusProfileNacQuarLog(i["log"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectAntivirusProfileNacQuarExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileNacQuarInfected(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileNacQuarLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileNntp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenObjectAntivirusProfileNntpArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenObjectAntivirusProfileNntpArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenObjectAntivirusProfileNntpAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenObjectAntivirusProfileNntpEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenObjectAntivirusProfileNntpExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectAntivirusProfileNntpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenObjectAntivirusProfileNntpOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenObjectAntivirusProfileNntpQuarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectAntivirusProfileNntpArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileNntpArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileNntpAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileNntpEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileNntpExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileNntpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileNntpOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileNntpQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileOutbreakPreventionArchiveScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenObjectAntivirusProfileOutbreakPreventionExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "ftgd_service"
	if _, ok := i["ftgd-service"]; ok {
		result["ftgd_service"] = flattenObjectAntivirusProfileOutbreakPreventionFtgdService(i["ftgd-service"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectAntivirusProfileOutbreakPreventionExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileOutbreakPreventionFtgdService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfilePop3(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenObjectAntivirusProfilePop3ArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenObjectAntivirusProfilePop3ArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenObjectAntivirusProfilePop3AvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "content_disarm"
	if _, ok := i["content-disarm"]; ok {
		result["content_disarm"] = flattenObjectAntivirusProfilePop3ContentDisarm(i["content-disarm"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenObjectAntivirusProfilePop3Emulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "executables"
	if _, ok := i["executables"]; ok {
		result["executables"] = flattenObjectAntivirusProfilePop3Executables(i["executables"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenObjectAntivirusProfilePop3ExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectAntivirusProfilePop3Options(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenObjectAntivirusProfilePop3OutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenObjectAntivirusProfilePop3Quarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectAntivirusProfilePop3ArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfilePop3ArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfilePop3AvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfilePop3ContentDisarm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfilePop3Emulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfilePop3Executables(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfilePop3ExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfilePop3Options(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfilePop3OutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfilePop3Quarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileScanMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSmtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenObjectAntivirusProfileSmtpArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenObjectAntivirusProfileSmtpArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenObjectAntivirusProfileSmtpAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "content_disarm"
	if _, ok := i["content-disarm"]; ok {
		result["content_disarm"] = flattenObjectAntivirusProfileSmtpContentDisarm(i["content-disarm"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenObjectAntivirusProfileSmtpEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "executables"
	if _, ok := i["executables"]; ok {
		result["executables"] = flattenObjectAntivirusProfileSmtpExecutables(i["executables"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenObjectAntivirusProfileSmtpExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectAntivirusProfileSmtpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenObjectAntivirusProfileSmtpOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenObjectAntivirusProfileSmtpQuarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectAntivirusProfileSmtpArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileSmtpArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileSmtpAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSmtpContentDisarm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSmtpEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSmtpExecutables(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSmtpExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSmtpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileSmtpOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSmtpQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSsh(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenObjectAntivirusProfileSshArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenObjectAntivirusProfileSshArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenObjectAntivirusProfileSshAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenObjectAntivirusProfileSshEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenObjectAntivirusProfileSshExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectAntivirusProfileSshOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenObjectAntivirusProfileSshOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenObjectAntivirusProfileSshQuarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectAntivirusProfileSshArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileSshArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileSshAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSshEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSshExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSshOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileSshOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSshQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectAntivirusProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("analytics_accept_filetype", flattenObjectAntivirusProfileAnalyticsAcceptFiletype(o["analytics-accept-filetype"], d, "analytics_accept_filetype")); err != nil {
		if vv, ok := fortiAPIPatch(o["analytics-accept-filetype"], "ObjectAntivirusProfile-AnalyticsAcceptFiletype"); ok {
			if err = d.Set("analytics_accept_filetype", vv); err != nil {
				return fmt.Errorf("Error reading analytics_accept_filetype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading analytics_accept_filetype: %v", err)
		}
	}

	if err = d.Set("analytics_bl_filetype", flattenObjectAntivirusProfileAnalyticsBlFiletype(o["analytics-bl-filetype"], d, "analytics_bl_filetype")); err != nil {
		if vv, ok := fortiAPIPatch(o["analytics-bl-filetype"], "ObjectAntivirusProfile-AnalyticsBlFiletype"); ok {
			if err = d.Set("analytics_bl_filetype", vv); err != nil {
				return fmt.Errorf("Error reading analytics_bl_filetype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading analytics_bl_filetype: %v", err)
		}
	}

	if err = d.Set("analytics_db", flattenObjectAntivirusProfileAnalyticsDb(o["analytics-db"], d, "analytics_db")); err != nil {
		if vv, ok := fortiAPIPatch(o["analytics-db"], "ObjectAntivirusProfile-AnalyticsDb"); ok {
			if err = d.Set("analytics_db", vv); err != nil {
				return fmt.Errorf("Error reading analytics_db: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading analytics_db: %v", err)
		}
	}

	if err = d.Set("analytics_ignore_filetype", flattenObjectAntivirusProfileAnalyticsIgnoreFiletype(o["analytics-ignore-filetype"], d, "analytics_ignore_filetype")); err != nil {
		if vv, ok := fortiAPIPatch(o["analytics-ignore-filetype"], "ObjectAntivirusProfile-AnalyticsIgnoreFiletype"); ok {
			if err = d.Set("analytics_ignore_filetype", vv); err != nil {
				return fmt.Errorf("Error reading analytics_ignore_filetype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading analytics_ignore_filetype: %v", err)
		}
	}

	if err = d.Set("analytics_max_upload", flattenObjectAntivirusProfileAnalyticsMaxUpload(o["analytics-max-upload"], d, "analytics_max_upload")); err != nil {
		if vv, ok := fortiAPIPatch(o["analytics-max-upload"], "ObjectAntivirusProfile-AnalyticsMaxUpload"); ok {
			if err = d.Set("analytics_max_upload", vv); err != nil {
				return fmt.Errorf("Error reading analytics_max_upload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading analytics_max_upload: %v", err)
		}
	}

	if err = d.Set("analytics_wl_filetype", flattenObjectAntivirusProfileAnalyticsWlFiletype(o["analytics-wl-filetype"], d, "analytics_wl_filetype")); err != nil {
		if vv, ok := fortiAPIPatch(o["analytics-wl-filetype"], "ObjectAntivirusProfile-AnalyticsWlFiletype"); ok {
			if err = d.Set("analytics_wl_filetype", vv); err != nil {
				return fmt.Errorf("Error reading analytics_wl_filetype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading analytics_wl_filetype: %v", err)
		}
	}

	if err = d.Set("av_block_log", flattenObjectAntivirusProfileAvBlockLog(o["av-block-log"], d, "av_block_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-block-log"], "ObjectAntivirusProfile-AvBlockLog"); ok {
			if err = d.Set("av_block_log", vv); err != nil {
				return fmt.Errorf("Error reading av_block_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_block_log: %v", err)
		}
	}

	if err = d.Set("av_virus_log", flattenObjectAntivirusProfileAvVirusLog(o["av-virus-log"], d, "av_virus_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-virus-log"], "ObjectAntivirusProfile-AvVirusLog"); ok {
			if err = d.Set("av_virus_log", vv); err != nil {
				return fmt.Errorf("Error reading av_virus_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_virus_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("cifs", flattenObjectAntivirusProfileCifs(o["cifs"], d, "cifs")); err != nil {
			if vv, ok := fortiAPIPatch(o["cifs"], "ObjectAntivirusProfile-Cifs"); ok {
				if err = d.Set("cifs", vv); err != nil {
					return fmt.Errorf("Error reading cifs: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading cifs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cifs"); ok {
			if err = d.Set("cifs", flattenObjectAntivirusProfileCifs(o["cifs"], d, "cifs")); err != nil {
				if vv, ok := fortiAPIPatch(o["cifs"], "ObjectAntivirusProfile-Cifs"); ok {
					if err = d.Set("cifs", vv); err != nil {
						return fmt.Errorf("Error reading cifs: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading cifs: %v", err)
				}
			}
		}
	}

	if err = d.Set("comment", flattenObjectAntivirusProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectAntivirusProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("content_disarm", flattenObjectAntivirusProfileContentDisarm(o["content-disarm"], d, "content_disarm")); err != nil {
			if vv, ok := fortiAPIPatch(o["content-disarm"], "ObjectAntivirusProfile-ContentDisarm"); ok {
				if err = d.Set("content_disarm", vv); err != nil {
					return fmt.Errorf("Error reading content_disarm: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading content_disarm: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("content_disarm"); ok {
			if err = d.Set("content_disarm", flattenObjectAntivirusProfileContentDisarm(o["content-disarm"], d, "content_disarm")); err != nil {
				if vv, ok := fortiAPIPatch(o["content-disarm"], "ObjectAntivirusProfile-ContentDisarm"); ok {
					if err = d.Set("content_disarm", vv); err != nil {
						return fmt.Errorf("Error reading content_disarm: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading content_disarm: %v", err)
				}
			}
		}
	}

	if err = d.Set("ems_threat_feed", flattenObjectAntivirusProfileEmsThreatFeed(o["ems-threat-feed"], d, "ems_threat_feed")); err != nil {
		if vv, ok := fortiAPIPatch(o["ems-threat-feed"], "ObjectAntivirusProfile-EmsThreatFeed"); ok {
			if err = d.Set("ems_threat_feed", vv); err != nil {
				return fmt.Errorf("Error reading ems_threat_feed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ems_threat_feed: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenObjectAntivirusProfileExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "ObjectAntivirusProfile-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("external_blocklist", flattenObjectAntivirusProfileExternalBlocklist(o["external-blocklist"], d, "external_blocklist")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-blocklist"], "ObjectAntivirusProfile-ExternalBlocklist"); ok {
			if err = d.Set("external_blocklist", vv); err != nil {
				return fmt.Errorf("Error reading external_blocklist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_blocklist: %v", err)
		}
	}

	if err = d.Set("external_blocklist_archive_scan", flattenObjectAntivirusProfileExternalBlocklistArchiveScan(o["external-blocklist-archive-scan"], d, "external_blocklist_archive_scan")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-blocklist-archive-scan"], "ObjectAntivirusProfile-ExternalBlocklistArchiveScan"); ok {
			if err = d.Set("external_blocklist_archive_scan", vv); err != nil {
				return fmt.Errorf("Error reading external_blocklist_archive_scan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_blocklist_archive_scan: %v", err)
		}
	}

	if err = d.Set("external_blocklist_enable_all", flattenObjectAntivirusProfileExternalBlocklistEnableAll(o["external-blocklist-enable-all"], d, "external_blocklist_enable_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-blocklist-enable-all"], "ObjectAntivirusProfile-ExternalBlocklistEnableAll"); ok {
			if err = d.Set("external_blocklist_enable_all", vv); err != nil {
				return fmt.Errorf("Error reading external_blocklist_enable_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_blocklist_enable_all: %v", err)
		}
	}

	if err = d.Set("feature_set", flattenObjectAntivirusProfileFeatureSet(o["feature-set"], d, "feature_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["feature-set"], "ObjectAntivirusProfile-FeatureSet"); ok {
			if err = d.Set("feature_set", vv); err != nil {
				return fmt.Errorf("Error reading feature_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if err = d.Set("ftgd_analytics", flattenObjectAntivirusProfileFtgdAnalytics(o["ftgd-analytics"], d, "ftgd_analytics")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftgd-analytics"], "ObjectAntivirusProfile-FtgdAnalytics"); ok {
			if err = d.Set("ftgd_analytics", vv); err != nil {
				return fmt.Errorf("Error reading ftgd_analytics: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftgd_analytics: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ftp", flattenObjectAntivirusProfileFtp(o["ftp"], d, "ftp")); err != nil {
			if vv, ok := fortiAPIPatch(o["ftp"], "ObjectAntivirusProfile-Ftp"); ok {
				if err = d.Set("ftp", vv); err != nil {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ftp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftp"); ok {
			if err = d.Set("ftp", flattenObjectAntivirusProfileFtp(o["ftp"], d, "ftp")); err != nil {
				if vv, ok := fortiAPIPatch(o["ftp"], "ObjectAntivirusProfile-Ftp"); ok {
					if err = d.Set("ftp", vv); err != nil {
						return fmt.Errorf("Error reading ftp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("http", flattenObjectAntivirusProfileHttp(o["http"], d, "http")); err != nil {
			if vv, ok := fortiAPIPatch(o["http"], "ObjectAntivirusProfile-Http"); ok {
				if err = d.Set("http", vv); err != nil {
					return fmt.Errorf("Error reading http: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading http: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("http"); ok {
			if err = d.Set("http", flattenObjectAntivirusProfileHttp(o["http"], d, "http")); err != nil {
				if vv, ok := fortiAPIPatch(o["http"], "ObjectAntivirusProfile-Http"); ok {
					if err = d.Set("http", vv); err != nil {
						return fmt.Errorf("Error reading http: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading http: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("imap", flattenObjectAntivirusProfileImap(o["imap"], d, "imap")); err != nil {
			if vv, ok := fortiAPIPatch(o["imap"], "ObjectAntivirusProfile-Imap"); ok {
				if err = d.Set("imap", vv); err != nil {
					return fmt.Errorf("Error reading imap: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading imap: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("imap"); ok {
			if err = d.Set("imap", flattenObjectAntivirusProfileImap(o["imap"], d, "imap")); err != nil {
				if vv, ok := fortiAPIPatch(o["imap"], "ObjectAntivirusProfile-Imap"); ok {
					if err = d.Set("imap", vv); err != nil {
						return fmt.Errorf("Error reading imap: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading imap: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mapi", flattenObjectAntivirusProfileMapi(o["mapi"], d, "mapi")); err != nil {
			if vv, ok := fortiAPIPatch(o["mapi"], "ObjectAntivirusProfile-Mapi"); ok {
				if err = d.Set("mapi", vv); err != nil {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mapi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mapi"); ok {
			if err = d.Set("mapi", flattenObjectAntivirusProfileMapi(o["mapi"], d, "mapi")); err != nil {
				if vv, ok := fortiAPIPatch(o["mapi"], "ObjectAntivirusProfile-Mapi"); ok {
					if err = d.Set("mapi", vv); err != nil {
						return fmt.Errorf("Error reading mapi: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			}
		}
	}

	if err = d.Set("mobile_malware_db", flattenObjectAntivirusProfileMobileMalwareDb(o["mobile-malware-db"], d, "mobile_malware_db")); err != nil {
		if vv, ok := fortiAPIPatch(o["mobile-malware-db"], "ObjectAntivirusProfile-MobileMalwareDb"); ok {
			if err = d.Set("mobile_malware_db", vv); err != nil {
				return fmt.Errorf("Error reading mobile_malware_db: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mobile_malware_db: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("nac_quar", flattenObjectAntivirusProfileNacQuar(o["nac-quar"], d, "nac_quar")); err != nil {
			if vv, ok := fortiAPIPatch(o["nac-quar"], "ObjectAntivirusProfile-NacQuar"); ok {
				if err = d.Set("nac_quar", vv); err != nil {
					return fmt.Errorf("Error reading nac_quar: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nac_quar: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nac_quar"); ok {
			if err = d.Set("nac_quar", flattenObjectAntivirusProfileNacQuar(o["nac-quar"], d, "nac_quar")); err != nil {
				if vv, ok := fortiAPIPatch(o["nac-quar"], "ObjectAntivirusProfile-NacQuar"); ok {
					if err = d.Set("nac_quar", vv); err != nil {
						return fmt.Errorf("Error reading nac_quar: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nac_quar: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectAntivirusProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectAntivirusProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("nntp", flattenObjectAntivirusProfileNntp(o["nntp"], d, "nntp")); err != nil {
			if vv, ok := fortiAPIPatch(o["nntp"], "ObjectAntivirusProfile-Nntp"); ok {
				if err = d.Set("nntp", vv); err != nil {
					return fmt.Errorf("Error reading nntp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nntp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nntp"); ok {
			if err = d.Set("nntp", flattenObjectAntivirusProfileNntp(o["nntp"], d, "nntp")); err != nil {
				if vv, ok := fortiAPIPatch(o["nntp"], "ObjectAntivirusProfile-Nntp"); ok {
					if err = d.Set("nntp", vv); err != nil {
						return fmt.Errorf("Error reading nntp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nntp: %v", err)
				}
			}
		}
	}

	if err = d.Set("outbreak_prevention_archive_scan", flattenObjectAntivirusProfileOutbreakPreventionArchiveScan(o["outbreak-prevention-archive-scan"], d, "outbreak_prevention_archive_scan")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-archive-scan"], "ObjectAntivirusProfile-OutbreakPreventionArchiveScan"); ok {
			if err = d.Set("outbreak_prevention_archive_scan", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_archive_scan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_archive_scan: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("outbreak_prevention", flattenObjectAntivirusProfileOutbreakPrevention(o["outbreak-prevention"], d, "outbreak_prevention")); err != nil {
			if vv, ok := fortiAPIPatch(o["outbreak-prevention"], "ObjectAntivirusProfile-OutbreakPrevention"); ok {
				if err = d.Set("outbreak_prevention", vv); err != nil {
					return fmt.Errorf("Error reading outbreak_prevention: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading outbreak_prevention: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("outbreak_prevention"); ok {
			if err = d.Set("outbreak_prevention", flattenObjectAntivirusProfileOutbreakPrevention(o["outbreak-prevention"], d, "outbreak_prevention")); err != nil {
				if vv, ok := fortiAPIPatch(o["outbreak-prevention"], "ObjectAntivirusProfile-OutbreakPrevention"); ok {
					if err = d.Set("outbreak_prevention", vv); err != nil {
						return fmt.Errorf("Error reading outbreak_prevention: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading outbreak_prevention: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("pop3", flattenObjectAntivirusProfilePop3(o["pop3"], d, "pop3")); err != nil {
			if vv, ok := fortiAPIPatch(o["pop3"], "ObjectAntivirusProfile-Pop3"); ok {
				if err = d.Set("pop3", vv); err != nil {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pop3: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pop3"); ok {
			if err = d.Set("pop3", flattenObjectAntivirusProfilePop3(o["pop3"], d, "pop3")); err != nil {
				if vv, ok := fortiAPIPatch(o["pop3"], "ObjectAntivirusProfile-Pop3"); ok {
					if err = d.Set("pop3", vv); err != nil {
						return fmt.Errorf("Error reading pop3: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			}
		}
	}

	if err = d.Set("replacemsg_group", flattenObjectAntivirusProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "ObjectAntivirusProfile-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("scan_mode", flattenObjectAntivirusProfileScanMode(o["scan-mode"], d, "scan_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-mode"], "ObjectAntivirusProfile-ScanMode"); ok {
			if err = d.Set("scan_mode", vv); err != nil {
				return fmt.Errorf("Error reading scan_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("smtp", flattenObjectAntivirusProfileSmtp(o["smtp"], d, "smtp")); err != nil {
			if vv, ok := fortiAPIPatch(o["smtp"], "ObjectAntivirusProfile-Smtp"); ok {
				if err = d.Set("smtp", vv); err != nil {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading smtp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("smtp"); ok {
			if err = d.Set("smtp", flattenObjectAntivirusProfileSmtp(o["smtp"], d, "smtp")); err != nil {
				if vv, ok := fortiAPIPatch(o["smtp"], "ObjectAntivirusProfile-Smtp"); ok {
					if err = d.Set("smtp", vv); err != nil {
						return fmt.Errorf("Error reading smtp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ssh", flattenObjectAntivirusProfileSsh(o["ssh"], d, "ssh")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssh"], "ObjectAntivirusProfile-Ssh"); ok {
				if err = d.Set("ssh", vv); err != nil {
					return fmt.Errorf("Error reading ssh: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssh: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssh"); ok {
			if err = d.Set("ssh", flattenObjectAntivirusProfileSsh(o["ssh"], d, "ssh")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssh"], "ObjectAntivirusProfile-Ssh"); ok {
					if err = d.Set("ssh", vv); err != nil {
						return fmt.Errorf("Error reading ssh: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssh: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectAntivirusProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectAntivirusProfileAnalyticsAcceptFiletype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileAnalyticsBlFiletype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileAnalyticsDb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileAnalyticsIgnoreFiletype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileAnalyticsMaxUpload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileAnalyticsWlFiletype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileAvBlockLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileAvVirusLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileCifs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-block"], _ = expandObjectAntivirusProfileCifsArchiveBlock(d, i["archive_block"], pre_append)
	} else {
		result["archive-block"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-log"], _ = expandObjectAntivirusProfileCifsArchiveLog(d, i["archive_log"], pre_append)
	} else {
		result["archive-log"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok {
		result["av-scan"], _ = expandObjectAntivirusProfileCifsAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok {
		result["emulator"], _ = expandObjectAntivirusProfileCifsEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok {
		result["external-blocklist"], _ = expandObjectAntivirusProfileCifsExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectAntivirusProfileCifsOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileCifsOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok {
		result["quarantine"], _ = expandObjectAntivirusProfileCifsQuarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandObjectAntivirusProfileCifsArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileCifsArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileCifsAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileCifsEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileCifsExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileCifsOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileCifsOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileCifsQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cover_page"
	if _, ok := d.GetOk(pre_append); ok {
		result["cover-page"], _ = expandObjectAntivirusProfileContentDisarmCoverPage(d, i["cover_page"], pre_append)
	}
	pre_append = pre + ".0." + "detect_only"
	if _, ok := d.GetOk(pre_append); ok {
		result["detect-only"], _ = expandObjectAntivirusProfileContentDisarmDetectOnly(d, i["detect_only"], pre_append)
	}
	pre_append = pre + ".0." + "error_action"
	if _, ok := d.GetOk(pre_append); ok {
		result["error-action"], _ = expandObjectAntivirusProfileContentDisarmErrorAction(d, i["error_action"], pre_append)
	}
	pre_append = pre + ".0." + "office_action"
	if _, ok := d.GetOk(pre_append); ok {
		result["office-action"], _ = expandObjectAntivirusProfileContentDisarmOfficeAction(d, i["office_action"], pre_append)
	}
	pre_append = pre + ".0." + "office_dde"
	if _, ok := d.GetOk(pre_append); ok {
		result["office-dde"], _ = expandObjectAntivirusProfileContentDisarmOfficeDde(d, i["office_dde"], pre_append)
	}
	pre_append = pre + ".0." + "office_embed"
	if _, ok := d.GetOk(pre_append); ok {
		result["office-embed"], _ = expandObjectAntivirusProfileContentDisarmOfficeEmbed(d, i["office_embed"], pre_append)
	}
	pre_append = pre + ".0." + "office_hylink"
	if _, ok := d.GetOk(pre_append); ok {
		result["office-hylink"], _ = expandObjectAntivirusProfileContentDisarmOfficeHylink(d, i["office_hylink"], pre_append)
	}
	pre_append = pre + ".0." + "office_linked"
	if _, ok := d.GetOk(pre_append); ok {
		result["office-linked"], _ = expandObjectAntivirusProfileContentDisarmOfficeLinked(d, i["office_linked"], pre_append)
	}
	pre_append = pre + ".0." + "office_macro"
	if _, ok := d.GetOk(pre_append); ok {
		result["office-macro"], _ = expandObjectAntivirusProfileContentDisarmOfficeMacro(d, i["office_macro"], pre_append)
	}
	pre_append = pre + ".0." + "original_file_destination"
	if _, ok := d.GetOk(pre_append); ok {
		result["original-file-destination"], _ = expandObjectAntivirusProfileContentDisarmOriginalFileDestination(d, i["original_file_destination"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_form"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdf-act-form"], _ = expandObjectAntivirusProfileContentDisarmPdfActForm(d, i["pdf_act_form"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_gotor"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdf-act-gotor"], _ = expandObjectAntivirusProfileContentDisarmPdfActGotor(d, i["pdf_act_gotor"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_java"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdf-act-java"], _ = expandObjectAntivirusProfileContentDisarmPdfActJava(d, i["pdf_act_java"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_launch"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdf-act-launch"], _ = expandObjectAntivirusProfileContentDisarmPdfActLaunch(d, i["pdf_act_launch"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_movie"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdf-act-movie"], _ = expandObjectAntivirusProfileContentDisarmPdfActMovie(d, i["pdf_act_movie"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_sound"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdf-act-sound"], _ = expandObjectAntivirusProfileContentDisarmPdfActSound(d, i["pdf_act_sound"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_embedfile"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdf-embedfile"], _ = expandObjectAntivirusProfileContentDisarmPdfEmbedfile(d, i["pdf_embedfile"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_hyperlink"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdf-hyperlink"], _ = expandObjectAntivirusProfileContentDisarmPdfHyperlink(d, i["pdf_hyperlink"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_javacode"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdf-javacode"], _ = expandObjectAntivirusProfileContentDisarmPdfJavacode(d, i["pdf_javacode"], pre_append)
	}

	return result, nil
}

func expandObjectAntivirusProfileContentDisarmCoverPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmDetectOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmErrorAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmOfficeAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmOfficeDde(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmOfficeEmbed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmOfficeHylink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmOfficeLinked(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmOfficeMacro(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmOriginalFileDestination(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfActForm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfActGotor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfActJava(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfActLaunch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfActMovie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfActSound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfEmbedfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfHyperlink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileContentDisarmPdfJavacode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileEmsThreatFeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectAntivirusProfileExternalBlocklistArchiveScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileExternalBlocklistEnableAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFeatureSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFtgdAnalytics(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-block"], _ = expandObjectAntivirusProfileFtpArchiveBlock(d, i["archive_block"], pre_append)
	} else {
		result["archive-block"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-log"], _ = expandObjectAntivirusProfileFtpArchiveLog(d, i["archive_log"], pre_append)
	} else {
		result["archive-log"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok {
		result["av-scan"], _ = expandObjectAntivirusProfileFtpAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok {
		result["emulator"], _ = expandObjectAntivirusProfileFtpEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok {
		result["external-blocklist"], _ = expandObjectAntivirusProfileFtpExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectAntivirusProfileFtpOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileFtpOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok {
		result["quarantine"], _ = expandObjectAntivirusProfileFtpQuarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandObjectAntivirusProfileFtpArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileFtpArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileFtpAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFtpEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFtpExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFtpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileFtpOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFtpQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-block"], _ = expandObjectAntivirusProfileHttpArchiveBlock(d, i["archive_block"], pre_append)
	} else {
		result["archive-block"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-log"], _ = expandObjectAntivirusProfileHttpArchiveLog(d, i["archive_log"], pre_append)
	} else {
		result["archive-log"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok {
		result["av-scan"], _ = expandObjectAntivirusProfileHttpAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "content_disarm"
	if _, ok := d.GetOk(pre_append); ok {
		result["content-disarm"], _ = expandObjectAntivirusProfileHttpContentDisarm(d, i["content_disarm"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok {
		result["emulator"], _ = expandObjectAntivirusProfileHttpEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok {
		result["external-blocklist"], _ = expandObjectAntivirusProfileHttpExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectAntivirusProfileHttpOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileHttpOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok {
		result["quarantine"], _ = expandObjectAntivirusProfileHttpQuarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandObjectAntivirusProfileHttpArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileHttpArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileHttpAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileHttpContentDisarm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileHttpEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileHttpExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileHttpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileHttpOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileHttpQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileImap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-block"], _ = expandObjectAntivirusProfileImapArchiveBlock(d, i["archive_block"], pre_append)
	} else {
		result["archive-block"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-log"], _ = expandObjectAntivirusProfileImapArchiveLog(d, i["archive_log"], pre_append)
	} else {
		result["archive-log"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok {
		result["av-scan"], _ = expandObjectAntivirusProfileImapAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "content_disarm"
	if _, ok := d.GetOk(pre_append); ok {
		result["content-disarm"], _ = expandObjectAntivirusProfileImapContentDisarm(d, i["content_disarm"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok {
		result["emulator"], _ = expandObjectAntivirusProfileImapEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "executables"
	if _, ok := d.GetOk(pre_append); ok {
		result["executables"], _ = expandObjectAntivirusProfileImapExecutables(d, i["executables"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok {
		result["external-blocklist"], _ = expandObjectAntivirusProfileImapExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectAntivirusProfileImapOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileImapOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok {
		result["quarantine"], _ = expandObjectAntivirusProfileImapQuarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandObjectAntivirusProfileImapArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileImapArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileImapAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileImapContentDisarm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileImapEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileImapExecutables(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileImapExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileImapOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileImapOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileImapQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileMapi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-block"], _ = expandObjectAntivirusProfileMapiArchiveBlock(d, i["archive_block"], pre_append)
	} else {
		result["archive-block"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-log"], _ = expandObjectAntivirusProfileMapiArchiveLog(d, i["archive_log"], pre_append)
	} else {
		result["archive-log"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok {
		result["av-scan"], _ = expandObjectAntivirusProfileMapiAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok {
		result["emulator"], _ = expandObjectAntivirusProfileMapiEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "executables"
	if _, ok := d.GetOk(pre_append); ok {
		result["executables"], _ = expandObjectAntivirusProfileMapiExecutables(d, i["executables"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok {
		result["external-blocklist"], _ = expandObjectAntivirusProfileMapiExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectAntivirusProfileMapiOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileMapiOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok {
		result["quarantine"], _ = expandObjectAntivirusProfileMapiQuarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandObjectAntivirusProfileMapiArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileMapiArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileMapiAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileMapiEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileMapiExecutables(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileMapiExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileMapiOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileMapiOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileMapiQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileMobileMalwareDb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileNacQuar(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "expiry"
	if _, ok := d.GetOk(pre_append); ok {
		result["expiry"], _ = expandObjectAntivirusProfileNacQuarExpiry(d, i["expiry"], pre_append)
	}
	pre_append = pre + ".0." + "infected"
	if _, ok := d.GetOk(pre_append); ok {
		result["infected"], _ = expandObjectAntivirusProfileNacQuarInfected(d, i["infected"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandObjectAntivirusProfileNacQuarLog(d, i["log"], pre_append)
	}

	return result, nil
}

func expandObjectAntivirusProfileNacQuarExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileNacQuarInfected(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileNacQuarLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileNntp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-block"], _ = expandObjectAntivirusProfileNntpArchiveBlock(d, i["archive_block"], pre_append)
	} else {
		result["archive-block"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-log"], _ = expandObjectAntivirusProfileNntpArchiveLog(d, i["archive_log"], pre_append)
	} else {
		result["archive-log"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok {
		result["av-scan"], _ = expandObjectAntivirusProfileNntpAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok {
		result["emulator"], _ = expandObjectAntivirusProfileNntpEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok {
		result["external-blocklist"], _ = expandObjectAntivirusProfileNntpExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectAntivirusProfileNntpOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileNntpOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok {
		result["quarantine"], _ = expandObjectAntivirusProfileNntpQuarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandObjectAntivirusProfileNntpArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileNntpArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileNntpAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileNntpEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileNntpExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileNntpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileNntpOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileNntpQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileOutbreakPreventionArchiveScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok {
		result["external-blocklist"], _ = expandObjectAntivirusProfileOutbreakPreventionExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "ftgd_service"
	if _, ok := d.GetOk(pre_append); ok {
		result["ftgd-service"], _ = expandObjectAntivirusProfileOutbreakPreventionFtgdService(d, i["ftgd_service"], pre_append)
	}

	return result, nil
}

func expandObjectAntivirusProfileOutbreakPreventionExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileOutbreakPreventionFtgdService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfilePop3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-block"], _ = expandObjectAntivirusProfilePop3ArchiveBlock(d, i["archive_block"], pre_append)
	} else {
		result["archive-block"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-log"], _ = expandObjectAntivirusProfilePop3ArchiveLog(d, i["archive_log"], pre_append)
	} else {
		result["archive-log"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok {
		result["av-scan"], _ = expandObjectAntivirusProfilePop3AvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "content_disarm"
	if _, ok := d.GetOk(pre_append); ok {
		result["content-disarm"], _ = expandObjectAntivirusProfilePop3ContentDisarm(d, i["content_disarm"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok {
		result["emulator"], _ = expandObjectAntivirusProfilePop3Emulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "executables"
	if _, ok := d.GetOk(pre_append); ok {
		result["executables"], _ = expandObjectAntivirusProfilePop3Executables(d, i["executables"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok {
		result["external-blocklist"], _ = expandObjectAntivirusProfilePop3ExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectAntivirusProfilePop3Options(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfilePop3OutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok {
		result["quarantine"], _ = expandObjectAntivirusProfilePop3Quarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandObjectAntivirusProfilePop3ArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfilePop3ArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfilePop3AvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfilePop3ContentDisarm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfilePop3Emulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfilePop3Executables(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfilePop3ExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfilePop3Options(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfilePop3OutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfilePop3Quarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileScanMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSmtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-block"], _ = expandObjectAntivirusProfileSmtpArchiveBlock(d, i["archive_block"], pre_append)
	} else {
		result["archive-block"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-log"], _ = expandObjectAntivirusProfileSmtpArchiveLog(d, i["archive_log"], pre_append)
	} else {
		result["archive-log"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok {
		result["av-scan"], _ = expandObjectAntivirusProfileSmtpAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "content_disarm"
	if _, ok := d.GetOk(pre_append); ok {
		result["content-disarm"], _ = expandObjectAntivirusProfileSmtpContentDisarm(d, i["content_disarm"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok {
		result["emulator"], _ = expandObjectAntivirusProfileSmtpEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "executables"
	if _, ok := d.GetOk(pre_append); ok {
		result["executables"], _ = expandObjectAntivirusProfileSmtpExecutables(d, i["executables"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok {
		result["external-blocklist"], _ = expandObjectAntivirusProfileSmtpExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectAntivirusProfileSmtpOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileSmtpOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok {
		result["quarantine"], _ = expandObjectAntivirusProfileSmtpQuarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandObjectAntivirusProfileSmtpArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileSmtpArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileSmtpAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSmtpContentDisarm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSmtpEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSmtpExecutables(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSmtpExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSmtpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileSmtpOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSmtpQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSsh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-block"], _ = expandObjectAntivirusProfileSshArchiveBlock(d, i["archive_block"], pre_append)
	} else {
		result["archive-block"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok {
		result["archive-log"], _ = expandObjectAntivirusProfileSshArchiveLog(d, i["archive_log"], pre_append)
	} else {
		result["archive-log"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok {
		result["av-scan"], _ = expandObjectAntivirusProfileSshAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok {
		result["emulator"], _ = expandObjectAntivirusProfileSshEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok {
		result["external-blocklist"], _ = expandObjectAntivirusProfileSshExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectAntivirusProfileSshOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileSshOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok {
		result["quarantine"], _ = expandObjectAntivirusProfileSshQuarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandObjectAntivirusProfileSshArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileSshArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileSshAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSshEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSshExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSshOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileSshOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSshQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectAntivirusProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("analytics_accept_filetype"); ok {
		t, err := expandObjectAntivirusProfileAnalyticsAcceptFiletype(d, v, "analytics_accept_filetype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-accept-filetype"] = t
		}
	}

	if v, ok := d.GetOk("analytics_bl_filetype"); ok {
		t, err := expandObjectAntivirusProfileAnalyticsBlFiletype(d, v, "analytics_bl_filetype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-bl-filetype"] = t
		}
	}

	if v, ok := d.GetOk("analytics_db"); ok {
		t, err := expandObjectAntivirusProfileAnalyticsDb(d, v, "analytics_db")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-db"] = t
		}
	}

	if v, ok := d.GetOk("analytics_ignore_filetype"); ok {
		t, err := expandObjectAntivirusProfileAnalyticsIgnoreFiletype(d, v, "analytics_ignore_filetype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-ignore-filetype"] = t
		}
	}

	if v, ok := d.GetOk("analytics_max_upload"); ok {
		t, err := expandObjectAntivirusProfileAnalyticsMaxUpload(d, v, "analytics_max_upload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-max-upload"] = t
		}
	}

	if v, ok := d.GetOk("analytics_wl_filetype"); ok {
		t, err := expandObjectAntivirusProfileAnalyticsWlFiletype(d, v, "analytics_wl_filetype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-wl-filetype"] = t
		}
	}

	if v, ok := d.GetOk("av_block_log"); ok {
		t, err := expandObjectAntivirusProfileAvBlockLog(d, v, "av_block_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-block-log"] = t
		}
	}

	if v, ok := d.GetOk("av_virus_log"); ok {
		t, err := expandObjectAntivirusProfileAvVirusLog(d, v, "av_virus_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-virus-log"] = t
		}
	}

	if v, ok := d.GetOk("cifs"); ok {
		t, err := expandObjectAntivirusProfileCifs(d, v, "cifs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectAntivirusProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("content_disarm"); ok {
		t, err := expandObjectAntivirusProfileContentDisarm(d, v, "content_disarm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-disarm"] = t
		}
	}

	if v, ok := d.GetOk("ems_threat_feed"); ok {
		t, err := expandObjectAntivirusProfileEmsThreatFeed(d, v, "ems_threat_feed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-threat-feed"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok {
		t, err := expandObjectAntivirusProfileExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("external_blocklist"); ok {
		t, err := expandObjectAntivirusProfileExternalBlocklist(d, v, "external_blocklist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-blocklist"] = t
		}
	}

	if v, ok := d.GetOk("external_blocklist_archive_scan"); ok {
		t, err := expandObjectAntivirusProfileExternalBlocklistArchiveScan(d, v, "external_blocklist_archive_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-blocklist-archive-scan"] = t
		}
	}

	if v, ok := d.GetOk("external_blocklist_enable_all"); ok {
		t, err := expandObjectAntivirusProfileExternalBlocklistEnableAll(d, v, "external_blocklist_enable_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-blocklist-enable-all"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok {
		t, err := expandObjectAntivirusProfileFeatureSet(d, v, "feature_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("ftgd_analytics"); ok {
		t, err := expandObjectAntivirusProfileFtgdAnalytics(d, v, "ftgd_analytics")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftgd-analytics"] = t
		}
	}

	if v, ok := d.GetOk("ftp"); ok {
		t, err := expandObjectAntivirusProfileFtp(d, v, "ftp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp"] = t
		}
	}

	if v, ok := d.GetOk("http"); ok {
		t, err := expandObjectAntivirusProfileHttp(d, v, "http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http"] = t
		}
	}

	if v, ok := d.GetOk("imap"); ok {
		t, err := expandObjectAntivirusProfileImap(d, v, "imap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imap"] = t
		}
	}

	if v, ok := d.GetOk("mapi"); ok {
		t, err := expandObjectAntivirusProfileMapi(d, v, "mapi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi"] = t
		}
	}

	if v, ok := d.GetOk("mobile_malware_db"); ok {
		t, err := expandObjectAntivirusProfileMobileMalwareDb(d, v, "mobile_malware_db")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mobile-malware-db"] = t
		}
	}

	if v, ok := d.GetOk("nac_quar"); ok {
		t, err := expandObjectAntivirusProfileNacQuar(d, v, "nac_quar")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-quar"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectAntivirusProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nntp"); ok {
		t, err := expandObjectAntivirusProfileNntp(d, v, "nntp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nntp"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_archive_scan"); ok {
		t, err := expandObjectAntivirusProfileOutbreakPreventionArchiveScan(d, v, "outbreak_prevention_archive_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-archive-scan"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention"); ok {
		t, err := expandObjectAntivirusProfileOutbreakPrevention(d, v, "outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("pop3"); ok {
		t, err := expandObjectAntivirusProfilePop3(d, v, "pop3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandObjectAntivirusProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("scan_mode"); ok {
		t, err := expandObjectAntivirusProfileScanMode(d, v, "scan_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-mode"] = t
		}
	}

	if v, ok := d.GetOk("smtp"); ok {
		t, err := expandObjectAntivirusProfileSmtp(d, v, "smtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp"] = t
		}
	}

	if v, ok := d.GetOk("ssh"); ok {
		t, err := expandObjectAntivirusProfileSsh(d, v, "ssh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh"] = t
		}
	}

	return &obj, nil
}
