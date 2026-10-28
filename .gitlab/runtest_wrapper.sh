#!/bin/bash

export LANG=en_US.UTF-8
export LC_ALL=en_US.UTF-8

python3.14 test/setup.py
python3.14 test/test_driver.py
rm -rf datasets/*

