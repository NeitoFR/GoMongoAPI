const w = require('winston')

// w.add(new w.transports.Console({
//     format: w.format.simple()
// }));

module.exports = {
    getPassport: function (callback) {
        w.info("Getting mongodb passport")
        // TODO : Implement MongoDB call
        callback('{"passport":"mongo"}')
    }
}