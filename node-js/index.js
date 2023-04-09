const express = require('express')
const cors = require('cors')
const morgan = require('morgan')
const ch = require('./rabbitmq')
const queue = 'q.go.service'
const app = express()

app.use(cors())
app.use(morgan('dev'))
app.use(express.json())
app.use(express.urlencoded({ extended: true }))

async function main() {
  const ch1 = await ch.RabbitMQ(queue, 'ex.sing', 'fanout')
  // Listener
    ch1.consume(queue, (msg) => {
      if (msg !== null) {
        console.log('Received:', msg.content.toString())
        ch1.ack(msg)
      } else {
        console.log('Consumer cancelled by server')
      }
    })

  app.get('/', (req, res) => {
    res.json({ message: 'Hello World!' })
  })

  app.post('/login', (req, res) => {
    ch1.sendToQueue(queue, Buffer.from(JSON.stringify(req.body)))
    res.json(req.body)
  })
}

app.listen(3000, () => {
  main()
  console.log('Server started on port 3000')
})
