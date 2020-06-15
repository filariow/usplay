#-------------------------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See https://go.microsoft.com/fwlink/?linkid=2090316 for license information.
#-------------------------------------------------------------------------------------------------------------

FROM golang:1.14

ENV DEBIAN_FRONTEND=noninteractive
# Avoid warnings by switching to noninteractive

# This Dockerfile adds a non-root user with sudo access. Use the "remoteUser"
# property in devcontainer.json to use it. On Linux, the container user's GID/UIDs
# will be updated to match your local UID/GID (when using the dockerFile property).
# See https://aka.ms/vscode-remote/containers/non-root-user for details.
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID
ARG DOCKER_GUID

ARG GHCLI_VERSION=0.10.0
ARG NODEJS_VERSION=14
ARG PROTOC_GRPCWEB_VERSION=1.1.0
ARG COMPOSE_VERSION=1.26.0
# Configure apt, install packages and tools
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils dialog 2>&1 \
    #
    # Verify git, process tools, lsb-release (common in install instructions for CLIs) installed
    && apt-get -y install git openssh-client less iproute2 procps lsb-release \
    #
    #
    # Install Go tools w/module support
    && mkdir -p /tmp/gotools \
    && cd /tmp/gotools \
    && GO111MODULE=on go get -v golang.org/x/tools/gopls@latest 2>&1 \
    && GO111MODULE=on go get -v \
        honnef.co/go/tools/...@latest \
        golang.org/x/tools/cmd/gorename@latest \
        golang.org/x/tools/cmd/goimports@latest \
        golang.org/x/tools/cmd/guru@latest \
        golang.org/x/lint/golint@latest \
        github.com/mdempsky/gocode@latest \
        github.com/cweill/gotests/...@latest \
        github.com/haya14busa/goplay/cmd/goplay@latest \
        github.com/sqs/goreturns@latest \
        github.com/josharian/impl@latest \
        github.com/davidrjenni/reftools/cmd/fillstruct@latest \
        github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest  \
        github.com/ramya-rao-a/go-outline@latest  \
        github.com/acroca/go-symbols@latest  \
        github.com/godoctor/godoctor@latest  \
        github.com/rogpeppe/godef@latest  \
        github.com/zmb3/gogetdoc@latest \
        github.com/fatih/gomodifytags@latest  \
        github.com/mgechev/revive@latest  \
        github.com/go-delve/delve/cmd/dlv@latest 2>&1 \
    #
    # Install Go tools w/o module support
    && go get -v github.com/alecthomas/gometalinter 2>&1 \
    #
    # Install gocode-gomod
    && go get -x -d github.com/stamblerre/gocode 2>&1 \
    && go build -o gocode-gomod github.com/stamblerre/gocode \
    && mv gocode-gomod $GOPATH/bin/ \
    #
    # Install golangci-lint
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin 2>&1 \
    #
    # Create a non-root user to use if preferred - see https://aka.ms/vscode-remote/containers/non-root-user.
    && groupadd --gid $USER_GID $USERNAME \
    && useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && mkdir -p /home/${USERNAME}/.config \
    && chown -R ${USER_UID}:${USER_GID} /home/${USERNAME} \
    # [Optional] Add sudo support
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME \
    # zsh, oh-my-zsh and tools
    && apt-get install -y vim less zsh tree \
    && mkdir -p /tmp/zsh \
    && cd /tmp/zsh \
    && su ${USERNAME} sh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" \
    && echo 'alias ff="find . -iname " >> /home/${USERNAME}/.zshrc' \
    && echo 'alias ls="ls -h --color=auto " >> /home/${USERNAME}/.zshrc' \
    && echo 'alias ll="ls -l " >> /home/${USERNAME}/.zshrc' \
    && echo 'alias l="ls -l " >> /home/${USERNAME}/.zshrc' \
    && echo 'alias la="ls -al " >> /home/${USERNAME}/.zshrc' \
    && echo 'alias l.="ls -al " >> /home/${USERNAME}/.zshrc' \
    #
    # GitHub CLI
    && mkdir -p /tmp/gh \
    && cd /tmp/gh \
    && curl -sLO  https://github.com/cli/cli/releases/download/v0.7.0/gh_${GHCLI_VERSION}_linux_amd64.deb \
    && apt install ./gh_${GHCLI_VERSION}_linux_amd64.deb \
    #
    # Permit VS Code find the installed extensions after container rebuild
    && mkdir -p /home/$USERNAME/.vscode-server/extensions \
        /home/$USERNAME/.vscode-server-insiders/extensions \
    && chown -R $USERNAME \
        /home/$USERNAME/.vscode-server \
        /home/$USERNAME/.vscode-server-insiders  \
    # Install kubectl
    && curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl \
    && chmod +x kubectl \
    && mv ./kubectl /usr/bin/kubectl \
    # 
    # Protobuf
    && apt-get install -y protobuf-compiler \
    && go get -u google.golang.org/grpc \
    && go get -u github.com/golang/protobuf/protoc-gen-go \
    && GO111MODULE=on go get github.com/uber/prototool/cmd/prototool@dev \
    # Add write permission for /go/pkg
    && chown -R ${USERNAME}:${USERNAME} /go \
    #
    # Install Docker CE CLI
    && apt-get install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common lsb-release \
    && curl -fsSL https://download.docker.com/linux/$(lsb_release -is | tr '[:upper:]' '[:lower:]')/gpg | (OUT=$(apt-key add - 2>&1) || echo $OUT) \
    && add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/$(lsb_release -is | tr '[:upper:]' '[:lower:]') $(lsb_release -cs) stable" \
    && apt-get update \
    && apt-get install -y docker-ce-cli \
    && groupadd -g ${DOCKER_GUID} docker \
    && usermod -aG docker ${USERNAME} \
    #
    # Install Docker Compose
    && curl -sSL "https://github.com/docker/compose/releases/download/${COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose \
    && chmod a+x /usr/local/bin/docker-compose \
    && echo ${DOCKER_GUID} > /home/${USERNAME}/dockerguid \
    #
    # Install NodeJS
    # Using Debian, as root
    && curl -sL https://deb.nodesource.com/setup_${NODEJS_VERSION}.x | bash - \
    && apt-get install -y nodejs \
    && npm install -g @vue/cli \
    #
    # gRPC Web
    && mkdir -p /tmp/proto \
    && cd /tmp/proto \
    && wget https://github.com/grpc/grpc-web/releases/download/${PROTOC_GRPCWEB_VERSION}/protoc-gen-grpc-web-${PROTOC_GRPCWEB_VERSION}-linux-x86_64  \
    && chmod +x protoc-gen-grpc-web-${PROTOC_GRPCWEB_VERSION}-linux-x86_64 \
    && mv protoc-gen-grpc-web-${PROTOC_GRPCWEB_VERSION}-linux-x86_64 /usr/local/bin/protoc-gen-grpc-web \ 
    #
    ## Helm
    && mkdir -p /tmp/helm \
    && cd /tmp/helm \
    && curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 \ 
    && chmod 700 get_helm.sh \ 
    && ./get_helm.sh \ 
    #
    # Clean up
    && cd / \
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/* /go/src /tmp/* 

ENV ZSH_THEME=agnoster

# Update this to "on" or "off" as appropriate
ENV GO111MODULE=auto

# Switch back to dialog for any ad-hoc use of apt-get
ENV DEBIAN_FRONTEND=dialog
