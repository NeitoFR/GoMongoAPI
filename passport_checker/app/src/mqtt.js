const w = require('winston')

// w.add(new w.transports.Console({
//     format: w.format.simple()
// }));

module.exports = {
    getPassport: function (callback) {
        w.info('Getting MQTT passport')
        // TODO : Implement MQTT call
        callback('{"passport":"mqtt"}')
    }
}