package web

import (
	"errors"
	"fmt"
	"my_goods/internal/entities"
	"my_goods/internal/service"
	"my_goods/pkg/logger"
)

const (
	getDish      = "get-dish"
	getAllDishes = "get-all-dishes"
	createDish   = "create-dish"
	updateDish   = "update-dish"
	deleteDish   = "delete-dish"

	getGoods    = "get-goods"
	getAllGoods = "get-all-goods"
	createGoods = "create-goods"
	updateGoods = "update-goods"
	deleteGoods = "delete-goods"

	getList     = "get-list"
	getAllLists = "get-all-lists"
	createList  = "create-list"
	updateList  = "update-list"
	deleteList  = "delete-list"
)

type DispatcherInterface interface {
	attach(command string, handler WsHandlerInterface)
	CallHandler(command string, payload []byte)
}

type Dispatcher struct {
	handlers map[string]WsHandlerInterface
}

func NewDispatcher() *Dispatcher {
	d := &Dispatcher{handlers: make(map[string]WsHandlerInterface)}
	d.MatchCommands()
	return d
}

func (d *Dispatcher) MatchCommands() {
	d.attach(getDish, &GetDishHandler{})
	d.attach(getAllDishes, &GetAllDishesHandler{})
	d.attach(createDish, &CreateDishHandler{})
	d.attach(updateDish, &UpdateDishHandler{})
	d.attach(deleteDish, &DeleteDishHandler{})

	d.attach(getGoods, &GetGoodsHandler{})
	d.attach(getAllGoods, &GetAllGoodsHandler{})
	d.attach(createGoods, &CreateGoodsHandler{})
	d.attach(updateGoods, &UpdateGoodsHandler{})
	d.attach(deleteGoods, &DeleteGoodsHandler{})

	d.attach(getList, &GetListHandler{})
	d.attach(getAllLists, &GetAllListHandler{})
	d.attach(createList, &CreateListHandler{})
	d.attach(updateList, &UpdateListHandler{})
	d.attach(deleteList, &DeleteListHandler{})
}

func (d *Dispatcher) attach(command string, handler WsHandlerInterface) {
	d.handlers[command] = handler
}

func (d *Dispatcher) CallHandler(command string, payload []byte) {
	handler, ok := d.handlers[command]
	if ok {
		handler.ProcessTask(payload)
	}
	logger.Error(errors.New(fmt.Sprintf("no handler provided for: %s", command)))
}

type WsHandlerInterface interface {
	ProcessTask(context []byte)
}

type BaseWsCommandHandler struct {
	services service.Service
}

// DISH WS HANDLERS

type GetDishHandler struct {
	BaseWsCommandHandler
}

func (h *GetDishHandler) ProcessTask(payload []byte) {
	dish, err := h.services.Dish.GetDish(1)
	fmt.Println(dish, err)
}

type GetAllDishesHandler struct {
	BaseWsCommandHandler
}

func (h *GetAllDishesHandler) ProcessTask(payload []byte) {
	h.services.Dish.GetAllDishes()
}

type CreateDishHandler struct {
	BaseWsCommandHandler
}

func (h *CreateDishHandler) ProcessTask(payload []byte) {
	h.services.Dish.CreateDish(&entities.Dish{})
}

type UpdateDishHandler struct {
	BaseWsCommandHandler
}

func (h *UpdateDishHandler) ProcessTask(payload []byte) {
	h.services.Dish.UpdateDish(&entities.Dish{})
}

type DeleteDishHandler struct {
	BaseWsCommandHandler
}

func (h *DeleteDishHandler) ProcessTask(payload []byte) {
	h.services.Dish.DeleteDish(1)
}

// GOODS WS HANDLERS

type GetGoodsHandler struct {
	BaseWsCommandHandler
}

func (h *GetGoodsHandler) ProcessTask(payload []byte) {
	h.services.Goods.GetGoods(1)
}

type GetAllGoodsHandler struct {
	BaseWsCommandHandler
}

func (h *GetAllGoodsHandler) ProcessTask(payload []byte) {
	h.services.Goods.GetAllGoods()
}

type CreateGoodsHandler struct {
	BaseWsCommandHandler
}

func (h *CreateGoodsHandler) ProcessTask(payload []byte) {
	h.services.Goods.CreateGoods(&entities.Goods{})
}

type UpdateGoodsHandler struct {
	BaseWsCommandHandler
}

func (h *UpdateGoodsHandler) ProcessTask(payload []byte) {
	h.services.Goods.UpdateGoods(&entities.Goods{})
}

type DeleteGoodsHandler struct {
	BaseWsCommandHandler
}

func (h *DeleteGoodsHandler) ProcessTask(payload []byte) {
	h.services.Goods.DeleteGoods(1)
}

// LIST WS HANDLERS

type GetListHandler struct {
	BaseWsCommandHandler
}

func (h *GetListHandler) ProcessTask(payload []byte) {
	h.services.List.GetList(1)
}

type GetAllListHandler struct {
	BaseWsCommandHandler
}

func (h *GetAllListHandler) ProcessTask(payload []byte) {
	h.services.List.GetAllLists()
}

type CreateListHandler struct {
	BaseWsCommandHandler
}

func (h *CreateListHandler) ProcessTask(payload []byte) {
	h.services.List.CreateList(&entities.List{})
}

type UpdateListHandler struct {
	BaseWsCommandHandler
}

func (h *UpdateListHandler) ProcessTask(payload []byte) {
	h.services.List.UpdateList(&entities.List{})
}

type DeleteListHandler struct {
	BaseWsCommandHandler
}

func (h *DeleteListHandler) ProcessTask(payload []byte) {
	h.services.List.DeleteList(1)
}
