// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Packages WebProxyPacPolicy

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePackagesWebProxyPacPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesWebProxyPacPolicyCreate,
		Read:   resourcePackagesWebProxyPacPolicyRead,
		Update: resourcePackagesWebProxyPacPolicyUpdate,
		Delete: resourcePackagesWebProxyPacPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"pkg_folder_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"pkg": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"pac_file_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pac_file_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourcePackagesWebProxyPacPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	obj, err := getObjectPackagesWebProxyPacPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesWebProxyPacPolicy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("policyid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadPackagesWebProxyPacPolicy(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdatePackagesWebProxyPacPolicy(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating PackagesWebProxyPacPolicy resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreatePackagesWebProxyPacPolicy(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating PackagesWebProxyPacPolicy resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesWebProxyPacPolicyRead(d, m)
}

func resourcePackagesWebProxyPacPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	obj, err := getObjectPackagesWebProxyPacPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesWebProxyPacPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdatePackagesWebProxyPacPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating PackagesWebProxyPacPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesWebProxyPacPolicyRead(d, m)
}

func resourcePackagesWebProxyPacPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	wsParams["adom"] = adomv

	err = c.DeletePackagesWebProxyPacPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesWebProxyPacPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesWebProxyPacPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	if pkg_folder_path == "" {
		pkg_folder_path = importOptionChecking(m.(*FortiClient).Cfg, "pkg_folder_path")
		if err = d.Set("pkg_folder_path", pkg_folder_path); err != nil {
			return fmt.Errorf("Error set params pkg_folder_path: %v", err)
		}
	}
	if pkg == "" {
		pkg = importOptionChecking(m.(*FortiClient).Cfg, "pkg")
		if pkg == "" {
			return fmt.Errorf("Parameter pkg is missing")
		}
		if err = d.Set("pkg", pkg); err != nil {
			return fmt.Errorf("Error set params pkg: %v", err)
		}
	}
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg

	o, err := c.ReadPackagesWebProxyPacPolicy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading PackagesWebProxyPacPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesWebProxyPacPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesWebProxyPacPolicy resource from API: %v", err)
	}
	return nil
}

func flattenPackagesWebProxyPacPolicyComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesWebProxyPacPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesWebProxyPacPolicyPacFileData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesWebProxyPacPolicyPacFileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesWebProxyPacPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesWebProxyPacPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesWebProxyPacPolicySrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesWebProxyPacPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesWebProxyPacPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comments", flattenPackagesWebProxyPacPolicyComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "PackagesWebProxyPacPolicy-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenPackagesWebProxyPacPolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "PackagesWebProxyPacPolicy-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("pac_file_data", flattenPackagesWebProxyPacPolicyPacFileData(o["pac-file-data"], d, "pac_file_data")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-data"], "PackagesWebProxyPacPolicy-PacFileData"); ok {
			if err = d.Set("pac_file_data", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_data: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_data: %v", err)
		}
	}

	if err = d.Set("pac_file_name", flattenPackagesWebProxyPacPolicyPacFileName(o["pac-file-name"], d, "pac_file_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-file-name"], "PackagesWebProxyPacPolicy-PacFileName"); ok {
			if err = d.Set("pac_file_name", vv); err != nil {
				return fmt.Errorf("Error reading pac_file_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_file_name: %v", err)
		}
	}

	if err = d.Set("policyid", flattenPackagesWebProxyPacPolicyPolicyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "PackagesWebProxyPacPolicy-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenPackagesWebProxyPacPolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "PackagesWebProxyPacPolicy-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenPackagesWebProxyPacPolicySrcaddr6(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "PackagesWebProxyPacPolicy-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesWebProxyPacPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesWebProxyPacPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenPackagesWebProxyPacPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesWebProxyPacPolicyComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesWebProxyPacPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesWebProxyPacPolicyPacFileData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesWebProxyPacPolicyPacFileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesWebProxyPacPolicyPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesWebProxyPacPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesWebProxyPacPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandPackagesWebProxyPacPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesWebProxyPacPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandPackagesWebProxyPacPolicyComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandPackagesWebProxyPacPolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_data"); ok || d.HasChange("pac_file_data") {
		t, err := expandPackagesWebProxyPacPolicyPacFileData(d, v, "pac_file_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-data"] = t
		}
	}

	if v, ok := d.GetOk("pac_file_name"); ok || d.HasChange("pac_file_name") {
		t, err := expandPackagesWebProxyPacPolicyPacFileName(d, v, "pac_file_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-file-name"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandPackagesWebProxyPacPolicyPolicyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandPackagesWebProxyPacPolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandPackagesWebProxyPacPolicySrcaddr6(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandPackagesWebProxyPacPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
