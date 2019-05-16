#!/usr/bin/env bash
xl apply -f /Users/bmoussaud/Workspace/bmoussaud/xl-source-to-prod/services/xebialabs/release-clean.yaml

kubectl delete rs -l tier=staging  -n xl-demo-staging
kubectl delete rs -l tier=production  -n xl-demo-production
kubectl delete rs -l application=xl-demo-uccm-app  -n xl-demo-production
kubectl delete rs -l application=xl-demo-uccm-app  -n xl-demo-staging
kubectl delete deployment -l application=xl-demo-uccm-app  -n xl-demo-staging
kubectl delete deployment -l application=xl-demo-uccm-app  -n xl-demo-production

