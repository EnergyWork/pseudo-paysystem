@echo off

if exist "swagger.yaml" (
    del swagger.yaml
)

echo generating the 'swagger' specification
swagger generate spec -o ./swagger.yaml --scan-models
