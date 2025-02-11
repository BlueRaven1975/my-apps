from flask import Flask, jsonify
import requests

app = Flask(__name__)

@app.route('/')
def hello():
    return "Hello World!"

@app.route('/quote')
def quote():
    try:
        response = requests.get("https://zenquotes.io/api/random", timeout=5)
        response.raise_for_status()
        data = response.json()
        return jsonify({"quote": data[0]["q"], "author": data[0]["a"]})
    except requests.RequestException as e:
        return jsonify({"error": "Failed to fetch quote", "details": str(e)}), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8081)
