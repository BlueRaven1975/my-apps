# Use the latest stable Python Alpine image from the Docker Hub
FROM python:3.13-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the requirements file
COPY requirements.txt .

# Install the requirements
RUN pip install --no-cache-dir -r requirements.txt

# Remove the requirements file
RUN rm -f requirements.txt

# Copy the application code
COPY main.py .

# Expose port 8081 to the outside world
EXPOSE 8081

# Run the application when the container launches
CMD ["python", "main.py"]