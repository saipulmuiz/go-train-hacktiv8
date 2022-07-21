package main

import (
	"fmt"
)

func main() {
	var f float32 = 1.000025013580322265625
	var str = `
		SELECT 
			u.id, u.name, p.name, t.amount, t.created_at
		FROM users as u 
		JOIN transactions as t 
			ON t.user_id = u.id 
		JOIN products as p 
			ON p.id = t.product_id 
		WHERE u.id = 10 
		LIMIT 10
`
	fmt.Printf("%.5f\n", f)
	fmt.Println(str)
}
