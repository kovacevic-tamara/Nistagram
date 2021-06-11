#!/bin/bash
cd .. && cd gateway && docker build -t gateway .
cd .. && cd backend
docker build -t user_service -f UserServiceDockerfile .
docker build -t content_service -f ContentServiceDockerfile .
docker build -t recommendation_service -f RecommendationServiceDockerfile .
docker-compose up