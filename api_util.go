
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var baseUrl = "https://api.coincap.io/v2/assets"
var defaultLimit = 10

type HistoryInterval struct {
	apiCode      string
	timeInterval time.Duration
}

var (
	MINUTES_1  HistoryInterval = HistoryInterval{"m1", time.Minute}
	MINUTES_5  HistoryInterval = HistoryInterval{"m5", time.Minute * 5}
	MINUTES_15 HistoryInterval = HistoryInterval{"m15", time.Minute * 15}
	MINUTES_30 HistoryInterval = HistoryInterval{"m30", time.Minute * 30}
	HOURS_1    HistoryInterval = HistoryInterval{"h1", time.Hour}
	HOURS_2    HistoryInterval = HistoryInterval{"h2", time.Hour * 2}
	HOURS_6    HistoryInterval = HistoryInterval{"h6", time.Hour * 6}
	HOURS_12   HistoryInterval = HistoryInterval{"h12", time.Hour * 12}
	DAYS_1     HistoryInterval = HistoryInterval{"d1", time.Hour * 24}
)

type AssetJson struct {
	Id                string
	Rank              string
	Symbol            string
	Name              string
	Supply            string
	MaxSupply         string
	MarketCapUsd      string
	VolumeUsd24Hr     string
	PriceUsd          string
	ChangePercent24Hr string
	Vwap24Hr          string
}

type SingleAssetJson struct {
	Data      AssetJson
	Timestamp int64
}

type AssetsJson struct {
	Data      []AssetJson
	Timestamp int64
}

type PriceJson struct {
	PriceUsd          string
	Time              int64
	CirculatingSupply string
	Date              string
}

type HistoryJson struct {
	Data      []PriceJson
	Timestamp int64
}