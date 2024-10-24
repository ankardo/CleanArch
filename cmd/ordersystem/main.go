package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ankardo/CleanArch/configs"
	"github.com/ankardo/CleanArch/internal/event/handler"
	"github.com/ankardo/CleanArch/internal/infra/graph"
	"github.com/ankardo/CleanArch/internal/infra/grpc/pb"
	"github.com/ankardo/CleanArch/internal/infra/grpc/service"
	"github.com/ankardo/CleanArch/internal/infra/web/webserver"
	"github.com/ankardo/CleanArch/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	findOrderUseCase := NewFindOrderUseCase(db)
	findAllOrdersUseCase := NewFindAllOrdersUseCase(db)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	createOrderHandler := NewOrderCreateHandler(db, eventDispatcher)
	findAllOrdersHandler := NewFindAllOrdersHandler(db)
	findOrderHandler := NewFindOrderHandler(db)

	webserver.AddHandler(http.MethodGet, "/orders/{id}", findOrderHandler.GetByID)
	webserver.AddHandler(http.MethodPost, "/orders", createOrderHandler.Create)
	webserver.AddHandler(http.MethodGet, "/orders", findAllOrdersHandler.Get)

	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	orderService := service.NewOrderService(*createOrderUseCase, *findAllOrdersUseCase, *findOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase:   *createOrderUseCase,
		FindAllOrdersUseCase: *findAllOrdersUseCase,
		FindOrderUseCase:     *findOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://appuser:appuserpassword@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
