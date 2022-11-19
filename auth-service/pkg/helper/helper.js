//Password encryption library
const bcrypt = require('bcrypt')

function generatePassword(length) {
    if (!length) return ""
    var result           = ''
    var characters       = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
    var charactersLength = characters.length
    for ( var i = 0; i < length; i++ ) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result
}

function hashingPassword(password) {
    return new Promise((resolve, reject) => {
        if (!password || typeof password == 'number') resolve("")
        //Hashing password using brcypt library
        bcrypt.genSalt(10, function(err, salt) {
            bcrypt.hash(password, salt, function(err, hash) {
                resolve(hash)
            });
        });
    })
}

module.exports = { generatePassword, hashingPassword}