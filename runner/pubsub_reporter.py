import atexit
import json
import base64
from messages import RunnerStatusMessage
from google.cloud import pubsub_v1
from google.auth import jwt
from reporter import EmptyReporter

# Reporter config is base64 encoded JSON with following structure:
# {
#     "project": "project-id",
#     "topic": "topic-name",
#     "credentials": base64 encoded service account JSON
# }

class PubSubReporter(EmptyReporter):
    def __init__(self, reporter_config, reference=None) -> None:
        self.credentials = base64.b64decode(reporter_config["credentials"]).decode('utf-8')
        self.credentials = json.loads(self.credentials)

        self.audience = "https://pubsub.googleapis.com/google.pubsub.v1.Publisher"

        credentials = jwt.Credentials.from_service_account_info(
            self.credentials, audience=self.audience)

        self.publisher = pubsub_v1.PublisherClient(credentials=credentials)
        self.topic = 'projects/{project_id}/topics/{topic}'.format(
            project_id=reporter_config["project"],
            topic=reporter_config["topic"],
        )
        self.reference = reference

        print(
            f"Reporter initialized for reference: {self.reference}, topic: {self.topic}")
        atexit.register(self.exit_handler)
        self.send(RunnerStatusMessage('start', 'success', {}).get())

    def send(self, json_message):
        json_message["reference"] = str(self.reference)
        try:
            future = self.publisher.publish(
                self.topic, json.dumps(json_message).encode('utf-8'))
            future.result()
        except Exception as e:
            print(e)

    def exit_handler(self):
        self.send(RunnerStatusMessage('exit', 'success', {}).get())
