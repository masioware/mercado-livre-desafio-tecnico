package distributioncenter

import (
	"fmt"
	"sync"

	config "github.com/masioware/mercado-livre-desafio-tecnico/order-api/config"
	dto "github.com/masioware/mercado-livre-desafio-tecnico/order-api/internal/model"
	http "github.com/masioware/mercado-livre-desafio-tecnico/order-api/pkg/httpclient"
)

const maxConcurrentRequests = 10

// result encapsula o retorno da consulta de um item a um CD.
type result struct {
	ItemID  int
	Centers []string
	Err     error
}

var DoRequestFunc = http.DoRequest

// RetrieveDistributionCenters consulta os centros de distribuição para cada item do pedido.
// Retorna dois mapas: um com os resultados válidos e outro com os erros por item.
func RetrieveDistributionCenters(order dto.OrderRequestDTO) (map[int][]string, map[int]string) {
	var (
		wg               sync.WaitGroup
		resultsChan      = make(chan result, len(order.Items))
		concurrencyLimit = make(chan struct{}, maxConcurrentRequests)
	)

	for _, item := range order.Items {
		wg.Add(1)

		go func(itemID int) {
			defer wg.Done()

			concurrencyLimit <- struct{}{}        // Acquire slot
			defer func() { <-concurrencyLimit }() // Release slot

			res := fetchDistributionCenter(itemID)
			resultsChan <- res
		}(item.ID)
	}

	wg.Wait()
	close(resultsChan)

	return organizeResults(resultsChan)
}

// fetchDistributionCenter realiza a chamada HTTP para um item e retorna o resultado.
func fetchDistributionCenter(itemID int) result {
	var apiResp dto.DistributionCenterResponseDTO

	err := DoRequestFunc(http.RequestOptions{
		Method:      "GET",
		URL:         config.GetDistributionCenterURL(),
		QueryParams: map[string]string{"itemId": fmt.Sprintf("%d", itemID)},
		Result:      &apiResp,
	})

	if err != nil {
		return result{ItemID: itemID, Err: err}
	}

	return result{ItemID: itemID, Centers: apiResp.DistributionCenters}
}

// organizeResults transforma o canal de resultados em dois mapas: um de sucesso e um de erro.
func organizeResults(resultsChan <-chan result) (map[int][]string, map[int]string) {
	distributionResults := make(map[int][]string)
	errors := make(map[int]string)

	for r := range resultsChan {
		if r.Err != nil {
			errors[r.ItemID] = r.Err.Error()
		} else {
			distributionResults[r.ItemID] = r.Centers
		}
	}

	return distributionResults, errors
}
