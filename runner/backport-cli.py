import os
from backport_runner import BackportRunner
from runner.pubsub_reporter import PubSubReporter

reference = os.environ.get('REFERENCE')
base_branch = os.environ.get('BASE_BRANCH')
target_branch = os.environ.get('TARGET_BRANCH')
commits = os.environ.get('COMMITS')
pubsubCredentials = os.environ.get('PUBSUB_CREDENTIALS_BASE64')
sourcePath = os.environ.get('SOURCE_PATH')

if reference is None:
    raise ValueError('REFERENCE environment variable is not set')

if base_branch is None:
    raise ValueError('BASE_BRANCH environment variable is not set')

if target_branch is None:
    raise ValueError('TARGET_BRANCH environment variable is not set')

if commits is None:
    raise ValueError('COMMITS environment variable is not set')

if pubsubCredentials is None:
    raise ValueError('PUBSUB_CREDENTIALS_BASE64 environment variable is not set')

if sourcePath is None:
    raise ValueError('SOURCE_PATH environment variable is not set')

commits = commits.split(',')

if len(commits) == 0:
    raise ValueError('No commits provided')

reporter = PubSubReporter(pubsubCredentials, reference=reference)

backport_runner = BackportRunner(source_path=sourcePath, destination=base_branch, branch_name=target_branch, commits=commits, tracker=reporter)

print('Starting runner with following parameters:')
print('Reference: ' + reference)
print('Base branch: ' + base_branch)
print('Target branch: ' + target_branch)
print('Commits: ' + str(commits))

backport_runner.run()
print('Backporting finished successfully!')
