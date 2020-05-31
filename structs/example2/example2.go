package example2

type user struct {
	name   string
	age    int
	weigth float32
	high   int
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

func NewUser(u user) user {
	return user{
		name:   u.name,
		age:    u.age,
		weigth: u.weigth,
		high:   u.high,
	}
}

func NewProduct(p Product) Product {
	return Product{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
	}
}
