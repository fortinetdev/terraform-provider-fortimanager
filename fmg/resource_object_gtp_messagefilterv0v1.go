// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Message filter for GTPv0/v1 messages.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectGtpMessageFilterV0V1() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectGtpMessageFilterV0V1Create,
		Read:   resourceObjectGtpMessageFilterV0V1Read,
		Update: resourceObjectGtpMessageFilterV0V1Update,
		Delete: resourceObjectGtpMessageFilterV0V1Delete,

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
			"create_mbms": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"create_pdp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data_record": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delete_aa_pdp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delete_mbms": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delete_pdp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"echo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_marker": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error_indication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failure_report": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwd_relocation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwd_srns_context": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtp_pdu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"identification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mbms_de_registration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mbms_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mbms_registration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mbms_session_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mbms_session_stop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mbms_session_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ms_info_change_notif": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"node_alive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"note_ms_present": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pdu_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ran_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"relocation_cancel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sgsn_context": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"support_extension": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_message_white_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"update_mbms": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_pdp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"v0_create_aa_pdp__v1_init_pdp_ctx": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version_not_support": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectGtpMessageFilterV0V1Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectGtpMessageFilterV0V1(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectGtpMessageFilterV0V1 resource while getting object: %v", err)
	}

	_, err = c.CreateObjectGtpMessageFilterV0V1(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectGtpMessageFilterV0V1 resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectGtpMessageFilterV0V1Read(d, m)
}

func resourceObjectGtpMessageFilterV0V1Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectGtpMessageFilterV0V1(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGtpMessageFilterV0V1 resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectGtpMessageFilterV0V1(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGtpMessageFilterV0V1 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectGtpMessageFilterV0V1Read(d, m)
}

func resourceObjectGtpMessageFilterV0V1Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectGtpMessageFilterV0V1(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectGtpMessageFilterV0V1 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectGtpMessageFilterV0V1Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectGtpMessageFilterV0V1(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGtpMessageFilterV0V1 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectGtpMessageFilterV0V1(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGtpMessageFilterV0V1 resource from API: %v", err)
	}
	return nil
}

func flattenObjectGtpMessageFilterV0V1CreateMbms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1CreatePdp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1DataRecord(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1DeleteAaPdp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1DeleteMbms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1DeletePdp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1Echo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1EndMarker(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1ErrorIndication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1FailureReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1FwdRelocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1FwdSrnsContext(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1GtpPdu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1Identification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1MbmsDeRegistration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1MbmsNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1MbmsRegistration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1MbmsSessionStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1MbmsSessionStop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1MbmsSessionUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1MsInfoChangeNotif(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpMessageFilterV0V1NodeAlive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1NoteMsPresent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1PduNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1RanInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1Redirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1RelocationCancel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1SendRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1SgsnContext(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1SupportExtension(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1UnknownMessage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1UnknownMessageWhiteList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectGtpMessageFilterV0V1UpdateMbms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1UpdatePdp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1V0CreateAaPdpV1InitPdpCtx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "deny",
			1: "allow",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectGtpMessageFilterV0V1VersionNotSupport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "allow",
			1: "deny",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectObjectGtpMessageFilterV0V1(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("create_mbms", flattenObjectGtpMessageFilterV0V1CreateMbms(o["create-mbms"], d, "create_mbms")); err != nil {
		if vv, ok := fortiAPIPatch(o["create-mbms"], "ObjectGtpMessageFilterV0V1-CreateMbms"); ok {
			if err = d.Set("create_mbms", vv); err != nil {
				return fmt.Errorf("Error reading create_mbms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading create_mbms: %v", err)
		}
	}

	if err = d.Set("create_pdp", flattenObjectGtpMessageFilterV0V1CreatePdp(o["create-pdp"], d, "create_pdp")); err != nil {
		if vv, ok := fortiAPIPatch(o["create-pdp"], "ObjectGtpMessageFilterV0V1-CreatePdp"); ok {
			if err = d.Set("create_pdp", vv); err != nil {
				return fmt.Errorf("Error reading create_pdp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading create_pdp: %v", err)
		}
	}

	if err = d.Set("data_record", flattenObjectGtpMessageFilterV0V1DataRecord(o["data-record"], d, "data_record")); err != nil {
		if vv, ok := fortiAPIPatch(o["data-record"], "ObjectGtpMessageFilterV0V1-DataRecord"); ok {
			if err = d.Set("data_record", vv); err != nil {
				return fmt.Errorf("Error reading data_record: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading data_record: %v", err)
		}
	}

	if err = d.Set("delete_aa_pdp", flattenObjectGtpMessageFilterV0V1DeleteAaPdp(o["delete-aa-pdp"], d, "delete_aa_pdp")); err != nil {
		if vv, ok := fortiAPIPatch(o["delete-aa-pdp"], "ObjectGtpMessageFilterV0V1-DeleteAaPdp"); ok {
			if err = d.Set("delete_aa_pdp", vv); err != nil {
				return fmt.Errorf("Error reading delete_aa_pdp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delete_aa_pdp: %v", err)
		}
	}

	if err = d.Set("delete_mbms", flattenObjectGtpMessageFilterV0V1DeleteMbms(o["delete-mbms"], d, "delete_mbms")); err != nil {
		if vv, ok := fortiAPIPatch(o["delete-mbms"], "ObjectGtpMessageFilterV0V1-DeleteMbms"); ok {
			if err = d.Set("delete_mbms", vv); err != nil {
				return fmt.Errorf("Error reading delete_mbms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delete_mbms: %v", err)
		}
	}

	if err = d.Set("delete_pdp", flattenObjectGtpMessageFilterV0V1DeletePdp(o["delete-pdp"], d, "delete_pdp")); err != nil {
		if vv, ok := fortiAPIPatch(o["delete-pdp"], "ObjectGtpMessageFilterV0V1-DeletePdp"); ok {
			if err = d.Set("delete_pdp", vv); err != nil {
				return fmt.Errorf("Error reading delete_pdp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delete_pdp: %v", err)
		}
	}

	if err = d.Set("echo", flattenObjectGtpMessageFilterV0V1Echo(o["echo"], d, "echo")); err != nil {
		if vv, ok := fortiAPIPatch(o["echo"], "ObjectGtpMessageFilterV0V1-Echo"); ok {
			if err = d.Set("echo", vv); err != nil {
				return fmt.Errorf("Error reading echo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading echo: %v", err)
		}
	}

	if err = d.Set("end_marker", flattenObjectGtpMessageFilterV0V1EndMarker(o["end-marker"], d, "end_marker")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-marker"], "ObjectGtpMessageFilterV0V1-EndMarker"); ok {
			if err = d.Set("end_marker", vv); err != nil {
				return fmt.Errorf("Error reading end_marker: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_marker: %v", err)
		}
	}

	if err = d.Set("error_indication", flattenObjectGtpMessageFilterV0V1ErrorIndication(o["error-indication"], d, "error_indication")); err != nil {
		if vv, ok := fortiAPIPatch(o["error-indication"], "ObjectGtpMessageFilterV0V1-ErrorIndication"); ok {
			if err = d.Set("error_indication", vv); err != nil {
				return fmt.Errorf("Error reading error_indication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading error_indication: %v", err)
		}
	}

	if err = d.Set("failure_report", flattenObjectGtpMessageFilterV0V1FailureReport(o["failure-report"], d, "failure_report")); err != nil {
		if vv, ok := fortiAPIPatch(o["failure-report"], "ObjectGtpMessageFilterV0V1-FailureReport"); ok {
			if err = d.Set("failure_report", vv); err != nil {
				return fmt.Errorf("Error reading failure_report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failure_report: %v", err)
		}
	}

	if err = d.Set("fwd_relocation", flattenObjectGtpMessageFilterV0V1FwdRelocation(o["fwd-relocation"], d, "fwd_relocation")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-relocation"], "ObjectGtpMessageFilterV0V1-FwdRelocation"); ok {
			if err = d.Set("fwd_relocation", vv); err != nil {
				return fmt.Errorf("Error reading fwd_relocation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_relocation: %v", err)
		}
	}

	if err = d.Set("fwd_srns_context", flattenObjectGtpMessageFilterV0V1FwdSrnsContext(o["fwd-srns-context"], d, "fwd_srns_context")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-srns-context"], "ObjectGtpMessageFilterV0V1-FwdSrnsContext"); ok {
			if err = d.Set("fwd_srns_context", vv); err != nil {
				return fmt.Errorf("Error reading fwd_srns_context: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_srns_context: %v", err)
		}
	}

	if err = d.Set("gtp_pdu", flattenObjectGtpMessageFilterV0V1GtpPdu(o["gtp-pdu"], d, "gtp_pdu")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp-pdu"], "ObjectGtpMessageFilterV0V1-GtpPdu"); ok {
			if err = d.Set("gtp_pdu", vv); err != nil {
				return fmt.Errorf("Error reading gtp_pdu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp_pdu: %v", err)
		}
	}

	if err = d.Set("identification", flattenObjectGtpMessageFilterV0V1Identification(o["identification"], d, "identification")); err != nil {
		if vv, ok := fortiAPIPatch(o["identification"], "ObjectGtpMessageFilterV0V1-Identification"); ok {
			if err = d.Set("identification", vv); err != nil {
				return fmt.Errorf("Error reading identification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading identification: %v", err)
		}
	}

	if err = d.Set("mbms_de_registration", flattenObjectGtpMessageFilterV0V1MbmsDeRegistration(o["mbms-de-registration"], d, "mbms_de_registration")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-de-registration"], "ObjectGtpMessageFilterV0V1-MbmsDeRegistration"); ok {
			if err = d.Set("mbms_de_registration", vv); err != nil {
				return fmt.Errorf("Error reading mbms_de_registration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_de_registration: %v", err)
		}
	}

	if err = d.Set("mbms_notification", flattenObjectGtpMessageFilterV0V1MbmsNotification(o["mbms-notification"], d, "mbms_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-notification"], "ObjectGtpMessageFilterV0V1-MbmsNotification"); ok {
			if err = d.Set("mbms_notification", vv); err != nil {
				return fmt.Errorf("Error reading mbms_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_notification: %v", err)
		}
	}

	if err = d.Set("mbms_registration", flattenObjectGtpMessageFilterV0V1MbmsRegistration(o["mbms-registration"], d, "mbms_registration")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-registration"], "ObjectGtpMessageFilterV0V1-MbmsRegistration"); ok {
			if err = d.Set("mbms_registration", vv); err != nil {
				return fmt.Errorf("Error reading mbms_registration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_registration: %v", err)
		}
	}

	if err = d.Set("mbms_session_start", flattenObjectGtpMessageFilterV0V1MbmsSessionStart(o["mbms-session-start"], d, "mbms_session_start")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-session-start"], "ObjectGtpMessageFilterV0V1-MbmsSessionStart"); ok {
			if err = d.Set("mbms_session_start", vv); err != nil {
				return fmt.Errorf("Error reading mbms_session_start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_session_start: %v", err)
		}
	}

	if err = d.Set("mbms_session_stop", flattenObjectGtpMessageFilterV0V1MbmsSessionStop(o["mbms-session-stop"], d, "mbms_session_stop")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-session-stop"], "ObjectGtpMessageFilterV0V1-MbmsSessionStop"); ok {
			if err = d.Set("mbms_session_stop", vv); err != nil {
				return fmt.Errorf("Error reading mbms_session_stop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_session_stop: %v", err)
		}
	}

	if err = d.Set("mbms_session_update", flattenObjectGtpMessageFilterV0V1MbmsSessionUpdate(o["mbms-session-update"], d, "mbms_session_update")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbms-session-update"], "ObjectGtpMessageFilterV0V1-MbmsSessionUpdate"); ok {
			if err = d.Set("mbms_session_update", vv); err != nil {
				return fmt.Errorf("Error reading mbms_session_update: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbms_session_update: %v", err)
		}
	}

	if err = d.Set("ms_info_change_notif", flattenObjectGtpMessageFilterV0V1MsInfoChangeNotif(o["ms-info-change-notif"], d, "ms_info_change_notif")); err != nil {
		if vv, ok := fortiAPIPatch(o["ms-info-change-notif"], "ObjectGtpMessageFilterV0V1-MsInfoChangeNotif"); ok {
			if err = d.Set("ms_info_change_notif", vv); err != nil {
				return fmt.Errorf("Error reading ms_info_change_notif: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ms_info_change_notif: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectGtpMessageFilterV0V1Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectGtpMessageFilterV0V1-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("node_alive", flattenObjectGtpMessageFilterV0V1NodeAlive(o["node-alive"], d, "node_alive")); err != nil {
		if vv, ok := fortiAPIPatch(o["node-alive"], "ObjectGtpMessageFilterV0V1-NodeAlive"); ok {
			if err = d.Set("node_alive", vv); err != nil {
				return fmt.Errorf("Error reading node_alive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading node_alive: %v", err)
		}
	}

	if err = d.Set("note_ms_present", flattenObjectGtpMessageFilterV0V1NoteMsPresent(o["note-ms-present"], d, "note_ms_present")); err != nil {
		if vv, ok := fortiAPIPatch(o["note-ms-present"], "ObjectGtpMessageFilterV0V1-NoteMsPresent"); ok {
			if err = d.Set("note_ms_present", vv); err != nil {
				return fmt.Errorf("Error reading note_ms_present: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading note_ms_present: %v", err)
		}
	}

	if err = d.Set("pdu_notification", flattenObjectGtpMessageFilterV0V1PduNotification(o["pdu-notification"], d, "pdu_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdu-notification"], "ObjectGtpMessageFilterV0V1-PduNotification"); ok {
			if err = d.Set("pdu_notification", vv); err != nil {
				return fmt.Errorf("Error reading pdu_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdu_notification: %v", err)
		}
	}

	if err = d.Set("ran_info", flattenObjectGtpMessageFilterV0V1RanInfo(o["ran-info"], d, "ran_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["ran-info"], "ObjectGtpMessageFilterV0V1-RanInfo"); ok {
			if err = d.Set("ran_info", vv); err != nil {
				return fmt.Errorf("Error reading ran_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ran_info: %v", err)
		}
	}

	if err = d.Set("redirection", flattenObjectGtpMessageFilterV0V1Redirection(o["redirection"], d, "redirection")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirection"], "ObjectGtpMessageFilterV0V1-Redirection"); ok {
			if err = d.Set("redirection", vv); err != nil {
				return fmt.Errorf("Error reading redirection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirection: %v", err)
		}
	}

	if err = d.Set("relocation_cancel", flattenObjectGtpMessageFilterV0V1RelocationCancel(o["relocation-cancel"], d, "relocation_cancel")); err != nil {
		if vv, ok := fortiAPIPatch(o["relocation-cancel"], "ObjectGtpMessageFilterV0V1-RelocationCancel"); ok {
			if err = d.Set("relocation_cancel", vv); err != nil {
				return fmt.Errorf("Error reading relocation_cancel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading relocation_cancel: %v", err)
		}
	}

	if err = d.Set("send_route", flattenObjectGtpMessageFilterV0V1SendRoute(o["send-route"], d, "send_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-route"], "ObjectGtpMessageFilterV0V1-SendRoute"); ok {
			if err = d.Set("send_route", vv); err != nil {
				return fmt.Errorf("Error reading send_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_route: %v", err)
		}
	}

	if err = d.Set("sgsn_context", flattenObjectGtpMessageFilterV0V1SgsnContext(o["sgsn-context"], d, "sgsn_context")); err != nil {
		if vv, ok := fortiAPIPatch(o["sgsn-context"], "ObjectGtpMessageFilterV0V1-SgsnContext"); ok {
			if err = d.Set("sgsn_context", vv); err != nil {
				return fmt.Errorf("Error reading sgsn_context: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sgsn_context: %v", err)
		}
	}

	if err = d.Set("support_extension", flattenObjectGtpMessageFilterV0V1SupportExtension(o["support-extension"], d, "support_extension")); err != nil {
		if vv, ok := fortiAPIPatch(o["support-extension"], "ObjectGtpMessageFilterV0V1-SupportExtension"); ok {
			if err = d.Set("support_extension", vv); err != nil {
				return fmt.Errorf("Error reading support_extension: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading support_extension: %v", err)
		}
	}

	if err = d.Set("unknown_message", flattenObjectGtpMessageFilterV0V1UnknownMessage(o["unknown-message"], d, "unknown_message")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-message"], "ObjectGtpMessageFilterV0V1-UnknownMessage"); ok {
			if err = d.Set("unknown_message", vv); err != nil {
				return fmt.Errorf("Error reading unknown_message: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_message: %v", err)
		}
	}

	if err = d.Set("unknown_message_white_list", flattenObjectGtpMessageFilterV0V1UnknownMessageWhiteList(o["unknown-message-white-list"], d, "unknown_message_white_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-message-white-list"], "ObjectGtpMessageFilterV0V1-UnknownMessageWhiteList"); ok {
			if err = d.Set("unknown_message_white_list", vv); err != nil {
				return fmt.Errorf("Error reading unknown_message_white_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_message_white_list: %v", err)
		}
	}

	if err = d.Set("update_mbms", flattenObjectGtpMessageFilterV0V1UpdateMbms(o["update-mbms"], d, "update_mbms")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-mbms"], "ObjectGtpMessageFilterV0V1-UpdateMbms"); ok {
			if err = d.Set("update_mbms", vv); err != nil {
				return fmt.Errorf("Error reading update_mbms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_mbms: %v", err)
		}
	}

	if err = d.Set("update_pdp", flattenObjectGtpMessageFilterV0V1UpdatePdp(o["update-pdp"], d, "update_pdp")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-pdp"], "ObjectGtpMessageFilterV0V1-UpdatePdp"); ok {
			if err = d.Set("update_pdp", vv); err != nil {
				return fmt.Errorf("Error reading update_pdp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_pdp: %v", err)
		}
	}

	if err = d.Set("v0_create_aa_pdp__v1_init_pdp_ctx", flattenObjectGtpMessageFilterV0V1V0CreateAaPdpV1InitPdpCtx(o["v0-create-aa-pdp--v1-init-pdp-ctx"], d, "v0_create_aa_pdp__v1_init_pdp_ctx")); err != nil {
		if vv, ok := fortiAPIPatch(o["v0-create-aa-pdp--v1-init-pdp-ctx"], "ObjectGtpMessageFilterV0V1-V0CreateAaPdpV1InitPdpCtx"); ok {
			if err = d.Set("v0_create_aa_pdp__v1_init_pdp_ctx", vv); err != nil {
				return fmt.Errorf("Error reading v0_create_aa_pdp__v1_init_pdp_ctx: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading v0_create_aa_pdp__v1_init_pdp_ctx: %v", err)
		}
	}

	if err = d.Set("version_not_support", flattenObjectGtpMessageFilterV0V1VersionNotSupport(o["version-not-support"], d, "version_not_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["version-not-support"], "ObjectGtpMessageFilterV0V1-VersionNotSupport"); ok {
			if err = d.Set("version_not_support", vv); err != nil {
				return fmt.Errorf("Error reading version_not_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version_not_support: %v", err)
		}
	}

	return nil
}

func flattenObjectGtpMessageFilterV0V1FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectGtpMessageFilterV0V1CreateMbms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1CreatePdp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1DataRecord(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1DeleteAaPdp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1DeleteMbms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1DeletePdp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1Echo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1EndMarker(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1ErrorIndication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1FailureReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1FwdRelocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1FwdSrnsContext(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1GtpPdu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1Identification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1MbmsDeRegistration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1MbmsNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1MbmsRegistration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1MbmsSessionStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1MbmsSessionStop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1MbmsSessionUpdate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1MsInfoChangeNotif(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1NodeAlive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1NoteMsPresent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1PduNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1RanInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1Redirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1RelocationCancel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1SendRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1SgsnContext(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1SupportExtension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1UnknownMessage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1UnknownMessageWhiteList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectGtpMessageFilterV0V1UpdateMbms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1UpdatePdp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1V0CreateAaPdpV1InitPdpCtx(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV0V1VersionNotSupport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectGtpMessageFilterV0V1(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("create_mbms"); ok {
		t, err := expandObjectGtpMessageFilterV0V1CreateMbms(d, v, "create_mbms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create-mbms"] = t
		}
	}

	if v, ok := d.GetOk("create_pdp"); ok {
		t, err := expandObjectGtpMessageFilterV0V1CreatePdp(d, v, "create_pdp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create-pdp"] = t
		}
	}

	if v, ok := d.GetOk("data_record"); ok {
		t, err := expandObjectGtpMessageFilterV0V1DataRecord(d, v, "data_record")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data-record"] = t
		}
	}

	if v, ok := d.GetOk("delete_aa_pdp"); ok {
		t, err := expandObjectGtpMessageFilterV0V1DeleteAaPdp(d, v, "delete_aa_pdp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delete-aa-pdp"] = t
		}
	}

	if v, ok := d.GetOk("delete_mbms"); ok {
		t, err := expandObjectGtpMessageFilterV0V1DeleteMbms(d, v, "delete_mbms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delete-mbms"] = t
		}
	}

	if v, ok := d.GetOk("delete_pdp"); ok {
		t, err := expandObjectGtpMessageFilterV0V1DeletePdp(d, v, "delete_pdp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delete-pdp"] = t
		}
	}

	if v, ok := d.GetOk("echo"); ok {
		t, err := expandObjectGtpMessageFilterV0V1Echo(d, v, "echo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["echo"] = t
		}
	}

	if v, ok := d.GetOk("end_marker"); ok {
		t, err := expandObjectGtpMessageFilterV0V1EndMarker(d, v, "end_marker")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-marker"] = t
		}
	}

	if v, ok := d.GetOk("error_indication"); ok {
		t, err := expandObjectGtpMessageFilterV0V1ErrorIndication(d, v, "error_indication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["error-indication"] = t
		}
	}

	if v, ok := d.GetOk("failure_report"); ok {
		t, err := expandObjectGtpMessageFilterV0V1FailureReport(d, v, "failure_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failure-report"] = t
		}
	}

	if v, ok := d.GetOk("fwd_relocation"); ok {
		t, err := expandObjectGtpMessageFilterV0V1FwdRelocation(d, v, "fwd_relocation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-relocation"] = t
		}
	}

	if v, ok := d.GetOk("fwd_srns_context"); ok {
		t, err := expandObjectGtpMessageFilterV0V1FwdSrnsContext(d, v, "fwd_srns_context")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-srns-context"] = t
		}
	}

	if v, ok := d.GetOk("gtp_pdu"); ok {
		t, err := expandObjectGtpMessageFilterV0V1GtpPdu(d, v, "gtp_pdu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp-pdu"] = t
		}
	}

	if v, ok := d.GetOk("identification"); ok {
		t, err := expandObjectGtpMessageFilterV0V1Identification(d, v, "identification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["identification"] = t
		}
	}

	if v, ok := d.GetOk("mbms_de_registration"); ok {
		t, err := expandObjectGtpMessageFilterV0V1MbmsDeRegistration(d, v, "mbms_de_registration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-de-registration"] = t
		}
	}

	if v, ok := d.GetOk("mbms_notification"); ok {
		t, err := expandObjectGtpMessageFilterV0V1MbmsNotification(d, v, "mbms_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-notification"] = t
		}
	}

	if v, ok := d.GetOk("mbms_registration"); ok {
		t, err := expandObjectGtpMessageFilterV0V1MbmsRegistration(d, v, "mbms_registration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-registration"] = t
		}
	}

	if v, ok := d.GetOk("mbms_session_start"); ok {
		t, err := expandObjectGtpMessageFilterV0V1MbmsSessionStart(d, v, "mbms_session_start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-session-start"] = t
		}
	}

	if v, ok := d.GetOk("mbms_session_stop"); ok {
		t, err := expandObjectGtpMessageFilterV0V1MbmsSessionStop(d, v, "mbms_session_stop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-session-stop"] = t
		}
	}

	if v, ok := d.GetOk("mbms_session_update"); ok {
		t, err := expandObjectGtpMessageFilterV0V1MbmsSessionUpdate(d, v, "mbms_session_update")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbms-session-update"] = t
		}
	}

	if v, ok := d.GetOk("ms_info_change_notif"); ok {
		t, err := expandObjectGtpMessageFilterV0V1MsInfoChangeNotif(d, v, "ms_info_change_notif")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ms-info-change-notif"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectGtpMessageFilterV0V1Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("node_alive"); ok {
		t, err := expandObjectGtpMessageFilterV0V1NodeAlive(d, v, "node_alive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["node-alive"] = t
		}
	}

	if v, ok := d.GetOk("note_ms_present"); ok {
		t, err := expandObjectGtpMessageFilterV0V1NoteMsPresent(d, v, "note_ms_present")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["note-ms-present"] = t
		}
	}

	if v, ok := d.GetOk("pdu_notification"); ok {
		t, err := expandObjectGtpMessageFilterV0V1PduNotification(d, v, "pdu_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdu-notification"] = t
		}
	}

	if v, ok := d.GetOk("ran_info"); ok {
		t, err := expandObjectGtpMessageFilterV0V1RanInfo(d, v, "ran_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ran-info"] = t
		}
	}

	if v, ok := d.GetOk("redirection"); ok {
		t, err := expandObjectGtpMessageFilterV0V1Redirection(d, v, "redirection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirection"] = t
		}
	}

	if v, ok := d.GetOk("relocation_cancel"); ok {
		t, err := expandObjectGtpMessageFilterV0V1RelocationCancel(d, v, "relocation_cancel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["relocation-cancel"] = t
		}
	}

	if v, ok := d.GetOk("send_route"); ok {
		t, err := expandObjectGtpMessageFilterV0V1SendRoute(d, v, "send_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-route"] = t
		}
	}

	if v, ok := d.GetOk("sgsn_context"); ok {
		t, err := expandObjectGtpMessageFilterV0V1SgsnContext(d, v, "sgsn_context")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sgsn-context"] = t
		}
	}

	if v, ok := d.GetOk("support_extension"); ok {
		t, err := expandObjectGtpMessageFilterV0V1SupportExtension(d, v, "support_extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["support-extension"] = t
		}
	}

	if v, ok := d.GetOk("unknown_message"); ok {
		t, err := expandObjectGtpMessageFilterV0V1UnknownMessage(d, v, "unknown_message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-message"] = t
		}
	}

	if v, ok := d.GetOk("unknown_message_white_list"); ok {
		t, err := expandObjectGtpMessageFilterV0V1UnknownMessageWhiteList(d, v, "unknown_message_white_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-message-white-list"] = t
		}
	}

	if v, ok := d.GetOk("update_mbms"); ok {
		t, err := expandObjectGtpMessageFilterV0V1UpdateMbms(d, v, "update_mbms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-mbms"] = t
		}
	}

	if v, ok := d.GetOk("update_pdp"); ok {
		t, err := expandObjectGtpMessageFilterV0V1UpdatePdp(d, v, "update_pdp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-pdp"] = t
		}
	}

	if v, ok := d.GetOk("v0_create_aa_pdp__v1_init_pdp_ctx"); ok {
		t, err := expandObjectGtpMessageFilterV0V1V0CreateAaPdpV1InitPdpCtx(d, v, "v0_create_aa_pdp__v1_init_pdp_ctx")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["v0-create-aa-pdp--v1-init-pdp-ctx"] = t
		}
	}

	if v, ok := d.GetOk("version_not_support"); ok {
		t, err := expandObjectGtpMessageFilterV0V1VersionNotSupport(d, v, "version_not_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version-not-support"] = t
		}
	}

	return &obj, nil
}
