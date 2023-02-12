#!/bin/bash

GITHUB_PERSONAL_ACCESS_TOKEN=''
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
TAG_TITLE=''
HAS_CHANGES="false"
VERSION_BUMP=

while getopts 't:b:' flag; do
  case "${flag}" in
    t) GITHUB_PERSONAL_ACCESS_TOKEN="${OPTARG}" ;;
    b) VERSION_BUMP="${OPTARG}" ;;
    *) echo test
       exit 1 ;;
  esac
done

if [[ -z $VERSION_BUMP || -z $GITHUB_PERSONAL_ACCESS_TOKEN ]]; then
    echo "please provide both version bump [-b] [major | minor | patch] and github personal access token [-t] flags"
    exit 1
fi

#uncomment to check for uncommitted changes
#if [[ -n $(git status -s) ]]; then
#    HAS_CHANGES="true"
#    echo "there are uncommitted changes"
#    exit 1
#fi

if [[ $CURRENT_BRANCH != "main" ]]; then
    echo "you are not in main branch"
    exit 1
fi


RELEASES=$(curl --silent \
  -u tsotosa:"${GITHUB_PERSONAL_ACCESS_TOKEN}" \
  -H "Accept: application/vnd.github.v3+json" \
  https://api.github.com/repos/tsotosa/atmm/releases \

  )
LATEST_TAG_VERSION="$(echo "$RELEASES" | LC_ALL=en_US.utf8 grep -P -o -e '[\d]{0,3}\.[\d]{0,3}\.[\d]{0,3}' | head -1)"

OVERRIDE_VERSION_BUMP=
if [[ -z $LATEST_TAG_VERSION ]]; then
    while true; do
        read -p "latest tag version is: [$LATEST_TAG_VERSION] which will produce unexpected version bump. proceed? [y,n,o] " yn
        case $yn in
            [Yy]* ) break;;
            [Nn]* ) exit;;
            [Oo]* ) OVERRIDE_VERSION_BUMP="y"; break;;
            * ) echo "Please answer yes or no.";;
        esac
    done
fi

NEW_VERSION=$( ./scripts/versionBump.sh $VERSION_BUMP $LATEST_TAG_VERSION )

if [[ $OVERRIDE_VERSION_BUMP == "y" ]]; then
    while true; do
        read -p "what is the new version? " yn
        case $yn in
            [Nn]* ) exit;;
            * ) NEW_VERSION="$yn"; break;;
        esac
    done
fi

TAG_TITLE="v${NEW_VERSION}"

./scripts/test.sh
if [ $? -eq 0 ]; then
    echo "tests ok"
else
    echo "failed to pass test suite"
    exit 1
fi

#uncomment to create tag locally
#git tag $TAG_TITLE
#if [ $? -eq 0 ]; then
#    echo "Created  release tag"
#else
#    echo "Failed to create tag"
#    exit 1
#fi

R=$(curl --silent \
-X POST \
-u tsotosa:"${GITHUB_PERSONAL_ACCESS_TOKEN}" \
-H "Accept: application/vnd.github.v3+json" \
https://api.github.com/repos/tsotosa/atmm/releases \
-d '{
    "tag_name":"'"$TAG_TITLE"'",
    "name":"'"$TAG_TITLE"'",
    "draft":true,
    "target_commitish": "main",
    "body":"",
    "generate_release_notes": true
    }'
  )

ASSETS="$(echo "$R" | LC_ALL=en_US.utf8 grep -P -o -e 'https:\/\/uploads\.github\.com\/.*assets' )"
#echo $ASSETS

rm ./bin/* 2> /dev/null
env GOOS=windows go build -ldflags="-s -w -X 'main.Version=${TAG_TITLE}'" -race -o  ./bin/atmm-windows-x64-${TAG_TITLE}.exe
env GOOS=linux go build -ldflags="-s -w -X 'main.Version=${TAG_TITLE}'" -o  ./bin/atmm-linux-x64-${TAG_TITLE}

FILES="bin/*"
## NOTE : Quote it else use array to avoid problems #
for f in $FILES
do
  echo "Processing $f file..."

  BASENAME="$(basename $f)"
#  echo $BASENAME

  RR=$(curl --silent \
  -X POST \
  -u tsotosa:"${GITHUB_PERSONAL_ACCESS_TOKEN}" \
  -H "Accept: application/vnd.github.v3+json" \
  -H "Content-Type: $(file -b --mime-type $f)" \
  -T "$f" \
  "${ASSETS}?name=$BASENAME" \
  | cat
    )

  echo "finished uploading $f file..."
done

echo "cleanup..deleting bin/*"
rm -r ./bin/*

