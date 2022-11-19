const authDatabase = require("../pkg/database/auth_database")
const helper = require('../pkg/helper/helper')
const bcrypt = require('bcrypt')
const jwt = require('jsonwebtoken')

//Schema Validation
const Ajv = require("ajv")
const ajv = new Ajv()

async function Register(req,res) {
    // Request body [phone,role,name]
    let requestBody = req.body

    //Preparing validation request body
    const schema = {
        type: "object",
        properties: {
          phone: {type: "string", minLength: 1},
          role: {type: "string", minLength: 1},
          name: {type: "string", minLength: 1}
        },
        required: ["phone","role","name"],
        additionalProperties: false
    }
    const validate = ajv.compile(schema)

    const valid = validate(requestBody)
    if (!valid) {
        res.status(400).json({
            message: "Register is failed",
            errors: validate.errors
        })
        return
    }

    requestBody["timestamp"] = String(Date.now())

    let userData = await authDatabase.SelectByPhone(requestBody.phone)
    if (userData) {
        res.status(400).json({
            message: "Sorry phone number has been registered"
        })
        return
    }

    let userPassword = helper.generatePassword(4)
    let userPasswordHashed = await helper.hashingPassword(userPassword)
    requestBody["password"] = userPasswordHashed

    let error = await authDatabase.InsertData(requestBody)
    if (error) {
        res.status(400).json({
            message: "Register is failed"
        })
        return
    }

    res.status(201).json({
        message: "Register is successfully",
        password: userPassword
    })
}

function Authorize(req,res) {
    const authHeader = req.headers.authorization
    let tokenHedaer = ""
    if (authHeader) {
        tokenHedaer = authHeader.split(' ')[1]
    }
    const token = req.query.token ?? tokenHedaer

    //Decode token and get private claims
    let decoded
    try {
        decoded = jwt.verify(token, 'shhhhh')
    } catch(err) {
        res.status(400).json({
            message: "JWT token invalid"
        })
        return
    }
 
    res.status(200).json({
        message: "JWT token Is valid",
        data: decoded
    })
}


async function Login(req,res) {
    let requestBody = req.body

    //Preparing validation request body
    const schema = {
        type: "object",
        properties: {
            phone: {type: "string", minLength: 1},
            password: {type: "string", minLength: 1}
        },
        required: ["phone","password"],
        additionalProperties: false
    }
    const validate = ajv.compile(schema)

    const valid = validate(requestBody)
    if (!valid) {
        res.status(400).json({
            message: "Register is failed",
            errors: validate.errors
        })
        return
    }
    
    let userData = await authDatabase.SelectByPhone(requestBody.phone)
    if (!userData) {
        res.status(400).json({
            message: "Phone number is not found"
        })
        return
    }

    //Checking password is match
    const match = await bcrypt.compare(requestBody.password, userData.password)
    if (!match) {
        res.status(400).json({
            message: "Sorry your password do not match"
        })
        return
    }

    //Generate jwt token
    delete userData.password
    let claimsData = {...userData, exp: Math.floor(Date.now() / 1000) + (60 * 60)}
    let token = jwt.sign(claimsData,'shhhhh')

    res.status(200).json({
        message: "Authentication success",
        jwt_token: token
    })
}

module.exports = { Register, Authorize, Login }