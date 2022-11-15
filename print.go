package main

import "fmt"

// Helper ---------------------------------------------------------------------

func printColorize(color string, text string) {
	fmt.Println(color + text + RESET)
}

// Colors ---------------------------------------------------------------------

func PrintRed(text string) {
	printColorize(RED, text)
}

func PrintGreen(text string) {
	printColorize(GREEN, text)
}

func PrintYellow(text string) {
	printColorize(YELLOW, text)
}

func PrintBlue(text string) {
	printColorize(BLUE, text)
}

func PrintMagenta(text string) {
	printColorize(MAGENTA, text)
}

func PrintCyan(text string) {
	printColorize(CYAN, text)
}

func PrintWhite(text string) {
	printColorize(WHITE, text)
}

func PrintGrey(text string) {
	printColorize(GREY, text)
}

func PrintBlack(text string) {
	printColorize(BLACK, text)
}

// Bright Colors --------------------------------------------------------------

func PrintRedBright(text string) {
	printColorize(RED_BRIGHT, text)
}

func PrintGreenBright(text string) {
	printColorize(GREEN_BRIGHT, text)
}

func PrintYellowBright(text string) {
	printColorize(YELLOW_BRIGHT, text)
}

func PrintBlueBright(text string) {
	printColorize(BLUE_BRIGHT, text)
}

func PrintMagentaBright(text string) {
	printColorize(MAGENTA_BRIGHT, text)
}

func PrintCyanBright(text string) {
	printColorize(CYAN_BRIGHT, text)
}

func PrintWhiteBright(text string) {
	printColorize(WHITE_BRIGHT, text)
}

func PrintGreyBright(text string) {
	printColorize(GREY_BRIGHT, text)
}

func PrintBlackBright(text string) {
	printColorize(BLACK_BRIGHT, text)
}

// Background Colors -----------------------------------------------------------

func PrintBgRed(text string) {
	printColorize(BG_RED, text)
}

func PrintBgGreen(text string) {
	printColorize(BG_GREEN, text)
}

func PrintBgYellow(text string) {
	printColorize(BG_YELLOW, text)
}

func PrintBgBlue(text string) {
	printColorize(BG_BLUE, text)
}

func PrintBgMagenta(text string) {
	printColorize(BG_MAGENTA, text)
}

func PrintBgCyan(text string) {
	printColorize(BG_CYAN, text)
}

func PrintBgWhite(text string) {
	printColorize(BG_WHITE, text)
}

func PrintBgGrey(text string) {
	printColorize(BG_GREY, text)
}

func PrintBgBlack(text string) {
	printColorize(BG_BLACK, text)
}

// Bright Background Colors ----------------------------------------------------

func PrintBgRedBright(text string) {
	printColorize(BG_RED_BRIGHT, text)
}

func PrintBgGreenBright(text string) {
	printColorize(BG_GREEN_BRIGHT, text)
}

func PrintBgYellowBright(text string) {
	printColorize(BG_YELLOW_BRIGHT, text)
}

func PrintBgBlueBright(text string) {
	printColorize(BG_BLUE_BRIGHT, text)
}

func PrintBgMagentaBright(text string) {
	printColorize(BG_MAGENTA_BRIGHT, text)
}

func PrintBgCyanBright(text string) {
	printColorize(BG_CYAN_BRIGHT, text)
}

func PrintBgWhiteBright(text string) {
	printColorize(BG_WHITE_BRIGHT, text)
}

func PrintBgGreyBright(text string) {
	printColorize(BG_GREY_BRIGHT, text)
}

func PrintBgBlackBright(text string) {
	printColorize(BG_BLACK_BRIGHT, text)
}

// Text Styles -----------------------------------------------------------------

func PrintBold(text string) {
	printColorize(BOLD, text)
}

func PrintDim(text string) {
	printColorize(DIM, text)
}

func PrintItalic(text string) {
	printColorize(ITALIC, text)
}

func PrintUnderline(text string) {
	printColorize(UNDERLINE, text)
}

func PrintStrikethrough(text string) {
	printColorize(STRIKETHROUGH, text)
}
