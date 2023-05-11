package main

import (
	_ "LinEngineRules/api"
	_ "LinEngineRules/api/exprrulerecord"
	_ "LinEngineRules/api/ruledatasource"
	"LinEngineRules/initdata"
	_ "LinEngineRules/initdata"
	_ "LinEngineRules/utils"
	"fmt"
	rest "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"os"
)

func main() {
	config := rest.Config{
		WebServices:                   restful.RegisteredWebServices(),
		APIPath:                       "/apiDocs.json",
		PostBuildSwaggerObjectHandler: initdata.EnrichSwaggerObject,
	}
	restful.DefaultContainer.Add(rest.NewOpenAPIService(config))
	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger-ui/dist"))))
	cors := restful.CrossOriginResourceSharing{
		//AllowedDomains: []string{"localhost"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
		CookiesAllowed: true,
		Container:      restful.DefaultContainer,
	}
	restful.DefaultContainer.Filter(cors.Filter)
	server := &http.Server{Addr: fmt.Sprintf(":" + os.Getenv("SERVER_PORT")), Handler: restful.DefaultContainer}
	log.Println("Server Running Listening:", os.Getenv("SERVER_PORT"))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(server.ListenAndServe())
		return
	}

}
