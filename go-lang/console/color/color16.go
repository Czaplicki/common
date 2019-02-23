package color


var Const colors = colors {

	"\u001b[0m",
	"\u001b[30m",
	"\u001b[31m",
	"\u001b[32m",
	"\u001b[33m",
	"\u001b[34m",
	"\u001b[35m",
	"\u001b[36m",
	"\u001b[37m",

	"\u001b[30;1m",
	"\u001b[31;1m",
	"\u001b[32;1m",
	"\u001b[33;1m",
	"\u001b[34;1m",
	"\u001b[35;1m",
	"\u001b[36;1m",
	"\u001b[37;1m",

}

type colors struct {


	Reset string
	Black,			Red,			Green,			Yellow			string
	BrightBlack,	BrightRed,		BrightGreen,	BrightYellow	string

	Blue,			Magenta,		Cyan,			White			string
	BrightBlue,		BrightMagenta,	BrightCyan,		BrightWhite 	string

}


func Black		(text string) string { return Const.Black	+ text + Const.Reset }
func Red		(text string) string { return Const.Red		+ text + Const.Reset }
func Green		(text string) string { return Const.Green	+ text + Const.Reset }
func Yellow		(text string) string { return Const.Yellow	+ text + Const.Reset }
func Blue		(text string) string { return Const.Blue	+ text + Const.Reset }
func Magenta	(text string) string { return Const.Magenta	+ text + Const.Reset }
func Cyan		(text string) string { return Const.Cyan	+ text + Const.Reset }
func White		(text string) string { return Const.White	+ text + Const.Reset }

func BrightBlack	(text string) string { return Const.BrightBlack		+ text + Const.Reset }
func BrightRed		(text string) string { return Const.BrightRed		+ text + Const.Reset }
func BrightGreen	(text string) string { return Const.BrightGreen		+ text + Const.Reset }
func BrightYellow	(text string) string { return Const.BrightYellow	+ text + Const.Reset }
func BrightBlue		(text string) string { return Const.BrightBlue		+ text + Const.Reset }
func BrightMagenta	(text string) string { return Const.BrightMagenta	+ text + Const.Reset }
func BrightCyan		(text string) string { return Const.BrightCyan		+ text + Const.Reset }
func BrightWhite	(text string) string { return Const.BrightWhite		+ text + Const.Reset }
