package forticlient

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

// sortObjectWirelessControllerBonjourProfilePolicyListItem contains the parameters for each item
type sortObjectWirelessControllerBonjourProfilePolicyListItem struct {
	policy_id int
}

func getEntryListObjectWirelessControllerBonjourProfilePolicyList(c *FortiSDKClient, inputModel *SortInputModel) (itemList []sortObjectWirelessControllerBonjourProfilePolicyListItem, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/bonjour-profile/{bonjour-profile}/policy-list"
	path, err = replaceParaWithValue(path, inputModel.URLParams)
	params := map[string]interface{}{
		"fields": []string{"policy-id"},
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

		var members []sortObjectWirelessControllerBonjourProfilePolicyListItem
		for _, v := range listTmp {
			c := v.(map[string]interface{})

			members = append(members,
				sortObjectWirelessControllerBonjourProfilePolicyListItem{
					policy_id: int(c["policy-id"].(float64)),
				})
		}

		itemList = members
	}

	return
}

func bEntryListSortedObjectWirelessControllerBonjourProfilePolicyList(itemList []sortObjectWirelessControllerBonjourProfilePolicyListItem, inputModel *SortInputModel) (bsorted bool) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	bsorted = true
	if sortby == "policy-id" {
		for i := 0; i < len(itemList)-1; i++ {
			if sortdirection == "ascending" {
				if itemList[i].policy_id > itemList[i+1].policy_id {
					bsorted = false
					return
				}
			} else if sortdirection == "descending" {
				if itemList[i].policy_id < itemList[i+1].policy_id {
					bsorted = false
					return
				}
			} else if sortdirection == "manual" {
				curItemMap := make(map[string]int)
				for index, item := range itemList {
					curKeyValue := strconv.Itoa(item.policy_id)
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

func moveAfterObjectWirelessControllerBonjourProfilePolicyList(idbefore, idafter int, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	idbefores := strconv.Itoa(idbefore)
	idafters := strconv.Itoa(idafter)
	path := "/pm/config/[*]/obj/wireless-controller/bonjour-profile/{bonjour-profile}/policy-list/"
	path, err = replaceParaWithValue(path, inputModel.URLParams)

	params := make(map[string]interface{})
	path += "/" + idbefores
	params["target"] = idafters
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

func sortEntryListObjectWirelessControllerBonjourProfilePolicyList(itemList []sortObjectWirelessControllerBonjourProfilePolicyListItem, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	var targetItemOrder []sortObjectWirelessControllerBonjourProfilePolicyListItem
	if sortby == "policy-id" {
		if sortdirection == "ascending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].policy_id < itemList[j].policy_id
			})
			targetItemOrder = itemList
		} else if sortdirection == "descending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].policy_id > itemList[j].policy_id
			})
			targetItemOrder = itemList
		} else if sortdirection == "manual" {
			curItemMap := make(map[string]sortObjectWirelessControllerBonjourProfilePolicyListItem)
			for _, item := range itemList {
				curIndex := strconv.Itoa(item.policy_id)
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
		err = moveAfterObjectWirelessControllerBonjourProfilePolicyList(targetItemOrder[i+1].policy_id, targetItemOrder[i].policy_id, c, inputModel)
		if err != nil {
			err = fmt.Errorf("Error move entry: %s", err)
			return
		}
	}

	return nil
}

// CreateUpdateObjectWirelessControllerBonjourProfilePolicyListSort API operation for FortiManager to sort the firewall policies.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateObjectWirelessControllerBonjourProfilePolicyListSort(inputModel *SortInputModel) (err error) {
	itemList, err := getEntryListObjectWirelessControllerBonjourProfilePolicyList(c, inputModel)
	log.Printf("[INFO] Entry ID list: %v", itemList)
	if err != nil {
		err = fmt.Errorf("Fail to get entries before sort: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectWirelessControllerBonjourProfilePolicyList(itemList, inputModel)
	if bsorted == true {
		return
	}

	err = sortEntryListObjectWirelessControllerBonjourProfilePolicyList(itemList, c, inputModel)
	if err != nil {
		err = fmt.Errorf("Error when sort entries: %s", err)
		return
	}

	return
}

// ReadObjectWirelessControllerBonjourProfilePolicyListSort API operation for FortiManager to read the firewall policies sort results
// Returns sort status
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadObjectWirelessControllerBonjourProfilePolicyListSort(inputModel *SortInputModel) (sorted bool, itemMapList []interface{}, err error) {
	itemList, err := getEntryListObjectWirelessControllerBonjourProfilePolicyList(c, inputModel)
	if err != nil {
		err = fmt.Errorf("Fail to read the entries: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectWirelessControllerBonjourProfilePolicyList(itemList, inputModel)
	if bsorted == true {
		sorted = true
		return
	}

	sorted = false
	for _, item := range itemList {
		curItemMap := make(map[string]interface{})
		curItemMap["policy-id"] = item.policy_id
		itemMapList = append(itemMapList, curItemMap)
	}

	return
}
