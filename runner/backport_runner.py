import git
from messages import RunnerStatusMessage
from tracker import EmptyTracker
import time


class BackportRunner:
    def __init__(self, source_path, destination, branch_name, commits, tracker=EmptyTracker()) -> None:
        self.destination = destination
        self.branch_name = branch_name
        self.commits = commits

        self.git = git.Git(source_path)
        self.tracker = tracker
        self.push_retry_count = 0
        self.push_retry_count_limit = 5

    def run(self):
        self.fetch()
        self.checkout(self.destination, clean=True)
        self.pull()
        self.checkout_new_branch(self.branch_name)
        for commit in self.commits:
            self.cherry_pick(commit, ['-m 1'])
        self.push()

    def checkout(self, branch, clean=False):
        print('Checking out branch: ' + branch)
        self.tracker.send(RunnerStatusMessage(
            'checkout', 'start', {"branch": branch}).get())
        try:
            if clean:
                self.git.reset('--hard')

            self.git.checkout(branch)
            
            if clean:
                self.git.reset('--hard')
                self.git.clean('-dfx')
            self.tracker.send(RunnerStatusMessage(
                'checkout', 'success', {"branch": branch}).get())
        except git.GitCommandError as e:
            print(
                f'Failed to check out to branch {branch}. Please make sure that branch exists.')
            print(e)
            self.tracker.send(RunnerStatusMessage(
                'checkout', 'failure', {"branch": branch, "error": str(e)}).get())
            exit(1)

    def checkout_new_branch(self, branch):
        print('Checking out new branch: ' + branch)
        self.tracker.send(RunnerStatusMessage(
            'checkout_new_branch', 'start', {"branch": branch}).get())
        try:
            self.git.checkout('-b', branch)
            self.tracker.send(RunnerStatusMessage(
                'checkout_new_branch', 'success', {"branch": branch}).get())
        except git.GitCommandError as e:
            print(
                f'Failed to check out to new branch {branch}. Please make sure that branch does not exist.')
            print
            self.tracker.send(RunnerStatusMessage(
                'checkout_new_branch', 'failure', {"branch": branch, "error": str(e)}).get())
            exit(1)

    def cherry_pick(self, commit, options):
        print('Cherry picking commit: ' + commit)
        self.tracker.send(RunnerStatusMessage(
            'cherry_pick', 'start', {"commit": commit}).get())
        try:
            self.git.cherry_pick(options, commit)
            self.tracker.send(RunnerStatusMessage(
                'cherry_pick', 'success', {"commit": commit}).get())
        except git.GitCommandError as e:
            if "CONFLICT" in e.stdout:
                print('Merge conflict was detected')
                print(e)
                self.tracker.send(RunnerStatusMessage('cherry_pick', 'failure', {
                    "commit": commit, "error": str(e)}).get())
                raise MergeConflictException
            else:
                print(f'Failed to cherry pick commit {commit}')
                print(e)
                self.tracker.send(RunnerStatusMessage('cherry_pick', 'failure', {
                    "commit": commit, "error": str(e)}).get())
                exit(1)

    def pull(self):
        print('Pulling from origin')
        self.tracker.send(RunnerStatusMessage(
            'pull', 'start', {}).get()
        )
        try:
            self.git.pull()
            self.tracker.send(RunnerStatusMessage(
                'pull', 'success', {}).get()
            )
        except git.GitCommandError as e:
            print('Failed to pull from origin')
            print(e)
            self.tracker.send(RunnerStatusMessage(
                'pull', 'failure', {"error": str(e)}).get()
            )
            exit(1)

    def fetch(self):
        print('Fetching from origin')
        self.tracker.send(RunnerStatusMessage(
            'fetch', 'start', {}).get()
        )
        try:
            self.git.fetch('--prune', '--all')
            self.tracker.send(RunnerStatusMessage(
                'fetch', 'success', {}).get()
            )
        except git.GitCommandError as e:
            print('Failed to fetch from origin')
            print(e)
            self.tracker.send(RunnerStatusMessage(
                'fetch', 'failure', {"error": str(e)}).get()
            )
            exit(1)

    def push(self):
        print('Pushing to origin')
        self.tracker.send(RunnerStatusMessage(
            'push', 'start', {"branch": self.branch_name}).get()
        )
        try:
            self.git.push('--set-upstream', 'origin', self.branch_name)
            self.tracker.send(RunnerStatusMessage(
                'push', 'success', {"branch": self.branch_name}).get()
            )
        except git.GitCommandError as e:
            print('Failed to push to origin')
            print(e)
            if self.push_retry_count < self.push_retry_count_limit:
                self.push_retry_count += 1
                self.push()
                return
            self.tracker.send(RunnerStatusMessage(
                'push', 'failure', {"branch": self.branch_name, "error": str(e)}).get()
            )
            exit(1)


class MergeConflictException(Exception):
    pass
