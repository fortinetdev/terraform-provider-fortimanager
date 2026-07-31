package forticlient

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

// sortObjectVideofilterProfileFortiguardCategoryFiltersItem contains the parameters for each item
type sortObjectVideofilterProfileFortiguardCategoryFiltersItem struct {
	id int
}

func getEntryListObjectVideofilterProfileFortiguardCategoryFilters(c *FortiSDKClient, inputModel *SortInputModel) (itemList []sortObjectVideofilterProfileFortiguardCategoryFiltersItem, err error) {
	path := "/pm/config/[*]/obj/videofilter/profile/{profile}/fortiguard-category/filters"
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

		var members []sortObjectVideofilterProfileFortiguardCategoryFiltersItem
		for _, v := range listTmp {
			c := v.(map[string]interface{})

			members = append(members,
				sortObjectVideofilterProfileFortiguardCategoryFiltersItem{
					id: int(c["id"].(float64)),
				})
		}

		itemList = members
	}

	return
}

func bEntryListSortedObjectVideofilterProfileFortiguardCategoryFilters(itemList []sortObjectVideofilterProfileFortiguardCategoryFiltersItem, inputModel *SortInputModel) (bsorted bool) {
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

func moveAfterObjectVideofilterProfileFortiguardCategoryFilters(idbefore, idafter int, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	idbefores := strconv.Itoa(idbefore)
	idafters := strconv.Itoa(idafter)
	path := "/pm/config/[*]/obj/videofilter/profile/{profile}/fortiguard-category/filters/"
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

func sortEntryListObjectVideofilterProfileFortiguardCategoryFilters(itemList []sortObjectVideofilterProfileFortiguardCategoryFiltersItem, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	var targetItemOrder []sortObjectVideofilterProfileFortiguardCategoryFiltersItem
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
			curItemMap := make(map[string]sortObjectVideofilterProfileFortiguardCategoryFiltersItem)
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
		err = moveAfterObjectVideofilterProfileFortiguardCategoryFilters(targetItemOrder[i+1].id, targetItemOrder[i].id, c, inputModel)
		if err != nil {
			err = fmt.Errorf("Error move entry: %s", err)
			return
		}
	}

	return nil
}

// CreateUpdateObjectVideofilterProfileFortiguardCategoryFiltersSort API operation for FortiManager to sort the firewall policies.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateObjectVideofilterProfileFortiguardCategoryFiltersSort(inputModel *SortInputModel) (err error) {
	itemList, err := getEntryListObjectVideofilterProfileFortiguardCategoryFilters(c, inputModel)
	log.Printf("[INFO] Entry ID list: %v", itemList)
	if err != nil {
		err = fmt.Errorf("Fail to get entries before sort: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectVideofilterProfileFortiguardCategoryFilters(itemList, inputModel)
	if bsorted == true {
		return
	}

	err = sortEntryListObjectVideofilterProfileFortiguardCategoryFilters(itemList, c, inputModel)
	if err != nil {
		err = fmt.Errorf("Error when sort entries: %s", err)
		return
	}

	return
}

// ReadObjectVideofilterProfileFortiguardCategoryFiltersSort API operation for FortiManager to read the firewall policies sort results
// Returns sort status
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadObjectVideofilterProfileFortiguardCategoryFiltersSort(inputModel *SortInputModel) (sorted bool, itemMapList []interface{}, err error) {
	itemList, err := getEntryListObjectVideofilterProfileFortiguardCategoryFilters(c, inputModel)
	if err != nil {
		err = fmt.Errorf("Fail to read the entries: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectVideofilterProfileFortiguardCategoryFilters(itemList, inputModel)
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
