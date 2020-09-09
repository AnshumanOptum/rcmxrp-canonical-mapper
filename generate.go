package main

//go:generate mkdir -p model/generated
//go:generate rm -rf model/generated/*
//go:generate $GOPATH/bin/gogen-avro --package model model/generated schema/*
