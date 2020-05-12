package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jinzhu/now"
	"os"
)

func main() {
	//colors()
	times()
}

func times()  {
	fmt.Println("All the beginnings...")
	fmt.Println("All the beginnings...")
	fmt.Println(now.BeginningOfMinute())
	fmt.Println(now.BeginningOfHour())
	fmt.Println(now.BeginningOfDay())
	fmt.Println(now.BeginningOfWeek())
	fmt.Println(now.BeginningOfMonth())
	fmt.Println(now.BeginningOfQuarter())
	fmt.Println(now.BeginningOfYear().Day())
}
func colors()  {
	color.Cyan("Prints text in cyan.")

	// a newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// More default foreground colors..
	color.Red("We have red")
	color.Yellow("Yellow color too!")
	color.Magenta("And many others ..")

	// Hi-intensity colors
	color.HiGreen("Bright green color.")
	color.HiBlack("Bright black means gray..")
	color.HiWhite("Shiny white color!")

	// Create a new color object
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")

	// Or just add them to New()
	d := color.New(color.FgCyan, color.Bold)
	d.Printf("This prints bold cyan %s\n", "too!.")

	// Mix up foreground and background colors, create new mixes!
	red := color.New(color.FgRed)

	boldRed := red.Add(color.Bold)
	boldRed.Println("This will print text in bold red.")

	whiteBackground := red.Add(color.BgWhite)
	whiteBackground.Println("Red text with White background.")

	// Use your own io.Writer output
	color.New(color.FgBlue).Fprintln(os.Stdout, "blue color!")

	blue := color.New(color.FgBlue)
	blue.Fprint(os.Stdout, "This will print text in blue.\n")
}
