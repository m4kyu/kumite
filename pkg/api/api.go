package api

import (
	"github.com/m4kyu/kumite/internal/utils"
)

type Organization struct {
	ID 		int 		`json:"orgId"`
	Name 	string 	`json:"orgName"`
	Flag 	string 	`json:"orgFlag"`
}


type Champ struct {
	ID 	string `json:"champId"` // Solve incoorect id type
	
	RegFrom 	string `json:"champRegFrom"`
	RegTo 		string `json:"champRegTo"`
	From 			string `json:"champFrom"`
	To 				string `json:"champTo"`
	
	Type 			string `json:"champType"`
	Tittle 		string `json:"title"`
	Email 		string `json:"emailorg"`
	EmailTech string `json:"emailtech"`
	OrgID 		string `json:"orgId"`	 
	ClubID 		string `json:"champClubId"`
}


type ChampInfo struct {
	Categories map[string]string `json:"Categories"`
	Subchamps  map[string]string `json:"subchamps"`
	Info  		 string            `json:"champInfo"`
	Address 	 string            `json:"address"`
}

type Club struct {
  ID		int    `json:"id"`
	Name 	string `json:"name"`
	Count	int    `json:"count"`
	Logo 	string `json:"logo"`
}

type Coach struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}


func Organizations() (*[]Organization, error) {
	return utils.FetchAPI[[]Organization]("https://alliance-kumite.net/api-get-organization")
}


func Champs() (*[]Champ, error) {
	champs, err := utils.FetchAPI[map[string]Champ]("https://alliance-kumite.net/api-champ-get-all-champs-info")
	if err != nil {
		return nil, err
	}

	return utils.ToStruct(champs), nil
}

func ChampClubs(champTittle string) (*[]Club, error) {
	data := map[string]string{
		"title": champTittle,
	} 

	clubs, err := utils.PostAPI[map[string]Club]("https://alliance-kumite.net/api-champ-get-clubs", data)
	if err != nil {
		return nil, err
	}

	return utils.ToStruct(clubs), nil
}


func ChampCoaches(champTittle string) (*[]Coach, error) {
	data := map[string]string{
		"title": champTittle,
	} 

	coahes, err := utils.PostAPI[map[string]Coach]("https://alliance-kumite.net/api-champ-get-coaches", data)
	if err != nil {
		return nil, err
	}


	return utils.ToStruct(coahes), nil
} 


