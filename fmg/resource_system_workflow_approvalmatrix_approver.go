// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Approver.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemWorkflowApprovalMatrixApprover() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemWorkflowApprovalMatrixApproverCreate,
		Read:   resourceSystemWorkflowApprovalMatrixApproverRead,
		Update: resourceSystemWorkflowApprovalMatrixApproverUpdate,
		Delete: resourceSystemWorkflowApprovalMatrixApproverDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"approval_matrix": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"seq_num": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemWorkflowApprovalMatrixApproverCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	approval_matrix := d.Get("approval_matrix").(string)
	paradict["approval_matrix"] = approval_matrix

	obj, err := getObjectSystemWorkflowApprovalMatrixApprover(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemWorkflowApprovalMatrixApprover resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystemWorkflowApprovalMatrixApprover(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemWorkflowApprovalMatrixApprover resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "seq_num")))

	return resourceSystemWorkflowApprovalMatrixApproverRead(d, m)
}

func resourceSystemWorkflowApprovalMatrixApproverUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	approval_matrix := d.Get("approval_matrix").(string)
	paradict["approval_matrix"] = approval_matrix

	obj, err := getObjectSystemWorkflowApprovalMatrixApprover(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemWorkflowApprovalMatrixApprover resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemWorkflowApprovalMatrixApprover(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemWorkflowApprovalMatrixApprover resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "seq_num")))

	return resourceSystemWorkflowApprovalMatrixApproverRead(d, m)
}

func resourceSystemWorkflowApprovalMatrixApproverDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	approval_matrix := d.Get("approval_matrix").(string)
	paradict["approval_matrix"] = approval_matrix

	wsParams["adom"] = adomv

	err = c.DeleteSystemWorkflowApprovalMatrixApprover(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemWorkflowApprovalMatrixApprover resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemWorkflowApprovalMatrixApproverRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	approval_matrix := d.Get("approval_matrix").(string)
	if approval_matrix == "" {
		approval_matrix = importOptionChecking(m.(*FortiClient).Cfg, "approval_matrix")
		if approval_matrix == "" {
			return fmt.Errorf("Parameter approval_matrix is missing")
		}
		if err = d.Set("approval_matrix", approval_matrix); err != nil {
			return fmt.Errorf("Error set params approval_matrix: %v", err)
		}
	}
	paradict["approval_matrix"] = approval_matrix

	o, err := c.ReadSystemWorkflowApprovalMatrixApprover(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemWorkflowApprovalMatrixApprover resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemWorkflowApprovalMatrixApprover(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemWorkflowApprovalMatrixApprover resource from API: %v", err)
	}
	return nil
}

func flattenSystemWorkflowApprovalMatrixApproverMember2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWorkflowApprovalMatrixApproverSeqNum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemWorkflowApprovalMatrixApprover(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("member", flattenSystemWorkflowApprovalMatrixApproverMember2edl(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "SystemWorkflowApprovalMatrixApprover-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("seq_num", flattenSystemWorkflowApprovalMatrixApproverSeqNum2edl(o["seq_num"], d, "seq_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["seq_num"], "SystemWorkflowApprovalMatrixApprover-SeqNum"); ok {
			if err = d.Set("seq_num", vv); err != nil {
				return fmt.Errorf("Error reading seq_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	return nil
}

func flattenSystemWorkflowApprovalMatrixApproverFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemWorkflowApprovalMatrixApproverMember2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWorkflowApprovalMatrixApproverSeqNum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemWorkflowApprovalMatrixApprover(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandSystemWorkflowApprovalMatrixApproverMember2edl(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("seq_num"); ok || d.HasChange("seq_num") {
		t, err := expandSystemWorkflowApprovalMatrixApproverSeqNum2edl(d, v, "seq_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq_num"] = t
		}
	}

	return &obj, nil
}
