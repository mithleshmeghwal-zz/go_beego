package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// Product ...
type Product struct {
	OrderID string `orm:"column(order_id);pk"`
	User    *User  `orm:"rel(fk)"`
	Price   string
}

// AddP ...
func AddP(p *Product) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(p)
	return id, err
}

// Data ...
func Data() {
	o := orm.NewOrm()
	userIds := "userN"
	var product []*Product
	num, err := o.QueryTable("product").Filter("user_id__in", userIds).RelatedSel().All(&product)
	fmt.Println("PROFJF", product[0].OrderID, num, err)
}
