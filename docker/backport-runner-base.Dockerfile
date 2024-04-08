FROM python:3.12.2-slim

# Clone the backports repository
RUN apt-get update && apt-get install -y git
RUN git clone https://github.com/4rgetlahm/backports.git

COPY . .

# Install requirements
WORKDIR /backports/runner
RUN pip install -r requirements.txt

# Set the entrypoint
ENTRYPOINT ["python", "backport_runner.py"]