package models

type ProductEntity struct {
	PK                int64  `gorm:"primaryKey:index:column:PK"`
	CbxUniqueStyleId  string `gorm:"column:p_pimcbxuniquestyleid"`
	ItemNumber        string `gorm:"column:p_pimitemnumber"`
	Code              string `gorm:"column:p_code"`
	CbxDisplayType    int64  `gorm:"column:p_pimcbxdisplaytype"`
	PimContainerSize  int    `gorm:"column:p_pimcontainersize"`
	AssemblyItemCount int    `gorm:"column:p_pimassemblyitemcount"`
	AssemblyType      int64  `gorm:"column:p_pimassemblytype"`
	CbxPillar         string `gorm:"column:p_pimcbxpillar"`
	IsDisplayItem     int    `gorm:"column:p_pimisdisplayitem"`

	// Get Merchandise Category
	// Get Brand
	// Get PIMNAME
	// Get MKT TEXT

}

func (ProductEntity) TableName() string {
	return "products"
}
