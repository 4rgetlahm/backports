import base64
import json
import os
from backport_runner_pipeline_git import GitBackportRunnerPipeline
from pubsub_reporter import PubSubReporter
from backport_runner_pipeline_mercurial import MercurialBackportRunnerPipeline

reference = os.environ.get('REFERENCE')
new_branch_name = os.environ.get('NEW_BRANCH_NAME')
target_branch_name = os.environ.get('TARGET_BRANCH_NAME')
commits = os.environ.get('COMMITS')
reporter_config = os.environ.get('REPORTER_CONFIG')
source_path = os.environ.get('SOURCE_PATH')
vcs = os.environ.get('VCS')

if reference is None:
    raise ValueError('REFERENCE environment variable is not set')

if new_branch_name is None:
    raise ValueError('NEW_BRANCH_NAME environment variable is not set')

if target_branch_name is None:
    raise ValueError('TARGET_BRANCH_NAME environment variable is not set')

if commits is None:
    raise ValueError('COMMITS environment variable is not set')

if reporter_config is None:
    raise ValueError('REPORTER_CONFIG environment variable is not set')

if source_path is None:
    raise ValueError('SOURCE_PATH environment variable is not set')

if vcs is None:
    raise ValueError('VCS environment variable is not set')

commits = commits.split(',')

if len(commits) == 0:
    raise ValueError('No commits provided')

reporter_config = base64.b64decode(reporter_config).decode('utf-8')
reporter_config = json.loads(reporter_config)
reporter = PubSubReporter(reporter_config, reference=reference)

pipeline = None
if vcs == 'git':
    pipeline = GitBackportRunnerPipeline(source_path=source_path, target_branch_name=target_branch_name, new_branch_name=new_branch_name, commits=commits, reporter=reporter)
elif vcs == 'hg':
    pipeline = MercurialBackportRunnerPipeline(source_path=source_path, target_branch_name=target_branch_name, new_branch_name=new_branch_name, commits=commits, reporter=reporter)

print('Starting runner with following parameters:')
print('Source path: ' + source_path)
print('Reference: ' + reference)
print('New branch name: ' + new_branch_name)
print('Target branch name: ' + target_branch_name)
print('Commits: ' + str(commits))

pipeline.run()
print('Backporting finished successfully!')
