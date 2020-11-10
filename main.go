package main

import (
	"TouchAll-Web/routers"
)

//var useMongodb = config.GetValue("mongodb.use").(bool)
//var useMysql = config.GetValue("mysql.use").(bool)

func main() {
	router := routers.NewRouter()
	router.Start()
}
