package example2_test

import (
	"basic/structs/example2"
	"testing"
)

func Test_call_NewUser_sholde_be_show_new_user(t *testing.T) {
	jone := example2.user{
		name:   "jone",
		age:    25,
		weigth: 70.00,
		high:   183,
	}

	newUser := example2.NewUser(jone)

	if newUser != jone {
		t.Error("fail")
	}
}

func Test_call_NewProduct_sholde_be_show_new_product(t *testing.T) {
	iPhone := example2.Product{
		ID:    1,
		Name:  "iPhone 8",
		Price: 32900.00,
	}

	newProduct := example2.NewProduct(iPhone)

	if newProduct != iPhone {
		t.Error("fail")
	}
}
