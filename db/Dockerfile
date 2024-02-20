FROM alpine as base 

ARG VERSION=3.11.2
RUN set -x \
    && apk add --no-cache curl \
    && ARCH=$(if [ "$(uname -m)" == "arm64" ] || [ "$(uname -m)" == "aarch64" ]; then echo "arm64"; else echo "x86_64"; fi) \
    && curl -o /bin/goose -sSL https://github.com/pressly/goose/releases/download/v${VERSION}/goose_linux_${ARCH} \
    && chmod +x /bin/goose

FROM scratch

WORKDIR /migraion

COPY . .
COPY --from=base /bin/goose /bin/
ENTRYPOINT [ "/bin/goose" ]
CMD ["--help"]
