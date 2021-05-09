package main

import (
    "fmt"
    "flag"
    "github.com/stianeikeland/go-rpio"
    mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
    morseDirections string
    clientId string
    username string
    password string
    host string
    port int
    gpin int
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	pin := rpio.Pin(gpin)
	pin.Output()

    if string(msg.Payload()) == "1" {
        pin.High()
        fmt.Printf("1\n")
    } else {
        pin.Low()
        fmt.Printf("0\n")
    }
}


var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    opts := client.OptionsReader()
    fmt.Printf("Connected to %s\n", opts.Servers()[0].String())
    sub(client, morseDirections)
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
    fmt.Printf("Connect lost: %v\n", err)
}

func main() {
    opts := mqtt.NewClientOptions()

    flag.StringVar(&host, "h", "morsecipher.xyz", "Specify a broker's host. Default is morsecipher.xyz")
    flag.IntVar(&port, "p", 1883, "Specify a broker's port. Default is 1883")
    flag.StringVar(&clientId, "c", "piled", "Specify a client_id for broker connection. Default is piled")
    flag.StringVar(&username, "u", "morsecipher", "Specify a username for broker connection. Default is morsecipher")
    flag.StringVar(&password, "P", "password", "Specify a password for broker connection. Default is password")
    flag.StringVar(&morseDirections, "t", "morse/msg", "Specify a custom topic to subscribe. Default is morse/msg")
    flag.IntVar(&gpin, "gpin", 4, "Specify a custom PIN to receive the signals on. Default is 4")
    
    flag.Parse()

    opts.AddBroker(fmt.Sprintf("tcp://%s:%d", host, port))
    opts.SetClientID(clientId)
    opts.SetUsername(username)
    opts.SetPassword(password)
    opts.SetDefaultPublishHandler(messagePubHandler)
    opts.SetAutoReconnect(true)
    opts.SetOrderMatters(true)

    opts.OnConnect = connectHandler
    opts.OnConnectionLost = connectLostHandler
    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    for {}
}

func sub(client mqtt.Client, topic string) {
    token := client.Subscribe(topic, 2, nil)
    token.Wait()
    fmt.Printf("Subscribed to topic: %s\n", topic)
}
