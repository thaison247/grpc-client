package main

import (
"context"
"io"
"log"
"main/employeepb"
"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())

	if err != nil {
		log.Fatal("error when dial to grpc server: ", err)
	}

	defer cc.Close()

	client := employeepb.NewEmployeeServiceClient(cc)

	//callGetEmployee(client, 109)
	// callGetListEmployees(client, 15, 0)

	newEmp := &employeepb.EmployeeRequest{
		EmployeeId:  120,
		FullName:    "Lê Dung Anh",
		Email:       "le.dug.anh@gmail.com",
		PhoneNumber: "0222 789 222",
		Address:     "Linh Xuân, Thủ Đức, TP.HCM",
		Gender:      true,
		JobTitle:    "designer",
	}

	callUpdateEmployee(client, newEmp)
	callGetEmployee(client, 120)

}

func callGetEmployee(c employeepb.EmployeeServiceClient, employeeId int) {
	log.Println("Client calls -Get employee api")

	res, err := c.GetEmployee(context.Background(), &employeepb.EmployeeRequest{EmployeeId: uint32(employeeId)})

	if err != nil {
		log.Fatalf("error when call get employee api %v", err)
	}

	log.Printf("Get employee api response: %v", res)
}

func callGetListEmployees(c employeepb.EmployeeServiceClient, limit int, offset int) {
	log.Println("Client calls -Get list employees api")

	stream, err := c.GetListEmployees(context.Background(), &employeepb.ListEmployeesRequest{Limit: uint32(limit), Offset: uint32(offset)})

	if err != nil {
		log.Fatalf("Error when client call get list employees api: %v", err)
	}

	for {
		emp, recvErr := stream.Recv()

		if recvErr == io.EOF {
			log.Println("Server finished sending")
			return
		}

		log.Printf("Emloyee data: %v\n", emp)
	}
}

func callCreateEmployee(c employeepb.EmployeeServiceClient, emp *employeepb.EmployeeRequest) {
	res, err := c.CreateEmployee(context.Background(), emp)

	if err != nil {
		log.Fatalf("Error when client call create employee api: %v\n", err)
	}

	log.Printf("Return status: %v\n", res)
}

func callDeleteEmployee(c employeepb.EmployeeServiceClient, emp *employeepb.EmployeeRequest) {
	res, err := c.DeleteEmployee(context.Background(), emp)

	if err != nil {
		log.Fatalf("Error when client call delete employee api: %v\n", err)
	}

	log.Printf("Return status: %v\n", res)
}

func callUpdateEmployee(c employeepb.EmployeeServiceClient, emp *employeepb.EmployeeRequest) {
	res, err := c.UpdateEmployee(context.Background(), emp)

	if err != nil {
		log.Fatalf("Error when client call update employee api: %v\n", err)
	}

	log.Printf("Updated employee data: %v\n", res)
}

