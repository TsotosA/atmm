#!/bin/bash
#-v version
#-t github personal access token

VERSION=''
GITHUB_PERSONAL_ACCESS_TOKEN=''
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
TAG_TITLE=''
HAS_CHANGES="false"

while getopts 't:v:' flag; do
  case "${flag}" in
    v) VERSION="${OPTARG}" ;;
    t) GITHUB_PERSONAL_ACCESS_TOKEN="${OPTARG}" ;;
    *) echo test
       exit 1 ;;
  esac
done

if [[ -z $VERSION || -z $GITHUB_PERSONAL_ACCESS_TOKEN ]]; then
    echo "please provide both version [-v] and github personal access token [-t] flags"
    exit 1
fi

#uncomment to check for uncommitted changes
#if [[ -n $(git status -s) ]]; then
#    HAS_CHANGES="true"
#    echo "there are uncommitted changes"
#    exit 1
#fi

if [[ $CURRENT_BRANCH != "master" ]]; then
    echo "you are not in master branch"
    exit 1
fi

TAG_TITLE="v${VERSION}"

./test.sh
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
    "target_commitish": "master",
    "body":"describe the release"
    }'
  )

ASSETS="$(echo "$R" | grep -P -o -e 'https:\/\/uploads\.github\.com\/.*assets' )"
#echo $ASSETS

rm ./bin/* 2> /dev/null
env GOOS=windows go build -ldflags="-X 'main.Version=${TAG_TITLE}'" -race -o  ./bin/atmm-windows-x64-${TAG_TITLE}.exe
env GOOS=linux go build -ldflags="-X 'main.Version=${TAG_TITLE}'" -o  ./bin/atmm-linux-x64-${TAG_TITLE}

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
