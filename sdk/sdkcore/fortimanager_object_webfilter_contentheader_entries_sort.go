package forticlient

import (
	"fmt"
	"log"
	"sort"
)

// sortObjectWebfilterContentHeaderEntriesItem contains the parameters for each item
type sortObjectWebfilterContentHeaderEntriesItem struct {
	pattern string
}

func getEntryListObjectWebfilterContentHeaderEntries(c *FortiSDKClient, inputModel *SortInputModel) (itemList []sortObjectWebfilterContentHeaderEntriesItem, err error) {
	path := "/pm/config/[*]/obj/webfilter/content-header/{content-header}/entries"
	path, err = replaceParaWithValue(path, inputModel.URLParams)
	params := map[string]interface{}{
		"fields": []string{"pattern"},
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

		var members []sortObjectWebfilterContentHeaderEntriesItem
		for _, v := range listTmp {
			c := v.(map[string]interface{})

			members = append(members,
				sortObjectWebfilterContentHeaderEntriesItem{
					pattern: conv2str(c["pattern"]),
				})
		}

		itemList = members
	}

	return
}

func bEntryListSortedObjectWebfilterContentHeaderEntries(itemList []sortObjectWebfilterContentHeaderEntriesItem, inputModel *SortInputModel) (bsorted bool) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	bsorted = true
	if sortby == "pattern" {
		for i := 0; i < len(itemList)-1; i++ {
			if sortdirection == "ascending" {
				if itemList[i].pattern > itemList[i+1].pattern {
					bsorted = false
					return
				}
			} else if sortdirection == "descending" {
				if itemList[i].pattern < itemList[i+1].pattern {
					bsorted = false
					return
				}
			} else if sortdirection == "manual" {
				curItemMap := make(map[string]int)
				for index, item := range itemList {
					curKeyValue := item.pattern
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

func moveAfterObjectWebfilterContentHeaderEntries(idbefore, idafter string, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	path := "/pm/config/[*]/obj/webfilter/content-header/{content-header}/entries/"
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

func sortEntryListObjectWebfilterContentHeaderEntries(itemList []sortObjectWebfilterContentHeaderEntriesItem, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	var targetItemOrder []sortObjectWebfilterContentHeaderEntriesItem
	if sortby == "pattern" {
		if sortdirection == "ascending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].pattern < itemList[j].pattern
			})
			targetItemOrder = itemList
		} else if sortdirection == "descending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].pattern > itemList[j].pattern
			})
			targetItemOrder = itemList
		} else if sortdirection == "manual" {
			curItemMap := make(map[string]sortObjectWebfilterContentHeaderEntriesItem)
			for _, item := range itemList {
				curIndex := item.pattern
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
		err = moveAfterObjectWebfilterContentHeaderEntries(targetItemOrder[i+1].pattern, targetItemOrder[i].pattern, c, inputModel)
		if err != nil {
			err = fmt.Errorf("Error move entry: %s", err)
			return
		}
	}

	return nil
}

// CreateUpdateObjectWebfilterContentHeaderEntriesSort API operation for FortiManager to sort the firewall policies.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateObjectWebfilterContentHeaderEntriesSort(inputModel *SortInputModel) (err error) {
	itemList, err := getEntryListObjectWebfilterContentHeaderEntries(c, inputModel)
	log.Printf("[INFO] Entry ID list: %v", itemList)
	if err != nil {
		err = fmt.Errorf("Fail to get entries before sort: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectWebfilterContentHeaderEntries(itemList, inputModel)
	if bsorted == true {
		return
	}

	err = sortEntryListObjectWebfilterContentHeaderEntries(itemList, c, inputModel)
	if err != nil {
		err = fmt.Errorf("Error when sort entries: %s", err)
		return
	}

	return
}

// ReadObjectWebfilterContentHeaderEntriesSort API operation for FortiManager to read the firewall policies sort results
// Returns sort status
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadObjectWebfilterContentHeaderEntriesSort(inputModel *SortInputModel) (sorted bool, itemMapList []interface{}, err error) {
	itemList, err := getEntryListObjectWebfilterContentHeaderEntries(c, inputModel)
	if err != nil {
		err = fmt.Errorf("Fail to read the entries: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectWebfilterContentHeaderEntries(itemList, inputModel)
	if bsorted == true {
		sorted = true
		return
	}

	sorted = false
	for _, item := range itemList {
		curItemMap := make(map[string]interface{})
		curItemMap["pattern"] = item.pattern
		itemMapList = append(itemMapList, curItemMap)
	}

	return
}
