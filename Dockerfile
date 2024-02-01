# 이미지 선택
FROM debian:buster-slim

# 패키지 설치
RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates curl netbase git && \
    rm -rf /var/lib/apt/lists/*

# Go 설치
ENV GO_VERSION go1.21.6
RUN curl -L https://golang.org/dl/${GO_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -xz

# 환경 변수 추가
ENV PATH="/usr/local/go/bin:${PATH}"

# 작업 디렉토리 정의
WORKDIR /app

# 기본 명령어 설정
CMD ["bash"]