#!/usr/bin/env bash

echo "checking code format"

files=($(gofmt -s -l -e **/*.go))

if [ ${#files[@]} -eq 0 ]; then
  echo "OK All *.go files formatted"
  exit 0
else
  echo "formatting errors:"
  echo "${files[@]}"
  echo "ERROR ${#files[@]} files not formatted"
  echo "run make format"
  exit 1
fi


