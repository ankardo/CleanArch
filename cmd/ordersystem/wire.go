//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/ankardo/CleanArch/internal/entity"
	"github.com/ankardo/CleanArch/internal/event"
	"github.com/ankardo/CleanArch/internal/infra/database"
	"github.com/ankardo/CleanArch/internal/infra/web"
	"github.com/ankardo/CleanArch/internal/usecase"
	"github.com/ankardo/CleanArch/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewFindOrderUseCase(db *sql.DB) *usecase.FindOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewFindOrderUseCase,
	)
	return &usecase.FindOrderUseCase{}
}

func NewFindAllOrdersUseCase(db *sql.DB) *usecase.FindAllOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewFindAllOrdersUseCase,
	)
	return &usecase.FindAllOrdersUseCase{}
}

func NewOrderCreateHandler(
	db *sql.DB,
	eventDispatcher events.EventDispatcherInterface,
) *web.OrderCreateHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewOrderCreateHandler,
	)
	return &web.OrderCreateHandler{}
}

func NewFindAllOrdersHandler(
	db *sql.DB,
) *web.FindAllOrdersHandler {
	wire.Build(
		setOrderRepositoryDependency,
		web.NewFindAllOrdersHandler,
	)
	return &web.FindAllOrdersHandler{}
}

func NewFindOrderHandler(
	db *sql.DB,
) *web.FindOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		web.NewFindOrderHandler,
	)
	return &web.FindOrderHandler{}
}
