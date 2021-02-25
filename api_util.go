
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
