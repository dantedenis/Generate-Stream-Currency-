package api

import (
	"fmt"
	mock_contract "generate_stream_currency/mock"
	"generate_stream_currency/pkg/model"
	_ "github.com/buaazp/fasthttprouter"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"net"
	"testing"
)

func Test_Health(t *testing.T) {
	port := ":1234"
	host := "http://localhost"
	serv := NewServer(nil)
	defer startServerOnPort(t, port, serv.health).Close()

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(host + port + "/health")
	req.Header.SetMethod("GET")
	req.Header.SetContentType("application/json")

	resp := fasthttp.AcquireResponse()

	assert.Nil(t, fasthttp.Do(req, resp))

	if resp.StatusCode() != 200 {
		t.Error("Error status code")
	}
	if len(resp.Body()) != 0 {
		t.Error("Error Body response")
	}
}

func Test_GetVal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ret := model.Currency{
		Value: map[string]*model.Rater{
			"test": model.NewRate(),
		},
	}

	apiService := mock_contract.NewMockApiService(ctrl)
	apiService.EXPECT().GetCurrency().Return(ret)

	serv := NewServer(apiService)
	port := ":1235"
	host := "http://localhost"
	defer startServerOnPort(t, port, serv.getValues).Close()

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(host + port + "/values?target=test")
	req.Header.SetMethod("GET")
	req.Header.SetContentType("application/json")

	resp := fasthttp.AcquireResponse()

	assert.Nil(t, fasthttp.Do(req, resp))

	if resp.StatusCode() != 200 {
		t.Error("Error status code")
	}

	log.Println(string(resp.Body()))
}

func startServerOnPort(t *testing.T, port string, h fasthttp.RequestHandler) io.Closer {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost%s", port))
	if err != nil {
		t.Fatalf("cannot start tcp server on port %s: %s", port, err)
	}

	go func() {
		err = fasthttp.Serve(ln, h)
		if err != nil {
			log.Println(err)
		}
	}()

	return ln
}
