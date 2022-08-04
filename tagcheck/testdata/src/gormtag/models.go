package main

import (
	"database/sql"
	"time"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
// NamedPet is a reference to a Named `Pets` (has many)
type User struct {
	Name      string
	Age       uint
	Birthday  *time.Time
	Account   Account
	Pets      []*Pet
	NamedPet  *Pet
	Toys      []Toy `gorm:"polymorphic:Owner"`
	CompanyID *int
	Company   Company
	ManagerID *uint
	Manager   *User
	Team      []User     `gorm:"foreignkey:ManagerID "`
	Languages []Language `gorm:"many2many:UserSpeak;"`
	Friends   []*User    `gorm:"many2many:user_friends;"`
	Active    bool
}

type Account struct {
	UserID sql.NullInt64
	Number string
}

type Pet struct {
	UserID *uint
	Name   string
	Toy    Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	Name      string
	OwnerID   string
	OwnerType string
}

type Company struct {
	ID   int
	Name string
}

type Language struct {
	Code string `gorm:"primarykey"`
	Name string
}

type Coupon struct {
	ID               int              `gorm:"primarykey; size:255"`
	AppliesToProduct []*CouponProduct `gorm:"foreignKey:CouponId;constraint:OnDelete:CASCADE"`
	AmountOff        uint32           `gorm:"amount_off"`  //@ diag(`not support Gorm option "amount_off"`)
	PercentOff       float32          `gorm:"percent_off"` //@ diag(`not support Gorm option "percent_off"`)
}

type CouponProduct struct {
	CouponId  int    `gorm:"primarykey;size:255"`
	ProductId string `gorm:"primarykey;size:255"`
	Desc      string
}

type Order struct {
	Num      string
	Coupon   *Coupon
	CouponID string
}

type Parent struct {
	FavChildID uint
	FavChild   *Child
	Children   []*Child
}

type Child struct {
	Name     string `gorm:"column:name;column:nam"` //@ diag(`duplicate Gorm option "column"`)
	Name2    string `gorm:"column"`                 //@ diag(`not support Gorm option "column" value "" can not be empty`)
	Size     int    `gorm:"size:10"`
	Size2    int    `gorm:"size:a"`             //@ diag(`not support Gorm option "size" value "a" not an uint`)
	Size3    int    `gorm:"size:-1"`            //@ diag(`not support Gorm option "size" value "-1" not an uint`)
	Bool     bool   `gorm:"autoIncrement:fals"` //@ diag(`not support Gorm option "autoIncrement" value "fals" not empty or bool`)
	Def      string `gorm:"default:"`           //@ diag(`not support Gorm option "default" value "" can not be empty`)
	ParentID *uint
	Parent   *Parent `gorm:"foreignKey"` //@ diag(`not support Gorm option "foreignKey" value "" can not be empty`)
}

func TestFn() {
	type MyCoupon struct {
		Id        int
		AmountOff uint32 `gorm:"amount_off"` //@ diag(`not support Gorm option "amount_off"`)
	}
}
