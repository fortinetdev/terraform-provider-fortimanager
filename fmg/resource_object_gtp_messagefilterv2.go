// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Message filter for GTPv2 messages.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectGtpMessageFilterV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectGtpMessageFilterV2Create,
		Read:   resourceObjectGtpMessageFilterV2Read,
		Update: resourceObjectGtpMessageFilterV2Update,
		Delete: resourceObjectGtpMessageFilterV2Delete,

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
			"bearer_resource_cmd_fail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"change_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"create_bearer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"create_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delete_bearer_cmd_fail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delete_bearer_req_resp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delete_pdn_connection_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delete_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"echo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"modify_bearer_cmd_fail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"modify_bearer_req_resp": &schema.Schema{
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
			"resume": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"suspend": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trace_session": &schema.Schema{
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
			"update_bearer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_pdn_connection_set": &schema.Schema{
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

func resourceObjectGtpMessageFilterV2Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectGtpMessageFilterV2(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectGtpMessageFilterV2 resource while getting object: %v", err)
	}

	_, err = c.CreateObjectGtpMessageFilterV2(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectGtpMessageFilterV2 resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectGtpMessageFilterV2Read(d, m)
}

func resourceObjectGtpMessageFilterV2Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectGtpMessageFilterV2(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGtpMessageFilterV2 resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectGtpMessageFilterV2(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGtpMessageFilterV2 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectGtpMessageFilterV2Read(d, m)
}

func resourceObjectGtpMessageFilterV2Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectGtpMessageFilterV2(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectGtpMessageFilterV2 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectGtpMessageFilterV2Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectGtpMessageFilterV2(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGtpMessageFilterV2 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectGtpMessageFilterV2(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGtpMessageFilterV2 resource from API: %v", err)
	}
	return nil
}

func flattenObjectGtpMessageFilterV2BearerResourceCmdFail(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2ChangeNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2CreateBearer(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2CreateSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2DeleteBearerCmdFail(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2DeleteBearerReqResp(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2DeletePdnConnectionSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2DeleteSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2Echo(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2ModifyBearerCmdFail(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2ModifyBearerReqResp(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpMessageFilterV2Resume(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2Suspend(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2TraceSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2UnknownMessage(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2UnknownMessageWhiteList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectGtpMessageFilterV2UpdateBearer(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2UpdatePdnConnectionSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectGtpMessageFilterV2VersionNotSupport(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func refreshObjectObjectGtpMessageFilterV2(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("bearer_resource_cmd_fail", flattenObjectGtpMessageFilterV2BearerResourceCmdFail(o["bearer-resource-cmd-fail"], d, "bearer_resource_cmd_fail")); err != nil {
		if vv, ok := fortiAPIPatch(o["bearer-resource-cmd-fail"], "ObjectGtpMessageFilterV2-BearerResourceCmdFail"); ok {
			if err = d.Set("bearer_resource_cmd_fail", vv); err != nil {
				return fmt.Errorf("Error reading bearer_resource_cmd_fail: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bearer_resource_cmd_fail: %v", err)
		}
	}

	if err = d.Set("change_notification", flattenObjectGtpMessageFilterV2ChangeNotification(o["change-notification"], d, "change_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["change-notification"], "ObjectGtpMessageFilterV2-ChangeNotification"); ok {
			if err = d.Set("change_notification", vv); err != nil {
				return fmt.Errorf("Error reading change_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading change_notification: %v", err)
		}
	}

	if err = d.Set("create_bearer", flattenObjectGtpMessageFilterV2CreateBearer(o["create-bearer"], d, "create_bearer")); err != nil {
		if vv, ok := fortiAPIPatch(o["create-bearer"], "ObjectGtpMessageFilterV2-CreateBearer"); ok {
			if err = d.Set("create_bearer", vv); err != nil {
				return fmt.Errorf("Error reading create_bearer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading create_bearer: %v", err)
		}
	}

	if err = d.Set("create_session", flattenObjectGtpMessageFilterV2CreateSession(o["create-session"], d, "create_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["create-session"], "ObjectGtpMessageFilterV2-CreateSession"); ok {
			if err = d.Set("create_session", vv); err != nil {
				return fmt.Errorf("Error reading create_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading create_session: %v", err)
		}
	}

	if err = d.Set("delete_bearer_cmd_fail", flattenObjectGtpMessageFilterV2DeleteBearerCmdFail(o["delete-bearer-cmd-fail"], d, "delete_bearer_cmd_fail")); err != nil {
		if vv, ok := fortiAPIPatch(o["delete-bearer-cmd-fail"], "ObjectGtpMessageFilterV2-DeleteBearerCmdFail"); ok {
			if err = d.Set("delete_bearer_cmd_fail", vv); err != nil {
				return fmt.Errorf("Error reading delete_bearer_cmd_fail: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delete_bearer_cmd_fail: %v", err)
		}
	}

	if err = d.Set("delete_bearer_req_resp", flattenObjectGtpMessageFilterV2DeleteBearerReqResp(o["delete-bearer-req-resp"], d, "delete_bearer_req_resp")); err != nil {
		if vv, ok := fortiAPIPatch(o["delete-bearer-req-resp"], "ObjectGtpMessageFilterV2-DeleteBearerReqResp"); ok {
			if err = d.Set("delete_bearer_req_resp", vv); err != nil {
				return fmt.Errorf("Error reading delete_bearer_req_resp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delete_bearer_req_resp: %v", err)
		}
	}

	if err = d.Set("delete_pdn_connection_set", flattenObjectGtpMessageFilterV2DeletePdnConnectionSet(o["delete-pdn-connection-set"], d, "delete_pdn_connection_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["delete-pdn-connection-set"], "ObjectGtpMessageFilterV2-DeletePdnConnectionSet"); ok {
			if err = d.Set("delete_pdn_connection_set", vv); err != nil {
				return fmt.Errorf("Error reading delete_pdn_connection_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delete_pdn_connection_set: %v", err)
		}
	}

	if err = d.Set("delete_session", flattenObjectGtpMessageFilterV2DeleteSession(o["delete-session"], d, "delete_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["delete-session"], "ObjectGtpMessageFilterV2-DeleteSession"); ok {
			if err = d.Set("delete_session", vv); err != nil {
				return fmt.Errorf("Error reading delete_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delete_session: %v", err)
		}
	}

	if err = d.Set("echo", flattenObjectGtpMessageFilterV2Echo(o["echo"], d, "echo")); err != nil {
		if vv, ok := fortiAPIPatch(o["echo"], "ObjectGtpMessageFilterV2-Echo"); ok {
			if err = d.Set("echo", vv); err != nil {
				return fmt.Errorf("Error reading echo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading echo: %v", err)
		}
	}

	if err = d.Set("modify_bearer_cmd_fail", flattenObjectGtpMessageFilterV2ModifyBearerCmdFail(o["modify-bearer-cmd-fail"], d, "modify_bearer_cmd_fail")); err != nil {
		if vv, ok := fortiAPIPatch(o["modify-bearer-cmd-fail"], "ObjectGtpMessageFilterV2-ModifyBearerCmdFail"); ok {
			if err = d.Set("modify_bearer_cmd_fail", vv); err != nil {
				return fmt.Errorf("Error reading modify_bearer_cmd_fail: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modify_bearer_cmd_fail: %v", err)
		}
	}

	if err = d.Set("modify_bearer_req_resp", flattenObjectGtpMessageFilterV2ModifyBearerReqResp(o["modify-bearer-req-resp"], d, "modify_bearer_req_resp")); err != nil {
		if vv, ok := fortiAPIPatch(o["modify-bearer-req-resp"], "ObjectGtpMessageFilterV2-ModifyBearerReqResp"); ok {
			if err = d.Set("modify_bearer_req_resp", vv); err != nil {
				return fmt.Errorf("Error reading modify_bearer_req_resp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modify_bearer_req_resp: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectGtpMessageFilterV2Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectGtpMessageFilterV2-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("resume", flattenObjectGtpMessageFilterV2Resume(o["resume"], d, "resume")); err != nil {
		if vv, ok := fortiAPIPatch(o["resume"], "ObjectGtpMessageFilterV2-Resume"); ok {
			if err = d.Set("resume", vv); err != nil {
				return fmt.Errorf("Error reading resume: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resume: %v", err)
		}
	}

	if err = d.Set("suspend", flattenObjectGtpMessageFilterV2Suspend(o["suspend"], d, "suspend")); err != nil {
		if vv, ok := fortiAPIPatch(o["suspend"], "ObjectGtpMessageFilterV2-Suspend"); ok {
			if err = d.Set("suspend", vv); err != nil {
				return fmt.Errorf("Error reading suspend: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading suspend: %v", err)
		}
	}

	if err = d.Set("trace_session", flattenObjectGtpMessageFilterV2TraceSession(o["trace-session"], d, "trace_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["trace-session"], "ObjectGtpMessageFilterV2-TraceSession"); ok {
			if err = d.Set("trace_session", vv); err != nil {
				return fmt.Errorf("Error reading trace_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trace_session: %v", err)
		}
	}

	if err = d.Set("unknown_message", flattenObjectGtpMessageFilterV2UnknownMessage(o["unknown-message"], d, "unknown_message")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-message"], "ObjectGtpMessageFilterV2-UnknownMessage"); ok {
			if err = d.Set("unknown_message", vv); err != nil {
				return fmt.Errorf("Error reading unknown_message: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_message: %v", err)
		}
	}

	if err = d.Set("unknown_message_white_list", flattenObjectGtpMessageFilterV2UnknownMessageWhiteList(o["unknown-message-white-list"], d, "unknown_message_white_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-message-white-list"], "ObjectGtpMessageFilterV2-UnknownMessageWhiteList"); ok {
			if err = d.Set("unknown_message_white_list", vv); err != nil {
				return fmt.Errorf("Error reading unknown_message_white_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_message_white_list: %v", err)
		}
	}

	if err = d.Set("update_bearer", flattenObjectGtpMessageFilterV2UpdateBearer(o["update-bearer"], d, "update_bearer")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-bearer"], "ObjectGtpMessageFilterV2-UpdateBearer"); ok {
			if err = d.Set("update_bearer", vv); err != nil {
				return fmt.Errorf("Error reading update_bearer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_bearer: %v", err)
		}
	}

	if err = d.Set("update_pdn_connection_set", flattenObjectGtpMessageFilterV2UpdatePdnConnectionSet(o["update-pdn-connection-set"], d, "update_pdn_connection_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-pdn-connection-set"], "ObjectGtpMessageFilterV2-UpdatePdnConnectionSet"); ok {
			if err = d.Set("update_pdn_connection_set", vv); err != nil {
				return fmt.Errorf("Error reading update_pdn_connection_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_pdn_connection_set: %v", err)
		}
	}

	if err = d.Set("version_not_support", flattenObjectGtpMessageFilterV2VersionNotSupport(o["version-not-support"], d, "version_not_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["version-not-support"], "ObjectGtpMessageFilterV2-VersionNotSupport"); ok {
			if err = d.Set("version_not_support", vv); err != nil {
				return fmt.Errorf("Error reading version_not_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version_not_support: %v", err)
		}
	}

	return nil
}

func flattenObjectGtpMessageFilterV2FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectGtpMessageFilterV2BearerResourceCmdFail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2ChangeNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2CreateBearer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2CreateSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2DeleteBearerCmdFail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2DeleteBearerReqResp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2DeletePdnConnectionSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2DeleteSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2Echo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2ModifyBearerCmdFail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2ModifyBearerReqResp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2Resume(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2Suspend(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2TraceSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2UnknownMessage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2UnknownMessageWhiteList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectGtpMessageFilterV2UpdateBearer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2UpdatePdnConnectionSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpMessageFilterV2VersionNotSupport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectGtpMessageFilterV2(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bearer_resource_cmd_fail"); ok {
		t, err := expandObjectGtpMessageFilterV2BearerResourceCmdFail(d, v, "bearer_resource_cmd_fail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bearer-resource-cmd-fail"] = t
		}
	}

	if v, ok := d.GetOk("change_notification"); ok {
		t, err := expandObjectGtpMessageFilterV2ChangeNotification(d, v, "change_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["change-notification"] = t
		}
	}

	if v, ok := d.GetOk("create_bearer"); ok {
		t, err := expandObjectGtpMessageFilterV2CreateBearer(d, v, "create_bearer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create-bearer"] = t
		}
	}

	if v, ok := d.GetOk("create_session"); ok {
		t, err := expandObjectGtpMessageFilterV2CreateSession(d, v, "create_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create-session"] = t
		}
	}

	if v, ok := d.GetOk("delete_bearer_cmd_fail"); ok {
		t, err := expandObjectGtpMessageFilterV2DeleteBearerCmdFail(d, v, "delete_bearer_cmd_fail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delete-bearer-cmd-fail"] = t
		}
	}

	if v, ok := d.GetOk("delete_bearer_req_resp"); ok {
		t, err := expandObjectGtpMessageFilterV2DeleteBearerReqResp(d, v, "delete_bearer_req_resp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delete-bearer-req-resp"] = t
		}
	}

	if v, ok := d.GetOk("delete_pdn_connection_set"); ok {
		t, err := expandObjectGtpMessageFilterV2DeletePdnConnectionSet(d, v, "delete_pdn_connection_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delete-pdn-connection-set"] = t
		}
	}

	if v, ok := d.GetOk("delete_session"); ok {
		t, err := expandObjectGtpMessageFilterV2DeleteSession(d, v, "delete_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delete-session"] = t
		}
	}

	if v, ok := d.GetOk("echo"); ok {
		t, err := expandObjectGtpMessageFilterV2Echo(d, v, "echo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["echo"] = t
		}
	}

	if v, ok := d.GetOk("modify_bearer_cmd_fail"); ok {
		t, err := expandObjectGtpMessageFilterV2ModifyBearerCmdFail(d, v, "modify_bearer_cmd_fail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modify-bearer-cmd-fail"] = t
		}
	}

	if v, ok := d.GetOk("modify_bearer_req_resp"); ok {
		t, err := expandObjectGtpMessageFilterV2ModifyBearerReqResp(d, v, "modify_bearer_req_resp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modify-bearer-req-resp"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectGtpMessageFilterV2Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("resume"); ok {
		t, err := expandObjectGtpMessageFilterV2Resume(d, v, "resume")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resume"] = t
		}
	}

	if v, ok := d.GetOk("suspend"); ok {
		t, err := expandObjectGtpMessageFilterV2Suspend(d, v, "suspend")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["suspend"] = t
		}
	}

	if v, ok := d.GetOk("trace_session"); ok {
		t, err := expandObjectGtpMessageFilterV2TraceSession(d, v, "trace_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trace-session"] = t
		}
	}

	if v, ok := d.GetOk("unknown_message"); ok {
		t, err := expandObjectGtpMessageFilterV2UnknownMessage(d, v, "unknown_message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-message"] = t
		}
	}

	if v, ok := d.GetOk("unknown_message_white_list"); ok {
		t, err := expandObjectGtpMessageFilterV2UnknownMessageWhiteList(d, v, "unknown_message_white_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-message-white-list"] = t
		}
	}

	if v, ok := d.GetOk("update_bearer"); ok {
		t, err := expandObjectGtpMessageFilterV2UpdateBearer(d, v, "update_bearer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-bearer"] = t
		}
	}

	if v, ok := d.GetOk("update_pdn_connection_set"); ok {
		t, err := expandObjectGtpMessageFilterV2UpdatePdnConnectionSet(d, v, "update_pdn_connection_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-pdn-connection-set"] = t
		}
	}

	if v, ok := d.GetOk("version_not_support"); ok {
		t, err := expandObjectGtpMessageFilterV2VersionNotSupport(d, v, "version_not_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version-not-support"] = t
		}
	}

	return &obj, nil
}
