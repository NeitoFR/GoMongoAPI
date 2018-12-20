const w = require('winston'),
http = require('http')

// w.add(new w.transports.Console({
//     format: w.format.simple()
// }));

module.exports = {
  getPassport: function (name, callback) {
    w.info('Getting MQTT passport for model : '+ name)
    var path = "/passports?model_Name=" + name
    http
      .get(
        {
          host: process.env.mqtt_api_host,
          port: process.env.mqtt_api_port,
          path: path
        },
        res => {
          var body = "";
          res.on("data", chunk => {
            body += chunk;
          });
          res.on("end", () => {
            try {
              var passport = JSON.parse(body);
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
        callback('{"error":"' + err + '"}')
      });
  }
}