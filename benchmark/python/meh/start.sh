#!/bin/sh
gunicorn -w 12 -k eventlet --bind 0.0.0.0:8000 meh.wsgi
