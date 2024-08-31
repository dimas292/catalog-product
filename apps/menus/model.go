package menus

import "time"

type Menu struct {
	Id int 
	Name string
	Category string
	Desc string
	Price int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(name, category, desc string, price int) Menu {
	return Menu{
		Name: name,
		Category: category,
		Desc: desc,
		Price: price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m Menu) WithId(id int)Menu{
	m.Id = id
	return m
}