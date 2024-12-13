### README ###
### Run: docker build -t accel_db_img .
###
# Use the official PostgreSQL image as the base
FROM postgres:15-alpine

# Set the environment variables for the PostgreSQL database
ENV POSTGRES_USER accelUser
ENV POSTGRES_PASSWORD accelPass
ENV POSTGRES_DB accel_db

# Expose the PostgreSQL port
EXPOSE 5432