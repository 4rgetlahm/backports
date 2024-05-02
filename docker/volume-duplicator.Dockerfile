FROM alpine:3.19.1

CMD ["ash", "-c", "cd /source ; cp -av . /target"]