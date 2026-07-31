package forticlient

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

// sortObjectDlpProfileRuleItem contains the parameters for each item
type sortObjectDlpProfileRuleItem struct {
	id int
}

func getEntryListObjectDlpProfileRule(c *FortiSDKClient, inputModel *SortInputModel) (itemList []sortObjectDlpProfileRuleItem, err error) {
	path := "/pm/config/[*]/obj/dlp/profile/{profile}/rule"
	path, err = replaceParaWithValue(path, inputModel.URLParams)
	params := map[string]interface{}{
		"fields": []string{"id"},
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

		var members []sortObjectDlpProfileRuleItem
		for _, v := range listTmp {
			c := v.(map[string]interface{})

			members = append(members,
				sortObjectDlpProfileRuleItem{
					id: int(c["id"].(float64)),
				})
		}

		itemList = members
	}

	return
}

func bEntryListSortedObjectDlpProfileRule(itemList []sortObjectDlpProfileRuleItem, inputModel *SortInputModel) (bsorted bool) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	bsorted = true
	if sortby == "id" {
		for i := 0; i < len(itemList)-1; i++ {
			if sortdirection == "ascending" {
				if itemList[i].id > itemList[i+1].id {
					bsorted = false
					return
				}
			} else if sortdirection == "descending" {
				if itemList[i].id < itemList[i+1].id {
					bsorted = false
					return
				}
			} else if sortdirection == "manual" {
				curItemMap := make(map[string]int)
				for index, item := range itemList {
					curKeyValue := strconv.Itoa(item.id)
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

func moveAfterObjectDlpProfileRule(idbefore, idafter int, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	idbefores := strconv.Itoa(idbefore)
	idafters := strconv.Itoa(idafter)
	path := "/pm/config/[*]/obj/dlp/profile/{profile}/rule/"
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

func sortEntryListObjectDlpProfileRule(itemList []sortObjectDlpProfileRuleItem, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	var targetItemOrder []sortObjectDlpProfileRuleItem
	if sortby == "id" {
		if sortdirection == "ascending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].id < itemList[j].id
			})
			targetItemOrder = itemList
		} else if sortdirection == "descending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].id > itemList[j].id
			})
			targetItemOrder = itemList
		} else if sortdirection == "manual" {
			curItemMap := make(map[string]sortObjectDlpProfileRuleItem)
			for _, item := range itemList {
				curIndex := strconv.Itoa(item.id)
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
		err = moveAfterObjectDlpProfileRule(targetItemOrder[i+1].id, targetItemOrder[i].id, c, inputModel)
		if err != nil {
			err = fmt.Errorf("Error move entry: %s", err)
			return
		}
	}

	return nil
}

// CreateUpdateObjectDlpProfileRuleSort API operation for FortiManager to sort the firewall policies.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateObjectDlpProfileRuleSort(inputModel *SortInputModel) (err error) {
	itemList, err := getEntryListObjectDlpProfileRule(c, inputModel)
	log.Printf("[INFO] Entry ID list: %v", itemList)
	if err != nil {
		err = fmt.Errorf("Fail to get entries before sort: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectDlpProfileRule(itemList, inputModel)
	if bsorted == true {
		return
	}

	err = sortEntryListObjectDlpProfileRule(itemList, c, inputModel)
	if err != nil {
		err = fmt.Errorf("Error when sort entries: %s", err)
		return
	}

	return
}

// ReadObjectDlpProfileRuleSort API operation for FortiManager to read the firewall policies sort results
// Returns sort status
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadObjectDlpProfileRuleSort(inputModel *SortInputModel) (sorted bool, itemMapList []interface{}, err error) {
	itemList, err := getEntryListObjectDlpProfileRule(c, inputModel)
	if err != nil {
		err = fmt.Errorf("Fail to read the entries: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectDlpProfileRule(itemList, inputModel)
	if bsorted == true {
		sorted = true
		return
	}

	sorted = false
	for _, item := range itemList {
		curItemMap := make(map[string]interface{})
		curItemMap["id"] = item.id
		itemMapList = append(itemMapList, curItemMap)
	}

	return
}
