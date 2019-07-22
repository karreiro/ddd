protoc -I proto/ proto/game.proto --go_out=plugins=grpc:proto
mvn exec:java -Dexec.mainClass=org.kie.grpc.KieGrpcServer
