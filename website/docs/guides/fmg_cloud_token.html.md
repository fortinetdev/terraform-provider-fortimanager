---
subcategory: ""
layout: "fortimanager"
page_title: "Generate an FortiCloud token for FortiManager Cloud"
description: |-
  Generate an FortiCloud token for FortiManager Cloud
---

# Define a FortiCloud API user
1. Navigate to https://support.fortinet.com
2. Login using your FortiCloud account
3. Click Services > IAM
4. Click Permission Profiles > Add New > Add Portal > Select FortiManager Cloud and fill all other info > Submit
5. Click Users > ADD New > API User > Select the permission profile that created in previous step and fill other info > Next > Confirm
6. Successful API User Registration click *DOWNLOAD CREDENTIALS*

You can use the apiId/password as the username/password on the provider configuration. If you want to user FortiCloud Token, see the following step.

# Obtain a FortiCloud token
Process is documented [here](https://docs.fortinet.com/document/fortiauthenticator/6.1.2/rest-api-solution-guide/498666/oauth-server-token-oauth-token)

It is required to build the following payload:
```
{
    "username": "{apiId}",
    "password": "{password}",
    "client_id": "FortiManager",
    "grant_type": "password"
}
```

Where username/password are the apiId/password from previous step. 

Following curl command can be used to obtain the FortiCloud token, the required payload are saved in file data.json:
```
curl -s -k -X POST -H "Content-Type: application/json" https://customerapiauth.fortinet.com/api/v1/oauth/token/ -d @data.json | jq
```
This is the output when the login succeeds:
```
{
    "access_token": "srgkxUG9SqqEG8h4XvNK5qG27tqktk",
    "expires_in": 3600,
    "message": "successfully authenticated",
    "refresh_token": "jie6v6qV62h7NsIITGwmL6GEcmZgst",
    "scope": "read write",
    "status": "success",
    "token_type": "Bearer"
}
```
Use `access_token` as `fmg_cloud_token` on provider configuration. *Note: `access_token` has an expire time, make sure it is valid before use it.*