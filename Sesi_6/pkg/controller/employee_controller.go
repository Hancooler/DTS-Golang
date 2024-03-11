package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

var employees = []Employee{
	{ID: "4ad5ccd2-560c-497d-ad54-9905958e7113", Name: "aril", Age: 23, Division: "curriculum developer"},
	{ID: "de3ecd37-1f20-42f2-942c-d900af5fce47", Name: "nanda", Age: 23, Division: "finance"},
	{ID: "c3b4f582-202c-47a1-b42c-991fc0db4f50", Name: "mail", Age: 23, Division: "marketing"},
}

func GetEmployeeList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, employees)
}
