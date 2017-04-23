FROM golang
ADD . /go/src/ridedonor
WORKDIR /go/src/ridedonor

RUN go get gopkg.in/gin-gonic/gin.v1
RUN go get -u github.com/jinzhu/gorm
RUN go get github.com/lib/pq
RUN go install ridedonor

ENTRYPOINT /go/bin/ridedonor
EXPOSE 8080
