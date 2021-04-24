# piled

A client application to interpret the messages received from the morsecipher webserver
To compile piled for the Raspberry Pi
```bash
env GOOS=linux GOARCH=arm GOARM=5 go build piled.go
```

```bash
pi@raspberrypi ~> ./piled --help
```
```
Usage of ./piled:
  -P int
        Specify a broker's port. Default is 1883 (default 1883)
  -c string
        Specify a client_id for broker connection. Default is piled (default "piled")
  -h string
        Specify a broker's host. Default is morsecipher.xyz (default "morsecipher.xyz")
  -p string
        Specify a password for broker connection. Default is password (default "password")
  -t string
        Specify a custom topic to subscribe. Default is morse/msg (default "morse/msg")
  -u string
        Specify a username for broker connection. Default is morsecipher (default "morsecipher")
```