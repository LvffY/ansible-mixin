FROM golang:1.17

ENV PATH="$PATH:$HOME/.porter" \
    VERSION="v1.0.0-beta.4"


    ## Install Ansible
RUN pip install --upgrade pip && \
    pip install --upgrade setuptools wheel && \
    pip install --upgrade ansible && \
    ## Install Porter
    curl -L https://cdn.porter.sh/$VERSION/install-mac.sh | bash


ENTRYPOINT []
CMD []

