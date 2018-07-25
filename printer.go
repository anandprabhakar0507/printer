/*
** Printer v0.1.0
** by Momo {M4ll0k} Outaadi
** https://github.com/m4ll0k
*/

package printer

import "fmt"

// colors
var black  = 30
var red    = 31
var green  = 32
var yellow = 33
var blue   = 34
var purple = 35
var cyan   = 36
var gray   = 37
var white  = 38
// labels
var info = "[i]"
var plus = "[+]"
var less = "[-]"
var test = "[*]"
var warn = "[!]"
var ask  = "[?]"


// [i] Hello Word
func Info(str string,colorable bool) {
	if colorable {
		fmt.Printf("\033[1;%dm%s\033[0m \033[1;%dm%s\033[0m\n",yellow,info,yellow,str)
	} else {
		fmt.Printf("\033[1;%dm%s\033[0m %s\n",yellow,info,str)
	}
}

// [+] Hello Word
func Plus(str string,colorable bool) {
	if colorable {
		fmt.Printf("\033[1;%dm%s\033[0m \033[1;%dm%s\033[0m\n",green,plus,green,str)
	} else {
		fmt.Printf("\033[1;%dm%s\033[0m %s\n",green,plus,str)
	}
}

// [-] Hello Word
func Less(str string,colorable bool) {
	if colorable {
		fmt.Printf("\033[0;%dm%s\033[0m \033[0;%dm%s\033[0m\n",red,less,red,str)
	} else {
		fmt.Printf("\033[0;%dm%s\033[0m %s\n",red,less,str)
	}
}

// [*] Hello Word
func Test(str string,colorable bool) {
	if colorable {
		fmt.Printf("\033[1;%dm%s\033[0m \033[1;%dm%s\033[0m\n",blue,test,blue,str)
	} else {
		fmt.Printf("\033[1;%dm%s\033[0m %s\n",blue,test,str)
	}
}

// [!] Hello Word
func Warn(str string,colorable bool) {
	if colorable {
		fmt.Printf("\033[1;%dm%s\033[0m \033[1;%dm%s\033[0m\n",red,warn,red,str)
	} else {
		fmt.Printf("\033[1;%dm%s\033[0m %s\n",red,warn,str)
	}
}

// [?] Hello Word
func Ask(str string,colorable bool) {
	if colorable {
		fmt.Printf("\033[1;%dm%s\033[0m \033[1;%dm%s\033[0m\n",cyan,ask,cyan,str)
	} else {
		fmt.Printf("\033[1;%dm%s\033[0m %s\n",cyan,ask,str)
	}
}