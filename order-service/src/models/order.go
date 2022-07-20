package models

type Order struct {
	// gorm.Model
	OrderID 		int 	`gorm:"primary_key;size:11;not null;" json:"order_id"`
	MadeByUserID   	int 	`gorm:"size:11;not null;" json:"made_by_id"`
	ProductDict		string	`gorm:"size:256;not null;" json:"prod_dict"`
	TotalPrice		int		`gorm:"size:11;not null;" json:"total_price"`
}

func (o *Order) SaveOrder() (*Order, error) {

	var err error
	err = DB.Create(&o).Error
	if err != nil {
		return &Order{}, err
	}
	return o, nil
}

func (o *Order) UpdateOrder() (*Order, error) {

	var err error
	err = DB.Save(&o).Error
	if err != nil {
		return &Order{}, err
	}
	return o, nil
}
