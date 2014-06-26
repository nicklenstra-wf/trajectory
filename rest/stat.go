package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lyddonb/trajectory/api"
)

type StatServices struct {
	api *api.StatAPI
}

func NewStatServices(statAPI *api.StatAPI) *StatServices {
	return &StatServices{statAPI}
}

func (s *StatServices) addStat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding stat")

	decoder := json.NewDecoder(r.Body)

	var statJson map[string]*json.RawMessage

	err := decoder.Decode(&statJson)

	if err != nil {
		SendJsonErrorResponse(w, &statJson, err.Error())
		return
	}

	request, stat := s.api.MakeRequestStats(statJson)

	fmt.Println(request.RequestId)
	fmt.Println(request.MachineInfo)
	fmt.Println(request.Url)

	//_, ok := task[db.REQUEST_ADDRESS]

	//if !ok {
	//task[db.REQUEST_ADDRESS] = r.Host
	//}

	timestamp, e := s.api.SaveRequestStats(request, stat)

	if e != nil {
		SendJsonErrorResponse(w, &statJson, e.Error())
		fmt.Println("Stat failed with %s at %s", e, timestamp)
		return
	}

	SendJsonResponse(w, timestamp, nil)

	fmt.Println("Stat saved at %s", timestamp)
}

func (s *StatServices) getAllStats(w http.ResponseWriter, r *http.Request) {
	SendJsonResponse(w, "Test", nil)
}

func (s *StatServices) getStatByRequestId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	requestId := params["requestId"]

	if requestId == "" {
		SendJsonErrorResponse(w, nil, "No request id passed in.")
	}

	stat, err := s.api.GetStatForRequest(requestId)

	SendJsonResponse(w, stat, err)
}
