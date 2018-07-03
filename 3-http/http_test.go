package main

import "testing"

func TestHttpGet(t *testing.T) {
	  get( []string {"https://www.google.com", "https://www.yahoo.com", "https://www.abc.dummy"} )
}
