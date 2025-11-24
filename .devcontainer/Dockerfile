FROM ubuntu:24.04

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

ENV DEBIAN_FRONTEND=noninteractive
ENV LANG=en_US.UTF-8

### DEV ENVIRONMENT ###

# Install dependencies
RUN apt-get update \
    && apt-get upgrade -y \
    && apt-get install -y --no-install-recommends \
    git \
    wget \
    gnupg \
    build-essential \
    ca-certificates \
    software-properties-common \
    shellcheck \
    && rm -rf /var/lib/apt/lists/*

# Install Hadolint
ENV HADOLINT_VERSION=v2.12.0
RUN wget -q https://github.com/hadolint/hadolint/releases/download/${HADOLINT_VERSION}/hadolint-Linux-x86_64 \
    && wget -q https://github.com/hadolint/hadolint/releases/download/${HADOLINT_VERSION}/hadolint-Linux-x86_64.sha256 \
    # Verify file
    && sha256sum -c hadolint-Linux-x86_64.sha256 \
    && rm hadolint-Linux-x86_64.sha256 \
    # Make executable
    && chmod +x hadolint-Linux-x86_64 \
    # Add to PATH
    && mv hadolint-Linux-x86_64 /bin/hadolint \
    # Verify installation
    && hadolint --version

### FRONTEND ###

RUN wget -q https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.3/install.sh -O /tmp/nvm_install.sh \
    && bash /tmp/nvm_install.sh \
    # in lieu of restarting the shell
    && source "${HOME}/.nvm/nvm.sh" \
    # Download and install Node.js:
    && nvm install 24 \
    # Verify the Node.js version:
    && node -v \
    && nvm current \
    && npm -v \
    # Download and install Yarn:
    && corepack enable yarn \
    # Verify Yarn version:
    && yarn -v

### BACKEND ###
