#!/user/bin/env bash


# Request URL:https://quay.io/api/v1/repository/dunnevan/ibm-spectrum-scale-csi-driver/tag/v1.1.0
# Request Method:PUT
# {"expiration":1586277660}

function quay_add_tag() {

    curl \
        -X PUT \
        --silent \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer itvkQEgBMdaZmHVU2TMkW4wPDs3q99A0yKCRG5WP" \
        --data '{"image":"a978b108b8c8ef4823bffe923b8f6f5436c640f85d30acf76a9a3da508f1858a","manifest_digest":"sha256:74f274d7c5bb6bbbee48bb15e98bb016d5dcd27d82bf7babb6f07bc9e3aeb7e0"}'\
        "https://quay.io/api/v1/repository/dunnevan/ibm-spectrum-scale-csi-driver/tag/test"
}



QUAY_OAUTH_TOKEN="itvkQEgBMdaZmHVU2TMkW4wPDs3q99A0yKCRG5WP"

# Request URL:https://quay.io/api/v1/repository/dunnevan/ibm-spectrum-scale-csi-driver/tag/latest
# Request Method:PUT
# {"image":"a978b108b8c8ef4823bffe923b8f6f5436c640f85d30acf76a9a3da508f1858a","manifest_digest":"sha256:74f274d7c5bb6bbbee48bb15e98bb016d5dcd27d82bf7babb6f07bc9e3aeb7e0"}

function quay_add_expiration() {

    curl \
        --silent \
        -X PUT \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer ${QUAY_OAUTH_TOKEN}" \
        --data '{"expiration":$(epoch_add_exp_days 14)}'\
        "https://quay.io/api/v1/repository/dunnevan/ibm-spectrum-scale-csi-driver/tag/test"
}


function epoch_add_exp_days() {
    days=$1;

    if date --version >/dev/null 2>&1 ; then
        date -d "+${days} days" +%s
    else
        date -v "+{$days}d" +%s
    fi
}