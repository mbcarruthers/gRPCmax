### gRPCmax

gRPC bidirectional stream client/server.
client streams a set of random int64 values to the server
and server responds with the current maximum value in the set.

to run:<br/>
<code> make clean_protoc </code><br/>
<code> make build_protoc </code><br/>
<code> make build_server </code><br/>
<code> make build_client </code><br/>

<code>./bin/max/server</code><br/>
<code>./bin/max/client</code><br/>

