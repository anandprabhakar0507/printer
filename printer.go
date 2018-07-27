/*
** Printer v0.1.0
** by Momo {M4ll0k} Outaadi
** https://github.com/m4ll0k
*/

package printer

import "fmt"

// colors
var BLACK   = "\033[1;30m"  
var LBLACK  = "\033[0;30m"
var RED     = "\033[1;31m" 
var LRED    = "\033[0;31m" 
var GREEN   = "\033[1;32m"
var LGREEN  = "\033[0;32m"
var YELLOW  = "\033[1;33m"
var LYELLOW = "\033[0;33m"
var BLUE    = "\033[1;34m"
var LBLUE   = "\033[0;34m"
var PURPLE  = "\033[1;35m"
var LPURPLE = "\033[0;35m"
var CYAN    = "\033[1;36m"
var LCYAN   = "\033[0;36m"
var GRAY    = "\033[1;37m"
var LGRAY   = "\033[0;37m"
var WHITE   = "\033[1;38m"
var LWHITE  = "\033[0;38m"
var RESET   = "\033[0m"
// labels
var INFO = "[i]"
var PLUS = "[+]"
var LESS = "[-]"
var TEST = "[*]"
var WARN = "[!]"
var ASK  = "[?]"
//
func Info(str string,color bool) {
	if color{
		fmt.Printf("%s%s%s %s%s%s\n",YELLOW,INFO,RESET,YELLOW,str,RESET)
	} else{
		fmt.Printf("%s%s%s %s\n",YELLOW,INFO,RESET,str)
	}
}
//
func Plus(str string,color bool) {
	if color{
		fmt.Printf("%s%s%s %s%s%s\n",GREEN,PLUS,RESET,GREEN,str,RESET)
	} else{
		fmt.Printf("%s%s%s %s\n",GREEN,PLUS,RESET,str)
	}
}
//
func Less(str string,color bool) {
	if color{
		fmt.Printf("%s%s%s %s%s%s\n",LRED,LESS,RESET,LRED,str,RESET)
	} else{
		fmt.Printf("%s%s%s %s\n",LRED,LESS,RESET,str)
	}
}
//
func Test(str string,color bool) {
	if color{
		fmt.Printf("%s%s%s %s%s%s\n",BLUE,TEST,RESET,BLUE,str,RESET)
	} else{
		fmt.Printf("%s%s%s %s\n",BLUE,TEST,RESET,str)
	}
}
//
func Warn(str string,color bool) {
	if color{
		fmt.Printf("%s%s%s %s%s%s\n",RED,WARN,RESET,RED,str,RESET)
	} else{
		fmt.Printf("%s%s%s %s\n",RED,WARN,RESET,str)
	}
}
//
func Ask(str string,color bool) {
	if color{
		fmt.Printf("%s%s%s %s%s%s\n",CYAN,ASK,RESET,CYAN,str,RESET)
	} else{
		fmt.Printf("%s%s%s %s\n",CYAN,ASK,RESET,str)
	}
}
