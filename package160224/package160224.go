package package160224

import "fmt"

func runprogram() {
	b2 := [4]bool{false, false,false,false}
	sum := 1
	for sum <= 10 {
	
		fmt.Println("Please select an option:")
		fmt.Println("1 – Print Covid19 cases in Pakistan") 
		fmt.Println("2 – Print Covid19 cases in SouthKorea")
		fmt.Println("3 – Print Covid19 cases in France")
		fmt.Println("4 – Print my personalized message to Coronavirus")
		fmt.Println("0 – Exit")
		var i int
			
		fmt.Scanf("%d", &i)
		fmt.Println(i)
		
		if i==0 {
			check:= true
			for j := 0; j < 4; j++ {
					if b2[j]==false{
						check=false
					}
			}
			if check==true{
			    fmt.Println("All options marked, termintaing program")
		        break
			}
			if check == false{
			    fmt.Println("Cant quit the program, view all options first\n\n")
			}
				
		}
		if i==1{
			fmt.Println("Total cases in Pakistan are 1523\n\n")
			b2[0]=true
		}
		if i==2 {
			fmt.Println("Total cases in South Korea are 9000\n\n")
		    b2[1]=true
		}
		if i==3 {
			fmt.Println("Total cases in France are 15000\n\n")
		    b2[2]=true
		}
		if i==4 {
			fmt.Println("Go, CORONA, GO\n\n")
		    b2[3]=true
		}
	}
}

