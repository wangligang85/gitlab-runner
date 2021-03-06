FROM golang:1.13.8

# Install required dependencies
RUN apt-get update -yq && \
    apt-get install -yq locales make xz-utils \
                        ruby ruby-dev python-pip \
                        dpkg-sig createrepo rpm \
                        zip libffi-dev jq && \
    rm -rf /var/lib/apt/lists/*

# Set default locale for the environment
RUN echo "en_US UTF-8" > /etc/locale.gen; \
    locale-gen en_US.UTF-8

ENV LANG=en_US.UTF-8
ENV LANGUAGE=en_US:en
ENV LC_ALL=en_US.UTF-8

ARG DOCKER_VERSION

# Install docker client
RUN wget -q https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz -O /tmp/docker.tar.gz && \
    tar -xzf /tmp/docker.tar.gz -C /tmp/ && \
    cp /tmp/docker/docker* /usr/bin && \
    chmod +x /usr/bin/docker* && \
    rm -rf /tmp/*

ARG BUILDX_VERSION

# Install docker buildx plugin
RUN mkdir -p ~/.docker/cli-plugins && \
    wget -q https://github.com/docker/buildx/releases/download/v${BUILDX_VERSION}/buildx-v${BUILDX_VERSION}.linux-amd64 -O ~/.docker/cli-plugins/docker-buildx && \
    chmod a+x ~/.docker/cli-plugins/docker-buildx

ARG GIT_LFS_VERSION
ARG GIT_LFS_256_CHECKSUM

COPY dockerfiles/ci/install_git_lfs /tmp/
RUN /tmp/install_git_lfs

ARG KUBECTL_VERSION

# Install kubectl
RUN wget -q https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl -O /usr/local/bin/kubectl && \
    chmod +x /usr/local/bin/kubectl

COPY Makefile* /tmp/
WORKDIR /tmp
RUN make deps package-deps packagecloud-deps

RUN go get -u github.com/jstemmer/go-junit-report
