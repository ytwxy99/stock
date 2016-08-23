/*
invoke python file 'download.py'
*/

package ftp

import (
        "os/exec"
        )

func Run() {
<<<<<<< HEAD
    cmd := exec.Command("ping 127.0.0.1") 
=======
    cmd := exec.Command("ping 127.0.0.1")
>>>>>>> 5d2447946630a55d7e0521c00cf28aaf3313fadc
    _, err := cmd.Output()
    if err != nil {
        panic(err.Error())
    }
<<<<<<< HEAD
                        
    if err := cmd.Start(); err != nil {
        panic(err.Error())
    }
                                              
=======

    if err := cmd.Start(); err != nil {
        panic(err.Error())
    }

>>>>>>> 5d2447946630a55d7e0521c00cf28aaf3313fadc
    if err := cmd.Wait(); err != nil {
        panic(err.Error())
    }
}
