package controllers

import (
	"encoding/json"
	"net/http"
	"ormTest/models"
	"strconv"

	"github.com/astaxie/beego"
)

// MainController ...
type MainController struct {
	beego.Controller
}

// AddUser ...
func (c *MainController) AddUser() {

	var params struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		Email    string `json:"email"`
	}
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)

	if params.Username == "" || params.Name == "" || params.Email == "" {
		ErrorResponse(c, http.StatusBadRequest, "Invalid Params")
		return
	}
	user := models.User{
		Username: params.Username,
		Name:     params.Name,
		Email:    params.Email,
	}
	id, err := models.Add(&user)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(c, http.StatusOK, "User added "+string(id))
}

// UpdateUser ...
func (c *MainController) UpdateUser() {
	var params struct {
		Username  string `json:"username"`
		Confirmed string `json:"confirmed"`
	}
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	_, err := strconv.Atoi(params.Confirmed)

	// fmt.Println(c.Ctx.Input.RequestBody, params.Confirmed, conf, err)

	if params.Username == "" || params.Confirmed == "" || err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid Params")
		return
	}

	user := &models.User{
		Username:  params.Username,
		Confirmed: 1,
	}
	go models.UserTransaction(user, "Confirmed")
	id, err := models.Update(user, "Confirmed")
	go models.Data()
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	SuccessResponse(c, http.StatusOK, "User Updated "+string(id))
}

// AddProduct ...
func (c *MainController) AddProduct() {
	var params struct {
		OrderID string `json:"orderId"`
		Price   string `json:"price"`
	}
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)

	// fmt.Println(c.Ctx.Input.RequestBody, params.Confirmed, conf, err)

	if params.OrderID == "" || params.Price == "" {
		ErrorResponse(c, http.StatusBadRequest, "Invalid Params")
		return
	}
	user := models.User{
		Username: "userN",
	}
	product := &models.Product{
		OrderID: params.OrderID,
		Price:   params.Price,
		User:    &user,
	}
	num, err := models.AddP(product)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	SuccessResponse(c, http.StatusOK, "Product Added "+string(num))
}

// ErrorResponse ...
func ErrorResponse(c *MainController, code int, message string) {
	ResponseJSON(c, code, map[string]string{"error": message})
}

// SuccessResponse ...
func SuccessResponse(c *MainController, code int, message string) {
	ResponseJSON(c, code, map[string]string{"message": message})
}

// ResponseJSON ...
func ResponseJSON(c *MainController, code int, paylaod interface{}) {
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.Data["json"] = paylaod
	c.ServeJSON()
}
