#/bin/bash

rm -rf test/*
cp /stor-shared/ubuntu/setup_csc_376_pa5.py test/setup.py

mkdir -p datasets
rm -rf datasets/*
cp -r /stor-shared/ubuntu/dataset1_client_server $PWD/datasets/dataset1_client_server


