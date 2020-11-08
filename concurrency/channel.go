package main
import (
    "fmt"
    "time"
)

type Cmd struct {
    Get bool
    Val int
    Ch chan<- int
}

// a Manager that can help you store a value, and whenever you need it,
// you can take the value back!
func Manager(cmd_ch <-chan Cmd){
    val := 0
    for {
        c := <-cmd_ch
        if c.Get {
            // give the value to the cmd channel
            c.Ch <- val
            fmt.Printf("[!] I am a manager. I put a value %d into the cmd\n", val)
        }else{
            val = c.Val
            fmt.Printf("[!] I am a manager. I store a value %d in me!\n", val)
        }
    }
}

func generate_cmds(cmd_ch chan<- Cmd){
    for i := 0; i < 10; i++ {
        var cmd Cmd
        if i % 2 == 0{
            // put value
            cmd = Cmd{false, i, nil}
            cmd_ch <- cmd
        }else{
            // get value
            ch := make(chan int)
            cmd = Cmd{true, 0, ch}
            cmd_ch <- cmd
            fmt.Printf("[!] Get a value: %d from the manager\n", <-ch)
        }
        // send to the manager
        time.Sleep(3 * time.Second)
    }
}

func main(){
    cmd_ch := make(chan Cmd)
    go Manager(cmd_ch)
    go generate_cmds(cmd_ch)
    for{
    }
}

