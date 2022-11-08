FROM mcr.microsoft.com/vscode/devcontainers/base:dev-focal

ARG GO_VERSION="1.19.2"
ARG TERRAFORM_VERSION="1.0.5"

COPY fix-sources.sh .
RUN chmod +x fix-sources.sh
RUN ./fix-sources.sh
RUN apt-key adv --refresh-keys --keyserver hkp://keyserver.ubuntu.com:80

ADD ctccacerts.crt /usr/local/share/ca-certificates/ctccacerts.crt
RUN chmod 644 /usr/local/share/ca-certificates/ctccacerts.crt & update-ca-certificates
ENV NODE_EXTRA_CA_CERTS=/etc/ssl/certs/ctccacerts.pem

RUN apt update

# install kubectl
RUN my_arch=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/) \ 
    && curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/$my_arch/kubectl" \
    && curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/$my_arch/kubectl.sha256" \
    && chmod +x ./kubectl \
    && mv ./kubectl /usr/local/bin

# Install Az Cli
RUN curl -sL https://aka.ms/InstallAzureCLIDeb | bash

# set env when using container
COPY .bash_profile /root/.bash_profile
RUN chmod +x /root/.bash_profile

# update env during docker build
ENV GOPATH=$HOME/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# force latest go installation against goreleaser
RUN my_arch=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/) \
    && wget --verbose --show-progress --no-check-certificate -L "https://golang.org/dl/go${GO_VERSION}.linux-$my_arch.tar.gz" \
    && tar -xf "go${GO_VERSION}.linux-$my_arch.tar.gz" \
    && rm -rvf /usr/local/go \
    && mv -v go /usr/local

# Terraform
RUN wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor | tee /usr/share/keyrings/hashicorp-archive-keyring.gpg \
    && echo "deb [trusted=yes signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/hashicorp.list \
    && apt update \
    && apt install terraform=${TERRAFORM_VERSION}

# install terraform docs
RUN my_arch=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/) \
    && wget --verbose --show-progress --no-check-certificate -O "terraform-docs.tar.gz" -L "https://github.com/terraform-docs/terraform-docs/releases/download/v0.16.0/terraform-docs-v0.16.0-$(uname)-$my_arch.tar.gz" \
    && tar -xf "terraform-docs.tar.gz" \
    && chmod +x ./terraform-docs \
    && mv terraform-docs /usr/local/bin

# install python
RUN apt update -y \
    && apt install -y python3-pip \
    && pip3 install pre-commit

# install goreleaser
RUN go install github.com/goreleaser/goreleaser@latest

# RUN go get -x <your-dependency-or-tool>
CMD ["bash", "-l"]