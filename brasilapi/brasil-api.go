package brasilapi

import (
	"encoding/json"
	"fmt"
	"multithreading-project/utils"
	"net/http"
)

const (
	// Accessed with https://brasilapi.com.br/api/cep/v1/{cep}
	BRASIL_API = "https://brasilapi.com.br/api/cep/v1"
)

type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func FetchFromBrasilAPI(cep string, resultChan chan<- utils.AddressResult) {
	url := fmt.Sprintf("%s/%s", BRASIL_API, cep)

	client := &http.Client{Timeout: utils.ONE_SECOND_TIMEOUT}
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return
	}

	var brasilResp BrasilAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&brasilResp)
	if err != nil {
		return
	}

	result := utils.AddressResult{
		Cep:          brasilResp.Cep,
		Street:       brasilResp.Street,
		Neighborhood: brasilResp.Neighborhood,
		City:         brasilResp.City,
		State:        brasilResp.State,
		API:          "BrasilAPI",
	}

	select {
	case resultChan <- result:
	default:
		// Discard
	}
}
