package main

import (
	"fmt"
	"os"
	"os/exec"

	"project-uts-mvc/view"
	customer "project-uts-mvc/view/viewcustomer"
	kasir "project-uts-mvc/view/viewkasir"
	menurestoran "project-uts-mvc/view/viewmenurestoran"
)

func clearCMD() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
menuUtama:
	var pilihanMenuUtama int = view.MenuUtama()
	for pilihanMenuUtama != 4 {
		switch pilihanMenuUtama {
		case 1: //Kasir
			clearCMD()
		menuSecondary1:
			var pilihanMenuKedua int = view.MenuKedua()

			//Insert
			if pilihanMenuKedua == 1 {
				clearCMD()
				kasir.InsertDataKasir()
				goto menuSecondary1

				//Update
			} else if pilihanMenuKedua == 2 {
				clearCMD()
				kasir.UpdateDataKasir()
				goto menuSecondary1

				//Delete
			} else if pilihanMenuKedua == 3 {
				clearCMD()
				kasir.DeleteDataKasir()
				goto menuSecondary1

				//View All
			} else if pilihanMenuKedua == 4 {
				clearCMD()
				kasir.ViewAllDataKasir()
				goto menuSecondary1

				//View By Id
			} else if pilihanMenuKedua == 5 {
				clearCMD()
				kasir.ViewByIdDataKasir()
				goto menuSecondary1

				//Back
			} else if pilihanMenuKedua == 6 {
				goto menuUtama
			}
			break

		case 2: //Customer
			clearCMD()
		menuSecondary2:
			var pilihanMenuKedua int = view.MenuKedua()

			//Insert
			if pilihanMenuKedua == 1 {
				clearCMD()
				customer.InsertDataCustomer()
				goto menuSecondary2

				//Update
			} else if pilihanMenuKedua == 2 {
				clearCMD()
				customer.UpdateDataCustomer()
				goto menuSecondary2

				//Delete
			} else if pilihanMenuKedua == 3 {
				clearCMD()
				customer.DeleteDataCustomer()
				goto menuSecondary2

				//View All
			} else if pilihanMenuKedua == 4 {
				clearCMD()
				customer.ViewAllDataCustomer()
				goto menuSecondary2

				//View By Id
			} else if pilihanMenuKedua == 5 {
				clearCMD()
				customer.ViewByIdDataCustomer()
				goto menuSecondary2

				//Back
			} else if pilihanMenuKedua == 6 {
				goto menuUtama
			}
			break

		case 3: //Customer
			clearCMD()
		menuSecondary3:
			var pilihanMenuKedua int = view.MenuKedua()

			//Insert
			if pilihanMenuKedua == 1 {
				clearCMD()
				menurestoran.InsertDataMenuRestoran()
				goto menuSecondary3

				//Update
			} else if pilihanMenuKedua == 2 {
				clearCMD()
				menurestoran.UpdateDataMenuRestoran()
				goto menuSecondary3

				//Delete
			} else if pilihanMenuKedua == 3 {
				clearCMD()
				menurestoran.DeleteDataMenuRestoran()
				goto menuSecondary3

				//View All
			} else if pilihanMenuKedua == 4 {
				clearCMD()
				menurestoran.ViewAllDataMenuRestoran()
				goto menuSecondary3

				//View By Id
			} else if pilihanMenuKedua == 5 {
				clearCMD()
				menurestoran.ViewByIdDataMenuRestoran()
				goto menuSecondary3

				//Back
			} else if pilihanMenuKedua == 6 {
				goto menuUtama
			}
			break
		}

	menuUlangUtama:
		var ulangMenuUtama string
		fmt.Print("Apakah Anda Ingin Mengulang (Y/t): ")
		fmt.Scan(&ulangMenuUtama)
		if ulangMenuUtama == "y" || ulangMenuUtama == "Y" {
			pilihanMenuUtama = view.MenuUtama()
		} else if ulangMenuUtama == "t" || ulangMenuUtama == "T" {
			break
		} else {
			goto menuUlangUtama
		}
	}
	fmt.Println("Akhir Program!")
}
