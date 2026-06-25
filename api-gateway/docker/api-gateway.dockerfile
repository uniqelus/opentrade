FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY bin/api-gateway app/api-gateway

USER nonroot:nonroot

ENTRYPOINT [ "/app/api-gateway" ]
CMD [ "--config", "/etc/api-gateway.yaml"]