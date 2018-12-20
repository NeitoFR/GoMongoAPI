const w = require("winston"),
  http = require("http");

module.exports = {
  getPassport: function(callback) {
    w.info("Getting mongodb passport");
    w.info(process.env);
    // TODO : Implement MongoDB call
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
            try {
              var passport = JSON.parse(body);
              //TODO What if you get an array ?
              delete passport._id;
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
