package main

import (
	"reflect"
	"testing"
	"time"
)

var googleHttp = Host{Network: "tcp", Address: "www.google.com:80"}
var googleTelnet = Host{Network: "tcp", Address: "www.google.com:23"}
var bogusHost = Host{Network: "tcp", Address: "zebeleo.com.br:80"}
var timeout = time.Second * 3

func TestCheckOpenHostAndPortReturnsOpen(t *testing.T) {
	expected := Check{state: open, host: googleHttp}
	result := check(&googleHttp, &timeout)
	if result != expected {
		t.Fatalf("got %v, expected %v", result, expected)
	}
}

func TestCheckClosedPortReturnsClosed(t *testing.T) {
	expected := Check{state: closed, host: googleTelnet}
	result := check(&googleTelnet, &timeout)
	if result != expected {
		t.Fatalf("got %v, expected %v", result, expected)
	}
}

func TestCheckBogusHostReturnsFailure(t *testing.T) {
	expected := Check{state: failure, host: bogusHost}
	result := check(&bogusHost, &timeout)
	if result != expected {
		t.Fatalf("got %v, expected %v", result, expected)
	}
}

func TestRunScannerIteratesConfiguredData(t *testing.T) {
	config := Config{
		Interval: time.Minute,
		Timeout:  time.Second,
		Hosts: []Host{
			googleHttp,
			googleTelnet,
			bogusHost,
		},
	}

	expected := []Check{
		{state: open, host: googleHttp},
		{state: closed, host: googleTelnet},
		{state: failure, host: bogusHost},
	}

	result, err := scan(&config)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
