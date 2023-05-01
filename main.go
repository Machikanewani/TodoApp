package main

import (
	"Users/xushao/Desktop/Apps/TodoApp/app/controllers"
	"Users/xushao/Desktop/Apps/TodoApp/app/models"
	_ "Users/xushao/Desktop/Apps/TodoApp/config"
	"fmt"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}
