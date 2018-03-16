#! /bin/bash

docker login -u $DHUSERNAME -p $DHPASSWORD

dockerfiles=$(ls | grep "Dockerfile.*")
for dockerfile in ${dockerfiles}; do
  tag=$(echo "$dockerfile" | cut -d '.' -f2)
  docker push "${DHUSERNAME}/${tag}"
done
