// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure ICAP profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectIcapProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectIcapProfileCreate,
		Read:   resourceObjectIcapProfileRead,
		Update: resourceObjectIcapProfileUpdate,
		Delete: resourceObjectIcapProfileDelete,

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
			"n204_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"n204_size_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"chunk_encap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"extension_feature": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"file_transfer": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"file_transfer_failure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_transfer_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"file_transfer_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"icap_block_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icap_headers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"base64_encoding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"methods": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"ocr_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preview": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preview_data_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_failure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"respmod_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"respmod_forward_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header_group": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"case_sensitivity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"header": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"header_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_resp_status_code": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_failure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_req_hdr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scan_progress_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"streaming_content_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceObjectIcapProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectIcapProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIcapProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectIcapProfile(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIcapProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectIcapProfileRead(d, m)
}

func resourceObjectIcapProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectIcapProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIcapProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectIcapProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIcapProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectIcapProfileRead(d, m)
}

func resourceObjectIcapProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectIcapProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectIcapProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectIcapProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectIcapProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIcapProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectIcapProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIcapProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectIcapProfile204Response(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfile204SizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileChunkEncap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileExtensionFeature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIcapProfileFileTransfer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIcapProfileFileTransferFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileFileTransferPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileFileTransferServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectIcapProfileIcapBlockLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileIcapHeaders(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if _, ok := i["base64-encoding"]; ok {
			v := flattenObjectIcapProfileIcapHeadersBase64Encoding(i["base64-encoding"], d, pre_append)
			tmp["base64_encoding"] = fortiAPISubPartPatch(v, "ObjectIcapProfile-IcapHeaders-Base64Encoding")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := i["content"]; ok {
			v := flattenObjectIcapProfileIcapHeadersContent(i["content"], d, pre_append)
			tmp["content"] = fortiAPISubPartPatch(v, "ObjectIcapProfile-IcapHeaders-Content")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectIcapProfileIcapHeadersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectIcapProfile-IcapHeaders-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectIcapProfileIcapHeadersName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectIcapProfile-IcapHeaders-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectIcapProfileIcapHeadersBase64Encoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileIcapHeadersContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileIcapHeadersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileIcapHeadersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileMethods(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIcapProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileOcrOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfilePreview(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfilePreviewDataLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectIcapProfileRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRequestFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRequestPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRequestServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectIcapProfileRespmodDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRespmodForwardRules(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectIcapProfileRespmodForwardRulesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectIcapProfile-RespmodForwardRules-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_group"
		if _, ok := i["header-group"]; ok {
			v := flattenObjectIcapProfileRespmodForwardRulesHeaderGroup(i["header-group"], d, pre_append)
			tmp["header_group"] = fortiAPISubPartPatch(v, "ObjectIcapProfile-RespmodForwardRules-HeaderGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {
			v := flattenObjectIcapProfileRespmodForwardRulesHost(i["host"], d, pre_append)
			tmp["host"] = fortiAPISubPartPatch(v, "ObjectIcapProfile-RespmodForwardRules-Host")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_resp_status_code"
		if _, ok := i["http-resp-status-code"]; ok {
			v := flattenObjectIcapProfileRespmodForwardRulesHttpRespStatusCode(i["http-resp-status-code"], d, pre_append)
			tmp["http_resp_status_code"] = fortiAPISubPartPatch(v, "ObjectIcapProfile-RespmodForwardRules-HttpRespStatusCode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectIcapProfileRespmodForwardRulesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectIcapProfile-RespmodForwardRules-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectIcapProfileRespmodForwardRulesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRespmodForwardRulesHeaderGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := i["case-sensitivity"]; ok {
			v := flattenObjectIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity(i["case-sensitivity"], d, pre_append)
			tmp["case_sensitivity"] = fortiAPISubPartPatch(v, "ObjectIcapProfileRespmodForwardRules-HeaderGroup-CaseSensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectIcapProfileRespmodForwardRulesHeaderGroupHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectIcapProfileRespmodForwardRules-HeaderGroup-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenObjectIcapProfileRespmodForwardRulesHeaderGroupHeaderName(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "ObjectIcapProfileRespmodForwardRules-HeaderGroup-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectIcapProfileRespmodForwardRulesHeaderGroupId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectIcapProfileRespmodForwardRules-HeaderGroup-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRespmodForwardRulesHeaderGroupHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRespmodForwardRulesHeaderGroupHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRespmodForwardRulesHeaderGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileRespmodForwardRulesHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectIcapProfileRespmodForwardRulesHttpRespStatusCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectIcapProfileRespmodForwardRulesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileResponseFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileResponsePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileResponseReqHdr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileResponseServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectIcapProfileScanProgressInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileStreamingContentBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectIcapProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("n204_response", flattenObjectIcapProfile204Response(o["204-response"], d, "n204_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["204-response"], "ObjectIcapProfile-204Response"); ok {
			if err = d.Set("n204_response", vv); err != nil {
				return fmt.Errorf("Error reading n204_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n204_response: %v", err)
		}
	}

	if err = d.Set("n204_size_limit", flattenObjectIcapProfile204SizeLimit(o["204-size-limit"], d, "n204_size_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["204-size-limit"], "ObjectIcapProfile-204SizeLimit"); ok {
			if err = d.Set("n204_size_limit", vv); err != nil {
				return fmt.Errorf("Error reading n204_size_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n204_size_limit: %v", err)
		}
	}

	if err = d.Set("chunk_encap", flattenObjectIcapProfileChunkEncap(o["chunk-encap"], d, "chunk_encap")); err != nil {
		if vv, ok := fortiAPIPatch(o["chunk-encap"], "ObjectIcapProfile-ChunkEncap"); ok {
			if err = d.Set("chunk_encap", vv); err != nil {
				return fmt.Errorf("Error reading chunk_encap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading chunk_encap: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectIcapProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectIcapProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("extension_feature", flattenObjectIcapProfileExtensionFeature(o["extension-feature"], d, "extension_feature")); err != nil {
		if vv, ok := fortiAPIPatch(o["extension-feature"], "ObjectIcapProfile-ExtensionFeature"); ok {
			if err = d.Set("extension_feature", vv); err != nil {
				return fmt.Errorf("Error reading extension_feature: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extension_feature: %v", err)
		}
	}

	if err = d.Set("file_transfer", flattenObjectIcapProfileFileTransfer(o["file-transfer"], d, "file_transfer")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-transfer"], "ObjectIcapProfile-FileTransfer"); ok {
			if err = d.Set("file_transfer", vv); err != nil {
				return fmt.Errorf("Error reading file_transfer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_transfer: %v", err)
		}
	}

	if err = d.Set("file_transfer_failure", flattenObjectIcapProfileFileTransferFailure(o["file-transfer-failure"], d, "file_transfer_failure")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-transfer-failure"], "ObjectIcapProfile-FileTransferFailure"); ok {
			if err = d.Set("file_transfer_failure", vv); err != nil {
				return fmt.Errorf("Error reading file_transfer_failure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_transfer_failure: %v", err)
		}
	}

	if err = d.Set("file_transfer_path", flattenObjectIcapProfileFileTransferPath(o["file-transfer-path"], d, "file_transfer_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-transfer-path"], "ObjectIcapProfile-FileTransferPath"); ok {
			if err = d.Set("file_transfer_path", vv); err != nil {
				return fmt.Errorf("Error reading file_transfer_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_transfer_path: %v", err)
		}
	}

	if err = d.Set("file_transfer_server", flattenObjectIcapProfileFileTransferServer(o["file-transfer-server"], d, "file_transfer_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-transfer-server"], "ObjectIcapProfile-FileTransferServer"); ok {
			if err = d.Set("file_transfer_server", vv); err != nil {
				return fmt.Errorf("Error reading file_transfer_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_transfer_server: %v", err)
		}
	}

	if err = d.Set("icap_block_log", flattenObjectIcapProfileIcapBlockLog(o["icap-block-log"], d, "icap_block_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["icap-block-log"], "ObjectIcapProfile-IcapBlockLog"); ok {
			if err = d.Set("icap_block_log", vv); err != nil {
				return fmt.Errorf("Error reading icap_block_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icap_block_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("icap_headers", flattenObjectIcapProfileIcapHeaders(o["icap-headers"], d, "icap_headers")); err != nil {
			if vv, ok := fortiAPIPatch(o["icap-headers"], "ObjectIcapProfile-IcapHeaders"); ok {
				if err = d.Set("icap_headers", vv); err != nil {
					return fmt.Errorf("Error reading icap_headers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading icap_headers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("icap_headers"); ok {
			if err = d.Set("icap_headers", flattenObjectIcapProfileIcapHeaders(o["icap-headers"], d, "icap_headers")); err != nil {
				if vv, ok := fortiAPIPatch(o["icap-headers"], "ObjectIcapProfile-IcapHeaders"); ok {
					if err = d.Set("icap_headers", vv); err != nil {
						return fmt.Errorf("Error reading icap_headers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading icap_headers: %v", err)
				}
			}
		}
	}

	if err = d.Set("methods", flattenObjectIcapProfileMethods(o["methods"], d, "methods")); err != nil {
		if vv, ok := fortiAPIPatch(o["methods"], "ObjectIcapProfile-Methods"); ok {
			if err = d.Set("methods", vv); err != nil {
				return fmt.Errorf("Error reading methods: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading methods: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectIcapProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectIcapProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ocr_only", flattenObjectIcapProfileOcrOnly(o["ocr-only"], d, "ocr_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["ocr-only"], "ObjectIcapProfile-OcrOnly"); ok {
			if err = d.Set("ocr_only", vv); err != nil {
				return fmt.Errorf("Error reading ocr_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ocr_only: %v", err)
		}
	}

	if err = d.Set("preview", flattenObjectIcapProfilePreview(o["preview"], d, "preview")); err != nil {
		if vv, ok := fortiAPIPatch(o["preview"], "ObjectIcapProfile-Preview"); ok {
			if err = d.Set("preview", vv); err != nil {
				return fmt.Errorf("Error reading preview: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preview: %v", err)
		}
	}

	if err = d.Set("preview_data_length", flattenObjectIcapProfilePreviewDataLength(o["preview-data-length"], d, "preview_data_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["preview-data-length"], "ObjectIcapProfile-PreviewDataLength"); ok {
			if err = d.Set("preview_data_length", vv); err != nil {
				return fmt.Errorf("Error reading preview_data_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preview_data_length: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenObjectIcapProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "ObjectIcapProfile-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("request", flattenObjectIcapProfileRequest(o["request"], d, "request")); err != nil {
		if vv, ok := fortiAPIPatch(o["request"], "ObjectIcapProfile-Request"); ok {
			if err = d.Set("request", vv); err != nil {
				return fmt.Errorf("Error reading request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading request: %v", err)
		}
	}

	if err = d.Set("request_failure", flattenObjectIcapProfileRequestFailure(o["request-failure"], d, "request_failure")); err != nil {
		if vv, ok := fortiAPIPatch(o["request-failure"], "ObjectIcapProfile-RequestFailure"); ok {
			if err = d.Set("request_failure", vv); err != nil {
				return fmt.Errorf("Error reading request_failure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading request_failure: %v", err)
		}
	}

	if err = d.Set("request_path", flattenObjectIcapProfileRequestPath(o["request-path"], d, "request_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["request-path"], "ObjectIcapProfile-RequestPath"); ok {
			if err = d.Set("request_path", vv); err != nil {
				return fmt.Errorf("Error reading request_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading request_path: %v", err)
		}
	}

	if err = d.Set("request_server", flattenObjectIcapProfileRequestServer(o["request-server"], d, "request_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["request-server"], "ObjectIcapProfile-RequestServer"); ok {
			if err = d.Set("request_server", vv); err != nil {
				return fmt.Errorf("Error reading request_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading request_server: %v", err)
		}
	}

	if err = d.Set("respmod_default_action", flattenObjectIcapProfileRespmodDefaultAction(o["respmod-default-action"], d, "respmod_default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["respmod-default-action"], "ObjectIcapProfile-RespmodDefaultAction"); ok {
			if err = d.Set("respmod_default_action", vv); err != nil {
				return fmt.Errorf("Error reading respmod_default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading respmod_default_action: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("respmod_forward_rules", flattenObjectIcapProfileRespmodForwardRules(o["respmod-forward-rules"], d, "respmod_forward_rules")); err != nil {
			if vv, ok := fortiAPIPatch(o["respmod-forward-rules"], "ObjectIcapProfile-RespmodForwardRules"); ok {
				if err = d.Set("respmod_forward_rules", vv); err != nil {
					return fmt.Errorf("Error reading respmod_forward_rules: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading respmod_forward_rules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("respmod_forward_rules"); ok {
			if err = d.Set("respmod_forward_rules", flattenObjectIcapProfileRespmodForwardRules(o["respmod-forward-rules"], d, "respmod_forward_rules")); err != nil {
				if vv, ok := fortiAPIPatch(o["respmod-forward-rules"], "ObjectIcapProfile-RespmodForwardRules"); ok {
					if err = d.Set("respmod_forward_rules", vv); err != nil {
						return fmt.Errorf("Error reading respmod_forward_rules: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading respmod_forward_rules: %v", err)
				}
			}
		}
	}

	if err = d.Set("response", flattenObjectIcapProfileResponse(o["response"], d, "response")); err != nil {
		if vv, ok := fortiAPIPatch(o["response"], "ObjectIcapProfile-Response"); ok {
			if err = d.Set("response", vv); err != nil {
				return fmt.Errorf("Error reading response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading response: %v", err)
		}
	}

	if err = d.Set("response_failure", flattenObjectIcapProfileResponseFailure(o["response-failure"], d, "response_failure")); err != nil {
		if vv, ok := fortiAPIPatch(o["response-failure"], "ObjectIcapProfile-ResponseFailure"); ok {
			if err = d.Set("response_failure", vv); err != nil {
				return fmt.Errorf("Error reading response_failure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading response_failure: %v", err)
		}
	}

	if err = d.Set("response_path", flattenObjectIcapProfileResponsePath(o["response-path"], d, "response_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["response-path"], "ObjectIcapProfile-ResponsePath"); ok {
			if err = d.Set("response_path", vv); err != nil {
				return fmt.Errorf("Error reading response_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading response_path: %v", err)
		}
	}

	if err = d.Set("response_req_hdr", flattenObjectIcapProfileResponseReqHdr(o["response-req-hdr"], d, "response_req_hdr")); err != nil {
		if vv, ok := fortiAPIPatch(o["response-req-hdr"], "ObjectIcapProfile-ResponseReqHdr"); ok {
			if err = d.Set("response_req_hdr", vv); err != nil {
				return fmt.Errorf("Error reading response_req_hdr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading response_req_hdr: %v", err)
		}
	}

	if err = d.Set("response_server", flattenObjectIcapProfileResponseServer(o["response-server"], d, "response_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["response-server"], "ObjectIcapProfile-ResponseServer"); ok {
			if err = d.Set("response_server", vv); err != nil {
				return fmt.Errorf("Error reading response_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading response_server: %v", err)
		}
	}

	if err = d.Set("scan_progress_interval", flattenObjectIcapProfileScanProgressInterval(o["scan-progress-interval"], d, "scan_progress_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-progress-interval"], "ObjectIcapProfile-ScanProgressInterval"); ok {
			if err = d.Set("scan_progress_interval", vv); err != nil {
				return fmt.Errorf("Error reading scan_progress_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_progress_interval: %v", err)
		}
	}

	if err = d.Set("streaming_content_bypass", flattenObjectIcapProfileStreamingContentBypass(o["streaming-content-bypass"], d, "streaming_content_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["streaming-content-bypass"], "ObjectIcapProfile-StreamingContentBypass"); ok {
			if err = d.Set("streaming_content_bypass", vv); err != nil {
				return fmt.Errorf("Error reading streaming_content_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading streaming_content_bypass: %v", err)
		}
	}

	if err = d.Set("timeout", flattenObjectIcapProfileTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "ObjectIcapProfile-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	return nil
}

func flattenObjectIcapProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectIcapProfile204Response(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfile204SizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileChunkEncap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileExtensionFeature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIcapProfileFileTransfer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIcapProfileFileTransferFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileFileTransferPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileFileTransferServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectIcapProfileIcapBlockLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileIcapHeaders(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["base64-encoding"], _ = expandObjectIcapProfileIcapHeadersBase64Encoding(d, i["base64_encoding"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["content"], _ = expandObjectIcapProfileIcapHeadersContent(d, i["content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectIcapProfileIcapHeadersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectIcapProfileIcapHeadersName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectIcapProfileIcapHeadersBase64Encoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileIcapHeadersContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileIcapHeadersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileIcapHeadersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileMethods(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIcapProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileOcrOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfilePreview(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfilePreviewDataLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectIcapProfileRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRequestFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRequestPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRequestServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectIcapProfileRespmodDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRespmodForwardRules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectIcapProfileRespmodForwardRulesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectIcapProfileRespmodForwardRulesHeaderGroup(d, i["header_group"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["header-group"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host"], _ = expandObjectIcapProfileRespmodForwardRulesHost(d, i["host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_resp_status_code"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-resp-status-code"], _ = expandObjectIcapProfileRespmodForwardRulesHttpRespStatusCode(d, i["http_resp_status_code"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectIcapProfileRespmodForwardRulesName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectIcapProfileRespmodForwardRulesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRespmodForwardRulesHeaderGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitivity"], _ = expandObjectIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity(d, i["case_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectIcapProfileRespmodForwardRulesHeaderGroupHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandObjectIcapProfileRespmodForwardRulesHeaderGroupHeaderName(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectIcapProfileRespmodForwardRulesHeaderGroupId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRespmodForwardRulesHeaderGroupHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRespmodForwardRulesHeaderGroupHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRespmodForwardRulesHeaderGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileRespmodForwardRulesHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectIcapProfileRespmodForwardRulesHttpRespStatusCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectIcapProfileRespmodForwardRulesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileResponseFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileResponsePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileResponseReqHdr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileResponseServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectIcapProfileScanProgressInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileStreamingContentBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectIcapProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("n204_response"); ok || d.HasChange("n204_response") {
		t, err := expandObjectIcapProfile204Response(d, v, "n204_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["204-response"] = t
		}
	}

	if v, ok := d.GetOk("n204_size_limit"); ok || d.HasChange("n204_size_limit") {
		t, err := expandObjectIcapProfile204SizeLimit(d, v, "n204_size_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["204-size-limit"] = t
		}
	}

	if v, ok := d.GetOk("chunk_encap"); ok || d.HasChange("chunk_encap") {
		t, err := expandObjectIcapProfileChunkEncap(d, v, "chunk_encap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chunk-encap"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectIcapProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("extension_feature"); ok || d.HasChange("extension_feature") {
		t, err := expandObjectIcapProfileExtensionFeature(d, v, "extension_feature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension-feature"] = t
		}
	}

	if v, ok := d.GetOk("file_transfer"); ok || d.HasChange("file_transfer") {
		t, err := expandObjectIcapProfileFileTransfer(d, v, "file_transfer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-transfer"] = t
		}
	}

	if v, ok := d.GetOk("file_transfer_failure"); ok || d.HasChange("file_transfer_failure") {
		t, err := expandObjectIcapProfileFileTransferFailure(d, v, "file_transfer_failure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-transfer-failure"] = t
		}
	}

	if v, ok := d.GetOk("file_transfer_path"); ok || d.HasChange("file_transfer_path") {
		t, err := expandObjectIcapProfileFileTransferPath(d, v, "file_transfer_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-transfer-path"] = t
		}
	}

	if v, ok := d.GetOk("file_transfer_server"); ok || d.HasChange("file_transfer_server") {
		t, err := expandObjectIcapProfileFileTransferServer(d, v, "file_transfer_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-transfer-server"] = t
		}
	}

	if v, ok := d.GetOk("icap_block_log"); ok || d.HasChange("icap_block_log") {
		t, err := expandObjectIcapProfileIcapBlockLog(d, v, "icap_block_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-block-log"] = t
		}
	}

	if v, ok := d.GetOk("icap_headers"); ok || d.HasChange("icap_headers") {
		t, err := expandObjectIcapProfileIcapHeaders(d, v, "icap_headers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-headers"] = t
		}
	}

	if v, ok := d.GetOk("methods"); ok || d.HasChange("methods") {
		t, err := expandObjectIcapProfileMethods(d, v, "methods")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["methods"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectIcapProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ocr_only"); ok || d.HasChange("ocr_only") {
		t, err := expandObjectIcapProfileOcrOnly(d, v, "ocr_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocr-only"] = t
		}
	}

	if v, ok := d.GetOk("preview"); ok || d.HasChange("preview") {
		t, err := expandObjectIcapProfilePreview(d, v, "preview")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preview"] = t
		}
	}

	if v, ok := d.GetOk("preview_data_length"); ok || d.HasChange("preview_data_length") {
		t, err := expandObjectIcapProfilePreviewDataLength(d, v, "preview_data_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preview-data-length"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandObjectIcapProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("request"); ok || d.HasChange("request") {
		t, err := expandObjectIcapProfileRequest(d, v, "request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request"] = t
		}
	}

	if v, ok := d.GetOk("request_failure"); ok || d.HasChange("request_failure") {
		t, err := expandObjectIcapProfileRequestFailure(d, v, "request_failure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-failure"] = t
		}
	}

	if v, ok := d.GetOk("request_path"); ok || d.HasChange("request_path") {
		t, err := expandObjectIcapProfileRequestPath(d, v, "request_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-path"] = t
		}
	}

	if v, ok := d.GetOk("request_server"); ok || d.HasChange("request_server") {
		t, err := expandObjectIcapProfileRequestServer(d, v, "request_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-server"] = t
		}
	}

	if v, ok := d.GetOk("respmod_default_action"); ok || d.HasChange("respmod_default_action") {
		t, err := expandObjectIcapProfileRespmodDefaultAction(d, v, "respmod_default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["respmod-default-action"] = t
		}
	}

	if v, ok := d.GetOk("respmod_forward_rules"); ok || d.HasChange("respmod_forward_rules") {
		t, err := expandObjectIcapProfileRespmodForwardRules(d, v, "respmod_forward_rules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["respmod-forward-rules"] = t
		}
	}

	if v, ok := d.GetOk("response"); ok || d.HasChange("response") {
		t, err := expandObjectIcapProfileResponse(d, v, "response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response"] = t
		}
	}

	if v, ok := d.GetOk("response_failure"); ok || d.HasChange("response_failure") {
		t, err := expandObjectIcapProfileResponseFailure(d, v, "response_failure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-failure"] = t
		}
	}

	if v, ok := d.GetOk("response_path"); ok || d.HasChange("response_path") {
		t, err := expandObjectIcapProfileResponsePath(d, v, "response_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-path"] = t
		}
	}

	if v, ok := d.GetOk("response_req_hdr"); ok || d.HasChange("response_req_hdr") {
		t, err := expandObjectIcapProfileResponseReqHdr(d, v, "response_req_hdr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-req-hdr"] = t
		}
	}

	if v, ok := d.GetOk("response_server"); ok || d.HasChange("response_server") {
		t, err := expandObjectIcapProfileResponseServer(d, v, "response_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-server"] = t
		}
	}

	if v, ok := d.GetOk("scan_progress_interval"); ok || d.HasChange("scan_progress_interval") {
		t, err := expandObjectIcapProfileScanProgressInterval(d, v, "scan_progress_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-progress-interval"] = t
		}
	}

	if v, ok := d.GetOk("streaming_content_bypass"); ok || d.HasChange("streaming_content_bypass") {
		t, err := expandObjectIcapProfileStreamingContentBypass(d, v, "streaming_content_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["streaming-content-bypass"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandObjectIcapProfileTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	return &obj, nil
}
