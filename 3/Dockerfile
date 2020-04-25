FROM alpine
LABEL maintainer="sharath.aradhyamath <sharatham94@gmai.com>"
# RUN groupadd -r -g 2200 gousr && \
#     useradd -rM -g gousr -u 2200 gousr
WORKDIR /app
#copying executable linux binary [binary was build in local system]
ADD goredis_linux_amd64_link .
# USER gousr
EXPOSE 8080
ENTRYPOINT [ "./goredis_linux_amd64_link" ]
