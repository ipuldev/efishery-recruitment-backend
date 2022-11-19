var helper = require("./helper")

describe('Generate Password', () => {
    let testTable = [
        {
            name: "Generate",
            param: 4,
            validate: (password) => {
                expect(password.length).toBe(4)
            }
        },
        {
            name: "Generate - Param Null",
            param: null,
            validate: (password) => {
                expect(password.length).toBe(0)
            }
        },
        {
            name: "Generate - Param 0",
            param: 0,
            validate: (password) => {
                expect(password.length).toBe(0)
            }
        }
    ]

    testTable.forEach(res => {
        test(res.name, () => {
            password = helper.generatePassword(res.param)
            res.validate(password)
        });    
    })
});

describe('Hashing Password', () => {
    let testTable = [
        {
            name: "Hash",
            param: "PASSWORD",
            validate: (hashPassword) => {
                expect(hashPassword.length).toBeGreaterThan(0)
            }
        },
        {
            name: "Hash - Param empty",
            param: "",
            validate: (hashPassword) => {
                expect(hashPassword.length).toBe(0)
            }
        },
        {
            name: "Hash - Param 0",
            param: 0,
            validate: (hashPassword) => {
                expect(hashPassword.length).toBe(0)
            }
        },
        {
            name: "Hash - Param Number",
            param: 1231,
            validate: (hashPassword) => {
                expect(hashPassword.length).toBe(0)
            }
        },
    ]

    testTable.forEach(res => {
        test(res.name, async () => {
            hashPassword = await helper.hashingPassword(res.param)
            res.validate(hashPassword)
        });    
    })
});