//--------------------------------
// Battleship - K.Andreev 20182907
// BSD 2.0 License
// https://blog.andreev.it
//--------------------------------
package main
 
import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strings"
    "time"
)
 
//Size of the board 10 x 10
var intBoardSize = 10
 
//Total number of ships on each side
var intTotalShips = 15
 
//Player's name
var sYourName = ""
 
//Holder for "press any key"
var sAnyKey = ""
 
//Number of player's turns
var intPlayerTurns = 0
 
//Number of computer's turns
var intComputerTurns = 0
 
//Who plays first
var bPlayerFirst = true
 
//How many ships left for the player
var intPlayerShipsLeft = 15
 
//How many ships left for the computer
var intComputerShipsLeft = 15
 
//Define a two dimensional array for the player
var arrPlayerBoard [10][10]int
 
//Define a two dimension array for the computer
var arrComputerBoard [10][10]int
 
//Define Unicode elements for on-screen presentation
//If there are any issues with the display, change them here
//const emptyElement = "\u25a2"     //replace with " "
//const shipElement = "\u25a3"      //replace with "X"
//const shipWreckElement = "\u25a4" //replace with "+"
//const missedElement = "\u25a9"    //replace with "o"
const emptyElement = " "    
const shipElement = "X"     
const shipWreckElement = "+"
const missedElement = "o"   
 
func main() {
    //Initialize the random generator
    rand.Seed(time.Now().UnixNano())
    //Print the initial greetings and instructions
    printInstructions()
    //Initialize the boards
    initBoard()
    //Get player's name
    inputYourName()
    //Print the boards
    printBoard()
    //Get the player's ships (coordinates)
    inputFleet()
    //Generate computer's fleet
    generateComputerFleet()
    //Print the boards
    printBoard()
    //Who plays first
    decideFirstPlayer()
    //MAIN GAME LOOP
    for {
        if bPlayerFirst {
            playerMove()
            computerMove()
        } else {
            computerMove()
            playerMove()
        }
    }
}
 
//Print's game instructions
func printInstructions() {
    fmt.Println("               ==========")
    fmt.Println("               Battleship")
    fmt.Println("               ==========")
    fmt.Println("")
    fmt.Println("The goal of this game is to sink all of the enemy's ships on the", intBoardSize, "X", intBoardSize, "board.")
    fmt.Println("You do that by guessing ship's location using coordinates (e.g. LAUNCH AT> C3).")
    fmt.Println("If you or the computer hit a ship, another chance is given until you miss again.")
    fmt.Println("The game is over when you or the opponent sinks all", intTotalShips, "ships.")
    fmt.Println("")
    fmt.Println("To make the game more difficult, you won't see where you missed on the computer board")
    fmt.Println("...and the computer won't hit the same place twice. HAVE FUN!!!")
    fmt.Println("")
    fmt.Println("K.Andreev - 20180729 - BSD Simplified License")
}
 
//Initialized both boards with zero values
//Sets the number of ships
func initBoard() {
    i := 0
    j := 0
    intPlayerShipsLeft = 15
    intComputerShipsLeft = 15
    intPlayerTurns = 0
    intComputerTurns = 0
    for i = 0; i < intBoardSize; i++ {
        for j = 0; j < intBoardSize; j++ {
            arrPlayerBoard[i][j] = 0
            arrComputerBoard[i][j] = 0
        }
    }
}
 
//Gets player's name
func inputYourName() {
    fmt.Println("")
    fmt.Print("Enter your name: ")
    fmt.Scanln(&sYourName)
    fmt.Println("")
    fmt.Println("Hello,", sYourName, "!")
    fmt.Println("")
    fmt.Println("Please place your fleet using coordinates (e.g. D5)")
    fmt.Println("If you make a mistake, re-enter the same coordinate.")
    fmt.Println("")
    pressAnyKey()
}
 
//Ask the player to enter the coordinates of the fleet
func inputFleet() {
    intShipCounter := 0
    intXCoord := 0
    intYCoord := 0
    sPosition := ""
    fmt.Println("")
    for {
    enter:
        fmt.Print("SHIP AT> ")
        fmt.Scan(&sPosition)
        //If the length is different than two, coordinate is wrong
        if len(sPosition) != 2 {
            fmt.Println("INVALID POSITION.")
            goto enter
        }
        //Convert the coordiantes in array position, e.g. A0 is (0,0), J9 is (9,9)
        sPosition = strings.ToUpper(sPosition)
        intXCoord = int(sPosition[0] - 65)
        intYCoord = int(sPosition[1] - 48)
        //If the array coordinate is not in [(0,0)..(9,9)] range, coordinate is wrong
        if intXCoord < 0 || intXCoord > 9 {
            fmt.Println("INVALID POSITION.")
            goto enter
        }
        //If the player already entered a ship at a coordinate, toggle it, make it empty
        if arrPlayerBoard[intXCoord][intYCoord] == 1 {
            arrPlayerBoard[intXCoord][intYCoord] = 0
        } else {
            arrPlayerBoard[intXCoord][intYCoord] = 1
            intShipCounter = intShipCounter + 1
        }
        printBoard()
        //Exit when all 15 ships are placed
        if intShipCounter == 15 {
            break
        }
    }
}
 
//Prints the boards on the screen
func printBoard() {
    fmt.Println(sYourName)
    i := 0
    j := 0
    intPlayerShipCounter := 0
    intComputerShipCounter := 0
    sChar := 'A'
    fmt.Println(" 0123456789")
    //Prints player's board
    for i = 0; i < intBoardSize; i++ {
        fmt.Print(string(sChar))
        sChar = sChar + 1
        for j = 0; j < intBoardSize; j++ {
            switch arrPlayerBoard[i][j] {
            case 0:
                fmt.Print(emptyElement)
            case 1:
                fmt.Print(shipElement)
                intPlayerShipCounter = intPlayerShipCounter + 1
            case 2:
                fmt.Print(shipWreckElement)
            case 3:
                fmt.Print(missedElement)
            }
        }
        fmt.Println("")
    }
    //Prints player's ships left and the total
    fmt.Println("SHIPS", intPlayerShipCounter, "/", intTotalShips)
    fmt.Println("")
    //Prints computer's board
    fmt.Println("COMPUTER")
    i = 0
    j = 0
    sChar = 'A'
    fmt.Println(" 0123456789")
    for i = 0; i < intBoardSize; i++ {
        fmt.Print(string(sChar))
        sChar = sChar + 1
        for j = 0; j < intBoardSize; j++ {
            switch arrComputerBoard[i][j] {
            case 0:
                fmt.Print(emptyElement)
            case 1:
                //Uncomment the line below to see computer's ships (cheat)
                //fmt.Print(shipElement)
                intComputerShipCounter = intComputerShipCounter + 1
            case 2:
                fmt.Print(shipWreckElement)
            case 3:
                fmt.Print(missedElement)
            }
        }
        fmt.Println("")
    }
    //Prints computer's ships left and the total
    fmt.Println("SHIPS", intComputerShipCounter, "/", intTotalShips)
}
 
//Places computer's fleet using random generator
func generateComputerFleet() {
    intShipCounter := 0
    intRandomShipatX := 0
    intRandomShipatY := 0
    for {
        intRandomShipatX = rand.Intn(10)
        intRandomShipatY = rand.Intn(10)
        if arrComputerBoard[intRandomShipatX][intRandomShipatY] == 0 {
            arrComputerBoard[intRandomShipatX][intRandomShipatY] = 1
            intShipCounter = intShipCounter + 1
        }
        if intShipCounter == 15 {
            break
        }
    }
}
 
//Who plays first. A number between 0 and 99 is drawn for both players.
//The bigger number plays first.
func decideFirstPlayer() {
    intPlayerFirst := 0
    intComputerFirst := 0
    fmt.Println("")
    fmt.Println("Prepare for the battle...")
    intPlayerFirst = rand.Intn(100)
    intComputerFirst = rand.Intn(100)
    if intPlayerFirst >= intComputerFirst {
        bPlayerFirst = true
        fmt.Println("You play first. Random numbers say", intPlayerFirst, "vs.", intComputerFirst)
        fmt.Scanln(&sAnyKey)
        pressAnyKey()
    } else {
        bPlayerFirst = false
        fmt.Println("Computer plays first. Random numbers say", intComputerFirst, "vs.", intPlayerFirst)
        fmt.Scanln(&sAnyKey)
        pressAnyKey()
    }
}
 
//Get the player to input a coordinate
func playerMove() {
    intXCoord := 0
    intYCoord := 0
    sPosition := ""
    fmt.Println("")
enter:
    for {
        fmt.Print("LAUNCH AT> ")
        fmt.Scan(&sPosition)
        //Same checks for invalid position
        if len(sPosition) != 2 {
            fmt.Println("INVALID POSITION.")
            goto enter
        }
        sPosition = strings.ToUpper(sPosition)
        intXCoord = int(sPosition[0] - 65)
        intYCoord = int(sPosition[1] - 48)
        if intXCoord < 0 || intXCoord > 9 {
            fmt.Println("INVALID POSITION.")
            goto enter
        }
        //Increase the turns number
        intPlayerTurns = intPlayerTurns + 1
        //This is a miss if the board has 0 value at the coordinate
        if arrComputerBoard[intXCoord][intYCoord] == 0 {
            fmt.Print("You missed at ", sPosition, ".")
            fmt.Scanln(&sAnyKey)
            pressAnyKey()
        } else {
            //This is a hit
            arrComputerBoard[intXCoord][intYCoord] = 2
            fmt.Print("YOU HIT A SHIP AT ", sPosition, ".")
            fmt.Scanln(&sAnyKey)
            pressAnyKey()
            printBoard()
            intComputerShipsLeft = intComputerShipsLeft - 1
            //Game over if all ships are hit
            if intComputerShipsLeft == 0 {
                fmt.Println("================")
                fmt.Println("[   YOU WON    ]")
                fmt.Println("================")
                fmt.Println(intPlayerTurns, "turns to complete the game...")
                os.Exit(0)
            }
            goto enter
        }
        printBoard()
        break
    }
}
 
//Computer logic for the play
func computerMove() {
    intRandomShipatX := 0
    intRandomShipatY := 0
enter:
    for {
        //Get a random coordinate
        intRandomShipatX = rand.Intn(10)
        intRandomShipatY = rand.Intn(10)
        //Increase the number of turns
        intComputerTurns = intComputerTurns + 1
        //If it's a miss, remember that position by placing "3" and never hit that position again
        if arrPlayerBoard[intRandomShipatX][intRandomShipatY] == 0 {
            arrPlayerBoard[intRandomShipatX][intRandomShipatY] = 3
            fmt.Println("Computer missed at", string(intRandomShipatX+65), string(intRandomShipatY+48))
            pressAnyKey()
            printBoard()
            break
        }
        //If it's a hit, mark the ship as hit by placing "2"
        if arrPlayerBoard[intRandomShipatX][intRandomShipatY] == 1 {
            arrPlayerBoard[intRandomShipatX][intRandomShipatY] = 2
            fmt.Println("COMPUTER HIT A SHIP at", string(intRandomShipatX+65), string(intRandomShipatY+48))
            pressAnyKey()
            printBoard()
            intPlayerShipsLeft = intPlayerShipsLeft - 1
            if intPlayerShipsLeft == 0 {
                fmt.Println("================")
                fmt.Println("[ COMPUTER WON ]")
                fmt.Println("================")
                fmt.Println(intComputerTurns, "turns to complete the game...")
                os.Exit(0)
            } else {
                goto enter
            }
        }
        //If the random coordinate is where the ship was already hit, go guess again
        if arrPlayerBoard[intRandomShipatX][intRandomShipatY] == 2 {
            goto enter
        }
        //If the random coordinate is where we missed, go guess again
        if arrPlayerBoard[intRandomShipatX][intRandomShipatY] == 3 {
            goto enter
        }
    }
}
 
//Wait for any key to be pressed
func pressAnyKey() {
    in := bufio.NewReader(os.Stdin)
    line, _ := in.ReadString('\n')
    _ = line
}
