package main

import (
	"internal"
)

func main() {
	internal.LogTime("start")
	internal.LogTime("end")
	internal.DisplayProfileData()
}
