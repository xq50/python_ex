
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