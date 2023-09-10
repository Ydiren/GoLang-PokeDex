package main

import "os"

func commandExit(_ *string) error {
	os.Exit(0)
	return nil
}
