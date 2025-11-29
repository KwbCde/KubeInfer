# KubeInfer

A custom Kubernetes Operator for running ML inference jobs using declarative CRDs.

This project introduces a custom `InferenceJob` CRD that allows users to submit inference requests by specifying:
- **model** – which ML model to run
- **input** – input source or payload
- **image** – container image that performs inference

The controller manages the lifecycle of each job by:
- initializing status phases (`Pending`, `Running`, `Succeeded`, `Failed`)
- reconciling state changes through a simple state machine
- (in progress) creating Pods for inference execution
- (in progress) tracking Pod status and updating job results

## Status

**Under development**  
- CRD and basic reconciliation logic implemented  
- Phase initialization system in place  
- Pod creation and job execution logic coming next  

## Goals

- Simple declarative interface for launching ML inference
- Automatic Pod orchestration and cleanup
- Extensible design for additional runtime backends
