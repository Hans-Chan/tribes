package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tribes/common/router"
)

func tribesHeader() {
	fmt.Println(
		`
   _|_|_|_|_|  _|_|_|    _|  _|_|_|    _|_|_|_|   _|_|_|
       _|      _|    _|  _|  _|    _|  _|        _|      
       _|      _|_|_|    _|  _|_|_|    _|_|_|_|    _|_|   
       _|      _|    _|  _|  _|    _|  _|              _| 
       _|      _|    _|  _|  _|_|_|    _|_|_|_|  _|_|_|    
	`)
}

func main() {
	// Set Variable
	r := httprouter.New()
	r.ServeFiles("/static/*filepath", http.Dir("view/asset"))
	tribesHeader()
	router.InitRoutes(r)
	log.Fatal(http.ListenAndServe(":8001", r))
}
