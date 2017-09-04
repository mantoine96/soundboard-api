package main

import (
	"testing"

	"gopkg.in/h2non/baloo.v1"
)

var test = baloo.New("http://localhost:9000")

func TestHomePage(t *testing.T) {
	test.Get("/status").
		Expect(t).
		Status(200).
		Type("json").
		JSON(map[string]string{"status": "ok"}).
		Done()
}
