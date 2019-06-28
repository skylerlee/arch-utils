# -*- coding: utf-8 -*-

import sys
import base64
from urllib import request


def fetch_feed(url):
    resp = request.urlopen(url)
    return resp.read()

def padding(data):
    length = len(data)
    pad = -length % 4
    return data + b'=' * pad

def get_lines(s):
    lines = s.split('\n')
    return filter(lambda i: i, lines)

def parse_ssr_url(url):
    protocol = 'ssr://'
    if not url.startswith(protocol):
        return None
    raw = url[len(protocol):]
    data = padding(raw.encode('utf-8'))
    text = base64.b64decode(data, '-_').decode('utf-8')
    return text.split(':')

def parse_feed(url):
    raw = fetch_feed(url)
    data = padding(raw)
    text = base64.decodebytes(data).decode('utf-8')
    for line in get_lines(text):
        segs = parse_ssr_url(line)
        print(segs)


if __name__ == '__main__':
    parse_feed(sys.argv[1])
