package routing_test;

import (
  "net/http"
  "net/http/httptest"
  "testing"
  
  "github.com/hrafn/gophercon2018/pkg/routing"
  
)

func TestBaseRouter(t *testing.T) {
  handler = routing.homeHandler()
  
  ts := httptest.NewServer(handler)
  defer ts.Close()
  
  res, err = http.Get(ts.URL + "/home")
  if err != nil {
    t.Fatal(err)
  }
  
  if res.StatusCode != http.StatusOK {
    t.Errorf("Wrong status code %d (%d expected)", res.StatusCode, http.StatusOK)
  }
}