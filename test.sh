#!/bin/bash
set -e
DURATION=$1
[ x$DURATION = x ] && ( echo "No duration (\$1) specified"; exit 1 )
IMAGE=akihirosuda/demo-funker-issue
NW=demo-nw
CALLEE=demo-callee

log(){
    echo -e "\e[104m\e[97m[TEST]\e[49m\e[39m $@"
}

cleanup() {
    log "Cleaning up service $CALLEE"
    set -x
    docker service rm $CALLEE || true
    set +x
    sleep 3
    log "Cleaning up network $NW"
    docker network rm $NW || true
}

trap cleanup EXIT INT

# just in case
cleanup

log "Creating a network $NW"
set -x
docker network create -d overlay --attachable $NW
set +x

image_id=$(docker image inspect -f '{{.Id}}' $IMAGE)
log "Using image $IMAGE ($image_id)"
log "Creating a callee service $CALLEE (arg=$DURATION)"
set -x
docker service create --name $CALLEE --network $NW $image_id app callee $DURATION
set +x

log "Waiting for a while"
# FIXME
set -x
sleep 10
set +x

log "Running a caller"
set -x
docker run --rm --network $NW $image_id app caller $CALLEE
set +x

log "Done"
