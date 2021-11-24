package server

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"my_goods/internal/auth"
	"my_goods/internal/dish"
	"my_goods/internal/goods"
	"my_goods/internal/list"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func Router(db *gorm.DB) *gin.Engine {
	handler := gin.New()
	handler.Use(gin.Logger())
	handler.Use(cors.Default())

	dishRepo := dish.NewDishRepo(db)
	dishService := dish.NewDishService(dishRepo)
	dishHandler := dish.NewDishHandler(dishService)

	goodsRepo := goods.NewGoodsRepo(db)
	goodsService := goods.NewGoodsService(*goodsRepo)
	goodsHandler := goods.NewGoodsHandler(goodsService)

	listRepo := list.NewListRepo(db)
	listService := list.NewListService(*listRepo)
	listHandler := list.NewListHandler(listService)

	authRepo := auth.NewAuthRepo(db)
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	a := handler.Group("/auth")
	{
		a.POST("/sign-in", authHandler.SignIn)
		a.POST("/sign-up", authHandler.SignUp)
	}
	api := handler.Group("/api", authHandler.AuthMiddleware)
	{
		d := api.Group("/dish")
		{
			d.GET("/:id", dishHandler.GetDish)
			d.GET("/", dishHandler.GetAllDishes)
			d.POST("/", dishHandler.CreateDish)
			d.PUT("/", dishHandler.UpdateDish)
			d.DELETE("/:id", dishHandler.DeleteDish)
		}
		g := api.Group("/goods")
		{
			g.GET("/:id", goodsHandler.GetGoods)
			g.GET("/", goodsHandler.GetAllGoods)
			g.POST("/", goodsHandler.CreateGoods)
			g.PUT("/", goodsHandler.UpdateGoods)
			g.DELETE("/:id", goodsHandler.DeleteGoods)
		}
		l := api.Group("/list")
		{
			l.GET("/:id", listHandler.GetList)
			l.GET("/", listHandler.GetAllLists)
			l.POST("/", listHandler.CreateList)
			l.PUT("/:id", listHandler.UpdateList)
			l.DELETE("/:id", listHandler.DeleteList)
		}
	}
	return handler
}
