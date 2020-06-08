package main  
  
import (  
 "fmt"  
 "strconv"  
)  
  
func ReverseNumber(number int) int {  
 strNumber := strconv.Itoa(number)  
 reverseStrNumber := ""  
 for length := len(strNumber); length > 0; length-- {  
  reverseStrNumber += string(strNumber[length-1])  
 }  
 reverseNum, error := strconv.Atoi(reverseStrNumber)  
 if error != nil {  
  fmt.Println("Failure to cast String to int")  
 }  
 return reverseNum  
}  
  
func main() {  
 var number int  
 fmt.Print("Enter number to reverse ")  
 fmt.Scanln(&number)  
 fmt.Printf("Reverse of number(%d) is %d\n", number, ReverseNumber(number))  
  
}  
