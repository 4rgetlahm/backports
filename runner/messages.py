class RunnerStatusMessage():
    def __init__(self, stage, status, payload) -> None:
        self.stage = stage
        self.status = status
        self.payload = payload

    def get(self):
        return {
            'stage': self.stage,
            'status': self.status,
            'payload': self.payload
        }
