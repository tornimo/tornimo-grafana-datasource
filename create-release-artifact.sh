#!/usr/bin/env bash

make
cp -r dist tornimo-datasource
zip tornimo-datasource.zip tornimo-datasource/*
zip tornimo-datasource.zip tornimo-datasource/*/*
rm -rf tornimo-datasource.zip tornimo-datasource
