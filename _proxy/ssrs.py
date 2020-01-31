# -*- coding: utf-8 -*-

import sys
import base64
from urllib import request, parse


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

def process_enc(s):
    result = parse.urlparse(s)
    path = result.path
    data = padding(path.encode('utf-8'))
    path2 = base64.b64decode(data, '-_').decode('utf-8')
    params = parse.parse_qsl(result.query)
    params2 = []
    for key, val in params:
        data = padding(val.encode('utf-8'))
        val2 = base64.b64decode(data, '-_').decode('utf-8')
        params2.append((key, val2))
    return {
        "path": path2,
        "params": params2,
    }

def parse_ssr_url(url):
    protocol = 'ssr://'
    if not url.startswith(protocol):
        return None
    raw = url[len(protocol):]
    data = padding(raw.encode('utf-8'))
    text = base64.b64decode(data, '-_').decode('utf-8')
    segs = text.split(':')
    segs[-1] = process_enc(segs[-1])
    return segs

def parse_feed(url):
    feed = fetch_feed(url)
    data = padding(feed)
    text = base64.decodebytes(data).decode('utf-8')
    for line in get_lines(text):
        segs = parse_ssr_url(line)
        print(segs)


if __name__ == '__main__':
    parse_feed(sys.argv[1])
