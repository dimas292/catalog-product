package employee

import (
	"context"
	"log"
	"time"
)

type repositoryContract interface{
	findAllEmployees(ctx context.Context)(res []Employee, err error)
}

type service struct{
	repo repositoryContract
}

func newService(repo repositoryContract) service{
	return service{repo: repo}
}

func(s service) listEmployee(ctx context.Context) (employees []listEmployeeResponse, err error){

	resp, err := s.repo.findAllEmployees(ctx)
	if err != nil {
		log.Println("[listEmployees, findALlEmployee] Error", err)
		return []listEmployeeResponse{}, nil
	}

	for _, v := range resp {
		var createdAt = v.CreatedAt.Format(time.DateOnly)
		emp := listEmployeeResponse{
			Id: v.Id,
			Name: v.Name,
			NIP: v.NIP,
			Address: v.Address,
			CreatedAt: createdAt,
		}

		employees = append(employees, emp)
	}

	return employees, nil
}

