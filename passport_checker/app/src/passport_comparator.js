w = require("winston"),
  _ = require('lodash');

// w.add(new w.transports.Console({
//     format: w.format.simple()
// }));

module.exports = {
  comparePassports: function (model_Name, pMqtt, pMongo, callback) {
    w.info("********** Starting comparaison for model : " + model_Name + " **********")
    var full_passport = JSON.parse(JSON.stringify(pMongo))
    //TODO Implement passport concatenation and comparaison
    pMongo.room_List.forEach((room, i)  => {
      room.esp_List.forEach((esp, y) => {
        w.info("Looking for MAC : "+ esp.mac_address+" name : "+ esp.esp_Hostname)
        var esp_temp = _.find(pMqtt, item => {
          return (item.mac_address === esp.mac_address) && (item.esp_Hostname === esp.esp_Hostname)
        }) 
        if(esp_temp) full_passport.room_List[i].esp_List[y] = esp_temp
        else full_passport.room_List[i].esp_List[y]["not_found"] = true
        // else full_passport.room_List[i].esp_List[y] = {not_Found : true}
      })
    });
    
    
    var final = { "MongoDB_Passport": pMongo, "MQTT_Passport": pMqtt || "none yet", "full_Passport": full_passport}
    callback(final);
  }
};
