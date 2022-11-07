ARG VERSION="1.19.2" # go version
FROM mcr.microsoft.com/vscode/devcontainers/base:dev-hirsute

COPY fix-sources.sh .
RUN chmod +x fix-sources.sh
RUN ./fix-sources.sh

RUN sudo apt update

# install kubectl
RUN my_arch=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/) && curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/$my_arch/kubectl" && curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/$my_arch/kubectl.sha256"
RUN chmod +x ./kubectl
RUN mv ./kubectl /usr/local/bin

RUN apt update -y \
    && apt install -y python3-pip \
    && pip install pre-commit

# Install Az Cli
RUN curl -sL https://aka.ms/InstallAzureCLIDeb | bash

# [Optional] Uncomment the next lines to use go get to install anything else you need

ENV VERSION="1.19.2"
RUN my_arch=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/) && wget --verbose --show-progress -L "https://golang.org/dl/go${VERSION}.linux-$my_arch.tar.gz" && tar -xf "go${VERSION}.linux-$my_arch.tar.gz"
#RUN chown -R root:root ./go
RUN mv -v go /usr/local

# set env when using container
COPY .bash_profile /root/.bash_profile
RUN chmod +x /root/.bash_profile

# update env during docker build
ENV GOPATH=$HOME/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

RUN go install github.com/goreleaser/goreleaser@latest

RUN sudo apt-key adv --refresh-keys --keyserver keyserver.ubuntu.com

# Terraform
RUN wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor | tee /usr/share/keyrings/hashicorp-archive-keyring.gpg
RUN echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/hashicorp.list
RUN apt update && apt install terraform=1.0.5

# install terraform docs
RUN go install github.com/terraform-docs/terraform-docs@v0.16.0


# RUN go get -x <your-dependency-or-tool>
CMD ["bash", "-l"]
#
# 
# # Kubectl
# RUN curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
# RUN echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | tee /etc/apt/sources.list.d/kubernetes.list
# RUN apt-get update 
# RUN apt-get install -y kubectl
# 
# # Helm
# RUN curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
#