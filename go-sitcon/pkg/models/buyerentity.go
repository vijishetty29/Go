package models

type BuyerEntity struct {
	PK       int64  `gorm:"column:PK"`
	Name     string `gorm:"column:p_name"`
	LastName string `gorm:"column:p_lastname"`
}

func (BuyerEntity) TableName() string {
	return "pimbuyer"
}

func GetBuyerNameByID(id int64) (string, error) {
	var buyer *BuyerEntity
	DB.Find(&buyer, "PK = ?", id)

	if buyer != nil {
		return buyer.Name + " " + buyer.LastName, nil
	}

	return "", nil
}
