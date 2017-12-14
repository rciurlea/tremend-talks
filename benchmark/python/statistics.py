import random
from flask import Flask, jsonify

num_samples = 100000
app = Flask(__name__)


@app.route("/")
def stats():
    data = generate_samples(num_samples)
    return jsonify({
        'data': data,
        'mean': mean(data),
        'variance': variance(data)
    })


@app.route("/hello")
def hello():
    return jsonify({'message': 'hello there!'})


def generate_samples(n):
    data = []
    for i in range(num_samples):
        data.append(random.random())
    return data


def mean(data):
    return sum(data) / len(data)


def variance(data):
    m = mean(data)
    return mean([(x - m) * (x - m) for x in data])


if __name__ == "__main__":
    app.run(host='0.0.0.0')
