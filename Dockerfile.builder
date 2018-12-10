FROM centos:7
LABEL maintainer "Devtools <devtools@redhat.com>"
LABEL author "Devtools <devtools@redhat.com>"
ENV LANG=en_US.utf8
ARG USE_GO_VERSION_FROM_WEBSITE=0

# Some packages might seem weird but they are required by the RVM installer.
RUN yum install epel-release -y \
    && yum --enablerepo=centosplus --enablerepo=epel install -y \
      findutils \
      git \
      $(test "$USE_GO_VERSION_FROM_WEBSITE" != 1 && echo "golang") \
      make \
      procps-ng \
      tar \
      wget \
      which \
    && yum clean all

RUN if [[ "$USE_GO_VERSION_FROM_WEBSITE" = 1 ]]; then cd /tmp \
    && wget --no-verbose https://dl.google.com/go/go1.10.5.linux-amd64.tar.gz \
    && echo "a035d9beda8341b645d3f45a1b620cf2d8fb0c5eb409be36b389c0fd384ecc3a go1.10.5.linux-amd64.tar.gz" > checksum \
    && sha256sum -c checksum \
    && tar -C /usr/local -xzf go1.10.5.linux-amd64.tar.gz \
    && rm -f go1.10.5.linux-amd64.tar.gz; \
    fi
ENV PATH=$PATH:/usr/local/go/bin

# Get dep for Go package management and make sure the directory has full rwz permissions for non-root users
ENV GOPATH /tmp/go
RUN mkdir -p $GOPATH/bin && chmod a+rwx $GOPATH
RUN cd $GOPATH/bin \
	&& curl -L -s https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 -o dep \
	&& echo "287b08291e14f1fae8ba44374b26a2b12eb941af3497ed0ca649253e21ba2f83 dep" > dep-linux-amd64.sha256 \
	&& sha256sum -c dep-linux-amd64.sha256

ENTRYPOINT ["/bin/bash"]
