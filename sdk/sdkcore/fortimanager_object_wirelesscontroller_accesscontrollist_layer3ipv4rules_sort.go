package forticlient

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

// sortObjectWirelessControllerAccessControlListLayer3Ipv4RulesItem contains the parameters for each item
type sortObjectWirelessControllerAccessControlListLayer3Ipv4RulesItem struct {
	rule_id int
}

func getEntryListObjectWirelessControllerAccessControlListLayer3Ipv4Rules(c *FortiSDKClient, inputModel *SortInputModel) (itemList []sortObjectWirelessControllerAccessControlListLayer3Ipv4RulesItem, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/access-control-list/{access-control-list}/layer3-ipv4-rules"
	path, err = replaceParaWithValue(path, inputModel.URLParams)
	params := map[string]interface{}{
		"fields": []string{"rule-id"},
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

		var members []sortObjectWirelessControllerAccessControlListLayer3Ipv4RulesItem
		for _, v := range listTmp {
			c := v.(map[string]interface{})

			members = append(members,
				sortObjectWirelessControllerAccessControlListLayer3Ipv4RulesItem{
					rule_id: int(c["rule-id"].(float64)),
				})
		}

		itemList = members
	}

	return
}

func bEntryListSortedObjectWirelessControllerAccessControlListLayer3Ipv4Rules(itemList []sortObjectWirelessControllerAccessControlListLayer3Ipv4RulesItem, inputModel *SortInputModel) (bsorted bool) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	bsorted = true
	if sortby == "rule-id" {
		for i := 0; i < len(itemList)-1; i++ {
			if sortdirection == "ascending" {
				if itemList[i].rule_id > itemList[i+1].rule_id {
					bsorted = false
					return
				}
			} else if sortdirection == "descending" {
				if itemList[i].rule_id < itemList[i+1].rule_id {
					bsorted = false
					return
				}
			} else if sortdirection == "manual" {
				curItemMap := make(map[string]int)
				for index, item := range itemList {
					curKeyValue := strconv.Itoa(item.rule_id)
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

func moveAfterObjectWirelessControllerAccessControlListLayer3Ipv4Rules(idbefore, idafter int, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	idbefores := strconv.Itoa(idbefore)
	idafters := strconv.Itoa(idafter)
	path := "/pm/config/[*]/obj/wireless-controller/access-control-list/{access-control-list}/layer3-ipv4-rules/"
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

func sortEntryListObjectWirelessControllerAccessControlListLayer3Ipv4Rules(itemList []sortObjectWirelessControllerAccessControlListLayer3Ipv4RulesItem, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	var targetItemOrder []sortObjectWirelessControllerAccessControlListLayer3Ipv4RulesItem
	if sortby == "rule-id" {
		if sortdirection == "ascending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].rule_id < itemList[j].rule_id
			})
			targetItemOrder = itemList
		} else if sortdirection == "descending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].rule_id > itemList[j].rule_id
			})
			targetItemOrder = itemList
		} else if sortdirection == "manual" {
			curItemMap := make(map[string]sortObjectWirelessControllerAccessControlListLayer3Ipv4RulesItem)
			for _, item := range itemList {
				curIndex := strconv.Itoa(item.rule_id)
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
		err = moveAfterObjectWirelessControllerAccessControlListLayer3Ipv4Rules(targetItemOrder[i+1].rule_id, targetItemOrder[i].rule_id, c, inputModel)
		if err != nil {
			err = fmt.Errorf("Error move entry: %s", err)
			return
		}
	}

	return nil
}

// CreateUpdateObjectWirelessControllerAccessControlListLayer3Ipv4RulesSort API operation for FortiManager to sort the firewall policies.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateObjectWirelessControllerAccessControlListLayer3Ipv4RulesSort(inputModel *SortInputModel) (err error) {
	itemList, err := getEntryListObjectWirelessControllerAccessControlListLayer3Ipv4Rules(c, inputModel)
	log.Printf("[INFO] Entry ID list: %v", itemList)
	if err != nil {
		err = fmt.Errorf("Fail to get entries before sort: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectWirelessControllerAccessControlListLayer3Ipv4Rules(itemList, inputModel)
	if bsorted == true {
		return
	}

	err = sortEntryListObjectWirelessControllerAccessControlListLayer3Ipv4Rules(itemList, c, inputModel)
	if err != nil {
		err = fmt.Errorf("Error when sort entries: %s", err)
		return
	}

	return
}

// ReadObjectWirelessControllerAccessControlListLayer3Ipv4RulesSort API operation for FortiManager to read the firewall policies sort results
// Returns sort status
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadObjectWirelessControllerAccessControlListLayer3Ipv4RulesSort(inputModel *SortInputModel) (sorted bool, itemMapList []interface{}, err error) {
	itemList, err := getEntryListObjectWirelessControllerAccessControlListLayer3Ipv4Rules(c, inputModel)
	if err != nil {
		err = fmt.Errorf("Fail to read the entries: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectWirelessControllerAccessControlListLayer3Ipv4Rules(itemList, inputModel)
	if bsorted == true {
		sorted = true
		return
	}

	sorted = false
	for _, item := range itemList {
		curItemMap := make(map[string]interface{})
		curItemMap["rule-id"] = item.rule_id
		itemMapList = append(itemMapList, curItemMap)
	}

	return
}
