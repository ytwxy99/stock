package main

import (
        "fmt"
        "tools"
        "compute"
        //"ftp"
)

func main(){
	contents, _, _ := tools.ReadLine("000007.csv", tools.Print)
    var t = 0
    for content := contents.Front(); content != nil; content = content.Next() {
        //fmt.Println(content.Value)
        t = t + 1
    }
    statistics := compute.Statistics(contents)
	for k, v := range statistics {
		fmt.Println(k, v)
	}
	fmt.Println("large", statistics["large"])
	fmt.Println("middle", statistics["middle"])
	fmt.Println("small", statistics["small"])
	fmt.Println("total", statistics["total"])
    
    //ftp.Run()
}
