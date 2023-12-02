package app

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gcors88/read_files_go/src/infra/database/entities"
	"github.com/gcors88/read_files_go/src/infra/database/repository"
)

var baseDir string

func ReadFiles() {
	currentDir, _ := os.Getwd()
	baseDir = filepath.Join(currentDir, "files")
	done := make(chan bool)
	go readFolder(baseDir, done)

	finish := <-done

	if finish {
		println("Files processed successfully")
	}
}

func readFolder(dir string, doneCnh chan bool, file ...fs.DirEntry) {
	var currentlyDir string
	if len(file) > 0 {
		currentlyDir = filepath.Join(dir, file[0].Name())
	} else {
		currentlyDir = dir
	}
	files, err := os.ReadDir(currentlyDir)

	if err != nil {
		fmt.Println("Error to open dir", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			readFile(currentlyDir, file)
		} else if file.IsDir() {
			readFolder(currentlyDir, doneCnh, file)
		}
	}

	doneCnh <- true
}

func readFile(dir string, file fs.DirEntry) {
	filePath := filepath.Join(dir, file.Name())

	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		row := fileScanner.Text()
		blocks := strings.Split(row, "^")
		age, _ := strconv.Atoi(strings.TrimSpace(blocks[1]))
		name := strings.TrimSpace(blocks[0])
		profession := strings.TrimSpace(blocks[2])

		user := repository.FindUserByName(strings.TrimSpace(blocks[0]))
		finalValue := fmt.Sprintf("Nome: %s Idade: %d Profissão: %s", strings.TrimSpace(blocks[0]), age, strings.TrimSpace(blocks[2]))
		fmt.Println("USER IN DATABASE", user)
		if len(user.Name) == 0 {
			userToCreate := entities.User{
				Name:       name,
				Age:        age,
				Profession: profession,
			}
			repository.CreateUser(userToCreate)
			fmt.Println("CRATED: ", finalValue)
		} else {
			user.Name = name
			user.Age = age
			user.Profession = profession
			repository.UpdateUser(user)
			fmt.Println("UPDATED: ", finalValue)
		}
	}
}
