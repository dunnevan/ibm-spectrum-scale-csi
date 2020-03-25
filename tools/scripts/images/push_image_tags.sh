#!/usr/bin/env bash

source tools/scripts/images/common.sh

#
# push_image_tags source_image push_base
#
# push_image_tags tags the source docker image with zero or more
# image tags based on TravisCI environment variables and the
# presence of git tags in the repository of the current working
# directory. If a second argument is present, it will be used as
# the base image name in pushed image tags.
#
function push_image_tags() {
  source_image=$1; shift || fatal "${FUNCNAME} needs source image."
  push_base=$1; shift || push_base=$source_image

  docker_login $push_base

  check_can_push || return 0

  push_images=$(get_image_tags $push_base)

  for push_image in $push_images; do
    docker tag "$source_image" "$push_image"
    docker push "$push_image"
  done
}

push_image_tags "$@"