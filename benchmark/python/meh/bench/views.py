import random
from django.http import JsonResponse


num_samples = 20


def hello(request):
    return JsonResponse({'message': 'Hello there!'})


def stats(request):
    data = generate_samples(num_samples)
    return JsonResponse({
        'data': data,
        'mean': mean(data),
        'variance': variance(data)
    })


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
