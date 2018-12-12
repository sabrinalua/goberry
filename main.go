package main 

import (
	"fmt"
	"github.com/sabrinalua/s3"
)

func main(){
	fmt.Println("go-app main")
	x:= gos3.InitiateMultipartUploadResult{}
	xml,_:= gos3.EncodeXml(true, x)
	fmt.Printf("xml %s\n", xml)
}