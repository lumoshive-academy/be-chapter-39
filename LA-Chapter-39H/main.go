package main

import (
	"golang-chapter-39/LA-Chapter-39H/infra"
	"golang-chapter-39/LA-Chapter-39H/routes"
	"log"
)

func main() {
	ctx := infra.NewServiceContext()

	r := routes.SetupRouter(*ctx)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
