package main

import (
	"log"
	"pujaprabha/internal/config"
	"pujaprabha/internal/domain/models"
)

func main() {
	RegisterPermissionURL()
}

func RegisterPermissionURL() {
	// var errList []string

	urlPaths := map[uint]string{
		1:  "/user/login/",
		2:  "/product/list/",
		3:  "/product/details/",
		4:  "/category/list/",
		5:  "/varient/list/",
		6:  "/cart/create/",
		7:  "/cart/list/",
		8:  "/cart/update/",
		9:  "/cart/delete/",
		10: "/slider/list/",
		11: "/blog/list/",
		12: "/blog/details/",
		13: "/blogCategory/list/",
	}

	db, err := config.ConfigDb()
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range urlPaths {
		err := db.Create(&models.CustomerPermission{
			ID:  k,
			Url: v,
		}).Error
		if err != nil {
			log.Fatal(err)
		}

	}

}
