FROM golang:1.17.7

ENV PATH="$PATH:$HOME/.porter"

    ## Install Ansible
RUN pip install --upgrade pip && \
    pip install --upgrade setuptools wheel && \
    pip install --upgrade ansible && \
    ## Install Porter
    curl -L https://cdn.porter.sh/latest/install-linux.sh | bash

ENTRYPOINT []
CMD []

