package models

import (
	"time"

	"gorm.io/gorm"
)

type SelectionItem struct {
	ID                      int64      `gorm:"column:pk"`
	Code                    string     `gorm:"column:p_code"`
	AdvertisingWeek         string     `gorm:"column:p_advertisingweek"`
	DeliveryWeek            string     `gorm:"column:p_deliveryweek"`
	DeliveryYear            string     `gorm:"column:p_deliveryyeear"`
	BusinessYear            string     `gorm:"column:p_businessyear"`
	Note                    string     `gorm:"column:p_note"`
	TargetPriceFrom         string     `gorm:"column:p_targetpricefrom"`
	TargetPriceTo           string     `gorm:"column:p_targetpriceto"`
	Repeating               int        `gorm:"column:p_repeating"`
	StoreCheck              string     `gorm:"column:p_storecheck"`
	TopSeller               int        `gorm:"column:p_topseller"`
	ArticleProposal         int        `gorm:"column:p_articleproposal"`
	CountryProposal         int        `gorm:"column:p_countryproposal"`
	CatalogArticle          int        `gorm:"column:p_catalogarticle"`
	BriefingHl              int        `gorm:"column:p_briefinghl"`
	BasicLine               int        `gorm:"column:p_basicline"`
	MinimumOrderPerQuantity int        `gorm:"column:p_minimumorderperquantity"`
	InitialOrder            int        `gorm:"column:p_initialorder"`
	StylesChecked           string     `gorm:"column:p_styleschecked"`
	NumOfSamples            string     `gorm:"column:p_numofsamples"`
	CountryOrderStatus      string     `gorm:"column:p_pimcountryorderstatus"`
	IdocStatus              string     `gorm:"column:p_pimidocstatus"`
	WorkflowStatus          string     `gorm:"column:p_pimworkflowstatus"`
	PurchasingNetPrice      string     `gorm:"column:p_purchasingnetprice"`
	TopicCode               string     `gorm:"column:p_topiccode"`
	ReturnsGuarantee        string     `gorm:"column:p_returnsguarantee"`
	Buyer                   string     `gorm:"column:p_buyer"`
	CbxLotNumber            string     `gorm:"column:p_pimcbxlotnumber"`
	Product                 string     `gorm:"column:p_product"`
	CreatedTS               *time.Time `json:"createdTS"`
	ModifiedTS              *time.Time `json:"modifiedTS"`
}

func (SelectionItem) TableName() string {
	return "pimselectionitem"
}

func GetModifiedSelectionItems(since time.Time, db *gorm.DB) ([]*SelectionItem, error) {
	var selectionItems []*SelectionItem
	db.Find(&selectionItems, "ModifiedTS > ?", since)

	return selectionItems, nil
}
