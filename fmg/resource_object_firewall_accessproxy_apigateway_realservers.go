// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Select the real servers that this Access Proxy will distribute traffic to.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallAccessProxyApiGatewayRealservers() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallAccessProxyApiGatewayRealserversCreate,
		Read:   resourceObjectFirewallAccessProxyApiGatewayRealserversRead,
		Update: resourceObjectFirewallAccessProxyApiGatewayRealserversUpdate,
		Delete: resourceObjectFirewallAccessProxyApiGatewayRealserversDelete,

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
			"access_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"api_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"health_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"health_check_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"holddown_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mappedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssh_client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_host_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_host_key_validation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"translate_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
	}
}

func resourceObjectFirewallAccessProxyApiGatewayRealserversCreate(d *schema.ResourceData, m interface{}) error {
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

	access_proxy := d.Get("access_proxy").(string)
	api_gateway := d.Get("api_gateway").(string)
	paradict["access_proxy"] = access_proxy
	paradict["api_gateway"] = api_gateway

	obj, err := getObjectObjectFirewallAccessProxyApiGatewayRealservers(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAccessProxyApiGatewayRealservers resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFirewallAccessProxyApiGatewayRealservers(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAccessProxyApiGatewayRealservers resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallAccessProxyApiGatewayRealserversRead(d, m)
}

func resourceObjectFirewallAccessProxyApiGatewayRealserversUpdate(d *schema.ResourceData, m interface{}) error {
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

	access_proxy := d.Get("access_proxy").(string)
	api_gateway := d.Get("api_gateway").(string)
	paradict["access_proxy"] = access_proxy
	paradict["api_gateway"] = api_gateway

	obj, err := getObjectObjectFirewallAccessProxyApiGatewayRealservers(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAccessProxyApiGatewayRealservers resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallAccessProxyApiGatewayRealservers(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAccessProxyApiGatewayRealservers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallAccessProxyApiGatewayRealserversRead(d, m)
}

func resourceObjectFirewallAccessProxyApiGatewayRealserversDelete(d *schema.ResourceData, m interface{}) error {
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

	access_proxy := d.Get("access_proxy").(string)
	api_gateway := d.Get("api_gateway").(string)
	paradict["access_proxy"] = access_proxy
	paradict["api_gateway"] = api_gateway

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallAccessProxyApiGatewayRealservers(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallAccessProxyApiGatewayRealservers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallAccessProxyApiGatewayRealserversRead(d *schema.ResourceData, m interface{}) error {
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

	access_proxy := d.Get("access_proxy").(string)
	api_gateway := d.Get("api_gateway").(string)
	if access_proxy == "" {
		access_proxy = importOptionChecking(m.(*FortiClient).Cfg, "access_proxy")
		if access_proxy == "" {
			return fmt.Errorf("Parameter access_proxy is missing")
		}
		if err = d.Set("access_proxy", access_proxy); err != nil {
			return fmt.Errorf("Error set params access_proxy: %v", err)
		}
	}
	if api_gateway == "" {
		api_gateway = importOptionChecking(m.(*FortiClient).Cfg, "api_gateway")
		if api_gateway == "" {
			return fmt.Errorf("Parameter api_gateway is missing")
		}
		if err = d.Set("api_gateway", api_gateway); err != nil {
			return fmt.Errorf("Error set params api_gateway: %v", err)
		}
	}
	paradict["access_proxy"] = access_proxy
	paradict["api_gateway"] = api_gateway

	o, err := c.ReadObjectFirewallAccessProxyApiGatewayRealservers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAccessProxyApiGatewayRealservers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallAccessProxyApiGatewayRealservers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAccessProxyApiGatewayRealservers resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversAddrType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversAddress3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversDomain3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversExternalAuth3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversHealthCheck3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversHealthCheckProto3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversHolddownInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversHttpHost3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversMappedport3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversSshClientCert3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversSshHostKey3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversTranslateHost3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversTunnelEncryption3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversWeight3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallAccessProxyApiGatewayRealservers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("addr_type", flattenObjectFirewallAccessProxyApiGatewayRealserversAddrType3rdl(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "ObjectFirewallAccessProxyApiGatewayRealservers-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("address", flattenObjectFirewallAccessProxyApiGatewayRealserversAddress3rdl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "ObjectFirewallAccessProxyApiGatewayRealservers-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("domain", flattenObjectFirewallAccessProxyApiGatewayRealserversDomain3rdl(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "ObjectFirewallAccessProxyApiGatewayRealservers-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("external_auth", flattenObjectFirewallAccessProxyApiGatewayRealserversExternalAuth3rdl(o["external-auth"], d, "external_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-auth"], "ObjectFirewallAccessProxyApiGatewayRealservers-ExternalAuth"); ok {
			if err = d.Set("external_auth", vv); err != nil {
				return fmt.Errorf("Error reading external_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_auth: %v", err)
		}
	}

	if err = d.Set("health_check", flattenObjectFirewallAccessProxyApiGatewayRealserversHealthCheck3rdl(o["health-check"], d, "health_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check"], "ObjectFirewallAccessProxyApiGatewayRealservers-HealthCheck"); ok {
			if err = d.Set("health_check", vv); err != nil {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check: %v", err)
		}
	}

	if err = d.Set("health_check_proto", flattenObjectFirewallAccessProxyApiGatewayRealserversHealthCheckProto3rdl(o["health-check-proto"], d, "health_check_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check-proto"], "ObjectFirewallAccessProxyApiGatewayRealservers-HealthCheckProto"); ok {
			if err = d.Set("health_check_proto", vv); err != nil {
				return fmt.Errorf("Error reading health_check_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check_proto: %v", err)
		}
	}

	if err = d.Set("holddown_interval", flattenObjectFirewallAccessProxyApiGatewayRealserversHolddownInterval3rdl(o["holddown-interval"], d, "holddown_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["holddown-interval"], "ObjectFirewallAccessProxyApiGatewayRealservers-HolddownInterval"); ok {
			if err = d.Set("holddown_interval", vv); err != nil {
				return fmt.Errorf("Error reading holddown_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading holddown_interval: %v", err)
		}
	}

	if err = d.Set("http_host", flattenObjectFirewallAccessProxyApiGatewayRealserversHttpHost3rdl(o["http-host"], d, "http_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-host"], "ObjectFirewallAccessProxyApiGatewayRealservers-HttpHost"); ok {
			if err = d.Set("http_host", vv); err != nil {
				return fmt.Errorf("Error reading http_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_host: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallAccessProxyApiGatewayRealserversId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallAccessProxyApiGatewayRealservers-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenObjectFirewallAccessProxyApiGatewayRealserversIp3rdl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectFirewallAccessProxyApiGatewayRealservers-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mappedport", flattenObjectFirewallAccessProxyApiGatewayRealserversMappedport3rdl(o["mappedport"], d, "mappedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["mappedport"], "ObjectFirewallAccessProxyApiGatewayRealservers-Mappedport"); ok {
			if err = d.Set("mappedport", vv); err != nil {
				return fmt.Errorf("Error reading mappedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mappedport: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectFirewallAccessProxyApiGatewayRealserversPort3rdl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectFirewallAccessProxyApiGatewayRealservers-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("ssh_client_cert", flattenObjectFirewallAccessProxyApiGatewayRealserversSshClientCert3rdl(o["ssh-client-cert"], d, "ssh_client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-client-cert"], "ObjectFirewallAccessProxyApiGatewayRealservers-SshClientCert"); ok {
			if err = d.Set("ssh_client_cert", vv); err != nil {
				return fmt.Errorf("Error reading ssh_client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_client_cert: %v", err)
		}
	}

	if err = d.Set("ssh_host_key", flattenObjectFirewallAccessProxyApiGatewayRealserversSshHostKey3rdl(o["ssh-host-key"], d, "ssh_host_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-host-key"], "ObjectFirewallAccessProxyApiGatewayRealservers-SshHostKey"); ok {
			if err = d.Set("ssh_host_key", vv); err != nil {
				return fmt.Errorf("Error reading ssh_host_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_host_key: %v", err)
		}
	}

	if err = d.Set("ssh_host_key_validation", flattenObjectFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation3rdl(o["ssh-host-key-validation"], d, "ssh_host_key_validation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-host-key-validation"], "ObjectFirewallAccessProxyApiGatewayRealservers-SshHostKeyValidation"); ok {
			if err = d.Set("ssh_host_key_validation", vv); err != nil {
				return fmt.Errorf("Error reading ssh_host_key_validation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_host_key_validation: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectFirewallAccessProxyApiGatewayRealserversStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectFirewallAccessProxyApiGatewayRealservers-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("translate_host", flattenObjectFirewallAccessProxyApiGatewayRealserversTranslateHost3rdl(o["translate-host"], d, "translate_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["translate-host"], "ObjectFirewallAccessProxyApiGatewayRealservers-TranslateHost"); ok {
			if err = d.Set("translate_host", vv); err != nil {
				return fmt.Errorf("Error reading translate_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading translate_host: %v", err)
		}
	}

	if err = d.Set("tunnel_encryption", flattenObjectFirewallAccessProxyApiGatewayRealserversTunnelEncryption3rdl(o["tunnel-encryption"], d, "tunnel_encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-encryption"], "ObjectFirewallAccessProxyApiGatewayRealservers-TunnelEncryption"); ok {
			if err = d.Set("tunnel_encryption", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_encryption: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFirewallAccessProxyApiGatewayRealserversType3rdl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallAccessProxyApiGatewayRealservers-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("weight", flattenObjectFirewallAccessProxyApiGatewayRealserversWeight3rdl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ObjectFirewallAccessProxyApiGatewayRealservers-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallAccessProxyApiGatewayRealserversFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallAccessProxyApiGatewayRealserversAddrType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversAddress3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversDomain3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversExternalAuth3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversHealthCheck3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversHealthCheckProto3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversHolddownInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversHttpHost3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversMappedport3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversSshClientCert3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversSshHostKey3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversTranslateHost3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversTunnelEncryption3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxyApiGatewayRealserversWeight3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallAccessProxyApiGatewayRealservers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversAddrType3rdl(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversAddress3rdl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversDomain3rdl(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("external_auth"); ok || d.HasChange("external_auth") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversExternalAuth3rdl(d, v, "external_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-auth"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversHealthCheck3rdl(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("health_check_proto"); ok || d.HasChange("health_check_proto") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversHealthCheckProto3rdl(d, v, "health_check_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-proto"] = t
		}
	}

	if v, ok := d.GetOk("holddown_interval"); ok || d.HasChange("holddown_interval") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversHolddownInterval3rdl(d, v, "holddown_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["holddown-interval"] = t
		}
	}

	if v, ok := d.GetOk("http_host"); ok || d.HasChange("http_host") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversHttpHost3rdl(d, v, "http_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-host"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversIp3rdl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("mappedport"); ok || d.HasChange("mappedport") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversMappedport3rdl(d, v, "mappedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedport"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversPort3rdl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("ssh_client_cert"); ok || d.HasChange("ssh_client_cert") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversSshClientCert3rdl(d, v, "ssh_client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-client-cert"] = t
		}
	}

	if v, ok := d.GetOk("ssh_host_key"); ok || d.HasChange("ssh_host_key") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversSshHostKey3rdl(d, v, "ssh_host_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-host-key"] = t
		}
	}

	if v, ok := d.GetOk("ssh_host_key_validation"); ok || d.HasChange("ssh_host_key_validation") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation3rdl(d, v, "ssh_host_key_validation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-host-key-validation"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("translate_host"); ok || d.HasChange("translate_host") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversTranslateHost3rdl(d, v, "translate_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["translate-host"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_encryption"); ok || d.HasChange("tunnel_encryption") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversTunnelEncryption3rdl(d, v, "tunnel_encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-encryption"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversType3rdl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandObjectFirewallAccessProxyApiGatewayRealserversWeight3rdl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
