FROM ubuntu:18.04
COPY ./ /bin/
CMD ["/bin/test.exe"]
