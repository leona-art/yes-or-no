package main

import (
	"context"
	"fmt"
	"gomini/infra"
)

func main() {
	ctx := context.Background()
	gomini, err := infra.NewGemini(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer gomini.Close()

	q := "それは楽しいですか？"
	gomini.GenerateText(ctx, q)

}
