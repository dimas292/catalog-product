package menus

import (
	"context"
	"log"
	"time"
)

type repositoryContract interface{
	insertMenu(ctx context.Context, model Menu)(err error)
	findAll(ctx context.Context) (model []Menu, err error)
	findByID(ctx context.Context, id int) (model Menu, err error)
}

type service struct{
	repo repositoryContract
}

func newService(repo repositoryContract) service{
	return service{
		repo: repo,
	}
}

func(s service) createMenu(ctx context.Context, req createMenuRequest){

	var model = Menu{
		Name: req.Name,
		Category: req.Category,
		Desc: req.Desc,
		Price: req.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.insertMenu(ctx, model); err != nil {
		log.Println("[createdmenu, insertmenu] error", err)
		return
	}
	return
}

func (s service) getListMenu(ctx context.Context)(list []listMenuResponse, err error){


	menus, err := s.repo.findAll(ctx)
	if err != nil {
		log.Println("[getlistmenu, getall] error", err)
		return
	}

	if len(menus) == 0 {
		return list, nil
	}

	for _, v := range menus {
		resp := listMenuResponse{
			Id:       menu.Id,
			Price:    menu.Price,
			Desc:     menu.Desc,
			Name:     menu.Name,
			Category: menu.Category,
		}
		list = append(list, resp)
	}

	return list, nil
}

func (s service) getMenuById(ctx context.Context, id int)(resp singgleMenuResponse, err error){
	menu, err := s.getMenuById(ctx, id)
	if err != nil {
		log.Println("[getmenubyid, findbyid] error", err)
	}

	resp = singgleMenuResponse{
		Id:        menu.Id,
		Price:     menu.Price,
		Desc:      menu.Desc,
		Name:      menu.Name,
		Category:  menu.Category,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
	}

	return resp, nil
}