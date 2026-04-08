package viacep

import (
	"encoding/json"
	"fmt"
	"multithreading-project/utils"
	"net/http"
)

const (
	// Accessed with http://viacep.com.br/ws/{cep}/json/
	VIACEP_API = "http://viacep.com.br/ws"
)

type ViaCEPResponse struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func FetchFromViaCEP(cep string, resultChan chan<- utils.AddressResult) {
	url := fmt.Sprintf("%s/%s/json/", VIACEP_API, cep)

	client := &http.Client{Timeout: utils.ONE_SECOND_TIMEOUT}
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return
	}

	var viaResp ViaCEPResponse
	err = json.NewDecoder(resp.Body).Decode(&viaResp)
	if err != nil {
		return
	}

	if viaResp.Cep == "" {
		return
	}

	result := utils.AddressResult{
		Cep:          viaResp.Cep,
		Street:       viaResp.Logradouro,
		Neighborhood: viaResp.Bairro,
		City:         viaResp.Localidade,
		State:        viaResp.Uf,
		API:          "ViaCEP",
	}

	select {
	case resultChan <- result:
	default:
		// Discard if channel already received another result
	}
}
