require 'sinatra'
require "sinatra/json"

def generate_samples(n)
  data = []
  n.times do
    data << Random.rand()
  end
  data
end

def mean(data)
  data.sum / data.length
end

def variance(data)
  m = mean(data)
  mean data.map { |x| (x - m) * (x - m) }
end

disable :logging

get '/' do
  data = generate_samples(100000)
  json data: data, mean: mean(data), variance: variance(data)
end

get '/hello' do
  json message: 'hello there!'
end
