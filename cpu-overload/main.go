package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
	"github.com/emersion/go-autostart"
	"github.com/shirou/gopsutil/v4/cpu"
)

var result float64


func main() {

	app := &autostart.App{
		Name: "test",
		DisplayName: "Just a Test App",
		Exec: []string{"apka", "-c", "echo autostart >> ~/main.exe"},
	}
	if app.IsEnabled(){
		fmt.Println("[!] disabling autostart")
		app.Disable()
	} else {
		fmt.Println("[!] enabling autostart")
		app.Enable()
	}

    numCPU := runtime.NumCPU()
    for i := 0; i < numCPU; i++ {
        go func() {
            for {
                result = math.Sqrt(12345.6789)
            }
        }()

		y, _ := cpu.Percent(time.Second, false)
		
		cpup := int(y[0])

		fmt.Println(cpup)
    }

	select {}
}
