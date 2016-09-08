package main
import "testing"
import (
	"fmt"
	"time"
)

func TestSomething(t *testing.T) {
	fmt.Println(logo)
	time.Sleep(time.Minute * 5)
}
