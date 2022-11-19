const bodyParser = require('body-parser')
const express = require('express')
const app = express()
app.use(bodyParser.json())
app.use(bodyParser.urlencoded({ extended: true }))

const authService = require("../services/auth_service")
app.post('/auth/register', authService.Register)
app.post('/auth/login', authService.Login)
app.get('/auth/authorize', authService.Authorize)

module.exports = app