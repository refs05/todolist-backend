package main

import (
	"log"
	"time"

	_routes "todo/app/routes"
	_todosUsecase "todo/business/todos"
	_todosController "todo/controller/todos"
	_todosRepo "todo/driver/database/todos"
	mysql_driver "todo/driver/mysql"

	"github.com/labstack/echo/v4"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}	
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_todosRepo.Todo{})
}

func main() {
	configDB := mysql_driver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitialDB()
	dbMigrate(db)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()
	// e.Use(middleware.CORS())

	todoRepo := _todosRepo.NewTodosRepository(db)
	todoUsecase := _todosUsecase.NewTodosUsecase(todoRepo, timeoutContext)
	todoCtrl := _todosController.NewTodosController(todoUsecase)

	routesInit := _routes.ControllerList{
		TodosController: *todoCtrl,
	}
	routesInit.RouteList(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}