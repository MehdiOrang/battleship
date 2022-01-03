package main

import (
	"fmt"
	"bytes"
	"regexp"
	"strconv"
	"strings"
	"math/rand"
	"time"
)

func main(){

	userBoard     := Board{}
	computerBoard := Board{}
	userBoard.newBoard()
	computerBoard.newBoard()
	battleShip := Battleship{userBoard: userBoard, computerBoard: computerBoard }
	battleShip.startGame()
		
}


type Board struct{
    board [][]string
    ships []Ship
}

type Ship struct{
    shipType    string
    location    []int
    orientation string
    dead        bool
}

type Battleship struct {
   userBoard     Board
   computerBoard Board
}

func convertCoordinates(A string) []int {
	letters := []byte{'A','B','C','D','E','F','G','H','I','J'}
	r       := regexp.MustCompile("[0-9]+")
 	column, _  := strconv.Atoi(r.FindString(A))
 	return []int{column-1, int(bytes.IndexAny([]byte(letters),string(A[0])))}
}

func reverse(A []int) string{
	letters := []byte{'A','B','C','D','E','F','G','H','I','J'}
	column := string(letters[A[1]])
	row    := strconv.Itoa(A[0] + 1)
	return column+row 
}


func(ship *Ship) die(){
    ship.dead = true
}

func(ship *Ship) shipSize() int{
    switch {
        case ship.shipType == "carrier":
            return 5
        case ship.shipType ==  "battleship":
            return 4 
        case ship.shipType ==  "cruiser":
            return 3 
        case strings.Contains(ship.shipType, "destroyer"): 
            return 2
        case strings.Contains(ship.shipType, "submarine"):
            return 1     
    }
    
    return 1
}

func(brd *Board) newBoard() {
    board := make([][]string,10)
    for i, _ := range board {
	    board[i] = []string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-"}
    }
   
    brd.board = board
    brd.ships = []Ship{}

}

func(brd *Board) shot(row int, column int) bool {
    mycase := brd.board[row][column]
    switch mycase{
        case "-":
            brd.board[row][column] = "/"
        case "0":
            brd.board[row][column] = "X" 
            fmt.Println("You hit a ship!")
        case "/", "X":
            fmt.Println("Repeated coordinate. Choose again")
            return false      
    }
    brd.checkIfShipSank()
    return true
}


func(brd *Board) insertShip(ship Ship) bool {

	   if brd.placeable(ship){
		  brd.ships = append(brd.ships, ship)
		  if ship.orientation == "horizontal" {
		  	 for i:=0; i < ship.shipSize(); i++{
	      			brd.board[ship.location[0]][ship.location[1]+i] = "0"
	      		 }
		  
		  }else if ship.orientation == "vertical" {
		  	 for i:=0; i < ship.shipSize(); i++{
	      			brd.board[ship.location[0]+i][ship.location[1]] = "0"
	      		 }
		  }     
	     return true
	   }

    return false

}
//?
func(brd *Board) placeable(ship Ship) bool{

	if ship.location[0] < 0 || ship.location[1] < 0 || ship.location[0] > 9 || ship.location[1] > 9{
		return false
	}
	
        fields := brd.shipFields(ship)
        if fields == nil {
        	return false
        }
        
        for _, x := range fields {
		if x == "0" {
		   return false
		}
    	}
 
	return true

}

func(brd *Board) shipFields(ship Ship) []string{
    field := make([]string,ship.shipSize())
    if ship.orientation == "horizontal"{
	      if ship.location[1]+ship.shipSize() > 10 {
	      	  return nil
	      }
	      shipRange := make([]int, (ship.location[1]+ship.shipSize()) - ship.location[1] )
	      incr := ship.location[1]
	      for i := 0; i < len(shipRange); i++{
	          field[i] = brd.board[ship.location[0]][incr+i]
	      } 
    }else if ship.orientation == "vertical" {
	     if ship.location[0]+ship.shipSize() > 10 {
	         return nil
	     }
	     shipRange := make([]int, (ship.location[0]+ship.shipSize()) - ship.location[0] )
	     incr := ship.location[0]
	     for i := 0; i < len(shipRange); i++{
	          field[i] = brd.board[incr+i][ship.location[1]]
	     } 
    }
    
    return field

}



 func(brd *Board) liveShips() int{
 	len    := len(brd.ships)
 	result := len
 	for i:=0; i< len; i++{
 		if brd.ships[i].dead { result--}
 	}
 
 	return result
 }
 
 func(brd *Board) print(A bool) {
  	fmt.Println("   A B C D E F G H I J ")
  	fmt.Println("  _____________________")
  	counter := 1
  	
  	for _, line := range brd.board {
  	         print := ""
  		 if counter < 10 {
  		 	print += strconv.Itoa(counter) + " "
  		 }else{
  		 	print += strconv.Itoa(counter)
  		 }
		
		for _, elm := range line {
			print += "|"
			if elm == "-"{ 
				print += "-" 
			}else if A == false { 
				print += elm
			}else{
				if elm == "0" {
				   elm = "-"
				}
				print += elm
			} 
    		}
    		print += "|"
    		fmt.Println(print)
    		counter += 1
    	}
    	
  	fmt.Println("  _____________________")

 }
 // ?
func(brd *Board) randomPlacement() {

    shipTypes := []string{"carrier", "battleship", "cruiser", "destroyer 1", "destroyer 2", "submarine 1", "submarine 2"}
    for _, ship := range shipTypes {
      placeable := true
      for ;placeable;{
         placeable = false
      	 location := []int{rand.Intn(9 - 0)+0,rand.Intn(9 - 0)+0}
      	 orientation := ""
      	 if rand.Intn(1 - 0)+0 == 1{
      	 	orientation += "horizontal"
      	 }else{
      	 	orientation += "horizontal"
      	 }
      	 myRandomship := Ship{shipType:ship, location: []int{location[0],location[1]}, orientation: orientation}
      	 placeable = brd.insertShip(myRandomship)
      }
    }
}
//?
func(brd *Board) checkIfShipSank() {
	fmt.Println("Dead ships:")
	for _, ship := range brd.ships {
	  field    := brd.shipFields(ship)
	  shipSize := ship.shipSize()
	  for _, elm := range field {
	    if elm == "x"{
	       shipSize-- 
	    }
	    if shipSize == 0{
	        fmt.Printf("%v has sunk \n", ship.shipType )
	        ship.die()
	    }
	  }
	}
}

func(battleship *Battleship) startGame() {
    fmt.Println("Hello and welcome to Battleship! We will start by placing ships. Type random to randomize placement.")
    battleship.userBoard.print(false)
    battleship.placeUserShips()
    battleship.computerBoard.randomPlacement()
    battleship.startShooting()
}

func(battleship *Battleship) placeUserShips() {
    shipTypes := []string{"carrier", "battleship", "cruiser", "destroyer 1", "destroyer 2", "submarine 1", "submarine 2"}
    for _, ship := range shipTypes {
      placeable := false
      for {
      
      		fmt.Printf("Where would you like to place your %v?\n", ship)
      		location := ""
	        if _, err := fmt.Scan(&location); location == "random" {
	       	if err != nil { fmt.Println(err) }
	       	battleship.userBoard.randomPlacement()
	       	fmt.Printf("This is your board: \n")
      			battleship.userBoard.print(false)
			return
		}

       
	       fmt.Printf("Would you like your %v to be horizontal or vertical?\n", ship)
	       orientation := ""
	       for {
		       if _, err := fmt.Scan(&orientation); orientation == "horizontal" || orientation == "vertical" {
		       	if err != nil { fmt.Println(err) }
				break
			}else{
				fmt.Println("Please input valid value")
			}
	       }
	       
	       placeable = battleship.userBoard.insertShip(Ship{shipType:ship, location:convertCoordinates(location), orientation: orientation})
	       if  placeable == false { fmt.Println("Not a valid position. Please choose again")}
	       if  placeable == true { break }
       }
       fmt.Println("Your ship has been placed.")
       battleship.userBoard.print(false)
    
    }
}


func(battleship *Battleship) userShot() { 
	fmt.Println("Call out your shot or type history")
	shot := ""
	fmt.Scan(&shot);
	matched, _ := regexp.MatchString(`^[A-J]\d0?`, shot )
        switch  {
           case shot == "view": //cheating!
      		battleship.computerBoard.print(false)
      		battleship.userShot()
		return
	    case shot == "history":
	        battleship.computerBoard.print(true)
	        battleship.userShot()
      		return
      	    case matched: 
      	        shoot := convertCoordinates(shot)
      	        battleship.computerBoard.shot(shoot[0], shoot[1])
      	    default:
      	         fmt.Println("Invalid coordinates. Please try again.")
      	         battleship.userShot()
      	     
	}
}

func(battleship *Battleship) computerShot() { 
  result := false
  
  for{
      time.Sleep(1 * time.Second)  
      location := []int{rand.Intn(9 - 0)+0,rand.Intn(9 - 0)+0}
      result = battleship.userBoard.shot(location[0],location[1])
      fmt.Printf("Opponent shot at %v\n", reverse(location))
      if result == true {
    
    	  battleship.userBoard.print(false)
    	  return
      }
  
  }


}

func(battleship *Battleship) startShooting() { 

	for {
		userShips := battleship.userBoard.liveShips()
		fmt.Println("user live ships", userShips)
		cpShips   := battleship.computerBoard.liveShips()
		fmt.Println("cmp live ships", cpShips)
		//fmt.Println(battleship.computerBoard.ships)
		
		if userShips <= 0 {
			fmt.Println("Computer Won.")
			return
		}
		if cpShips <= 0 {
			fmt.Println("You Won.")
			return
		}
		
		for i:=0; i <= userShips; i++{
		      battleship.userShot()
		}
		
		for i:=0; i <= userShips; i++{
		      battleship.computerShot()
		}
	
	}
}






   



