// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure ingress ACL policies to be applied on managed FortiSwitch ports.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerAclIngress() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerAclIngressCreate,
		Read:   resourceObjectSwitchControllerAclIngressRead,
		Update: resourceObjectSwitchControllerAclIngressUpdate,
		Delete: resourceObjectSwitchControllerAclIngressDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"classifier": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_ip_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dst_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"src_ip_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"src_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSwitchControllerAclIngressCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSwitchControllerAclIngress(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerAclIngress resource while getting object: %v", err)
	}

	v, err := c.CreateObjectSwitchControllerAclIngress(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerAclIngress resource: %v", err)
	}

	if v != nil && v["id"] != nil {
		if vidn, ok := v["id"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceObjectSwitchControllerAclIngressRead(d, m)
		} else {
			return fmt.Errorf("Error creating ObjectSwitchControllerAclIngress resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSwitchControllerAclIngressRead(d, m)
}

func resourceObjectSwitchControllerAclIngressUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSwitchControllerAclIngress(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerAclIngress resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerAclIngress(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerAclIngress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSwitchControllerAclIngressRead(d, m)
}

func resourceObjectSwitchControllerAclIngressDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSwitchControllerAclIngress(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerAclIngress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerAclIngressRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSwitchControllerAclIngress(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerAclIngress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerAclIngress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerAclIngress resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerAclIngressAction(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "count"
	if _, ok := i["count"]; ok {
		result["count"] = flattenObjectSwitchControllerAclIngressActionCount(i["count"], d, pre_append)
	}

	pre_append = pre + ".0." + "drop"
	if _, ok := i["drop"]; ok {
		result["drop"] = flattenObjectSwitchControllerAclIngressActionDrop(i["drop"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSwitchControllerAclIngressActionCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerAclIngressActionDrop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerAclIngressClassifier(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dst_ip_prefix"
	if _, ok := i["dst-ip-prefix"]; ok {
		result["dst_ip_prefix"] = flattenObjectSwitchControllerAclIngressClassifierDstIpPrefix(i["dst-ip-prefix"], d, pre_append)
	}

	pre_append = pre + ".0." + "dst_mac"
	if _, ok := i["dst-mac"]; ok {
		result["dst_mac"] = flattenObjectSwitchControllerAclIngressClassifierDstMac(i["dst-mac"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_ip_prefix"
	if _, ok := i["src-ip-prefix"]; ok {
		result["src_ip_prefix"] = flattenObjectSwitchControllerAclIngressClassifierSrcIpPrefix(i["src-ip-prefix"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_mac"
	if _, ok := i["src-mac"]; ok {
		result["src_mac"] = flattenObjectSwitchControllerAclIngressClassifierSrcMac(i["src-mac"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenObjectSwitchControllerAclIngressClassifierVlan(i["vlan"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSwitchControllerAclIngressClassifierDstIpPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenObjectSwitchControllerAclIngressClassifierDstMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerAclIngressClassifierSrcIpPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convipstringlist2ipmask(v)
}

func flattenObjectSwitchControllerAclIngressClassifierSrcMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerAclIngressClassifierVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerAclIngressDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerAclIngressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerAclIngress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if isImportTable() {
		if err = d.Set("action", flattenObjectSwitchControllerAclIngressAction(o["action"], d, "action")); err != nil {
			if vv, ok := fortiAPIPatch(o["action"], "ObjectSwitchControllerAclIngress-Action"); ok {
				if err = d.Set("action", vv); err != nil {
					return fmt.Errorf("Error reading action: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading action: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("action"); ok {
			if err = d.Set("action", flattenObjectSwitchControllerAclIngressAction(o["action"], d, "action")); err != nil {
				if vv, ok := fortiAPIPatch(o["action"], "ObjectSwitchControllerAclIngress-Action"); ok {
					if err = d.Set("action", vv); err != nil {
						return fmt.Errorf("Error reading action: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading action: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("classifier", flattenObjectSwitchControllerAclIngressClassifier(o["classifier"], d, "classifier")); err != nil {
			if vv, ok := fortiAPIPatch(o["classifier"], "ObjectSwitchControllerAclIngress-Classifier"); ok {
				if err = d.Set("classifier", vv); err != nil {
					return fmt.Errorf("Error reading classifier: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading classifier: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("classifier"); ok {
			if err = d.Set("classifier", flattenObjectSwitchControllerAclIngressClassifier(o["classifier"], d, "classifier")); err != nil {
				if vv, ok := fortiAPIPatch(o["classifier"], "ObjectSwitchControllerAclIngress-Classifier"); ok {
					if err = d.Set("classifier", vv); err != nil {
						return fmt.Errorf("Error reading classifier: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading classifier: %v", err)
				}
			}
		}
	}

	if err = d.Set("description", flattenObjectSwitchControllerAclIngressDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectSwitchControllerAclIngress-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectSwitchControllerAclIngressId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSwitchControllerAclIngress-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerAclIngressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerAclIngressAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["count"], _ = expandObjectSwitchControllerAclIngressActionCount(d, i["count"], pre_append)
	}
	pre_append = pre + ".0." + "drop"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["drop"], _ = expandObjectSwitchControllerAclIngressActionDrop(d, i["drop"], pre_append)
	}

	return result, nil
}

func expandObjectSwitchControllerAclIngressActionCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerAclIngressActionDrop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerAclIngressClassifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dst_ip_prefix"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dst-ip-prefix"], _ = expandObjectSwitchControllerAclIngressClassifierDstIpPrefix(d, i["dst_ip_prefix"], pre_append)
	}
	pre_append = pre + ".0." + "dst_mac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dst-mac"], _ = expandObjectSwitchControllerAclIngressClassifierDstMac(d, i["dst_mac"], pre_append)
	}
	pre_append = pre + ".0." + "src_ip_prefix"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-ip-prefix"], _ = expandObjectSwitchControllerAclIngressClassifierSrcIpPrefix(d, i["src_ip_prefix"], pre_append)
	}
	pre_append = pre + ".0." + "src_mac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-mac"], _ = expandObjectSwitchControllerAclIngressClassifierSrcMac(d, i["src_mac"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlan"], _ = expandObjectSwitchControllerAclIngressClassifierVlan(d, i["vlan"], pre_append)
	}

	return result, nil
}

func expandObjectSwitchControllerAclIngressClassifierDstIpPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerAclIngressClassifierDstMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerAclIngressClassifierSrcIpPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerAclIngressClassifierSrcMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerAclIngressClassifierVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerAclIngressDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerAclIngressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerAclIngress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectSwitchControllerAclIngressAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("classifier"); ok || d.HasChange("classifier") {
		t, err := expandObjectSwitchControllerAclIngressClassifier(d, v, "classifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["classifier"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectSwitchControllerAclIngressDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectSwitchControllerAclIngressId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
