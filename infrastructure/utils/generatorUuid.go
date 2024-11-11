package uuids

import "github.com/google/uuid"

func GenerateUuid() any {
	return uuid.New().String()
}
