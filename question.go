package main

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"time"
)

type question struct {
	Key    uint64 `json:"key,omitempty"`
	Value  string `json:"value,omitempty"`
	Answer string `json:"answer,omitempty"`
}

func (a *App) Save(questionnaire []question) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	saveOtp := runtime.SaveDialogOptions{
		DefaultDirectory:           pwd,
		DefaultFilename:            fmt.Sprintf("questionnaire_%s", time.Now().Format("20060102_150405")),
		Title:                      "Save This Questionnaire",
		Filters:                    nil,
		ShowHiddenFiles:            false,
		CanCreateDirectories:       true,
		TreatPackagesAsDirectories: false,
	}
	path, err := runtime.SaveFileDialog(a.ctx, saveOtp)
	if err != nil {
		return "", err
	}
	if len(path) == 0 {
		return "", nil
	}
	data, err := jsoniter.Marshal(questionnaire)
	if err != nil {
		return "", errors.New("failed to save the questionnaire")
	}
	err = os.WriteFile(fmt.Sprintf("%s.zzd", path), data, 0666)
	return "save success", err
}

func (a *App) OpenQuestionnaire() ([]*question, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	openOpt := runtime.OpenDialogOptions{
		DefaultDirectory: pwd,
		DefaultFilename:  "",
		Title:            "",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Questionnaire File",
				Pattern:     "*.zzd",
			},
		},
		ShowHiddenFiles:            false,
		CanCreateDirectories:       true,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	}
	path, err := runtime.OpenFileDialog(a.ctx, openOpt)
	if err != nil {
		return nil, err
	}
	if len(path) == 0 {
		return nil, nil
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var questionList []*question
	if len(file) > 0 {
		err = jsoniter.Unmarshal(file, &questionList)
	}
	return questionList, err
}
