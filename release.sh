#!/usr/bin/env bash

# reference: https://sourcegraph.com/github.com/sourcegraph/src-cli@main/-/blob/release.sh

set -euf -o pipefail

read -p "What version do you want to release? " VERSION

if ! echo "$VERSION" | grep -Eq '^[0-9]+\.[0-9]+\.[0-9]+(-rc\.[0-9]+)?$'; then
    echo -e "\033[31m\$VERSION is not in MAJOR.MINOR.PATCH format\033[0m"
  exit 1
fi

# Create a new tag and push it, this will trigger the goreleaser workflow in .github/workflows/goreleaser.yml
git tag "${VERSION}" -a -m "release v${VERSION}"
# We use `--atomic` so that we push the tag and the commit if the commit was or wasn't pushed before
git push --atomic origin main "${VERSION}"

sed -i "4s/VERSION=\"[0-9]*\.[0-9]*\.[0-9]*\"/VERSION=\"${VERSION}\"/" install.sh

git add install.sh
git commit -m "update: ${VERSION}"
git push origin main
