package main

import (
	"encoding/json"
	"fmt"
	pb "github.com/gawth/go-proto/hotel"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
)

//go:generate msgp
type HotelDetailsResponseModel struct {
	Hotels []Hotel
}

type Hotel struct {
	Id               int
	HotelCode        string
	HotelStatus      string
	HotelCurrency    string
	MinutesOffsetUtc int
	ProviderName     string
	Rooms            []Room
}

type Room struct {
	Id           int
	InvCode      string
	RackRate     string
	HideRackRate bool
}

func hit(method string, url string, body string, headers http.Header) (resp *http.Response, err error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = headers
	return client.Do(req)
}

func hitUrl(contentType string, url string) []byte {
	h := http.Header{}
	h.Add("ContentType", contentType)
	res, err := hit("GET", url, "", h)
	if err != nil {
		log.Fatal(err)
	}

	payload, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	return payload
}
func processMsgPack(raw []byte) HotelDetailsResponseModel {

	hd := HotelDetailsResponseModel{}

	_, err := hd.UnmarshalMsg(raw)
	if err != nil {
		log.Fatal(err)
	}

	return hd
}
func processJson(raw []byte) HotelDetailsResponseModel {
	hd := HotelDetailsResponseModel{}

	err := json.Unmarshal(raw, &hd)
	if err != nil {
		log.Fatal(err)
	}
	return hd

}

func processProto(raw []byte) pb.HotelDetailsResponseModel {
	hd := pb.HotelDetailsResponseModel{}

	err := proto.Unmarshal(raw, &hd)
	if err != nil {
		log.Fatal(err)
	}

	return hd
}

func main() {
	tmp := hitUrl("application/x-msgpack", "http://10.211.55.3:8888/testmsgpack")
	res := processMsgPack(tmp)
	fmt.Printf("Hotels: %v", res)
}
