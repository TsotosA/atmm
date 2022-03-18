#!/bin/sh

# ex. ./versionBump.sh [major | minor | patch] [1.1.1]

var=$2
IFS=. read -r version minor patch <<EOF
$var
EOF
case "$1" in
	patch) tag="$version.$minor.$((patch+1))"; ;;
	major) tag="$((version+1)).0.0"; ;;
	minor)     tag="$version.$((minor+1)).0"; ;;
esac

echo $tag