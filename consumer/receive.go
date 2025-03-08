package main

import (
    "log"
    "github.com/go-resty/resty/v2"
    amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Panicf("%s: %s", msg, err)
    }
}

func main() {
    client := resty.New()

    conn, err := amqp.Dial("amqp://keren:keren12345@44.216.244.37:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    err = ch.ExchangeDeclare(
        "broker1",   // name
        "fanout", // type
        true,     // durable
        false,    // auto-deleted
        false,    // internal
        false,    // no-wait
        nil,      // arguments
    )
    failOnError(err, "Failed to declare an exchange")

    q, err := ch.QueueDeclare(
        "citas",    // name
        false, // durable
        false, // delete when unused
        false,  // exclusive
        false, // no-wait
        nil,   // arguments
    )
    failOnError(err, "Failed to declare a queue")

    err = ch.QueueBind(
        q.Name, // queue name
        "",     // routing key
        "broker1", // exchange
        false,
        nil,
    )
    failOnError(err, "Failed to bind a queue")

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    failOnError(err, "Failed to register a consumer")

    var forever chan struct{}

    go func() {
        for d := range msgs {
            log.Printf(" [x] Received message: %s", d.Body)

            // Aquí se imprime la solicitud POST antes de enviarla
            log.Printf("Enviando request POST con los siguientes detalles:")
            log.Printf("URL: http://localhost:8082/citas/")
            log.Printf("Headers: %v", map[string]string{"Content-Type": "text/plain"})
            log.Printf("Body: %s", d.Body)

            for i := 0; i < 3; i++ {
                resp, err := client.R().
                    SetHeader("Content-Type", "text/plain").
                    SetBody(d.Body).
                    Post("http://localhost:8082/citas/")

                if err == nil && resp.StatusCode() < 500 {
                    log.Printf("Request exitoso. Código de estado: %d", resp.StatusCode())
                    break
                }

                log.Printf("Intento %d falló, reintentando...", i+1)
            }
        }
    }()

    log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
    <-forever
}