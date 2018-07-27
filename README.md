Printer 
--
Printer provides a minimal and powerful interface to print colored labels and text in the terminal.

![screen](https://raw.githubusercontent.com/m4ll0k/printer/master/screen.png)

Installation
--
```
$ go get github.com/m4ll0k/printer
```
__OR:__

```
$ wget https://raw.githubusercontent.com/m4ll0k/printer/master/printer.go
$ mv printer.go myproject/lib
```

Supported
--
__Supported Labels:__

```go
// [i] ...
printer.Info(str string, color bool)
// [+] ...
printer.Plus(str string, color bool)
// [-] ..
printer.Less(str string, color bool)
// [*] ..
printer.Test(str string, color bool)
// [!] ...
printer.Warn(str string, color bool)
// [?] ...
printer.Ask(str string, color bool)
```

Usage
--
```go
// print the string without color
printer.Info("Hello Word..",false)

// print the string with color
printer.Info("Hello Word..",true)

// print sigle color

fmt.Printf("%sMy String Here%s\n",printer.YELLOW,printer.RESET)

// Costum your print
fmt.Printf("%s%s%s %sHello Word!%s\n",printer.RED,printer.WARN,printer.RESET,printer.LRED,printer.RESET)
```
