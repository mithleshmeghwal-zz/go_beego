package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

// User ...
type User struct {
	Username  string `orm:"pk"`
	Name      string
	Email     string    `orm:"unique"`
	Confirmed uint8     `orm:"default(0)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime);null"`
}

func init() {
	orm.RegisterModel(new(User), new(Product))
}

// Add ...
func Add(u *User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(u)
	return id, err
}

// Update ...
func Update(u *User, params ...string) (int64, error) {
	o := orm.NewOrm()
	err := o.Read(u)
	u.Confirmed = 3
	if err == nil {
		num, err := o.Update(u, params...)
		return num, err
	}
	return 0, err
}

// UserByUsername ...
func UserByUsername(u *User, params ...string) (int64, error) {
	o := orm.NewOrm()
	err := o.Read(u)
	return 0, err
}

// UserTransaction ...
func UserTransaction(u *User, params ...string) {
	fmt.Println("TRANSACTION")
	o := orm.NewOrm()

	o.Begin()
	err := o.Read(u)
	if err != nil {
		o.Rollback()
		fmt.Println("READING ERROR")
	}
	u.Confirmed = 5
	num, err := o.Update(u, params...)
	if err != nil {
		o.Rollback()
		fmt.Println("UPDATING ERROR")
	}
	o.Commit()
	fmt.Println("USER TRANSACTION COMPLETED", num)
}
