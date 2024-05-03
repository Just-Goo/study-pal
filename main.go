/*
Copyright © 2024 Just-Goo
*/
package main

import (
	"github.com/Just-Goo/study-pal/cmd"
	"github.com/Just-Goo/study-pal/data"
)

func main() {
	data.OpenDatabase()
	
	cmd.Execute()
}
