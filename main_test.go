package main

import (
	"testing"
)

var host = "http://10.211.55.3:8888/"

func BenchmarkTLRGJsonEndToEnd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processJson(hitUrl("application/json", host+"testjson"))
	}
}
func BenchmarkTLRGMsgpackEndToEnd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processMsgPack(hitUrl("application/x-msgpack", host+"testmsgpack"))

	}
}
func BenchmarkTLRGProtoBufEndToEnd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processProto(hitUrl("application/x-protobuf", host+"testprotobuf"))
	}
}

func BenchmarkTLRGJsonEndpoint(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hitUrl("application/json", host+"testjson")
	}
}
func BenchmarkTLRGMPEndpoint(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hitUrl("application/x-msgpack", host+"testmsgpack")
	}
}
func BenchmarkTLRGMPFastEndpoint(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hitUrl("application/x-msgpack", host+"testmsgpack2")
	}
}
func BenchmarkTLRGPBEndpoint(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hitUrl("application/x-protbuf", host+"testprotobuf")
	}
}

func BenchmarkTLRGJsonProcessing(b *testing.B) {
	data := hitUrl("application/json", host+"testjson")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		processJson(data)
	}
}
func BenchmarkTLRGMPProcessing(b *testing.B) {
	data := hitUrl("application/x-msgpack", host+"testmsgpack")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		processMsgPack(data)
	}
}
func BenchmarkTLRGPBProcessing(b *testing.B) {
	data := hitUrl("application/x-protobuf", host+"testprotobuf")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		processProto(data)
	}
}
