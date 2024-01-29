package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[2]
	var fileName string
	var fileExt string
	
	// var fileName strings.Builder 
	// var fileExt strings.Builder

	filePthSplit := strings.Split(filePth, "/")
	lastSplit := filePthSplit[len(filePthSplit)-1] //получаем название файла и расширение

	// if strings.LastIndex(lastSplit, ".") == -1{
	// 	fileName.WriteString(fmt.Sprintf(lastSplit))
	// } else {
	// 	for ind, v := range lastSplit {
	// 		if ind < strings.LastIndex(lastSplit, ".") {
	// 			fileName.WriteString(fmt.Sprintf(string(v)))
	// 		} else if ind > strings.LastIndex(lastSplit, "."){
	// 			fileExt.WriteString(fmt.Sprintf(string(v)))
	// 		}
	// 	}
	// }
	var indOfPoint int = strings.LastIndex(lastSplit, ".")
	fileName = lastSplit[0:indOfPoint]
	fileExt = lastSplit[indOfPoint:]
	
	
	// Напишите код, который выведет следующее
	// filename: <fileName>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex
	// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
	// ) Они могут помочь для проверки решения
	fmt.Printf("filename: %v\n", fileName)
	fmt.Printf("filename: %v\n", fileExt)
	// fmt.Printf("filename: %v\n", fileName.String())
	// fmt.Printf("extension: %v\n", fileExt.String())
}