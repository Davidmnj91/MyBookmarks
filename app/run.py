from flask import Flask, jsonify
from pymongo.mongo_client import MongoClient

from app.bookmarks import bp

app = Flask(__name__)

app.register_blueprint(bp)

@app.route('/')
def index():
    return jsonify("Hello from App!!")

if __name__ == "__main__":
    app.run(debug=True)
