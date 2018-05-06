package main

import (
	"fmt"

	"github.com/CurtisVermeeren/GoRecipesAProblemSolvingApproach/1-beginning-go/sharedlibrarypackage"
)

func main() {
	// Using the shared library package
	str1 := "hello"
	fmt.Println("To Upper first", sharedlibrarypackage.ToFirstUpper(str1))
	fmt.Println("To Lower", sharedlibrarypackage.ToLowerCase("HeLlO"))
	fmt.Println("To Upper", sharedlibrarypackage.ToUpperCase(str1))
}
