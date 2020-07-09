import pprint
import time
import pymongo
from pymongo import MongoClient
from pymongo.read_preferences import Secondary
from pymongo import ReadPreference
client = pymongo.MongoClient("mongodb+srv://veronica:<PASSWORD>@pluto.e3h8i.mongodb.net/sample_mflix?retryWrites=true&w=majority")

col = client.test.data
start_time = time.time()
pprint.pprint(col.find_one({"ident":"state-gcp-noregion-develop"}))
print("--- %s seconds ---" % (time.time() - start_time))



client_us = pymongo.MongoClient('mongodb://veronica:<PASSWORD>@pluto-shard-00-01.e3h8i.mongodb.net:27017/sample_mflix?ssl=true&authSource=admin', read_preference=ReadPreference.SECONDARY)
db= client_us.test
col = db.data
start_time = time.time()
pprint.pprint(col.find_one({"ident":"state-gcp-noregion-develop"}))
print("US Secondary read --- %s seconds ---" % (time.time() - start_time))



client_sea = pymongo.MongoClient('mongodb://veronica:<PASSWORD>@pluto-shard-00-02.e3h8i.mongodb.net:27017/sample_mflix?ssl=true&authSource=admin', read_preference=ReadPreference.SECONDARY)
db = client_sea.test
col = db.data
start_time = time.time()
pprint.pprint(col.find_one({"ident":"state-gcp-noregion-develop"}))
print("SEA Secondary read --- %s seconds ---" % (time.time() - start_time))

