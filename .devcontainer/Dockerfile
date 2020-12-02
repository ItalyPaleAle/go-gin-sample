# Go version
ARG VARIANT=1.15
FROM mcr.microsoft.com/vscode/devcontainers/go:${VARIANT}

# Install Node.js
ARG INSTALL_NODE="true"
ARG NODE_VERSION="lts/*"
RUN if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "source /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# Workdir
WORKDIR /home/vscode/workspace
