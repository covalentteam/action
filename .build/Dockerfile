# Retrieve the `golang:alpine` image to provide us the necessary Golang
# tooling for building Go binaries.
FROM golang:1.17.0-alpine3.14 AS build

WORKDIR /src
COPY . /src

# This environment variables define the experimental mode fature flags
#
# 0.                        GO111MODULE:
#                           It's change how Go imports packagescan be very useful
#                           for enabling the Go Modules behavior selecting a
#                           specific version based on git tags.
#
# 1.                        GO15VENDOREXPERIMENT:
#                           In short it allows you to put a package at A/vendor/B
#                           and import it as A from B.
ENV GO111MODULE             'on'
ENV GO15VENDOREXPERIMENT    1

# This environment variables define the target architecture of our binary
# and permit the static build to reduce drastically the binary size.
ENV GOOS                    'linux'
ENV GOARCH                  'amd64'
ENV CGO_LDFLAGS             '-s -w -extldflags "-static"'
ENV CGO_ENABLED             0

# 0.                        Golang will use git to download the dependencies
#                           source code.
# 1.                        Dowload our project dependencies.
# 2.                        Perform the go build with some flags to make our
#                           build produce a static binary.
RUN apk add git
RUN go get -d -v ./action
RUN go build -a -installsuffix cgo -trimpath -tags netgo -o  /bin/action -mod readonly ./action

# This steps contains tasks to minimize the artifact size and setup the final
# layer user.
RUN cp action.yml /etc/action.yml \
    && echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

# Create the `final` stage with the most basic that we need. Using scratch which
# will contains only our binary, allowing us to start with a fat build image
# and end up with a very small runtime image.
#
# In order to take advantage of this simplicity we handle copy our certificates
# from BUILD step and the binary of course.
FROM scratch as final

# Copy only the necessary staff to run the binary.
#                  | OPTIONS          | ORIGIN                                | DESTINATION
COPY --from=build                       /etc/ssl/certs/ca-certificates.crt      /etc/ssl/certs/
COPY --from=build                       /etc_passwd                             /etc/passwd
COPY --from=build   --chown=65534:0     /bin/action                             /bin/action

# Define arguments used to enrich image labels description.
ARG RELEASE_CREATED_AT
ARG RELEASE_REFERENCE
ARG RELEASE_REVISION
ARG RELEASE_VERSION
ARG RELEASE_URL

# Convert the build arguments in environment variables
ENV RELEASE_CREATED_AT  $RELEASE_CREATED_AT
ENV RELEASE_REFERENCE   $RELEASE_REFERENCE
ENV RELEASE_REVISION    $RELEASE_REVISION
ENV RELEASE_VERSION     $RELEASE_VERSION
ENV RELEASE_URL         $RELEASE_URL

# Covalentteam labels this section define labels usefull to customize you
# GitHub Actions Release.
LABEL com.covalentteam.template="true"
LABEL com.covalentteam.project="4hdk4hd-8fh42nv-3jhan5-j9ah3s"

# Marketplace labels this section define labels usefull to customize you
# GitHub Actions Release.
LABEL com.github.actions.name="covalentteam/template"
LABEL com.github.actions.description="A dead simple template to build robust GitHub Actions made with GO."
LABEL com.github.actions.icon="copy"
LABEL com.github.actions.color="blue"

# Containers labels this section define labels that can be very usefull to
# retrieving container informations on analitics projects.
LABEL org.opencontainers.image.authors="Vidal Vasconcelos"
LABEL org.opencontainers.image.base.name="covalentteam/template"
LABEL org.opencontainers.image.description="A dead simple template to build robust GitHub Actions made with Go."
LABEL org.opencontainers.image.documentation="https://github.com/covalentteam/template#readme"
LABEL org.opencontainers.image.licenses="Copyright © 2021 Covalentteam"
LABEL org.opencontainers.image.source="https://github.com/covalentteam/template"
LABEL org.opencontainers.image.title="Covalentteam template action"
LABEL org.opencontainers.image.vendor="covalentteam"
LABEL org.opencontainers.image.url="${RELEASE_URL}"
LABEL org.opencontainers.image.created="${RELEASE_CREATED_AT}"
LABEL org.opencontainers.image.ref.name="${RELEASE_REFERENCE}"
LABEL org.opencontainers.image.revision="${RELEASE_REVISION}"
LABEL org.opencontainers.image.version="${RELEASE_VERSION}"

# Set the user
USER nobody

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "/bin/action" ]
