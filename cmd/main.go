package main

import (
	"fmt"
	"multithreading-project/brasilapi"
	"multithreading-project/utils"
	"multithreading-project/viacep"
	"os"
	"time"
)

func main() {
	cep := os.Args[1]
	resultChan := make(chan utils.AddressResult)
	timeout := time.After(utils.ONE_SECOND_TIMEOUT)

	go brasilapi.FetchFromBrasilAPI(cep, resultChan)
	go viacep.FetchFromViaCEP(cep, resultChan)

	select {
	case result := <-resultChan:
		fmt.Println("Address found!")
		fmt.Printf("API: %s\n", result.API)
		fmt.Printf("CEP: %s\n", result.Cep)
		fmt.Printf("Street: %s\n", result.Street)
		fmt.Printf("Neighborhood: %s\n", result.Neighborhood)
		fmt.Printf("City: %s\n", result.City)
		fmt.Printf("State: %s\n", result.State)

	case <-timeout:
		fmt.Println("\nError: Timeout of 1 second exceeded")
		fmt.Println("No API responded within the time limit.")
	}
}
