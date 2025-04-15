// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Enable/disable services provided by the built-in FortiGuard.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFmupdateService() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateServiceUpdate,
		Read:   resourceFmupdateServiceRead,
		Update: resourceFmupdateServiceUpdate,
		Delete: resourceFmupdateServiceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"avips": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_antispam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_antivirus": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_filequery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_ioc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_iot": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_iot_collection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_iot_vulnerability": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_outbreak_prevention": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_webfilter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_https_traversal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateServiceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectFmupdateService(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateService resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFmupdateService(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateService")

	return resourceFmupdateServiceRead(d, m)
}

func resourceFmupdateServiceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteFmupdateService(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateServiceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadFmupdateService(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateService resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateServiceAvips(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServiceQueryAntispam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServiceQueryAntivirus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServiceQueryFilequery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServiceQueryIoc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServiceQueryIot(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServiceQueryIotCollection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServiceQueryIotVulnerability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServiceQueryOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServiceQueryWebfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServiceWebfilterHttpsTraversal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("avips", flattenFmupdateServiceAvips(o["avips"], d, "avips")); err != nil {
		if vv, ok := fortiAPIPatch(o["avips"], "FmupdateService-Avips"); ok {
			if err = d.Set("avips", vv); err != nil {
				return fmt.Errorf("Error reading avips: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading avips: %v", err)
		}
	}

	if err = d.Set("query_antispam", flattenFmupdateServiceQueryAntispam(o["query-antispam"], d, "query_antispam")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-antispam"], "FmupdateService-QueryAntispam"); ok {
			if err = d.Set("query_antispam", vv); err != nil {
				return fmt.Errorf("Error reading query_antispam: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_antispam: %v", err)
		}
	}

	if err = d.Set("query_antivirus", flattenFmupdateServiceQueryAntivirus(o["query-antivirus"], d, "query_antivirus")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-antivirus"], "FmupdateService-QueryAntivirus"); ok {
			if err = d.Set("query_antivirus", vv); err != nil {
				return fmt.Errorf("Error reading query_antivirus: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_antivirus: %v", err)
		}
	}

	if err = d.Set("query_filequery", flattenFmupdateServiceQueryFilequery(o["query-filequery"], d, "query_filequery")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-filequery"], "FmupdateService-QueryFilequery"); ok {
			if err = d.Set("query_filequery", vv); err != nil {
				return fmt.Errorf("Error reading query_filequery: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_filequery: %v", err)
		}
	}

	if err = d.Set("query_ioc", flattenFmupdateServiceQueryIoc(o["query-ioc"], d, "query_ioc")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-ioc"], "FmupdateService-QueryIoc"); ok {
			if err = d.Set("query_ioc", vv); err != nil {
				return fmt.Errorf("Error reading query_ioc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_ioc: %v", err)
		}
	}

	if err = d.Set("query_iot", flattenFmupdateServiceQueryIot(o["query-iot"], d, "query_iot")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-iot"], "FmupdateService-QueryIot"); ok {
			if err = d.Set("query_iot", vv); err != nil {
				return fmt.Errorf("Error reading query_iot: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_iot: %v", err)
		}
	}

	if err = d.Set("query_iot_collection", flattenFmupdateServiceQueryIotCollection(o["query-iot-collection"], d, "query_iot_collection")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-iot-collection"], "FmupdateService-QueryIotCollection"); ok {
			if err = d.Set("query_iot_collection", vv); err != nil {
				return fmt.Errorf("Error reading query_iot_collection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_iot_collection: %v", err)
		}
	}

	if err = d.Set("query_iot_vulnerability", flattenFmupdateServiceQueryIotVulnerability(o["query-iot-vulnerability"], d, "query_iot_vulnerability")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-iot-vulnerability"], "FmupdateService-QueryIotVulnerability"); ok {
			if err = d.Set("query_iot_vulnerability", vv); err != nil {
				return fmt.Errorf("Error reading query_iot_vulnerability: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_iot_vulnerability: %v", err)
		}
	}

	if err = d.Set("query_outbreak_prevention", flattenFmupdateServiceQueryOutbreakPrevention(o["query-outbreak-prevention"], d, "query_outbreak_prevention")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-outbreak-prevention"], "FmupdateService-QueryOutbreakPrevention"); ok {
			if err = d.Set("query_outbreak_prevention", vv); err != nil {
				return fmt.Errorf("Error reading query_outbreak_prevention: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_outbreak_prevention: %v", err)
		}
	}

	if err = d.Set("query_webfilter", flattenFmupdateServiceQueryWebfilter(o["query-webfilter"], d, "query_webfilter")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-webfilter"], "FmupdateService-QueryWebfilter"); ok {
			if err = d.Set("query_webfilter", vv); err != nil {
				return fmt.Errorf("Error reading query_webfilter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_webfilter: %v", err)
		}
	}

	if err = d.Set("webfilter_https_traversal", flattenFmupdateServiceWebfilterHttpsTraversal(o["webfilter-https-traversal"], d, "webfilter_https_traversal")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-https-traversal"], "FmupdateService-WebfilterHttpsTraversal"); ok {
			if err = d.Set("webfilter_https_traversal", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_https_traversal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_https_traversal: %v", err)
		}
	}

	return nil
}

func flattenFmupdateServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateServiceAvips(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServiceQueryAntispam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServiceQueryAntivirus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServiceQueryFilequery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServiceQueryIoc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServiceQueryIot(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServiceQueryIotCollection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServiceQueryIotVulnerability(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServiceQueryOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServiceQueryWebfilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServiceWebfilterHttpsTraversal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("avips"); ok || d.HasChange("avips") {
		t, err := expandFmupdateServiceAvips(d, v, "avips")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["avips"] = t
		}
	}

	if v, ok := d.GetOk("query_antispam"); ok || d.HasChange("query_antispam") {
		t, err := expandFmupdateServiceQueryAntispam(d, v, "query_antispam")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-antispam"] = t
		}
	}

	if v, ok := d.GetOk("query_antivirus"); ok || d.HasChange("query_antivirus") {
		t, err := expandFmupdateServiceQueryAntivirus(d, v, "query_antivirus")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-antivirus"] = t
		}
	}

	if v, ok := d.GetOk("query_filequery"); ok || d.HasChange("query_filequery") {
		t, err := expandFmupdateServiceQueryFilequery(d, v, "query_filequery")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-filequery"] = t
		}
	}

	if v, ok := d.GetOk("query_ioc"); ok || d.HasChange("query_ioc") {
		t, err := expandFmupdateServiceQueryIoc(d, v, "query_ioc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-ioc"] = t
		}
	}

	if v, ok := d.GetOk("query_iot"); ok || d.HasChange("query_iot") {
		t, err := expandFmupdateServiceQueryIot(d, v, "query_iot")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-iot"] = t
		}
	}

	if v, ok := d.GetOk("query_iot_collection"); ok || d.HasChange("query_iot_collection") {
		t, err := expandFmupdateServiceQueryIotCollection(d, v, "query_iot_collection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-iot-collection"] = t
		}
	}

	if v, ok := d.GetOk("query_iot_vulnerability"); ok || d.HasChange("query_iot_vulnerability") {
		t, err := expandFmupdateServiceQueryIotVulnerability(d, v, "query_iot_vulnerability")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-iot-vulnerability"] = t
		}
	}

	if v, ok := d.GetOk("query_outbreak_prevention"); ok || d.HasChange("query_outbreak_prevention") {
		t, err := expandFmupdateServiceQueryOutbreakPrevention(d, v, "query_outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("query_webfilter"); ok || d.HasChange("query_webfilter") {
		t, err := expandFmupdateServiceQueryWebfilter(d, v, "query_webfilter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-webfilter"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_https_traversal"); ok || d.HasChange("webfilter_https_traversal") {
		t, err := expandFmupdateServiceWebfilterHttpsTraversal(d, v, "webfilter_https_traversal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-https-traversal"] = t
		}
	}

	return &obj, nil
}
