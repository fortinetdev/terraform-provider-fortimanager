package forticlient

import (
	"fmt"
	"log"
	"sort"
)

// sortObjectSystemSdnConnectorCompartmentListItem contains the parameters for each item
type sortObjectSystemSdnConnectorCompartmentListItem struct {
	compartment_id string
}

func getEntryListObjectSystemSdnConnectorCompartmentList(c *FortiSDKClient, inputModel *SortInputModel) (itemList []sortObjectSystemSdnConnectorCompartmentListItem, err error) {
	path := "/pm/config/[*]/obj/system/sdn-connector/{sdn-connector}/compartment-list"
	path, err = replaceParaWithValue(path, inputModel.URLParams)
	params := map[string]interface{}{
		"fields": []string{"compartment-id"},
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

		var members []sortObjectSystemSdnConnectorCompartmentListItem
		for _, v := range listTmp {
			c := v.(map[string]interface{})

			members = append(members,
				sortObjectSystemSdnConnectorCompartmentListItem{
					compartment_id: conv2str(c["compartment-id"]),
				})
		}

		itemList = members
	}

	return
}

func bEntryListSortedObjectSystemSdnConnectorCompartmentList(itemList []sortObjectSystemSdnConnectorCompartmentListItem, inputModel *SortInputModel) (bsorted bool) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	bsorted = true
	if sortby == "compartment-id" {
		for i := 0; i < len(itemList)-1; i++ {
			if sortdirection == "ascending" {
				if itemList[i].compartment_id > itemList[i+1].compartment_id {
					bsorted = false
					return
				}
			} else if sortdirection == "descending" {
				if itemList[i].compartment_id < itemList[i+1].compartment_id {
					bsorted = false
					return
				}
			} else if sortdirection == "manual" {
				curItemMap := make(map[string]int)
				for index, item := range itemList {
					curKeyValue := item.compartment_id
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

func moveAfterObjectSystemSdnConnectorCompartmentList(idbefore, idafter string, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	path := "/pm/config/[*]/obj/system/sdn-connector/{sdn-connector}/compartment-list/"
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

func sortEntryListObjectSystemSdnConnectorCompartmentList(itemList []sortObjectSystemSdnConnectorCompartmentListItem, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	var targetItemOrder []sortObjectSystemSdnConnectorCompartmentListItem
	if sortby == "compartment-id" {
		if sortdirection == "ascending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].compartment_id < itemList[j].compartment_id
			})
			targetItemOrder = itemList
		} else if sortdirection == "descending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].compartment_id > itemList[j].compartment_id
			})
			targetItemOrder = itemList
		} else if sortdirection == "manual" {
			curItemMap := make(map[string]sortObjectSystemSdnConnectorCompartmentListItem)
			for _, item := range itemList {
				curIndex := item.compartment_id
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
		err = moveAfterObjectSystemSdnConnectorCompartmentList(targetItemOrder[i+1].compartment_id, targetItemOrder[i].compartment_id, c, inputModel)
		if err != nil {
			err = fmt.Errorf("Error move entry: %s", err)
			return
		}
	}

	return nil
}

// CreateUpdateObjectSystemSdnConnectorCompartmentListSort API operation for FortiManager to sort the firewall policies.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateObjectSystemSdnConnectorCompartmentListSort(inputModel *SortInputModel) (err error) {
	itemList, err := getEntryListObjectSystemSdnConnectorCompartmentList(c, inputModel)
	log.Printf("[INFO] Entry ID list: %v", itemList)
	if err != nil {
		err = fmt.Errorf("Fail to get entries before sort: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectSystemSdnConnectorCompartmentList(itemList, inputModel)
	if bsorted == true {
		return
	}

	err = sortEntryListObjectSystemSdnConnectorCompartmentList(itemList, c, inputModel)
	if err != nil {
		err = fmt.Errorf("Error when sort entries: %s", err)
		return
	}

	return
}

// ReadObjectSystemSdnConnectorCompartmentListSort API operation for FortiManager to read the firewall policies sort results
// Returns sort status
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadObjectSystemSdnConnectorCompartmentListSort(inputModel *SortInputModel) (sorted bool, itemMapList []interface{}, err error) {
	itemList, err := getEntryListObjectSystemSdnConnectorCompartmentList(c, inputModel)
	if err != nil {
		err = fmt.Errorf("Fail to read the entries: %s", err)
		return
	}

	bsorted := bEntryListSortedObjectSystemSdnConnectorCompartmentList(itemList, inputModel)
	if bsorted == true {
		sorted = true
		return
	}

	sorted = false
	for _, item := range itemList {
		curItemMap := make(map[string]interface{})
		curItemMap["compartment-id"] = item.compartment_id
		itemMapList = append(itemMapList, curItemMap)
	}

	return
}
