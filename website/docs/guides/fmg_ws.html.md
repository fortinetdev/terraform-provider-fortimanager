---
subcategory: ""
layout: "fortimanager"
page_title: "To Set the Permission Level for RPC-Permit"
description: |-
  To Set the Permission Level for RPC-Permit
---


# To Set the Permission Level for RPC-Permit

Set the permission level for RPC-Permit:

```
config system admin user
    edit "apiusertest"
        set password testadmin123
        set rpc-permit read-write
        profileid "Super_User"
    end
end
```

FortiManager is now configured and ready to accept terraform requests.