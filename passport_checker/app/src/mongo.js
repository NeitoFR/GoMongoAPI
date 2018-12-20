const w = require("winston"),
  http = require("http");

module.exports = {
  getPassport: function (name, callback) {
    w.info("Getting MongoDB passport for model : " + name);
    var path = "/passports?model_Name=" + name
    http
      .get(
        {
          host: process.env.mongo_api_host,
          port: process.env.mongo_api_port,
          path: "/passports"
        },
        res => {
          var body = "";
          res.on("data", chunk => {
            body += chunk;
          });
          res.on("end", () => {
            console.log('Req ended message is', body);

            try {
              var passport = JSON.parse(body);
              //TODO What if you get an array ?
              if (passport._id) delete passport._id;
            } catch (e) {
              w.info(
                "Cannot parse the response from : " +
                process.env.mongo_api_host +
                " | error  : " +
                e
              );
            }
            callback(passport);
            return;
          });
        }
      )
      .on("error", err => {
        w.info(
          "Error getting passport from : " +
          process.env.mongo_api_host +
          " : " +
          err
        );
      });
  }
};
