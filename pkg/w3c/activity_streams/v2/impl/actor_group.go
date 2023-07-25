package impl

import (
	"fmt"
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/activity_streams/v2/vocab"
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
)

const ActorGroupTypeValue = "Group"

type ActorGroup struct {
	ActorProperties
	alias *string
}

func DeserializeActorGroup(d map[string]interface{}, ldAliases map[string]string) (vocab.ActorGroup, error) {
	alias := helpers.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	s, ok := helpers.GetType(d, alias)

	if !ok {
		return nil, fmt.Errorf("type is not defined")
	} else if *s != ActorGroupTypeValue {
		return nil, nil
	}

	ret := ActorGroup{}

	if err := deserializeActorProperties(d, ldAliases, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
