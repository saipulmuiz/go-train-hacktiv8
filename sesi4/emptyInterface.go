package main

func main() {
	// -- Empty interface Example --
	var randomValues interface{}
	_ = randomValues
	randomValues = "Jalan Sudirman"
	randomValues = 20
	randomValues = true
	randomValues = []string{"Airell", "Nanda"}

	// -- Empty interface (Type assertion) --
	var v interface{}
	v = 20
	// v = v * 9 // Error

	if value, ok := v.(int); ok == true {
		v = value * 9
	}

	// -- Empty interface (Map & Slice with Empty Interface) --
	rs := []interface{}{1, "Airell", true, 2, "Ananda", true}

	rm := map[string]interface{}{
		"Name":   "Airell",
		"Status": true,
		"Age":    23,
	}

	_, _ = rs, rm

}
