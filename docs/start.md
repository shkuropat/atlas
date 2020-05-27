# How to start

## Entry points / executables

**Atlas** boilerplate project consists of the following executable components/entry points:

1. Service. Serves gRPC calls from/to client(s) and consumer(s).
   - Entry point: [cmd/service](../cmd/service)
   - Config file: [config/service.yaml](../config/service.yaml)

   The following commands are supported out of the box:
   - [Main entry point] Serve requests command: [`serve`](../cmd/service/cmd/serve.go)
   - [Supplementary] Display parsed config command: [`config`](../cmd/service/cmd/config.go)
   - [Supplementary] Display software version command: [`version`](../cmd/service/cmd/version.go)
1. Client. Makes gRPC calls to service
   - Entry point: [cmd/client](../cmd/client)
   - Config file: [config/client.yaml](../config/client.yaml)

   The following commands are supported out of the box:
   - [Main entry point] Send file or STDIN from client to service command: [`send`](../cmd/client/cmd/send.go)
   - [User management] Register OAuth client on OAuth server command: [`register`](../cmd/client/cmd/register.go)
   - [Supplementary] Display parsed config command: [`config`](../cmd/client/cmd/config.go)
   - [Supplementary] Display software version command: [`version`](../cmd/client/cmd/version.go)
1. Consumer. Consumes data replayed by the service
   - Entry point: [cmd/consumer](../cmd/consumer)
   - Config file: [config/consumer.yaml](../config/consumer.yaml)

   The following commands are supported out of the box:
   - [Main entry point] Consume data accumulated by the service command: [`consume`](../cmd/consumer/cmd/consume.go)
   - [Supplementary] Display parsed config command: [`config`](../cmd/consumer/cmd/config.go)
   - [Supplementary] Display software version command: [`version`](../cmd/consumer/cmd/version.go)

## Data exchange format and specification

Client and Service exchange data in `protobuf` format. Protobuf specs are located in [pkg/api/atlas](../pkg/api/atlas) folder.
Main entry point is [service_control_plane.proto](../pkg/api/atlas/service_control_plane.proto) file which defines `ControlPlane` data exchange.
 
## Main subsystems
 
**Atlas** consists of the following subsystems:
1. **Auth**. Based on OAuth2. Client/Server auth with DCRP support.  [pkg/auth](../pkg/auth)
1. **Config**. Support for easily extensible files, ENV vars and CLI options config. [pkg/config](../pkg/config)
1. **Controller**. Entry point for gRPC request/response parties. Sort of **Client/Server API**. [pkg/controller](../pkg/controller)
1. **Kafka**. Interface component with Apache Kafka MOM. [pkg/kafka](../pkg/kafka)
1. **MinIO**. Interface component with MinIO object storage. [pkg/minio](../pkg/minio)
1. **SoftwareID**. Software identification. [pkg/softewareid](../pkg/softwareid)
1. **Transport**. gRPC/TLS machinery. [pkg/transport](../pkg/transport)
