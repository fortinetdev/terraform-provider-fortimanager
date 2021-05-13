// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure priorities for FortiGate units accessing antivirus updates and web filtering services.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFmupdateServerAccessPriorities() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateServerAccessPrioritiesUpdate,
		Read:   resourceFmupdateServerAccessPrioritiesRead,
		Update: resourceFmupdateServerAccessPrioritiesUpdate,
		Delete: resourceFmupdateServerAccessPrioritiesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"access_public": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_ips": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"private_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"time_zone": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"web_spam": &schema.Schema{
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

func resourceFmupdateServerAccessPrioritiesUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateServerAccessPriorities(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateServerAccessPriorities resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateServerAccessPriorities(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateServerAccessPriorities resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateServerAccessPriorities")

	return resourceFmupdateServerAccessPrioritiesRead(d, m)
}

func resourceFmupdateServerAccessPrioritiesDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateServerAccessPriorities(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateServerAccessPriorities resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateServerAccessPrioritiesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateServerAccessPriorities(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateServerAccessPriorities resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateServerAccessPriorities(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateServerAccessPriorities resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateServerAccessPrioritiesAccessPublicFsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenFmupdateServerAccessPrioritiesAvIpsFsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenFmupdateServerAccessPrioritiesPrivateServerFsa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFmupdateServerAccessPrioritiesPrivateServerIdFsa(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FmupdateServerAccessPriorities-PrivateServer-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFmupdateServerAccessPrioritiesPrivateServerIpFsa(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FmupdateServerAccessPriorities-PrivateServer-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenFmupdateServerAccessPrioritiesPrivateServerIp6Fsa(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "FmupdateServerAccessPriorities-PrivateServer-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time_zone"
		if _, ok := i["time_zone"]; ok {
			v := flattenFmupdateServerAccessPrioritiesPrivateServerTimeZoneFsa(i["time_zone"], d, pre_append)
			tmp["time_zone"] = fortiAPISubPartPatch(v, "FmupdateServerAccessPriorities-PrivateServer-TimeZone")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFmupdateServerAccessPrioritiesPrivateServerIdFsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServerAccessPrioritiesPrivateServerIpFsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServerAccessPrioritiesPrivateServerIp6Fsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServerAccessPrioritiesPrivateServerTimeZoneFsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServerAccessPrioritiesWebSpamFsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func refreshObjectFmupdateServerAccessPriorities(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("access_public", flattenFmupdateServerAccessPrioritiesAccessPublicFsa(o["access-public"], d, "access_public")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-public"], "FmupdateServerAccessPriorities-AccessPublic"); ok {
			if err = d.Set("access_public", vv); err != nil {
				return fmt.Errorf("Error reading access_public: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_public: %v", err)
		}
	}

	if err = d.Set("av_ips", flattenFmupdateServerAccessPrioritiesAvIpsFsa(o["av-ips"], d, "av_ips")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-ips"], "FmupdateServerAccessPriorities-AvIps"); ok {
			if err = d.Set("av_ips", vv); err != nil {
				return fmt.Errorf("Error reading av_ips: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_ips: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("private_server", flattenFmupdateServerAccessPrioritiesPrivateServerFsa(o["private-server"], d, "private_server")); err != nil {
			if vv, ok := fortiAPIPatch(o["private-server"], "FmupdateServerAccessPriorities-PrivateServer"); ok {
				if err = d.Set("private_server", vv); err != nil {
					return fmt.Errorf("Error reading private_server: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading private_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("private_server"); ok {
			if err = d.Set("private_server", flattenFmupdateServerAccessPrioritiesPrivateServerFsa(o["private-server"], d, "private_server")); err != nil {
				if vv, ok := fortiAPIPatch(o["private-server"], "FmupdateServerAccessPriorities-PrivateServer"); ok {
					if err = d.Set("private_server", vv); err != nil {
						return fmt.Errorf("Error reading private_server: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading private_server: %v", err)
				}
			}
		}
	}

	if err = d.Set("web_spam", flattenFmupdateServerAccessPrioritiesWebSpamFsa(o["web-spam"], d, "web_spam")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-spam"], "FmupdateServerAccessPriorities-WebSpam"); ok {
			if err = d.Set("web_spam", vv); err != nil {
				return fmt.Errorf("Error reading web_spam: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_spam: %v", err)
		}
	}

	return nil
}

func flattenFmupdateServerAccessPrioritiesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateServerAccessPrioritiesAccessPublicFsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServerAccessPrioritiesAvIpsFsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServerAccessPrioritiesPrivateServerFsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandFmupdateServerAccessPrioritiesPrivateServerIdFsa(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandFmupdateServerAccessPrioritiesPrivateServerIpFsa(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip6"], _ = expandFmupdateServerAccessPrioritiesPrivateServerIp6Fsa(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time_zone"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["time_zone"], _ = expandFmupdateServerAccessPrioritiesPrivateServerTimeZoneFsa(d, i["time_zone"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFmupdateServerAccessPrioritiesPrivateServerIdFsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServerAccessPrioritiesPrivateServerIpFsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServerAccessPrioritiesPrivateServerIp6Fsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServerAccessPrioritiesPrivateServerTimeZoneFsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServerAccessPrioritiesWebSpamFsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateServerAccessPriorities(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_public"); ok {
		t, err := expandFmupdateServerAccessPrioritiesAccessPublicFsa(d, v, "access_public")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-public"] = t
		}
	}

	if v, ok := d.GetOk("av_ips"); ok {
		t, err := expandFmupdateServerAccessPrioritiesAvIpsFsa(d, v, "av_ips")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-ips"] = t
		}
	}

	if v, ok := d.GetOk("private_server"); ok {
		t, err := expandFmupdateServerAccessPrioritiesPrivateServerFsa(d, v, "private_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-server"] = t
		}
	}

	if v, ok := d.GetOk("web_spam"); ok {
		t, err := expandFmupdateServerAccessPrioritiesWebSpamFsa(d, v, "web_spam")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-spam"] = t
		}
	}

	return &obj, nil
}
