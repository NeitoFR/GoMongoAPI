w = require('winston')

// w.add(new w.transports.Console({
//     format: w.format.simple()
// }));

module.exports = {
    comparePassports: function (pMqtt, pMongo,callback) {
        w.info(pMongo+ pMqtt)
        callback(JSON.stringify(pMongo) + JSON.stringify(pMqtt))
    }
}