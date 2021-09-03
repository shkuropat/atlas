#!/bin/bash

rm -rf node_modules
npm install

rm -rf ../../node_modules
mv node_modules ../..
