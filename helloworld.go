// Run this file by 
// go run helloworld.go it will run the program from editor 
// if you want to build the code only 
// go build helloworld.go it will produce file helloworld.exe or main file  in linux or mac os  you can execute it by .\helloworld.exe or .\helloworld

package main // Every go file must start with the package name statement , Packages are used to provide code compartmentalization and reusability
             // اى مشروع جديد لازم يكون فية بكدج كانك بتجمع كل الكود فى حاجة واحدة مش شرط يكون اسم ال بكدج main
            // اى حاجة هيتعملها Run لازم هتكون ف func  ال main

import "fmt" // دى package فيها func بتعمل println 

func main() { // اى func مهمة بيتعمل ليها print هنا 

	fmt.Println("Hello World") 
  
}
