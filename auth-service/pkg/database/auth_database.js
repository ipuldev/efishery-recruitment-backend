var conf = require("../config/config.json")
const sqlite3 = require('sqlite3').verbose()

let db = new sqlite3.Database(`${__dirname}/../../db/${conf.db.name}`,sqlite3.OPEN_READWRITE,(err) => {
    if (err) {
        console.log(err.message)
    }
})

async function InsertData(userData) {
    return new Promise((resolve, reject) => {
        db.run(`INSERT INTO auth(name,phone,role,password,timestamp) VALUES(?,?,?,?,?)`,[userData.name,userData.phone,userData.role,userData.password,userData.timestamp],function(err) {
            if (err) {
                resolve(err);
                return console.log(err.message)
            }
            resolve(null)
            return null
        })
    })
}

async function SelectByPhone(phone) {
    return new Promise((resolve, reject) => {
        db.get(`SELECT * FROM auth WHERE phone = ?`,[phone],function(err,row) {
            if (err) {
                resolve(null)
                return console.log(err.message)
            }
            resolve(row)
        })
    })
}
module.exports = { InsertData, SelectByPhone}