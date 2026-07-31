package forticlient

import (
	"fmt"
	"log"
	"sort"
)

// sortObjectSystemSdnConnectorExternalAccountListItem contains the parameters for each item
type sortObjectSystemSdnConnectorExternalAccountListItem struct {
	role_arn string
}

func getEntryListObjectSystemSdnConnectorExternalAccountList(c *FortiSDKClient, inputModel *SortInputModel) (itemList []sortObjectSystemSdnConnectorExternalAccountListItem, err error) {
	path := "/pm/config/[*]/obj/system/sdn-connector/{sdn-connector}/external-account-list"
	path, err = replaceParaWithValue(path, inputModel.URLParams)
	params := map[string]interface{}{
		"fields": []string{"role-arn"},
	}

	requestInput := &requestInput{}

	requestInput.fortiSDKClient = c
	requestInput.method = "get"
	requestInput.path = path
	requestInput.bodyParams = &params
	requestInput.bMove = true

	listTmp, err := readMove(requestInput)

	if err == nil {
		if listTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		var members []sortObjectSystemSdnConnectorExternalAccountListItem
		for _, v := range listTmp {
			c := v.(map[string]interface{})

			members = append(members,
				sortObjectSystemSdnConnectorExternalAccountListItem{
					role_arn: conv2str(c["role-arn"]),
				})
		}

		itemList = members
	}

	return
}

func bEntryListSortedObjectSystemSdnConnectorExternalAccountList(itemList []sortObjectSystemSdnConnectorExternalAccountListItem, inputModel *SortInputModel) (bsorted bool) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	bsorted = true
	if sortby == "role-arn" {
		for i := 0; i < len(itemList)-1; i++ {
			if sortdirection == "ascending" {
				if itemList[i].role_arn > itemList[i+1].role_arn {
					bsorted = false
					return
				}
			} else if sortdirection == "descending" {
				if itemList[i].role_arn < itemList[i+1].role_arn {
					bsorted = false
					return
				}
			} else if sortdirection == "manual" {
				curItemMap := make(map[string]int)
				for index, item := range itemList {
					curKeyValue := item.role_arn
					curItemMap[curKeyValue] = index
				}
				for j := 0; j < len(manual_order)-1; j++ {
					indexL, okL := curItemMap[manual_order[j].(string)]
					indexR, okR := curItemMap[manual_order[j+1].(string)]
					if okL && okR && indexL > indexR {
						bsorted = false
						return
					}
				}
			}
		}
	}

	return
}

func moveAfterObjectSystemSdnConnectorExternalAccountList(idbefore, idafter string, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	path := "/pm/config/[*]/obj/system/sdn-connector/{sdn-connector}/external-account-list/"
	path, err = replaceParaWithValue(path, inputModel.URLParams)

	params := make(map[string]interface{})
	path += "/" + idbefore
	params["target"] = idafter
	params["option"] = "after"

	requestInput := &requestInput{}

	requestInput.fortiSDKClient = c
	requestInput.method = "move"
	requestInput.path = path
	requestInput.bodyParams = &params
	requestInput.wsParams = inputModel.WSParams
	requestInput.bMove = true

	_, err = createUpdate(requestInput)

	return
}

func sortEntryListObjectSystemSdnConnectorExternalAccountList(itemList []sortObjectSystemSdnConnectorExternalAccountListItem, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	var targetItemOrder []sortObjectSystemSdnConnectorExternalAccountListItem
	if sortby == "role-arn" {
		if sortdirection == "ascending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].role_arn < itemList[j].role_arn
			})
			targetItemOrder = itemList
		} else if sortdirection == "descending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].role_arn > itemList[j].role_arn
			})
			targetItemOrder = itemList
		} else if sortdirection == "manual" {
			curItemMap := make(map[string]sortObjectSystemSdnConnectorExternalAccountListItem)
			for _, item := range itemList {
				curIndex := item.role_arn
				curItemMap[curIndex] = item
			}
			for _, val := range manual_order {
				if item, ok := curItemMap[val.(string)]; ok {
					targetItemOrder = append(targetItemOrder, item)
				}
			}
		}
	}

	for i := 0; i < len(targetItemOrder)-1; i++ {
		err = moveAfterObjectSystemSdnConnectorExternalAccountList(targetItemOrder[i+1].role_arn, targetItemOrder[i].role_arn, c, inputModel)
		if err != nil {
			err = fmt.Errorf("Error move entry: %s", err)
			return
		}
	}

	return nil
}

// CreateUpdateObjectSystemSdnConnectorExternalAccountListSort API operation for FortiManager to sort the firewall policies.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateObjectSystemSdnConnectorExternalAccountListSort(inputModel *SortInputModel) (err error) {
	itemList, err := getEntryListObjectSystemSdnConnectorExternalAccountList(c, inputModel)
	log.Printf("[INFO] Entry ID list: %v", itemList)
	if err != nil {
		err = fmt.Errorf("Fail to get entries before sort: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectSystemSdnConnectorExternalAccountList(itemList, inputModel)
	if bsorted == true {
		return
	}

	err = sortEntryListObjectSystemSdnConnectorExternalAccountList(itemList, c, inputModel)
	if err != nil {
		err = fmt.Errorf("Error when sort entries: %s", err)
		return
	}

	return
}

// ReadObjectSystemSdnConnectorExternalAccountListSort API operation for FortiManager to read the firewall policies sort results
// Returns sort status
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadObjectSystemSdnConnectorExternalAccountListSort(inputModel *SortInputModel) (sorted bool, itemMapList []interface{}, err error) {
	itemList, err := getEntryListObjectSystemSdnConnectorExternalAccountList(c, inputModel)
	if err != nil {
		err = fmt.Errorf("Fail to read the entries: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectSystemSdnConnectorExternalAccountList(itemList, inputModel)
	if bsorted == true {
		sorted = true
		return
	}

	sorted = false
	for _, item := range itemList {
		curItemMap := make(map[string]interface{})
		curItemMap["role-arn"] = item.role_arn
		itemMapList = append(itemMapList, curItemMap)
	}

	return
}
