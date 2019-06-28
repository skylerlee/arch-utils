# -*- coding: utf-8 -*-

import sys
import base64
from urllib import request


def fetch_feed(url):
    resp = request.urlopen(url)
    return resp.read()

def padding(data):
    length = len(data)
    pad = 4 - length % 4
    return data + b'=' * pad

def get_lines(s):
    lines = s.split('\n')
    return filter(lambda i: i, lines)

def parse_ssr_url(url):
    pass

def parse_feed(url):
    raw = fetch_feed(url)
    data = padding(raw)
    text = base64.decodebytes(data).decode('utf-8')
    for line in get_lines(text):
        parse_ssr_url(line)


if __name__ == '__main__':
    parse_feed(sys.argv[1])
