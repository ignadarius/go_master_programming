#!/bin/bash
protoc calculator/calculatorpb/calc.proto --go_out=plugins=grpc:.