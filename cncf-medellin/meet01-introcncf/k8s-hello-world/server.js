var express = require('express');
var app = express();


const {MY_ENV1, MY_ENV2} = process.env

app.get('/', function (req, res) {
  res.send(`Hello Medellin Meetup, ${MY_ENV2}`);
});


app.get('/shutdown', function (req, res) {
    
    res.send(`Server Shutted Down', ${MY_ENV1}`);
    process.exit(1);
    
  });

app.listen(8080, function () {

  console.log('This is an example app on NodeJs ');

});
