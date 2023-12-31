package handler

import (
	"os"
	"strconv"

	"github.com/ardin2001/backend-pemilu/controllers/admin/student"
	"github.com/ardin2001/backend-pemilu/controllers/admin/student/usecase"
	"github.com/ardin2001/backend-pemilu/helper"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type StudentHandlerInterface interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type StudentHandlerStruct struct {
	StudentUsecase usecase.StudentUsecaseInterface
}

func NewStudentHandler(StudentUsecase usecase.StudentUsecaseInterface) StudentHandlerInterface {
	return &StudentHandlerStruct{
		StudentUsecase: StudentUsecase,
	}
}

func (sh *StudentHandlerStruct) GetAll(c echo.Context) error {
	var getRequestParam student.GetRequestParam
	c.Bind(&getRequestParam)

	students, statusCode, err := sh.StudentUsecase.GetAll(getRequestParam.NIM)
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData(err.Error(), statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to get all students data", statusCode, students))
}

func (sh *StudentHandlerStruct) GetById(c echo.Context) error {
	id := c.Param("id")

	student, statusCode, err := sh.StudentUsecase.GetById(id)
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData(err.Error(), statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to get student detail data", statusCode, student))
}

func (sh *StudentHandlerStruct) Create(c echo.Context) error {
	var student student.CreateStudent
	godotenv.Load()
	user := os.Getenv("ROLE_DEFAULT")
	userInt, _ := strconv.Atoi(user)
	student.RoleId = userInt
	id := uuid.New().String()
	c.Bind(&student)
	student.ID = id

	err := isRequestValid(student)
	if err != nil {
		return c.JSON(401, helper.ResponseData(err.Error(), 401, nil))
	}

	statusCode, err := sh.StudentUsecase.Create(&student)
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData(err.Error(), statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to create student data", statusCode, nil))
}

func (sh *StudentHandlerStruct) Update(c echo.Context) error {
	id := c.Param("id")
	var student student.UpdateStudent
	c.Bind(&student)
	student.ID = id

	statusCode, err := sh.StudentUsecase.Update(&student)
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData(err.Error(), statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to update student data", statusCode, nil))
}

func (sh *StudentHandlerStruct) Delete(c echo.Context) error {
	id := c.Param("id")

	statusCode, err := sh.StudentUsecase.Delete(id)
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData(err.Error(), statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to delete student data", statusCode, nil))
}
