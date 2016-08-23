/*
invoke python file 'download.py'
*/

package ftp

import (
        "os/exec"
        )

func Run() {
    cmd := exec.Command("ping 127.0.0.1") 
    _, err := cmd.Output()
    if err != nil {
        panic(err.Error())
    }
                        
    if err := cmd.Start(); err != nil {
        panic(err.Error())
    }
                                              
    if err := cmd.Wait(); err != nil {
        panic(err.Error())
    }
}
