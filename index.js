const express = require('express')
const bodyParser = require('body-parser')
const db = require('./queries')
const app = express()

app.use(bodyParser.json())
app.use(bodyParser.urlencoded({ extended: true }))

app.get('/', (req, res) => {
  res.json({ info: 'test database with node and postgreSQL' })
})

app.get('/user', db.getUsers)
app.get('/user/:id', db.getUserById)
app.post('/users', db.createUser)
app.put('/users/:id', db.updateUser)
app.delete('/users/:id', db.deleteUser)


app.listen(3000, () => {
  console.log('App running on port 3000')
})
  
