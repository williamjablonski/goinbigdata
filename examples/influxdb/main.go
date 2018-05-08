//  https://tutorialedge.net/golang/parsing-xml-with-golang/

/*
-------
Duracao:  =(getEndTime-getStartTime)

getEndTime("<item name="Responsavel">williamjablonski@dellemc.com</item>
</metadata>
<status status="PASS" endtime="Value" starttime=")

getStartTime(<item name="Responsavel">williamjablonski@dellemc.com</item>
</metadata>
<status status="PASS" endtime="*" starttime="Value"></status>)
-------
Suite: count("<suite>")
-------
Test: count("<test")
-------
Keyword: count("<kw")
-------
*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"math/rand"
	"time"
	"io/ioutil"
	"os"
	"encoding/xml"
)

const (
	database = "nodes"
	username = "monitor"
	password = "secret"
)

var clusters = []string{"public", "private"}

func main() {
	//create conection
	c := influxDBClient()
	//parse XML
	parseXml()
	//send data
	// ?
}

func influxDBClient() client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	return c
}

func parseXml(){
	
	// Open our xmlFile
	xmlFile, err := os.Open("reports/meuvivo/output.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened output.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()
	
	
}
