package main

import (
	"golang-chapter-39/LA-Chapter-39H-I/infra"
	"golang-chapter-39/LA-Chapter-39H-I/routes"
	"log"
)

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	r := routes.SetupRouter(*ctx)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
