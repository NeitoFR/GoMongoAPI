const express = require('express');
const app = express();

app.get('/hello', function (req, res) {
    res.send("Hello from passport_checker").end()
})

app.listen(9002, function () {
  console.log('Example app listening on port 9002!');
})