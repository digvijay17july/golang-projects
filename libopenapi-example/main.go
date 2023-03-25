package main

import (
	"fmt"
	"os"

	"github.com/pb33f/libopenapi"
)

type API struct {
	paths    []Path
	metaData MetaData
}
type Path struct {
	path   string
	method string
}
type MetaData struct {
	description string
	title       string
}

func main() {

	//Load the file or file content
	openApi3, _ := os.ReadFile("./openapi.json")

	// lets create a document to parse the content
	openApidDcument, err := libopenapi.NewDocument(openApi3)

	// catch the error
	if err != nil {
		panic(fmt.Sprintf("Error in creating the Open API Doc: %e", err))
	}
	apiSpec, errors := openApidDcument.BuildV3Model()

	if len(errors) > 0 {
		for i := range errors {
			fmt.Printf("Issue while reading Spec: %e\n", errors[i])
		}

	}
	paths := []Path{}
	api := &API{}
	metaData := &MetaData{}
	metaData.description = apiSpec.Model.Info.Description
	metaData.title = apiSpec.Model.Info.Title
	for pathName, _ := range apiSpec.Model.Paths.PathItems {
		path := &Path{}
		path.path = pathName
		paths = append(paths, *path)
	}
	api.metaData = *metaData
	api.paths = paths
	printPretty(*api)
}

func printPretty(api API) {
	fmt.Println("API Title is : " + api.metaData.title)
	fmt.Println("API Description is : " + api.metaData.title)
	for _, path := range api.paths {
		fmt.Println("API  Path: ", path)
	}
}
