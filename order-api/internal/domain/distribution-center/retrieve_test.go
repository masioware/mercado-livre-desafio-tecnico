package distributioncenter

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	model "github.com/masioware/mercado-livre-desafio-tecnico/order-api/internal/model"
	httpclient "github.com/masioware/mercado-livre-desafio-tecnico/order-api/pkg/httpclient"
)

// substitui temporariamente a função real por um mock
var originalDoRequestFunc = DoRequestFunc

func teardownMock() {
	DoRequestFunc = originalDoRequestFunc
}

func TestOrganizeResults_ShouldSeparateSuccessAndErrors(t *testing.T) {
	results := make(chan result, 3)
	results <- result{ItemID: 1, Centers: []string{"CD1", "CD2"}}
	results <- result{ItemID: 2, Err: errors.New("timeout")}
	results <- result{ItemID: 3, Centers: []string{"CD3"}}
	close(results)

	success, errors := organizeResults(results)

	assert.Equal(t, 2, len(success))
	assert.Equal(t, []string{"CD1", "CD2"}, success[1])
	assert.Equal(t, []string{"CD3"}, success[3])

	assert.Equal(t, 1, len(errors))
	assert.Equal(t, "timeout", errors[2])
}

func TestRetrieveDistributionCenters_WithMixedResults(t *testing.T) {
	defer teardownMock()

	DoRequestFunc = func(options httpclient.RequestOptions) error {
		itemID := options.QueryParams["itemId"]
		if itemID == "2" {
			return fmt.Errorf("mocked error")
		}

		res := options.Result.(*model.DistributionCenterResponseDTO)
		res.DistributionCenters = []string{"CD_TEST"}
		return nil
	}

	order := model.OrderRequestDTO{
		Items: []model.ItemDTO{
			{ID: 1, Name: "Item A", Price: 10.0},
			{ID: 2, Name: "Item B", Price: 20.0},
		},
	}

	success, errors := RetrieveDistributionCenters(order)

	assert.Len(t, success, 1)
	assert.Equal(t, []string{"CD_TEST"}, success[1])
	assert.Len(t, errors, 1)
	assert.Contains(t, errors[2], "mocked error")
}
