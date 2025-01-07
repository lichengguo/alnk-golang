package main

// UUID 生成器

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuidStr, err := uuidGenerator()
	if err != nil {
		fmt.Println("uuidGenerator error:", err)
		return
	}

	fmt.Println("uuidStr:", uuidStr)

}

// uuidGenerator 生成 UUID
func uuidGenerator() (string, error) {
	uUIdObj, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uUIdObj.String(), nil
}
