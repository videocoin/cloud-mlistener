#!/bin/bash

readonly CHART_NAME=mlistener
readonly CHART_DIR=./deploy/helm

CONSUL_ADDR=${CONSUL_ADDR:=127.0.0.1:8500}
ENV=${ENV:=dev}
VERSION=${VERSION:=`git describe --abbrev=0`-`git rev-parse --short HEAD`}

function log {
  local readonly level="$1"
  local readonly message="$2"
  local readonly timestamp=$(date +"%Y-%m-%d %H:%M:%S")
  >&2 echo -e "${timestamp} [${level}] [$SCRIPT_NAME] ${message}"
}

function log_info {
  local readonly message="$1"
  log "INFO" "$message"
}

function log_warn {
  local readonly message="$1"
  log "WARN" "$message"
}

function log_error {
  local readonly message="$1"
  log "ERROR" "$message"
}

function update_deps() {
    log_info "Syncing dependencies..."
    helm dependencies update --kube-context ${KUBE_CONTEXT} ${CHART_DIR}
}

function has_jq {
  [ -n "$(command -v jq)" ]
}

function has_consul {
  [ -n "$(command -v consul)" ]
}

function has_helm {
  [ -n "$(command -v helm)" ]
}

function get_vars() {
    log_info "Getting variables..."
    readonly KUBE_CONTEXT=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/common/kube_context`
    readonly DB_URI=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/dbUri`
    readonly MQ_URI=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/mqUri`
    readonly REDIS_URI=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/redisUri`
    readonly CONTRACT_ADDR=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/contractAddr`
    readonly NODE_HTTP_ADDR=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/nodeHttpAddr`
    readonly SENTRY_DSN=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/sentryDsn`
}

function get_vars_ci() {
    log_info "Getting ci variables..."
    readonly KUBE_CONTEXT=`curl --silent --user ${CONSUL_AUTH} http://consul.${ENV}.videocoin.network/v1/kv/config/${ENV}/common/kube_context?raw`
    readonly DB_URI=`curl --silent --user ${CONSUL_AUTH} http://consul.${ENV}.videocoin.network/v1/kv/config/${ENV}/services/${CHART_NAME}/secrets/dbUri?raw`
    readonly MQ_URI=`curl --silent --user ${CONSUL_AUTH} http://consul.${ENV}.videocoin.network/v1/kv/config/${ENV}/services/${CHART_NAME}/secrets/mqUri?raw`
    readonly REDIS_URI=`curl --silent --user ${CONSUL_AUTH} http://consul.${ENV}.videocoin.network/v1/kv/config/${ENV}/services/${CHART_NAME}/secrets/redisUri?raw`
    readonly CONTRACT_ADDR=`curl --silent --user ${CONSUL_AUTH} http://consul.${ENV}.videocoin.network/v1/kv/config/${ENV}/services/${CHART_NAME}/secrets/contractAddr?raw`
    readonly NODE_HTTP_ADDR=`curl --silent --user ${CONSUL_AUTH} http://consul.${ENV}.videocoin.network/v1/kv/config/${ENV}/services/${CHART_NAME}/secrets/nodeHttpAddr?raw`
    readonly SENTRY_DSN=`curl --silent --user ${CONSUL_AUTH} http://consul.${ENV}.videocoin.network/v1/kv/config/${ENV}/services/${CHART_NAME}/secrets/sentryDsn?raw`
}

function deploy() {
    log_info "Deploying ${CHART_NAME} version ${VERSION}"
    helm upgrade \
        --kube-context "${KUBE_CONTEXT}" \
        --install \
        --set image.tag="${VERSION}" \
        --set secrets.dbUri="${DB_URI}" \
        --set secrets.mqUri="${MQ_URI}" \
        --set secrets.redisUri="${REDIS_URI}" \
        --set secrets.contractAddr="${CONTRACT_ADDR}" \
        --set secrets.nodeHttpAddr="${NODE_HTTP_ADDR}" \
        --set secrets.sentryDsn="${SENTRY_DSN}" \
        --wait ${CHART_NAME} ${CHART_DIR}
}

if ! $(has_jq); then
    log_error "Could not find jq"
    exit 1
fi

if ! $(has_consul); then
    log_error "Could not find consul"
    exit 1
fi

if ! $(has_helm); then
    log_error "Could not find helm"
    exit 1
fi

if [ "${CI_ENABLED}" = "1" ]; then
  get_vars_ci
else
  get_vars
fi

update_deps
deploy

exit $?