const request = require("supertest");
const helper = require('../pkg/helper/helper')
const app = require("../cmd/index");
const { param } = require("../cmd/index");


describe('Register', () => {
    let testTable = [
        {
            name: "Register",
            params: {
                name: "Testing Bosss",
                phone: helper.generatePassword(4),
                role: "admin"
            },
            expected: 201
        },
    ]

    testTable.forEach(res => {
        test(res.name, async () => {
            await request(app)
            .post("/register")
            .send(res.params)
            .expect(res.expected)
        });    
    })
});

describe('Login', () => {
    let testTable = [
        {
            name: "Login",
            params: {
                password: helper.generatePassword(4),
                phone: "0203123213"
            },
            expected: 400
        },
    ]

    testTable.forEach(res => {
        test(res.name, async () => {
            await request(app)
            .post("/login")
            .send(res.params)
            .expect(res.expected)
        });    
    })
});

describe('Authorize', () => {
    let testTable = [
        {
            name: "Authorize",
            params: "sdfsdfsfdsfsdf",
            expected: 400
        },
    ]

    testTable.forEach(res => {
        test(res.name, async () => {
            await request(app)
            .get(`/authorize?token=${res.params}`)
            .send(res.params)
            .expect(res.expected)
        });    
    })
});