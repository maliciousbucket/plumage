FROM golang:1.23 AS build
LABEL authors="maliciousbucket"

WORKDIR /app

COPY go.mod go.sum  ./

RUN go mod download

COPY *.go ./
COPY imports/ ./imports/

RUN --mount=type=bind,target=. go build -o /plumage .

RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs \
    && mkdir -p /node-install \
    && cp -r /usr/bin/node /usr/bin/npm /usr/lib/node_modules /node-install/

FROM scratch

COPY --from=build /plumage /plumage

COPY --from=build /node-install/ /usr/local/


ENTRYPOINT ["/plumage"]