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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			},
			"analytics_bl_filetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"analytics_db": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"analytics_ignore_filetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"analytics_max_upload": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"analytics_wl_filetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
				Computed: true,
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
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortisandbox": &schema.Schema{
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
			},
			"content_disarm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
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
				Type:     schema.TypeSet,
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
			"fortiai_error_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiai_timeout_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortindr_error_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortindr_timeout_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortisandbox_error_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortisandbox_max_upload": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fortisandbox_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortisandbox_timeout_action": &schema.Schema{
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
				Computed: true,
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
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortisandbox": &schema.Schema{
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
				Computed: true,
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
						"av_optimize": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortisandbox": &schema.Schema{
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
						"unknown_content_encoding": &schema.Schema{
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
				Computed: true,
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
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortisandbox": &schema.Schema{
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
			"inspection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mapi": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
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
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"executables": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortisandbox": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
				Computed: true,
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
			},
			"nntp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
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
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortisandbox": &schema.Schema{
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
				Computed: true,
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
				Computed: true,
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
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortisandbox": &schema.Schema{
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
			},
			"scan_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"smb": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
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
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
					},
				},
			},
			"smtp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
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
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortisandbox": &schema.Schema{
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
				Computed: true,
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
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortisandbox": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectAntivirusProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectAntivirusProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectAntivirusProfile(obj, paradict)

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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectAntivirusProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectAntivirusProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectAntivirusProfile(obj, mkey, paradict)
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	err = c.DeleteObjectAntivirusProfile(mkey, paradict)
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	o, err := c.ReadObjectAntivirusProfile(mkey, paradict)
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
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectAntivirusProfileAnalyticsBlFiletype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectAntivirusProfileAnalyticsDb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileAnalyticsIgnoreFiletype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectAntivirusProfileAnalyticsMaxUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileAnalyticsWlFiletype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
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

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenObjectAntivirusProfileCifsFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenObjectAntivirusProfileCifsFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenObjectAntivirusProfileCifsFortisandbox(i["fortisandbox"], d, pre_append)
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

func flattenObjectAntivirusProfileCifsFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileCifsFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileCifsFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectAntivirusProfileFortiaiErrorAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFortiaiTimeoutAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFortindrErrorAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFortindrTimeoutAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFortisandboxErrorAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFortisandboxMaxUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFortisandboxMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFortisandboxTimeoutAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenObjectAntivirusProfileFtpFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenObjectAntivirusProfileFtpFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenObjectAntivirusProfileFtpFortisandbox(i["fortisandbox"], d, pre_append)
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

func flattenObjectAntivirusProfileFtpFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFtpFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileFtpFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	pre_append = pre + ".0." + "av_optimize"
	if _, ok := i["av-optimize"]; ok {
		result["av_optimize"] = flattenObjectAntivirusProfileHttpAvOptimize(i["av-optimize"], d, pre_append)
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

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenObjectAntivirusProfileHttpFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenObjectAntivirusProfileHttpFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenObjectAntivirusProfileHttpFortisandbox(i["fortisandbox"], d, pre_append)
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

	pre_append = pre + ".0." + "unknown_content_encoding"
	if _, ok := i["unknown-content-encoding"]; ok {
		result["unknown_content_encoding"] = flattenObjectAntivirusProfileHttpUnknownContentEncoding(i["unknown-content-encoding"], d, pre_append)
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

func flattenObjectAntivirusProfileHttpAvOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
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

func flattenObjectAntivirusProfileHttpFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileHttpFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileHttpFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectAntivirusProfileHttpUnknownContentEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenObjectAntivirusProfileImapFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenObjectAntivirusProfileImapFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenObjectAntivirusProfileImapFortisandbox(i["fortisandbox"], d, pre_append)
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

func flattenObjectAntivirusProfileImapFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileImapFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileImapFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectAntivirusProfileInspectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenObjectAntivirusProfileMapiFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenObjectAntivirusProfileMapiFortisandbox(i["fortisandbox"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenObjectAntivirusProfileMapiFortiai(i["fortiai"], d, pre_append)
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

func flattenObjectAntivirusProfileMapiFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileMapiFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileMapiFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenObjectAntivirusProfileNntpFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenObjectAntivirusProfileNntpFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenObjectAntivirusProfileNntpFortisandbox(i["fortisandbox"], d, pre_append)
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

func flattenObjectAntivirusProfileNntpFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileNntpFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileNntpFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenObjectAntivirusProfilePop3Fortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenObjectAntivirusProfilePop3Fortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenObjectAntivirusProfilePop3Fortisandbox(i["fortisandbox"], d, pre_append)
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

func flattenObjectAntivirusProfilePop3Fortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfilePop3Fortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfilePop3Fortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
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
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectAntivirusProfileScanMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSmb(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenObjectAntivirusProfileSmbArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenObjectAntivirusProfileSmbArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenObjectAntivirusProfileSmbEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectAntivirusProfileSmbOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenObjectAntivirusProfileSmbOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectAntivirusProfileSmbArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileSmbArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileSmbEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSmbOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAntivirusProfileSmbOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenObjectAntivirusProfileSmtpFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenObjectAntivirusProfileSmtpFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenObjectAntivirusProfileSmtpFortisandbox(i["fortisandbox"], d, pre_append)
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

func flattenObjectAntivirusProfileSmtpFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSmtpFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSmtpFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenObjectAntivirusProfileSshFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenObjectAntivirusProfileSshFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenObjectAntivirusProfileSshFortisandbox(i["fortisandbox"], d, pre_append)
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

func flattenObjectAntivirusProfileSshFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSshFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileSshFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

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

	if err = d.Set("fortiai_error_action", flattenObjectAntivirusProfileFortiaiErrorAction(o["fortiai-error-action"], d, "fortiai_error_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiai-error-action"], "ObjectAntivirusProfile-FortiaiErrorAction"); ok {
			if err = d.Set("fortiai_error_action", vv); err != nil {
				return fmt.Errorf("Error reading fortiai_error_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiai_error_action: %v", err)
		}
	}

	if err = d.Set("fortiai_timeout_action", flattenObjectAntivirusProfileFortiaiTimeoutAction(o["fortiai-timeout-action"], d, "fortiai_timeout_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiai-timeout-action"], "ObjectAntivirusProfile-FortiaiTimeoutAction"); ok {
			if err = d.Set("fortiai_timeout_action", vv); err != nil {
				return fmt.Errorf("Error reading fortiai_timeout_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiai_timeout_action: %v", err)
		}
	}

	if err = d.Set("fortindr_error_action", flattenObjectAntivirusProfileFortindrErrorAction(o["fortindr-error-action"], d, "fortindr_error_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortindr-error-action"], "ObjectAntivirusProfile-FortindrErrorAction"); ok {
			if err = d.Set("fortindr_error_action", vv); err != nil {
				return fmt.Errorf("Error reading fortindr_error_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortindr_error_action: %v", err)
		}
	}

	if err = d.Set("fortindr_timeout_action", flattenObjectAntivirusProfileFortindrTimeoutAction(o["fortindr-timeout-action"], d, "fortindr_timeout_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortindr-timeout-action"], "ObjectAntivirusProfile-FortindrTimeoutAction"); ok {
			if err = d.Set("fortindr_timeout_action", vv); err != nil {
				return fmt.Errorf("Error reading fortindr_timeout_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortindr_timeout_action: %v", err)
		}
	}

	if err = d.Set("fortisandbox_error_action", flattenObjectAntivirusProfileFortisandboxErrorAction(o["fortisandbox-error-action"], d, "fortisandbox_error_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortisandbox-error-action"], "ObjectAntivirusProfile-FortisandboxErrorAction"); ok {
			if err = d.Set("fortisandbox_error_action", vv); err != nil {
				return fmt.Errorf("Error reading fortisandbox_error_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortisandbox_error_action: %v", err)
		}
	}

	if err = d.Set("fortisandbox_max_upload", flattenObjectAntivirusProfileFortisandboxMaxUpload(o["fortisandbox-max-upload"], d, "fortisandbox_max_upload")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortisandbox-max-upload"], "ObjectAntivirusProfile-FortisandboxMaxUpload"); ok {
			if err = d.Set("fortisandbox_max_upload", vv); err != nil {
				return fmt.Errorf("Error reading fortisandbox_max_upload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortisandbox_max_upload: %v", err)
		}
	}

	if err = d.Set("fortisandbox_mode", flattenObjectAntivirusProfileFortisandboxMode(o["fortisandbox-mode"], d, "fortisandbox_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortisandbox-mode"], "ObjectAntivirusProfile-FortisandboxMode"); ok {
			if err = d.Set("fortisandbox_mode", vv); err != nil {
				return fmt.Errorf("Error reading fortisandbox_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortisandbox_mode: %v", err)
		}
	}

	if err = d.Set("fortisandbox_timeout_action", flattenObjectAntivirusProfileFortisandboxTimeoutAction(o["fortisandbox-timeout-action"], d, "fortisandbox_timeout_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortisandbox-timeout-action"], "ObjectAntivirusProfile-FortisandboxTimeoutAction"); ok {
			if err = d.Set("fortisandbox_timeout_action", vv); err != nil {
				return fmt.Errorf("Error reading fortisandbox_timeout_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortisandbox_timeout_action: %v", err)
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

	if err = d.Set("inspection_mode", flattenObjectAntivirusProfileInspectionMode(o["inspection-mode"], d, "inspection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["inspection-mode"], "ObjectAntivirusProfile-InspectionMode"); ok {
			if err = d.Set("inspection_mode", vv); err != nil {
				return fmt.Errorf("Error reading inspection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inspection_mode: %v", err)
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
		if err = d.Set("smb", flattenObjectAntivirusProfileSmb(o["smb"], d, "smb")); err != nil {
			if vv, ok := fortiAPIPatch(o["smb"], "ObjectAntivirusProfile-Smb"); ok {
				if err = d.Set("smb", vv); err != nil {
					return fmt.Errorf("Error reading smb: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading smb: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("smb"); ok {
			if err = d.Set("smb", flattenObjectAntivirusProfileSmb(o["smb"], d, "smb")); err != nil {
				if vv, ok := fortiAPIPatch(o["smb"], "ObjectAntivirusProfile-Smb"); ok {
					if err = d.Set("smb", vv); err != nil {
						return fmt.Errorf("Error reading smb: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading smb: %v", err)
				}
			}
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
	return convstr2list(v, nil), nil
}

func expandObjectAntivirusProfileAnalyticsBlFiletype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectAntivirusProfileAnalyticsDb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileAnalyticsIgnoreFiletype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectAntivirusProfileAnalyticsMaxUpload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileAnalyticsWlFiletype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
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
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandObjectAntivirusProfileCifsArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandObjectAntivirusProfileCifsArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandObjectAntivirusProfileCifsAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandObjectAntivirusProfileCifsEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandObjectAntivirusProfileCifsExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandObjectAntivirusProfileCifsFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandObjectAntivirusProfileCifsFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandObjectAntivirusProfileCifsFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandObjectAntivirusProfileCifsOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileCifsOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
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

func expandObjectAntivirusProfileCifsFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileCifsFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileCifsFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cover-page"], _ = expandObjectAntivirusProfileContentDisarmCoverPage(d, i["cover_page"], pre_append)
	}
	pre_append = pre + ".0." + "detect_only"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["detect-only"], _ = expandObjectAntivirusProfileContentDisarmDetectOnly(d, i["detect_only"], pre_append)
	}
	pre_append = pre + ".0." + "error_action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["error-action"], _ = expandObjectAntivirusProfileContentDisarmErrorAction(d, i["error_action"], pre_append)
	}
	pre_append = pre + ".0." + "office_action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["office-action"], _ = expandObjectAntivirusProfileContentDisarmOfficeAction(d, i["office_action"], pre_append)
	}
	pre_append = pre + ".0." + "office_dde"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["office-dde"], _ = expandObjectAntivirusProfileContentDisarmOfficeDde(d, i["office_dde"], pre_append)
	}
	pre_append = pre + ".0." + "office_embed"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["office-embed"], _ = expandObjectAntivirusProfileContentDisarmOfficeEmbed(d, i["office_embed"], pre_append)
	}
	pre_append = pre + ".0." + "office_hylink"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["office-hylink"], _ = expandObjectAntivirusProfileContentDisarmOfficeHylink(d, i["office_hylink"], pre_append)
	}
	pre_append = pre + ".0." + "office_linked"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["office-linked"], _ = expandObjectAntivirusProfileContentDisarmOfficeLinked(d, i["office_linked"], pre_append)
	}
	pre_append = pre + ".0." + "office_macro"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["office-macro"], _ = expandObjectAntivirusProfileContentDisarmOfficeMacro(d, i["office_macro"], pre_append)
	}
	pre_append = pre + ".0." + "original_file_destination"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["original-file-destination"], _ = expandObjectAntivirusProfileContentDisarmOriginalFileDestination(d, i["original_file_destination"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_form"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-act-form"], _ = expandObjectAntivirusProfileContentDisarmPdfActForm(d, i["pdf_act_form"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_gotor"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-act-gotor"], _ = expandObjectAntivirusProfileContentDisarmPdfActGotor(d, i["pdf_act_gotor"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_java"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-act-java"], _ = expandObjectAntivirusProfileContentDisarmPdfActJava(d, i["pdf_act_java"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_launch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-act-launch"], _ = expandObjectAntivirusProfileContentDisarmPdfActLaunch(d, i["pdf_act_launch"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_movie"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-act-movie"], _ = expandObjectAntivirusProfileContentDisarmPdfActMovie(d, i["pdf_act_movie"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_sound"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-act-sound"], _ = expandObjectAntivirusProfileContentDisarmPdfActSound(d, i["pdf_act_sound"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_embedfile"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-embedfile"], _ = expandObjectAntivirusProfileContentDisarmPdfEmbedfile(d, i["pdf_embedfile"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_hyperlink"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-hyperlink"], _ = expandObjectAntivirusProfileContentDisarmPdfHyperlink(d, i["pdf_hyperlink"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_javacode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
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
	return expandStringList(v.(*schema.Set).List()), nil
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

func expandObjectAntivirusProfileFortiaiErrorAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFortiaiTimeoutAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFortindrErrorAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFortindrTimeoutAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFortisandboxErrorAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFortisandboxMaxUpload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFortisandboxMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFortisandboxTimeoutAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandObjectAntivirusProfileFtpArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandObjectAntivirusProfileFtpArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandObjectAntivirusProfileFtpAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandObjectAntivirusProfileFtpEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandObjectAntivirusProfileFtpExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandObjectAntivirusProfileFtpFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandObjectAntivirusProfileFtpFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandObjectAntivirusProfileFtpFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandObjectAntivirusProfileFtpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileFtpOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
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

func expandObjectAntivirusProfileFtpFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFtpFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileFtpFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandObjectAntivirusProfileHttpArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandObjectAntivirusProfileHttpArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_optimize"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-optimize"], _ = expandObjectAntivirusProfileHttpAvOptimize(d, i["av_optimize"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandObjectAntivirusProfileHttpAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "content_disarm"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["content-disarm"], _ = expandObjectAntivirusProfileHttpContentDisarm(d, i["content_disarm"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandObjectAntivirusProfileHttpEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandObjectAntivirusProfileHttpExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandObjectAntivirusProfileHttpFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandObjectAntivirusProfileHttpFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandObjectAntivirusProfileHttpFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandObjectAntivirusProfileHttpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileHttpOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["quarantine"], _ = expandObjectAntivirusProfileHttpQuarantine(d, i["quarantine"], pre_append)
	}
	pre_append = pre + ".0." + "unknown_content_encoding"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unknown-content-encoding"], _ = expandObjectAntivirusProfileHttpUnknownContentEncoding(d, i["unknown_content_encoding"], pre_append)
	}

	return result, nil
}

func expandObjectAntivirusProfileHttpArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileHttpArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileHttpAvOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
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

func expandObjectAntivirusProfileHttpFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileHttpFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileHttpFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

func expandObjectAntivirusProfileHttpUnknownContentEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandObjectAntivirusProfileImapArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandObjectAntivirusProfileImapArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandObjectAntivirusProfileImapAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "content_disarm"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["content-disarm"], _ = expandObjectAntivirusProfileImapContentDisarm(d, i["content_disarm"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandObjectAntivirusProfileImapEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "executables"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["executables"], _ = expandObjectAntivirusProfileImapExecutables(d, i["executables"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandObjectAntivirusProfileImapExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandObjectAntivirusProfileImapFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandObjectAntivirusProfileImapFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandObjectAntivirusProfileImapFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandObjectAntivirusProfileImapOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileImapOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
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

func expandObjectAntivirusProfileImapFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileImapFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileImapFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

func expandObjectAntivirusProfileInspectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandObjectAntivirusProfileMapiArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandObjectAntivirusProfileMapiArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandObjectAntivirusProfileMapiAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandObjectAntivirusProfileMapiEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "executables"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["executables"], _ = expandObjectAntivirusProfileMapiExecutables(d, i["executables"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandObjectAntivirusProfileMapiExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandObjectAntivirusProfileMapiFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandObjectAntivirusProfileMapiFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandObjectAntivirusProfileMapiFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandObjectAntivirusProfileMapiOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileMapiOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
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

func expandObjectAntivirusProfileMapiFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileMapiFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileMapiFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["expiry"], _ = expandObjectAntivirusProfileNacQuarExpiry(d, i["expiry"], pre_append)
	}
	pre_append = pre + ".0." + "infected"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["infected"], _ = expandObjectAntivirusProfileNacQuarInfected(d, i["infected"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
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
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandObjectAntivirusProfileNntpArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandObjectAntivirusProfileNntpArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandObjectAntivirusProfileNntpAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandObjectAntivirusProfileNntpEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandObjectAntivirusProfileNntpExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandObjectAntivirusProfileNntpFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandObjectAntivirusProfileNntpFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandObjectAntivirusProfileNntpFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandObjectAntivirusProfileNntpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileNntpOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
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

func expandObjectAntivirusProfileNntpFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileNntpFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileNntpFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandObjectAntivirusProfileOutbreakPreventionExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "ftgd_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
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
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandObjectAntivirusProfilePop3ArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandObjectAntivirusProfilePop3ArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandObjectAntivirusProfilePop3AvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "content_disarm"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["content-disarm"], _ = expandObjectAntivirusProfilePop3ContentDisarm(d, i["content_disarm"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandObjectAntivirusProfilePop3Emulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "executables"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["executables"], _ = expandObjectAntivirusProfilePop3Executables(d, i["executables"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandObjectAntivirusProfilePop3ExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandObjectAntivirusProfilePop3Fortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandObjectAntivirusProfilePop3Fortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandObjectAntivirusProfilePop3Fortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandObjectAntivirusProfilePop3Options(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfilePop3OutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
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

func expandObjectAntivirusProfilePop3Fortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfilePop3Fortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfilePop3Fortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
	return convstr2list(v, nil), nil
}

func expandObjectAntivirusProfileScanMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSmb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandObjectAntivirusProfileSmbArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandObjectAntivirusProfileSmbArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandObjectAntivirusProfileSmbEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandObjectAntivirusProfileSmbOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileSmbOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}

	return result, nil
}

func expandObjectAntivirusProfileSmbArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileSmbArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileSmbEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSmbOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAntivirusProfileSmbOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandObjectAntivirusProfileSmtpArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandObjectAntivirusProfileSmtpArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandObjectAntivirusProfileSmtpAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "content_disarm"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["content-disarm"], _ = expandObjectAntivirusProfileSmtpContentDisarm(d, i["content_disarm"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandObjectAntivirusProfileSmtpEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "executables"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["executables"], _ = expandObjectAntivirusProfileSmtpExecutables(d, i["executables"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandObjectAntivirusProfileSmtpExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandObjectAntivirusProfileSmtpFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandObjectAntivirusProfileSmtpFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandObjectAntivirusProfileSmtpFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandObjectAntivirusProfileSmtpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileSmtpOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
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

func expandObjectAntivirusProfileSmtpFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSmtpFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSmtpFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandObjectAntivirusProfileSshArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandObjectAntivirusProfileSshArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandObjectAntivirusProfileSshAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandObjectAntivirusProfileSshEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandObjectAntivirusProfileSshExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandObjectAntivirusProfileSshFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandObjectAntivirusProfileSshFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandObjectAntivirusProfileSshFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandObjectAntivirusProfileSshOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandObjectAntivirusProfileSshOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
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

func expandObjectAntivirusProfileSshFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSshFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileSshFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

	if v, ok := d.GetOk("analytics_accept_filetype"); ok || d.HasChange("analytics_accept_filetype") {
		t, err := expandObjectAntivirusProfileAnalyticsAcceptFiletype(d, v, "analytics_accept_filetype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-accept-filetype"] = t
		}
	}

	if v, ok := d.GetOk("analytics_bl_filetype"); ok || d.HasChange("analytics_bl_filetype") {
		t, err := expandObjectAntivirusProfileAnalyticsBlFiletype(d, v, "analytics_bl_filetype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-bl-filetype"] = t
		}
	}

	if v, ok := d.GetOk("analytics_db"); ok || d.HasChange("analytics_db") {
		t, err := expandObjectAntivirusProfileAnalyticsDb(d, v, "analytics_db")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-db"] = t
		}
	}

	if v, ok := d.GetOk("analytics_ignore_filetype"); ok || d.HasChange("analytics_ignore_filetype") {
		t, err := expandObjectAntivirusProfileAnalyticsIgnoreFiletype(d, v, "analytics_ignore_filetype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-ignore-filetype"] = t
		}
	}

	if v, ok := d.GetOk("analytics_max_upload"); ok || d.HasChange("analytics_max_upload") {
		t, err := expandObjectAntivirusProfileAnalyticsMaxUpload(d, v, "analytics_max_upload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-max-upload"] = t
		}
	}

	if v, ok := d.GetOk("analytics_wl_filetype"); ok || d.HasChange("analytics_wl_filetype") {
		t, err := expandObjectAntivirusProfileAnalyticsWlFiletype(d, v, "analytics_wl_filetype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-wl-filetype"] = t
		}
	}

	if v, ok := d.GetOk("av_block_log"); ok || d.HasChange("av_block_log") {
		t, err := expandObjectAntivirusProfileAvBlockLog(d, v, "av_block_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-block-log"] = t
		}
	}

	if v, ok := d.GetOk("av_virus_log"); ok || d.HasChange("av_virus_log") {
		t, err := expandObjectAntivirusProfileAvVirusLog(d, v, "av_virus_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-virus-log"] = t
		}
	}

	if v, ok := d.GetOk("cifs"); ok || d.HasChange("cifs") {
		t, err := expandObjectAntivirusProfileCifs(d, v, "cifs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectAntivirusProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("content_disarm"); ok || d.HasChange("content_disarm") {
		t, err := expandObjectAntivirusProfileContentDisarm(d, v, "content_disarm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-disarm"] = t
		}
	}

	if v, ok := d.GetOk("ems_threat_feed"); ok || d.HasChange("ems_threat_feed") {
		t, err := expandObjectAntivirusProfileEmsThreatFeed(d, v, "ems_threat_feed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-threat-feed"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandObjectAntivirusProfileExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("external_blocklist"); ok || d.HasChange("external_blocklist") {
		t, err := expandObjectAntivirusProfileExternalBlocklist(d, v, "external_blocklist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-blocklist"] = t
		}
	}

	if v, ok := d.GetOk("external_blocklist_archive_scan"); ok || d.HasChange("external_blocklist_archive_scan") {
		t, err := expandObjectAntivirusProfileExternalBlocklistArchiveScan(d, v, "external_blocklist_archive_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-blocklist-archive-scan"] = t
		}
	}

	if v, ok := d.GetOk("external_blocklist_enable_all"); ok || d.HasChange("external_blocklist_enable_all") {
		t, err := expandObjectAntivirusProfileExternalBlocklistEnableAll(d, v, "external_blocklist_enable_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-blocklist-enable-all"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok || d.HasChange("feature_set") {
		t, err := expandObjectAntivirusProfileFeatureSet(d, v, "feature_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("fortiai_error_action"); ok || d.HasChange("fortiai_error_action") {
		t, err := expandObjectAntivirusProfileFortiaiErrorAction(d, v, "fortiai_error_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiai-error-action"] = t
		}
	}

	if v, ok := d.GetOk("fortiai_timeout_action"); ok || d.HasChange("fortiai_timeout_action") {
		t, err := expandObjectAntivirusProfileFortiaiTimeoutAction(d, v, "fortiai_timeout_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiai-timeout-action"] = t
		}
	}

	if v, ok := d.GetOk("fortindr_error_action"); ok || d.HasChange("fortindr_error_action") {
		t, err := expandObjectAntivirusProfileFortindrErrorAction(d, v, "fortindr_error_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortindr-error-action"] = t
		}
	}

	if v, ok := d.GetOk("fortindr_timeout_action"); ok || d.HasChange("fortindr_timeout_action") {
		t, err := expandObjectAntivirusProfileFortindrTimeoutAction(d, v, "fortindr_timeout_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortindr-timeout-action"] = t
		}
	}

	if v, ok := d.GetOk("fortisandbox_error_action"); ok || d.HasChange("fortisandbox_error_action") {
		t, err := expandObjectAntivirusProfileFortisandboxErrorAction(d, v, "fortisandbox_error_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortisandbox-error-action"] = t
		}
	}

	if v, ok := d.GetOk("fortisandbox_max_upload"); ok || d.HasChange("fortisandbox_max_upload") {
		t, err := expandObjectAntivirusProfileFortisandboxMaxUpload(d, v, "fortisandbox_max_upload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortisandbox-max-upload"] = t
		}
	}

	if v, ok := d.GetOk("fortisandbox_mode"); ok || d.HasChange("fortisandbox_mode") {
		t, err := expandObjectAntivirusProfileFortisandboxMode(d, v, "fortisandbox_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortisandbox-mode"] = t
		}
	}

	if v, ok := d.GetOk("fortisandbox_timeout_action"); ok || d.HasChange("fortisandbox_timeout_action") {
		t, err := expandObjectAntivirusProfileFortisandboxTimeoutAction(d, v, "fortisandbox_timeout_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortisandbox-timeout-action"] = t
		}
	}

	if v, ok := d.GetOk("ftgd_analytics"); ok || d.HasChange("ftgd_analytics") {
		t, err := expandObjectAntivirusProfileFtgdAnalytics(d, v, "ftgd_analytics")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftgd-analytics"] = t
		}
	}

	if v, ok := d.GetOk("ftp"); ok || d.HasChange("ftp") {
		t, err := expandObjectAntivirusProfileFtp(d, v, "ftp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp"] = t
		}
	}

	if v, ok := d.GetOk("http"); ok || d.HasChange("http") {
		t, err := expandObjectAntivirusProfileHttp(d, v, "http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http"] = t
		}
	}

	if v, ok := d.GetOk("imap"); ok || d.HasChange("imap") {
		t, err := expandObjectAntivirusProfileImap(d, v, "imap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imap"] = t
		}
	}

	if v, ok := d.GetOk("inspection_mode"); ok || d.HasChange("inspection_mode") {
		t, err := expandObjectAntivirusProfileInspectionMode(d, v, "inspection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspection-mode"] = t
		}
	}

	if v, ok := d.GetOk("mapi"); ok || d.HasChange("mapi") {
		t, err := expandObjectAntivirusProfileMapi(d, v, "mapi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi"] = t
		}
	}

	if v, ok := d.GetOk("mobile_malware_db"); ok || d.HasChange("mobile_malware_db") {
		t, err := expandObjectAntivirusProfileMobileMalwareDb(d, v, "mobile_malware_db")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mobile-malware-db"] = t
		}
	}

	if v, ok := d.GetOk("nac_quar"); ok || d.HasChange("nac_quar") {
		t, err := expandObjectAntivirusProfileNacQuar(d, v, "nac_quar")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-quar"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectAntivirusProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nntp"); ok || d.HasChange("nntp") {
		t, err := expandObjectAntivirusProfileNntp(d, v, "nntp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nntp"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_archive_scan"); ok || d.HasChange("outbreak_prevention_archive_scan") {
		t, err := expandObjectAntivirusProfileOutbreakPreventionArchiveScan(d, v, "outbreak_prevention_archive_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-archive-scan"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention"); ok || d.HasChange("outbreak_prevention") {
		t, err := expandObjectAntivirusProfileOutbreakPrevention(d, v, "outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("pop3"); ok || d.HasChange("pop3") {
		t, err := expandObjectAntivirusProfilePop3(d, v, "pop3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandObjectAntivirusProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("scan_mode"); ok || d.HasChange("scan_mode") {
		t, err := expandObjectAntivirusProfileScanMode(d, v, "scan_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-mode"] = t
		}
	}

	if v, ok := d.GetOk("smb"); ok || d.HasChange("smb") {
		t, err := expandObjectAntivirusProfileSmb(d, v, "smb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smb"] = t
		}
	}

	if v, ok := d.GetOk("smtp"); ok || d.HasChange("smtp") {
		t, err := expandObjectAntivirusProfileSmtp(d, v, "smtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp"] = t
		}
	}

	if v, ok := d.GetOk("ssh"); ok || d.HasChange("ssh") {
		t, err := expandObjectAntivirusProfileSsh(d, v, "ssh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh"] = t
		}
	}

	return &obj, nil
}
