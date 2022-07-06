package models

type Payment struct {
	// gorm.Model
	PaymentID 		int 	`gorm:"primary_key;size:11;not null;" json:"payment_id"`
	MadeByUserID   	int 	`gorm:"size:11;not null;" json:"made_by_id"`
	Amount			int 	`gorm:"size:11;not null;" json:"amount"`
	Paid			bool	`gorm:"not null;" json:"paid"`
}

func (p *Payment) SavePayment() (*Payment, error) {

	var err error
	err = DB.Create(&p).Error
	if err != nil {
		return &Payment{}, err
	}
	return p, nil
}

func (p *Payment) UpdatePayment(pay_id int, paid bool) (*Payment, error) {

	var err error
	err = DB.Model(&p).Where("payment_id = ?", pay_id).Update("paid", true).Error
	if err != nil {
		return &Payment{}, err
	}
	return p, nil
}


