package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/codeskyblue/go-uuid"
	"github.com/golang/glog"
	"github.com/golang/snappy"
	"github.com/grafana/worldping-gw/event_publish"
	"gopkg.in/raintank/schema.v1"
	"gopkg.in/raintank/schema.v1/msg"
)

func Events(ctx *Context) {
	contentType := ctx.Req.Header.Get("Content-Type")
	switch contentType {
	case "rt-metric-binary":
		eventsBinary(ctx, false)
	case "rt-metric-binary-snappy":
		eventsBinary(ctx, true)
	case "application/json":
		eventsJson(ctx)
	default:
		ctx.JSON(400, fmt.Sprintf("unknown content-type: %s", contentType))
	}
}

func eventsJson(ctx *Context) {
	defer ctx.Req.Request.Body.Close()
	if ctx.Req.Request.Body != nil {
		body, err := ioutil.ReadAll(ctx.Req.Request.Body)
		if err != nil {
			glog.Errorf("unable to read request body. %s", err)
		}
		event := new(schema.ProbeEvent)
		err = json.Unmarshal(body, event)
		if err != nil {
			ctx.JSON(400, fmt.Sprintf("unable to parse request body. %s", err))
			return
		}
		if !ctx.IsAdmin {
			event.OrgId = int64(ctx.OrgId)
		}

		u := uuid.NewUUID()
		event.Id = u.String()

		err = event_publish.Publish([]*schema.ProbeEvent{event})
		if err != nil {
			glog.Errorf("failed to publish event. %s", err)
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, "ok")
		return
	}
	ctx.JSON(400, "no data included in request.")
}

func eventsBinary(ctx *Context, compressed bool) {
	var body io.ReadCloser
	if compressed {
		body = ioutil.NopCloser(snappy.NewReader(ctx.Req.Request.Body))
	} else {
		body = ctx.Req.Request.Body
	}
	defer body.Close()
	if ctx.Req.Request.Body != nil {
		body, err := ioutil.ReadAll(body)
		if err != nil {
			glog.Errorf("unable to read request body. %s", err)
		}
		ms, err := msg.ProbeEventFromMsg(body)
		if err != nil {
			glog.Errorf("event payload not Event. %s", err)
			ctx.JSON(500, err)
			return
		}

		err = ms.DecodeProbeEvent()
		if err != nil {
			glog.Errorf("failed to unmarshal EventData. %s", err)
			ctx.JSON(500, err)
			return
		}
		if !ctx.IsAdmin {
			ms.Event.OrgId = int64(ctx.OrgId)
		}
		u := uuid.NewUUID()
		ms.Event.Id = u.String()

		err = event_publish.Publish([]*schema.ProbeEvent{ms.Event})
		if err != nil {
			glog.Errorf("failed to publish Event. %s", err)
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, "ok")
		return
	}
	ctx.JSON(400, "no data included in request.")
}
