/*
Functions for file's operations
*/

package tools

import (
	   //"fmt"
       "bufio"
	   "io"
	   "os"
	   "strings"
   	   "container/list"
       )

func ReadLine(fileName string, handler func(int, string, *list.List)) (content *list.List, count int, err error) {
	/*
	Read "filename" file by line.
   
    :param fileName: the file path.
    :param handler: the function for handling file content.
    */

	count = 0
    content = list.New()
    f, err := os.Open(fileName)  
    if err != nil {  
        return content, count, err  
    }  
    buf := bufio.NewReader(f)  
    for {  
        line, err := buf.ReadString('\n')  
        line = strings.TrimSpace(line)  
		count = count + 1
        handler(count, line, content)
        if err != nil {  
            if err == io.EOF {  
                return content, count, nil  
            }  
            return content, count, err
        }  
    }  
    return content, count, nil  
}  
  
func Print(count int, line string, content *list.List){  
    if count > 4 {
    	content.PushBack(line)
	}
}  
