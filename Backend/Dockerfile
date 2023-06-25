FROM ubuntu:latest
EXPOSE 3000
COPY genLink /opt/genLink
RUN chmod +x /opt/genLink
COPY entrypoint.sh /opt/entrypoint.sh
RUN chmod +x /opt/entrypoint.sh
ENTRYPOINT ["/opt/entrypoint.sh"]