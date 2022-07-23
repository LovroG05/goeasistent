package goeasistent

import (
	"encoding/json"

	"github.com/LovroG05/goeasistent/endpoints"
	"github.com/LovroG05/goeasistent/objects"
)

type Goeasistent interface {
	Login(username string, password string) error
	RefreshToken(sessionToken string) error
	GetChild() (objects.ChildData, error)
	GetTimetable(from string, to string) (objects.Timetable, error)
}

type Instance struct {
	UserData objects.UserData
}

func (i *Instance) Login(username string, password string) error {
	userjson, err := endpoints.Login(username, password)
	if err != nil {
		return err
	}

	var userdata objects.UserData

	err = json.Unmarshal([]byte(userjson), &userdata)
	if err != nil {
		return err
	}

	i.UserData = userdata

	return nil
}

func (i *Instance) RefreshToken(sessionToken string) error {
	tokens, err := endpoints.RefreshToken(i.UserData.Tokens.RefreshToken, sessionToken)
	if err != nil {
		return err
	}

	i.UserData.Tokens = tokens

	return nil
}

func (i *Instance) GetChild() (objects.ChildData, error) {
	childjson, err := endpoints.GetChild(i.UserData.Tokens.AccessToken.Token)
	if err != nil {
		return objects.ChildData{}, err
	}

	var childdata objects.ChildData

	err = json.Unmarshal([]byte(childjson), &childdata)
	if err != nil {
		return objects.ChildData{}, err
	}

	return childdata, nil
}

func (i *Instance) GetTimetable(from string, to string) (objects.Timetable, error) {
	timetablejson, err := endpoints.Timetable(i.UserData.Tokens.AccessToken.Token, from, to)
	if err != nil {
		return objects.Timetable{}, err
	}

	var timetable objects.Timetable

	err = json.Unmarshal([]byte(timetablejson), &timetable)
	if err != nil {
		return objects.Timetable{}, err
	}

	return timetable, nil
}
