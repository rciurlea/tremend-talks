const express = require('express');
const cluster = require('cluster');

function generate_samples(n) {
  const data = [];
  for (let i = 0; i < n; i++) {
    data.push(Math.random())
  }
  return data;
}

function mean(data) {
  return data.reduce((acc, curr) => acc + curr) / data.length;
}

function variance(data) {
  const m = mean(data);
  return mean(data.map(x => (x - m) * (x - m)));
}

if (cluster.isMaster) {
  const cpuCount = 2 * require('os').cpus().length;

  for (var i = 0; i < cpuCount; i += 1) {
     cluster.fork();
  }
} else {
  const app = express();

  app.get('/hello', (req, res) => {
      res.json({'message': 'hello there!'});
  });

  app.get('/stats', (req, res) => {
    const data = generate_samples(20);
    res.json({
      data,
      mean: mean(data),
      variance: variance(data),
    });
  });

  app.listen(3000);
  console.log('Application running!');
}
