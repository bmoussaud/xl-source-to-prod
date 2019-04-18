#!/usr/bin/env bash

curl http://localhost:8001/api/v1/proxy/namespaces/staging/services/frontend:80/
curl http://localhost:8001/api/v1/proxy/namespaces/xl-demo-staging/services/frontend:80/
