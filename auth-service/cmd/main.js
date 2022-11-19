const conf = require("../pkg/config/config.json")
const port = conf.server.port
const app = require("./index");

app.listen(port, () => {
    console.log(`Auth service listening on port ${port}`);
})