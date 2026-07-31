package forticlient

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

// sortWantempSystemSdwanMembersItem contains the parameters for each item
type sortWantempSystemSdwanMembersItem struct {
	seq_num int
}

func getEntryListWantempSystemSdwanMembers(c *FortiSDKClient, inputModel *SortInputModel) (itemList []sortWantempSystemSdwanMembersItem, err error) {
	path := "/pm/config/[*]/wanprof/{wanprof}/system/sdwan/members"
	path, err = replaceParaWithValue(path, inputModel.URLParams)
	params := map[string]interface{}{
		"fields": []string{"seq-num"},
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

		var members []sortWantempSystemSdwanMembersItem
		for _, v := range listTmp {
			c := v.(map[string]interface{})

			members = append(members,
				sortWantempSystemSdwanMembersItem{
					seq_num: int(c["seq-num"].(float64)),
				})
		}

		itemList = members
	}

	return
}

func bEntryListSortedWantempSystemSdwanMembers(itemList []sortWantempSystemSdwanMembersItem, inputModel *SortInputModel) (bsorted bool) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	bsorted = true
	if sortby == "seq-num" {
		for i := 0; i < len(itemList)-1; i++ {
			if sortdirection == "ascending" {
				if itemList[i].seq_num > itemList[i+1].seq_num {
					bsorted = false
					return
				}
			} else if sortdirection == "descending" {
				if itemList[i].seq_num < itemList[i+1].seq_num {
					bsorted = false
					return
				}
			} else if sortdirection == "manual" {
				curItemMap := make(map[string]int)
				for index, item := range itemList {
					curKeyValue := strconv.Itoa(item.seq_num)
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

func moveAfterWantempSystemSdwanMembers(idbefore, idafter int, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	idbefores := strconv.Itoa(idbefore)
	idafters := strconv.Itoa(idafter)
	path := "/pm/config/[*]/wanprof/{wanprof}/system/sdwan/members/"
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

func sortEntryListWantempSystemSdwanMembers(itemList []sortWantempSystemSdwanMembersItem, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	var targetItemOrder []sortWantempSystemSdwanMembersItem
	if sortby == "seq-num" {
		if sortdirection == "ascending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].seq_num < itemList[j].seq_num
			})
			targetItemOrder = itemList
		} else if sortdirection == "descending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].seq_num > itemList[j].seq_num
			})
			targetItemOrder = itemList
		} else if sortdirection == "manual" {
			curItemMap := make(map[string]sortWantempSystemSdwanMembersItem)
			for _, item := range itemList {
				curIndex := strconv.Itoa(item.seq_num)
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
		err = moveAfterWantempSystemSdwanMembers(targetItemOrder[i+1].seq_num, targetItemOrder[i].seq_num, c, inputModel)
		if err != nil {
			err = fmt.Errorf("Error move entry: %s", err)
			return
		}
	}

	return nil
}

// CreateUpdateWantempSystemSdwanMembersSort API operation for FortiManager to sort the firewall policies.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateWantempSystemSdwanMembersSort(inputModel *SortInputModel) (err error) {
	itemList, err := getEntryListWantempSystemSdwanMembers(c, inputModel)
	log.Printf("[INFO] Entry ID list: %v", itemList)
	if err != nil {
		err = fmt.Errorf("Fail to get entries before sort: %s", err)
		return
	}

	bsorted := bEntryListSortedWantempSystemSdwanMembers(itemList, inputModel)
	if bsorted == true {
		return
	}

	err = sortEntryListWantempSystemSdwanMembers(itemList, c, inputModel)
	if err != nil {
		err = fmt.Errorf("Error when sort entries: %s", err)
		return
	}

	return
}

// ReadWantempSystemSdwanMembersSort API operation for FortiManager to read the firewall policies sort results
// Returns sort status
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadWantempSystemSdwanMembersSort(inputModel *SortInputModel) (sorted bool, itemMapList []interface{}, err error) {
	itemList, err := getEntryListWantempSystemSdwanMembers(c, inputModel)
	if err != nil {
		err = fmt.Errorf("Fail to read the entries: %s", err)
		return
	}

	bsorted := bEntryListSortedWantempSystemSdwanMembers(itemList, inputModel)
	if bsorted == true {
		sorted = true
		return
	}

	sorted = false
	for _, item := range itemList {
		curItemMap := make(map[string]interface{})
		curItemMap["seq-num"] = item.seq_num
		itemMapList = append(itemMapList, curItemMap)
	}

	return
}
