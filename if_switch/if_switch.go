package ifswitch

import "fmt"

func IfSwitchMain() {
    arturoName := "Arturo"
    arturoAge := 30

    santiagoName := "Santiago"
    santiagoAge := 17

    if arturoAge >= 18 {
        fmt.Println(arturoName, "can drink")
    }  

    if santiagoAge >= 18 { // santiagoAge -> 17 => cannot drink
        fmt.Println(santiagoName, "can drink")
    } else {
        fmt.Println(santiagoName, "cannot drink")
    }


    switch {
        case arturoAge >= 50:
            fmt.Println("Can drink but in lower amounts")
        case arturoAge >= 40 && arturoAge < 50:
            fmt.Println("Can drink with no problem")
        case arturoAge >= 30 && arturoAge < 40:
            fmt.Println("Can drink AND have a discount")
        case arturoAge > 18:
            fmt.Println("You cannot drink")
        default:
            fmt.Println("Check with bar tender")
    }

    grade := "B-"

    switch grade {
        case "A":
            fmt.Println("You have 94%")
        case "A-":
            fmt.Println("You have 92%")
        case "B+":
            fmt.Println("You have 89%")
        case "B-":
            fmt.Println("You have 82%")
        case "C":
            fmt.Println("You have 73%")
        default:
            fmt.Println("Go check with the Teacher")
    }

}
