package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	customerReqReps "rest-api/domain/customer/other"
	"rest-api/entity"
	"rest-api/usecase/customer"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	customerServices customer.ServiceCustomer
}

var (
	requestBody entity.Customer
)

func newCustomerHandler(customerService customer.ServiceCustomer) *customerHandler {
	return &customerHandler{customerServices: customerService}
}

func (b *customerHandler) getAllDataCustomer() func(c *gin.Context) {
	return func(c *gin.Context) {
		result, err := b.customerServices.GetAllCustomer()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprint(err))
		} else {
			reps, err2 := json.Marshal(result)
			if err2 != nil {
				c.String(http.StatusInternalServerError, fmt.Sprint(err2))
			} else {
				c.String(http.StatusOK, string(reps))
			}
		}
	}
}

func (b *customerHandler) customerSaveData() func(c *gin.Context) {
	return func(c *gin.Context) {
		err := c.BindJSON(&requestBody)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprint(err))
		} else {
			switch {
			case requestBody.FirstName == "":
				reps, err := json.Marshal("firstName is required")
				if err != nil {
					c.String(http.StatusInternalServerError, fmt.Sprint(err))
				} else {
					c.String(http.StatusBadRequest, string(reps))
				}
			case requestBody.LastName == "":
				reps, err := json.Marshal("lastName is required")
				if err != nil {
					c.String(http.StatusInternalServerError, fmt.Sprint(err))
				} else {
					c.String(http.StatusBadRequest, string(reps))
				}
			default:
				err := b.customerServices.SaveCustomer(requestBody)

				if err != nil {
					c.String(http.StatusInternalServerError, fmt.Sprint(err))
				} else {
					a := customerReqReps.CustomerReps{
						Code:   200,
						Status: "success",
					}
					jsonByte, _ := json.Marshal(a)
					c.String(http.StatusOK, string(jsonByte))
				}
			}
		}
	}
}

func (b *customerHandler) getDataCustomerById() func(c *gin.Context) {
	return func(c *gin.Context) {
		err := c.BindJSON(&requestBody)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprint(err))
		} else {
			switch {
			case requestBody.CustomerID == 0:
				reps, err := json.Marshal("Customer ID is required")
				if err != nil {
					c.String(http.StatusInternalServerError, fmt.Sprint(err))
				} else {
					c.String(http.StatusBadRequest, string(reps))
				}
			default:
				result, err := b.customerServices.GetCustomerById(requestBody.CustomerID)
				if err != nil {
					c.String(http.StatusInternalServerError, fmt.Sprint(err))
				} else {
					reps, err2 := json.Marshal(result)
					if err2 != nil {
						c.String(http.StatusInternalServerError, fmt.Sprint(err2))
					} else {
						c.String(http.StatusOK, string(reps))
					}
				}
			}
		}
	}
}

func (b *customerHandler) customerUpdate() func(c *gin.Context) {
	return func(c *gin.Context) {
		err := c.BindJSON(&requestBody)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprint(err))
		} else {
			switch {
			case requestBody.CustomerID == 0:
				reps, err := json.Marshal("Customer ID is required")
				if err != nil {
					c.String(http.StatusInternalServerError, fmt.Sprint(err))
				} else {
					c.String(http.StatusBadRequest, string(reps))
				}
			default:
				err := b.customerServices.UpdateCustomerById(requestBody.CustomerID, requestBody)
				if err != nil {
					c.String(http.StatusInternalServerError, fmt.Sprint(err))
				} else {
					a := customerReqReps.CustomerReps{
						Code:   200,
						Status: "success",
					}
					jsonByte, _ := json.Marshal(a)
					c.String(http.StatusOK, string(jsonByte))
				}
			}
		}
	}
}

func (b *customerHandler) customerDelete() func(c *gin.Context) {
	return func(c *gin.Context) {
		err := c.BindJSON(&requestBody)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprint(err))
		} else {
			switch {
			case requestBody.CustomerID == 0:
				reps, err := json.Marshal("Customer ID is required")
				if err != nil {
					c.String(http.StatusInternalServerError, fmt.Sprint(err))
				} else {
					c.String(http.StatusBadRequest, string(reps))
				}
			default:
				err := b.customerServices.DeleteCustomerById(requestBody.CustomerID)
				if err != nil {
					c.String(http.StatusInternalServerError, fmt.Sprint(err))
				} else {
					a := customerReqReps.CustomerReps{
						Code:   200,
						Status: "success",
					}
					jsonByte, _ := json.Marshal(a)
					c.String(http.StatusOK, string(jsonByte))
				}
			}
		}
	}
}
