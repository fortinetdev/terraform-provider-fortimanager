---
subcategory: ""
layout: "fortimanager"
page_title: "To Enable Web Services"
description: |-
  To Enable Web Services
---


# To Enable Web Services

FortiManager provides WebServices API access to the console via several different methods. However,
this access is disabled by default. How to enable WebServices access with FortiManager?

## Enable web service of the management interface

Add `webservice` to allowaccess:

```
config system interface
    edit "port2"
        set ip 192.168.52.178 255.255.255.0
        set allowaccess ping https ssh snmp webservice
    next
end
```

## Set the permission level for rpc-permit

```
config system admin user
    edit "apiusertest"
        set password testadmin123
        set rpc-permit read-write
        profileid "Super_User"
    end
end
```

FortiManager is now configured and ready to accept WebService requests.