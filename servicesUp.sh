#!/bin/bash

case "$1" in
start)
  echo "Starting Services..."
  docker compose up -d
  sleep 10
  go run ./services/persistdata/main.go
  ;;
stop)
  echo "Stopping Services..."
  docker compose down
  ;;
logs)
  echo "Fetching logs for Compose Stack LanChat..."
  docker compose logs
  ;;
*)
  echo "Usage: $0 {start|stop|logs}"
  exit 1
  ;;
esac
