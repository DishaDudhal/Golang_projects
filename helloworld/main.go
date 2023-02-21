package main
 
import "fmt"
 
type laptopSize float64
 
func (this laptopSize) getSizeOfLaptop() laptopSize {
    return this
}



