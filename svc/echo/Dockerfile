FROM golang:1.16 AS builder

# Args
ARG VERSION
ENV VERSION ${VERSION}

# Make app
WORKDIR $GOPATH/src/github.com/monorepo/svc/echo
COPY . ./
RUN make build

FROM ubuntu:bionic

# add Certs
RUN set -ex \
    && echo '\
path-exclude /usr/share/doc/*\n\
path-include /usr/share/doc/*/copyright\n\
path-exclude /usr/share/man/*\n\
path-exclude /usr/share/groff/*\n\
path-exclude /usr/share/info/*\n\
path-exclude /usr/share/lintian/*\n\
path-exclude /usr/share/linda/*\n\
' > /etc/dpkg/dpkg.cfg.d/01_nodoc \
    && echo '\
APT::Install-Recommends "false";\n\
' > /etc/apt/apt.conf.d/99synaptic \
    && apt-get update && apt-get install -y ca-certificates tzdata \
    && apt-get clean \
    && apt-get autoremove \
    && rm -rf \
        /var/lib/apt/lists/* \
        /tmp/* \
        /var/tmp/* \
        /var/cache/apt/archives \
        /usr/share/doc/ \
        /usr/share/man/ \
        /usr/share/info/ \
        /usr/share/groff/ \
        /usr/share/locale/

# Copy complied binary
COPY --from=builder /go/src/github.com/monorepo/svc/echo/out/echo-api /bin/

# Run app and expose port
ARG PORT=8000
EXPOSE ${PORT}

ENTRYPOINT ["/bin/echo-api"]
