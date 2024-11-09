package operations

import (
	"encoding/json"
	"errors"
	"log"
)

func NewOperation(Type string, Data string, User string, Rank int) error {
	operation := Operation{Type: Type, Data: Data, User: User, Rank: Rank}
	jsonOperation, err := json.Marshal(operation)
	if err != nil {
		return errors.New(EMO)
	}

	log.Println(jsonOperation)
	return nil
}
