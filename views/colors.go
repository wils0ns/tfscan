package views

import "github.com/fatih/color"

// ColorBlue formats text with this color
var ColorBlue = color.New(color.FgBlue).SprintFunc()

// ColorCyan formats text with this color
var ColorCyan = color.New(color.FgCyan).Set().SprintFunc()

// ColorGreen formats text with this color
var ColorGreen = color.New(color.FgGreen).Set().SprintFunc()

// ColorMagenta formats text with this color
var ColorMagenta = color.New(color.FgMagenta).Set().SprintFunc()

// ColorYellow formats text with this color
var ColorYellow = color.New(color.FgYellow).Set().SprintFunc()

// ColorRed formats text with this color
var ColorRed = color.New(color.FgRed).Set().SprintFunc()

// ColorWhite formats text with this color
// var ColorWhite = color.New(color.FgWhite).Set().SprintFunc()

// ColorModuleAddress defines the default color when printing this type of object
var ColorModuleAddress = ColorMagenta

// ColorResourceAddress defines the default color when printing this type of object
var ColorResourceAddress = ColorBlue

// ColorResourceIndex defines the default color when printing this type of object
var ColorResourceIndex = ColorYellow
