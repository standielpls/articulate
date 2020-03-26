from flask import Flask, escape, request, jsonify
from pymongo import MongoClient
import uuid
import pprint

client = MongoClient()
db = client["notes"]
col = db["notes"]
pp = pprint.PrettyPrinter(indent=4)

app = Flask(__name__)

@app.route('/')
def hello():
    name = request.args.get("name", "World")
    return f'Hello, {escape(name)}!'

@app.route('/notes', methods=['GET', 'POST'])
def postNotes():
    if request.method == 'GET':
        data = []
        notes = col.find().sort("time", -1)
        for n in notes:
            n['_id'] = ""
            data.append(n)
        pp.pprint(data)
        resp = jsonify(data)
        resp.headers.add("Access-Control-Allow-Origin", "*")
        return resp

    resp = request.get_json()
    if resp['url'] == None: 
        return jsonify({"error": {"message": "url must be provided"}})
    if resp['quote'] == None: 
        return jsonify({"error": {"message": "quote must be provided"}})
    print(resp)
    resp['id'] = uuid.uuid4()
    try: 
        id = col.insert_one(resp).inserted_id
        print(id)
    except Exception as e:
        print("unable to insert:{0}".format(e))
        return jsonify({"error": {"message": "unable to insert: {0}".format(e)}})
    return jsonify(success=True)

