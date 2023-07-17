package filereaderr

import (
	"fmt"
	"os"

	"github.com/yousefzinsazk78/test_2_web_app/model/post"
	"github.com/yousefzinsazk78/test_2_web_app/model/user"
)

// CreateFile method create file if doesn't exists...
func CreateFile(filename string) error {
	file, err := os.Create("./stored_file/" + filename + ".txt")
	if err != nil {
		return fmt.Errorf("i can't create this file because \n%s", err)
	}
	file.Close()
	return nil
}

func WriteData(filename string, user user.User) error {
	file, err := os.OpenFile("./stored_file/"+filename+".txt", os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(user.String())
	if err != nil {
		return err
	}
	fmt.Println("write operation completed!")
	return nil
}

func WriteDataBlog(filename string, postBlog post.Post) error {
	file, err := os.OpenFile("./stored_file/"+filename+".txt", os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(postBlog.String())
	if err != nil {
		return err
	}
	fmt.Println("write operation completed!")
	return nil
}

// ReadFile method read filename given and return data and error
func ReadFile(filename string) (string, error) {
	readFile, err := os.ReadFile("./stored_file/" + filename + ".txt")
	if err != nil {
		return "", err
	}
	return string(readFile), nil
}
