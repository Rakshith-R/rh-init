package main

import "testing"

func TestCheatsheet(t *testing.T) {
	result := HelloWorld("m")
	if result != "Hello World !! m" {
		t.Errorf("String was incorrect")
	}
	if result != "Hello World !! k" {
		t.Errorf("String was incorrect")
	}
}
func Test2Cheatsheet(t *testing.T) {
	result := HelloWorld("m")
	if result != "Hello World !! m" {
		t.Errorf("String was incorrect")
	}
	if result != "Hello World !! k" {
		t.Errorf("String was incorrect")
	}
}
