package activity_streams_v2_impl

import (
	"fmt"
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/activity_streams/v2/vocab"
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/json_ld/v1"
)

const ActivityLikeTypeValue = "Like"

type ActivityLike struct {
	ActivityProperties
	alias *string
}

func DeserializeActivityLike(d map[string]interface{}, ldAliases map[string]string) (activity_streams_v2_vocab.ActivityLike, error) {
	alias := json_ld_v1.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	s, ok := json_ld_v1.GetType(d, alias)

	if !ok {
		return nil, fmt.Errorf("type is not defined")
	} else if *s != ActivityLikeTypeValue {
		return nil, nil
	}

	ret := ActivityLike{}

	if err := deserializeActivityProperties(d, ldAliases, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}