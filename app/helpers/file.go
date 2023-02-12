package helpers

import (
	"fmt"
	"os"
	"strings"
)

func Model(fileName string) {
	f, err := os.Create("app/models/" + strings.Title(fileName) + ".go")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	d := []string{
		"package models",
		"",
		"import (",
		"\t\"time\"",
		")",
		"",
		"type " + strings.Title(fileName) + " struct {",
		"\tID          string         `json:\"id\" gorm:\"uniqueIndex;not null\"`",
		"\tCreatedAt   time.Time      `json:\"created_at\" gorm:\"autoCreateTime\"`",
		"\tUpdatedAt   time.Time      `json:\"updated_at\" gorm:\"autoUpdateTime\"`",
		"}",
	}

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Controller(fileName string) {
	f, err := os.Create("app/controllers/" + strings.Title(fileName) + "Controller.go")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	d := []string{
		"package controllers",
		"",
		"type " + strings.Title(fileName) + "Controller interface {}",
		"",
		"type " + strings.ToLower(fileName) + "Controller struct {",
		"\t" + strings.ToLower(fileName) + "Service services." + strings.Title(fileName) + "Service",
		"}",
		"",
		"func New" + strings.Title(fileName) + "Controller(" + strings.ToLower(fileName) + "Serv services." + strings.Title(fileName) + "Service) " + strings.Title(fileName) + "Controller {",
		"\treturn &" + strings.ToLower(fileName) + "Controller{",
		"\t" + strings.ToLower(fileName) + "Service: " + strings.ToLower(fileName) + "Serv,",
		"\t}",
		"}",
	}

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Service(fileName string) {
	f, err := os.Create("app/services/" + strings.Title(fileName) + "Service.go")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	d := []string{
		"package services",
		"",
		"type " + strings.Title(fileName) + "Service interface {}",
		"",
		"type " + strings.ToLower(fileName) + "Service struct {",
		"\t" + strings.ToLower(fileName) + "Repository repositories." + strings.Title(fileName) + "Repository",
		"}",
		"",
		"func New" + strings.Title(fileName) + "Service(" + strings.ToLower(fileName) + "Repo repositories." + strings.Title(fileName) + "Repository) " + strings.Title(fileName) + "Service {",
		"\treturn &" + strings.ToLower(fileName) + "Service{",
		"\t" + strings.ToLower(fileName) + "Repository: " + strings.ToLower(fileName) + "Repo,",
		"\t}",
		"}",
	}

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Repository(fileName string) {
	f, err := os.Create("app/repositories/" + strings.Title(fileName) + "Repository.go")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	d := []string{
		"package repositories",
		"",
		"import (\"gorm.io/gorm\")",
		"",
		"type " + strings.Title(fileName) + "Repository interface {}",
		"",
		"type " + strings.ToLower(fileName) + "Connection struct {connection *gorm.DB}",
		"",
		"func New" + strings.Title(fileName) + "Repository(dbConn *gorm.DB) " + strings.Title(fileName) + "Repository { return &" + strings.ToLower(fileName) + "Connection{connection: dbConn,}}",
	}

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Dto(fileName string) {
	f, err := os.Create("app/dtos/" + strings.Title(fileName) + "DTO.go")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	d := []string{
		"package dtos",
		"",
		"type " + strings.Title(fileName) + "CreateDTO struct {",
		"\tID string `json:\"id\" form:\"id\"`",
		"}",
		"",
		"type " + strings.Title(fileName) + "UpdateDTO struct {",
		"ID string `json:\"id\" form:\"id\"`",
		"}",
	}

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
