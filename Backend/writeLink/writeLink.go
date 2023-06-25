package writeLink

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetSHA1Hash(text string) string {
	algorithm := sha1.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

func AddNew(dest string) (status string, e error) {
	dest_SHA1 := GetSHA1Hash(dest)

	url := "http://172.16.238.11:5984/link/" + dest_SHA1

	data := map[string]interface{}{
		"ref":       dest_SHA1[:8],
		"dest":      dest,
		"timestamp": time.Now().Unix(),
	}

	payload, err := json.Marshal(data)
	if err != nil {
		e = fmt.Errorf("error marshalling data: %s", err)
		return
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		e = fmt.Errorf("error creating request: %s", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("LinkWriter", "passwd")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		e = fmt.Errorf("error making request: %s", err)
		return
	}
	defer resp.Body.Close()

	return resp.Status, nil
}
