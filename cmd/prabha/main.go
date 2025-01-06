package main

import (
	"pujaprabha/internal/config"
	"pujaprabha/internal/interfaces"
	_authHandler "pujaprabha/pkg/v1/auth/handler/http"
	"pujaprabha/pkg/v1/auth/middleware"
	_authRepo "pujaprabha/pkg/v1/auth/repository"
	_authUsecase "pujaprabha/pkg/v1/auth/usecase"

	_productHandler "pujaprabha/pkg/v1/product/handler"
	_productRepo "pujaprabha/pkg/v1/product/repository"
	_productUsecase "pujaprabha/pkg/v1/product/usecase"

	_categoryHandler "pujaprabha/pkg/v1/category/handler"
	_categoryRepo "pujaprabha/pkg/v1/category/repository"
	_categoryUsecase "pujaprabha/pkg/v1/category/usecase"

	_varientHandler "pujaprabha/pkg/v1/varient_product/handler"
	_varientRepo "pujaprabha/pkg/v1/varient_product/repository"
	_varientUsecase "pujaprabha/pkg/v1/varient_product/usecase"

	_cartHandler "pujaprabha/pkg/v1/cart/handler"
	_cartRepo "pujaprabha/pkg/v1/cart/repository"
	_cartUsecase "pujaprabha/pkg/v1/cart/usecase"

	_sliderHandler "pujaprabha/pkg/v1/slider/handler"
	_sliderRepo "pujaprabha/pkg/v1/slider/repository"
	_sliderUsecase "pujaprabha/pkg/v1/slider/usecase"

	_blogHandler "pujaprabha/pkg/v1/blog/handler"
	_blogRepo "pujaprabha/pkg/v1/blog/repository"
	_blogUsecase "pujaprabha/pkg/v1/blog/usecase"

	_blogCategoryHandler "pujaprabha/pkg/v1/blogcategory/handler"
	_blogCategoryRepo "pujaprabha/pkg/v1/blogcategory/repository"
	_blogCategoryUsecase "pujaprabha/pkg/v1/blogcategory/usecase"

	_reviewHandler "pujaprabha/pkg/v1/review/handler"
	_reviewRepo "pujaprabha/pkg/v1/review/repository"
	_reviewUsecase "pujaprabha/pkg/v1/review/usecase"

	_verifyPaymentHandler "pujaprabha/pkg/v1/verifyPayment/handler"
	_verifyPaymentRepo "pujaprabha/pkg/v1/verifyPayment/repository"
	_verifyPaymentUsecase "pujaprabha/pkg/v1/verifyPayment/usecase"

	_orderHandler "pujaprabha/pkg/v1/order/handler"
	_orderRepo "pujaprabha/pkg/v1/order/repository"
	_orderUsecase "pujaprabha/pkg/v1/order/usecase"

	_orderProductHandler "pujaprabha/pkg/v1/orderProduct/handler"
	_orderProductRepo "pujaprabha/pkg/v1/orderProduct/repository"
	_orderProductUsecase "pujaprabha/pkg/v1/orderProduct/usecase"

	"log"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"gorm.io/gorm"
)

type server struct {
	DB  *gorm.DB
	App *fiber.App
}
type usecases struct {
	productUsecase        interfaces.ProductUsecase
	authUsecase           interfaces.UserUsecase
	categoryUsecase       interfaces.CategoryUsecase
	varientProductUsecase interfaces.VarientProductUsecase
	cartUsecase           interfaces.CartUsecase
	sliderUsecase         interfaces.SliderUsecase
	blogUsecase           interfaces.BlogUsecase
	blogCategoryUsecase   interfaces.BlogCategoryUsecase
	reviewUsecase         interfaces.ReviewUsecase
	verifyPaymentUsecase  interfaces.VerifyPaymentUsecase
	orderUsecase          interfaces.OrderUsecase
	orderProductUsecase   interfaces.OrderProductUsecase
}

func main() {
	//var c *fiber.Ctx
	config.ConfigViper()
	db, err := config.ConfigDb()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "PujaPrabha",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	productRepo := _productRepo.New(db)
	authRepo := _authRepo.New(db)
	categoryRepo := _categoryRepo.New(db)
	varientRepo := _varientRepo.New(db)
	cartRepo := _cartRepo.New(db)
	sliderRepo := _sliderRepo.New(db)
	blogRepo := _blogRepo.New(db)
	blogCategoryRepo := _blogCategoryRepo.New(db)
	reviewRepo := _reviewRepo.New(db)
	verifyPaymentRepo := _verifyPaymentRepo.New(db)
	orderRepo := _orderRepo.New(db)
	orderProductRepo := _orderProductRepo.New(db)

	authMW := middleware.Protected()
	jwtMW := middleware.ValidateJWT(db, jwtware.Config{})

	usecases := &usecases{
		productUsecase:        _productUsecase.New(productRepo, categoryRepo),
		authUsecase:           _authUsecase.New(authRepo),
		categoryUsecase:       _categoryUsecase.New(categoryRepo),
		varientProductUsecase: _varientUsecase.New(varientRepo),
		cartUsecase:           _cartUsecase.New(cartRepo, varientRepo, productRepo),
		sliderUsecase:         _sliderUsecase.New(sliderRepo),
		blogUsecase:           _blogUsecase.New(blogRepo, blogCategoryRepo),
		blogCategoryUsecase:   _blogCategoryUsecase.New(blogCategoryRepo),
		reviewUsecase:         _reviewUsecase.New(reviewRepo),
		verifyPaymentUsecase:  _verifyPaymentUsecase.New(verifyPaymentRepo),
		orderUsecase:          _orderUsecase.New(orderRepo),
		orderProductUsecase:   _orderProductUsecase.New(orderProductRepo, orderRepo),
	}
	// server := &server{DB: db}

	authRoute := app.Group("/auth/v1/")
	protectedRoutes := app.Group("/api/v1/", authMW, jwtMW)

	// //this code will give output in 8080 port
	InitRoutes(app, authRoute, protectedRoutes, usecases)
	log.Fatal(app.Listen(":8080"))
}

func InitRoutes(app *fiber.App, authRoute fiber.Router, protectedRoutes fiber.Router, usecases *usecases) {
	// authHandler.AuthRoutes(app, c)
	_productHandler.New(authRoute, protectedRoutes, usecases.productUsecase)
	_authHandler.New(authRoute, protectedRoutes, usecases.authUsecase)
	_categoryHandler.New(authRoute, protectedRoutes, usecases.categoryUsecase)
	_varientHandler.New(authRoute, protectedRoutes, usecases.varientProductUsecase)
	_cartHandler.New(protectedRoutes, usecases.cartUsecase)
	_sliderHandler.New(authRoute, protectedRoutes, usecases.sliderUsecase)
	_blogHandler.New(authRoute, protectedRoutes, usecases.blogUsecase)
	_blogCategoryHandler.New(authRoute, protectedRoutes, usecases.blogCategoryUsecase)
	_reviewHandler.New(authRoute, protectedRoutes, usecases.reviewUsecase)
	_verifyPaymentHandler.New(protectedRoutes, usecases.verifyPaymentUsecase)
	_orderHandler.New(protectedRoutes, usecases.orderUsecase)
	_orderProductHandler.New(protectedRoutes, usecases.orderProductUsecase)
}
