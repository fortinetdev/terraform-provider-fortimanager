// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure AntiSpam profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSpamfilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSpamfilterProfileCreate,
		Read:   resourceObjectSpamfilterProfileRead,
		Update: resourceObjectSpamfilterProfileUpdate,
		Delete: resourceObjectSpamfilterProfileDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"flow_based": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gmail": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tag_msg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tag_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
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
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"msn_hotmail": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"pop3": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tag_msg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tag_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
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
			"smtp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"hdrip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"local_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tag_msg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tag_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"spam_bwl_table": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"spam_bword_table": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"spam_bword_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"spam_filtering": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"spam_iptrust_table": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"spam_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"spam_log_fortiguard_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"spam_mheader_table": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"spam_rbl_table": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectSpamfilterProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSpamfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSpamfilterProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSpamfilterProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSpamfilterProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSpamfilterProfileRead(d, m)
}

func resourceObjectSpamfilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSpamfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSpamfilterProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSpamfilterProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSpamfilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSpamfilterProfileRead(d, m)
}

func resourceObjectSpamfilterProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSpamfilterProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSpamfilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSpamfilterProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSpamfilterProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSpamfilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSpamfilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSpamfilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectSpamfilterProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileFlowBased(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileGmail(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectSpamfilterProfileGmailLog(i["log"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSpamfilterProfileGmailLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileImap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectSpamfilterProfileImapAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectSpamfilterProfileImapLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_msg"
	if _, ok := i["tag-msg"]; ok {
		result["tag_msg"] = flattenObjectSpamfilterProfileImapTagMsg(i["tag-msg"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_type"
	if _, ok := i["tag-type"]; ok {
		result["tag_type"] = flattenObjectSpamfilterProfileImapTagType(i["tag-type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSpamfilterProfileImapAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileImapLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileImapTagMsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileImapTagType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSpamfilterProfileMapi(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectSpamfilterProfileMapiAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectSpamfilterProfileMapiLog(i["log"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSpamfilterProfileMapiAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileMapiLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileMsnHotmail(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectSpamfilterProfileMsnHotmailLog(i["log"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSpamfilterProfileMsnHotmailLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSpamfilterProfilePop3(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectSpamfilterProfilePop3Action(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectSpamfilterProfilePop3Log(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_msg"
	if _, ok := i["tag-msg"]; ok {
		result["tag_msg"] = flattenObjectSpamfilterProfilePop3TagMsg(i["tag-msg"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_type"
	if _, ok := i["tag-type"]; ok {
		result["tag_type"] = flattenObjectSpamfilterProfilePop3TagType(i["tag-type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSpamfilterProfilePop3Action(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfilePop3Log(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfilePop3TagMsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfilePop3TagType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSpamfilterProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileSmtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectSpamfilterProfileSmtpAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "hdrip"
	if _, ok := i["hdrip"]; ok {
		result["hdrip"] = flattenObjectSpamfilterProfileSmtpHdrip(i["hdrip"], d, pre_append)
	}

	pre_append = pre + ".0." + "local_override"
	if _, ok := i["local-override"]; ok {
		result["local_override"] = flattenObjectSpamfilterProfileSmtpLocalOverride(i["local-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectSpamfilterProfileSmtpLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_msg"
	if _, ok := i["tag-msg"]; ok {
		result["tag_msg"] = flattenObjectSpamfilterProfileSmtpTagMsg(i["tag-msg"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_type"
	if _, ok := i["tag-type"]; ok {
		result["tag_type"] = flattenObjectSpamfilterProfileSmtpTagType(i["tag-type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSpamfilterProfileSmtpAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileSmtpHdrip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileSmtpLocalOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileSmtpLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileSmtpTagMsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileSmtpTagType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSpamfilterProfileSpamBwlTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileSpamBwordTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileSpamBwordThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileSpamFiltering(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileSpamIptrustTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileSpamLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileSpamLogFortiguardResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileSpamMheaderTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterProfileSpamRblTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSpamfilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comment", flattenObjectSpamfilterProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectSpamfilterProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("external", flattenObjectSpamfilterProfileExternal(o["external"], d, "external")); err != nil {
		if vv, ok := fortiAPIPatch(o["external"], "ObjectSpamfilterProfile-External"); ok {
			if err = d.Set("external", vv); err != nil {
				return fmt.Errorf("Error reading external: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external: %v", err)
		}
	}

	if err = d.Set("flow_based", flattenObjectSpamfilterProfileFlowBased(o["flow-based"], d, "flow_based")); err != nil {
		if vv, ok := fortiAPIPatch(o["flow-based"], "ObjectSpamfilterProfile-FlowBased"); ok {
			if err = d.Set("flow_based", vv); err != nil {
				return fmt.Errorf("Error reading flow_based: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flow_based: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("gmail", flattenObjectSpamfilterProfileGmail(o["gmail"], d, "gmail")); err != nil {
			if vv, ok := fortiAPIPatch(o["gmail"], "ObjectSpamfilterProfile-Gmail"); ok {
				if err = d.Set("gmail", vv); err != nil {
					return fmt.Errorf("Error reading gmail: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading gmail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("gmail"); ok {
			if err = d.Set("gmail", flattenObjectSpamfilterProfileGmail(o["gmail"], d, "gmail")); err != nil {
				if vv, ok := fortiAPIPatch(o["gmail"], "ObjectSpamfilterProfile-Gmail"); ok {
					if err = d.Set("gmail", vv); err != nil {
						return fmt.Errorf("Error reading gmail: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading gmail: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("imap", flattenObjectSpamfilterProfileImap(o["imap"], d, "imap")); err != nil {
			if vv, ok := fortiAPIPatch(o["imap"], "ObjectSpamfilterProfile-Imap"); ok {
				if err = d.Set("imap", vv); err != nil {
					return fmt.Errorf("Error reading imap: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading imap: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("imap"); ok {
			if err = d.Set("imap", flattenObjectSpamfilterProfileImap(o["imap"], d, "imap")); err != nil {
				if vv, ok := fortiAPIPatch(o["imap"], "ObjectSpamfilterProfile-Imap"); ok {
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
		if err = d.Set("mapi", flattenObjectSpamfilterProfileMapi(o["mapi"], d, "mapi")); err != nil {
			if vv, ok := fortiAPIPatch(o["mapi"], "ObjectSpamfilterProfile-Mapi"); ok {
				if err = d.Set("mapi", vv); err != nil {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mapi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mapi"); ok {
			if err = d.Set("mapi", flattenObjectSpamfilterProfileMapi(o["mapi"], d, "mapi")); err != nil {
				if vv, ok := fortiAPIPatch(o["mapi"], "ObjectSpamfilterProfile-Mapi"); ok {
					if err = d.Set("mapi", vv); err != nil {
						return fmt.Errorf("Error reading mapi: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("msn_hotmail", flattenObjectSpamfilterProfileMsnHotmail(o["msn-hotmail"], d, "msn_hotmail")); err != nil {
			if vv, ok := fortiAPIPatch(o["msn-hotmail"], "ObjectSpamfilterProfile-MsnHotmail"); ok {
				if err = d.Set("msn_hotmail", vv); err != nil {
					return fmt.Errorf("Error reading msn_hotmail: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading msn_hotmail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("msn_hotmail"); ok {
			if err = d.Set("msn_hotmail", flattenObjectSpamfilterProfileMsnHotmail(o["msn-hotmail"], d, "msn_hotmail")); err != nil {
				if vv, ok := fortiAPIPatch(o["msn-hotmail"], "ObjectSpamfilterProfile-MsnHotmail"); ok {
					if err = d.Set("msn_hotmail", vv); err != nil {
						return fmt.Errorf("Error reading msn_hotmail: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading msn_hotmail: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectSpamfilterProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSpamfilterProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("options", flattenObjectSpamfilterProfileOptions(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "ObjectSpamfilterProfile-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("pop3", flattenObjectSpamfilterProfilePop3(o["pop3"], d, "pop3")); err != nil {
			if vv, ok := fortiAPIPatch(o["pop3"], "ObjectSpamfilterProfile-Pop3"); ok {
				if err = d.Set("pop3", vv); err != nil {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pop3: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pop3"); ok {
			if err = d.Set("pop3", flattenObjectSpamfilterProfilePop3(o["pop3"], d, "pop3")); err != nil {
				if vv, ok := fortiAPIPatch(o["pop3"], "ObjectSpamfilterProfile-Pop3"); ok {
					if err = d.Set("pop3", vv); err != nil {
						return fmt.Errorf("Error reading pop3: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			}
		}
	}

	if err = d.Set("replacemsg_group", flattenObjectSpamfilterProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "ObjectSpamfilterProfile-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("smtp", flattenObjectSpamfilterProfileSmtp(o["smtp"], d, "smtp")); err != nil {
			if vv, ok := fortiAPIPatch(o["smtp"], "ObjectSpamfilterProfile-Smtp"); ok {
				if err = d.Set("smtp", vv); err != nil {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading smtp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("smtp"); ok {
			if err = d.Set("smtp", flattenObjectSpamfilterProfileSmtp(o["smtp"], d, "smtp")); err != nil {
				if vv, ok := fortiAPIPatch(o["smtp"], "ObjectSpamfilterProfile-Smtp"); ok {
					if err = d.Set("smtp", vv); err != nil {
						return fmt.Errorf("Error reading smtp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			}
		}
	}

	if err = d.Set("spam_bwl_table", flattenObjectSpamfilterProfileSpamBwlTable(o["spam-bwl-table"], d, "spam_bwl_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-bwl-table"], "ObjectSpamfilterProfile-SpamBwlTable"); ok {
			if err = d.Set("spam_bwl_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_bwl_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_bwl_table: %v", err)
		}
	}

	if err = d.Set("spam_bword_table", flattenObjectSpamfilterProfileSpamBwordTable(o["spam-bword-table"], d, "spam_bword_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-bword-table"], "ObjectSpamfilterProfile-SpamBwordTable"); ok {
			if err = d.Set("spam_bword_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_bword_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_bword_table: %v", err)
		}
	}

	if err = d.Set("spam_bword_threshold", flattenObjectSpamfilterProfileSpamBwordThreshold(o["spam-bword-threshold"], d, "spam_bword_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-bword-threshold"], "ObjectSpamfilterProfile-SpamBwordThreshold"); ok {
			if err = d.Set("spam_bword_threshold", vv); err != nil {
				return fmt.Errorf("Error reading spam_bword_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_bword_threshold: %v", err)
		}
	}

	if err = d.Set("spam_filtering", flattenObjectSpamfilterProfileSpamFiltering(o["spam-filtering"], d, "spam_filtering")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-filtering"], "ObjectSpamfilterProfile-SpamFiltering"); ok {
			if err = d.Set("spam_filtering", vv); err != nil {
				return fmt.Errorf("Error reading spam_filtering: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_filtering: %v", err)
		}
	}

	if err = d.Set("spam_iptrust_table", flattenObjectSpamfilterProfileSpamIptrustTable(o["spam-iptrust-table"], d, "spam_iptrust_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-iptrust-table"], "ObjectSpamfilterProfile-SpamIptrustTable"); ok {
			if err = d.Set("spam_iptrust_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_iptrust_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_iptrust_table: %v", err)
		}
	}

	if err = d.Set("spam_log", flattenObjectSpamfilterProfileSpamLog(o["spam-log"], d, "spam_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-log"], "ObjectSpamfilterProfile-SpamLog"); ok {
			if err = d.Set("spam_log", vv); err != nil {
				return fmt.Errorf("Error reading spam_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_log: %v", err)
		}
	}

	if err = d.Set("spam_log_fortiguard_response", flattenObjectSpamfilterProfileSpamLogFortiguardResponse(o["spam-log-fortiguard-response"], d, "spam_log_fortiguard_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-log-fortiguard-response"], "ObjectSpamfilterProfile-SpamLogFortiguardResponse"); ok {
			if err = d.Set("spam_log_fortiguard_response", vv); err != nil {
				return fmt.Errorf("Error reading spam_log_fortiguard_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_log_fortiguard_response: %v", err)
		}
	}

	if err = d.Set("spam_mheader_table", flattenObjectSpamfilterProfileSpamMheaderTable(o["spam-mheader-table"], d, "spam_mheader_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-mheader-table"], "ObjectSpamfilterProfile-SpamMheaderTable"); ok {
			if err = d.Set("spam_mheader_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_mheader_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_mheader_table: %v", err)
		}
	}

	if err = d.Set("spam_rbl_table", flattenObjectSpamfilterProfileSpamRblTable(o["spam-rbl-table"], d, "spam_rbl_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-rbl-table"], "ObjectSpamfilterProfile-SpamRblTable"); ok {
			if err = d.Set("spam_rbl_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_rbl_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_rbl_table: %v", err)
		}
	}

	return nil
}

func flattenObjectSpamfilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSpamfilterProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileExternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileFlowBased(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileGmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectSpamfilterProfileGmailLog(d, i["log"], pre_append)
	}

	return result, nil
}

func expandObjectSpamfilterProfileGmailLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileImap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectSpamfilterProfileImapAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectSpamfilterProfileImapLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "tag_msg"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag-msg"], _ = expandObjectSpamfilterProfileImapTagMsg(d, i["tag_msg"], pre_append)
	}
	pre_append = pre + ".0." + "tag_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag-type"], _ = expandObjectSpamfilterProfileImapTagType(d, i["tag_type"], pre_append)
	}

	return result, nil
}

func expandObjectSpamfilterProfileImapAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileImapLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileImapTagMsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileImapTagType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSpamfilterProfileMapi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectSpamfilterProfileMapiAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectSpamfilterProfileMapiLog(d, i["log"], pre_append)
	}

	return result, nil
}

func expandObjectSpamfilterProfileMapiAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileMapiLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileMsnHotmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectSpamfilterProfileMsnHotmailLog(d, i["log"], pre_append)
	}

	return result, nil
}

func expandObjectSpamfilterProfileMsnHotmailLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSpamfilterProfilePop3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectSpamfilterProfilePop3Action(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectSpamfilterProfilePop3Log(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "tag_msg"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag-msg"], _ = expandObjectSpamfilterProfilePop3TagMsg(d, i["tag_msg"], pre_append)
	}
	pre_append = pre + ".0." + "tag_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag-type"], _ = expandObjectSpamfilterProfilePop3TagType(d, i["tag_type"], pre_append)
	}

	return result, nil
}

func expandObjectSpamfilterProfilePop3Action(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfilePop3Log(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfilePop3TagMsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfilePop3TagType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSpamfilterProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileSmtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectSpamfilterProfileSmtpAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "hdrip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["hdrip"], _ = expandObjectSpamfilterProfileSmtpHdrip(d, i["hdrip"], pre_append)
	}
	pre_append = pre + ".0." + "local_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["local-override"], _ = expandObjectSpamfilterProfileSmtpLocalOverride(d, i["local_override"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandObjectSpamfilterProfileSmtpLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "tag_msg"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag-msg"], _ = expandObjectSpamfilterProfileSmtpTagMsg(d, i["tag_msg"], pre_append)
	}
	pre_append = pre + ".0." + "tag_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag-type"], _ = expandObjectSpamfilterProfileSmtpTagType(d, i["tag_type"], pre_append)
	}

	return result, nil
}

func expandObjectSpamfilterProfileSmtpAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileSmtpHdrip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileSmtpLocalOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileSmtpLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileSmtpTagMsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileSmtpTagType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSpamfilterProfileSpamBwlTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileSpamBwordTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileSpamBwordThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileSpamFiltering(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileSpamIptrustTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileSpamLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileSpamLogFortiguardResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileSpamMheaderTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterProfileSpamRblTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSpamfilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectSpamfilterProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("external"); ok || d.HasChange("external") {
		t, err := expandObjectSpamfilterProfileExternal(d, v, "external")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external"] = t
		}
	}

	if v, ok := d.GetOk("flow_based"); ok || d.HasChange("flow_based") {
		t, err := expandObjectSpamfilterProfileFlowBased(d, v, "flow_based")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flow-based"] = t
		}
	}

	if v, ok := d.GetOk("gmail"); ok || d.HasChange("gmail") {
		t, err := expandObjectSpamfilterProfileGmail(d, v, "gmail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gmail"] = t
		}
	}

	if v, ok := d.GetOk("imap"); ok || d.HasChange("imap") {
		t, err := expandObjectSpamfilterProfileImap(d, v, "imap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imap"] = t
		}
	}

	if v, ok := d.GetOk("mapi"); ok || d.HasChange("mapi") {
		t, err := expandObjectSpamfilterProfileMapi(d, v, "mapi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi"] = t
		}
	}

	if v, ok := d.GetOk("msn_hotmail"); ok || d.HasChange("msn_hotmail") {
		t, err := expandObjectSpamfilterProfileMsnHotmail(d, v, "msn_hotmail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msn-hotmail"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSpamfilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandObjectSpamfilterProfileOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("pop3"); ok || d.HasChange("pop3") {
		t, err := expandObjectSpamfilterProfilePop3(d, v, "pop3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandObjectSpamfilterProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("smtp"); ok || d.HasChange("smtp") {
		t, err := expandObjectSpamfilterProfileSmtp(d, v, "smtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp"] = t
		}
	}

	if v, ok := d.GetOk("spam_bwl_table"); ok || d.HasChange("spam_bwl_table") {
		t, err := expandObjectSpamfilterProfileSpamBwlTable(d, v, "spam_bwl_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bwl-table"] = t
		}
	}

	if v, ok := d.GetOk("spam_bword_table"); ok || d.HasChange("spam_bword_table") {
		t, err := expandObjectSpamfilterProfileSpamBwordTable(d, v, "spam_bword_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bword-table"] = t
		}
	}

	if v, ok := d.GetOk("spam_bword_threshold"); ok || d.HasChange("spam_bword_threshold") {
		t, err := expandObjectSpamfilterProfileSpamBwordThreshold(d, v, "spam_bword_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bword-threshold"] = t
		}
	}

	if v, ok := d.GetOk("spam_filtering"); ok || d.HasChange("spam_filtering") {
		t, err := expandObjectSpamfilterProfileSpamFiltering(d, v, "spam_filtering")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-filtering"] = t
		}
	}

	if v, ok := d.GetOk("spam_iptrust_table"); ok || d.HasChange("spam_iptrust_table") {
		t, err := expandObjectSpamfilterProfileSpamIptrustTable(d, v, "spam_iptrust_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-iptrust-table"] = t
		}
	}

	if v, ok := d.GetOk("spam_log"); ok || d.HasChange("spam_log") {
		t, err := expandObjectSpamfilterProfileSpamLog(d, v, "spam_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-log"] = t
		}
	}

	if v, ok := d.GetOk("spam_log_fortiguard_response"); ok || d.HasChange("spam_log_fortiguard_response") {
		t, err := expandObjectSpamfilterProfileSpamLogFortiguardResponse(d, v, "spam_log_fortiguard_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-log-fortiguard-response"] = t
		}
	}

	if v, ok := d.GetOk("spam_mheader_table"); ok || d.HasChange("spam_mheader_table") {
		t, err := expandObjectSpamfilterProfileSpamMheaderTable(d, v, "spam_mheader_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-mheader-table"] = t
		}
	}

	if v, ok := d.GetOk("spam_rbl_table"); ok || d.HasChange("spam_rbl_table") {
		t, err := expandObjectSpamfilterProfileSpamRblTable(d, v, "spam_rbl_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-rbl-table"] = t
		}
	}

	return &obj, nil
}
