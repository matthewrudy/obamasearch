package obamasearch

import (
	"reflect"
	"testing"
)

func Test_NewTwitterClient(t *testing.T) {
	api := NewTwitterClient()

	apiType := reflect.TypeOf(api)

	if apiType.String() != "*anaconda.TwitterApi" {
		t.Error("not a Twitter Api:", apiType)
	}
}
