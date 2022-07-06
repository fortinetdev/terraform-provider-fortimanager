// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: workflow approval matrix.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemWorkflowApprovalMatrix() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemWorkflowApprovalMatrixCreate,
		Read:   resourceSystemWorkflowApprovalMatrixRead,
		Update: resourceSystemWorkflowApprovalMatrixUpdate,
		Delete: resourceSystemWorkflowApprovalMatrixDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"adom_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"approver": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"seq_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mail_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"notify": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceSystemWorkflowApprovalMatrixCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemWorkflowApprovalMatrix(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemWorkflowApprovalMatrix resource while getting object: %v", err)
	}

	_, err = c.CreateSystemWorkflowApprovalMatrix(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemWorkflowApprovalMatrix resource: %v", err)
	}

	d.SetId(getStringKey(d, "adom_name"))

	return resourceSystemWorkflowApprovalMatrixRead(d, m)
}

func resourceSystemWorkflowApprovalMatrixUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemWorkflowApprovalMatrix(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemWorkflowApprovalMatrix resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemWorkflowApprovalMatrix(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemWorkflowApprovalMatrix resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "adom_name"))

	return resourceSystemWorkflowApprovalMatrixRead(d, m)
}

func resourceSystemWorkflowApprovalMatrixDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemWorkflowApprovalMatrix(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemWorkflowApprovalMatrix resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemWorkflowApprovalMatrixRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemWorkflowApprovalMatrix(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemWorkflowApprovalMatrix resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemWorkflowApprovalMatrix(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemWorkflowApprovalMatrix resource from API: %v", err)
	}
	return nil
}

func flattenSystemWorkflowApprovalMatrixAdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWorkflowApprovalMatrixApprover(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := i["member"]; ok {
			v := flattenSystemWorkflowApprovalMatrixApproverMember(i["member"], d, pre_append)
			tmp["member"] = fortiAPISubPartPatch(v, "SystemWorkflowApprovalMatrix-Approver-Member")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := i["seq_num"]; ok {
			v := flattenSystemWorkflowApprovalMatrixApproverSeqNum(i["seq_num"], d, pre_append)
			tmp["seq_num"] = fortiAPISubPartPatch(v, "SystemWorkflowApprovalMatrix-Approver-SeqNum")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemWorkflowApprovalMatrixApproverMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWorkflowApprovalMatrixApproverSeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWorkflowApprovalMatrixMailServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWorkflowApprovalMatrixNotify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemWorkflowApprovalMatrix(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("adom_name", flattenSystemWorkflowApprovalMatrixAdomName(o["adom-name"], d, "adom_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-name"], "SystemWorkflowApprovalMatrix-AdomName"); ok {
			if err = d.Set("adom_name", vv); err != nil {
				return fmt.Errorf("Error reading adom_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("approver", flattenSystemWorkflowApprovalMatrixApprover(o["approver"], d, "approver")); err != nil {
			if vv, ok := fortiAPIPatch(o["approver"], "SystemWorkflowApprovalMatrix-Approver"); ok {
				if err = d.Set("approver", vv); err != nil {
					return fmt.Errorf("Error reading approver: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading approver: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("approver"); ok {
			if err = d.Set("approver", flattenSystemWorkflowApprovalMatrixApprover(o["approver"], d, "approver")); err != nil {
				if vv, ok := fortiAPIPatch(o["approver"], "SystemWorkflowApprovalMatrix-Approver"); ok {
					if err = d.Set("approver", vv); err != nil {
						return fmt.Errorf("Error reading approver: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading approver: %v", err)
				}
			}
		}
	}

	if err = d.Set("mail_server", flattenSystemWorkflowApprovalMatrixMailServer(o["mail-server"], d, "mail_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["mail-server"], "SystemWorkflowApprovalMatrix-MailServer"); ok {
			if err = d.Set("mail_server", vv); err != nil {
				return fmt.Errorf("Error reading mail_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mail_server: %v", err)
		}
	}

	if err = d.Set("notify", flattenSystemWorkflowApprovalMatrixNotify(o["notify"], d, "notify")); err != nil {
		if vv, ok := fortiAPIPatch(o["notify"], "SystemWorkflowApprovalMatrix-Notify"); ok {
			if err = d.Set("notify", vv); err != nil {
				return fmt.Errorf("Error reading notify: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notify: %v", err)
		}
	}

	return nil
}

func flattenSystemWorkflowApprovalMatrixFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemWorkflowApprovalMatrixAdomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWorkflowApprovalMatrixApprover(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member"], _ = expandSystemWorkflowApprovalMatrixApproverMember(d, i["member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq_num"], _ = expandSystemWorkflowApprovalMatrixApproverSeqNum(d, i["seq_num"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemWorkflowApprovalMatrixApproverMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWorkflowApprovalMatrixApproverSeqNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWorkflowApprovalMatrixMailServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWorkflowApprovalMatrixNotify(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemWorkflowApprovalMatrix(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("adom_name"); ok || d.HasChange("adom_name") {
		t, err := expandSystemWorkflowApprovalMatrixAdomName(d, v, "adom_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-name"] = t
		}
	}

	if v, ok := d.GetOk("approver"); ok || d.HasChange("approver") {
		t, err := expandSystemWorkflowApprovalMatrixApprover(d, v, "approver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["approver"] = t
		}
	}

	if v, ok := d.GetOk("mail_server"); ok || d.HasChange("mail_server") {
		t, err := expandSystemWorkflowApprovalMatrixMailServer(d, v, "mail_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mail-server"] = t
		}
	}

	if v, ok := d.GetOk("notify"); ok || d.HasChange("notify") {
		t, err := expandSystemWorkflowApprovalMatrixNotify(d, v, "notify")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notify"] = t
		}
	}

	return &obj, nil
}
