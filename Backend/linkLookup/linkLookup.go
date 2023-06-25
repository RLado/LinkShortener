package linkLookup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Dest struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DestResponse struct {
	TotalRows int    `json:"total_rows"`
	Offset    int    `json:"offset"`
	Rows      []Dest `json:"rows"`
}

type Time struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value int64  `json:"value"`
}

type TimeResponse struct {
	TotalRows int    `json:"total_rows"`
	Offset    int    `json:"offset"`
	Rows      []Time `json:"rows"`
}

func GetDest(ref string) (DestResponse, error) {
	var destResponse DestResponse

	url := "http://172.16.238.11:5984/link/_design/LinkRefView/_view/LinkRef"

	data := map[string]string{
		"key": ref,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return destResponse, fmt.Errorf("error marshalling data: %s", err)
	}
	fmt.Println(string(payload))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return destResponse, fmt.Errorf("error creating request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("LinkWriter", "passwd")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return destResponse, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()

	// fmt.Println("Response Status:", resp.Status)

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&destResponse)
	if err != nil {
		return destResponse, fmt.Errorf("error decoding response: %s", err)
	}

	return destResponse, nil
}

func GetTimestamp(ref string) (TimeResponse, error) {
	var timeResponse TimeResponse

	url := "http://172.16.238.11:5984/link/_design/LinkTime/_view/LinkTime"

	data := map[string]string{
		"key": ref,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return timeResponse, fmt.Errorf("error marshalling data: %s", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return timeResponse, fmt.Errorf("error creating request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("LinkWriter", "passwd")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return timeResponse, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()

	// fmt.Println("Response Status:", resp.Status)

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&timeResponse)
	if err != nil {
		return timeResponse, fmt.Errorf("error decoding response: %s", err)
	}

	return timeResponse, nil
}
