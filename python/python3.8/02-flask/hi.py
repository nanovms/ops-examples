from flask import Flask
import requests

app = Flask(__name__)

@app.route("/")
def hello():
    response = requests.get("https://api.ipify.org?format=json")
    return f"Hello World! Your IP is {response.json().get('ip')}."