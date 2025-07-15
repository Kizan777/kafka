package key_generator

import "github.com/google/uuid"

func GenerateUUIDString(numberOfKeys int) []string {
	var uuids []string
	for i := 0; i < numberOfKeys; i++ {
		uuids = append(uuids, uuid.NewString())
	}
	return uuids
}
