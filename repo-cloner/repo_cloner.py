import os
from git import Repo

credentials = os.environ.get('CREDENTIALS')
clone_url = os.environ.get('CLONE_URL')

if credentials is None:
    raise ValueError('CREDENTIALS environment variable is not set')

if clone_url is None:
    raise ValueError('CLONE URL environment variable is not set')

clone_url = clone_url.replace('https://', f'https://{credentials}@')
Repo.clone_from(clone_url, '/repo')