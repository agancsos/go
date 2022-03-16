/*
 * Name        : foobar.rb
 * Author      : Abel Gancsos                                                  
 * Version     : 1.0.0.0                
 * Description : Just something I like to port to all languages
 */
package main
import (
    "os"
    "strconv"
    "fmt"
)

func main() {
    var maxFoobar = 17;
    var err error;
    if (os.Args[1] != "") {
        maxFoobar, err = strconv.Atoi(os.Args[1]);
        if err != nil {
            println("Failed to convert string to int.  Falling back to default....");
            maxFoobar = 17;
        }
    }

    for i := 1; i <= maxFoobar; i++ {
        print(fmt.Sprintf("%010s ", strconv.Itoa(i)));
        if i % 15 == 0 {
            println("FOOBAR");
        } else if i % 2 == 0 {
            println("FOO");
        } else if i % 3 == 0 {
            println("BAR");
        } else {
            println("BARFOO");
        }
    }

    os.Exit(0);
}
