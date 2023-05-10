# Crayon ✏️

Crayon is small module for colorizing strings in golang for terminal output.

## Installation

```bash
go get github.com/k3y0708/crayon
```

# Usage

The usage is very simple. You can use `crayon.COLOR(input)` (e.g. `crayon.Red("Hello World!")`) to colorize a string. You can also use `crayon.PrintCOLOR(input)` (e.g. `crayon.PrintRed("Hello World!")`) to print a colorized string. Anothere option is to change the background color with `crayon.BgCOLOR(input)` (e.g. `crayon.BgRed("Hello World!")`) or `crayon.PrintBgCOLOR(input)` (e.g. `crayon.PrintBgRed("Hello World!")`).

Instead of a color you can also use a style (e.g. `crayon.Bold("Hello World!")` or `crayon.PrintBold("Hello World!")`).

You can also combine styles and colors (e.g. `crayon.Red(crayon.Bold("Hello World!"))` or `crayon.PrintRed(crayon.Bold("Hello World!"))`).

## Colorized strings

### Colors

```go
fmt.Println(crayon.Red("This is red"))
fmt.Println(crayon.Green("This is green"))
fmt.Println(crayon.Blue("This is blue"))
fmt.Println(crayon.Yellow("This is yellow"))
fmt.Println(crayon.Magenta("This is magenta"))
fmt.Println(crayon.Cyan("This is cyan"))
fmt.Println(crayon.White("This is white"))
fmt.Println(crayon.Black("This is black"))
fmt.Println(crayon.Grey("This is grey"))
```

### Bright Colors

```go
fmt.Println(crayon.BrightRed("This is bright red"))
fmt.Println(crayon.BrightGreen("This is bright green"))
fmt.Println(crayon.BrightBlue("This is bright blue"))
fmt.Println(crayon.BrightYellow("This is bright yellow"))
fmt.Println(crayon.BrightMagenta("This is bright magenta"))
fmt.Println(crayon.BrightCyan("This is bright cyan"))
fmt.Println(crayon.BrightWhite("This is bright white"))
fmt.Println(crayon.BrightBlack("This is bright black"))
fmt.Println(crayon.BrightGrey("This is bright grey"))
```

### Background Colors

```go
fmt.Println(crayon.BgRed("This is red background"))
fmt.Println(crayon.BgGreen("This is green background"))
fmt.Println(crayon.BgBlue("This is blue background"))
fmt.Println(crayon.BgYellow("This is yellow background"))
fmt.Println(crayon.BgMagenta("This is magenta background"))
fmt.Println(crayon.BgCyan("This is cyan background"))
fmt.Println(crayon.BgWhite("This is white background"))
fmt.Println(crayon.BgBlack("This is black background"))
fmt.Println(crayon.BgGrey("This is grey background"))
```

### Bright Background Colors

```go
fmt.Println(crayon.BgBrightRed("This is bright red background"))
fmt.Println(crayon.BgBrightGreen("This is bright green background"))
fmt.Println(crayon.BgBrightBlue("This is bright blue background"))
fmt.Println(crayon.BgBrightYellow("This is bright yellow background"))
fmt.Println(crayon.BgBrightMagenta("This is bright magenta background"))
fmt.Println(crayon.BgBrightCyan("This is bright cyan background"))
fmt.Println(crayon.BgBrightWhite("This is bright white background"))
fmt.Println(crayon.BgBrightBlack("This is bright black background"))
fmt.Println(crayon.BgBrightGrey("This is bright grey background"))
```

### Styles

```go
fmt.Println(crayon.Bold("This is bold"))
fmt.Println(crayon.Dim("This is dim"))
fmt.Println(crayon.Italic("This is italic"))
fmt.Println(crayon.Underline("This is underline"))
fmt.Println(crayon.Strikethrough("This is strikethrough"))
```

## Colorized Println

### Colors

```go
crayon.PrintRed("This is red")
crayon.PrintGreen("This is green")
crayon.PrintBlue("This is blue")
crayon.PrintYellow("This is yellow")
crayon.PrintMagenta("This is magenta")
crayon.PrintCyan("This is cyan")
crayon.PrintWhite("This is white")
crayon.PrintBlack("This is black")
crayon.PrintGrey("This is grey")
```

### Bright Colors

```go
crayon.PrintBrightRed("This is bright red")
crayon.PrintBrightGreen("This is bright green")
crayon.PrintBrightBlue("This is bright blue")
crayon.PrintBrightYellow("This is bright yellow")
crayon.PrintBrightMagenta("This is bright magenta")
crayon.PrintBrightCyan("This is bright cyan")
crayon.PrintBrightWhite("This is bright white")
crayon.PrintBrightBlack("This is bright black")
crayon.PrintBrightGrey("This is bright grey")
```

### Background Colors

```go
crayon.PrintBgRed("This is red background")
crayon.PrintBgGreen("This is green background")
crayon.PrintBgBlue("This is blue background")
crayon.PrintBgYellow("This is yellow background")
crayon.PrintBgMagenta("This is magenta background")
crayon.PrintBgCyan("This is cyan background")
crayon.PrintBgWhite("This is white background")
crayon.PrintBgBlack("This is black background")
crayon.PrintBgGrey("This is grey background")
```

### Bright Background Colors

```go
crayon.PrintBgBrightRed("This is bright red background")
crayon.PrintBgBrightGreen("This is bright green background")
crayon.PrintBgBrightBlue("This is bright blue background")
crayon.PrintBgBrightYellow("This is bright yellow background")
crayon.PrintBgBrightMagenta("This is bright magenta background")
crayon.PrintBgBrightCyan("This is bright cyan background")
crayon.PrintBgBrightWhite("This is bright white background")
crayon.PrintBgBrightBlack("This is bright black background")
crayon.PrintBgBrightGrey("This is bright grey background")
```

### Styles

```go
crayon.PrintBold("This is bold")
crayon.PrintDim("This is dim")
crayon.PrintItalic("This is italic")
crayon.PrintUnderline("This is underline")
crayon.PrintStrikethrough("This is strikethrough")
```
