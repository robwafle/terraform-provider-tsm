# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.241.1/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.18, 1.17, 1-bullseye, 1.18-bullseye, 1.17-bullseye, 1-buster, 1.18-buster, 1.17-buster
ARG VARIANT="1.19-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

#[Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive 
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1

# install kubectl
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
RUN chmod +x ./kubectl
RUN mv ./kubectl /usr/local/bin

COPY docker-in-docker-debian.sh /usr/local/bin/
RUN /usr/local/bin/docker-in-docker-debian.sh "true" "${USERNAME}" "false" "latest"

# [Optional] Uncomment the next lines to use go get to install anything else you need
USER vscode
RUN go install github.com/goreleaser/goreleaser@latest
# RUN go get -x <your-dependency-or-tool>

COPY docker-bind-mount.sh /usr/local/share/
COPY devcontainer-init.sh /usr/local/bin/

# Mount for docker-in-docker
VOLUME [ "/var/lib/docker" ]

ENTRYPOINT [ "/usr/local/bin/devcontainer-init.sh" ]
CMD [ "sleep", "infinity" ]