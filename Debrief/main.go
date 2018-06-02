package main

import (
	"sort"
	"math"
	"time"
	"math/rand"
    "bufio"
    "fmt"
    "log"
    "os"
)

var Path = "C:/Users/Ramso/Desktop/Shredder/logs/"



var RoundSize int = 90

var MartingalaLimit = 7


func main() {

	var totalReds int
	var totalBlacks int

	var redCounter int
	var blackCounter int
	var lastColor string



	var redStrikes []int
	var redStrikesIndex []int
	var blackStrikes []int
	var blackStrikesIndex []int

	var CritIndexesPerRound []int


	//var CritDistribution []
	var concatenatedStrikes []ConcatenatedStrike

	var totalRedStrikesRecord int
	var totalBlackStrikesRecord int

	var turns int
	
	files := []string{"jan1","jan2","jan3","jan4","feb1","feb2","feb3","mai1","mai2","mai3","mai4",
		"mai5","mar1","mar2","mar3","mar4","apr1","apr2","apr3","jun1","jun2","jun3","jun4","jul1","jul2","jul3","aug1","aug2","aug3"}
	
	var totalwins int
	var totalloses int
	
	for ifile,file := range files{

	fmt.Println("---------------",file,"---------------------")
	var wins int
	var loses int
	
	
    file, err := os.Open(Path+file+".txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	
	var CritIndexInRound int  //counts how many indexes in the actual round lead to fail

    for scanner.Scan() {
		var checkString = scanner.Text()
		if checkString != "1" && checkString != "2" && checkString != "0"{
			if checkString=="Ganancia"{
				wins++
				totalwins++
			}else if checkString=="Perdida"{
				loses++
				totalloses++
			}
			
			continue
		}
		var color = scanner.Text()


		
		if color=="1" || (lastColor=="1" && color=="0"){

			
			if color=="1"{
				totalReds++
			}
			
			if lastColor=="2" {
				
				if redCounter >= MartingalaLimit{
					redStrikes=append(redStrikes,redCounter)
					redStrikesIndex = append(redStrikesIndex,turns-redCounter)
					CritIndexInRound += redCounter-MartingalaLimit+1
				}
				redCounter=0
				
			}

			if blackCounter>0{
				
				if blackCounter >= MartingalaLimit{
					blackStrikes=append(blackStrikes,blackCounter)
					blackStrikesIndex = append(blackStrikesIndex,turns-blackCounter)
					CritIndexInRound += blackCounter-MartingalaLimit+1
				}
				blackCounter=0
			}
			redCounter++ //the position here matters
		}else if color=="2" || (lastColor=="2" && color=="0"){
			
			if color=="2"{
				totalBlacks++
			}
			
			if lastColor=="1"{
				if blackCounter >= MartingalaLimit{
					blackStrikes=append(blackStrikes,blackCounter)
					blackStrikesIndex = append(blackStrikesIndex,turns-blackCounter)
					CritIndexInRound += blackCounter-MartingalaLimit+1
				}
				blackCounter=0
			}

			if redCounter>0{
				if redCounter >= MartingalaLimit{
					redStrikes=append(redStrikes,redCounter)
					redStrikesIndex = append(redStrikesIndex,turns-redCounter)
					CritIndexInRound += redCounter-MartingalaLimit+1
				}
				redCounter=0
			}
			blackCounter++ //the position here matters
		}
		lastColor=color

		if turns>0 && turns%RoundSize==0{
			CritIndexesPerRound = append(CritIndexesPerRound,CritIndexInRound)
			CritIndexInRound=0 //reset counter
		}
		//test for printing , do not delete
		/*
		if turns>=236 && turns<260{
			fmt.Println(turns,color,redCounter,blackCounter)
		}*/

		turns++
		
		
	}
	
	if blackCounter >= MartingalaLimit{
		blackStrikes=append(blackStrikes,blackCounter)
		blackStrikesIndex = append(blackStrikesIndex,turns-blackCounter-1)
}
if redCounter >= MartingalaLimit{
	redStrikes=append(redStrikes,redCounter)
	redStrikesIndex = append(redStrikesIndex,turns-redCounter-1)
}

var mediaStrikesPerRound []int
var arrayindex int = 0
for j:=RoundSize; j<turns; j+=RoundSize{
	mediaStrikesPerRound = append(mediaStrikesPerRound,0)
	for k:=0; k<len(redStrikesIndex);k++{
		if redStrikesIndex[k]<=j && redStrikesIndex[k]>j-RoundSize{
			mediaStrikesPerRound[arrayindex]++
		}
	}
	for k:=0; k<len(blackStrikesIndex);k++{
		if blackStrikesIndex[k]<=j && blackStrikesIndex[k]>j-RoundSize{
			mediaStrikesPerRound[arrayindex]++
		}
	}
	arrayindex++
}

	fmt.Println("turns:",turns,"rounds",toFixed(float64(turns)/float64(RoundSize),0))
	fmt.Println("redStrikes:",len(redStrikes),"->",redStrikes, "with indexes,", redStrikesIndex)
	fmt.Println("blackStrikes:",len(blackStrikes),"->",blackStrikes, "with indexes,", blackStrikesIndex)
	fmt.Println("mediaStrikesPerRound:",media(mediaStrikesPerRound),"->",mediaStrikesPerRound)
	fmt.Println("total reds:",totalReds," total blacks",totalBlacks)
fmt.Println("CritIndexesPerRound:",media(CritIndexesPerRound),"->",CritIndexesPerRound)
var failChancePerRound []float64
for _,val := range CritIndexesPerRound{
	chance := float64(val)/float64(RoundSize)/2*2 //index/number of turns/redorblack*numberoftimesibet
	failChancePerRound = append(failChancePerRound,chance)
}
fmt.Println("Fail chance per round:",media64(failChancePerRound),"Max:",Max64(failChancePerRound),"min:",Min64(failChancePerRound),"->",failChancePerRound)

//Concatenated red strikes. This ,means: red strike, the 1 black, then red strike
for index,iVal := range redStrikesIndex{

	if index<=len(redStrikesIndex)-2 && iVal+(redStrikes[index]-1)+2 == redStrikesIndex[index+1] {
		fmt.Println(iVal+(redStrikes[index]-1)+2)

		var cs ConcatenatedStrike
		cs.File=files[ifile]
		cs.Color="red"
		cs.start=iVal
		cs.cut=iVal+(redStrikes[index]-1)+1
		cs.end=cs.cut+redStrikes[index+1]
		concatenatedStrikes = append(concatenatedStrikes,cs)
	}
}
//Concatenated black strikes. This ,means: black strike, the 1 red, then black strike
for index,iVal := range blackStrikesIndex{
	
	if index<len(blackStrikesIndex)-2 && iVal+(blackStrikes[index]-1)+2 == blackStrikesIndex[index+1] {
		var cs ConcatenatedStrike
		cs.File=files[ifile]
		cs.Color="black"
		cs.start=iVal
		cs.cut=iVal+(blackStrikes[index]-1)+1
		cs.end=cs.cut+blackStrikes[index+1]
		concatenatedStrikes = append(concatenatedStrikes,cs)
	}
}

//total strikes records
totalRedStrikesRecord += len(redStrikes)
totalBlackStrikesRecord += len(blackStrikes)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	//every time a file ends, complete the empty turns
	for{
		if turns%RoundSize!=0{
			turns++
		}else{
			break
		}

	}


	 totalReds = 0
	 totalBlacks = 0

	 redCounter = 0
	 blackCounter = 0
	 lastColor = ""

	 redStrikes = redStrikes[:0]
	 redStrikesIndex = redStrikesIndex[:0]
	 blackStrikes = blackStrikes[:0]
	 blackStrikesIndex = blackStrikesIndex[:0]

	 CritIndexesPerRound = CritIndexesPerRound[:0]

	 turns=0

	 fmt.Println("Wins",wins,"Loses",loses)
	wins=0
	loses=0
}
fmt.Println("-------- Total Records -------------")
fmt.Println("totalStrikesRecord",totalRedStrikesRecord,"totalBlackStrikesRecord",totalBlackStrikesRecord)
fmt.Println("concatenatedStrikes",len(concatenatedStrikes),"-->",concatenatedStrikes)
fmt.Println("total wins",totalwins,"total loses",totalloses, "money:",(totalwins*5)-(totalloses*640)-640)


}


func media(array []int) float64{
	var sum float64
	
  for index := range array{
	
	sum += float64(array[index])
  }
  
  
  return sum/float64(len(array))
}

func media64(array []float64) float64{
	var sum float64
	
  for index := range array{
	
	sum += array[index]
  }
  
  
  return sum/float64(len(array))
}

func Min64(v []float64) float64{
	var m float64 = 1000
	for _, e := range v {
		if e < m {
			m = e
		}
	}
	return m
}
func Max64(v []float64) float64{
	var m float64
	for _, e := range v {
		if e > m {
			m = e
		}
	}
	return m
}

func GenerateSerie(){
	var array []int

	for i:=0; i<RoundSize*10;i+=RoundSize{
		rand.Seed(time.Now().UTC().UnixNano()*int64(i+1))
		
		num := randInt(i, i+90)
		num2 := randInt(i, i+90)
		if math.Abs(float64(num)-float64(num2)) <=7{
			i-=RoundSize
		}else{
			array = append(array,num)
			array = append(array,num2)
		}
		
	}
	sort.Ints(array)
	fmt.Println(array)
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}
func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}

//ConcatenatedStrike
type ConcatenatedStrike struct{
	File string
	Color string
	start int
	cut int
	end int
}