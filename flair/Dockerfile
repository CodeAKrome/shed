# Use an official Python runtime as a parent image
FROM python:3.9-slim

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY flair_docker.py /app
COPY requirements.txt /app

# Install any needed packages specified in requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

# Make port 5000 available to the world outside this container
EXPOSE 1337

# Define environment variable
ENV PORT=1337

# Run flair_sentiment.py when the container launches
CMD ["python", "flair_docker.py"]
