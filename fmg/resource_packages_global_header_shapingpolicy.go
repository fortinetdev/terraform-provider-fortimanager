// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure shaping policies.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourcePackagesGlobalHeaderShapingPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesGlobalHeaderShapingPolicyCreate,
		Read:   resourcePackagesGlobalHeaderShapingPolicyRead,
		Update: resourcePackagesGlobalHeaderShapingPolicyUpdate,
		Delete: resourcePackagesGlobalHeaderShapingPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"pkg": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"app_category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"class_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode_rev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_custom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_custom_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_custom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_custom_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_ip_shaper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_shaper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_shaper_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"users": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourcePackagesGlobalHeaderShapingPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	obj, err := getObjectPackagesGlobalHeaderShapingPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesGlobalHeaderShapingPolicy resource while getting object: %v", err)
	}

	_, err = c.CreatePackagesGlobalHeaderShapingPolicy(obj, adomv, paralist)

	if err != nil {
		return fmt.Errorf("Error creating PackagesGlobalHeaderShapingPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, ""))

	return resourcePackagesGlobalHeaderShapingPolicyRead(d, m)
}

func resourcePackagesGlobalHeaderShapingPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	obj, err := getObjectPackagesGlobalHeaderShapingPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesGlobalHeaderShapingPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesGlobalHeaderShapingPolicy(obj, adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error updating PackagesGlobalHeaderShapingPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, ""))

	return resourcePackagesGlobalHeaderShapingPolicyRead(d, m)
}

func resourcePackagesGlobalHeaderShapingPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	err = c.DeletePackagesGlobalHeaderShapingPolicy(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesGlobalHeaderShapingPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesGlobalHeaderShapingPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	pkg := d.Get("pkg").(string)
	if pkg == "" {
		pkg = importOptionChecking(m.(*FortiClient).Cfg, "pkg")
		if err = d.Set("pkg", pkg); err != nil {
			return fmt.Errorf("Error set params pkg: %v", err)
		}
	}
	var paralist []string
	paralist = append(paralist, pkg)

	o, err := c.ReadPackagesGlobalHeaderShapingPolicy(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error reading PackagesGlobalHeaderShapingPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesGlobalHeaderShapingPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesGlobalHeaderShapingPolicy resource from API: %v", err)
	}
	return nil
}

func flattenPackagesGlobalHeaderShapingPolicyAppCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyAppGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenPackagesGlobalHeaderShapingPolicyClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyDiffservForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesGlobalHeaderShapingPolicyDiffservReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesGlobalHeaderShapingPolicyDiffservcodeForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyDiffservcodeRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyInternetService(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesGlobalHeaderShapingPolicyInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyInternetServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyInternetServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyInternetServiceSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesGlobalHeaderShapingPolicyInternetServiceSrcCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyInternetServiceSrcCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyInternetServiceSrcGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyInternetServiceSrcId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyInternetServiceSrcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyIpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			7: "4",
			8: "6",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyPerIpShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicySchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicySrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicySrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesGlobalHeaderShapingPolicyTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyTosMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyTosNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesGlobalHeaderShapingPolicyTrafficShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyTrafficShaperReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyUrlCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesGlobalHeaderShapingPolicyUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesGlobalHeaderShapingPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("app_category", flattenPackagesGlobalHeaderShapingPolicyAppCategory(o["app-category"], d, "app_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-category"], "PackagesGlobalHeaderShapingPolicy-AppCategory"); ok {
			if err = d.Set("app_category", vv); err != nil {
				return fmt.Errorf("Error reading app_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_category: %v", err)
		}
	}

	if err = d.Set("app_group", flattenPackagesGlobalHeaderShapingPolicyAppGroup(o["app-group"], d, "app_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-group"], "PackagesGlobalHeaderShapingPolicy-AppGroup"); ok {
			if err = d.Set("app_group", vv); err != nil {
				return fmt.Errorf("Error reading app_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_group: %v", err)
		}
	}

	if err = d.Set("application", flattenPackagesGlobalHeaderShapingPolicyApplication(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "PackagesGlobalHeaderShapingPolicy-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("class_id", flattenPackagesGlobalHeaderShapingPolicyClassId(o["class-id"], d, "class_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-id"], "PackagesGlobalHeaderShapingPolicy-ClassId"); ok {
			if err = d.Set("class_id", vv); err != nil {
				return fmt.Errorf("Error reading class_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_id: %v", err)
		}
	}

	if err = d.Set("comment", flattenPackagesGlobalHeaderShapingPolicyComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "PackagesGlobalHeaderShapingPolicy-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("diffserv_forward", flattenPackagesGlobalHeaderShapingPolicyDiffservForward(o["diffserv-forward"], d, "diffserv_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-forward"], "PackagesGlobalHeaderShapingPolicy-DiffservForward"); ok {
			if err = d.Set("diffserv_forward", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_forward: %v", err)
		}
	}

	if err = d.Set("diffserv_reverse", flattenPackagesGlobalHeaderShapingPolicyDiffservReverse(o["diffserv-reverse"], d, "diffserv_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-reverse"], "PackagesGlobalHeaderShapingPolicy-DiffservReverse"); ok {
			if err = d.Set("diffserv_reverse", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_reverse: %v", err)
		}
	}

	if err = d.Set("diffservcode_forward", flattenPackagesGlobalHeaderShapingPolicyDiffservcodeForward(o["diffservcode-forward"], d, "diffservcode_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode-forward"], "PackagesGlobalHeaderShapingPolicy-DiffservcodeForward"); ok {
			if err = d.Set("diffservcode_forward", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode_forward: %v", err)
		}
	}

	if err = d.Set("diffservcode_rev", flattenPackagesGlobalHeaderShapingPolicyDiffservcodeRev(o["diffservcode-rev"], d, "diffservcode_rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode-rev"], "PackagesGlobalHeaderShapingPolicy-DiffservcodeRev"); ok {
			if err = d.Set("diffservcode_rev", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode_rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode_rev: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenPackagesGlobalHeaderShapingPolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "PackagesGlobalHeaderShapingPolicy-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenPackagesGlobalHeaderShapingPolicyDstaddr6(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "PackagesGlobalHeaderShapingPolicy-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenPackagesGlobalHeaderShapingPolicyDstintf(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "PackagesGlobalHeaderShapingPolicy-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("groups", flattenPackagesGlobalHeaderShapingPolicyGroups(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "PackagesGlobalHeaderShapingPolicy-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("fosid", flattenPackagesGlobalHeaderShapingPolicyId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "PackagesGlobalHeaderShapingPolicy-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("internet_service", flattenPackagesGlobalHeaderShapingPolicyInternetService(o["internet-service"], d, "internet_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service"], "PackagesGlobalHeaderShapingPolicy-InternetService"); ok {
			if err = d.Set("internet_service", vv); err != nil {
				return fmt.Errorf("Error reading internet_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", flattenPackagesGlobalHeaderShapingPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom"], "PackagesGlobalHeaderShapingPolicy-InternetServiceCustom"); ok {
			if err = d.Set("internet_service_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_custom_group", flattenPackagesGlobalHeaderShapingPolicyInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom-group"], "PackagesGlobalHeaderShapingPolicy-InternetServiceCustomGroup"); ok {
			if err = d.Set("internet_service_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_group", flattenPackagesGlobalHeaderShapingPolicyInternetServiceGroup(o["internet-service-group"], d, "internet_service_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-group"], "PackagesGlobalHeaderShapingPolicy-InternetServiceGroup"); ok {
			if err = d.Set("internet_service_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_group: %v", err)
		}
	}

	if err = d.Set("internet_service_id", flattenPackagesGlobalHeaderShapingPolicyInternetServiceId(o["internet-service-id"], d, "internet_service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-id"], "PackagesGlobalHeaderShapingPolicy-InternetServiceId"); ok {
			if err = d.Set("internet_service_id", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_id: %v", err)
		}
	}

	if err = d.Set("internet_service_name", flattenPackagesGlobalHeaderShapingPolicyInternetServiceName(o["internet-service-name"], d, "internet_service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-name"], "PackagesGlobalHeaderShapingPolicy-InternetServiceName"); ok {
			if err = d.Set("internet_service_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_name: %v", err)
		}
	}

	if err = d.Set("internet_service_src", flattenPackagesGlobalHeaderShapingPolicyInternetServiceSrc(o["internet-service-src"], d, "internet_service_src")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src"], "PackagesGlobalHeaderShapingPolicy-InternetServiceSrc"); ok {
			if err = d.Set("internet_service_src", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom", flattenPackagesGlobalHeaderShapingPolicyInternetServiceSrcCustom(o["internet-service-src-custom"], d, "internet_service_src_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-custom"], "PackagesGlobalHeaderShapingPolicy-InternetServiceSrcCustom"); ok {
			if err = d.Set("internet_service_src_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom_group", flattenPackagesGlobalHeaderShapingPolicyInternetServiceSrcCustomGroup(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-custom-group"], "PackagesGlobalHeaderShapingPolicy-InternetServiceSrcCustomGroup"); ok {
			if err = d.Set("internet_service_src_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src_group", flattenPackagesGlobalHeaderShapingPolicyInternetServiceSrcGroup(o["internet-service-src-group"], d, "internet_service_src_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-group"], "PackagesGlobalHeaderShapingPolicy-InternetServiceSrcGroup"); ok {
			if err = d.Set("internet_service_src_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src_id", flattenPackagesGlobalHeaderShapingPolicyInternetServiceSrcId(o["internet-service-src-id"], d, "internet_service_src_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-id"], "PackagesGlobalHeaderShapingPolicy-InternetServiceSrcId"); ok {
			if err = d.Set("internet_service_src_id", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_id: %v", err)
		}
	}

	if err = d.Set("internet_service_src_name", flattenPackagesGlobalHeaderShapingPolicyInternetServiceSrcName(o["internet-service-src-name"], d, "internet_service_src_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-name"], "PackagesGlobalHeaderShapingPolicy-InternetServiceSrcName"); ok {
			if err = d.Set("internet_service_src_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_name: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenPackagesGlobalHeaderShapingPolicyIpVersion(o["ip-version"], d, "ip_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-version"], "PackagesGlobalHeaderShapingPolicy-IpVersion"); ok {
			if err = d.Set("ip_version", vv); err != nil {
				return fmt.Errorf("Error reading ip_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("name", flattenPackagesGlobalHeaderShapingPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesGlobalHeaderShapingPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("per_ip_shaper", flattenPackagesGlobalHeaderShapingPolicyPerIpShaper(o["per-ip-shaper"], d, "per_ip_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-ip-shaper"], "PackagesGlobalHeaderShapingPolicy-PerIpShaper"); ok {
			if err = d.Set("per_ip_shaper", vv); err != nil {
				return fmt.Errorf("Error reading per_ip_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_ip_shaper: %v", err)
		}
	}

	if err = d.Set("schedule", flattenPackagesGlobalHeaderShapingPolicySchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "PackagesGlobalHeaderShapingPolicy-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("service", flattenPackagesGlobalHeaderShapingPolicyService(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "PackagesGlobalHeaderShapingPolicy-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenPackagesGlobalHeaderShapingPolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "PackagesGlobalHeaderShapingPolicy-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenPackagesGlobalHeaderShapingPolicySrcaddr6(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "PackagesGlobalHeaderShapingPolicy-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenPackagesGlobalHeaderShapingPolicySrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "PackagesGlobalHeaderShapingPolicy-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesGlobalHeaderShapingPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesGlobalHeaderShapingPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tos", flattenPackagesGlobalHeaderShapingPolicyTos(o["tos"], d, "tos")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos"], "PackagesGlobalHeaderShapingPolicy-Tos"); ok {
			if err = d.Set("tos", vv); err != nil {
				return fmt.Errorf("Error reading tos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenPackagesGlobalHeaderShapingPolicyTosMask(o["tos-mask"], d, "tos_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos-mask"], "PackagesGlobalHeaderShapingPolicy-TosMask"); ok {
			if err = d.Set("tos_mask", vv); err != nil {
				return fmt.Errorf("Error reading tos_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("tos_negate", flattenPackagesGlobalHeaderShapingPolicyTosNegate(o["tos-negate"], d, "tos_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos-negate"], "PackagesGlobalHeaderShapingPolicy-TosNegate"); ok {
			if err = d.Set("tos_negate", vv); err != nil {
				return fmt.Errorf("Error reading tos_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos_negate: %v", err)
		}
	}

	if err = d.Set("traffic_shaper", flattenPackagesGlobalHeaderShapingPolicyTrafficShaper(o["traffic-shaper"], d, "traffic_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper"], "PackagesGlobalHeaderShapingPolicy-TrafficShaper"); ok {
			if err = d.Set("traffic_shaper", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper: %v", err)
		}
	}

	if err = d.Set("traffic_shaper_reverse", flattenPackagesGlobalHeaderShapingPolicyTrafficShaperReverse(o["traffic-shaper-reverse"], d, "traffic_shaper_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper-reverse"], "PackagesGlobalHeaderShapingPolicy-TrafficShaperReverse"); ok {
			if err = d.Set("traffic_shaper_reverse", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
		}
	}

	if err = d.Set("url_category", flattenPackagesGlobalHeaderShapingPolicyUrlCategory(o["url-category"], d, "url_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-category"], "PackagesGlobalHeaderShapingPolicy-UrlCategory"); ok {
			if err = d.Set("url_category", vv); err != nil {
				return fmt.Errorf("Error reading url_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_category: %v", err)
		}
	}

	if err = d.Set("users", flattenPackagesGlobalHeaderShapingPolicyUsers(o["users"], d, "users")); err != nil {
		if vv, ok := fortiAPIPatch(o["users"], "PackagesGlobalHeaderShapingPolicy-Users"); ok {
			if err = d.Set("users", vv); err != nil {
				return fmt.Errorf("Error reading users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading users: %v", err)
		}
	}

	if err = d.Set("uuid", flattenPackagesGlobalHeaderShapingPolicyUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "PackagesGlobalHeaderShapingPolicy-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenPackagesGlobalHeaderShapingPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesGlobalHeaderShapingPolicyAppCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyAppGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandPackagesGlobalHeaderShapingPolicyClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyDiffservForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyDiffservReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyDiffservcodeForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyDiffservcodeRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyInternetService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyInternetServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyInternetServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyInternetServiceSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyInternetServiceSrcCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyInternetServiceSrcCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyInternetServiceSrcGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyInternetServiceSrcId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyInternetServiceSrcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyIpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyPerIpShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicySchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicySrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyTos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyTosMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyTosNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyTrafficShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyTrafficShaperReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyUrlCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesGlobalHeaderShapingPolicyUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesGlobalHeaderShapingPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("app_category"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyAppCategory(d, v, "app_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-category"] = t
		}
	}

	if v, ok := d.GetOk("app_group"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyAppGroup(d, v, "app_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-group"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("class_id"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyClassId(d, v, "class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_forward"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyDiffservForward(d, v, "diffserv_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_reverse"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyDiffservReverse(d, v, "diffserv_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-reverse"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_forward"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyDiffservcodeForward(d, v, "diffservcode_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_rev"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyDiffservcodeRev(d, v, "diffservcode_rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-rev"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyDstaddr6(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyDstintf(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyGroups(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyInternetService(d, v, "internet_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyInternetServiceCustom(d, v, "internet_service_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom_group"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyInternetServiceCustomGroup(d, v, "internet_service_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_group"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyInternetServiceGroup(d, v, "internet_service_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_id"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyInternetServiceId(d, v, "internet_service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_name"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyInternetServiceName(d, v, "internet_service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyInternetServiceSrc(d, v, "internet_service_src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyInternetServiceSrcCustom(d, v, "internet_service_src_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom_group"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyInternetServiceSrcCustomGroup(d, v, "internet_service_src_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_group"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyInternetServiceSrcGroup(d, v, "internet_service_src_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_id"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyInternetServiceSrcId(d, v, "internet_service_src_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_name"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyInternetServiceSrcName(d, v, "internet_service_src_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-name"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyIpVersion(d, v, "ip_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("per_ip_shaper"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyPerIpShaper(d, v, "per_ip_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-ip-shaper"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicySchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicySrcaddr6(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicySrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyTos(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyTosMask(d, v, "tos_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	if v, ok := d.GetOk("tos_negate"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyTosNegate(d, v, "tos_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-negate"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyTrafficShaper(d, v, "traffic_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper_reverse"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyTrafficShaperReverse(d, v, "traffic_shaper_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper-reverse"] = t
		}
	}

	if v, ok := d.GetOk("url_category"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyUrlCategory(d, v, "url_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-category"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyUsers(d, v, "users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandPackagesGlobalHeaderShapingPolicyUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
