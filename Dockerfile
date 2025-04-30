FROM hub.expvent.com.cn:1111/expvent/builder/golang:1.22 as builder
LABEL maintainer=expvent@expvent.com

WORKDIR /workspace
COPY . .

RUN --mount=type=cache,target=/go/pkg/mod,sharing=locked --mount=type=cache,target=/root/.cache,sharing=locked \
    make -e BUILD_DEST_DIR=build build

FROM hub.expvent.com.cn:1111/expvent/base/ubuntu:20.04
LABEL maintainer=expvent@expvent.com

WORKDIR /
USER root
COPY --from=builder /workspace/build/* ./

# 运行二进制文件
CMD ["./everai-test-go"]


