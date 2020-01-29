package main

import "fmt"

type RequiredFlagIsEmpty string

func (e RequiredFlagIsEmpty) Error() string {
	return fmt.Sprintf("required flag %q is empty", string(e))
}
