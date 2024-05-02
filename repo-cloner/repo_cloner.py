import os
import subprocess

vcs = os.environ.get('VCS')
clone_url = os.environ.get('CLONE_URL')

if clone_url is None:
    raise ValueError('CLONE URL environment variable is not set')

if vcs == 'git':
    subprocess.run(['git', 'clone', clone_url, '/repo'])

if vcs == 'hg':
    subprocess.run(['hg', 'clone', clone_url, '/repo'])

print('Repo cloned successfully')