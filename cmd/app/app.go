package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"my_goods/internal/delivery/web"
	"my_goods/internal/delivery/web/dish_http"
	"my_goods/internal/delivery/web/goods_http"
	"my_goods/internal/delivery/web/lists_http"
	"my_goods/internal/service/dish_service"
	"my_goods/internal/service/goods_service"
	"my_goods/internal/service/lists_service"
	"my_goods/internal/storage/postgres/dish_repository"
	"my_goods/internal/storage/postgres/goods_repository"
	"my_goods/internal/storage/postgres/lists_repository"
	"my_goods/pkg/database"
	"my_goods/pkg/env"
)

type App struct {
	server *web.Server
	dbPool *pgxpool.Pool
}

func NewApplication() *App {
	if err := env.InitEnvVariables(); err != nil {
		log.Fatal(err)
	}
	db, err := database.NewDatabasePostgres()
	if err != nil {
		log.Fatalf("error initialising database: %s", err.Error())
	}
	app := &App{
		server: web.NewServer(),
		dbPool: db,
	}
	return app
}

func (a *App) Run() {
	defer a.dbPool.Close()
	a.server.Run()
}

func (a *App) InitApp() {
	api := a.server.Router.Group("/api")
	a.initLists(a.dbPool, api)
	a.initGoods(a.dbPool, api)
	a.initDish(a.dbPool, api)
}

func (a *App) initGoods(db *pgxpool.Pool, api *gin.RouterGroup) {
	repo := goods_repository.NewGoodsRepository(db)
	service := goods_service.NewGoodsService(repo)
	httpHandler := goods_http.NewGoodsHttpHandler(service)
	httpHandler.RegisterRoutes(api)
}

func (a *App) initDish(db *pgxpool.Pool, api *gin.RouterGroup) {
	repo := dish_repository.NewDishRepository(db)
	service := dish_service.NewDishService(repo)
	httpHandler := dish_http.NewDishHttpHandler(service)
	httpHandler.RegisterRoutes(api)
}

func (a *App) initLists(db *pgxpool.Pool, api *gin.RouterGroup) {
	repo := lists_repository.NewListRepository(db)
	service := lists_service.NewListService(repo)
	httpHandler := lists_http.NewListsHttpHandler(service)
	httpHandler.RegisterRoutes(api)
}
