#!/bin/bash
versionprefix="v0.1.0"
revisioncount=`git log --oneline | wc -l | tr -d ' '`
projectversion=`git describe --tags --long`
cleanversion=${projectversion%%-*}

# echo "$projectversion-$revisioncount"
# echo "$cleanversion.$revisioncount"
version=$versionprefix.$revisioncount
gitmessage="Auto deployment $version"

# update version number into go func
sed "s/{{.VERSION}}/$version/g" util/version > util/version.go

git add .
git commit -m "$gitmessage"
git push origin master
git tag -a $version -m "$gitmessage"
git push origin $version