FROM scylladb/scylla:latest

# Copy the CQL file into the container
COPY init.cql /docker-entrypoint-initdb.d/init.cql

# Run the CQL file on startup
ENTRYPOINT ["/docker-entrypoint.py"]