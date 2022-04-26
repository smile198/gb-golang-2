package main

import (
    "fmt"
    "gb-golang-2/Lesson1/HomeWork/task1"
    "gb-golang-2/Lesson1/HomeWork/task2"
    "gb-golang-2/Lesson1/HomeWork/task3"
    "gb-golang-2/Lesson1/HomeWork/task4"
)

func main() {
    fmt.Println("task1")
    task1.Run()
    fmt.Printf("task1 complete\n\n")

    fmt.Println("task2")
    task2err := task2.Run()
    if task2err != nil {
        fmt.Println(task2err)
    }
    fmt.Printf("task2 complete\n\n")

    // Исходя из полученной ошибки, максималькое кол-во открытых фалов моей ОС - 10237 (+ уже открытые файлы)
    // Чтобы не насиловать комп (опять), для выполнения задачи снижаю лимит до 10300
    fmt.Println("task3")
    //fileLimit := 1_000_000
    fileLimit := 10_300
    task3err := task3.Run("./millionFileDir", fileLimit)
    if task3err != nil {
        // open ./millionFileDir/file_10237.txt: too many open files
        fmt.Println(task3err)
    }
    // rm -rf millionFileDir/*.txt
    fmt.Printf("task3 complete\n\n")

    fmt.Println("task4")
    task4.Run()
    fmt.Printf("task4 complete\n\n")
}
