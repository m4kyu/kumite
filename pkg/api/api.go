package api

import (
	"github.com/m4kyu/kumite/internal/utils"
	"github.com/m4kyu/kumite/internal/types"
)

type Organization struct {
	ID 		int 		`json:"orgId"`
	Name 	string 	`json:"orgName"`
	Flag 	string 	`json:"orgFlag"`
}


type Champ struct {
	ID types.KumiteInt `json:"champId"` 
	
	RegFrom 	string `json:"champRegFrom"`
	RegTo 		string `json:"champRegTo"`
	From 			string `json:"champFrom"`
	To 				string `json:"champTo"`
	
	Tittle 		string `json:"title"`
	Email 		string `json:"emailorg"`
	EmailTech string `json:"emailtech"`
	
	Type 			types.KumiteInt `json:"champType"`
  OrgID 		types.KumiteInt `json:"orgId"`	 
	ClubID 		types.KumiteInt `json:"champClubId"`
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

type Participant struct {
  FIO	 string `json:"FIO"`
	FIO1 string `json:"FIO1"`
	FIO2 string `json:"FIO2"`
	FIO3 string `json:"FIO3"`
  FIO4 string `json:"FIO4"`

  Dan   string `json:"DAN"`
  Birth string `json:"DateBR"`
  

	Age	int `json:"Age"`
  Photo	string `json:"Photo"`
  
  Weight    types.KumiteFloat	`json:"Weight"`
	Kumite    types.KumiteInt   `json:"Kumite"`
  Kata      types.KumiteInt   `json:"Kata"`
  KataGroup	types.KumiteInt   `json:"KataGroup"`
  Country	  types.KumiteInt   `json:"CountryId"` 
  Region	  types.KumiteInt   `json:"RegionId"`
  Club	    types.KumiteInt   `json:"ClubId"`
  Coach	    types.KumiteInt   `json:"CoachId"`
  
	Subtournament	types.KumiteInt   `json:"subtournamentId"`
  Categories    types.KumiteSlice `json:"categoryId"`

  SoundUA     string `json:"SoundUA"` 	
  ClubName	  string `json:"ClubName"`
	ClubCity 	  string `json:"ClubCity"`
  CoachName	  string `json:"CoachName"`
  CountryFlag	string `json:"CountryFlag"`
  ClubLogo    string `json:"ClubLogo"`
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
	data := map[string]any {
		"title": champTittle,
	} 

	clubs, err := utils.PostAPI[map[string]Club]("https://alliance-kumite.net/api-champ-get-clubs", data)
	if err != nil {
		return nil, err
	}

	return utils.ToStruct(clubs), nil
}


func ChampCoaches(champTittle string) (*[]Coach, error) {
	data := map[string]any {
		"title": champTittle,
	} 

	coahes, err := utils.PostAPI[map[string]Coach]("https://alliance-kumite.net/api-champ-get-coaches", data)
	if err != nil {
		return nil, err
	}


	return utils.ToStruct(coahes), nil
} 

func ParticipantsCount(champTittle string) (int, error) {
	data := map[string]any {
		"title": champTittle,
		"category": nil, 
		"coach": nil,
		"club": nil,
		"coachname": nil,
		"clubname": nil,
		"country": nil,
	}

		
	count, err := utils.PostAPI[struct {Count int `json:"count"`}]("https://alliance-kumite.net/api-champ-get-participants-count", data)
	if err != nil {
		return 0, err
	}

	return count.Count, nil
}

func Participants(champTittle string) (*[]Participant, error) {
	data := map[string]any {
    "title": champTittle,
    "category": nil,
    "coach": nil,
    "club": nil,
    "coachname": nil,
    "clubname": nil,
    "country": nil,
	}

	fighters, err := utils.PostAPI[struct {Participants map[string]Participant}]("https://alliance-kumite.net/api-champ-get-participants", data)
	if err != nil {
		return nil, err
	}

  return utils.ToStruct(&fighters.Participants), nil
}



