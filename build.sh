#! /bin/bash

dockerfiles=$(ls | grep "Dockerfile.*")
	for dockerfile in ${dockerfiles}; do
		tag=$(echo "$dockerfile" | cut -d '.' -f2)
		docker build -f $dockerfile . -t "${DHUSERNAME}/${tag}"
done
