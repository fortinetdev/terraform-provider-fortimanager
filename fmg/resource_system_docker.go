// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Docker host.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemDocker() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDockerUpdate,
		Read:   resourceSystemDockerRead,
		Update: resourceSystemDockerUpdate,
		Delete: resourceSystemDockerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"cpu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"default_address_pool_base": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"default_address_pool_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"docker_user_login_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fortiaiops": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiauthenticator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiportal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortisigconverter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortisoar": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiwlm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsmcollector": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mem": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"policyanalyzer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdwancontroller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"universalconnector": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDockerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemDocker(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDocker resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDocker(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemDocker resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemDocker")

	return resourceSystemDockerRead(d, m)
}

func resourceSystemDockerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemDocker(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDocker resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDockerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemDocker(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemDocker resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDocker(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDocker resource from API: %v", err)
	}
	return nil
}

func flattenSystemDockerCpu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDockerDefaultAddressPoolBase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDockerDefaultAddressPoolSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDockerDockerUserLoginMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDockerFortiaiops(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDockerFortiauthenticator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDockerFortiportal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDockerFortisigconverter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDockerFortisoar(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDockerFortiwlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDockerFsmcollector(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDockerMem(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDockerPolicyanalyzer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDockerSdwancontroller(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDockerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDockerUniversalconnector(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDocker(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cpu", flattenSystemDockerCpu(o["cpu"], d, "cpu")); err != nil {
		if vv, ok := fortiAPIPatch(o["cpu"], "SystemDocker-Cpu"); ok {
			if err = d.Set("cpu", vv); err != nil {
				return fmt.Errorf("Error reading cpu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cpu: %v", err)
		}
	}

	if err = d.Set("default_address_pool_base", flattenSystemDockerDefaultAddressPoolBase(o["default-address-pool_base"], d, "default_address_pool_base")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-address-pool_base"], "SystemDocker-DefaultAddressPoolBase"); ok {
			if err = d.Set("default_address_pool_base", vv); err != nil {
				return fmt.Errorf("Error reading default_address_pool_base: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_address_pool_base: %v", err)
		}
	}

	if err = d.Set("default_address_pool_size", flattenSystemDockerDefaultAddressPoolSize(o["default-address-pool_size"], d, "default_address_pool_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-address-pool_size"], "SystemDocker-DefaultAddressPoolSize"); ok {
			if err = d.Set("default_address_pool_size", vv); err != nil {
				return fmt.Errorf("Error reading default_address_pool_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_address_pool_size: %v", err)
		}
	}

	if err = d.Set("docker_user_login_max", flattenSystemDockerDockerUserLoginMax(o["docker-user-login-max"], d, "docker_user_login_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["docker-user-login-max"], "SystemDocker-DockerUserLoginMax"); ok {
			if err = d.Set("docker_user_login_max", vv); err != nil {
				return fmt.Errorf("Error reading docker_user_login_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading docker_user_login_max: %v", err)
		}
	}

	if err = d.Set("fortiaiops", flattenSystemDockerFortiaiops(o["fortiaiops"], d, "fortiaiops")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiaiops"], "SystemDocker-Fortiaiops"); ok {
			if err = d.Set("fortiaiops", vv); err != nil {
				return fmt.Errorf("Error reading fortiaiops: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiaiops: %v", err)
		}
	}

	if err = d.Set("fortiauthenticator", flattenSystemDockerFortiauthenticator(o["fortiauthenticator"], d, "fortiauthenticator")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiauthenticator"], "SystemDocker-Fortiauthenticator"); ok {
			if err = d.Set("fortiauthenticator", vv); err != nil {
				return fmt.Errorf("Error reading fortiauthenticator: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiauthenticator: %v", err)
		}
	}

	if err = d.Set("fortiportal", flattenSystemDockerFortiportal(o["fortiportal"], d, "fortiportal")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiportal"], "SystemDocker-Fortiportal"); ok {
			if err = d.Set("fortiportal", vv); err != nil {
				return fmt.Errorf("Error reading fortiportal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiportal: %v", err)
		}
	}

	if err = d.Set("fortisigconverter", flattenSystemDockerFortisigconverter(o["fortisigconverter"], d, "fortisigconverter")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortisigconverter"], "SystemDocker-Fortisigconverter"); ok {
			if err = d.Set("fortisigconverter", vv); err != nil {
				return fmt.Errorf("Error reading fortisigconverter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortisigconverter: %v", err)
		}
	}

	if err = d.Set("fortisoar", flattenSystemDockerFortisoar(o["fortisoar"], d, "fortisoar")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortisoar"], "SystemDocker-Fortisoar"); ok {
			if err = d.Set("fortisoar", vv); err != nil {
				return fmt.Errorf("Error reading fortisoar: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortisoar: %v", err)
		}
	}

	if err = d.Set("fortiwlm", flattenSystemDockerFortiwlm(o["fortiwlm"], d, "fortiwlm")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiwlm"], "SystemDocker-Fortiwlm"); ok {
			if err = d.Set("fortiwlm", vv); err != nil {
				return fmt.Errorf("Error reading fortiwlm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiwlm: %v", err)
		}
	}

	if err = d.Set("fsmcollector", flattenSystemDockerFsmcollector(o["fsmcollector"], d, "fsmcollector")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsmcollector"], "SystemDocker-Fsmcollector"); ok {
			if err = d.Set("fsmcollector", vv); err != nil {
				return fmt.Errorf("Error reading fsmcollector: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsmcollector: %v", err)
		}
	}

	if err = d.Set("mem", flattenSystemDockerMem(o["mem"], d, "mem")); err != nil {
		if vv, ok := fortiAPIPatch(o["mem"], "SystemDocker-Mem"); ok {
			if err = d.Set("mem", vv); err != nil {
				return fmt.Errorf("Error reading mem: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mem: %v", err)
		}
	}

	if err = d.Set("policyanalyzer", flattenSystemDockerPolicyanalyzer(o["policyanalyzer"], d, "policyanalyzer")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyanalyzer"], "SystemDocker-Policyanalyzer"); ok {
			if err = d.Set("policyanalyzer", vv); err != nil {
				return fmt.Errorf("Error reading policyanalyzer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyanalyzer: %v", err)
		}
	}

	if err = d.Set("sdwancontroller", flattenSystemDockerSdwancontroller(o["sdwancontroller"], d, "sdwancontroller")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdwancontroller"], "SystemDocker-Sdwancontroller"); ok {
			if err = d.Set("sdwancontroller", vv); err != nil {
				return fmt.Errorf("Error reading sdwancontroller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdwancontroller: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemDockerStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemDocker-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("universalconnector", flattenSystemDockerUniversalconnector(o["universalconnector"], d, "universalconnector")); err != nil {
		if vv, ok := fortiAPIPatch(o["universalconnector"], "SystemDocker-Universalconnector"); ok {
			if err = d.Set("universalconnector", vv); err != nil {
				return fmt.Errorf("Error reading universalconnector: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading universalconnector: %v", err)
		}
	}

	return nil
}

func flattenSystemDockerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDockerCpu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDockerDefaultAddressPoolBase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemDockerDefaultAddressPoolSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDockerDockerUserLoginMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDockerFortiaiops(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDockerFortiauthenticator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDockerFortiportal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDockerFortisigconverter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDockerFortisoar(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDockerFortiwlm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDockerFsmcollector(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDockerMem(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDockerPolicyanalyzer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDockerSdwancontroller(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDockerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDockerUniversalconnector(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDocker(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cpu"); ok {
		t, err := expandSystemDockerCpu(d, v, "cpu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpu"] = t
		}
	}

	if v, ok := d.GetOk("default_address_pool_base"); ok {
		t, err := expandSystemDockerDefaultAddressPoolBase(d, v, "default_address_pool_base")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-address-pool_base"] = t
		}
	}

	if v, ok := d.GetOk("default_address_pool_size"); ok {
		t, err := expandSystemDockerDefaultAddressPoolSize(d, v, "default_address_pool_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-address-pool_size"] = t
		}
	}

	if v, ok := d.GetOk("docker_user_login_max"); ok {
		t, err := expandSystemDockerDockerUserLoginMax(d, v, "docker_user_login_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["docker-user-login-max"] = t
		}
	}

	if v, ok := d.GetOk("fortiaiops"); ok {
		t, err := expandSystemDockerFortiaiops(d, v, "fortiaiops")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiaiops"] = t
		}
	}

	if v, ok := d.GetOk("fortiauthenticator"); ok {
		t, err := expandSystemDockerFortiauthenticator(d, v, "fortiauthenticator")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiauthenticator"] = t
		}
	}

	if v, ok := d.GetOk("fortiportal"); ok {
		t, err := expandSystemDockerFortiportal(d, v, "fortiportal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiportal"] = t
		}
	}

	if v, ok := d.GetOk("fortisigconverter"); ok {
		t, err := expandSystemDockerFortisigconverter(d, v, "fortisigconverter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortisigconverter"] = t
		}
	}

	if v, ok := d.GetOk("fortisoar"); ok {
		t, err := expandSystemDockerFortisoar(d, v, "fortisoar")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortisoar"] = t
		}
	}

	if v, ok := d.GetOk("fortiwlm"); ok {
		t, err := expandSystemDockerFortiwlm(d, v, "fortiwlm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiwlm"] = t
		}
	}

	if v, ok := d.GetOk("fsmcollector"); ok {
		t, err := expandSystemDockerFsmcollector(d, v, "fsmcollector")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsmcollector"] = t
		}
	}

	if v, ok := d.GetOk("mem"); ok {
		t, err := expandSystemDockerMem(d, v, "mem")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mem"] = t
		}
	}

	if v, ok := d.GetOk("policyanalyzer"); ok {
		t, err := expandSystemDockerPolicyanalyzer(d, v, "policyanalyzer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyanalyzer"] = t
		}
	}

	if v, ok := d.GetOk("sdwancontroller"); ok {
		t, err := expandSystemDockerSdwancontroller(d, v, "sdwancontroller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdwancontroller"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemDockerStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("universalconnector"); ok {
		t, err := expandSystemDockerUniversalconnector(d, v, "universalconnector")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["universalconnector"] = t
		}
	}

	return &obj, nil
}
