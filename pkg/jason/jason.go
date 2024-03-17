package jason

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/senyc/jason-cli/pkg/auth"
	"github.com/senyc/jason-cli/pkg/types"
)

const (
	url = "http://172.20.0.4:8080"
)

func AddNewTask(title string, priority int, date string) error {
	newTaskPayload := types.NewTaskPayload{Title: title, Priority: int16(priority), Due: date}
	j, err := json.Marshal(newTaskPayload)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/api/tasks/new", url), bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	// Get the users stored key
	key, err := auth.GetKeyFromFS()
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", key)

	// Send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer response.Body.Close()
	return nil
}

