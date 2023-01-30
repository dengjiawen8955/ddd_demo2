FROM alpine  as runner
WORKDIR /app
# 创建目录
RUN mkdir -p /app
# 解决没有上海时区问题
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai
COPY ./bin/dc2 /app/
COPY ./config.yaml /app/
CMD ["./dc2", "--config", "./config.yaml"]
