// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SSH filter profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSshFilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSshFilterProfileCreate,
		Read:   resourceObjectSshFilterProfileRead,
		Update: resourceObjectSshFilterProfileUpdate,
		Delete: resourceObjectSshFilterProfileDelete,

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
			"block": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"default_command_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"shell_commands": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"alert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
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

func resourceObjectSshFilterProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSshFilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSshFilterProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSshFilterProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSshFilterProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSshFilterProfileRead(d, m)
}

func resourceObjectSshFilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSshFilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSshFilterProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSshFilterProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSshFilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSshFilterProfileRead(d, m)
}

func resourceObjectSshFilterProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSshFilterProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSshFilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSshFilterProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSshFilterProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSshFilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSshFilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSshFilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectSshFilterProfileBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:   "x11",
			2:   "shell",
			4:   "exec",
			8:   "port-forward",
			16:  "tun-forward",
			32:  "sftp",
			64:  "unknown",
			128: "scp",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectSshFilterProfileDefaultCommandLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectSshFilterProfileLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:   "x11",
			2:   "shell",
			4:   "exec",
			8:   "port-forward",
			16:  "tun-forward",
			32:  "sftp",
			64:  "unknown",
			128: "scp",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectSshFilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSshFilterProfileShellCommands(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenObjectSshFilterProfileShellCommandsAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectSshFilterProfile-ShellCommands-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := i["alert"]; ok {
			v := flattenObjectSshFilterProfileShellCommandsAlert(i["alert"], d, pre_append)
			tmp["alert"] = fortiAPISubPartPatch(v, "ObjectSshFilterProfile-ShellCommands-Alert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSshFilterProfileShellCommandsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSshFilterProfile-ShellCommands-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectSshFilterProfileShellCommandsLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectSshFilterProfile-ShellCommands-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenObjectSshFilterProfileShellCommandsPattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "ObjectSshFilterProfile-ShellCommands-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenObjectSshFilterProfileShellCommandsSeverity(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "ObjectSshFilterProfile-ShellCommands-Severity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectSshFilterProfileShellCommandsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectSshFilterProfile-ShellCommands-Type")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSshFilterProfileShellCommandsAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "block",
			1: "allow",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSshFilterProfileShellCommandsAlert(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectSshFilterProfileShellCommandsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSshFilterProfileShellCommandsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectSshFilterProfileShellCommandsPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSshFilterProfileShellCommandsSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			2: "low",
			3: "medium",
			4: "high",
			5: "critical",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSshFilterProfileShellCommandsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "regex",
			2: "simple",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectObjectSshFilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("block", flattenObjectSshFilterProfileBlock(o["block"], d, "block")); err != nil {
		if vv, ok := fortiAPIPatch(o["block"], "ObjectSshFilterProfile-Block"); ok {
			if err = d.Set("block", vv); err != nil {
				return fmt.Errorf("Error reading block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block: %v", err)
		}
	}

	if err = d.Set("default_command_log", flattenObjectSshFilterProfileDefaultCommandLog(o["default-command-log"], d, "default_command_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-command-log"], "ObjectSshFilterProfile-DefaultCommandLog"); ok {
			if err = d.Set("default_command_log", vv); err != nil {
				return fmt.Errorf("Error reading default_command_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_command_log: %v", err)
		}
	}

	if err = d.Set("log", flattenObjectSshFilterProfileLog(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ObjectSshFilterProfile-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSshFilterProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSshFilterProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("shell_commands", flattenObjectSshFilterProfileShellCommands(o["shell-commands"], d, "shell_commands")); err != nil {
			if vv, ok := fortiAPIPatch(o["shell-commands"], "ObjectSshFilterProfile-ShellCommands"); ok {
				if err = d.Set("shell_commands", vv); err != nil {
					return fmt.Errorf("Error reading shell_commands: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading shell_commands: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("shell_commands"); ok {
			if err = d.Set("shell_commands", flattenObjectSshFilterProfileShellCommands(o["shell-commands"], d, "shell_commands")); err != nil {
				if vv, ok := fortiAPIPatch(o["shell-commands"], "ObjectSshFilterProfile-ShellCommands"); ok {
					if err = d.Set("shell_commands", vv); err != nil {
						return fmt.Errorf("Error reading shell_commands: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading shell_commands: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectSshFilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSshFilterProfileBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSshFilterProfileDefaultCommandLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSshFilterProfileLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSshFilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSshFilterProfileShellCommands(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandObjectSshFilterProfileShellCommandsAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["alert"], _ = expandObjectSshFilterProfileShellCommandsAlert(d, i["alert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectSshFilterProfileShellCommandsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandObjectSshFilterProfileShellCommandsLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pattern"], _ = expandObjectSshFilterProfileShellCommandsPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["severity"], _ = expandObjectSshFilterProfileShellCommandsSeverity(d, i["severity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandObjectSshFilterProfileShellCommandsType(d, i["type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSshFilterProfileShellCommandsAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSshFilterProfileShellCommandsAlert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSshFilterProfileShellCommandsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSshFilterProfileShellCommandsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSshFilterProfileShellCommandsPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSshFilterProfileShellCommandsSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSshFilterProfileShellCommandsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSshFilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("block"); ok {
		t, err := expandObjectSshFilterProfileBlock(d, v, "block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block"] = t
		}
	}

	if v, ok := d.GetOk("default_command_log"); ok {
		t, err := expandObjectSshFilterProfileDefaultCommandLog(d, v, "default_command_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-command-log"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok {
		t, err := expandObjectSshFilterProfileLog(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectSshFilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("shell_commands"); ok {
		t, err := expandObjectSshFilterProfileShellCommands(d, v, "shell_commands")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shell-commands"] = t
		}
	}

	return &obj, nil
}
