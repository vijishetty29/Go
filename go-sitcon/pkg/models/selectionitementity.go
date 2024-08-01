package models

import (
	"time"
)

type SelectionItemEntity struct {
	PK                      int64         `gorm:"column:pk"`
	Code                    string        `gorm:"column:p_code"`
	AdvertisingWeek         string        `gorm:"column:p_advertisingweek"`
	DeliveryWeek            string        `gorm:"column:p_deliveryweek"`
	DeliveryYear            string        `gorm:"column:p_deliveryyeear"`
	BusinessYear            string        `gorm:"column:p_businessyear"`
	Note                    string        `gorm:"column:p_note"`
	TargetPriceFrom         string        `gorm:"column:p_targetpricefrom"`
	TargetPriceTo           string        `gorm:"column:p_targetpriceto"`
	Repeating               int           `gorm:"column:p_repeating"`
	StoreCheck              string        `gorm:"column:p_storecheck"`
	TopSeller               int           `gorm:"column:p_topseller"`
	ArticleProposal         int           `gorm:"column:p_articleproposal"`
	CountryProposal         string        `gorm:"column:p_countryproposal"`
	CatalogArticle          int           `gorm:"column:p_catalogarticle"`
	BriefingHl              int           `gorm:"column:p_briefinghl"`
	BasicLine               int           `gorm:"column:p_basicline"`
	MinimumOrderPerQuantity int           `gorm:"column:p_minimumorderperquantity"`
	InitialOrder            int           `gorm:"column:p_initialorder"`
	StylesChecked           string        `gorm:"column:p_styleschecked"`
	NumOfSamples            string        `gorm:"column:p_numofsamples"`
	CountryOrderStatus      int64         `gorm:"column:p_pimcountryorderstatus"`
	IdocStatus              int64         `gorm:"column:p_pimidocstatus"`
	WorkflowStatus          int64         `gorm:"column:p_pimworkflowstatus"`
	PurchasingNetPrice      string        `gorm:"column:p_purchasingnetprice"`
	TopicCode               string        `gorm:"column:p_topiccode"`
	ReturnsGuarantee        string        `gorm:"column:p_returnsguarantee"`
	Buyer                   int64         `gorm:"column:p_buyer"`
	CbxLotNumber            string        `gorm:"column:p_pimcbxlotnumber"`
	ProductCategory         int64         `gorm:"column:p_productcategory"`
	InnovationArticle       int           `gorm:"column:p_innovationarticle"`
	ArticleClassification   int64         `gorm:"column:p_articleclassification"`
	ArticleType             int64         `gorm:"column:p_articletype"`
	ProductPK               int64         `gorm:"column:p_product"`
	Product                 ProductEntity `gorm:"foreignKey:ProductPK"`
	CreatedTS               *time.Time    `json:"createdTS"`
	ModifiedTS              *time.Time    `json:"modifiedTS"`
}

func (SelectionItemEntity) TableName() string {
	return "pimselectionitem"
}

func GetModifiedSelectionItems(since time.Time) ([]*SelectionItemEntity, error) {
	var selectionItems []*SelectionItemEntity

	DB.Preload("Product").Find(&selectionItems, "ModifiedTS > ?", since)

	// DB.Find(&selectionItems, "ModifiedTS > ?", since)

	return selectionItems, nil
}
