FROM golang:1.17.12-buster

RUN mkdir /build
WORKDIR /build

RUN apt-get update && apt-get -y install --no-install-recommends protobuf-compiler-grpc python3-grpc-tools libprotobuf-dev protobuf-compiler python3-pip python3-dev python3-setuptools  && rm -rf /var/lib/apt/lists/*

RUN go mod init proto && \
    go install github.com/envoyproxy/protoc-gen-validate@v0.3.0-java && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install github.com/srikrsna/protoc-gen-gotag@v0.5.0 && \
    go get google.golang.org/grpc@v1.28.0 && \
    go get github.com/google/googleapis@v0.0.0-20200324113624-36c0febd0fa7 && \
    go get go.unistack.org/micro/v3@v3.10.10 && \
    go get github.com/presnalex/protoc-gen-micro@v0.0.2 && \
    go get github.com/go-bindata/go-bindata/...

CMD rm -rf python && mkdir python && rm -rf go && mkdir go && \
        python3 -m grpc_tools.protoc \
          --validate_out=lang=go:go \
          --micro_out=go \
          --go_out=go \
          --python_out=python \
          --grpc_python_out=python \
          -I=. \
          -I=/usr/include \
          -I=/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.3.0-java \
          -I=/go/pkg/mod/github.com/srikrsna/protoc-gen-gotag@v0.5.0 \
          -I=/go/pkg/mod/github.com/google/googleapis@v0.0.0-20200324113624-36c0febd0fa7 \
          */*.proto && \
        cd go && \
        protoc \
        -I=../ \
        -I=/usr/include \
        -I=/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.3.0-java \
        -I=/go/pkg/mod/github.com/srikrsna/protoc-gen-gotag@v0.5.0 \
        -I=/go/pkg/mod/github.com/google/googleapis@v0.0.0-20200324113624-36c0febd0fa7 \
        --gotag_out=xxx="xml+\"-\"":. \
        ../*/*.proto  && \
        cd .. 