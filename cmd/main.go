package main

import (
	"fmt"
	"github.com/AshiishKarhade/url-shortner-go/pkg/hashing"
	"github.com/AshiishKarhade/url-shortner-go/pkg/snowflake"
	"log"
)

func main() {
	sf, err := snowflake.NewSnowflake(1)
	if err != nil {
		log.Fatalf("Failed to initialize snowflake: %v", err)
	}
	for i := 0; i < 10; i++ {
		nextID := sf.NextID()
		hashed := hashing.GenerateShortURL(nextID)
		fmt.Printf("Generated hash ID: %s\n", hashed)
	}
}
