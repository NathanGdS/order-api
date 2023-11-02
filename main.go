package main

import (
	"encoding/json"
	"fmt"

	"github.com/nathangds/order-api/models"
)

func main() {
	var category = models.NewCategory("1", "description", models.CLOTHES)
	res1B, _ := json.Marshal(models.CreateCategoryResponseDto(category))
	fmt.Println(string(res1B))
}
