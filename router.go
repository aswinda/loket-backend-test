package main

import (
	"sync"

	"github.com/go-chi/chi"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	eventController := ServiceContainer().InjectEventController()

	r := chi.NewRouter()
	r.Get("/event/get_info", eventController.GetEventDetailAction)
	r.Post("/event/create", eventController.CreateEventAction)

	locationController := ServiceContainer().InjectLocationController()
	r.Post("/location/create", locationController.CreateLocationAction)

	ticketController := ServiceContainer().InjectTicketController()
	r.Post("/ticket/create", ticketController.CreateTicketAction)

	transactionController := ServiceContainer().InjectTransactionController()
	r.Post("/transaction/purchase", transactionController.PurchaseTransactionAction)
	r.Get("/transaction/get_info", transactionController.GetTransactionDetailAction)

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
