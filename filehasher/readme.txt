other user:
    docker pull grahamwichhh/filehasher:latest
    docker run --rm -v $(pwd):/data grahamwichhh/filehasher:latest --file=input.txt --algo=md5 --output=/data/output.txt



build/push later versions:
    docker build -t grahamwichhh/filehasher:v2 .
    docker push grahamwichhh/filehasher:v2


build with both intel and arm (mac):
    docker buildx create --use
    docker buildx build --platform linux/amd64,linux/arm64 -t grahamwichhh/filehasher:latest --push .
