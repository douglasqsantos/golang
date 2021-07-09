package msgs

import (
	"fmt"
	"time"
)

// Define the type Color
type Color string

// Defining the colors
const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed    Color = "\u001b[31m"
	ColorGreen  Color = "\u001b[32m"
	ColorYellow Color = "\u001b[33m"
	ColorBlue   Color = "\u001b[34m"
	ColorPurple Color = "\u001b[35m"
	ColorCyan   Color = "\u001b[36m"
	ColorGrey   Color = "\u001b[37m"
	ColorWhite  Color = "\u001b[39m"
	ColorReset  Color = "\u001b[0m"
)

// Main msg function to handle the header of each message with date and time
func msg(module string) string {
	now := time.Now()
	date := fmt.Sprintf("%02d/%02d/%d", now.Day(), now.Month(), now.Year())
	time := fmt.Sprintf("%02d:%02d:%02d", now.Hour(), now.Minute(), now.Second())
	return fmt.Sprintf("%v [%v %v - %v %v][%v %v %v] %v> %v", string(ColorRed), string(ColorWhite), date, time, string(ColorRed), string(ColorBlue), module, string(ColorRed), string(ColorWhite), string(ColorReset))
}

// Warning messages
func Warn(module string, message string) {
	msg_module := msg(module)
	fmt.Printf("%v %v[!]%v %v", msg_module, string(ColorYellow), string(ColorReset), message)
}

// Error messages
func Error(module string, message string) {
	msg_module := msg(module)
	fmt.Printf("%v %v[-]%v %v", msg_module, string(ColorRed), string(ColorReset), message)
}

// Successfull messages
func Ok(module string, message string) {
	msg_module := msg(module)
	fmt.Printf("%v %v[+]%v %v", msg_module, string(ColorGreen), string(ColorReset), message)
}

// Highlighted
func Focus(message string) string {
	return fmt.Sprintf("%v[%v %v %v]%v", string(ColorRed), string(ColorWhite), message, string(ColorRed), string(ColorReset))
}
