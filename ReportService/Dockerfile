FROM rust

WORKDIR /app

COPY ./ ./

RUN apt update
RUN apt install pkg-config

RUN rustup override set nightly
RUN cargo build --release

EXPOSE 8086

CMD ["./target/release/ReportService"]