package main  
import (  
 "fmt"  
)  
 
//FUNCTION TO CHECK PRIME
 
func CheckPrime(number int) {  
 isPrime := true  
 if number == 0 || number == 1 {  
  fmt.Printf(" %d is not a  prime number\n", number)  
 } else {  
  for i := 2; i <= number/2; i++ {  
   if number%i == 0 {  
    fmt.Printf(" %d is not a  prime number\n", number)  
    isPrime = false  
    break  
   }  
  }  
  if isPrime == true {  
   fmt.Printf(" %d is a prime number\n", number)  
  }  
 }  
}  
func main() {  
 var number int  
 fmt.Print("Enter Number To check Prime:")  
 fmt.Scanln(&number)  
 CheckPrime(number)  
}  
