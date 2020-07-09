const MongoClient = require('mongodb').MongoClient;
const assert = require('assert');

// Connection URL
const url = 'mongodb+srv://veronica:<PASSWORD>@pluto.e3h8i.mongodb.net/test?retryWrites=true&w=majority';

// Database Name
const dbName = 'test';

// Create a new MongoClient
const client = new MongoClient(url);

// Use connect method to connect to the Server
client.connect(function(err) {
  assert.equal(null, err);
  console.log("Connected successfully to server");

  const db = client.db(dbName);
  var runQuery = function(coll, query, iterations=1) {
      var start = Date.now();

      for (i = 0; i < iterations; i++) {
          res = db.collection(coll).find(query).toArray();
          }
      var end = Date.now();
      return end-start;
     };

   var a = runQuery("test.data", {"ident":"state-gcp-noregion-develop"},100);
   console.log(a);;
  client.close();
});
