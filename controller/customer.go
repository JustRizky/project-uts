package controller

import (
	"project-uts-mvc/model"
	"project-uts-mvc/view"
)

type CustomerController struct {
	Model *model.CustomerModel
	View  *view.CustomerView
}
