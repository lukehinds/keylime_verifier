package config

import (
	"testing"
)

func TestMapMatch(t *testing.T) {
	c := GetConfig("cloud_verifier", "my_cert")

	if match["cloud_verifier"] != "my_cert" {
		t.Error("Section does not match")
	}
	if match["Key"] != "KEY_1" {
		t.Error("Key does not match")
	}
}
