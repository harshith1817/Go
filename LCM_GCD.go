package main
import ("fmt")
func gcd(num1 int, num2 int) int {
if num2 == 0 {
return num1
} return gcd(num2, num1%num2)
}func lcm(num1 int, num2 int) int {
return (num1 * num2) / gcd(num1, num2)
}func main() {
var num1, num2, choice int
fmt.Println("Enter number one:")
fmt.Scanln(&num1)
fmt.Println("Enter number two:")
fmt.Scanln(&num2)
fmt.Println("Enter choice: 1 for GCD and 2 for LCM")
fmt.Scanln(&choice)
switch choice {
case 1:
fmt.Printf("GCD of %d and %d is %d\n", num1, num2, 
gcd(num1, num2))
case 2:
fmt.Printf("LCM of %d and %d is %d\n", num1, num2, 
lcm(num1, num2))
default:
fmt.Println("Invalid choice. Please enter 1 for GCD or 2 for 
LCM.")
}