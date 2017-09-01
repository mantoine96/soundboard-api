package soundboardApi

import (
	"testing"

	"gopkg.in/h2non/baloo.v1"
)

var test = baloo.New("http://localhost:8080")

func TestHomePage(t *testing.T) {
	test.Get("/").
		Expect(t).
		Status(200).
		Type("json").
		JSON(map[string]string{"message": "coucou"}).
		Done()
}
