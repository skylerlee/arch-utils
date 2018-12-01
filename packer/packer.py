"""
utility for archiving and extracting files.
"""

import argparse


def process(action, input, output, passwd):
    print(action, input, output, passwd)


def main():
    parser = argparse.ArgumentParser(prog='packer', description=__doc__)
    parser.add_argument('action', choices=['a', 'x'], help='action flag')
    parser.add_argument('input', help='input file or directory')
    parser.add_argument('-o', dest='output', metavar='output', help='path to place output')
    parser.add_argument('-p', dest='passwd', action='store_true', help='whether to apply password')
    args = parser.parse_args()
    process(**args.__dict__)


if __name__ == '__main__':
    main()
