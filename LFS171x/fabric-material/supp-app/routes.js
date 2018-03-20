//SPDX-License-Identifier: Apache-2.0

var supp = require('./controller.js');

module.exports = function(app){

  app.get('/get_supp/:id', function(req, res){
    supp.get_supp(req, res);
  });
  app.get('/add_supp/:supp', function(req, res){
    supp.add_supp(req, res);
  });
  app.get('/get_all_supp', function(req, res){
    supp.get_all_supp(req, res);
  });
  app.get('/change_holder/:holder', function(req, res){
    supp.change_holder(req, res);
  });
}
