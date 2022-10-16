package main

import (
	"fmt"
	namingV1 "github.com/nomospace/baby-naming/v1"
	namingV2 "github.com/nomospace/baby-naming/v2/naming"
	namingV3 "github.com/nomospace/baby-naming/v3/naming"
)

func main() {
	name := namingV1.CreateBabyName()
	fmt.Println("Print baby name: ", name)

	// v2 
	fmt.Println("====")
	fmt.Println("Print baby name:", namingV2.CreateBabyName(true))
	fmt.Println("Print baby name:", namingV2.CreateBabyName(false))

	// v3
	fmt.Println("====")
	boy, boyErr := namingV3.CreateBabyName(true, 4)
	if boyErr == nil {
		fmt.Printf("Print baby(min length %v): %v\n", 4, boy)
	}
	girl, girlErr := namingV3.CreateBabyName(false, 3)
	if girlErr == nil {
		fmt.Printf("Print baby(min length %v): %v\n", 3, girl)
	}

}