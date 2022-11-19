const bodyParser = require('body-parser')
const express = require('express')
const app = express()
app.use(bodyParser.json())
app.use(bodyParser.urlencoded({ extended: true }))

const authService = require("../services/auth_service")
app.post('/register', authService.Register)
app.post('/login', authService.Login)
app.get('/authorize', authService.Authorize)

module.exports = app