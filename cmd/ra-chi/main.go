package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nukkua/ra-chi/internal/app/router"
)

func main (){

	r:= router.SetupRouter();
	
	port:= os.Getenv("PORT");
	fmt.Println("Initializing server");
	fmt.Println("Serving at port: "+port);

	if port == ""{
		port = "3000"
	}


	log.Fatal(http.ListenAndServe("localhost:" + port, r));
}
