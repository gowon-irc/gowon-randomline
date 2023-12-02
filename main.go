package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gowon-irc/go-gowon"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Broker string `short:"b" long:"broker" env:"GOWON_BROKER" default:"localhost:1883" description:"mqtt broker"`
}

const (
	moduleName               = "randomline"
	mqttConnectRetryInternal = 5
	mqttDisconnectTimeout    = 1000
)

func defaultPublishHandler(c mqtt.Client, msg mqtt.Message) {
	log.Printf("unexpected message:  %s\n", msg)
}

func onConnectionLostHandler(c mqtt.Client, err error) {
	log.Println("connection to broker lost")
}

func onRecconnectingHandler(c mqtt.Client, opts *mqtt.ClientOptions) {
	log.Println("attempting to reconnect to broker")
}

func onConnectHandler(c mqtt.Client) {
	log.Println("connected to broker")
}

func main() {
	log.Printf("%s starting\n", moduleName)

	opts := Options{}
	if _, err := flags.Parse(&opts); err != nil {
		log.Fatal(err)
	}

	mqttOpts := mqtt.NewClientOptions()
	mqttOpts.AddBroker(fmt.Sprintf("tcp://%s", opts.Broker))
	mqttOpts.SetClientID(fmt.Sprintf("gowon_%s", moduleName))
	mqttOpts.SetConnectRetry(true)
	mqttOpts.SetConnectRetryInterval(mqttConnectRetryInternal * time.Second)
	mqttOpts.SetAutoReconnect(true)

	mqttOpts.DefaultPublishHandler = defaultPublishHandler
	mqttOpts.OnConnectionLost = onConnectionLostHandler
	mqttOpts.OnReconnecting = onRecconnectingHandler
	mqttOpts.OnConnect = onConnectHandler

	mr := gowon.NewMessageRouter()

	h := newBuilder().inputText("handling").build()
	mr.AddRegex("test", h.handle)

	mr.Subscribe(mqttOpts, moduleName)

	log.Print("connecting to broker")

	c := mqtt.NewClient(mqttOpts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	log.Print("connected to broker")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	log.Println("signal caught, exiting")
	c.Disconnect(mqttDisconnectTimeout)
	log.Println("shutdown complete")
}
