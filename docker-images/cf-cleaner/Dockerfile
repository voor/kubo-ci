FROM alpine:latest
RUN apk add --update openssl jq && \
  rm -rf /var/cache/apk/* && \
  wget 'https://cli.run.pivotal.io/stable?release=linux64-binary&version=6.30.0&source=github-rel' -O cf.tgz && \
  tar -xvf cf.tgz && \
  rm cf.tgz && \
  install cf /usr/local/bin && \
  wget 'https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-2.0.44-linux-amd64' -O bosh && \
  install bosh /usr/local/bin && \
  ln -s /usr/local/bin/bosh /usr/local/bin/bosh-cli
