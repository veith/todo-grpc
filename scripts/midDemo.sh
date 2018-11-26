#! /bin/bash

echo "generating files from mid"
echo "generating sqlite schema"
simple-generator -d api/mid/mid.json -t api/mid/schema.sql.tmpl > api/mid_generates/schema.sql
echo "generating java compatible messages proto (without gogoproto.moretags options)"
simple-generator -d api/mid/mid.json -t api/mid/messages.proto.java.tmpl > api/mid_generates/messages.java.proto
echo "generating go compatible messages proto"
simple-generator -d api/mid/mid.json -t api/mid/messages.proto.go.upper.tmpl > api/mid_generates/messages.proto
echo "generating services proto for transcoder and api"
simple-generator -d api/mid/mid.json -t api/mid/services.proto.tmpl > api/mid_generates/services.proto
echo "done"