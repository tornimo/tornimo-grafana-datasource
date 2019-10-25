#!/usr/bin/env bash

make
cp -r dist tornimo-datasource

zip tornimo-datasource.zip tornimo-datasource
zip tornimo-datasource.zip tornimo-datasource/*
zip tornimo-datasource.zip tornimo-datasource/img/*
zip tornimo-datasource.zip tornimo-datasource/partials/*

rm -rf tornimo-datasource

