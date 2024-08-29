package employee

import (
	"context"
	"log"
	"time"
)

type repositoryContact interface {
	findAllEmployees(ctx context.Context) (res []Employee, err error)
}

type service struct {
	repo repositoryContact
}

func NewService(repo repositoryContact) service {
	return service{
		repo: repo,
	}
}

func (s service) listEmployees(ctx context.Context) (employees []listEmployeeResponse, err error) {
    // pertama, kita perlu memanggil data dari datasource
    // jadi kita memanggil method pada contract
	resp, err := s.repo.findAllEmployees(ctx)
	if err != nil {
		log.Println("[listEmployees, findAllEmployees] error :", err)
		return []listEmployeeResponse{}, err
	}
    

    // object yang dikembalikan dari datasource adalah Employee
    // sedangkan kita membutuhkan object dalam bentuk listEmployeeResponse
    //
    // jadi kita perlu melakukan convert ke object listEmployeeResponse dari Employee
	for _, res := range resp {
		var createdAt = res.CreatedAt.Format(time.DateOnly)
		emp := listEmployeeResponse{
			Id:        res.Id,
			Name:      res.Name,
			NIP:       res.NIP,
			Address:   res.Address,
			CreatedAt: createdAt,
		}
		employees = append(employees, emp)
	}

	return employees, nil
}
