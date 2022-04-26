package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

type HelloWorldStats struct {
    Count     int
    FileNames []string
}

var dirPath = flag.String("dirPath", "", "Path to dir with files")
var substrRequired = "Hello World"

func main() {
    flag.Parse()
    if *dirPath == "" {
        panic("dirPath cannot be empty")
    }

    stats, err := getStats(*dirPath)
    if err != nil {
        fmt.Printf("%s\n", err)
    }

    fmt.Println(stats.Count, stats.FileNames)
}

func hasSubstr(filePath string) (bool, error) {
    file, err := os.Open(filePath)
    defer file.Close()
    if err != nil {
        return false, err
    }

    b, err := ioutil.ReadAll(file)
    if err != nil {
        return false, err
    }

    return strings.Contains(string(b), substrRequired), nil
}

func getStats(dirPath string) (*HelloWorldStats, error) {
    stats := &HelloWorldStats{
        Count:     0,
        FileNames: make([]string, 0),
    }

    allFiles, err := ioutil.ReadDir(dirPath)
    if err != nil {
        return nil, err
    }

    for _, f := range allFiles {
        filePath := dirPath + "/" + f.Name()

        ok, err := hasSubstr(filePath)
        if err != nil {
            return nil, ErrFailedToGetStats{path: filePath, innerError: err.Error()}
        }

        if ok {
            stats.Count++
            stats.FileNames = append(stats.FileNames, f.Name())
        }
    }

    return stats, nil
}

type ErrFailedToGetStats struct {
    path       string
    innerError string
}

func (e ErrFailedToGetStats) Error() string {
    fmt.Sprintf("Failed to get stats for path: %s, reason: %s", e.path, e.innerError)
}
