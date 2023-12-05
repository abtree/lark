package utils

import "github.com/google/uuid"

//生成一个UUID
func NewUUID() string {
	return uuid.New().String()
}
