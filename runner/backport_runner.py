import base64
import json
import os
from backport_runner_pipeline import BackportRunnerPipeline
from pubsub_reporter import PubSubReporter

reference = os.environ.get('REFERENCE')
base_branch = os.environ.get('BASE_BRANCH')
target_branch = os.environ.get('TARGET_BRANCH')
commits = os.environ.get('COMMITS')
reporter_config = os.environ.get('REPORTER_CONFIG')
source_path = os.environ.get('SOURCE_PATH')

if reference is None:
    raise ValueError('REFERENCE environment variable is not set')

if base_branch is None:
    raise ValueError('BASE_BRANCH environment variable is not set')

if target_branch is None:
    raise ValueError('TARGET_BRANCH environment variable is not set')

if commits is None:
    raise ValueError('COMMITS environment variable is not set')

if reporter_config is None:
    raise ValueError('REPORTER_CONFIG environment variable is not set')

if source_path is None:
    raise ValueError('SOURCE_PATH environment variable is not set')

commits = commits.split(',')

if len(commits) == 0:
    raise ValueError('No commits provided')

reporter_config = base64.b64decode(reporter_config).decode('utf-8')
reporter_config = json.loads(reporter_config)
reporter = PubSubReporter(reporter_config, reference=reference)

pipeline = BackportRunnerPipeline(source_path=source_path, destination=base_branch, branch_name=target_branch, commits=commits, reporter=reporter)

print('Starting runner with following parameters:')
print('Reference: ' + reference)
print('Base branch: ' + base_branch)
print('Target branch: ' + target_branch)
print('Commits: ' + str(commits))

pipeline.run()
print('Backporting finished successfully!')
