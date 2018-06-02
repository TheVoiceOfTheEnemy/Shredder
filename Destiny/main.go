package main

import (
	"sort"

	"time"
	"math/rand"
    "fmt"

)



var Path = "C:/Users/Ramso/Desktop/Shredder/logs/"


var AffordableReps = 6
var NoReturnLimit = 8

//real scenery is 180 turns/hour -> 7 hours playtime -> 1260
var interations = 90
var betIndexes []int

var Apuesta = 7
var ApuestasAlDia = Apuesta * 2 
var MaximaPerdida = -889//- 128//-635 //-889
var ColateralLoses = 1 // cada hora tirando 0.1 euros pierdo 1 euro por simple deriva

func main() {

	rand.Seed(time.Now().UnixNano())
	
	
	var money int
	money=0
	for i := 0; i < 30*12; i++ {
	var wii = randInt(0,1000)
	money = money + doRun(int64(wii))
	}
	fmt.Println("Results!!! Year money:", money,"Expected:", ApuestasAlDia*30*12)
	
	
	var longmoney []int
	for k := 0; k < 30; k++ {
	var money2 int
	money2=0
	for j := 0; j < 30*12; j++ {
	var wii = randInt(0,1000)
	money2 = money2 + doRun(int64(wii))
	}
	longmoney = append(longmoney,money2)
	}
	fmt.Println("Results!!! ", longmoney)

}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}


func doRun(seed int64) int{

	fmt.Println("-----------Day Simulation-------------")
	
	rand.Seed(time.Now().UnixNano()*seed)

	var table1 []int
	var table2 []int

	var serieCounter int
	var destiny []int

	
	

	/*
	var redCounter int
	var blackCounter int
	var lastColor int

	var redStrikes []int
	var blackStrikes []int*/

	
	
	
	for i := 0; i < interations; i++ {
		number := randInt(0,36)
		if number == 0{
			table1 = append(table1,0)
		}else if number%2 == 0{
			table1 = append(table1,1)
		}else{
			table1 = append(table1,2)
		}
	}
	for i := 0; i < interations; i++ {
		number := randInt(0,36)

		if number == 0{
			table2 = append(table2,0)
		}else if number%2 == 0{
			table2 = append(table2,1)
		}else{
			table2 = append(table2,2)
		}
	}
	//fmt.Println(table1)
	//fmt.Println(table2)

	for i := 0; i < interations-1; i++ {
		//if (table1[i]==table2[i] || table1[i]==0 || table2[i]==0) && (table1[i]==table1[i+1] || table1[i]==0 || table1[i+1]==0) && (table2[i]==table2[i+1] || table2[i]==0 || table2[i+1]==0) {
			if (table1[i]==table2[i] || table1[i]==0 || table2[i]==0) && (table1[i]==table1[i+1] || table1[i]==0 || table1[i+1]==0) && (table1[i]==table2[i+1] || table1[i]==0 || table2[i+1]==0) && (table2[i]==table2[i+1] || table2[i]==0 || table2[i+1]==0)&& (table2[i]==table1[i+1] || table2[i]==0 || table1[i+1]==0) {
			
				//fmt.Println(table1[i],table2[i])
				serieCounter++
		}else{
			//fmt.Println("seriecounter>",serieCounter)
			if serieCounter>AffordableReps{
				destiny = append(destiny,serieCounter)
				betIndexes = append(betIndexes,i-serieCounter)
				
				/*
				fmt.Println("From index",i-serieCounter)
				for x := i-serieCounter; x <= i; x++ {
					fmt.Println(table1[x],table2[x])
					}*/

				if table2[i+1]==table2[i]{  //series finished, but does my table go further with the serie?
					furtherCounter:=0
					for j:=i; j<interations;j++{
						if table2[j+1]==table2[j] || table2[j+1]==0 || table2[j]==0 {
							furtherCounter++
						}else{
							if serieCounter+furtherCounter > NoReturnLimit{
								fmt.Println("No return found", serieCounter+furtherCounter, "at index",i)
								/*
								for l := 1; l <= furtherCounter; l++ { //check wich range of indexes are not betable
									betIndexes = append(betIndexes,i+l)
								}*/
							}
							
							break
						}
					}
				}
			}
			serieCounter=0
		}
	}

	//loop finished, append last serie or not?
	if serieCounter>5{
		destiny = append(destiny,serieCounter)
	}

	//fmt.Println("table1")
	//checkTableStrikes(table1)
	fmt.Println("table2")
	var money int
	money = checkTableStrikes(table2)

	fmt.Println("Repetitions between tables found:", destiny, "Total", len(destiny), "Hit chance:",float64(len(destiny))/float64(interations))
	fmt.Println("Critical indexes:",betIndexes)
	fmt.Println("Hit chance:",float64(len(betIndexes))/float64(interations))


	betIndexes = nil

	//Debug(table1,table2)

	return money
}
func checkTableStrikes(table []int) int {
	var redCounter int
	var blackCounter int
	var lastColor int
	var redStrikes []int
	var blackStrikes []int
	var redBetIndexes []int
	var blackBetIndexes []int
		//check how many red and black strikes occur in table 2
		for i := 0; i < interations; i++ {
		
			var color = table[i]
			
			if color==1 || color==0{
				
				redCounter++
			
			}
			lastColor=color
	
			if lastColor==2{
				if redCounter > AffordableReps{
					redStrikes=append(redStrikes,redCounter)
					redBetIndexes=append(redBetIndexes,i-redCounter)
				}
				redCounter=0
				
			}
			
		}
	
	
	if redCounter > AffordableReps{
		redStrikes=append(redStrikes,redCounter)
		redBetIndexes=append(redBetIndexes,interations-redCounter)
	}
	
	
	for i := 0; i < interations; i++ {
			
		var color = table[i]
	
	if color==2 || color==0{
		blackCounter++
		
	}
	lastColor=color
	
	if lastColor==1{
		if blackCounter > AffordableReps{
			blackStrikes=append(blackStrikes,blackCounter)
			blackBetIndexes=append(blackBetIndexes,i-blackCounter)
		}
		blackCounter=0
		
	}
	
	}
	if blackCounter > AffordableReps{
		blackStrikes=append(blackStrikes,blackCounter)
		blackBetIndexes=append(blackBetIndexes,interations-blackCounter)
	}

	var tits []int
	tits = append(redBetIndexes,blackBetIndexes...)
	sort.Ints(tits[:])
	
	

	
	fmt.Println("redStrikes:",len(redStrikes))
	fmt.Println("indexes:",redBetIndexes)
	fmt.Println("blackStrikes:",len(blackStrikes))
	fmt.Println("indexes:",blackBetIndexes)

	var pick1 int
	var pick2 int
	if len(redBetIndexes)>1{
	pick1 = redBetIndexes[randInt(0,len(redBetIndexes)-1)]
	}else if len(redBetIndexes)==1{
		pick1 = redBetIndexes[0]
	}else{
		pick1 = interations-8
	}
	if len(blackBetIndexes)>1{
	pick2 = blackBetIndexes[randInt(0,len(blackBetIndexes)-1)]
	}else if len(blackBetIndexes)==1{
		pick2 = blackBetIndexes[0]
	}else{
		pick2 = 0
	}
	pick3 :=0

	fmt.Println("indexes:",pick1,pick2)

	for h := 0; h < len(betIndexes); h++ {
		if pick1==betIndexes[h] || pick2==betIndexes[h]  || pick3==betIndexes[h]{
			fmt.Println("TOTAL FAILURE")
			return MaximaPerdida
		}
	}


	return ApuestasAlDia - ColateralLoses
}

func Debug(table1 []int,  table2 []int){

	//reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter index: ")
	//index, _ := reader.ReadString('\n')
	//fmt.Println(index)
	//indexInt, _ := strconv.Atoi(index)
	var indexInt int
	fmt.Scan(&indexInt)

	fmt.Println("Debug index",indexInt)
	for i := indexInt; i <= indexInt+AffordableReps; i++ {
	fmt.Println(table1[i],table2[i])
	}

}