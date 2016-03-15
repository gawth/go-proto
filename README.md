# go-proto
Noddy app that consumes and benchmarks some hard coded endpoints for JSON and Protocol Buffers (and message pack)

This is part of an exercise to trial performance of alternatives to JSON over the wire.  The predominant success criteria is speed.

There are a number of Github repositories:

[Golang benchmark tests and primary write up of results](https://github.com/gawth/go-proto)

[Node.js server implementation](https://github.com/gawth/protosrv)

[.Net server implementation](https://github.com/gawth/proto-service)

[Protocol Buffer schema](https://github.com/gawth/hotel-proto)

[Node.js consumer plus benchmarking](https://github.com/gawth/node_proto)

##Setup
This app's job in life is to hit an end point using a number of different protocols, deserialize the data and then report metrics.

Its all a bit hard coded.  The target servers were running on my local VM with an IP address of 10.211.55.3.

To get up and running you will need to have a server running the app from [.Net server implementation](https://github.com/gawth/proto-service) and then you will need to change the hard coded IP (probably).

And a node.js server is also needed from [Node.js server implementation](https://github.com/gawth/protosrv)

Once you've got your target servers up and running you will need to:

1. go get
2. cd hotel-proto; git submodule init; git submodule update; cd ..
3. go test bench=TLRG

That last command runs all benchmark tests whose title includes "TLRG".  Without that you will get a bunch of other tests running - no harm but takes a while.

##What's Going On

Go provides benchmark functionality out of the box.  It takes care of repeatedly running the tests and collating the results.

    BenchmarkTLRGJsonEndToEnd-4          200          12097761 ns/op    Full end to end against JSON endpoint
    BenchmarkTLRGProtoBufEndToEnd-4     1000           2029807 ns/op    Full end to end against protocol buffers end point
    BenchmarkTLRGJsonEndpoint-4          200           6375379 ns/op    JSON Endpoint 
    BenchmarkTLRGPBEndpoint-4           1000           1680085 ns/op    Protcol Buffer Endpoint
    BenchmarkTLRGJsonProcessing-4       1000           2187234 ns/op    Golang processing of JSON
    BenchmarkTLRGMPProcessing-4         5000            279339 ns/op    Golang processing of Message Pack
    BenchmarkTLRGPBProcessing-4         5000            365211 ns/op    Golang processing of Protocol Buffers

    BenchmarkTLRG2JsonEndpoint-4        1000           2086541 ns/op    JSON node.js end point 
    BenchmarkTLRG2PBEndpoint-4          2000            603525 ns/op    Protocol Buffers node.js end point

    BenchmarkTLRGMsgpackEndToEnd-4       500           3723422 ns/op    Message pack end to end before I moved to different .Net implementation (see below)
    BenchmarkTLRGMPEndpoint-4            100          10303550 ns/op    Message pack end point only
    BenchmarkTLRGMPFastEndpoint-4       1000           1452170 ns/op    Message pack using arrays

###Message Pack Notes
At the start of the eval I included evaluation of message pack.  I've since let the message pack support lapse a little.

In summary, message pack operates in two modes.  The default is to pack the message in to an array. Alternatively you can use a map.  

The array is the most space efficient but does rely on ordering being consistent between server and client.
It also means that any additions to the message must be added to the end.

Maps are a little more flexible in that ordering doesn't matter and additions are easier but the drawback of using a map is that it increases message size and is slower.

Given than I am looking at this to operate across teams using different technologies expecting everyone to be consistent in their message format without the help of a schema felt ambitious so I pretty much abandoned message pack in this context.

###Conclusion

* For .Net server we're looking at Protocol Buffers being 5x faster than JSON.
* For Node.js we're looking at Protocol Buffers being 2x faster than JSON.
* And for Golang its a 5x speed up when handling message pack 

##TODO

Try out 0mq as an alternative to HTTP

