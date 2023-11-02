package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenValidParams_WhenCallCreateTask_ThenSuccess(t *testing.T) {

	command := &CreateTaskCommand{
		Name:    "English class",
		DueTime: "18:00",
	}

	task, err := CreateTask(command)

	assert.Nil(t, err)
	assert.NotNil(t, task)
	assert.IsType(t, &Task{}, task)
	assert.NotNil(t, task.Id)
	assert.NotNil(t, task.CreatedAt)
	assert.NotNil(t, task.UpdatedAt)
	assert.Equal(t, command.Name, task.Name)
	assert.Equal(t, command.DueTime, task.DueTime)
	assert.Equal(t, false, task.IsArchive)
}

func Test_GivenEmptyName_WhenCallCreateTask_ThenError(t *testing.T) {

	command := &CreateTaskCommand{
		Name:    "",
		DueTime: "18:00",
	}
	errorMessage := "Task name is required"

	task, err := CreateTask(command)

	assert.NotNil(t, err)
	assert.Nil(t, task)

	assert.Equal(t, errorMessage, err.Error())
}

func Test_GivenEmptyTime_WhenCallCreateTask_ThenError(t *testing.T) {

	command := &CreateTaskCommand{
		Name:    "English class",
		DueTime: "18:00",
	}
	errorMessage := "Task due time is required"

	task, err := CreateTask(command)

	assert.NotNil(t, err)
	assert.Nil(t, task)

	assert.Equal(t, errorMessage, err.Error())
}
