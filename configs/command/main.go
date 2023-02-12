package main

import (
	"fmt"
	"intoduction/app/helpers"
	"intoduction/database"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 && args[1] != "app/handlers/command.go" {
		first := args[1]
		second := ""
		if len(args) > 2 {
			second = args[2]
		}

		if first == "migrate" {
			database.InitialMigration()
		} else if first == "model" {
			if second == "" {
				fmt.Println("make model <filename>")
				os.Exit(0)
			}
			helpers.Model(second)
		} else if first == "controller" {
			if second == "" {
				fmt.Println("make controller <filename>")
				os.Exit(0)
			}
			helpers.Controller(second)
		} else if first == "service" {
			if second == "" {
				fmt.Println("make service <filename>")
				os.Exit(0)
			}
			helpers.Service(second)
		} else if first == "repository" {
			if second == "" {
				fmt.Println("make repository <filename>")
				os.Exit(0)
			}
			helpers.Repository(second)
		} else if first == "dto" {
			if second == "" {
				fmt.Println("make dto <filename>")
				os.Exit(0)
			}
			helpers.Dto(second)
		} else if first == "all" {
			if second == "" {
				fmt.Println("make all <filename>")
				os.Exit(0)
			}
			helpers.Model(second)
			helpers.Dto(second)
			helpers.Repository(second)
			helpers.Service(second)
			helpers.Controller(second)
		}

		if first != "" && second == "" {
			os.Exit(0)
		}
	}
}
