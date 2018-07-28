package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/emmanuelq/lotus/dao"
	"github.com/emmanuelq/lotus/models"
)

type LimitedFighter struct {
	ID int `json:"id"`
}

func main() {
	// fighter data is not complete hence the variable name
	limitedFighters, err := getFighters("http://ufc-data-api.ufc.com/api/v1/us/fighters.json")
	if err != nil {
		log.Fatal(err)
	}
	fightersPool(limitedFighters)
}

func getFighters(url string) ([]LimitedFighter, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var limitedFighters []LimitedFighter
	json.Unmarshal(body, &limitedFighters)
	return limitedFighters, nil
}

func fightersPool(limitedFighters []LimitedFighter) {
	fightersChan := make(chan fighter.Fighter)

	var wg sync.WaitGroup
	wg.Add(len(limitedFighters))

	for _, limitedFighter := range limitedFighters {
		go fightersWorker(fightersChan, limitedFighter, &wg)
	}

	go storeFighters(fightersChan)

	wg.Wait()
}

func fightersWorker(fightersChan chan fighter.Fighter, limitedFighter LimitedFighter, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get(
		"http://ufc-data-api.ufc.com/api/v3/us/fighters/" + strconv.Itoa(limitedFighter.ID) + ".json",
	)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var fighter fighter.Fighter
	json.Unmarshal(body, &fighter)
	fightersChan <- fighter
}

func storeFighters(fightersChan <-chan fighter.Fighter) {
	var fightersDAO fightersdao.FightersDAO
	fightersDAO.Connect()

	for fighter := range fightersChan {
		err := fightersDAO.Insert(fighter)
		if err != nil {
			log.Fatal(err)
		}
	}
}
