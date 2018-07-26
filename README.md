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

Usage
--
```
// print the string without color
printer.Info("Hello Word..",false)

// print the string with color
printer.Info("Hello Word..",true)

```
