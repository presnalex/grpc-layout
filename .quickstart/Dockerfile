FROM docker:git
RUN git clone https://github.com/presnalex/grpc-layout.git
RUN apk add curl
RUN apk add --no-cache make
COPY --from=golang:alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
WORKDIR /grpc-layout/.quickstart