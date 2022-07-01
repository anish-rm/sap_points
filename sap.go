package main

import "fmt"

func main() {
	var a int
	for i := 0; i < 10; i++ {
		fmt.Println("--------------------------- Welcome to SAP Point Calculation -----------------------")
		fmt.Println("Choose what you want to calculate")
		fmt.Println("1.Paper Presentations")
		fmt.Println("2.Project Presentations")
		fmt.Println("3.Techno Managerial Events")
		fmt.Println("4.Social Activities")
		fmt.Println("5.Online Courses")
		fmt.Println("6.Sports")
		fmt.Println("7.Membership")
		fmt.Println("8.Entrepreneurship")
		fmt.Println("9.Calculate SAP")
		fmt.Println("10.Exit")
		fmt.Scan(&a)
		grand_total := 0
		var total [8]int
		switch a {
		case 1:
			fmt.Println("----- PAPER PRESENTATION -----")
			var p int

			var papersubmitted int
			var paperpresentedin int
			var paperpresentedout int
			var paperpresentedpr int
			var paperprizein int
			var paperprizeout int
			var paperprizepr int

			for i := 0; i < 5; i++ {

				fmt.Println("1.paper submitted")
				fmt.Println("2.paper presented")
				fmt.Println("3.prize winner")
				fmt.Println("4.exit")
				fmt.Println("Enter a option")
				fmt.Scan(&p)
				switch p {
				case 1:
					var ps int
					fmt.Println("Enter the number of papers submitted")
					fmt.Scan(&ps)

					papersubmitted = ps * 2

				case 2:
					var paprin int
					var paprout int
					var paprpr int
					fmt.Println("Enter how many papers in inside college ")
					fmt.Scan(&paprin)
					fmt.Println("Enter how many papers in outside college")
					fmt.Scan(&paprout)
					fmt.Println("Enter how many papers in premier institutions")
					fmt.Scan(&paprpr)

					paperpresentedin = paprin * 5
					paperpresentedout = paprout * 10
					paperpresentedpr = paprpr * 20
				case 3:
					var paprzin int
					var paprzout int
					var paprzpr int
					fmt.Println("Enter how many prizes won in papers presentation  in inside college ")
					fmt.Scan(&paprzin)
					fmt.Println("Enter how many prizes won in papers presentation  in outside college")
					fmt.Scan(&paprzout)
					fmt.Println("Enter how many prizes won in papers presentation  in premier institutions")
					fmt.Scan(&paprzpr)

					paperprizein = paprzin * 20
					paperprizeout = paprzout * 30
					paperprizepr = paprzpr * 50
				case 4:
					break
				default:
					fmt.Println("Invalid option")
				}
				if p == 4 {
					break
				}
			}
			var papersap int
			papersap = papersubmitted + paperpresentedin + paperpresentedout + paperpresentedpr + paperprizein + paperprizeout + paperprizepr
			if papersap < 75 {
				fmt.Printf("...............The total points in  paper presentation is %d................", papersap)
			} else {
				fmt.Println("The total points in paper presentation should be less than 75.")
			}
		case 2:
			fmt.Println("----- PROJECT PRESENTATION -----")

			var pr int

			var projectsubmitted int
			var projectpresentedin int
			var projectpresentedout int
			var projectpresentedpr int
			var projectprizein int
			var projectprizeout int
			var projectprizepr int

			for i := 0; i < 5; i++ {

				fmt.Println("1.project submitted")
				fmt.Println("2.project presented")
				fmt.Println("3.project winner")
				fmt.Println("4.exit")
				fmt.Println("Enter a option")
				fmt.Scan(&pr)
				switch pr {
				case 1:
					var prs int
					fmt.Println("Enter the number of project submitted")
					fmt.Scan(&prs)

					projectsubmitted = prs * 5

				case 2:
					var prprin int
					var prprout int
					var prprpr int
					fmt.Println("Enter how many project in inside college ")
					fmt.Scan(&prprin)
					fmt.Println("Enter how many project in outside college")
					fmt.Scan(&prprout)
					fmt.Println("Enter how many project in premier institutions")
					fmt.Scan(&prprpr)

					projectpresentedin = prprin * 10
					projectpresentedout = prprout * 15
					projectpresentedpr = prprpr * 20
				case 3:
					var prprzin int
					var prprzout int
					var prprzpr int
					fmt.Println("Enter how many prizes won in project presentation  in inside college ")
					fmt.Scan(&prprzin)
					fmt.Println("Enter how many prizes won in project presentation  in outside college")
					fmt.Scan(&prprzout)
					fmt.Println("Enter how many prizes won in project presentation  in premier institutions")
					fmt.Scan(&prprzpr)

					projectprizein = prprzin * 20
					projectprizeout = prprzout * 30
					projectprizepr = prprzpr * 50
				case 4:
					break
				default:
					fmt.Println("Invalid option")
				}
				if pr == 4 {
					break
				}
			}
			var projectsap int
			projectsap = projectsubmitted + projectpresentedin + projectpresentedout + projectpresentedpr + projectprizein + projectprizeout + projectprizepr
			if projectsap < 100 {
				fmt.Printf("...............The total points in  project presentation is %d................", projectsap)
			} else {
				fmt.Println("The total points in project presentation should be less than 100.")
			}

		case 3:
			var techno_count [8]int
			techno_pts := [8]int{2, 5, 10, 20, 10, 20, 30, 50}
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
			for j := 7; j >= 0; j-- {
				for k := techno_count[j]; k > 0; k-- {
					total[2] = total[2] + techno_pts[j]
					if total[2] > 75 {
						total[2] = total[2] - techno_pts[j]
					}
				}
			}
			grand_total = grand_total + total[2]
		case 4:
			var social_count [3]int
			social_pts := [3]int{5, 30, 50}
			fmt.Println("----- SOCIAL ACTIVITIES -----")
			fmt.Print("No.of BLOOD DONATIONS : ")
			fmt.Scan(&social_count[0])
			fmt.Print("No.of 1 to 2 weeks of NSS/NCC CAMP : ")
			fmt.Scan(&social_count[1])
			fmt.Print("No.of more than 2 weeks of NSS/NCC CAMP : ")
			fmt.Scan(&social_count[2])
			for j := 2; j >= 0; j-- {
				for k := social_count[j]; k > 0; k-- {
					total[3] = total[3] + social_pts[j]
					if total[3] > 50 {
						total[3] = total[3] - social_pts[j]
					}
				}
			}
			grand_total = grand_total + total[3]
		case 5:
			fmt.Println("----- ONLINE COURCES -----")
			var course_count[4]int
			course_pts:=[4]int{5,10,20,30}					   
			fmt.Print("VALUE ADDED COURSE: ")
			fmt.Scan(&course_count[0])
			fmt.Print("ONE CREDIT COURSES:")
			fmt.Scan(&course_count[1])
			fmt.Print("TWO CREDIT COURSES:")
			fmt.Scan(&course_count[2])
			fmt.Print("MORE THAN TWO CREDIT COURSES:")
			fmt.Scan(&course_count[3])
			for j :=3; j >= 0; j-- {
				for k:=course_count[j];k>0;k-- {
					total[4]=total[4]+course_pts[j]
					if(total[4]>100){
						total[4]=total[4]-course_pts[j]
				}
			}
		}
		grand_total=grand_total+total[4]
		case 6:
			var sports_count[8]int
			sports_pts:=[8]int{2,10,20,50,5,20,40,100}
			fmt.Println("----- SPORTS AND GAMES -----")
			fmt.Print("No.of INSIDE PARTICIPATION : ")
			fmt.Scan(&sports_count[0])
			fmt.Print("No.of ZONE/OUTSIDE PARTICIPATION : ")
			fmt.Scan(&sports_count[1])
			fmt.Print("No.of STATE/INTER LEVEL PARTICIPATION : ")
			fmt.Scan(&sports_count[2])
			fmt.Print("No.of NATIONAL/INTERNATIONAL LEVEL PARTICIPATION : ")
			fmt.Scan(&sports_count[3])
			fmt.Print("No.of INSIDE PRIZES : ")
			fmt.Scan(&sports_count[4])
			fmt.Print("No.of ZONE/OUTSIDE PRIZES : ")
			fmt.Scan(&sports_count[5])
			fmt.Print("No.of STATE/INTER LEVEL PRIZES : ")
			fmt.Scan(&sports_count[6])
			fmt.Print("No.of NATIONAL/INTERNATIONAL LEVEL PRIZES : ")
			fmt.Scan(&sports_count[7])
			for j :=7; j >= 0; j-- {
				for k:=sports_count[j];k>0;k-- {
					total[5]=total[5]+sports_pts[j]
					if(total[5]>100){
						total[5]=total[5]-sports_pts[j]
				}
			}
		}
		grand_total=grand_total+total[5]
		case 7:
			fmt.Println("----- MEMBERSHIP -----")
			var membership_count [3]int
			membership_pts := [3]int{20, 5, 2}
			fmt.Print("NCC&NSS: ")
			fmt.Scan(&membership_count[0])
			fmt.Print("Proffessional Society : ")
			fmt.Scan(&membership_count[1])
			fmt.Print("Clubs:")
			fmt.Scan(&membership_count[2])
			for j := 2; j >= 0; j-- {
				for k := membership_count[j]; k > 0; k-- {
					total[6] = total[6] + membership_pts[j]
					if total[6] > 50 {
						total[6] = total[6] - membership_pts[j]
					}
				}
			}
			grand_total = grand_total + total[6]
			
		case 8:
			fmt.Println("----- ENTERPRENEURSHIP -----")
			var entrepreneurship_count [3]int
			entrepreneurship_pts := [3]int{10, 50, 100}
			fmt.Print("No.of workshop attended" )
			fmt.Scan(&entrepreneurship_count[0])
			fmt.Print("Registered for startup: ")
			fmt.Scan(&entrepreneurship_count[1])
			fmt.Print("Product Released:")
			fmt.Scan(&entrepreneurship_count[2])
			for j := 2; j >= 0; j-- {
				for k := entrepreneurship_count[j]; k > 0; k-- {
					total[7] = total[7] + entrepreneurship_pts[j]
					if (total[7] > 100) {
						total[7] = total[7] - entrepreneurship_pts[j]
					}
				}
			}
			grand_total = grand_total + total[7]
			
		case 9:
			fmt.Println("----- CALCULATE SAP -----")
			fmt.Println(" Paper Presentations      : ",total[0])
			fmt.Println(" Project Presentations    : ",total[1])
			fmt.Println(" Techno Managerial Events : ",total[2])
			fmt.Println(" Social Activities        : ",total[3])
			fmt.Println(" Online Courses           : ",total[4])
			fmt.Println(" Sports                   : ",total[5])
			fmt.Println(" Membership               : ",total[6])
			fmt.Println(" Entrepreneurship         : ",total[7])
			fmt.Println("-- STUDENT ACTIVITY POINTS SCORED = ", grand_total," --")
			// UNABLE TO GET grand_total!!!! after calculating Techno & social activities
		case 10:
			fmt.Println("------------------------------------Exiting-----------------------------------------------------")
		default:
			fmt.Println("Invalid Choice")
		}
		if a == 10 {
			break
		}
	}
}
