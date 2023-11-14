package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nukkua/ra-chi/internal/mvc/router"
)

func main (){

	r:= router.SetupRouter();
	
	port:=  os.Getenv("PORT")
	fmt.Println("Initializing server");
	fmt.Println("Serving at port: "+port);

	if port == ""{
		port = "3000"
	}

	log.Fatal(http.ListenAndServe("0.0.0.0:" + port, r));
}
