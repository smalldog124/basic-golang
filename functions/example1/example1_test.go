package example1_test

import (
	"basic/functions/example1"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateVat_total_price_100_percen_7_should_be_7(t *testing.T) {
	expectedVat := 70.00
	totalPrice := 1000.00
	percen := 7.00

	actualVat := example1.CalculateVat(totalPrice, percen)

	assert.Equal(t, expectedVat, actualVat)
}

func Test_JsonProduct_shuold_be_product_json_format(t *testing.T) {
	expected := `{"id":1,"name":"macbook pro","price":39000}`
	p := example1.Product{
		ID:    1,
		Nmae:  "macbook pro",
		Price: 39000,
	}

	actual, _ := example1.JsonProduct(p)

	assert.Equal(t, expected, actual)
}
