package pointerpkg

func dereferencing {
	var count1 *int 	// nil pointer
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}


	if count1 != nil {
		fmt.Printf("Count1: %#v\n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}
}
