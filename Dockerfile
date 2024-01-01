FROM golang:alpine as builder 

# CREATE A FOLDER FOR OUR BUILD
RUN mkdir /build
ADD . /build/
WORKDIR /build


# GENERATE THE EXE BUILD FILE
RUN go mod tidy
RUN go build -o main .

# STAGE 2 
FROM alpine

# CREATE A USER
RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000

# RUN MAIN
CMD ["./main"]