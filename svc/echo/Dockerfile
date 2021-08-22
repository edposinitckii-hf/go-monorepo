# Monorepo args
ARG REPO=${GOPATH}/src/github.com/monorepo
ARG SVC_PATH=svc/echo

FROM golang:1.16 AS builder

# Monorepo args
ARG REPO
ARG SVC_PATH

# Args
ARG VERSION
ENV VERSION ${VERSION}

# Make app
WORKDIR ${REPO}
COPY . ./

WORKDIR ${REPO}/${SVC_PATH}
RUN make build

FROM ubuntu:bionic

# Monorepo args
ARG REPO
ARG SVC_PATH

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
COPY --from=builder ${REPO}/${SVC_PATH}/out/echo-server /bin/

# Run app and expose port
ARG PORT=9001
EXPOSE ${PORT}

ENTRYPOINT ["/bin/echo-server"]
