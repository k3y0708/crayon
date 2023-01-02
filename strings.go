package crayon

// Helper ---------------------------------------------------------------------

func colorize(color string, text string) string {
	return color + text + RESET
}

// Colors ---------------------------------------------------------------------

func Red(text string) string {
	return colorize(RED, text)
}

func Green(text string) string {
	return colorize(GREEN, text)
}

func Yellow(text string) string {
	return colorize(YELLOW, text)
}

func Blue(text string) string {
	return colorize(BLUE, text)
}

func Magenta(text string) string {
	return colorize(MAGENTA, text)
}

func Cyan(text string) string {
	return colorize(CYAN, text)
}

func White(text string) string {
	return colorize(WHITE, text)
}

func Grey(text string) string {
	return colorize(GREY, text)
}

func Black(text string) string {
	return colorize(BLACK, text)
}

// Bright Colors --------------------------------------------------------------

func RedBright(text string) string {
	return colorize(RED_BRIGHT, text)
}

func GreenBright(text string) string {
	return colorize(GREEN_BRIGHT, text)
}

func YellowBright(text string) string {
	return colorize(YELLOW_BRIGHT, text)
}

func BlueBright(text string) string {
	return colorize(BLUE_BRIGHT, text)
}

func MagentaBright(text string) string {
	return colorize(MAGENTA_BRIGHT, text)
}

func CyanBright(text string) string {
	return colorize(CYAN_BRIGHT, text)
}

func WhiteBright(text string) string {
	return colorize(WHITE_BRIGHT, text)
}

func GreyBright(text string) string {
	return colorize(GREY_BRIGHT, text)
}

func BlackBright(text string) string {
	return colorize(BLACK_BRIGHT, text)
}

// Background Colors ----------------------------------------------------------

func BgRed(text string) string {
	return colorize(BG_RED, text)
}

func BgGreen(text string) string {
	return colorize(BG_GREEN, text)
}

func BgYellow(text string) string {
	return colorize(BG_YELLOW, text)
}

func BgBlue(text string) string {
	return colorize(BG_BLUE, text)
}

func BgMagenta(text string) string {
	return colorize(BG_MAGENTA, text)
}

func BgCyan(text string) string {
	return colorize(BG_CYAN, text)
}

func BgWhite(text string) string {
	return colorize(BG_WHITE, text)
}

func BgBlack(text string) string {
	return colorize(BG_BLACK, text)
}

func BgGrey(text string) string {
	return colorize(BG_GREY, text)
}

// Background Bright Colors ---------------------------------------------------

func BgRedBright(text string) string {
	return colorize(BG_RED_BRIGHT, text)
}

func BgGreenBright(text string) string {
	return colorize(BG_GREEN_BRIGHT, text)
}

func BgYellowBright(text string) string {
	return colorize(BG_YELLOW_BRIGHT, text)
}

func BgBlueBright(text string) string {
	return colorize(BG_BLUE_BRIGHT, text)
}

func BgMagentaBright(text string) string {
	return colorize(BG_MAGENTA_BRIGHT, text)
}

func BgCyanBright(text string) string {
	return colorize(BG_CYAN_BRIGHT, text)
}

func BgWhiteBright(text string) string {
	return colorize(BG_WHITE_BRIGHT, text)
}

func BgBlackBright(text string) string {
	return colorize(BG_BLACK_BRIGHT, text)
}

func BgGreyBright(text string) string {
	return colorize(BG_GREY_BRIGHT, text)
}

// Style ----------------------------------------------------------------------

func Bold(text string) string {
	return colorize(BOLD, text)
}

func Dim(text string) string {
	return colorize(DIM, text)
}

func Italic(text string) string {
	return colorize(ITALIC, text)
}

func Underline(text string) string {
	return colorize(UNDERLINE, text)
}

func Strikethrough(text string) string {
	return colorize(STRIKETHROUGH, text)
}
