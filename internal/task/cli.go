package task

import(
	"os"
)


func getArgs() []string{
	argsWithoutProg := os.Args[1:]
	return argsWithoutProg
}

