// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure replacement message groups.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSystemReplacemsgGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemReplacemsgGroupCreate,
		Read:   resourceObjectSystemReplacemsgGroupRead,
		Update: resourceObjectSystemReplacemsgGroupUpdate,
		Delete: resourceObjectSystemReplacemsgGroupDelete,

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
			"admin": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"alertmail": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"auth": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"automation": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_message": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"device_detection_portal": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ec": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"fortiguard_wf": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ftp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"group_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"icap": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"mail": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"mm1": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"add_smil": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"charset": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"class": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"from": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"from_sender": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"image": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rsp_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rsp_text": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sender_visibility": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"smil_part": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"subject": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"mm3": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"add_html": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"charset": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"from": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"from_sender": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"html_part": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"image": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"subject": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"mm4": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"add_smil": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"charset": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"class": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"domain": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"from": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"from_sender": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"image": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rsp_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"smil_part": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"subject": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"mm7": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"add_smil": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"allow_content_adaptation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"charset": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"class": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"from": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"from_sender": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"image": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rsp_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"smil_part": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"subject": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"mms": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"charset": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"image": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"nac_quar": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
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
			"nntp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"spam": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"sslvpn": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"traffic_quota": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"utm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"webproxy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectSystemReplacemsgGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemReplacemsgGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemReplacemsgGroup resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemReplacemsgGroup(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemReplacemsgGroup resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemReplacemsgGroupRead(d, m)
}

func resourceObjectSystemReplacemsgGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemReplacemsgGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemReplacemsgGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemReplacemsgGroup(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemReplacemsgGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemReplacemsgGroupRead(d, m)
}

func resourceObjectSystemReplacemsgGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemReplacemsgGroup(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemReplacemsgGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemReplacemsgGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemReplacemsgGroup(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemReplacemsgGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemReplacemsgGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemReplacemsgGroup resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemReplacemsgGroupAdmin(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupAdminBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Admin-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupAdminFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Admin-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupAdminHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Admin-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupAdminMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Admin-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupAdminBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupAdminFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupAdminHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupAdminMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupAlertmail(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupAlertmailBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Alertmail-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupAlertmailFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Alertmail-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupAlertmailHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Alertmail-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupAlertmailMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Alertmail-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupAlertmailBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupAlertmailFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupAlertmailHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupAlertmailMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupAuth(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupAuthBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Auth-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupAuthFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Auth-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupAuthHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Auth-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupAuthMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Auth-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupAuthBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupAuthFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupAuthHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupAuthMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupAutomation(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupAutomationBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Automation-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupAutomationFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Automation-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupAutomationHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Automation-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupAutomationMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Automation-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupAutomationBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupAutomationFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupAutomationHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupAutomationMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupCustomMessage(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupCustomMessageBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-CustomMessage-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupCustomMessageFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-CustomMessage-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupCustomMessageHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-CustomMessage-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupCustomMessageMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-CustomMessage-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupCustomMessageBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupCustomMessageFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupCustomMessageHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupCustomMessageMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupDeviceDetectionPortal(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupDeviceDetectionPortalBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-DeviceDetectionPortal-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupDeviceDetectionPortalFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-DeviceDetectionPortal-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupDeviceDetectionPortalHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-DeviceDetectionPortal-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupDeviceDetectionPortalMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-DeviceDetectionPortal-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupDeviceDetectionPortalBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupDeviceDetectionPortalFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupDeviceDetectionPortalHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupDeviceDetectionPortalMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupEc(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupEcBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Ec-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupEcFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Ec-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupEcHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Ec-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupEcMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Ec-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupEcBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupEcFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupEcHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupEcMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupFortiguardWf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupFortiguardWfBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-FortiguardWf-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupFortiguardWfFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-FortiguardWf-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupFortiguardWfHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-FortiguardWf-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupFortiguardWfMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-FortiguardWf-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupFortiguardWfBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupFortiguardWfFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupFortiguardWfHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupFortiguardWfMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupFtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupFtpBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Ftp-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupFtpFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Ftp-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupFtpHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Ftp-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupFtpMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Ftp-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupFtpBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupFtpFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupFtpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupFtpMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupGroupType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupHttp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupHttpBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Http-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupHttpFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Http-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupHttpHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Http-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupHttpMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Http-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupHttpBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupHttpFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupHttpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupHttpMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupIcap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupIcapBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Icap-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupIcapFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Icap-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupIcapHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Icap-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupIcapMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Icap-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupIcapBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupIcapFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupIcapHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupIcapMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMail(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupMailBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mail-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupMailFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mail-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupMailHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mail-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupMailMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mail-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupMailBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMailFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMailHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMailMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_smil"
		if _, ok := i["add-smil"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1AddSmil(i["add-smil"], d, pre_append)
			tmp["add_smil"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-AddSmil")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "charset"
		if _, ok := i["charset"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1Charset(i["charset"], d, pre_append)
			tmp["charset"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-Charset")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class"
		if _, ok := i["class"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1Class(i["class"], d, pre_append)
			tmp["class"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-Class")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1Format(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from"
		if _, ok := i["from"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1From(i["from"], d, pre_append)
			tmp["from"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-From")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from_sender"
		if _, ok := i["from-sender"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1FromSender(i["from-sender"], d, pre_append)
			tmp["from_sender"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-FromSender")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1Header(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "image"
		if _, ok := i["image"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1Image(i["image"], d, pre_append)
			tmp["image"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-Image")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "message"
		if _, ok := i["message"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1Message(i["message"], d, pre_append)
			tmp["message"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-Message")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1MsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-MsgType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1Priority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsp_status"
		if _, ok := i["rsp-status"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1RspStatus(i["rsp-status"], d, pre_append)
			tmp["rsp_status"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-RspStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsp_text"
		if _, ok := i["rsp-text"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1RspText(i["rsp-text"], d, pre_append)
			tmp["rsp_text"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-RspText")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sender_visibility"
		if _, ok := i["sender-visibility"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1SenderVisibility(i["sender-visibility"], d, pre_append)
			tmp["sender_visibility"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-SenderVisibility")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smil_part"
		if _, ok := i["smil-part"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1SmilPart(i["smil-part"], d, pre_append)
			tmp["smil_part"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-SmilPart")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subject"
		if _, ok := i["subject"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm1Subject(i["subject"], d, pre_append)
			tmp["subject"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm1-Subject")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupMm1AddSmil(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1Charset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1Class(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1Format(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1From(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1FromSender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1Header(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1Image(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1Message(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1MsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1Priority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1RspStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1RspText(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1SenderVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1SmilPart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm1Subject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm3(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_html"
		if _, ok := i["add-html"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm3AddHtml(i["add-html"], d, pre_append)
			tmp["add_html"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm3-AddHtml")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "charset"
		if _, ok := i["charset"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm3Charset(i["charset"], d, pre_append)
			tmp["charset"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm3-Charset")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm3Format(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm3-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from"
		if _, ok := i["from"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm3From(i["from"], d, pre_append)
			tmp["from"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm3-From")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from_sender"
		if _, ok := i["from-sender"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm3FromSender(i["from-sender"], d, pre_append)
			tmp["from_sender"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm3-FromSender")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm3Header(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm3-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "html_part"
		if _, ok := i["html-part"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm3HtmlPart(i["html-part"], d, pre_append)
			tmp["html_part"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm3-HtmlPart")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "image"
		if _, ok := i["image"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm3Image(i["image"], d, pre_append)
			tmp["image"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm3-Image")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "message"
		if _, ok := i["message"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm3Message(i["message"], d, pre_append)
			tmp["message"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm3-Message")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm3MsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm3-MsgType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm3Priority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm3-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subject"
		if _, ok := i["subject"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm3Subject(i["subject"], d, pre_append)
			tmp["subject"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm3-Subject")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupMm3AddHtml(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm3Charset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm3Format(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm3From(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm3FromSender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm3Header(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm3HtmlPart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm3Image(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm3Message(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm3MsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm3Priority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm3Subject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm4(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_smil"
		if _, ok := i["add-smil"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm4AddSmil(i["add-smil"], d, pre_append)
			tmp["add_smil"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm4-AddSmil")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "charset"
		if _, ok := i["charset"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm4Charset(i["charset"], d, pre_append)
			tmp["charset"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm4-Charset")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class"
		if _, ok := i["class"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm4Class(i["class"], d, pre_append)
			tmp["class"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm4-Class")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm4Domain(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm4-Domain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm4Format(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm4-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from"
		if _, ok := i["from"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm4From(i["from"], d, pre_append)
			tmp["from"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm4-From")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from_sender"
		if _, ok := i["from-sender"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm4FromSender(i["from-sender"], d, pre_append)
			tmp["from_sender"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm4-FromSender")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm4Header(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm4-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "image"
		if _, ok := i["image"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm4Image(i["image"], d, pre_append)
			tmp["image"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm4-Image")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "message"
		if _, ok := i["message"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm4Message(i["message"], d, pre_append)
			tmp["message"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm4-Message")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm4MsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm4-MsgType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm4Priority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm4-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsp_status"
		if _, ok := i["rsp-status"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm4RspStatus(i["rsp-status"], d, pre_append)
			tmp["rsp_status"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm4-RspStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smil_part"
		if _, ok := i["smil-part"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm4SmilPart(i["smil-part"], d, pre_append)
			tmp["smil_part"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm4-SmilPart")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subject"
		if _, ok := i["subject"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm4Subject(i["subject"], d, pre_append)
			tmp["subject"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm4-Subject")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupMm4AddSmil(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm4Charset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm4Class(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm4Domain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm4Format(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm4From(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm4FromSender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm4Header(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm4Image(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm4Message(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm4MsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm4Priority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm4RspStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm4SmilPart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm4Subject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_smil"
		if _, ok := i["add-smil"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7AddSmil(i["add-smil"], d, pre_append)
			tmp["add_smil"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-AddSmil")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := i["addr-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7AddrType(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allow_content_adaptation"
		if _, ok := i["allow-content-adaptation"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7AllowContentAdaptation(i["allow-content-adaptation"], d, pre_append)
			tmp["allow_content_adaptation"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-AllowContentAdaptation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "charset"
		if _, ok := i["charset"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7Charset(i["charset"], d, pre_append)
			tmp["charset"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-Charset")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class"
		if _, ok := i["class"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7Class(i["class"], d, pre_append)
			tmp["class"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-Class")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7Format(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from"
		if _, ok := i["from"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7From(i["from"], d, pre_append)
			tmp["from"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-From")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from_sender"
		if _, ok := i["from-sender"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7FromSender(i["from-sender"], d, pre_append)
			tmp["from_sender"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-FromSender")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7Header(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "image"
		if _, ok := i["image"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7Image(i["image"], d, pre_append)
			tmp["image"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-Image")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "message"
		if _, ok := i["message"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7Message(i["message"], d, pre_append)
			tmp["message"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-Message")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7MsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-MsgType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7Priority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsp_status"
		if _, ok := i["rsp-status"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7RspStatus(i["rsp-status"], d, pre_append)
			tmp["rsp_status"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-RspStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smil_part"
		if _, ok := i["smil-part"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7SmilPart(i["smil-part"], d, pre_append)
			tmp["smil_part"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-SmilPart")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subject"
		if _, ok := i["subject"]; ok {
			v := flattenObjectSystemReplacemsgGroupMm7Subject(i["subject"], d, pre_append)
			tmp["subject"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mm7-Subject")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupMm7AddSmil(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7AddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7AllowContentAdaptation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7Charset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7Class(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7Format(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7From(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7FromSender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7Header(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7Image(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7Message(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7MsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7Priority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7RspStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7SmilPart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMm7Subject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMms(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupMmsBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mms-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "charset"
		if _, ok := i["charset"]; ok {
			v := flattenObjectSystemReplacemsgGroupMmsCharset(i["charset"], d, pre_append)
			tmp["charset"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mms-Charset")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupMmsFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mms-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupMmsHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mms-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "image"
		if _, ok := i["image"]; ok {
			v := flattenObjectSystemReplacemsgGroupMmsImage(i["image"], d, pre_append)
			tmp["image"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mms-Image")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupMmsMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Mms-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupMmsBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMmsCharset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMmsFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMmsHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMmsImage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupMmsMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupNacQuar(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupNacQuarBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-NacQuar-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupNacQuarFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-NacQuar-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupNacQuarHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-NacQuar-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupNacQuarMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-NacQuar-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupNacQuarBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupNacQuarFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupNacQuarHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupNacQuarMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupNntp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupNntpBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Nntp-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupNntpFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Nntp-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupNntpHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Nntp-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupNntpMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Nntp-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupNntpBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupNntpFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupNntpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupNntpMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupSpam(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupSpamBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Spam-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupSpamFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Spam-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupSpamHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Spam-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupSpamMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Spam-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupSpamBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupSpamFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupSpamHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupSpamMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupSslvpn(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupSslvpnBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Sslvpn-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupSslvpnFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Sslvpn-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupSslvpnHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Sslvpn-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupSslvpnMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Sslvpn-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupSslvpnBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupSslvpnFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupSslvpnHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupSslvpnMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupTrafficQuota(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupTrafficQuotaBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-TrafficQuota-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupTrafficQuotaFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-TrafficQuota-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupTrafficQuotaHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-TrafficQuota-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupTrafficQuotaMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-TrafficQuota-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupTrafficQuotaBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupTrafficQuotaFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupTrafficQuotaHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupTrafficQuotaMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupUtm(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupUtmBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Utm-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupUtmFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Utm-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupUtmHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Utm-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupUtmMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Utm-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupUtmBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupUtmFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupUtmHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupUtmMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupWebproxy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenObjectSystemReplacemsgGroupWebproxyBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Webproxy-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenObjectSystemReplacemsgGroupWebproxyFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Webproxy-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectSystemReplacemsgGroupWebproxyHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Webproxy-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenObjectSystemReplacemsgGroupWebproxyMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "ObjectSystemReplacemsgGroup-Webproxy-MsgType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemReplacemsgGroupWebproxyBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupWebproxyFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupWebproxyHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupWebproxyMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemReplacemsgGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("admin", flattenObjectSystemReplacemsgGroupAdmin(o["admin"], d, "admin")); err != nil {
			if vv, ok := fortiAPIPatch(o["admin"], "ObjectSystemReplacemsgGroup-Admin"); ok {
				if err = d.Set("admin", vv); err != nil {
					return fmt.Errorf("Error reading admin: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading admin: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("admin"); ok {
			if err = d.Set("admin", flattenObjectSystemReplacemsgGroupAdmin(o["admin"], d, "admin")); err != nil {
				if vv, ok := fortiAPIPatch(o["admin"], "ObjectSystemReplacemsgGroup-Admin"); ok {
					if err = d.Set("admin", vv); err != nil {
						return fmt.Errorf("Error reading admin: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading admin: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("alertmail", flattenObjectSystemReplacemsgGroupAlertmail(o["alertmail"], d, "alertmail")); err != nil {
			if vv, ok := fortiAPIPatch(o["alertmail"], "ObjectSystemReplacemsgGroup-Alertmail"); ok {
				if err = d.Set("alertmail", vv); err != nil {
					return fmt.Errorf("Error reading alertmail: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading alertmail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("alertmail"); ok {
			if err = d.Set("alertmail", flattenObjectSystemReplacemsgGroupAlertmail(o["alertmail"], d, "alertmail")); err != nil {
				if vv, ok := fortiAPIPatch(o["alertmail"], "ObjectSystemReplacemsgGroup-Alertmail"); ok {
					if err = d.Set("alertmail", vv); err != nil {
						return fmt.Errorf("Error reading alertmail: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading alertmail: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("auth", flattenObjectSystemReplacemsgGroupAuth(o["auth"], d, "auth")); err != nil {
			if vv, ok := fortiAPIPatch(o["auth"], "ObjectSystemReplacemsgGroup-Auth"); ok {
				if err = d.Set("auth", vv); err != nil {
					return fmt.Errorf("Error reading auth: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading auth: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("auth"); ok {
			if err = d.Set("auth", flattenObjectSystemReplacemsgGroupAuth(o["auth"], d, "auth")); err != nil {
				if vv, ok := fortiAPIPatch(o["auth"], "ObjectSystemReplacemsgGroup-Auth"); ok {
					if err = d.Set("auth", vv); err != nil {
						return fmt.Errorf("Error reading auth: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading auth: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("automation", flattenObjectSystemReplacemsgGroupAutomation(o["automation"], d, "automation")); err != nil {
			if vv, ok := fortiAPIPatch(o["automation"], "ObjectSystemReplacemsgGroup-Automation"); ok {
				if err = d.Set("automation", vv); err != nil {
					return fmt.Errorf("Error reading automation: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading automation: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("automation"); ok {
			if err = d.Set("automation", flattenObjectSystemReplacemsgGroupAutomation(o["automation"], d, "automation")); err != nil {
				if vv, ok := fortiAPIPatch(o["automation"], "ObjectSystemReplacemsgGroup-Automation"); ok {
					if err = d.Set("automation", vv); err != nil {
						return fmt.Errorf("Error reading automation: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading automation: %v", err)
				}
			}
		}
	}

	if err = d.Set("comment", flattenObjectSystemReplacemsgGroupComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectSystemReplacemsgGroup-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_message", flattenObjectSystemReplacemsgGroupCustomMessage(o["custom-message"], d, "custom_message")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-message"], "ObjectSystemReplacemsgGroup-CustomMessage"); ok {
				if err = d.Set("custom_message", vv); err != nil {
					return fmt.Errorf("Error reading custom_message: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_message: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_message"); ok {
			if err = d.Set("custom_message", flattenObjectSystemReplacemsgGroupCustomMessage(o["custom-message"], d, "custom_message")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-message"], "ObjectSystemReplacemsgGroup-CustomMessage"); ok {
					if err = d.Set("custom_message", vv); err != nil {
						return fmt.Errorf("Error reading custom_message: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_message: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("device_detection_portal", flattenObjectSystemReplacemsgGroupDeviceDetectionPortal(o["device-detection-portal"], d, "device_detection_portal")); err != nil {
			if vv, ok := fortiAPIPatch(o["device-detection-portal"], "ObjectSystemReplacemsgGroup-DeviceDetectionPortal"); ok {
				if err = d.Set("device_detection_portal", vv); err != nil {
					return fmt.Errorf("Error reading device_detection_portal: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading device_detection_portal: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("device_detection_portal"); ok {
			if err = d.Set("device_detection_portal", flattenObjectSystemReplacemsgGroupDeviceDetectionPortal(o["device-detection-portal"], d, "device_detection_portal")); err != nil {
				if vv, ok := fortiAPIPatch(o["device-detection-portal"], "ObjectSystemReplacemsgGroup-DeviceDetectionPortal"); ok {
					if err = d.Set("device_detection_portal", vv); err != nil {
						return fmt.Errorf("Error reading device_detection_portal: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading device_detection_portal: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ec", flattenObjectSystemReplacemsgGroupEc(o["ec"], d, "ec")); err != nil {
			if vv, ok := fortiAPIPatch(o["ec"], "ObjectSystemReplacemsgGroup-Ec"); ok {
				if err = d.Set("ec", vv); err != nil {
					return fmt.Errorf("Error reading ec: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ec: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ec"); ok {
			if err = d.Set("ec", flattenObjectSystemReplacemsgGroupEc(o["ec"], d, "ec")); err != nil {
				if vv, ok := fortiAPIPatch(o["ec"], "ObjectSystemReplacemsgGroup-Ec"); ok {
					if err = d.Set("ec", vv); err != nil {
						return fmt.Errorf("Error reading ec: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ec: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("fortiguard_wf", flattenObjectSystemReplacemsgGroupFortiguardWf(o["fortiguard-wf"], d, "fortiguard_wf")); err != nil {
			if vv, ok := fortiAPIPatch(o["fortiguard-wf"], "ObjectSystemReplacemsgGroup-FortiguardWf"); ok {
				if err = d.Set("fortiguard_wf", vv); err != nil {
					return fmt.Errorf("Error reading fortiguard_wf: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading fortiguard_wf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fortiguard_wf"); ok {
			if err = d.Set("fortiguard_wf", flattenObjectSystemReplacemsgGroupFortiguardWf(o["fortiguard-wf"], d, "fortiguard_wf")); err != nil {
				if vv, ok := fortiAPIPatch(o["fortiguard-wf"], "ObjectSystemReplacemsgGroup-FortiguardWf"); ok {
					if err = d.Set("fortiguard_wf", vv); err != nil {
						return fmt.Errorf("Error reading fortiguard_wf: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading fortiguard_wf: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ftp", flattenObjectSystemReplacemsgGroupFtp(o["ftp"], d, "ftp")); err != nil {
			if vv, ok := fortiAPIPatch(o["ftp"], "ObjectSystemReplacemsgGroup-Ftp"); ok {
				if err = d.Set("ftp", vv); err != nil {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ftp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftp"); ok {
			if err = d.Set("ftp", flattenObjectSystemReplacemsgGroupFtp(o["ftp"], d, "ftp")); err != nil {
				if vv, ok := fortiAPIPatch(o["ftp"], "ObjectSystemReplacemsgGroup-Ftp"); ok {
					if err = d.Set("ftp", vv); err != nil {
						return fmt.Errorf("Error reading ftp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			}
		}
	}

	if err = d.Set("group_type", flattenObjectSystemReplacemsgGroupGroupType(o["group-type"], d, "group_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-type"], "ObjectSystemReplacemsgGroup-GroupType"); ok {
			if err = d.Set("group_type", vv); err != nil {
				return fmt.Errorf("Error reading group_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("http", flattenObjectSystemReplacemsgGroupHttp(o["http"], d, "http")); err != nil {
			if vv, ok := fortiAPIPatch(o["http"], "ObjectSystemReplacemsgGroup-Http"); ok {
				if err = d.Set("http", vv); err != nil {
					return fmt.Errorf("Error reading http: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading http: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("http"); ok {
			if err = d.Set("http", flattenObjectSystemReplacemsgGroupHttp(o["http"], d, "http")); err != nil {
				if vv, ok := fortiAPIPatch(o["http"], "ObjectSystemReplacemsgGroup-Http"); ok {
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
		if err = d.Set("icap", flattenObjectSystemReplacemsgGroupIcap(o["icap"], d, "icap")); err != nil {
			if vv, ok := fortiAPIPatch(o["icap"], "ObjectSystemReplacemsgGroup-Icap"); ok {
				if err = d.Set("icap", vv); err != nil {
					return fmt.Errorf("Error reading icap: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading icap: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("icap"); ok {
			if err = d.Set("icap", flattenObjectSystemReplacemsgGroupIcap(o["icap"], d, "icap")); err != nil {
				if vv, ok := fortiAPIPatch(o["icap"], "ObjectSystemReplacemsgGroup-Icap"); ok {
					if err = d.Set("icap", vv); err != nil {
						return fmt.Errorf("Error reading icap: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading icap: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mail", flattenObjectSystemReplacemsgGroupMail(o["mail"], d, "mail")); err != nil {
			if vv, ok := fortiAPIPatch(o["mail"], "ObjectSystemReplacemsgGroup-Mail"); ok {
				if err = d.Set("mail", vv); err != nil {
					return fmt.Errorf("Error reading mail: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mail"); ok {
			if err = d.Set("mail", flattenObjectSystemReplacemsgGroupMail(o["mail"], d, "mail")); err != nil {
				if vv, ok := fortiAPIPatch(o["mail"], "ObjectSystemReplacemsgGroup-Mail"); ok {
					if err = d.Set("mail", vv); err != nil {
						return fmt.Errorf("Error reading mail: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mail: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mm1", flattenObjectSystemReplacemsgGroupMm1(o["mm1"], d, "mm1")); err != nil {
			if vv, ok := fortiAPIPatch(o["mm1"], "ObjectSystemReplacemsgGroup-Mm1"); ok {
				if err = d.Set("mm1", vv); err != nil {
					return fmt.Errorf("Error reading mm1: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mm1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mm1"); ok {
			if err = d.Set("mm1", flattenObjectSystemReplacemsgGroupMm1(o["mm1"], d, "mm1")); err != nil {
				if vv, ok := fortiAPIPatch(o["mm1"], "ObjectSystemReplacemsgGroup-Mm1"); ok {
					if err = d.Set("mm1", vv); err != nil {
						return fmt.Errorf("Error reading mm1: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mm1: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mm3", flattenObjectSystemReplacemsgGroupMm3(o["mm3"], d, "mm3")); err != nil {
			if vv, ok := fortiAPIPatch(o["mm3"], "ObjectSystemReplacemsgGroup-Mm3"); ok {
				if err = d.Set("mm3", vv); err != nil {
					return fmt.Errorf("Error reading mm3: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mm3: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mm3"); ok {
			if err = d.Set("mm3", flattenObjectSystemReplacemsgGroupMm3(o["mm3"], d, "mm3")); err != nil {
				if vv, ok := fortiAPIPatch(o["mm3"], "ObjectSystemReplacemsgGroup-Mm3"); ok {
					if err = d.Set("mm3", vv); err != nil {
						return fmt.Errorf("Error reading mm3: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mm3: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mm4", flattenObjectSystemReplacemsgGroupMm4(o["mm4"], d, "mm4")); err != nil {
			if vv, ok := fortiAPIPatch(o["mm4"], "ObjectSystemReplacemsgGroup-Mm4"); ok {
				if err = d.Set("mm4", vv); err != nil {
					return fmt.Errorf("Error reading mm4: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mm4: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mm4"); ok {
			if err = d.Set("mm4", flattenObjectSystemReplacemsgGroupMm4(o["mm4"], d, "mm4")); err != nil {
				if vv, ok := fortiAPIPatch(o["mm4"], "ObjectSystemReplacemsgGroup-Mm4"); ok {
					if err = d.Set("mm4", vv); err != nil {
						return fmt.Errorf("Error reading mm4: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mm4: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mm7", flattenObjectSystemReplacemsgGroupMm7(o["mm7"], d, "mm7")); err != nil {
			if vv, ok := fortiAPIPatch(o["mm7"], "ObjectSystemReplacemsgGroup-Mm7"); ok {
				if err = d.Set("mm7", vv); err != nil {
					return fmt.Errorf("Error reading mm7: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mm7: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mm7"); ok {
			if err = d.Set("mm7", flattenObjectSystemReplacemsgGroupMm7(o["mm7"], d, "mm7")); err != nil {
				if vv, ok := fortiAPIPatch(o["mm7"], "ObjectSystemReplacemsgGroup-Mm7"); ok {
					if err = d.Set("mm7", vv); err != nil {
						return fmt.Errorf("Error reading mm7: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mm7: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mms", flattenObjectSystemReplacemsgGroupMms(o["mms"], d, "mms")); err != nil {
			if vv, ok := fortiAPIPatch(o["mms"], "ObjectSystemReplacemsgGroup-Mms"); ok {
				if err = d.Set("mms", vv); err != nil {
					return fmt.Errorf("Error reading mms: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mms: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mms"); ok {
			if err = d.Set("mms", flattenObjectSystemReplacemsgGroupMms(o["mms"], d, "mms")); err != nil {
				if vv, ok := fortiAPIPatch(o["mms"], "ObjectSystemReplacemsgGroup-Mms"); ok {
					if err = d.Set("mms", vv); err != nil {
						return fmt.Errorf("Error reading mms: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mms: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("nac_quar", flattenObjectSystemReplacemsgGroupNacQuar(o["nac-quar"], d, "nac_quar")); err != nil {
			if vv, ok := fortiAPIPatch(o["nac-quar"], "ObjectSystemReplacemsgGroup-NacQuar"); ok {
				if err = d.Set("nac_quar", vv); err != nil {
					return fmt.Errorf("Error reading nac_quar: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nac_quar: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nac_quar"); ok {
			if err = d.Set("nac_quar", flattenObjectSystemReplacemsgGroupNacQuar(o["nac-quar"], d, "nac_quar")); err != nil {
				if vv, ok := fortiAPIPatch(o["nac-quar"], "ObjectSystemReplacemsgGroup-NacQuar"); ok {
					if err = d.Set("nac_quar", vv); err != nil {
						return fmt.Errorf("Error reading nac_quar: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nac_quar: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectSystemReplacemsgGroupName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemReplacemsgGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("nntp", flattenObjectSystemReplacemsgGroupNntp(o["nntp"], d, "nntp")); err != nil {
			if vv, ok := fortiAPIPatch(o["nntp"], "ObjectSystemReplacemsgGroup-Nntp"); ok {
				if err = d.Set("nntp", vv); err != nil {
					return fmt.Errorf("Error reading nntp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nntp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nntp"); ok {
			if err = d.Set("nntp", flattenObjectSystemReplacemsgGroupNntp(o["nntp"], d, "nntp")); err != nil {
				if vv, ok := fortiAPIPatch(o["nntp"], "ObjectSystemReplacemsgGroup-Nntp"); ok {
					if err = d.Set("nntp", vv); err != nil {
						return fmt.Errorf("Error reading nntp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nntp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("spam", flattenObjectSystemReplacemsgGroupSpam(o["spam"], d, "spam")); err != nil {
			if vv, ok := fortiAPIPatch(o["spam"], "ObjectSystemReplacemsgGroup-Spam"); ok {
				if err = d.Set("spam", vv); err != nil {
					return fmt.Errorf("Error reading spam: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading spam: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("spam"); ok {
			if err = d.Set("spam", flattenObjectSystemReplacemsgGroupSpam(o["spam"], d, "spam")); err != nil {
				if vv, ok := fortiAPIPatch(o["spam"], "ObjectSystemReplacemsgGroup-Spam"); ok {
					if err = d.Set("spam", vv); err != nil {
						return fmt.Errorf("Error reading spam: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading spam: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("sslvpn", flattenObjectSystemReplacemsgGroupSslvpn(o["sslvpn"], d, "sslvpn")); err != nil {
			if vv, ok := fortiAPIPatch(o["sslvpn"], "ObjectSystemReplacemsgGroup-Sslvpn"); ok {
				if err = d.Set("sslvpn", vv); err != nil {
					return fmt.Errorf("Error reading sslvpn: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sslvpn: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sslvpn"); ok {
			if err = d.Set("sslvpn", flattenObjectSystemReplacemsgGroupSslvpn(o["sslvpn"], d, "sslvpn")); err != nil {
				if vv, ok := fortiAPIPatch(o["sslvpn"], "ObjectSystemReplacemsgGroup-Sslvpn"); ok {
					if err = d.Set("sslvpn", vv); err != nil {
						return fmt.Errorf("Error reading sslvpn: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sslvpn: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("traffic_quota", flattenObjectSystemReplacemsgGroupTrafficQuota(o["traffic-quota"], d, "traffic_quota")); err != nil {
			if vv, ok := fortiAPIPatch(o["traffic-quota"], "ObjectSystemReplacemsgGroup-TrafficQuota"); ok {
				if err = d.Set("traffic_quota", vv); err != nil {
					return fmt.Errorf("Error reading traffic_quota: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading traffic_quota: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("traffic_quota"); ok {
			if err = d.Set("traffic_quota", flattenObjectSystemReplacemsgGroupTrafficQuota(o["traffic-quota"], d, "traffic_quota")); err != nil {
				if vv, ok := fortiAPIPatch(o["traffic-quota"], "ObjectSystemReplacemsgGroup-TrafficQuota"); ok {
					if err = d.Set("traffic_quota", vv); err != nil {
						return fmt.Errorf("Error reading traffic_quota: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading traffic_quota: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("utm", flattenObjectSystemReplacemsgGroupUtm(o["utm"], d, "utm")); err != nil {
			if vv, ok := fortiAPIPatch(o["utm"], "ObjectSystemReplacemsgGroup-Utm"); ok {
				if err = d.Set("utm", vv); err != nil {
					return fmt.Errorf("Error reading utm: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading utm: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("utm"); ok {
			if err = d.Set("utm", flattenObjectSystemReplacemsgGroupUtm(o["utm"], d, "utm")); err != nil {
				if vv, ok := fortiAPIPatch(o["utm"], "ObjectSystemReplacemsgGroup-Utm"); ok {
					if err = d.Set("utm", vv); err != nil {
						return fmt.Errorf("Error reading utm: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading utm: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("webproxy", flattenObjectSystemReplacemsgGroupWebproxy(o["webproxy"], d, "webproxy")); err != nil {
			if vv, ok := fortiAPIPatch(o["webproxy"], "ObjectSystemReplacemsgGroup-Webproxy"); ok {
				if err = d.Set("webproxy", vv); err != nil {
					return fmt.Errorf("Error reading webproxy: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading webproxy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("webproxy"); ok {
			if err = d.Set("webproxy", flattenObjectSystemReplacemsgGroupWebproxy(o["webproxy"], d, "webproxy")); err != nil {
				if vv, ok := fortiAPIPatch(o["webproxy"], "ObjectSystemReplacemsgGroup-Webproxy"); ok {
					if err = d.Set("webproxy", vv); err != nil {
						return fmt.Errorf("Error reading webproxy: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading webproxy: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectSystemReplacemsgGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemReplacemsgGroupAdmin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupAdminBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupAdminFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupAdminHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupAdminMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupAdminBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupAdminFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupAdminHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupAdminMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupAlertmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupAlertmailBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupAlertmailFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupAlertmailHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupAlertmailMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupAlertmailBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupAlertmailFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupAlertmailHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupAlertmailMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupAuthBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupAuthFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupAuthHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupAuthMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupAuthBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupAuthFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupAuthHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupAuthMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupAutomation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupAutomationBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupAutomationFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupAutomationHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupAutomationMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupAutomationBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupAutomationFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupAutomationHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupAutomationMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupCustomMessage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupCustomMessageBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupCustomMessageFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupCustomMessageHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupCustomMessageMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupCustomMessageBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupCustomMessageFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupCustomMessageHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupCustomMessageMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupDeviceDetectionPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupDeviceDetectionPortalBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupDeviceDetectionPortalFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupDeviceDetectionPortalHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupDeviceDetectionPortalMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupDeviceDetectionPortalBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupDeviceDetectionPortalFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupDeviceDetectionPortalHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupDeviceDetectionPortalMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupEc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupEcBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupEcFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupEcHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupEcMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupEcBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupEcFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupEcHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupEcMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupFortiguardWf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupFortiguardWfBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupFortiguardWfFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupFortiguardWfHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupFortiguardWfMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupFortiguardWfBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupFortiguardWfFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupFortiguardWfHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupFortiguardWfMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupFtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupFtpBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupFtpFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupFtpHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupFtpMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupFtpBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupFtpFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupFtpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupFtpMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupGroupType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupHttpBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupHttpFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupHttpHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupHttpMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupHttpBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupHttpFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupHttpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupHttpMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupIcap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupIcapBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupIcapFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupIcapHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupIcapMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupIcapBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupIcapFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupIcapHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupIcapMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupMailBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupMailFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupMailHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupMailMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupMailBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMailFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMailHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMailMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_smil"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["add-smil"], _ = expandObjectSystemReplacemsgGroupMm1AddSmil(d, i["add_smil"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "charset"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["charset"], _ = expandObjectSystemReplacemsgGroupMm1Charset(d, i["charset"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["class"], _ = expandObjectSystemReplacemsgGroupMm1Class(d, i["class"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupMm1Format(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["from"], _ = expandObjectSystemReplacemsgGroupMm1From(d, i["from"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from_sender"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["from-sender"], _ = expandObjectSystemReplacemsgGroupMm1FromSender(d, i["from_sender"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupMm1Header(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "image"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["image"], _ = expandObjectSystemReplacemsgGroupMm1Image(d, i["image"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "message"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["message"], _ = expandObjectSystemReplacemsgGroupMm1Message(d, i["message"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupMm1MsgType(d, i["msg_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectSystemReplacemsgGroupMm1Priority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsp_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsp-status"], _ = expandObjectSystemReplacemsgGroupMm1RspStatus(d, i["rsp_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsp_text"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsp-text"], _ = expandObjectSystemReplacemsgGroupMm1RspText(d, i["rsp_text"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sender_visibility"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sender-visibility"], _ = expandObjectSystemReplacemsgGroupMm1SenderVisibility(d, i["sender_visibility"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smil_part"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["smil-part"], _ = expandObjectSystemReplacemsgGroupMm1SmilPart(d, i["smil_part"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subject"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subject"], _ = expandObjectSystemReplacemsgGroupMm1Subject(d, i["subject"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupMm1AddSmil(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1Charset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1Class(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1Format(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1From(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1FromSender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1Header(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1Image(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1Message(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1MsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1Priority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1RspStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1RspText(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1SenderVisibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1SmilPart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm1Subject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_html"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["add-html"], _ = expandObjectSystemReplacemsgGroupMm3AddHtml(d, i["add_html"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "charset"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["charset"], _ = expandObjectSystemReplacemsgGroupMm3Charset(d, i["charset"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupMm3Format(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["from"], _ = expandObjectSystemReplacemsgGroupMm3From(d, i["from"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from_sender"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["from-sender"], _ = expandObjectSystemReplacemsgGroupMm3FromSender(d, i["from_sender"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupMm3Header(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "html_part"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["html-part"], _ = expandObjectSystemReplacemsgGroupMm3HtmlPart(d, i["html_part"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "image"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["image"], _ = expandObjectSystemReplacemsgGroupMm3Image(d, i["image"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "message"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["message"], _ = expandObjectSystemReplacemsgGroupMm3Message(d, i["message"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupMm3MsgType(d, i["msg_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectSystemReplacemsgGroupMm3Priority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subject"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subject"], _ = expandObjectSystemReplacemsgGroupMm3Subject(d, i["subject"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupMm3AddHtml(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm3Charset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm3Format(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm3From(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm3FromSender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm3Header(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm3HtmlPart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm3Image(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm3Message(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm3MsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm3Priority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm3Subject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_smil"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["add-smil"], _ = expandObjectSystemReplacemsgGroupMm4AddSmil(d, i["add_smil"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "charset"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["charset"], _ = expandObjectSystemReplacemsgGroupMm4Charset(d, i["charset"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["class"], _ = expandObjectSystemReplacemsgGroupMm4Class(d, i["class"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain"], _ = expandObjectSystemReplacemsgGroupMm4Domain(d, i["domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupMm4Format(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["from"], _ = expandObjectSystemReplacemsgGroupMm4From(d, i["from"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from_sender"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["from-sender"], _ = expandObjectSystemReplacemsgGroupMm4FromSender(d, i["from_sender"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupMm4Header(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "image"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["image"], _ = expandObjectSystemReplacemsgGroupMm4Image(d, i["image"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "message"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["message"], _ = expandObjectSystemReplacemsgGroupMm4Message(d, i["message"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupMm4MsgType(d, i["msg_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectSystemReplacemsgGroupMm4Priority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsp_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsp-status"], _ = expandObjectSystemReplacemsgGroupMm4RspStatus(d, i["rsp_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smil_part"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["smil-part"], _ = expandObjectSystemReplacemsgGroupMm4SmilPart(d, i["smil_part"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subject"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subject"], _ = expandObjectSystemReplacemsgGroupMm4Subject(d, i["subject"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupMm4AddSmil(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm4Charset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm4Class(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm4Domain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm4Format(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm4From(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm4FromSender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm4Header(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm4Image(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm4Message(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm4MsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm4Priority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm4RspStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm4SmilPart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm4Subject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_smil"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["add-smil"], _ = expandObjectSystemReplacemsgGroupMm7AddSmil(d, i["add_smil"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-type"], _ = expandObjectSystemReplacemsgGroupMm7AddrType(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allow_content_adaptation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allow-content-adaptation"], _ = expandObjectSystemReplacemsgGroupMm7AllowContentAdaptation(d, i["allow_content_adaptation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "charset"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["charset"], _ = expandObjectSystemReplacemsgGroupMm7Charset(d, i["charset"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["class"], _ = expandObjectSystemReplacemsgGroupMm7Class(d, i["class"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupMm7Format(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["from"], _ = expandObjectSystemReplacemsgGroupMm7From(d, i["from"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from_sender"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["from-sender"], _ = expandObjectSystemReplacemsgGroupMm7FromSender(d, i["from_sender"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupMm7Header(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "image"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["image"], _ = expandObjectSystemReplacemsgGroupMm7Image(d, i["image"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "message"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["message"], _ = expandObjectSystemReplacemsgGroupMm7Message(d, i["message"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupMm7MsgType(d, i["msg_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectSystemReplacemsgGroupMm7Priority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsp_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsp-status"], _ = expandObjectSystemReplacemsgGroupMm7RspStatus(d, i["rsp_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smil_part"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["smil-part"], _ = expandObjectSystemReplacemsgGroupMm7SmilPart(d, i["smil_part"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subject"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subject"], _ = expandObjectSystemReplacemsgGroupMm7Subject(d, i["subject"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupMm7AddSmil(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7AddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7AllowContentAdaptation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7Charset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7Class(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7Format(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7From(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7FromSender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7Header(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7Image(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7Message(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7MsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7Priority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7RspStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7SmilPart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMm7Subject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupMmsBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "charset"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["charset"], _ = expandObjectSystemReplacemsgGroupMmsCharset(d, i["charset"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupMmsFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupMmsHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "image"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["image"], _ = expandObjectSystemReplacemsgGroupMmsImage(d, i["image"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupMmsMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupMmsBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMmsCharset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMmsFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMmsHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMmsImage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupMmsMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupNacQuar(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupNacQuarBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupNacQuarFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupNacQuarHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupNacQuarMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupNacQuarBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupNacQuarFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupNacQuarHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupNacQuarMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupNntp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupNntpBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupNntpFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupNntpHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupNntpMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupNntpBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupNntpFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupNntpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupNntpMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupSpam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupSpamBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupSpamFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupSpamHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupSpamMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupSpamBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupSpamFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupSpamHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupSpamMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupSslvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupSslvpnBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupSslvpnFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupSslvpnHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupSslvpnMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupSslvpnBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupSslvpnFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupSslvpnHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupSslvpnMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupTrafficQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupTrafficQuotaBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupTrafficQuotaFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupTrafficQuotaHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupTrafficQuotaMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupTrafficQuotaBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupTrafficQuotaFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupTrafficQuotaHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupTrafficQuotaMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupUtm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupUtmBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupUtmFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupUtmHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupUtmMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupUtmBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupUtmFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupUtmHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupUtmMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupWebproxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupWebproxyBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupWebproxyFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupWebproxyHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandObjectSystemReplacemsgGroupWebproxyMsgType(d, i["msg_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemReplacemsgGroupWebproxyBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupWebproxyFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupWebproxyHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgGroupWebproxyMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemReplacemsgGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("admin"); ok || d.HasChange("admin") {
		t, err := expandObjectSystemReplacemsgGroupAdmin(d, v, "admin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin"] = t
		}
	}

	if v, ok := d.GetOk("alertmail"); ok || d.HasChange("alertmail") {
		t, err := expandObjectSystemReplacemsgGroupAlertmail(d, v, "alertmail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alertmail"] = t
		}
	}

	if v, ok := d.GetOk("auth"); ok || d.HasChange("auth") {
		t, err := expandObjectSystemReplacemsgGroupAuth(d, v, "auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("automation"); ok || d.HasChange("automation") {
		t, err := expandObjectSystemReplacemsgGroupAutomation(d, v, "automation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["automation"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectSystemReplacemsgGroupComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("custom_message"); ok || d.HasChange("custom_message") {
		t, err := expandObjectSystemReplacemsgGroupCustomMessage(d, v, "custom_message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-message"] = t
		}
	}

	if v, ok := d.GetOk("device_detection_portal"); ok || d.HasChange("device_detection_portal") {
		t, err := expandObjectSystemReplacemsgGroupDeviceDetectionPortal(d, v, "device_detection_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-detection-portal"] = t
		}
	}

	if v, ok := d.GetOk("ec"); ok || d.HasChange("ec") {
		t, err := expandObjectSystemReplacemsgGroupEc(d, v, "ec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ec"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_wf"); ok || d.HasChange("fortiguard_wf") {
		t, err := expandObjectSystemReplacemsgGroupFortiguardWf(d, v, "fortiguard_wf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-wf"] = t
		}
	}

	if v, ok := d.GetOk("ftp"); ok || d.HasChange("ftp") {
		t, err := expandObjectSystemReplacemsgGroupFtp(d, v, "ftp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp"] = t
		}
	}

	if v, ok := d.GetOk("group_type"); ok || d.HasChange("group_type") {
		t, err := expandObjectSystemReplacemsgGroupGroupType(d, v, "group_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-type"] = t
		}
	}

	if v, ok := d.GetOk("http"); ok || d.HasChange("http") {
		t, err := expandObjectSystemReplacemsgGroupHttp(d, v, "http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http"] = t
		}
	}

	if v, ok := d.GetOk("icap"); ok || d.HasChange("icap") {
		t, err := expandObjectSystemReplacemsgGroupIcap(d, v, "icap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap"] = t
		}
	}

	if v, ok := d.GetOk("mail"); ok || d.HasChange("mail") {
		t, err := expandObjectSystemReplacemsgGroupMail(d, v, "mail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mail"] = t
		}
	}

	if v, ok := d.GetOk("mm1"); ok || d.HasChange("mm1") {
		t, err := expandObjectSystemReplacemsgGroupMm1(d, v, "mm1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1"] = t
		}
	}

	if v, ok := d.GetOk("mm3"); ok || d.HasChange("mm3") {
		t, err := expandObjectSystemReplacemsgGroupMm3(d, v, "mm3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm3"] = t
		}
	}

	if v, ok := d.GetOk("mm4"); ok || d.HasChange("mm4") {
		t, err := expandObjectSystemReplacemsgGroupMm4(d, v, "mm4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm4"] = t
		}
	}

	if v, ok := d.GetOk("mm7"); ok || d.HasChange("mm7") {
		t, err := expandObjectSystemReplacemsgGroupMm7(d, v, "mm7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7"] = t
		}
	}

	if v, ok := d.GetOk("mms"); ok || d.HasChange("mms") {
		t, err := expandObjectSystemReplacemsgGroupMms(d, v, "mms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms"] = t
		}
	}

	if v, ok := d.GetOk("nac_quar"); ok || d.HasChange("nac_quar") {
		t, err := expandObjectSystemReplacemsgGroupNacQuar(d, v, "nac_quar")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-quar"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSystemReplacemsgGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nntp"); ok || d.HasChange("nntp") {
		t, err := expandObjectSystemReplacemsgGroupNntp(d, v, "nntp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nntp"] = t
		}
	}

	if v, ok := d.GetOk("spam"); ok || d.HasChange("spam") {
		t, err := expandObjectSystemReplacemsgGroupSpam(d, v, "spam")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn"); ok || d.HasChange("sslvpn") {
		t, err := expandObjectSystemReplacemsgGroupSslvpn(d, v, "sslvpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn"] = t
		}
	}

	if v, ok := d.GetOk("traffic_quota"); ok || d.HasChange("traffic_quota") {
		t, err := expandObjectSystemReplacemsgGroupTrafficQuota(d, v, "traffic_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-quota"] = t
		}
	}

	if v, ok := d.GetOk("utm"); ok || d.HasChange("utm") {
		t, err := expandObjectSystemReplacemsgGroupUtm(d, v, "utm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm"] = t
		}
	}

	if v, ok := d.GetOk("webproxy"); ok || d.HasChange("webproxy") {
		t, err := expandObjectSystemReplacemsgGroupWebproxy(d, v, "webproxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy"] = t
		}
	}

	return &obj, nil
}
