package forticlient

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

// sortPackagesPblockFirewallLocalInPolicyItem contains the parameters for each item
type sortPackagesPblockFirewallLocalInPolicyItem struct {
	policyid int
}

func getEntryListPackagesPblockFirewallLocalInPolicy(c *FortiSDKClient, inputModel *SortInputModel) (itemList []sortPackagesPblockFirewallLocalInPolicyItem, err error) {
	path := "/pm/config/[*]/pblock/{pblock}/firewall/local-in-policy"
	path, err = replaceParaWithValue(path, inputModel.URLParams)
	params := map[string]interface{}{
		"fields": []string{"policyid"},
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

		var members []sortPackagesPblockFirewallLocalInPolicyItem
		for _, v := range listTmp {
			c := v.(map[string]interface{})

			members = append(members,
				sortPackagesPblockFirewallLocalInPolicyItem{
					policyid: int(c["policyid"].(float64)),
				})
		}

		itemList = members
	}

	return
}

func bEntryListSortedPackagesPblockFirewallLocalInPolicy(itemList []sortPackagesPblockFirewallLocalInPolicyItem, inputModel *SortInputModel) (bsorted bool) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	bsorted = true
	if sortby == "policyid" {
		for i := 0; i < len(itemList)-1; i++ {
			if sortdirection == "ascending" {
				if itemList[i].policyid > itemList[i+1].policyid {
					bsorted = false
					return
				}
			} else if sortdirection == "descending" {
				if itemList[i].policyid < itemList[i+1].policyid {
					bsorted = false
					return
				}
			} else if sortdirection == "manual" {
				curItemMap := make(map[string]int)
				for index, item := range itemList {
					curKeyValue := strconv.Itoa(item.policyid)
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

func moveAfterPackagesPblockFirewallLocalInPolicy(idbefore, idafter int, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	idbefores := strconv.Itoa(idbefore)
	idafters := strconv.Itoa(idafter)
	path := "/pm/config/[*]/pblock/{pblock}/firewall/local-in-policy/"
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

func sortEntryListPackagesPblockFirewallLocalInPolicy(itemList []sortPackagesPblockFirewallLocalInPolicyItem, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	var targetItemOrder []sortPackagesPblockFirewallLocalInPolicyItem
	if sortby == "policyid" {
		if sortdirection == "ascending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].policyid < itemList[j].policyid
			})
			targetItemOrder = itemList
		} else if sortdirection == "descending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].policyid > itemList[j].policyid
			})
			targetItemOrder = itemList
		} else if sortdirection == "manual" {
			curItemMap := make(map[string]sortPackagesPblockFirewallLocalInPolicyItem)
			for _, item := range itemList {
				curIndex := strconv.Itoa(item.policyid)
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
		err = moveAfterPackagesPblockFirewallLocalInPolicy(targetItemOrder[i+1].policyid, targetItemOrder[i].policyid, c, inputModel)
		if err != nil {
			err = fmt.Errorf("Error move entry: %s", err)
			return
		}
	}

	return nil
}

// CreateUpdatePackagesPblockFirewallLocalInPolicySort API operation for FortiManager to sort the firewall policies.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdatePackagesPblockFirewallLocalInPolicySort(inputModel *SortInputModel) (err error) {
	itemList, err := getEntryListPackagesPblockFirewallLocalInPolicy(c, inputModel)
	log.Printf("[INFO] Entry ID list: %v", itemList)
	if err != nil {
		err = fmt.Errorf("Fail to get entries before sort: %s", err)
		return
	}

	bsorted := bEntryListSortedPackagesPblockFirewallLocalInPolicy(itemList, inputModel)
	if bsorted == true {
		return
	}

	err = sortEntryListPackagesPblockFirewallLocalInPolicy(itemList, c, inputModel)
	if err != nil {
		err = fmt.Errorf("Error when sort entries: %s", err)
		return
	}

	return
}

// ReadPackagesPblockFirewallLocalInPolicySort API operation for FortiManager to read the firewall policies sort results
// Returns sort status
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadPackagesPblockFirewallLocalInPolicySort(inputModel *SortInputModel) (sorted bool, itemMapList []interface{}, err error) {
	itemList, err := getEntryListPackagesPblockFirewallLocalInPolicy(c, inputModel)
	if err != nil {
		err = fmt.Errorf("Fail to read the entries: %s", err)
		return
	}

	bsorted := bEntryListSortedPackagesPblockFirewallLocalInPolicy(itemList, inputModel)
	if bsorted == true {
		sorted = true
		return
	}

	sorted = false
	for _, item := range itemList {
		curItemMap := make(map[string]interface{})
		curItemMap["policyid"] = item.policyid
		itemMapList = append(itemMapList, curItemMap)
	}

	return
}
