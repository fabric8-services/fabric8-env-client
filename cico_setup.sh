#!/bin/bash

set -x
set -e

# Source environment variables of the jenkins slave that might interest this worker
function load_jenkins_vars() {
  if [ -e "jenkins-env" ]; then
    cat jenkins-env \
      | grep -E "(DEVSHIFT_TAG_LEN|JENKINS_URL|GIT_BRANCH|GIT_COMMIT|BUILD_NUMBER|ghprbSourceBranch|ghprbActualCommit|BUILD_URL|ghprbPullId)=" \
      | sed 's/^/export /g' \
      > ~/.jenkins-env
    source ~/.jenkins-env
  fi
}

function install_deps() {
  # We need to disable selinux for now, XXX
  /usr/sbin/setenforce 0 || :

  # Get all the deps in
  yum -y install \
    docker \
    make \
    curl

  service docker start

  echo 'CICO: Dependencies installed'
}


function build() {
  make docker-start
  make docker-deps
  make docker-build
  echo 'CICO: Build complete'
}

function cico_setup() {
  load_jenkins_vars;
  install_deps;
}
