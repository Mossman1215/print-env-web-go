FROM quay.io/fedora/fedora:38

COPY main /usr/bin/print-env-web

CMD print-env-web
