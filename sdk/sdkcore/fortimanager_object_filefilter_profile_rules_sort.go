package forticlient

import (
	"fmt"
	"log"
	"sort"
)

// sortObjectFileFilterProfileRulesItem contains the parameters for each item
type sortObjectFileFilterProfileRulesItem struct {
	name string
}

func getEntryListObjectFileFilterProfileRules(c *FortiSDKClient, inputModel *SortInputModel) (itemList []sortObjectFileFilterProfileRulesItem, err error) {
	path := "/pm/config/[*]/obj/file-filter/profile/{profile}/rules"
	path, err = replaceParaWithValue(path, inputModel.URLParams)
	params := map[string]interface{}{
		"fields": []string{"name"},
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

		var members []sortObjectFileFilterProfileRulesItem
		for _, v := range listTmp {
			c := v.(map[string]interface{})

			members = append(members,
				sortObjectFileFilterProfileRulesItem{
					name: conv2str(c["name"]),
				})
		}

		itemList = members
	}

	return
}

func bEntryListSortedObjectFileFilterProfileRules(itemList []sortObjectFileFilterProfileRulesItem, inputModel *SortInputModel) (bsorted bool) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	bsorted = true
	if sortby == "name" {
		for i := 0; i < len(itemList)-1; i++ {
			if sortdirection == "ascending" {
				if itemList[i].name > itemList[i+1].name {
					bsorted = false
					return
				}
			} else if sortdirection == "descending" {
				if itemList[i].name < itemList[i+1].name {
					bsorted = false
					return
				}
			} else if sortdirection == "manual" {
				curItemMap := make(map[string]int)
				for index, item := range itemList {
					curKeyValue := item.name
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

func moveAfterObjectFileFilterProfileRules(idbefore, idafter string, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	path := "/pm/config/[*]/obj/file-filter/profile/{profile}/rules/"
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

func sortEntryListObjectFileFilterProfileRules(itemList []sortObjectFileFilterProfileRulesItem, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	var targetItemOrder []sortObjectFileFilterProfileRulesItem
	if sortby == "name" {
		if sortdirection == "ascending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].name < itemList[j].name
			})
			targetItemOrder = itemList
		} else if sortdirection == "descending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].name > itemList[j].name
			})
			targetItemOrder = itemList
		} else if sortdirection == "manual" {
			curItemMap := make(map[string]sortObjectFileFilterProfileRulesItem)
			for _, item := range itemList {
				curIndex := item.name
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
		err = moveAfterObjectFileFilterProfileRules(targetItemOrder[i+1].name, targetItemOrder[i].name, c, inputModel)
		if err != nil {
			err = fmt.Errorf("Error move entry: %s", err)
			return
		}
	}

	return nil
}

// CreateUpdateObjectFileFilterProfileRulesSort API operation for FortiManager to sort the firewall policies.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateObjectFileFilterProfileRulesSort(inputModel *SortInputModel) (err error) {
	itemList, err := getEntryListObjectFileFilterProfileRules(c, inputModel)
	log.Printf("[INFO] Entry ID list: %v", itemList)
	if err != nil {
		err = fmt.Errorf("Fail to get entries before sort: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectFileFilterProfileRules(itemList, inputModel)
	if bsorted == true {
		return
	}

	err = sortEntryListObjectFileFilterProfileRules(itemList, c, inputModel)
	if err != nil {
		err = fmt.Errorf("Error when sort entries: %s", err)
		return
	}

	return
}

// ReadObjectFileFilterProfileRulesSort API operation for FortiManager to read the firewall policies sort results
// Returns sort status
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadObjectFileFilterProfileRulesSort(inputModel *SortInputModel) (sorted bool, itemMapList []interface{}, err error) {
	itemList, err := getEntryListObjectFileFilterProfileRules(c, inputModel)
	if err != nil {
		err = fmt.Errorf("Fail to read the entries: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectFileFilterProfileRules(itemList, inputModel)
	if bsorted == true {
		sorted = true
		return
	}

	sorted = false
	for _, item := range itemList {
		curItemMap := make(map[string]interface{})
		curItemMap["name"] = item.name
		itemMapList = append(itemMapList, curItemMap)
	}

	return
}
