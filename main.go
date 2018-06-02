package main
import (
    "fmt"
    "sync"
)
const Filename = "test.ppm"

func main() {
	var params = map[string]int{
		"r":0,
		"g":0,
		"b":0,
	}
    //stack tester
    stack := Stack{&node{2,nil},sync.Mutex{}}
    fmt.Println(stack.Peek())
    stack.Push(1)
    fmt.Println(stack.Pop())
    fmt.Println(stack.Peek())

    //cstack tester
    worlds := MakeWorldStack()
    worlds.GetWorld().PrintMatrix()
    worlds.PushWorld()
    worlds.GetWorld().RotX(90)
    worlds.GetWorld().PrintMatrix()
    worlds.PopWorld()
    worlds.GetWorld().PrintMatrix()
	screen = MakeGrid(Width, Height)
    fmt.Println("MESSAGES THAT SAY ERROR IN FRONT OF THEM WERE COMPILER ERRORS, OTHER PRINTED COMMANDS ARE ONES THAT ARE INTERPRETED BY THE COMPILER BUT DO NOT DO MUCH OF ANYTHING YET")
    FillGrid(255,255,255)
    ParseFile("test.mdl",params)
    screen = MakeGrid(Width,Height)
    FillGrid(255,255,255)
    ParseFile("robot.mdl",params)
    screen = MakeGrid(Width,Height)
    FillGrid(0,0,0)
    ParseFile("script.mdl",params)

}
