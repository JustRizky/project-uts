package main

import (
	"os"
	"os/exec"
	"project-uts-mvc/controller"
	"project-uts-mvc/model"
	"project-uts-mvc/view"
)

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	customerModel := &model.CustomerModel{}
	customerView := &view.CustomerView{}
	customerController := &controller.CustomerController{
		Model: customerModel,
		View:  customerView,
	}

}
