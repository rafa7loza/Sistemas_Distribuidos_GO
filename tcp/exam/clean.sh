#!/bin/sh
for f in "clients/"; do
  echo "Removing $f"
  rm -r $f
done

rm server.log
