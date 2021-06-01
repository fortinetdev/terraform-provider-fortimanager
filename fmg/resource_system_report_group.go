// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Report group.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemReportGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemReportGroupCreate,
		Read:   resourceSystemReportGroupRead,
		Update: resourceSystemReportGroupUpdate,
		Delete: resourceSystemReportGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fmgadom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"case_insensitive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"chart_alternative": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"chart_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"chart_replace": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"group_by": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"var_expression": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"var_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"var_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"group_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"report_like": &schema.Schema{
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

func resourceSystemReportGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemReportGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemReportGroup resource while getting object: %v", err)
	}

	_, err = c.CreateSystemReportGroup(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemReportGroup resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "group_id")))

	return resourceSystemReportGroupRead(d, m)
}

func resourceSystemReportGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemReportGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemReportGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemReportGroup(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemReportGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "group_id")))

	return resourceSystemReportGroupRead(d, m)
}

func resourceSystemReportGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemReportGroup(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemReportGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReportGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemReportGroup(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemReportGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemReportGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemReportGroup resource from API: %v", err)
	}
	return nil
}

func flattenSystemReportGroupAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportGroupCaseInsensitive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemReportGroupChartAlternative(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chart_name"
		if _, ok := i["chart-name"]; ok {
			v := flattenSystemReportGroupChartAlternativeChartName(i["chart-name"], d, pre_append)
			tmp["chart_name"] = fortiAPISubPartPatch(v, "SystemReportGroup-ChartAlternative-ChartName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chart_replace"
		if _, ok := i["chart-replace"]; ok {
			v := flattenSystemReportGroupChartAlternativeChartReplace(i["chart-replace"], d, pre_append)
			tmp["chart_replace"] = fortiAPISubPartPatch(v, "SystemReportGroup-ChartAlternative-ChartReplace")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemReportGroupChartAlternativeChartName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportGroupChartAlternativeChartReplace(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportGroupGroupBy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "var_expression"
		if _, ok := i["var-expression"]; ok {
			v := flattenSystemReportGroupGroupByVarExpression(i["var-expression"], d, pre_append)
			tmp["var_expression"] = fortiAPISubPartPatch(v, "SystemReportGroup-GroupBy-VarExpression")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "var_name"
		if _, ok := i["var-name"]; ok {
			v := flattenSystemReportGroupGroupByVarName(i["var-name"], d, pre_append)
			tmp["var_name"] = fortiAPISubPartPatch(v, "SystemReportGroup-GroupBy-VarName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "var_type"
		if _, ok := i["var-type"]; ok {
			v := flattenSystemReportGroupGroupByVarType(i["var-type"], d, pre_append)
			tmp["var_type"] = fortiAPISubPartPatch(v, "SystemReportGroup-GroupBy-VarType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemReportGroupGroupByVarExpression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportGroupGroupByVarName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportGroupGroupByVarType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "integer",
			4: "string",
			5: "enum",
			6: "ip",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemReportGroupGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportGroupReportLike(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemReportGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fmgadom", flattenSystemReportGroupAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "SystemReportGroup-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if err = d.Set("case_insensitive", flattenSystemReportGroupCaseInsensitive(o["case-insensitive"], d, "case_insensitive")); err != nil {
		if vv, ok := fortiAPIPatch(o["case-insensitive"], "SystemReportGroup-CaseInsensitive"); ok {
			if err = d.Set("case_insensitive", vv); err != nil {
				return fmt.Errorf("Error reading case_insensitive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading case_insensitive: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("chart_alternative", flattenSystemReportGroupChartAlternative(o["chart-alternative"], d, "chart_alternative")); err != nil {
			if vv, ok := fortiAPIPatch(o["chart-alternative"], "SystemReportGroup-ChartAlternative"); ok {
				if err = d.Set("chart_alternative", vv); err != nil {
					return fmt.Errorf("Error reading chart_alternative: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading chart_alternative: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("chart_alternative"); ok {
			if err = d.Set("chart_alternative", flattenSystemReportGroupChartAlternative(o["chart-alternative"], d, "chart_alternative")); err != nil {
				if vv, ok := fortiAPIPatch(o["chart-alternative"], "SystemReportGroup-ChartAlternative"); ok {
					if err = d.Set("chart_alternative", vv); err != nil {
						return fmt.Errorf("Error reading chart_alternative: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading chart_alternative: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("group_by", flattenSystemReportGroupGroupBy(o["group-by"], d, "group_by")); err != nil {
			if vv, ok := fortiAPIPatch(o["group-by"], "SystemReportGroup-GroupBy"); ok {
				if err = d.Set("group_by", vv); err != nil {
					return fmt.Errorf("Error reading group_by: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading group_by: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("group_by"); ok {
			if err = d.Set("group_by", flattenSystemReportGroupGroupBy(o["group-by"], d, "group_by")); err != nil {
				if vv, ok := fortiAPIPatch(o["group-by"], "SystemReportGroup-GroupBy"); ok {
					if err = d.Set("group_by", vv); err != nil {
						return fmt.Errorf("Error reading group_by: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading group_by: %v", err)
				}
			}
		}
	}

	if err = d.Set("group_id", flattenSystemReportGroupGroupId(o["group-id"], d, "group_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-id"], "SystemReportGroup-GroupId"); ok {
			if err = d.Set("group_id", vv); err != nil {
				return fmt.Errorf("Error reading group_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_id: %v", err)
		}
	}

	if err = d.Set("report_like", flattenSystemReportGroupReportLike(o["report-like"], d, "report_like")); err != nil {
		if vv, ok := fortiAPIPatch(o["report-like"], "SystemReportGroup-ReportLike"); ok {
			if err = d.Set("report_like", vv); err != nil {
				return fmt.Errorf("Error reading report_like: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading report_like: %v", err)
		}
	}

	return nil
}

func flattenSystemReportGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemReportGroupAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportGroupCaseInsensitive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportGroupChartAlternative(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chart_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["chart-name"], _ = expandSystemReportGroupChartAlternativeChartName(d, i["chart_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chart_replace"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["chart-replace"], _ = expandSystemReportGroupChartAlternativeChartReplace(d, i["chart_replace"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReportGroupChartAlternativeChartName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportGroupChartAlternativeChartReplace(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportGroupGroupBy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "var_expression"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["var-expression"], _ = expandSystemReportGroupGroupByVarExpression(d, i["var_expression"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "var_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["var-name"], _ = expandSystemReportGroupGroupByVarName(d, i["var_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "var_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["var-type"], _ = expandSystemReportGroupGroupByVarType(d, i["var_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReportGroupGroupByVarExpression(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportGroupGroupByVarName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportGroupGroupByVarType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportGroupGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportGroupReportLike(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReportGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok {
		t, err := expandSystemReportGroupAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("case_insensitive"); ok {
		t, err := expandSystemReportGroupCaseInsensitive(d, v, "case_insensitive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-insensitive"] = t
		}
	}

	if v, ok := d.GetOk("chart_alternative"); ok {
		t, err := expandSystemReportGroupChartAlternative(d, v, "chart_alternative")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chart-alternative"] = t
		}
	}

	if v, ok := d.GetOk("group_by"); ok {
		t, err := expandSystemReportGroupGroupBy(d, v, "group_by")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-by"] = t
		}
	}

	if v, ok := d.GetOk("group_id"); ok {
		t, err := expandSystemReportGroupGroupId(d, v, "group_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-id"] = t
		}
	}

	if v, ok := d.GetOk("report_like"); ok {
		t, err := expandSystemReportGroupReportLike(d, v, "report_like")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report-like"] = t
		}
	}

	return &obj, nil
}
