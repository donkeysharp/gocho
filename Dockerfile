FROM debian:latest

COPY ./dist/gocho /usr/local/bin/gocho

RUN chmod +x /usr/local/bin/gocho \
    && mkdir -p /root/public \
    && echo 'file1' > /root/public/file1 \
    && echo 'file2' > /root/public/file2 \
    && echo 'file3' > /root/public/file3 \
    && echo 'file4' > /root/public/file4 \
    && echo 'file5' > /root/public/file5 \
    && echo 'NodeId: root' > /root/.gocho.conf \
    && echo 'WebPort: "5555"' >> /root/.gocho.conf \
    && echo 'LocalPort: "1337"' >> /root/.gocho.conf \
    && echo 'ShareDirectory: "/root/public"' >> /root/.gocho.conf

CMD ["/usr/local/bin/gocho", "start"]
