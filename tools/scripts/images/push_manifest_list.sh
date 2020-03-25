#!/usr/bin/env bash

source tools/scripts/images/common.sh

#
# push_manifest_list base_image [platforms...]
#
# note: [platforms...] like "linux/amd64 linux/s390x linux/ppc64le" or "amd64 s390x ppc64le"
function push_manifest_list() {
  manifest=$1; shift || fatal "${FUNCNAME} needs base image (manifest)."
  platforms=${@:-"linux/amd64"}

  docker_login $manifest

  check_can_push || return 0

  tags=$(get_image_tags)
  for tag in $tags; do
    images=$(get_arch_images $manifest $tag $platforms)
    DOCKER_CLI_EXPERIMENTAL="enabled" docker manifest create $manifest:$tag $images
    DOCKER_CLI_EXPERIMENTAL="enabled" docker manifest push --purge $manifest:$tag
  done
}

#
# get_arch_images base_image tag [platforms...]
#
function get_arch_images() {
    base=$1; shift || fatal "${FUNCNAME} needs base image."
    tag=$1; shift || fatal "${FUNCNAME} needs tag."
    platforms="$@"
    for plat in $platforms; do
        # grab last split on '/'
        arch=${plat##*\/}
        echo "$base-$arch :$tag"
    done
}

push_manifest_list "$@"