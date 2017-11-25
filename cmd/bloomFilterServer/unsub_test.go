/* This file will contain functionality for a simulated
client. The client is responsible for makine http requests
to the dbServer and bloomFilterServer
*/

package bloomFilterServer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"Inf191BloomFilter/bloomDataGenerator"
	"Inf191BloomFilter/databaseAccessObj"
)

const membershipEndpoint = "http://localhost:9090/filterUnsubscribed"

func TestUnsub(t *testing.T) {
	// var payload Payload
	var dataSum []string
	buff := new(bytes.Buffer)

	// Generate random id_email pairs (positives) and save it in a var
	inDB := bloomDataGenerator.GenData(1, 100, 200)

	dao := databaseAccessObj.New("bloom:test@/unsubscribed")
	// Insert new data in the db
	dao.Insert(inDB)

	// Call BF server to update the bit array
	// res, err := http.Get(update_bit_array_endpoint)
	// checkErr(err)

	// Generate more raandom id_email pairs (negatives) and save it ina var
	notInDB := bloomDataGenerator.GenData(1, 50, 100)

	for userid, emails := range inDB {
		dataSum = append(emails, notInDB[userid]...)
	}

	// Put values into payload to be sent to the server later
	payload = Payload{0, dataSum}

	// convert to json
	data, err := json.Marshal(payload)
	checkErr(err)

	// encode to bytes buffer
	err = json.NewEncoder(buff).Encode(data)
	checkErr(err)

	var deres []int8
	var arr Result
	res, _ := http.Post(membershipEndpoint, "application/json; charset=utf-8", buff)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(body)

	err = json.NewDecoder(res.Body).Decode(&deres)

	err = json.Unmarshal(body, &arr)
	checkErr(err)
	fmt.Println(arr.Trues)
	}
}

