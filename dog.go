package dog

import (
	"fmt"
	"strings"
)

func WhenGrowUp(s string) string {
	return "Saat puppy tumbuh dia berkata" + strings.ToUpper(s)
}


func Update() {
	fmt.Println("update version v1.1.0")
}