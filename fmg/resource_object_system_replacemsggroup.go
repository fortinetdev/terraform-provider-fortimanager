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
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
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
			"custom_message": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"group_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupAdminHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupAlertmailHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupAuthHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupAuthMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupCustomMessageHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupDeviceDetectionPortalHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupDeviceDetectionPortalMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupFortiguardWfHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupFtpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupFtpMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgGroupGroupType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "default",
			1: "utm",
			2: "auth",
		}
		res := getEnumVal(v, emap)
		return res
	}
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupHttpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupIcapHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupMailHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupMailMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupNacQuarHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupNntpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupSpamHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupSslvpnHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupTrafficQuotaHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupUtmHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
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
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "text",
			3: "html",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupWebproxyHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "http",
			3: "8bit",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSystemReplacemsgGroupWebproxyMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemReplacemsgGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupAdminBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupAdminFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupAdminHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupAlertmailBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupAlertmailFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupAlertmailHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupAuthBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupAuthFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupAuthHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupCustomMessageBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupCustomMessageFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupCustomMessageHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupDeviceDetectionPortalBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupDeviceDetectionPortalFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupDeviceDetectionPortalHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupFortiguardWfBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupFortiguardWfFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupFortiguardWfHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupFtpBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupFtpFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupFtpHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupHttpBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupHttpFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupHttpHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupIcapBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupIcapFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupIcapHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupMailBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupMailFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupMailHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupNacQuarBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupNacQuarFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupNacQuarHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupNntpBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupNntpFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupNntpHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupSpamBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupSpamFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupSpamHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupSslvpnBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupSslvpnFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupSslvpnHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupTrafficQuotaBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupTrafficQuotaFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupTrafficQuotaHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupUtmBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupUtmFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupUtmHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandObjectSystemReplacemsgGroupWebproxyBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandObjectSystemReplacemsgGroupWebproxyFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandObjectSystemReplacemsgGroupWebproxyHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
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

	if v, ok := d.GetOk("admin"); ok {
		t, err := expandObjectSystemReplacemsgGroupAdmin(d, v, "admin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin"] = t
		}
	}

	if v, ok := d.GetOk("alertmail"); ok {
		t, err := expandObjectSystemReplacemsgGroupAlertmail(d, v, "alertmail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alertmail"] = t
		}
	}

	if v, ok := d.GetOk("auth"); ok {
		t, err := expandObjectSystemReplacemsgGroupAuth(d, v, "auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectSystemReplacemsgGroupComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("custom_message"); ok {
		t, err := expandObjectSystemReplacemsgGroupCustomMessage(d, v, "custom_message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-message"] = t
		}
	}

	if v, ok := d.GetOk("device_detection_portal"); ok {
		t, err := expandObjectSystemReplacemsgGroupDeviceDetectionPortal(d, v, "device_detection_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-detection-portal"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_wf"); ok {
		t, err := expandObjectSystemReplacemsgGroupFortiguardWf(d, v, "fortiguard_wf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-wf"] = t
		}
	}

	if v, ok := d.GetOk("ftp"); ok {
		t, err := expandObjectSystemReplacemsgGroupFtp(d, v, "ftp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp"] = t
		}
	}

	if v, ok := d.GetOk("group_type"); ok {
		t, err := expandObjectSystemReplacemsgGroupGroupType(d, v, "group_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-type"] = t
		}
	}

	if v, ok := d.GetOk("http"); ok {
		t, err := expandObjectSystemReplacemsgGroupHttp(d, v, "http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http"] = t
		}
	}

	if v, ok := d.GetOk("icap"); ok {
		t, err := expandObjectSystemReplacemsgGroupIcap(d, v, "icap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap"] = t
		}
	}

	if v, ok := d.GetOk("mail"); ok {
		t, err := expandObjectSystemReplacemsgGroupMail(d, v, "mail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mail"] = t
		}
	}

	if v, ok := d.GetOk("nac_quar"); ok {
		t, err := expandObjectSystemReplacemsgGroupNacQuar(d, v, "nac_quar")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-quar"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectSystemReplacemsgGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nntp"); ok {
		t, err := expandObjectSystemReplacemsgGroupNntp(d, v, "nntp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nntp"] = t
		}
	}

	if v, ok := d.GetOk("spam"); ok {
		t, err := expandObjectSystemReplacemsgGroupSpam(d, v, "spam")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn"); ok {
		t, err := expandObjectSystemReplacemsgGroupSslvpn(d, v, "sslvpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn"] = t
		}
	}

	if v, ok := d.GetOk("traffic_quota"); ok {
		t, err := expandObjectSystemReplacemsgGroupTrafficQuota(d, v, "traffic_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-quota"] = t
		}
	}

	if v, ok := d.GetOk("utm"); ok {
		t, err := expandObjectSystemReplacemsgGroupUtm(d, v, "utm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm"] = t
		}
	}

	if v, ok := d.GetOk("webproxy"); ok {
		t, err := expandObjectSystemReplacemsgGroupWebproxy(d, v, "webproxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy"] = t
		}
	}

	return &obj, nil
}
