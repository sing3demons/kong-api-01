const amqplib = require('amqplib')

async function RabbitMQ(queue, exchange, type) {
  //   const queue = 'q.sing.service'

  const conn = await amqplib.connect('amqp://rabbitmq:1jj395qu@localhost:5672')
  const ch1 = await conn.createChannel()
  await ch1.assertQueue(queue)
  await ch1.assertExchange(exchange, type, { durable: true })

  ch1.bindQueue(queue, exchange, '')

  return ch1
}

module.exports = { RabbitMQ }
