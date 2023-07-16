FROM golang:1.20 as build
ENV GO111MODULE=on
WORKDIR /app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -o /app/license /app/api/main.go

# Unprivileged users can execute
RUN chgrp 0 /app/license
RUN chmod g+x /app/license

FROM scratch
COPY --from=build /app/license .
USER 65534
CMD ["/license"]