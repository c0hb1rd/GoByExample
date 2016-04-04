package main

import "os"
import F "fmt"


var P = F.Printf

type point struct {
    x, y int
}

func main() {
    p := point{1, 2}

    P("[*]Source is type of custom structure \"point\": point{1, 2}\n")
    P("[-->]Formatting by %s: %v\n", "%v", p)
    P("[-->]Formatting by %s: %+v\n", "%+v", p)
    P("[-->]Formatting by %s: %#v\n", "%#v", p)
    P("[-->]Formatting by %s: %T\n", "%T", p)
    P("[-->]%s print the variable pointer: %p\n", "%p", &p)

    P("\n[*]Source is boolean value \"true\"\n")
    P("[-->]Formatting by %s: %t\n", "%t", true)

    P("\n[*]Source is integer value \"250\"\n")
    P("[-->]Formatting by %s: %d\n", "%d", 250)
    P("[-->]Formatting by %s: %b\n", "%b", 250)
    P("[-->]Formatting by %s: %c\n", "%c", 250)
    P("[-->]Formatting by %s: %x\n", "%x", 250)
    P("[-->]Formatting by %s: |%6d|\n", "%6d", 250)
    P("[-->]Formatting by %s: |%-6d|\n", "%-6d", 250)

    P("\n[*]Source is float value \"250.25\"\n")
    P("[-->]Formatting by %s: %f\n", "%f", 250.25)
    P("[-->]Formatting by %s: %e\n", "%e", 250.25)
    P("[-->]Formatting by %s: %E\n", "%E", 250.25)
    P("[-->]Formatting by %s: |%6.1f|\n", "%6.1f", 250.25)
    P("[-->]Formatting by %s: |%-6.1f|\n", "%-6.1f", 250.25)


    P("\n[*]Source is a string value \"This is a string\"\n")
    P("[-->]Formatting by %s: %s\n", "%s", "This is a string")
    P("[-->]Formatting by %s: %q\n", "%q", "This is a string")
    P("[-->]Formatting by %s: %x\n", "%x", "This is a string")
    P("[-->]Formatting by %s: |%20s|\n", "%20s", "This is a string")
    P("[-->]Formatting by %s: |%-20s|\n", "%-20s", "This is a string")

    P("\n[*]Print data to standard error buffer(os.Stderr)\n")
    F.Fprintf(os.Stderr, "[-->]%s\n", "An error")

}
