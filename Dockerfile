FROM scratch
COPY logserver /
ENTRYPOINT ["/logserver"]