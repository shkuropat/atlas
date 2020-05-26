# How to start

**Atlas** boilerplate project consists of the following executable components/entry points:

1. Service. Serves gRPC calls from/to client and consumer(s).
   - Entry point: [cmd/service](../cmd/service)
   - Config file: [config/service.yaml](../config/service.yaml)
   The following commands are supported out of the box:
   - [Main entry point] Serve requests [serve](../cmd/service/cmd/serve.go)
   - [Supplementary] Display parsed config [config](../cmd/service/cmd/config.go)
   - [Supplementary] Display software version [version](../cmd/service/cmd/version.go)
1. Client. Makes gRPC calls to service
   - Entry point: [cmd/client](../cmd/client)
   - Config file: [config/client.yaml](../config/client.yaml)
   The following commands are supported out of the box:
   - [Main entry point] Send file or STDIN from client to service [send](../cmd/client/cmd/send.go)
   - [User management] Register OAuth client on OAuth server [register](../cmd/client/cmd/register.go)
   - [Supplementary] Display parsed config [config](../cmd/client/cmd/config.go)
   - [Supplementary] Display software version [version](../cmd/client/cmd/version.go)
1. Consumer. Consumes data replayed by the service
   - Entry point: [cmd/consumer](../cmd/consumer)
   - Config file: [config/consumer.yaml](../config/consumer.yaml)
   The following commands are supported out of the box:
   - [Main entry point] Consume data accumulated by the service [consume](../cmd/consumer/cmd/consume.go)
   - [Supplementary] Display parsed config [config](../cmd/consumer/cmd/config.go)
   - [Supplementary] Display software version [version](../cmd/consumer/cmd/version.go)

Client and Service exchange data in `protobuf` format. Protobuf specs are located in [pkg/api/atlas](../pkg/api/atlas) folder.
Main entry point is [service_control_plane.proto](../pkg/api/atlas/service_control_plane.proto) file which defines `ControlPlane` data exchange.
 
**Atlas** consists of the following subsystems:
1. Auth  [pkg/auth](../pkg/auth)
1. Controller - is an entry point for gRPC request/response parties. [pkg/contoller](../pkg/controller)
1. Kafka - interface component with Apache Kafka MOM [pkg/kafka](../pkg/kafka)
1. MinIO - interface component with MinIO object storage [pkg/minio](../pkg/minio)
1. Transport - grpc/TLS machinery [pkg/transport](../pkg/transport)

