package main

import (
	"encoding/json"
	"fmt"
	"github.com/cleberson/enumeration/stringer_generate/model"
	"log"
)

func main() {
	results := []model.Result{
		{
			Action:     "save",
			StatusCode: model.Success,
		},
		{
			Action:     "get",
			StatusCode: model.NoContent,
		},
		{
			Action:     "update",
			StatusCode: model.BadRequest,
		},
		{
			Action:     "put",
			StatusCode: model.NotFound,
		},
		{
			Action:     "post",
			StatusCode: model.InternalServerError,
		},
	}

	bytes, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	payload := string(bytes)
	fmt.Println(payload)
	fmt.Println(results)

	fmt.Println("=================Unmarshal - decode =================")

	var httpStatusCodes []model.Result
	err = json.Unmarshal([]byte(payload), &httpStatusCodes)
	if err != nil {
		log.Fatal(err)
	}
	for _, element := range httpStatusCodes {
		fmt.Printf("%v\n", element.StatusCode)
	}

}
