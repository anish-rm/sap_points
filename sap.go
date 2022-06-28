package main
import "fmt"
func main(){
	var a int
	for i:=0; i < 10; i++{
	fmt.Println("---------------------------Welcome to SAP Point Calculation-----------------------")
	fmt.Println("Choose what you want to calculate")
	fmt.Println("1.Paper Presentation")
	fmt.Println("2.Project Presentation")
	fmt.Println("3.Techno Managerial Events")
	fmt.Println("4.Social Activity")
	fmt.Println("5.Online Courses")
	fmt.Println("6.Sports")
	fmt.Println("7.Membership")
	fmt.Println("8.Entrepreneurship")
	fmt.Println("9.Calculate SAP")
	fmt.Println("10.Exit")	
	fmt.Scan(&a)
	switch a {
  	case 1:
    	fmt.Println("You have chosen Paper Presentation")
		fmt.Println("fill ur code here")
  	case 2:
    	fmt.Println("Project Presentation")
		fmt.Println("fill ur code here")
  	case 3:
    	fmt.Println("Techno Managerial Events")
		fmt.Println("fill ur code here")
  	case 4:
    	fmt.Println("Social Activity")
		fmt.Println("fill ur code here")
  	case 5:
    	fmt.Println("Online Courses")
		fmt.Println("fill ur code here")
  	case 6:
    	fmt.Println("Sports")
		fmt.Println("fill ur code here")
  	case 7:
    	fmt.Println("Membership")
		fmt.Println("fill ur code here")
	case 8:
    	fmt.Println("Entrepreneurship")
		fmt.Println("fill ur code here")
	case 9:
    	fmt.Println("Calculate SAP")
		fmt.Println("fill ur code here")
	case 10:
		fmt.Println("------------------------------------Exiting-----------------------------------------------------")
	default:
		fmt.Println("Invalid Choice")
  	}
	if(a==10){
		break
	}
	}
}