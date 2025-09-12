package functions

import structure "PROJETRED/structure"

func NewMerchant(merchantItems *[]structure.MerchantItems, DeletedItem string) {
	newMerchantItem := []structure.MerchantItems{}
	for _, element := range *merchantItems {
		if element.Name != DeletedItem {
			newMerchantItem = append(newMerchantItem, element)
		}
	}
	*merchantItems = newMerchantItem
}
