package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/firerplayer/hexagonal-arch-go/internal/infra/akafka"
	"github.com/firerplayer/hexagonal-arch-go/internal/infra/repository"
	"github.com/firerplayer/hexagonal-arch-go/internal/infra/web"
	"github.com/firerplayer/hexagonal-arch-go/internal/usecase"
	"github.com/go-chi/chi/v5"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repo)
	listProductsUseCase := usecase.NewListProductsUseCase(repo)

	// Web
	productHandlers := web.NewProductHandlers(createProductUseCase, listProductsUseCase)
	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProduct)
	r.Get("/products/all", productHandlers.ListProducts)

	go http.ListenAndServe(":8000", r)

	// Kafka
	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)
	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = createProductUseCase.Execute(dto)
		if err != nil {
			panic(err)
		}
	}
}
