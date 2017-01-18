package main

/**
 * Instrucciones
 * $ export GOPATH=$HOME/go
 * $ go get -u github.com/labstack/echo
 * $ go get -u github.com/labstack/echo/middleware
 * $ go get -u github.com/jinzhu/gorm
 * $ go get -u github.com/go-sql-driver/mysql
 */
import (
    "log"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"

    "./models"
)


func main() {
    i := Impl{}
    i.InitDB()
    // i.InitSchema() // Enable AutoMigrate

	// Echo instance
	e := echo.New()

	// Middleware
    e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {	
        var pedimentos []models.Pedimentos
        err := i.DB.Find(&pedimentos).Error

        if err != nil {
            log.Fatalf("Got error when query, the error is '%v'", err)
            return c.JSON(404, err)
        }

        return c.JSON(200, pedimentos)
	})
	// Start server
	e.Logger.Fatal(e.Start(":8001"))
}

func (i *Impl) InitDB() {
    var err error
    i.DB, err = gorm.Open("mysql", "USER:PASSWORD@tcp(IP_SERVER:3306)/DATABASE?charset=utf8&parseTime=True")
    if err != nil {
        log.Fatalf("Got error when connect database, the error is '%v'", err)
    }

    i.DB.LogMode(true)
    i.DB.SingularTable(true)
}

func (i *Impl) InitSchema() {
    i.DB.AutoMigrate(&models.Pedimentos{}, &models.Departamento{})
}

type Impl struct {
    DB *gorm.DB
}