package employee

import (
	"context"
	"log"
	"time"
)

type repositoryContract interface{
	findAllEmployees(ctx context.Context)(res []Employee, err error)

	newEmployee(ctx context.Context, req Employee) (err error)
	deleteEmployee(ctx context.Context, id string) (err error)
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

func (s service) createNewEmployee(ctx context.Context, req createNewEmpoyeesRequest) (err error){

	var emp = Employee{
		Name: req.Name,
		NIP: req.NIP,
		Address: req.Address,
	}

	err = s.repo.newEmployee(ctx, emp)
	if err != nil {
		log.Println("[createNewEmployee, newEmployee] error :", err)
		return err
	}

	return nil

}

func(s service) deleteEmployeeByID(ctx context.Context, id string)(err error){

	err = s.repo.deleteEmployee(ctx, id)
	if err != nil {
		log.Println("[deleted, delete by id] error", err)
		return err
	}

	return nil
}


