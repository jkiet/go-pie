package listener

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/jkiet/go-pie/driver"
	"net/http"
)

type Response struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"_meta"`
}

type ListenerSvc struct {
	Section *driver.Section
}

func NewListenerSvc(s *driver.Section) *ListenerSvc {
	return &ListenerSvc{Section: s}
}

func (svc *ListenerSvc) reload(request *restful.Request, response *restful.Response) {
	meta := make(map[string]string)
	var data map[string]string
	err := request.ReadEntity(&data)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		meta["error"] = fmt.Sprintf("Bad request: %v", err)
		response.WriteAsJson(Response{Meta: meta})
		return
	}
	meta["status"] = "ok"
	response.WriteEntity(Response{Meta: meta, Data: svc.Section.Reload(data)})
}

func (svc *ListenerSvc) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/lamps").
		Doc("lamps service").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.POST("/reload/").To(svc.reload).
		Doc("reload commands for lamps").
		Operation("reload"))
	container.Add(ws)
}
