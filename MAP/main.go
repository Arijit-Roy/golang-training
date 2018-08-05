package main

import(
    "fmt"
)

type colorMap map[string]string
func main(){
	color := colorMap{
        "red": "#235677",
        "greeen": "#344354",
	}


	color["white"] = "#ffffff"
	color.print()

	cl := "red"
	delete(color, cl)
	
	color.print()

}

func (c colorMap) print(){
	for color, hc := range c {
		fmt.Println("color ", color, " has code ", hc )
	}
}