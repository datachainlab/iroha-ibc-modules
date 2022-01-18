FROM my-iroha-burrow:1.3.0

COPY entrypoint.sh /
COPY wait-for-it.sh /

ENTRYPOINT ["/entrypoint.sh"]
CMD ["irohad"]
