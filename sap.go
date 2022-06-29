package main
import "fmt"
func main(){
	var a int
	for i:=0; i < 10; i++{
	fmt.Println("--------------------------- Welcome to SAP Point Calculation -----------------------")
	fmt.Println("Choose what you want to calculate")
	fmt.Println("1.Paper Presentation")
	fmt.Println("2.Project Presentation")
	fmt.Println("3.Techno Managerial Events")
	fmt.Println("4.Social Activities")
	fmt.Println("5.Online Courses")
	fmt.Println("6.Sports")
	fmt.Println("7.Membership")
	fmt.Println("8.Entrepreneurship")
	fmt.Println("9.Calculate SAP")
	fmt.Println("10.Exit")	
	fmt.Scan(&a)
	grand_total:=0
	var total[8]int
	switch a {
  	case 1:
    	fmt.Println("----- PAPER PRESENTATION -----")
		fmt.Println("fill ur code here")
  	case 2:
    	fmt.Println("----- PROJECT PRESENTATION -----")
		fmt.Println("fill ur code here")
  	case 3:
		var techno_count[8]int
		techno_pts:=[8]int{2,5,10,20,10,20,30,50}
		fmt.Println("----- TECHNO MANAGERIAL EVENTS -----")
		fmt.Print("No.of INSIDE PARTICIPATION : ")
		fmt.Scan(&techno_count[0])
		fmt.Print("No.of OUTSIDE PARTICIPATION : ")
		fmt.Scan(&techno_count[1])
		fmt.Print("No.of STATE LEVEL PARTICIPATION : ")
		fmt.Scan(&techno_count[2])
		fmt.Print("No.of NATIONAL/INTERNATIONAL LEVEL PARTICIPATION : ")
		fmt.Scan(&techno_count[3])
		fmt.Print("No.of INSIDE PRIZES : ")
		fmt.Scan(&techno_count[4])
		fmt.Print("No.of OUTSIDE PRIZES : ")
		fmt.Scan(&techno_count[5])
		fmt.Print("No.of STATE LEVEL PRIZES : ")
		fmt.Scan(&techno_count[6])
		fmt.Print("No.of NATIONAL/INTERNATIONAL LEVEL PRIZES : ")
		fmt.Scan(&techno_count[7])
		for j :=7; j >= 0; j-- {
			for k:=techno_count[j];k>0;k-- {
				total[2]=total[2]+techno_pts[j]
				if(total[2]>75){
					total[2]=total[2]-techno_pts[j]
				}
			}
		}
		grand_total=grand_total+total[2]
  	case 4:
    	var social_count[3]int
		social_pts:=[3]int{5,30,50}
		fmt.Println("----- SOCIAL ACTIVITIES -----")
		fmt.Print("No.of BLOOD DONATIONS : ")
		fmt.Scan(&social_count[0])
		fmt.Print("No.of 1 to 2 weeks of NSS/NCC CAMP : ")
		fmt.Scan(&social_count[1])
		fmt.Print("No.of more than 2 weeks of NSS/NCC CAMP : ")
		fmt.Scan(&social_count[2])
		for j :=2; j >= 0; j-- {
			for k:=social_count[j];k>0;k-- {
				total[3]=total[3]+social_pts[j]
				if(total[3]>50){
					total[3]=total[3]-social_pts[j]
				}
			}
		}
		grand_total=grand_total+total[3]
  	case 5:
    	fmt.Println("----- ONLINE COURCES -----")
		fmt.Println("fill ur code here")
  	case 6:
    	fmt.Println("----- SPORTS -----")
		fmt.Println("fill ur code here")
  	case 7:
    	fmt.Println("----- MEMBERSHIP -----")
		fmt.Println("fill ur code here")
	case 8:
    	fmt.Println("----- ENTERPRENEURSHIP -----")
		fmt.Println("fill ur code here")
	case 9:
    	fmt.Println("----- CALCULATE SAP -----")
		fmt.Println("STUDENT ACTIVITY POINTS SCORED = ",grand_total)
		// UNABLE TO GET grand_total!!!! after calculating Techno & social activities
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