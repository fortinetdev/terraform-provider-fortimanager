// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WAN optimization profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWanoptProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWanoptProfileCreate,
		Read:   resourceObjectWanoptProfileRead,
		Update: resourceObjectWanoptProfileUpdate,
		Delete: resourceObjectWanoptProfileDelete,

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
			"auth_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cifs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"byte_caching": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefer_chunking": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol_opt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secure_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnel_sharing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"byte_caching": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefer_chunking": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol_opt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secure_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnel_sharing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"http": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"byte_caching": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefer_chunking": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol_opt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secure_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnel_sharing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mapi": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"byte_caching": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secure_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnel_sharing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"tcp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"byte_caching": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"byte_caching_opt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secure_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_port": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnel_sharing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"transparent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWanoptProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWanoptProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWanoptProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWanoptProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWanoptProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWanoptProfileRead(d, m)
}

func resourceObjectWanoptProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWanoptProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWanoptProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWanoptProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWanoptProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWanoptProfileRead(d, m)
}

func resourceObjectWanoptProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWanoptProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWanoptProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWanoptProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWanoptProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWanoptProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWanoptProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWanoptProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectWanoptProfileAuthGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileCifs(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "byte_caching"
	if _, ok := i["byte-caching"]; ok {
		result["byte_caching"] = flattenObjectWanoptProfileCifsByteCaching(i["byte-caching"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_traffic"
	if _, ok := i["log-traffic"]; ok {
		result["log_traffic"] = flattenObjectWanoptProfileCifsLogTraffic(i["log-traffic"], d, pre_append)
	}

	pre_append = pre + ".0." + "prefer_chunking"
	if _, ok := i["prefer-chunking"]; ok {
		result["prefer_chunking"] = flattenObjectWanoptProfileCifsPreferChunking(i["prefer-chunking"], d, pre_append)
	}

	pre_append = pre + ".0." + "protocol_opt"
	if _, ok := i["protocol-opt"]; ok {
		result["protocol_opt"] = flattenObjectWanoptProfileCifsProtocolOpt(i["protocol-opt"], d, pre_append)
	}

	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := i["secure-tunnel"]; ok {
		result["secure_tunnel"] = flattenObjectWanoptProfileCifsSecureTunnel(i["secure-tunnel"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWanoptProfileCifsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := i["tunnel-sharing"]; ok {
		result["tunnel_sharing"] = flattenObjectWanoptProfileCifsTunnelSharing(i["tunnel-sharing"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWanoptProfileCifsByteCaching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileCifsLogTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileCifsPreferChunking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileCifsProtocolOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileCifsSecureTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileCifsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileCifsTunnelSharing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileFtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "byte_caching"
	if _, ok := i["byte-caching"]; ok {
		result["byte_caching"] = flattenObjectWanoptProfileFtpByteCaching(i["byte-caching"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_traffic"
	if _, ok := i["log-traffic"]; ok {
		result["log_traffic"] = flattenObjectWanoptProfileFtpLogTraffic(i["log-traffic"], d, pre_append)
	}

	pre_append = pre + ".0." + "prefer_chunking"
	if _, ok := i["prefer-chunking"]; ok {
		result["prefer_chunking"] = flattenObjectWanoptProfileFtpPreferChunking(i["prefer-chunking"], d, pre_append)
	}

	pre_append = pre + ".0." + "protocol_opt"
	if _, ok := i["protocol-opt"]; ok {
		result["protocol_opt"] = flattenObjectWanoptProfileFtpProtocolOpt(i["protocol-opt"], d, pre_append)
	}

	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := i["secure-tunnel"]; ok {
		result["secure_tunnel"] = flattenObjectWanoptProfileFtpSecureTunnel(i["secure-tunnel"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl"
	if _, ok := i["ssl"]; ok {
		result["ssl"] = flattenObjectWanoptProfileFtpSsl(i["ssl"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWanoptProfileFtpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := i["tunnel-sharing"]; ok {
		result["tunnel_sharing"] = flattenObjectWanoptProfileFtpTunnelSharing(i["tunnel-sharing"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWanoptProfileFtpByteCaching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileFtpLogTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileFtpPreferChunking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileFtpProtocolOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileFtpSecureTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileFtpSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileFtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileFtpTunnelSharing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileHttp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "byte_caching"
	if _, ok := i["byte-caching"]; ok {
		result["byte_caching"] = flattenObjectWanoptProfileHttpByteCaching(i["byte-caching"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_traffic"
	if _, ok := i["log-traffic"]; ok {
		result["log_traffic"] = flattenObjectWanoptProfileHttpLogTraffic(i["log-traffic"], d, pre_append)
	}

	pre_append = pre + ".0." + "prefer_chunking"
	if _, ok := i["prefer-chunking"]; ok {
		result["prefer_chunking"] = flattenObjectWanoptProfileHttpPreferChunking(i["prefer-chunking"], d, pre_append)
	}

	pre_append = pre + ".0." + "protocol_opt"
	if _, ok := i["protocol-opt"]; ok {
		result["protocol_opt"] = flattenObjectWanoptProfileHttpProtocolOpt(i["protocol-opt"], d, pre_append)
	}

	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := i["secure-tunnel"]; ok {
		result["secure_tunnel"] = flattenObjectWanoptProfileHttpSecureTunnel(i["secure-tunnel"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl"
	if _, ok := i["ssl"]; ok {
		result["ssl"] = flattenObjectWanoptProfileHttpSsl(i["ssl"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWanoptProfileHttpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := i["tunnel-sharing"]; ok {
		result["tunnel_sharing"] = flattenObjectWanoptProfileHttpTunnelSharing(i["tunnel-sharing"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWanoptProfileHttpByteCaching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileHttpLogTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileHttpPreferChunking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileHttpProtocolOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileHttpSecureTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileHttpSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileHttpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileHttpTunnelSharing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileMapi(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "byte_caching"
	if _, ok := i["byte-caching"]; ok {
		result["byte_caching"] = flattenObjectWanoptProfileMapiByteCaching(i["byte-caching"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_traffic"
	if _, ok := i["log-traffic"]; ok {
		result["log_traffic"] = flattenObjectWanoptProfileMapiLogTraffic(i["log-traffic"], d, pre_append)
	}

	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := i["secure-tunnel"]; ok {
		result["secure_tunnel"] = flattenObjectWanoptProfileMapiSecureTunnel(i["secure-tunnel"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWanoptProfileMapiStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := i["tunnel-sharing"]; ok {
		result["tunnel_sharing"] = flattenObjectWanoptProfileMapiTunnelSharing(i["tunnel-sharing"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWanoptProfileMapiByteCaching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileMapiLogTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileMapiSecureTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileMapiStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileMapiTunnelSharing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileTcp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "byte_caching"
	if _, ok := i["byte-caching"]; ok {
		result["byte_caching"] = flattenObjectWanoptProfileTcpByteCaching(i["byte-caching"], d, pre_append)
	}

	pre_append = pre + ".0." + "byte_caching_opt"
	if _, ok := i["byte-caching-opt"]; ok {
		result["byte_caching_opt"] = flattenObjectWanoptProfileTcpByteCachingOpt(i["byte-caching-opt"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_traffic"
	if _, ok := i["log-traffic"]; ok {
		result["log_traffic"] = flattenObjectWanoptProfileTcpLogTraffic(i["log-traffic"], d, pre_append)
	}

	pre_append = pre + ".0." + "port"
	if _, ok := i["port"]; ok {
		result["port"] = flattenObjectWanoptProfileTcpPort(i["port"], d, pre_append)
	}

	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := i["secure-tunnel"]; ok {
		result["secure_tunnel"] = flattenObjectWanoptProfileTcpSecureTunnel(i["secure-tunnel"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl"
	if _, ok := i["ssl"]; ok {
		result["ssl"] = flattenObjectWanoptProfileTcpSsl(i["ssl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_port"
	if _, ok := i["ssl-port"]; ok {
		result["ssl_port"] = flattenObjectWanoptProfileTcpSslPort(i["ssl-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectWanoptProfileTcpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := i["tunnel-sharing"]; ok {
		result["tunnel_sharing"] = flattenObjectWanoptProfileTcpTunnelSharing(i["tunnel-sharing"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWanoptProfileTcpByteCaching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileTcpByteCachingOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileTcpLogTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileTcpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileTcpSecureTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileTcpSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileTcpSslPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectWanoptProfileTcpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileTcpTunnelSharing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptProfileTransparent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWanoptProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_group", flattenObjectWanoptProfileAuthGroup(o["auth-group"], d, "auth_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-group"], "ObjectWanoptProfile-AuthGroup"); ok {
			if err = d.Set("auth_group", vv); err != nil {
				return fmt.Errorf("Error reading auth_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("cifs", flattenObjectWanoptProfileCifs(o["cifs"], d, "cifs")); err != nil {
			if vv, ok := fortiAPIPatch(o["cifs"], "ObjectWanoptProfile-Cifs"); ok {
				if err = d.Set("cifs", vv); err != nil {
					return fmt.Errorf("Error reading cifs: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading cifs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cifs"); ok {
			if err = d.Set("cifs", flattenObjectWanoptProfileCifs(o["cifs"], d, "cifs")); err != nil {
				if vv, ok := fortiAPIPatch(o["cifs"], "ObjectWanoptProfile-Cifs"); ok {
					if err = d.Set("cifs", vv); err != nil {
						return fmt.Errorf("Error reading cifs: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading cifs: %v", err)
				}
			}
		}
	}

	if err = d.Set("comments", flattenObjectWanoptProfileComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "ObjectWanoptProfile-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ftp", flattenObjectWanoptProfileFtp(o["ftp"], d, "ftp")); err != nil {
			if vv, ok := fortiAPIPatch(o["ftp"], "ObjectWanoptProfile-Ftp"); ok {
				if err = d.Set("ftp", vv); err != nil {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ftp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftp"); ok {
			if err = d.Set("ftp", flattenObjectWanoptProfileFtp(o["ftp"], d, "ftp")); err != nil {
				if vv, ok := fortiAPIPatch(o["ftp"], "ObjectWanoptProfile-Ftp"); ok {
					if err = d.Set("ftp", vv); err != nil {
						return fmt.Errorf("Error reading ftp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("http", flattenObjectWanoptProfileHttp(o["http"], d, "http")); err != nil {
			if vv, ok := fortiAPIPatch(o["http"], "ObjectWanoptProfile-Http"); ok {
				if err = d.Set("http", vv); err != nil {
					return fmt.Errorf("Error reading http: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading http: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("http"); ok {
			if err = d.Set("http", flattenObjectWanoptProfileHttp(o["http"], d, "http")); err != nil {
				if vv, ok := fortiAPIPatch(o["http"], "ObjectWanoptProfile-Http"); ok {
					if err = d.Set("http", vv); err != nil {
						return fmt.Errorf("Error reading http: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading http: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mapi", flattenObjectWanoptProfileMapi(o["mapi"], d, "mapi")); err != nil {
			if vv, ok := fortiAPIPatch(o["mapi"], "ObjectWanoptProfile-Mapi"); ok {
				if err = d.Set("mapi", vv); err != nil {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mapi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mapi"); ok {
			if err = d.Set("mapi", flattenObjectWanoptProfileMapi(o["mapi"], d, "mapi")); err != nil {
				if vv, ok := fortiAPIPatch(o["mapi"], "ObjectWanoptProfile-Mapi"); ok {
					if err = d.Set("mapi", vv); err != nil {
						return fmt.Errorf("Error reading mapi: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectWanoptProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWanoptProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tcp", flattenObjectWanoptProfileTcp(o["tcp"], d, "tcp")); err != nil {
			if vv, ok := fortiAPIPatch(o["tcp"], "ObjectWanoptProfile-Tcp"); ok {
				if err = d.Set("tcp", vv); err != nil {
					return fmt.Errorf("Error reading tcp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tcp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tcp"); ok {
			if err = d.Set("tcp", flattenObjectWanoptProfileTcp(o["tcp"], d, "tcp")); err != nil {
				if vv, ok := fortiAPIPatch(o["tcp"], "ObjectWanoptProfile-Tcp"); ok {
					if err = d.Set("tcp", vv); err != nil {
						return fmt.Errorf("Error reading tcp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tcp: %v", err)
				}
			}
		}
	}

	if err = d.Set("transparent", flattenObjectWanoptProfileTransparent(o["transparent"], d, "transparent")); err != nil {
		if vv, ok := fortiAPIPatch(o["transparent"], "ObjectWanoptProfile-Transparent"); ok {
			if err = d.Set("transparent", vv); err != nil {
				return fmt.Errorf("Error reading transparent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transparent: %v", err)
		}
	}

	return nil
}

func flattenObjectWanoptProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWanoptProfileAuthGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileCifs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "byte_caching"
	if _, ok := d.GetOk(pre_append); ok {
		result["byte-caching"], _ = expandObjectWanoptProfileCifsByteCaching(d, i["byte_caching"], pre_append)
	}
	pre_append = pre + ".0." + "log_traffic"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-traffic"], _ = expandObjectWanoptProfileCifsLogTraffic(d, i["log_traffic"], pre_append)
	}
	pre_append = pre + ".0." + "prefer_chunking"
	if _, ok := d.GetOk(pre_append); ok {
		result["prefer-chunking"], _ = expandObjectWanoptProfileCifsPreferChunking(d, i["prefer_chunking"], pre_append)
	}
	pre_append = pre + ".0." + "protocol_opt"
	if _, ok := d.GetOk(pre_append); ok {
		result["protocol-opt"], _ = expandObjectWanoptProfileCifsProtocolOpt(d, i["protocol_opt"], pre_append)
	}
	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := d.GetOk(pre_append); ok {
		result["secure-tunnel"], _ = expandObjectWanoptProfileCifsSecureTunnel(d, i["secure_tunnel"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectWanoptProfileCifsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := d.GetOk(pre_append); ok {
		result["tunnel-sharing"], _ = expandObjectWanoptProfileCifsTunnelSharing(d, i["tunnel_sharing"], pre_append)
	}

	return result, nil
}

func expandObjectWanoptProfileCifsByteCaching(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileCifsLogTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileCifsPreferChunking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileCifsProtocolOpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileCifsSecureTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileCifsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileCifsTunnelSharing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileFtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "byte_caching"
	if _, ok := d.GetOk(pre_append); ok {
		result["byte-caching"], _ = expandObjectWanoptProfileFtpByteCaching(d, i["byte_caching"], pre_append)
	}
	pre_append = pre + ".0." + "log_traffic"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-traffic"], _ = expandObjectWanoptProfileFtpLogTraffic(d, i["log_traffic"], pre_append)
	}
	pre_append = pre + ".0." + "prefer_chunking"
	if _, ok := d.GetOk(pre_append); ok {
		result["prefer-chunking"], _ = expandObjectWanoptProfileFtpPreferChunking(d, i["prefer_chunking"], pre_append)
	}
	pre_append = pre + ".0." + "protocol_opt"
	if _, ok := d.GetOk(pre_append); ok {
		result["protocol-opt"], _ = expandObjectWanoptProfileFtpProtocolOpt(d, i["protocol_opt"], pre_append)
	}
	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := d.GetOk(pre_append); ok {
		result["secure-tunnel"], _ = expandObjectWanoptProfileFtpSecureTunnel(d, i["secure_tunnel"], pre_append)
	}
	pre_append = pre + ".0." + "ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl"], _ = expandObjectWanoptProfileFtpSsl(d, i["ssl"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectWanoptProfileFtpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := d.GetOk(pre_append); ok {
		result["tunnel-sharing"], _ = expandObjectWanoptProfileFtpTunnelSharing(d, i["tunnel_sharing"], pre_append)
	}

	return result, nil
}

func expandObjectWanoptProfileFtpByteCaching(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileFtpLogTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileFtpPreferChunking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileFtpProtocolOpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileFtpSecureTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileFtpSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileFtpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileFtpTunnelSharing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "byte_caching"
	if _, ok := d.GetOk(pre_append); ok {
		result["byte-caching"], _ = expandObjectWanoptProfileHttpByteCaching(d, i["byte_caching"], pre_append)
	}
	pre_append = pre + ".0." + "log_traffic"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-traffic"], _ = expandObjectWanoptProfileHttpLogTraffic(d, i["log_traffic"], pre_append)
	}
	pre_append = pre + ".0." + "prefer_chunking"
	if _, ok := d.GetOk(pre_append); ok {
		result["prefer-chunking"], _ = expandObjectWanoptProfileHttpPreferChunking(d, i["prefer_chunking"], pre_append)
	}
	pre_append = pre + ".0." + "protocol_opt"
	if _, ok := d.GetOk(pre_append); ok {
		result["protocol-opt"], _ = expandObjectWanoptProfileHttpProtocolOpt(d, i["protocol_opt"], pre_append)
	}
	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := d.GetOk(pre_append); ok {
		result["secure-tunnel"], _ = expandObjectWanoptProfileHttpSecureTunnel(d, i["secure_tunnel"], pre_append)
	}
	pre_append = pre + ".0." + "ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl"], _ = expandObjectWanoptProfileHttpSsl(d, i["ssl"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectWanoptProfileHttpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := d.GetOk(pre_append); ok {
		result["tunnel-sharing"], _ = expandObjectWanoptProfileHttpTunnelSharing(d, i["tunnel_sharing"], pre_append)
	}

	return result, nil
}

func expandObjectWanoptProfileHttpByteCaching(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileHttpLogTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileHttpPreferChunking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileHttpProtocolOpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileHttpSecureTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileHttpSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileHttpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileHttpTunnelSharing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileMapi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "byte_caching"
	if _, ok := d.GetOk(pre_append); ok {
		result["byte-caching"], _ = expandObjectWanoptProfileMapiByteCaching(d, i["byte_caching"], pre_append)
	}
	pre_append = pre + ".0." + "log_traffic"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-traffic"], _ = expandObjectWanoptProfileMapiLogTraffic(d, i["log_traffic"], pre_append)
	}
	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := d.GetOk(pre_append); ok {
		result["secure-tunnel"], _ = expandObjectWanoptProfileMapiSecureTunnel(d, i["secure_tunnel"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectWanoptProfileMapiStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := d.GetOk(pre_append); ok {
		result["tunnel-sharing"], _ = expandObjectWanoptProfileMapiTunnelSharing(d, i["tunnel_sharing"], pre_append)
	}

	return result, nil
}

func expandObjectWanoptProfileMapiByteCaching(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileMapiLogTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileMapiSecureTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileMapiStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileMapiTunnelSharing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileTcp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "byte_caching"
	if _, ok := d.GetOk(pre_append); ok {
		result["byte-caching"], _ = expandObjectWanoptProfileTcpByteCaching(d, i["byte_caching"], pre_append)
	}
	pre_append = pre + ".0." + "byte_caching_opt"
	if _, ok := d.GetOk(pre_append); ok {
		result["byte-caching-opt"], _ = expandObjectWanoptProfileTcpByteCachingOpt(d, i["byte_caching_opt"], pre_append)
	}
	pre_append = pre + ".0." + "log_traffic"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-traffic"], _ = expandObjectWanoptProfileTcpLogTraffic(d, i["log_traffic"], pre_append)
	}
	pre_append = pre + ".0." + "port"
	if _, ok := d.GetOk(pre_append); ok {
		result["port"], _ = expandObjectWanoptProfileTcpPort(d, i["port"], pre_append)
	}
	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := d.GetOk(pre_append); ok {
		result["secure-tunnel"], _ = expandObjectWanoptProfileTcpSecureTunnel(d, i["secure_tunnel"], pre_append)
	}
	pre_append = pre + ".0." + "ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl"], _ = expandObjectWanoptProfileTcpSsl(d, i["ssl"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-port"], _ = expandObjectWanoptProfileTcpSslPort(d, i["ssl_port"], pre_append)
	} else {
		result["ssl-port"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectWanoptProfileTcpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := d.GetOk(pre_append); ok {
		result["tunnel-sharing"], _ = expandObjectWanoptProfileTcpTunnelSharing(d, i["tunnel_sharing"], pre_append)
	}

	return result, nil
}

func expandObjectWanoptProfileTcpByteCaching(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileTcpByteCachingOpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileTcpLogTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileTcpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileTcpSecureTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileTcpSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileTcpSslPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectWanoptProfileTcpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileTcpTunnelSharing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptProfileTransparent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWanoptProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_group"); ok {
		t, err := expandObjectWanoptProfileAuthGroup(d, v, "auth_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-group"] = t
		}
	}

	if v, ok := d.GetOk("cifs"); ok {
		t, err := expandObjectWanoptProfileCifs(d, v, "cifs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandObjectWanoptProfileComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("ftp"); ok {
		t, err := expandObjectWanoptProfileFtp(d, v, "ftp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp"] = t
		}
	}

	if v, ok := d.GetOk("http"); ok {
		t, err := expandObjectWanoptProfileHttp(d, v, "http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http"] = t
		}
	}

	if v, ok := d.GetOk("mapi"); ok {
		t, err := expandObjectWanoptProfileMapi(d, v, "mapi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectWanoptProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("tcp"); ok {
		t, err := expandObjectWanoptProfileTcp(d, v, "tcp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp"] = t
		}
	}

	if v, ok := d.GetOk("transparent"); ok {
		t, err := expandObjectWanoptProfileTransparent(d, v, "transparent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transparent"] = t
		}
	}

	return &obj, nil
}
