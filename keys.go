package main

var shiftkeys = map[int][]string{
	50: {"Shift_L"},
	62: {"Shift_R"},
}

var ctrlkeys = map[int][]string{
	37:  {"Control_L"},
	66:  {"Control_L"},
	105: {"Control_R"},
}

var altkeys = map[int][]string{
	108: {"Alt_R"},
	64:  {"Alt_L"},
}

var keymap = map[int][]string{
	8:   {" ", " "},
	9:   {"ESC", " "},
	10:  {"1", "!"},
	11:  {"2", "@"},
	12:  {"3", "#"},
	13:  {"4", "$"},
	14:  {"5", "%"},
	15:  {"6", "^"},
	16:  {"7", "&"},
	17:  {"8", "*"},
	18:  {"9", "("},
	19:  {"0", ")"},
	20:  {"-", "_"},
	21:  {"=", "+"},
	22:  {"⌫", " "},
	23:  {"⌦", " "},
	24:  {"q", "Q"},
	25:  {"w", "W"},
	26:  {"e", "E"},
	27:  {"r", "R"},
	28:  {"t", "T"},
	29:  {"y", "Y"},
	30:  {"u", "U"},
	31:  {"i", "I"},
	32:  {"o", "O"},
	33:  {"p", "P"},
	34:  {"[", "{"},
	35:  {"]", "}"},
	36:  {"⏎", " "},
	38:  {"a", "A"},
	39:  {"s", "S"},
	40:  {"d", "D"},
	41:  {"f", "F"},
	42:  {"g", "G"},
	43:  {"h", "H"},
	44:  {"j", "J"},
	45:  {"k", "K"},
	46:  {"l", "L"},
	47:  {";", ":"},
	48:  {"'", "\""},
	49:  {"~", "`"},
	51:  {"\\", "|"},
	52:  {"z", "Z"},
	53:  {"x", "X"},
	54:  {"c", "C"},
	55:  {"v", "V"},
	56:  {"b", "B"},
	57:  {"n", "N"},
	58:  {"m", "M"},
	59:  {",", "<"},
	60:  {".", ">"},
	61:  {"/", "?"},
	65:  {" ", " "},
	67:  {"F1", " "},
	68:  {"F2", " "},
	69:  {"F3", " "},
	70:  {"F4", " "},
	71:  {"F5", " "},
	72:  {"F6", " "},
	73:  {"F7", " "},
	74:  {"F8", " "},
	75:  {"F9", " "},
	76:  {"F10", " "},
	94:  {"<", ">"},
	95:  {"F11", " "},
	96:  {"F12", " "},
	107: {"Print", " "},
	110: {"Home", " "},
	111: {"Up", " "},
	112: {"Prior", " "},
	113: {"Left", " "},
	114: {"Right", " "},
	115: {"End", " "},
	116: {"Down", " "},
	117: {"Next", " "},
	118: {"Insert", " "},
	119: {"Delete", " "},
}

const (
	Shift_L int = 50
	Shift_R int = 62
)
