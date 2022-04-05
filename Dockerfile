## We specify the base image we need for our
## go application
FROM golang:1.17-alpine

ENV REPO_URL=github.com/amitabhprasad/bookstore-app/bookstore-oauth-api/src

ENV GOPATH=/app

ENV APP_PATH=${GOPATH}/src/${REPO_URL}

ENV WORKPATH=$APP_PATH/src

COPY src ${WORKPATH}

WORKDIR $WORKPATH
RUN go build -o oauth-api .

# expose port 8082
EXPOSE 8082

CMD ["./oauth-api"]