// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure queue assignment on NP7.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuNpQueues() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuNpQueuesUpdate,
		Read:   resourceObjectSystemNpuNpQueuesRead,
		Update: resourceObjectSystemNpuNpQueuesUpdate,
		Delete: resourceObjectSystemNpuNpQueuesDelete,

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
			"ethernet_type": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"queue": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ip_protocol": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"queue": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ip_service": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"queue": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"sport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"profile": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cos0": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos5": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos7": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp0": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp10": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp11": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp12": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp13": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp14": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp15": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp16": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp17": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp18": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp19": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp20": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp21": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp22": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp23": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp24": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp25": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp26": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp27": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp28": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp29": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp30": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp31": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp32": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp33": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp34": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp35": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp36": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp37": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp38": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp39": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp40": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp41": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp42": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp43": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp44": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp45": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp46": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp47": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp48": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp49": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp5": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp50": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp51": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp52": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp53": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp54": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp55": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp56": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp57": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp58": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp59": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp60": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp61": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp62": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp63": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp7": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp8": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp9": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"scheduler": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
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

func resourceObjectSystemNpuNpQueuesUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuNpQueues(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpQueues resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemNpuNpQueues(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpQueues resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuNpQueues")

	return resourceObjectSystemNpuNpQueuesRead(d, m)
}

func resourceObjectSystemNpuNpQueuesDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectSystemNpuNpQueues(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuNpQueues resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuNpQueuesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuNpQueues(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpQueues resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuNpQueues(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpQueues resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuNpQueuesEthernetType2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemNpuNpQueuesEthernetTypeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := i["queue"]; ok {
			v := flattenObjectSystemNpuNpQueuesEthernetTypeQueue2edl(i["queue"], d, pre_append)
			tmp["queue"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Queue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectSystemNpuNpQueuesEthernetTypeType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesEthernetTypeWeight2edl(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesEthernetTypeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesEthernetTypeQueue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesEthernetTypeType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesEthernetTypeWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocol2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpProtocolName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpProtocolProtocol2edl(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := i["queue"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpProtocolQueue2edl(i["queue"], d, pre_append)
			tmp["queue"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Queue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpProtocolWeight2edl(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesIpProtocolName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocolProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocolQueue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocolWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpService2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dport"
		if _, ok := i["dport"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceDport2edl(i["dport"], d, pre_append)
			tmp["dport"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Dport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceProtocol2edl(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := i["queue"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceQueue2edl(i["queue"], d, pre_append)
			tmp["queue"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Queue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sport"
		if _, ok := i["sport"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceSport2edl(i["sport"], d, pre_append)
			tmp["sport"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Sport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceWeight2edl(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesIpServiceDport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceQueue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceSport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfile2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos0"
		if _, ok := i["cos0"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos02edl(i["cos0"], d, pre_append)
			tmp["cos0"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos0")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos1"
		if _, ok := i["cos1"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos12edl(i["cos1"], d, pre_append)
			tmp["cos1"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos2"
		if _, ok := i["cos2"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos22edl(i["cos2"], d, pre_append)
			tmp["cos2"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos3"
		if _, ok := i["cos3"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos32edl(i["cos3"], d, pre_append)
			tmp["cos3"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos4"
		if _, ok := i["cos4"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos42edl(i["cos4"], d, pre_append)
			tmp["cos4"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos5"
		if _, ok := i["cos5"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos52edl(i["cos5"], d, pre_append)
			tmp["cos5"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos5")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos6"
		if _, ok := i["cos6"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos62edl(i["cos6"], d, pre_append)
			tmp["cos6"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos7"
		if _, ok := i["cos7"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos72edl(i["cos7"], d, pre_append)
			tmp["cos7"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos7")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp0"
		if _, ok := i["dscp0"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp02edl(i["dscp0"], d, pre_append)
			tmp["dscp0"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp0")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp1"
		if _, ok := i["dscp1"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp12edl(i["dscp1"], d, pre_append)
			tmp["dscp1"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp10"
		if _, ok := i["dscp10"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp102edl(i["dscp10"], d, pre_append)
			tmp["dscp10"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp10")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp11"
		if _, ok := i["dscp11"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp112edl(i["dscp11"], d, pre_append)
			tmp["dscp11"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp11")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp12"
		if _, ok := i["dscp12"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp122edl(i["dscp12"], d, pre_append)
			tmp["dscp12"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp12")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp13"
		if _, ok := i["dscp13"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp132edl(i["dscp13"], d, pre_append)
			tmp["dscp13"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp13")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp14"
		if _, ok := i["dscp14"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp142edl(i["dscp14"], d, pre_append)
			tmp["dscp14"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp14")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp15"
		if _, ok := i["dscp15"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp152edl(i["dscp15"], d, pre_append)
			tmp["dscp15"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp15")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp16"
		if _, ok := i["dscp16"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp162edl(i["dscp16"], d, pre_append)
			tmp["dscp16"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp16")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp17"
		if _, ok := i["dscp17"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp172edl(i["dscp17"], d, pre_append)
			tmp["dscp17"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp17")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp18"
		if _, ok := i["dscp18"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp182edl(i["dscp18"], d, pre_append)
			tmp["dscp18"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp18")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp19"
		if _, ok := i["dscp19"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp192edl(i["dscp19"], d, pre_append)
			tmp["dscp19"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp19")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp2"
		if _, ok := i["dscp2"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp22edl(i["dscp2"], d, pre_append)
			tmp["dscp2"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp20"
		if _, ok := i["dscp20"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp202edl(i["dscp20"], d, pre_append)
			tmp["dscp20"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp20")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp21"
		if _, ok := i["dscp21"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp212edl(i["dscp21"], d, pre_append)
			tmp["dscp21"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp21")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp22"
		if _, ok := i["dscp22"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp222edl(i["dscp22"], d, pre_append)
			tmp["dscp22"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp22")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp23"
		if _, ok := i["dscp23"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp232edl(i["dscp23"], d, pre_append)
			tmp["dscp23"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp23")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp24"
		if _, ok := i["dscp24"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp242edl(i["dscp24"], d, pre_append)
			tmp["dscp24"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp24")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp25"
		if _, ok := i["dscp25"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp252edl(i["dscp25"], d, pre_append)
			tmp["dscp25"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp25")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp26"
		if _, ok := i["dscp26"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp262edl(i["dscp26"], d, pre_append)
			tmp["dscp26"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp26")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp27"
		if _, ok := i["dscp27"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp272edl(i["dscp27"], d, pre_append)
			tmp["dscp27"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp27")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp28"
		if _, ok := i["dscp28"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp282edl(i["dscp28"], d, pre_append)
			tmp["dscp28"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp28")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp29"
		if _, ok := i["dscp29"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp292edl(i["dscp29"], d, pre_append)
			tmp["dscp29"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp29")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp3"
		if _, ok := i["dscp3"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp32edl(i["dscp3"], d, pre_append)
			tmp["dscp3"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp30"
		if _, ok := i["dscp30"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp302edl(i["dscp30"], d, pre_append)
			tmp["dscp30"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp30")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp31"
		if _, ok := i["dscp31"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp312edl(i["dscp31"], d, pre_append)
			tmp["dscp31"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp31")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp32"
		if _, ok := i["dscp32"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp322edl(i["dscp32"], d, pre_append)
			tmp["dscp32"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp32")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp33"
		if _, ok := i["dscp33"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp332edl(i["dscp33"], d, pre_append)
			tmp["dscp33"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp33")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp34"
		if _, ok := i["dscp34"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp342edl(i["dscp34"], d, pre_append)
			tmp["dscp34"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp34")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp35"
		if _, ok := i["dscp35"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp352edl(i["dscp35"], d, pre_append)
			tmp["dscp35"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp35")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp36"
		if _, ok := i["dscp36"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp362edl(i["dscp36"], d, pre_append)
			tmp["dscp36"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp36")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp37"
		if _, ok := i["dscp37"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp372edl(i["dscp37"], d, pre_append)
			tmp["dscp37"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp37")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp38"
		if _, ok := i["dscp38"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp382edl(i["dscp38"], d, pre_append)
			tmp["dscp38"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp38")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp39"
		if _, ok := i["dscp39"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp392edl(i["dscp39"], d, pre_append)
			tmp["dscp39"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp39")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp4"
		if _, ok := i["dscp4"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp42edl(i["dscp4"], d, pre_append)
			tmp["dscp4"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp40"
		if _, ok := i["dscp40"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp402edl(i["dscp40"], d, pre_append)
			tmp["dscp40"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp40")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp41"
		if _, ok := i["dscp41"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp412edl(i["dscp41"], d, pre_append)
			tmp["dscp41"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp41")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp42"
		if _, ok := i["dscp42"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp422edl(i["dscp42"], d, pre_append)
			tmp["dscp42"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp42")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp43"
		if _, ok := i["dscp43"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp432edl(i["dscp43"], d, pre_append)
			tmp["dscp43"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp43")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp44"
		if _, ok := i["dscp44"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp442edl(i["dscp44"], d, pre_append)
			tmp["dscp44"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp44")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp45"
		if _, ok := i["dscp45"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp452edl(i["dscp45"], d, pre_append)
			tmp["dscp45"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp45")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp46"
		if _, ok := i["dscp46"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp462edl(i["dscp46"], d, pre_append)
			tmp["dscp46"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp46")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp47"
		if _, ok := i["dscp47"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp472edl(i["dscp47"], d, pre_append)
			tmp["dscp47"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp47")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp48"
		if _, ok := i["dscp48"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp482edl(i["dscp48"], d, pre_append)
			tmp["dscp48"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp48")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp49"
		if _, ok := i["dscp49"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp492edl(i["dscp49"], d, pre_append)
			tmp["dscp49"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp49")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp5"
		if _, ok := i["dscp5"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp52edl(i["dscp5"], d, pre_append)
			tmp["dscp5"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp5")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp50"
		if _, ok := i["dscp50"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp502edl(i["dscp50"], d, pre_append)
			tmp["dscp50"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp50")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp51"
		if _, ok := i["dscp51"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp512edl(i["dscp51"], d, pre_append)
			tmp["dscp51"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp51")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp52"
		if _, ok := i["dscp52"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp522edl(i["dscp52"], d, pre_append)
			tmp["dscp52"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp52")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp53"
		if _, ok := i["dscp53"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp532edl(i["dscp53"], d, pre_append)
			tmp["dscp53"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp53")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp54"
		if _, ok := i["dscp54"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp542edl(i["dscp54"], d, pre_append)
			tmp["dscp54"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp54")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp55"
		if _, ok := i["dscp55"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp552edl(i["dscp55"], d, pre_append)
			tmp["dscp55"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp55")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp56"
		if _, ok := i["dscp56"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp562edl(i["dscp56"], d, pre_append)
			tmp["dscp56"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp56")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp57"
		if _, ok := i["dscp57"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp572edl(i["dscp57"], d, pre_append)
			tmp["dscp57"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp57")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp58"
		if _, ok := i["dscp58"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp582edl(i["dscp58"], d, pre_append)
			tmp["dscp58"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp58")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp59"
		if _, ok := i["dscp59"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp592edl(i["dscp59"], d, pre_append)
			tmp["dscp59"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp59")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp6"
		if _, ok := i["dscp6"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp62edl(i["dscp6"], d, pre_append)
			tmp["dscp6"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp60"
		if _, ok := i["dscp60"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp602edl(i["dscp60"], d, pre_append)
			tmp["dscp60"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp60")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp61"
		if _, ok := i["dscp61"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp612edl(i["dscp61"], d, pre_append)
			tmp["dscp61"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp61")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp62"
		if _, ok := i["dscp62"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp622edl(i["dscp62"], d, pre_append)
			tmp["dscp62"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp62")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp63"
		if _, ok := i["dscp63"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp632edl(i["dscp63"], d, pre_append)
			tmp["dscp63"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp63")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp7"
		if _, ok := i["dscp7"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp72edl(i["dscp7"], d, pre_append)
			tmp["dscp7"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp7")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp8"
		if _, ok := i["dscp8"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp82edl(i["dscp8"], d, pre_append)
			tmp["dscp8"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp8")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp9"
		if _, ok := i["dscp9"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp92edl(i["dscp9"], d, pre_append)
			tmp["dscp9"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp9")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileWeight2edl(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesProfileCos02edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos52edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos72edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp02edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp102edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp112edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp122edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp132edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp142edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp152edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp162edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp172edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp182edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp192edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp202edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp212edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp222edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp232edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp242edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp252edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp262edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp272edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp282edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp292edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp302edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp312edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp322edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp332edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp342edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp352edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp362edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp372edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp382edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp392edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp402edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp412edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp422edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp432edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp442edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp452edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp462edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp472edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp482edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp492edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp52edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp502edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp512edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp522edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp532edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp542edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp552edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp562edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp572edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp582edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp592edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp602edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp612edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp622edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp632edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp72edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp82edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp92edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesScheduler2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {
			v := flattenObjectSystemNpuNpQueuesSchedulerMode2edl(i["mode"], d, pre_append)
			tmp["mode"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Scheduler-Mode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemNpuNpQueuesSchedulerName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Scheduler-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesSchedulerMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesSchedulerName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuNpQueues(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("ethernet_type", flattenObjectSystemNpuNpQueuesEthernetType2edl(o["ethernet-type"], d, "ethernet_type")); err != nil {
			if vv, ok := fortiAPIPatch(o["ethernet-type"], "ObjectSystemNpuNpQueues-EthernetType"); ok {
				if err = d.Set("ethernet_type", vv); err != nil {
					return fmt.Errorf("Error reading ethernet_type: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ethernet_type: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ethernet_type"); ok {
			if err = d.Set("ethernet_type", flattenObjectSystemNpuNpQueuesEthernetType2edl(o["ethernet-type"], d, "ethernet_type")); err != nil {
				if vv, ok := fortiAPIPatch(o["ethernet-type"], "ObjectSystemNpuNpQueues-EthernetType"); ok {
					if err = d.Set("ethernet_type", vv); err != nil {
						return fmt.Errorf("Error reading ethernet_type: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ethernet_type: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ip_protocol", flattenObjectSystemNpuNpQueuesIpProtocol2edl(o["ip-protocol"], d, "ip_protocol")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-protocol"], "ObjectSystemNpuNpQueues-IpProtocol"); ok {
				if err = d.Set("ip_protocol", vv); err != nil {
					return fmt.Errorf("Error reading ip_protocol: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_protocol: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_protocol"); ok {
			if err = d.Set("ip_protocol", flattenObjectSystemNpuNpQueuesIpProtocol2edl(o["ip-protocol"], d, "ip_protocol")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-protocol"], "ObjectSystemNpuNpQueues-IpProtocol"); ok {
					if err = d.Set("ip_protocol", vv); err != nil {
						return fmt.Errorf("Error reading ip_protocol: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_protocol: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ip_service", flattenObjectSystemNpuNpQueuesIpService2edl(o["ip-service"], d, "ip_service")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-service"], "ObjectSystemNpuNpQueues-IpService"); ok {
				if err = d.Set("ip_service", vv); err != nil {
					return fmt.Errorf("Error reading ip_service: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_service"); ok {
			if err = d.Set("ip_service", flattenObjectSystemNpuNpQueuesIpService2edl(o["ip-service"], d, "ip_service")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-service"], "ObjectSystemNpuNpQueues-IpService"); ok {
					if err = d.Set("ip_service", vv); err != nil {
						return fmt.Errorf("Error reading ip_service: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_service: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("profile", flattenObjectSystemNpuNpQueuesProfile2edl(o["profile"], d, "profile")); err != nil {
			if vv, ok := fortiAPIPatch(o["profile"], "ObjectSystemNpuNpQueues-Profile"); ok {
				if err = d.Set("profile", vv); err != nil {
					return fmt.Errorf("Error reading profile: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading profile: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("profile"); ok {
			if err = d.Set("profile", flattenObjectSystemNpuNpQueuesProfile2edl(o["profile"], d, "profile")); err != nil {
				if vv, ok := fortiAPIPatch(o["profile"], "ObjectSystemNpuNpQueues-Profile"); ok {
					if err = d.Set("profile", vv); err != nil {
						return fmt.Errorf("Error reading profile: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading profile: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("scheduler", flattenObjectSystemNpuNpQueuesScheduler2edl(o["scheduler"], d, "scheduler")); err != nil {
			if vv, ok := fortiAPIPatch(o["scheduler"], "ObjectSystemNpuNpQueues-Scheduler"); ok {
				if err = d.Set("scheduler", vv); err != nil {
					return fmt.Errorf("Error reading scheduler: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading scheduler: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("scheduler"); ok {
			if err = d.Set("scheduler", flattenObjectSystemNpuNpQueuesScheduler2edl(o["scheduler"], d, "scheduler")); err != nil {
				if vv, ok := fortiAPIPatch(o["scheduler"], "ObjectSystemNpuNpQueues-Scheduler"); ok {
					if err = d.Set("scheduler", vv); err != nil {
						return fmt.Errorf("Error reading scheduler: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading scheduler: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectSystemNpuNpQueuesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuNpQueuesEthernetType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSystemNpuNpQueuesEthernetTypeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["queue"], _ = expandObjectSystemNpuNpQueuesEthernetTypeQueue2edl(d, i["queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectSystemNpuNpQueuesEthernetTypeType2edl(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesEthernetTypeWeight2edl(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeQueue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSystemNpuNpQueuesIpProtocolName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandObjectSystemNpuNpQueuesIpProtocolProtocol2edl(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["queue"], _ = expandObjectSystemNpuNpQueuesIpProtocolQueue2edl(d, i["queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesIpProtocolWeight2edl(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolQueue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dport"], _ = expandObjectSystemNpuNpQueuesIpServiceDport2edl(d, i["dport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSystemNpuNpQueuesIpServiceName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandObjectSystemNpuNpQueuesIpServiceProtocol2edl(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["queue"], _ = expandObjectSystemNpuNpQueuesIpServiceQueue2edl(d, i["queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sport"], _ = expandObjectSystemNpuNpQueuesIpServiceSport2edl(d, i["sport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesIpServiceWeight2edl(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesIpServiceDport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceQueue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceSport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos0"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos0"], _ = expandObjectSystemNpuNpQueuesProfileCos02edl(d, i["cos0"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos1"], _ = expandObjectSystemNpuNpQueuesProfileCos12edl(d, i["cos1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos2"], _ = expandObjectSystemNpuNpQueuesProfileCos22edl(d, i["cos2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos3"], _ = expandObjectSystemNpuNpQueuesProfileCos32edl(d, i["cos3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos4"], _ = expandObjectSystemNpuNpQueuesProfileCos42edl(d, i["cos4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos5"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos5"], _ = expandObjectSystemNpuNpQueuesProfileCos52edl(d, i["cos5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos6"], _ = expandObjectSystemNpuNpQueuesProfileCos62edl(d, i["cos6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos7"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos7"], _ = expandObjectSystemNpuNpQueuesProfileCos72edl(d, i["cos7"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp0"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp0"], _ = expandObjectSystemNpuNpQueuesProfileDscp02edl(d, i["dscp0"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp1"], _ = expandObjectSystemNpuNpQueuesProfileDscp12edl(d, i["dscp1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp10"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp10"], _ = expandObjectSystemNpuNpQueuesProfileDscp102edl(d, i["dscp10"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp11"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp11"], _ = expandObjectSystemNpuNpQueuesProfileDscp112edl(d, i["dscp11"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp12"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp12"], _ = expandObjectSystemNpuNpQueuesProfileDscp122edl(d, i["dscp12"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp13"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp13"], _ = expandObjectSystemNpuNpQueuesProfileDscp132edl(d, i["dscp13"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp14"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp14"], _ = expandObjectSystemNpuNpQueuesProfileDscp142edl(d, i["dscp14"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp15"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp15"], _ = expandObjectSystemNpuNpQueuesProfileDscp152edl(d, i["dscp15"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp16"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp16"], _ = expandObjectSystemNpuNpQueuesProfileDscp162edl(d, i["dscp16"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp17"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp17"], _ = expandObjectSystemNpuNpQueuesProfileDscp172edl(d, i["dscp17"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp18"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp18"], _ = expandObjectSystemNpuNpQueuesProfileDscp182edl(d, i["dscp18"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp19"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp19"], _ = expandObjectSystemNpuNpQueuesProfileDscp192edl(d, i["dscp19"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp2"], _ = expandObjectSystemNpuNpQueuesProfileDscp22edl(d, i["dscp2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp20"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp20"], _ = expandObjectSystemNpuNpQueuesProfileDscp202edl(d, i["dscp20"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp21"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp21"], _ = expandObjectSystemNpuNpQueuesProfileDscp212edl(d, i["dscp21"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp22"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp22"], _ = expandObjectSystemNpuNpQueuesProfileDscp222edl(d, i["dscp22"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp23"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp23"], _ = expandObjectSystemNpuNpQueuesProfileDscp232edl(d, i["dscp23"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp24"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp24"], _ = expandObjectSystemNpuNpQueuesProfileDscp242edl(d, i["dscp24"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp25"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp25"], _ = expandObjectSystemNpuNpQueuesProfileDscp252edl(d, i["dscp25"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp26"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp26"], _ = expandObjectSystemNpuNpQueuesProfileDscp262edl(d, i["dscp26"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp27"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp27"], _ = expandObjectSystemNpuNpQueuesProfileDscp272edl(d, i["dscp27"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp28"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp28"], _ = expandObjectSystemNpuNpQueuesProfileDscp282edl(d, i["dscp28"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp29"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp29"], _ = expandObjectSystemNpuNpQueuesProfileDscp292edl(d, i["dscp29"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp3"], _ = expandObjectSystemNpuNpQueuesProfileDscp32edl(d, i["dscp3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp30"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp30"], _ = expandObjectSystemNpuNpQueuesProfileDscp302edl(d, i["dscp30"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp31"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp31"], _ = expandObjectSystemNpuNpQueuesProfileDscp312edl(d, i["dscp31"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp32"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp32"], _ = expandObjectSystemNpuNpQueuesProfileDscp322edl(d, i["dscp32"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp33"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp33"], _ = expandObjectSystemNpuNpQueuesProfileDscp332edl(d, i["dscp33"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp34"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp34"], _ = expandObjectSystemNpuNpQueuesProfileDscp342edl(d, i["dscp34"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp35"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp35"], _ = expandObjectSystemNpuNpQueuesProfileDscp352edl(d, i["dscp35"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp36"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp36"], _ = expandObjectSystemNpuNpQueuesProfileDscp362edl(d, i["dscp36"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp37"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp37"], _ = expandObjectSystemNpuNpQueuesProfileDscp372edl(d, i["dscp37"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp38"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp38"], _ = expandObjectSystemNpuNpQueuesProfileDscp382edl(d, i["dscp38"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp39"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp39"], _ = expandObjectSystemNpuNpQueuesProfileDscp392edl(d, i["dscp39"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp4"], _ = expandObjectSystemNpuNpQueuesProfileDscp42edl(d, i["dscp4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp40"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp40"], _ = expandObjectSystemNpuNpQueuesProfileDscp402edl(d, i["dscp40"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp41"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp41"], _ = expandObjectSystemNpuNpQueuesProfileDscp412edl(d, i["dscp41"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp42"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp42"], _ = expandObjectSystemNpuNpQueuesProfileDscp422edl(d, i["dscp42"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp43"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp43"], _ = expandObjectSystemNpuNpQueuesProfileDscp432edl(d, i["dscp43"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp44"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp44"], _ = expandObjectSystemNpuNpQueuesProfileDscp442edl(d, i["dscp44"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp45"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp45"], _ = expandObjectSystemNpuNpQueuesProfileDscp452edl(d, i["dscp45"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp46"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp46"], _ = expandObjectSystemNpuNpQueuesProfileDscp462edl(d, i["dscp46"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp47"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp47"], _ = expandObjectSystemNpuNpQueuesProfileDscp472edl(d, i["dscp47"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp48"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp48"], _ = expandObjectSystemNpuNpQueuesProfileDscp482edl(d, i["dscp48"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp49"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp49"], _ = expandObjectSystemNpuNpQueuesProfileDscp492edl(d, i["dscp49"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp5"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp5"], _ = expandObjectSystemNpuNpQueuesProfileDscp52edl(d, i["dscp5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp50"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp50"], _ = expandObjectSystemNpuNpQueuesProfileDscp502edl(d, i["dscp50"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp51"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp51"], _ = expandObjectSystemNpuNpQueuesProfileDscp512edl(d, i["dscp51"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp52"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp52"], _ = expandObjectSystemNpuNpQueuesProfileDscp522edl(d, i["dscp52"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp53"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp53"], _ = expandObjectSystemNpuNpQueuesProfileDscp532edl(d, i["dscp53"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp54"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp54"], _ = expandObjectSystemNpuNpQueuesProfileDscp542edl(d, i["dscp54"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp55"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp55"], _ = expandObjectSystemNpuNpQueuesProfileDscp552edl(d, i["dscp55"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp56"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp56"], _ = expandObjectSystemNpuNpQueuesProfileDscp562edl(d, i["dscp56"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp57"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp57"], _ = expandObjectSystemNpuNpQueuesProfileDscp572edl(d, i["dscp57"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp58"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp58"], _ = expandObjectSystemNpuNpQueuesProfileDscp582edl(d, i["dscp58"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp59"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp59"], _ = expandObjectSystemNpuNpQueuesProfileDscp592edl(d, i["dscp59"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp6"], _ = expandObjectSystemNpuNpQueuesProfileDscp62edl(d, i["dscp6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp60"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp60"], _ = expandObjectSystemNpuNpQueuesProfileDscp602edl(d, i["dscp60"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp61"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp61"], _ = expandObjectSystemNpuNpQueuesProfileDscp612edl(d, i["dscp61"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp62"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp62"], _ = expandObjectSystemNpuNpQueuesProfileDscp622edl(d, i["dscp62"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp63"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp63"], _ = expandObjectSystemNpuNpQueuesProfileDscp632edl(d, i["dscp63"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp7"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp7"], _ = expandObjectSystemNpuNpQueuesProfileDscp72edl(d, i["dscp7"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp8"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp8"], _ = expandObjectSystemNpuNpQueuesProfileDscp82edl(d, i["dscp8"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp9"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp9"], _ = expandObjectSystemNpuNpQueuesProfileDscp92edl(d, i["dscp9"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectSystemNpuNpQueuesProfileId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectSystemNpuNpQueuesProfileType2edl(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesProfileWeight2edl(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesProfileCos02edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos52edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos72edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp02edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp102edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp112edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp122edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp132edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp142edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp152edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp162edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp172edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp182edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp192edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp202edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp212edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp222edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp232edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp242edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp252edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp262edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp272edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp282edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp292edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp302edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp312edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp322edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp332edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp342edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp352edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp362edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp372edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp382edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp392edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp402edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp412edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp422edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp432edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp442edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp452edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp462edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp472edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp482edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp492edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp52edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp502edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp512edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp522edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp532edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp542edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp552edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp562edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp572edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp582edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp592edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp602edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp612edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp622edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp632edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp72edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp82edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp92edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesScheduler2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mode"], _ = expandObjectSystemNpuNpQueuesSchedulerMode2edl(d, i["mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSystemNpuNpQueuesSchedulerName2edl(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesSchedulerMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesSchedulerName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuNpQueues(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ethernet_type"); ok || d.HasChange("ethernet_type") {
		t, err := expandObjectSystemNpuNpQueuesEthernetType2edl(d, v, "ethernet_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ethernet-type"] = t
		}
	}

	if v, ok := d.GetOk("ip_protocol"); ok || d.HasChange("ip_protocol") {
		t, err := expandObjectSystemNpuNpQueuesIpProtocol2edl(d, v, "ip_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-protocol"] = t
		}
	}

	if v, ok := d.GetOk("ip_service"); ok || d.HasChange("ip_service") {
		t, err := expandObjectSystemNpuNpQueuesIpService2edl(d, v, "ip_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-service"] = t
		}
	}

	if v, ok := d.GetOk("profile"); ok || d.HasChange("profile") {
		t, err := expandObjectSystemNpuNpQueuesProfile2edl(d, v, "profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile"] = t
		}
	}

	if v, ok := d.GetOk("scheduler"); ok || d.HasChange("scheduler") {
		t, err := expandObjectSystemNpuNpQueuesScheduler2edl(d, v, "scheduler")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scheduler"] = t
		}
	}

	return &obj, nil
}
