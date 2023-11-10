set -e
# only exit with zero if all commands of the pipeline exit successfully
set -o pipefail

export REGISTRIES="${REGISTRIES:-"458176070654.dkr.ecr.us-east-1.amazonaws.com"}"

export IMAGE_OPERATOR="${IMAGE_OPERATOR:-"coveord/prometheus-operator"}"
export IMAGE_RELOADER="${IMAGE_RELOADER:-"coveord/prometheus-config-reloader"}"
export IMAGE_WEBHOOK="${IMAGE_WEBHOOK:="coveord/admission-webhook"}"

# Authenticate to dev ECR
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 458176070654.dkr.ecr.us-east-1.amazonaws.com

./scripts/push-docker-image.sh
